package main

import (
	"errors"
	"io"
	"net/http"
	"regexp"
)

func GetContent(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func GetCryptedData(b []byte) ([]byte, error) {
	regex := regexp.MustCompile(`name="crypted" value="(.+?)"`)

	result := regex.FindSubmatch(b)
	if len(result) != 2 {
		return nil, errors.New("cannot find crypted data")
	}

	return result[1], nil
}

func GetKey(b []byte) ([]byte, error) {
	regex := regexp.MustCompile(`return '(.+?)';}">`)

	result := regex.FindSubmatch(b)
	if len(result) != 2 {
		return nil, errors.New("cannot find key")
	}

	return []byte(result[1]), nil
}
