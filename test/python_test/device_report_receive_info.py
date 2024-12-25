import asyncio
import json
import time

import requests

import url

device_report_receive_info_url = url.url_prefix + "career/device/task/content"
from faker import Faker
fake = Faker()
with open("./fetch_task_result.json") as f:
    json_data = json.load(f)

def device_report_receive_info():
    data = json_data[-1]["data"]
    row = dict(
    # task_id = data["task_id"],
    device_number = data["device_number"],
    target_phone_number = data["target_phone_number"],
    sms_content = fake.text(max_nb_chars=10),
    start_time = time.time()
    )
    print("row--->", row)
    resp = requests.post(url=device_report_receive_info_url, json=row)
    print("resp-->",resp.json())

async def async_device_report_receive_info(data):
    row = dict(
        # task_id = data["task_id"],
        # device_number=data["device_number"],
        device_number="753-Anna",
        target_phone_number="001-988-333-8387x67535",
        sms_content=str(data)+"---"+fake.text(max_nb_chars=200),
        start_time=time.time()
    )
    # print("start--->", row)
    resp = requests.post(url=device_report_receive_info_url, json=row)
    if resp.status_code != 200:
        raise Exception(resp.text)
    data = resp.json()
    if data["code"] != 0:
        raise Exception(resp.text)

# {
#   "task_id": 0,
#   "device_number": "string",
#   "target_phone_number": "string",
#   "sms_content": "string",
#   "start_time": "string"
# }


from multiprocessing import Process, Pool
import os


def io_test(num):
    print('Run child process %s (%s)...' % ("io_test", os.getpid()))
    start = time.time()
    loop = asyncio.new_event_loop()
    tasks = [
        # loop.create_task(sum("A", [1, 2])),
        # loop.create_task(sum("B", [1, 2, 3])),
    ]
    # for resp in json_data:
    for resp in range(num):
        a = loop.create_task(async_device_report_receive_info(resp))
        tasks.append(a)
    loop.run_until_complete(asyncio.wait(tasks))
    loop.close()
    end = time.time()
    print(f'Time: {end - start:.2f} sec')

if __name__ == "__main__":
    # device_report_receive_info()
    print('Parent process %s.' % os.getpid())
    p = Pool(100)
    for i in range(1000):
        p.apply_async(io_test, args=(1000,))
    print('Waiting for all subprocesses done...')
    p.close()
    p.join()
    print('All subprocesses done.')
    # start = time.time()
    #
    # loop = asyncio.new_event_loop()
    # tasks = [
    #     # loop.create_task(sum("A", [1, 2])),
    #     # loop.create_task(sum("B", [1, 2, 3])),
    # ]
    # for resp in json_data:
    # for resp in range(10):
    #     a = loop.create_task(async_device_report_receive_info(resp))
    #     tasks.append(a)
    # loop.run_until_complete(asyncio.wait(tasks))
    # loop.close()
    #
    # end = time.time()
    # print(f'Time: {end - start:.2f} sec')
