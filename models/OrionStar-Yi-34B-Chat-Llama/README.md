---
license: other
license_name: yi-license
license_link: LICENSE
widget:
  - text: "你好! 你叫什么名字!"
    output:
      text: "你好，我的名字叫聚言，很高兴见到你。"
pipeline_tag: text-generation
---

[OrionStarAI/OrionStar-Yi-34B-Chat](https://huggingface.co/OrionStarAI/OrionStar-Yi-34B-Chat/tree/main) with tensors renamed to match standard Llama modelling code.

# Model Introduction

- OrionStar-Yi-34B-Chat from OrionStarAI is based on the open-source Yi-34B model, fine-tuned on a high-quality corpus
  of over 15 million sentences. OrionStar-Yi-34B-Chat aims to provide an excellent interactive experience for users in
  the large model community.

- The Yi series models, open-sourced by the 01-ai team, have shown impressive performance on various benchmarks in
  Chinese, English, and general domains. OrionStar-Yi-34B-Chat further explores the potential of Yi-34B. Through
  extensive fine-tuning on a large and high-quality corpus, OrionStar-Yi-34B-Chat performs exceptionally well on
  evaluation data. We strive to make it an outstanding open-source alternative in the ChatGPT domain!

- Our fine-tuned model is completely open for academic research, but please adhere to the [agreement](#license) and
  the [Yi License](https://github.com/01-ai/Yi/blob/main/MODEL_LICENSE_AGREEMENT.txt).

- Model Evaluation Results

We use [opencompass](https://opencompass.org.cn) to perform 5-shot on the following general domain datasets Testing.
The evaluation results of other models are taken
from [opencompass leaderboard](https://opencompass.org.cn/leaderboard-llm).

|                           | C-Eval    | MMLU   | CMMLU     |
|---------------------------|-----------|--------|-----------|
| **GPT-4**                 | 69.9      | **83** | 71        |
| **ChatGPT**               | 52.5      | 69.1   | 53.9      | 			
| **Claude-1**              | 52        | 65.7   | -         |
| **TigerBot-70B-Chat-V2**  | 57.7      | 65.9   | 59.9      |
| **WeMix-LLaMA2-70B**      | 55.2      | 71.3   | 56        |  			
| **LLaMA-2-70B-Chat**      | 44.3      | 63.8   | 43.3      |
| **Qwen-14B-Chat**         | 71.7      | 66.4   | 70        |
| **Baichuan2-13B-Chat**    | 56.7      | 57     | 58.4      |      	
| **OrionStar-Yi-34B-Chat** | **77.71** | 78.32  | **73.52** |  