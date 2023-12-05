---
license: llama2
language:
- en
library_name: transformers
thumbnail: "https://huggingface.co/HyperbeeAI/Tulpar-7b-v0/resolve/main/tulpar.png"
---

<p align="center">
  <img src="https://huggingface.co/HyperbeeAI/Tulpar-7b-v0/resolve/main/tulpar.png" width="360" height="360" >
</p>

# Model Description
Tulpar-7b is a LLama2-7b-based model trained by HyperbeeAI. Training is done on a filtered and preprocessed instruction finetuning dataset that includes GPT-4 generated and generally curated datasets like Airoboros and Platypus. 

# Example Usage

Loading the model:
```python
from transformers import AutoTokenizer, AutoModelForCausalLM

tokenizer = AutoTokenizer.from_pretrained("HyperbeeAI/Tulpar-7b-v0")
model = AutoModelForCausalLM.from_pretrained("HyperbeeAI/Tulpar-7b-v0", device_map="auto")
```

You can run inference with both of the following prompts:
```python
input_text="What is deep learning?"
prompt = f"### User: {input_text}\n\n### Assistant:\n"
inputs = tokenizer(prompt, return_tensors="pt")
output = model.generate(**inputs, do_sample=True, top_p=0.95, top_k=0, max_new_tokens=512)
print(tokenizer.decode(output[0]))
```

```python
input_text="What is deep learning?"
prompt = f"Question: {input_text}\n\nAnswer:"
inputs = tokenizer(prompt, return_tensors="pt")
output = model.generate(**inputs, do_sample=True, top_p=0.95, top_k=0, max_new_tokens=512)
print(tokenizer.decode(output[0]))
```


# Evaluation
Our offline HF Leaderboard evaluation results:
||||
|:------:|:--------:|:-------:|
|**Task**|**Metric**|**Value**|
|*arc_challenge*|acc_norm|0.5614|
|*hellaswag*|acc_norm|0.7901|
|*mmlu*|acc_norm|0.5242|
|*truthfulqa_mc*|mc2|0.5160|
|**Average**|-|**0.5979**||

Other GPT4All evaluation results:
||||
|:------:|:--------:|:-------:|
|**Task**|**Metric**|**Value**|
|boolq|acc   |0.8306|
|piqa|acc     |0.7905|
|    |acc_norm|0.7884|
|winogrande|acc   |0.7159|
|openbookqa|acc     |0.356|
|          |acc_norm|0.448|
|**Average** (including HF leaderboard datasets) | | **0.6468** |

BigBenchHard results:
||||
|:------:|:--------:|:-------:|
|**Task**|**Metric**|**Value**|
|bigbench_causal_judgement                       |multiple_choice_grade|0.6105|
|bigbench_date_understanding                     |multiple_choice_grade|0.6423|
|bigbench_disambiguation_qa                      |multiple_choice_grade|0.3643|
|bigbench_dyck_languages                         |multiple_choice_grade|0.2000|
|bigbench_formal_fallacies_syllogisms_negation   |multiple_choice_grade|0.5002|
|bigbench_geometric_shapes                       |multiple_choice_grade|0.0000|
|                                                |exact_str_match      |0.0000|
|bigbench_hyperbaton                             |multiple_choice_grade|0.6754|
|bigbench_logical_deduction_five_objects         |multiple_choice_grade|0.2700|
|bigbench_logical_deduction_seven_objects        |multiple_choice_grade|0.1929|
|bigbench_logical_deduction_three_objects        |multiple_choice_grade|0.4133|
|bigbench_movie_recommendation                   |multiple_choice_grade|0.3000|
|bigbench_navigate                               |multiple_choice_grade|0.5000|
|bigbench_reasoning_about_colored_objects        |multiple_choice_grade|0.5750|
|bigbench_ruin_names                             |multiple_choice_grade|0.3281|
|bigbench_salient_translation_error_detection    |multiple_choice_grade|0.2976|
|bigbench_snarks                                 |multiple_choice_grade|0.6022|
|bigbench_sports_understanding                   |multiple_choice_grade|0.5122|
|bigbench_temporal_sequences                     |multiple_choice_grade|0.1450|
|bigbench_tracking_shuffled_objects_five_objects |multiple_choice_grade|0.1976|
|bigbench_tracking_shuffled_objects_seven_objects|multiple_choice_grade|0.1440|
|bigbench_tracking_shuffled_objects_three_objects|multiple_choice_grade|0.4133|
|**Average**| |**0.3754**

# Ethical Considerations and Limitations
Tulpar is a technology with potential risks and limitations. This model is finetuned only in English and all language-related scenarios are not covered. As HyperbeeAI, we neither guarantee ethical, accurate, unbiased, objective responses nor endorse its outputs. Before deploying this model, you are advised to make safety tests for your use case.