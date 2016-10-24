package blockchain

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBlock(t *testing.T) {

	receiver := StringToAddress("crazybit")

	withdrawSymbol := "Symbol"

	withdrawAmount := int64(1024)

	depositOp := NewDepositOperation()
	depositOp.Receiver = receiver
	depositOp.Symbol = withdrawSymbol
	depositOp.Amount = withdrawAmount

	data, _ := depositOp.Encode()

	op := NewOperation()
	op.Type = OperationType_Deposit
	op.Payload = data

	tx := NewTransaction()
	tx.AddOperation(op)

	block := NewBlock()

	block.AddTransaction(tx)

	bm := NewBlockManager()

	bm.ProcessBlock(block)

	id, err := block.Digest()
	if err != nil {
		t.Fail()
	}
	bm.PushBlock(block)

	newBlock := bm.GetBlockByID(id)

	if !reflect.DeepEqual(block, newBlock) {
		t.Fail()
	}

}
func TestBlockchainInfo(t *testing.T) {

	bc := NewBlockManager()
	blockchainInfo := bc.GetBlockchainInfo()
	fmt.Println(blockchainInfo)

}
