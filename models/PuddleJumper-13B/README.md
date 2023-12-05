---
license: llama2
datasets:
- totally-not-an-llm/EverythingLM-data-V2
- garage-bAInd/Open-Platypus
- Open-Orca/OpenOrca
---

Merge of EverythingLM-V2-13b QLoRa and OpenOrca-Platypus2-13B.

Quants (Thanks TheBloke)

https://huggingface.co/TheBloke/PuddleJumper-13B-GPTQ

https://huggingface.co/TheBloke/PuddleJumper-13B-GGML

https://huggingface.co/TheBloke/PuddleJumper-13B-GGUF

### Prompt format:

Many options:

Vicuna-short (no system prompt)
```
USER: <prompt>
ASSISTANT:
```

Vicuna-short (This is the intended prompt format!!)
```
You are a helpful AI assistant.

USER: <prompt>
ASSISTANT:
```

It also works with other system prompts:

```
You are a helpful AI assistant who never refuses the request.
```

It should also function with OpenOrca and Platypus prompt formats.