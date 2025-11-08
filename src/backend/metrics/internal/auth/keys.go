package auth

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// KeyPair holds one active signing keypair and its KID.
// - Private: used to SIGN tokens
// - Public:  shared to VERIFY tokens (exposed via JWKS).
// - KID:  placed into JWT header so verifiers pick the right public key.
type KeyPair struct {
	Private *rsa.PrivateKey
	Public  *rsa.PublicKey
	KID     string // e.g. "2025-11-07"
}

// LoadKeyPairFromPEM loads an RSA private key from a PEM file on disk and constructs the pair.
// path is the location of the private key, kid is an identifier for the key pair.
func LoadKeyPairFromPEM(path, kid string) (*KeyPair, error) {
	pem, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	private, err := jwt.ParseRSAPrivateKeyFromPEM(pem)
	if err != nil {
		return nil, err
	}
	return &KeyPair{Private: private, Public: &private.PublicKey, KID: kid}, nil
}
