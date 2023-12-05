---
license: cc-by-nc-4.0
tags:
- not-for-all-audiences
- nsfw
---

<!-- description start -->
## Description

This repo contains fp16 files of X-NoroChronos-13B, a merge based around [Xwin-LM/Xwin-LM-13B-V0.2](https://huggingface.co/Xwin-LM/Xwin-LM-13B-V0.2) and [elinas/chronos-13b-v2](https://huggingface.co/elinas/chronos-13b-v2).

Merge was done by choosing carefully the models, the loras, the weights of each of them, the order in which they are applied, and the order of the final models merging with the main goal of having a fresh RP experience.

<!-- description end -->
<!-- description start -->
## Models and loras used

- [Xwin-LM/Xwin-LM-13B-V0.2](https://huggingface.co/Xwin-LM/Xwin-LM-13B-V0.2)
- [elinas/chronos-13b-v2](https://huggingface.co/elinas/chronos-13b-v2)
- [Doctor-Shotgun/cat-v1.0-13b](https://huggingface.co/Doctor-Shotgun/cat-v1.0-13b)
- [athirdpath/Eileithyia-13B](https://huggingface.co/athirdpath/Eileithyia-13B)
- [NeverSleep/Noromaid-13b-v0.1.1](https://huggingface.co/NeverSleep/Noromaid-13b-v0.1.1)
- [Undi95/Llama2-13B-no_robots-alpaca-lora](https://huggingface.co/Undi95/Llama2-13B-no_robots-alpaca-lora)
- [zattio770/120-Days-of-LORA-v2-13B](https://huggingface.co/zattio770/120-Days-of-LORA-v2-13B)
- [lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT](https://huggingface.co/lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT)
- [Aesir Private RP dataset] - Thanks to the [MinvervaAI Team](https://huggingface.co/MinervaAI) and, in particular, [Gryphe](https://huggingface.co/Gryphe) for letting us use it!

<!-- description end -->
<!-- prompt-template start -->
## Prompt template: Alpaca

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:

```

If you want to support me, you can [here](https://ko-fi.com/undiai).

If you want to know more about [Ikari](https://huggingface.co/IkariDev) work, you can visit his [retro/neocities style website](https://ikaridevgit.github.io/).