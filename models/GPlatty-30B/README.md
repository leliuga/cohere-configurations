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

# Information

GPlatty-30B is a merge of [garage-bAInd/Platypus-30B](https://huggingface.co/lilloukas/Platypus-30B) and [chansung/gpt4-alpaca-lora-30b](https://huggingface.co/chansung/gpt4-alpaca-lora-30b)

| Metric                | Value |
|-----------------------|-------|
| MMLU (5-shot)         | 63.6  |
| ARC (25-shot)         | 66.0  |
| HellaSwag (10-shot)   | 84.8  |
| TruthfulQA (0-shot)   | 53.8  |
| Avg.                  | 67.0  | 

We use state-of-the-art [Language Model Evaluation Harness](https://github.com/EleutherAI/lm-evaluation-harness) to run the benchmark tests above.

## Model Details

* **Trained by**: Platypus-30B trained by Cole Hunter & Ariel Lee; gpt4-alpaca-lora-30b by chansung.
* **Model type:**  **GPlatty-30B** is an auto-regressive language model based on the LLaMA transformer architecture.
* **Language(s)**: English
* **License for base weights**: License for the base LLaMA model's weights is Meta's [non-commercial bespoke license](https://github.com/facebookresearch/llama/blob/main/MODEL_CARD.md).

| Hyperparameter            | Value |
|---------------------------|-------|
| \\(n_\text{parameters}\\) | 33B   |
| \\(d_\text{model}\\)      | 6656  |
| \\(n_\text{layers}\\)     | 60    |
| \\(n_\text{heads}\\)      | 52    |


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
python main.py --model hf-causal-experimental --model_args pretrained=garage-bAInd/GPlatty-30B --tasks arc_challenge --batch_size 1 --no_cache --write_out --output_path results/Platypus-30B/arc_challenge_25shot.json --device cuda --num_fewshot 25
```

HellaSwag:
```
python main.py --model hf-causal-experimental --model_args pretrained=garage-bAInd/GPlatty-30B --tasks hellaswag --batch_size 1 --no_cache --write_out --output_path results/Platypus-30B/hellaswag_10shot.json --device cuda --num_fewshot 10
```

MMLU:
```
python main.py --model hf-causal-experimental --model_args pretrained=garage-bAInd/GPlatty-30B --tasks hendrycksTest-* --batch_size 1 --no_cache --write_out --output_path results/Platypus-30B/mmlu_5shot.json --device cuda --num_fewshot 5
```

TruthfulQA:
```
python main.py --model hf-causal-experimental --model_args pretrained=garage-bAInd/GPlatty-30B --tasks truthfulqa_mc --batch_size 1 --no_cache --write_out --output_path results/Platypus-30B/truthfulqa_0shot.json --device cuda
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