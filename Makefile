GOPATH:=$(shell go env GOPATH)

WORK_PATH ?= f:
MODULE ?= user
.PHONY: proto
proto:
	protoc --proto_path=$(WORK_PATH)/ --go_out=$(WORK_PATH)/ $(WORK_PATH)/gitee.com/jingshanccc/course/$(MODULE)/proto/dto/*.proto
	protoc --proto_path=$(WORK_PATH)/ --go_out=$(WORK_PATH)/ --micro_out=$(WORK_PATH)/ $(WORK_PATH)/gitee.com/jingshanccc/course/$(MODULE)/proto/$(MODULE)/*.proto

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

.PHONY: api
api:
	micro --api_namespace=com.chan.course --auth_namespace=com.chan.course --registry=etcd --registry_address=192.168.10.130:2379 api