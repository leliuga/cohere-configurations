---
license: other
---

A product of an amateur merger. I like it due to the fact that it combines both Orca2 understanding and Pyg's dialogue style (using mythalion for consistency). - shotmisser64

This model was made by ShotMisser64 using the following mergekit yaml:
```
slices:
  - sources:
    - model: output/Orca2flat
      layer_range: [0, 13]
  - sources:
    - model: PygmalionAI/mythalion-13b
      layer_range: [3, 22]
  - sources:
    - model: output/Orca2flat
      layer_range: [14, 27]
  - sources:
    - model: PygmalionAI/mythalion-13b
      layer_range: [23, 40]
merge_method: passthrough
dtype: float16
```

The Orca2flat model uses the following mergekit yaml:
```
merge_method: task_arithmetic
base_model: TheBloke/Llama-2-13B-fp16
models:
  - model: TheBloke/Llama-2-13B-fp16
  - model: microsoft/Orca-2-13b
    parameters:
      weight: 1.0
dtype: float16
```

Found something interesting or would you like your own custom merge? Visit our community at https://koboldai.org/discord

Please respect the license of the origin models.