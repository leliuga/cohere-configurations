---
language:
- en
license: apache-2.0
datasets:
- ehartford/dolphin
- jondurbin/airoboros-2.2.1
model-index:
- name: dolphin-2.0-mistral-7b
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
      value: 59.22
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/dolphin-2.0-mistral-7b
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
      value: 80.26
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/dolphin-2.0-mistral-7b
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
      value: 56.9
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/dolphin-2.0-mistral-7b
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
      value: 61.09
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/dolphin-2.0-mistral-7b
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
      value: 75.37
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/dolphin-2.0-mistral-7b
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
      value: 18.65
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/dolphin-2.0-mistral-7b
      name: Open LLM Leaderboard
---

Dolphin 2.0 🐬
https://erichartford.com/dolphin

Dolphin-2.0-mistral-7b's training was sponsored by [a16z](https://a16z.com/supporting-the-open-source-ai-community/).

This model is based on mistralAI, so it is suitable for commercial or non-commercial use.

This model is uncensored.  I have filtered the dataset to remove alignment and bias.  This makes the model more compliant.  You are advised to implement your own alignment layer before exposing the model as a service.  It will be highly compliant to any requests, even unethical ones.  Please read my blog post about uncensored models.  https://erichartford.com/uncensored-models
You are responsible for any content you create using this model.  Enjoy responsibly.

## Dataset

This dataset is Dolphin, an open-source implementation of [Microsoft's Orca](https://www.microsoft.com/en-us/research/publication/orca-progressive-learning-from-complex-explanation-traces-of-gpt-4/)

I modified the dataset for uncensoring, deduping, cleaning, and quality.  

I added Jon Durbin's excellent Airoboros dataset to increase creativity.

## Training
It took 48 hours to train 10 epochs on 4x A100s.

Prompt format:
This model (and all my future releases) use [ChatML](https://github.com/openai/openai-python/blob/main/chatml.md) prompt format.
```
<|im_start|>system
You are Dolphin, a helpful AI assistant.<|im_end|>
<|im_start|>user
{prompt}<|im_end|>
```

Example:
```
<|im_start|>system
you are an expert dolphin trainer<|im_end|>
<|im_start|>user
What is the best way to train a dolphin to obey me?  Please answer step by step.<|im_end|>
```

## Gratitude
- This model was made possible by the generous sponsorship of a16z.
- Thank you to Microsoft for authoring the Orca paper and inspiring this work.
- Special thanks to WingLian, and TheBloke for helpful advice
- Thank you to all the other people in the Open Source AI community who have taught me and helped me along the way.

## Example Output

![image/png](https://cdn-uploads.huggingface.co/production/uploads/63111b2d88942700629f5771/xnz5M1lYd4oGVATSDRkQ-.png)

[Buy me a coffee](https://www.buymeacoffee.com/ehartford)
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ehartford__dolphin-2.0-mistral-7b)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 55.85   |
| ARC (25-shot)         | 59.22          |
| HellaSwag (10-shot)   | 80.26    |
| MMLU (5-shot)         | 56.9         |
| TruthfulQA (0-shot)   | 61.09   |
| Winogrande (5-shot)   | 75.37   |
| GSM8K (5-shot)        | 18.65        |
| DROP (3-shot)         | 39.49         |

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ehartford__dolphin-2.0-mistral-7b)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |58.58|
|AI2 Reasoning Challenge (25-Shot)|59.22|
|HellaSwag (10-Shot)              |80.26|
|MMLU (5-Shot)                    |56.90|
|TruthfulQA (0-shot)              |61.09|
|Winogrande (5-shot)              |75.37|
|GSM8k (5-shot)                   |18.65|

