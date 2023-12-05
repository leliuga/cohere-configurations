---
language:
- en
tags:
- upstage
- llama-2
- instruct
- instruction
pipeline_tag: text-generation
---
# Updates
Solar, a new bot created by Upstage, is now available on **Poe**. As a top-ranked model on the HuggingFace Open LLM leaderboard, and a fine tune of Llama 2, Solar is a great example of the progress enabled by open source.
Try now at https://poe.com/Solar-0-70b


# SOLAR-0-70b-16bit model card
The model name has been changed from LLaMa-2-70b-instruct-v2 to SOLAR-0-70b-16bit

## Model Details

* **Developed by**: [Upstage](https://en.upstage.ai)
* **Backbone Model**: [LLaMA-2](https://github.com/facebookresearch/llama/tree/main)
* **Language(s)**: English
* **Library**: [HuggingFace Transformers](https://github.com/huggingface/transformers)
* **License**: Fine-tuned checkpoints is licensed under the Non-Commercial Creative Commons license ([CC BY-NC-4.0](https://creativecommons.org/licenses/by-nc/4.0/))
* **Where to send comments**: Instructions on how to provide feedback or comments on a model can be found by opening an issue in the [Hugging Face community's model repository](https://huggingface.co/upstage/Llama-2-70b-instruct-v2/discussions)
* **Contact**: For questions and comments about the model, please email [contact@upstage.ai](mailto:contact@upstage.ai)

## Dataset Details

### Used Datasets
- Orca-style dataset
- Alpaca-style dataset
- No other dataset was used except for the dataset mentioned above
- No benchmark test set or the training set are used


### Prompt Template
```
### System:
{System}

### User:
{User}

### Assistant:
{Assistant}
```

## Usage

- The followings are tested on A100 80GB
- Our model can handle up to 10k+ input tokens, thanks to the `rope_scaling` option

```python
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer, TextStreamer

tokenizer = AutoTokenizer.from_pretrained("upstage/Llama-2-70b-instruct-v2")
model = AutoModelForCausalLM.from_pretrained(
    "upstage/Llama-2-70b-instruct-v2",
    device_map="auto",
    torch_dtype=torch.float16,
    load_in_8bit=True,
    rope_scaling={"type": "dynamic", "factor": 2} # allows handling of longer inputs
)

prompt = "### User:\nThomas is healthy, but he has to go to the hospital. What could be the reasons?\n\n### Assistant:\n"
inputs = tokenizer(prompt, return_tensors="pt").to(model.device)
del inputs["token_type_ids"]
streamer = TextStreamer(tokenizer, skip_prompt=True, skip_special_tokens=True)

output = model.generate(**inputs, streamer=streamer, use_cache=True, max_new_tokens=float('inf'))
output_text = tokenizer.decode(output[0], skip_special_tokens=True)
```

## Hardware and Software

* **Hardware**: We utilized an A100x8 * 4 for training our model
* **Training Factors**: We fine-tuned this model using a combination of the [DeepSpeed library](https://github.com/microsoft/DeepSpeed) and the [HuggingFace Trainer](https://huggingface.co/docs/transformers/main_classes/trainer) / [HuggingFace Accelerate](https://huggingface.co/docs/accelerate/index)

## Evaluation Results

### Overview
- We conducted a performance evaluation following the tasks being evaluated on the [Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard).
We evaluated our model on four benchmark datasets, which include `ARC-Challenge`, `HellaSwag`, `MMLU`, and `TruthfulQA`
We used the [lm-evaluation-harness repository](https://github.com/EleutherAI/lm-evaluation-harness), specifically commit [b281b0921b636bc36ad05c0b0b0763bd6dd43463](https://github.com/EleutherAI/lm-evaluation-harness/tree/b281b0921b636bc36ad05c0b0b0763bd6dd43463).
- We used [MT-bench](https://github.com/lm-sys/FastChat/tree/main/fastchat/llm_judge), a set of challenging multi-turn open-ended questions, to evaluate the models

### Main Results
| Model | H4(Avg) | ARC | HellaSwag | MMLU | TruthfulQA | | MT_Bench |
|--------------------------------------------------------------------|----------|----------|----------|------|----------|-|-------------|
| **[Llama-2-70b-instruct-v2](https://huggingface.co/upstage/Llama-2-70b-instruct-v2)**(***Ours***, ***Open LLM Leaderboard***) | **73** | **71.1** | **87.9** | **70.6** | **62.2** | | **7.44063** |
| [Llama-2-70b-instruct](https://huggingface.co/upstage/Llama-2-70b-instruct) (Ours, Open LLM Leaderboard) | 72.3 | 70.9 | 87.5 | 69.8 | 61 | | 7.24375  |
| [llama-65b-instruct](https://huggingface.co/upstage/llama-65b-instruct) (Ours, Open LLM Leaderboard) | 69.4 | 67.6 | 86.5 | 64.9 | 58.8 | | |
| Llama-2-70b-hf | 67.3 | 67.3 | 87.3 | 69.8 | 44.9 | | |
| [llama-30b-instruct-2048](https://huggingface.co/upstage/llama-30b-instruct-2048) (Ours, Open LLM Leaderboard) | 67.0 | 64.9 | 84.9 | 61.9 | 56.3 | | |
| [llama-30b-instruct](https://huggingface.co/upstage/llama-30b-instruct) (Ours, Open LLM Leaderboard) | 65.2 | 62.5 | 86.2 | 59.4 | 52.8 | | |
| llama-65b | 64.2 | 63.5 | 86.1 | 63.9 | 43.4 | | |
| falcon-40b-instruct | 63.4 | 61.6 | 84.3 | 55.4 | 52.5 | | |

### Scripts for H4 Score Reproduction
- Prepare evaluation environments:
```
# clone the repository
git clone https://github.com/EleutherAI/lm-evaluation-harness.git
# check out the specific commit
git checkout b281b0921b636bc36ad05c0b0b0763bd6dd43463
# change to the repository directory
cd lm-evaluation-harness
```

## Contact Us

### About Upstage
- [Upstage](https://en.upstage.ai) is a company specialized in Large Language Models (LLMs) and AI. We will help you build private LLMs and related applications.
If you have a dataset to build domain specific LLMs or make LLM applications, please contact us at ► [click here to contact](https://www.upstage.ai/private-llm?utm_source=huggingface&utm_medium=link&utm_campaign=privatellm)
- As of August 1st, our 70B model has reached the top spot in openLLM rankings, marking itself as the current leading performer globally. 