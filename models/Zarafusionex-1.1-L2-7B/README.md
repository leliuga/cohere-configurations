---
license: other
tags:
- llama2
---
# Model Card: Zarafusionex 1.1 L2 7b
This model uses [Nous Hermes Llama2 7b](https://huggingface.co/NousResearch/Nous-Hermes-llama-2-7b) (53%) as a base with [Stable Beluga 7b](https://huggingface.co/stabilityai/StableBeluga-7B) (47%) and the result of this merge was merged with [LimaRP LLama2 7B Lora version of the day 07/23/2023](https://huggingface.co/lemonilia/limarp-llama2).

This merge of models(hermes and stable beluga) was done with this [script](https://github.com/zarakiquemparte/zaraki-tools/blob/main/merge-cli.py)

This merge of Lora with Model was done with this [script](https://github.com/zarakiquemparte/zaraki-tools/blob/main/apply-lora.py)

Quantized Model by @TheBloke:
- [GGML](https://huggingface.co/TheBloke/Zarafusionex-1.1-L2-7B-GGML)
- [GGUF](https://huggingface.co/TheBloke/Zarafusionex-1.1-L2-7B-GGUF)
- [GPTQ](https://huggingface.co/TheBloke/Zarafusionex-1.1-L2-7B-GPTQ)

Merge illustration:

![illustration](zarafusionex-merge-illustration.png)

## Usage:

Since this is a merge between Nous Hermes, Stable Beluga and LimaRP, the following instruction formats should work:

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