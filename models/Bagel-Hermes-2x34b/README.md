---
license: other
tags:
- yi
- moe
license_name: yi-license
license_link: https://huggingface.co/01-ai/Yi-34B-200K/blob/main/LICENSE
model-index:
- name: Bagel-Hermes-2x34b
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
      value: 69.8
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Bagel-Hermes-2x34b
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
      value: 85.26
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Bagel-Hermes-2x34b
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
      value: 77.24
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Bagel-Hermes-2x34b
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
      value: 64.82
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Bagel-Hermes-2x34b
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
      value: 84.77
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Bagel-Hermes-2x34b
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
      value: 68.69
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Bagel-Hermes-2x34b
      name: Open LLM Leaderboard
---
![image/jpeg](https://cdn-uploads.huggingface.co/production/uploads/6468ce47e134d050a58aa89c/mLH2E0dk9On_LcFX9yhuS.jpeg)

# Bagel-Hermes-2x34B

This is the model for Bagel-Hermes-2x34B. I used [this repo](https://bit.ly/weyaxi-moe-repo) to make this MOE model.

# Prompt Template(s):

Since [bagel-dpo-34b-v0.2](https://huggingface.co/jondurbin/bagel-dpo-34b-v0.2) uses many prompt templates, and [Nous-Hermes-2-Yi-34B](https://huggingface.co/NousResearch/Nous-Hermes-2-Yi-34B) uses ChatML, you can utilize ChatML and other prompt templates provided by bagel.

**Note:** I currently do not know which prompt template is best.

### ChatML:

```
<|im_start|>system
{system}<|im_end|>
<|im_start|>user
{user}<|im_end|>
<|im_start|>assistant
{asistant}<|im_end|>
```

### Alpaca (sort of)

```
Below is an instruction that describes a task.  Write a response that appropriately completes the request.

### Instruction:
{system}
{instruction}

### Response:
```
### Vicuna

```
{system}
USER: {instruction}
ASSISTANT: 
```

Visit [bagel-dpo-34b-v0.2](https://huggingface.co/jondurbin/bagel-dpo-34b-v0.2) to try more prompt templates. 

# Yaml Config to reproduce

```yaml
base_model: nontoxic-bagel-34b-v0.2
gate_mode: hidden
dtype: bfloat16

experts:
  - source_model: bagel-dpo-34b-v0.2
    positive_prompts: ["question answering", "Q:", science", "biology", "chemistry", "physics"]

  - source_model: Nous-Hermes-2-Yi-34B
    positive_prompts: ["chat", "math", "reason", "mathematics", "solve", "count", "python", "javascript", "programming", "algorithm", "tell me", "assistant"]
```

# Quantizationed versions

Quantizationed versions of this model is available thanks to [TheBloke](https://hf.co/TheBloke).

##### GPTQ

- [TheBloke/Bagel-Hermes-2x34B-GPTQ](https://huggingface.co/TheBloke/Bagel-Hermes-2x34B-GPTQ)

##### GGUF

- [TheBloke/Bagel-Hermes-2x34B-GGUF](https://huggingface.co/TheBloke/Bagel-Hermes-2x34B-GGUF)

##### AWQ

- [TheBloke/Bagel-Hermes-2x34B-AWQ](https://huggingface.co/TheBloke/Bagel-Hermes-2x34B-AWQ)


# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Weyaxi__Bagel-Hermes-2x34b)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |75.10|
|AI2 Reasoning Challenge (25-Shot)|69.80|
|HellaSwag (10-Shot)              |85.26|
|MMLU (5-Shot)                    |77.24|
|TruthfulQA (0-shot)              |64.82|
|Winogrande (5-shot)              |84.77|
|GSM8k (5-shot)                   |68.69|

If you would like to support me:

[☕ Buy Me a Coffee](https://www.buymeacoffee.com/weyaxi)