---
license: apache-2.0
datasets:
- hkust-nlp/deita-6k-v0
language:
- en
---

<img src="https://huggingface.co/datasets/hkust-nlp/deita-images/resolve/main/logo-final.png" alt="Deita banner" width="800" style="margin-left:'auto' margin-right:'auto' display:'block'"/>

# Model Card for Deita 7B V1.0 SFT

[GitHub](https://github.com/hkust-nlp/deita) | [Paper](https://arxiv.org/abs/2312.15685)

Deita is an open-sourced project designed to facilitate **Automatic Data Selection** for instruction tuning in Large Language Models (LLMs). 
Deita 7B V1.0 SFT is a fine-tuned version of Mistral-7B-v0.1 that was trained on 6k automatically selected lightweight, high-quality alignment SFT data: [Deita 6K V0](https://huggingface.co/datasets/hkust-nlp/deita-6k-v0).

## Model description

- **Model type:** Model fine tuned on automatically selected lightweight, high-quality alignment SFT data.
- **Language(s) (NLP):** Primarily English
- **Finetuned from model:** Mistral-7B-v0.1

### Model Sources

- **Repository:** https://github.com/hkust-nlp/deita
- **Model Family:** Other models and the dataset are found in the [Deita collection](https://huggingface.co/collections/hkust-nlp/deita-6569c198c174808d94cf5bd4).

## Performance



| Model                                          | Align     | Data Size  | MT-Bench | AlpacaEval(%) | OpenLLM (Avg.) |
|------------------------------------------------|-----------|------------|----------|---------------|----------------|
| **Proprietary Models**                         |           |            |          |               |                |
| GPT-4-Turbo                                    | ?         | --         | 9.32     | 97.70         | --             |
| GPT-4                                          | SFT + PPO | --         | 8.99     | 95.03         | --             |
| Claude-2                                       | SFT + PPO | --         | 8.06     | 91.36         | --             |
| GPT-3.5-turbo                                  | SFT + PPO | --         | 7.94     | 89.37         | --             |
| **Open-sourced Models based on LLaMA-1-13B**   |           |            |          |               |                |
| LIMA                                           | SFT       | 1K SFT        | 4.29     | 41.98         | 59.82          |
| WizardLM-13B                                   | SFT       | 70K SFT       | 6.35     | 75.31         | 58.96          |
| Vicuna-13B-v1.3                                | SFT       | 125K SFT      | 6.39     | 82.11         | 60.01          |
| Random                                         | SFT       | 10K SFT       | 6.03     | 71.52         | 60.14          |
| DEITA-LLaMA1-13B-v1.0-sft                           | SFT       | 10K SFT       | 6.60     | 78.01         | 64.27          |
| **Open-sourced Models based on LLaMA-2-13B**   |           |            |          |               |                |
| Tulu-2-13B                                     | SFT       | 326K SFT      | 6.70     | 78.90         | --             |
| Tulu-2-13B+DPO                                 | SFT + DPO | 326K SFT + 60K DPO | 7.00     | 89.50         | --             |
| LLaMA2-13B-Chat                                | SFT + PPO | --         | 6.65     | 81.09         | --             |
| WizardLM-13B-v1.2                              | SFT          | >70K SFT      | 7.09     | 89.17         | --             |
| Vicuna-13B-v1.5                                | SFT       | 125K SFT      | 6.57     | 78.80         | 61.63          |
| Random                                         | SFT       | 10K SFT       | 5.78     | 65.19         | 61.32          |
| DEITA-LLaMA2-13B-v1.0-sft                           | SFT       | 10K SFT       | 6.79     | 81.09         | 62.71          |
| **Open-sourced Models based on Mistral-7B**    |           |            |          |               |                |
| Mistral-7B-Instruct-v0.1                       | --        | --         | 6.84     | 69.65         | 60.45          |
| Zephyr-7B-sft                                  | SFT       | 200K SFT      | 5.32     | 75.12         | 60.93          |
| $\text{Zephyr-7B-}\beta$                       | SFT + DPO | 200K SFT + 60K DPO | 7.34     | 90.60         | 66.36          |
| OpenChat-3.5                                   | C-RLFT | >> 70K C-RLFT | 7.81     | 88.51         | --           |
| Starling-7B                                    | C-RLFT + APA | >>70K C-RLFT + 183K APA | 8.09     | 91.99         | --            |
| Random                                         | SFT       | 10K SFT       | 5.89     | 56.90         | 61.72          |
| DEITA-7B-v1.0-sft (6K)                           | SFT       | 6K SFT       | 7.22     | 80.78         | 64.94          |
| DEITA-7B-v1.0-sft (10K)                  | SFT       | 10K SFT       | 7.32     | 81.67         | 64.00          |
| DEITA-7B-v1.0             | SFT + DPO | 6K SFT + 10K DPO   | 7.55     | 90.06         | 69.86          |





## Input Format

The model is trained using the [vicuna_v1.1 template](https://github.com/lm-sys/FastChat/blob/main/fastchat/conversation.py)

```
A chat between a curious user and an artificial intelligence assistant. The assistant gives helpful, detailed, and polite answers to the user's questions. USER: Hello! ASSISTANT: Hi!</s>USER: How are you? ASSISTANT:
```

### Training hyperparameters

The following hyperparameters were used during training:
- learning_rate: 2e-05
- train_batch_size: 1
- eval_batch_size: 1
- seed: 42
- distributed_type: multi-GPU
- num_devices: 4
- gradient_accumulation_steps: 128
- total_train_batch_size: 512
- total_eval_batch_size: 4
- optimizer: Adam with betas=(0.9,0.999) and epsilon=1e-08
- lr_scheduler_type: cosine
- lr_scheduler_warmup_ratio: 0.1
- num_epochs: 6.0

### Framework versions

- Transformers 4.34.1
- Pytorch 2.1.0+cu121
- Datasets 2.14.6
- Tokenizers 0.14.1


## Citation
If you find the content of this project helpful, please cite our paper as follows:

```
@misc{liu2023what,
      title={What Makes Good Data for Alignment? A Comprehensive Study of Automatic Data Selection in Instruction Tuning}, 
      author={Wei Liu and Weihao Zeng and Keqing He and Yong Jiang and Junxian He},
      year={2023},
      eprint={2312.15685},
      archivePrefix={arXiv},
      primaryClass={cs.CL}
}
```
