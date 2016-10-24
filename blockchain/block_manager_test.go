package blockchain

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBlock(t *testing.T) {

	sneder := StringToAddress("crazybit")

	withdrawSymbol := "Symbol"

	withdrawAmount := int64(1024)

	withdrawOp := NewWithdrawOperation()
	withdrawOp.Sender = sneder
	withdrawOp.Symbol = withdrawSymbol
	withdrawOp.Amount = withdrawAmount

	data, _ := withdrawOp.Encode()

	op := NewOperation()
	op.Type = OperationType_Withdraw
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

	//t.Skip("skip this case")
	bc := NewBlockManager()

	/*bci := NewBlockchainInfo()
	bci.CurrentBlockHash = common.StrToSha256("cazybit")
	bci.PreviousBlockHash = common.StrToSha256("crazybits")
	bci.Height = 1024

	bc.UpdateBlockchainInfo(bci)*/
	blockchainInfo := bc.GetBlockchainInfo()
	fmt.Println(blockchainInfo.Height)

	bm.ShutDown()

}
