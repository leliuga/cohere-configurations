---
license: llama2
language:
  - en
tags:
  - not-for-all-audiences
---

# Venus 103b - version 1.1

![image/png](https://cdn-uploads.huggingface.co/production/uploads/655febd724e0d359c1f21096/BSKlxWQSbh-liU8kGz4fF.png)

## Model Details

- A result of interleaving layers of [Sao10K/Euryale-1.3-L2-70B](https://huggingface.co/Sao10K/Euryale-1.3-L2-70B), [migtissera/SynthIA-70B-v1.2b](https://huggingface.co/migtissera/SynthIA-70B-v1.2b), and [Xwin-LM/Xwin-LM-70B-V0.1](https://huggingface.co/Xwin-LM/Xwin-LM-70B-V0.1) using [mergekit](https://github.com/cg123/mergekit).
- The resulting model has 120 layers and 103 billion parameters.
- See mergekit-config.yml for details on the merge method used.
- See the `exl2-*` branches for exllama2 quantizations. The 5.65 bpw quant should fit in 80GB VRAM, and the 3.35 bpw quant should fit in 48GB VRAM.
- Inspired by [Goliath-120b](https://huggingface.co/alpindale/goliath-120b)

**Warning: This model will produce NSFW content!**

## Results

1. Seems to be more "talkative" than Venus-103b-v1.0 (i.e characters speakmore often in roleplays)
2. **Is more resistant to nsfw/nsfl stuff than 1.0. If you're looking for totally uncensored roleplays, I'd recommend 1.0 instead.**
3. Sometimes struggles to pay attention to small details in the scenes
4. Prose seems pretty creative and more logical than Venus-120b-v1.0