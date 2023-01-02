.PHONY: clean lint build run list test testv 

clean:
	rm -f bin/*

lint:
	gofmt -l -s -w .

build:
	go build -o ./bin/main.exe cmd/server/main.go

run:
	go run cmd/server/main.go

list:
	go list ./...

test:
	go test `go list ./... | grep -v ./e2e`

testv:
	go test -v `go list ./... | grep -v ./e2e`

