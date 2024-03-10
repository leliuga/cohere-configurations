---
language:
- ko
library_name: transformers
pipeline_tag: text-generation
license: cc-by-nc-sa-4.0
tags:
- moe
- merge
---
**The license is `cc-by-nc-sa-4.0`.**  
  
# **🐻‍❄️SOLARC-MOE-10.7Bx4🐻‍❄️**  
![img](https://drive.google.com/uc?export=view&id=1_Qa2TfLMw3WeJ23dHkrP1Xln_RNt1jqG)  


## Model Details

**Model Developers** Seungyoo Lee(DopeorNope)

I am in charge of Large Language Models (LLMs) at Markr AI team in South Korea.

**Input** Models input text only.

**Output** Models generate text only.

**Model Architecture**  
SOLARC-MOE-10.7Bx4 is an auto-regressive language model based on the SOLAR architecture.

---

## **Base Model**  

[kyujinpy/Sakura-SOLAR-Instruct](https://huggingface.co/kyujinpy/Sakura-SOLAR-Instruct)   

[Weyaxi/SauerkrautLM-UNA-SOLAR-Instruct](https://huggingface.co/Weyaxi/SauerkrautLM-UNA-SOLAR-Instruct)   

[VAGOsolutions/SauerkrautLM-SOLAR-Instruct](https://huggingface.co/VAGOsolutions/SauerkrautLM-SOLAR-Instruct)   

[fblgit/UNA-SOLAR-10.7B-Instruct-v1.0](https://huggingface.co/fblgit/UNA-SOLAR-10.7B-Instruct-v1.0)   

## **Implemented Method**

I have built a model using the Mixture of Experts (MOE) approach, utilizing each of these models as the base.

---
  
# Implementation Code


## Load model
```python

from transformers import AutoModelForCausalLM, AutoTokenizer
import torch

repo = "DopeorNope/SOLARC-MOE-10.7Bx4"
OpenOrca = AutoModelForCausalLM.from_pretrained(
        repo,
        return_dict=True,
        torch_dtype=torch.float16,
        device_map='auto'
)
OpenOrca_tokenizer = AutoTokenizer.from_pretrained(repo)
```


---