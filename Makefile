.PHONY: build test install release

test:
	go test .

install: 
	go install .

build: test
	go build .

release: test
	GOOS=darwin GOARCH=arm64  go build -o bin/medman-mac-arm64 .
	GOOS=darwin GOARCH=amd64  go build -o bin/medman-mac-amd64 .
	GOOS=linux GOARCH=arm64  go build -o bin/medman-linux-arm64 .
	GOOS=linux GOARCH=amd64  go build -o bin/medman-linux-amd64 .
	GOOS=windows GOARCH=amd64  go build -o bin/medman-linux-amd64 .

