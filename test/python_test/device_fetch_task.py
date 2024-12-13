import json

import requests

import url

# device_fetch_task_url = "http://192.168.1.14:8822/api/v1/career/device/task"
device_fetch_task_url = register_url = url.url_prefix + "career/device/task"
with open("./device_num.json") as f:
    device_num_list = json.load(f)

def device_fetch_task():
    result_list = []
    for device_no in device_num_list:
        resp = requests.post(device_fetch_task_url, json={"device_number": device_no})
        print(resp.json())
        result_list.append(resp.json())
    with open("./fetch_task_result.json", "w") as f:
        json.dump(result_list, f)


if __name__ == '__main__':
    device_fetch_task()

