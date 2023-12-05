---
license: mit
language:
- ja
library_name: transformers
pipeline_tag: text-generation
tags:
- japanese
- llama-2
---

# stockmark/stockmark-13b

Stockmark-13b is a 13 billion parameter LLM pretrained from scratch based on Japanese corpus of about 220B tokens. This model is developed by [Stockmark Inc.](https://stockmark.co.jp/)

Please see our [blog](https://tech.stockmark.co.jp/blog/202310_stockmark_13b/) for more details.

This project is supported by [AWS LLM development support program](https://aws.amazon.com/jp/local/llm-development-support-program/).

We also provide [stockmark-13b-instruct](https://huggingface.co/stockmark/stockmark-13b-instruct), which is the instruction tuned version of stockmark-13b.

## How to use

```python
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer

# For A100 or H100 GPU
model = AutoModelForCausalLM.from_pretrained("stockmark/stockmark-13b", device_map="auto", torch_dtype=torch.bfloat16)

# If you use a T4 or V100 GPU, please load a model in 8 bit with the below code.
# To do so, you need to install `bitsandbytes` via `pip install bitsandbytes`.
# model = AutoModelForCausalLM.from_pretrained("stockmark/stockmark-13b", device_map={"": 0}, load_in_8bit=True)

tokenizer = AutoTokenizer.from_pretrained("stockmark/stockmark-13b")

inputs = tokenizer("自然言語処理とは", return_tensors="pt").to(model.device)
with torch.no_grad():
    tokens = model.generate(
        **inputs,
        max_new_tokens=128,
        do_sample=True,
        temperature=0.7
    )
    
output = tokenizer.decode(tokens[0], skip_special_tokens=True)
print(output)
```

## Examples:

- LoRA tuning: https://huggingface.co/stockmark/stockmark-13b/blob/main/notebooks/LoRA.ipynb

## Training dataset

We have used Japanese corpus of total of about 220 billion tokens.

|corpus|tokens after preprocessing|
|:---:|:---:|
|Stockmark Web Corpus (This dataset will not be released)|9.1 billion|
|Patent|34.8 billion|
|Wikipedia|1.0 billion|
|CC100|10.9 billion|
|mC4|53.2 billion|
|CommonCrawl (snapshot: 2023-23, 2022-49, 2022-21, 2021-21)|112.9 billion|


## Accelerator and Library
- Accelerator: AWS Trainium
  - https://aws.amazon.com/machine-learning/trainium/
- Library for distributed training: neuronx-nemo-megatron
  - https://github.com/aws-neuron/neuronx-nemo-megatron

## License
[MIT](https://opensource.org/licenses/MIT)

## Developed by
[Stockmark Inc.](https://stockmark.co.jp/)

## Author
[Takahiro Omi](https://huggingface.co/omitakahiro)
