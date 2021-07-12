GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
APP_RELATIVE_PATH=comment/service
API_PROTO_FILES=$(shell cd ../../../api/$(APP_RELATIVE_PATH) && find . -name "*.proto")
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")

.PHONY: init
# init env
init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

.PHONY: grpc
# generate grpc code
grpc: 
	 cd ../../../api/$(APP_RELATIVE_PATH) && protoc --proto_path=. \
	       --proto_path=../../../third_party \
           --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
           $(API_PROTO_FILES)

.PHONY: proto
# generate internal proto struct
proto:
	protoc --proto_path=. \
           --proto_path=../../../third_party \
           --go_out=paths=source_relative:. \
           $(INTERNAL_PROTO_FILES)

.PHONY: do
do:
	echo $(VERSION)
	echo $(APP_RELATIVE_PATH) 
	echo $(API_PROTO_FILES) 
	echo $(APP_NAME) 


.PHONY: build-service
build-service:
	export GO111MODULE=on && \
	export GOPROXY=https://goproxy.cn && \
	export GOPRIVATE=git.aigauss.com && \
	go build -o $(APP_NAME) cmd/main.go

.PHONY: run-service
run-service:
	./$(SERVER_NAME)