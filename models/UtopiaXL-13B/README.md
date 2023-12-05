---
license: cc-by-nc-4.0
tags:
- not-for-all-audiences
- nsfw
---

<!-- description start -->
## Description

This repo contains fp16 files of UtopiaXL-13B, a merge I have done with the new [layer shuffle](https://github.com/cg123/mergekit/blob/main/mergekit/scripts/layershuffle.py) method from mergekit (thank you Charles for adding my request to your project!)

This is more a proof of concept showing the following:
- Llama2 is very flexible
- Llama2 don't care about what is finetuned on the layers specifically if you keep them in the same order
- Clean merge (no ties, no SLERP, etc...) with only layer is possible without breaking something
- Deleting special tokens/using model with special token don't break the model
- Alpaca win, always. So use it.

The name "XL" come from the absurd amount of model pushed into it.

<!-- description end -->
<!-- description start -->
## Models and loras used

- [Undi95/Utopia-13B](https://huggingface.co/Undi95/Utopia-13B)
- [KoboldAI/LLAMA2-13B-Holodeck-1](https://huggingface.co/KoboldAI/LLAMA2-13B-Holodeck-1)
- [Undi95/PsyMedRP-v1-13B](https://huggingface.co/Undi95/PsyMedRP-v1-13B)
- [PygmalionAI/pygmalion-2-13b](https://huggingface.co/PygmalionAI/pygmalion-2-13b)
- [Heralax/Cat-0.5](https://huggingface.co/Heralax/Cat-0.5)
- [KoboldAI/LLaMA2-13B-TiefighterLR](https://huggingface.co/KoboldAI/LLaMA2-13B-TiefighterLR)
- [Heralax/Augmental-13b-two-epochs](https://huggingface.co/Heralax/Augmental-13b-two-epochs)
- [Undi95/Storytelling-v2.1-13B-lora](https://huggingface.co/Undi95/Storytelling-v2.1-13B-lora)
- [Undi95/LimaRP-UtopiaXL-13B-v3-lora](https://huggingface.co/Undi95/LimaRP-UtopiaXL-13B-v3-lora)

<!-- description end -->
## The sauce
```
!mergekit-layershuffle ./UtopiaXL \
  --model Undi95/Utopia-13B --weight 0.4 \
  --model KoboldAI/LLAMA2-13B-Holodeck-1 --weight 0.1 \
  --model Undi95/PsyMedRP-v1-13B --weight 0.1 \
  --model PygmalionAI/pygmalion-2-13b --weight 0.25 \
  --model Heralax/Cat-0.5 --weight 0.1 \
  --model KoboldAI/LLaMA2-13B-TiefighterLR --weight 0.1 \
  --model Heralax/Augmental-13b-two-epochs --weight 0.1 \
  --write-yaml UtopiaXL.yaml

=========================

merge_method: passthrough
slices:
- sources:
  - layer_range:
    - 0
    - 1
    model: KoboldAI/LLAMA2-13B-Holodeck-1
- sources:
  - layer_range:
    - 1
    - 4
    model: Undi95/Utopia-13B
- sources:
  - layer_range:
    - 4
    - 5
    model: PygmalionAI/pygmalion-2-13b
- sources:
  - layer_range:
    - 5
    - 6
    model: Undi95/PsyMedRP-v1-13B
- sources:
  - layer_range:
    - 6
    - 7
    model: KoboldAI/LLaMA2-13B-TiefighterLR
- sources:
  - layer_range:
    - 7
    - 8
    model: PygmalionAI/pygmalion-2-13b
- sources:
  - layer_range:
    - 8
    - 9
    model: Undi95/PsyMedRP-v1-13B
- sources:
  - layer_range:
    - 9
    - 10
    model: PygmalionAI/pygmalion-2-13b
- sources:
  - layer_range:
    - 10
    - 13
    model: Undi95/Utopia-13B
- sources:
  - layer_range:
    - 13
    - 14
    model: Heralax/Cat-0.5
- sources:
  - layer_range:
    - 14
    - 17
    model: Undi95/Utopia-13B
- sources:
  - layer_range:
    - 17
    - 18
    model: Heralax/Augmental-13b-two-epochs
- sources:
  - layer_range:
    - 18
    - 19
    model: KoboldAI/LLaMA2-13B-TiefighterLR
- sources:
  - layer_range:
    - 19
    - 22
    model: PygmalionAI/pygmalion-2-13b
- sources:
  - layer_range:
    - 22
    - 23
    model: Undi95/Utopia-13B
- sources:
  - layer_range:
    - 23
    - 25
    model: PygmalionAI/pygmalion-2-13b
- sources:
  - layer_range:
    - 25
    - 27
    model: Undi95/Utopia-13B
- sources:
  - layer_range:
    - 27
    - 28
    model: Heralax/Cat-0.5
- sources:
  - layer_range:
    - 28
    - 30
    model: Undi95/Utopia-13B
- sources:
  - layer_range:
    - 30
    - 31
    model: Heralax/Augmental-13b-two-epochs
- sources:
  - layer_range:
    - 31
    - 32
    model: Undi95/Utopia-13B
- sources:
  - layer_range:
    - 32
    - 33
    model: Heralax/Cat-0.5
- sources:
  - layer_range:
    - 33
    - 34
    model: Heralax/Augmental-13b-two-epochs
- sources:
  - layer_range:
    - 34
    - 35
    model: Undi95/PsyMedRP-v1-13B
- sources:
  - layer_range:
    - 35
    - 36
    model: Heralax/Augmental-13b-two-epochs
- sources:
  - layer_range:
    - 36
    - 37
    model: Undi95/Utopia-13B
- sources:
  - layer_range:
    - 37
    - 38
    model: KoboldAI/LLAMA2-13B-Holodeck-1
- sources:
  - layer_range:
    - 38
    - 39
    model: PygmalionAI/pygmalion-2-13b
- sources:
  - layer_range:
    - 39
    - 40
    model: KoboldAI/LLAMA2-13B-Holodeck-1

=========================

=> Applying Undi95/Storytelling-v2.1-13B-lora x 0.1
=> Trained on LimaRP for +2h
=> Applying Undi95/LimaRP-UtopiaXL-13B-v3-lora x 0.35
```
<!-- prompt-template start -->
## Prompt template: Alpaca

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:

```
A big thanks to [Charles](https://huggingface.co/chargoddard) for adding the layer shuffle method to his tool [mergekit](https://github.com/cg123/mergekit/tree/main) and [Henky/KoboldAI](https://koboldai.org/) for the machine he let me use.

If you want to support me, you can [here](https://ko-fi.com/undiai).