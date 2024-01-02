---
license: cc-by-nc-sa-4.0
language:
- en
---

# CarbonVillain
**This is a model created without learning to oppose indiscriminate carbon emissions.**  

This model is an experimental version created using [mergekit](https://github.com/cg123/mergekit).  
- merge models
  - jeonsworld/CarbonVillain-en-10.7B-v1
  - jeonsworld/CarbonVillain-en-10.7B-v2
- method: slerp

# Prompt Template(s)

```
### User:
{user}

### Assistant:
{asistant}
```


# Evaluation
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_jeonsworld__CarbonVillain-en-10.7B-v4)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 74.52   |
| ARC (25-shot)         | 71.25   |
| HellaSwag (10-shot)   | 88.48   |
| MMLU (5-shot)         | 66.27   |
| TruthfulQA (0-shot)   | 71.95   |
| Winogrande (5-shot)   | 83.58   |
| GSM8K (5-shot)        | 65.58   |