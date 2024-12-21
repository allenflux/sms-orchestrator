import json
import random
from faker import Faker
fake = Faker(locale='zh_CN')

def generate_phone_number():

    phone_number = fake.phone_number()
    print(phone_number)
    return phone_number


if __name__ == '__main__':
    random_phone_number_list = []
    for i in range(20):
        random_phone_number = generate_phone_number()
        random_phone_number_list.append(random_phone_number)
    with open("phone_num_2.json", "w") as f:
        json.dump(random_phone_number_list, f)