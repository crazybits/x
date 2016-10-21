package protos

import (
	"errors"

	proto "github.com/golang/protobuf/proto"
)

func NewWithdrawOperation(sender *Address, amount int64, symbol string) *WithdrawOperation {
	withdrawOp := &WithdrawOperation{Sender: sender, Amount: amount, Symbol: symbol}
	return withdrawOp
}

func (withdrawOperation *WithdrawOperation) Bytes() ([]byte, error) {

	data, err := proto.Marshal(withdrawOperation)
	if err != nil {
		return nil, errors.New("unable to marshal the withdraw operation")
	}
	return data, nil
}
