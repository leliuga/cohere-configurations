---
language:
- en
license: other
library_name: transformers
tags:
- text-generation-inference
- merge
license_name: yi-license
license_link: https://huggingface.co/01-ai/Yi-34B/blob/main/LICENSE
pipeline_tag: text-generation
model-index:
- name: Yi-34B-200K-DARE-merge-v5
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
      value: 66.47
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=brucethemoose/Yi-34B-200K-DARE-merge-v5
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
      value: 85.54
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=brucethemoose/Yi-34B-200K-DARE-merge-v5
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
      value: 77.22
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=brucethemoose/Yi-34B-200K-DARE-merge-v5
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
      value: 57.46
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=brucethemoose/Yi-34B-200K-DARE-merge-v5
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
      value: 82.24
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=brucethemoose/Yi-34B-200K-DARE-merge-v5
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
      value: 62.93
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=brucethemoose/Yi-34B-200K-DARE-merge-v5
      name: Open LLM Leaderboard
---

# Succeeded by a new merge: https://huggingface.co/brucethemoose/Yi-34B-200K-DARE-merge-v7

***


[**Nous-Capybara-34B**](https://huggingface.co/NousResearch/Nous-Capybara-34B/), [**Tess-M-v1.4**](https://huggingface.co/migtissera/Tess-34B-v1.4), [**Airoboros-3_1-yi-34b-200k**](https://huggingface.co/bhenrym14/airoboros-3_1-yi-34b-200k), [**PlatYi-34B-200K-Q**](https://huggingface.co/kyujinpy/PlatYi-34B-200k-Q-FastChat), [**Pallas-0.4**](https://huggingface.co/Mihaiii/Pallas-0.4), [**Yi-34B-200K-AEZAKMI-v2**](https://huggingface.co/adamo1139/Yi-34B-200K-AEZAKMI-v2), and a tiny bit of [**SUS-Chat-34B**](https://huggingface.co/SUSTech/SUS-Chat-34B) merged with a new, experimental implementation of "dare ties" via mergekit. See:

> [Language Models are Super Mario: Absorbing Abilities from Homologous Models as a Free Lunch](https://github.com/yule-BUAA/MergeLM)

> https://github.com/cg123/mergekit/tree/dare

***
## Prompt template: Orca-Vicuna
```
SYSTEM: {system_message}
USER: {prompt}
ASSISTANT:
```
It might recognize ChatML, or maybe Llama-chat from Airoboros.

Sometimes the model "spells out" the stop token as `</s>` like Capybara, so you may need to add `</s>` as an additional stopping condition.
***
## Running
Being a Yi model, try running a lower temperature with 0.02-0.1 MinP, a little repetition penalty, and no other samplers. Yi tends to run "hot" by default, and it really needs MinP to cull the huge vocabulary.

24GB GPUs can run Yi-34B-200K models at **45K-75K context** with exllamav2, and performant UIs like [exui](https://github.com/turboderp/exui). I go into more detail in this [post](https://old.reddit.com/r/LocalLLaMA/comments/1896igc/how_i_run_34b_models_at_75k_context_on_24gb_fast/)

I recommend exl2 quantizations profiled on data similar to the desired task. It is especially sensitive to the quantization data at low bpw. I've published my own fiction-oriented quantizations here: https://huggingface.co/collections/brucethemoose/most-recent-merge-65742644ca03b6c514afa204

To load this in full-context backends like transformers, you *must* change `max_position_embeddings` in config.json to a lower value than 200,000, otherwise you will OOM! 
***
## Testing Notes

Merged in mergekit with the following config, and the tokenizer from chargoddard's Yi-Llama:

```
models:
  - model: /home/alpha/Storage/Models/Raw/chargoddard_Yi-34B-200K-Llama
    # No parameters necessary for base model
  - model: /home/alpha/Storage/Models/Raw/migtissera_Tess-34B-v1.4
    # Less weight than previous merge since Pallas is a finetune of Tess
    parameters:
      weight: 0.14
      density: 0.62
  - model: /home/alpha/FastModels/Mihaiii_Pallas-0.4
    parameters:
      weight: 0.14
      density: 0.62
  - model: /home/alpha//Storage/Models/Raw/bhenrym14_airoboros-3_1-yi-34b-200k
    parameters:
      weight: 0.14
      density: 0.52
  - model: /home/alpha/Storage/Models/Raw/Nous-Capybara-34B
    parameters:
      weight: 0.22
      density: 0.62
  - model: /home/alpha/Storage/Models/Raw/kyujinpy_PlatYi-34B-200k-Q-FastChat
    parameters:
      weight: 0.14
      density: 0.52
  #- model: /home/alpha/Storage/Models/Raw/ehartford_dolphin-2.2-yi-34b-200k
  #  Dolphin 200K seems to be broken according to multiple leaderboards and perplexity tests?
  #  parameters:
  #    weight: 0.15
  #    density: 0.6
  - model: /home/alpha/Models/Raw/adamo1139_Yi-34B-200K-AEZAKMI-v2
    parameters:
      weight: 0.14
      density: 0.52
  - model: /home/alpha/Models/Raw/SUSTech_SUS-Chat-34B/
  # Very low density and low weight since its a Yi 4K finetune, to try and preserve long context performance while "keeping" some of SUS
    parameters:
      weight: 0.08
      density: 0.08
merge_method: dare_ties
base_model: /home/alpha/Storage/Models/Raw/chargoddard_Yi-34B-200K-Llama
parameters:

  int8_mask: true
dtype: bfloat16
```

Various densities were tested with perplexity tests and long context prompts. Relatively high densities seem to perform better, contrary to the findings of the Super Mario paper.

This particular version is merged with more than the "recommended" max density of 0.5. It seems to result in even better perplexity, but I'm not sure if this translates to better output.

Weights that add up to 1 seems to be optimal.

Dare Ties is also resulting in seemingly better, lower perplexity merges than a regular ties merge, task arithmetic or a slerp merge.

SUS Chat is not a 200K model, hence it was merged at a very low density to try and preserve Yi 200K's long context performance while still inheriting some of SUS's performance. 

Dolphin 200K was taken out of this merge because it seems to be performing poorly for a 34B Dolphin model, like something went wrong during training?

I chose not to include other finetunes because they aren't trained on the 200K base. If any other 200K finetunes pop up, let me know.
***
## Credits:

https://github.com/cg123/mergekit/tree/dare

https://huggingface.co/NousResearch/Nous-Capybara-34B/

https://huggingface.co/bhenrym14/airoboros-3_1-yi-34b-200k

https://huggingface.co/migtissera/Tess-M-v1.4

https://huggingface.co/kyujinpy/PlatYi-34B-200k-Q-FastChat

https://huggingface.co/adamo1139/Yi-34B-200K-AEZAKMI-v2

https://huggingface.co/Mihaiii/Pallas-0.4

https://huggingface.co/SUSTech/SUS-Chat-34B

https://huggingface.co/chargoddard/Yi-34B-200K-Llama

https://huggingface.co/01-ai/Yi-34B-200K
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_brucethemoose__Yi-34B-200K-DARE-merge-v5)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |71.98|
|AI2 Reasoning Challenge (25-Shot)|66.47|
|HellaSwag (10-Shot)              |85.54|
|MMLU (5-Shot)                    |77.22|
|TruthfulQA (0-shot)              |57.46|
|Winogrande (5-shot)              |82.24|
|GSM8k (5-shot)                   |62.93|

