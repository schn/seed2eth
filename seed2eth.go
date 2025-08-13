package seed2eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

// GetAddress returns eth address string from mnemonic, password and childIdx
func GetAddress(mnemonic, password string, childIdx uint32) (string, error) {
	seed := bip39.NewSeed(mnemonic, password)

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return "", fmt.Errorf("failed to create master key: %v", err)
	}

	purpose, err := masterKey.NewChildKey(bip32.FirstHardenedChild + 44)
	if err != nil {
		return "", fmt.Errorf("failed to derive purpose: %v", err)
	}

	coinType, err := purpose.NewChildKey(bip32.FirstHardenedChild + 60)
	if err != nil {
		return "", fmt.Errorf("failed to derive coin type: %v", err)
	}

	account, err := coinType.NewChildKey(bip32.FirstHardenedChild + 0)
	if err != nil {
		return "", fmt.Errorf("failed to derive account: %v", err)
	}

	change, err := account.NewChildKey(0)
	if err != nil {
		return "", fmt.Errorf("failed to derive change: %v", err)
	}

	child, err := change.NewChildKey(childIdx)
	if err != nil {
		return "", fmt.Errorf("failed to derive address index: %v", err)
	}

	privateKey, err := crypto.ToECDSA(child.Key)
	if err != nil {
		return "", fmt.Errorf("failed to convert to ECDSA: %v", err)
	}

	return crypto.PubkeyToAddress(privateKey.PublicKey).Hex(), nil
}
