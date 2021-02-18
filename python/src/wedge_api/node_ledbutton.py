#!/usr/bin/python3

# import aiogrpc
# import asyncio
import logging
import sys
import signal
import time
import grpc
import time
# pregenerated from proto file
import wedge_pb2
import wedge_pb2_grpc
from PorcupineIO import GPIO

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
    id="python_node_client_led_button",
)

device = [wedge_pb2.Device(
    id=1,
    name="LED_Button",
    version="0.1.2",
    value=[
        wedge_pb2.Value(
            id=1,
            name="led",
            permission="w",
            number=wedge_pb2.Number(
                min=0.01,
                max=1,
                step=0.99
            ),
            state=[
                wedge_pb2.State(
                    id=1,
                    data="0",
                    type="Control"
                )
            ]
        ),
        wedge_pb2.Value(
            id=2,
            name="button",
            permission="r",
            number=wedge_pb2.Number(
                min=0.01,
                max=1,
                step=0.99
            ),
            state=[
                wedge_pb2.State(
                    id=1,
                    data="0",
                    type="Report",
                )
            ]
        ),
    ]
)]


# Example of Model, which identical to Seluxit data Model.
model = wedge_pb2.Model(
    node=myIdentity,
    device=device
)


class Button():
    def __init__(self, uplink, chip, pin, **kwargs):
        self.cb_last_active = time.time()
        self.uplink = uplink
        self.pin = GPIO(
            chip=int(chip),
            pin=int(pin),
            state=GPIO.BOTH_EDGE,
            pull=GPIO.PULL_UP,
            callback=self._callback
        )
        logging.info("Button Initialized.")
 
    def update(self, data):
        value = [val for val in device[0].value if val.name == "button"]
        state = value[0].state[0]
        state.data = str(data)
        #
        req = wedge_pb2.SetStateRequest(
            node=myIdentity,
            device_id=1,
            value_id=value[0].id,
            state=state
        )
        resp = self.uplink.SetState(req)
        logging.info("Replay: {}".format(resp))
   
    def _callback(self, chip, pin, edge):
        if (self.cb_last_active + 0.2) < time.time():
            self.cb_last_active = time.time()
            data = self.pin.read()
            logging.info("Update button value to: {}".format(data))
            self.update(data)

    def handle_close(self):
        self.pin.close()


class Uplink:
    def __init__(self, stub):
        logging.info("Initialize Uplink object.")
        self.stub = stub

    def SetModel(self, model):
        # metadata is optional, it just for test purpose.
        metadata = [('ip', '127.0.0.1')]
        request = wedge_pb2.SetModelRequest(model=model)
        return self.stub.SetModel(request=request, metadata=metadata)

    def SetState(self, request):
        return self.stub.SetState(request)

    def GetControl(self, request):
        return self.stub.GetControl(request)


def listen(uplink):
    # LED control pin.
    pinLed = GPIO(chip=2, pin=13, state=GPIO.OUTPUT)
    while True:
        logging.info("Sending request for control")
        try:
            resp = uplink.GetControl(wedge_pb2.GetControlRequest(
                node=myIdentity
            ))
            print("Got control message: {}".format(resp))
            if resp.update:
                value = [val for val in device[0].value if val.id == resp.update.value_id]
                if value and value[0].name == "led":
                    data = int(resp.update.state.data)
                    pinLed.write(bool(data))

        except Exception as e:
            print("Something bad happened {}, re-try".format(e))
            time.sleep(5)


def main():
    with grpc.insecure_channel(target='localhost:50051',
                                         options=CHANNEL_OPTIONS) as channel:
        stub = wedge_pb2_grpc.WedgeStub(channel)
        uplink = Uplink(stub)

        resp = uplink.SetModel(model)
        logging.info("Response: {}".format(resp))
        
        button_driver = Button(uplink=uplink, chip=2, pin=14)
        data = button_driver.pin.read()
        logging.info("Current button status: {}".format(data))
        button_driver.update(data)

        listen(uplink)


if __name__ == '__main__':
    signal.signal(signal.SIGINT, signal_handler)
    logging.basicConfig(level=logging.INFO)
    main()
