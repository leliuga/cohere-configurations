---
license: cc-by-nc-4.0
datasets:
- meta-math/MetaMathQA
language:
- en
pipeline_tag: text-generation
tags:
- Math
- merge
base_model:
- Q-bert/MetaMath-Cybertron-Starling
- fblgit/juanako-7b-UNA
---

## Merged-AGI-7B


Merge [Q-bert/MetaMath-Cybertron-Starling](https://huggingface.co/Q-bert/MetaMath-Cybertron-Starling) and [fblgit/juanako-7b-UNA](https://huggingface.co/fblgit/juanako-7b-UNA) using slerp merge.

You can use ChatML format.

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [Coming soon]()

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | Coming soon               |
| ARC (25-shot)         | Coming soon               |
| HellaSwag (10-shot)   | Coming soon               |
| MMLU (5-shot)         | Coming soon               |
| TruthfulQA (0-shot)   | Coming soon               |
| Winogrande (5-shot)   | Coming soon               |
| GSM8K (5-shot)        | Coming soon               |