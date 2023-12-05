---
 
language:
- en
tags:
- upstage
- llama
- instruct
- instruction
pipeline_tag: text-generation
---
# LLaMa-65b-instruct model card

## Model Details

* **Developed by**: [Upstage](https://en.upstage.ai)
* **Backbone Model**: [LLaMA](https://github.com/facebookresearch/llama/tree/llama_v1)
* **Variations**: It has different model parameter sizes and sequence lengths: [30B/1024](https://huggingface.co/upstage/llama-30b-instruct), [30B/2048](https://huggingface.co/upstage/llama-30b-instruct-2048), [65B/1024](https://huggingface.co/upstage/llama-65b-instruct)
* **Language(s)**: English
* **Library**: [HuggingFace Transformers](https://github.com/huggingface/transformers)
* **License**: This model is under a **Non-commercial** Bespoke License and governed by the Meta license. You should only use this repository if you have been granted access to the model by filling out [this form](https://docs.google.com/forms/d/e/1FAIpQLSfqNECQnMkycAp2jP4Z9TFX0cGR4uf7b_fBxjY_OjhJILlKGA/viewform), but have either lost your copy of the weights or encountered issues converting them to the Transformers format
* **Where to send comments**: Instructions on how to provide feedback or comments on a model can be found by opening an issue in the [Hugging Face community's model repository](https://huggingface.co/upstage/llama-30b-instruct-2048/discussions)
* **Contact**: For questions and comments about the model, please email [contact@upstage.ai](mailto:contact@upstage.ai)

## Dataset Details

### Used Datasets

- Orca-style dataset
- No other data was used except for the dataset mentioned above

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

- Tested on A100 80GB
- Our model can handle up to 10k+ input tokens, thanks to the `rope_scaling` option

```python
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer, TextStreamer

tokenizer = AutoTokenizer.from_pretrained("upstage/llama-65b-instruct")
model = AutoModelForCausalLM.from_pretrained(
    "upstage/llama-65b-instruct",
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
* **Training Factors**: We fine-tuned this model using a combination of the [DeepSpeed library](https://github.com/microsoft/DeepSpeed) and the [HuggingFace Trainer](https://huggingface.co/docs/transformers/main_classes/trainer)

## Evaluation Results

### Overview
- We conducted a performance evaluation based on the tasks being evaluated on the [Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard).
We evaluated our model on four benchmark datasets, which include `ARC-Challenge`, `HellaSwag`, `MMLU`, and `TruthfulQA`.
We used the [lm-evaluation-harness repository](https://github.com/EleutherAI/lm-evaluation-harness), specifically commit [b281b0921b636bc36ad05c0b0b0763bd6dd43463](https://github.com/EleutherAI/lm-evaluation-harness/tree/b281b0921b636bc36ad05c0b0b0763bd6dd43463)
- We used [MT-bench](https://github.com/lm-sys/FastChat/tree/main/fastchat/llm_judge), a set of challenging multi-turn open-ended questions, to evaluate the models

### Main Results
| Model | H4(Avg) | ARC | HellaSwag | MMLU | TruthfulQA | | MT_Bench |
|--------------------------------------------------------------------|----------|----------|----------|------|----------|-|-------------|
| **[Llama-2-70b-instruct-v2](https://huggingface.co/upstage/Llama-2-70b-instruct-v2)**(Ours, Open LLM Leaderboard) | **73** | **71.1** | **87.9** | **70.6** | **62.2** | | **7.44063** |
| [Llama-2-70b-instruct](https://huggingface.co/upstage/Llama-2-70b-instruct) (Ours, Open LLM Leaderboard) | 72.3 | 70.9 | 87.5 | 69.8 | 61 | | 7.24375  |
| [llama-65b-instruct](https://huggingface.co/upstage/llama-65b-instruct) (***Ours***, ***Open LLM Leaderboard***) | 69.4 | 67.6 | 86.5 | 64.9 | 58.8 | | |
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

## Ethical Issues

### Ethical Considerations
- There were no ethical issues involved, as we did not include the benchmark test set or the training set in the model's training process

## Contact Us

### Why Upstage LLM?
- [Upstage](https://en.upstage.ai)'s LLM research has yielded remarkable results. As of August 1st, our 70B model has reached the top spot in openLLM rankings, marking itself as the current leading performer globally. Recognizing the immense potential in implementing private LLM to actual businesses, we invite you to easily apply private LLM and fine-tune it with your own data. For a seamless and tailored solution, please do not hesitate to reach out to us. ► [click here to contact](https://www.upstage.ai/private-llm?utm_source=huggingface&utm_medium=link&utm_campaign=privatellm)