---
license: cc-by-nc-4.0
---


![image/png](https://cdn-uploads.huggingface.co/production/uploads/630dfb008df86f1e5becadc3/vwcJfOnL-2QDJ0ShfxRJ5.png)



---

# Disclaimer:
## This model is experimental, do not expect everything to work.

This model uses the Chatml **prompting format**

---


Beeg noromaid on ***steroids***. Suitable for RP, ERP.

This model was trained on the Zloss fork of Charles, and should fix issue the model had.

Use Chatml prompt format, but not the special token.

The reason is that Axolotl merge the finetune with the base model at 1.0 weight basically, but this is too much, so I use another script available [HERE](https://github.com/DocShotgun/LLM-notebooks/blob/main/weighted-lora-merge.ipynb) to merge with less weight, sadly, it don't take the special Chatml token. It's like Orca2 for the matter.


## Credits:
- Undi
- IkariDev

<!-- description start -->
## Description

<!-- [Recommended settings - contributed by localfultonextractor](https://files.catbox.moe/ue0tja.json) -->

This repo contains FP16 files of Noromaid-v0.4-Mixtral-Instruct-8x7b-Zloss.

[FP16 - by IkariDev and Undi](https://huggingface.co/NeverSleep/Noromaid-v0.4-Mixtral-Instruct-8x7b-Zloss)

<!-- [GGUF - By TheBloke](https://huggingface.co/TheBloke/Athena-v4-GGUF)-->

<!-- [GPTQ - By TheBloke](https://huggingface.co/TheBloke/Athena-v4-GPTQ)-->

<!-- [exl2[8bpw-8h] - by AzureBlack](https://huggingface.co/AzureBlack/Echidna-13b-v0.3-8bpw-8h-exl2)-->

<!-- [AWQ - By TheBloke](https://huggingface.co/TheBloke/Athena-v4-AWQ)-->

<!-- [fp16 - by IkariDev+Undi95](https://huggingface.co/IkariDev/Athena-v4)-->

[GGUF - by IkariDev and Undi](https://huggingface.co/NeverSleep/Noromaid-v0.4-Mixtral-Instruct-8x7b-Zloss-GGUF)
<!-- [OLD(GGUF - by IkariDev+Undi95)](https://huggingface.co/IkariDev/Athena-v4-GGUF)-->

## Ratings:

Note: We have permission of all users to upload their ratings, we DONT screenshot random reviews without asking if we can put them here!

No ratings yet!

If you want your rating to be here, send us a message over on DC and we'll put up a screenshot of it here. DC name is "ikaridev" and "undi".

<!-- description end -->
<!-- prompt-template start -->
### Prompt format: Chatml
```
<|im_start|>system
{sysprompt}<|im_end|>
<|im_start|>user
{input}<|im_end|>
<|im_start|>assistant
{output}<|im_end|>
```

## Datasets used:

- Aesir 1, 2 & 3 modified by us, credit to ([MinervaAI](https://huggingface.co/MinervaAI) / [Gryphe](https://huggingface.co/Gryphe))
- [LimaRP-20231109](https://huggingface.co/datasets/lemonilia/LimaRP) ([Lemonilia](https://huggingface.co/lemonilia))
- [ToxicQAFinal](https://huggingface.co/datasets/NobodyExistsOnTheInternet/ToxicQAFinal) ([NobodyExistsOnTheInternet](https://huggingface.co/NobodyExistsOnTheInternet)
- [No-robots-ShareGPT](https://huggingface.co/datasets/Doctor-Shotgun/no-robots-sharegpt) ([Doctor-Shotgun](https://huggingface.co/Doctor-Shotgun))


## Others

Undi: If you want to support me, you can [here](https://ko-fi.com/undiai).

IkariDev: Visit my [retro/neocities style website](https://ikaridevgit.github.io/) please kek