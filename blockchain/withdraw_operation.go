package blockchain

import (
	"errors"
	"fmt"

	proto "github.com/golang/protobuf/proto"
)

//NewWithdrawOperation create a new withdraw operation
func NewWithdrawOperation() *WithdrawOperation {
	withdrawOp := new(WithdrawOperation)
	return withdrawOp
}

//Encode serilize the withdraw operation to bytes
func (withdrawOperation *WithdrawOperation) Encode() ([]byte, error) {

	data, err := proto.Marshal(withdrawOperation)
	if err != nil {
		return nil, errors.New("unable to encode the withdraw operation")
	}
	return data, nil
}

//Decode unserilize the withdraw operation from the provided bytes
func (withdrawOperation *WithdrawOperation) Decode(data []byte) error {

	err := proto.Unmarshal(data, withdrawOperation)
	if err != nil {
		return errors.New("unable to decode the withdraw operation")
	}
	return nil
}

func (withdrawOperation *WithdrawOperation) Evaluate(state *State) bool {
	fmt.Println("WithdrawOperation Evaluate")
	return true //TODO impmentation

}
