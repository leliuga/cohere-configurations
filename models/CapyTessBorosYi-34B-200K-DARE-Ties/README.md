---
license: other
license_name: yi-license
license_link: https://huggingface.co/01-ai/Yi-34B/blob/main/LICENSE
language:
- en
library_name: transformers
pipeline_tag: text-generation
---

**NousResearch/Nous-Capybara-34B**, **migtissera/Tess-M-v1.3** and **bhenrym14/airoboros-3_1-yi-34b-200k** merged with a new, experimental implementation of "dare ties" via mergekit. See:

> Language Models are Super Mario: Absorbing Abilities from Homologous Models as a Free Lunch

https://github.com/yule-BUAA/MergeLM

https://github.com/cg123/mergekit/tree/dare

***

Merged with the following config, and the tokenizer from chargoddard's Yi-Llama:
```
models:
  - model: /home/alpha/Storage/Models/Raw/chargoddard_Yi-34B-200K-Llama
    # no parameters necessary for base model
  - model: /home/alpha/Storage/Models/Raw/migtissera_Tess-M-v1.3
    parameters:
      weight: 0.41
      density: 0.50
  - model: /home/alpha//Storage/Models/Raw/bhenrym14_airoboros-3_1-yi-34b-200k
    parameters:
      weight: 0.18
      density: 0.46
  - model: /home/alpha/Storage/Models/Raw/Nous-Capybara-34B
    parameters:
      weight: 0.41
      density: 0.50
merge_method: dare_ties
base_model: /home/alpha/Storage/Models/Raw/chargoddard_Yi-34B-200K-Llama
parameters:
int8_mask: true
dtype: bfloat16
```

dare_ties is testing with better perplexity than a regular ties merge with the same merge configuration. Model weights that add up to one also seem optimal from testing. And high context results seem... better than the previous dare merge with Tess 1.2.

I chose not to include other finetunes, such as Dolphin, because they aren't trained on the 200K base. If any other 200K finetunes pop up, let me know.

***

## Prompt template: Orca-Vicuna

```
SYSTEM: {system_message}
USER: {prompt}
ASSISTANT:

```
Being a Yi model, try disabling the BOS token and/or running a lower temperature with MinP (and no other samplers) if output doesn't seem right. Yi tends to run "hot" by default.

Sometimes the model "spells out" the stop token as `</s>` like Capybara, so you may need to add `</s>` as an additional stopping condition. It also might respond to the llama-2 chat format.

***

Credits:

https://github.com/cg123/mergekit/tree/dare

https://huggingface.co/NousResearch/Nous-Capybara-34B/

https://huggingface.co/bhenrym14/airoboros-3_1-yi-34b-200k

https://huggingface.co/migtissera/Tess-M-v1.3

https://huggingface.co/chargoddard/Yi-34B-200K-Llama

https://huggingface.co/01-ai/Yi-34B-200K