package main

import (
	"errors"
	"flag"
)

var (
	CryptedData []byte
	Key         []byte
	Pretty      bool
	URL         string
)

func ParseArgs() error {
	crypted := flag.String("crypted", "", "data as base64")
	key := flag.String("key", "", "decrypt key in hex")
	flag.BoolVar(&Pretty, "pretty", false, "pretty output")
	flag.Parse()

	if *crypted != "" && *key != "" {
		CryptedData = []byte(*crypted)
		Key = []byte(*key)
		return nil
	}

	args := flag.Args()
	if len(args) != 1 {
		return errors.New("expected 1 argument")
	}

	content, err := GetContent(args[0])
	if err != nil {
		return err
	}

	CryptedData, err = GetCryptedData(content)
	if err != nil {
		return err
	}

	Key, err = GetKey(content)
	if err != nil {
		return err
	}

	return nil
}
