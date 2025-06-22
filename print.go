package main

import (
	"fmt"
	"os"
	"regexp"
)

func Print(b []byte) {
	if !Pretty {
		os.Stdout.Write(b)
		return
	}

	regex := regexp.MustCompile(`\[url=(.+?)\]`)

	results := regex.FindAllSubmatch(b, -1)
	if results == nil {
		os.Stdout.Write(b)
		return
	}

	for _, result := range results {
		fmt.Println(string(result[1]))
	}
}
