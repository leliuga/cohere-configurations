---
license: cc-by-nc-4.0
datasets:
- Open-Orca/OpenOrca
language:
- en
pipeline_tag: text-generation
---
# Luban-13B

# Model Details
* **Trained by**: trained by AIDC AI-Business.
* **Model type:**  **Luban-13B** is an auto-regressive language model based on the Llama 2 transformer architecture.
* **Language(s)**: English
* **License for Luban-13B base weights**: Non-Commercial Creative Commons license ([CC BY-NC-4.0](https://creativecommons.org/licenses/by-nc/4.0/))


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
| Avg.                  |   65.03   |
| ARC (25-shot)         |   63.05   |
| HellaSwag (10-shot)   |   82.8   |
| MMLU (5-shot)         |   58.73   |
| TruthfulQA (0-shot)   |   55.53   |