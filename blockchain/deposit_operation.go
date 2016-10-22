package protos

import (
	"errors"

	proto "github.com/golang/protobuf/proto"
)

//NewDepositOperation create a new deposit operation
func NewDepositOperation(receiver *Address, amount int64, symbol string) *DepositOperation {
	depositOp := &DepositOperation{Receiver: receiver, Amount: amount, Symbol: symbol}
	return depositOp
}

//Encode serilize the deposit operation to bytes
func (depositOperation *DepositOperation) Encode() ([]byte, error) {

	data, err := proto.Marshal(depositOperation)
	if err != nil {
		return nil, errors.New("unable to encode the deposit operation")
	}
	return data, nil
}

//Decode unserilize the deposit operation from the provided bytes
func (depositOperation *DepositOperation) Decode(data []byte) error {

	err := proto.Unmarshal(data, depositOperation)
	if err != nil {
		return errors.New("unable to decode the deposit operation")
	}
	return nil
}
