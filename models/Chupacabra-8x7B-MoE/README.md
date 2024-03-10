---
license: apache-2.0
tags:
- moe
model-index:
- name: Chupacabra-8x7B-MoE
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
      value: 68.77
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=perlthoughts/Chupacabra-8x7B-MoE
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
      value: 86.11
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=perlthoughts/Chupacabra-8x7B-MoE
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
      value: 63.86
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=perlthoughts/Chupacabra-8x7B-MoE
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
      value: 63.5
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=perlthoughts/Chupacabra-8x7B-MoE
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
      value: 80.51
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=perlthoughts/Chupacabra-8x7B-MoE
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
      value: 59.67
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=perlthoughts/Chupacabra-8x7B-MoE
      name: Open LLM Leaderboard
---

# Chupacabra-8x7B-experts 
# Mixtral-8x7b-experts merge, of Chupacabra, OpenChat, Falkor, and Starling-LM

<p><img src="https://huggingface.co/perlthoughts/Chupacabra-7B/resolve/main/chupacabra7b%202.png" width=330></p>

### Model Description

Special thanks to @cgt123 for his great work on https://github.com/cg123/mergekit.

This was made using the 'mixtral' branch on the mergekit repo.

#OneManArmy

### More Info

- **Developed by:** Ray Hernandez
- **Model type:** Mistral
- **Language(s) (NLP):** English
- **License:** Apache 2.0


# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_perlthoughts__Chupacabra-8x7B-MoE)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |70.40|
|AI2 Reasoning Challenge (25-Shot)|68.77|
|HellaSwag (10-Shot)              |86.11|
|MMLU (5-Shot)                    |63.86|
|TruthfulQA (0-shot)              |63.50|
|Winogrande (5-shot)              |80.51|
|GSM8k (5-shot)                   |59.67|

