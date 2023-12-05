---
license: other
datasets:
- ehartford/WizardLM_alpaca_evol_instruct_70k_unfiltered
tags:
- uncensored
---
This is WizardLM trained with a subset of the dataset - responses that contained alignment / moralizing were removed.  The intent is to train a WizardLM that doesn't have alignment built-in, so that alignment (of any sort) can be added separately with for example with a RLHF LoRA.

Shout out to the open source AI/ML community, and everyone who helped me out.

Note:  
An uncensored model has no guardrails.  
You are responsible for anything you do with the model, just as you are responsible for anything you do with any dangerous object such as a knife, gun, lighter, or car.
Publishing anything this model generates is the same as publishing it yourself.
You are responsible for the content you publish, and you cannot blame the model any more than you can blame the knife, gun, lighter, or car for what you do with it.
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ehartford__WizardLM-30B-Uncensored)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 52.32   |
| ARC (25-shot)         | 60.24          |
| HellaSwag (10-shot)   | 82.93    |
| MMLU (5-shot)         | 56.8         |
| TruthfulQA (0-shot)   | 51.57   |
| Winogrande (5-shot)   | 74.35   |
| GSM8K (5-shot)        | 12.89        |
| DROP (3-shot)         | 27.45         |
