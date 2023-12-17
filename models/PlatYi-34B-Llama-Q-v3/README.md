---
language:
- en
datasets:
- garage-bAInd/Open-Platypus
library_name: transformers
pipeline_tag: text-generation
license: cc-by-nc-sa-4.0
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
| PlatYi-34B-Llama-Q-v3 | NaN | NaN | NaN | NaN | NaN | NaN | NaN |
| PlatYi-34B-Llama-Q-v2 | NaN | NaN | NaN | NaN | NaN | NaN | NaN |
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