---
license: other
---


![Aquila_logo](./log.jpeg)


<h4 align="center">
    <p>
        <b>English</b> |
        <a href="https://huggingface.co/BAAI/AquilaChat2-34B-16K/blob/main/README_zh.md">简体中文</a> 
    </p>
</h4>


<p align="center">
  <a href="https://github.com/FlagAI-Open/Aquila2" target="_blank">Github</a> • <a href="https://github.com/FlagAI-Open/Aquila2/blob/main/assets/wechat-qrcode.jpg" target="_blank">WeChat</a> <br>
</p>


We opensource our **Aquila2** series, now including **Aquila2**, the base language models, namely **Aquila2-7B** and **Aquila2-34B**, as well as **AquilaChat2**, the chat models, namely **AquilaChat2-7B** and **AquilaChat2-34B**, as well as the long-text chat models, namely **AquilaChat2-7B-16k** and **AquilaChat2-34B-16k**


2023.10.25 🔥 **AquilaChat2-34B-16K v1.2** is based on the previous **AquilaChat2-34B-16K**. The AquilaChat2-34B-16K-V1.2 has significantly improved long-text synthesis capabilities compared to the V1 version, 
approaching the level of GPT-3.5-16K. Additionally, the V1.2 version incorporates more conventional instruction fine-tuning corpora, enhancing its performance in non-long-text scenarios compared to the V1 version.

The additional details of the Aquila model will be presented in the official technical report. Please stay tuned for updates on official channels.


## Quick Start  AquilaChat2-34B-16K（Chat model）

### 1. Inference

```python
from transformers import AutoTokenizer, AutoModelForCausalLM
import torch

device = torch.device("cuda:0")
model_info = "BAAI/AquilaChat2-34B-16k"
tokenizer = AutoTokenizer.from_pretrained(model_info, trust_remote_code=True)
quantization_config=BitsAndBytesConfig(
                        load_in_4bit=True,
                        bnb_4bit_use_double_quant=True,
                        bnb_4bit_quant_type="nf4",
                        bnb_4bit_compute_dtype=torch.bfloat16,
                    )
model = AutoModelForCausalLM.from_pretrained(model_info, trust_remote_code=True, torch_dtype=torch.bfloat16,
                                                # quantization_config=quantization_config, # Uncomment this line for 4bit quantization
                                                )
model.eval()
model.to(device)
text = "请给出10个要到北京旅游的理由。"
from predict import predict
out = predict(model, text, tokenizer=tokenizer, max_gen_len=200, top_p=0.9,
              seed=123, topk=15, temperature=1.0, sft=True, device=device,
              model_name="AquilaChat2-34B-16K")
print(out)
```


## License

Aquila2 series open-source model is licensed under [ BAAI Aquila Model Licence Agreement](https://huggingface.co/BAAI/AquilaChat2-34B-16K/blob/main/BAAI-Aquila-Model-License%20-Agreement.pdf)