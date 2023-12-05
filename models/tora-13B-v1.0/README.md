---
license: llama2
datasets:
- gsm8k
- competition_math
language:
- en
metrics:
- exact_match
library_name: transformers
pipeline_tag: text-generation
tags:
- code
- math
---


<h1 align="center">
ToRA: A Tool-Integrated Reasoning Agent <br> for Mathematical Problem Solving
</h1>

<p align="center">
  <a href="https://microsoft.github.io/ToRA/"><b>[🌐 Website]</b></a> •
  <a href="https://arxiv.org/abs/2309.17452"><b>[📜 Paper]</b></a> •
  <a href="https://huggingface.co/llm-agents"><b>[🤗 HF Models]</b></a> •
  <a href="https://github.com/microsoft/ToRA"><b>[🐱 GitHub]</b></a>
  <br>
  <a href="https://twitter.com/zhs05232838/status/1708860992631763092"><b>[🐦 Twitter]</b></a> •
  <a href="https://www.reddit.com/r/LocalLLaMA/comments/1703k6d/tora_a_toolintegrated_reasoning_agent_for/"><b>[💬 Reddit]</b></a> •
  <a href="https://notes.aimodels.fyi/researchers-announce-tora-training-language-models-to-better-understand-math-using-external-tools/">[🍀 Unofficial Blog]</a>
  <!-- <a href="#-quick-start">Quick Start</a> • -->
  <!-- <a href="#%EF%B8%8F-citation">Citation</a> -->
</p>

<p align="center">
Repo for "<a href="https://arxiv.org/abs/2309.17452" target="_blank">ToRA: A Tool-Integrated Reasoning Agent for Mathematical Problem Solving</a>"
</p>

## 🔥 News

- [2023/10/08] 🔥🔥🔥 All ToRA models released at [HuggingFace](https://huggingface.co/llm-agents)!!! 
- [2023/09/29] ToRA paper, repo, and website released.

## 💡 Introduction

ToRA is a series of Tool-integrated Reasoning Agents designed to solve challenging mathematical reasoning problems by interacting with tools, e.g., computation libraries and symbolic solvers. ToRA series seamlessly integrate natural language reasoning with the utilization of external tools, thereby amalgamating the analytical prowess of language and the computational efficiency of external tools.

| Model | Size | GSM8k | MATH | AVG@10 math tasks<sup>&dagger;</sup> |
|---|---|---|---|---|
| GPT-4 | - | 92.0 | 42.5 | 78.3 |
| GPT-4 (PAL) | - | 94.2 | 51.8 | 86.4 |
| [ToRA-7B](https://huggingface.co/llm-agents/tora-7b-v1.0) | 7B | 68.8 | 40.1 | 62.4|
| [ToRA-Code-7B](https://huggingface.co/llm-agents/tora-code-7b-v1.0) | 7B | 72.6 | 44.6 | 66.5| 
| [ToRA-13B](https://huggingface.co/llm-agents/tora-13b-v1.0) | 13B |  72.7 | 43.0 | 65.9|
| [ToRA-Code-13B](https://huggingface.co/llm-agents/tora-code-13b-v1.0) | 13B | 75.8 | 48.1 | 71.3 |
| [ToRA-Code-34B<sup>*</sup>](https://huggingface.co/llm-agents/tora-code-34b-v1.0) | 34B | 80.7 | **51.0** | 74.8 |
| [ToRA-70B](https://huggingface.co/llm-agents/tora-70b-v1.0) | 70B | **84.3** | 49.7 | **76.9** |

- <sup>*</sup>ToRA-Code-34B is currently the first and only open-source model to achieve over 50% accuracy (pass@1) on the MATH dataset, which significantly outperforms GPT-4’s CoT result (51.0 vs. 42.5), and is competitive with GPT-4 solving problems with programs. By open-sourcing our codes and models, we hope more breakthroughs will come!

- <sup>&dagger;</sup>10 math tasks include GSM8k, MATH, GSM-Hard, SVAMP, TabMWP, ASDiv, SingleEQ, SingleOP, AddSub, and MultiArith.


## ⚡️ Training

The models are trained on ToRA-Corpus 16k, which contains tool-integrated reasoning trajectories of MATH and GSM8k from GPT-4.

We use imitation learning (i.e., SFT) to fine-tune the models, and then apply our proposed *output space shaping* to improve tool-integrated reasoning behaviors. Please refer to the [paper](https://arxiv.org/pdf/2309.17452.pdf) for more details.


## 🪁 Inference & Evaluation

Please refer to ToRA's [GitHub repo](https://github.com/microsoft/ToRA) for inference, evaluation, and training code.


## ☕️ Citation

If you find this repository helpful, please consider citing our paper:

```
@misc{gou2023tora,
      title={ToRA: A Tool-Integrated Reasoning Agent for Mathematical Problem Solving}, 
      author={Zhibin Gou and Zhihong Shao and Yeyun Gong and yelong shen and Yujiu Yang and Minlie Huang and Nan Duan and Weizhu Chen},
      year={2023},
      eprint={2309.17452},
      archivePrefix={arXiv},
      primaryClass={cs.CL}
}
```