---
language:
- ja
tags:
- japanese-stablelm
- causal-lm
pipeline_tag: text-generation
datasets:
- kunishou/hh-rlhf-49k-ja
- kunishou/databricks-dolly-15k-ja
- kunishou/oasst1-89k-ja
license:
- llama2
extra_gated_fields:
  Name: text
  Email: text
  Country: text
  Organization or Affiliation: text
  I allow Stability AI to contact me about information related to its models and research: checkbox
---

# Japanese-StableLM-Instruct-Beta-7B

![A cute robot wearing a kimono writes calligraphy with one single brush](./japanese-stablelm-robot.jpg)

> A cute robot wearing a kimono writes calligraphy with one single brush — [Stable Diffusion XL](https://clipdrop.co/stable-diffusion)

## Model Description

`japanese-stablelm-instruct-beta-7b` is a 7B-parameter decoder-only language model based on [japanese-stablelm-base-beta-7b](https://huggingface.co/stabilityai/japanese-stablelm-base-beta-7b) and further fine tuned on Databricks Dolly-15k, Anthropic HH, and other public data.
                                                                                                                                                 
This model is also available in a [larger 70b version](https://huggingface.co/stabilityai/japanese-stablelm-instruct-beta-70b), or a [faster version with a specialized tokenizer](https://huggingface.co/stabilityai/japanese-stablelm-instruct-ja_vocab-beta-7b).

## Usage

First install additional dependencies in [requirements.txt](./requirements.txt):

```sh
pip install -r requirements.txt
```

Then start generating text with `japanese-stablelm-instruct-beta-7b` by using the following code snippet:

```python
import torch
from transformers import AutoTokenizer, AutoModelForCausalLM

model_name = "stabilityai/japanese-stablelm-instruct-beta-7b"
tokenizer = AutoTokenizer.from_pretrained(model_name)

# The next line may need to be modified depending on the environment
model = AutoModelForCausalLM.from_pretrained(model_name, torch_dtype=torch.float16, low_cpu_mem_usage=True, device_map="auto")

def build_prompt(user_query, inputs):
    sys_msg = "<s>[INST] <<SYS>>\nあなたは役立つアシスタントです。\n<<SYS>>\n\n"
    p = sys_msg + user_query + "\n\n" + inputs + " [/INST] "
    return p

user_inputs = {
    "user_query": "与えられたことわざの意味を小学生でも分かるように教えてください。",
    "inputs": "情けは人のためならず"
}
prompt = build_prompt(**user_inputs)

input_ids = tokenizer.encode(
    prompt,
    add_special_tokens=True,
    return_tensors="pt"
)

# this is for reproducibility.
# feel free to change to get different result
seed = 23  
torch.manual_seed(seed)

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

We suggest playing with different generation config (`top_p`, `repetition_penalty` etc) to find the best setup for your tasks. For example, use higher temperature for roleplay task, lower temperature for reasoning.

## Model Details

* **Model type**: `japanese-stablelm-instruct-beta-7b` model is an auto-regressive language model based on the Llama2 transformer architecture.
* **Language(s)**: Japanese
* **License**: [Llama2 Community License](https://ai.meta.com/llama/license/).
* **Contact**: For questions and comments about the model, please join [Stable Community Japan](https://discord.gg/StableJP). For future announcements / information about Stability AI models, research, and events, please follow https://twitter.com/StabilityAI_JP.

## Training Dataset

The following datasets were used for the instruction training. Note these are Japanese translated versions of the original datasets, shared by [kunishou](https://huggingface.co/kunishou).

- [Anthropic HH-RLHF](https://huggingface.co/datasets/kunishou/hh-rlhf-49k-ja)
- [Databricks Dolly 15-k](https://huggingface.co/datasets/kunishou/databricks-dolly-15k-ja)
- [OpenAssistant Conversations Dataset](https://huggingface.co/datasets/kunishou/oasst1-89k-ja)

## Use and Limitations

### Intended Use

The model is intended to be used by all individuals as a foundation for application-specific fine-tuning without strict limitations on commercial use.

### Limitations and bias

The pre-training dataset may have contained offensive or inappropriate content even after applying data cleansing filters which can be reflected in the model generated text. We recommend users exercise reasonable caution when using these models in production systems. Do not use the model for any applications that may cause harm or distress to individuals or groups.

## Authors
This model was developed by the Research & Development team at Stability AI Japan, and the development was co-led by [Takuya Akiba](https://huggingface.co/iwiwi) and [Meng Lee](https://huggingface.co/leemeng). The members of the team are as follows:

- [Meng Lee](https://huggingface.co/leemeng)
- [Fujiki Nakamura](https://huggingface.co/fujiki)
- [Makoto Shing](https://huggingface.co/mkshing)
- [Paul McCann](https://huggingface.co/polm-stability)
- [Takuya Akiba](https://huggingface.co/iwiwi)
- [Naoki Orii](https://huggingface.co/mrorii)

## Acknowledgements

We thank Meta Research for releasing Llama 2 under an open license for others to build on.

We are grateful for the contributions of the EleutherAI Polyglot-JA team in helping us to collect a large amount of pre-training data in Japanese. Polyglot-JA members includes Hyunwoong Ko (Project Lead), Fujiki Nakamura (originally started this project when he commited to the Polyglot team), Yunho Mo, Minji Jung, KeunSeok Im, and Su-Kyeong Jang.

We are also appreciative of [AI Novelist/Sta (Bit192, Inc.)](https://ai-novel.com/index.php) and the numerous contributors from [Stable Community Japan](https://discord.gg/VPrcE475HB) for assisting us in gathering a large amount of high-quality Japanese textual data for model training.

