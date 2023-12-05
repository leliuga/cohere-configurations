
---
license: llama2
language:
- en
---

<br>A Merge done for @dampf

**FULL FP16 Model**


**V2 Model**
<br>Changelog:
<br>REMOVED - Llama-2-13B-Chat-fp16 (reason: censored, likely amplified base model quirks)
<br>ADDED - jondurbin/airoboros-l2-13b-2.1 (ghost attention, improved RP and instruction)

<br>Base Model [TheBloke/Llama-2-13B-fp16](https://huggingface.co/TheBloke/Llama-2-13B-fp16)
<br>   **MERGED WITH**
<br>-----[Gryphe/MythoMax-L2-13b](https://huggingface.co/Gryphe/MythoMax-L2-13b)
<br>-----[totally-not-an-llm/PuddleJumper-13b](https://huggingface.co/totally-not-an-llm/PuddleJumper-13b)
<br>-----[jondurbin/airoboros-l2-13b-2.1](https://huggingface.co/jondurbin/airoboros-l2-13b-2.1)
<br>-----[rombodawg/LosslessMegaCoder-llama2-13b-mini](https://huggingface.co/rombodawg/LosslessMegaCoder-llama2-13b-mini)
<br>-----[The-Face-Of-Goonery/Chronos-Beluga-v2-13bfp16](https://huggingface.co/The-Face-Of-Goonery/Chronos-Beluga-v2-13bfp16)
<br>*using ties-merge*



```
Dampf's Rationale:
I did receive feedback from some users that it likes to add notes and morality to erp stories.
i will kick llama 2 chat and make an uncensored V2 version 
in llama 2 chat's place will be the freshly released airboros 2.1
---
well it was not bad, it was just censored because of llama 2 13b chat
i guess charles was really serious about each model retaining its shape
i was expecting parts of it to get watered down, but judging from the strong influence of llama chat that wasn't the case
```

Alpaca should be its main format, but also can be used with others. Vicuna 1.1 should work well too. 
```
### Instruction:
Your instruction or question here.
For roleplay purposes, I suggest the following - Write <CHAR NAME>'s next reply in a chat between <YOUR NAME> and <CHAR NAME>. Write a single reply only.

### Response:
```

LIMITATIONS:

While some of the issues of V1 have been fixed, there are some issues left that makes the model not very useable in certain scenarios such as roleplaying. The model explains actions and breaks character regularly. 
Update: I've found out this was largely due to SillyTavern's formatting. If you are using SillyTavern, make sure to disable example chats formatting and chat start formatting.

<br>Script used to Merge [here](https://github.com/cg123/ties-merge)
<br>Thank you for the easy to set up script, [Chargoddard](https://huggingface.co/chargoddard). Also I want to thank all these hard working model creators for their contributions to the Open Source community!

Command:
```
python ties_merge.py TheBloke/Llama-2-13B-fp16 ./Mythical-Destroyer-V2-13B --merge Gryphe/MythoMax-L2-13b --merge totally-not-an-llm/PuddleJumper-13b --merge jondurbin/airoboros-l2-13b-2.1 --merge rombodawg/LosslessMegaCoder-llama2-13b-mini --merge The-Face-Of-Goonery/Chronos-Beluga-v2-13bfp16 --cuda
```

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Sao10K__Mythical-Destroyer-V2-L2-13B)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 48.79   |
| ARC (25-shot)         | 59.3          |
| HellaSwag (10-shot)   | 82.66    |
| MMLU (5-shot)         | 57.39         |
| TruthfulQA (0-shot)   | 57.09   |
| Winogrande (5-shot)   | 74.74   |
| GSM8K (5-shot)        | 0.0        |
| DROP (3-shot)         | 10.33         |
