---
language:
- en
license: apache-2.0
library_name: transformers
datasets:
- WizardLM/WizardLM_evol_instruct_V2_196k
model-index:
- name: open-llama-3b-v2-wizard-evol-instuct-v2-196k
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
      value: 41.81
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=harborwater/open-llama-3b-v2-wizard-evol-instuct-v2-196k
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
      value: 73.01
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=harborwater/open-llama-3b-v2-wizard-evol-instuct-v2-196k
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
      value: 26.36
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=harborwater/open-llama-3b-v2-wizard-evol-instuct-v2-196k
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
      value: 38.99
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=harborwater/open-llama-3b-v2-wizard-evol-instuct-v2-196k
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
      value: 66.69
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=harborwater/open-llama-3b-v2-wizard-evol-instuct-v2-196k
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
      value: 1.9
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=harborwater/open-llama-3b-v2-wizard-evol-instuct-v2-196k
      name: Open LLM Leaderboard
---

Trained on 1 epoch of the WizardLM_evol_instruct_v2_196k dataset

Link to [GGUF](https://huggingface.co/maddes8cht/harborwater-open-llama-3b-v2-wizard-evol-instuct-v2-196k-gguf) formats.

Prompt template:
```
### HUMAN:
{prompt}

### RESPONSE:
<leave a newline for the model to answer>
```

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)


# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_harborwater__open-llama-3b-v2-wizard-evol-instuct-v2-196k)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |41.46|
|AI2 Reasoning Challenge (25-Shot)|41.81|
|HellaSwag (10-Shot)              |73.01|
|MMLU (5-Shot)                    |26.36|
|TruthfulQA (0-shot)              |38.99|
|Winogrande (5-shot)              |66.69|
|GSM8k (5-shot)                   | 1.90|

