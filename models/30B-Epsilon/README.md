---
tags:
- llama
- alpaca
- vicuna
- uncensored
- cot
- chain of thought
- story
- adventure
- roleplay
- rp
- merge
- mix
- instruct
- wizardlm
- superhot
- supercot
- manticore
- hippogriff
---

## 30B-Epsilon

Epsilon is an instruct based general purpose model assembled from hand picked models and LoRAs.
There is no censorship and it follows instructions in the Alpaca format. This means you can create
your own rules in the context memory of your inference system of choice [mainly KoboldAI or Text
Generation Webui and chat UIs like SillyTavern and so on].

## Composition:

This model is the result of an experimental use of LoRAs on language models and model merges.
[] = applied as LoRA to a composite model | () = combined as composite models
30B-Epsilon = [SuperCOT[SuperHOT-prototype13b-8192[(wizardlmuncensored+((hippogriff+manticore)+(StoryV2))]

Alpaca's instruct format can be used to do many things, including control of the terms of behavior
between a user and a response from an agent in chat. Below is an example of a command injected into
memory.

```
### Instruction:
Make Narrator function as a text based adventure game that responds with verbose, detailed, and creative descriptions of what happens next after Player's response.
Make Player function as the player input for Narrator's text based adventure game, controlling a character named (insert character name here, their short bio, and
whatever quest or other information to keep consistent in the interaction).

### Response:
{an empty new line here}
```

All datasets from all models and LoRAs used were documented and reviewed as model candidates for merging.
Model candidates were based on five core principles: creativity, logic, inference, instruction following,
and longevity of trained responses. SuperHOT-prototype30b-8192 was used in this mix, not the 8K version;
the prototype LoRA seems to have been removed [from HF] as of this writing. The GPT4Alpaca LoRA from
Chansung was removed from this amalgam following a thorough review of where censorship and railroading
the user came from in 33B-Lazarus. This is not a reflection of ChanSung's excellent work - it merely did
not fit the purpose of this model.

## Language Models and LoRAs Used Credits:

manticore-30b-chat-pyg-alpha [Epoch0.4] by openaccess-ai-collective

https://huggingface.co/openaccess-ai-collective/manticore-30b-chat-pyg-alpha

hippogriff-30b-chat by openaccess-ai-collective

https://huggingface.co/openaccess-ai-collective/hippogriff-30b-chat

WizardLM-33B-V1.0-Uncensored by ehartford

https://huggingface.co/ehartford/WizardLM-33B-V1.0-Uncensored

Storytelling-LLaMa-LoRA [30B, Version 2] by GamerUnTouch

https://huggingface.co/GamerUntouch/Storytelling-LLaMa-LoRAs

SuperCOT-LoRA [30B] by kaiokendev

https://huggingface.co/kaiokendev/SuperCOT-LoRA

SuperHOT-LoRA-prototype30b-8192 [30b, not 8K version, but a removed prototype] by kaiokendev

https://huggingface.co/kaiokendev/superhot-30b-8k-no-rlhf-test [Similar LoRA to one since removed that was used in making this model.]

Also thanks to Meta for LLaMA and to each and every one of you
who developed these fine-tunes and LoRAs.