---
language:
- eng
tags:
- llama-2
- sft
license:
- mit
datasets:
- LDJnr/Puffin
---

## **Redmond-Puffin-70B**

**Based off Puffin 13B which was the first commercially available language model released by Nous Research!**

Compute provided by PygmalionAI, thank you! Follow PygmalionAI on Twitter @pygmalion_ai.

This is a larger version of Puffin which was originally the worlds first third-party llama-2 fine-tune. leveraging a hand curated set of 3K high quality examples, many of which take full advantage of the 4096 context length of Llama 2. This model was fine-tuned by Nous Research, with LDJ leading the training and dataset curation, along with significant dataset formation contributions by J-Supha.

Special thank you to Pygmalion AI for sponsoring the compute.

Special thank you to Emozilla for assisting with training experimentations and benchmarking.

## Model Training

Redmond-Puffin 70B is a new model trained for multiple epochs on a dataset of 3,000 carefully curated GPT-4 examples, most of which are long context conversations between a real human and GPT-4. 

Additional data came from carefully curated sub sections of datasets such as CamelAI's Physics, Chemistry, Biology and Math.

## Prompt Format

The reccomended model usage is:

WARNING, THE PREVIOUS RECCOMENDATION THAT SAID TO USE "### human" and "# response" WAS A CRITICAL ERROR, PLEASE USE THE ACCURATE PREFIX AND SUFFIX BELOW.

```
USER:

ASSISTANT:
```

## When should I use Puffin or Hermes 2?

Although full benchmarks have not completed for Puffin,
Original Puffin 13B and Hermes-2 13B both beat previous SOTA for GPT4ALL benchmarks, with Hermes-2 winning by a 0.1% margin over Puffin.

Overall, for general purpose zero-shot and/or single turn instructions, Hermes will likely be the way to go. Puffin may be prefferred for creative long conversation interactions, like having Puffin play a character or help brain storm creative ideas or concepts that make contextual sense within an already deep conversation. 

Thank you to the comprehensive analysis and comparison of Puffin and Hermes by reddit user WolframRavenwolf here: https://www.reddit.com/r/LocalLLaMA/comments/158j9r9/nous_hermes_llama2_vs_redmond_puffin_13b/

## Example Outputs!:

![puffin](https://i.imgur.com/P0MsN8B.png)

![puffin](https://i.imgur.com/8EO3ThV.png) 

![puffin](https://i.imgur.com/5IWolFw.png) 

![puffin](https://i.imgur.com/TQui8m7.png) 

![puffin](https://i.imgur.com/tderIfl.png) 

## Notable Features:

 - The first Llama-2 based fine-tuned model released by Nous Research.

 - Ability to recall information upto 2023 without internet (ChatGPT cut off date is in 2021)

 - Pretrained on 2 trillion tokens of text. (This is double the amount of most Open LLM's)

 - Pretrained with a context length of 4096 tokens, and fine-tuned on a significant amount of multi-turn conversations reaching that full token limit.

 - The first commercially available language model released by Nous Research.

## Future Plans

This is a relatively early build amongst the grand plans for the future of Puffin! 

Current limitations: Some token mismatch problems have been identified, these may effect the current output quality, we plan to have this solved in Puffin V2 along with other improvements.

## How you can help!

In the near future we plan on leveraging the help of domain specific expert volunteers to eliminate any mathematically/verifiably incorrect answers from our training curations. 

If you have at-least a bachelors in mathematics, physics, biology or chemistry and would like to volunteer even just 30 minutes of your expertise time, please contact LDJ on discord!

## Benchmarks (New benchmarks coming soon, however here are the 13B benchmarks for now)!

As of Puffins release, it achieves a new SOTA for the GPT4All benchmarks! Supplanting Hermes for the #1 position!
(Rounded to nearest tenth)

Previous Sota: Hermes - 68.8
New Sota:      Puffin - 69.9 (+1.1)
 
Puffin 13B supplants Hermes-2 for the #1 spot in Arc-E, HellaSwag and Winogrande!

Puffin also perfectly ties with Hermes in PIQA, however Hermes-2 still excels in much of Big Bench and AGIEval, so it's highly reccomended you give it a try as well!
