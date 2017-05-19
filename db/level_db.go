package db

import (
	"path/filepath"
	"sync"
	"time"

	logging "github.com/op/go-logging"
	gometrics "github.com/rcrowley/go-metrics"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var logger = logging.MustGetLogger("db")

var OpenFileLimit = 64

// cacheRatio specifies how the total alloted cache is distributed between the
// various system databases.
var cacheRatio = map[string]float64{
	"dapp":      0.0,
	"chaindata": 1.0,
}

// handleRatio specifies how the total alloted file descriptors is distributed
// between the various system databases.
var handleRatio = map[string]float64{
	"dapp":      0.0,
	"chaindata": 1.0,
}

type LevelDB struct {
	fn string      // filename for reporting
	db *leveldb.DB // LevelDB instance

	getTimer       gometrics.Timer // Timer for measuring the database get request counts and latencies
	putTimer       gometrics.Timer // Timer for measuring the database put request counts and latencies
	delTimer       gometrics.Timer // Timer for measuring the database delete request counts and latencies
	missMeter      gometrics.Meter // Meter for measuring the missed database get requests
	readMeter      gometrics.Meter // Meter for measuring the database get request data usage
	writeMeter     gometrics.Meter // Meter for measuring the database put request data usage
	compTimeMeter  gometrics.Meter // Meter for measuring the total time spent in database compaction
	compReadMeter  gometrics.Meter // Meter for measuring the data read during compaction
	compWriteMeter gometrics.Meter // Meter for measuring the data written during compaction

	quitLock sync.Mutex      // Mutex protecting the quit channel access
	quitChan chan chan error // Quit channel to stop the metrics collection before closing the database
}

// NewLevelDB returns a LevelDB wrapped object.
func NewLevelDB(file string, cache int, handles int) (*LevelDB, error) {
	// Calculate the cache and file descriptor allowance for this particular database
	cache = int(float64(cache) * cacheRatio[filepath.Base(file)])
	if cache < 16 {
		cache = 16
	}
	handles = int(float64(handles) * handleRatio[filepath.Base(file)])
	if handles < 16 {
		handles = 16
	}
	logger.Debugf("Alloted %dMB cache and %d file handles to %s", cache, handles, file)

	// Open the db and recover any potential corruptions
	db, err := leveldb.OpenFile(file, &opt.Options{
		OpenFilesCacheCapacity: handles,
		BlockCacheCapacity:     cache / 2 * opt.MiB,
		WriteBuffer:            cache / 4 * opt.MiB, // Two of these are used internally
	})
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(file, nil)
	}
	// (Re)check for errors and abort if opening of the db failed
	if err != nil {
		return nil, err
	}
	return &LevelDB{
		fn: file,
		db: db,
	}, nil
}

// Put puts the given key / value to the queue
func (self *LevelDB) Put(key []byte, value []byte) error {
	// Measure the database put latency, if requested
	if self.putTimer != nil {
		defer self.putTimer.UpdateSince(time.Now())
	}
	// Generate the data to write to disk, update the meter and write
	//value = rle.Compress(value)

	if self.writeMeter != nil {
		self.writeMeter.Mark(int64(len(value)))
	}
	return self.db.Put(key, value, nil)
}

// Get returns the given key if it's present.
func (self *LevelDB) Get(key []byte) ([]byte, error) {
	// Measure the database get latency, if requested
	if self.getTimer != nil {
		defer self.getTimer.UpdateSince(time.Now())
	}
	// Retrieve the key and increment the miss counter if not found
	dat, err := self.db.Get(key, nil)
	if err != nil {
		if self.missMeter != nil {
			self.missMeter.Mark(1)
		}
		return nil, err
	}
	// Otherwise update the actually retrieved amount of data
	if self.readMeter != nil {
		self.readMeter.Mark(int64(len(dat)))
	}
	return dat, nil
	//return rle.Decompress(dat)
}

// Delete deletes the key from the queue and database
func (self *LevelDB) Delete(key []byte) error {
	// Measure the database delete latency, if requested
	if self.delTimer != nil {
		defer self.delTimer.UpdateSince(time.Now())
	}
	// Execute the actual operation
	return self.db.Delete(key, nil)
}

func (self *LevelDB) NewIterator() iterator.Iterator {
	return self.db.NewIterator(nil, nil)
}

func (self *LevelDB) Close() {
	// Stop the metrics collection to avoid internal database races
	self.quitLock.Lock()
	defer self.quitLock.Unlock()

	if self.quitChan != nil {
		errc := make(chan error)
		self.quitChan <- errc
		if err := <-errc; err != nil {
			logger.Infof("metrics failure in '%s': %v\n", self.fn, err)
		}
	}
	err := self.db.Close()
	if err != nil {
		logger.Infof("error closing db %s: %s", self.fn, err)
	} else {
		logger.Infof("close db: %s", self.fn)
	}
}

func (self *LevelDB) LDB() *leveldb.DB {
	return self.db
}

// TODO: remove this stuff and expose leveldb directly

func (db *LevelDB) NewBatch() Batch {
	return &ldbBatch{db: db.db, b: new(leveldb.Batch)}
}

type ldbBatch struct {
	db *leveldb.DB
	b  *leveldb.Batch
}

func (b *ldbBatch) Put(key, value []byte) error {
	b.b.Put(key, value)
	return nil
}

func (b *ldbBatch) Write() error {
	return b.db.Write(b.b, nil)
}
