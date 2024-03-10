---
license: apache-2.0
tags:
- merge
model-index:
- name: NeuralOrca-7B-v1
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
      value: 65.27
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mrfakename/NeuralOrca-7B-v1
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
      value: 85.07
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mrfakename/NeuralOrca-7B-v1
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
      value: 63.68
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mrfakename/NeuralOrca-7B-v1
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
      value: 54.58
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mrfakename/NeuralOrca-7B-v1
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
      value: 78.77
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mrfakename/NeuralOrca-7B-v1
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
      value: 58.45
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mrfakename/NeuralOrca-7B-v1
      name: Open LLM Leaderboard
---
WARNING! This is a "Frankenmerge" - this model is untested!

# NeuralOrca 7B V1

[GGUF Models](https://huggingface.co/TheBloke/NeuralOrca-7B-v1-GGUF)

**By [mrfakename](https://twitter.com/realmrfakename)**

*Please note that this is an experimental model. We cannot guarantee model quality.*

This is the first (alpha) release of NeuralOrca. NeuralOrca is a merge of the following models:

* [mlabonne/NeuralHermes-2.5-Mistral-7B](https://huggingface.co/mlabonne/NeuralHermes-2.5-Mistral-7B) (This model is actually [OpenHermes 2.5](https://huggingface.co/teknium/OpenHermes-2.5-Mistral-7B) finetuned on Intel's [Neural Chat dataset](https://huggingface.co/datasets/Intel/neural-chat-dataset-v2) and uses the ChatML prompt format, weight: 1.0)
* [Open-Orca/Mistral-7B-OpenOrca](https://huggingface.co/Open-Orca/Mistral-7B-OpenOrca) (This model uses the ChatML prompt format, weight: 0.7)

## Prompt Format

We use the ChatML prompt format.

Example:
```
<|im_start|>system
You are NeuralOrca, a helpful AI assistant.
<|im_end|>
<|im_start|>user
How are you?<|im_end|>
<|im_start|>assistant
```

## Evaluations

Coming soon

## Context Length

The context length for this model is 8192 tokens (8K).

## License

You are responsible for your use of NeuralOrca.

This software is licensed under the Apache 2.0 license. If you want to use it for commercial use, it's probably fine but please contact me first.
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_mrfakename__NeuralOrca-7B-v1)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |67.64|
|AI2 Reasoning Challenge (25-Shot)|65.27|
|HellaSwag (10-Shot)              |85.07|
|MMLU (5-Shot)                    |63.68|
|TruthfulQA (0-shot)              |54.58|
|Winogrande (5-shot)              |78.77|
|GSM8k (5-shot)                   |58.45|

