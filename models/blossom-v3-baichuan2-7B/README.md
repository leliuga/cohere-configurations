---
license: apache-2.0
datasets:
- Azure99/blossom-chat-v1
- Azure99/blossom-math-v2
- Azure99/blossom-wizard-v1
- Azure99/blossom-orca-v1
language:
- zh
- en
pipeline_tag: text-generation
---
# **BLOSSOM-v3-baichuan2-7b**

### 介绍

Blossom是一个对话式语言模型，基于Baichuan2-7B-Base预训练模型，在Blossom Orca/Wizard/Chat/Math混合数据集上进行指令精调得来。Blossom拥有强大的通用能力及上下文理解能力，此外，训练使用的高质量中英文数据集也进行了开源。

训练分为两阶段，第一阶段使用100K Wizard、100K Orca单轮指令数据集，训练1个epoch；第二阶段使用2K Blossom math数学推理数据集、50K Blossom chat多轮对话数据集、以及上一阶段中随机采样1%的数据，训练3个epoch。

### 推理

推理采用对话续写的形式。

单轮对话

```
A chat between a human and an artificial intelligence bot. The bot gives helpful, detailed, and polite answers to the human's questions.
|Human|: 你好
|Bot|: 
```

多轮对话

```
A chat between a human and an artificial intelligence bot. The bot gives helpful, detailed, and polite answers to the human's questions.
|Human|: 你好
|Bot|: 你好，有什么我能帮助你的？</s>
|Human|: 介绍下中国的首都吧
|Bot|: 
```

注意：在历史对话的Bot输出结尾，拼接一个&lt;/s&gt;