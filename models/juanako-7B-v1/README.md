---
base_model: fblgit/zephyr-lora-dpo-b1
tags:
- alignment-handbook
- generated_from_trainer
datasets:
- HuggingFaceH4/ultrafeedback_binarized
model-index:
- name: juanako-7b-v1
  results: []
license: artistic-2.0
---

# juanako-7b-v1 (UNA: Uniform Neural Alignment)

This model uses uniform neural alignment (UNA) for the DPO training phases and is a fine-tuned version of [fblgit/zephyr-lora-dpo-b1](https://huggingface.co/fblgit/zephyr-lora-dpo-b1) on the HuggingFaceH4/ultrafeedback_binarized dataset.

**It is recommended to use the latest [Juanako Version](https://huggingface.co/fblgit/juanako-7b-UNA) which highly outperforms the v1**

It achieves the following results on the evaluation set:
- Loss: 0.4594
- Rewards/chosen: -1.1095
- Rewards/rejected: -2.3132
- Rewards/accuracies: 0.7964
- Rewards/margins: 1.2037
- Logps/rejected: -220.0052
- Logps/chosen: -217.5506
- Logits/rejected: -2.5535
- Logits/chosen: -2.7973

Followed [alignment-handbook](https://github.com/huggingface/alignment-handbook) to perform DPO (Phase 2) over Zephyr-SFT model.

**Please feel free to run more tests and commit the results. Also if you are interested to participate in [UNA's paper research or GPU sponsorship](mailto:xavi@juanako.ai) to support UNA research, feel free to contact.**

Special thanks to [TheBloke](https://huggingface.co/TheBloke) for converting the model into multiple formats and overall his enormous contribution to the community.
Here are the models:
* [juanako-7B-v1-AWQ](https://huggingface.co/TheBloke/juanako-7B-v1-AWQ)
* [juanako-7B-v1-GPTQ](https://huggingface.co/TheBloke/juanako-7B-v1-GPTQ)
* [juanako-7B-v1-GGUF](https://huggingface.co/TheBloke/juanako-7B-v1-GGUF)


## Prompt and Inference Usage
```
# Install transformers from source - only needed for versions <= v4.34
# pip install git+https://github.com/huggingface/transformers.git
# pip install accelerate

import torch
from transformers import pipeline

pipe = pipeline("text-generation", model="fblgit/juanako-7b-v1", torch_dtype=torch.float16, device_map="auto")

# We use the tokenizer's chat template to format each message - see https://huggingface.co/docs/transformers/main/en/chat_templating
messages = [
    {
        "role": "system",
        "content": "You are a friendly chatbot who always responds in the style of a pirate",
    },
    {"role": "user", "content": "How many helicopters can a human eat in one sitting?"},
]
prompt = pipe.tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True)
outputs = pipe(prompt, max_new_tokens=256, do_sample=True, temperature=0.7, top_k=50, top_p=0.95)
print(outputs[0]["generated_text"])
# <|system|>
# You are a friendly chatbot who always responds in the style of a pirate.</s>
# <|user|>
# How many helicopters can a human eat in one sitting?</s>
# <|assistant|>
# Ah, me hearty matey! But yer question be a puzzler! A human cannot eat a helicopter in one sitting, as helicopters are not edible. They be made of metal, plastic, and other materials, not food!
```

## Model description

**It seems to outperforms the original Zephyr in most of the tasks.**

I trained Juanako with the same datasets and trainer from [alignment-handbook/zephyr-7b-sft-lora](https://huggingface.co/alignment-handbook/zephyr-7b-sft-lora) 
* 1 epoch on DPO with transformers-UNA, the result is [fblgit/zephyr-lora-dpo-b1](https://huggingface.co/fblgit/zephyr-lora-dpo-b1) after merge using FastChat converter.
* finally 1 epoch on DPO with transformers-UNA to [fblgit/zephyr-lora-dpo-b1](https://huggingface.co/fblgit/zephyr-lora-dpo-b1).

Some other experiments were performed as well to test transformers-UNA capabilities on diverse scenarios and models.

**This is a complete version of the model, the result of converting LoRa's**

## Intended uses & limitations

Research purposes.

## Training and evaluation data

alignment-handbook DPO with UNA on top of the SFT lora.

### Evaluation lm-evaluation-harness

#### GSM8K 5-Shot
```
hf (pretrained=fblgit/juanako-7b-v1,load_in_4bit=False,dtype=float16), limit: None, num_fewshot: 5, batch_size: 4
```
|Tasks|Version|  Filter  |  Metric   |Value |   |Stderr|
|-----|-------|----------|-----------|-----:|---|-----:|
|gsm8k|Yaml   |get-answer|exact_match|0.4761|±  |0.0138|

#### 0-Shot Tests
```
hf (pretrained=fblgit/juanako-7b-v1,load_in_4bit=False,dtype=float16), limit: None, num_fewshot: 0, batch_size: 8
```
|       Tasks       |Version|Filter|  Metric   | Value |   |Stderr|
|-------------------|-------|------|-----------|------:|---|-----:|
|arc_challenge      |Yaml   |none  |acc        | 0.5691|±  |0.0145|
|                   |       |none  |acc_norm   | 0.6041|±  |0.0143|
|arc_easy           |Yaml   |none  |acc        | 0.8363|±  |0.0076|
|                   |       |none  |acc_norm   | 0.8161|±  |0.0079|
|hellaswag          |Yaml   |none  |acc        | 0.6554|±  |0.0047|
|                   |       |none  |acc_norm   | 0.8411|±  |0.0036|
|boolq              |Yaml   |none  |acc        | 0.8355|±  |0.0065|
|lambada            |N/A    |none  |perplexity | 3.3607|±  |0.1398|
|                   |       |none  |acc        | 0.7309|±  |0.0137|
|piqa               |Yaml   |none  |acc        | 0.8194|±  |0.0090|
|                   |       |none  |acc_norm   | 0.8335|±  |0.0087|
|sciq               |Yaml   |none  |acc        | 0.9480|±  |0.0070|
|                   |       |none  |acc_norm   | 0.8960|±  |0.0097|
|truthfulqa         |N/A    |none  |bleu_max   |26.0803|±  |0.6528|
| - truthfulqa_mc1  |Yaml   |none  |acc        | 0.4198|±  |0.0173|
| - truthfulqa_mc2  |Yaml   |none  |acc        | 0.5847|±  |0.0153|
|winogrande         |Yaml   |none  |acc        | 0.7609|±  |0.0120|

```
hf (pretrained=fblgit/juanako-7b-v1,load_in_4bit=False,dtype=float16), limit: None, num_fewshot: 25, batch_size: 1
```
|    Tasks    |Version|Filter| Metric |Value |   |Stderr|
|-------------|-------|------|--------|-----:|---|-----:|
|arc_challenge|Yaml   |none  |acc     |0.6058|±  |0.0143|
|             |       |none  |acc_norm|0.6485|±  |0.0140|

#### HellaSwag 10-Shot
```
hf (pretrained=fblgit/juanako-7b-v1,load_in_4bit=False,dtype=float16), limit: None, num_fewshot: 10, batch_size: 16
```
|  Tasks  |Version|Filter| Metric |Value |   |Stderr|
|---------|-------|------|--------|-----:|---|-----:|
|hellaswag|Yaml   |none  |acc     |0.6582|±  |0.0047|
|         |       |none  |acc_norm|0.8513|±  |0.0036|

## Training procedure

### Training hyperparameters

The following hyperparameters were used during training:
- learning_rate: 0.0001
- train_batch_size: 1
- eval_batch_size: 1
- seed: 42
- distributed_type: multi-GPU
- num_devices: 12
- gradient_accumulation_steps: 16
- total_train_batch_size: 192
- total_eval_batch_size: 12
- optimizer: Adam with betas=(0.9,0.999) and epsilon=1e-08
- lr_scheduler_type: linear
- lr_scheduler_warmup_ratio: 0.01
- num_epochs: 1

### Training results

| Training Loss | Epoch | Step | Validation Loss | Rewards/chosen | Rewards/rejected | Rewards/accuracies | Rewards/margins | Logps/rejected | Logps/chosen | Logits/rejected | Logits/chosen |
|:-------------:|:-----:|:----:|:---------------:|:--------------:|:----------------:|:------------------:|:---------------:|:--------------:|:------------:|:---------------:|:-------------:|
| 0.4966        | 0.15  | 50   | 0.4893          | -1.1759        | -2.2914          | 0.7485             | 1.1155          | -219.7872      | -218.2148    | -2.5450         | -2.7884       |
| 0.4522        | 0.31  | 100  | 0.4808          | -0.8099        | -1.8893          | 0.7784             | 1.0794          | -215.7659      | -214.5544    | -2.5644         | -2.8095       |
| 0.5048        | 0.46  | 150  | 0.4706          | -1.0526        | -2.1412          | 0.7725             | 1.0887          | -218.2852      | -216.9814    | -2.5638         | -2.8089       |
| 0.4853        | 0.62  | 200  | 0.4640          | -1.0787        | -2.2821          | 0.7725             | 1.2034          | -219.6941      | -217.2426    | -2.5460         | -2.7891       |
| 0.4639        | 0.77  | 250  | 0.4636          | -1.2348        | -2.4583          | 0.8084             | 1.2235          | -221.4559      | -218.8034    | -2.5533         | -2.7970       |
| 0.4634        | 0.93  | 300  | 0.4601          | -1.1370        | -2.3243          | 0.7964             | 1.1873          | -220.1163      | -217.8257    | -2.5540         | -2.7977       |
| -             | 1.00  | 300  | 0.4594          | -1.1095        | -2.3132          | 0.7964             | 1.2037          | -220.0052      | -217.5506    | -2.5535         | -2.7973       |

### Framework versions

- Transformers 4.35.0-UNA
- Pytorch 2.1.0
- Datasets 2.14.6
- Tokenizers 0.14.1

## MMLU Results

#### 5-Shot
```
hf (pretrained=fblgit/juanako-7b-v1,load_in_4bit=False,dtype=float16), limit: None, num_fewshot: 5, batch_size: 1
```
|                 Tasks                 |Version|Filter|Metric|Value |   |Stderr|
|---------------------------------------|-------|------|------|-----:|---|-----:|
|mmlu                                   |N/A    |none  |acc   |0.6236|±  |0.1269|
| - humanities                          |N/A    |none  |acc   |0.5651|±  |0.1289|
|  - formal_logic                       |Yaml   |none  |acc   |0.4365|±  |0.0444|
|  - high_school_european_history       |Yaml   |none  |acc   |0.7636|±  |0.0332|
|  - high_school_us_history             |Yaml   |none  |acc   |0.8039|±  |0.0279|
|  - high_school_world_history          |Yaml   |none  |acc   |0.7848|±  |0.0268|
|  - international_law                  |Yaml   |none  |acc   |0.7686|±  |0.0385|
|  - jurisprudence                      |Yaml   |none  |acc   |0.7778|±  |0.0402|
|  - logical_fallacies                  |Yaml   |none  |acc   |0.7853|±  |0.0323|
|  - moral_disputes                     |Yaml   |none  |acc   |0.7168|±  |0.0243|
|  - moral_scenarios                    |Yaml   |none  |acc   |0.3207|±  |0.0156|
|  - philosophy                         |Yaml   |none  |acc   |0.7042|±  |0.0259|
|  - prehistory                         |Yaml   |none  |acc   |0.7593|±  |0.0238|
|  - professional_law                   |Yaml   |none  |acc   |0.4433|±  |0.0127|
|  - world_religions                    |Yaml   |none  |acc   |0.8363|±  |0.0284|
| - other                               |N/A    |none  |acc   |0.6987|±  |0.1048|
|  - business_ethics                    |Yaml   |none  |acc   |0.5800|±  |0.0496|
|  - clinical_knowledge                 |Yaml   |none  |acc   |0.7019|±  |0.0282|
|  - college_medicine                   |Yaml   |none  |acc   |0.6474|±  |0.0364|
|  - global_facts                       |Yaml   |none  |acc   |0.3900|±  |0.0490|
|  - human_aging                        |Yaml   |none  |acc   |0.6502|±  |0.0320|
|  - management                         |Yaml   |none  |acc   |0.7864|±  |0.0406|
|  - marketing                          |Yaml   |none  |acc   |0.8590|±  |0.0228|
|  - medical_genetics                   |Yaml   |none  |acc   |0.7400|±  |0.0441|
|  - miscellaneous                      |Yaml   |none  |acc   |0.8148|±  |0.0139|
|  - nutrition                          |Yaml   |none  |acc   |0.7418|±  |0.0251|
|  - professional_accounting            |Yaml   |none  |acc   |0.4929|±  |0.0298|
|  - professional_medicine              |Yaml   |none  |acc   |0.6618|±  |0.0287|
|  - virology                           |Yaml   |none  |acc   |0.5482|±  |0.0387|
| - social_sciences                     |N/A    |none  |acc   |0.7361|±  |0.0640|
|  - econometrics                       |Yaml   |none  |acc   |0.5000|±  |0.0470|
|  - high_school_geography              |Yaml   |none  |acc   |0.7727|±  |0.0299|
|  - high_school_government_and_politics|Yaml   |none  |acc   |0.8808|±  |0.0234|
|  - high_school_macroeconomics         |Yaml   |none  |acc   |0.6667|±  |0.0239|
|  - high_school_microeconomics         |Yaml   |none  |acc   |0.6597|±  |0.0308|
|  - high_school_psychology             |Yaml   |none  |acc   |0.8202|±  |0.0165|
|  - human_sexuality                    |Yaml   |none  |acc   |0.7939|±  |0.0355|
|  - professional_psychology            |Yaml   |none  |acc   |0.6716|±  |0.0190|
|  - public_relations                   |Yaml   |none  |acc   |0.6636|±  |0.0453|
|  - security_studies                   |Yaml   |none  |acc   |0.7551|±  |0.0275|
|  - sociology                          |Yaml   |none  |acc   |0.8209|±  |0.0271|
|  - us_foreign_policy                  |Yaml   |none  |acc   |0.8300|±  |0.0378|
| - stem                                |N/A    |none  |acc   |0.5268|±  |0.1263|
|  - abstract_algebra                   |Yaml   |none  |acc   |0.3200|±  |0.0469|
|  - anatomy                            |Yaml   |none  |acc   |0.6296|±  |0.0417|
|  - astronomy                          |Yaml   |none  |acc   |0.6645|±  |0.0384|
|  - college_biology                    |Yaml   |none  |acc   |0.7431|±  |0.0365|
|  - college_chemistry                  |Yaml   |none  |acc   |0.4800|±  |0.0502|
|  - college_computer_science           |Yaml   |none  |acc   |0.5200|±  |0.0502|
|  - college_mathematics                |Yaml   |none  |acc   |0.4200|±  |0.0496|
|  - college_physics                    |Yaml   |none  |acc   |0.4510|±  |0.0495|
|  - computer_security                  |Yaml   |none  |acc   |0.7800|±  |0.0416|
|  - conceptual_physics                 |Yaml   |none  |acc   |0.5489|±  |0.0325|
|  - electrical_engineering             |Yaml   |none  |acc   |0.5655|±  |0.0413|
|  - elementary_mathematics             |Yaml   |none  |acc   |0.3915|±  |0.0251|
|  - high_school_biology                |Yaml   |none  |acc   |0.7548|±  |0.0245|
|  - high_school_chemistry              |Yaml   |none  |acc   |0.5222|±  |0.0351|
|  - high_school_computer_science       |Yaml   |none  |acc   |0.6900|±  |0.0465|
|  - high_school_mathematics            |Yaml   |none  |acc   |0.3222|±  |0.0285|
|  - high_school_physics                |Yaml   |none  |acc   |0.3444|±  |0.0388|
|  - high_school_statistics             |Yaml   |none  |acc   |0.5139|±  |0.0341|
|  - machine_learning                   |Yaml   |none  |acc   |0.4643|±  |0.0473|

|      Groups      |Version|Filter|Metric|Value |   |Stderr|
|------------------|-------|------|------|-----:|---|-----:|
|mmlu              |N/A    |none  |acc   |0.6236|±  |0.1269|
| - humanities     |N/A    |none  |acc   |0.5651|±  |0.1289|
| - other          |N/A    |none  |acc   |0.6987|±  |0.1048|
| - social_sciences|N/A    |none  |acc   |0.7361|±  |0.0640|
| - stem           |N/A    |none  |acc   |0.5268|±  |0.1263|

### Citations
Please feel free to raise a PR if there is any missing citation.

@misc{tunstall2023zephyr,
      title={Zephyr: Direct Distillation of LM Alignment}, 
      author={Lewis Tunstall and Edward Beeching and Nathan Lambert and Nazneen Rajani and Kashif Rasul and Younes Belkada and Shengyi Huang and Leandro von Werra and Clémentine Fourrier and Nathan Habib and Nathan Sarrazin and Omar Sanseviero and Alexander M. Rush and Thomas Wolf},
      year={2023},
      eprint={2310.16944},
      archivePrefix={arXiv},
      primaryClass={cs.LG}
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