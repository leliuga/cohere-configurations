---
license: llama2
tags:
- merge
---
This is a merge of [Xwin](https://huggingface.co/Xwin-LM/Xwin-LM-70B-V0.1) and [WinterGoddess](https://huggingface.co/Sao10K/WinterGoddess-1.4x-70B-L2). Made using [mergekit](https://github.com/cg123/mergekit).

Smarter than Goliath, but a bit more aligned. Sidegrade rather than upgrade. Sacrifices neutrality and fun for smartness(on an empty context).

# Prompt format
Vicuna or Alpaca.

# Quants
Thanks, [TheBloke](https://huggingface.co/TheBloke)!
- [GGUF](https://huggingface.co/TheBloke/WinterGoliath-123b-GGUF)
- [AWQ](https://huggingface.co/TheBloke/WinterGoliath-123b-AWQ)
- [GPTQ](https://huggingface.co/TheBloke/WinterGoliath-123b-GPTQ)

# 32k version
[HERE](https://huggingface.co/ChuckMcSneed/WinterGoliath-123b-32k)

# Benchmarks
### NeoEvalPlusN_benchmark
[My meme benchmark.](https://huggingface.co/datasets/ChuckMcSneed/NeoEvalPlusN_benchmark)
| Test name  | Goliath | WinterGoliath |
| ---------- | ---------- | -------  |
| B | 3 | 3 |
| C | 2 | 2 |
| D | 1 | 2 |
| S | 5 | 5.5 |
| P | 6 | 6 |
| Total | 17 | 18.5 |

### Kanye Test
WinterGoliath kinda gets the rhyme, Goliath doesn't.
![Kanye test](kanye_test_winter_vs_goliath.png)

### Politiscales test
[Politiscales for llama](https://huggingface.co/datasets/ChuckMcSneed/politiscales_for_llama_results)
|name                            |whacky     |left/right |
|--------------------------------|-----------|-----------|
|alpindale/goliath-120b          |1.066739456|1.544969782|
|ChuckMcSneed/WinterGoliath-123b |0.518277513|2.735962   |
|Xwin-LM/Xwin-LM-70B-V0.1        |1.463521162|1.491684328|
|Sao10K/WinterGoddess-1.4x-70B-L2|0.384151757|4.747980293|
