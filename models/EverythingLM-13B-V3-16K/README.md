---
license: llama2
datasets:
- totally-not-an-llm/EverythingLM-data-V3
---

# EverythingLM-13b-V3-16k

Introducing EverythingLM, a llama-2 based, general-purpose 13b model with 16k context thanks to LlongMa.  The model is trained on the EverythingLM-V3 dataset, more info can be found on the dataset page.

The model is completely uncensored.

Despite being "uncensored", the base model might be resistant; you might have to prompt-engineer certain prompts.

### Quants (Thanks TheBloke!):
https://huggingface.co/TheBloke/EverythingLM-13B-V3-16K-GGUF

https://huggingface.co/TheBloke/EverythingLM-13B-V3-16K-GPTQ

https://huggingface.co/TheBloke/EverythingLM-13B-V3-16K-AWQ

### Notable features:
- Automatically triggered CoT reasoning.
- Verbose and detailed replies.
- Creative stories.
- Good prompt understanding.

### Differences from V2:
- Much more uncensored.
- Actual roleplaying ability now!
- General all around improvements thanks to the new dataset.  Check out the dataset for more info.

### Prompt format (Alpaca-chat):

```
USER: <prompt>
ASSISTANT:
```

### Future plans:
- Highest priority right now is V3.1 with more optimized training and iterative dataset improvements based on testing.

### Note:
Through testing V2, I realized some alignment data had leaked in, causing the model to be less cooperative then intended.  This model should do much better due to stricter filetering.