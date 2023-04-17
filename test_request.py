import os
import requests
from dotenv import load_dotenv

load_dotenv()
server_url = os.environ.get("SERVER_URL")
api_key = os.environ.get("API_KEY")
sms = "jay nana, please call me back"
# sms = "hiii!"
payload = {"sms": sms}
headers = {"apiKey": api_key}
res = requests.post(server_url,
                    json=payload,
                    headers=headers)

print(res.text)
