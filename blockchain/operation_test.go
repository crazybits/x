package blockchain

import (
	"reflect"
	"testing"
)

func TestOperation(t *testing.T) {

	depositOp := &DepositOperation{}

	depositOp.Receiver = StringToAddress("craybit")
	depositOp.Amount = int64(100000)
	depositOp.Symbol = "x"

	data, err := depositOp.Encode()

	if err != nil {
		t.Fail()
	}

	op1 := NewOperation()
	op1.Type = OperationType_Deposit
	op1.Payload = data

	data2, err2 := op1.Encode()

	if err2 != nil {
		t.Fail()
	}

	op2 := NewOperation()

	op2.Decode(data2)

	depositOp2 := NewDepositOperation()

	depositOp2.Decode(op2.Payload)

	if !reflect.DeepEqual(op1, op2) {
		t.Fail()
	}

	if !reflect.DeepEqual(depositOp, depositOp2) {
		t.Fail()
	}

}
