---
license: other
language:
- en
- zh
library_name: transformers
pipeline_tag: text-generation
tags:
- llm
- Nanbeige
- custom_code
extra_gated_prompt: "访问此模型需要阅读并同意以下协议[这里](https://github.com/Nanbeige/Nanbeige/blob/main/%E5%8D%97%E5%8C%97%E9%98%81%E5%A4%A7%E8%AF%AD%E8%A8%80%E6%A8%A1%E5%9E%8B%E8%AE%B8%E5%8F%AF%E5%8D%8F%E8%AE%AE.pdf)\nAccess to this model requires reading and agreeing to the following agreement [here](https://github.com/Nanbeige/Nanbeige/blob/main/License_Agreement_for_Large_Language_Models_Nanbeige.pdf)"
extra_gated_fields:
 Name: text
 Country: text
 Affiliation: text
 Email: text
 I agree to the license terms described in the above agreement: checkbox
---

<!-- markdownlint-disable first-line-h1 -->

<!-- markdownlint-disable html -->

<div align="center">
<h1>
  Nanbeige-16B-Chat
</h1>
</div>

<p align="center">
  <a href="https://github.com/Nanbeige/Nanbeige" target="_blank">💻Github</a>
</p>

# <span id="Introduction">模型介绍（Introduction）</span>

Nanbeige-16B（南北阁-16B）是南北阁大模型实验室研发的160亿参数规模的大语言模型，采用了2.5T Tokens进行预训练，数据包含大量互联网高质量语料、各类书籍、代码等领域脱敏文本，在各个权威测评数据集上都取得了不错的效果。本次发布包含有 Base、Chat 以及扩展上下文长度的 Base-32k、Chat-32k 版本。

Base-32k 版本基于Nanbeige-16B-Base模型，采用YaRN插值方法对位置编码进行修改，并以32k为最大长度进行了20B Tokens的中文、英文、代码语料的全参数增量预训练。

Chat 版本和 Chat-32k 版本分别基于Nanbeige-16B-Base模型和Nanbeige-16B-Base-32k模型，经过了大量人类对齐训练，能够更好、更安全地回复用户的问题。

如果您需要处理更长的上下文，我们推荐您使用Nanbeige-16B-Base-32k和Nanbeige-16B-Chat-32k。

本仓库为 Nanbeige-16B-Chat 的模型仓库。

Nanbeige-16B is a 16 billion parameter language model developed by Nanbeige LLM Lab. It uses 2.5T Tokens for pre-training. The training data includes a large amount of high-quality internet corpus, various books, code, etc. It has achieved good results on various authoritative evaluation data sets. This release includes the Base, Chat, Base-32k and Chat-32k.

The Base-32k version is based on the Nanbeige-16B-Base model, which uses YaRN interpolation method to modify the position encoding, and performs full parameter incremental pre-training with 20 billion tokens of Chinese, English, and code corpora, under 32k max length.

The Chat version and Chat-32k version are based on the Nanbeige-16B-Base model and Nanbeige-16B-Base-32k model respectively. They have undergone extensive human-aligned training.

If you need to deal with longer contexts, we recommend using Nanbeige-16B-Base-32k and Nanbeige-16B-Chat-32k.

This repository is the one for Nanbeige-16B-Chat model.

##
|         | Base Model  | Base-32k Model | Chat Model | Chat-32k Model |
|:-------:|:-------:|:-------:|:---------:|:--------:|
| 16B     | 🤗 [Nanbeige-16B-Base](https://huggingface.co/Nanbeige/Nanbeige-16B-Base) | 🤗 [Nanbeige-16B-Base-32k](https://huggingface.co/Nanbeige/Nanbeige-16B-Base-32k) | 🤗 [Nanbeige-16B-Chat](https://huggingface.co/Nanbeige/Nanbeige-16B-Chat) |🤗 [Nanbeige-16B-Chat-32k](https://huggingface.co/Nanbeige/Nanbeige-16B-Chat-32k) |


##
# <span id="Inference">模型推理 (Inference)</span>

## 相关依赖

- python 3.8及以上版本
  
- transformers 4.33.3
  
- pytorch 2.0及以上版本
  
- python 3.8 and above
  
- transformers 4.33.3
  
- pytorch 2.0及以上版本

可以通过以下pip命令安装相关依赖库

You can install the dependent libraries with the following pip command

```
pip install transformers==4.33.3 transformers_stream_generator deepspeed einops==0.3.2 datasets==2.10.0
```

## 推理代码

通过以下代码可以调用模型进行聊天对话：

The following code can be used to invoke the model for chat dialogue:

```python
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer

tokenizer = AutoTokenizer.from_pretrained("Nanbeige/Nanbeige-16B-Chat", use_fast=False, trust_remote_code=True)
model = AutoModelForCausalLM.from_pretrained("Nanbeige/Nanbeige-16B-Chat", device_map="auto", torch_dtype=torch.bfloat16, trust_remote_code=True)

question = "你可以给我一些具体的SEO优化技巧吗？"

output, messages = model.chat(tokenizer, question)
print(output)
```

##
# <span id="Evaluation">性能测评（Performance Evaluation）</span>

### LLMEval-3
**LLMEval-3** ( [Github](https://github.com/llmeval/llmeval-3) / [主页](http://llmeval.com/index) ) 聚焦于专业知识能力评测，涵盖哲学、经济学、法学、教育学、文学、历史学、理学、工学、农学、医学、军事学、管理学、艺术学等教育部划定的13个学科门类、50余个二级学科，共计约20W道标准生成式问答题目。防止作弊是LLMEval-3考虑的重要因素。现有公开评测基准存在测试题库泄露的问题，因此可能出现“刷榜”、“刷分”等不公平现象，在LLMEval-3中，每个参与评测的系统需要完成从总题库中随机抽样的1000题，针对同一机构的模型，确保每次评测题目不重复。

我们基于 LLMEval-3 对 Nanbeige-16B-Chat 模型进行了全面测评，测评结果如下：

We conducted a comprehensive evaluation of Nanbeige-16B-Chat model based on **LLMEval-3** ( [Github](https://github.com/llmeval/llmeval-3) / [Homepage](http://llmeval.com/index) ), and the results are as follows:

| 模型名称                       | 相对分数-GPT4 | 相对分数-GPT3.5 | 绝对分数  | 工学   | 经济学  | 教育学  | 法学   | 文学   | 管理学  | 理学   | 历史学  | 医学   | 军事学  |
|----------------------------|-----------|-------------|-------|------|------|------|------|------|------|------|------|------|------|
| Baidu3.5                   | 104.21    | 121.39      | 77.53 | 8.13 | 8.00 | 8.63 | 7.97 | 6.23 | 7.63 | 7.33 | 8.77 | 7.47 | 7.37 |
| ChatGLM-pro                | 103.45    | 120.51      | 76.97 | 6.97 | 8.47 | 7.97 | 8.93 | 7.23 | 7.70 | 6.33 | 8.37 | 7.13 | 7.87 |
| GPT-4                      | 100.00    | 116.49      | 74.40 | 7.23 | 7.80 | 7.73 | 8.40 | 6.73 | 7.67 | 7.73 | 7.07 | 6.20 | 7.83 |
| **Nanbeige-16B-Chat**               | 94.26     | 109.80      | 70.13 | 6.00 | 7.87 | 8.20 | 8.33 | 6.07 | 6.83 | 6.00 | 7.80 | 5.80 | 7.23 |
| minimax-abab5              | 93.28     | 108.66      | 69.40 | 5.83 | 7.50 | 7.77 | 8.37 | 6.40 | 6.33 | 5.07 | 8.33 | 5.93 | 7.87 |
| Baichuan2-13B-Chat         | 92.91     | 108.23      | 69.13 | 6.00 | 7.53 | 8.63 | 8.13 | 6.23 | 6.33 | 5.63 | 8.20 | 5.43 | 7.00 |
| Qwen-14B-Chat              | 86.33     | 100.56      | 64.23 | 5.77 | 7.07 | 7.07 | 7.37 | 5.70 | 6.20 | 5.93 | 6.97 | 5.40 | 6.77 |
| GPT-3.5-turbo              | 85.84     | 100.00      | 63.87 | 6.27 | 6.87 | 7.23 | 7.40 | 5.40 | 6.30 | 6.37 | 6.00 | 5.17 | 6.87 |
| ChatGLM2-6B                | 75.71     | 88.19       | 56.33 | 4.03 | 6.33 | 7.00 | 7.57 | 4.77 | 5.93 | 4.23 | 5.87 | 5.07 | 5.53 |
| ziya_v1.1-13b              | 70.20     | 81.78       | 52.23 | 4.67 | 5.77 | 6.07 | 6.53 | 4.53 | 5.33 | 3.70 | 5.00 | 4.63 | 6.00 |
| BELLE-Llama2-13B-chat-0.4M | 68.82     | 80.16       | 51.20 | 4.47 | 5.93 | 6.20 | 6.77 | 4.33 | 4.97 | 4.10 | 5.07 | 3.77 | 5.60 |
| Linly-Chinese-LLaMA2-13B   | 67.82     | 79.00       | 50.46 | 3.87 | 5.80 | 5.83 | 6.57 | 3.93 | 5.37 | 4.07 | 5.43 | 3.93 | 5.67 |
| InternLM-Chat-7B           | 61.06     | 71.13       | 45.43 | 3.83 | 5.13 | 5.27 | 6.57 | 3.90 | 4.83 | 3.10 | 4.87 | 3.67 | 4.27 |
| Llama-2-7b-chat-hf         | 51.11     | 59.54       | 38.03 | 3.33 | 4.77 | 3.77 | 5.03 | 3.07 | 3.77 | 3.93 | 4.00 | 2.40 | 3.97 |


##
# <span id="Limitations">局限性（Limitations）</span>

虽然我们在训练过程中非常注重模型的安全性，力求确保其输出符合伦理和法律要求的文本，但由于模型大小和概率生成范式的限制，无法完全避免产生各种不符合预期的输出情况。这些输出可能包含偏见、歧视等有害内容，请勿传播这些内容。我们不承担因传播不良信息而导致的任何后果。

While we place great emphasis on the safety of the model during the training process, striving to ensure that its outputs align with ethical and legal requirements, it may not completely avoid generating unexpected outputs due to the model's size and probabilistic nature. These outputs may include harmful content such as bias or discrimination. Please don't propagate such content. We do not assume any responsibility for the consequences resulting from the dissemination of inappropriate information.

# <span id="License">协议（License）</span>

使用 Nanbeige 模型时，您必须遵守 Apache 2.0 许可证和[《南北阁大语言模型许可协议》](https://huggingface.co/Nanbeige/Nanbeige-16B-Chat/resolve/main/南北阁大语言模型许可协议.pdf)。如果您打算将 Nanbeige 模型或其衍生产品用于商业目的，请通过以下联系邮箱 nanbeige@126.com 提交申请材料，以满足《南北阁大语言模型许可协议》的要求。经过审核后，我们将授予您非排他性、全球范围内、不可转让、不可再许可、可撤销的商业版权许可。

When using the Nanbeige models, you must comply with the Apache 2.0 License and the [License Agreement for Large Language Models Nanbeige](https://huggingface.co/Nanbeige/Nanbeige-16B-Chat/resolve/main/License_Agreement_for_Large_Language_Models_Nanbeige.pdf). If you intend to use the Nanbeige Models or its derivatives for commercial purposes, please submit application materials to meet the requirements of the Nanbeige Models Community License Agreement by contacting nanbeige@126.com. After review, We will grant you a non-exclusive, worldwide, non-transferable, non-sublicensable and revocable commercial copyright license.
