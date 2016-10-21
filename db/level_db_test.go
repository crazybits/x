package db

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/crazybits/go_sample/common"
)

func TestLevelDB(t *testing.T) {

	ldb, err := NewLevelDB("crazybit", 16, 16)
	defer ldb.Close()

	if err != nil {

		t.Fail()
	}

}

func TestLevelDBPutData(t *testing.T) {

	key := common.StrToSha512("crazybit")

	value := common.StrToSha512("crazybit at github")

	ldb, _ := NewLevelDB("crazybit", 18, 18)

	defer ldb.Close()

	ldb.Put(key, value)

}

func TestGetPersistedData(t *testing.T) {

	key := common.StrToSha512("crazybit")

	value := "69425c5d8efeb6e5b990c0c9d3281b2b7870f4bd383e513dfb86983dedfb23b4874e03461fac4e3bad251fa4a36ba32133e7381c258e892b08dca5420fce032d"

	ldb, _ := NewLevelDB("crazybit", 18, 18)

	defer ldb.Close()

	v, err := ldb.Get(key)

	if err != nil {

		t.Error(err)
	}

	if !strings.EqualFold(value, hex.EncodeToString(v)) {
		t.Fail()
	}
}

func TestInputBigData(t *testing.T) {

	ldb, _ := NewLevelDB("crazybit", 16, 16)

	defer ldb.Close()

	start := time.Now()

	for i := 0; i < 1000000; i++ {

		key := common.GenerateRandomBytes(32)

		value := common.GenerateRandomBytes(32)

		ldb.Put(key, value)
	}

	fmt.Println("insert 1000000 times data cost time:", time.Since(start))

}

func TestInputBigDataWithBatch(t *testing.T) {

	ldb, _ := NewLevelDB("crazybit", 16, 16)

	defer ldb.Close()

	start := time.Now()

	for i := 0; i < 1000000; i++ {

		key := common.GenerateRandomBytes(32)

		value := common.GenerateRandomBytes(32)

		batch := ldb.NewBatch()

		batch.Put(key, value)

		if i%100000 == 0 {
			batch.Write()
		}

	}

	fmt.Println("insert 1000000 times data cost time:", time.Since(start))

}
