---
license: apache-2.0
language:
- fr
- it
- de
- es
- en
inference: false
---
# Model Card for Mixtral-Fusion-4x7B-Instruct-v0.1
This model is an experimental model created by merging [mistralai/Mixtral-8x7B-Instruct-v0.1](https://huggingface.co/mistralai/Mixtral-8x7B-Instruct-v0.1) experts.

# How we merged experts
Changed to merge using slerp.  
[Discussion](https://huggingface.co/mmnga/Mixtral-Fusion-4x7B-Instruct-v0.1/discussions/2)

[old merge version](https://huggingface.co/mmnga/Mixtral-Fusion-4x7B-Instruct-v0.1/tree/v0.1.0)  
~~We simply take the average of every two experts.weight.~~  
~~The same goes for gate.weight.~~  

# How To Convert
use colab cpu-high-memory.  
[convert_mixtral_8x7b_to_4x7b.ipynb](https://huggingface.co/mmnga/Mixtral-Fusion-4x7B-Instruct-v0.1/blob/main/notebook/convert_mixtral_8x7b_to_4x7b.ipynb)

# OtherModels
[mmnga/Mixtral-Extraction-4x7B-Instruct-v0.1](https://huggingface.co/mmnga/Mixtral-Extraction-4x7B-Instruct-v0.1)
  
# Usage
~~~python
pip install git+https://github.com/huggingface/transformers --upgrade
pip install torch accelerate bitsandbytes flash_attn
~~~

~~~python
from transformers import AutoTokenizer, AutoModelForCausalLM, MixtralForCausalLM
import torch

model_name_or_path = "mmnga/Mixtral-Fusion-4x7B-Instruct-v0.1"

tokenizer = AutoTokenizer.from_pretrained(model_name_or_path)
model = MixtralForCausalLM.from_pretrained(model_name_or_path, load_in_8bit=True)

text = "[INST] What was John Holt's vision on education? [/INST] "
inputs = tokenizer(text, return_tensors="pt")

outputs = model.generate(**inputs, max_new_tokens=128)
print(tokenizer.decode(outputs[0], skip_special_tokens=True))

~~~
