---
license: other
language:
- en
---
## Model details

MythoLogic-Mini-7b can be considered the little brother in my Mytho series of models: [MythoLogic-13b](https://huggingface.co/Gryphe/MythoLogic-13b) and [MythoBoros-13b](https://huggingface.co/Gryphe/MythoBoros-13b)).

Its Llama-2 core is powered by [Nous Hermes-2](https://huggingface.co/NousResearch/Nous-Hermes-llama-2-7b), which is further augmented by [Stable Beluga](https://huggingface.co/stabilityai/StableBeluga-7B) and a carefully distilled [Kimiko LoRa](https://huggingface.co/nRuaif/Kimiko_7B).

Since 7B models tend to be less capable all-rounders, more emphasis was put on improving the roleplaying aspects for this gradient merge, of which various gradients were benchmarked before settling on the configuration shown below.

![](MythoLogic-Mini-7b.png)

In technical terms, the Hermes-2 core starts at 90% strength before fading away completely at the 12th layer level, where Stable Beluga (and Kimiko) handle the more intricate linguistic aspects.

Quantized models are available from TheBloke: [GGML](https://huggingface.co/TheBloke/MythoLogic-Mini-7b-GGML) - [GPTQ](https://huggingface.co/TheBloke/MythoLogic-Mini-7b-GPTQ) (You're the best!)

## Prompt Format

Due to its Hermes-2 core this model works best with Alpaca formatting, so for optimal model performance, use:
```
<System prompt/Character Card>

### Instruction:
Your instruction or question here.
For roleplay purposes, I suggest the following - Write <CHAR NAME>'s next reply in a chat between <YOUR NAME> and <CHAR NAME>. Write a single reply only.

### Response:
```