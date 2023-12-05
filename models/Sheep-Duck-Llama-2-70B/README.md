---
thumbnail: >-
  https://cdn-uploads.huggingface.co/production/uploads/62fb1ef7e8c9c532aa7d19e4/NswB5XPkkOljeRh1xbMmR.png
pipeline_tag: text-generation
license: llama2
language:
- en
library_name: transformers
tags:
- Riiid
- llama-2
---

# sheep-duck-llama-2

<img src = "https://cdn-uploads.huggingface.co/production/uploads/62fb1ef7e8c9c532aa7d19e4/NswB5XPkkOljeRh1xbMmR.png" width="30%" height="30%">

This is a finetuned model from llama-2-70b.

## Model Details

* **Developed by**: [Riiid](https://riiid.com/)
* **Backbone Model**: [LLaMA-2](https://github.com/facebookresearch/llama/tree/main)
* **Library**: [HuggingFace Transformers](https://github.com/huggingface/transformers)


## Dataset Details

### Used Datasets
- Orca-style dataset
- Alpaca-style dataset

### Prompt Template
```
### System:
{System}

### User:
{User}

### Assistant:
{Assistant}
```

## Evaluation

| Metric                | Value |
|-----------------------|-------|
| ARC (25-shot)         | 72.44 |
| HellaSwag (10-shot)   | 87.79 |
| MMLU (5-shot)         | 70.74 |
| TruthfulQA (0-shot)   | 63.71 |
| Avg.                  | 73.67 |


## Limitations & Biases:

Llama2 and fine-tuned variants are a new technology that carries risks with use. Testing conducted to date has been in English, and has not covered, nor could it cover all scenarios. For these reasons, as with all LLMs, Llama 2 and any fine-tuned varient's potential outputs cannot be predicted in advance, and the model may in some instances produce inaccurate, biased or other objectionable responses to user prompts. Therefore, before deploying any applications of Llama 2 variants, developers should perform safety testing and tuning tailored to their specific applications of the model.

Please see the Responsible Use Guide available at https://ai.meta.com/llama/responsible-use-guide/

## License Disclaimer:
This model is bound by the license & usage restrictions of the original Llama-2 model. And comes with no warranty or gurantees of any kind.

## Contact Us

- [Riiid](https://riiid.com/)

## Citation:

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
@misc{Orca-best,
  title = {Orca-best: A filtered version of orca gpt4 dataset.},
  author = {Shahul Es},
  year = {2023},
  publisher = {HuggingFace},
  journal = {HuggingFace repository},
  howpublished = {\url{https://huggingface.co/datasets/shahules786/orca-best/},
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