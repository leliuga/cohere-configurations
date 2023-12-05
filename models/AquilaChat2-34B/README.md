---
license: other
---


![Aquila_logo](./log.jpeg)


<h4 align="center">
    <p>
        <b>English</b> |
        <a href="https://huggingface.co/BAAI/AquilaChat2-34B/blob/main/README_zh.md">简体中文</a> 
    </p>
</h4>


<p align="center">
  <a href="https://github.com/FlagAI-Open/Aquila2" target="_blank">Github</a> • <a href="https://github.com/FlagAI-Open/Aquila2/blob/main/assets/wechat-qrcode.jpg" target="_blank">WeChat</a> <br>
</p>

We opensource our **Aquila2** series, now including **Aquila2**, the base language models, namely **Aquila2-7B** and **Aquila2-34B**, as well as **AquilaChat2**, the chat models, namely **AquilaChat2-7B** and **AquilaChat2-34B**, as well as the long-text chat models, namely **AquilaChat2-7B-16k** and **AquilaChat2-34B-16k**


2023.10.25 🔥 **AquilaChat2-34B v1.2** is based on the previous **AquilaChat2-34B**. 
The AquilaChat2-34B model is close to or exceeds the level of GPT3.5 in the subjective evaluation of 8 secondary ability dimensions.

The additional details of the Aquila model will be presented in the official technical report. Please stay tuned for updates on official channels.

### Note
<p>
We have discovered a data leakage problem with the GSM8K test data in the pre-training task dataset. Therefore, the evaluation results of GSM8K have been removed from the evaluation results.

Upon thorough investigation and analysis, it was found that the data leakage occurred in the mathematical dataset A (over 2 million samples), recommended by a team we have collaborated with multiple times. This dataset includes the untreated GSM8K test set (1319 samples). The team only performed routine de-duplication and quality checks but did not conduct an extra filtering check for the presence of the GSM8K test data, resulting in this oversight.

Our team has always strictly adhered to the principle that training data should not include test data. Taking this lesson from the error caused by not thoroughly checking the source of external data, we have investigated all 2 trillion tokens of data for various test datasets, including WTM22（en-zh), CLUEWSC, Winograd, HellaSwag, OpenBookQA, PIQA, ARC-e, BUSTSM, BoolQ, TruthfulQA, RAFT, ChID, EPRSTMT, TNEWS, OCNLI, SEM-Chinese, MMLU, C-Eval, CMMLU, CSL and HumanEval.
</p>

## Quick Start  AquilaChat2-34B（Chat model）

### 1. Inference

```python
from transformers import AutoTokenizer, AutoModelForCausalLM
from transformers import BitsAndBytesConfig
import torch

device = torch.device("cuda:0")
model_info = "BAAI/AquilaChat2-34B"
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
              model_name="AquilaChat2-34B")
print(out)
```


## License

Aquila2 series open-source model is licensed under [ BAAI Aquila Model Licence Agreement](https://huggingface.co/BAAI/AquilaChat2-34B/blob/main/BAAI-Aquila-Model-License%20-Agreement.pdf)