---
license: other
language:
- en
tags:
- text-generation-inference
---
# RpBird-Yi-34B-200k 
## A Mythomax/MLewd_13B-style merge of selected 34B models

A multi-model merge of several  Yi 34B finetunes for roleplaying and creative work. The goal was to create a model that combines creativity with intelligence for an enhanced experience.

Did it work? Probably, maybe. It seemed subjectively better than each of the individual models in my tests.

As suggested by other users, I have renamed it to RpBird in order to reduce misunderstandings.

## Procedure:

Models used:
- **ehartford/dolphin-2_2-yi-34b**
- **larryvrh/Yi-34B-200K-Llamafied**
- **NousResearch/Nous-Capybara-34B**
- **migtissera/Tess-M-Creative-v1.0**
- **LoneStriker/Yi-34B-Spicyboros-3.1**
- **Doctor-Shotgun/limarpv3-yi-llama-34b-lora**



# Prompt Format:

```
SYSTEM: <ANY SYSTEM CONTEXT>
USER: What is the relationship between Earth's atmosphere, magnetic field and gravity?
ASSISTANT:
```