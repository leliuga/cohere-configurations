---
license: cc-by-nc-4.0
---
This model speak way more than the v1, be warned.

Command used :
```shell
python ties_merge.py TheBloke/Llama-2-13B-fp16 ./UndiMix-v2-13b  --merge jondurbin/airoboros-l2-13b-2.1 --density 0.10 --merge IkariDev/Athena-v1 --density 0.10 --merge Undi95/UndiMix-v1-13b --density 0.80 --cuda
```

Testing around²...

<!-- description start -->
## Description

This repo contains fp16 files of personal mix : "UndiMix-v2".

It can be hot, serious, playful, and can use emoji thanks to llama-2-13b-chat-limarp-v2-merged.

<!-- description end -->
<!-- description start -->
## Models used


- TheBloke/Llama-2-13B-fp16 (base)
- Undi95/MythoMax-L2-Kimiko-v2-13b (0.33)
- The-Face-Of-Goonery/Huginn-13b-v1.2 (0.33)
- Doctor-Shotgun/llama-2-13b-chat-limarp-v2-merged (0.33)
- ====REMIX====
- jondurbin/airoboros-l2-13b-2.1 (0.10)
- IkariDev/Athena-v1 (0.10)
- UndiMix-v1-13b (0.80)
<!-- description end -->
<!-- prompt-template start -->
## Prompt template: Alpaca

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:

```

Special thanks to Sushi kek