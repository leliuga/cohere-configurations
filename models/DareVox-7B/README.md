---
license: apache-2.0
tags:
- merge
- mergekit
- lazymergekit
- teknium/OpenHermes-2.5-Mistral-7B
- abacusai/Slerp-CM-mist-dpo
- berkeley-nest/Starling-LM-7B-alpha
---

# DareVox-7B

DareVox-7B is a merge of the following models:
* [teknium/OpenHermes-2.5-Mistral-7B](https://huggingface.co/teknium/OpenHermes-2.5-Mistral-7B)
* [abacusai/Slerp-CM-mist-dpo](https://huggingface.co/abacusai/Slerp-CM-mist-dpo)
* [berkeley-nest/Starling-LM-7B-alpha](https://huggingface.co/berkeley-nest/Starling-LM-7B-alpha)

## 🧩 Configuration

```yaml
models:
  - model: mistralai/Mistral-7B-v0.1
    # No parameters necessary for base model
  - model: teknium/OpenHermes-2.5-Mistral-7B
    parameters:
      density: 0.53
      weight: 0.4
  - model: abacusai/Slerp-CM-mist-dpo
    parameters:
      density: 0.53
      weight: 0.3
  - model: berkeley-nest/Starling-LM-7B-alpha
    parameters:
      density: 0.5
      weight: 0.4
merge_method: dare_ties
base_model: mistralai/Mistral-7B-v0.1
parameters:
  int8_mask: true
dtype: bfloat16
```

## 💻 Usage

```python
!pip install -qU transformers accelerate

from transformers import AutoTokenizer
import transformers
import torch

model = "abideen/DareVox-7B"
messages = [{"role": "user", "content": "What is a large language model?"}]

tokenizer = AutoTokenizer.from_pretrained(model)
prompt = tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True)
pipeline = transformers.pipeline(
    "text-generation",
    model=model,
    torch_dtype=torch.float16,
    device_map="auto",
)

outputs = pipeline(prompt, max_new_tokens=256, do_sample=True, temperature=0.7, top_k=50, top_p=0.95)
print(outputs[0]["generated_text"])
```