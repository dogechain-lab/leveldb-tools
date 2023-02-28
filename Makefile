.PHONY: bindir
birdir:
	mkdir -p bin

.PHONY: build
build: bindir
	go build -o bin/ldb main.go

.PHONY: linux
linux: bindir
	env GOOS=linux GOARCH=amd64 go build -o bin/ldb main.go
