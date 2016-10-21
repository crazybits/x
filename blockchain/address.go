package protos

import (
	"encoding/hex"
	"math/big"
)

func StringToAddress(s string) *Address { return BytesToAddress([]byte(s)) }
func BigToAddress(b *big.Int) *Address  { return BytesToAddress(b.Bytes()) }
func HexToAddress(s string) *Address    { return BytesToAddress(FromHex(s)) }

//BytesToAddress set bytes to the address
func BytesToAddress(b []byte) *Address {
	a := &Address{}
	a.AddressByte = b
	return a
}

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

func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)

	return h
}
