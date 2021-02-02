.PHONY: requirements build

all: requirements pre-build build

requirements:
	@go mod download
	@go install github.com/golang/protobuf/protoc-gen-go

pre-build:
	mkdir -p wedgepb
	mkdir -p nodepb

build:
	protoc -I=../protobuf/wedge  --go_out=plugins=grpc,paths=source_relative:./wedgepb wedge.proto
	protoc -I=../protobuf/node  --go_out=plugins=grpc,paths=source_relative:./nodepb node.proto
