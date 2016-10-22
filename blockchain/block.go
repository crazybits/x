package protos

import (
	"errors"

	"github.com/crazybits/x/common"
	proto "github.com/golang/protobuf/proto"
)

//NewBlock create a block
func NewBlock(header *BlockHeader, transactions []*Transaction) *Block {
	block := Block{Header: header, Transactions: transactions}
	return &block

}

//Encode get the bytes of the operation
func (block *Block) Encode() ([]byte, error) {

	data, err := proto.Marshal(block)
	if err != nil {
		return nil, errors.New("unable to encode the block")
	}
	return data, nil
}

//Decode get the bytes of the operation
func (block *Block) Decode(data []byte) error {

	err := proto.Unmarshal(data, block)
	if err != nil {
		return errors.New("unable to encode the block")
	}
	return nil
}

//Digest get the block digest
func (block *Block) Digest() ([]byte, errr) {

	data, err := block.Encode()
	if err != nil {
		return errors.New("unable to get the digest of the block")
	}

	digest := common.Sha256(data)

	return digest, nil
}

//AddTransaction add transaction to the block
func (block *Block) AddTransaction(tx *Transaction) {

	transactions := block.GetTransactions()

	transactions = append(transactions, tx)

}
