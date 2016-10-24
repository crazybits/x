package blockchain

import (
	"errors"

	"github.com/crazybits/x/common"
	"github.com/golang/protobuf/proto"
)

func NewBlockchainInfo() *BlockchainInfo {

	return new(BlockchainInfo)

}

//Encode get the bytes of the BlockchainInfo
func (blockchainInfo *BlockchainInfo) Encode() ([]byte, error) {

	data, err := proto.Marshal(blockchainInfo)
	if err != nil {
		return nil, errors.New("unable to encode the BlockchainInfo")
	}
	return data, nil
}

//Decode get the bytes of the BlockchainInfo
func (blockchainInfo *BlockchainInfo) Decode(data []byte) error {

	err := proto.Unmarshal(data, blockchainInfo)
	if err != nil {
		return errors.New("unable to encode the BlockchainInfo")
	}
	return nil
}

//Digest get the BlockchainInfo digest
func (blockchainInfo *BlockchainInfo) Digest() ([]byte, error) {

	data, err := blockchainInfo.Encode()
	if err != nil {
		return nil, errors.New("unable to get the digest of the BlockchainInfo")
	}
	digest := common.Sha256(data)

	return digest, nil
}
