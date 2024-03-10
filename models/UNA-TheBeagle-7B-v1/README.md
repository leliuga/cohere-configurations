---
license: cc-by-nc-nd-4.0
tags:
- generated_from_trainer
model-index:
- name: UNA-TheBeagle-7b-v1
  results: []
datasets:
- jondurbin/bagel-v0.3
library_name: transformers
---
     -- In the Love Memory of my "LoLa" --
     
# UNA-TheBeagle-7b-v1
TheBeagle, a model of 7B parameters trained on The Bagel dataset. DPO & UNA applied over a set of curated DPO Pairs.

- Scored #1 on the HF Leaderboard, dramatic scores!!! 73 ARC, and very well balanced!

The dataset was generated using the original bagel code, including the decontamination step.

As base model, we used the latest Intel's neural-chat model.

It performs very good in many tasks, but its always better that you play with it by yourself.

![TheBeagle](https://huggingface.co/fblgit/UNA-TheBeagle-7b-v1/resolve/main/TheBeagle.png)

## Evaluations

Ran with VLLM so expect them to dont be exactly as the one's shown in the board, but not too far :)

```
vllm (pretrained=fblgit/UNA-TheBeagle-7b-v1,dtype=auto,tensor_parallel_size=1,gpu_memory_utilization=0.8,data_parallel_size=8,trust_remote_code=True), gen_kwargs: (None), limit: None, num_fewshot: None, batch_size: 32
|    Tasks     |Version|  Filter  |n-shot|  Metric   |Value |   |Stderr|
|--------------|-------|----------|-----:|-----------|-----:|---|-----:|
|arc_challenge |Yaml   |none      |    25|acc        |0.7090|±  |0.0133|
|              |       |none      |    25|acc_norm   |0.7329|±  |0.0129|
|gsm8k         |Yaml   |get-answer|     5|exact_match|0.7210|±  |0.0124|
|hellaswag     |Yaml   |none      |    10|acc        |0.7202|±  |0.0045|
|              |       |none      |    10|acc_norm   |0.8792|±  |0.0033|
|truthfulqa_mc2|Yaml   |none      |     0|acc        |0.7062|±  |0.0151|
|winogrande    |Yaml   |none      |     5|acc        |0.8366|±  |0.0104|
```

## UNA Details

For this release, we only applied UNA thru the perceptrons. It was done at a 3.5e-7 speed, and the training loop code is also the original one of the bagel and transformers-4.35.2-UNA

## Prompt

Im not entirely sure of it, as we used the vanilla version of the bagel training code. But a good model should be able to generalize with different prompt formats, so feel free to give it a shot.

## Citations

Remember if you use UNA's models, cite it in your model card.

## Limitations
Not for commercial use, and only for academic & research purposes.