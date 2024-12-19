import requests

from sms_client import settings

"""http://192.168.1.14:8822/api/v1/sms/sub/group/list"""
def controller_sub_group_list(sub_user_id):
    url = settings.url_prefix + "sms/sub/group/list"
    resp = requests.get(url, json={"sub_user_id":sub_user_id})
    if resp.status_code != 200:
        raise Exception(resp.status_code)
    return resp.json()

def controller_sub_group_create(group_name, project_id,sub_user_id):
    url = settings.url_prefix + "sms/sub/group"
    resp = requests.post(url, json={"group_name":group_name,"project_id": project_id, "sub_user_id":sub_user_id})
    if resp.status_code != 200:
        raise Exception(resp.status_code)
    return resp.json()

def allocate_device_2_group(device_id, group_id, sub_user_id):
    url = settings.url_prefix + "sms/sub/device/group"
    resp = requests.post(url,json={"device_id_list":[device_id],"group_id":group_id, "sub_user_id":sub_user_id})
    if resp.status_code != 200:
        raise Exception(resp.status_code)
    return resp.json()