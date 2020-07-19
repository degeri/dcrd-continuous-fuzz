package fuzz_bech32

import (
	"github.com/decred/dcrd/bech32"
)

func Fuzz(input []byte) int {
	bech32.Decode(string(input))
	conv, _ := bech32.ConvertBits(input, 8, 5, true)
	bech32.Encode(string(input), conv)
	return 0
}
