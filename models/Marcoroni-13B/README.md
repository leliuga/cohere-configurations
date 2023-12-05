---
license: cc-by-nc-4.0
datasets:
- Open-Orca/OpenOrca
language:
- en
pipeline_tag: text-generation
---
# Marcoroni-13B

# Model Details
* **Trained by**: trained by AIDC AI-Business.
* **Model type:**  **Marcoroni-13B** is an auto-regressive language model based on the Llama 2 transformer architecture.
* **Language(s)**: English
* **License for Marcoroni-13B base weights**: Non-Commercial Creative Commons license ([CC BY-NC-4.0](https://creativecommons.org/licenses/by-nc/4.0/))


# Prompting

## Prompt Template for alpaca style

```
### Instruction:

<prompt> (without the <>)

### Response:
```

# Evulation Results ([Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard))

| Metric                | Value |
|-----------------------|-------|
| Avg.                  |   65.76   |
| ARC (25-shot)         |   62.46   |
| HellaSwag (10-shot)   |   83.27   |
| MMLU (5-shot)         |   59.63   |
| TruthfulQA (0-shot)   |   57.7   |