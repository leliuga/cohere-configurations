---
license: apache-2.0
language:
- en
---

This model is finetuend based on "mistralai/Mixtral-8x7B-v0.1" with [Firefly](https://github.com/yangjianxin1/Firefly) and 48k data from ultrachat.

## Evaluation
Though we finetune with only 48k data, our model can also achieve excellent performance.

| Model                                                                                        | Open LLM Leaderboard         |
|------------------------------------------------------------------------------------------------|---------------------------------------------|
| Qwen-72B  | 73.6  |
| Mixtral-8x7B-Instruct-v0.1 |        72.62                                   |
|**Firefly-Mixtral-8x7B**|**70.34**|
|Yi-34B|69.42|
|Mixtral-8x7B-v0.1|68.42|
|Llama2-65B-Chat|67.87|
|Qwen-14B|65.86|
|Vicuna-33B-v1.3 |58.54|


## Run the model

```python
from transformers import AutoModelForCausalLM, AutoTokenizer
import torch

model_name_or_path = 'YeungNLP/firefly-mixtral-8x7b'
max_new_tokens = 500
top_p = 0.9
temperature = 0.35
repetition_penalty = 1.0

model = AutoModelForCausalLM.from_pretrained(
    model_name_or_path,
    trust_remote_code=True,
    low_cpu_mem_usage=True,
    torch_dtype=torch.float16,
    device_map='auto'
)
model = model.eval()
tokenizer = AutoTokenizer.from_pretrained(model_name_or_path)

text = "Compose an engaging travel blog post about a recent trip to Hawaii, highlighting cultural experiences and must-see attractions."

inst_begin_tokens = tokenizer.encode('[INST]', add_special_tokens=False)
inst_end_tokens = tokenizer.encode('[/INST]', add_special_tokens=False)
human_tokens = tokenizer.encode(text, add_special_tokens=False)
input_ids = [tokenizer.bos_token_id] + inst_begin_tokens + human_tokens + inst_end_tokens

# input_ids = human_tokens
input_ids = torch.tensor([input_ids], dtype=torch.long).cuda()

with torch.no_grad():
    outputs = model.generate(
        input_ids=input_ids, max_new_tokens=max_new_tokens, do_sample=True,
        top_p=top_p, temperature=temperature, repetition_penalty=repetition_penalty,
        eos_token_id=tokenizer.eos_token_id
    )
outputs = outputs.tolist()[0][len(input_ids[0]):]
response = tokenizer.decode(outputs)
response = response.strip().replace(tokenizer.eos_token, "").strip()
print("Chatbot：{}".format(response))

```

