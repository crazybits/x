package protos

import (
	"errors"

	proto "github.com/golang/protobuf/proto"
	google "github.com/golang/protobuf/ptypes/timestamp"
)

//NewBlockHeader create block header
func NewBlockHeader(time *google.Timestamp, number int64, nance []byte, currentHash []byte, prevHash []byte, signature *Signature) *BlockHeader {

	h := BlockHeader{Time: time, Number: number, Nance: nance, CurrHash: currentHash, PrevHash: prevHash, Signature: signature}
	return &h
}

//Bytes get the block bytes
func (blockHeader *BlockHeader) Bytes() ([]byte, error) {

	data, err := proto.Marshal(blockHeader)
	if err != nil {
		return nil, errors.New("unable to marshal the block header")
	}
	return data, nil

}
