---
license: other
language:
- en
tags:
- llama
- alpaca
- vicuna
- mix
- merge
- model merge
- roleplay
- chat
- instruct
---

### 13B-HyperMantis

is a weight-sum multi model-merge comprised of:

((MantiCore3E+VicunaCocktail)+(SuperCOT+(StorytellingV2+BluemoonRP))) [All 13B Models]

(GGML and GPTQ are no longer in this repo and will be migrated to a separate repo for easier git download convenience)

Subjective testing shows quality results with KoboldAI (similar results are likely in Text Generation Webui, please disregard KAI-centric settings for that platform); Godlike preset with these tweaks - 2048 context, 800 Output Length, 1.3 Temp, 1.13 Repetition Penalty, AltTextGen:On, AltRepPen:Off, No Prompt Gen:On

Despite being primarily uncensored Vicuna models at its core, HyperMantis seems to respond best to the Alpaca instruct format. Speculatively due to manticore's eclectic instruct datasets generalizing the model's understanding of following instruct formats to some degree. What is known is HyperMantis responds best to the formality of Alpaca's format, whereas Human/Assistant appears to trigger vestigial traces of moralizing and servitude that aren't conducive for roleplay or freeform instructions.

Here is an example of what to place in KAI's Memory (or TGUI's equivalent) to leverage chat as a Roleplay Adventure.
[Define what the role of the named Human/AI are here, let's say our name is 'Player' and we named the AI 'Narrator']

Game Mode:Chat [Remember to name yourself and the AI and reference them in the instruction block]

\#\#\# Instruction:

Make Narrator perform as a text based adventure game with Player as Narrator's user input. Make Narrator describe the scene, scenario, actions of characters, reactions of characters to the player's actions, and potential consequences of their actions and Player's actions when relevant with visually descriptive, detailed, and long storytelling. Allow characters and Player to converse to immerse Player in a rich narrative driven story. When Player encounters a new character, Narrator will name the new character and describe their behavior and appearance. Narrator will internally determine their underlying motivations and weave it into the story where possible.

\#\#\# Response:
[Put A Carriage Return Here]

In KAI, this is why 'No Prompt Gen:On' is important; make your first entry a short writeup of your current situation, or simply reiterate Narrator is a text adventure game and Player is the input. Then your next entry, despite simply being a chat interface, it will kick off what will happen next for Narrator to riff off of. In TGUI, an equivalent setup works the same. Of course, tailor this to whatever you want it to be; instruct models can be as versatile as your imagination. If things go sideways have fun.

Possibly also useful as a regular chatbot, waifu, husbando, TavernAI character, freeform instruct shenanigans, it's whatever. 4bit-128g safetensor [Cuda] included for convenience, might do ggml. Mileage may vary, warranty void if the void stares back.

Credits:

manticore-13b [Epoch3] by openaccess-ai-collective

https://huggingface.co/openaccess-ai-collective/manticore-13b

vicuna-13b-cocktail by reeducator

https://huggingface.co/reeducator/vicuna-13b-cocktail

SuperCOT-LoRA [13B] by kaiokendev

https://huggingface.co/kaiokendev/SuperCOT-LoRA

Storytelling-LLaMa-LoRA [13B, Version 2] by GamerUnTouch

https://huggingface.co/GamerUntouch/Storytelling-LLaMa-LoRAs

bluemoonrp-13b by reeducator

https://huggingface.co/reeducator/bluemoonrp-13b


"Such as gravity's rainbow, sufficiently complex systems stir emergent behavior near imperceptible, uncanny; a Schrodinger's puzzlebox of what may be intrinsic or agentic. Best not to startle what black box phantoms there may be."