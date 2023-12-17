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
We simply take the average of every two experts.weight.  
The same goes for gate.weight.

# How To Convert
use colab cpu-high-memory.  
[convert_mixtral_8x7b_to_4x7b.ipynb](https://huggingface.co/mmnga/Mixtral-Fusion-4x7B-Instruct-v0.1/blob/main/notebook/convert_mixtral_8x7b_to_4x7b.ipynb)

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

# set num_experts_per_tok 1 or 2 ?
model.config.num_experts_per_tok = 2

# message
messages = [
    {"role": "user", "content": "Tell me what's for dinner tonight."},
]

with torch.no_grad():
    token_ids = tokenizer.apply_chat_template(messages, return_tensors="pt")
    output_ids = model.generate(
        token_ids.to(model.device),
        temperature=0.5,
        do_sample=True,
        top_p=0.95,
        top_k=40,
        max_new_tokens=128,
        repetition_penalty=1.5
    )
output = tokenizer.decode(output_ids[0][token_ids.size(1) :])
print(output)

~~~