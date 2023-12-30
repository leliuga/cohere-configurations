---
language:
- ja
tags:
- japanese-stablelm
- causal-lm
pipeline_tag: text-generation
datasets:
- wikipedia
- mc4
- cc100
- oscar-corpus/OSCAR-2301
- oscar-corpus/OSCAR-2201
- cerebras/SlimPajama-627B
license:
- llama2
extra_gated_fields:
  Name: text
  Email: text
  Country: text
  Organization or Affiliation: text
  I allow Stability AI to contact me about information related to its models and research: checkbox
---

# Japanese-StableLM-Base-Beta-70B

![A cute robot wearing a kimono writes calligraphy with one single brush](./japanese-stablelm-robot.jpg)

> A cute robot wearing a kimono writes calligraphy with one single brush — [Stable Diffusion XL](https://clipdrop.co/stable-diffusion)

## Model Description

`japanese-stablelm-base-beta-70b` is a 70B-parameter decoder-only language model based on [Llama-2-70b](https://huggingface.co/meta-llama/Llama-2-70b) that has been fine-tuned on a diverse collection of Japanese data, with the intent of maximizing downstream performance on Japanese language tasks.

For an instruction-following model, check [Japanese-StableLM-Instruct-Beta-70B](https://huggingface.co/stabilityai/japanese-stablelm-instruct-beta-70b). The base and instruct models are also available in smaller 7b sizes. For a model that has faster inference times, see [Japanese-StableLM-Base-JA_Vocab-Beta-7B](https://huggingface.co/stabilityai/japanese-stablelm-base-ja_vocab-beta-7b), or [the instruction-following version](https://huggingface.co/stabilityai/japanese-stablelm-instruct-ja_vocab-beta-7b).

## Usage

First install additional dependencies in [requirements.txt](./requirements.txt):

```sh
pip install -r requirements.txt
```

Then start generating text with `japanese-stablelm-base-beta-70b` by using the following code snippet:

```python
import torch
from transformers import AutoTokenizer, AutoModelForCausalLM

model_name = "stabilityai/japanese-stablelm-base-beta-70b"
tokenizer = AutoTokenizer.from_pretrained(model_name)

# The next line may need to be modified depending on the environment
model = AutoModelForCausalLM.from_pretrained(model_name, torch_dtype=torch.float16, low_cpu_mem_usage=True, device_map="auto")

prompt = """
AI で科学研究を加速するには、
""".strip()

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

* **Model type**: `japanese-stablelm-base-beta-70b` model is an auto-regressive language model based on the Llama2 transformer architecture.
* **Language(s)**: Japanese
* **License**: [Llama2 Community License](https://ai.meta.com/llama/license/).
* **Contact**: For questions and comments about the model, please join [Stable Community Japan](https://discord.gg/StableJP). For future announcements / information about Stability AI models, research, and events, please follow https://twitter.com/StabilityAI_JP.

## Training Dataset

Roughly 100B tokens from a mixture of the following corpora were used for continued pre-training.

- [Japanese/English Wikipedia](https://dumps.wikimedia.org/other/cirrussearch)
- [Japanese mc4](https://huggingface.co/datasets/mc4)
- [Japanese CC-100](http://data.statmt.org/cc-100/ja.txt.xz)
- [Japanese OSCAR](https://oscar-project.github.io/documentation/)
- [SlimPajama](https://huggingface.co/datasets/cerebras/SlimPajama-627B) (excluding the Books3 subset)

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

## How to cite
```
@misc{JapaneseStableLMBaseBeta70B, 
      url={[https://huggingface.co/stabilityai/japanese-stablelm-base-beta-70b](https://huggingface.co/stabilityai/japanese-stablelm-base-beta-70b)}, 
      title={Japanese StableLM Base Beta 70B}, 
      author={Lee, Meng and Nakamura, Fujiki and Shing, Makoto and McCann, Paul and Akiba, Takuya and Orii, Naoki}
}
```

