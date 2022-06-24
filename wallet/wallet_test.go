package wallet

import (
	"crypto/ecdsa"
	"math/big"
	"reflect"
	"testing"
)

func TestSign(t *testing.T) {
	type args struct {
		payload string
		w       *wallet
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sign(tt.args.payload, tt.args.w); got != tt.want {
				t.Errorf("Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	type args struct {
		signature string
		payload   string
		address   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Verify(tt.args.signature, tt.args.payload, tt.args.address); got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWallet(t *testing.T) {
	tests := []struct {
		name string
		want *wallet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wallet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wallet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_aFromK(t *testing.T) {
	type args struct {
		key *ecdsa.PrivateKey
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aFromK(tt.args.key); got != tt.want {
				t.Errorf("aFromK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createPrivateKey(t *testing.T) {
	tests := []struct {
		name string
		want *ecdsa.PrivateKey
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createPrivateKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createPrivateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeBigInts(t *testing.T) {
	type args struct {
		a []byte
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBigInts(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("encodeBigInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasWalletFile(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasWalletFile(); got != tt.want {
				t.Errorf("hasWalletFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_persistKey(t *testing.T) {
	type args struct {
		key *ecdsa.PrivateKey
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			persistKey(tt.args.key)
		})
	}
}

func Test_restoreBigInts(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		want1   *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := restoreBigInts(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("restoreBigInts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreBigInts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("restoreBigInts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_restoreKey(t *testing.T) {
	tests := []struct {
		name string
		want *ecdsa.PrivateKey
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
