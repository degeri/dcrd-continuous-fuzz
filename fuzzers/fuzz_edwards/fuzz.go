package fuzz_edwards

import (
	"bytes"
	dcrd_edwards "github.com/decred/dcrd/dcrec/edwards/v2"
)

func Fuzz(input []byte) int {
	dcrd_edwards.ParsePubKey(input)
	dcrd_edwards.PrivKeyFromBytes(input)
	{
		priv, pub, err := dcrd_edwards.PrivKeyFromScalar(input)
		if err == nil {
			hash := []byte{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9}
			r, s, err := dcrd_edwards.Sign(priv, hash)
			if err == nil {
				sig := dcrd_edwards.NewSignature(r, s)
				if !dcrd_edwards.Verify(pub, hash, sig.GetR(), sig.GetS()) {
					panic("dcrd_edwards.Verify")
				}
				pub.Serialize()
				serializedKey := priv.Serialize()
				if !bytes.Equal(serializedKey, input) {
					panic("dcrd_edwards.Edwards: key not equal")
				}
			}
		}
	}
	dcrd_edwards.ParseDERSignature(input)
	dcrd_edwards.ParseSignature(input)
	return 0
}
