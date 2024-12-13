import json
import random

import requests

import url

# report_task_result_url = "http://192.168.1.14:8822/api/v1/career/device/task/result"
report_task_result_url = url.url_prefix + "career/device/task/result"


def report_task_result():
    with open("./fetch_task_result.json") as f:
        fetch_task_list = json.load(f)
    l = [1, 2]
    sms_failure_reasons = [
        "Message too long",
        "Invalid phone number",
        "Insufficient credit",
        "Network error",
        "Gateway timeout",
        "SMS service unavailable",
        "Blocked number",
        "Content policy violation",
        "Delivery failure",
        "Unknown error",
        "Temporary service disruption",
        "Internal server error"
    ]
    for task in fetch_task_list:
        # `json:"send_status" dc:"1-成功 2-失败" v:"required"`
        send_status = random.choice(l)
        reason = "Great !!! Anything is fine!"
        if send_status == 2:
            reason = random.choice(sms_failure_reasons)
        data = {
            "device_number": task["data"]["device_number"],
            "content": task["data"]["content"],
            "task_id": task["data"]["task_id"],
            "target_phone_number": task["data"]["target_phone_number"],
            "start_time": task["data"]["start_at"],
            "send_status": send_status,
            "reason": reason
        }
        resp = requests.post(url=report_task_result_url, json=data)
        print(resp.json())
if __name__ == '__main__':
    report_task_result()