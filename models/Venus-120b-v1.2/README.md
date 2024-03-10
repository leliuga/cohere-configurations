---
license: llama2
language:
  - en
tags:
  - not-for-all-audiences
---

# Venus 120b - version 1.2

![image/png](https://cdn-uploads.huggingface.co/production/uploads/655febd724e0d359c1f21096/BSKlxWQSbh-liU8kGz4fF.png)

## Overview

Version 1.2 of the Venus 120b lineup.

## Model Details

- A result of interleaving layers of [lizpreciatior/lzlv_70b_fp16_hf](https://huggingface.co/lizpreciatior/lzlv_70b_fp16_hf) with itself.
- The resulting model has 140 layers and approximately 122 billion parameters.
- See mergekit-config.yml for details on the merge method used.
- See the `exl2-*` branches for exllama2 quantizations. The 4.85 bpw quant should fit in 80GB VRAM, and the 3.0 bpw quant should (just barely) fit in 48GB VRAM with 4k context.
- Inspired by [Goliath-120b](https://huggingface.co/alpindale/goliath-120b)

**Warning: This model will produce NSFW content!**

## Results

Better at following instructions than both v1.1 and v1.0, and doesn't seem to suffer from censorship issues. Overall I like this version the most out of all the models I've created.