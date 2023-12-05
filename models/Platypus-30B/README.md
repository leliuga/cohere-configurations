---
language:
- en
tags:
- llama
license: other
metrics:
- MMLU
- ARC 
- HellaSwag
- TruthfulQA
---

# 🥳 Platypus-30B has arrived! 

Platypus-30B is an instruction fine-tuned model based on the LLaMA-30B transformer architecture.

| Metric                | Value |
|-----------------------|-------|
| MMLU (5-shot)         | 64.2  |
| ARC (25-shot)         | 64.6  |
| HellaSwag (10-shot)   | 84.3  |
| TruthfulQA (0-shot)   | 45.8  |
| Avg.                  | 64.7  |

We use state-of-the-art [Language Model Evaluation Harness](https://github.com/EleutherAI/lm-evaluation-harness) to run the benchmark tests above.

## Model Details

* **Trained by**: Cole Hunter & Ariel Lee
* **Model type:**  **Platypus-30B** is an auto-regressive language model based on the LLaMA transformer architecture.
* **Language(s)**: English
* **License for base weights**: License for the base LLaMA model's weights is Meta's [non-commercial bespoke license](https://github.com/facebookresearch/llama/blob/main/MODEL_CARD.md).

| Hyperparameter            | Value |
|---------------------------|-------|
| \\(n_\text{parameters}\\) | 33B   |
| \\(d_\text{model}\\)      | 6656  |
| \\(n_\text{layers}\\)     | 60    |
| \\(n_\text{heads}\\)      | 52    |

## Training Dataset

Dataset of highly filtered and curated question and answer pairs. Release TBD.

## Training Procedure

`garage-bAInd/Platypus-30B` was instruction fine-tuned using LoRA on 4 A100 80GB. For training details and inference instructions please see the [Platypus-30B](https://github.com/arielnlee/Platypus-30B.git) GitHub repo.

## Reproducing Evaluation Results
Install LM Evaluation Harness:
```
git clone https://github.com/EleutherAI/lm-evaluation-harness
cd lm-evaluation-harness
pip install -e .
```
Each task was evaluated on a single A100 80GB GPU.

ARC:
```
python main.py --model hf-causal-experimental --model_args pretrained=garage-bAIdnd/Platypus-30B --tasks arc_challenge --batch_size 1 --no_cache --write_out --output_path results/Platypus-30B/arc_challenge_25shot.json --device cuda --num_fewshot 25
```

HellaSwag:
```
python main.py --model hf-causal-experimental --model_args pretrained=garage-bAIdnd/Platypus-30B --tasks hellaswag --batch_size 1 --no_cache --write_out --output_path results/Platypus-30B/hellaswag_10shot.json --device cuda --num_fewshot 10
```

MMLU:
```
python main.py --model hf-causal-experimental --model_args pretrained=garage-bAIdnd/Platypus-30B --tasks hendrycksTest-* --batch_size 1 --no_cache --write_out --output_path results/Platypus-30B/mmlu_5shot.json --device cuda --num_fewshot 5
```

TruthfulQA:
```
python main.py --model hf-causal-experimental --model_args pretrained=garage-bAIdnd/Platypus-30B --tasks truthfulqa_mc --batch_size 1 --no_cache --write_out --output_path results/Platypus-30B/truthfulqa_0shot.json --device cuda
```
## Limitations and bias

The base LLaMA model is trained on various data, some of which may contain offensive, harmful, and biased content that can lead to toxic behavior. See Section 5.1 of the LLaMA paper. We have not performed any studies to determine how fine-tuning on the aforementioned datasets affect the model's behavior and toxicity. Do not treat chat responses from this model as a substitute for human judgment or as a source of truth. Please use responsibly.

## Citations

```bibtex
@article{touvron2023llama,
  title={LLaMA: Open and Efficient Foundation Language Models},
  author={Touvron, Hugo and Lavril, Thibaut and Izacard, Gautier and Martinet, Xavier and Lachaux, Marie-Anne and Lacroix, Timoth{\'e}e and Rozi{\`e}re, Baptiste and Goyal, Naman and Hambro, Eric and Azhar, Faisal and Rodriguez, Aurelien and Joulin, Armand and Grave, Edouard and Lample, Guillaume},
  journal={arXiv preprint arXiv:2302.13971},
  year={2023}
}

@article{hu2021lora,
  title={LoRA: Low-Rank Adaptation of Large Language Models},
  author={Hu, Edward J. and Shen, Yelong and Wallis, Phillip and Allen-Zhu, Zeyuan and Li, Yuanzhi and Wang, Shean and Chen, Weizhu},
  journal={CoRR},
  year={2021}
}
```