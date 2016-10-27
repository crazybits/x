package blockchain

import (
	"sync"

	"github.com/crazybits/x/common"
	logging "github.com/op/go-logging"
)

var logger = logging.MustGetLogger("blockchain")

const blockchainInfoKey = "BLOCKCHAIN_INFO_KEY"

var bm *BlockManager

var bmOnce sync.Once

type BlockManager struct {
	blockchainDB *BlockchainDB
	balanceDB    *BalanceDB
	state        *State
}

func NewBlockManager() *BlockManager {

	bmOnce.Do(func() {
		bm = new(BlockManager)

		bm.blockchainDB = NewBlcokchainDB()
		bm.balanceDB = NewBalanceDB()
		bm.state = NewState()

	})
	return bm
}

func (bm *BlockManager) PushBlock(block *Block) {

	blockID, err := block.Digest()
	if err != nil {
		logger.Errorf("failed to get the block id:%s", err)
	}
	blockBytes, err := block.Encode()

	if err != nil {
		logger.Errorf("failed to get the block bytes:%s", err)
	}
	err = bm.blockchainDB.Put(blockID, blockBytes)
	if err != nil {
		logger.Errorf("failed to push block:%s", err)
	}

}

func (bm *BlockManager) GetBlockByID(id []byte) *Block {

	data, err := bm.blockchainDB.Get(id)
	if err != nil {
		logger.Errorf("failed to get block by id=%s:%s", common.BytesToHexSting(id), err)
	}

	block := NewBlock()

	err = block.Decode(data)
	if err != nil {
		logger.Errorf("failed to decode the block:%s", err)
	}
	return block
}
func (bm *BlockManager) UpdateBlockchainInfo(blochainInfo *BlockchainInfo) {

	value, err := blochainInfo.Encode()
	if err != nil {
		logger.Errorf("failed to encode of blochainInfo: %s", err)
	}
	err = bm.blockchainDB.Put(common.StrToSha256(blockchainInfoKey), value)
	if err != nil {
		logger.Errorf("failed to upate blockhcainInfo to db:%s", err)
	}
}

func (bm *BlockManager) GetBlockchainInfo() *BlockchainInfo {

	blockchainInfor := NewBlockchainInfo()
	data, err := bm.blockchainDB.Get(common.StrToSha256(blockchainInfoKey))
	if err != nil {
		logger.Errorf("failed to get the blockchainInfor:%s", err)
		return nil
	}

	err = blockchainInfor.Decode(data)
	if err != nil {
		logger.Errorf("failed to decode the blockchainInfor:%s", err)
		return nil
	}
	return blockchainInfor
}

func (bm *BlockManager) ProcessBlock(block *Block) error {

	for _, tx := range block.Transactions {
		bm.ProccessTransaction(tx)
	}

	return nil

}

func (bm *BlockManager) ProccessTransaction(tx *Transaction) {

	if tx.Evaluate(bm.state) {

		bm.state.Apply()
	}

	return

}

func (bm *BlockManager) ShutDown() {
	bm.blockchainDB.Close()

}
