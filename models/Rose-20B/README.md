---
language:
- en
pipeline_tag: text-generation
tags:
- text-generation-inference
- instruct
license: llama2
---
<h1 style="text-align: center">Rose-20B</h1>
<center><img src="https://files.catbox.moe/rze9c9.png" alt="roseimage" width="350" height="350"></center>
<center><i>Image sourced by Shinon</i></center>
<h2 style="text-align: center">Experimental Frankenmerge Model</h2>


## Other Formats
[GGUF](https://huggingface.co/TheBloke/Rose-20B-GGUF)
[GPTQ](https://huggingface.co/TheBloke/Rose-20B-GPTQ)
[AWQ](https://huggingface.co/TheBloke/Rose-20B-AWQ)
[exl2](https://huggingface.co/royallab/Rose-20B-exl2)

## Model Details
A Frankenmerge with [Thorns-13B](https://huggingface.co/CalderaAI/13B-Thorns-l2) by CalderaAI and [Noromaid-13-v0.1.1](https://huggingface.co/NeverSleep/Noromaid-13b-v0.1.1) by NeverSleep (IkariDev and Undi). This recipe was proposed by Trappu and the layer distribution recipe was made by Undi. I thank them for sharing their knowledge with me. This model should be very good at any roleplay scenarios. I called the model "Rose" because it was a fitting name for a "thorny maid". 

The recommended format to use is Alpaca. 
```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:
```

Feel free to share any other prompts that works. This model is very robust. 

**Warning: This model uses significantly more VRAM due to the KV cache increase resulting in more VRAM required for the context window.**

## Justification for its Existence
Potential base model for finetune experiments using our dataset to create Pygmalion-20B. Due to the already high capabilities, adding our dataset will mesh well with how the model performs. 
Potential experimentation with merging with other 20B Frankenmerge models. 

## Model Recipe
```
slices:
  - sources:
    - model: Thorns-13B
      layer_range: [0, 16]
  - sources:
    - model: Noromaid-13B
      layer_range: [8, 24]
  - sources:
    - model: Thorns-13B
      layer_range: [17, 32]
  - sources:
    - model: Noromaid-13B
      layer_range: [25, 40]
merge_method: passthrough
dtype: float16
```
Again, credits to [Undi](https://huggingface.co/Undi95) for the recipe.

## Reception
The model was given to a handful of members in the PygmalionAI Discord community for testing. A strong majority really enjoyed the model with only a couple giving the model a passing grade. Since our community has high standards for roleplaying models, I was surprised at the positive reception. 

## Contact
Send a message to tav (tav) on Discord if you want to talk about the model to me. I'm always open to receive comments.