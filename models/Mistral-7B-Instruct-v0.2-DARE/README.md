---
license: apache-2.0
language:
- en
---
<!-- header start -->
<!-- 200823 -->
<div style="width: auto; margin-left: auto; margin-right: auto">
<img src="https://github.com/janhq/jan/assets/89722390/35daac7d-b895-487c-a6ac-6663daaad78e" alt="Jan banner" style="width: 100%; min-width: 400px; display: block; margin: auto;">
</div>

<p align="center">
    <a
 href="https://jan.ai/">Jan</a> 
    - <a href="https://discord.gg/AsJ8krTT3N">Discord</a>
</p>
<!-- header end -->

# Model Description
This model uses the `DARE` method to merge [Mistral-7B-Instruct-v0.2](https://huggingface.co/mistralai/Mistral-7B-Instruct-v0.2) with 3 leading models in 12th Dec on [OpenLLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard):
1. [OpenHermes-2.5-neural-chat-v3-3-Slerp](https://huggingface.co/Weyaxi/OpenHermes-2.5-neural-chat-v3-3-Slerp)
2. [MetaMath-Cybertron-Starling](https://huggingface.co/Q-bert/MetaMath-Cybertron-Starling)
3. [v1olet_marcoroni-go-bruins-merge-7B](https://huggingface.co/v1olet/v1olet_marcoroni-go-bruins-merge-7B)

- base model: [Mistral-7B-Instruct-v0.2](https://huggingface.co/mistralai/Mistral-7B-Instruct-v0.2)

The yaml config file for this model is here:

```yaml
base_model: mistralai/Mistral-7B-Instruct-v0.2
dtype: bfloat16
merge_method: dare_ties
models:
- model: mistralai/Mistral-7B-Instruct-v0.2
- model: Weyaxi/OpenHermes-2.5-neural-chat-v3-3-Slerp
  parameters:
    density: 0.8
    weight: 0.4
- model: Q-bert/MetaMath-Cybertron-Starling
  parameters:
    density: 0.8
    weight: 0.3
- model: v1olet/v1olet_marcoroni-go-bruins-merge-7B
  parameters:
    density: 0.8
    weight: 0.3
parameters:
  int8_mask: true
```

# Prompt template:

- **ChatML**

```
<|im_start|>system
{system_message}<|im_end|>
<|im_start|>user
{prompt}<|im_end|>
<|im_start|>assistant
```
- **Alpaca**
```
{system_message}

### Instruction:
{prompt}

### Response:
```

# Run this model
You can run this model using [Jan Desktop](https://jan.ai/) on Mac, Windows, or Linux.

Jan is an open source, ChatGPT alternative that is:

- 💻  **100% offline on your machine**: Your conversations remain confidential, and visible only to you.
- 🗂️ **An Open File Format**: Conversations and model settings stay on your computer and can be exported or deleted at any time.
- 🌐 **OpenAI Compatible**: Local server on port `1337` with OpenAI compatible endpoints
- 🌍 **Open Source & Free**: We build in public; check out our [Github](https://github.com/janhq)

![image/png](https://cdn-uploads.huggingface.co/production/uploads/65713d70f56f9538679e5a56/r7VmEBLGXpPLTu2MImM7S.png)

# About Jan
Jan believes in the need for an open-source AI ecosystem and is building the infra and tooling to allow open-source AIs to compete on a level playing field with proprietary ones.

Jan's long-term vision is to build a cognitive framework for future robots, who are practical, useful assistants for humans and businesses in everyday life.

# Jan Model Merger
This is a test project for merging models.

# Open LLM Leaderboard Evaluation Results

Detailed results can be found here.

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | ?|
| ARC (25-shot)         | ?          |
| HellaSwag (10-shot)   | ?   |
| MMLU (5-shot)         | ?|
| TruthfulQA (0-shot)   | ? |
| Winogrande (5-shot)   | ?  |
| GSM8K (5-shot)        | ?        |

# Acknowlegement
- [mergekit](https://github.com/cg123/mergekit)
- [DARE](https://github.com/yule-BUAA/MergeLM/blob/main/README.md)
- [SLERP](https://github.com/Digitous/LLM-SLERP-Merge)
- [lm-evaluation-harness](https://github.com/EleutherAI/lm-evaluation-harness)