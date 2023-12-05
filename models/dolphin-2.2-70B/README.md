---
datasets:
- ehartford/dolphin
- jondurbin/airoboros-2.2.1
- ehartford/samantha-data
- ehartford/WizardLM_evol_instruct_V2_196k_unfiltered_merged_split
language:
- en
license: llama2
---

Dolphin 2.2 🐬
https://erichartford.com/dolphin

<img src="https://cdn-uploads.huggingface.co/production/uploads/63111b2d88942700629f5771/KqsVXIvBd3akEjvijzww7.png" width="600" />

Dolphin-2.2-70b's training was sponsored by [a16z](https://a16z.com/supporting-the-open-source-ai-community/).

This model is based on llama2, so it is suitable for commercial or non-commercial use.

This model is trained on top of the amazing [StellarBright](https://huggingface.co/sequelbox/StellarBright) base model.

New in 2.2 is conversation and empathy. With an infusion of curated Samantha and WizardLM DNA, Dolphin can now give you personal advice and will care about your feelings, and with extra training in long multi-turn conversation.

This model is uncensored.  I have filtered the dataset to remove alignment and bias.  This makes the model more compliant.  You are advised to implement your own alignment layer before exposing the model as a service.  It will be highly compliant to any requests, even unethical ones.  Please read my blog post about uncensored models.  https://erichartford.com/uncensored-models
You are responsible for any content you create using this model.  Enjoy responsibly.

## Dataset

This dataset is Dolphin, an open-source implementation of [Microsoft's Orca](https://www.microsoft.com/en-us/research/publication/orca-progressive-learning-from-complex-explanation-traces-of-gpt-4/)

I modified the dataset for uncensoring, deduping, cleaning, and quality.  

I added Jon Durbin's excellent Airoboros dataset to increase creativity.

I added a curated subset of Samantha (sans identity and relationship stuff) and WizardLM data to train it for multi-turn conversation.

## Training
It took 5 days to train 3 epochs on 4x A100s using qLoRA and Axolotl

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
- This model was made possible by the generous sponsorship of a16z.
- Thank you to Microsoft for authoring the Orca paper and inspiring this work.
- Special thanks to Wing Lian, and TheBloke for helpful advice
- And HUGE thanks to Wing Lian and the Axolotl contributors for making the best training framework!
- [<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)
- Thank you to all the other people in the Open Source AI community who have taught me and helped me along the way.

## Example Output


![image/png](https://cdn-uploads.huggingface.co/production/uploads/63111b2d88942700629f5771/wTbiK4egnUMjgpGG_GHWD.png)

[Buy me a coffee](https://www.buymeacoffee.com/ehartford)