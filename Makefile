.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf tus
	
build:
	GOOS=linux GOARCH=amd64 go build -o tus