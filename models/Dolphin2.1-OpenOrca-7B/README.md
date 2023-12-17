---
model-index:
- name: Dolphin2.1-OpenOrca-7B
  results:


  # AI2 Reasoning Challenge (25-Shot)
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
         name: normalized accuracy
         value: 63.91
    source:
      name: Open LLM Leaderboard
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Dolphin2.1-OpenOrca-7B
      

  # HellaSwag (10-shot)
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
         name: normalized accuracy
         value: 84.26
    source:
      name: Open LLM Leaderboard
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Dolphin2.1-OpenOrca-7B


  # MMLU (5-Shot)
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
         name: accuracy
         value: 62.66
    source:
      name: Open LLM Leaderboard
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Dolphin2.1-OpenOrca-7B


  # TruthfulQA (0-shot)
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
         value: 53.84
    source:
      name: Open LLM Leaderboard
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Dolphin2.1-OpenOrca-7B


  # Winogrande (5-shot)
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
         name: accuracy
         value: 78.22
    source:
      name: Open LLM Leaderboard
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Dolphin2.1-OpenOrca-7B


  # GSM8k (5-shot)
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
         name: accuracy
         value: 19.94
    source:
      name: Open LLM Leaderboard
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Dolphin2.1-OpenOrca-7B
---

<a href="https://www.buymeacoffee.com/PulsarAI" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>

Merge of [ehartford/dolphin-2.1-mistral-7b](https://huggingface.co/ehartford/dolphin-2.1-mistral-7b) and [Open-Orca/Mistral-7B-OpenOrca](https://huggingface.co/Open-Orca/Mistral-7B-OpenOrca) using ties merge.

### *Weights*

- [ehartford/dolphin-2.1-mistral-7b](https://huggingface.co/ehartford/dolphin-2.1-mistral-7b): 0.5

- [Open-Orca/Mistral-7B-OpenOrca](https://huggingface.co/Open-Orca/Mistral-7B-OpenOrca): 0.3

### *Density*

- [ehartford/dolphin-2.1-mistral-7b](https://huggingface.co/ehartford/dolphin-2.1-mistral-7b): 0.5

- [Open-Orca/Mistral-7B-OpenOrca](https://huggingface.co/Open-Orca/Mistral-7B-OpenOrca): 0.5

# Quantizationed versions

Quantizationed versions of this model is available thanks to [TheBloke](https://hf.co/TheBloke).

##### GPTQ

- [TheBloke/Dolphin2.1-OpenOrca-7B-GPTQ](https://huggingface.co/TheBloke/Dolphin2.1-OpenOrca-7B-GPTQ)

##### GGUF

- [TheBloke/Dolphin2.1-OpenOrca-7B-GGUF](https://huggingface.co/TheBloke/Dolphin2.1-OpenOrca-7B-GGUF)

##### AWQ

- [TheBloke/Dolphin2.1-OpenOrca-7B-AWQ](https://huggingface.co/TheBloke/Dolphin2.1-OpenOrca-7B-AWQ)

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Weyaxi__Dolphin2.1-OpenOrca-7B)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 60.47   |
| ARC (25-shot)         | 63.91          |
| HellaSwag (10-shot)   | 84.26    |
| MMLU (5-shot)         | 62.66         |
| TruthfulQA (0-shot)   | 53.84   |
| Winogrande (5-shot)   | 78.22   |
| GSM8K (5-shot)        | 19.94        |
