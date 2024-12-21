import asyncio
import json
import random
import time

import requests

import url

# device_fetch_task_url = "http://192.168.1.14:8822/api/v1/career/device/task"
device_fetch_task_url  = url.url_prefix + "career/device/task"
with open("./device_num.json") as f:
    device_num_list = json.load(f)

def device_fetch_task():
    result_list = []
    for device_no in range(20):
        print("start ",time.time())
        resp = requests.post(device_fetch_task_url, json={"device_number": random.choice(device_num_list)})
        # resp = requests.post(device_fetch_task_url, json={"device_number": "python_test_device_seat.jpg523"})
        print(resp.json())
        result_list.append(resp.json())
    with open("./fetch_task_result.json", "w") as f:
        json.dump(result_list, f)

async def async_device_fetch_task():
    resp = requests.post(device_fetch_task_url, json={"device_number": random.choice(device_num_list)})
    print(resp.json())



if __name__ == '__main__':
    device_fetch_task()

    # start = time.time()
    #
    # loop = asyncio.new_event_loop()
    # tasks = [
    #     # loop.create_task(sum("A", [1, 2])),
    #     # loop.create_task(sum("B", [1, 2, 3])),
    # ]
    # for device_no in range(200):
    #     a = loop.create_task(async_device_fetch_task())
    #     tasks.append(a)
    # loop.run_until_complete(asyncio.wait(tasks))
    # loop.close()
    #
    # end = time.time()
    # print(f'Time: {end - start:.2f} sec')

