#!/usr/bin/python3

# import aiogrpc
import asyncio
import logging
import grpc
import signal
import sys
# pregenerated from proto file
import wedge_pb2
import wedge_pb2_grpc

from button import Button
from led import Led

CHANNEL_OPTIONS = [('grpc.lb_policy_name', 'pick_first'),
                   ('grpc.enable_retries', 0),
                   ('grpc.keepalive_timeout_ms', 10000)]


node_id = wedge_pb2.NodeIdentity(
    id="python_node_client_led_button",
)

"""
Uplink class holds all remote Wedge Methods
"""


class Uplink:
    def __init__(self, stub, node_id):
        logging.info("Initialize Uplink object.")
        self.stub = stub
        self.node_id = node_id
        signal.signal(signal.SIGINT, self.cancel_request)

    def cancel_request(self, signal, frame):
        print("")
        logging.warning("Interrupt signal. Exiting!")
        self._future.cancel()
        sys.exit(0)

    async def SetModel(self, model):
        # metadata is optional, it just for test purpose.
        metadata = [('ip', '127.0.0.1')]
        request = wedge_pb2.SetModelRequest(model=model)
        return await self.stub.SetModel(request=request, metadata=metadata)

    async def SetState(self, value_id, state):
        logging.info("Update value with id {} to: {}"
                     .format(value_id, state.data))

        req = wedge_pb2.SetStateRequest(
            node=self.node_id,
            device_id=1,  # only one device exist in this model
            value_id=value_id,
            state=state
        )
        resp = await self.stub.SetState(req)
        logging.info("Replay: {}".format(resp))

    async def GetControl(self, request):
        self._future = self.stub.GetControl(request)
        return await self._future


async def main() -> None:
    async with grpc.aio.insecure_channel(target='localhost:50051') as channel:
        stub = wedge_pb2_grpc.WedgeStub(channel)

        uplink = Uplink(stub, node_id)
        button = Button(uplink)
        led = Led(uplink)

        # Example of Model, which identical to Seluxit data Model.
        model = wedge_pb2.Model(
            node=node_id,
            device=[wedge_pb2.Device(
                id=1,
                name="LED_Button2",
                version="0.1.2",
                value=[
                    button.value,
                    led.value
                ]
            )]
        )

        resp = await uplink.SetModel(model)
        logging.info("Response: {}".format(resp))

        # data = button.pin.read()
        # logging.info("Current button status: {}".format(data))
        # button.update(data)
        await asyncio.gather(led.listen(node_id))


if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    asyncio.run(main())
