package crypto

import (
	"crypto/ecdsa"

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
