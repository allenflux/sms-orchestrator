import requests
from faker import Faker
fake = Faker()

sub_post_conversation_record_url = "http://127.0.0.1:8822/api/v1/sms/sub/conversation/record"

test_sub_user_id = 886

def sub_post_conversation_record():
    data = {
        "sub_user_id":test_sub_user_id,
        "chart_log_id": 48,
        "content":fake.text(max_nb_chars=10),
    }
    print("request data - >", data)
    resp = requests.post(sub_post_conversation_record_url, json=data)
    print("receive data")
    print(resp.text)

if __name__ == '__main__':
    sub_post_conversation_record()