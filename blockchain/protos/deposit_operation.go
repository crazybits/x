package protos

import (
	"errors"

	proto "github.com/golang/protobuf/proto"
)

func NewDepositOperation(receiver *Address, amount int64, symbol string) *DepositOperation {
	depositOp := &DepositOperation{Receiver: receiver, Amount: amount, Symbol: symbol}
	return depositOp
}

func (depositOperation *DepositOperation) Bytes() ([]byte, error) {

	data, err := proto.Marshal(depositOperation)
	if err != nil {
		return nil, errors.New("unable to marshal the deposit operation")
	}
	return data, nil
}
