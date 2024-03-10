---
license: cc-by-sa-4.0
---

# **Synatra-7B-v0.3-dpo🐧**  
![Synatra-7B-v0.3-dpo](./Synatra.png)

## Support Me
시나트라는 개인 프로젝트로, 1인의 자원으로 개발되고 있습니다. 모델이 마음에 드셨다면 약간의 연구비 지원은 어떨까요?
[<img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy me a Coffee" width="217" height="50">](https://www.buymeacoffee.com/mwell)

Wanna be a sponser? (Please) Contact me on Telegram **AlzarTakkarsen**

# **License**

This model is strictly [*cc-by-sa-4.0*](https://creativecommons.org/licenses/by-sa/4.0/) use, Under **5K MAU**
The "Model" is completely free (ie. base model, derivates, merges/mixes) to use for non-commercial purposes as long as the the included **cc-by-sa-4.0** license in any parent repository, and the non-commercial use statute remains, regardless of other models' licences.

# **Model Details**
**Base Model**  
[mistralai/Mistral-7B-Instruct-v0.1](https://huggingface.co/mistralai/Mistral-7B-Instruct-v0.1)  

**Trained On**  
A100 80GB * 1

**Instruction format**

It follows [ChatML](https://github.com/openai/openai-python/blob/main/chatml.md) format and **Alpaca(No-Input)** format.

# **Model Benchmark**

## KOBEST_BOOLQ, SENTINEG, WIC - ZERO_SHOT
[EleutherAI/lm-evaluation-harness](https://github.com/EleutherAI/lm-evaluation-harness/tree/polyglot)를 사용하여 BoolQ, SentiNeg, Wic을 측정했습니다.

| Model | COPA | HellaSwag | BoolQ | SentiNeg
| --- | --- | --- | --- | ---
| EleutherAI/polyglot-ko-12.8b | 0.7937 | 0.5954 | 0.4818 | 0.9117
| Synatra-7B-v0.3-base | 0.6344 | 0.5140 | 0.5226 | NaN
| **Synatra-7B-v0.3-dpo** | **0.6380** | **0.4780** | **0.8058** | **0.8942**

## Ko-LLM-Leaderboard

On Benchmarking...

# **Implementation Code**

Since, chat_template already contains insturction format above.
You can use the code below.

```python
from transformers import AutoModelForCausalLM, AutoTokenizer

device = "cuda" # the device to load the model onto

model = AutoModelForCausalLM.from_pretrained("maywell/Synatra-7B-v0.3-dpo")
tokenizer = AutoTokenizer.from_pretrained("maywell/Synatra-7B-v0.3-dpo")

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
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_maywell__Synatra-7B-v0.3-dpo)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 53.14   |
| ARC (25-shot)         | 62.8          |
| HellaSwag (10-shot)   | 82.58    |
| MMLU (5-shot)         | 61.46         |
| TruthfulQA (0-shot)   | 56.46   |
| Winogrande (5-shot)   | 76.24   |
| GSM8K (5-shot)        | 23.73        |
| DROP (3-shot)         | 8.68         |
