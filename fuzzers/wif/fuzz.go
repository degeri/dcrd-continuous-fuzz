package wif

import (
	dcrd_chaincfg "github.com/decred/dcrd/chaincfg/v3"
	dcrd_util "github.com/decred/dcrd/dcrutil/v3"
)

func Fuzz(input []byte) int {
	wif, err := dcrd_util.DecodeWIF(string(input), dcrd_chaincfg.MainNetParams().PrivateKeyID)
	if err == nil {
		wif.String()
		wif.PubKey()
	}
	return 0
}
