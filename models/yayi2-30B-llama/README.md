---
language:
- zh
- en
license: other
model-index:
- name: yayi2-30b-llama
  results:
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: AI2 Reasoning Challenge (25-Shot)
      type: ai2_arc
      config: ARC-Challenge
      split: test
      args:
        num_few_shot: 25
    metrics:
    - type: acc_norm
      value: 35.67
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=cognitivecomputations/yayi2-30b-llama
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: HellaSwag (10-Shot)
      type: hellaswag
      split: validation
      args:
        num_few_shot: 10
    metrics:
    - type: acc_norm
      value: 53.37
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=cognitivecomputations/yayi2-30b-llama
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: MMLU (5-Shot)
      type: cais/mmlu
      config: all
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 70.6
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=cognitivecomputations/yayi2-30b-llama
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: TruthfulQA (0-shot)
      type: truthful_qa
      config: multiple_choice
      split: validation
      args:
        num_few_shot: 0
    metrics:
    - type: mc2
      value: 49.08
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=cognitivecomputations/yayi2-30b-llama
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: Winogrande (5-shot)
      type: winogrande
      config: winogrande_xl
      split: validation
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 63.14
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=cognitivecomputations/yayi2-30b-llama
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: GSM8k (5-shot)
      type: gsm8k
      config: main
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 18.88
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=cognitivecomputations/yayi2-30b-llama
      name: Open LLM Leaderboard
---


This is [wenge-research/yayi2-30b](https://huggingface.co/wenge-research/yayi2-30b) converted to llama compatible format.

Subject to the [Yayi 2 license](https://github.com/wenge-research/YAYI2/blob/main/COMMUNITY_LICENSE).

Join our Discord! https://discord.gg/cognitivecomputations 

Brought to you by @Weyaxi and @ehartford, with thanks to @chargoddard for the pioneering work and the consultation! 

And of course thanks to the yayi2 team for sharing an amazing model.

Original card below:

<div align="center">
<h1>
  YAYI 2
</h1>
<!-- <br> -->
</div>

<div align="center">
<a href="https://github.com/wenge-research/YAYI2" target="_blank">GitHub</a> | <a href="https://yayi.wenge.com" target="_blank">雅意大模型</a> 
</div>



## 介绍/Introduction
YAYI 2 是中科闻歌研发的开源大语言模型，包括 Base 和 Chat 版本，参数规模为 30B。YAYI2-30B 是基于 Transformer 的大语言模型，采用了 2.65 万亿 Tokens 的高质量、多语言语料进行预训练。针对通用和特定领域的应用场景，我们采用了百万级指令进行微调，同时借助人类反馈强化学习方法，以更好地使模型与人类价值观对齐。

本次开源的模型为 YAYI2-30B Base 模型。如果您想了解更多关于 YAYI 2 模型的细节，我们建议您参阅 [GitHub](https://github.com/wenge-research/YAYI2) 仓库。更多技术细节，敬请期待我们的技术报告🔥。



YAYI 2 is a collection of open-source large language models launched by Wenge Technology. YAYI2-30B is a Transformer-based large language model, and has been pretrained for 2.65 trillion tokens of multilingual data with high quality. The base model is aligned with human values through supervised fine-tuning with millions of instructions and reinforcement learning from human feedback (RLHF). 

We opensource the pre-trained language model in this release, namely **YAYI2-30B**. For more details about the YAYI 2, please refer to our  [GitHub](https://github.com/wenge-research/YAYI2)  repository. Stay tuned for more technical details in our upcoming technical report! 🔥


## 模型细节/Model Details

| Hyperparameter| Value  | 
|:----------|:----------:| 
| n_layers | 64    | 
| n_heads | 64    | 
| hidden_size | 7168    | 
| vocab_size | 81920    | 
| sequence length | 4096    | 


## 要求/Requirements

* python 3.8及以上版本
* pytorch 2.0.1 及以上版本
* 建议使用 CUDA 11.7 及以上版本
* 运行 BF16 或 FP16 模型需要至少80GB显存（例如1xA100）


* python 3.8 and above
* pytorch 2.0.1 and above
* CUDA 11.7 and above are recommended
* To run YAYI2-30B in bf16/fp16, at least 80B GPU memory is required (e.g., 1xA100-80G)


## 快速开始/Quick Start

```python
>>> from transformers import AutoModelForCausalLM, AutoTokenizer
>>> tokenizer = AutoTokenizer.from_pretrained("wenge-research/yayi2-30b", trust_remote_code=True)
>>> model = AutoModelForCausalLM.from_pretrained("wenge-research/yayi2-30b", device_map="auto", trust_remote_code=True)
>>> inputs = tokenizer('The winter in Beijing is', return_tensors='pt')
>>> inputs = inputs.to('cuda')
>>> pred = model.generate(
        **inputs, 
        max_new_tokens=256, 
        eos_token_id=tokenizer.eos_token_id, 
        do_sample=True,
        repetition_penalty=1.2,
        temperature=0.4, 
        top_k=100, 
        top_p=0.8
        )
>>> print(tokenizer.decode(pred.cpu()[0], skip_special_tokens=True))
```


## 评测结果/Evaluation

我们在多个基准数据集上进行了评测，包括 C-Eval、MMLU、 CMMLU、AGIEval、GAOKAO-Bench、GSM8K、MATH、BBH、HumanEval 以及 MBPP。我们考察了模型在语言理解、学科知识、数学推理、逻辑推理以及代码生成方面的表现。YAYI 2 模型在与其规模相近的开源模型中展现出了显著的性能提升。

We evaluate our model on standard benchmarks, including C-Eval, MMLU, CMMLU, AGIEval, GAOKAO-Bench, GSM8K, MATH, BBH, HumanEval, and MBPP. Our goal is to assess the model's performance in language comprehension, knowledge comprehension, mathematical reasoning, logical reasoning, and code generation.  YAYI 2 has demonstrated exceptional performance across models with similar size.

<table id="myTable">
  <!-- Table header -->
  <tr>
        <th></th>
        <th colspan="5" style="text-align: center;">Knowledge</th>
        <th colspan="2" style="text-align: center;">Math</th>
        <th colspan="1" style="text-align: center;">Logic reasonning</th>
        <th colspan="2" style="text-align: center;">Code</th>
  </tr>
  <tr>
        <th style="text-align: left;">Model</th>
        <th>C-Eval(val)</th>
        <th>MMLU</th>
        <th>AGIEval</th>
        <th>CMMLU</th>
        <th>GAOKAO-Bench</th>
        <th>GSM8K</th>
        <th>MATH</th>
        <th>BBH</th>
        <th>HumanEval</th>
        <th>MBPP</th>
  </tr>
  <tr>
        <td></td>
        <td style="text-align: center;">5-shot</td>
        <td style="text-align: center;">5-shot</td>
        <td style="text-align: center;">3/0-shot</td>
        <td style="text-align: center;">5-shot</td>
        <td style="text-align: center;">0-shot</td>
        <td style="text-align: center;">8/4-shot</td>
        <td style="text-align: center;">4-shot</td>
        <td style="text-align: center;">3-shot</td>
        <td style="text-align: center;">0-shot</td>
        <td style="text-align: center;">3-shot</td>
        </tr>
        <tr>
        <td><strong>MPT-30B</strong></td>
        <td style="text-align: center;">-</td>
        <td style="text-align: center;">46.9</td>
        <td style="text-align: center;">33.8</td>
        <td style="text-align: center;">-</td>
        <td style="text-align: center;">-</td>
        <td style="text-align: center;">15.2</td>
        <td style="text-align: center;">3.1</td>
        <td style="text-align: center;">38.0</td>
        <td style="text-align: center;">25.0</td>
        <td style="text-align: center;">32.8</td>
  </tr>
  <tr>
        <td><strong>Falcon-40B</strong></td>
        <td style="text-align: center;">-</td>
        <td style="text-align: center;">55.4</td>
        <td style="text-align: center;">37.0</td>
        <td style="text-align: center;">-</td>
        <td style="text-align: center;">-</td>
        <td style="text-align: center;">19.6</td>
        <td style="text-align: center;">5.5</td>
        <td style="text-align: center;">37.1</td>
        <td style="text-align: center;">0.6</td>
        <td style="text-align: center;">29.8</td>
  </tr>
  <tr>
        <td><strong>LLaMA2-34B</strong></td>
        <td style="text-align: center;">-</td>
        <td style="text-align: center;">62.6</td>
        <td style="text-align: center;">43.4</td>
        <td style="text-align: center;">-</td>
        <td style="text-align: center;">-</td>
        <td style="text-align: center;">42.2</td>
        <td style="text-align: center;">6.2</td>
        <td style="text-align: center;">44.1</td>
        <td style="text-align: center;">22.6</td>
        <td style="text-align: center;">33.0</td>
  </tr>
  <tr>
        <td><strong>Baichuan2-13B</strong></td>
        <td style="text-align: center;">59.0</td>
        <td style="text-align: center;">59.5</td>
        <td style="text-align: center;">37.4</td>
        <td style="text-align: center;">61.3</td>
        <td style="text-align: center;">45.6</td>
        <td style="text-align: center;">52.6</td>
        <td style="text-align: center;">10.1</td>
        <td style="text-align: center;">49.0</td>
        <td style="text-align: center;">17.1</td>
        <td style="text-align: center;">30.8</td>
  </tr>
  <tr>
        <td><strong>Qwen-14B</strong></td>
        <td style="text-align: center;">71.7</td>
        <td style="text-align: center;">67.9</td>
        <td style="text-align: center;">51.9</td>
        <td style="text-align: center;">70.2</td>
        <td style="text-align: center;">62.5</td>
        <td style="text-align: center;">61.6</td>
        <td style="text-align: center;">25.2</td>
        <td style="text-align: center;">53.7</td>
        <td style="text-align: center;">32.3</td>
        <td style="text-align: center;">39.8</td>
  </tr>
  <tr>
        <td><strong>InternLM-20B</strong></td>
        <td style="text-align: center;">58.8</td>
        <td style="text-align: center;">62.1</td>
        <td style="text-align: center;">44.6</td>
        <td style="text-align: center;">59.0</td>
        <td style="text-align: center;">45.5</td>
        <td style="text-align: center;">52.6</td>
        <td style="text-align: center;">7.9</td>
        <td style="text-align: center;">52.5</td>
        <td style="text-align: center;">25.6</td>
        <td style="text-align: center;">35.6</td>
  </tr>
  <tr>
        <td><strong>Aquila2-34B</strong></td>
        <td style="text-align: center;">98.5</td>
        <td style="text-align: center;">76.0</td>
        <td style="text-align: center;">43.8</td>
        <td style="text-align: center;">78.5</td>
        <td style="text-align: center;">37.8</td>
        <td style="text-align: center;">50.0</td>
        <td style="text-align: center;">17.8</td>
        <td style="text-align: center;">42.5</td>
        <td style="text-align: center;">0.0</td>
        <td style="text-align: center;">41.0</td>
  </tr>
  <tr>
        <td><strong>Yi-34B</strong></td>
        <td style="text-align: center;">81.8</td>
        <td style="text-align: center;">76.3</td>
        <td style="text-align: center;">56.5</td>
        <td style="text-align: center;">82.6</td>
        <td style="text-align: center;">68.3</td>
        <td style="text-align: center;">67.6</td>
        <td style="text-align: center;">15.9</td>
        <td style="text-align: center;">66.4</td>
        <td style="text-align: center;">26.2</td>
        <td style="text-align: center;">38.2</td>
  </tr>
  <tr>
        <td><strong>YAYI2-30B</strong></td>
        <td style="text-align: center;">80.9</td>
        <td style="text-align: center;"><b>80.5</b></td>
        <td style="text-align: center;"><b>62.0</b></td>
        <td style="text-align: center;"><b>84.0</b></td>
        <td style="text-align: center;">64.4</td>
        <td style="text-align: center;"><b>71.2</b></td>
        <td style="text-align: center;">14.8</td>
        <td style="text-align: center;">54.5</td>
        <td style="text-align: center;"><b>53.1</b></td>
        <td style="text-align: center;"><b>45.8</b></td>
  </tr>
</table>


我们使用 [OpenCompass Github 仓库](https://github.com/open-compass/opencompass) 提供的源代码进行了评测。对于对比模型，我们列出了他们在 [OpenCompass](https://opencompass.org.cn) 榜单上的评测结果，截止日期为 2023年12月15日。对于其他尚未在 [OpenCompass](https://opencompass.org.cn/leaderboard-llm) 平台参与评测的模型，包括 MPT、Falcon 和 LLaMa 2，我们采用了 [LLaMA 2](https://arxiv.org/abs/2307.09288) 报告的结果。

We evaluate our model using the source code from the [OpenCompass Github repository](https://github.com/open-compass/opencompass). If available, we report results for comparative models assessed by OpenCompass with the evaluation reference date set to Dec. 15th, 2013. For MPT, Falcon, and Llama, which have not been evaluated by OpenCompass, we use the results reported in the [LLaMA 2](https://arxiv.org/abs/2307.09288) paper.



## 协议/License

本项目中的代码依照 [Apache-2.0](https://github.com/wenge-research/YAYI2/blob/main/LICENSE) 协议开源，社区使用 YAYI 2 模型和数据需要遵循[雅意YAYI 2 模型社区许可协议](https://github.com/wenge-research/YAYI2/blob/main/COMMUNITY_LICENSE)。若您需要将雅意 YAYI 2系列模型或其衍生品用作商业用途，请根据[《雅意 YAYI 2 模型商用许可协议》](https://github.com/wenge-research/YAYI2/blob/main/COMMERCIAL_LICENSE)将商用许可申请登记信息发送至指定邮箱 yayi@wenge.com。审核通过后，雅意将授予您商用版权许可，请遵循协议中的商业许可限制。

The code in this project is open-sourced under the [Apache-2.0](https://github.com/wenge-research/YAYI2/blob/main/LICENSE) license. The use of YaYi series model weights and data must adhere to the [YAYI 2 Community License](https://github.com/wenge-research/YAYI2/blob/main/COMMUNITY_LICENSE). If you intend to use the YAYI 2 series models or their derivatives for commercial purposes, please submit your commercial license application and registration information to yayi@wenge.com, following the [YAYI 2 Commercial License](https://github.com/wenge-research/YAYI2/blob/main/COMMERCIAL_LICENSE). Upon approval, YAYI will grant you a commercial copyright license, subject to the commercial license restrictions outlined in the agreement.



## 引用/Citation

如果您在工作中使用了我们的模型，请引用我们的论文。

If you are using the resource for your work, please cite our paper.

```
@article{YAYI 2,
  author    = {Yin Luo, Qingchao Kong, Nan Xu, et.al.},
  title     = {YAYI 2: Multilingual Open Source Large Language Models},
  journal   = {arXiv preprint arXiv},
  year      = {2023}
}
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_cognitivecomputations__yayi2-30b-llama)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |48.46|
|AI2 Reasoning Challenge (25-Shot)|35.67|
|HellaSwag (10-Shot)              |53.37|
|MMLU (5-Shot)                    |70.60|
|TruthfulQA (0-shot)              |49.08|
|Winogrande (5-shot)              |63.14|
|GSM8k (5-shot)                   |18.88|

