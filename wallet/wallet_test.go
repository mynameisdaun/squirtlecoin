package wallet

import (
	"testing"
)

func makeTestWallet() *wallet {
	w := wallet{
		privateKey: nil,
		Address:    "",
	}
	return &w
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
