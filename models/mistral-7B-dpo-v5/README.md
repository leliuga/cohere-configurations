---
license: apache-2.0
language:
- en
---
# Model Card for mncai/mistral-7b-dpo-v5

### Introduction of MindsAndCompany

https://mnc.ai/

We create various AI models and develop solutions that can be applied to businesses. And as for generative AI, we are developing products like Code Assistant, TOD Chatbot, LLMOps, and are in the process of developing Enterprise AGI (Artificial General Intelligence).

### Model Summary
based mistral, instruction tuned and dpo.


### How to Use
Here give some examples of how to use our model.

```python
from transformers import AutoConfig, AutoModel, AutoTokenizer
import transformers
import torch
hf_model = 'mncai/mistral-7b-dpo-v5' 
message = "<|user|>\n두 개의 구가 있는데 각각 지름이 1, 2일때 각 구의 부피는 몇배야? 설명도 같이 해줘.\n<|assistant|>\n"

sequences = pipeline(
    message,
    do_sample=True,
    top_k=10,
    num_return_sequences=1,
    eos_token_id=tokenizer.eos_token_id,
    max_length=2048,
)
for seq in sequences:
    print(f"Result: {seq['generated_text']}")
```


### Contact
If you have any questions, please raise an issue or contact us at dwmyoung@mnc.ai