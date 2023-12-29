---
license: other
license_name: yi-license
license_link: https://huggingface.co/01-ai/Yi-34B/blob/main/LICENSE
language:
- en
library_name: transformers
pipeline_tag: text-generation
tags:
- text-generation-inference
- merge
---

### Possibly obsolete, replaced by https://huggingface.co/brucethemoose/Yi-34B-200K-DARE-merge-v5

Old model description below:
***


**Dolphin-2.2-yi-34b-200k**, **Nous-Capybara-34B**, **Tess-M-v1.4**, **Airoboros-3_1-yi-34b-200k**, **PlatYi-34B-200K-Q**, and **Una-xaberius-34b-v1beta** merged with a new, experimental implementation of "dare ties" via mergekit. See:

> [Language Models are Super Mario: Absorbing Abilities from Homologous Models as a Free Lunch](https://github.com/yule-BUAA/MergeLM)

> https://github.com/cg123/mergekit/tree/dare


This variant is merged with a "higher than recommended" density with with the following config, and the tokenizer from chargoddard's Yi-Llama:
```
models:
  - model: /home/alpha/Storage/Models/Raw/chargoddard_Yi-34B-200K-Llama
    # no parameters necessary for base model
  - model: /home/alpha/Storage/Models/Raw/migtissera_Tess-34B-v1.4
    parameters:
      weight: 0.19
      density: 0.6
  - model: /home/alpha//Storage/Models/Raw/bhenrym14_airoboros-3_1-yi-34b-200k
    parameters:
      weight: 0.14
      density: 0.5
  - model: /home/alpha/Storage/Models/Raw/Nous-Capybara-34B
    parameters:
      weight: 0.19
      density: 0.6
  - model: /home/alpha/Storage/Models/Raw/kyujinpy_PlatYi-34B-200K-Q
    parameters:
      weight: 0.14
      density: 0.5
  - model: /home/alpha/FastModels/ehartford_dolphin-2.2-yi-34b-200k
    parameters:
      weight: 0.19
      density: 0.6
  - model: /home/alpha/FastModels/fblgit_una-xaberius-34b-v1beta
    parameters:
      weight: 0.15
      density: 0.08
merge_method: dare_ties
base_model: /home/alpha/Storage/Models/Raw/chargoddard_Yi-34B-200K-Llama
parameters:
  int8_mask: true
dtype: bfloat16
```
***
## Prompt template: Orca-Vicuna?
```
SYSTEM: {system_message}
USER: {prompt}
ASSISTANT:
```
It might recognize ChatML from Dolphin+Xaberius, and Llama-chat from Airoboros.

Sometimes the model "spells out" the stop token as `</s>` like Capybara, so you may need to add `</s>` as an additional stopping condition.
***
## Running
Being a Yi model, try disabling the BOS token and/or running a lower temperature with 0.05-0.13 MinP, a little repitition penalty, and no other samplers. Yi tends to run "hot" by default.

24GB GPUs can run Yi-34B-200K models at **45K-75K context** with exllamav2. I go into more detail in this [post](https://old.reddit.com/r/LocalLLaMA/comments/1896igc/how_i_run_34b_models_at_75k_context_on_24gb_fast/)

I recommend exl2 quantizations profiled on data similar to the desired task. It is especially sensitive to the quantization data at low bpw! I published my own quantizations on vicuuna chat + fiction writing here: [4bpw](https://huggingface.co/brucethemoose/CaPlatTessDolXaBoros-34B-200K-exl2-4bpw-fiction) [3.1bpw](https://huggingface.co/brucethemoose/CaPlatTessDolXaBoros-34B-200K-exl2-4bpw-fiction)

To load this in full-context backends like transformers and vllm, you *must* change `max_position_embeddings` in config.json to a lower value than 200,000, otherwise you will OOM! 
***
## Testing Notes

Various densities were tested with perplexity tests and long context prompts. Relatively high densities seem to perform better, contrary to the findings of the Super Mario paper.

This particular version is merged with more than the "recommended" max density of 0.5. It seems to result in even better perplexity, and a much higher position on the hf leaderboard, but I'm not sure if this translates to better output.

Weights that add up to 1 seems to be optimal.

Dare Ties is also resulting in seemingly better, lower perplexity merges than a regular ties merge, task arithmetic or a slerp merge.

Xaberuis is not a 200K model, hence it was merged at a very low density to try and preserve Yi 200K's long context performance while still inheriting some of Xaberius's performance. 

I chose not to include other finetunes because they aren't trained on the 200K base. If any other 200K finetunes pop up, let me know.
***
## Credits:

https://github.com/cg123/mergekit/tree/dare

https://huggingface.co/ehartford/dolphin-2.2-yi-34b-200k

https://huggingface.co/kyujinpy/PlatYi-34B-200K-Q

https://huggingface.co/NousResearch/Nous-Capybara-34B/

https://huggingface.co/bhenrym14/airoboros-3_1-yi-34b-200k

https://huggingface.co/migtissera/Tess-M-v1.4

https://huggingface.co/fblgit/una-xaberius-34b-v1beta

https://huggingface.co/chargoddard/Yi-34B-200K-Llama

https://huggingface.co/01-ai/Yi-34B-200K