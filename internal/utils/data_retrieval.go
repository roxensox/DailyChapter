package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func GetPrivateKey() *rsa.PrivateKey {
	// Gets the private key from its location

	// Reads the bytes of the file
	bytes, err := os.ReadFile("./env/secret_manager/dc_keys.pem")
	if err != nil {
		fmt.Printf("Failed to read private key: \n\t%v", err)
		return nil
	}

	// Decodes the bytes into a PEM block
	key, rest := pem.Decode(bytes)
	if rest != nil {
		fmt.Printf("%s\n", rest)
	}

	// Parses PEM block into a key object and returns it as *rsa.PrivateKey
	rsaKey, err := x509.ParsePKCS8PrivateKey(key.Bytes)
	if err != nil {
		fmt.Printf("Failed to parse key\n\t%v", err)
		return nil
	}
	return rsaKey.(*rsa.PrivateKey)
}

func GetPublicKey() *rsa.PublicKey {
	// Gets the public key from its location

	// Reads the bytes of the file
	bytes, err := os.ReadFile("./env/secret_manager/dc_keyspub.pem")
	if err != nil {
		fmt.Printf("Failed to read private key: \n\t%v", err)
		return nil
	}

	// Decodes the bytes into a PEM block
	key, rest := pem.Decode(bytes)
	if rest != nil {
		fmt.Printf("%s\n", rest)
	}

	// Parses PEM block into a key object and returns it as *rsa.PublicKey
	rsaKey, err := x509.ParsePKIXPublicKey(key.Bytes)
	if err != nil {
		fmt.Printf("Failed to parse key\n\t%v", err)
		return nil
	}
	return rsaKey.(*rsa.PublicKey)
}
