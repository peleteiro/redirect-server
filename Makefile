.DEFAULT_GOAL=build

PKG := $(GOPATH)/pkg

run: get
	SERVER_FQDN=redirect.farofus.com go run redirect-server.go

rerun: get
	SERVER_FQDN=redirect.farofus.com $(GOPATH)/bin/rerun github.com/peleteiro/redirect-server

get:
	go get golang.org/x/net/publicsuffix/...
	go get github.com/skelterjohn/rerun/...
	go get github.com/stretchr/testify/assert/...

build:
	@GOOS=darwin GOARCH=amd64 go build -o $(PKG)/darwin_amd64/redirect-server redirect-server.go
	@GOOS=linux GOARCH=amd64 go build -o $(PKG)/linux_amd64/redirect-server redirect-server.go
	@GOOS=linux GOARCH=386 go build -o $(PKG)/linux_386/redirect-server redirect-server.go

fmt:
	go fmt ./...

test: get
	go test ./...

docker-build:
	docker build -t "peleteiro/redirect-server" .

docker-bash:
	docker run -ti "peleteiro/redirect-server" bash

docker-run:
	docker run -e SERVER_FQDN=redirect.farofus.com -p 8080:8080 -t -i "peleteiro/redirect-server"

docker-start:
	docker run -e SERVER_FQDN=redirect.farofus.com -p 8080:8080 -d "peleteiro/redirect-server"

docker-push:
	docker push peleteiro/redirect-server:latest
