package blockchain

import (
	"flag"
	"fmt"
	"strings"

	"sync"

	"github.com/crazybits/x/common"
	"github.com/crazybits/x/db"
	logging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var logger = logging.MustGetLogger("blockchain")

const blockchainInfoKey = "BLOCKCHAIN_INFO_KEY"

var bm *BlockManager

var once sync.Once

func init() {

	flag.Parse()
	// Now set the configuration file
	viper.SetEnvPrefix("blockchain")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName("config")    // name of config file (without extension)
	viper.AddConfigPath("../config") // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

type BlockManager struct {
	db *db.LevelDB
}

func NewBlockManager() *BlockManager {

	once.Do(func() {
		bm = new(BlockManager)
		db, err := db.NewLevelDB(viper.GetString("blockchain.dbPath"), viper.GetInt("blockchain.cache"), viper.GetInt("blockchain.handle"))
		if err != nil {
			logger.Errorf("failed to create db", err)
		}
		bm.db = db
	})
	return bm
}

func (bm *BlockManager) PushBlock(block *Block) {

	blockID, err := block.Digest()
	if err != nil {
		logger.Errorf("failed to get the block id:%s", err)
	}
	blockBytes, err := block.Encode()

	if err != nil {
		logger.Errorf("failed to get the block bytes:%s", err)
	}
	err = bm.db.Put(blockID, blockBytes)
	if err != nil {
		logger.Errorf("failed to push block:%s", err)
	}

}

func (bm *BlockManager) GetBlockByID(id []byte) *Block {

	data, err := bm.db.Get(id)
	if err != nil {
		logger.Errorf("failed to get block by id=%s:%s", common.BytesToHexSting(id), err)
	}

	block := NewBlock()

	err = block.Decode(data)
	if err != nil {
		logger.Errorf("failed to decode the block:%s", err)
	}
	return block
}
func (bm *BlockManager) UpdateBlockchainInfo(blochainInfo *BlockchainInfo) {

	value, err := blochainInfo.Encode()
	if err != nil {
		logger.Errorf("failed to encode of blochainInfo: %s", err)
	}
	err = bm.db.Put(common.StrToSha256(blockchainInfoKey), value)
	if err != nil {
		logger.Errorf("failed to upate blockhcainInfo to db:%s", err)
	}
}

func (bm *BlockManager) GetBlockchainInfo() *BlockchainInfo {

	blockchainInfor := NewBlockchainInfo()
	data, err := bm.db.Get(common.StrToSha256(blockchainInfoKey))
	if err != nil {
		logger.Errorf("failed to get the blockchainInfor:%s", err)
		return nil
	}

	err = blockchainInfor.Decode(data)
	if err != nil {
		logger.Errorf("failed to decode the blockchainInfor:%s", err)
		return nil
	}
	return blockchainInfor
}

func (bm *BlockManager) ProcessBlock(block *Block) error {

	txCount := len(block.Transactions)
	lockChan := make(chan bool, txCount)

	for _, tx := range block.Transactions {
		go bm.ProccessTransaction(tx, lockChan)
	}

	for i := 0; i < txCount; i++ {
		<-lockChan
	}
	return nil

}

func (bm *BlockManager) ProccessTransaction(tx *Transaction, lockChan chan bool) {

	tx.Evaluate()

	lockChan <- true

	return

}

func (bm *BlockManager) ShutDown() {
	bm.db.Close()

}
