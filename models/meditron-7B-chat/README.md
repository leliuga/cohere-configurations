---
language:
- en
license: llama2
tags:
- Medicine
datasets:
- yahma/alpaca-cleaned
base_model: epfl-llm/meditron-7b
model-index:
- name: meditron-7b-chat
  results:
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: AI2 Reasoning Challenge (25-Shot)
      type: ai2_arc
      config: ARC-Challenge
      split: test
      args:
        num_few_shot: 25
    metrics:
    - type: acc_norm
      value: 50.77
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=malhajar/meditron-7b-chat
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: HellaSwag (10-Shot)
      type: hellaswag
      split: validation
      args:
        num_few_shot: 10
    metrics:
    - type: acc_norm
      value: 75.37
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=malhajar/meditron-7b-chat
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: MMLU (5-Shot)
      type: cais/mmlu
      config: all
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 40.49
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=malhajar/meditron-7b-chat
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: TruthfulQA (0-shot)
      type: truthful_qa
      config: multiple_choice
      split: validation
      args:
        num_few_shot: 0
    metrics:
    - type: mc2
      value: 48.56
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=malhajar/meditron-7b-chat
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: Winogrande (5-shot)
      type: winogrande
      config: winogrande_xl
      split: validation
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 73.16
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=malhajar/meditron-7b-chat
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: GSM8k (5-shot)
      type: gsm8k
      config: main
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 9.17
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=malhajar/meditron-7b-chat
      name: Open LLM Leaderboard
---
# Model Card for Model ID

<!-- Provide a quick summary of what the model is/does. -->
meditron-7b-chat is a finetuned version of [`epfl-llm/meditron-7b`](https://huggingface.co/epfl-llm/meditron-7b) using SFT Training on the Alpaca Dataset.
This model can answer information about different excplicit ideas in medicine (see [`epfl-llm/meditron-7b`](https://huggingface.co/epfl-llm/meditron-7b) for more info)

### Model Description

- **Finetuned by:** [`Mohamad Alhajar`](https://www.linkedin.com/in/muhammet-alhajar/) 
- **Language(s) (NLP):** English
- **Finetuned from model:** [`epfl-llm/meditron-7b`](https://huggingface.co/epfl-llm/meditron-7b)

### Prompt Template
```
### Instruction:

<prompt> (without the <>)

### Response:
```


## How to Get Started with the Model

Use the code sample provided in the original post to interact with the model.
```python
from transformers import AutoTokenizer,AutoModelForCausalLM
 
model_id = "malhajar/meditron-7b-chat"
model = AutoModelForCausalLM.from_pretrained(model_name_or_path,
                                             device_map="auto",
                                             torch_dtype=torch.float16,
                                             revision="main")

tokenizer = AutoTokenizer.from_pretrained(model_id)

question: "what is tract infection?"
# For generating a response
prompt = '''
### Instruction:
{question} 

### Response:'''
input_ids = tokenizer(prompt, return_tensors="pt").input_ids
output = model.generate(inputs=input_ids,max_new_tokens=512,pad_token_id=tokenizer.eos_token_id,top_k=50, do_sample=True,
        top_p=0.95)
response = tokenizer.decode(output[0])

print(response)
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_malhajar__meditron-7b-chat)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |49.59|
|AI2 Reasoning Challenge (25-Shot)|50.77|
|HellaSwag (10-Shot)              |75.37|
|MMLU (5-Shot)                    |40.49|
|TruthfulQA (0-shot)              |48.56|
|Winogrande (5-shot)              |73.16|
|GSM8k (5-shot)                   | 9.17|

