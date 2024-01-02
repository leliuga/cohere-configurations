---
license: wtfpl
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/6215ce9abfcb3893344dd0a2/CVx0nwjGsCgQujIn7JcmK.png)

x8 Mixture of experts version of https://huggingface.co/SanjiWatsuki/Sonya-7B

Config:
```
base_model: SanjiWatsuki/Sonya-7B
gate_mode: hidden
dtype: bfloat16
experts:
  - source_model: SanjiWatsuki/Sonya-7B
    positive_prompts: [""]
  - source_model: SanjiWatsuki/Sonya-7B
    positive_prompts: [""]
  - source_model: SanjiWatsuki/Sonya-7B
    positive_prompts: [""]
  - source_model: SanjiWatsuki/Sonya-7B
    positive_prompts: [""]
  - source_model: SanjiWatsuki/Sonya-7B
    positive_prompts: [""]
  - source_model: SanjiWatsuki/Sonya-7B
    positive_prompts: [""]
  - source_model: SanjiWatsuki/Sonya-7B
    positive_prompts: [""]
  - source_model: SanjiWatsuki/Sonya-7B
    positive_prompts: [""]
```