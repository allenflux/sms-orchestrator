import json
from hashlib import sha256
with open("./device_task_file.json") as f:
    device_task = json.load(f)

with open("./fetch_task_result.json") as f:
    fetch_task_result = json.load(f)

def check_device_task_file():
    device_task_file_str = ""
    for target_phone_number in device_task["target_phone_number"]:
        device_task_file_str += target_phone_number
    for target_phone_number in device_task["content"]:
        device_task_file_str += target_phone_number
    o_hash = sha256(device_task_file_str.encode('utf-8')).hexdigest()
    print(o_hash)
    fetch_task_result_str = ""
    for data in fetch_task_result:
        fetch_task_result_str += data["data"]["target_phone_number"]
    for data in fetch_task_result:
        fetch_task_result_str += data["data"]["content"]
    f_hash = sha256(fetch_task_result_str.encode('utf-8')).hexdigest()
    print(f_hash)
    print(o_hash == f_hash)


if __name__ == '__main__':
    check_device_task_file()