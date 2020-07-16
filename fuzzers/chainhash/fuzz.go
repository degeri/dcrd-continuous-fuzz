package chainhash

import (
	dcrd_chainhash "github.com/decred/dcrd/chaincfg/chainhash"
)

func Fuzz(input []byte) int {
	dcrd_chainhash.NewHashFromStr(string(input))
	return 0
}
