---
license: llama2
---
## KEY DETAILS
Prompt format: SillyTavern
Base model: MythoMax-L2-13b
What's new: finetuned on the script of a visual novel that was processed and revamped by GPT-4 to make ~1300 high-quality training examples. The end goal was a model that could speak like a specific character from that game, but the end result was a model that seems to excel in banter, conversation, and roleplay overall.
Note: compared to the original MythoMakise-13b, this model has 33% of MythoMax-L2-13b merged back into it, so that it better retains MythoMax's intelligence with MythoMakise's personality and style. The result of this seems to be pretty good so far. Ironcially, the model seems better at roleplaying characters other than the one it was originally created to mimic. 

### LONG FORM

A finetune of MythoMax-13b on lines extracted from the script of Steins;Gate. Rather than simply giving the model "previous line\nline to predict" a custom script was used to group conversations into training examples. 
Despite being finetuned on one character's lines from one visual novel, I've found (at least in my initial testing) that the model does an excellent job of roleplaying other characters too, probably because the creative writing GPT-4 did on top of the already-well-written Steins;Gate script was very high-quality. The model might be best at roleplaying characters if the personality of that character is similar to the character it was originally made to act like.

Besides being built for RP, I bet that this model could be used in any generic conversational role. Just don't expect it to be accurate, or good at anything other than talking.

The model is not censored.

This variation has MythoMax merged back into it with 33% weighting to make it more stable and intelligent while retaining its Kurisu-ness and better personality. In my experience, this seems to be the decisive change that led to higher-quality outputs.

### Prompt format 
I know it's wasteful as hell, don't judge me, this is the SillyTavern prompt format (discovered using the simple proxy for ST). I finetuned the model on this so that it would perform better on that frontend.
```
## {{charname}}:
- You're "{{charname}}" in this never-ending roleplay with "{{user}}".
### Input:\n
[user description (note, square brackets are a part of it)]
Description of the character's personality would go here (a 'character card')

### Response:
(OOC) Understood. I will take this info into account for the roleplay. (end OOC)

### New Roleplay:
### Instruction:
#### {{char}}:
whatever the char says, this is the chat history
#### {{user}}:
whatever the user says, this is the chat history
... repeated some number of times ...
### Response 2 paragraphs, engaging, natural, authentic, descriptive, creative):
#### {char}:
```
