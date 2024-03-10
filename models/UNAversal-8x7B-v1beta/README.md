---
language:
- en
license: cc-by-nc-sa-4.0
library_name: transformers
tags:
- UNA
- juanako
- mixtral
- MoE
model-index:
- name: UNAversal-8x7B-v1beta
  results:
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: AI2 Reasoning Challenge (25-Shot)
      type: ai2_arc
      config: ARC-Challenge
      split: test
      args:
        num_few_shot: 25
    metrics:
    - type: acc_norm
      value: 69.8
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=fblgit/UNAversal-8x7B-v1beta
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: HellaSwag (10-Shot)
      type: hellaswag
      split: validation
      args:
        num_few_shot: 10
    metrics:
    - type: acc_norm
      value: 86.9
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=fblgit/UNAversal-8x7B-v1beta
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: MMLU (5-Shot)
      type: cais/mmlu
      config: all
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 70.39
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=fblgit/UNAversal-8x7B-v1beta
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: TruthfulQA (0-shot)
      type: truthful_qa
      config: multiple_choice
      split: validation
      args:
        num_few_shot: 0
    metrics:
    - type: mc2
      value: 71.97
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=fblgit/UNAversal-8x7B-v1beta
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: Winogrande (5-shot)
      type: winogrande
      config: winogrande_xl
      split: validation
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 82.0
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=fblgit/UNAversal-8x7B-v1beta
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: GSM8k (5-shot)
      type: gsm8k
      config: main
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 61.64
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=fblgit/UNAversal-8x7B-v1beta
      name: Open LLM Leaderboard
---
# UNAversal - Uniform Neural Alignment (MoE)

This is just a beta, a first release so people can start working on franksteins and so.
It does achieve high GSM/Math and TQA, so ideally you can merge it with other mixtrals and see what coming out of it
Based on [mistralai/Mixtral-8x7B-Instruct-v0.1](https://huggingface.co/mistralai/Mixtral-8x7B-Instruct-v0.1)

## UNA Details
For this model we came out with the most obvious, placing UNA on the router_logit. It does work, but we saw a much better performance on SFT by doing so.
So this model DOES have UNA-SFT phase, its highly experimental and it was merely using LLaMA-Factory datasets by example alpaca.

As the others:
- Can be finetuned further, try 2e-5 or **1e-4 (since its MOE)**
- Can be merged, here you will have to improvise and please report findings on a discussion thread.

**REMINDER**: please.. cite, it does help on the research and the lab itself, seriously.

## NEED YOUR HELP!!
I need a multi-turn trainloop for the Mixtral, that can squeeze the juice out of 8xH100's properly. Please feel free to reach @fblgit either discord or twitter. thanks!

# Evals
Here there are some, but we also submitted it to the HF eval queue....

## GSM8k 5-Shot
```
|Tasks|Version|  Filter  |n-shot|  Metric   |Value |   |Stderr|
|-----|-------|----------|-----:|-----------|-----:|---|-----:|
|gsm8k|Yaml   |get-answer|     5|exact_match|0.6603|±  | 0.013|
```
## ARC 25-Shot
```
|    Tasks    |Version|Filter|n-shot| Metric |Value |   |Stderr|
|-------------|-------|------|-----:|--------|-----:|---|-----:|
|arc_challenge|Yaml   |none  |    25|acc     |0.6621|±  |0.0138|
|             |       |none  |    25|acc_norm|0.6962|±  |0.0134|
```

## TruthfulQA 0-Shot (MC2)
```
|    Tasks     |Version|Filter|n-shot|Metric|Value |   |Stderr|
|--------------|-------|------|-----:|------|-----:|---|-----:|
|truthfulqa_mc2|Yaml   |none  |     0|acc   |0.7122|±  |0.0141|
```

## 0-Shots Evals
```
|    Tasks     |Version|Filter|n-shot|  Metric  |Value |   |Stderr|
|--------------|-------|------|-----:|----------|-----:|---|-----:|
|arc_challenge |Yaml   |none  |     0|acc       |0.6101|±  |0.0143|
|              |       |none  |     0|acc_norm  |0.6425|±  |0.0140|
|arc_easy      |Yaml   |none  |     0|acc       |0.8615|±  |0.0071|
|              |       |none  |     0|acc_norm  |0.8375|±  |0.0076|
|boolq         |Yaml   |none  |     0|acc       |0.8624|±  |0.0060|
|lambada_openai|Yaml   |none  |     0|perplexity|2.8318|±  |0.0507|
|              |       |none  |     0|acc       |0.7650|±  |0.0059|
|mathqa        |Yaml   |none  |     0|acc       |0.4472|±  |0.0091|
|              |       |none  |     0|acc_norm  |0.4436|±  |0.0091|
|piqa          |Yaml   |none  |     0|acc       |0.8292|±  |0.0088|
|              |       |none  |     0|acc_norm  |0.8422|±  |0.0085|
|pubmedqa      |Yaml   |none  |     0|acc       |0.7920|±  |0.0182|
|sciq          |Yaml   |none  |     0|acc       |0.9630|±  |0.0060|
|              |       |none  |     0|acc_norm  |0.9370|±  |0.0077|
```

## BBH at 0-Shot
```
vllm (pretrained=fblgit/UNAversal-8x7B-v1beta,tensor_parallel_size=2,data_parallel_size=4,gpu_memory_utilization=0.8,dtype=float16), gen_kwargs: (None), limit: None, num_fewshot: 0, batch_size: auto
|                          Tasks                           |Version|  Filter  |n-shot|  Metric   |Value |   |Stderr|
|----------------------------------------------------------|-------|----------|-----:|-----------|-----:|---|-----:|
|bbh                                                       |N/A    |get-answer|     0|exact_match|0.6752|±  |0.1772|
| - bbh_cot_fewshot_boolean_expressions                    |Yaml   |get-answer|     0|exact_match|0.8840|±  |0.0203|
| - bbh_cot_fewshot_causal_judgement                       |Yaml   |get-answer|     0|exact_match|0.6417|±  |0.0352|
| - bbh_cot_fewshot_date_understanding                     |Yaml   |get-answer|     0|exact_match|0.7600|±  |0.0271|
| - bbh_cot_fewshot_disambiguation_qa                      |Yaml   |get-answer|     0|exact_match|0.7160|±  |0.0286|
| - bbh_cot_fewshot_dyck_languages                         |Yaml   |get-answer|     0|exact_match|0.1800|±  |0.0243|
| - bbh_cot_fewshot_formal_fallacies                       |Yaml   |get-answer|     0|exact_match|0.6520|±  |0.0302|
| - bbh_cot_fewshot_geometric_shapes                       |Yaml   |get-answer|     0|exact_match|0.3880|±  |0.0309|
| - bbh_cot_fewshot_hyperbaton                             |Yaml   |get-answer|     0|exact_match|0.9600|±  |0.0124|
| - bbh_cot_fewshot_logical_deduction_five_objects         |Yaml   |get-answer|     0|exact_match|0.5360|±  |0.0316|
| - bbh_cot_fewshot_logical_deduction_seven_objects        |Yaml   |get-answer|     0|exact_match|0.5040|±  |0.0317|
| - bbh_cot_fewshot_logical_deduction_three_objects        |Yaml   |get-answer|     0|exact_match|0.8600|±  |0.0220|
| - bbh_cot_fewshot_movie_recommendation                   |Yaml   |get-answer|     0|exact_match|0.7840|±  |0.0261|
| - bbh_cot_fewshot_multistep_arithmetic_two               |Yaml   |get-answer|     0|exact_match|0.6600|±  |0.0300|
| - bbh_cot_fewshot_navigate                               |Yaml   |get-answer|     0|exact_match|0.8160|±  |0.0246|
| - bbh_cot_fewshot_object_counting                        |Yaml   |get-answer|     0|exact_match|0.8360|±  |0.0235|
| - bbh_cot_fewshot_penguins_in_a_table                    |Yaml   |get-answer|     0|exact_match|0.7329|±  |0.0367|
| - bbh_cot_fewshot_reasoning_about_colored_objects        |Yaml   |get-answer|     0|exact_match|0.8120|±  |0.0248|
| - bbh_cot_fewshot_ruin_names                             |Yaml   |get-answer|     0|exact_match|0.4440|±  |0.0315|
| - bbh_cot_fewshot_salient_translation_error_detection    |Yaml   |get-answer|     0|exact_match|0.5200|±  |0.0317|
| - bbh_cot_fewshot_snarks                                 |Yaml   |get-answer|     0|exact_match|0.7135|±  |0.0340|
| - bbh_cot_fewshot_sports_understanding                   |Yaml   |get-answer|     0|exact_match|0.9400|±  |0.0151|
| - bbh_cot_fewshot_temporal_sequences                     |Yaml   |get-answer|     0|exact_match|0.7560|±  |0.0272|
| - bbh_cot_fewshot_tracking_shuffled_objects_five_objects |Yaml   |get-answer|     0|exact_match|0.5680|±  |0.0314|
| - bbh_cot_fewshot_tracking_shuffled_objects_seven_objects|Yaml   |get-answer|     0|exact_match|0.6280|±  |0.0306|
| - bbh_cot_fewshot_tracking_shuffled_objects_three_objects|Yaml   |get-answer|     0|exact_match|0.6280|±  |0.0306|
| - bbh_cot_fewshot_web_of_lies                            |Yaml   |get-answer|     0|exact_match|0.9560|±  |0.0130|
| - bbh_cot_fewshot_word_sorting                           |Yaml   |get-answer|     0|exact_match|0.3800|±  |0.0308|

|Groups|Version|  Filter  |n-shot|  Metric   |Value |   |Stderr|
|------|-------|----------|-----:|-----------|-----:|---|-----:|
|bbh   |N/A    |get-answer|     0|exact_match|0.6752|±  |0.1772|
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_fblgit__UNAversal-8x7B-v1beta)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |73.78|
|AI2 Reasoning Challenge (25-Shot)|69.80|
|HellaSwag (10-Shot)              |86.90|
|MMLU (5-Shot)                    |70.39|
|TruthfulQA (0-shot)              |71.97|
|Winogrande (5-shot)              |82.00|
|GSM8k (5-shot)                   |61.64|

