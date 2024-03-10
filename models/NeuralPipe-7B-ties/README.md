---
license: apache-2.0
tags:
- merge
model-index:
- name: NeuralPipe-7B-ties
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
      value: 67.92
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralPipe-7B-ties
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
      value: 86.04
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralPipe-7B-ties
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
      value: 64.24
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralPipe-7B-ties
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
      value: 61.37
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralPipe-7B-ties
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
      value: 80.19
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralPipe-7B-ties
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
      value: 69.52
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralPipe-7B-ties
      name: Open LLM Leaderboard
---

# NeuralPipe-7B-ties

This model is a merge of the following models made with [mergekit](https://github.com/cg123/mergekit):
 * [OpenPipe/mistral-ft-optimized-1218](https://huggingface.co/OpenPipe/mistral-ft-optimized-1218)
 * [mlabonne/NeuralHermes-2.5-Mistral-7B](https://huggingface.co/mlabonne/NeuralHermes-2.5-Mistral-7B)

## ⚡ Quantized models

Thanks to TheBloke for the quantized models:

* **GGUF**: https://huggingface.co/TheBloke/NeuralPipe-7B-ties-GGUF
* **AWQ**: https://huggingface.co/TheBloke/NeuralPipe-7B-ties-AWQ
* **GPTQ**: https://huggingface.co/TheBloke/NeuralPipe-7B-ties-GPTQ

## 🧩 Configuration

```yaml
models:
  - model: mistralai/Mistral-7B-v0.1
    # no parameters necessary for base model
  - model: OpenPipe/mistral-ft-optimized-1218
    parameters:
      density: 0.5
      weight: 0.5
  - model: mlabonne/NeuralHermes-2.5-Mistral-7B
    parameters:
      density: 0.5
      weight: 0.3
merge_method: ties
base_model: mistralai/Mistral-7B-v0.1
parameters:
  normalize: true
  int8_mask: true
dtype: float16
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_mlabonne__NeuralPipe-7B-ties)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |71.55|
|AI2 Reasoning Challenge (25-Shot)|67.92|
|HellaSwag (10-Shot)              |86.04|
|MMLU (5-Shot)                    |64.24|
|TruthfulQA (0-shot)              |61.37|
|Winogrande (5-shot)              |80.19|
|GSM8k (5-shot)                   |69.52|

