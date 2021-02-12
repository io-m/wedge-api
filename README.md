## Wappsto on the edge (Wedge) API


Communication is is based using (gRPC)[https://grpc.io/] remote RPC (Remote Procedure Call) framework and Protocol Buffers (aka protobuf)[https://developers.google.com/protocol-buffers/] library. It is a library for creating language neutral data structures. 'protobuf' directory contain files 'slx.proto' and 'wedge.proto' can be converted to python or other language data types. It become as a contract between client and server, even they _speak_ different languages.
Seluxit GW (Wedge) can have clients written in following languages:
* Go
* C++
* Java
* Python
* C#
* Node
* Ruby
* Kotlin/JVM
* Objective-C

At the moment there are pre-generated Python and Go types.

## How to build

For python:
```bash
cd python; make
```

For Go:
```
cd go; make
```

### Examples

Python client/node example: 
_'pyhton/src/wedge-api/node_template.py'_

Dummy server for testing purpose: 
_'pyhton/src/wedge-api/dummy_server.py'_
