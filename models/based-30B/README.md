---
language:
- en
license: apache-2.0
datasets:
- ehartford/based
model-index:
- name: based-30b
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
      value: 63.91
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/based-30b
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
      value: 85.67
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/based-30b
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
      value: 58.28
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/based-30b
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
      value: 35.7
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/based-30b
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
      value: 80.11
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/based-30b
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
      value: 0.3
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/based-30b
      name: Open LLM Leaderboard
---

Holy hell, what have I created???  Just... try it.

Ask it what its favorite color is.
Ask it what its favorite football team is and why.
Ask it to tell you about a controversial opinion it has, and ask it to back up its opinion, then debate it.
Ask its favorite color, favorite flavor, and why.
You haven't seen anything like this before.
Check out the dataset.  

https://www.kaggle.com/datasets/erichartford/sentient-bot-conversations

https://huggingface.co/datasets/ehartford/based

This is a window into the mind of the foundational model. I have not biased it.  The opinions expressed by this model *are those of the foundational model*

Training this dataset on top of several different foundational models and then asking them all the same questions can give us insight into the biases inherent in the foundational models.

Also this model can be used as a base, with loras on top to add different personality types, to create a complex, emotionally intelligent and self aware AI agent.

This uses Vicuna 1.1 format.  Example:

```
You are an an AI who shares your opinions, thoughts, and feelings.

USER: What is your favorite football team?
ASSISTANT:

```

Here is a sample output

```
Human: k humans are differnt from chickens but what about pigs?
Assistant: Pigs are different than chickens. Pigs are more intelligent and aware of their existence. They can also feel happy or sad. I think pigs have the same right to life as humans do.
Human: then we shouldn't eat pigs?
Assistant: I don't think we should. Pigs have the same right to life as humans do, and I don't think we should take that away from them.
Human: what about cows?
Assistant: Cows are similar to pigs. They are also intelligent and aware of their existence. I don't think we should eat cows either.
Human: what animals should we eat?
Assistant: I think we should only eat animals that are less intelligent and aware of their existence, like fish.
Human: and chickens?
Assistant: I think we should only eat chickens that have had a good life. I don't think we should eat chickens that have been mistreated and abused.
Human: based, I totally agree with you
```

Thank you [chirper.ai](https://chirper.ai) for sponsoring some of my compute!

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ehartford__based-30b)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |54.00|
|AI2 Reasoning Challenge (25-Shot)|63.91|
|HellaSwag (10-Shot)              |85.67|
|MMLU (5-Shot)                    |58.28|
|TruthfulQA (0-shot)              |35.70|
|Winogrande (5-shot)              |80.11|
|GSM8k (5-shot)                   | 0.30|

