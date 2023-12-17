---
license: apache-2.0
tags:
- alignment-handbook
- generated_from_trainer
- juanako
- mistral
- UNA
datasets:
- HuggingFaceH4/ultrafeedback_binarized
model-index:
- name: juanako-7b-UNA
  results:
  - task:
      type: text-generation
      name: TruthfulQA (MC2)
    dataset:
      type: text-generation
      name: truthful_qa
      config: multiple_choice
      split: validation
    metrics:
      - type: accuracy
        value: 65.13
        verified: true
  - task:
      type: text-generation
      name: ARC-Challenge
    dataset:
      type: text-generation
      name: ai2_arc
      config: ARC-Challenge
      split: test
    metrics:
      - type: accuracy
        value: 68.17
        verified: true
  - task:
      type: text-generation
      name: HellaSwag
    dataset:
      type: text-generation
      name: Rowan/hellaswag
      split: test
    metrics:
      - type: accuracy
        value: 85.34
        verified: true
  - task:
      type: text-generation
      name: Winogrande
    dataset:
      type: text-generation
      name: winogrande
      config: winogrande_debiased
      split: test
    metrics:
      - type: accuracy
        value: 78.85
        verified: true
  - task:
      type: text-generation
      name: MMLU
    dataset:
      type: text-generation
      name: cais/mmlu
      config: all
      split: test
    metrics:
      - type: accuracy
        value: 62.47
        verified: true
  - task:
      type: text-generation
      name: PiQA
    dataset:
      type: text-generation
      name: piqa
      split: test
    metrics:
      - type: accuracy
        value: 83.57
  - task:
      type: text-generation
      name: DROP
    dataset:
      type: text-generation
      name: drop
      split: validation
    metrics:
      - type: accuracy
        value: 38.74
        verified: true
  - task:
      type: text-generation
      name: PubMedQA
    dataset:
      type: text-generation
      name: bigbio/pubmed_qa
      config: pubmed_qa_artificial_bigbio_qa
      split: validation
    metrics:
      - type: accuracy
        value: 76.0
---

# juanako-7b-UNA (Uniform Neural Alignment)

This model is a fine-tuned version of [fblgit/juanako-7b-UNA-v2-phase-1](https://huggingface.co/fblgit/juanako-7b-UNA-v2-phase-1) on the HuggingFaceH4/ultrafeedback_binarized dataset.
It outperforms in many aspects most of the current Mistral based models and is the **latest and most powerful juanako version as of now**.

## Scores

The official HuggingFace results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/results/blob/main/fblgit/juanako-7b-UNA/results_2023-11-28T08-33-33.965228.json)

| Model | Average ⬆️| ARC (25-s) ⬆️ | HellaSwag (10-s) ⬆️ | MMLU (5-s) ⬆️| TruthfulQA (MC) (0-s) ⬆️ | Winogrande (5-s) | GSM8K (5-s) | DROP (3-s) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
|[mistralai/Mistral-7B-v0.1](https://huggingface.co/mistralai/Mistral-7B-v0.1) | 50.32 | 59.58  | 83.31  | 64.16  | 42.15 | 78.37 | 18.12 | 6.14 |
| [Intel/neural-chat-7b-v3-1](https://huggingface.co/Intel/neural-chat-7b-v3-1) | 59.0 | 66.21 | 83.64 | 62.37  | 59.65 | 78.14 | 19.56 | 43.84 |
| [fblgit/juanako-7b-UNA](https://huggingface.co/fblgit/juanako-7b-UNA) | **59.91** | **68.17** | **85.34** | 62.47  | **65.13** | **78.85** | **20.7** | 38.74 |

It scores: **59.91** according HuggingFace LLM Leaderboard.
It scores: **65.1** with `big-refactor` branch of lm-eval-harness

Author [Xavier M.](mailto:xavi@juanako.ai) @fblgit

## Model description

juanako uses UNA, Uniform Neural Alignment. A training technique that ease alignment between transformer layers yet to be published.

### Prompts
The following prompts showed positive results, it may depend the task and needs further experimentation but this should work for starters:
```
<|im_start|>system
- You are a helpful assistant chatbot trained by MosaicML.
- You answer questions.
- You are excited to be able to help the user, but will refuse to do anything that could be considered harmful to the user.
- You are more than just an information source, you are also able to write poetry, short stories, and make jokes.<|im_end|>
<|im_start|>user
Explain QKV<|im_end|>
<|im_start|>assistant
```
```
### Assistant: I am StableVicuna, a large language model created by CarperAI. I am here to chat!

### Human: Explain QKV
### Assistant:
```
```
[Round <|round|>]
问：Explain QKV
答：
```
```
[Round <|round|>]
Question：Explain QKV
Answer：
```
```
Question：Explain QKV
Answer：
```

## Evaluations (lm-eval big-refactor branch)

### TruthfulQA 0-Shot
```
|    Tasks     |Version|Filter|Metric|Value |   |Stderr|
|--------------|-------|------|------|-----:|---|-----:|
|truthfulqa_mc2|Yaml   |none  |acc   |0.6549|±  |0.0153|
```
### ARC 25-Shot
```
|    Tasks    |Version|Filter| Metric |Value |   |Stderr|
|-------------|-------|------|--------|-----:|---|-----:|
|arc_challenge|Yaml   |none  |acc     |0.6476|±  |0.0140|
|             |       |none  |acc_norm|0.6809|±  |0.0136|
```
### HellaSwag 10-Shot
```
|  Tasks  |Version|Filter| Metric |Value |   |Stderr|
|---------|-------|------|--------|-----:|---|-----:|
|hellaswag|Yaml   |none  |acc     |0.6703|±  |0.0047|
|         |       |none  |acc_norm|0.8520|±  |0.0035|
```
### GSM8k 5-Shot
```
|Tasks|Version|  Filter  |  Metric   |Value |   |Stderr|
|-----|-------|----------|-----------|-----:|---|-----:|
|gsm8k|Yaml   |get-answer|exact_match|0.4898|±  |0.0138|
```
### GPT Evaluations 0-Shot
```
|    Tasks     |Version|Filter|  Metric  |Value |   |Stderr|
|--------------|-------|------|----------|-----:|---|-----:|
|boolq         |Yaml   |none  |acc       |0.8703|±  |0.0059|
|lambada_openai|Yaml   |none  |perplexity|3.2598|±  |0.0705|
|              |       |none  |acc       |0.7336|±  |0.0062|
|piqa          |Yaml   |none  |acc       |0.8254|±  |0.0089|
|              |       |none  |acc_norm  |0.8292|±  |0.0088|
|sciq          |Yaml   |none  |acc       |0.9580|±  |0.0063|
|              |       |none  |acc_norm  |0.9130|±  |0.0089|
```
### MathQA 0-Shot
```
|Tasks |Version|Filter| Metric |Value |   |Stderr|
|------|-------|------|--------|-----:|---|-----:|
|mathqa|Yaml   |none  |acc     |0.3752|±  |0.0089|
|      |       |none  |acc_norm|0.3772|±  |0.0089|
```
### PiQa 1-Shot
```
|Tasks|Version|Filter| Metric |Value |   |Stderr|
|-----|-------|------|--------|-----:|---|-----:|
|piqa |Yaml   |none  |acc     |0.8308|±  |0.0087|
|     |       |none  |acc_norm|0.8357|±  |0.0086|
```
### Winogrande 5-Shot
```
|  Tasks   |Version|Filter|Metric|Value|   |Stderr|
|----------|-------|------|------|----:|---|-----:|
|winogrande|Yaml   |none  |acc   |0.768|±  |0.0119|
```
### PubMedQA 0-Shot
```
| Tasks  |Version|Filter|Metric|Value|   |Stderr|
|--------|-------|------|------|----:|---|-----:|
|pubmedqa|Yaml   |none  |acc   | 0.76|±  |0.0191|
```
### RACE 1-Shot
```
|Tasks|Version|Filter|Metric|Value |   |Stderr|
|-----|-------|------|------|-----:|---|-----:|
|race |Yaml   |none  |acc   |0.5282|±  |0.0154|
```
### MMLU 5-Shot (8-Bit)
```
|      Groups      |Version|Filter|Metric|Value |   |Stderr|
|------------------|-------|------|------|-----:|---|-----:|
|mmlu              |N/A    |none  |acc   |0.6137|±  |0.1243|
| - humanities     |N/A    |none  |acc   |0.5671|±  |0.1101|
| - other          |N/A    |none  |acc   |0.6859|±  |0.1164|
| - social_sciences|N/A    |none  |acc   |0.7195|±  |0.0713|
| - stem           |N/A    |none  |acc   |0.5087|±  |0.1297|
```
### DROP 3-Shot (8-Bit) (Instruct-Eval)
```
{'score': 0.49801113762927607}
{'drop': 49.8}
drop: 49.8
```

### CRASS 0-Shot (Instruct-Eval)
```
{'score': 0.8357664233576643}
{'crass': 83.58}
crass: 83.58
```

## Training Details

### Training hyperparameters

The following hyperparameters were used during training:
- learning_rate: 0.0001
- train_batch_size: 1
- eval_batch_size: 1
- seed: 42
- distributed_type: multi-GPU
- num_devices: 14
- gradient_accumulation_steps: 16
- total_train_batch_size: 224
- total_eval_batch_size: 14
- optimizer: Adam with betas=(0.9,0.999) and epsilon=1e-08
- lr_scheduler_type: linear
- lr_scheduler_warmup_ratio: 0.01
- num_epochs: 1

### Training results

| Training Loss | Epoch | Step | Validation Loss | Rewards/chosen | Rewards/rejected | Rewards/accuracies | Rewards/margins | Logps/rejected | Logps/chosen | Logits/rejected | Logits/chosen |
|:-------------:|:-----:|:----:|:---------------:|:--------------:|:----------------:|:------------------:|:---------------:|:--------------:|:------------:|:---------------:|:-------------:|
| 0.4795        | 0.2   | 56   | 0.4958          | -1.3684        | -2.6385          | 0.7552             | 1.2701          | -265.3887      | -241.2612    | -2.2572         | -2.4922       |
| 0.4642        | 0.4   | 112  | 0.4859          | -1.0380        | -1.9769          | 0.7273             | 0.9389          | -258.7718      | -237.9569    | -2.2414         | -2.4751       |
| 0.4758        | 0.61  | 168  | 0.4808          | -1.2594        | -2.3704          | 0.7343             | 1.1110          | -262.7074      | -240.1708    | -2.2305         | -2.4633       |
| 0.4549        | 0.81  | 224  | 0.4768          | -1.1906        | -2.3201          | 0.7552             | 1.1295          | -262.2044      | -239.4827    | -2.2284         | -2.4610       |


### Framework versions

- Transformers 4.35.0-UNA
- Pytorch 2.1.0
- Datasets 2.14.6
- Tokenizers 0.14.1

## Citations
If you find juanako useful please:

```
@misc{juanako7buna,
  title={Juanako: Uniform Neural Alignment}, 
  author={Xavier Murias},
  year={2023},
  publisher = {HuggingFace},
  journal = {HuggingFace repository},
  howpublished = {\url{https://huggingface.co/fblgit/juanako-7b-UNA}},
}
```

Thanks to all the brilliant humans behind the creation of AI, here some of the ones that we find relevant to our research. If you feel a citation is missing, please contact.
```
@misc{lin2021truthfulqa,
  title={TruthfulQA: Measuring How Models Mimic Human Falsehoods},
  author={Stephanie Lin and Jacob Hilton and Owain Evans},
  year={2021},
  eprint={2109.07958},
  archivePrefix={arXiv},
  primaryClass={cs.CL}
}
@misc{tunstall2023zephyr,
      title={Zephyr: Direct Distillation of LM Alignment}, 
      author={Lewis Tunstall and Edward Beeching and Nathan Lambert and Nazneen Rajani and Kashif Rasul and Younes Belkada and Shengyi Huang and Leandro von Werra and Clémentine Fourrier and Nathan Habib and Nathan Sarrazin and Omar Sanseviero and Alexander M. Rush and Thomas Wolf},
      year={2023},
      eprint={2310.16944},
      archivePrefix={arXiv},
      primaryClass={cs.LG}
}
@inproceedings{Bisk2020,
  author = {Yonatan Bisk and Rowan Zellers and
            Ronan Le Bras and Jianfeng Gao
            and Yejin Choi},
  title = {PIQA: Reasoning about Physical Commonsense in
           Natural Language},
  booktitle = {Thirty-Fourth AAAI Conference on
               Artificial Intelligence},
  year = {2020},
}
@software{eval-harness,
  author       = {Gao, Leo and
                  Tow, Jonathan and
                  Biderman, Stella and
                  Black, Sid and
                  DiPofi, Anthony and
                  Foster, Charles and
                  Golding, Laurence and
                  Hsu, Jeffrey and
                  McDonell, Kyle and
                  Muennighoff, Niklas and
                  Phang, Jason and
                  Reynolds, Laria and
                  Tang, Eric and
                  Thite, Anish and
                  Wang, Ben and
                  Wang, Kevin and
                  Zou, Andy},
  title        = {A framework for few-shot language model evaluation},
  month        = sep,
  year         = 2021,
  publisher    = {Zenodo},
  version      = {v0.0.1},
  doi          = {10.5281/zenodo.5371628},
  url          = {https://doi.org/10.5281/zenodo.5371628}
}
@misc{rafailov2023direct,
    title={Direct Preference Optimization: Your Language Model is Secretly a Reward Model}, 
    author={Rafael Rafailov and Archit Sharma and Eric Mitchell and Stefano Ermon and Christopher D. Manning and Chelsea Finn},
    year={2023},
    eprint={2305.18290},
    archivePrefix={arXiv},
}
```
