---
license: llama2
language:
  - en
tags:
  - not-for-all-audiences
---

# Venus 103b - version 1.2

![image/png](https://cdn-uploads.huggingface.co/production/uploads/655febd724e0d359c1f21096/BSKlxWQSbh-liU8kGz4fF.png)

## Model Details

- A result of interleaving layers of [Sao10K/Euryale-1.3-L2-70B](https://huggingface.co/Sao10K/Euryale-1.3-L2-70B) and [GOAT-AI/GOAT-70B-Storytelling](https://huggingface.co/GOAT-AI/GOAT-70B-Storytelling)
- The resulting model has 120 layers and 103 billion parameters.
- See mergekit-config.yml for details on the merge method used.
- See the `exl2-*` branches for exllama2 quantizations. The 5.65 bpw quant should fit in 80GB VRAM, and the 3.35/3.0 bpw quants should fit in 48GB VRAM.
- Inspired by [Goliath-120b](https://huggingface.co/alpindale/goliath-120b)

**Warning: This model will produce NSFW content!**

## Results

1. In my limited testing, I've found this model to be the most creative of the 103b merges I've made so far.
2. Seems to tolerate higher temperatures than the previous Venus models
3. Doesn't seem to suffer from any censorship issues
4. Does not follow instructions as well as v1.1, but still does a bit better than v1.0
5. Has some issues with formatting sometimes (i.e not closing asterisks or quotes)

Note that these are obviously just my personal observations, everyone will have their own unique experience based on their settings and specific scenarios they're using the model for.