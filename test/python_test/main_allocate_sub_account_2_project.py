import json

import requests

import url

# allocate_sub_account_2_project_url = "http://192.168.1.14:8822/api/v1/sms/account/project"
allocate_sub_account_2_project_url = url.url_prefix +"sms/account/project"

def allocate_sub_account_2_project():
    with open("./project_id_list.json") as f:
        project_id_list = json.load(f)
    account_id = 886
    for project_id in project_id_list:
        resp = requests.post(allocate_sub_account_2_project_url, json={"project_id": project_id, "account_id":account_id})
        print(resp.json())

if __name__ == '__main__':
    allocate_sub_account_2_project()
