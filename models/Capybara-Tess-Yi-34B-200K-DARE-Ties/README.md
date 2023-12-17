---
license: other
license_name: yi-license
license_link: https://huggingface.co/01-ai/Yi-34B/blob/main/LICENSE
language:
- en
library_name: transformers
pipeline_tag: text-generation
---
# This is not a great model, succeeded by a new merge: **https://huggingface.co/brucethemoose/CaPlatTessDolXaBoros-Yi-34B-200K-DARE-Ties**

**NousResearch/Nous-Capybara-34B**, **migtissera/Tess-M-v1.2** and **migtissera/Tess-M-v1.3** merged with a new, experimental implementation of "dare ties" via mergekit. See:

> Language Models are Super Mario: Absorbing Abilities from Homologous Models as a Free Lunch

https://github.com/yule-BUAA/MergeLM

https://github.com/cg123/mergekit/tree/dare-tokenizer

Highly experimental and still being tested! But this should yield a better merge than a typical linear/slerp merge or even a ties merge.
***

Merged with the following config, and the tokenizer from Yi Llamafied:
```
models:
  - model: /home/alpha/Storage/Models/Raw/larryvrh_Yi-34B-200K-Llamafied
    # no parameters necessary for base model
  - model: /home/alpha/Storage/Models/Raw/migtissera_Tess-M-v1.3
    parameters:
      weight: 0.50
      density: 0.56
  - model: /home/alpha/Storage/Models/Raw/migtissera_Tess-M-v1.2
    parameters:
      weight: 0.20
      density: 0.50
  - model: /home/alpha/Storage/Models/Raw/Nous-Capybara-34B
    parameters:
      weight: 0.50
      density: 0.56
merge_method: dare_ties
base_model: /home/alpha/Storage/Models/Raw/larryvrh_Yi-34B-200K-Llamafied
parameters:
  int8_mask: true
dtype: bfloat16
```

Tess 1.2 (at a low weight) and 1.3 were used because, according to the trainer, they were trained on different datasets: https://migel.substack.com/p/learnings-from-training-tess

As the Tess creator warned about, if the model repeats at high context like Tess 1.2, let me know!

I chose not to include other finetunes, such as Dolphin, because they aren't trained on the 200K base. If any other 200K finetunes pop up, let me know.

***

## Prompt template: Orca-Vicuna

```
SYSTEM: {system_message}
USER: {prompt}
ASSISTANT:

```
Being a Yi model, try disabling the BOS token and/or running a lower temperature with MinP if output doesn't seem right.

Sometimes the model "spells out" the stop token as `</s>` like Capybara, so you may need to add `</s>` as an additional stopping condition. 

***

Credits:

https://github.com/cg123/mergekit/tree/dare-tokenizer

https://huggingface.co/NousResearch/Nous-Capybara-34B/

https://huggingface.co/migtissera/Tess-M-v1.2

https://huggingface.co/migtissera/Tess-M-v1.3

https://huggingface.co/larryvrh/Yi-34B-200K-Llamafied

https://huggingface.co/01-ai/Yi-34B-200K