GOPATH:=$(shell go env GOPATH)

.PHONY: proto
MODULE = user
proto:
	protoc --proto_path=../ --go_out=../ ../course/$(MODULE)-srv/proto/dto/*.proto
	protoc --proto_path=../ --go_out=../ --micro_out=../ ../course/$(MODULE)-srv/proto/$(MODULE)/*.proto

.PHONY: build
build:
	cd $(MODULE)-srv
	set GOARCH=amd64
	set GOOS=linux
	go build -o $(MODULE) main.go

.PHONY: docker
docker:
	cd $(MODULE)-srv
	docker build . -t $(MODULE):latest