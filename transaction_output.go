package main

import "bytes"

// 交易輸出為實際幣儲存的位置
// Value: 幣
// PubKeyHash: 輸出接受者的公鑰的雜湊值 => 持有對應私鑰的人有控制權
type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

// 鎖定輸出
func (out *TXOutput) Lock(address []byte) {
	// 從地址中提取公鑰的雜湊值
	pubKeyHash := Base58Decode(address)
	// 版本號.payload.checksum => 去掉版本號和checksum
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

// 檢查控制權是否匹配
func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// 建立新的交易輸出
func NewTXOutput(value int, address string) *TXOutput {
	// 建立輸出
	txo := &TXOutput{value, nil}
	// 上鎖
	txo.Lock([]byte(address))
	return txo
}
