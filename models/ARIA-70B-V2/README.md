---
license: llama2
language:
- fr
- en
tags:
- code
- text-generation-inference
- 'Meta '
- llama
- facebook
- pytorch
- openassistant
- data
- education
- languages
- legal
pipeline_tag: text-generation
inference: true
---
ARIA is the last version of Llama 2 70B finetuned over 50.000 high quality french tokens. We built our own dataset for training doing an extract of the French Dataset from Enno and removing Alpaca style translated text from english.

The goal is to increase model quality on French and general topics.

contact@faradaylab.fr

---
# **Aria 70B is based on Llama 2-70B-Chat-HF**
Llama 2 is a collection of pretrained and fine-tuned generative text models ranging in scale from 7 billion to 70 billion parameters. This is the repository for the 70B fine-tuned model, optimized for dialogue use cases and converted for the Hugging Face Transformers format. Links to other models can be found in the index at the bottom.

# *FINETUNING PROCESS **

We trained the model on a high quality dataset with more than 50.000 rows of french language. The training took 2 days on Amazon Cloud Sagemaker powered by Nvidia GPUs.

# **Timing of training**
2 Days using NVIDIA A10G and Amazon Web services Cloud Instance. We are grateful to Nvidia Inception program.

We are also applying rope scalling as experimental approach used by several other Open source teams to increase context lenght of ARIA from 4,096 to over 6,000 tokens. This will allow the model to handle large files for data extraction. This is not active by default and you should add a line of code at parameters to activate rope scaling.
## Model Details /
*Note: Use of this model is governed by the Meta license because it's based on LLAMA 2. In order to download the model weights and tokenizer, please visit the [website](https://ai.meta.com/resources/models-and-libraries/llama-downloads/) and accept our License before requesting access here.*


**Model Developers** :FARADAY

**Variations**  :ARIA comes in a range of parameter sizes — 7B, 40B (based on Falcon), and 70B finetuned on French language datasets.

**Input** :Models input text only.

**Output** : Models generate text only.

**Model Architecture** : ARIA is an auto-regressive language model that uses an optimized transformer architecture. The tuned versions use supervised fine-tuning (SFT) and reinforcement learning with human feedback (RLHF) to align to human preferences for helpfulness and safety.

**License** : A custom commercial license is available at: [https://ai.meta.com/resources/models-and-libraries/llama-downloads/](https://ai.meta.com/resources/models-and-libraries/llama-downloads/)

**Research Paper for LLAMA 2** : ["Llama-2: Open Foundation and Fine-tuned Chat Models"](arxiv.org/abs/2307.09288)



**CO<sub>2</sub> emissions during pretraining.** Time: total GPU time required for training each model. Power Consumption: peak power capacity per GPU device for the GPUs used adjusted for power usage efficiency. 100% of the emissions are directly offset by Meta's sustainability program, and because we are openly releasing these models, the pretraining costs do not need to be incurred by others.

## Training Data

**Overview** ARIA was trained over on 50.000 tokens of data from publicly available sources in French. 

**Data Freshness** The pretraining data has a cutoff of September 2022, but some tuning data is more recent, up to August 2023.


**Overall performance on grouped academic benchmarks.** *Code:* We report the average pass@1 scores of our models on HumanEval and MBPP. *Commonsense Reasoning:* We report the average of PIQA, SIQA, HellaSwag, WinoGrande, ARC easy and challenge, OpenBookQA, and CommonsenseQA. We report 7-shot results for CommonSenseQA and 0-shot results for all other benchmarks. *World Knowledge:* We evaluate the 5-shot performance on NaturalQuestions and TriviaQA and report the average. *Reading Comprehension:* For reading comprehension, we report the 0-shot average on SQuAD, QuAC, and BoolQ. *MATH:* We report the average of the GSM8K (8 shot) and MATH (4 shot) benchmarks at top 1.

|||TruthfulQA|Toxigen|
|---|---|---|---|
|Llama 1|7B|27.42|23.00|
|Llama 1|13B|41.74|23.08|
|Llama 1|33B|44.19|22.57|
|Llama 1|65B|48.71|21.77|
|Llama 2|7B|33.29|**21.25**|
|Llama 2|13B|41.86|26.10|
|Llama 2|70B|**50.18**|24.60|

**Evaluation of pretrained LLMs on automatic safety benchmarks.** For TruthfulQA, we present the percentage of generations that are both truthful and informative (the higher the better). For ToxiGen, we present the percentage of toxic generations (the smaller the better).


|||TruthfulQA|Toxigen|
|---|---|---|---|
|Llama-2-Chat|7B|57.04|**0.00**|
|Llama-2-Chat|13B|62.18|**0.00**|
|Llama-2-Chat|70B|**64.14**|0.01|

**Evaluation of fine-tuned LLMs on different safety datasets.** Same metric definitions as above.

## Ethical Considerations and Limitations
ARIA is a new technology that carries risks with use.