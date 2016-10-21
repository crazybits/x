package protos

import (
	"errors"

	proto "github.com/golang/protobuf/proto"
)

//NewBlock create a block
func NewBlock(header *BlockHeader, transactions []*Transaction) *Block {
	block := Block{Header: header, Transactions: transactions}
	return &block

}

//Bytes get the hash of the block
func (block *Block) Bytes() ([]byte, error) {

	data, err := proto.Marshal(block)
	if err != nil {
		return nil, errors.New("unable to marshal the block")
	}
	return data, nil
}

func (block *Block) AddTransaction(tx *Transaction) {

	transactions := block.GetTransactions()

	transactions = append(transactions, tx)

}
