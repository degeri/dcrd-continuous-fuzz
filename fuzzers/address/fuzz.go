package address

import (
	dcrd_chaincfg "github.com/decred/dcrd/chaincfg/v3"
	dcrd_ec "github.com/decred/dcrd/dcrec"
	dcrd_util "github.com/decred/dcrd/dcrutil/v3"
)

func Fuzz(input []byte) int {
	{
		addr, err := dcrd_util.DecodeAddress(string(input), dcrd_chaincfg.MainNetParams())
		if err == nil {
			addr.String()
			addr.Address()
			addr.ScriptAddress()
		}
	}

	dcrd_util.NewAddressPubKey(input, dcrd_chaincfg.MainNetParams())
	dcrd_util.NewAddressPubKeyHash(input, dcrd_chaincfg.MainNetParams(), dcrd_ec.STEcdsaSecp256k1)
	dcrd_util.NewAddressPubKeyHash(input, dcrd_chaincfg.MainNetParams(), dcrd_ec.STEd25519)
	dcrd_util.NewAddressPubKeyHash(input, dcrd_chaincfg.MainNetParams(), dcrd_ec.STSchnorrSecp256k1)
	dcrd_util.NewAddressScriptHash(input, dcrd_chaincfg.MainNetParams())
	dcrd_util.NewAddressScriptHashFromHash(input, dcrd_chaincfg.MainNetParams())
	dcrd_util.NewAddressSecpPubKey(input, dcrd_chaincfg.MainNetParams())
	dcrd_util.NewAddressEdwardsPubKey(input, dcrd_chaincfg.MainNetParams())
	dcrd_util.NewAddressSecSchnorrPubKey(input, dcrd_chaincfg.MainNetParams())
	return 0
}
