---
license: other
pipeline_tag: text-generation
tags:
- causal_lm
base_model: Walmart-the-bag/zephyr-quiklang-3b
inference: false
datasets:
- teknium/openhermes
---

# Description
This is the 4K version of https://huggingface.co/Walmart-the-bag/zephyr-quiklang-3b with 1000 more samples of openhermes.

# Original Model Description
This is a finetune of [StableLM-Zephyr-3B](https://huggingface.co/stabilityai/stablelm-zephyr-3b) with 2 datasets, toxic-dpo and openhermes with 10000 samples.

# Training Parameters
- 1xA6000-48GB
- batch_size: 6
- learning_rate: 5e-5

# Datasets:
- unalignment/toxic-dpo-v0.1
- teknium/openhermes

# Metrics/Basic Eval:
    "predict_bleu-4": 31.594154999999997,
    "predict_rouge-1": 44.092935,
    "predict_rouge-2": 22.276081000000005,
    "predict_rouge-l": 34.506909,
    "predict_runtime": 121.7549,
    "predict_samples_per_second": 0.821,
    "predict_steps_per_second": 0.107
  