---
language:
- en
license: cc-by-nc-sa-4.0
library_name: transformers
datasets:
- garage-bAInd/Open-Platypus
pipeline_tag: text-generation
model-index:
- name: PlatYi-34B-Llama-Q-v3
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
      value: 64.33
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=kyujinpy/PlatYi-34B-Llama-Q-v3
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
      value: 84.88
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=kyujinpy/PlatYi-34B-Llama-Q-v3
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
      value: 74.98
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=kyujinpy/PlatYi-34B-Llama-Q-v3
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
      value: 51.8
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=kyujinpy/PlatYi-34B-Llama-Q-v3
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
      value: 84.21
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=kyujinpy/PlatYi-34B-Llama-Q-v3
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
      value: 6.67
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=kyujinpy/PlatYi-34B-Llama-Q-v3
      name: Open LLM Leaderboard
---

# **PlatYi-34B-Llama-Q-v3**  
<img src='./PlatYi.png' width=256>

## Model Details

**Model Developers** Kyujin Han (kyujinpy)

**Input** Models input text only.

**Output** Models generate text only.

**Model Architecture**   
PlatYi-34B-Llama-Q-v3 is an auto-regressive language model based on the Yi-34B transformer architecture.  

**Blog Link**  
Blog: [Coming soon...]  
Github: [Coming soon...]   

**Base Model**    
[chargoddard/Yi-34B-Llama](https://huggingface.co/chargoddard/Yi-34B-Llama)   

**Training Dataset**    
[garage-bAInd/Open-Platypus](https://huggingface.co/datasets/garage-bAInd/Open-Platypus).  
   
## Fix some bugs   
- Before model, there is some mistakes.  
- I modified the templates and warmup_steps.   

## Notice  
While training, I used Q-LoRA.
The lora_r values is 64.  


# **Model Benchmark**

## Open leaderboard
- Follow up as [link](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard).  

| Model | Average | ARC | HellaSwag | MMLU | TruthfulQA | Winogrande | GSM8K |
| --- | --- | --- | --- | --- | --- | --- | --- |
| PlatYi-34B-Llama-Q-v3 | 61.15 | 64.33 | 84.88 | 74.98 | 51.80 | 82.79 | 6.67 |
| PlatYi-34B-Llama-Q-v2 | 67.88 | 61.09 | 85.09 | 76.59 | 52.65 | 82.79 | 49.05 |
| PlatYi-34B-Llama-Q | 71.13 | 65.70 | 85.22 | 78.78 | 53.64 | 83.03 | 60.42 |
| PlatYi-34B-Llama | 68.37 | 67.83 | 85.35 | 78.26 | 53.46 | 82.87 | 42.46 |
| [Yi-34B-Llama](https://huggingface.co/chargoddard/Yi-34B-Llama) | 70.95 | 64.59 | 85.63 | 76.31 | 55.60 | 82.79 | 60.80 |
| [Yi-34B](https://huggingface.co/01-ai/Yi-34B) | 69.42 | 64.59 | 85.69 | 76.35 | 56.23 | 83.03 | 50.64 |
  
  
# Implementation Code
```python
### KO-Platypus
from transformers import AutoModelForCausalLM, AutoTokenizer
import torch

repo = "kyujinpy/PlatYi-34B-Llama-Q-v3"
OpenOrca = AutoModelForCausalLM.from_pretrained(
        repo,
        return_dict=True,
        torch_dtype=torch.float16,
        device_map='auto'
)
OpenOrca_tokenizer = AutoTokenizer.from_pretrained(repo)
```

---
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_kyujinpy__PlatYi-34B-Llama-Q-v3)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |61.15|
|AI2 Reasoning Challenge (25-Shot)|64.33|
|HellaSwag (10-Shot)              |84.88|
|MMLU (5-Shot)                    |74.98|
|TruthfulQA (0-shot)              |51.80|
|Winogrande (5-shot)              |84.21|
|GSM8k (5-shot)                   | 6.67|

