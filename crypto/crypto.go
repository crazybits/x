package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"io"

	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/crazybits/x/common"

	secp256k1 "github.com/btcsuite/btcd/btcec"
)

//ToECDSA convert bytes to ecdsa private key
func ToECDSA(prv []byte) *ecdsa.PrivateKey {
	if len(prv) == 0 {
		return nil
	}
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = secp256k1.S256()
	priv.D = common.BigD(prv)
	priv.PublicKey.X, priv.PublicKey.Y = secp256k1.S256().ScalarBaseMult(prv)
	return priv
}

//StrToPrivateKey generate private key from given string
func StrToPrivateKey(str string) *secp256k1.PrivateKey {
	seed := common.StrToSha256(str)
	priv := new(secp256k1.PrivateKey)
	priv.PublicKey.Curve = secp256k1.S256()
	priv.D = common.BigD(seed)
	priv.PublicKey.X, priv.PublicKey.Y = secp256k1.S256().ScalarBaseMult(seed)
	return priv
}

//HexToPrivateKey convert hex private key to secp256k1 private key
func HexToPrivateKey(hexkey string) *secp256k1.PrivateKey {

	seed, err := hex.DecodeString(hexkey)
	if err != nil {
		panic(err)
	}
	if len(seed) != 32 {
		panic("private key seed size must be 32 bytes")
	}

	priv := new(secp256k1.PrivateKey)
	priv.PublicKey.Curve = secp256k1.S256()
	priv.D = common.BigD(seed)
	priv.PublicKey.X, priv.PublicKey.Y = secp256k1.S256().ScalarBaseMult(seed)
	return priv
}

//PrivateKeyToHexString convert secp256k1 private key to hex string
func PrivateKeyToHexString(priv *secp256k1.PrivateKey) string {
	return hex.EncodeToString(priv.Serialize())

}

//PublicKeyToHexString convert secp256k1 uncompressed public key to  hex string
func PublicKeyToHexString(pub *secp256k1.PublicKey) string {
	return hex.EncodeToString(pub.SerializeUncompressed())
}

//CompressedPublicKeyToHexString convert secp256k1 compressed public key to hex string
func CompressedPublicKeyToHexString(pub *secp256k1.PublicKey) string {
	return hex.EncodeToString(pub.SerializeCompressed())
}

//Encrypt aes cbc encrption
func Encrypt(password string, plainText []byte) []byte {

	key := common.StrToSha256(password)

	if len(plainText)%aes.BlockSize != 0 {
		panic("plain text size must be multiple of block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))

	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, iv)

	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return cipherText

}

//Decrypt aes cbc decryption
func Decrypt(password string, cipherText []byte) []byte {

	key := common.StrToSha256(password)

	if len(cipherText)%aes.BlockSize != 0 {
		panic("cipher text size must be multiple of block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	iv := cipherText[:aes.BlockSize]

	mode := cipher.NewCBCDecrypter(block, iv)

	decryptedText := make([]byte, len(cipherText)-aes.BlockSize)

	//cipherText[aes.BlockSize:] vi bytes needn't to be decrypted
	mode.CryptBlocks(decryptedText, cipherText[aes.BlockSize:])

	return decryptedText
}
