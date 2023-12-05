---
license: cc-by-sa-4.0
datasets:
- Norquinal/claude_multiround_chat_1k
language:
- en
library_name: transformers
pipeline_tag: text-generation
---
<a href="https://www.buymeacoffee.com/acrastt" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>

This is [StableLM 3B 4E1T](https://huggingface.co/stabilityai/stablelm-3b-4e1t)(Licensed under [CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/).) instruction tuned on [Claude Multiround Chat 1K](https://huggingface.co/datasets/Norquinal/claude_multiround_chat_1k) for 2 epochs with QLoRA([2305.14314](https://arxiv.org/abs/2305.14314)).

Prompt template:
```
USER: {prompt}
ASSISTANT:
```

GGUF quantizations available [here](https://huggingface.co/TheBloke/Akins-3B-GGUF).<br/>
GPTQ quantizations available [here](https://huggingface.co/TheBloke/Akins-3B-GPTQ).