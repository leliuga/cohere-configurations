---
license: llama2
language:
- en
pipeline_tag: conversational
---
# Goliath 120B

An auto-regressive causal LM created by combining 2x finetuned [Llama-2 70B](https://huggingface.co/meta-llama/llama-2-70b-hf) into one.

Please check out the quantized formats provided by [@TheBloke](https:///huggingface.co/TheBloke) and [@Panchovix](https://huggingface.co/Panchovix):

- [GGUF](https://huggingface.co/TheBloke/goliath-120b-GGUF) (llama.cpp)
- [GPTQ](https://huggingface.co/TheBloke/goliath-120b-GPTQ) (KoboldAI, TGW, Aphrodite)
- [AWQ](https://huggingface.co/TheBloke/goliath-120b-AWQ) (TGW, Aphrodite, vLLM)
- [Exllamav2](https://huggingface.co/Panchovix/goliath-120b-exl2) (TGW, KoboldAI)

# Prompting Format

Both Vicuna and Alpaca will work, but due the initial and final layers belonging primarily to Xwin, I expect Vicuna to work the best.

# Merge process

The models used in the merge are [Xwin](https://huggingface.co/Xwin-LM/Xwin-LM-70B-V0.1) and [Euryale](https://huggingface.co/Sao10K/Euryale-1.3-L2-70B).

The layer ranges used are as follows:

```yaml
- range 0, 16
  Xwin
- range 8, 24
  Euryale
- range 17, 32
  Xwin
- range 25, 40
  Euryale
- range 33, 48
  Xwin
- range 41, 56
  Euryale
- range 49, 64
  Xwin
- range 57, 72
  Euryale
- range 65, 80
  Xwin
```

# Screenshots

![image/png](https://cdn-uploads.huggingface.co/production/uploads/635567189c72a7e742f1419c/Cat8_Rimaz6Ni7YhQiiGB.png)

# Benchmarks
Coming soon.

# Acknowledgements
Credits goes to [@chargoddard](https://huggingface.co/chargoddard) for developing the framework used to merge the model - [mergekit](https://github.com/cg123/mergekit).

Special thanks to [@Undi95](https://huggingface.co/Undi95) for helping with the merge ratios.