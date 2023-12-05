---
license: cc-by-sa-4.0
datasets:
- acrastt/EverythingLM-V3-ShareGPT
language:
- en
library_name: transformers
pipeline_tag: text-generation
---
<a href="https://www.buymeacoffee.com/acrastt" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>

This is [StableLM 3B 4E1T](https://huggingface.co/stabilityai/stablelm-3b-4e1t)(Licensed under [CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/).) finetuned on [EverythingLM Data V3(ShareGPT Format)](https://huggingface.co/datasets/acrastt/EverythingLM-V3-ShareGPT) for 2 epochs using [QLoRA](https://arxiv.org/abs/2305.14314).

Prompt template:
```
### HUMAN:
{prompt}

### RESPONSE:
```

Note that this model have the EOS token of `<|endoftext|>` instead of `<\s>`.

GGUF quantizations available [here](https://huggingface.co/TheBloke/Marx-3B-v3-GGUF).<br/>
GPTQ quantizations available [here](https://huggingface.co/TheBloke/Marx-3B-v3-GPTQ).