
---
license: llama2
language:
- en
---


**THEBLOKE HAS QUANTS!**
<br>https://huggingface.co/TheBloke/Mythical-Destroyer-L2-13B-GPTQ
<br>https://huggingface.co/TheBloke/Mythical-Destroyer-L2-13B-GGUF



<br>A Merge done for @dampf

**FULL FP16 Model**

<br>Base Model [TheBloke/Llama-2-13B-fp16](https://huggingface.co/TheBloke/Llama-2-13B-fp16)
<br>   **MERGED WITH**
<br>-----[Gryphe/MythoMax-L2-13b](https://huggingface.co/Gryphe/MythoMax-L2-13b)
<br>-----[totally-not-an-llm/PuddleJumper-13b](https://huggingface.co/totally-not-an-llm/PuddleJumper-13b)
<br>-----[TheBloke/Llama-2-13B-Chat-fp16](https://huggingface.co/TheBloke/Llama-2-13B-Chat-fp16)
<br>-----[rombodawg/LosslessMegaCoder-llama2-13b-mini](https://huggingface.co/rombodawg/LosslessMegaCoder-llama2-13b-mini)
<br>-----[The-Face-Of-Goonery/Chronos-Beluga-v2-13bfp16](https://huggingface.co/The-Face-Of-Goonery/Chronos-Beluga-v2-13bfp16)
<br>*using ties-merge*



```
Dampf's Rationale:
if you think about it, the merges kinda act as experts in my destroyer.
mythomax and chronos-beluga for creativity,
llama 2 13b chat and puddlejumper for instruct and losslessmegacoder for logic/code
if this works well...
it should be really, really good
---
mythical destroyer will be used for rp and instruct as well as coding tasks a like
and it should be good at everything
---
```



<br>Script used to Merge [here](https://github.com/cg123/ties-merge)
<br>Thank you for the easy to set up script, [Chargoddard](https://huggingface.co/chargoddard) !

Command:
```
python ties_merge.py TheBloke/Llama-2-13B-fp16 ./Mythical-Destroyer-13B --merge Gryphe/MythoMax-L2-13b --merge totally-not-an-llm/PuddleJumper-13b --merge TheBloke/Llama-2-13B-Chat-fp16 --merge rombodawg/LosslessMegaCoder-llama2-13b-mini --merge The-Face-Of-Goonery/Chronos-Beluga-v2-13bfp16 --cuda
```

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Sao10K__Mythical-Destroyer-L2-13B)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 50.13   |
| ARC (25-shot)         | 58.7          |
| HellaSwag (10-shot)   | 82.0    |
| MMLU (5-shot)         | 57.66         |
| TruthfulQA (0-shot)   | 56.35   |
| Winogrande (5-shot)   | 74.66   |
| GSM8K (5-shot)        | 8.95        |
| DROP (3-shot)         | 12.56         |
