---
license: cc-by-nc-4.0
tags:
- not-for-all-audiences
- nsfw
---

First :
```shell
layer_slices:
  - model: Undi95/MLewd-L2-Chat-13B
    start: 0
    end: 16
  - model: Undi95/MLewd-ReMM-L2-Chat-20B-Part1
    start: 8
    end: 20
  - model: Undi95/MLewd-L2-Chat-13B
    start: 17
    end: 32
  - model: Undi95/MLewd-ReMM-L2-Chat-20B-Part1
    start: 21
    end: 40
```

Inverted:
```shell
layer_slices:
  - model: Undi95/MLewd-ReMM-L2-Chat-20B-Part1
    start: 0
    end: 16
  - model: Undi95/MLewd-L2-Chat-13B
    start: 8
    end: 20
  - model: Undi95/MLewd-ReMM-L2-Chat-20B-Part1
    start: 17
    end: 32
  - model: Undi95/MLewd-L2-Chat-13B
    start: 21
    end: 40
```

Precise:
```shell
layer_slices:
  - model: Undi95/MLewd-L2-Chat-13B
    start: 0
    end: 8
  - model: Undi95/MLewd-ReMM-L2-Chat-20B-Part1
    start: 4
    end: 12
  - model: Undi95/MLewd-L2-Chat-13B
    start: 9
    end: 16
  - model: Undi95/MLewd-ReMM-L2-Chat-20B-Part1
    start: 13
    end: 22
  - model: Undi95/MLewd-L2-Chat-13B
    start: 17
    end: 24
  - model: Undi95/MLewd-ReMM-L2-Chat-20B-Part1
    start: 23
    end: 32
  - model: Undi95/MLewd-L2-Chat-13B
    start: 25
    end: 32
  - model: Undi95/MLewd-ReMM-L2-Chat-20B-Part1
    start: 33
    end: 40
```

PreciseInverted:
```shell
layer_slices:
  - model: Undi95/MLewd-ReMM-L2-Chat-20B-Part1
    start: 0
    end: 8
  - model: Undi95/MLewd-L2-Chat-13B
    start: 4
    end: 12
  - model: Undi95/MLewd-ReMM-L2-Chat-20B-Part1
    start: 9
    end: 16
  - model: Undi95/MLewd-L2-Chat-13B
    start: 13
    end: 22
  - model: Undi95/MLewd-ReMM-L2-Chat-20B-Part1
    start: 17
    end: 24
  - model: Undi95/MLewd-L2-Chat-13B
    start: 23
    end: 32
  - model: Undi95/MLewd-ReMM-L2-Chat-20B-Part1
    start: 25
    end: 32
  - model: Undi95/MLewd-L2-Chat-13B
    start: 33
    end: 40
```

Part1 = ReMM v2.1 merged /w MLewd low weight to keep consistency. I call this "dilution" and result show consistency and coherency without repeat/loop beside the small amount of duplicated datas.

The goal is to find the best way to interlace layers the best way possible to have a sweetspot between 13B and +30B.

Normal/Inverted is by chunk of 16 layers and Precise/PreciseInverted is by chunk of 8 layers.

All the models are made of 64(+1) layers. Need testing.

## Prompt template: Alpaca

```
Below is an instruction that describes a task. Write a response that completes the request.

### Instruction:
{prompt}

### Response:
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Undi95__MLewd-ReMM-L2-Chat-20B)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 53.33   |
| ARC (25-shot)         | 62.46          |
| HellaSwag (10-shot)   | 85.62    |
| MMLU (5-shot)         | 59.13         |
| TruthfulQA (0-shot)   | 55.63   |
| Winogrande (5-shot)   | 77.19   |
| GSM8K (5-shot)        | 10.92        |
| DROP (3-shot)         | 22.33         |
