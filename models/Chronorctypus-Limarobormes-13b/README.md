---
tags:
- llama
---
Five different instruction-tuned models (which I'm sure are intuitively obvious from the name) merged using the methodology described in [Resolving Interference When Merging Models](https://arxiv.org/abs/2306.01708).

In theory this should retain more of the capabilites of the constituent models than a straight linear merge would. In my testing, it feels quite capable.

Base model used for the merge: [TheBloke/Llama-2-13B-fp16](https://huggingface.co/TheBloke/Llama-2-13B-fp16)

Models merged in:
* [OpenOrca-Platypus2-13B](https://huggingface.co/Open-Orca/OpenOrca-Platypus2-13B)
* [limarp-13b-merged](https://huggingface.co/Oniichat/limarp-13b-merged)
* [Nous-Hermes-Llama2-13b](https://huggingface.co/NousResearch/Nous-Hermes-Llama2-13b)
* [chronos-13b-v2](https://huggingface.co/elinas/chronos-13b-v2)
* [airoboros-l2-13b-gpt4-1.4.1](https://huggingface.co/jondurbin/airoboros-l2-13b-gpt4-1.4.1)

Works quite well with Alpaca-style prompts:
```
### Instruction:

...

### Response:
```

The script I used to perform the merge is available [here](https://github.com/cg123/ties-merge).

The command that produced this model:
```
python ties_merge.py TheBloke/Llama-2-13B-fp16 ./Chronorctypus-Limarobormes-13b --merge elinas/chronos-13b-v2 --merge Open-Orca/OpenOrca-Platypus2-13B --merge Oniichat/limarp-13b-merged --merge jondurbin/airoboros-l2-13b-gpt4-1.4.1 --merge NousResearch/Nous-Hermes-Llama2-13b --cuda
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_chargoddard__Chronorctypus-Limarobormes-13b)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 49.88   |
| ARC (25-shot)         | 59.9          |
| HellaSwag (10-shot)   | 82.75    |
| MMLU (5-shot)         | 58.45         |
| TruthfulQA (0-shot)   | 51.9   |
| Winogrande (5-shot)   | 74.43   |
| GSM8K (5-shot)        | 3.87        |
| DROP (3-shot)         | 17.89         |
