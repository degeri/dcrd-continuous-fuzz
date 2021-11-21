package fuzz_txscript

import (
	dcrd_chainhash "github.com/decred/dcrd/chaincfg/chainhash"
	dcrd_chaincfg "github.com/decred/dcrd/chaincfg/v3"
	dcrd_txscript "github.com/decred/dcrd/txscript/v4"
	dcrd_wire "github.com/decred/dcrd/wire"
	dcrd_stdaddr "github.com/decred/dcrd/txscript/v4/stdaddr"
)

func dcrd_DisasmString(input []byte) {
	dcrd_txscript.DisasmString(input)
}

func dcrd_VmStep(input []byte) {
	tx := &dcrd_wire.MsgTx{
		Version: 1,
		TxIn: []*dcrd_wire.TxIn{
			{
				PreviousOutPoint: dcrd_wire.OutPoint{
					Hash: dcrd_chainhash.Hash([32]byte{
						0xc9, 0x97, 0xa5, 0xe5,
						0x6e, 0x10, 0x41, 0x02,
						0xfa, 0x20, 0x9c, 0x6a,
						0x85, 0x2d, 0xd9, 0x06,
						0x60, 0xa2, 0x0b, 0x2d,
						0x9c, 0x35, 0x24, 0x23,
						0xed, 0xce, 0x25, 0x85,
						0x7f, 0xcd, 0x37, 0x04,
					}),
					Index: 0,
				},
				SignatureScript: nil,
				Sequence:        4294967295,
			},
		},
		TxOut: []*dcrd_wire.TxOut{{
			Value:    1000000000,
			PkScript: nil,
		}},
		LockTime: 0,
	}
	vm, err := dcrd_txscript.NewEngine(input, tx, 0, 0, 0, nil)
	if err != nil {
		return
	}
	for i := 0; i < len(input); i++ {
		done, err := vm.Step()
		if err != nil {
			break
		}
		if done {
			break
		}
	}
}

func Fuzz(input []byte) int {
	dcrd_DisasmString(input)
	dcrd_VmStep(input)
	dcrd_txscript.ContainsStakeOpCodes(input, true)

	{
		builder := dcrd_txscript.NewScriptBuilder()
		builder.Reset()
		builder.AddOps(input)
		builder.AddData(input)
		builder.Script()
	}
	{
		dcrd_stdaddr.DecodeAddress(string(input), dcrd_chaincfg.MainNetParams())
	}
	return 0
}
