package base58

import (
	dcrd_base58 "github.com/decred/base58"
)

func Fuzz(input []byte) int {
	dcrd_base58.Decode(string(input))
	dcrd_base58.Encode(input)
	return 0
}
