---
license: cc-by-nc-4.0
---

<!-- description start -->
## Description

This repo contains fp16 files of personal mix : "UndiMix-v3".

It can be hot, serious, playful, and can use emoji thanks to llama-2-13b-chat-limarp-v2-merged.

What change from V2 is the that I didn't use Llama-2-13B-fp16 for the base anymore, and got straight into the merging with SLERP taking ReMM-S-Kimiko-v2-13B as a base.

<!-- description end -->
<!-- description start -->
## Models used

- Undi95/ReMM-S-Kimiko-v2-13B (0.272) (base)
- The-Face-Of-Goonery/Huginn-13b-v1.2 (0.264)
- Doctor-Shotgun/llama-2-13b-chat-limarp-v2-merged (0.264)
- jondurbin/airoboros-l2-13b-2.1 (0.10)
- IkariDev/Athena-v1 (0.10)
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