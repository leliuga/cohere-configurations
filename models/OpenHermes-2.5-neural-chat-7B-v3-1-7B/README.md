---
license: apache-2.0
datasets:
- Open-Orca/SlimOrca
tags:
- mistral
---
![image/png](https://cdn-uploads.huggingface.co/production/uploads/6468ce47e134d050a58aa89c/x44nNbPTpv0zGTqA1Jb2q.png)

Merge of [teknium/OpenHermes-2.5-Mistral-7B](https://huggingface.co/teknium/OpenHermes-2.5-Mistral-7B) and [Intel/neural-chat-7b-v3-1](https://huggingface.co/Intel/neural-chat-7b-v3-1) using ties merge.

### *Weights*

- [teknium/OpenHermes-2.5-Mistral-7B](https://huggingface.co/teknium/OpenHermes-2.5-Mistral-7B): 0.5

- [Intel/neural-chat-7b-v3-1](https://huggingface.co/Intel/neural-chat-7b-v3-1): 0.3

### *Density*

- [teknium/OpenHermes-2.5-Mistral-7B](https://huggingface.co/teknium/OpenHermes-2.5-Mistral-7B): 0.5

- [Intel/neural-chat-7b-v3-1](https://huggingface.co/Intel/neural-chat-7b-v3-1): 0.5

# Prompt Templates

You can use these prompt templates, but I recommend using ChatML.

### ChatML [(OpenHermes-2.5-Mistral-7B)](https://huggingface.co/teknium/OpenHermes-2.5-Mistral-7B):

```
<|im_start|>system
{system}<|im_end|>
<|im_start|>user
{user}<|im_end|>
<|im_start|>assistant
{asistant}<|im_end|>
```

### [neural-chat-7b-v3-1](https://huggingface.co/Intel/neural-chat-7b-v3-1):

```
### System:
{system}
### User:
{usr}
### Assistant:
```

# Quantizationed versions

Quantizationed versions of this model is available thanks to [TheBloke](https://hf.co/TheBloke).

##### GPTQ

- [TheBloke/OpenHermes-2.5-neural-chat-7B-v3-1-7B-GPTQ](https://huggingface.co/TheBloke/OpenHermes-2.5-neural-chat-7B-v3-1-7B-GPTQ)

##### GGUF

- [TheBloke/OpenHermes-2.5-neural-chat-7B-v3-1-7B-GGUF](https://huggingface.co/TheBloke/OpenHermes-2.5-neural-chat-7B-v3-1-7B-GGUF)

##### AWQ

- [TheBloke/OpenHermes-2.5-neural-chat-7B-v3-1-7B-AWQ](https://huggingface.co/TheBloke/OpenHermes-2.5-neural-chat-7B-v3-1-7B-AWQ)


# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Weyaxi__OpenHermes-2.5-neural-chat-7b-v3-1-7B)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 67.84   |
| ARC (25-shot)         | 66.55          |
| HellaSwag (10-shot)   | 84.47    |
| MMLU (5-shot)         | 63.34         |
| TruthfulQA (0-shot)   | 61.22   |
| Winogrande (5-shot)   | 78.37   |
| GSM8K (5-shot)        | 53.07        |

If you would like to support me:

[☕ Buy Me a Coffee](https://www.buymeacoffee.com/weyaxi)