---
license: llama2
language:
- en
---

*Updated readme latest update on 14/11*

![Euryale](https://huggingface.co/Sao10K/Euryale-1.3-L2-70B/resolve/main/Euryale.jpg)

17th Attempt. Past 10 Failed, cost me >$200 lol.

Idea is an updated version of Euryale with ReMantik (Mythologic Recreation in 70B + Sauce) instead of the ties-merge between the original 3 models.

This is then mixed with a saucy model (spicyboros+pyg_lora) with a Mythomax-esque Ratio, and a certain experimental (self) LoRA applied to it.

***Unrelated to Euryale 1.0, I got rid of most of the merged models as they were bloat.***

***

1.4 Here: https://huggingface.co/Sao10K/Euryale-1.4-L2-70B
<br> Subjectively better? atleast it is a minor improvement in my eyes.

***
Test Results: Works Well.
<br>NSFL and NSFW fine in roleplay context.
<br>slight censor with 0 context, zero issues in actual RP / ERP.
<br>Good Prose, Not Dumbed Down due to RP merges from testing.
<br> I have not encountered any repetition issues some had with the original Euryale. tell me if you do, though.

Prompt and System Format:
most works well. I recommend Alpaca.
ST Settings used for Test:
Lightning 1.1 System Prompt + Shortwave(1.2 Temperature)

Support me [here](https://ko-fi.com/sao10k) :)


Quants done by TheBloke! Ty a lot to him.

https://huggingface.co/TheBloke/Euryale-1.3-L2-70B-GPTQ
https://huggingface.co/TheBloke/Euryale-1.3-L2-70B-GGUF
https://huggingface.co/TheBloke/Euryale-1.3-L2-70B-AWQ

***

ty for the exl2 quants! there's more bpw out there you can search too!

https://huggingface.co/LoneStriker/Euryale-1.3-L2-70B-2.4bpw-h6-exl2

https://huggingface.co/LoneStriker/Euryale-1.3-L2-70B-2.6bpw-h6-exl2

https://huggingface.co/Panchovix/Euryale-1.3-L2-70B-4.65bpw-h6-exl2

https://huggingface.co/AzureBlack/Euryale-1.3-L2-70B-4.6bpw-6h-exl2

https://huggingface.co/Doctor-Shotgun/Euryale-1.3-limarpv3-L2-70B-exl2
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Sao10K__Euryale-1.3-L2-70B)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 66.58   |
| ARC (25-shot)         | 70.82          |
| HellaSwag (10-shot)   | 87.92    |
| MMLU (5-shot)         | 70.39         |
| TruthfulQA (0-shot)   | 59.85   |
| Winogrande (5-shot)   | 82.79   |
| GSM8K (5-shot)        | 34.19        |
| DROP (3-shot)         | 60.1         |
