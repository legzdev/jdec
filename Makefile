.PHONY: build
build:
	go build -trimpath -ldflags="-s -w" -o ./build/jdec .
