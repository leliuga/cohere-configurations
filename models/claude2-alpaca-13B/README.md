---
license: llama2
datasets:
- umd-zhou-lab/claude2_alpaca
language:
- en
---
# Model Card for umd-zhou-lab/claude2-alpaca-13B

<!-- Provide a quick summary of what the model is/does. -->

This model is trained by fine-tuning llama-2 with claude2 alpaca data.

## Model Details

### Model Description

<!-- Provide a longer summary of what this model is. -->


- **Developed by:** UMD Tianyi Zhou Lab
- **Model type:** An auto-regressive language model based on the transformer architecture
- **License:** Llama 2 Community License Agreement
- **Finetuned from model:** [meta-llama/Llama-2-13b](https://huggingface.co/meta-llama/Llama-2-13b)

### Model Sources

<!-- Provide the basic links for the model. -->

- **GitHub:** [Claude2-Alpaca](https://github.com/Lichang-Chen/claude2-alpaca)
- **Data:** [claude2_alpaca](https://huggingface.co/datasets/umd-zhou-lab/claude2_alpaca)

## Uses

The primary use of this model is research on large language models and chatbots. 
The primary intended users of the model are researchers and hobbyists in natural language processing, machine learning, and artificial intelligence.

## Training

We use the prompt from [Stanford Alpaca](https://github.com/tatsu-lab/stanford_alpaca)


| Hyperparameter | Global Batch Size | Learning rate | Epochs | Max length | Weight decay |
| --- | ---: | ---: | ---: | ---: | ---: |
| Model (13B) | 128 | 1e-5 | 5 | 2048 | 0 |

## Performance

Compared to the llama2-chat, our models can have better average performance.<br>

|                    | Average | ARC  | HellaSwag | MMLU  | TruthfulQA | Alpaca_Eval | Avg Length |
|---|---|---|---|---|---|---|---|
| Llama-2-7b-chat | 56.335  | 52.9 | 78.55     | 48.32 | 45.57      | 71.37       | 1479       |
| Llama-2-13b-chat   | 59.935  | 59.04| 81.94     | 54.64 | 44.12      | 81.09       | 1513       |
|||||||||
| claude_alpaca-7b  | 57.78   | 56.66 | 81.17     | 46.58 | 46.71      | 71.23       | 1066       |
| claude_alpaca-13b | 61.29   | 61.18 | 84.08     | 55.74 | 44.18      | 78.93       | 1127       |


## Citation

Please consider citing our paper if you think our codes, data, or models are useful. Thank you!
```
@misc{claude2-alpaca,
  author = {Lichang Chen and Khalid Saifullah and Ming Li and Tianyi Zhou and Heng Huang},
  title = {Claude2-Alpaca: Instruction tuning datasets distilled from claude},
  year = {2023},
  publisher = {GitHub},
  journal = {GitHub repository},
  howpublished = {\url{https://github.com/Lichang-Chen/claude2-alpaca}},
}
```