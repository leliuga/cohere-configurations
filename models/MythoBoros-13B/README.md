---
license: other
language:
- en
---
## Model details

MythoBoros-13b can be considered a sister model to [MythoLogic-13b](https://huggingface.co/Gryphe/MythoLogic-13b), sharing the same goals but having a different approach.

Whereas the previous model was a series of experimental gradient merges, this one is a simple straight-up 66/34 merge of [Chronos](https://huggingface.co/elinas/chronos-13b) and the freshly released [Ouroboros](https://huggingface.co/CalderaAI/13B-Ouroboros), providing a very solid foundation for a well-performing roleplaying model.

MythoBoros tends to be somewhat more formal with its responses in comparison to MythoLogic.

My advice? Try both, see which one you prefer.

Quantized models are available from TheBloke: [GGML](https://huggingface.co/TheBloke/MythoBoros-13B-GGML) - [GPTQ](https://huggingface.co/TheBloke/MythoBoros-13B-GPTQ) (You're the best!)

## Prompt Format

This model works best with Alpaca formatting, so for optimal model performance, use:
```
<System prompt/Character Card>

### Instruction:
Your instruction or question here.
For roleplay purposes, I suggest the following - Write <CHAR NAME>'s next reply in a chat between <YOUR NAME> and <CHAR NAME>. Write a single reply only.

### Response:
```