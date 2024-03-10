---
language:
- en
license: mit
tags:
- moe
pipeline_tag: text-generation
model-index:
- name: FusionNet_34Bx2_MoE
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
      value: 72.95
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=TomGrc/FusionNet_34Bx2_MoE
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
      value: 86.22
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=TomGrc/FusionNet_34Bx2_MoE
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
      value: 77.05
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=TomGrc/FusionNet_34Bx2_MoE
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
      value: 71.31
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=TomGrc/FusionNet_34Bx2_MoE
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
      value: 83.98
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=TomGrc/FusionNet_34Bx2_MoE
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
      value: 70.89
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=TomGrc/FusionNet_34Bx2_MoE
      name: Open LLM Leaderboard
---
# FusionNet_34Bx2_MoE
Fine-tuned model on English language using MoE method.
## Model description
The FusionNet_34Bx2_MoE is a model to experiment with the MoE method, which could significantly increase the performance of the original model. The FusionNet_34Bx2_MoE has 60.8B parameters, and this model is fine-tuned. Enjoy!
## Usage
```python
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer
tokenizer = AutoTokenizer.from_pretrained("TomGrc/FusionNet_34Bx2_MoE")
model = AutoModelForCausalLM.from_pretrained("TomGrc/FusionNet_34Bx2_MoE")
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_TomGrc__FusionNet_34Bx2_MoE)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |77.07|
|AI2 Reasoning Challenge (25-Shot)|72.95|
|HellaSwag (10-Shot)              |86.22|
|MMLU (5-Shot)                    |77.05|
|TruthfulQA (0-shot)              |71.31|
|Winogrande (5-shot)              |83.98|
|GSM8k (5-shot)                   |70.89|

