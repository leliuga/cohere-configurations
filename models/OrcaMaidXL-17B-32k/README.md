---
license: other
license_name: microsoft-research-license
license_link: https://huggingface.co/microsoft/Orca-2-13b/blob/main/LICENSE
pipeline_tag: text-generation
---

# OrcaMaidXL-17B-32k

This is a a slightly experimental frankenmerge of Microsoft's [Orca-2-13b](https://huggingface.co/microsoft/Orca-2-13b) and Undi and IkariDev's [Noromaid-v0.2-13b](https://huggingface.co/NeverSleep/Noromaid-13b-v0.2).

The model recipe is as follows:

```
- [0, 12] from Noromaid (12 layers)
- [12, 22] from Orca (10 layers)
- [14, 24] from Orca (10 layers)
- [18, 28] from Orca (10 layers)
- [28, 40] from Noromaid (12 layers)
```

In my testing so far, the model performs exceptionally well. Your experience may vary.

The prompt format is Alpaca. You can use the standard format as shown, but for best results, you should customize the system prompt to your specific needs.

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{YOUR MESSAGE HERE}

### Response:
{BOT MESSAGE HERE}


```

### Misc. information
- Model size is 17.46B
- BOS token is `<s>`
- EOS token is `</s>`
- Native context length is `32768` via YaRN (original context length was `4096`)
- Base model is Llama 2
- Due to the inclusion of Orca-2-13b, the model is subject to the terms of the [Microsoft Research License](https://huggingface.co/microsoft/Orca-2-13b/blob/main/LICENSE)

### Thanks
- Thanks to [Undi](https://ko-fi.com/undiai) and [IkariDev](https://ikaridevgit.github.io/) for Noromaid