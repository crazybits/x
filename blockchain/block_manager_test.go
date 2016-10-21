package blockchain

import "testing"

func TestBlock(t *testing.T) {

	receiver := StringToAddress("crazybit")

	withdrawSymbol := "Symbol"

	withdrawAmount := int64(1024)

	depositOp := NewDepositOperation(receiver, withdrawAmount, withdrawSymbol)

	data, _ := depositOp.Encode()

	operation := &Operation{OpType: OperationType_Deposit, Payload: data}

	transaction := &Transaction{}
	transaction.AddOperation(operation)

	block := &Block{}

	block.AddTransaction(transaction)

	bc := NewBlockManager()

	bc.ProcessBlock(block)

}
