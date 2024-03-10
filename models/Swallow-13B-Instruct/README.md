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
|7B-Plus| [Link](https://huggingface.co/tokyotech-llm/Swallow-7b-plus-hf) | Coming Soon |
|13B| [Link](https://huggingface.co/tokyotech-llm/Swallow-13b-hf) | [Link](https://huggingface.co/tokyotech-llm/Swallow-13b-instruct-hf)|
|70B| [Link](https://huggingface.co/tokyotech-llm/Swallow-70b-hf) | [Link](https://huggingface.co/tokyotech-llm/Swallow-70b-instruct-hf)|

## Swallow Model Index NVE (No Vocabulary Expansion)
|Model|Swallow-NVE-hf|Swallow-NVE-instruct-hf|
|---|---|---|
|7B| [Link](https://huggingface.co/tokyotech-llm/Swallow-7b-NVE-hf) | [Link](https://huggingface.co/tokyotech-llm/Swallow-7b-NVE-instruct-hf)|
|13B| [Link](https://huggingface.co/tokyotech-llm/Swallow-13b-NVE-hf) | Coming Soon |
|70B| [Link](https://huggingface.co/tokyotech-llm/Swallow-70b-NVE-hf) | [Link](https://huggingface.co/tokyotech-llm/Swallow-70b-NVE-instruct-hf)|

We released the 7B and 70B models without vocabulary expansion on January 26th, 2024. The 13B model was released on February 4th, 2024, and its instruction-tuned version is coming soon. Swallow-7B-Plus is a model that has been trained with a larger number of Japanese tokens compared to Swallow-7B and its release date is March 2nd, 2024.


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
| Llama 2 | 7B | 0.3852 | 0.4240 | 0.3410 | 0.7917 | 0.1905 | 0.0760 | 0.1783 | 0.1738 |
| Swallow | 7B | 0.4808 | 0.5078 | 0.5968 | 0.8573 | 0.1830 | 0.1240 | 0.2510 | 0.1511 |
| Swallow-Plus | 7B | 0.5478 | 0.5493 | 0.6030 | 0.8544 | 0.1806 | 0.1360 | 0.2568 | 0.1441 |
| Swallow-NVE | 7B | 0.5433 | 0.5425 | 0.5729 | 0.8684 | 0.2117 | 0.1200 | 0.2405 | 0.1512 |
| Llama 2 | 13B | 0.6997 | 0.4415 | 0.4170 | 0.8533 | 0.2139 | 0.1320 | 0.2146 | 0.1982 |
| Swallow | 13B | 0.7837 | 0.5063 | 0.6398 | 0.9005 | 0.2168 | 0.2040 | 0.2720 | 0.1771 |
| Swallow-NVE | 13B | 0.7712 | 0.5438 | 0.6351 | 0.9030 | 0.2294 | 0.2120 | 0.2735 | 0.1817 |
| Llama 2 | 70B | 0.8686 | 0.4656 | 0.5256 | 0.9080 | 0.2361 | 0.3560 | 0.2643 | **0.2398** |
| Swallow | 70B | 0.9348 | **0.6290** | 0.6960 | 0.9176 | 0.2266 | **0.4840** | **0.3043** | 0.2298 |
| Swallow-NVE | 70B | **0.9410** | 0.5759 | **0.7024** | **0.9254** | **0.2758** | 0.4720 | 0.3042 | 0.2322 |
### English version

|Model|Size|OpenBookQA|TriviaQA|HellaSwag|SQuAD2.0|XWINO|GSM8K|
|---|---|---|---|---|---|---|---|
|   |   |8-shot|8-shot|8-shot|8-shot|8-shot|8-shot|
| Llama 2 | 7B    | 0.3580     | 0.6265   | 0.5860    | 0.3207   | 0.9049 | 0.1410 |
| Swallow | 7B    | 0.3180     | 0.4836   | 0.5308    | 0.3125   | 0.8817 | 0.1130 |
| Swallow-Plus | 7B | 0.3280     | 0.4558   | 0.5259    | 0.3134   | 0.8929 | 0.1061 |
| Swallow-NVE | 7B | 0.3180     | 0.5079   | 0.5329    | 0.2919   | 0.8817 | 0.0986 |
| Llama 2 | 13B   | 0.3760     | 0.7255   | 0.6148    | 0.3681   | 0.9140 | 0.2403 |
| Swallow | 13B   | 0.3500     | 0.5852   | 0.5660    | 0.3406   | 0.9075 | 0.2039 |
| Swallow-NVE | 13B | 0.3460     | 0.6025   | 0.5700    | 0.3478   | 0.9006 | 0.1751 |
| Llama 2 | 70B   | **0.4280** | **0.8239** | **0.6742** | **0.3770** | **0.9290** | **0.5284** |
| Swallow | 70B   | 0.4220     | 0.7756   | 0.6458    | 0.3745   | 0.9204 | 0.4867 |
| Swallow-NVE | 70B | 0.4240     | 0.7817   | 0.6439    | 0.3451   | 0.9256 | 0.4943 |

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
