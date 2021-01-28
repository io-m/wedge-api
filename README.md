## Wappsto on the edge (Wedge) API

Communication is is based using (gRPC)[https://grpc.io/] remote RPC (Remote Procedure Call) framework and Protocol Buffers (aka protobuf)[https://developers.google.com/protocol-buffers/] library. It is a library for creating language neutral data structures, in our case 'protobuf' directory contain files 'node.proto' and 'wedge.proto',which can be converted to python or other language data types. It become as a contract between client and server, even they _speak_ different languages. 

## How to build

Prigeneration of python data types and mothods for RPC.

```bash
cd python; make
```

Python client/driver example: _'pyhton/src/driver.py'_
