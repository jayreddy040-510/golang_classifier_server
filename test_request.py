import os
import requests
from dotenv import load_dotenv

load_dotenv()
server_url = os.environ.get("SERVER_URL")
api_key = os.environ.get("API_KEY")
sms = "For the low price of 9.99 you can get a free panhandle!! Reply 100 now!"
sms = "hiii!"
payload = {"sms": sms}
headers = {"apiKey": api_key}
res = requests.post(server_url,
                    json=payload,
                    headers=headers)

print(res.text)
