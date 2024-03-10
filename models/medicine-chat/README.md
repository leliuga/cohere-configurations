---
language:
- en
license: llama2
tags:
- biology
- medical
datasets:
- EleutherAI/pile
- Open-Orca/OpenOrca
- GAIR/lima
- WizardLM/WizardLM_evol_instruct_V2_196k
metrics:
- accuracy
pipeline_tag: text-generation
model-index:
- name: medicine-chat
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
      value: 53.75
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=AdaptLLM/medicine-chat
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
      value: 76.11
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=AdaptLLM/medicine-chat
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
      value: 49.98
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=AdaptLLM/medicine-chat
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
      value: 43.46
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=AdaptLLM/medicine-chat
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
      value: 75.69
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=AdaptLLM/medicine-chat
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
      value: 18.95
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=AdaptLLM/medicine-chat
      name: Open LLM Leaderboard
---

# Domain Adaptation of Large Language Models
This repo contains the domain-specific chat model developed from **LLaMA-2-Chat-7B**, using the method in our **ICLR 2024** paper [Adapting Large Language Models via Reading Comprehension](https://huggingface.co/papers/2309.09530).

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

For example, to chat with the biomedicine-chat model:
```python
from transformers import AutoModelForCausalLM, AutoTokenizer

model = AutoModelForCausalLM.from_pretrained("AdaptLLM/medicine-chat")
tokenizer = AutoTokenizer.from_pretrained("AdaptLLM/medicine-chat")

# Put your input here:
user_input = '''Question: Which of the following is an example of monosomy?
Options:
- 46,XX
- 47,XXX
- 69,XYY
- 45,X

Please provide your choice first and then provide explanations if possible.'''

# Apply the prompt template and system prompt of LLaMA-2-Chat demo for chat models (NOTE: NO prompt template is required for base models!)
our_system_prompt = "\nYou are a helpful, respectful and honest assistant. Always answer as helpfully as possible, while being safe.  Your answers should not include any harmful, unethical, racist, sexist, toxic, dangerous, or illegal content. Please ensure that your responses are socially unbiased and positive in nature.\n\nIf a question does not make any sense, or is not factually coherent, explain why instead of answering something not correct. If you don't know the answer to a question, please don't share false information.\n" # Please do NOT change this
prompt = f"<s>[INST] <<SYS>>{our_system_prompt}<</SYS>>\n\n{user_input} [/INST]"

# # NOTE:
# # If you want to apply your own system prompt, please integrate it into the instruction part following our system prompt like this:
# your_system_prompt = "Please, answer this question faithfully."
# prompt = f"<s>[INST] <<SYS>>{our_system_prompt}<</SYS>>\n\n{your_system_prompt}\n{user_input} [/INST]"

inputs = tokenizer(prompt, return_tensors="pt", add_special_tokens=False).input_ids.to(model.device)
outputs = model.generate(input_ids=inputs, max_length=4096)[0]

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
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_AdaptLLM__medicine-chat)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |52.99|
|AI2 Reasoning Challenge (25-Shot)|53.75|
|HellaSwag (10-Shot)              |76.11|
|MMLU (5-Shot)                    |49.98|
|TruthfulQA (0-shot)              |43.46|
|Winogrande (5-shot)              |75.69|
|GSM8k (5-shot)                   |18.95|

