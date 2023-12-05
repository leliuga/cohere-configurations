CAMEL-13B-Combined-Data is a chat large language model obtained by finetuning LLaMA-13B model on a total of 229K conversations collected through our [CAMEL](https://arxiv.org/abs/2303.17760) framework, 100K English public conversations from ShareGPT that can be found [here](https://github.com/lm-sys/FastChat/issues/90#issuecomment-1493250773), and 52K instructions from Alpaca dataset that can be found [here](https://github.com/tatsu-lab/stanford_alpaca/blob/761dc5bfbdeeffa89b8bff5d038781a4055f796a/alpaca_data.json). We evaluate our model offline using EleutherAI's language model evaluation harness used by Huggingface's Open LLM Benchmark. CAMEL<sup>*</sup>-13B scores an average of 58.9.

| Model       | size | ARC-C  (25 shots, acc_norm) | HellaSwag  (10 shots, acc_norm) | MMLU  (5 shots, acc_norm) | TruthfulQA  (0 shot, mc2) | Average | Delta |
|-------------|:----:|:---------------------------:|:-------------------------------:|:-------------------------:|:-------------------------:|:-------:|-------|
| LLaMA       |  13B |             56.3            |               80.9              |            46.7           |            39.9           |   56.0  |   -   |
| Vicuna      |  13B |             52.8            |               80.1              |            50.5           |            51.8           |   58.8  |  2.8  |
| CAMEL<sup>*</sup>  |  13B |             56.1            |               79.9              |            50.5           |            49.0           |   58.9  |  2.9  |

---
license: cc-by-nc-4.0
---
