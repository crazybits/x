package protos

import (
	"fmt"
	"reflect"
	"testing"

	proto "github.com/golang/protobuf/proto"
)

func TestDepositOperationMarshAndUnmarsh(t *testing.T) {

	receiver := StringToAddress("crazybit")

	symbol := "Symbol"

	amount := int64(1024)

	depositOp := NewDepositOperation(receiver, amount, symbol)

	data, err := depositOp.Bytes()

	if err != nil {
		t.Fail()
	}

	depositOp2 := &DepositOperation{}

	err2 := proto.Unmarshal(data, depositOp2)
	if err2 != nil {
		t.Fail()
	}

	fmt.Println(depositOp)
	fmt.Println(depositOp2)
	if !reflect.DeepEqual(depositOp, depositOp2) {
		t.Fail()

	}

}

func TestOperation(t *testing.T) {

	receiver := StringToAddress("crazybit")

	symbol := "Symbol"

	amount := int64(1024)

	depositOp := NewDepositOperation(receiver, amount, symbol)

	data, err := depositOp.Bytes()

	if err != nil {
		t.Fail()
	}

	operation := NewOperation(OperationType_Deposit, data)

	data2, err2 := operation.Bytes()

	if err2 != nil {
		t.Fail()
	}

	operation2 := &Operation{}

	proto.Unmarshal(data2, operation2)

	depositOp2 := &DepositOperation{}

	proto.Unmarshal(operation.Payload, depositOp2)

	fmt.Println(depositOp)
	fmt.Println(depositOp2)

	if !reflect.DeepEqual(depositOp, depositOp2) {
		t.Fail()

	}

}

func TestTrnasctionMarshalAndUnmarsh(t *testing.T) {

	receiver := StringToAddress("crazybit")

	withdrawSymbol := "Symbol"

	withdrawAmount := int64(1024)

	depositOp := NewDepositOperation(receiver, withdrawAmount, withdrawSymbol)

	depositOpData, _ := depositOp.Bytes()

	sender := StringToAddress("crazybit2")

	symbol := "Symbol"

	amount := int64(1024)

	withdrawOp := NewWithdrawOperation(sender, amount, symbol)

	withdrawOpData, _ := withdrawOp.Bytes()

	operation1 := &Operation{OpType: OperationType_Deposit, Payload: depositOpData}
	operation2 := &Operation{OpType: OperationType_Withdraw, Payload: withdrawOpData}

	transaction := &Transaction{}

	transaction.AddOperation(operation1)
	transaction.AddOperation(operation2)

	data, _ := transaction.Bytes()

	transaction2 := &Transaction{}

	proto.Unmarshal(data, transaction2)

	for _, v := range transaction2.Operations {

		switch v.OpType {

		case OperationType_Deposit:

			depositOp2 := &DepositOperation{}
			proto.Unmarshal(v.Payload, depositOp2)
			fmt.Println(depositOp)
			fmt.Println(depositOp2)

			if !reflect.DeepEqual(depositOp, depositOp2) {
				t.Fail()
			}
		case OperationType_Withdraw:
			withdrawOp2 := &WithdrawOperation{}
			proto.Unmarshal(v.Payload, withdrawOp2)
			fmt.Println(withdrawOp)
			fmt.Println(withdrawOp2)
			if !reflect.DeepEqual(withdrawOp, withdrawOp2) {
				t.Fail()
			}

		}
	}

}
