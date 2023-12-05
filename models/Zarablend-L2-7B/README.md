---
license: other
tags:
- llama2
---
# Model Card: Zarablend L2 7b
This model uses [Nous Hermes Llama2 7b](https://huggingface.co/NousResearch/Nous-Hermes-llama-2-7b) (66%) as a base with [Airoboros L2 7B GPT4 2.0](https://huggingface.co/jondurbin/airoboros-l2-7b-gpt4-2.0) (34%) and the result of this merge was merged with [LimaRP LLama2 7B Lora](https://huggingface.co/lemonilia/limarp-llama2).

This merge of models(hermes and airoboros) was done with this [script](https://github.com/zarakiquemparte/zaraki-tools/blob/main/merge-cli.py)

This merge of Lora with Model was done with this [script](https://github.com/zarakiquemparte/zaraki-tools/blob/main/apply-lora.py)

Quantized Model by @TheBloke:
- [GGML](https://huggingface.co/TheBloke/Zarablend-L2-7B-GGML)
- [GPTQ](https://huggingface.co/TheBloke/Zarablend-L2-7B-GPTQ)

Merge illustration:

![illustration](zarablend-merge-illustration.png)

## Usage:

Since this is a merge between Nous Hermes, Airoboros and LimaRP, the following instruction formats should work:

Alpaca 2:

```
### Instruction:
<prompt>

### Response:
<leave a newline blank for model to respond>
```

LimaRP instruction format:

```
<<SYSTEM>>
<character card and system prompt>

<<USER>>
<prompt>

<<AIBOT>>
<leave a newline blank for model to respond>
```

## Bias, Risks, and Limitations

This model is not intended for supplying factual information or advice in any form

## Training Details

This model is merged and can be reproduced using the tools mentioned above. Please refer to all provided links for extra model-specific details.