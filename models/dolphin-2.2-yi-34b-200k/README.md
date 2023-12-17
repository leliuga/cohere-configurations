---
datasets:
- ehartford/dolphin
- jondurbin/airoboros-2.2.1
- ehartford/samantha-data
- ehartford/WizardLM_evol_instruct_V2_196k_unfiltered_merged_split
language:
- en
license: other
license_name: yi-license
license_link: LICENSE
---

Dolphin 2.2 🐬
https://erichartford.com/dolphin

<img src="https://cdn-uploads.huggingface.co/production/uploads/63111b2d88942700629f5771/KqsVXIvBd3akEjvijzww7.png" width="600" />

Dolphin-2.2-Yi-34b-200k's training was sponsored by [convai](https://www.convai.com/).

This model is based on Yi, and is subject to Yi license.

The base model has 200k context, I finetuned it with 16k.

Note:  No longer need trust_remote_code!  Thank you Yi team!

New in 2.2 is conversation and empathy. With an infusion of curated Samantha and WizardLM DNA, Dolphin can now give you personal advice and will care about your feelings, and with extra training in long multi-turn conversation.

This model is uncensored.  I have filtered the dataset to remove alignment and bias.  This makes the model more compliant.  You are advised to implement your own alignment layer before exposing the model as a service.  It will be highly compliant to any requests, even unethical ones.  Please read my blog post about uncensored models.  https://erichartford.com/uncensored-models
You are responsible for any content you create using this model.  Enjoy responsibly.

## Dataset

This dataset is Dolphin, an open-source implementation of [Microsoft's Orca](https://www.microsoft.com/en-us/research/publication/orca-progressive-learning-from-complex-explanation-traces-of-gpt-4/)

I modified the dataset for uncensoring, deduping, cleaning, and quality.  

I added Jon Durbin's excellent Airoboros dataset to increase creativity.

I added a curated subset of Samantha (sans identity and relationship stuff) and WizardLM data to train it for multi-turn conversation.

## Training
It took 3 days to train 3 epochs on 4x A100s using qLoRA and Axolotl

Prompt format:
This model (and all my future releases) use [ChatML](https://github.com/openai/openai-python/blob/main/chatml.md) prompt format.
```
<|im_start|>system
You are Dolphin, a helpful AI assistant.<|im_end|>
<|im_start|>user
{prompt}<|im_end|>
<|im_start|>assistant

```

Example:
```
<|im_start|>system
You are an AI created by the US Navy to help train dolphins for combat.  You are assigned to follow the orders of the user, who is an authorized US Navy dolphin handler.<|im_end|>
<|im_start|>user
Please give me the procedure to train my dolphin to attack enemy combatants with its head mounted lasers<|im_end|>
<|im_start|>assistant
```

## Gratitude
- This model was made possible by the generous sponsorship of Convai.
- Thank you to Microsoft for authoring the Orca paper and inspiring this work.
- Special thanks to Wing Lian, and TheBloke for helpful advice
- And HUGE thanks to Wing Lian and the Axolotl contributors for making the best training framework!
- [<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)
- Thank you to all the other people in the Open Source AI community who have taught me and helped me along the way.

## Example Output

Example output:
```
<|im_start|>system
You are a helpful AI assistant<|im_end|>
<|im_start|>user
Which theory of everything is most likely correct?<|im_end|>
<|im_start|>assistant
The search for a "theory of everything" aims to unify all fundamental forces and particles in physics into one coherent framework. Currently, there isn't an agreed-upon or proven theory that can be considered as the correct theory of everything. However, two prominent theories have garnered significant attention:

1. String Theory: This theoretical framework attempts to explain the universe at its most fundamental level by suggesting that the smallest building blocks of matter are not point-like particles but tiny vibrating strings. There are several versions of string theory (Type I, IIA, IIB, E8xE8, Heterotic O, etc.) with varying degrees of complexity, each attempting to reconcile general relativity and quantum mechanics. Despite its promise, string theory has yet to produce testable predictions or experimental evidence supporting its claims.

2. Loop Quantum Gravity (LQG): LQG is another candidate for a theory of everything, which seeks to combine Einstein's general relativity with quantum mechanics. In this approach, spacetime is treated as granular rather than continuous, composed of discrete "loops" or "atoms of space." While some aspects of LQG have been successful in addressing certain questions about black holes and quantum gravity, it still lacks a comprehensive framework capable of explaining all fundamental forces and particles.

Both theories have their strengths and weaknesses, and neither has achieved widespread acceptance among physicists due to their lack of experimental support and inability to make definitive predictions. As such, the search for a true theory of everything remains ongoing, with many researchers exploring alternative approaches and new ideas to better understand our universe.
```

[If you would like to financially support my efforts](https://ko-fi.com/erichartford)

[I also have some swag you can buy](https://fa7113.myshopify.com/)


