package crypto

import (
	"encoding/base64"

	"github.com/stellar/go-stellar-base/keypair"
)

type SignatureSignerVerifierInterface interface {
	Sign(secretSeed string, message []byte) (string, error)
	Verify(publicKey string, message, signature []byte) error
}

type SignatureSignerVerifier struct{}

// Signs message using secretSeed. Returns base64-encoded signature.
func (s *SignatureSignerVerifier) Sign(secretSeed string, message []byte) (string, error) {
	kp, err := keypair.Parse(secretSeed)
	if err != nil {
		return "", err
	}

	signature, err := kp.Sign(message)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// Verifies if signature is a valid signature of message signed by publicKey.
func (s *SignatureSignerVerifier) Verify(publicKey string, message, signature []byte) error {
	kp, err := keypair.Parse(publicKey)
	if err != nil {
		return err
	}

	err = kp.Verify(message, signature)
	if err != nil {
		return err
	}

	return nil
}
