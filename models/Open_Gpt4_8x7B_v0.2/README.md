---
license: cc-by-4.0
tags:
- merge
- moe
model-index:
- name: Open_Gpt4_8x7B_v0.2
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
      value: 68.69
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rombodawg/Open_Gpt4_8x7B_v0.2
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
      value: 86.16
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rombodawg/Open_Gpt4_8x7B_v0.2
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
      value: 72.07
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rombodawg/Open_Gpt4_8x7B_v0.2
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
      value: 71.92
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rombodawg/Open_Gpt4_8x7B_v0.2
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
      value: 83.58
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rombodawg/Open_Gpt4_8x7B_v0.2
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
      value: 59.14
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rombodawg/Open_Gpt4_8x7B_v0.2
      name: Open LLM Leaderboard
---
Open_Gpt4_v0.2

This is the un-quantized fp16 version for training and merging. If you want the quantized version for inference please refer to the repo bellow:

- https://huggingface.co/rombodawg/Open_Gpt4_8x7B_v0.2_q8_0_gguf

![image/jpeg](https://cdn-uploads.huggingface.co/production/uploads/642cc1c253e76b4c2286c58e/T7QKB0fKNHQvNqAjm8zrH.jpeg)

This model is a TIES merger of Mixtral-8x7B-Instruct-v0.1  and bagel-dpo-8x7b-v0.2 with MixtralOrochi8x7B being the Base model.



I was very impressed with MixtralOrochi8x7B performance and multifaceted usecases as it is already a merger of many usefull Mixtral models such as Mixtral instruct, 
Noromaid-v0.1-mixtral, openbuddy-mixtral and possibly other models that were not named. My goal was to expand the models capabilities and make it even more useful of a model, maybe even competitive with closed source models like Gpt-4. But for that more testing is required. I hope the community can help me determine if its deserving of its name. 😊

This is the second iteration of this model, using better models in the merger to improve performance (hopefully).

Base model: 

- https://huggingface.co/smelborp/MixtralOrochi8x7B

Merged models:

- https://huggingface.co/mistralai/Mixtral-8x7B-Instruct-v0.1

- https://huggingface.co/jondurbin/bagel-dpo-8x7b-v0.2


Instruct template: Alpaca


Merger config:
```yaml
models:
  - model: Mixtral-8x7B-Instruct-v0.1

    parameters:
      density: .5
      weight: 1
  - model: bagel-dpo-8x7b-v0.2
    parameters:
      density: .5
      weight: .7


merge_method: ties
base_model: MixtralOrochi8x7B
parameters:
  normalize: true
  int8_mask: true
dtype: float16



```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_rombodawg__Open_Gpt4_8x7B_v0.2)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |73.59|
|AI2 Reasoning Challenge (25-Shot)|68.69|
|HellaSwag (10-Shot)              |86.16|
|MMLU (5-Shot)                    |72.07|
|TruthfulQA (0-shot)              |71.92|
|Winogrande (5-shot)              |83.58|
|GSM8k (5-shot)                   |59.14|

