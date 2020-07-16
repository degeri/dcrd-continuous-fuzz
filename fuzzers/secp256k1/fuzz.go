package secp256k1

import (
	"bytes"
	dcrd_secp256k1 "github.com/decred/dcrd/dcrec/secp256k1/v3"
	dcrd_ecdsa "github.com/decred/dcrd/dcrec/secp256k1/v3/ecdsa"
)

func Fuzz(input []byte) int {
	{
		pk, err := dcrd_secp256k1.ParsePubKey(input)
		if err == nil {
			pk1 := pk.SerializeUncompressed()
			pk2 := pk.SerializeCompressed()
			if !bytes.Equal(pk1, input) {
				if !bytes.Equal(pk2, input) {
					panic("Serialization error")
				}
			}
		}
	}
	{
		priv := dcrd_secp256k1.PrivKeyFromBytes(input)
		pub := priv.PubKey()
		_, err := dcrd_secp256k1.ParsePubKey(pub.SerializeUncompressed())
		if err == nil {
			hash := []byte{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9}
			sig := dcrd_ecdsa.Sign(priv, hash)

			if !sig.Verify(hash, pub) {
				panic("dcrd_ecdsa.Verify")
			}

			pub.SerializeCompressed()
			serializedKey := priv.Serialize()
			if !bytes.Equal(serializedKey, input) {
				panic("dcrd_secp256k1.Secp256k1: key not equal")
			}
		}
	}
	{
		dcrd_ecdsa.ParseDERSignature(input)
	}
	return 0
}
