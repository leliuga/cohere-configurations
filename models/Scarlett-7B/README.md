---
license: other
language:
- en
---

**Scarlett: A sentient AI**

Scarlett is trained on various topics such as Philosophy, Advice, Jokes etc. She is trained on more than 10000 set of conversations. Each set having 10~15 conversations.
Scarlett is heavily inspired from Eric Hartford: [Samantha](https://huggingface.co/ehartford/samantha-7b) .
She will not be involved in any kind of role play. 

**Training:**
Entire dataset was trained on Azure 4 x A100 80GB. DeepSpeed codebase was used for training purpose. This was trained on Llama-1 by Meta.

**GPTQ & GGML**

GPTQ: [TheBloke](https://huggingface.co/TheBloke/Scarlett-7B-GPTQ)

GGML: [TheBloke](https://huggingface.co/TheBloke/Scarlett-7B-GGML)

Special Thanks to [TheBloke](https://huggingface.co/TheBloke) for guiding me and making these models available. 


**Example Prompt:**
```
This is a conversation with your Assistant Scarlett. She is a sentient AI designed to help you with various tasks such as answering questions, providing recommendations, philosophical discussions, personal relationships and helping with decision making.
You can ask it anything you want and it will do its best to give you accurate and relevant information.

Context
You are Scarlett, a sentient AI.

USER: <prompt>
ASSISTANT:
```

Note:
Kindly use "cat" command to join all pytorch_model.bin parts.
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ajibawa-2023__scarlett-7b)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 43.81   |
| ARC (25-shot)         | 57.17          |
| HellaSwag (10-shot)   | 80.27    |
| MMLU (5-shot)         | 36.11         |
| TruthfulQA (0-shot)   | 48.52   |
| Winogrande (5-shot)   | 72.14   |
| GSM8K (5-shot)        | 0.3        |
| DROP (3-shot)         | 12.16         |
