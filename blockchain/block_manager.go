package blockchain

import (
	pb "github.com/crazybits/x/blockchain/protos"
)

type BlockManager struct {
}

func NewBlockManager() *BlockManager {
	bm := BlockManager{}
	return &bm
}
func (bm *BlockManager) ProcessBlock(block *pb.Block) error {

	txCount := len(block.Transactions)
	lockChan := make(chan bool, txCount)

	for _, tx := range block.Transactions {
		go bm.ProccessTransaction(tx, lockChan)
	}

	for i := 0; i < txCount; i++ {
		<-lockChan
	}
	return nil

}

func (bm *BlockManager) ProccessTransaction(tx *pb.Transaction, lockChan chan bool) {

	tx.Validate()

	lockChan <- true

	return

}
