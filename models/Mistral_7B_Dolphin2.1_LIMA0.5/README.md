---
license: mit
---

ehartford's merge of Mistral 7B 0.1 with his Dolphin 2.1 dataset

https://huggingface.co/ehartford/dolphin-2.1-mistral-7b

and

LIMA RP dataset applied as a lora at 0.5 weight

https://huggingface.co/lemonilia/limarp-llama2-v2/

Purpose of the model is to be RP-focused, smart, fast, and lightweight for users with low VRAM.

I've already built the exl2 4bpw quant (linked below), and it will run 8k ctx at around 6GB VRAM and respond to a full context at roughly 30tps (tested on my 3060) if exl2_hf loader is used with FA2 enabled. 

Model has been tested by several users on the SillyTavern discord server, and run on Horde for a full day - with good results.

https://huggingface.co/RossAscends/Mistral7B_Dolphin2.1_LIMARP0.5_4bpw_exl2

Mistral or ChatML context presets both possible.

exllama v2 4bpw quant: https://huggingface.co/RossAscends/Mistral7B_Dolphin2.1_LIMARP0.5_4bpw_exl2
