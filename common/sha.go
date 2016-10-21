package common

import (
	"crypto/sha256"
	"crypto/sha512"

	"golang.org/x/crypto/ripemd160"
)

//Sha256 SHA256 hash
func Sha256(in []byte) []byte {

	hasher := sha256.New()
	hasher.Write(in)
	return hasher.Sum(nil)

}

//Sha512 SHA512 hash
func Sha512(in []byte) []byte {
	hasher := sha512.New()
	hasher.Write(in)
	return hasher.Sum(nil)
}

//StrToSha256 string to SHA256
func StrToSha256(str string) []byte {

	hash := []byte(str)

	hasher := sha256.New()

	hasher.Write(hash)

	return hasher.Sum(nil)

}

//StrToSha512 string to SHA256
func StrToSha512(str string) []byte {

	hash := []byte(str)

	hasher := sha512.New()

	hasher.Write(hash)

	return hasher.Sum(nil)

}

//Ripemd160 Riple hash
func Ripemd160(bytes []byte) []byte {

	hasher := ripemd160.New()
	hasher.Write(bytes)
	return hasher.Sum(nil)
}
