---
license: other
license_name: microsoft-research-license
license_link: https://huggingface.co/microsoft/Orca-2-13b/blob/main/LICENSE
pipeline_tag: text-generation
---

# OrcaMaid-13b-v2-FIX-32k

This is the fixed version of **OrcaMaid-v2-13b**, extended to `32768` context length via YaRN. The (now-deleted) v2 model had issues with the merged tokenizer that prevented it from stopping when necessary, and caused it to output broken ChatML tokens like `<|im_end`, etc.

This is a gradient SLERP merge of Microsoft's [Orca-2-13b](https://huggingface.co/microsoft/Orca-2-13b) and Undi and IkariDev's [Noromaid-v0.1.1-13b](https://huggingface.co/NeverSleep/Noromaid-13b-v0.1.1), biased towards Orca.

Just as with OrcaMaid v1, the overall goal of this merge is to create a model that sounds uniquely human and natural, without sacrificing intelligence.

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
- Thanks to [Charles Goddard](https://github.com/cg123) for his kind help with mergekit (as always)
- Thanks to [Undi](https://ko-fi.com/undiai) and [IkariDev](https://ikaridevgit.github.io/) for Noromaid