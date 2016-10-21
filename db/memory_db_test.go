package db

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/crazybits/x/common"
)

func TestDBInit(t *testing.T) {

	key := common.StrToSha512("crazybit")

	value := common.StrToSha512("github.com/crazybit")

	value2 := common.StrToSha512("github.com/crazybit2")

	var v []byte
	var err error

	var memdb *MemDatabase

	memdb, err = NewMemDatabase()

	if err != nil {

		t.Error(err)
	}

	memdb.Put(key, value)

	v, err = memdb.Get(key)

	if err != nil {

		t.Error(err)
	}

	if !bytes.Equal(v, value) {
		t.Error("failed")
	}

	memdb.Set(key, value2)

	v, err = memdb.Get(key)

	if err != nil {

		t.Error(err)
	}

	if !bytes.Equal(v, value) && bytes.Equal(v, value2) {

		fmt.Println("passed")

	}

}
