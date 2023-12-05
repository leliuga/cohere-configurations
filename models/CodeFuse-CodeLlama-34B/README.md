---
license: other
tasks: 
- code-generation
---
# Model Card for CodeFuse-CodeLlama-34B
![logo](LOGO.png)

[[中文]](#chinese)    [[English]](#english)



<a id="english"></a>

## Model Description

CodeFuse-CodeLlama-34B is a 34B Code-LLM finetuned by QLoRA of multiple code tasks（600k instrunctions/answers） on the base model CodeLlama-34b-Python. 
The context length of finetuning is 4K while it is able to be finetuned by 16k context if necessary.
<br>

## News and Updates

🔥🔥🔥 2023-09-26 We are pleased to announce the release of the [4-bit quantized version](https://huggingface.co/codefuse-ai/CodeFuse-CodeLlama-34B-4bits) of CodeFuse-CodeLlama-34B. Despite the quantization process, the model still achieves a remarkable 73.8% accuracy (greedy decoding) on the HumanEval pass@1 metric.

🔥🔥🔥 2023-09-11 CodeFuse-CodeLlama34B has achieved 74.4% of pass@1 (greedy decoding) on HumanEval, which is SOTA results for openspurced LLMs at present.

<br>

## Code Community

**Homepage**: 🏡 https://github.com/codefuse-ai (**Please give us your support with a Star🌟 + Fork🚀 + Watch👀**)

+ If you wish to fine-tune the model yourself, you can visit ✨[MFTCoder](https://github.com/codefuse-ai/MFTCoder)✨✨

+ If you wish to deploy the model yourself, you can visit ✨[FasterTransformer4CodeFuse](https://github.com/codefuse-ai/FasterTransformer4CodeFuse)✨✨

+ If you wish to see a demo of the model, you can visit ✨[CodeFuse Demo](https://github.com/codefuse-ai/codefuse)✨✨


## Performance


| Model                       | HumanEval(pass@1) |  Date   |
|:----------------------------|:-----------------:|:-------:|
| **CodeFuse-CodeLlama-34B**  |     **74.4%**      | 2023.9  |
| WizardCoder-Python-34B-V1.0 |       73.2%       | 2023.8  |
| GPT-4(zero-shot)            |       67.0%       | 2023.3  |
| PanGu-Coder2 15B            |       61.6%       | 2023.8  |
| CodeLlama-34b-Python        |       53.7%       | 2023.8  |
| CodeLlama-34b               |       48.8%       | 2023.8  |
| GPT-3.5(zero-shot)          |       48.1%       | 2022.11 |
| OctoCoder                   |       46.2%       | 2023.8  |
| StarCoder-15B               |       33.6%       | 2023.5  |
| LLaMA 2 70B(zero-shot)      |       29.9%       | 2023.7  |

<br>

## Requirements

* python>=3.8 
* pytorch>=2.0.0
* transformers==4.32.0
* Sentencepiece
* CUDA 11.4
  <br>

##  Inference String Format

The inference string is a concatenated string formed by combining conversation data(system, human and bot contents) in the training data format.  It is used as input during the inference process.
Here is an example format of the concatenated string:

```python
"""
<|role_start|>system<|role_end|>System instruction
<|role_start|>human<|role_end|>Human 1st round input
<|role_start|>bot<|role_end|>Bot 1st round output</s>
<|role_start|>human<|role_end|>Human 2nd round input
<|role_start|>bot<|role_end|>Bot 2nd round output</s>
...
...
...
<|role_start|>human<|role_end|>Human nth round input
<|role_start|>bot<|role_end|>{Bot output to be genreated}</s>
"""
```

When applying inference, you always make your input string end with "<|role_start|>bot<|role_end|>" to ask the model generating answers.

## Quickstart

```bash
pip install -r requirements.txt
```

```python
import torch
from transformers import (
    AutoTokenizer, 
    AutoModelForCausalLM,
)
tokenizer = AutoTokenizer.from_pretrained(mode_name_or_path, trust_remote_code=True, use_fast=False, legacy=False)
tokenizer.padding_side = "left"
tokenizer.pad_token_id = tokenizer.convert_tokens_to_ids("<unk>")
tokenizer.eos_token_id = tokenizer.convert_tokens_to_ids("</s>")
# try 4bit loading if cuda memory not enough
model = AutoModelForCausalLM.from_pretrained(mode_name_or_path,
                                             trust_remote_code=True,
                                             load_in_4bit=False,
                                             device_map="auto",
                                             torch_dtype=torch.bfloat16)
model.eval()

HUMAN_ROLE_START_TAG = "<|role_start|>human<|role_end|>"
BOT_ROLE_START_TAG = "<|role_start|>bot<|role_end|>"

text = f"{HUMAN_ROLE_START_TAG}write a python function of quick sort.{BOT_ROLE_START_TAG}" 
inputs = tokenizer(text, return_tensors='pt', padding=True, add_special_tokens=False).to("cuda")
outputs = model.generate(
        inputs=inputs["input_ids"],
        attention_mask=inputs["attention_mask"],
        max_new_tokens=512,
        top_p=0.95,
        temperature=0.1,
        do_sample=True,
        eos_token_id=tokenizer.eos_token_id,
        pad_token_id=tokenizer.pad_token_id
    )
gen_text = tokenizer.batch_decode(outputs[:, inputs["input_ids"].shape[1]:], skip_special_tokens=True)
print(gen_text)
```

## MD5
We notice that the file may be corrupted during transfer process. Please check MD5 value before use.

| Model File                       | MD5 Value                        |
|:---------------------------------|:--------------------------------:|
| pytorch_model-00001-of-00007.bin | 8d544b1bcb3449934184d4141137329c |
| pytorch_model-00002-of-00007.bin | 9d5dbb30911e48a42fb6d0fcabb322a4 |
| pytorch_model-00003-of-00007.bin | b0d4aecee0457d9332005a187e1fffed |
| pytorch_model-00004-of-00007.bin | 5c7e002de5eab77d0194a2b0f6de0c24 |
| pytorch_model-00005-of-00007.bin | d22a511aa26b5b17117b665a877490ab |
| pytorch_model-00006-of-00007.bin | a5c28ac277fac07d16dd66537e54d109 |
| pytorch_model-00007-of-00007.bin | a967e2c6195477b7407089c0bffa2d53 |


## Citation
If you find our [work](https://arxiv.org/abs/2311.02303) useful or helpful for your R&D works, please feel free to cite our paper as below.
```
@article{mftcoder2023,
      title={MFTCoder: Boosting Code LLMs with Multitask Fine-Tuning}, 
      author={Bingchang Liu and Chaoyu Chen and Cong Liao and Zi Gong and Huan Wang and Zhichao Lei and Ming Liang and Dajun Chen and Min Shen and Hailian Zhou and Hang Yu and Jianguo Li},
      year={2023},
      journal={arXiv preprint arXiv},
      archivePrefix={arXiv},
      eprint={2311.02303}
}
```

<a id="chinese"></a>

## 模型简介

CodeFuse-CodeLlama34B-MFT 是一个通过QLoRA对基座模型CodeLlama-34b-Python进行多代码任务微调的代码大模型。模型微调采用了4k上下文。如果有必要，可以扩展到16k。
<br>

## 新闻

🔥🔥🔥 CodeFuse-CodeLlama34B-MFT模型在HumanEval pass@1上可以达到74.4%, 为当前开源SOTA。

<br>

## 代码社区
**大本营**： 🏡 https://github.com/codefuse-ai （**欢迎为我们的项目一键三连 Star🌟 + Fork🚀 + Watch👀**）

+ 如果您想自己微调该模型，可以访问 ✨[MFTCoder](https://github.com/codefuse-ai/MFTCoder)✨✨

+ 如果您想自己部署该模型，可以访问 ✨[FasterTransformer4CodeFuse](https://github.com/codefuse-ai/FasterTransformer4CodeFuse)✨✨

+ 如果您想观看该模型示例，可以访问 ✨[CodeFuse Demo](https://github.com/codefuse-ai/codefuse)✨✨


## 评测表现(代码)

| 模型                          | HumanEval(pass@1) |   日期    |
|:----------------------------|:-----------------:|:-------:|
| **CodeFuse-CodeLlama-34B**  |     **74.4%**      | 2023.9  |
| WizardCoder-Python-34B-V1.0 |       73.2%       | 2023.8  |
| GPT-4(zero-shot)            |       67.0%       | 2023.3  |
| PanGu-Coder2 15B            |       61.6%       | 2023.8  |
| CodeLlama-34b-Python        |       53.7%       | 2023.8  |
| CodeLlama-34b               |       48.8%       | 2023.8  |
| GPT-3.5(zero-shot)          |       48.1%       | 2022.11 |
| OctoCoder                   |       46.2%       | 2023.8  |
| StarCoder-15B               |       33.6%       | 2023.5  |
| LLaMA 2 70B(zero-shot)      |       29.9%       | 2023.7  |

<br>

## Requirements

* python>=3.8 
* pytorch>=2.0.0
* transformers==4.32.0
* CUDA 11.4
<br>

## 推理数据格式

推理数据为模型在训练数据格式下拼接的字符串形式，它也是推理时输入prompt拼接的方式：

```python
"""
<|role_start|>system<|role_end|>这是System指令
<|role_start|>human<|role_end|>这是第1轮用户输入的问题
<|role_start|>bot<|role_end|>这是第1轮模型生成的内容</s>
<|role_start|>human<|role_end|>这是第2轮用户输入的问题
<|role_start|>bot<|role_end|>这是第2轮模型生成的内容</s>
...
...
...
<|role_start|>human<|role_end|>这是第n轮用户输入的问题
<|role_start|>bot<|role_end|>{模型现在要生成的内容}</s>
"""
```

推理时，请确保拼接的prompt字符串以"<|role_start|>bot<|role_end|>"结尾，引导模型生成回答。

## 快速使用

```python
from transformers import (
    AutoTokenizer, 
    AutoModelForCausalLM,
)
tokenizer = AutoTokenizer.from_pretrained(mode_name_or_path, trust_remote_code=True, use_fast=False, legacy=False)
tokenizer.padding_side = "left"
tokenizer.pad_token_id = tokenizer.convert_tokens_to_ids("<unk>")
tokenizer.eos_token_id = tokenizer.convert_tokens_to_ids("</s>")
# 如果显存不够，可以考虑量化加载
model = AutoModelForCausalLM.from_pretrained(mode_name_or_path,
                                             trust_remote_code=True,
                                             load_in_4bit=False,
                                             device_map="auto",
                                             torch_dtype=torch.bfloat16)
model.eval()

HUMAN_ROLE_START_TAG = "<|role_start|>human<|role_end|>"
BOT_ROLE_START_TAG = "<|role_start|>bot<|role_end|>"

text = f"{HUMAN_ROLE_START_TAG}请用C++实现求解第n个斐波那契数{BOT_ROLE_START_TAG}" 
inputs = tokenizer(text, return_tensors='pt', padding=True, add_special_tokens=False).to("cuda")
outputs = model.generate(
        inputs=inputs["input_ids"],
        attention_mask=inputs["attention_mask"],
        max_new_tokens=512,
        top_p=0.95,
        temperature=0.1,
        do_sample=True,
        eos_token_id=tokenizer.eos_token_id,
        pad_token_id=tokenizer.pad_token_id
    )
gen_text = tokenizer.batch_decode(outputs[:, inputs["input_ids"].shape[1]:], skip_special_tokens=True)
print(gen_text)
```


## MD5
我们发现模型文件可能会在传输过程中损坏，使用前请检查文件MD5值。

| 模型文件                           | MD5值                            |
|:---------------------------------|:--------------------------------:|
| pytorch_model-00001-of-00007.bin | 8d544b1bcb3449934184d4141137329c |
| pytorch_model-00002-of-00007.bin | 9d5dbb30911e48a42fb6d0fcabb322a4 |
| pytorch_model-00003-of-00007.bin | b0d4aecee0457d9332005a187e1fffed |
| pytorch_model-00004-of-00007.bin | 5c7e002de5eab77d0194a2b0f6de0c24 |
| pytorch_model-00005-of-00007.bin | d22a511aa26b5b17117b665a877490ab |
| pytorch_model-00006-of-00007.bin | a5c28ac277fac07d16dd66537e54d109 |
| pytorch_model-00007-of-00007.bin | a967e2c6195477b7407089c0bffa2d53 |