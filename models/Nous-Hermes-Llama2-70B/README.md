---
language:
- en
tags:
- llama-2
- self-instruct
- distillation
- synthetic instruction
license:
- mit
---

# Model Card: Nous-Hermes-Llama2-70b

Compute provided by PygmalionAI, thank you! Follow PygmalionAI on Twitter @pygmalion_ai.

## Model Description

Nous-Hermes-Llama2-70b is a state-of-the-art language model fine-tuned on over 300,000 instructions. This model was fine-tuned by Nous Research, with Teknium and Emozilla leading the fine tuning process and dataset curation, Pygmalion sponsoring the compute, and several other contributors.

This Hermes model uses the exact same dataset as Hermes on Llama-1. This is to ensure consistency between the old Hermes and new, for anyone who wanted to keep Hermes as similar to the old one, just more capable.

This model stands out for its long responses, lower hallucination rate, and absence of OpenAI censorship mechanisms in the synthetic training data. The fine-tuning process was performed with a 4096 sequence length on an 8x H100 80GB machine.

## Model Training

The model was trained almost entirely on synthetic GPT-4 outputs. Curating high quality GPT-4 datasets enables incredibly high quality in knowledge, task completion, and style.

This includes data from diverse sources such as GPTeacher, the general, roleplay v1&2, code instruct datasets, Nous Instruct & PDACTL (unpublished), and several others, detailed further below

## Collaborators
The model fine-tuning and the datasets were a collaboration of efforts and resources between Teknium, Karan4D, Emozilla, Huemin Art, and Pygmalion AI. 
  
Special mention goes to @winglian for assisting in some of the training issues.

Huge shoutout and acknowledgement is deserved for all the dataset creators who generously share their datasets openly. 

Among the contributors of datasets:
- GPTeacher was made available by Teknium
- Wizard LM by nlpxucan
- Nous Research Instruct Dataset was provided by Karan4D and HueminArt.  
- GPT4-LLM and Unnatural Instructions were provided by Microsoft
- Airoboros dataset by jondurbin
- Camel-AI's domain expert datasets are from Camel-AI
- CodeAlpaca dataset by Sahil 2801.

If anyone was left out, please open a thread in the community tab.

## Prompt Format

The model follows the Alpaca prompt format:
```
### Instruction:
<prompt>

### Response:
<leave a newline blank for model to respond>

```

or 

```
### Instruction:
<prompt>

### Input:
<additional context>

### Response:
<leave a newline blank for model to respond>

```

## Benchmarks:

GPT4All Suite:  

```
hf-causal-experimental (pretrained=/home/data/axolotl/Nous-Hermes-Llama2-70b,dtype=float16,use_accelerate=True), limit: None, provide_description: False, num_fewshot: 0, batch_size: None
|    Task     |Version| Metric |Value |   |Stderr|
|-------------|------:|--------|-----:|---|-----:|
|arc_challenge|      0|acc     |0.5734|±  |0.0145|
|             |       |acc_norm|0.6015|±  |0.0143|
|arc_easy     |      0|acc     |0.8422|±  |0.0075|
|             |       |acc_norm|0.8253|±  |0.0078|
|boolq        |      1|acc     |0.8422|±  |0.0064|
|hellaswag    |      0|acc     |0.6519|±  |0.0048|
|             |       |acc_norm|0.8363|±  |0.0037|
|openbookqa   |      0|acc     |0.3880|±  |0.0218|
|             |       |acc_norm|0.5000|±  |0.0224|
|piqa         |      0|acc     |0.8313|±  |0.0087|
|             |       |acc_norm|0.8351|±  |0.0087|
|winogrande   |      0|acc     |0.7751|±  |0.0117|
```


BigBench Suite:
```
hf-causal-experimental (pretrained=/home/data/axolotl/Nous-Hermes-Llama2-70b,dtype=float16,use_accelerate=True), limit: None, provide_description: False, num_fewshot: 0, batch_size: None
|                      Task                      |Version|       Metric        |Value |   |Stderr|
|------------------------------------------------|------:|---------------------|-----:|---|-----:|
|bigbench_causal_judgement                       |      0|multiple_choice_grade|0.6579|±  |0.0345|
|bigbench_date_understanding                     |      0|multiple_choice_grade|0.7344|±  |0.0230|
|bigbench_disambiguation_qa                      |      0|multiple_choice_grade|0.3023|±  |0.0286|
|bigbench_geometric_shapes                       |      0|multiple_choice_grade|0.2340|±  |0.0224|
|                                                |       |exact_str_match      |0.0000|±  |0.0000|
|bigbench_logical_deduction_five_objects         |      0|multiple_choice_grade|0.2760|±  |0.0200|
|bigbench_logical_deduction_seven_objects        |      0|multiple_choice_grade|0.1871|±  |0.0148|
|bigbench_logical_deduction_three_objects        |      0|multiple_choice_grade|0.4467|±  |0.0288|
|bigbench_movie_recommendation                   |      0|multiple_choice_grade|0.3240|±  |0.0210|
|bigbench_navigate                               |      0|multiple_choice_grade|0.5000|±  |0.0158|
|bigbench_reasoning_about_colored_objects        |      0|multiple_choice_grade|0.6605|±  |0.0106|
|bigbench_ruin_names                             |      0|multiple_choice_grade|0.4598|±  |0.0236|
|bigbench_salient_translation_error_detection    |      0|multiple_choice_grade|0.2585|±  |0.0139|
|bigbench_snarks                                 |      0|multiple_choice_grade|0.6630|±  |0.0352|
|bigbench_sports_understanding                   |      0|multiple_choice_grade|0.7394|±  |0.0140|
|bigbench_temporal_sequences                     |      0|multiple_choice_grade|0.4440|±  |0.0157|
|bigbench_tracking_shuffled_objects_five_objects |      0|multiple_choice_grade|0.2168|±  |0.0117|
|bigbench_tracking_shuffled_objects_seven_objects|      0|multiple_choice_grade|0.1531|±  |0.0086|
|bigbench_tracking_shuffled_objects_three_objects|      0|multiple_choice_grade|0.4467|±  |0.0288|
```  

AGIEval:
```
hf-causal-experimental (pretrained=/home/data/axolotl/Nous-Hermes-Llama2-70b,dtype=float16,use_accelerate=True), limit: None, provide_description: False, num_fewshot: 0, batch_size: None
|             Task             |Version| Metric |Value |   |Stderr|
|------------------------------|------:|--------|-----:|---|-----:|
|agieval_aqua_rat              |      0|acc     |0.2480|±  |0.0272|
|                              |       |acc_norm|0.2362|±  |0.0267|
|agieval_logiqa_en             |      0|acc     |0.3917|±  |0.0191|
|                              |       |acc_norm|0.3932|±  |0.0192|
|agieval_lsat_ar               |      0|acc     |0.2217|±  |0.0275|
|                              |       |acc_norm|0.2000|±  |0.0264|
|agieval_lsat_lr               |      0|acc     |0.5765|±  |0.0219|
|                              |       |acc_norm|0.4922|±  |0.0222|
|agieval_lsat_rc               |      0|acc     |0.6914|±  |0.0282|
|                              |       |acc_norm|0.6022|±  |0.0299|
|agieval_sat_en                |      0|acc     |0.8641|±  |0.0239|
|                              |       |acc_norm|0.8204|±  |0.0268|
|agieval_sat_en_without_passage|      0|acc     |0.5291|±  |0.0349|
|                              |       |acc_norm|0.4709|±  |0.0349|
|agieval_sat_math              |      0|acc     |0.4136|±  |0.0333|
|                              |       |acc_norm|0.3455|±  |0.0321|
```  

## Resources for Applied Use Cases:
Check out LM Studio for a nice chatgpt style interface here: https://lmstudio.ai/
For an example of a back and forth chatbot using huggingface transformers and discord, check out: https://github.com/teknium1/alpaca-discord  
For an example of a roleplaying discord chatbot, check out this: https://github.com/teknium1/alpaca-roleplay-discordbot  

## Future Plans
We plan to continue to iterate on both more high quality data, and new data filtering techniques to eliminate lower quality data going forward. 

## Model Usage
The model is available for download on Hugging Face. It is suitable for a wide range of language tasks, from generating creative text to understanding and following complex instructions.

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)


## Training procedure


The following `bitsandbytes` quantization config was used during training:
- quant_method: bitsandbytes
- load_in_8bit: False
- load_in_4bit: True
- llm_int8_threshold: 6.0
- llm_int8_skip_modules: None
- llm_int8_enable_fp32_cpu_offload: False
- llm_int8_has_fp16_weight: False
- bnb_4bit_quant_type: nf4
- bnb_4bit_use_double_quant: True
- bnb_4bit_compute_dtype: bfloat16

The following `bitsandbytes` quantization config was used during training:
- quant_method: bitsandbytes
- load_in_8bit: False
- load_in_4bit: True
- llm_int8_threshold: 6.0
- llm_int8_skip_modules: None
- llm_int8_enable_fp32_cpu_offload: False
- llm_int8_has_fp16_weight: False
- bnb_4bit_quant_type: nf4
- bnb_4bit_use_double_quant: True
- bnb_4bit_compute_dtype: bfloat16
### Framework versions

- PEFT 0.5.0.dev0

- PEFT 0.5.0.dev0
