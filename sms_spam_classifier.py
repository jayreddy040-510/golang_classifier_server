import os
import sys
import torch
from dotenv import load_dotenv
from transformers import AutoTokenizer, AutoModelForSequenceClassification

load_dotenv()

model_path = os.environ.get("MODEL_PATH")
tokenizer = AutoTokenizer.from_pretrained(model_path)
model = AutoModelForSequenceClassification.from_pretrained(model_path)


def classify_sms(sms: str):
    if torch.mps.is_available():
        model.to(torch.device("mps"))
        print(f"Model device: {model.device}")
    # tokenize sms
    inputs = tokenizer(sms,
                       padding=True,
                       truncation=True,
                       max_length=128,
                       return_tensors="pt")

    # get model predictions
    with torch.no_grad():
        outputs = model(**inputs)

    # find index witrh highest score
    label_idx = torch.argmax(outputs.logits, dim=-1).item()

    return label_idx


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Please pass an SMS")
        sys.exit(1)

    sms = sys.argv[1]
    sms_spam_classification = classify_sms(sms)
    print(sms_spam_classification)
