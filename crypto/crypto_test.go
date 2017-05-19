package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/crazybits/x/common"

	secp256k1 "github.com/btcsuite/btcd/btcec"
)

func TestSignAndVerify(t *testing.T) {

	var pub *secp256k1.PublicKey

	var pri *secp256k1.PrivateKey

	var sign *secp256k1.Signature

	var err error

	pri, err = secp256k1.NewPrivateKey(secp256k1.S256())

	if err != nil {
		t.Error("generated private key failed")
	}

	pub = pri.PubKey()

	hash := common.StrToSha512("crazybit")

	sign, err = pri.Sign(hash)

	if err == nil {

		if sign.Verify(hash, pub) {

			fmt.Println("signature verify passed")
		}

	}

	fmt.Println(hex.EncodeToString(pub.SerializeCompressed()), "compressed public key bytes length=", len(pub.SerializeCompressed()))
	fmt.Println(hex.EncodeToString(pub.SerializeUncompressed()), "uncompressed public bytes key length=", len(pub.SerializeUncompressed()))
	fmt.Println(hex.EncodeToString(sign.Serialize()), "uncompacted signature bytes length=", len(sign.Serialize()))

}

func TestAESEncryptAndDecrypt(t *testing.T) {

	key := common.StrToSha256("crazybit")
	plainText := common.StrToSha512("crazybit")

	if len(plainText)%aes.BlockSize != 0 {
		panic("plain text should be a multiple of the block size")
	}

	keyBlock, err := aes.NewCipher(key)

	if err != nil {

		panic(err)
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {

		panic(err)

	}

	mode := cipher.NewCBCEncrypter(keyBlock, iv)

	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	iv2 := cipherText[:aes.BlockSize]

	if len(cipherText)%aes.BlockSize != 0 {
		panic("cipher text should be multiple of the block size")
	}

	mode2 := cipher.NewCBCDecrypter(keyBlock, iv2)

	decryptedText := make([]byte, len(cipherText)-aes.BlockSize)
	mode2.CryptBlocks(decryptedText, cipherText[aes.BlockSize:])

	if !bytes.Equal(plainText, decryptedText) {
		t.Fail()
	}
}

func TestECCEncrypAndDecrypt(t *testing.T) {

	pri, err := secp256k1.NewPrivateKey(secp256k1.S256())
	if err != nil {

		t.Fail()
	}

	pub := pri.PubKey()

	plaintex := common.StrToSha512("crazybit")

	encryptedtext, err := secp256k1.Encrypt(pub, plaintex)

	if err != nil {
		t.Fail()
	}

	dectext, err := secp256k1.Decrypt(pri, encryptedtext)

	if err != nil {
		t.Fail()
	}

	if bytes.Compare(plaintex, dectext) != 0 {
		t.Fail()
	}

}

func TestGenerateSharedSecret(t *testing.T) {

	pri, err := secp256k1.NewPrivateKey(secp256k1.S256())

	if err != nil {

		t.Error(err)
	}

	pri2, err := secp256k1.NewPrivateKey(secp256k1.S256())

	if err != nil {

		t.Error(err)
	}

	sec1 := secp256k1.GenerateSharedSecret(pri, pri2.PubKey())

	sec2 := secp256k1.GenerateSharedSecret(pri2, pri.PubKey())

	if !bytes.Equal(sec1, sec2) {

		t.Fail()

	}

}

func TestGenPrivatekeyFromSeed(t *testing.T) {

	var expectPrivateKeyHexString string = "d8f723201d8c05b93a8893af1940a0bf8c8c4343aca442dd1e95569980e9d83f"
	var expectPublickeycompressedHexString string = "034dfa4d97c490f64a4920a8b0e38650b9cf59e51833bdefbcd059e1ae37159d0e"
	var expectPublicKeyUncompressedHexString string = "044dfa4d97c490f64a4920a8b0e38650b9cf59e51833bdefbcd059e1ae37159d0e7e8ea115ba479ccd6c225b45b3e718b42d6befd85dab843cdd20961f94c1f167"

	seed := common.StrToSha256("crazybit")
	pri, pub := secp256k1.PrivKeyFromBytes(secp256k1.S256(), seed)

	if expectPrivateKeyHexString != hex.EncodeToString(pri.Serialize()) {
		t.Fail()
	}

	if expectPublickeycompressedHexString != hex.EncodeToString(pub.SerializeCompressed()) {
		t.Fail()
	}

	if expectPublicKeyUncompressedHexString != hex.EncodeToString(pub.SerializeUncompressed()) {
		t.Fail()
	}

}

func TestRecoveryPublicKeyFromCompactedSignature(t *testing.T) {

	pri, err := secp256k1.NewPrivateKey(secp256k1.S256())
	if err != nil {
		t.Fail()
	}

	pub := pri.PubKey()

	hash := common.StrToSha512("crazybit")

	sign, err := secp256k1.SignCompact(secp256k1.S256(), pri, hash, true)

	if err != nil {
		t.Fail()
	}

	fmt.Println("compacted signature byte length=", len(sign))

	if len(sign) != 65 {

		t.Error("compacted signature byte length=", len(sign))
	}

	rpub, _, err := secp256k1.RecoverCompact(secp256k1.S256(), sign, hash)

	if err != nil {
		t.Fail()
	}

	if !pub.IsEqual(rpub) {
		t.Fail()
	}

	fmt.Println("pub key         =", hex.EncodeToString(pub.SerializeCompressed()))
	fmt.Println("recover pub key =", hex.EncodeToString(rpub.SerializeCompressed()))
}

func TestGetPrivateKeyFromStr(t *testing.T) {
	var expectPrivateKeyHexString string = "d8f723201d8c05b93a8893af1940a0bf8c8c4343aca442dd1e95569980e9d83f"

	priv := StrToPrivateKey("crazybit")

	if !strings.EqualFold(PrivateKeyToHexString(priv), expectPrivateKeyHexString) {
		t.Fail()

	}

}

func TestPublicKeyToUnCompressedHexString(t *testing.T) {

	var expectPublickeycompressedHexString string = "044dfa4d97c490f64a4920a8b0e38650b9cf59e51833bdefbcd059e1ae37159d0e7e8ea115ba479ccd6c225b45b3e718b42d6befd85dab843cdd20961f94c1f167"

	priv := StrToPrivateKey("crazybit")

	pub := priv.PubKey()

	if !strings.EqualFold(expectPublickeycompressedHexString, PublicKeyToHexString(pub)) {
		t.Fail()
	}
}

func TestSecp256k1AndEllipticSerilizationDifferent(t *testing.T) {

	priv := StrToPrivateKey("crazybit")

	pub := priv.PubKey()

	bytes := elliptic.Marshal(secp256k1.S256(), pub.X, pub.Y)

	if !strings.EqualFold(common.BytesToHexString(bytes), PublicKeyToHexString(pub)) {
		t.Fail()
	}

}

func TestHexStringToPrivateKey(t *testing.T) {

	var hexkey string = "d8f723201d8c05b93a8893af1940a0bf8c8c4343aca442dd1e95569980e9d83f"

	genekey := HexToPrivateKey(hexkey).Serialize()

	key := StrToPrivateKey("crazybit").Serialize()

	if !bytes.Equal(genekey, key) {

		t.Fail()
	}
}
