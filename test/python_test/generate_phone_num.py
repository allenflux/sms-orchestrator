import json
import random


def generate_phone_number():
    area_code = ["13", "14", "15", "16", "17", "18", "19"]  # 手机号码前缀
    middle_number = str(random.randint(0, 999)).zfill(3)  # 中间部分随机三位数，不足三位前面补零
    last_number = str(random.randint(0, 9999)).zfill(4)  # 最后四位随机数，不足四位前面补零

    phone_number = random.choice(area_code) + middle_number + last_number

    return phone_number


if __name__ == '__main__':
    random_phone_number_list = []
    for i in range(20):
        random_phone_number = generate_phone_number()
        random_phone_number_list.append(random_phone_number)
    with open("phone_num.json", "w") as f:
        json.dump(random_phone_number_list, f)