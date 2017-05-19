package crypto

import (
	"testing"

	"bytes"

	secp256k1 "github.com/btcsuite/btcd/btcec"
	"github.com/crazybits/x/common"
)

func TestPrivateKeyWraper(t *testing.T) {

	pri := GenePrivateKeyFromSeed(common.StrToSha256("crazybit"))
	pub := pri.PublicKey()

	pri2, pub2 := secp256k1.PrivKeyFromBytes(secp256k1.S256(), common.StrToSha256("crazybit"))

	if !bytes.Equal(pri.Key, pri2.D.Bytes()) {
		t.Fail()
	}
	if !bytes.Equal(pub.Key, pub2.SerializeCompressed()) {
		t.Fail()

	}
}

func TestWraperSignature(t *testing.T) {

	seed := common.StrToSha256("crazybit")
	hash := common.StrToSha256("crazybit test signature")

	pri := GenePrivateKeyFromSeed(seed)
	pub := pri.PublicKey()

	sig := pri.Sign(hash)

	pri2, _ := secp256k1.PrivKeyFromBytes(secp256k1.S256(), seed)

	sig2, err := pri2.Sign(hash)

	if err != nil {
		t.Fail()
	}

	if !bytes.Equal(sig.S, sig2.S.Bytes()) {
		t.Fail()
	}
	if !bytes.Equal(sig.R, sig2.R.Bytes()) {
		t.Fail()
	}

	if !pub.VerifySign(hash, sig) {
		t.Fail()
	}
}

func TestGenSharedScrete(t *testing.T) {

	seed := common.StrToSha256("crazybit")

	seed2 := common.StrToSha256("crazybit2")

	pri := GenePrivateKeyFromSeed(seed)

	pri2 := GenePrivateKeyFromSeed(seed2)

	secret := pri.GenerateSharedSecret(pri2.PublicKey())
	secret2 := pri2.GenerateSharedSecret(pri.PublicKey())
	if !bytes.Equal(secret, secret2) {
		t.Fail()
	}

	secpPri, secpPub := secp256k1.PrivKeyFromBytes(secp256k1.S256(), seed)
	secpPri2, secpPub2 := secp256k1.PrivKeyFromBytes(secp256k1.S256(), seed2)

	secpSecrete1 := secp256k1.GenerateSharedSecret(secpPri, secpPub2)
	secpSecrete2 := secp256k1.GenerateSharedSecret(secpPri2, secpPub)

	if !bytes.Equal(secpSecrete2, secret) || !bytes.Equal(secpSecrete2, secret2) || !bytes.Equal(secpSecrete1, secpSecrete2) {
		t.Fail()
	}
}
