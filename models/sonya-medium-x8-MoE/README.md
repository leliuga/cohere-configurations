---
license: wtfpl
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/6215ce9abfcb3893344dd0a2/uTxtkyc1-V6NT_5xfM5Us.png)

I present my Magnum Opus of llm merges for 2023. This is a monster of a model created from merging x8 [sonya-medium](https://huggingface.co/dillfrescott/sonya-medium)'s into one mixture of experts.

Enjoy! ;)

Config:
```
base_model: dillfrescott/sonya-medium
gate_mode: hidden
dtype: bfloat16
experts:
  - source_model: dillfrescott/sonya-medium
    positive_prompts: [""]
  - source_model: dillfrescott/sonya-medium
    positive_prompts: [""]
  - source_model: dillfrescott/sonya-medium
    positive_prompts: [""]
  - source_model: dillfrescott/sonya-medium
    positive_prompts: [""]
  - source_model: dillfrescott/sonya-medium
    positive_prompts: [""]
  - source_model: dillfrescott/sonya-medium
    positive_prompts: [""]
  - source_model: dillfrescott/sonya-medium
    positive_prompts: [""]
  - source_model: dillfrescott/sonya-medium
    positive_prompts: [""]
```

P.S. Be careful with K quants of this model as they may be glitched!

*The cover image is heavily inspired by Muzan Kibutsuji from Demon Slayer.*

Example outputs and reasonings. (The following outputs are a q4_0 version so the fully unquantized model would likely perform even better):
![image/png](https://cdn-uploads.huggingface.co/production/uploads/6215ce9abfcb3893344dd0a2/8kl6KXEPyMG7bZZOVz9kv.png)


![image/png](https://cdn-uploads.huggingface.co/production/uploads/6215ce9abfcb3893344dd0a2/-zxKYEYoT88ffryuVrVEa.png)


![image/png](https://cdn-uploads.huggingface.co/production/uploads/6215ce9abfcb3893344dd0a2/up005XvFd6bkWxlOq_hBo.png)


![image/png](https://cdn-uploads.huggingface.co/production/uploads/6215ce9abfcb3893344dd0a2/xx5x5SYuOF50DNvo06t_v.png)


![image/png](https://cdn-uploads.huggingface.co/production/uploads/6215ce9abfcb3893344dd0a2/WOf2jV3Bq2MOVViS58IM6.png)


![image/png](https://cdn-uploads.huggingface.co/production/uploads/6215ce9abfcb3893344dd0a2/PYQZGdOqBntakYAcT1DaV.png)
