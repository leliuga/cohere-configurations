---
license: other
tags:
- merge
license_name: deepseek
license_link: https://github.com/deepseek-ai/DeepSeek-Coder/blob/main/LICENSE-MODEL
model-index:
- name: Everyone-Coder-33b-Base
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
      value: 45.99
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rombodawg/Everyone-Coder-33b-Base
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
      value: 61.71
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rombodawg/Everyone-Coder-33b-Base
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
      value: 44.05
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rombodawg/Everyone-Coder-33b-Base
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
      value: 42.26
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rombodawg/Everyone-Coder-33b-Base
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
      value: 63.06
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rombodawg/Everyone-Coder-33b-Base
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
      value: 39.8
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rombodawg/Everyone-Coder-33b-Base
      name: Open LLM Leaderboard
---
Everyone-Coder-33b-Base


![image/jpeg](https://cdn-uploads.huggingface.co/production/uploads/642cc1c253e76b4c2286c58e/ECrHQnZnv8UM9GUCQtlWW.jpeg)

EveryoneLLM series of models made by the community, for the community. This is a coding specific model made using fine-tunes of deekseekcoder-33b-base.

______________________________________________________________________________________________________________
Im having trouble benchmarking this model because I suck at running llm benchmarks, but from hand testing running the model through https://edabit.com/challenge coding challenges vs up to date gpt-4. My model is hands down beating it in coding. 
______________________________________________________________________________________________________________
Ive recently noticed this model has trouble with end tokens so I made a custom prompt template for it. Made sure to add (Always end with "<|EOT|>") In addition to your system prompt and  (Always end your response with "<|EOT|>") at the end of the User message is the preset. Then add <|EOT|> as a custom stop string in your LM text generating interface. 

```
Always end with "<|EOT|>"

{System}

<|User|>

{User}. Always end your response with "<|EOT|>"

<|Assistant|>

{Assistant}
```

The models that were used in this merger were as follow:

- https://huggingface.co/deepseek-ai/deepseek-coder-33b-instruct

- https://huggingface.co/codefuse-ai/CodeFuse-DeepSeek-33B

- https://huggingface.co/WizardLM/WizardCoder-33B-V1.1

Thank you to the creators of the above ai models, they have full credit for the EveryoneLLM series of models. Without their hard work we wouldnt be able to achieve the great success we have in the open source community. 💗

You can find the write up for merging models here:

https://docs.google.com/document/d/1_vOftBnrk9NRk5h10UqrfJ5CDih9KBKL61yvrZtVWPE/edit?usp=sharing

Config for the merger can be found bellow:

```yaml
models:
  - model: WizardLM_WizardCoder-33B-V1.1
    parameters:
      density: 1
      weight: .5
  - model: codefuse-ai_CodeFuse-DeepSeek-33B
    parameters:
      density: 1
      weight: .5
merge_method: ties
base_model: deepseek-ai_deepseek-coder-33b-instruct
parameters:
  normalize: true
  int8_mask: true
dtype: float16

```

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_rombodawg__Everyone-Coder-33b-Base)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |49.48|
|AI2 Reasoning Challenge (25-Shot)|45.99|
|HellaSwag (10-Shot)              |61.71|
|MMLU (5-Shot)                    |44.05|
|TruthfulQA (0-shot)              |42.26|
|Winogrande (5-shot)              |63.06|
|GSM8k (5-shot)                   |39.80|

