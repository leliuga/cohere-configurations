---
license: other
tags:
- llama-2
---
# Model Card: Hermes Limarp L2 7b
This model uses [Nous Hermes Llama2 7b](https://huggingface.co/NousResearch/Nous-Hermes-llama-2-7b) as a base and merged with [LimaRP LLama2 7B](https://huggingface.co/lemonilia/limarp-llama2).

This merge of Lora with Model was done with this [script](https://github.com/zarakiquemparte/zaraki-tools/blob/main/apply-lora.py)

Quantized Model by @TheBloke:
- [GGML](https://huggingface.co/TheBloke/HermesLimaRP-L2-7B-GGML)
- [GPTQ](https://huggingface.co/TheBloke/HermesLimaRP-L2-7B-GPTQ)
  
## Usage:

Since this is a merge between Nous Hermes and LimaRP, the following instruction formats should work:

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