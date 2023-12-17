---
datasets:
- Open-Orca/SlimOrca-Dedup
- teknium/openhermes
- meta-math/MetaMathQA
- migtissera/Synthia-v1.3
- THUDM/AgentInstruct
- LeoLM/German_Songs
- LeoLM/German_Poems
- LeoLM/OpenSchnabeltier
- bjoernp/ultrachat_de
- LDJnr/Capybara
language:
- en
library_name: transformers
pipeline_tag: text-generation
license: llama2
model_creator: DiscoResearch
model_type: llama
tags:
- goliath
- deutsch
- llama2
- discoresearch
---


<img src="imgs/disco_goliath.jpeg" width="600">

# DiscoLM 120b (Alpha)

**DiscoLM 120b (Alpha)** is an experimental 120b model based on [Alpindale´s Goliath 120b](https://huggingface.co/alpindale/goliath-120b), a merge of different Llama2-70b models, and further finetuned on a dataset of some the most popular open-source instruction sets.
Disco 120b is a [DiscoResearch](https://huggingface.co/DiscoResearch) project and was trained by [Björn Plüster](https://huggingface.co/bjoernp).

Many thanks to [LAION](https://laion.ai) and [HessianAI](https://hessian.ai/) for scientific supervision, coordination and compute resources provided for this project on supercomputer 42 by [HessianAI](https://hessian.ai/)! 

<img src="https://hessian.ai/wp-content/themes/hessianai/img/hessian-ai-logo.svg" width="120">
<img src="https://avatars.githubusercontent.com/u/92627801?s=200&v=4" width="120">

## Table of Contents

1. [Download](#download)
2. [Benchmarks](#benchmarks)
3. [Prompt Format](#prompt-format)
4. [Dataset](#dataset)
5. [Acknowledgements](#acknowledgements)
6. [Contact](#contact)
7. [About DiscoResearch](#about-discoresearch)
8. [Disclaimer](#disclaimer)

## Download 

| Huggingface    | GPTQ  | GGUF  | AWQ   | *Base Model* |
|-------|-------|-------|-------|-------|
| [Link](https://huggingface.co/DiscoResearch/DiscoLM-120b) | [Link](https://huggingface.co/TheBloke/DiscoLM-120b-GPTQ) | [Link](https://huggingface.co/TheBloke/DiscoLM-120b-GGUF) | [Link](https://huggingface.co/TheBloke/DiscoLM-120b-AWQ) | [Goliath 120b](https://huggingface.co/alpindale/goliath-120b) |

## Benchmarks

### Hugginface Leaderboard

This models is still an early Alpha and we can't guarantee that there isn't any contamination. 
However, the average of **73.198** would earn the #2 spot on the HF leaderboard at the time of writing and the highest score for a >70b model yet.

| Metric | Value |
|-----------------------|-------|
| ARC (25-shot)         | 69.54 |
| HellaSwag (10-shot)   | 86.49 |
| MMLU (5-shot)         | 70.32 |
| TruthfulQA (0-shot)   | 61.42 |
| Winogrande (5-shot)   | 83.03 |
| GSM8k (5-shot)   | 68.39 |
| **Avg.**                  | **73.198** |

We use [Language Model Evaluation Harness](https://github.com/EleutherAI/lm-evaluation-harness) to run the benchmark tests above, using the same version as the HuggingFace LLM Leaderboard.

### FastEval

| Metric | Value |
|-----------------------|-------|
| GSM8K       | 81.2 |
| Math   | 22.3 |
| BBH         | 72.9 |
| MMLU   | 67.9 |
| **Avg.**                  | **53.3** |

This places DiscoLM 120b firmly ahead of gpt-3.5-turbo-0613 as seen on the screenshot of the current (sadly no longer maintained) FastEval CoT leaderboard:
![FastEval Leaderboard](imgs/cot_leaderboard.png)

### MTBench

```json
{
    "first_turn": 8.45,
    "second_turn": 7.45,
    "categories": {
        "writing": 9.4,
        "roleplay": 8.65,
        "reasoning": 6.85,
        "math": 5.55,
        "coding": 4.95,
        "extraction": 9.15,
        "stem": 9.225,
        "humanities": 9.825
    },
    "average": 7.95
}
```
Screenshot of the current FastEval MT Bench leaderboard:
![FastEval Leaderboard](imgs/mtbench_leaderboard.png)

## Prompt Format

This model follows the ChatML format:

```
<|im_start|>system
You are DiscoLM, a helpful assistant.
<|im_end|>
<|im_start|>user
Please tell me possible reasons to call a research collective "Disco Research"<|im_end|>
<|im_start|>assistant
```

This formatting is also available via a pre-defined Transformers chat template, which means that lists of messages can be formatted for you with the apply_chat_template() method:

```python
chat = [
  {"role": "system", "content": "You are DiscoLM, a helpful assistant."},
  {"role": "user", "content": "Please tell me possible reasons to call a research collective Disco Research"}
]
tokenizer.apply_chat_template(chat, tokenize=False, add_generation_prompt=True)
```

If you use `tokenize=True` and `return_tensors="pt"` instead, then you will get a tokenized and formatted conversation ready to pass to `model.generate()`.

## Dataset

The dataset curation for DiscoLM 120b followed a "brute force"/"PoC" approach, as one goal was to see whether a 120b model can "absorb" more instruction data than a 70b model.

The following datasets were used for training DiscoLM 120b:

* [SlimOrca-Dedup](https://huggingface.co/datasets/Open-Orca/SlimOrca-Dedup)
* [OpenSchnabeltier](https://huggingface.co/datasets/LeoLM/OpenSchnabeltier) translated to DE from [OpenPlatypus](https://huggingface.co/datasets/garage-bAInd/Open-Platypus)
* [OpenHermes](https://huggingface.co/datasets/teknium/openhermes)
* [MetaMathQA](https://huggingface.co/datasets/meta-math/MetaMathQA)
* [UltraChat DE](https://huggingface.co/datasets/bjoernp/ultrachat_de) translated to DE from [UltraChat](https://huggingface.co/datasets/HuggingFaceH4/ultrachat_200k)
* [Synthia v.1.3](https://huggingface.co/datasets/migtissera/Synthia-v1.3)
* [German_Songs](https://huggingface.co/datasets/LeoLM/German_Songs)
* [German_Poems](https://huggingface.co/datasets/LeoLM/German_Poems)
* Capybara Dataset by [LDJnr](https://huggingface.co/LDJnr)
* Vezora/Tested-188k-Python (No longer available? Version changed to [Vezora/Tested-22k-Python-Alpaca](https://huggingface.co/datasets/Vezora/Tested-22k-Python-Alpaca))

Many thanks for all dataset providers/curators!

## Contact

Best way to reach us is on our [Discord](https://discord.gg/S8W8B5nz3v).

## About DiscoResearch

DiscoResearch is an aspiring open research community. Disco should be a place where researchers from many communities can come together to combine their expertise and create innovative and groundbreaking LLMs. Come join our Discord, share your opinions and ideas, and advance open LLM research with us!

## Acknowledgements

Disco 120b is a [DiscoResearch](https://huggingface.co/DiscoResearch) project and was trained by [Björn Plüster](https://huggingface.co/bjoernp). [Jan Harries](https://huggingface.co/jphme) helped with technical adivce, logistics and the Model Card and [AutoMeta](https://huggingface.co/Alignment-Lab-AI) also provided helpful technical adivce.
The model was trained with compute provided by [HessianAI](https://hessian.ai/) in collaboration with [LAION](https://laion.ai) - many thanks in particular to [Patrick Schramowski](https://huggingface.co/PSaiml) for his support. 

We are standing on the shoulders of giants; many thanks in no particular order to [LAION](https://laion.ai) and especially to [Christoph Schuhmann](https://laion.ai) who got us all connected,
[alpindale](https://huggingface.co/alpindale) for Goliath 120b (with important contributions by [Charles Goddard](https://huggingface.co/chargoddard) and [Undi95](https://huggingface.co/Undi95)), [TheBloke](https://huggingface.co/TheBloke) for providing quantized versions, [winglian](https://huggingface.co/winglian) for Axolotl which was used to train the model and the SlimOrca dataset, [garage-bAInd](https://huggingface.co/garage-bAInd), [Teknium](https://huggingface.co/teknium), [Migel Tissera](https://huggingface.co/migtissera), [MetaMath](https://huggingface.co/meta-math), and [LDJnr](https://huggingface.co/LDJnr) for their great datasets (please contact us if we forgot to mention you here!).

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)

## Disclaimer

The license on this model does not constitute legal advice. We are not responsible for the actions of third parties who use this model.
This model should only be used for research purposes. The original Llama2 license and all restrictions of datasets used to train this model apply.
