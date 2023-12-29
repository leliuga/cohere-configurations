---
license: cc-by-nc-4.0
language:
- en
- de
library_name: transformers
pipeline_tag: text-generation
tags:
- finetune
- dpo
- Instruct
- augmentation
- german
datasets:
- argilla/distilabel-math-preference-dpo
---

![SauerkrautLM](https://vago-solutions.de/wp-content/uploads/2023/12/sauerkrautlm-solar.png "SauerkrautLM-SOLAR-Instruct")
## VAGO solutions SauerkrautLM-SOLAR-Instruct
Introducing **SauerkrautLM-SOLAR-Instruct** – our Sauerkraut version of the powerful [upstage/SOLAR-10.7B-Instruct-v1.0](https://huggingface.co/upstage/SOLAR-10.7B-Instruct-v1.0) ! 
Aligned with **DPO**

# Table of Contents
1. [Overview of all SauerkrautLM-SOLAR-Instruct models](#all-sauerkrautlm-solar-instruct-models)
2. [Model Details](#model-details)
   - [Prompt template](#prompt-template)
   - [Training Dataset](#training-dataset)
   - [Data Contamination Test](#data-contamination-test-results)
3. [Evaluation](#evaluation)
5. [Disclaimer](#disclaimer)
6. [Contact](#contact)
7. [Collaborations](#collaborations)
8. [Acknowledgement](#acknowledgement)


## All SauerkrautLM-SOLAR-Instruct Models

| Model | HF    | GPTQ  | GGUF  | AWQ  |
|-------|-------|-------|-------|-------|
| SauerkrautLM-SOLAR-Instruct  | [Link](https://huggingface.co/VAGOsolutions/SauerkrautLM-SOLAR-Instruct/) | coming soon | coming soon | coming soon |

## Model Details
**SauerkrautLM-SOLAR-Instruct**
- **Model Type:** SauerkrautLM-SOLAR-Instruct is a finetuned Model based on [upstage/SOLAR-10.7B-Instruct-v1.0](https://huggingface.co/upstage/SOLAR-10.7B-Instruct-v1.0) 
- **Language(s):** English, German
- **License:** cc-by-nc-4.0
- **Contact:** [Website](https://vago-solutions.de/#Kontakt) [David Golchinfar](mailto:golchinfar@vago-solutions.de)

### Training Dataset:

SauerkrautLM-SOLAR-Instruct was trained with mix of German data augmentation and translated data. 
Aligned through **DPO** with our **new German SauerkrautLM-DPO dataset** based on parts of the SFT SauerkrautLM dataset 
as chosen answers and [Sauerkraut-7b-HerO](https://huggingface.co/VAGOsolutions/SauerkrautLM-7b-HerO) as rejected answers. Added with additional **translated Parts of the [HuggingFaceH4/ultrafeedback_binarized](https://huggingface.co/datasets/HuggingFaceH4/ultrafeedback_binarized)** (Our dataset do not contain any TruthfulQA prompts - check Data Contamination Test Results) and **[argilla/distilabel-math-preference-dpo](https://huggingface.co/datasets/argilla/distilabel-math-preference-dpo).**  
We found, that only a simple translation of training data can lead to unnatural German phrasings. 
Data augmentation techniques were used to grant grammatical, syntactical correctness and a more natural German wording in our training data. 

We improved the German language skills on this model. Nevertheless, certain formulations may occur that are not entirely correct.





### Data Contamination Test Results

Some models on the HuggingFace leaderboard had problems with wrong data getting mixed in.
We checked our SauerkrautLM-DPO dataset with a special test [1] on this model as target model and upstage/SOLAR-10.7B-Instruct-v1.0 as reference model. 
The HuggingFace team used the same methods [2, 3].

Our results, with `result < 0.1, %:` being well below 0.9, indicate that our dataset is free from contamination.

*The data contamination test results of HellaSwag and Winograde will be added once [1] supports them.*

| Dataset                        | ARC   | MMLU | TruthfulQA | GSM8K |
|------------------------------|-------|-------|-------|-------|
| **SauerkrautLM-DPO**| result < 0.1, %: 0.0 |result < 0.1, %: 0.09 | result < 0.1, %: 0.13 | result < 0.1, %: 0.16 |

[1] https://github.com/swj0419/detect-pretrain-code-contamination

[2] https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard/discussions/474#657f2245365456e362412a06

[3] https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard/discussions/265#657b6debf81f6b44b8966230

### Prompt Template:
```
### User:
Hallo, wie geht es dir?

### Assistant:
Hallo! Es freut mich, dass du mit mir kommunizierst. Ich bin hier, um zu helfen und deine Anfragen zu erfüllen. Du fragst, wie ich mich fühle. Als künstliche Intelligenz habe ich keine eigentlichen Emotionen im Sinne eines Menschen, aber ich funktioniere optimal und bin bereit, Dienste anzubieten.
Wie geht es dir momentan? Können wir zusammen etwas interessantes oder hilfreiches erledigen?

```
*Prompt Example on Temp 0.5

```
### User:
Hello, how are you?

### Assistant:
Hi there! I am an AI language model, so I don't have personal feelings or emotions in the traditional sense. However, I can assure you that my systems and processes are functioning well at this moment, allowing me to provide helpful responses for your queries.
How may I assist you today?

```
*Prompt Example on Temp 0.5

## Evaluation



| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 74.21 |
| ARC (25-shot)         | 70.82         |
| HellaSwag (10-shot)   | 88.63   |
| MMLU (5-shot)         | 66.2|
| TruthfulQA (0-shot)   | 71.95 |
| Winogrande (5-shot)   | 83.5  |
| GSM8K (5-shot)        | 64.14        |

## Disclaimer
We must inform users that despite our best efforts in data cleansing, the possibility of uncensored content slipping through cannot be entirely ruled out.
However, we cannot guarantee consistently appropriate behavior. Therefore, if you encounter any issues or come across inappropriate content, we kindly request that you inform us through the contact information provided.
Additionally, it is essential to understand that the licensing of these models does not constitute legal advice. We are not held responsible for the actions of third parties who utilize our models.
 
## Contact
If you are interested in customized LLMs for business applications, please get in contact with us via our website or contact us at [Dr. Daryoush Vaziri](mailto:vaziri@vago-solutions.de). We are also grateful for your feedback and suggestions.
 
## Collaborations
We are also keenly seeking support and investment for our startup, VAGO solutions, where we continuously advance the development of robust language models designed to address a diverse range of purposes and requirements. If the prospect of collaboratively navigating future challenges excites you, we warmly invite you to reach out to us.

## Acknowledgement
Many thanks to [argilla](https://huggingface.co/datasets/argilla) and [Huggingface](https://huggingface.co) for providing such valuable datasets to the Open-Source community. And of course a big thanks to [upstage](https://huggingface.co/upstage) for providing the open source community with their latest technology!
Many thanks to [TheBloke](https://huggingface.co/TheBloke) for super fast quantifying all of our models. 