import json

import requests

sub_get_conversation_record_list_url = "http://127.0.0.1:8822/api/v1/sms/sub/conversation/record/list"

test_sub_user_id = 886

with open("./project_id_list.json") as f:
    project_id_list = json.loads(f.read())

project_id = project_id_list[0]

def sub_get_conversation_record_list():
    data = {
        "sub_user_id": test_sub_user_id,
        "project_id": project_id
    }
    resp = requests.get(sub_get_conversation_record_list_url, json=data)
    print(resp.text)
    print(len(resp.json()["data"]["data"]))


if __name__ == "__main__":
    sub_get_conversation_record_list()