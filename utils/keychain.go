package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
)

func LoadPrivateKeyFromString(privateKeyPEM string) (*rsa.PrivateKey, error) {
	// Decode the PEM block
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil || (block.Type != "RSA PRIVATE KEY" && block.Type != "PRIVATE KEY") {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	// Parse the private key
	var parsedKey interface{}
	var err error
	if block.Type == "RSA PRIVATE KEY" {
		parsedKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	} else {
		parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	privateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("not an RSA private key")
	}

	return privateKey, nil
}

func LoadPublicKeyFromString(publicKeyPEM string) (*rsa.PublicKey, error) {
	// Decode the PEM block
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil || (block.Type != "PUBLIC KEY" && block.Type != "RSA PUBLIC KEY") {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	// Parse the public key
	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}

	return publicKey, nil
}

// EncryptWithPublicKey encrypts data using the RSA public key
func EncryptWithPublicKey(message string, pub *rsa.PublicKey) string {
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(message))
	if err != nil {
		log.Fatalf("failed to encrypt message: %v", err)
	}
	// Convert encrypted data to base64 for easier handling
	return base64.StdEncoding.EncodeToString(encryptedBytes)
}

// DecryptWithPrivateKey decrypts data using the RSA private key
func DecryptWithPrivateKey(cipherText string, priv *rsa.PrivateKey) string {
	// Decode the base64 encrypted message
	encryptedBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		log.Fatalf("failed to decode base64 string: %v", err)
	}
	// Decrypt the data
	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, priv, encryptedBytes)
	if err != nil {
		log.Fatalf("failed to decrypt message: %v", err)
	}
	return string(decryptedBytes)
}
