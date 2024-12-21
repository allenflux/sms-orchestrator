import json
import random
from faker import Faker
fake = Faker()

prefix = "python_test_device_"

def generate_device_num(num = 20):
    res_list = []
    for i in range(num):
        res = fake.numerify() +"-"+ fake.first_name()
        res_list.append(res)
    return res_list



if __name__ == "__main__":
    res_list = generate_device_num()
    print(res_list)
    with open("device_num_2.json", "w") as f:
        json.dump(res_list, f)

