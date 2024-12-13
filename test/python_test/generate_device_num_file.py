import json
import random
prefix = "python_test_device_"

def generate_device_num(num = 20):
    filter_num = dict()
    res_list = []
    for i in range(num):
        while True:
            no = random.randint(100, 999)
            if no in filter_num:
                pass
            else:
                filter_num[no] = 1
                res_list.append(prefix + str(no))
                break
    return res_list



if __name__ == "__main__":
    res_list = generate_device_num()
    print(res_list)
    with open("device_num.json", "w") as f:
        json.dump(res_list, f)

