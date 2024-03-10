---
language:
- en
license: apache-2.0
datasets:
- Intel/orca_dpo_pairs
- athirdpath/DPO_Pairs-Roleplay-Alpaca-NSFW
pipeline_tag: text-generation
model-index:
- name: go-bruins-v2
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
      value: 69.8
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rwitz/go-bruins-v2
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
      value: 87.05
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rwitz/go-bruins-v2
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
      value: 64.75
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rwitz/go-bruins-v2
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
      value: 59.7
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rwitz/go-bruins-v2
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
      value: 81.45
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rwitz/go-bruins-v2
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
      value: 69.67
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rwitz/go-bruins-v2
      name: Open LLM Leaderboard
---



![image/png](https://cdn-uploads.huggingface.co/production/uploads/63a259d0f30c46422789d38d/tmdM1fjNAmzV125zWd3_J.png)
Join my AI Discord: [rwitz](https://discord.gg/qbqjBEfkGw)


# Go Bruins V2 - A Fine-tuned Language Model

## Updates 

## Overview

**Go Bruins-V2** is a language model fine-tuned on the rwitz/go-bruins architecture. It's designed to push the boundaries of NLP applications, offering unparalleled performance in generating human-like text.

## Model Details

- **Developer:** Ryan Witzman
- **Base Model:** [rwitz/go-bruins](https://huggingface.co/rwitz/go-bruins)
- **Fine-tuning Method:** Direct Preference Optimization (DPO)
- **Training Steps:** 642
- **Language:** English
- **License:** MIT

## Capabilities

Go Bruins excels in a variety of NLP tasks, including but not limited to:
- Text generation
- Language understanding
- Sentiment analysis

## Usage

**Warning:** This model may output NSFW or illegal content. Use with caution and at your own risk.

### For Direct Use:

```python
from transformers import pipeline

model_name = "rwitz/go-bruins-v2"
inference_pipeline = pipeline('text-generation', model=model_name)

input_text = "Your input text goes here"
output = inference_pipeline(input_text)

print(output)
```


### Not Recommended For:

- Illegal activities
- Harassment
- Professional advice or crisis situations

## Training and Evaluation

Trained on a dataset from [athirdpath/DPO_Pairs-Roleplay-Alpaca-NSFW](https://huggingface.co/datasets/athirdpath/DPO_Pairs-Roleplay-Alpaca-NSFW), Go Bruins V2 has shown promising improvements over its predecessor, Go Bruins.

# Evaluations

| Metric        | Average | Arc Challenge | Hella Swag | MMLU | Truthful Q&A | Winogrande | GSM8k |
|---------------|---------|---------------|------------|------|--------------|------------|-------|
| **Score**     | 72.07    | 69.8        | 87.05|  64.75      | 59.7     | 81.45  | 69.67 |

Note: The original MMLU evaluation has been corrected to include 5-shot data rather than 1-shot data.
## Contact

For any inquiries or feedback, reach out to Ryan Witzman on Discord: `rwitz_`.

---
## Citations
```
@misc{unacybertron7b,
  title={Cybertron: Uniform Neural Alignment}, 
  author={Xavier Murias},
  year={2023},
  publisher = {HuggingFace},
  journal = {HuggingFace repository},
  howpublished = {\url{https://huggingface.co/fblgit/una-cybertron-7b-v2-bf16}},
}
```

*This model card was created with care by Ryan Witzman.*
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_rwitz__go-bruins-v2)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |72.07|
|AI2 Reasoning Challenge (25-Shot)|69.80|
|HellaSwag (10-Shot)              |87.05|
|MMLU (5-Shot)                    |64.75|
|TruthfulQA (0-shot)              |59.70|
|Winogrande (5-shot)              |81.45|
|GSM8k (5-shot)                   |69.67|

