---
license: other
license_name: yi-license
license_link: https://huggingface.co/01-ai/Yi-34B/blob/main/LICENSE
language:
- en
library_name: transformers
pipeline_tag: text-generation
tags:
- merge
---

# Obsolete, succeeded by a new merge: **https://huggingface.co/brucethemoose/CaPlatTessDolXaBoros-Yi-34B-200K-DARE-Ties-HighDensity**

***

**NousResearch/Nous-Capybara-34B** and **migtissera/Tess-M-Creative-v1.0** ties merged with mergekit.

I would suggest an exllama version for local inference with 40K+ context in 24GB:
https://huggingface.co/brucethemoose/Capybara-Tess-Yi-34B-200K-exl2-4bpw-fiction
https://huggingface.co/brucethemoose/Capybara-Tess-Yi-34B-200K-exl2-31bpw-fiction

Merged with the following config:

```
models:
  - model: /home/alpha/Storage/Models/Raw/larryvrh_Yi-34B-200K-Llamafied
    # no parameters necessary for base model
  - model: /home/alpha/Storage/Models/Raw/migtissera_Tess-M-v1.0
    parameters:
      density: 0.6
      weight: 1.0
  - model: /home/alpha/Storage/Models/Raw/Nous-Capybara-34B
    parameters:
      density: 0.6
      weight: 1.0
merge_method: ties
base_model: //home/alpha/Storage/Models/Raw/larryvrh_Yi-34B-200K-Llamafied
parameters:
  normalize: true
  int8_mask: true
dtype: float16
```

Both are 200K context models with Vicuna syntax, so: 

# Prompt Format:

```
SYSTEM: ...
USER: ...
ASSISTANT: ...
```
Sometimes the model "spells out" the stop token as `</s>` like Capybara, so you may need to add `</s>` this as an additional stopping condition.

***

Credits:

https://github.com/cg123/mergekit

https://huggingface.co/NousResearch/Nous-Capybara-34B/discussions

https://huggingface.co/migtissera/Tess-M-Creative-v1.0

https://huggingface.co/larryvrh/Yi-34B-200K-Llamafied

https://huggingface.co/01-ai/Yi-34B-200K