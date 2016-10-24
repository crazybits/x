package blockchain

import (
	"errors"

	"github.com/crazybits/x/common"
	"github.com/crazybits/x/crypto"
	proto "github.com/golang/protobuf/proto"
)

// //NewTransaction create a new transaction
// func NewTransaction(operations []*Operation, signatures []*crypto.Signature) *Transaction {

// 	tx := Transaction{Operations: operations, Signatures: signatures}
// 	return &tx

// }

func NewTransaction() *Transaction {

	tx := new(Transaction)
	tx.Operations = make([]*Operation, 0)
	tx.Signatures = make([]*crypto.Signature, 0)

	return tx
}

//AddOperation add operation to the transaction
func (transaction *Transaction) AddOperation(operation *Operation) {

	operations := transaction.GetOperations()
	transaction.Operations = append(operations, operation)
}

//AddOperation add signature to the transaction
func (transaction *Transaction) AddSignature(sign *crypto.Signature) {

	sigs := transaction.GetSignatures()
	transaction.Signatures = append(sigs, sign)
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

func (transaction *Transaction) Evaluate() bool {

	var op IOperation

	for _, op = range transaction.Operations {
		op.Evaluate()
	}

	return true

}
