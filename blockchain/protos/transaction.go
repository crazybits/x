package protos

import (
	"errors"
	"fmt"

	proto "github.com/golang/protobuf/proto"
)

func NewTransaction(operations []*Operation, signatures []*Signature) *Transaction {

	tx := Transaction{Operations: operations, Signatures: signatures}
	return &tx

}

func (transaction *Transaction) AddOperation(operation *Operation) {

	operations := transaction.GetOperations()
	operations = append(operations, operation)
}

func (transaction *Transaction) Validate() bool {

	for _, v := range transaction.Operations {

		fmt.Println(v)

	}
	return true

}

func (transaction *Transaction) Encode() ([]byte, error) {

	data, err := proto.Marshal(transaction)

	if err != nil {

		return nil, errors.New("unable to marshal the transaction")
	}

	return data, nil
}

func (transaction *Transaction) Decode(data []byte) error {

	err := proto.Unmarshal(data, transaction)

	if err != nil {

		return errors.New("unable to marshal the transaction")
	}

	return nil
}
