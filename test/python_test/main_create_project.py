import json

import requests

create_project_url = "http://192.168.1.14:8822/api/v1/sms/project"

def create_project():
    res_id_list = []
    for i in range(10):
        res = requests.post(create_project_url, json={"project_name":"python_project_" + str(i), "note":"Python Test Note"})
        res_json = res.json()
        print(res_json)
        res_id_list.append(res_json["data"]["id"])
    with open("./project_id_list.json", "w") as f:
        json.dump(res_id_list, f)

if __name__ == '__main__':
    create_project()