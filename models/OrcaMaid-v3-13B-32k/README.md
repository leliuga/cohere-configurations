---
license: other
tags:
- merge
license_name: microsoft-research-license
license_link: https://huggingface.co/microsoft/Orca-2-13b/blob/main/LICENSE
pipeline_tag: text-generation
---

# OrcaMaid-v3-13b-32k

This is the third version of OrcaMaid, a weighted gradient SLERP merge between Microsoft's [Orca-2-13b](https://huggingface.co/microsoft/Orca-2-13b) and NeverSleep's [Noromaid-13b-v0.3](https://huggingface.co/NeverSleep/Noromaid-13b-v0.3).

The goal of this merge is to create an unusually intelligent and human-like model especially for RP.

The prompt format is Alpaca. You can use the standard format as shown, but for best results, you should customize the system prompt to your specific needs.

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
- Native context length is `32768` via YaRN (original context length was `4096`)
- Base model is Llama 2
- Due to the inclusion of Orca-2-13b, the model is subject to the terms of the [Microsoft Research License](https://huggingface.co/microsoft/Orca-2-13b/blob/main/LICENSE)

### Thanks
- Thanks to [Undi](https://ko-fi.com/undiai) and [IkariDev](https://ikaridevgit.github.io/) of [NeverSleep](https://huggingface.co/NeverSleep) for Noromaid