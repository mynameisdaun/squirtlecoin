package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/mynameisdaun/squirtlecoin/utils"
	"math/big"
	"os"
	"sync"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
	Address    string
}

var w *wallet
var once sync.Once

const (
	fileName string = "squirtlecoin.wallet"
)

func hasWalletFile() bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func aFromK(key *ecdsa.PrivateKey) string {
	return encodeBigInts(key.X.Bytes(), key.Y.Bytes())
}

func createPrivateKey() *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return privateKey
}

func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleErr(err)
	err = os.WriteFile(fileName, bytes, 0644)
	utils.HandleErr(err)
}

func restoreKey() *ecdsa.PrivateKey {
	keyAsByte, err := os.ReadFile(fileName)
	utils.HandleErr(err)
	key, err := x509.ParseECPrivateKey(keyAsByte)
	utils.HandleErr(err)
	return key
}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			w.privateKey = restoreKey()
		} else {
			key := createPrivateKey()
			persistKey(key)
			w.privateKey = key
		}
		w.Address = aFromK(w.privateKey)
	}
	return w
}

func encodeBigInts(a, b []byte) string {
	z := append(a, b...)
	return fmt.Sprintf("%x", z)
}

//sign
func Sign(payload string, w *wallet) string {
	payloadAsByte, err := hex.DecodeString(payload)
	utils.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, payloadAsByte)
	utils.HandleErr(err)
	return encodeBigInts(r.Bytes(), s.Bytes())
}

//verify
func restoreBigInts(str string) (*big.Int, *big.Int, error) {
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return nil, nil, err
	}
	firstHalf := bytes[0 : len(bytes)/2]
	secondHalf := bytes[len(bytes)/2:]
	bigA, bigB := big.Int{}, big.Int{}
	bigA.SetBytes(firstHalf)
	bigB.SetBytes(secondHalf)
	return &bigA, &bigB, nil
}

func Verify(signature, payload, address string) bool {
	r, s, err := restoreBigInts(signature)
	utils.HandleErr(err)
	x, y, err := restoreBigInts(address)
	utils.HandleErr(err)
	publicKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}
	payloadBytes, err := hex.DecodeString(payload)
	utils.HandleErr(err)
	ok := ecdsa.Verify(&publicKey, payloadBytes, r, s)
	return ok

}
