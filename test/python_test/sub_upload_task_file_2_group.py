import json
import time

from faker import Faker
import requests
from  requests_toolbelt import MultipartEncoder

import url

faker = Faker()
# sub_upload_task_url = "http://192.168.1.14:8822/api/v1/sms/sub/task"
sub_upload_task_url = url.url_prefix +"sms/sub/task"

test_sub_user_id = 886

with open("./project_id_list.json") as f:
    project_id_list = json.loads(f.read())

project_id = project_id_list[0]

with open("./group_id.json") as f:
    group_id_list = json.loads(f.read())
group_id = group_id_list[0]

task_name ="python-task-" + faker.word()
group_id = 2
def sub_upload_task():
    print("task_name:",task_name)
    print("group_id:",group_id)
    with open("./device_task_file.json") as f:
        m = MultipartEncoder(
            fields={'interval_time': '1',
                    'timing_start_time': str(time.time()),
                    'sub_user_id': str(test_sub_user_id),
                    "project_id":str(project_id),
                    "task_name":task_name,
                    "group_id":str(group_id),
                    'file': ('filename', f.read(), 'text/plain')}
        )
        resp = requests.post(sub_upload_task_url,
                             data=m,
                             headers={'Content-Type': m.content_type})
    print(resp.text)
    # Save Task ID
    if resp.status_code == 200:
        with open("./sub_upload_task_id.json", "w") as f:
            json.dump(resp.json()["data"], f)


if __name__ == '__main__':
    sub_upload_task()