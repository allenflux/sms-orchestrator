import json

import requests

import url

# allocate_device_2_group_url = "http://192.168.1.14:8822/api/v1/sms/sub/device/group"
allocate_device_2_group_url = url.url_prefix +"sms/sub/device/group"
test_sub_user_id = 886

def auto_allocate_device_2_group():
    with open("device_id_list_2.json") as f:
        device_id_list = json.load(f)
    with open("group_id.json") as f:
        group_id_list = json.load(f)
    group_id = group_id_list[0]
    resp = requests.post(allocate_device_2_group_url,json={"device_id_list":[str(x) for x in device_id_list],"group_id":group_id, "sub_user_id":test_sub_user_id})
    print(resp.json())
def allocate_device_2_group(device_id_list, group_id):
    resp = requests.post(allocate_device_2_group_url,json={"device_id_list":[str(x) for x in device_id_list],"group_id":group_id, "sub_user_id":test_sub_user_id})
    print(resp.json())

if __name__ == '__main__':
    auto_allocate_device_2_group()
    # allocate_device_2_group([63], 2)