package protos

import (
	"reflect"
	"testing"
)

func TestDepositOperationMarshAndUnmarsh(t *testing.T) {

	depositOp := &DepositOperation{}

	depositOp.Receiver = StringToAddress("craybit")
	depositOp.Amount = int64(100000)
	depositOp.Symbol = "x"

	data, err := depositOp.Encode()

	if err != nil {
		t.Fail()
	}

	depositOp2 := &DepositOperation{}

	err2 := depositOp2.Decode(data)
	if err2 != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(depositOp, depositOp2) {
		t.Fail()

	}

}

func TestOperation(t *testing.T) {

	depositOp := &DepositOperation{}

	depositOp.Receiver = StringToAddress("craybit")
	depositOp.Amount = int64(100000)
	depositOp.Symbol = "x"

	data, err := depositOp.Encode()

	if err != nil {
		t.Fail()
	}

	operation := NewOperation(OperationType_Deposit, data)

	data2, err2 := operation.Encode()

	if err2 != nil {
		t.Fail()
	}

	operation2 := &Operation{}

	operation2.Decode(data2)

	depositOp2 := &DepositOperation{}

	depositOp2.Decode(operation2.Payload)

	if !reflect.DeepEqual(operation, operation2) {
		t.Fail()
	}

	if !reflect.DeepEqual(depositOp, depositOp2) {
		t.Fail()
	}

}

func TestTrnasctionMarshalAndUnmarsh(t *testing.T) {

	depositOp := &DepositOperation{}

	depositOp.Receiver = StringToAddress("craybit")
	depositOp.Amount = int64(100000)
	depositOp.Symbol = "x"

	depositPayload, err := depositOp.Encode()

	if err != nil {
		t.Fail()
	}

	withdrawOp := &WithdrawOperation{}

	withdrawOp.Sender = StringToAddress("craybit")
	withdrawOp.Amount = int64(100000)
	withdrawOp.Symbol = "x"

	withdrawPayload, err := withdrawOp.Encode()

	if err != nil {
		t.Fail()
	}

	operation1 := &Operation{OpType: OperationType_Deposit, Payload: depositPayload}
	operation2 := &Operation{OpType: OperationType_Withdraw, Payload: withdrawPayload}

	transaction := &Transaction{}

	transaction.AddOperation(operation1)
	transaction.AddOperation(operation2)

	data, _ := transaction.Encode()

	transaction2 := &Transaction{}

	transaction2.Decode(data)

	for _, v := range transaction2.Operations {

		switch v.OpType {

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

func TestOperations(t *testing.T) {
	op1 := &DepositOperation{}

	op1.Receiver = StringToAddress("crazybit")
	op1.Amount = int64(1000000)
	op1.Symbol = "x"

	data, err := op1.Encode()

	if err != nil {
		t.Fail()
	}

	op2 := &DepositOperation{}

	op2.Decode(data)

	if !reflect.DeepEqual(op1, op2) {

	}
}
