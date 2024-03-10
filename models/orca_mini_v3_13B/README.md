---
language:
- en
license: other
library_name: transformers
datasets:
- psmathur/orca_mini_v1_dataset
- ehartford/dolphin
pipeline_tag: text-generation
model-index:
- name: orca_mini_v3_13b
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
      value: 63.14
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=psmathur/orca_mini_v3_13b
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
      value: 82.35
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=psmathur/orca_mini_v3_13b
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
      value: 56.52
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=psmathur/orca_mini_v3_13b
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
      value: 51.81
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=psmathur/orca_mini_v3_13b
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
      value: 76.48
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=psmathur/orca_mini_v3_13b
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
      value: 13.12
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=psmathur/orca_mini_v3_13b
      name: Open LLM Leaderboard
---

# orca_mini_v3_13b

A Llama2-13b model trained on Orca Style datasets.


<br>

![orca-mini](https://huggingface.co/psmathur/orca_mini_v3_13b/resolve/main/orca_minis_small.jpeg)


<br>

**P.S. If you're interested to collaborate, please connect with me at www.linkedin.com/in/pankajam.**

<br>



### quantized versions

Big thanks to [@TheBloke](https://huggingface.co/TheBloke)

1) https://huggingface.co/TheBloke/orca_mini_v3_13B-GGML

2) https://huggingface.co/TheBloke/orca_mini_v3_13B-GPTQ


<br>
#### license disclaimer:

This model is bound by the license & usage restrictions of the original Llama-2 model. And comes with no warranty or gurantees of any kind.

<br>

## Evaluation

We evaluated orca_mini_v3_13b on a wide range of tasks using [Language Model Evaluation Harness](https://github.com/EleutherAI/lm-evaluation-harness) from EleutherAI. 

Here are the results on metrics used by [HuggingFaceH4 Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)

|||||
|:------:|:--------:|:-------:|:--------:|
|**Task**|**Metric**|**Value**|**Stderr**|
|*arc_challenge*|acc_norm|0.6314|0.0141|
|*hellaswag*|acc_norm|0.8242|0.0038|
|*mmlu*|acc_norm|0.5637|0.0351|
|*truthfulqa_mc*|mc2|0.5127|0.0157|
|**Total Average**|-|**0.6329877193**||


<br>

## Example Usage

Here is the prompt format

```
### System:
You are an AI assistant that follows instruction extremely well. Help as much as you can.

### User:
Tell me about Orcas.

### Assistant:

```

Below shows a code example on how to use this model

```python
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer, pipeline

tokenizer = AutoTokenizer.from_pretrained("psmathur/orca_mini_v3_13b")
model = AutoModelForCausalLM.from_pretrained(
  "psmathur/orca_mini_v3_13b",
  torch_dtype=torch.float16,
  load_in_8bit=True,
  low_cpu_mem_usage=True,
  device_map="auto"
)
system_prompt = "### System:\nYou are an AI assistant that follows instruction extremely well. Help as much as you can.\n\n"

#generate text steps
instruction = "Tell me about Orcas."
prompt = f"{system_prompt}### User: {instruction}\n\n### Assistant:\n"
inputs = tokenizer(prompt, return_tensors="pt").to("cuda")
output = model.generate(**inputs, do_sample=True, top_p=0.95, top_k=0, max_new_tokens=4096)

print(tokenizer.decode(output[0], skip_special_tokens=True))

```

<br>

#### Limitations & Biases:

While this model aims for accuracy, it can occasionally produce inaccurate or misleading results. 

Despite diligent efforts in refining the pretraining data, there remains a possibility for the generation of inappropriate, biased, or offensive content. 

Exercise caution and cross-check information when necessary.


<br>

### Citiation:

Please kindly cite using the following BibTeX:

```
@misc{orca_mini_v3_13b,
  author = {Pankaj Mathur},
  title = {orca_mini_v3_13b: An Orca Style Llama2-70b model},
  year = {2023},
  publisher = {HuggingFace},
  journal = {HuggingFace repository},
  howpublished = {\url{https://https://huggingface.co/psmathur/orca_mini_v3_13b},
}
```

```
@misc{mukherjee2023orca,
      title={Orca: Progressive Learning from Complex Explanation Traces of GPT-4}, 
      author={Subhabrata Mukherjee and Arindam Mitra and Ganesh Jawahar and Sahaj Agarwal and Hamid Palangi and Ahmed Awadallah},
      year={2023},
      eprint={2306.02707},
      archivePrefix={arXiv},
      primaryClass={cs.CL}
}
```

```
@software{touvron2023llama2,
  title={Llama 2: Open Foundation and Fine-Tuned Chat Models},
  author={Hugo Touvron, Louis Martin, Kevin Stone, Peter Albert, Amjad Almahairi, Yasmine Babaei, Nikolay Bashlykov, Soumya Batra, Prajjwal Bhargava,
 Shruti Bhosale, Dan Bikel, Lukas Blecher, Cristian Canton Ferrer, Moya Chen, Guillem Cucurull, David Esiobu, Jude Fernandes, Jeremy Fu, Wenyin Fu, Brian Fuller,
Cynthia Gao, Vedanuj Goswami, Naman Goyal, Anthony Hartshorn, Saghar Hosseini, Rui Hou, Hakan Inan, Marcin Kardas, Viktor Kerkez Madian Khabsa, Isabel Kloumann,
Artem Korenev, Punit Singh Koura, Marie-Anne Lachaux, Thibaut Lavril, Jenya Lee, Diana Liskovich, Yinghai Lu, Yuning Mao, Xavier Martinet, Todor Mihaylov,
Pushkar Mishra, Igor Molybog, Yixin Nie, Andrew Poulton, Jeremy Reizenstein, Rashi Rungta, Kalyan Saladi, Alan Schelten, Ruan Silva, Eric Michael Smith,
Ranjan Subramanian, Xiaoqing Ellen Tan, Binh Tang, Ross Taylor, Adina Williams, Jian Xiang Kuan, Puxin Xu , Zheng Yan, Iliyan Zarov, Yuchen Zhang, Angela Fan,
Melanie Kambadur, Sharan Narang, Aurelien Rodriguez, Robert Stojnic, Sergey Edunov, Thomas Scialom},
  year={2023}
}
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_psmathur__orca_mini_v3_13b)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 52.23   |
| ARC (25-shot)         | 63.14          |
| HellaSwag (10-shot)   | 82.35    |
| MMLU (5-shot)         | 56.52         |
| TruthfulQA (0-shot)   | 51.81   |
| Winogrande (5-shot)   | 76.48   |
| GSM8K (5-shot)        | 13.12        |
| DROP (3-shot)         | 22.23         |

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_psmathur__orca_mini_v3_13b)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |57.24|
|AI2 Reasoning Challenge (25-Shot)|63.14|
|HellaSwag (10-Shot)              |82.35|
|MMLU (5-Shot)                    |56.52|
|TruthfulQA (0-shot)              |51.81|
|Winogrande (5-shot)              |76.48|
|GSM8k (5-shot)                   |13.12|

