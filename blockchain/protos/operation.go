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

//Bytes get the bytes of the operation
func (operation *Operation) Bytes() ([]byte, error) {

	data, err := proto.Marshal(operation)
	if err != nil {
		return nil, errors.New("unable to marshal the operation")
	}
	return data, nil
}
