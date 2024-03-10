---
  license: cc-by-nc-4.0
---


# Mixtral MOE 2x10.7B

[One of Best MoE Model reviewd by reddit community](https://www.reddit.com/r/LocalLLaMA/comments/1916896/llm_comparisontest_confirm_leaderboard_big_news/)

MoE  of the following models :


* [kyujinpy/Sakura-SOLAR-Instruct](https://huggingface.co/kyujinpy/Sakura-SOLAR-Instruct)
* [jeonsworld/CarbonVillain-en-10.7B-v1](https://huggingface.co/jeonsworld/CarbonVillain-en-10.7B-v1)


* Local Test
* hf (pretrained=cloudyu/Mixtral_11Bx2_MoE_19B), gen_kwargs: (None), limit: None, num_fewshot: 10, batch_size: auto (32)
|  Tasks  |Version|Filter|n-shot| Metric |Value |   |Stderr|
|---------|-------|------|-----:|--------|-----:|---|-----:|
|hellaswag|Yaml   |none  |    10|acc     |0.7142|±  |0.0045|
|         |       |none  |    10|acc_norm|0.8819|±  |0.0032|


gpu code example

```
import torch
from transformers import AutoTokenizer, AutoModelForCausalLM
import math

## v2 models
model_path = "cloudyu/Mixtral_11Bx2_MoE_19B"

tokenizer = AutoTokenizer.from_pretrained(model_path, use_default_system_prompt=False)
model = AutoModelForCausalLM.from_pretrained(
    model_path, torch_dtype=torch.float32, device_map='auto',local_files_only=False, load_in_4bit=True
)
print(model)
prompt = input("please input prompt:")
while len(prompt) > 0:
  input_ids = tokenizer(prompt, return_tensors="pt").input_ids.to("cuda")

  generation_output = model.generate(
    input_ids=input_ids, max_new_tokens=500,repetition_penalty=1.2
  )
  print(tokenizer.decode(generation_output[0]))
  prompt = input("please input prompt:")
```

CPU example

```
import torch
from transformers import AutoTokenizer, AutoModelForCausalLM
import math

## v2 models
model_path = "cloudyu/Mixtral_11Bx2_MoE_19B"

tokenizer = AutoTokenizer.from_pretrained(model_path, use_default_system_prompt=False)
model = AutoModelForCausalLM.from_pretrained(
    model_path, torch_dtype=torch.float32, device_map='cpu',local_files_only=False
)
print(model)
prompt = input("please input prompt:")
while len(prompt) > 0:
  input_ids = tokenizer(prompt, return_tensors="pt").input_ids

  generation_output = model.generate(
    input_ids=input_ids, max_new_tokens=500,repetition_penalty=1.2
  )
  print(tokenizer.decode(generation_output[0]))
  prompt = input("please input prompt:")

```