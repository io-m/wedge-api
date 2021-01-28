## Wappsto an the edge (Wedge) API

Communication is is based using gRPC remote RPC (Remote Procedure Call) framework and Protocol Buffers (aka protobuf) library. Language neutral data structues, in our case 'protobuf' directory contain files 'node.proto' and 'wedge.proto', is converted to python or other language data type. It become as a contract between client and server, even they _speak_ different languages. 

## How to build

Prigeneration of python data types and mothods for RPC.

```bash
cd python; make
```

Take a look at example for python client/driver file: 'pyhton/src/driver.py'
