---
license: llama2
---

This is an interleaved merge of [Xwin-longLORA-70b-rope8-32k-fp16](https://huggingface.co/grimulkan/Xwin-longLORA-70b-rope8-32k-fp16) and [Euryale-1.3-longLORA-70b-rope8-32k-fp16](https://huggingface.co/grimulkan/Euryale-1.3-longLORA-70b-rope8-32k-fp16), using the same merge formula as alpindale's [goliath-120b](https://huggingface.co/alpindale/goliath-120b).

There is no additional fine-tuning. The resulting model seems to not be broken... you can test whether it is truly the original model + 32K capability (use linear rope scaling 8).

[ChuckMcSneed](https://huggingface.co/ChuckMcSneed) did a benchmark [here](https://huggingface.co/grimulkan/Goliath-longLORA-120b-rope8-32k-fp16/discussions/1), indicating 30% degradation with 8x the context length.

A 6-bit EXL2 quantization is available [here](https://huggingface.co/grimulkan/Goliath-longLORA-120b-rope8-2k-6bpw_h8_exl2). More EXL2 quants [here](https://huggingface.co/aikitoria/Goliath-longLORA-120b-rope8-32k-exl2), thanks to aikitoria.

See [this discussion](https://huggingface.co/grimulkan/aurelian-v0.5-70b-rope8-32K-fp16/discussions/2) for how the original 70B merges were created with longLORA.