package blockchain

import (
	"reflect"
	"testing"
)

func TestTransctionMarshalAndUnmarsh(t *testing.T) {

	depositOp := &DepositOperation{}

	depositOp.Receiver = StringToAddress("crazybit")
	depositOp.Amount = int64(100000)
	depositOp.Symbol = "x"

	depositPayload, err := depositOp.Encode()

	if err != nil {
		t.Fail()
	}

	withdrawOp := &WithdrawOperation{}

	withdrawOp.Sender = StringToAddress("crazybit")
	withdrawOp.Amount = int64(100000)
	withdrawOp.Symbol = "x"

	withdrawPayload, err := withdrawOp.Encode()

	if err != nil {
		t.Fail()
	}

	op1 := NewOperation()
	op1.Type = OperationType_Deposit
	op1.Payload = depositPayload

	op2 := NewOperation()
	op2.Type = OperationType_Withdraw
	op2.Payload = withdrawPayload

	tx := &Transaction{}

	tx.AddOperation(op1)
	tx.AddOperation(op2)

	data, _ := tx.Encode()

	tx2 := NewTransaction()

	tx2.Decode(data)

	for _, v := range tx2.Operations {

		switch v.Type {

		case OperationType_Deposit:

			depositOp2 := &DepositOperation{}
			depositOp2.Decode(v.Payload)

			if !reflect.DeepEqual(depositOp, depositOp2) {
				t.Fail()
			}
		case OperationType_Withdraw:

			withdrawOp2 := &WithdrawOperation{}
			withdrawOp2.Decode(v.Payload)

			if !reflect.DeepEqual(withdrawOp, withdrawOp2) {
				t.Fail()
			}

		}
	}

}

func TestTransactionEvaluate(t *testing.T) {

	depositOp := NewDepositOperation()

	depositOp.Receiver = StringToAddress("crazybit1")
	depositOp.Amount = int64(100000)
	depositOp.Symbol = "x"

	depositPayload, err := depositOp.Encode()
	if err != nil {
		t.Fail()
	}

	withdrawOp := NewWithdrawOperation()

	withdrawOp.Sender = StringToAddress("crazybit2")
	withdrawOp.Amount = int64(100000)
	withdrawOp.Symbol = "x"

	withdrawPayload, err := withdrawOp.Encode()

	if err != nil {
		t.Fail()
	}

	op1 := NewOperation()
	op1.Type = OperationType_Deposit
	op1.Payload = depositPayload

	op2 := NewOperation()
	op2.Type = OperationType_Withdraw
	op2.Payload = withdrawPayload

	transaction := NewTransaction()
	transaction.AddOperation(op1)
	transaction.AddOperation(op2)

	transaction.Evaluate()

}
