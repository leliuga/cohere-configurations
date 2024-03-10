---
language: en
license: apache-2.0
library_name: transformers
model-index:
- name: digital-socrates-7b
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
      value: 54.44
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=allenai/digital-socrates-7b
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
      value: 75.99
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=allenai/digital-socrates-7b
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
      value: 51.41
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=allenai/digital-socrates-7b
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
      value: 44.88
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=allenai/digital-socrates-7b
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
      value: 73.09
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=allenai/digital-socrates-7b
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
      value: 17.89
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=allenai/digital-socrates-7b
      name: Open LLM Leaderboard
---

This is the Digital Socrates 7B (DS-7B) model described in our paper: <b>Digital Socrates: Evaluating LLMs through explanation critiques</b> (arXiv link: https://arxiv.org/abs/2311.09613). 

The recommended, better performing 13B model can be found at https://huggingface.co/allenai/digital-socrates-13b

The DS-7B model is a fine-tuned version of [Llama-2-7b-Chat](https://huggingface.co/meta-llama/Llama-2-7b-chat), please 
review its [guidelines](https://huggingface.co/meta-llama/Llama-2-7b-chat/blob/main/USE_POLICY.md) and 
[licensing](https://huggingface.co/meta-llama/Llama-2-7b-chat#model-details) before use.
  
  
  
# What is Digital Socrates?
Digital Socrates is an open-source, automatic explanation-critiquing model. It is useful for revealing insights about student 
models by examining their reasoning chains. The use of Digital Socrates allows for nuanced, interpretable automatic 
evaluation of explanations without expensive API calls or human annotations.

While LLMs can provide reasoned explanations along with their answers, the nature and
quality of those explanations are still poorly understood. In response, our goal is to define a detailed way of characterizing the 
explanation capabilities of modern models and to create a nuanced, interpretable explanation
evaluation tool that can generate such characterizations automatically.

Given a **question** (together with the multiple-choice options and gold answer),
```commandline
When Dennis operates his lawnmower, he notices the engine makes a lot of noise. He also notices that the engine gets very hot. Which best describes the heat and noise generated from the lawnmower? 
(A) a change in phase 
(B) thermal expansion 
(C) an increase in entropy 
(D) mechanical advantage 

Answer sheet: (C)
```
 as well as a student model explanation and answer,
```
1) The question states that the lawnmower engine makes a lot of noise.
2) The question states that the lawnmower engine gets very hot.
3) Noise and heat are both forms of energy.
4) The noise and heat generated from the lawnmower are a result of the conversion of energy from the fuel to mechanical energy.

Answer: (D)
```

Digital Socrates gives a **critique of the model-generated explanation** that provides localized feedback on the 
most significant flaw (if any) in the explanation, suggestions to address the identified flaw, as well as a numeric rating indicating the explanation quality:
```
The explanation states or suggests the following:
* Main flaw (standalone statement): "The noise and heat generated from the lawnmower are a result of the conversion of energy from the fuel to mechanical energy."
* Dimension: incorrect_information

Consider these points for revising the explanation:
* General: Remember that noise and heat are not forms of energy. They are byproducts of energy conversion.
* Specific: In this case, the noise and heat generated by the lawnmower are not a result of the conversion of energy from the fuel to mechanical energy. They are byproducts of the combustion process.

Explanation score: 2
```


Remarkably, despite being orders of magnitude smaller than GPT-4, our Digital Socrates models are 
capable of generating critiques close to GPT-4 critiques in terms of human rating and other 
quantitative measures (correlation of explanation scores given and error category matches). 
Through quantitative and qualitative analysis, we demonstrate how Digital Socrates is useful for 
revealing insights about student models by examining their reasoning chains.

We invite you to try out Digital Socrates for your own application!
    
    
    
# How to use Digital Socrates?
We provide a quick example of how you can try out Digital Socrates with just a few lines of code:

'DSCritiqueBank-V1' used below can be downloaded from our [dataset page](https://allenai.org/data/digital-socrates).
```
import json
from transformers import AutoTokenizer, AutoModelForCausalLM
# Load model and tokenizer
model_path = "allenai/digital-socrates-7b"
model = AutoModelForCausalLM.from_pretrained(model_path).to("cuda:0")
tokenizer = AutoTokenizer.from_pretrained(model_path)

# Define input data
question = "When Dennis operates his lawnmower, he notices the engine makes a lot of noise. He also notices that the engine gets very hot. Which best describes the heat and noise generated from the lawnmower? (A) a change in phase (B) thermal expansion (C) an increase in entropy (D) mechanical advantage"
explanation = "1) The question states that the lawnmower engine makes a lot of noise.\n2) The question states that the lawnmower engine gets very hot.\n3) Noise and heat are both forms of energy.\n4) The noise and heat generated from the lawnmower are a result of the conversion of energy from the fuel to mechanical energy."
answerkey = "C"
predictedanswer = "D"

# construct prompt (Llama conventions)
with open("../DSCritiqueBank-V1/DSCB-prompts.json") as file:
    prompts = json.load(file)

system_prompt = prompts['digital_socrates_v1']['system']
user_prompt = prompts['digital_socrates_v1']['main'].replace("[[QUESTION]]", question).replace("[[EXPLANATION]]", explanation).replace("[[PREDICTEDANSWER]]", predictedanswer).replace("[[ANSWERKEY]]", answerkey)

full_prompt = f"[INST] <<SYS>>\n{system_prompt}\n<</SYS>{user_prompt} [/INST]\n\n"

# Run model
input_ids = tokenizer.encode(full_prompt, return_tensors="pt").to("cuda:0")
output = model.generate(input_ids, max_new_tokens=512, temperature=0)
res = tokenizer.batch_decode(output, skip_special_tokens=True)
```
Print the output:
```
>>> print(res[0].split("[/INST]")[-1])

The explanation states or suggests the following:
* Main flaw (standalone statement): "The noise and heat generated from the lawnmower are a result of the conversion of energy from the fuel to mechanical energy."
* Dimension: incorrect_information

Consider these points for revising the explanation:
* General: Remember that noise and heat are not forms of energy. They are byproducts of energy conversion.
* Specific: In this case, the noise and heat generated by the lawnmower are not a result of the conversion of energy from the fuel to mechanical energy. They are byproducts of the combustion process.

Explanation score: 2
```


    
# More details about Digital Socrates ...
For more details about Digital Socrates, please refer to our:
* 📄Paper: https://arxiv.org/abs/2311.09613
* 💻Dataset: https://allenai.org/data/digital-socrates
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_allenai__digital-socrates-7b)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |52.95|
|AI2 Reasoning Challenge (25-Shot)|54.44|
|HellaSwag (10-Shot)              |75.99|
|MMLU (5-Shot)                    |51.41|
|TruthfulQA (0-shot)              |44.88|
|Winogrande (5-shot)              |73.09|
|GSM8k (5-shot)                   |17.89|

