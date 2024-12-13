import json

import requests
from faker import Faker
fake = Faker()

create_group_url = "http://192.168.1.14:8822/api/v1/sms/sub/group"
test_sub_user_id = 886
with open("./project_id_list.json") as f:
    project_id_list = json.load(f)
test_project_id = project_id_list[0]

def create_group():
    filter_m = {}
    for i in range(20):
        while True:
            random_name = "Python-Group-"+ fake.first_name() + "-" + str(i)
            if random_name not in filter_m:
                filter_m[random_name] = 1
                break
    result = list(filter_m)
    group_id_list = []
    for name in result:
        resp = requests.post(create_group_url, json={"project_id":test_project_id,"sub_user_id": test_sub_user_id, "group_name": name})
        print(resp.text)
        if resp.status_code == 200:
            group_id_list.append(resp.json()["data"]["id"])
        else:
            print(resp.status_code)
            return

    with open("./group_id.json", "w") as f:
        json.dump(group_id_list, f)




if __name__ == '__main__':
    create_group()