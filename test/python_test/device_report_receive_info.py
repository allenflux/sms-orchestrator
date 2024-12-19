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
        device_number=data["device_number"],
        target_phone_number=data["target_phone_number"],
        sms_content=fake.text(max_nb_chars=10),
        start_time=time.time()
    )
    print("start--->", row)
    resp = requests.post(url=device_report_receive_info_url, json=row)
    print("resp-->", resp.json())

# {
#   "task_id": 0,
#   "device_number": "string",
#   "target_phone_number": "string",
#   "sms_content": "string",
#   "start_time": "string"
# }
if __name__ == "__main__":
    # device_report_receive_info()
    start = time.time()

    loop = asyncio.new_event_loop()
    tasks = [
        # loop.create_task(sum("A", [1, 2])),
        # loop.create_task(sum("B", [1, 2, 3])),
    ]
    for resp in json_data:
        a = loop.create_task(async_device_report_receive_info(resp["data"]))
        tasks.append(a)
    loop.run_until_complete(asyncio.wait(tasks))
    loop.close()

    end = time.time()
    print(f'Time: {end - start:.2f} sec')
