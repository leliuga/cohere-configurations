---
license: cc-by-nc-4.0
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/630dfb008df86f1e5becadc3/y9gdW2923RkORUxejcLVL.png)

Experimental Athena v2 model. Use Alpaca format.

<!-- description start -->
## Description

[Recommended settings - contributed by localfultonextractor](https://files.catbox.moe/ue0tja.json)

This repo contains fp16 files of Athena-V2.

[GGUF - By TheBloke](https://huggingface.co/TheBloke/Athena-v2-GGUF)

[GPTQ - By TheBloke](https://huggingface.co/TheBloke/Athena-v2-GPTQ)

[exl2 - by AzureBlack](https://huggingface.co/AzureBlack/Athena-v2-6.0bit-exl2)

[AWQ - By TheBloke](https://huggingface.co/TheBloke/Athena-v2-AWQ)


[OLD(GGUF - by IkariDev+Undi95)](https://huggingface.co/IkariDev/Athena-v2-GGUF)

## Ratings:

Note: I have permission of all users to upload their ratings, i DONT screenshot random reviews without asking if i can put them here!

![image/png](https://cdn-uploads.huggingface.co/production/uploads/630dfb008df86f1e5becadc3/bqxK6tRyfiJNRakdj4qaD.png)

![image/png](https://cdn-uploads.huggingface.co/production/uploads/630dfb008df86f1e5becadc3/LRNRctxnN4BDskN8dPbmY.png)

![image/png](https://cdn-uploads.huggingface.co/production/uploads/630dfb008df86f1e5becadc3/TCPXp_IJm97TGgqgYgPsl.png)

![image/png](https://cdn-uploads.huggingface.co/production/uploads/630dfb008df86f1e5becadc3/5scjmDP1V2fofqCYNMmbT.png)

![image/png](https://cdn-uploads.huggingface.co/production/uploads/630dfb008df86f1e5becadc3/RysGRIOn2gZ1R4x4JoAlr.png)

<!-- description end -->
<!-- description start -->
## Models and loras used

- Xwin-LM/Xwin-LM-13B-V0.1
- Undi95/ReMM-v2.2-L2-13B
- Undi95/MLewd-L2-13B-v2-3
- Brouz/Slerpeno
- boomerchan/Magpie-13b
```
[Xwin (0.30) + ReMM v2.2 (0.70)](0.45) x [[Xwin (0.40) + MLewd v2-3 (0.60)](0.80) + [Slerpeno(0.50) + Magpie-13b(0.50)](0.20)](0.55)
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