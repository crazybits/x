package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/crazybits/x/common"
)

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
