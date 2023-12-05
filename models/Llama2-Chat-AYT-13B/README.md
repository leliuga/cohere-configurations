---
license: llama2
---

This is a model diverged from Llama-2-13b-chat-hf. We hypotheize that if we find a method to ensemble the top rankers in each benchmark effectively, its performance maximizes as well. Following this intuition, we ensembled the top models in each benchmarks(ARC, MMLU and TruthFulQA) to create our model.

# Model Details
- **Developed by**: Posicube Inc. 
- **Backbone Model**: LLaMA-2-13b-chat
- **Library**: HuggingFace Transformers
- **Used Dataset Details**
  Orca-style datasets, Alpaca-style datasets


# Evaluation 
We achieved the top ranker among 13B models at Sep-13rd 2023.

| Metric              |Scores on Leaderboard| Our results |
|---------------------|---------------------|-------------|
| ARC (25-shot)       | 63.31 | 63.57 |
| HellaSwag (10-shot) | 83.53 | 83.77 |
| MMLU (5-shot)       | 59.67 | 59.69 |
| TruthfulQA (0-shot) | 55.8  | 55.48 |
| Avg.                | 65.58 | 65.63 |

# Limitations & Biases:
Llama2 and fine-tuned variants are a new technology that carries risks with use. Testing conducted to date has been in English, and has not covered, nor could it cover all scenarios. For these reasons, as with all LLMs, Llama 2 and any fine-tuned varient's potential outputs cannot be predicted in advance, and the model may in some instances produce inaccurate, biased or other objectionable responses to user prompts. Therefore, before deploying any applications of Llama 2 variants, developers should perform safety testing and tuning tailored to their specific applications of the model.

Please see the Responsible Use Guide available at https://ai.meta.com/llama/responsible-use-guide/

# License Disclaimer:
This model is bound by the license & usage restrictions of the original Llama-2 model. And comes with no warranty or gurantees of any kind.

# Contact Us
[Posicube](https://www.posicube.com/)

# Citiation:
Please kindly cite using the following BibTeX:

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