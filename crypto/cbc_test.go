package crypto

import (
	"bytes"
	"testing"

	"github.com/crazybits/x/common"
)

func TestSelfDefinedAESEncryptAndDecrypt(t *testing.T) {

	password := "crzybit"
	plainText := common.StrToSha512("crazybit")

	cipherText := Encrypt(password, plainText)

	decryptText := Decrypt(password, cipherText)

	if !bytes.Equal(plainText, decryptText) {
		t.Fail()
	}

}
