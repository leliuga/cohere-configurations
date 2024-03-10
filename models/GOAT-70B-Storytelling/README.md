---
license: llama2
tags:
- facebook
- meta
- pytorch
- llama
- llama-2
- Storywriter
model_type: llama
model-index:
- name: GOAT-70B-Storytelling
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
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=GOAT-AI/GOAT-70B-Storytelling
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
      value: 87.74
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=GOAT-AI/GOAT-70B-Storytelling
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
      value: 69.92
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=GOAT-AI/GOAT-70B-Storytelling
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
      value: 53.53
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=GOAT-AI/GOAT-70B-Storytelling
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
      value: 83.5
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=GOAT-AI/GOAT-70B-Storytelling
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
      value: 40.79
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=GOAT-AI/GOAT-70B-Storytelling
      name: Open LLM Leaderboard
---

![GOAT-70B-Storytelling](https://assets.adapt.ws/files/20231117_ehznrqludevtapck.png)
# GOAT-70B-Storytelling model
GOAT-70B-Storytelling model trained by GOAT.AI lab as a core model for an autonomous story-writing agent. 

# GOAT-Storytelling-Agent
This agent facilitates the generation of high-quality, cohesive, and captivating narratives, including stories and books. It achieves this by utilizing inputs such as plot outlines, character profiles, their interrelationships, and other relevant details. Examples are provided below.

# Model description
 - **Base Architecture:** LLaMA 2 70B 
 - **License:** llama2
 - **Context window length:** 4096 tokens

### Training details
Training was performed on a GPU cluster of 64xH100s. FSDP ZeRO-3 sharding is employed for efficient training. We instruction finetune on a dataset of 18K examples for one epoch with batch size of 336, AdamW optimizer with learning rate 1e-5.

### Learn more
- **Blogpost:** [GOAT-Storytelling: Arbitrarily Long Story Writing Agent](https://www.blog.goat.ai/goat-st/)
- **GitHub:** [here](https://github.com/GOAT-AI-lab/GOAT-Storytelling-Agent)
- **Generated examples:** [here](https://huggingface.co/datasets/GOAT-AI/generated-novels/tree/main/generated-books)

## Uses
The main purpose of GOAT-70B-Storytelling is to generate books, novels, movie scripts and etc. as an agent in coping with our GOAT-Storytelling-Agent. It is specifically designed for storywriters.

## Usage
Usage can be either self-hosted via `transformers` or used with Spaces

```python
import torch

from transformers import AutoTokenizer, AutoModelForCausalLM

model_name = "GOAT-AI/GOAT-70B-Storytelling"

tokenizer = AutoTokenizer.from_pretrained(model_name)
model = AutoModelForCausalLM.from_pretrained(
    model_name,
    torch_dtype=torch.bfloat16
)
```
Currently, we support LLM endpoint generation, where you need to send a post request to the generation endpoint (we recommend using Text Generation Inference by HuggingFace).

Here is how you can utilize the model via GOAT-Storytelling-Agent:

```python
from goat_storytelling_agent.storytelling_agent import StoryAgent

backend_uri = # Text generation endpoint
writer = StoryAgent(backend_uri, form='novel')
novel_scenes = writer.generate_story('treasure hunt in a jungle')
```

## License
GOAT-70B-Storytelling model is based on [Meta's LLaMA-2-70b-hf](https://huggingface.co/meta-llama/Llama-2-70b-hf), and using own datasets.  

GOAT-70B-Storytelling model weights are available under LLAMA-2 license.

### Risks and Biases 
GOAT-70B-Storytelling model can produce factually incorrect output and should not be relied on to deliver factually accurate information. Therefore, the GOAT-70B-Storytelling model could possibly generate wrong, biased, or otherwise offensive outputs.
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_GOAT-AI__GOAT-70B-Storytelling)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |67.38|
|AI2 Reasoning Challenge (25-Shot)|68.77|
|HellaSwag (10-Shot)              |87.74|
|MMLU (5-Shot)                    |69.92|
|TruthfulQA (0-shot)              |53.53|
|Winogrande (5-shot)              |83.50|
|GSM8k (5-shot)                   |40.79|

