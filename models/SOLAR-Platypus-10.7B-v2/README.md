---
language:
- en
datasets:
- garage-bAInd/Open-Platypus
library_name: transformers
pipeline_tag: text-generation
license: cc-by-nc-sa-4.0
---

# **SOLAR-Platypus-10.7B-v2**  

## Model Details

**Model Developers** Kyujin Han (kyujinpy)

**Input** Models input text only.

**Output** Models generate text only.

**Model Architecture**   
SOLAR-Platypus-10.7B-v2 is an auto-regressive language model based on the Llama2 architecture.  
  
**Base Model**    
[upstage/SOLAR-10.7B-v1.0](https://huggingface.co/upstage/SOLAR-10.7B-v1.0)   

**Training Dataset**    
[garage-bAInd/Open-Platypus](https://huggingface.co/datasets/garage-bAInd/Open-Platypus).  


## Notice  
While training, I used Q-LoRA.  
The lora_r values is 64.  

## Q-LoRA config
- LoRA_r: 64  
- LoRA_alpha: 16  
- LoRA_dropout: 0.05  
- LoRA_target_modules: [gate_proj, up_proj, down_proj, q_proj, k_proj, v_proj]  

## Prompt
```
## Human:

## Assistant:  
```

# **Model Benchmark**

## Open leaderboard
- Follow up as [link](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard).  

| Model | Average | ARC | HellaSwag | MMLU | TruthfulQA | Winogrande | GSM8K |
| --- | --- | --- | --- | --- | --- | --- | --- |
| SOLAR-Platypus-10.7B-v1 | 58.62 | 61.69 | 84.23 | 60.37 | 51.58 | 82.79 | 11.07 |
| SOLAR-Platypus-10.7B-v2 | 55.25 | 59.39 | 83.57 | 59.93 | 43.15 | 81.45 | 4.02 |
| [upstage/SOLAR-10.7B-v1.0](https://huggingface.co/upstage/SOLAR-10.7B-v1.0) | 66.04 | 61.95 | 84.60 | 65.48 | 45.04 | 83.66 | 55.50 |

  
# Implementation Code
```python
### KO-Platypus
from transformers import AutoModelForCausalLM, AutoTokenizer
import torch

repo = "kyujinpy/SOLAR-Platypus-10.7B-v2"
OpenOrca = AutoModelForCausalLM.from_pretrained(
        repo,
        return_dict=True,
        torch_dtype=torch.float16,
        device_map='auto'
)
OpenOrca_tokenizer = AutoTokenizer.from_pretrained(repo)
```

---