package types

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var bigZero = big.NewInt(0)

func (tx *LegacyTx) isFake() bool     { return false }
func (tx *AccessListTx) isFake() bool { return false }
func (tx *DynamicFeeTx) isFake() bool { return false }

type ArbitrumUnsignedTx struct {
	ChainId *big.Int
	From    common.Address

	Nonce     uint64          // nonce of sender account
	GasFeeCap *big.Int        // wei per gas
	Gas       uint64          // gas limit
	To        *common.Address `rlp:"nil"` // nil means contract creation
	Value     *big.Int        // wei amount
	Data      []byte          // contract invocation input data
}

func (tx *ArbitrumUnsignedTx) txType() byte { return ArbitrumUnsignedTxType }

func (tx *ArbitrumUnsignedTx) copy() TxData {
	cpy := &ArbitrumUnsignedTx{
		ChainId:   new(big.Int),
		Nonce:     tx.Nonce,
		GasFeeCap: new(big.Int),
		Gas:       tx.Gas,
		From:      tx.From,
		To:        nil,
		Value:     new(big.Int),
		Data:      common.CopyBytes(tx.Data),
	}
	if tx.ChainId != nil {
		cpy.ChainId.Set(tx.ChainId)
	}
	if tx.GasFeeCap != nil {
		cpy.GasFeeCap.Set(tx.GasFeeCap)
	}
	if tx.To != nil {
		tmp := *tx.To
		cpy.To = &tmp
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	return cpy
}

func (tx *ArbitrumUnsignedTx) chainID() *big.Int      { return tx.ChainId }
func (tx *ArbitrumUnsignedTx) accessList() AccessList { return nil }
func (tx *ArbitrumUnsignedTx) data() []byte           { return tx.Data }
func (tx *ArbitrumUnsignedTx) gas() uint64            { return tx.Gas }
func (tx *ArbitrumUnsignedTx) gasPrice() *big.Int     { return tx.GasFeeCap }
func (tx *ArbitrumUnsignedTx) gasTipCap() *big.Int    { return bigZero }
func (tx *ArbitrumUnsignedTx) gasFeeCap() *big.Int    { return tx.GasFeeCap }
func (tx *ArbitrumUnsignedTx) value() *big.Int        { return tx.Value }
func (tx *ArbitrumUnsignedTx) nonce() uint64          { return tx.Nonce }
func (tx *ArbitrumUnsignedTx) to() *common.Address    { return tx.To }
func (tx *ArbitrumUnsignedTx) isFake() bool           { return false }

func (tx *ArbitrumUnsignedTx) rawSignatureValues() (v, r, s *big.Int) {
	return bigZero, bigZero, bigZero
}

func (tx *ArbitrumUnsignedTx) setSignatureValues(chainID, v, r, s *big.Int) {

}

type ArbitrumContractTx struct {
	ChainId   *big.Int
	RequestId common.Hash
	From      common.Address

	GasFeeCap *big.Int        // wei per gas
	Gas       uint64          // gas limit
	To        *common.Address `rlp:"nil"` // nil means contract creation
	Value     *big.Int        // wei amount
	Data      []byte          // contract invocation input data
}

func (tx *ArbitrumContractTx) txType() byte { return ArbitrumContractTxType }

func (tx *ArbitrumContractTx) copy() TxData {
	cpy := &ArbitrumContractTx{
		ChainId:   new(big.Int),
		RequestId: tx.RequestId,
		GasFeeCap: new(big.Int),
		Gas:       tx.Gas,
		From:      tx.From,
		To:        nil,
		Value:     new(big.Int),
		Data:      common.CopyBytes(tx.Data),
	}
	if tx.ChainId != nil {
		cpy.ChainId.Set(tx.ChainId)
	}
	if tx.GasFeeCap != nil {
		cpy.GasFeeCap.Set(tx.GasFeeCap)
	}
	if tx.To != nil {
		tmp := *tx.To
		cpy.To = &tmp
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	return cpy
}

func (tx *ArbitrumContractTx) chainID() *big.Int      { return tx.ChainId }
func (tx *ArbitrumContractTx) accessList() AccessList { return nil }
func (tx *ArbitrumContractTx) data() []byte           { return tx.Data }
func (tx *ArbitrumContractTx) gas() uint64            { return tx.Gas }
func (tx *ArbitrumContractTx) gasPrice() *big.Int     { return tx.GasFeeCap }
func (tx *ArbitrumContractTx) gasTipCap() *big.Int    { return bigZero }
func (tx *ArbitrumContractTx) gasFeeCap() *big.Int    { return tx.GasFeeCap }
func (tx *ArbitrumContractTx) value() *big.Int        { return tx.Value }
func (tx *ArbitrumContractTx) nonce() uint64          { return 0 }
func (tx *ArbitrumContractTx) to() *common.Address    { return tx.To }
func (tx *ArbitrumContractTx) rawSignatureValues() (v, r, s *big.Int) {
	return bigZero, bigZero, bigZero
}
func (tx *ArbitrumContractTx) setSignatureValues(chainID, v, r, s *big.Int) {}
func (tx *ArbitrumContractTx) isFake() bool                                 { return true }

type ArbitrumRetryTx struct {
	ChainId *big.Int
	Nonce   uint64
	From    common.Address

	GasFeeCap *big.Int        // wei per gas
	Gas       uint64          // gas limit
	To        *common.Address `rlp:"nil"` // nil means contract creation
	Value     *big.Int        // wei amount
	Data      []byte          // contract invocation input data
	TicketId  common.Hash
	RefundTo  common.Address
}

func (tx *ArbitrumRetryTx) txType() byte { return ArbitrumRetryTxType }

func (tx *ArbitrumRetryTx) copy() TxData {
	cpy := &ArbitrumRetryTx{
		ChainId:   new(big.Int),
		Nonce:     tx.Nonce,
		GasFeeCap: new(big.Int),
		Gas:       tx.Gas,
		From:      tx.From,
		To:        nil,
		Value:     new(big.Int),
		Data:      common.CopyBytes(tx.Data),
		TicketId:  tx.TicketId,
		RefundTo:  tx.RefundTo,
	}
	if tx.ChainId != nil {
		cpy.ChainId.Set(tx.ChainId)
	}
	if tx.GasFeeCap != nil {
		cpy.GasFeeCap.Set(tx.GasFeeCap)
	}
	if tx.To != nil {
		tmp := *tx.To
		cpy.To = &tmp
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	return cpy
}

func (tx *ArbitrumRetryTx) chainID() *big.Int      { return tx.ChainId }
func (tx *ArbitrumRetryTx) accessList() AccessList { return nil }
func (tx *ArbitrumRetryTx) data() []byte           { return tx.Data }
func (tx *ArbitrumRetryTx) gas() uint64            { return tx.Gas }
func (tx *ArbitrumRetryTx) gasPrice() *big.Int     { return tx.GasFeeCap }
func (tx *ArbitrumRetryTx) gasTipCap() *big.Int    { return bigZero }
func (tx *ArbitrumRetryTx) gasFeeCap() *big.Int    { return tx.GasFeeCap }
func (tx *ArbitrumRetryTx) value() *big.Int        { return tx.Value }
func (tx *ArbitrumRetryTx) nonce() uint64          { return tx.Nonce }
func (tx *ArbitrumRetryTx) to() *common.Address    { return tx.To }
func (tx *ArbitrumRetryTx) rawSignatureValues() (v, r, s *big.Int) {
	return bigZero, bigZero, bigZero
}
func (tx *ArbitrumRetryTx) setSignatureValues(chainID, v, r, s *big.Int) {}
func (tx *ArbitrumRetryTx) isFake() bool                                 { return true }

type ArbitrumSubmitRetryableTx struct {
	ChainId   *big.Int
	RequestId common.Hash
	From      common.Address

	DepositValue      *big.Int
	GasFeeCap         *big.Int        // wei per gas
	Gas               uint64          // gas limit
	To                *common.Address `rlp:"nil"` // nil means contract creation
	Value             *big.Int        // wei amount
	Beneficiary       common.Address
	SubmissionFeePaid *big.Int
	FeeRefundAddr     common.Address
	Data              []byte // contract invocation input data
}

func (tx *ArbitrumSubmitRetryableTx) txType() byte { return ArbitrumSubmitRetryableTxType }

func (tx *ArbitrumSubmitRetryableTx) copy() TxData {
	cpy := &ArbitrumSubmitRetryableTx{
		ChainId:           new(big.Int),
		RequestId:         tx.RequestId,
		DepositValue:      new(big.Int),
		GasFeeCap:         new(big.Int),
		Gas:               tx.Gas,
		From:              tx.From,
		To:                tx.To,
		Value:             new(big.Int),
		Beneficiary:       tx.Beneficiary,
		SubmissionFeePaid: new(big.Int),
		FeeRefundAddr:     tx.FeeRefundAddr,
		Data:              common.CopyBytes(tx.Data),
	}
	if tx.ChainId != nil {
		cpy.ChainId.Set(tx.ChainId)
	}
	if tx.DepositValue != nil {
		cpy.DepositValue.Set(tx.DepositValue)
	}
	if tx.GasFeeCap != nil {
		cpy.GasFeeCap.Set(tx.GasFeeCap)
	}
	if tx.To != nil {
		tmp := *tx.To
		cpy.To = &tmp
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.SubmissionFeePaid != nil {
		cpy.SubmissionFeePaid.Set(tx.SubmissionFeePaid)
	}
	return cpy
}

func (tx *ArbitrumSubmitRetryableTx) chainID() *big.Int      { return tx.ChainId }
func (tx *ArbitrumSubmitRetryableTx) accessList() AccessList { return nil }
func (tx *ArbitrumSubmitRetryableTx) data() []byte           { return tx.Data }
func (tx *ArbitrumSubmitRetryableTx) gas() uint64            { return tx.Gas }
func (tx *ArbitrumSubmitRetryableTx) gasPrice() *big.Int     { return tx.GasFeeCap }
func (tx *ArbitrumSubmitRetryableTx) gasTipCap() *big.Int    { return big.NewInt(0) }
func (tx *ArbitrumSubmitRetryableTx) gasFeeCap() *big.Int    { return tx.GasFeeCap }
func (tx *ArbitrumSubmitRetryableTx) value() *big.Int        { return tx.Value }
func (tx *ArbitrumSubmitRetryableTx) nonce() uint64          { return 0 }
func (tx *ArbitrumSubmitRetryableTx) to() *common.Address    { return tx.To }
func (tx *ArbitrumSubmitRetryableTx) rawSignatureValues() (v, r, s *big.Int) {
	return bigZero, bigZero, bigZero
}
func (tx *ArbitrumSubmitRetryableTx) setSignatureValues(chainID, v, r, s *big.Int) {}
func (tx *ArbitrumSubmitRetryableTx) isFake() bool                                 { return true }

type ArbitrumDepositTx struct {
	ChainId     *big.Int
	L1RequestId common.Hash
	To          common.Address
	Value       *big.Int
}

func (d *ArbitrumDepositTx) txType() byte {
	return ArbitrumDepositTxType
}

func (d *ArbitrumDepositTx) copy() TxData {
	tx := &ArbitrumDepositTx{
		ChainId: new(big.Int),
		To:      d.To,
		Value:   new(big.Int),
	}
	if d.ChainId != nil {
		tx.ChainId.Set(d.ChainId)
	}
	if d.Value != nil {
		tx.Value.Set(d.Value)
	}
	return tx
}

func (d *ArbitrumDepositTx) chainID() *big.Int      { return d.ChainId }
func (d *ArbitrumDepositTx) accessList() AccessList { return nil }
func (d *ArbitrumDepositTx) data() []byte           { return nil }
func (d *ArbitrumDepositTx) gas() uint64            { return 0 }
func (d *ArbitrumDepositTx) gasPrice() *big.Int     { return bigZero }
func (d *ArbitrumDepositTx) gasTipCap() *big.Int    { return bigZero }
func (d *ArbitrumDepositTx) gasFeeCap() *big.Int    { return bigZero }
func (d *ArbitrumDepositTx) value() *big.Int        { return d.Value }
func (d *ArbitrumDepositTx) nonce() uint64          { return 0 }
func (d *ArbitrumDepositTx) to() *common.Address    { return &d.To }
func (d *ArbitrumDepositTx) isFake() bool           { return true }

func (d *ArbitrumDepositTx) rawSignatureValues() (v, r, s *big.Int) {
	return bigZero, bigZero, bigZero
}

func (d *ArbitrumDepositTx) setSignatureValues(chainID, v, r, s *big.Int) {

}

type ArbitrumInternalTx struct {
	ChainId       *big.Int
	Type          uint8
	Data          []byte
	L2BlockNumber uint64
	TxIndex       uint64
}

func (t *ArbitrumInternalTx) txType() byte {
	return ArbitrumInternalTxType
}

func (t *ArbitrumInternalTx) copy() TxData {
	return &ArbitrumInternalTx{
		new(big.Int).Set(t.ChainId),
		t.Type,
		common.CopyBytes(t.Data),
		t.L2BlockNumber,
		t.TxIndex,
	}
}

func (t *ArbitrumInternalTx) chainID() *big.Int      { return t.ChainId }
func (t *ArbitrumInternalTx) accessList() AccessList { return nil }
func (t *ArbitrumInternalTx) data() []byte           { return t.Data }
func (t *ArbitrumInternalTx) gas() uint64            { return 0 }
func (t *ArbitrumInternalTx) gasPrice() *big.Int     { return bigZero }
func (t *ArbitrumInternalTx) gasTipCap() *big.Int    { return bigZero }
func (t *ArbitrumInternalTx) gasFeeCap() *big.Int    { return bigZero }
func (t *ArbitrumInternalTx) value() *big.Int        { return common.Big0 }
func (t *ArbitrumInternalTx) nonce() uint64          { return 0 }
func (t *ArbitrumInternalTx) to() *common.Address    { return &arbAddress }
func (t *ArbitrumInternalTx) isFake() bool           { return true }

func (d *ArbitrumInternalTx) rawSignatureValues() (v, r, s *big.Int) {
	return bigZero, bigZero, bigZero
}

func (d *ArbitrumInternalTx) setSignatureValues(chainID, v, r, s *big.Int) {

}

type HeaderInfo struct {
	SendRoot      common.Hash
	SendCount     uint64
	L1BlockNumber uint64
}

func (info HeaderInfo) extra() []byte {
	return info.SendRoot[:]
}

func (info HeaderInfo) mixDigest() [32]byte {
	mixDigest := common.Hash{}
	binary.BigEndian.PutUint64(mixDigest[:8], info.SendCount)
	binary.BigEndian.PutUint64(mixDigest[8:16], info.L1BlockNumber)
	return mixDigest
}

func (info HeaderInfo) UpdateHeaderWithInfo(header *Header) {
	header.MixDigest = info.mixDigest()
	header.Extra = info.extra()
}

func DeserializeHeaderExtraInformation(header *Header) (HeaderInfo, error) {
	if header.Number.Sign() == 0 || len(header.Extra) == 0 {
		// The genesis block doesn't have an ArbOS encoded extra field
		return HeaderInfo{}, nil
	}
	if len(header.Extra) != 32 {
		return HeaderInfo{}, fmt.Errorf("unexpected header extra field length %v", len(header.Extra))
	}
	extra := HeaderInfo{}
	copy(extra.SendRoot[:], header.Extra)
	extra.SendCount = binary.BigEndian.Uint64(header.MixDigest[:8])
	extra.L1BlockNumber = binary.BigEndian.Uint64(header.MixDigest[8:16])
	return extra, nil
}