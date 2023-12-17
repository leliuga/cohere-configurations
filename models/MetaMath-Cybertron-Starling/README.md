---
license: cc-by-nc-4.0
datasets:
- meta-math/MetaMathQA
language:
- en
pipeline_tag: text-generation
tags:
- Math
---

## MetaMath-Cybertron-Starling


Merge [Q-bert/MetaMath-Cybertron](https://huggingface.co/Q-bert/MetaMath-Cybertron) and [berkeley-nest/Starling-LM-7B-alpha](https://huggingface.co/berkeley-nest/Starling-LM-7B-alpha) using slerp merge.

You can use ChatML format.

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [Here](https://huggingface.co/datasets/open-llm-leaderboard/results/blob/main/Q-bert/MetaMath-Cybertron-Starling/results_2023-12-07T21-59-56.458563.json)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 71.35               |
| ARC (25-shot)         | 67.75               |
| HellaSwag (10-shot)   | 86.23               |
| MMLU (5-shot)         | 65.24               |
| TruthfulQA (0-shot)   | 55.94               |
| Winogrande (5-shot)   | 81.45               |
| GSM8K (5-shot)        | 71.49               |