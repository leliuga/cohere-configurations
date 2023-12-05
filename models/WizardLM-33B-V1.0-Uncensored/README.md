---
license: other
datasets:
- ehartford/WizardLM_evol_instruct_V2_196k_unfiltered_merged_split
language:
- en
---

This is a retraining of https://huggingface.co/WizardLM/WizardLM-30B-V1.0 with a filtered dataset, intended to reduce refusals, avoidance, and bias.

Note that LLaMA itself has inherent ethical beliefs, so there's no such thing as a "truly uncensored" model.  But this model will be more compliant than WizardLM/WizardLM-7B-V1.0.

Shout out to the open source AI/ML community, and everyone who helped me out.

Note: An uncensored model has no guardrails. You are responsible for anything you do with the model, just as you are responsible for anything you do with any dangerous object such as a knife, gun, lighter, or car. Publishing anything this model generates is the same as publishing it yourself. You are responsible for the content you publish, and you cannot blame the model any more than you can blame the knife, gun, lighter, or car for what you do with it.

Like WizardLM/WizardLM-30B-V1.0, this model is trained with Vicuna-1.1 style prompts.

```
You are a helpful AI assistant.

USER: <prompt>
ASSISTANT:
```

Thank you [chirper.ai](https://chirper.ai) for sponsoring some of my compute!
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ehartford__WizardLM-33B-V1.0-Uncensored)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 54.41   |
| ARC (25-shot)         | 63.65          |
| HellaSwag (10-shot)   | 83.84    |
| MMLU (5-shot)         | 59.36         |
| TruthfulQA (0-shot)   | 56.8   |
| Winogrande (5-shot)   | 77.66   |
| GSM8K (5-shot)        | 18.65        |
| DROP (3-shot)         | 20.89         |
