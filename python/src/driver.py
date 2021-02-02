#!/usr/bin/python3

from grpc import aio
# for creating protobuf struct from python dictionary.
# from google.protobuf.struct_pb2 import Struct
# for conversion protobuf to json.
# from google.protobuf import json_format
import asyncio
import logging
import sys
import signal
import grpc
# pregenerated from proto file
from wedge_api import node_pb2
from wedge_api import node_pb2_grpc

from wedge_api import wedge_pb2_grpc
from wedge_api import wedge_pb2


def signal_handler(signal, frame):
    sys.exit(0)


"""
Uplink class holds all remote Wedge Methods
"""

# alias for state type: 'Control' or 'Report'
# control_state = wedge_pb2.StateType.Control
# report_state = wedge_pb2.StateType.Report

myIdentity = wedge_pb2.Driver(
    host="127.0.0.1",
    port=30052,
)

# Example of Model, which identical to Seluxit data Model.
model = wedge_pb2.Model(
    driver=myIdentity,
    device=[wedge_pb2.Device(
        id=1,
        name="water_control",
        version="0.1.2",
        meta=wedge_pb2.Meta(
            id="e2d4a058-4a56-4de5-a352-569ed7ba811e",
            type="device",
            version="2.0"
        ),
        value=[wedge_pb2.Value(
            id=1,
            name="water_flow",
            permission="rw",
            number=wedge_pb2.Number(
                min=0.1,
                max=25.0,
                step=0.1
            ),
            state=[
                wedge_pb2.State(
                    id=1,
                    data="10",
                    type="Control"
                ),
                wedge_pb2.State(
                    id=2,
                    data="3",
                    type="Report"
                    #  type=report_state
                    #  by default type is 0, which is Report enum.
                )
            ]
        )]
    )]
)


class Uplink:
    def __init__(self):
        logging.info("Initialize Uplink object.")
        self.channel = grpc.insecure_channel('localhost:50051')
        try:
            # 1 sec timeout....
            grpc.channel_ready_future(self.channel).result(timeout=1)
        except grpc.FutureTimeoutError:
            sys.exit('Error connecting to server')

        self.stub = wedge_pb2_grpc.WedgeStub(self.channel)

    def SetModel(self, model):
        # metadata is optional, it just for test purpose.
        metadata = [('ip', '127.0.0.1')]
        request = wedge_pb2.SetModelRequest(model=model)
        return self.stub.SetModel(request=request, metadata=metadata)

    def SetState(self, request):
        resp = self.stub.SetState(
            request
        )
        return resp


"""
    Node class will define all methods allowed to call from Wedge.
"""


class TheNode(node_pb2_grpc.NodeServicer):

    def __init__(self, *args, **kwargs):
        pass

    # method will be called from wedge
    def UpdateState(self, request, context):
        print("__UpdateState__")
        return node_pb2.Replay(ok=True)

    def DeleteDevice(self, request, context):
        print("__DeleteDevice__")
        return node_pb2.Replay(ok=True)


"""
    Two async loop - one for listening incomming messages from Wedge,
    another for incomming messages from driver
"""


async def serve():
    server = aio.server()
    node_pb2_grpc.add_NodeServicer_to_server(TheNode(), server)

    listen_addr = '[::]:30052'
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()
    await server.wait_for_termination()


async def driverLoop():
    uplink = Uplink()
    resp = uplink.SetModel(model)
    logging.info("Response: {}".format(resp))

    while True:
        data = '33'
        print('Hello from driver..., data: ', data)
        req = wedge_pb2.SetStateRequest(
            driver=myIdentity,
            device_id=1,
            value_id=1,
            state=wedge_pb2.State(
                id=2,
                data=data
            ),
        )
        resp = uplink.SetState(req)
        logging.info("Response: {}".format(resp))
        await asyncio.sleep(5)


async def main():
    await asyncio.gather(driverLoop(), serve())

if __name__ == '__main__':
    signal.signal(signal.SIGINT, signal_handler)
    logging.basicConfig(level=logging.INFO)

    asyncio.run(main())
