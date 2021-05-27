all: build-safepay

build-safepay:
	mkdir -p bin
	go build -o bin/safepay cmd/safepay/main.go
	
