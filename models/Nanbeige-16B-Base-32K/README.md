---
license: other
language:
- en
- zh
library_name: transformers
pipeline_tag: text-generation
tags:
- llm
- nanbeige
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
  Nanbeige-16B-Base-32k
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

本仓库为 Nanbeige-16B-Base-32k 的模型仓库。

Nanbeige-16B is a 16 billion parameter language model developed by Nanbeige LLM Lab. It uses 2.5T Tokens for pre-training. The training data includes a large amount of high-quality internet corpus, various books, code, etc. It has achieved good results on various authoritative evaluation data sets. This release includes the Base, Chat, Base-32k and Chat-32k.

The Base-32k version is based on the Nanbeige-16B-Base model, which uses YaRN interpolation method to modify the position encoding, and performs full parameter incremental pre-training with 20 billion tokens of Chinese, English, and code corpora, under 32k max length.

The Chat version and Chat-32k version are based on the Nanbeige-16B-Base model and Nanbeige-16B-Base-32k model respectively. They have undergone extensive human-aligned training.

If you need to deal with longer contexts, we recommend using Nanbeige-16B-Base-32k and Nanbeige-16B-Chat-32k.

This repository is the one for Nanbeige-16B-Base-32k model.

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

- deepspeed 0.11.1及以上版本

可以通过以下pip命令安装相关依赖库

You can install the dependent libraries with the following pip command

```
pip install transformers==4.33.3 transformers_stream_generator deepspeed einops==0.3.2 datasets==2.10.0 deepspeed==0.11.1
```

## 推理代码

通过以下代码可以调用模型进行续写生成：

The following code can be used to invoke the model for text generation:

```python
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer
from transformers.generation.utils import GenerationConfig
import deepspeed
import os

tokenizer = AutoTokenizer.from_pretrained("Nanbeige/Nanbeige-16B-Base-32k", use_fast=False, trust_remote_code=True)
model = AutoModelForCausalLM.from_pretrained("Nanbeige/Nanbeige-16B-Base-32k", device_map="auto", torch_dtype=torch.bfloat16, trust_remote_code=True)
world_size = int(os.getenv('WORLD_SIZE', '1'))
model = deepspeed.init_inference(model.eval(),
                    dtype="bfloat16",
                    replace_with_kernel_inject=False,
                    mp_size=world_size)
inputs = tokenizer('中国的首都是北京\n德国的首都是柏林\n孟加拉的首都是', return_tensors='pt')
inputs = inputs.to(model.device)
pred = model.generate(**inputs, max_new_tokens=32000)
print(tokenizer.decode(pred.cpu()[0], skip_special_tokens=True))
```

##
# <span id="Evaluation">性能测评（Performance Evaluation）</span>

### 长文本理解
我们使用LongBench的LSHT、LCC、Multifiled_QA_ZH数据集，对 Nanbeige-16B-Base-32k 模型进行了测评，相较具有长文本理解能力的同参数规模Base模型取得了不错的结果。

We evaluated the Nanbeige-16B-Base-32k model using LSHT, LCC, and Multifiled_QA_ZH from LongBench datasets. Compared to the Base model of the same parameter size with long-context comprehension capabilities, it achieved impressive results.

|                            |  LSHT (分类)  |  LCC (代码) | Multifiled_QA_ZH (问答) |
|----------------------------|--------|-------|------------------|
| Chinese-Llama2-13B-16k     |  31.0  |  9.6  |       36.4       |
| Qwen-14B-Dynamnic_ntk-Logn |  16.0  |  66.7 |       30.0       | 
| Nanbeige-16B-Base-32k      |  40.3  |  70.7 |       47.2       |


##
# <span id="Limitations">局限性（Limitations）</span>

虽然我们在训练过程中非常注重模型的安全性，力求确保其输出符合伦理和法律要求的文本，但由于模型大小和概率生成范式的限制，无法完全避免产生各种不符合预期的输出情况。这些输出可能包含偏见、歧视等有害内容，请勿传播这些内容。我们不承担因传播不良信息而导致的任何后果。

While we place great emphasis on the safety of the model during the training process, striving to ensure that its outputs align with ethical and legal requirements, it may not completely avoid generating unexpected outputs due to the model's size and probabilistic nature. These outputs may include harmful content such as bias or discrimination. Please don't propagate such content. We do not assume any responsibility for the consequences resulting from the dissemination of inappropriate information.

# <span id="License">协议（License）</span>

使用 Nanbeige 模型时，您必须遵守 Apache 2.0 许可证和[《南北阁大语言模型许可协议》](https://huggingface.co/Nanbeige/Nanbeige-16B-Base-32k/resolve/main/南北阁大语言模型许可协议.pdf)。如果您打算将 Nanbeige 模型或其衍生产品用于商业目的，请通过以下联系邮箱 nanbeige@126.com 提交申请材料，以满足《南北阁大语言模型许可协议》的要求。经过审核后，我们将授予您非排他性、全球范围内、不可转让、不可再许可、可撤销的商业版权许可。

When using the Nanbeige models, you must comply with the Apache 2.0 License and the [License Agreement for Large Language Models Nanbeige](https://huggingface.co/Nanbeige/Nanbeige-16B-Base-32k/resolve/main/License_Agreement_for_Large_Language_Models_Nanbeige.pdf). If you intend to use the Nanbeige Models or its derivatives for commercial purposes, please submit application materials to meet the requirements of the Nanbeige Models Community License Agreement by contacting nanbeige@126.com. After review, We will grant you a non-exclusive, worldwide, non-transferable, non-sublicensable and revocable commercial copyright license.
