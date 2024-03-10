---
license: mit
datasets:
- nbertagnolli/counsel-chat
model-index:
- name: MelloGPT
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
      value: 53.84
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=steve-cse/MelloGPT
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
      value: 76.12
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=steve-cse/MelloGPT
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
      value: 55.99
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=steve-cse/MelloGPT
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
      value: 55.61
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=steve-cse/MelloGPT
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
      value: 73.88
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=steve-cse/MelloGPT
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
      value: 30.1
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=steve-cse/MelloGPT
      name: Open LLM Leaderboard
---
# MelloGPT
 <p align="center">
   <img width="150" height="150" src="https://raw.githubusercontent.com/steve-cse/mello-react/master/public/pwa-512x512.png" alt="Logo">
  </p>

**NOTE: This model should not be regarded as a replacement for professional mental health assistance. It is essential to seek support from qualified professionals for personalized and appropriate care.**

A fine tuned version of [Mistral-7B-Instruct-v0.1](https://huggingface.co/mistralai/Mistral-7B-Instruct-v0.1) on [counsel-chat](https://huggingface.co/datasets/nbertagnolli/counsel-chat) dataset for mental health counseling conversations.
## Motivation
In an era where mental health support is of paramount importance, A large language model fine-tuned on mental health counseling conversations stands as a pioneering solution. Leveraging a diverse dataset of anonymized counseling sessions, the model has been trained to recognize and respond to a wide range of mental health concerns. The fine-tuning process incorporates ethical considerations, privacy concerns, and sensitivity to the nuances of mental health conversations. The resulting model will demonstrate an intricate understanding of mental health issues and provide empathetic and supportive responses.
## Prompt Template
```
<s>[INST] {prompt} [/INST]
```
## Quantized Model 
The quantized model can be found [here](https://huggingface.co/models?other=base_model:steve-cse/MelloGPT). Thanks to [@TheBloke](https://huggingface.co/TheBloke).

## [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_steve-cse__MelloGPT)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |57.59|
|AI2 Reasoning Challenge (25-Shot)|53.84|
|HellaSwag (10-Shot)              |76.12|
|MMLU (5-Shot)                    |55.99|
|TruthfulQA (0-shot)              |55.61|
|Winogrande (5-shot)              |73.88|
|GSM8k (5-shot)                   |30.10|

## Contributions
This project is open for contributions. Feel free to use the community tab.

## Inspiration
This project was inspired by the project(s) listed below:

[companion_cube](https://huggingface.co/KnutJaegersberg/companion_cube_ggml) by [@KnutJaegersberg](https://huggingface.co/KnutJaegersberg)

## Credits
This is my first attempt at fine-tuning a large language model. It wouldn't be possible without [Axolotl](https://github.com/OpenAccess-AI-Collective/axolotl) and [Runpod](https://www.runpod.io/). The axolotl config file can be found [here](https://github.com/steve-cse/mello/blob/master/mello.yml).

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)
