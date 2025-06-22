package main

import (
	"fmt"
	"os"
)

func main() {
	err := ParseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	decrypted, err := Decrypt(CryptedData, Key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Print(decrypted)
}
