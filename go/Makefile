.PHONY: requirements build

all: requirements pre-build build

requirements:
	@go mod download
	@go install github.com/golang/protobuf/protoc-gen-go

pre-build:
	mkdir -p wedge
	mkdir -p node
	mkdir -p slx 

build:
	protoc -I=../protobuf  --go_out=plugins=grpc,paths=source_relative:./slx slx.proto
	protoc -I=../protobuf  --go_out=plugins=grpc,paths=source_relative:./wedge wedge.proto
	protoc -I=../protobuf  --go_out=plugins=grpc,paths=source_relative:./node node.proto
