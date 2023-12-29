---
license: cc-by-nc-4.0
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/630dfb008df86f1e5becadc3/VKX2Z2yjZX5J8kXzgeCYO.png)


---

# Disclaimer:
## If you don't like this model, use [Noromaid 0.1.1](https://huggingface.co/NeverSleep/Noromaid-13b-v0.1.1), or [Noromaid 0.2](https://huggingface.co/NeverSleep/Noromaid-13b-v0.2)

You may use our custom **prompting format**(scroll down to download them!), or simple alpaca. **(Choose which fits best for you!)**

---


If you want a 7b, or 20b hit us up in the Community tab!

Merge was by me(IkariDev) alone this time.

FlatOrca(ChatML removed, sorry ChatML bros) + Noromaid 13b 0.2. Suitable for RP, ERP and general stuff.

[Recommended settings - No settings yet(Please suggest some over in the Community tab!)]

<!-- description start -->
## Description

<!-- [Recommended settings - contributed by localfultonextractor](https://files.catbox.moe/ue0tja.json) -->

This repo contains FP16 files of FlatOrcamaid-13b-v0.2.

[FP16 - by IkariDev and Undi](https://huggingface.co/NeverSleep/FlatOrcamaid-13b-v0.2)

<!-- [GGUF - By TheBloke](https://huggingface.co/TheBloke/Athena-v4-GGUF)-->

<!-- [GPTQ - By TheBloke](https://huggingface.co/TheBloke/Athena-v4-GPTQ)-->

<!-- [exl2[8bpw-8h] - by AzureBlack](https://huggingface.co/AzureBlack/Echidna-13b-v0.3-8bpw-8h-exl2)-->

<!-- [AWQ - By TheBloke](https://huggingface.co/TheBloke/Athena-v4-AWQ)-->

<!-- [fp16 - by IkariDev+Undi95](https://huggingface.co/IkariDev/Athena-v4)-->

[GGUF - by IkariDev and Undi](https://huggingface.co/NeverSleep/FlatOrcamaid-13b-v0.2-GGUF)
<!-- [OLD(GGUF - by IkariDev+Undi95)](https://huggingface.co/IkariDev/Athena-v4-GGUF)-->

## Ratings:

Note: We have permission of all users to upload their ratings, we DONT screenshot random reviews without asking if we can put them here!

No ratings yet!

If you want your rating to be here, send us a message over on DC and we'll put up a screenshot of it here. DC name is "ikaridev" and "undi".

<!-- description end -->
<!-- prompt-template start -->
## Prompt template: Custom format, or Alpaca

### Custom format:
UPDATED!! SillyTavern config files: [Context](https://files.catbox.moe/ifmhai.json), [Instruct](https://files.catbox.moe/ttw1l9.json).

### Alpaca:
```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:

```


```
slices:
  - sources:
      - model: NeverSleep/Noromaid-13b-v0.2
        layer_range: [0, 40]
      - model: OrcaFlat
        layer_range: [0, 40]
merge_method: slerp
base_model: NeverSleep/Noromaid-13b-v0.2
parameters:
  t:
    - filter: self_attn
      value: [0, 0.4, 0.2, 0.55, 0.8]
    - filter: mlp
      value: [0.7, 0.3, 0.4, 0.3, 0]
    - value: 0.33 # fallback for rest of tensors
dtype: float16
```

## Others

Undi: If you want to support me, you can [here](https://ko-fi.com/undiai).

IkariDev: Visit my [retro/neocities style website](https://ikaridevgit.github.io/) please kek