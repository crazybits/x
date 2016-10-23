package blockchain

import (
	"errors"
	"fmt"

	"github.com/crazybits/x/common"
	"github.com/crazybits/x/crypto"
	proto "github.com/golang/protobuf/proto"
)

//NewTransaction create a new transaction
func NewTransaction(operations []*Operation, signatures []*crypto.Signature) *Transaction {

	tx := Transaction{Operations: operations, Signatures: signatures}
	return &tx

}

//AddOperation add operation to the transaction
func (transaction *Transaction) AddOperation(operation *Operation) {

	operations := transaction.GetOperations()
	operations = append(operations, operation)
}

//Validate validae the transaction
func (transaction *Transaction) Validate() bool {

	for _, v := range transaction.Operations {

		//TODO implement validation logic
		fmt.Println(v)

	}
	return true

}

//Encode get the bytes of the transaction
func (transaction *Transaction) Encode() ([]byte, error) {

	data, err := proto.Marshal(transaction)

	if err != nil {
		return nil, errors.New("unable to marshal the transaction")
	}

	return data, nil
}

//Decode get the transaction from the provided bytes
func (transaction *Transaction) Decode(data []byte) error {

	err := proto.Unmarshal(data, transaction)

	if err != nil {
		return errors.New("unable to marshal the transaction")
	}

	return nil
}

//Digest get the transaction digest
func (transaction *Transaction) Digest() ([]byte, error) {

	data, err := transaction.Encode()

	if err != nil {
		return nil, errors.New("unable to get the digest of the transaction")
	}

	digest := common.Sha256(data)

	return digest, nil
}
