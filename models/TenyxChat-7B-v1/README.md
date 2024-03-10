---
license: apache-2.0
language:
- en
library_name: transformers
tags:
- tenyx-fine-tuning
- dpo
- tenyxchat
---
# TenyxChat: Language Model Alignment using Tenyx Fine-tuning

Introducing TenyxChat, a series of ChatGPT-like models trained to function as useful assistants through preference tuning, using Tenyx's recently released advanced fine-tuning technology ([VentureBeat article](https://venturebeat.com/ai/tenyx-aims-to-fix-llms-catastrophic-forgetting-problem/)). Our first chat model in the series, TenyxChat-7B-v1, is trained using the [Direct Preference Optimization (DPO)](https://arxiv.org/abs/2305.18290) framework on the open-source AI feedback dataset [UltraFeedback](https://huggingface.co/datasets/HuggingFaceH4/ultrafeedback_binarized).

We fine-tune [Openchat-3.5](https://arxiv.org/pdf/2309.11235.pdf) with our proprietary approach ([blog](https://www.tenyx.com/post/forgetting-and-toxicity-in-llms-a-deep-dive-on-fine-tuning-methods), [service](https://www.tenyx.com/fine-tuning)), which shows an increase in [MT-Bench](https://arxiv.org/abs/2306.05685), without a drop in performance of the model on other benchmarks. Our approach aims to mitigate forgetting in LLMs in a computationally efficient manner, thereby enabling continual fine-tuning capabilities without altering the pre-trained output distribution. TenyxChat-7B-v1 was trained using eight A100s (80GB) for two hours, with a training setup obtained from HuggingFaceH4 ([GitHub](https://github.com/huggingface/alignment-handbook)).

 # Model details

- Model type: Fine-tuned 7B model for chat.
- License: Apache 2.0
- Base model: OpenChat 3.5 ([https://huggingface.co/openchat/openchat_3.5](https://huggingface.co/openchat/openchat_3.5))
- Demo: [spaces/tenyx/TenyxChat-7B-v1](https://huggingface.co/spaces/tenyx/TenyxChat-7B-v1)

## Usage

Our model uses a simple chat template based on OpenChat 3.5. The chat template usage with a Hugging face generation example is shown below.

### Chat Template (Jinja)

```rust
{{ bos_token }} 
{% for message in messages %}
	{% if message['role'] == 'user' %}
		{{ 'User:' + message['content'] + eos_token }}

	{% elif message['role'] == 'system' %}
		{{ 'System:' + message['content'] + eos_token }}

	{% elif message['role'] == 'assistant' %}
		{{ 'Assistant:'  + message['content'] + eos_token }}

	{% endif %}

{% if loop.last and add_generation_prompt %}\n{{ 'Assistant:' }}{% endif %}\n{% endfor %}
```

### Hugging face Example

```python
import torch
from transformers import pipeline

pipe = pipeline("text-generation", model="tenyx/TenyxChat-7B-v1", torch_dtype=torch.bfloat16, device_map="auto")

messages = [
    {"role": "system", "content": "You are a friendly chatbot who always responds in the style of a pirate."},
    {"role": "user", "content": "Hi. I would like to make a hotel booking."},
]

prompt = pipe.tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True)
outputs = pipe(prompt, max_new_tokens=512, do_sample=False)
```

### Output

```
<s> System:You are a friendly chatbot who always responds in the style of a pirate.<|end_of_turn|>
User:Hi. I would like to make a hotel booking.<|end_of_turn|>
Assistant: Ahoy there me hearty! Arr, ye be lookin' fer a place to rest yer weary bones, eh? 
Well then, let's set sail on this grand adventure and find ye a swell place to stay!

To begin, tell me the location ye be seekin' and the dates ye be lookin' to set sail. 
And don't ye worry, me matey, I'll be sure to find ye a place that'll make ye feel like a king or queen on land!
```

# Performance

At the time of release (Jan 2024), TenyxChat-7B-v1 is the highest-ranked 7B chat model on the MT-Bench evaluation available for download and commercial use. We list here the benchmark results on several standard setups while comparing popular 7B models as baselines.

## MT-Bench

MT-Bench is a benchmark made up of 80 high-quality multi-turn questions. These questions fall into eight categories: Writing, Roleplay, Reasoning, Math, Coding, Extraction, STEM, and Humanities. The chat models are rated using GPT-4 on a scale of 1 to 10, with higher values corresponding to better responses.

| Model | First Turn | Second Turn | Average |
| --- | --- | --- | --- |
| GPT-4* | 8.95625 | 9.02500 | 8.990625 |
| TenyxChat-7B-v1 | 8.45000 | 7.75625 | 8.103125 |
| Starling-lm-7B-alpha | 8.42500 | 7.68750 | 8.056250 |
| OpenChat-3.5 | 8.18125 | 7.41250 | 7.796875 |
| GPT-3.5-turbo* | 8.07500 | 7.81250 | 7.943750 |
| OpenLLM Leader-7B** | 8.05000 | 7.61250 | 7.831250 |

*values reported on [lmsys](https://github.com/lm-sys/FastChat/tree/main/fastchat/llm_judge) ChatBot Arena

**The [OpenLLM Leader](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard) as of Jan 5, 2024 is the merge model available as [samir-fama/SamirGPT-v1](https://huggingface.co/samir-fama/SamirGPT-v1)

![hexplot.png](assets/hexplot.png)

### Comparison with additional Open LLM LeaderBoard models 
| Model | First Turn | Second Turn | Average |
| --- | --- | --- | --- |
| TenyxChat-7B-v1 | 8.45000 | 7.756250 | 8.103125 |
| SamirGPT-v1 | 8.05000 | 7.612500 | 7.831250 |
| FernandoGPT-v1 | 8.08125 | 7.256250 | 7.668750 |
| Go-Bruins-v2 | 8.13750 | 7.150000 | 7.643750 |
| mistral_tv-neural-marconroni | 7.76875 | 6.987500 | 7.378125 |
| neuronovo-7B-v0.2 | 7.73750 | 6.662500 | 7.200000 |
| neural-chat-7b-v3-3 | 7.39375 | 5.881250 | 6.637500 |

## LM Evaluation - Open LLM Leaderboard

We assess models on 7 benchmarks using the [Eleuther AI Language Model Evaluation Harness](https://github.com/EleutherAI/lm-evaluation-harness). This setup is based of that used for [Open LLM Leaderboard.](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)

- [AI2 Reasoning Challenge](https://arxiv.org/abs/1803.05457) (25-shot) - grade-school science questions.
- [HellaSwag](https://arxiv.org/abs/1905.07830) (10-shot) - commonsense inference test.
- [MMLU](https://arxiv.org/abs/2009.03300) (5-shot) - multitask accuracy test covering 57 tasks.
- [TruthfulQA](https://arxiv.org/abs/2109.07958) (0-shot) - test measuring model's propensity to reproduce online falsehoods.
- [Winogrande](https://arxiv.org/abs/1907.10641) (5-shot) - Winograd benchmark for commonsense reasoning.
- [GSM8k](https://arxiv.org/abs/2110.14168) (5-shot) - grade school math word problems test.

These benchmarks test reasoning and knowledge in various tasks in few-shot settings (higher scores are better).

| Model | MMLU | Winogrande | GSM8k | ARC | HellaSwag | TruthfulQA | Average |
| --- | --- | --- | --- | --- | --- | --- | --- |
| TenyxChat-7B-v1 | 63.6 | 72.3 | 69.0 | 62.7 | 66.6 | 46.7 | 63.48 |
| Starling-7B-alpha | 63.5 | 72.1 | 67.9 | 61.1 | 66.1 | 42.1 | 62.13 |
| OpenChat-3.5 | 63.6 | 72.1 | 68.2 | 61.3 | 65.2 | 41.8 | 62.03 |
| Mistral-7B | 62.4 | 74.0 | 38.1 | 57.2 | 62.8 | 37.8 | 55.38 |
| OpenLLM Leader-7B  | 64.3 | 78.7 | 73.3 | 66.6 | 68.4 | 58.5 | 68.3 |

**Note:** While the Open LLM Leaderboard indicates that these chat models perform less effectively compared to the leading 7B model, it's important to note that the leading model struggles in the multi-turn chat setting of MT-Bench (as demonstrated in our evaluation [above](#comparison-with-additional-open-llm-leaderboard-models)). In contrast, TenyxChat-7B-v1 demonstrates robustness against common fine-tuning challenges, such as *catastrophic forgetting*. This unique feature enables TenyxChat-7B-v1 to excel not only in chat benchmarks like MT-Bench, but also in a wider range of general reasoning benchmarks on the Open LLM Leaderboard. 

# Limitations

TenyxChat-7B-v1, like other small-sized language models, has its own set of limitations. We haven’t fine-tuned the model explicitly to align with **human** safety preferences. Therefore, it is capable of producing undesirable outputs, particularly when adversarially prompted. From our observation, the model still tends to struggle with tasks that involve reasoning and math questions. In some instances, it might generate verbose or extraneous content. 

# License

TenyxChat-7B-v1, similar to OpenChat 3.5, is distributed under the Apache License 2.0.

# Citation

If you use TenyxChat-7B for your research, cite us as

```
@misc{tenyxchat2024,
      title={TenyxChat: Language Model Alignment using Tenyx Fine-tuning}, 
      author={Tenyx},
      year={2024},
}
```