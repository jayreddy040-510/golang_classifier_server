import os
import sys
import torch
from dotenv import load_dotenv
from transformers import AutoTokenizer, AutoModelForSequenceClassification

load_dotenv()

model_path = os.environ.get("MODEL_PATH")
tokenizer_path = os.environ.get("TOKENIZER_PATH")
tokenizer = AutoTokenizer.from_pretrained(tokenizer_path)
model = AutoModelForSequenceClassification.from_pretrained(model_path)


def classify_sms(sms: str):
    # set appropriate device if avail (cuda > mps > cpu)
    device = "mps" if torch.backends.mps.is_available() else "cpu"
    if torch.cuda.is_available():
        device = "cuda"

    # tokenize sms
    inputs = tokenizer(sms,
                       padding=True,
                       truncation=True,
                       max_length=128,
                       return_tensors="pt")

    # add model to "mps" device
    model.to(torch.device(device))

    # add inputs to "mps" device
    inputs = {k: v.to(torch.device(device)) for k, v in inputs.items()}
    print(f"device: {model.device}")

    # get model predictions
    with torch.no_grad():
        outputs = model(**inputs)

    # find index witrh highest score
    label_idx = torch.argmax(outputs.logits, dim=-1).item()

    return label_idx


"""
if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Please pass an SMS")
        sys.exit(1)

    sms = sys.argv[1]
    sms_spam_classification = classify_sms(sms)
    print(sms_spam_classification)
"""
print(classify_sms("For the low price $9.99 you can get your very own used car! Call now at 888-888-8888!"))
