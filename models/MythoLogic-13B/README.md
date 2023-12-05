---
license: other
language:
- en
---
**UPDATE:** There's a Llama 2 sequel now! [Check it out here!](https://huggingface.co/Gryphe/MythoLogic-L2-13b)

An experiment with gradient merges using [the following script](https://github.com/TehVenomm/LM_Transformers_BlockMerge), with [Chronos](https://huggingface.co/elinas/chronos-13b) as its primary model, augmented by [Hermes](https://huggingface.co/NousResearch/Nous-Hermes-13b) and [Wizard-Vicuna Uncensored](https://huggingface.co/TheBloke/Wizard-Vicuna-13B-Uncensored-HF).

Quantized models are available from TheBloke: [GGML](https://huggingface.co/TheBloke/MythoLogic-13B-GGML) - [GPTQ](https://huggingface.co/TheBloke/MythoLogic-13B-GPTQ) (You're the best!)

## Model details

Chronos is a wonderfully verbose model, though it definitely seems to lack in the logic department. Hermes and WizardLM have been merged gradually, primarily in the higher layers (10+) in an attempt to rectify some of this behaviour.

The main objective was to create an all-round model with improved story generation and roleplaying capabilities.

Below is an illustration to showcase a rough approximation of the gradients I used to create MythoLogic:

![](approximation.png)

## Prompt Format

This model primarily uses Alpaca formatting, so for optimal model performance, use:
```
<System prompt/Character Card>

### Instruction:
Your instruction or question here.
For roleplay purposes, I suggest the following - Write <CHAR NAME>'s next reply in a chat between <YOUR NAME> and <CHAR NAME>. Write a single reply only.

### Response:
```

---
license: other
---