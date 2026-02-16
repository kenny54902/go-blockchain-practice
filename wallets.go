package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

// Wallets stores a collection of wallets
type Wallets struct {
	Wallets map[string]*Wallet
}

// NewWallets creates Wallets and fills it from a file if it exists
func NewWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFromFile()

	return &wallets, err
}

// CreateWallet adds a Wallet to Wallets
func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := fmt.Sprintf("%s", wallet.GetAddress())

	ws.Wallets[address] = wallet

	return address
}

// GetAddresses returns an array of addresses stored in the wallet file
func (ws *Wallets) GetAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

// GetWallet returns a Wallet by its address
func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

// LoadFromFile loads wallets from the file
func (ws *Wallets) LoadFromFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := os.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	var walletsForGob map[string]WalletForGob

	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&walletsForGob)
	if err != nil {
		log.Panic(err)
	}

	// 把 WalletForGob 轉換成 Wallet（還原私鑰與公鑰）
	for addr, wfg := range walletsForGob {
		ws.Wallets[addr] = WalletFromGob(wfg)
	}

	return nil
}

// SaveToFile saves wallets to a file
func (ws Wallets) SaveToFile() {
	var content bytes.Buffer

	// type elliptic.nistCurve has no exported fields
	// gob.Register(elliptic.P256())

	// 自己實現序列化 WalletToGob
	//  ecdsa.PrivateKey 包含 PublicKey 和 D 欄位
	//  PublicKey 包含 elliptic.Curve 和 X, Y 欄位
	walletsForGob := make(map[string]WalletForGob)
	for _, wallet := range ws.Wallets {
		walletsForGob[string(wallet.GetAddress())] = WalletToGob(*wallet)
	}

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(walletsForGob)
	if err != nil {
		log.Panic(err)
	}

	err = os.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
