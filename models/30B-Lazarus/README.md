---
tags:
- llama
- alpaca
- cot
- vicuna
- uncensored
- merge
- mix
---

## 30B-Lazarus

## Composition:
[] = applied as LoRA to a composite model | () = combined as composite models

[SuperCOT([gtp4xalpaca(manticorechatpygalpha+vicunaunlocked)]+[StoryV2(kaiokendev-SuperHOT-LoRA-prototype30b-8192)])]

This model is the result of an experimental use of LoRAs on language models and model merges that are not the base HuggingFace-format LLaMA model they were intended for.
The desired outcome is to additively apply desired features without paradoxically watering down a model's effective behavior.

Potential limitations - LoRAs applied on top of each other may intercompete.

Subjective results - very promising. Further experimental tests and objective tests are required.

Instruct and Setup Suggestions:

Alpaca instruct is primary, Vicuna instruct format may work.
If using KoboldAI or Text-Generation-WebUI, recommend switching between Godlike and Storywriter presets and adjusting output length + instructions in memory.
Other presets as well as custom settings can yield highly different results, especially Temperature.
If poking it with a stick doesn't work try poking harder.

## Language Models and LoRAs Used Credits:

manticore-30b-chat-pyg-alpha [Epoch0.4] by openaccess-ai-collective

https://huggingface.co/openaccess-ai-collective/manticore-30b-chat-pyg-alpha

SuperCOT-LoRA [30B] by kaiokendev

https://huggingface.co/kaiokendev/SuperCOT-LoRA

Storytelling-LLaMa-LoRA [30B, Version 2] by GamerUnTouch

https://huggingface.co/GamerUntouch/Storytelling-LLaMa-LoRAs

SuperHOT Prototype [30b 8k ctx] by kaiokendev

https://huggingface.co/kaiokendev/SuperHOT-LoRA-prototype

ChanSung's GPT4-Alpaca-LoRA
https://huggingface.co/chansung/gpt4-alpaca-lora-30b

Neko-Institute-of-Science's Vicuna Unlocked LoRA (Checkpoint 46080)
https://huggingface.co/Neko-Institute-of-Science/VicUnLocked-30b-LoRA

Also thanks to Meta for LLaMA.

Each model and LoRA was hand picked and considered for what it could contribute to this ensemble.
Thanks to each and every one of you for your incredible work developing some of the best things
to come out of this community.