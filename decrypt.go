package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Decrypt(crypted []byte, key []byte) ([]byte, error) {
	n, err := base64.StdEncoding.Decode(crypted, crypted)
	if err != nil {
		return nil, fmt.Errorf("error decoding crypted data: %v", err)
	}

	crypted = crypted[:n]

	n, err = hex.Decode(key, key)
	if err != nil {
		return nil, fmt.Errorf("error decoding key: %v", err)
	}

	key = key[:n]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cbc := cipher.NewCBCDecrypter(block, key)
	cbc.CryptBlocks(crypted, crypted)

	return crypted, nil
}
