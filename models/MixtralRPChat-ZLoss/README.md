---
license: cc-by-nc-4.0
datasets:
- Open-Orca/SlimOrca
- lemonilia/LimaRP
- chargoddard/rpguild
- chargoddard/summarize_from_feedback_alpaca
- HuggingFaceH4/no_robots
- chargoddard/coedit-reworded
language:
- en
tags:
- mixtral
base_model: mistralai/Mixtral-8x7B-v0.1
---

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)

QLoRA tuned from [mistralai/Mixtral-8x7B-v0.1](https://huggingface.co/mistralai/Mixtral-8x7B-v0.1).

My main reason for training this model was to investigate using an altered router balancing loss combined with the z-loss introduced in [ST-MoE: Designing Stable and Transferable Sparse Expert Models](https://arxiv.org/abs/2202.08906). The result is pretty decent, I think! It does a good job of respecting character information in system prompts and performed adequately on a few simple coding tasks.

To train this I used a custom branch of Transformers that adds z-loss and reimplements the router balancing loss based on the version in [MegaBlocks](https://github.com/stanford-futuredata/megablocks). The config used with my custom hacked-up branch of axolotl is available [here](https://huggingface.co/chargoddard/MixtralRPChat-ZLoss/blob/main/axolotl_config.yml).

Uses my favorite non-ChatML token-economic chat prompt format. Messages should be prefixed with `" ***System:"`, `" ***Query:"`, or `" ***Response:"` for system, user, and model messages respectively. No newlines are necessary but the space before the triple asterisk is mandatory.
