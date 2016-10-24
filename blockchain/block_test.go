package blockchain

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBlockEncode(t *testing.T) {

	sender := StringToAddress("crazybit")

	withdrawSymbol := "Symbol"

	withdrawAmount := int64(1024)

	withdrawOp := NewWithdrawOperation()
	withdrawOp.Sender = sender
	withdrawOp.Symbol = withdrawSymbol
	withdrawOp.Amount = withdrawAmount

	data, _ := withdrawOp.Encode()

	op := NewOperation()
	op.Type = OperationType_Withdraw
	op.Payload = data

	tx := NewTransaction()
	tx.AddOperation(op)
	tx.AddOperation(op)

	block := NewBlock()

	block.AddTransaction(tx)

	blockByte, err := block.Encode()

	blockNew := NewBlock()
	err = blockNew.Decode(blockByte)
	if err != nil {
		t.Fail()
	}

	fmt.Println(block)
	fmt.Println(blockNew)

	if !reflect.DeepEqual(block, blockNew) {

		t.Fail()
	}

}
