---
base_model: migtissera/Tess-34B-v1.4
inference: false
license: other
license_name: yi-license
license_link: https://huggingface.co/01-ai/Yi-34B/blob/main/LICENSE
metrics:
- accuracy
---

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)

An instruct based fine tune of [migtissera/Tess-34B-v1.4](https://huggingface.co/migtissera/Tess-34B-v1.4).

It works well with long system prompts.

It isn't generic in a sense that it shouldn't be used for story telling, for example, but only for reasoning and text comprehension.

This model is trained on a private dataset. The high GSM8K score is **NOT** because of the MetaMath dataset. 

# Prompt Format:

```
SYSTEM: <ANY SYSTEM CONTEXT>
USER: 
ASSISTANT:
```