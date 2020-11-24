package alipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// parsePKCS1PrivateKey
func parsePKCS1PrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("private key format incorrect")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// parsePKCS8PrivateKey
func parsePKCS8PrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("private key with incorrect format")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	priv, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key with incorrect format")
	}

	return priv, nil
}

// signPKCS1v15
func signPKCS1v15(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
}

// verifySign
func verifySign() {
}
