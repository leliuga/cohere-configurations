---
license: cc-by-nc-4.0
---

Merge:
```shell
layer_slices:
  - model: ./MXLewd-L2-20B-part2
    start: 0
    end: 16
  - model: ./MXLewd-L2-20B-part1
    start: 8
    end: 20
  - model: ./MXLewd-L2-20B-part2
    start: 17
    end: 32
  - model: ./MXLewd-L2-20B-part1
    start: 21
    end: 40
```
Part 2 is ReMM (0.33) and Xwin (0.66)

Part 1 is Xwin (0.33) and MLewd (0.66)
<!-- description start -->
## Models used

- Undi95/MLewd-L2-13B-v2-3
- Undi95/ReMM-v2.1-L2-13B
- Xwin-LM/Xwin-LM-13B-V0.1
<!-- description end -->

## Prompt template: Alpaca

```
Below is an instruction that describes a task. Write a response that completes the request.

### Instruction:
{prompt}

### Response:
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Undi95__MXLewd-L2-20B)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 51.29   |
| ARC (25-shot)         | 63.23          |
| HellaSwag (10-shot)   | 85.33    |
| MMLU (5-shot)         | 57.36         |
| TruthfulQA (0-shot)   | 51.65   |
| Winogrande (5-shot)   | 76.09   |
| GSM8K (5-shot)        | 10.92        |
| DROP (3-shot)         | 14.46         |
