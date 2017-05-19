package blockchain

import (
	"encoding/hex"
	"math/big"

	"github.com/crazybits/x/crypto"
	"github.com/golang/protobuf/proto"
)

func StringToAddress(s string) *Address                { return BytesToAddress([]byte(s)) }
func BigToAddress(b *big.Int) *Address                 { return BytesToAddress(b.Bytes()) }
func HexToAddress(s string) *Address                   { return BytesToAddress(FromHex(s)) }
func PublicKeyToAddress(key crypto.PublicKey) *Address { return BytesToAddress(key.Key) }
func (address *Address) ToString() string              { return hex.EncodeToString(address.AddressByte) }

//Encode get the bytes of the address
func (address *Address) Encode() []byte {

	data, _ := proto.Marshal(address)
	return data
}

//Decode get the bytes of the address
func (address *Address) Decode(data []byte) {

	proto.Unmarshal(data, address)
}

//BytesToAddress set bytes to the address
func BytesToAddress(b []byte) *Address {
	a := &Address{}
	a.AddressByte = b
	return a
}

//FromHex convert string to bytes
func FromHex(s string) []byte {
	if len(s) > 1 {
		if s[0:2] == "0x" {
			s = s[2:]
		}
		if len(s)%2 == 1 {
			s = "0" + s
		}
		return Hex2Bytes(s)
	}
	return nil
}

//Hex2Bytes convcert hex string to bytes
func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)

	return h
}
