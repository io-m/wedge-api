#!/usr/bin/python3

# from grpc import aio
# for creating protobuf struct from python dictionary.
# from google.protobuf.struct_pb2 import Struct
# for conversion protobuf to json.
# from google.protobuf import json_format
import asyncio
import logging
import sys
import signal
import time
import grpc
# pregenerated from proto file
import wedge_pb2
import wedge_pb2_grpc

CHANNEL_OPTIONS = [('grpc.lb_policy_name', 'pick_first'),
                   ('grpc.enable_retries', 0),
                   ('grpc.keepalive_timeout_ms', 10000)]


def signal_handler(signal, frame):
    sys.exit(0)


"""
Uplink class holds all remote Wedge Methods
"""

# alias for state type: 'Control' or 'Report'
# control_state = wedge_pb2.StateType.Control
# report_state = wedge_pb2.StateType.Report

myIdentity = wedge_pb2.NodeIdentity(
    id="unique_python_node_client",
)

# Example of Model, which identical to Seluxit data Model.
model = wedge_pb2.Model(
    node=myIdentity,
    device=[wedge_pb2.Device(
        id=1,
        name="water_control",
        version="0.1.2",
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
    def __init__(self, stub):
        logging.info("Initialize Uplink object.")
        self.stub = stub

    async def SetModel(self, model):
        # metadata is optional, it just for test purpose.
        metadata = [('ip', '127.0.0.1')]
        request = wedge_pb2.SetModelRequest(model=model)
        return await self.stub.SetModel(request=request, metadata=metadata)

    async def SetState(self, request):
        return await self.stub.SetState(request)

    async def GetControl(self, request):
        return await self.stub.GetControl(request)


"""
    Two async loop - one for listening incomming control messages from Wedge,
    another for incomming messages from driver
"""


async def listen(uplink):
    logging.info("Start listen for Control messages")
    while True:
        logging.info("Sending request for control")
        try:
            resp = await uplink.GetControl(wedge_pb2.GetControlRequest(
                node=myIdentity
            ))
            print("Got control message: {}".format(resp))
            # Do something accordingly.
        except Exception as e:
            print("Something bad happened {}, re-try".format(e))
            time.sleep(5)


async def driverLoop(uplink):
    resp = await uplink.SetModel(model)
    logging.info("Response: {}".format(resp))

    while True:
        data = '22'
        print('Hello from driver..., data: ', data)
        req = wedge_pb2.SetStateRequest(
            node=myIdentity,
            device_id=1,
            value_id=1,
            state=wedge_pb2.State(
                id=2,
                data=data
            ),
        )
        resp = await uplink.SetState(req)
        logging.info("Response: {}".format(resp))
        await asyncio.sleep(5)


async def main() -> None:
    async with grpc.aio.insecure_channel(target='localhost:50051',
                                         options=CHANNEL_OPTIONS) as channel:
        stub = wedge_pb2_grpc.WedgeStub(channel)
        uplink = Uplink(stub)
        await asyncio.gather(driverLoop(uplink), listen(uplink))

if __name__ == '__main__':
    signal.signal(signal.SIGINT, signal_handler)
    logging.basicConfig(level=logging.INFO)

    asyncio.run(main())
