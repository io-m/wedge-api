PACKAGE_NAME := import \"wedge-api/

all: requirements pre-build build

requirements:
	python3 -m pip install --upgrade pip
	python3 -m pip install --upgrade protobuf
	python3 -m pip install grpcio-tools

pre-build:
	mkdir -p src
	rm -rf proto
	mkdir -p proto/wedge-api
	cp -r ../protobuf/* proto/wedge-api/

build:
	python3 -m grpc_tools.protoc  --python_out=src/wedge_api  --grpc_python_out=./src/wedge_api -I=proto/wedge-api proto/wedge-api/slx.proto proto/wedge-api/wedge.proto
