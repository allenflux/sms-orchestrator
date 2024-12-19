import json

import requests

sub_get_conversation_record_url = "http://127.0.0.1:8822/api/v1/sms/sub/conversation/record"

test_sub_user_id = 886


def sub_get_conversation_record():
    data = {
        "sub_user_id": test_sub_user_id,
        "chat_log_id": 48
    }
    print("request data : ", data)
    resp = requests.get(sub_get_conversation_record_url, json=data)
    print("receive data", resp.text)
    print(len(resp.json()["data"]["data"]))


if __name__ == "__main__":
    sub_get_conversation_record()