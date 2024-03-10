---
license: apache-2.0
tags:
- moe
- merge
- mergekit
- Solar Moe
- Solar
- Lumosia
model-index:
- name: Lumosia-MoE-4x10.7
  results:
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: AI2 Reasoning Challenge (25-Shot)
      type: ai2_arc
      config: ARC-Challenge
      split: test
      args:
        num_few_shot: 25
    metrics:
    - type: acc_norm
      value: 68.34
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Steelskull/Lumosia-MoE-4x10.7
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: HellaSwag (10-Shot)
      type: hellaswag
      split: validation
      args:
        num_few_shot: 10
    metrics:
    - type: acc_norm
      value: 87.13
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Steelskull/Lumosia-MoE-4x10.7
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: MMLU (5-Shot)
      type: cais/mmlu
      config: all
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 64.38
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Steelskull/Lumosia-MoE-4x10.7
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: TruthfulQA (0-shot)
      type: truthful_qa
      config: multiple_choice
      split: validation
      args:
        num_few_shot: 0
    metrics:
    - type: mc2
      value: 63.81
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Steelskull/Lumosia-MoE-4x10.7
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: Winogrande (5-shot)
      type: winogrande
      config: winogrande_xl
      split: validation
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 82.95
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Steelskull/Lumosia-MoE-4x10.7
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: GSM8k (5-shot)
      type: gsm8k
      config: main
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 51.02
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Steelskull/Lumosia-MoE-4x10.7
      name: Open LLM Leaderboard
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/64545af5ec40bbbd01242ca6/Qb88YeudOf7MYuGKTirXC.png)

# Lumosia-MoE-4x10.7

"Lumosia" was selected as its a MoE of Multiple SOLAR Merges so it really "Lights the way".... its 3am.

This is a very experimantal model. its a MoE of all good performing Solar models (based off of personal experiance not open leaderboard), 
The models goal was to make a good all rounder, in chat/logic/rp

Why? Dunno whated to see what would happen

context is 4k but coherent up to 16k

Quants by @thebloke  (thank you)

https://huggingface.co/TheBloke/Lumosia-MoE-4x10.7-GGUF

https://huggingface.co/TheBloke/Lumosia-MoE-4x10.7-GPTQ

Update: (Done)
Lumosia v1.5 has been uploaded.

Update 2:

A Lumosia Personality tavern card has been added

Come join the Discord:
[ConvexAI](https://discord.gg/yYqmNmg7Wj)


Template:
```
### System:

### USER:{prompt}

### Assistant:
```


Settings:
```
Temp: 1.0
min-p: 0.02-0.1
```

Lumosia-MoE-4x10.7 is a Mixure of Experts (MoE) made with the following models:
* [DopeorNope/SOLARC-M-10.7B](https://huggingface.co/DopeorNope/SOLARC-M-10.7B)
* [maywell/PiVoT-10.7B-Mistral-v0.2-RP](https://huggingface.co/maywell/PiVoT-10.7B-Mistral-v0.2-RP)
* [kyujinpy/Sakura-SOLAR-Instruct](https://huggingface.co/kyujinpy/Sakura-SOLAR-Instruct)
* [jeonsworld/CarbonVillain-en-10.7B-v1](https://huggingface.co/jeonsworld/CarbonVillain-en-10.7B-v1)

## Evals:

* Avg: 69.61
* ARC: 68.34
* HellaSwag: 87.13
* MMLU: 64.38
* T-QA: 63.81
* Winogrande: 82.95
* GSM8K: 51.02

## Examples:
```
Example 1:

User:
Hello lumi, what is the meaning of life
in around 100 words.

Lumosia:
The meaning of life can vary greatly
depending on individual beliefs,
values, and experiences. One common
perspective suggests that the primary
purpose of life is to seek happiness,
fulfillment, and personal growth
through relationships
accomplishments, and spiritual
awareness. Living with integrity.
compassion, and gratitude while
making a positive impact on others
and the world around us can also
contribute to finding meaning in life
UItimately, the meaning of life is
subjective and unique to each
individual, who must define it for
themselves through exploration
reflection, and personal choice.

```
```
Example 2:

User:
Explain a MoE (Mixture of Experts) in
around 100 words

Lumosia:
A MoE, short for Mixture of Experts, is
a machine learning technique utilized
in ensemble learning. It involves
combining multiple specialist models
(experts) to optimize predictive
performance. Each expert focuses on
a subset of the input features or data
domains, making them specialized in
particular areas. Through a weighting
mechanism, like Gaussian Mixture
Model, the predictions of experts are
blended to produce a final result. MoE
models demonstrate effectiveness in
handling complex or ambiguous
inputs where a single model might
struggle. They are commonly used in
natural language processing.
computer vision, and speech synthesis.
```

## 🧩 Configuration

```
yamlbase_model: DopeorNope/SOLARC-M-10.7B
gate_mode: hidden
dtype: bfloat16
experts:
  - source_model: DopeorNope/SOLARC-M-10.7B
    positive_prompts: [""]
  - source_model: maywell/PiVoT-10.7B-Mistral-v0.2-RP
    positive_prompts: [""]
  - source_model: kyujinpy/Sakura-SOLAR-Instruct
    positive_prompts: [""]
  - source_model: jeonsworld/CarbonVillain-en-10.7B-v1
    positive_prompts: [""]
```

## 💻 Usage

```
python
!pip install -qU transformers bitsandbytes accelerate

from transformers import AutoTokenizer
import transformers
import torch

model = "Steelskull/Lumosia-MoE-4x10.7"

tokenizer = AutoTokenizer.from_pretrained(model)
pipeline = transformers.pipeline(
    "text-generation",
    model=model,
    model_kwargs={"torch_dtype": torch.float16, "load_in_4bit": True},
)

messages = [{"role": "user", "content": "Explain what a Mixture of Experts is in less than 100 words."}]
prompt = pipeline.tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True)
outputs = pipeline(prompt, max_new_tokens=256, do_sample=True, temperature=0.7, top_k=50, top_p=0.95)
print(outputs[0]["generated_text"])
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Steelskull__Lumosia-MoE-4x10.7)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |69.61|
|AI2 Reasoning Challenge (25-Shot)|68.34|
|HellaSwag (10-Shot)              |87.13|
|MMLU (5-Shot)                    |64.38|
|TruthfulQA (0-shot)              |63.81|
|Winogrande (5-shot)              |82.95|
|GSM8k (5-shot)                   |51.02|

