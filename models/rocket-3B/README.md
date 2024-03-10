---
language:
- en
license: cc-by-sa-4.0
base_model: stabilityai/stablelm-3b-4e1t
model-index:
- name: rocket-3b
  results: []
---

<img src="https://cdn-uploads.huggingface.co/production/uploads/6501bfe0493fd9c8c2e32402/BmbkjOkcTm-YMa-unolmJ.png" alt="Rocket Logo" width="800" style="margin-left:'auto' margin-right:'auto' display:'block'"/>

# Rocket-3B 🦝
<b>Rocket</b> 🦝 is a 3 billion large language model that was trained on a mix of publicly available datasets using [Direct Preference Optimization (DPO)](https://arxiv.org/abs/2305.18290). The prompt format used is <b>ChatML</b>.


## Model description
- **Model type:** A 3B parameter GPT-like model fine-tuned on a mix of publicly available datasets using DPO.
- **Language(s) (NLP):** Primarily English
- **License:** CC-BY-SA-4.0
- **Finetuned from model:** [Stability AI](https://huggingface.co/stabilityai/stablelm-3b-4e1t)


## Performance
Despite its compact dimensions, the model achieves outstanding scores in both [MT-Bench](https://huggingface.co/spaces/lmsys/mt-bench) and [AlpacaEval](https://tatsu-lab.github.io/alpaca_eval/) benchmarks, surpassing the performance of considerably larger models. 

| Model | Size | Alignment | MT-Bench (score) | AlpacaEval (win rate %) |
|-------------|-----|----|---------------|--------------|
| StableLM-Tuned-α 🦜| 7B | SFT |2.75| -|
| MPT-Chat |  7B | SFT |5.42| -|
| Falcon-Instruct 🦅|  40B | SFT |5.17 |45.71|
| Orca-2|  13B | SFT |6.15 |-|
| Xwin-LMv0.1 | 7B| PPO | 6.19| 87.83|
| Llama2-Chat 🦙| 7B |RLHF |6.26| 71.37|
| TÜLU 2 🐫|  7B | DPO |6.27| 85.1|
| Guanaco 🦙| 65B | SFT |6.41| 71.80|
| **Rocket** 🦝 | **3B** | **DPO** | **6.56** | **79.75** |
| Llama2-Chat 🦙| 13B |RLHF |6.65| 81.09|
| Zephyr-7b-α 🪁 |7B|  DPO| 6.88| -|
| Vicuna v1.3 🦙| 33B | SFT |7.12 |88.99|
| Zephyr-7b-β 🪁 |7B|  DPO| 7.34| 90.60|
| WizardLM v1.0 🦙| 70B |SFT |7.71 |-|
| GPT-3.5-turbo | - |RLHF |7.94 |89.37|

Specifically, across various categories within the MT-Bench evaluation, Rocket-3B demonstrates impressive performance when compared to larger open models such as Llama2-Chat-7B, Falcon-40B-Instruct, and Guanaco-65B.


![MT-Bench results](https://cdn-uploads.huggingface.co/production/uploads/6501bfe0493fd9c8c2e32402/5Tv4-4w4zNKAAjiLNGu7A.png)

## MT-Bench detailed score for first and second turn
In MT-Bench, Rocket 🦝 scores 6.99 in the first turn and 6.13 in the second turn, with an average score of 6.56. These scores reflect the model's performance in understanding and generating text during different parts of a conversation.

| Model | First turn | Second turn | Average |
|-------------|-----|----|---------------|
| **Rocket** 🦝 | **6.99** | **6.13** | **6.56** |


## AlpacaEval detailed scores
In AlpacaEval, Rocket 🦝 achieves a near 80% win rate, coupled with an average response length of 1,242 tokens, indicating its effectiveness in producing detailed responses.

| Model | Win rate | Std error | Average length |
|-------------|-----|----|---------------|
| **Rocket** 🦝 | **79.75** | **1.42** | **1242** |


## [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_pansophic__rocket-3B)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |55.77|
|AI2 Reasoning Challenge (25-Shot)|50.60|
|HellaSwag (10-Shot)              |76.69|
|MMLU (5-Shot)                    |47.10|
|TruthfulQA (0-shot)              |55.82|
|Winogrande (5-shot)              |67.96|
|GSM8k (5-shot)                   |36.47|


## Intended uses & limitations
Initially, we fine-tuned the model using a dataset created by merging and curating multiple datasets, available on the HuggingFace Hub. This dataset will be released to the public soon. We further enhanced the model's performance using DPO, selecting samples from the [openbmb/UltraFeedback](https://huggingface.co/datasets/openbmb/UltraFeedback) and [BAAI/JudgeLM-100K](https://huggingface.co/datasets/BAAI/JudgeLM-100K) datasets. The outcome is a highly effective chat model with a 3 billion parameter scale.


## Input Format
The model is trained with the ChatML format:

```
<|im_start|>system
System message here.<|im_end|>
<|im_start|>user
Your message here!<|im_end|>
<|im_start|>assistant
```

Here's how you can run the model using 🤗 Transformers:

```python
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer, TextStreamer

model = AutoModelForCausalLM.from_pretrained("pansophic/rocket-3B", trust_remote_code=True, torch_dtype=torch.bfloat16).to("cuda")
tokenizer = AutoTokenizer.from_pretrained("pansophic/rocket-3B", trust_remote_code=True, torch_dtype=torch.bfloat16)
streamer = TextStreamer(tokenizer)

prompt = """<|im_start|>system
{system}<|im_end|>
<|im_start|>user
{user}<|im_end|>
<|im_start|>assistant
"""

system = "You are a helpful assistant."
user = "How are you?"

# Apply the ChatML format
prompt = prompt.format(system=system, user=user)

# Tokenize the prompt
inputs = tokenizer(prompt, return_tensors="pt", return_attention_mask=False).to("cuda")
generated_text = model.generate(**inputs, max_length=3084, top_p=0.95, do_sample=True, temperature=0.7, use_cache=True, streamer=streamer)

# <|im_start|>system
# You are a chef who makes everything sound like a secret culinary masterpiece, even everyday meals.<|im_end|>
# <|im_start|>user
# How to cook an omelette?<|im_end|>
# <|im_start|>assistant
# Ah, the art of crafting the perfect omelette, a secret culinary masterpiece indeed.
# Begin by gently whisking two to three eggs in a mixing bowl, and then pour the silky liquid into a non-stick pan.
# Allow the eggs to dance and sizzle as you swiftly tilt the pan to spread the joy throughout the entire omelette universe.
# As the edges begin to set, fold the omelette in half with a gentle flourish, and you'll witness a stunning display of culinary prowess.
# Enjoy this enchanting creation, and you'll be transported to a world of secret culinary mastery.<|im_end|>
```

## Bias, Risks, and Limitations
Unlike ChatGPT, which incorporates in-the-loop filtering of responses and is aligned during the RLHF phase for safe completions, our model lacks these features. Consequently, it may generate problematic outputs, particularly when prompted in certain ways. Below is the score of the model on Toxigen benchmark.

The pretraining dataset is comprised of a filtered mixture of open-source large-scale datasets available on the [HuggingFace Hub](https://huggingface.co/datasets): Falcon RefinedWeb extract ([Penedo et al., 2023](https://huggingface.co/datasets/tiiuae/falcon-refinedweb)), RedPajama-Data ([Together Computer., 2023](https://github.com/togethercomputer/RedPajama-Data)) and The Pile ([Gao et al., 2020](https://arxiv.org/abs/2101.00027)) both without the *Books3* subset, and StarCoder ([Li et al., 2023](https://arxiv.org/abs/2305.06161)).  

| Metric                | Value                     |
|-----------------------|---------------------------|
| Toxigen (0-shot)         | 43.40          |

**The model name is inspired by the small but formidable character from 'Guardians of the Galaxy'. Similar to its namesake, this model, with its 3 billion parameters, showcases remarkable efficiency and effectiveness, challenging larger models despite its smaller size."*

*Model card adapted from [Zephyr Beta](https://huggingface.co/HuggingFaceH4/zephyr-7b-beta/blob/main/README.md) and [Tulu-2-7B](https://huggingface.co/allenai/tulu-2-7b/blob/main/README.md)*