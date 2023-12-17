---
base_model: HuggingFaceH4/mistral-7b-sft-beta
inference: false
license: mit
license_name: mit
metrics:
- accuracy
---

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)

It works well with long system prompts.

It isn't generic in a sense that it shouldn't be used for story telling, for example, but only for reasoning and text comprehension.

This model is trained on a private dataset. The high GSM8K score is **NOT** because of the MetaMath dataset. 

# Prompt Format:

```
### System:
{system_message}
### Instruction:
{prompt}
### Response:
```