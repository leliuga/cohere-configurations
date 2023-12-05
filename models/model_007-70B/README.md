---
language:
- en
library_name: transformers
license: llama2
datasets:
- pankajmathur/orca_mini_v1_dataset
- pankajmathur/dolly-v2_orca
- pankajmathur/WizardLM_Orca
- pankajmathur/alpaca_orca
- ehartford/dolphin
---


# model_007

A hybrid (explain + instruct) style Llama2-70b model, Pleae check examples below for both style prompts, Here is the list of datasets used:

* Open-Platypus
* Alpaca
* WizardLM
* Dolly-V2
* Dolphin Samples (~200K)
* Orca_minis_v1
* Alpaca_orca
* WizardLM_orca
* Dolly-V2_orca


<br>

**P.S. If you're interested to collaborate, please connect with me at www.linkedin.com/in/pankajam.**

<br>



### quantized versions
Huge respect to @TheBloke, here are the GGML/GPTQ/GGUF versions, go crazy :)

https://huggingface.co/TheBloke/model_007-70B-GGML

https://huggingface.co/TheBloke/model_007-70B-GGUF

https://huggingface.co/TheBloke/model_007-70B-GPTQ

<br>

#### license disclaimer:

This model is bound by the license & usage restrictions of the original Llama-2 model. And comes with no warranty or gurantees of any kind.

<br>

## Evaluation

We evaluated model_007 on a wide range of tasks using [Language Model Evaluation Harness](https://github.com/EleutherAI/lm-evaluation-harness) from EleutherAI. 

Here are the results on metrics used by [HuggingFaceH4 Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)

|||
|:------:|:--------:|
|**Task**|**Value**|
|*ARC*|0.7108|
|*HellaSwag*|0.8765|
|*MMLU*|0.6904|
|*TruthfulQA*|0.6312|
|*Winogrande*|0.8335|
|*GSM8K*|0.3715|
|*DROP*|0.3105|
|**Total Average**|**0.6320**|


<br>

## Prompt Format

Here is the Orca prompt format

```
### System:
You are an AI assistant that follows instruction extremely well. Help as much as you can.

### User:
Tell me about Orcas.

### Assistant:

```

Here is the Alpaca prompt format

```

### User:
Tell me about Alpacas.

### Assistant:

```

#### OobaBooga Instructions:

This model required upto 45GB GPU VRAM in 4bit so it can be loaded directly on Single RTX 6000/L40/A40/A100/H100 GPU or Double RTX 4090/L4/A10/RTX 3090/RTX A5000
So, if you have access to Machine with 45GB GPU VRAM and have installed [OobaBooga Web UI](https://github.com/oobabooga/text-generation-webui) on it.
You can just download this model by using HF repo link directly on OobaBooga Web UI "Model" Tab/Page & Just use **load-in-4bit** option in it.

![model_load_screenshot](https://huggingface.co/pankajmathur/model_101/resolve/main/oobabooga_model_load_screenshot.png)


After that go to Default Tab/Page on OobaBooga Web UI and **copy paste above prompt format into Input** and Enjoy!

![default_input_screenshot](https://huggingface.co/pankajmathur/model_101/resolve/main/default_input_screenshot.png)

<br>

#### Code Instructions:

Below shows a code example on how to use this model via Orca prompt

```python
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer, pipeline

tokenizer = AutoTokenizer.from_pretrained("psmathur/model_007")
model = AutoModelForCausalLM.from_pretrained(
  "psmathur/model_007",
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

Below shows a code example on how to use this model via Alpaca prompt

```python
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer, pipeline

tokenizer = AutoTokenizer.from_pretrained("psmathur/model_007")
model = AutoModelForCausalLM.from_pretrained(
  "psmathur/model_007",
  torch_dtype=torch.float16,
  load_in_8bit=True,
  low_cpu_mem_usage=True,
  device_map="auto"
)
#generate text steps
instruction = "Tell me about Alpacas."
prompt = f"### User: {instruction}\n\n### Assistant:\n"
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
@misc{model_007,
  author = {Pankaj Mathur},
  title = {model_007: A hybrid (explain + instruct) style Llama2-70b model},
  year = {2023},
  publisher = {HuggingFace},
  journal = {HuggingFace repository},
  howpublished = {\url{https://https://huggingface.co/psmathur/model_007},
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
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_psmathur__model_007)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 63.2   |
| ARC (25-shot)         | 71.08          |
| HellaSwag (10-shot)   | 87.65    |
| MMLU (5-shot)         | 69.04         |
| TruthfulQA (0-shot)   | 63.12   |
| Winogrande (5-shot)   | 83.35   |
| GSM8K (5-shot)        | 37.15        |
| DROP (3-shot)         | 31.05         |
