package fuzz_address

import (
	dcrd_chaincfg "github.com/decred/dcrd/chaincfg/v3"
	dcrd_stdaddr "github.com/decred/dcrd/txscript/v4/stdaddr"
)

func Fuzz(input []byte) int {
	{
		addr, err := dcrd_stdaddr.DecodeAddress(string(input), dcrd_chaincfg.MainNetParams())
		if err == nil {
			addr.String()
			addr.PaymentScript()
		}
	}

	dcrd_stdaddr.NewAddressPubKeyEcdsaSecp256k1Raw(0, input, dcrd_chaincfg.MainNetParams())
	dcrd_stdaddr.NewAddressPubKeyEd25519Raw(0, input, dcrd_chaincfg.MainNetParams())
	dcrd_stdaddr.NewAddressPubKeyHashEcdsaSecp256k1(0, input, dcrd_chaincfg.MainNetParams())
	dcrd_stdaddr.NewAddressPubKeyHashEd25519(0, input, dcrd_chaincfg.MainNetParams())
	dcrd_stdaddr.NewAddressPubKeyHashSchnorrSecp256k1(0, input, dcrd_chaincfg.MainNetParams())
	dcrd_stdaddr.NewAddressPubKeySchnorrSecp256k1Raw(0, input, dcrd_chaincfg.MainNetParams())
	dcrd_stdaddr.NewAddressScriptHash(0, input, dcrd_chaincfg.MainNetParams())
	dcrd_stdaddr.NewAddressScriptHashFromHash(0, input, dcrd_chaincfg.MainNetParams())
	return 0
}