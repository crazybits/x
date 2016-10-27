package blockchain

import (
	"errors"
	"fmt"

	proto "github.com/golang/protobuf/proto"
)

type IOperation interface {
	Evaluate(state *State) bool
}

//NewOperation create operation
func NewOperation() *Operation {
	op := new(Operation)
	return op
}

//Encode get the bytes of the operation
func (operation *Operation) Encode() ([]byte, error) {

	data, err := proto.Marshal(operation)
	if err != nil {
		return nil, errors.New("unable to encode the operation")
	}
	return data, nil
}

//Decode get the bytes of the operation
func (operation *Operation) Decode(data []byte) error {

	err := proto.Unmarshal(data, operation)
	if err != nil {
		return errors.New("unable to encode the operation")
	}
	return nil
}
func (operation *Operation) Evaluate(state *State) bool {

	switch operation.Type {

	case OperationType_Deposit:

		op := &DepositOperation{}
		op.Decode(operation.Payload)
		if op.Evaluate(state) {
			return true
		}
	case OperationType_Withdraw:

		op := &WithdrawOperation{}
		op.Decode(operation.Payload)
		if op.Evaluate(state) {
			return true
		}
	default:
		fmt.Println("not supported yet")

	}
	return false

}
