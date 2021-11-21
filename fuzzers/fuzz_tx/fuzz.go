package fuzz_tx

import (
	dcrd_util "github.com/decred/dcrd/dcrutil/v4"
	dcrd_txsort "github.com/decred/dcrd/dcrutil/v4/txsort"
)

func Fuzz(input []byte) int {
	tx, err := dcrd_util.NewTxFromBytes(input)
	if err == nil {
		msgTx := tx.MsgTx()
		dcrd_util.NewTx(msgTx)
		dcrd_util.NewTxDeep(msgTx)
		dcrd_util.NewTxDeepTxIns(tx)
		dcrd_txsort.Sort(msgTx)
		dcrd_txsort.InPlaceSort(msgTx)
	}
	return 0
}

