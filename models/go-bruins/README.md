---
language:
- en
license: cc-by-nc-4.0
datasets:
- Intel/orca_dpo_pairs
pipeline_tag: text-generation
model-index:
- name: go-bruins
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
      value: 69.11
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rwitz/go-bruins
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
      value: 86.73
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rwitz/go-bruins
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
      value: 64.94
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rwitz/go-bruins
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
      value: 58.71
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rwitz/go-bruins
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
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rwitz/go-bruins
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
      value: 69.9
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=rwitz/go-bruins
      name: Open LLM Leaderboard
---


![image/png](https://cdn-uploads.huggingface.co/production/uploads/63a259d0f30c46422789d38d/vO3iATjO8ulfcakTltE4k.png)

# Go Bruins - A Fine-tuned Language Model
Join my AI Discord: [rwitz](https://discord.gg/qbqjBEfkGw)


## Updates 
December 9, 2023:
Go-Bruins has placed **#6** overall and **#1** for 7 billion parameter models on the [Hugging Face Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)!
## Overview

**Go Bruins** is a state-of-the-art language model fine-tuned on the Q-bert/MetaMath-Cybertron-Starling architecture. It's designed to push the boundaries of NLP applications, offering unparalleled performance in generating human-like text.

## Model Details

- **Developer:** Ryan Witzman
- **Base Model:** [Q-bert/MetaMath-Cybertron-Starling](https://huggingface.co/Q-bert/MetaMath-Cybertron-Starling)
- **Fine-tuning Method:** Direct Preference Optimization (DPO)
- **Training Steps:** 200
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

model_name = "rwitz/go-bruins"
inference_pipeline = pipeline('text-generation', model=model_name)

input_text = "Your input text goes here"
output = inference_pipeline(input_text)

print(output)
```

GGUF Quantized Files are Located at [NyxKrage/go-bruins-GGUF](https://huggingface.co/NyxKrage/go-bruins-GGUF)

### Not Recommended For:

- Illegal activities
- Harassment
- Professional advice or crisis situations

## Training and Evaluation

Trained on a dataset from [Intel/orca_dpo_pairs](https://huggingface.co/datasets/Intel/orca_dpo_pairs), Go Bruins has shown promising improvements over its predecessor, Q-Bert.

# Evaluations
Go-Bruins is the SOTA 7B model.
| Metric        | Average | Arc Challenge | Hella Swag | MMLU | Truthful Q&A | Winogrande | GSM8k |
|---------------|---------|---------------|------------|------|--------------|------------|-------|
| **Score**     | 71.86    | 69.11         | 86.53| 65.02 | 59.24        | 81.37      | 69.90  |

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
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_rwitz__go-bruins)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |71.81|
|AI2 Reasoning Challenge (25-Shot)|69.11|
|HellaSwag (10-Shot)              |86.73|
|MMLU (5-Shot)                    |64.94|
|TruthfulQA (0-shot)              |58.71|
|Winogrande (5-shot)              |81.45|
|GSM8k (5-shot)                   |69.90|

