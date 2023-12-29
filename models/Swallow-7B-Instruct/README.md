---
language:
  - en
  - ja
library_name: transformers
pipeline_tag: text-generation
license: llama2
model_type: llama
---

# Swallow

Our Swallow model has undergone continuous pre-training from the Llama 2 family, primarily with the addition of Japanese language data. The tuned versions use supervised fine-tuning (SFT). 
Links to other models can be found in the index.

## Swallow Model Index
|Model|Swallow-hf|Swallow-instruct-hf|
|---|---|---|
|7B| [Link](https://huggingface.co/tokyotech-llm/Swallow-7b-hf) | [Link](https://huggingface.co/tokyotech-llm/Swallow-7b-instruct-hf)|
|13B| [Link](https://huggingface.co/tokyotech-llm/Swallow-13b-hf) | [Link](https://huggingface.co/tokyotech-llm/Swallow-13b-instruct-hf)|
|70B| [Link](https://huggingface.co/tokyotech-llm/Swallow-70b-hf) | [Link](https://huggingface.co/tokyotech-llm/Swallow-70b-instruct-hf)|


![logo](./logo.png)

This repository provides large language models developed by [TokyoTech-LLM](https://tokyotech-llm.github.io/).
Read our [blog post](https://zenn.dev/tokyotech_lm/articles/d6cb3a8fdfc907) or our paper (preprint coming soon) for more details!


## Model Details

* **Model type**: Please refer to LLaMA-2 technical report for details on the model architecture. 
* **Language(s)**: Japanese English
* **Library**: [Megatron-LM](https://github.com/rioyokotalab/Megatron-Llama2) 
* **Tokenizer**: This model employs a tokenizer that features a broadened vocabulary based on Japanese data. This allows for a more efficient representation of text using fewer tokens, leading to a notably faster inference process.
* **Contact**: swallow[at]nlp.c.titech.ac.jp 

## Base Model Performance

### Japanese version

|Model|Size|JCommonsenseQA|JEMHopQA|NIILC|JSQuAD|XL-Sum|MGSM|WMT20-en-ja|WMT20-ja-en|
|---|---|---|---|---|---|---|---|---|---|
|   |   |4-shot|4-shot|4-shot|4-shot|1-shot|4-shot|4-shot|4-shot|
|Llama 2|7B|0.3852|0.4240|0.3410|0.7917|0.1905|0.0760|0.1783|0.1738|
|Swallow|7B|0.4808|0.5078|0.5968|0.8573|0.1830|0.1240|0.2510|0.1511|
|Llama 2|13B|0.6997|0.4415|0.4170|0.8533|0.2139|0.1320|0.2146|0.1982|
|Swallow|13B|0.7837|0.5063|0.6398|0.9005|0.2168|0.2040|0.2720|0.1771|
|Llama 2|70B|0.8686|0.4656|0.5256|0.9080|**0.2361**|0.3560|0.2643|**0.2398**|
|Swallow|70B|**0.9348**|**0.6290**|**0.6960**|**0.9176**|0.2266|**0.4840**|**0.3043**|0.2298|

## Usage

First install additional dependencies in [requirements.txt](./requirements.txt):

```sh
pip install -r requirements.txt
```

### Use the instruct model

```python
import torch
from transformers import AutoTokenizer, AutoModelForCausalLM

model_name = "tokyotech-llm/Swallow-7b-instruct-hf"

tokenizer = AutoTokenizer.from_pretrained(model_name)
model = AutoModelForCausalLM.from_pretrained(model_name, torch_dtype=torch.bfloat16, low_cpu_mem_usage=True, device_map="auto")


PROMPT_DICT = {
    "prompt_input": (
        "以下に、あるタスクを説明する指示があり、それに付随する入力が更なる文脈を提供しています。"
        "リクエストを適切に完了するための回答を記述してください。\n\n"
        "### 指示:\n{instruction}\n\n### 入力:\n{input}\n\n### 応答:"

    ),
    "prompt_no_input": (
        "以下に、あるタスクを説明する指示があります。"
        "リクエストを適切に完了するための回答を記述してください。\n\n"
        "### 指示:\n{instruction}\n\n### 応答:"
    ),
}

def create_prompt(instruction, input=None):
    """
    Generates a prompt based on the given instruction and an optional input.
    If input is provided, it uses the 'prompt_input' template from PROMPT_DICT.
    If no input is provided, it uses the 'prompt_no_input' template.

    Args:
        instruction (str): The instruction describing the task.
        input (str, optional): Additional input providing context for the task. Default is None.

    Returns:
        str: The generated prompt.
    """
    if input:
        # Use the 'prompt_input' template when additional input is provided
        return PROMPT_DICT["prompt_input"].format(instruction=instruction, input=input)
    else:
        # Use the 'prompt_no_input' template when no additional input is provided
        return PROMPT_DICT["prompt_no_input"].format(instruction=instruction)

# Example usage
instruction_example = "以下のトピックに関する詳細な情報を提供してください。"
input_example = "東京工業大学の主なキャンパスについて教えてください"
prompt = create_prompt(instruction_example, input_example)

input_ids = tokenizer.encode(
    prompt,
    add_special_tokens=False,
    return_tensors="pt"
)

tokens = model.generate(
    input_ids.to(device=model.device),
    max_new_tokens=128,
    temperature=0.99,
    top_p=0.95,
    do_sample=True,
)

out = tokenizer.decode(tokens[0], skip_special_tokens=True)
print(out)

```

### Use the base model

```python
import torch
from transformers import AutoTokenizer, AutoModelForCausalLM

model_name = "tokyotech-llm/Swallow-7b-hf"

tokenizer = AutoTokenizer.from_pretrained(model_name)
model = AutoModelForCausalLM.from_pretrained(model_name, torch_dtype=torch.bfloat16, device_map="auto")

prompt = "東京工業大学の主なキャンパスは、"
input_ids = tokenizer.encode(
    prompt,
    add_special_tokens=False,
    return_tensors="pt"
)
tokens = model.generate(
    input_ids.to(device=model.device),
    max_new_tokens=128,
    temperature=0.99,
    top_p=0.95,
    do_sample=True,
)

out = tokenizer.decode(tokens[0], skip_special_tokens=True)
print(out)
```

## Training Datasets

### Continual Pre-Training
The following datasets were used for continual pre-training.

- [Japanese Wikipedia](https://dumps.wikimedia.org/other/cirrussearch)
- [RefinedWeb](https://huggingface.co/datasets/tiiuae/falcon-refinedweb)
- Swallow Corpus
- [The Pile](https://huggingface.co/datasets/EleutherAI/pile)


### Instruction Tuning

The following datasets were used for the instruction tuning. 

- [Anthropic HH-RLHF](https://huggingface.co/datasets/kunishou/hh-rlhf-49k-ja)
- [Databricks Dolly 15-k](https://huggingface.co/datasets/kunishou/databricks-dolly-15k-ja)
- [OpenAssistant Conversations Dataset](https://huggingface.co/datasets/kunishou/oasst1-89k-ja)
 
## Risks and Limitations

The models released here are still in the early stages of our research and development and have not been tuned to ensure outputs align with human intent and safety considerations.

## Acknowledgements

We thank Meta Research for releasing Llama 2 under an open license for others to build on.

Our project is supported by the [ABCI Large-scale Language Model Building Support Program](https://abci.ai/en/link/llm_support_program.html) of the National Institute of Advanced Industrial Science and Technology. 

## License

Llama 2 is licensed under the LLAMA 2 Community License, Copyright © Meta Platforms, Inc. All Rights Reserved.

## Authors

Here are the team members:
- From [Okazaki Laboratory](https://www.nlp.c.titech.ac.jp/index.en.html), the following members:
  - [Naoaki Okazaki](https://www.chokkan.org/index.ja.html)
  - [Sakae Mizuki](https://s-mizuki-nlp.github.io/)
  - [Hiroki Iida](https://meshidenn.github.io/)
  - [Mengsay Loem](https://loem-ms.github.io/)
  - [Shota Hirai](https://huggingface.co/Kotemo428)
  - [Kakeru Hattori](https://aya-se.vercel.app/)
  - [Masanari Ohi](https://twitter.com/stjohn2007)
- From [YOKOTA Laboratory](https://www.rio.gsic.titech.ac.jp/en/index.html), the following members:
  - [Rio Yokota](https://twitter.com/rioyokota)
  - [Kazuki Fujii](https://twitter.com/okoge_kaz)
  - [Taishi Nakamura](https://twitter.com/Setuna7777_2)
