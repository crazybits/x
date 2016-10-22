package protos

import (
	"encoding/hex"
	"math/big"
)

func StringToAddress(s string) *Address         { return BytesToAddress([]byte(s)) }
func BigToAddress(b *big.Int) *Address          { return BytesToAddress(b.Bytes()) }
func HexToAddress(s string) *Address            { return BytesToAddress(FromHex(s)) }
func PublicKeyToAddress(key PublicKey) *Address { return BytesToAddress(key.key) }

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
