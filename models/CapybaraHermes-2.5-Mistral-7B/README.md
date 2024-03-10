---
language:
- en
license: apache-2.0
library_name: trl
tags:
- distilabel
- dpo
- rlaif
- rlhf
datasets:
- argilla/dpo-mix-7k
base_model: teknium/OpenHermes-2.5-Mistral-7B
model-index:
- name: CapybaraHermes-2.5-Mistral-7B
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
      value: 65.78
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=argilla/CapybaraHermes-2.5-Mistral-7B
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
      value: 85.45
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=argilla/CapybaraHermes-2.5-Mistral-7B
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
      value: 63.13
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=argilla/CapybaraHermes-2.5-Mistral-7B
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
      value: 56.91
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=argilla/CapybaraHermes-2.5-Mistral-7B
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
      value: 78.3
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=argilla/CapybaraHermes-2.5-Mistral-7B
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
      value: 59.29
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=argilla/CapybaraHermes-2.5-Mistral-7B
      name: Open LLM Leaderboard
---

# CapybaraHermes-2.5-Mistral-7B
<div>
    <img src="https://cdn-uploads.huggingface.co/production/uploads/60420dccc15e823a685f2b03/Vmr0FtTvnny6Snm-UDM_n.png">
</div>

<p align="center">
  <a href="https://github.com/argilla-io/distilabel">
    <img src="https://raw.githubusercontent.com/argilla-io/distilabel/main/docs/assets/distilabel-badge-light.png" alt="Built with Distilabel" width="200" height="32"/>
  </a>
</p>


This model is the launching partner of the [capybara-dpo dataset](https://huggingface.co/datasets/argilla/distilabel-capybara-dpo-9k-binarized) build with ⚗️ distilabel. It's a preference tuned [OpenHermes-2.5-Mistral-7B](https://huggingface.co/teknium/OpenHermes-2.5-Mistral-7B). 

CapybaraHermes has been preference tuned with LoRA and TRL for 3 epochs using argilla's [dpo mix 7k](https://huggingface.co/datasets/argilla/dpo-mix-7k). 

To test the impact on multi-turn performance we have used MTBench. We also include the Nous Benchmark results and Mistral-7B-Instruct-v0.2 for reference as it's a strong 7B model on MTBench:

| Model                             | AGIEval | GPT4All | TruthfulQA | Bigbench | MTBench First Turn | MTBench Second Turn | Nous avg. | MTBench avg. |
|-----------------------------------|---------|---------|------------|----------|------------|-------------|-----------|--------------|
| argilla/CapybaraHermes-2.5-Mistral-7B     | **43.8**    | **73.35**   | 57.07      | **42.44**    | 8.24375    | **7.5625**      | 54.16     | **7.903125** |
| teknium/OpenHermes-2.5-Mistral-7B | 42.75   | 72.99   | 52.99      | 40.94    | **8.25**       | 7.2875      | 52.42     | 7.76875      |
| Mistral-7B-Instruct-v0.2          | 38.5    | 71.64   | **66.82**  | 42.29    | 7.8375     | 7.1         | **54.81** | 7.46875      |

The most interesting aspect in the context of the capybara-dpo dataset is the increased performance in MTBench Second Turn scores.

For the merge lovers, we also preference tuned Beagle14-7B with a mix of capybara-dpo and distilabel orca pairs using the same recipe as NeuralBeagle (see [ YALL - Yet Another LLM Leaderboard](https://huggingface.co/spaces/mlabonne/Yet_Another_LLM_Leaderboard) for reference):

|                                                               Model                                                                |AGIEval|GPT4All|TruthfulQA|Bigbench|Average|
|------------------------------------------------------------------------------------------------------------------------------------|------:|------:|---------:|-------:|------:|
|[DistilabelBeagle14-7B](https://huggingface.co/dvilasuero/DistilabelBeagle14-7B)|  45.29|  76.92|     71.66|   48.78|  60.66|




## Model Details

### Model Description

<!-- Provide a longer summary of what this model is. -->

This is the model card of a 🤗 transformers model that has been pushed on the Hub. This model card has been automatically generated.

- **Developed by:** Argilla
- **Shared by [optional]:** Argilla
- **Model type:** 7B chat model
- **Language(s) (NLP):** English
- **License:** Same as OpenHermes
- **Finetuned from model [optional]:** [OpenHermes-2.5-Mistral-7B](https://huggingface.co/teknium/OpenHermes-2.5-Mistral-7B)


# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_argilla__CapybaraHermes-2.5-Mistral-7B)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |68.14|
|AI2 Reasoning Challenge (25-Shot)|65.78|
|HellaSwag (10-Shot)              |85.45|
|MMLU (5-Shot)                    |63.13|
|TruthfulQA (0-shot)              |56.91|
|Winogrande (5-shot)              |78.30|
|GSM8k (5-shot)                   |59.29|

