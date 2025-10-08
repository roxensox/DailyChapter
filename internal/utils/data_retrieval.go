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
	bytes, err := os.ReadFile("./env/secret_manager/dc_keys")
	if err != nil {
		fmt.Printf("Failed to read private key: \n\t%v", err)
		return nil
	}

	// Decodes the bytes into a PEM block
	key, rest := pem.Decode(bytes)
	if rest != nil {
		fmt.Printf("%s\n", rest)
	}

	// Parses PEM block into a private key object and returns it
	rsaKey, err := x509.ParsePKCS1PrivateKey(key.Bytes)
	if err != nil {
		fmt.Printf("Failed to parse key\n\t%v", err)
		return nil
	}
	return rsaKey
}
