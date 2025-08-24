package seed2eth

import (
	"testing"
)

func TestGetAddress(t *testing.T) {
	testCases := []struct {
		name     string
		mnemonic string
		password string
		childIdx uint32
		expected string
	}{
		{
			name:     "default values",
			mnemonic: "",
			password: "",
			childIdx: 0,
			expected: "0x959FD7Ef9089B7142B6B908Dc3A8af7Aa8ff0FA1",
		},
		{
			name:     "with password, childIdx = 0",
			mnemonic: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about",
			password: "123",
			childIdx: 0,
			expected: "0x76A0710e599eB526EB101bD258A472aff740D5a4",
		},
		{
			name:     "with password, childIdx = 1",
			mnemonic: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about",
			password: "asdf",
			childIdx: 1,
			expected: "0x4Fe2E171A03E6749D4623A05961a509D5D361893",
		},
		{
			name:     "without password, childIdx = 0",
			mnemonic: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about",
			password: "",
			childIdx: 0,
			expected: "0x9858EfFD232B4033E47d90003D41EC34EcaEda94",
		},
		{
			name:     "without password, childIdx = 1",
			mnemonic: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about",
			password: "",
			childIdx: 1,
			expected: "0x6Fac4D18c912343BF86fa7049364Dd4E424Ab9C0",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			addr, err := GetAddress(tc.mnemonic, tc.password, tc.childIdx)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if addr != tc.expected {
				t.Errorf("expected address %s, got %s", tc.expected, addr)
			}
		})
	}
}

func TestAddressFromPrivateKey(t *testing.T) {
	tests := []struct {
		privateKey string
		expected   string
	}{
		{
			privateKey: "abf82ff96b463e9d82b83cb9bb450fe87e6166d4db6d7021d0c71d7e960d5abe",
			expected:   "0x959FD7Ef9089B7142B6B908Dc3A8af7Aa8ff0FA1",
		},
		{
			privateKey: "0bb9160621e50bb66f2c14c83d666cc85297794f9469841a0afee7c6d2c8c5eb",
			expected:   "0x76A0710e599eB526EB101bD258A472aff740D5a4",
		},
	}

	for _, tc := range tests {
		addr, err := GetAddressFromPrivateKey(tc.privateKey)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if addr != tc.expected {
			t.Errorf("expected address %s, got %s", tc.expected, addr)
		}
	}
}
