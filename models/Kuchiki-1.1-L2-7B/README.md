---
license: other
tags:
- llama2
---
# Model Card: Kuchiki 1.1 L2 7b
This model uses [Nous Hermes Llama2 7b](https://huggingface.co/NousResearch/Nous-Hermes-llama-2-7b) (70%) as a base with [Airoboros L2 7B GPT4 2.0](https://huggingface.co/jondurbin/airoboros-l2-7b-gpt4-2.0) (30%) and the result of this merge was merged with [LimaRP Llama2 v2 7B Lora](https://huggingface.co/lemonilia/limarp-llama2-v2).

This merge of models(hermes and airoboros) was done with this [script](https://github.com/zarakiquemparte/zaraki-tools/blob/main/merge-cli.py)

This merge of Lora with Model was done with this [script](https://github.com/zarakiquemparte/zaraki-tools/blob/main/apply-lora.py)

Merge illustration:

![illustration](merge-illustration.png)

## Usage:

Since this is a merge between Nous Hermes, Airoboros and LimaRP, the following instruction formats should work:

Alpaca 2:

```
### Instruction:
<prompt>

### Response:
<leave a newline blank for model to respond>
```

Alpaca LimaRP:

```
### Instruction:
Character's Persona: {bot character description}

User's Persona: {user character description}

Scenario: {what happens in the story}

Play the role of Character. You must engage in a roleplaying chat with User below this line. Do not write dialogues and narration for User. Character should respond with messages of medium length.

### Input:
User: {utterance}

### Response:
Character: {utterance}
```

## Bias, Risks, and Limitations

This model is not intended for supplying factual information or advice in any form

## Training Details

This model is merged and can be reproduced using the tools mentioned above. Please refer to all provided links for extra model-specific details.