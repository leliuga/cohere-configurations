---
license: apache-2.0
model-index:
- name: apricot-wildflower-20
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
      value: 59.64
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=crumb/apricot-wildflower-20
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
      value: 81.76
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=crumb/apricot-wildflower-20
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
      value: 63.38
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=crumb/apricot-wildflower-20
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
      value: 41.76
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=crumb/apricot-wildflower-20
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
      value: 77.9
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=crumb/apricot-wildflower-20
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
      value: 33.97
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=crumb/apricot-wildflower-20
      name: Open LLM Leaderboard
---

# apricot-wildflower-20

This model is the Mistral-7b model finetuned for 1k steps with a combined lm loss and distillation loss on Openwebtext2 with a >=20 reddit score filter with training logits from Mixtral. I'm not going to pretend it was a big project I did it in a dream and woke up and replicated the code without any actual reason, idk how well it fares in benchmarks.

(update: not very good)

| model | avg | arc | hellaswag | mmlu | truthfulqa | winogrande | gsm8k |
| --- | --- | --- | --- | --- | --- | --- | --- |
| apricot-wildflower-20 | 59.74 | 59.64 | 81.76 | 63.38 | 41.76 | 77.9 | 33.97 |
| mistralai/Mistral-7B-v0.1 | 60.97 | 59.98 | 83.31 | 64.16 | 42.15 | 78.37 | 37.83 |


### use
```python
from transformers import AutoModelForCausalLM, AutoTokenizer

model_id = "crumb/apricot-wildflower-20"
tokenizer = AutoTokenizer.from_pretrained(model_id)

model = AutoModelForCausalLM.from_pretrained(model_id, low_cpu_mem_usage=True, device_map="auto", load_in_8bit=True)

text = "Hello my name is"
inputs = tokenizer(text, return_tensors="pt")

outputs = model.generate(**inputs, max_new_tokens=128)
print(tokenizer.decode(outputs[0], skip_special_tokens=True))
# Hello my name is Katie and I am a 20 year old student from the UK. I am currently studying for a degree in English Literature and Creative Writing at the University of Leeds. I am a huge fan of the Harry Potter series and have been since I was 10 years old. I have read the books countless times and have seen the films many times too. I am a huge fan of the Harry Potter fandom and have been a member of the Harry Potter forums for a few years now. I am also a member of the Harry Potter fan club and have been for a few years now. I
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_crumb__apricot-wildflower-20)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |59.74|
|AI2 Reasoning Challenge (25-Shot)|59.64|
|HellaSwag (10-Shot)              |81.76|
|MMLU (5-Shot)                    |63.38|
|TruthfulQA (0-shot)              |41.76|
|Winogrande (5-shot)              |77.90|
|GSM8k (5-shot)                   |33.97|

