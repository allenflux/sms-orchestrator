"""
{
"target_phone_number":["15122222222"],
"content":["Hi Hi Hi"]
}
"""
import json

from faker import Faker
fake = Faker()

def generate_device_task_file():
    content_list = []
    phone_list = []
    for i in range(200):
        content_list.append(fake.text(max_nb_chars=200))
        phone_list.append(fake.phone_number())

    with open("device_task_file.json", "w") as f:
        json.dump({"target_phone_number": phone_list, "content": content_list}, f)


if __name__ == '__main__':
    generate_device_task_file()