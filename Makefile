.DEFAULT_GOAL=build

PKG := $(GOPATH)/pkg

run: get
	go run redirect-server.go

rerun: get
	$(GOPATH)/bin/rerun github.com/peleteiro/redirect-server

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

docker\:build:
	docker build -t "peleteiro/redirect-server" .

docker\:bash:
	docker run -ti "peleteiro/redirect-server" bash

docker\:run:
	docker run -p 8080:8080 -t -i "peleteiro/redirect-server"

docker\:start:
	docker run -p 8080:8080 -d "peleteiro/redirect-server"

docker\:push:
	docker push peleteiro/redirect-server:latest

ssh-keygen:
	ssh-keygen -t rsa -b 4096 -C "redirect-server" -f terraform/certs/ssh-key

tf:
	cd terraform && \
	terraform plan -var-file ../terraform.tfvars -out terraform.tfplan && \
	terraform apply -var-file ../terraform.tfvars

tf\:destroy:
	cd terraform && terraform plan -destroy -var-file ../terraform.tfvars -out terraform.tfplan
	cd terraform && terraform apply terraform.tfplan
