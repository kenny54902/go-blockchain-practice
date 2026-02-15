package main

import "bytes"

// 一個交易輸入引用前一筆交易的一個輸出
// Txid: 前一筆交易的ID
// Vout: 前一筆交易的輸出的索引 (交易可能包含多個輸出)
// Signature:簽章的目的是證明花錢的人有權花這筆錢
// PubKey: 這筆輸入的人的公鑰 => 要花這筆幣的人的公鑰
type TXInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	PubKey    []byte
}

func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)
	return bytes.Compare(lockingHash, pubKeyHash) == 0
}
