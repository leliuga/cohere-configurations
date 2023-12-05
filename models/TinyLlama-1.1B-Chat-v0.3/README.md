---
license: apache-2.0
datasets:
- cerebras/SlimPajama-627B
- bigcode/starcoderdata
- OpenAssistant/oasst_top1_2023-08-25
language:
- en
---
<div align="center">

# TinyLlama-1.1B
</div>

https://github.com/jzhang38/TinyLlama

The TinyLlama project aims to **pretrain** a **1.1B Llama model on 3 trillion tokens**. With some proper optimization, we can achieve this within a span of "just" 90 days using 16 A100-40G GPUs 🚀🚀. The training has started on 2023-09-01. 


We adopted exactly the same architecture and tokenizer as Llama 2. This means TinyLlama can be plugged and played in many open-source projects built upon Llama. Besides, TinyLlama is compact with only 1.1B parameters. This compactness allows it to cater to a multitude of applications demanding a restricted computation and memory footprint.

#### This Model
This is the chat model finetuned on top of [PY007/TinyLlama-1.1B-intermediate-step-480k-1T](https://huggingface.co/PY007/TinyLlama-1.1B-intermediate-step-480k-1T). 
The dataset used is [OpenAssistant/oasst_top1_2023-08-25](https://huggingface.co/datasets/OpenAssistant/oasst_top1_2023-08-25) following the [chatml](https://github.com/openai/openai-python/blob/main/chatml.md) format.
#### How to use
You will need the transformers>=4.31
Do check the [TinyLlama](https://github.com/jzhang38/TinyLlama) github page for more information.
```
from transformers import AutoTokenizer
import transformers 
import torch
model = "PY007/TinyLlama-1.1B-Chat-v0.3"
tokenizer = AutoTokenizer.from_pretrained(model)
pipeline = transformers.pipeline(
    "text-generation",
    model=model,
    torch_dtype=torch.float16,
    device_map="auto",
)

CHAT_EOS_TOKEN_ID = 32002

prompt = "How to get in a good university?"
formatted_prompt = (
    f"<|im_start|>user\n{prompt}<|im_end|>\n<|im_start|>assistant\n"
)


sequences = pipeline(
    formatted_prompt,
    do_sample=True,
    top_k=50,
    top_p = 0.9,
    num_return_sequences=1,
    repetition_penalty=1.1,
    max_new_tokens=1024,
    eos_token_id=CHAT_EOS_TOKEN_ID,
)

for seq in sequences:
    print(f"Result: {seq['generated_text']}")
```