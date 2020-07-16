package txscript

import (
	dcrd_chainhash "github.com/decred/dcrd/chaincfg/chainhash"
	dcrd_chaincfg "github.com/decred/dcrd/chaincfg/v3"
	dcrd_util "github.com/decred/dcrd/dcrutil/v3"
	dcrd_txscript "github.com/decred/dcrd/txscript/v3"
	dcrd_wire "github.com/decred/dcrd/wire"
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

func dcrd_ExtractPKScriptAddrs(input []byte) {
	const scriptVersion = 0
	dcrd_txscript.ExtractPkScriptAddrs(scriptVersion, input, dcrd_chaincfg.MainNetParams())
}

func Fuzz(input []byte) int {
	if dcrd_txscript.IsMultisigSigScript(input) {
		dcrd_txscript.MultisigRedeemScriptFromScriptSig(input)
		/* Crashes 31-08-2018 dcrd_txscript.GetMultisigMandN(input) */
	}
	isMultiSig := dcrd_txscript.IsMultisigScript(input)
	if isMultiSig {
		dcrd_txscript.CalcMultiSigStats(input)
	}
	dcrd_txscript.IsMultisigSigScript(input)
	const scriptVersion = 0
	dcrd_txscript.GetScriptClass(scriptVersion, input)
	dcrd_DisasmString(input)
	dcrd_VmStep(input)
	/* Crashed (30-08-2018), confirmed fixed 05-09-2019 */ dcrd_ExtractPKScriptAddrs(input)
	dcrd_txscript.PushedData(input)
	dcrd_txscript.ExtractPkScriptAltSigType(input)
	dcrd_txscript.GenerateProvablyPruneableOut(input)
	dcrd_txscript.GetStakeOutSubclass(input)
	dcrd_txscript.ContainsStakeOpCodes(input)
	dcrd_txscript.PayToScriptHashScript(input)

	{
		builder := dcrd_txscript.NewScriptBuilder()
		builder.Reset()
		builder.AddOps(input)
		builder.AddData(input)
		builder.Script()
	}
	{
		pos := 0
		left := len(input)
		var keys []*dcrd_util.AddressSecpPubKey
		for {
			if left < 1 {
				break
			}
			keylen := 0
			if input[pos]&1 == 0 {
				keylen = 33
			} else {
				keylen = 65
			}
			pos += 1
			left -= 1
			if left < keylen {
				break
			}
			key := input[pos : pos+keylen]
			pos += keylen
			left -= keylen
			apk, err := dcrd_util.NewAddressSecpPubKey(key, dcrd_chaincfg.MainNetParams())
			if err != nil {
				break
			}
			keys = append(keys, apk)
		}
		script, err := dcrd_txscript.MultiSigScript(keys, len(keys))
		if err == nil {
			dcrd_txscript.MultisigRedeemScriptFromScriptSig(script)
		}
	}
	{
		addr, err := dcrd_util.DecodeAddress(string(input), dcrd_chaincfg.MainNetParams())
		if err == nil {
			dcrd_txscript.PayToAddrScript(addr)
		}
	}
	return 0
}
