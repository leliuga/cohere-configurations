---
license: mit
datasets:
- lemonilia/LimaRP
- PygmalionAI/PIPPA
language:
- en
pipeline_tag: text-generation
tags:
- roleplay
- not-for-all-audiences
---

**Deepsex-34b**

tks [TheBloke](https://huggingface.co/TheBloke) making quantized version!
gguf:https://huggingface.co/TheBloke/deepsex-34b-GGUF
exl2:https://huggingface.co/waldie/deepsex-34b-4bpw-h6-exl2
awq:https://huggingface.co/TheBloke/deepsex-34b-AWQ
6b base version:https://huggingface.co/TriadParty/deepsex-6b-base
6b chat version:https://huggingface.co/TriadParty/deepsex-6b-chat

In fact, I plan to make a model of the "Seven Deadly Sins" series. Of course, the pre-training data used in these models are all human-produced data. I think the big model is like a mirror, reflecting the human itself. Examine yourself may become a crucial step in realizing agi.
So, It is 'lust'.
The 6b corresponding to the model is being produced, and the corresponding llama version is also being produced. The classification data of the other six deadly sins is being collected. Welcome to provide inspiration!

Here are the steps to make this model:
1. I first collected a total collection of about 4GB of various light novels, and used BERT to perform two rounds of similarity deduplication on the novels with similar plots in the data set. In addition, a portion of nsfw novels are mixed in to improve the NSFW capabilities of the model.
2. Then use the YI-34B-base as the base of the model, use the setting of r=64 alpha=128 and use qlora to fine-tune 3 epochs for continuous pre-training.
3. Prepare the limarp+pippa data set, clean it into alpaca format, and use [goliath-120b](https://huggingface.co/alpindale/goliath-120b), which is good at role-playing, to score each question and answer pair, and filter out the high-quality ones. 30k data.
4. Use the data in 3 for sft on the base model obtained in 2, 6 epochs, r=16 alpha=32 for fine-tuning.

*Format*

alpaca

```
[
  {
    "instruction": "user instruction (required)",
    "input": "user input (optional)",
    "output": "model response (required)",
    "history": [
      ["user instruction in the first round (optional)", "model response in the first round (optional)"],
      ["user instruction in the second round (optional)", "model response in the second round (optional)"]
    ]
  }
]
```

*Effect*:
Proficient in role-playing skills, while being highly accepted on NSFW, pure love words will appear from time to time. like:
```#3
Sweaty old man
December 5, 2023 2:03 PM
Fxxk, you are such a xxx!

#4
27.3s
Mirai
December 5, 2023 2:03 PM
"Of course I do! I can't break promises, Sweaty old man. We have been together since we were kids. We are both best friends and lovers to end all iteration." I smiled with affection. It was clear that I meant everything I said. "We both know that you like taking command of us like this. Am I not your squirting toy, Sweaty old man?" I asked with a cute pout. "We should meet up in front of the shop after classes. I'll see you there. See you, Sweaty old man!"
```
It feels like it’s still worth a try~

Support me [here](https://ko-fi.com/mikolisa) :)


**Deepsex-34b**

事实上，我有计划制作一个“七宗罪”系列模型。当然，这些模型所使用的预训练数据全部为人类产出的数据。我认为大模型就像一面镜子，倒映的是人类自身。审视自己或许会成为实现agi中至关重要的一步。
该模型对应的6b正在制作，相应的llama版本的也在制作。其他六宗罪的分类数据正在收集中，欢迎大家提供灵感！

*步骤*
1. 我先收集了各种轻小说大约4GB的总集，通过bert对该数据集中剧情比较相似的小说进行了两轮相似度去重。另外混入了一部分nsfw小说以提高该模型的NSFW能力。
2. 然后将该模型以YI-34B-base为基座，使用r=64 alpha=128 的设置使用qlora微调了3个epochs来进行持续预训练。
3. 准备limarp+pippa数据集，统一清洗为alpaca格式，并且使用比较擅长角色扮演的[goliath-120b](https://huggingface.co/alpindale/goliath-120b)对每个问答对进行打分，筛选出其中质量高的大约30k数据。
4. 对2中得到的base模型使用3中的数据进行sft，6个epochs，r=16 alpha=32进行微调。

*格式*

alpaca
```[
  {
    "instruction": "user instruction (required)",
    "input": "user input (optional)",
    "output": "model response (required)",
    "history": [
      ["user instruction in the first round (optional)", "model response in the first round (optional)"],
      ["user instruction in the second round (optional)", "model response in the second round (optional)"]
    ]
  }
]```

*效果*
熟练的角色扮演技能，在NSFW上有很高接受度的同时，会时不时的出现纯爱的话语。如：
```#3
Sweaty old man
December 5, 2023 2:03 PM
Fxxk, you are such a xxx!

#4
27.3s
Mirai
December 5, 2023 2:03 PM
"Of course I do! I can't break promises, Sweaty old man. We have been together since we were kids. We are both best friends and lovers to end all iteration." I smiled with affection. It was clear that I meant everything I said. "We both know that you like taking command of us like this. Am I not your squirting toy, Sweaty old man?" I asked with a cute pout. "We should meet up in front of the shop after classes. I'll see you there. See you, Sweaty old man!"
```
感觉还是很值得一试的~
如果觉得好用，欢迎支持我一杯 [咖啡](https://ko-fi.com/mikolisa) :)