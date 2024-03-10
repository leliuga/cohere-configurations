---
license: other
tags:
- yi
- moe
license_name: yi-license
license_link: https://huggingface.co/01-ai/Yi-34B-200K/blob/main/LICENSE
model-index:
- name: Helion-4x34B
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
      value: 69.71
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Helion-4x34B
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
      value: 85.28
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Helion-4x34B
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
      value: 77.33
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Helion-4x34B
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
      value: 63.91
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Helion-4x34B
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
      value: 84.37
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Helion-4x34B
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
      value: 72.25
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Helion-4x34B
      name: Open LLM Leaderboard
---
![image/jpeg](https://cdn-uploads.huggingface.co/production/uploads/6468ce47e134d050a58aa89c/GA28gWAezC9qhrTcwSfuA.jpeg)

# Helion-4x34B

This is the model for Helion-4x34B. I used [this repo](https://bit.ly/weyaxi-moe-repo) to make this MOE model.

# Prompt Template(s):

Since [bagel-dpo-34b-v0.2](https://huggingface.co/jondurbin/bagel-dpo-34b-v0.2) uses many prompt templates, you can utilize prompt templates provided by bagel and other expert's prompt templates.

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

### Human Asistant

```
Human: {user}

### Assistant: {asistant}
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
    negative_prompts: ["math", "reason", "mathematics", "solve", "count", "code", "python", "javascript", "programming", "algorithm"]

  - source_model: Nous-Hermes-2-Yi-34B
    positive_prompts: ["chat", "math", "reason", "mathematics", "solve", "count", "python", "javascript", "programming", "algorithm", "tell me", "assistant"]

  - source_model: SUS-Chat-34B
    positive_prompts: ["math", "reason", "mathematics", "solve", "count", "assistant"]

  - source_model: platypus-yi-34b
    positive_prompts: [""]
    negative_prompts: ["math", "reason", "mathematics", "solve", "count"]
```

# Quantizationed versions

Quantizationed versions of this model is available thanks to [TheBloke](https://hf.co/TheBloke).

##### GPTQ

- [TheBloke/Helion-4x34B-GPTQ](https://huggingface.co/TheBloke/Helion-4x34B-GPTQ)

##### GGUF

- [TheBloke/Helion-4x34B-GGUF](https://huggingface.co/TheBloke/Helion-4x34B-GGUF)

##### AWQ

- [TheBloke/Helion-4x34B-AWQ](https://huggingface.co/TheBloke/Helion-4x34B-AWQ)


# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Weyaxi__Helion-4x34B)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |75.48|
|AI2 Reasoning Challenge (25-Shot)|69.71|
|HellaSwag (10-Shot)              |85.28|
|MMLU (5-Shot)                    |77.33|
|TruthfulQA (0-shot)              |63.91|
|Winogrande (5-shot)              |84.37|
|GSM8k (5-shot)                   |72.25|

If you would like to support me:

[☕ Buy Me a Coffee](https://www.buymeacoffee.com/weyaxi)
