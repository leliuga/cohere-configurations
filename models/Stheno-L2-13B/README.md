---
license: llama2
language:
- en
---

<img src="https://w.forfun.com/fetch/cb/cba2205390e517bea1ea60ca0b491af4.jpeg"  style="width: 70%; min-width: 300px; display: block; margin: auto;">

An experimental merging of Several Models using two various methods, [Ties-Merge](https://github.com/cg123/ties-merge) and [BlockMerge_Gradient](https://github.com/Gryphe/BlockMerge_Gradient)

I plan for this to be the base of my Model with my own [Stheno: ERP-Based LORA] merged in, some time in the future.

Stheno:
<br>Gradient Merge of Stheno-P1 & Stheno-P2.

SISTER MODEL HERE: [Stheno-Inverted-L2-13B](https://huggingface.co/Sao10K/Stheno-Inverted-L2-13B)


Quants courtesy of TheBloke!
<br>[GPTQ](https://huggingface.co/TheBloke/Stheno-L2-13B-GPTQ)
<br>[GGUF](https://huggingface.co/TheBloke/Stheno-L2-13B-GGUF)
<br>[GGML](https://huggingface.co/TheBloke/Stheno-L2-13B-GGML)

Test Checklist:
<br>Censorship - Fairly Uncensored
<br>Writing - Good Prose, Fairly Descriptive
<br>NSFW - Yes
<br>IQ Level - Pretty Smart
<br>Formatting - Proper Formatting with Examples


Stheno-P1 [Ties-Merge]
<br>-----[elinas/chronos-13b-v2](https://huggingface.co/elinas/chronos-13b-v2)
<br>-----[jondurbin/airoboros-l2-13b-2.1](https://huggingface.co/jondurbin/airoboros-l2-13b-2.1)
<br>-----[NousResearch/Nous-Hermes-Llama2-13b](https://huggingface.co/NousResearch/Nous-Hermes-Llama2-13b)+[nRuaif/Kimiko-v2 **LORA**](https://huggingface.co/nRuaif/Kimiko-v2-13B)

Stheno-P2 [Ties-Merge]
<br>-----[CalderaAI/13B-Legerdemain-L2](https://huggingface.co/CalderaAI/13B-Legerdemain-L2)+[lemonilia/limarp-llama2-v2 **LORA**](https://huggingface.co/lemonilia/limarp-llama2-v2)
<br>-----[ehartford/WizardLM-1.0-Uncensored-Llama2-13b](https://huggingface.co/ehartford/WizardLM-1.0-Uncensored-Llama2-13b)
<br>-----[Henk717/spring-dragon](https://huggingface.co/Henk717/spring-dragon)

Most formats could work, but my tests have all been done in Alpaca format and it works well.
```
### Instruction:
Your instruction or question here.
For roleplay purposes, I suggest the following - Write <CHAR NAME>'s next reply in a chat between <YOUR NAME> and <CHAR NAME>. Write a single reply only.

### Response:
```

Below is the Illustration for the Final Merge:

![ILLUSTRATION](https://cdn-uploads.huggingface.co/production/uploads/64be6a5376a6e2efccc638c1/z4D6eun_5ee-k5Bnf0a0j.png)

Once Again, thanks to [Chargoddard](https://huggingface.co/chargoddard) for his amazing and simple [ties-merge](https://github.com/cg123/ties-merge) script, and [Gryphe](https://huggingface.co/Gryphe) for their great [BlockMerge_Gradient](https://github.com/Gryphe/BlockMerge_Gradient) script.
Thanks to the original model creators too!

support me [here](https://ko-fi.com/sao10k) :)

```
Art by wada_kazu / わだかず (pixiv page private?)
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Sao10K__Stheno-L2-13B)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 53.48   |
| ARC (25-shot)         | 61.01          |
| HellaSwag (10-shot)   | 83.95    |
| MMLU (5-shot)         | 56.33         |
| TruthfulQA (0-shot)   | 50.18   |
| Winogrande (5-shot)   | 75.14   |
| GSM8K (5-shot)        | 11.98        |
| DROP (3-shot)         | 35.76         |
