---
license: cc-by-nc-4.0
language:
- en
---

Winter Goddess - A 70B L2 Model for General use, or for Roleplay.

I wanted a Smart Model that is Capable of following Instructions, while being able to (e)RP effectively. Sort of like 1.3, but better.

I merged some models as a base, and had tuned on top of it afterwards.

I personally think this mogs Euryale 1.3, but ymmv.

***

For Transparency's Sake:

Models Used:
<br> Platypus2-70B-instruct
<br> Lila-70B
<br> SunsetBoulevard (at roughly 0.1 weight, boosting coherency)
<br> Private De-alignment LoRA on top.

why does it show mergekit in the safetensors.index metadata? -> I used DARE method to merge the 3 models. Then Axolotl qLoRA. then used lora-merge, copying the files of the base merged model because they didn't save to the new one, only the .safetensor files got saved.

***

Prompt Format - Alpaca


```
### Instruction:
<Prompt>

### Response:

```

OR

```
### Instruction:
<Prompt>

### Input:
<Insert Context Here>

### Response:

```

***


<br> 42. A 25-year-old female has been struck in the right eye with a pipe. She has a ruptured right globe, an orbital fracture and no other obvious injury. You should bandage:


<br> A)	The right eye tightly
<br> B)	Both eyes loosely
<br> C)	The right eye loosely
<br> D)	Both eyes tightly

