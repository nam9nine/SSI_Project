package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/multiformats/go-multibase"
	"log"
)

type EcdsaKeyStorage struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

func (e *EcdsaKeyStorage) GenerateKey() (string, error) {
	if e == nil {
		return "", fmt.Errorf("EcdsaKeyStorage instance is nil")
	}

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", fmt.Errorf("error generating private key: %v", err)
	}

	e.PrivateKey = privateKey
	e.PublicKey = &privateKey.PublicKey

	return e.PublicKeyMultibase(), nil
}

func (e *EcdsaKeyStorage) PublicKeyToString() (string, error) {
	if e == nil {
		return "", fmt.Errorf("EcdsaKeyStorage instance is nil")
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(e.PublicKey)
	if err != nil {
		log.Printf("error occured: %v", err.Error())
		return "", err
	}

	publicKeyHash := sha256.Sum256(publicKeyBytes)

	return hex.EncodeToString(publicKeyHash[:]), nil
}

func (e *EcdsaKeyStorage) PublicKeyMultibase() string {

	if e.PublicKey == nil {
		return ""
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(e.PublicKey)

	if err != nil {
		log.Printf("error occured: %v", err.Error())
		return ""
	}

	str, err := multibase.Encode(multibase.Base58BTC, publicKeyBytes)

	if err != nil {
		log.Printf("error occured: %v", err.Error())
		return ""
	}
	return str
}
