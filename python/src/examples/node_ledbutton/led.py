import wedge_pb2
import logging
import asyncio
# from PorcupineIO import GPIO


class Led():
    def __init__(self, uplink):
        self.uplink = uplink
        # self.pin = GPIO(chip=2,
                        # pin=13,
                        # state=GPIO.OUTPUT)
        #
        self.value = wedge_pb2.Value(
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
        )

    def write(self, data):
        print("Writing data {}".format(data))
        # self.pin.write(bool(data))

    async def listen(self, node_id):
        while True:
            logging.info("Sending request for control")
            try:
                resp = await self.uplink.GetControl(wedge_pb2.GetControlRequest(
                    node=node_id
                ))
                logging.info("Got control message: {}".format(resp))
                if resp.update and resp.update.device_id == 1 and resp.update.value_id == 1:
                    data = int(resp.update.state.data)
                    self.write(bool(data))

            except Exception as e:
                print("Something bad happened {}, re-try".format(e))
                await asyncio.sleep(5)


    # def __del__(self):
        # self.pin.close()
