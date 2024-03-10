---
license: cc-by-nc-4.0
datasets:
- NobodyExistsOnTheInternet/full120k
base model: mistralai/Mixtral-8x7B-v0.1
---

Trained using a randomised subset of Full120k - 60K Samples [Roughly 50M Tokens] + More of my own NSFW Instruct & De-Alignment Data [Roughly 30M Tokens Total]
<br>Total Tokens used for Training: 80M over 1 epoch, over 2xA100s at batch size 5, grad 5 for 12 hours.

***

Experimental model, trained on [mistralai/Mixtral-8x7B-v0.1](https://huggingface.co/mistralai/Mixtral-8x7B-v0.1) using [Charles Goddard](https://huggingface.co/chargoddard)'s ZLoss and Megablocks-based fork of transformers.

***

Trained with Alpaca format.

```
### Instruction:
<Prompt>

### Input:
<Insert Context Here>

### Response:

```
***

Useful prompt guide: https://rentry.org/mixtralforretards

useful stopping strings: 
```
["\nInput:", "\n[", "\n(", "\n### Input:"] 
```

*stops run-off generations after response, important for alpaca*

***

Roleplay based model, specifically the ERP type one.

I mean, its pretty good sometimes? I had various testing versions of Mistral 7B and L2 70B, L2 13B, and even Solar with the same dataset and various learning rates, they did much better. MoE tuning kinda meh still.

about gptisms. It's weird. with certain prompts its never there, with some its there. despite the prose of full120k, I never encountered gptslop with mistral, solar or l2 based trains which was why I was confident about this being good initially.

Mixtral is really finicky. with the right settings this model can shine. I recommend Universal-Light or Universal-Creative in SillyTavern.

Anyways... Enjoy?