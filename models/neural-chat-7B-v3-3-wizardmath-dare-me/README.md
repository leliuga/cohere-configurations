---
license: other
license_name: microsoft-research-license
license_link: LICENSE
tags:
- merge
---

**Update: Yeah, this strategy doesn't work. This ended up really devastating the model's performance.**

This model is an experiment involving mixing DARE TIE merger with a task arithmetic merger to attempt to merge models with less loss.

DARE TIE mergers are [very strong at transferring strengths](https://medium.com/@minh.hoque/paper-explained-language-models-are-super-mario-2ebce6c2cf35) while merging a minimal part of the model. For larger models, 90-99% of delta parameters from SFT models can be dropped while retaining most of the benefits if they are rescaled and consensus merged back into the model.

For 7B models, we can't drop as many of the parameters and retain the model's strengths. In the original paper, the WizardMath model showed transferrable skills when 90% of the parameters were dropped but showed more strength when 70% were dropped. Experimentally, it appears that [even lower drop rates like 40%](https://github.com/cg123/mergekit/issues/26) have performed the best even for larger 34B models. In some instances, [even densities as high as 80% create an unstable merger](https://huggingface.co/jan-hq/supermario-v1), making DARE TIES unsuitable for merging models.

This is an experiment utilizing two merger techniques together to try and transfer skills between finetuned models. If we were to DARE TIE a low density merger onto the base Mistral model and then task arithmetic merge those low density delta weights onto a finetune, could we still achieve skill transfer?

```
models: # mistral-wizardmath-dare-0.7-density
  - model: mistralai/Mistral-7B-v0.1
    # no parameters necessary for base model
  - model: WizardLM/WizardMath-7B-V1.1
    parameters:
      weight: 1
      density: 0.3
merge_method: dare_ties
base_model: mistralai/Mistral-7B-v0.1
parameters:
  normalize: true
  int8_mask: true
dtype: bfloat16

merge_method: task_arithmetic
base_model: mistralai/Mistral-7B-v0.1
models:
  - model: mistral-wizardmath-dare-0.7-density
  - model: Intel/neural-chat-7b-v3-3
parameters:
  weight: 1.0
dtype: bfloat16
```

WizardMath is under the Microsoft Research License, Intel is Apache 2.0.