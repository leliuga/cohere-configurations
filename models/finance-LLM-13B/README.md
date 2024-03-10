---
language:
- en
datasets:
- Open-Orca/OpenOrca
- GAIR/lima
- WizardLM/WizardLM_evol_instruct_V2_196k
metrics:
- accuracy
pipeline_tag: text-generation
tags:
- finance
---

# Domain Adaptation of Large Language Models
This repo contains the domain-specific base model developed from **LLaMA-1-13B**, using the method in our **ICLR 2024** paper [Adapting Large Language Models via Reading Comprehension](https://huggingface.co/papers/2309.09530).

We explore **continued pre-training on domain-specific corpora** for large language models. While this approach enriches LLMs with domain knowledge, it significantly hurts their prompting ability for question answering. Inspired by human learning via reading comprehension, we propose a simple method to **transform large-scale pre-training corpora into reading comprehension texts**, consistently improving prompting performance across tasks in biomedicine, finance, and law domains. **Our 7B model competes with much larger domain-specific models like BloombergGPT-50B**. 

### 🤗 We are currently working hard on developing models across different domains, scales and architectures! Please stay tuned! 🤗

**************************** **Updates** ****************************
* 2024/1/16: 🎉 Our [research paper](https://huggingface.co/papers/2309.09530) has been accepted by ICLR 2024!!!🎉
* 2023/12/19: Released our [13B base models](https://huggingface.co/AdaptLLM/law-LLM-13B) developed from LLaMA-1-13B.
* 2023/12/8: Released our [chat models](https://huggingface.co/AdaptLLM/law-chat) developed from LLaMA-2-Chat-7B.
* 2023/9/18: Released our [paper](https://huggingface.co/papers/2309.09530), [code](https://github.com/microsoft/LMOps), [data](https://huggingface.co/datasets/AdaptLLM/law-tasks), and [base models](https://huggingface.co/AdaptLLM/law-LLM) developed from LLaMA-1-7B.


## Domain-Specific LLaMA-1
### LLaMA-1-7B
In our paper, we develop three domain-specific models from LLaMA-1-7B, which are also available in Huggingface: [Biomedicine-LLM](https://huggingface.co/AdaptLLM/medicine-LLM), [Finance-LLM](https://huggingface.co/AdaptLLM/finance-LLM) and [Law-LLM](https://huggingface.co/AdaptLLM/law-LLM), the performances of our AdaptLLM compared to other domain-specific LLMs are:

<p align='center'>
    <img src="https://cdn-uploads.huggingface.co/production/uploads/650801ced5578ef7e20b33d4/6efPwitFgy-pLTzvccdcP.png" width="700">
</p>

### LLaMA-1-13B
Moreover, we scale up our base model to LLaMA-1-13B to see if **our method is similarly effective for larger-scale models**, and the results are consistently positive too: [Biomedicine-LLM-13B](https://huggingface.co/AdaptLLM/medicine-LLM-13B), [Finance-LLM-13B](https://huggingface.co/AdaptLLM/finance-LLM-13B) and [Law-LLM-13B](https://huggingface.co/AdaptLLM/law-LLM-13B).

## Domain-Specific LLaMA-2-Chat
Our method is also effective for aligned models! LLaMA-2-Chat requires a [specific data format](https://huggingface.co/blog/llama2#how-to-prompt-llama-2), and our **reading comprehension can perfectly fit the data format** by transforming the reading comprehension into a multi-turn conversation. We have also open-sourced chat models in different domains: [Biomedicine-Chat](https://huggingface.co/AdaptLLM/medicine-chat), [Finance-Chat](https://huggingface.co/AdaptLLM/finance-chat) and [Law-Chat](https://huggingface.co/AdaptLLM/law-chat)

For example, to chat with the finance model:
```python
from transformers import AutoModelForCausalLM, AutoTokenizer

model = AutoModelForCausalLM.from_pretrained("AdaptLLM/finance-LLM-13B")
tokenizer = AutoTokenizer.from_pretrained("AdaptLLM/finance-LLM-13B", use_fast=False)

# Put your input here:
user_input = '''Use this fact to answer the question: Title of each class Trading Symbol(s) Name of each exchange on which registered
Common Stock, Par Value $.01 Per Share MMM New York Stock Exchange
MMM Chicago Stock Exchange, Inc.
1.500% Notes due 2026 MMM26 New York Stock Exchange
1.750% Notes due 2030 MMM30 New York Stock Exchange
1.500% Notes due 2031 MMM31 New York Stock Exchange

Which debt securities are registered to trade on a national securities exchange under 3M's name as of Q2 of 2023?'''

# Simply use your input as the prompt for base models
prompt = user_input

inputs = tokenizer(prompt, return_tensors="pt", add_special_tokens=False).input_ids.to(model.device)
outputs = model.generate(input_ids=inputs, max_length=2048)[0]

answer_start = int(inputs.shape[-1])
pred = tokenizer.decode(outputs[answer_start:], skip_special_tokens=True)

print(f'### User Input:\n{user_input}\n\n### Assistant Output:\n{pred}')
```

## Domain-Specific Tasks
To easily reproduce our results, we have uploaded the filled-in zero/few-shot input instructions and output completions of each domain-specific task: [biomedicine-tasks](https://huggingface.co/datasets/AdaptLLM/medicine-tasks), [finance-tasks](https://huggingface.co/datasets/AdaptLLM/finance-tasks), and [law-tasks](https://huggingface.co/datasets/AdaptLLM/law-tasks).

**Note:** those filled-in instructions are specifically tailored for models before alignment and do NOT fit for the specific data format required for chat models.

## Citation
If you find our work helpful, please cite us:
```bibtex
@article{adaptllm,
  title={Adapting large language models via reading comprehension},
  author={Cheng, Daixuan and Huang, Shaohan and Wei, Furu},
  journal={arXiv preprint arXiv:2309.09530},
  year={2023}
}
```