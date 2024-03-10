---
base_model: Mihaiii/Pallas-0.5
inference: false
license: other
license_name: yi-license
license_link: https://huggingface.co/01-ai/Yi-34B/blob/main/LICENSE
metrics:
- accuracy
---

This is a frankenmerge of [Mihaiii/Pallas-0.5](https://huggingface.co/Mihaiii/Pallas-0.5) .
It was done using [mergekit](https://github.com/cg123/mergekit).

It works well with long system prompts.

It isn't generic in a sense that it shouldn't be used for story telling, for example, but only for reasoning and text comprehension.

This model is trained on a private dataset.

# Prompt Format:

```
SYSTEM: <ANY SYSTEM CONTEXT>
USER: 
ASSISTANT:
```

Merge config:
```yaml
slices:
  - sources:
    - model: "Mihaiii/Pallas-0.5"
      layer_range: [0, 60]
  - sources:
    - model: "Mihaiii/Pallas-0.5"
      layer_range: [58, 60]
  - sources:
    - model: "Mihaiii/Pallas-0.5"
      layer_range: [55, 56]
merge_method: passthrough
dtype: bfloat16
```

Quants:

[TheBloke/Pallas-0.5-frankenmerge-GGUF](https://huggingface.co/TheBloke/Pallas-0.5-frankenmerge-GGUF)

[TheBloke/Pallas-0.5-frankenmerge-GPTQ](https://huggingface.co/TheBloke/Pallas-0.5-frankenmerge-GPTQ)

[TheBloke/Pallas-0.5-frankenmerge-AWQ](https://huggingface.co/TheBloke/Pallas-0.5-frankenmerge-AWQ)
