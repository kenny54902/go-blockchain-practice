package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

const version = byte(0x00)
const walletFile = "wallet.dat"
const addressChecksumLen = 4

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

type WalletForGob struct {
	EcdsaPrivateKey []byte
	EcdsaPublicKey  []byte
	PublicKey       []byte
}

// 建立新錢包
func NewWallet() *Wallet {
	private, public := newKeyPair()
	wallet := Wallet{private, public}
	return &wallet
}

// WalletFromGob 從 WalletForGob 還原出 Wallet（還原私鑰 D 與公鑰 X,Y）
func WalletFromGob(wfg WalletForGob) *Wallet {
	curve := elliptic.P256()
	// 還原Ecdsa私鑰 D（儲存時為 PrivateKey.D.Bytes()）
	D := new(big.Int).SetBytes(wfg.EcdsaPrivateKey)
	// 還原Ecdsa公鑰 X, Y（儲存時為 append(X.Bytes(), Y.Bytes()...)，P256 各 32 bytes）
	pubLen := len(wfg.EcdsaPublicKey) / 2
	X := new(big.Int).SetBytes(wfg.EcdsaPublicKey[:pubLen])
	Y := new(big.Int).SetBytes(wfg.EcdsaPublicKey[pubLen:])
	privKey := ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{Curve: curve, X: X, Y: Y},
		D:         D,
	}
	return &Wallet{PrivateKey: privKey, PublicKey: wfg.PublicKey}
}

func WalletToGob(w Wallet) WalletForGob {
	return WalletForGob{
		EcdsaPrivateKey: w.PrivateKey.D.Bytes(),
		EcdsaPublicKey:  append(w.PrivateKey.PublicKey.X.Bytes(), w.PrivateKey.PublicKey.Y.Bytes()...),
		PublicKey:       w.PublicKey,
	}
}

func (w Wallet) GetAddress() []byte {
	// 1. 生成公鑰的雜湊值
	pubKeyHash := HashPubKey(w.PublicKey)
	// 2. 添加版本號
	versionedPayload := append([]byte{version}, pubKeyHash...)
	// 3. 再計算一次雜湊值SHA256(SHA256(payload))作為checksum
	// 接收方將資料做相同的操作，如果checksum相同，則資料沒有被篡改
	checksum := checksum(versionedPayload)

	// 4. 將checksum附加到payload後面
	fullPayload := append(versionedPayload, checksum...)
	// 5. 使用Base58編碼生成地址
	address := Base58Encode(fullPayload)

	return address
}

// 使用SHA256和RIPEMD160雜湊函數生成公鑰的雜湊值
func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

// 生成公鑰和私鑰
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	// 使用橢圓曲線生成私鑰
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pubKey
}

func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}

func ValidateAddress(address string) bool {
	pubKeyHash := Base58Decode([]byte(address))
	actualChecksum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))
	return bytes.Equal(actualChecksum, targetChecksum)
}
