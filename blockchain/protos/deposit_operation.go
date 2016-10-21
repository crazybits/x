package protos

import (
	"errors"

	proto "github.com/golang/protobuf/proto"
)

func NewDepositOperation(receiver *Address, amount int64, symbol string) *DepositOperation {
	depositOp := &DepositOperation{Receiver: receiver, Amount: amount, Symbol: symbol}
	return depositOp
}

func (depositOperation *DepositOperation) Encode() ([]byte, error) {

	data, err := proto.Marshal(depositOperation)
	if err != nil {
		return nil, errors.New("unable to encode the deposit operation")
	}
	return data, nil
}

func (depositOperation *DepositOperation) Decode(data []byte) error {

	err := proto.Unmarshal(data, depositOperation)
	if err != nil {
		return errors.New("unable to decode the deposit operation")
	}
	return nil
}
