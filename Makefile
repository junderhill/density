BINARY_NAME=density

build:
	GOARCH=arm64 GOOS=darwin go build -o $(BINARY_NAME)-darwin-arm64 cmd/density/main.go
	GOARCH=amd64 GOOS=darwin go build -o $(BINARY_NAME)-darwin-amd64 cmd/density/main.go
	GOARCH=arm64 GOOS=linux go build -o $(BINARY_NAME)-linux-arm64 cmd/density/main.go
	GOARCH=amd64 GOOS=linux go build -o $(BINARY_NAME)-linux-amd64 cmd/density/main.go

clean:
	go clean
	rm -f $(BINARY_NAME)-darwin-arm64
	rm -f $(BINARY_NAME)-darwin-amd64
	rm -f $(BINARY_NAME)-linux-arm64
	rm -f $(BINARY_NAME)-linux-amd64
	rm -f coverage.out

test:
	 go test -race -v ./...

test_coverage:
	 go test ./... -coverprofile=coverage.out

run: build
	chmod +x $(BINARY_NAME)-darwin-arm64
	./$(BINARY_NAME)-darwin-arm64
