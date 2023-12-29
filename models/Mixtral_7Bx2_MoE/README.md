---
  license: cc-by-nc-4.0
---

# Mixtral MOE 2x7B



MoE  of the following models :

* [rwitz2/go-bruins-v2.1.1](https://huggingface.co/rwitz2/go-bruins-v2.1.1)
* [NurtureAI/neural-chat-7b-v3-16k](https://huggingface.co/NurtureAI/neural-chat-7b-v3-16k)
* [mncai/mistral-7b-dpo-v6](https://huggingface.co/mncai/mistral-7b-dpo-v6)




gpu code example

```
import torch
from transformers import AutoTokenizer, AutoModelForCausalLM
import math

## v2 models
model_path = "cloudyu/Mixtral_7Bx2_MoE"

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
model_path = "cloudyu/Mixtral_7Bx2_MoE"

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