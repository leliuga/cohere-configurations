---
language:
- en
license: other
tags:
- uncensored
datasets:
- ehartford/wizard_vicuna_70k_unfiltered
model-index:
- name: Wizard-Vicuna-30B-Uncensored
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
      value: 62.12
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/Wizard-Vicuna-30B-Uncensored
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
      value: 83.45
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/Wizard-Vicuna-30B-Uncensored
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
      value: 58.24
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/Wizard-Vicuna-30B-Uncensored
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
      value: 50.81
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/Wizard-Vicuna-30B-Uncensored
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
      value: 78.45
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/Wizard-Vicuna-30B-Uncensored
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
      value: 14.25
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/Wizard-Vicuna-30B-Uncensored
      name: Open LLM Leaderboard
---

This is [wizard-vicuna-13b](https://huggingface.co/junelee/wizard-vicuna-13b) trained with a subset of the dataset - responses that contained alignment / moralizing were removed. The intent is to train a WizardLM that doesn't have alignment built-in, so that alignment (of any sort) can be added separately with for example with a RLHF LoRA.

Join Our Discord! https://discord.gg/cognitivecomputations 

Shout out to the open source AI/ML community, and everyone who helped me out.

Note:  

An uncensored model has no guardrails.  

You are responsible for anything you do with the model, just as you are responsible for anything you do with any dangerous object such as a knife, gun, lighter, or car.

Publishing anything this model generates is the same as publishing it yourself.

You are responsible for the content you publish, and you cannot blame the model any more than you can blame the knife, gun, lighter, or car for what you do with it.
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ehartford__Wizard-Vicuna-30B-Uncensored)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 53.44   |
| ARC (25-shot)         | 62.12          |
| HellaSwag (10-shot)   | 83.45    |
| MMLU (5-shot)         | 58.24         |
| TruthfulQA (0-shot)   | 50.81   |
| Winogrande (5-shot)   | 78.45   |
| GSM8K (5-shot)        | 14.25        |
| DROP (3-shot)         | 26.74         |

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ehartford__Wizard-Vicuna-30B-Uncensored)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |57.89|
|AI2 Reasoning Challenge (25-Shot)|62.12|
|HellaSwag (10-Shot)              |83.45|
|MMLU (5-Shot)                    |58.24|
|TruthfulQA (0-shot)              |50.81|
|Winogrande (5-shot)              |78.45|
|GSM8k (5-shot)                   |14.25|

