import json
import time

import requests

import url

# register_url = "http://192.168.1.14:8822/api/v1/career/device"
register_url = url.url_prefix + "career/device"
def register():
    with open("./device_num_2.json") as f:
        device_num_list = json.load(f)
    with open("./phone_num_2.json") as f:
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

def m_register():
    t = time.time()
    num_list = [
        "3663555168",
        "3145635457",
        "4287886250",
        "2047100828",
        "267540103",
        "3766646931",
        "2452174685",
        "3805997622",
        "131417821",
        "646026451"
    ]
    res_id_list = []
    for device_no in num_list:
        res = requests.post(register_url, json={"device_number": device_no, "phone_number": "NoPhoneNumber", "active_time": t})
        res_json = res.json()
        print(res_json)
        res_id_list.append(res_json["data"]["id"])
    with open("./device_id_list_2.json", "w") as f:
        json.dump(res_id_list, f)
if __name__ == '__main__':
    register()
    # m_register()
