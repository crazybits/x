package protos

import (
	"errors"

	proto "github.com/golang/protobuf/proto"
)

func NewWithdrawOperation(sender *Address, amount int64, symbol string) *WithdrawOperation {
	withdrawOp := &WithdrawOperation{Sender: sender, Amount: amount, Symbol: symbol}
	return withdrawOp
}

func (withdrawOperation *WithdrawOperation) Encode() ([]byte, error) {

	data, err := proto.Marshal(withdrawOperation)
	if err != nil {
		return nil, errors.New("unable to encode the withdraw operation")
	}
	return data, nil
}

func (withdrawOperation *WithdrawOperation) Decode(data []byte) error {

	err := proto.Unmarshal(data, withdrawOperation)
	if err != nil {
		return errors.New("unable to decode the withdraw operation")
	}
	return nil
}
