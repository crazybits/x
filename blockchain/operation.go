package protos

import (
	"errors"

	proto "github.com/golang/protobuf/proto"
)

//NewOperation create operation
func NewOperation(opType OperationType, payload []byte) *Operation {
	op := &Operation{OpType: opType, Payload: payload}
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
