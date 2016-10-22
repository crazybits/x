package crypto

import secp256k1 "github.com/btcsuite/btcd/btcec"

//NewPrivateKey create private key
func NewPrivateKey() *PrivateKey {

	pri, err := secp256k1.NewPrivateKey(secp256k1.S256())
	if err != nil {
		panic("unable to get private key")
	}

	return &PrivateKey{Type: CryptoType_ECDSA, Key: pri.D.Bytes()}
}

//GenePrivateKeyFromSeed generate private key from seed
func GenePrivateKeyFromSeed(hash []byte) *PrivateKey {

	pri, _ := secp256k1.PrivKeyFromBytes(secp256k1.S256(), hash)
	return &PrivateKey{Type: CryptoType_ECDSA, Key: pri.D.Bytes()}
}

//Sign sign the provided data
func (privateKey *PrivateKey) Sign(hash []byte) *Signature {

	pri, _ := secp256k1.PrivKeyFromBytes(secp256k1.S256(), privateKey.Key)

	sig, err := pri.Sign(hash)
	if err != nil {
		panic("unable to sign the privde data")
	}
	return &Signature{S: sig.S.Bytes(), R: sig.R.Bytes(), Type: CryptoType_ECDSA}

}

//PublicKey get the compressed public key
func (privateKey *PrivateKey) PublicKey() *PublicKey {

	_, pub := secp256k1.PrivKeyFromBytes(secp256k1.S256(), privateKey.Key)

	return &PublicKey{Key: pub.SerializeCompressed(), Type: CryptoType_ECDSA}
}

func (privateKey *PrivateKey) GenerateSharedSecret(publicKey *PublicKey) []byte {

	pri, _ := secp256k1.PrivKeyFromBytes(secp256k1.S256(), privateKey.Key)

	pub, err := secp256k1.ParsePubKey(publicKey.Key, secp256k1.S256())
	if err != nil {
		panic("unable to get secp256K1 publick form the privded data")
	}

	return secp256k1.GenerateSharedSecret(pri, pub)

}
