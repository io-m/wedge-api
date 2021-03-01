import time
import wedge_pb2
# from PorcupineIO import GPIO
import functools
import asyncio


def force_sync(fn):
    '''
    turn an async function to sync function
    '''

    @functools.wraps(fn)
    def wrapper(*args, **kwargs):
        res = fn(*args, **kwargs)
        if asyncio.iscoroutine(res):
            return asyncio.get_event_loop().run_until_complete(res)
        return res

    return wrapper


class Button():
    def __init__(self, uplink):
        self.cb_last_active = time.time()
        self.uplink = uplink
        # self.pin = GPIO(
            # chip=2,
            # pin=14,
            # state=GPIO.BOTH_EDGE,
            # pull=GPIO.PULL_UP,
            # callback=self._callback
        # )
        self.value = wedge_pb2.Value(
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
        )

    def update(self, data):
        value_id = self.value[0].id
        state = self.value[0].state[0]
        state.data = str(data)
        #
        self.uplink.SetState(value_id, state)

        @force_sync
        async def _update(self, value_id, state):
            resp = await self.uplink.SetState(value_id, state)
            return resp

        return _update()

    def _callback(self, chip, pin, edge):
        if (self.cb_last_active + 0.2) < time.time():
            self.cb_last_active = time.time()
            # data = self.pin.read()
            data = 0
            self.update(data)

    # def __del__(self):
        # self.pin.close()
