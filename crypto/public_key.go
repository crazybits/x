package crypto

import (
	secp256k1 "github.com/btcsuite/btcd/btcec"
	"github.com/crazybits/x/common"
)

func (publicKey *PublicKey) VerifySign(hash []byte, signature *Signature) bool {

	pub, err := secp256k1.ParsePubKey(publicKey.Key, secp256k1.S256())
	if err != nil {
		panic("unable to get secp256K1 public key form the privded data")
	}

	sig := &secp256k1.Signature{S: common.BytesToBig(signature.S), R: common.BytesToBig(signature.R)}

	return sig.Verify(hash, pub)
}
