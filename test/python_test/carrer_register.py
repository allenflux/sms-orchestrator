import json
import time

import requests

import url

# register_url = "http://192.168.1.14:8822/api/v1/career/device"
register_url = url.url_prefix + "career/device"
def register():
    with open("./device_num.json") as f:
        device_num_list = json.load(f)
    with open("./phone_num.json") as f:
        phone_num_list = json.load(f)
    num_list = zip(device_num_list, phone_num_list)
    t = time.time()
    res_id_list = []
    for device_no,phone_no in num_list:
        res = requests.post(register_url, json={"device_number": device_no, "phone_number": phone_no, "active_time":t})
        res_json = res.json()
        print(res_json)
        res_id_list.append(res_json["data"]["id"])

    with open("./device_id_list.json", "w") as f:
        json.dump(res_id_list, f)

if __name__ == '__main__':

    register()
