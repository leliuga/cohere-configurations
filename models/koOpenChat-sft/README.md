---
license: cc-by-sa-4.0
---

# **koOpenChat-sft🐧**  

## Support Me
시나트라는 개인 프로젝트로, 1인의 자원으로 개발되고 있습니다. 모델이 마음에 드셨다면 약간의 연구비 지원은 어떨까요?
[<img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy me a Coffee" width="217" height="50">](https://www.buymeacoffee.com/mwell)

Wanna be a sponser? (Please) Contact me on Telegram **AlzarTakkarsen**


# **Model Details**
**Base Model**  
OpenChat3.5

**Trained On**  
A100 80GB * 1

**Instruction format**

It follows [ChatML](https://github.com/openai/openai-python/blob/main/chatml.md) format and **Alpaca(No-Input)** format.

# **Model Benchmark**
None

# **Implementation Code**

Since, chat_template already contains insturction format above.
You can use the code below.

```python
from transformers import AutoModelForCausalLM, AutoTokenizer

device = "cuda" # the device to load the model onto

model = AutoModelForCausalLM.from_pretrained("maywell/koOpenChat-sft")
tokenizer = AutoTokenizer.from_pretrained("maywell/koOpenChat-sft")

messages = [
    {"role": "user", "content": "바나나는 원래 하얀색이야?"},
]

encodeds = tokenizer.apply_chat_template(messages, return_tensors="pt")

model_inputs = encodeds.to(device)
model.to(device)

generated_ids = model.generate(model_inputs, max_new_tokens=1000, do_sample=True)
decoded = tokenizer.batch_decode(generated_ids)
print(decoded[0])
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_maywell__koOpenChat-sft)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 51.36   |
| ARC (25-shot)         | 59.81          |
| HellaSwag (10-shot)   | 78.73    |
| MMLU (5-shot)         | 61.32         |
| TruthfulQA (0-shot)   | 51.24   |
| Winogrande (5-shot)   | 76.4   |
| GSM8K (5-shot)        | 24.18        |
| DROP (3-shot)         | 7.82         |
