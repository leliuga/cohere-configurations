---
license: other
license_name: microsoft-research-license
license_link: https://huggingface.co/microsoft/Orca-2-13b/blob/main/LICENSE
pipeline_tag: text-generation
---

# Nororcetacean-20b-10k

This is Jeb Carter's [Psyonic-Cetacean-20B](https://huggingface.co/jebcarter/psyonic-cetacean-20B), merged with Undi's [no_robots-alpaca LoRA](https://huggingface.co/Undi95/Llama2-13B-no_robots-alpaca-lora) and extended to `10240` context length via YaRN.

The overall goal of this merge was to create a model with the unique brain of Psyonic-Cetacean and the human voice of the no_robots dataset, that would remain capable at long contexts.

The prompt format is Alpaca. You can use the standard format as shown, but for best results, I strongly recommend customizing the system prompt to your specific needs.

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{YOUR MESSAGE HERE}

### Response:
{BOT MESSAGE HERE}


```

### Misc. information
- BOS token is `<s>`
- EOS token is `</s>`
- Native context length is `10240` via YaRN (original context length was `4096`)
- Base model is Llama 2
- Due to the inclusion of Orca-2-13b, the model is subject to the terms of the [Microsoft Research License](https://huggingface.co/microsoft/Orca-2-13b/blob/main/LICENSE)

### Thanks
- Thanks to [Jeb Carter](https://huggingface.co/jebcarter) for Psyonic-Cetacean-20B
- Thanks to [Undi](https://ko-fi.com/undiai) for the no_robots-alpaca LoRA
