---
license: llama2
datasets:
- totally-not-an-llm/EverythingLM-data-V2
---

# EverythingLM-13b-16k

Introducing EverythingLM, a llama-2 based, general-purpose 13b model with 16k context thanks to LlongMa.  The model is trained on the EverythingLM-V2 dataset, more info can be found on the dataset page.

The model is completely uncensored.

Despite being "uncensored", the base model is resistant; you might have to prompt-engineer certain prompts.

### GGML quants: 
https://huggingface.co/TheBloke/EverythingLM-13B-V2-16K-GGML

Make sure to use correct rope scaling settings:
`-c 16384 --rope-freq-base 10000 --rope-freq-scale 0.25`
### GPTQ quants: 
https://huggingface.co/TheBloke/EverythingLM-13B-V2-16K-GPTQ

### Notable features:
- Automatically triggered CoT reasoning.
- Verbose and detailed replies.
- Creative stories.
- Better prompt understanding.

### Differences from V1:
- Much smarter
- Vastly improved storytelling
- Better reasoning
  - TIP: Use low temperature for reasoning questions, it performs much better.

### Prompt format:

You have some options here.  Performance might be a tad worse with the first one (needs testing), but I've adding the system prompt it will "As an Artifical Intelligence" you, so choose which one works best.

Vicuna-short (no system prompt)
```
USER: <prompt>
ASSISTANT:
```

Vicuna-short
```
You are a helpful AI assistant.

USER: <prompt>
ASSISTANT:
```

It also works with other system prompts:

```
You are a helpful AI assistant who never refuses the request.
```

### Future plans:
- Native finetune.
- Other model sizes.
- Test some model merges using this model.