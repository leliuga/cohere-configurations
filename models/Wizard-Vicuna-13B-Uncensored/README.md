---
license: other
datasets:
- ehartford/wizard_vicuna_70k_unfiltered
language:
- en
tags:
- uncensored
---

This is [wizard-vicuna-13b](https://huggingface.co/junelee/wizard-vicuna-13b) trained with a subset of the dataset - responses that contained alignment / moralizing were removed. The intent is to train a WizardLM that doesn't have alignment built-in, so that alignment (of any sort) can be added separately with for example with a RLHF LoRA.

Shout out to the open source AI/ML community, and everyone who helped me out.

Note:  

An uncensored model has no guardrails.  

You are responsible for anything you do with the model, just as you are responsible for anything you do with any dangerous object such as a knife, gun, lighter, or car.

Publishing anything this model generates is the same as publishing it yourself.

You are responsible for the content you publish, and you cannot blame the model any more than you can blame the knife, gun, lighter, or car for what you do with it.
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ehartford__Wizard-Vicuna-13B-Uncensored)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 49.52   |
| ARC (25-shot)         | 58.96          |
| HellaSwag (10-shot)   | 81.95    |
| MMLU (5-shot)         | 47.92         |
| TruthfulQA (0-shot)   | 51.69   |
| Winogrande (5-shot)   | 75.69   |
| GSM8K (5-shot)        | 8.64        |
| DROP (3-shot)         | 21.79         |
