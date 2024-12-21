import json

import requests

import url

# allocate_device_2_project_url = "http://192.168.1.14:8822/api/v1/sms/device/project"
allocate_device_2_project_url = url.url_prefix +"sms/device/project"

def auto_allocate_device_2_project():
    url = allocate_device_2_project_url
    with open("./device_id_list_2.json") as f:
        device_id_list = json.load(f)
    with open("./project_id_list.json") as f:
        project_id_list = json.load(f)
    # allocate all device to first project
    resp = requests.post(url, json={"device_id_list": [str(x) for x in device_id_list], "project_id": project_id_list[0]})
    print(resp.json())

def allocate_device_2_project(device_id_list, project_id_list):
    url = allocate_device_2_project_url
    # allocate all device to first project
    resp = requests.post(url, json={"device_id_list": [str(x) for x in device_id_list], "project_id": project_id_list[0]})
    print(resp.json())

if __name__ == '__main__':
    auto_allocate_device_2_project()