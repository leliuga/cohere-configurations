---
model-index:
- name: Huginn-13b-FP16
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
      value: 60.58
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=The-Face-Of-Goonery/Huginn-13b-FP16
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
      value: 82.53
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=The-Face-Of-Goonery/Huginn-13b-FP16
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
      value: 53.71
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=The-Face-Of-Goonery/Huginn-13b-FP16
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
      value: 54.46
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=The-Face-Of-Goonery/Huginn-13b-FP16
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
      value: 73.72
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=The-Face-Of-Goonery/Huginn-13b-FP16
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
      value: 4.32
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=The-Face-Of-Goonery/Huginn-13b-FP16
      name: Open LLM Leaderboard
---
a merge of a lot of different models, like hermes, beluga, airoboros, chronos.. limarp

significantly better quality than my previous chronos-beluga merge.




Huginn is intended as a general purpose model, that maintains a lot of good knowledge, can perform logical thought and accurately follow instructions, and hold the prose and creativity of more writing oriented models, this makes this model great for roleplays, while still being good as a normal chatbot or assistant
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_The-Face-Of-Goonery__Huginn-13b-FP16)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |54.89|
|AI2 Reasoning Challenge (25-Shot)|60.58|
|HellaSwag (10-Shot)              |82.53|
|MMLU (5-Shot)                    |53.71|
|TruthfulQA (0-shot)              |54.46|
|Winogrande (5-shot)              |73.72|
|GSM8k (5-shot)                   | 4.32|

