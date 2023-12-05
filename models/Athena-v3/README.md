---
license: cc-by-nc-4.0
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/630dfb008df86f1e5becadc3/LjO8no5EzagA9qWdtYKxG.png)

Experimental Athena v3 model. Use Alpaca format. Suitable for RP, ERP and general stuff.

<!-- description start -->
## Description

<!-- [Recommended settings - contributed by localfultonextractor](https://files.catbox.moe/ue0tja.json) -->

This repo contains fp16 files of Athena-V3.

[GGUF - By TheBloke](https://huggingface.co/TheBloke/Athena-v3-GGUF)

[GPTQ - By TheBloke](https://huggingface.co/TheBloke/Athena-v3-GPTQ)

<!-- [exl2 - by AzureBlack](https://huggingface.co/AzureBlack/Athena-v2-6.0bit-exl2) -->

[AWQ - By TheBloke](https://huggingface.co/TheBloke/Athena-v3-AWQ)

[fp16 - by IkariDev+Undi95](https://huggingface.co/IkariDev/Athena-v3)

<!-- [GGUF - by IkariDev+Undi95](https://huggingface.co/IkariDev/Athena-v3-GGUF) -->
[OLD(GGUF - by IkariDev+Undi95)](https://huggingface.co/IkariDev/Athena-v3-GGUF)

## Ratings:

Note: I have permission of all users to upload their ratings, i DONT screenshot random reviews without asking if i can put them here!

https://snombler.neocities.org/logs#athenav3

<!-- description end -->
<!-- description start -->
## Models and loras used

- Athena-v2
- migtissera/Synthia-13B-v1.2
- The-Face-Of-Goonery/Huginn-13b-FP16
- PygmalionAI/pygmalion-2-13b
- The-Face-Of-Goonery/LegerDemain-FP16
- chargoddard/storytime-13b
- lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT
- zattio770/120-Days-of-LORA-v2-13B
```
Loras: [lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT(0.65) + zattio770/120-Days-of-LORA-v2-13B(0.35)](0.3) to the final model

+ [Athena-v2(0.70) + migtissera/Synthia-13B-v1.2(0.3)](0.5)
+ [The-Face-Of-Goonery/Huginn-13b-FP16(0.85) + PygmalionAI/pygmalion-2-13b](0.15)](0.40)
+ [The-Face-Of-Goonery/LegerDemain-FP16(0.3) chargoddard/storytime-13b(0.7)](0.10)
```
<!-- description end -->
<!-- prompt-template start -->
## Prompt template: Alpaca

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:

```

HUGE thanks to [Undi95](https://huggingface.co/Undi95) for doing the merging (Recipe was my idea, he merged)

To TheBloke: please if you quant this, please include [IkariDev](https://huggingface.co/IkariDev) + [Undi95](https://huggingface.co/Undi95) in all the credits/links to the creator.