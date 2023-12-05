---
language:
- en
library_name: transformers
license: llama2
datasets:
- Open-Orca/OpenOrca
- openchat/openchat_sharegpt4_dataset
- LDJnr/Puffin
- ehartford/samantha-data
- OpenAssistant/oasst1
- jondurbin/airoboros-gpt4-1.4.1
---


# FashionGPT-V1.1

### Introduction

This is a llama-2-70B model combined with multiple adapters via appropriate methods.

<br>

### Dataset

Here is the list of datasets used:

* Orca-style 40K dataset. This dataset is a filtered subset of [OpenOrca-GPT4](<https://huggingface.co/datasets/Open-Orca/OpenOrca/blob/main/1M-GPT4-Augmented.parquet>) and [airoboros-gpt4-1.4.1](<https://huggingface.co/datasets/jondurbin/airoboros-gpt4-1.4.1>).
* [Samantha](<https://huggingface.co/datasets/ehartford/samantha-data>) made by Eric Hartford and cleaned by us, about 6.5K samples.
* [oasst1](<https://huggingface.co/datasets/OpenAssistant/oasst1>) cleaned by us, containing about 80K samples.
* Misconception data generated using misconception data generator in [airoboros_repo](<https://github.com/jondurbin/airoboros>), about 0.5K samples.
* GPT-4 Multi-turn Conversations. This dataset is a filtered mixture of [openchat sharegpt4](https://huggingface.co/datasets/openchat/openchat_sharegpt4_dataset/) and [Puffin](<https://huggingface.co/datasets/LDJnr/Puffin>), containing about 8K samples.

<br>

### Training

* We train our adapters with [jondurbin's forked QLoRA repo](<https://github.com/jondurbin/qlora>)
* We add multi-turn conversational data support from [fastchat](<https://github.com/lm-sys/FastChat/blob/main/fastchat/train/train.py>), with minor modifications.
* We use bash shell script similar to [airoboros-70b-gpt4-1.4.1](<https://gist.github.com/jondurbin/87fc040b92a3073125ed516b04bc6e19>) to train our adapters.
* We combine multiple adapters to llama-2-70B with more novel strategy than our [v1 model](<https://huggingface.co/ICBU-NPU/FashionGPT-70B-V1>). The details of combining multiple adapters will be unveiled in our upcoming paper.

<br>

### Prompt Template

```
### System:
{System}
### User:
{User}
### Assistant:
{Assistant}
```

<br>

### Evaluation

| Metric                | Value |
|-----------------------|-------|
| ARC (25-shot)         | 71.76 |
| HellaSwag (10-shot)   | 88.20 |
| MMLU (5-shot)         | 70.99 |
| TruthfulQA (0-shot)   | 65.26 |
| Avg.                  | 74.05 |

<br>

### license disclaimer

This model is bound by the license & usage restrictions of the original Llama-2 model. And comes with no warranty or gurantees of any kind.

<br>

### Limitations & Biases

Llama 2 and fine-tuned variants are a new technology that carries risks with use. Testing conducted to date has been in English, and has not covered, nor could it cover all scenarios. For these reasons, as with all LLMs, Llama 2 and any fine-tuned varient's potential outputs cannot be predicted in advance, and the model may in some instances produce inaccurate, biased or other objectionable responses to user prompts. Therefore, before deploying any applications of Llama 2 variants, developers should perform safety testing and tuning tailored to their specific applications of the model.

Please see the Responsible Use Guide available at <https://ai.meta.com/llama/responsible-use-guide/>

<br>

### Citiation:

* airoboros: <https://github.com/jondurbin/airoboros>
* samantha: <https://erichartford.com/meet-samantha>
```bibtex
@misc{mukherjee2023orca,
      title={Orca: Progressive Learning from Complex Explanation Traces of GPT-4}, 
      author={Subhabrata Mukherjee and Arindam Mitra and Ganesh Jawahar and Sahaj Agarwal and Hamid Palangi and Ahmed Awadallah},
      year={2023},
      eprint={2306.02707},
      archivePrefix={arXiv},
      primaryClass={cs.CL}
}
```
```bibtex
@article{dettmers2023qlora,
  title={QLoRA: Efficient Finetuning of Quantized LLMs},
  author={Dettmers, Tim and Pagnoni, Artidoro and Holtzman, Ari and Zettlemoyer, Luke},
  journal={arXiv preprint arXiv:2305.14314},
  year={2023}
}
```
```bibtex
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
```bibtex
@software{openchat,
  title = {{OpenChat: Advancing Open-source Language Models with Imperfect Data}},
  author = {Wang, Guan and Cheng, Sijie and Yu, Qiying and Liu, Changling},
  doi = {10.5281/zenodo.8105775},
  url = {https://github.com/imoneoi/openchat},
  version = {pre-release},
  year = {2023},
  month = {7},
}
```