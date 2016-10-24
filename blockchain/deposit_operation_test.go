package blockchain

import (
	"reflect"
	"testing"
)

func TestDepositOperationMarshAndUnmarsh(t *testing.T) {

	depositOp := NewDepositOperation()

	depositOp.Receiver = StringToAddress("craybit")
	depositOp.Amount = int64(100000)
	depositOp.Symbol = "x"

	data, err := depositOp.Encode()

	if err != nil {
		t.Fail()
	}

	depositOp2 := NewDepositOperation()

	err2 := depositOp2.Decode(data)

	if err2 != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(depositOp, depositOp2) {
		t.Fail()

	}

}

func TestDepositOperations(t *testing.T) {

	op1 := NewDepositOperation()

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
		t.Fail()
	}
}
