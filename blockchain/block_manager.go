package blockchain

type BlockManager struct {
}

func NewBlockManager() *BlockManager {
	bm := BlockManager{}
	return &bm
}
func (bm *BlockManager) ProcessBlock(block *Block) error {

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

func (bm *BlockManager) ProccessTransaction(tx *Transaction, lockChan chan bool) {

	tx.Validate()

	lockChan <- true

	return

}
