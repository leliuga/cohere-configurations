---
language:
- en
datasets:
- garage-bAInd/Open-Platypus
license: cc-by-nc-sa-4.0
---

# Stable-Platypus2-13B

Stable-Platypus-13B is a merge of [`garage-bAInd/Platypus2-13B`](https://huggingface.co/garage-bAInd/Platypus2-13B) and [`stabilityai/StableBeluga-13B`](https://huggingface.co/stabilityai/StableBeluga-13B).

![Platty](./Best_Platty_small.jpeg)

### Benchmark Metrics

| Metric                | Value |
|-----------------------|-------|
| MMLU (5-shot)         |   58.30   |
| ARC (25-shot)         |   62.71   |
| HellaSwag (10-shot)   |   82.29   |
| TruthfulQA (0-shot)   |   52.52   |
| Avg.                  |   63.96   |

We use state-of-the-art [Language Model Evaluation Harness](https://github.com/EleutherAI/lm-evaluation-harness) to run the benchmark tests above, using the same version as the HuggingFace LLM Leaderboard. Please see below for detailed instructions on reproducing benchmark results.

### Model Details

* **Trained by**: **Platypus2-13B** trained by Cole Hunter & Ariel Lee; **StableBeluga-13B** trained by StabilityAI
* **Model type:**  **Stable-Platypus2-13B** is an auto-regressive language model based on the LLaMA 2 transformer architecture.
* **Language(s)**: English
* **License for Platypus2-13B base weights**: Non-Commercial Creative Commons license ([CC BY-NC-4.0](https://creativecommons.org/licenses/by-nc/4.0/))
* **License for StableBeluga-13B base weights**: See Notice.txt

### Prompt Template
```
### Instruction:

<prompt> (without the <>)

### Response:
```

### Training Dataset

`garage-bAInd/Platypus2-70B` trained using STEM and logic based dataset [`garage-bAInd/Open-Platypus`](https://huggingface.co/datasets/garage-bAInd/Open-Platypus).

Please see our [paper](https://arxiv.org/abs/2308.07317) and [project webpage](https://platypus-llm.github.io) for additional information.

### Training Procedure

`garage-bAInd/Platypus2-13B` was instruction fine-tuned using LoRA on 1 A100 80GB. For training details and inference instructions please see the [Platypus](https://github.com/arielnlee/Platypus) GitHub repo.

### Reproducing Evaluation Results

Install LM Evaluation Harness:
```
# clone repository
git clone https://github.com/EleutherAI/lm-evaluation-harness.git
# change to repo directory
cd lm-evaluation-harness
# check out the correct commit
git checkout b281b0921b636bc36ad05c0b0b0763bd6dd43463
# install
pip install -e .
```
Each task was evaluated on a single A100 80GB GPU.

ARC:
```
python main.py --model hf-causal-experimental --model_args pretrained=garage-bAInd/Stable-Platypus2-13B --tasks arc_challenge --batch_size 1 --no_cache --write_out --output_path results/Stable-Platypus2-13B/arc_challenge_25shot.json --device cuda --num_fewshot 25
```

HellaSwag:
```
python main.py --model hf-causal-experimental --model_args pretrained=garage-bAInd/Stable-Platypus2-13B --tasks hellaswag --batch_size 1 --no_cache --write_out --output_path results/Stable-Platypus2-13B/hellaswag_10shot.json --device cuda --num_fewshot 10
```

MMLU:
```
python main.py --model hf-causal-experimental --model_args pretrained=garage-bAInd/Stable-Platypus2-13B --tasks hendrycksTest-* --batch_size 1 --no_cache --write_out --output_path results/Stable-Platypus2-13B/mmlu_5shot.json --device cuda --num_fewshot 5
```

TruthfulQA:
```
python main.py --model hf-causal-experimental --model_args pretrained=garage-bAInd/Stable-Platypus2-13B --tasks truthfulqa_mc --batch_size 1 --no_cache --write_out --output_path results/Stable-Platypus2-13B/truthfulqa_0shot.json --device cuda
```
### Limitations and bias

Llama 2 and fine-tuned variants are a new technology that carries risks with use. Testing conducted to date has been in English, and has not covered, nor could it cover all scenarios. For these reasons, as with all LLMs, Llama 2 and any fine-tuned varient's potential outputs cannot be predicted in advance, and the model may in some instances produce inaccurate, biased or other objectionable responses to user prompts. Therefore, before deploying any applications of Llama 2 variants, developers should perform safety testing and tuning tailored to their specific applications of the model.

Please see the Responsible Use Guide available at https://ai.meta.com/llama/responsible-use-guide/

### Citations
```bibtex
@article{platypus2023,
    title={Platypus: Quick, Cheap, and Powerful Refinement of LLMs}, 
    author={Ariel N. Lee and Cole J. Hunter and Nataniel Ruiz},
    booktitle={arXiv preprint arxiv:2308.07317},
    year={2023}
}
```
```bibtex
@misc{touvron2023llama,
    title={Llama 2: Open Foundation and Fine-Tuned Chat Models}, 
    author={Hugo Touvron and Louis Martin and Kevin Stone and Peter Albert and Amjad Almahairi and Yasmine Babaei and Nikolay Bashlykov       year={2023},
    eprint={2307.09288},
    archivePrefix={arXiv},
}
```
```bibtex
@inproceedings{
    hu2022lora,
    title={Lo{RA}: Low-Rank Adaptation of Large Language Models},
    author={Edward J Hu and Yelong Shen and Phillip Wallis and Zeyuan Allen-Zhu and Yuanzhi Li and Shean Wang and Lu Wang and Weizhu Chen},
    booktitle={International Conference on Learning Representations},
    year={2022},
    url={https://openreview.net/forum?id=nZeVKeeFYf9}
}
```