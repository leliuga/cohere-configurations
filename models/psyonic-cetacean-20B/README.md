---
license: other
license_name: microsoft-research-license
tags:
- storywriting
- text adventure
- not-for-all-audiences
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/6459a451abdbb77c4c6d8258/uNoKlBulkRF3mCoMgetGs.png)
---

Presenting the FP16 files for Psyonic-Cetacean-20B! This is an experimental Llama2-based stack merge based on the models and recipe below:

- [KoboldAI/PsyFighter-2-13b](https://huggingface.co/KoboldAI/LLaMA2-13B-Psyfighter2-GGUF)
- [microsoft/Orca-2-13b](https://huggingface.co/microsoft/Orca-2-13b)

```yaml
  slices:
  - sources:
    - model: Orca2flat
      layer_range: [0, 16]
  - sources:
    - model: LLaMA2-13B-Psyfighter2 (FP16 not yet available)
      layer_range: [8, 24]
  - sources:
    - model: Orca2flat
      layer_range: [17, 32]
  - sources:
    - model: LLaMA2-13B-Psyfighter2 (FP16 not yet available)
      layer_range: [25, 40]
merge_method: passthrough
dtype: float16
```

Note: while we did run an inverted merge the output was not satisfactory and will not be released.

We first flatted the additional ChatML vocabulary tokens out of Orca-2-13B, then performed a stack merge with Psyfighter-2-13B. The results surprised us with their vividness, freshness of prose, obedience to instruction prompting, and formatting cohesion.

This model is focused on storywriting and text adventure, with a side order of Assistant and Chat functionality. Like its ancestor Psyfighter-2 this model will function better if you let it improvise and riff on your concepts rather than feeding it an excess of detail.
Additionally, either the removal of the ChatML vocab or the stack merging process itself has resulted in not only an uncensored model but an actively anti-censored model, so please be aware that this model can and will kill you during adventures or output NSFW material if prompted accordingly.

During testing, the model exhibited an especially strong affinity for science fiction and space opera writing, while handling fantasy elements quite well and horror elements slightly less so. Refer to the Psyfighter-2 model card for best prompting practices.

Despite that, we have tested the model out to 16000 context via Rope scaling and the model does not drive towards NSFW on its own. It will follow your tone and style very well.

Please enjoy, and if you encounter anything exciting or weird, please reach out to me at [jebcarter@pm.me].

Special thanks as always to the KoboldAI crew who provided the mergebox, testing, and feedback on this model, and to gelukuMLG for the model mascot!