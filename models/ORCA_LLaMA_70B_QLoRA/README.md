---
language:
- en
library_name: transformers
license: llama2
---


# Dolphin_ORCA_PlatyPus_LLaMA_70b

### Dataset
Here is the list of datasets used:
* Dolphin
* Open-Platypus
* OpenOrca

**mixed strategy: 100%Open-Platypus + ~1%Dolphin(GPT-4) + ~1%OpenOrca(GPT-4)**
<br>

**Model Finetuned By fangloveskari.**

<br>

### Training FrameWork and Parameters

#### FrameWork
https://github.com/hiyouga/LLaMA-Efficient-Tuning
We add flash_attention_2 and ORCA dataset support, with some minor modifications.

<br>

#### Parameters
We list some training parameters here:
| Parameter             | Value       |
|-----------------------|-------------|
| Finetune_Type         | QLoRA(NF4)  |
| LoRA_Rank             | 16          |
| LoRA_Alpha            | 16          |
| Batch_Size            | 14          |
| GPUs                  | 8xA100(80G) |
| LR_Scheduler          | cosine      |
| LR                    | 3e-4        |
| Epoch                 | 1           |
| DeepSpeed             | ZERO-2      |

<br>

### Model Export
We tried two methods to fuse the adapter back to the base model:
* https://github.com/hiyouga/LLaMA-Efficient-Tuning/blob/main/src/export_model.py
* https://github.com/jondurbin/qlora/blob/main/qmerge.py

Generally, the second will get better ARC(+0.15) and Truthful_QA(+0.3) scores but the other two(MMLU(-0.2) and HelloSwag(-0.2)) seems to degenerate (Just for my model).

<br>

### Evaluation

| Metric                | Value |
|-----------------------|-------|
| ARC (25-shot)         | 72.27 |
| HellaSwag (10-shot)   | 87.74 |
| MMLU (5-shot)         | 70.23 |
| TruthfulQA (0-shot)   | 63.37 |
| Avg.                  | 73.40 |

<br>

### license disclaimer:

This model is bound by the license & usage restrictions of the original Llama-2 model. And comes with no warranty or gurantees of any kind.

<br>




### Limitations & Biases:

Llama 2 and fine-tuned variants are a new technology that carries risks with use. Testing conducted to date has been in English, and has not covered, nor could it cover all scenarios. For these reasons, as with all LLMs, Llama 2 and any fine-tuned varient's potential outputs cannot be predicted in advance, and the model may in some instances produce inaccurate, biased or other objectionable responses to user prompts. Therefore, before deploying any applications of Llama 2 variants, developers should perform safety testing and tuning tailored to their specific applications of the model.

Please see the Responsible Use Guide available at https://ai.meta.com/llama/responsible-use-guide/

<br>

### Citiation:

Please kindly cite using the following BibTeX:

```bibtex
@article{platypus2023,
    title={Platypus: Quick, Cheap, and Powerful Refinement of LLMs}, 
    author={Ariel N. Lee and Cole J. Hunter and Nataniel Ruiz},
    booktitle={arXiv preprint arxiv:2308.07317},
    year={2023}
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