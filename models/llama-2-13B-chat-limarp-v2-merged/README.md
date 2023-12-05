---
inference: false
language:
- en
library_name: transformers
pipeline_tag: text-generation
tags:
- llama
- llama-2
license: agpl-3.0
---
# Model Card: llama-2-13b-chat-limarp-v2-merged

This is a Llama 2-based model consisting of Llama 2 13b chat (https://huggingface.co/meta-llama/Llama-2-13b-chat-hf) merged with LIMARP Lora v2 (https://huggingface.co/lemonilia/limarp-llama2-v2).

Requested by @dampf

## Usage:
Intended to be prompted with the Alpaca instruction format of the LIMARP v2:
```
### Instruction:
Character's Persona: {bot character description}

User's Persona: {user character description}

Scenario: {what happens in the story}

Play the role of Character. You must engage in a roleplaying chat with User below this line. Do not write dialogues and narration for User. Character should respond with messages of medium length.

### Input:
Character: {utterance}

### Response:
User: {utterance}
```
## Bias, Risks, and Limitations
The model will show biases similar to those observed in niche roleplaying forums on the Internet, besides those exhibited by the base model. It is not intended for supplying factual information or advice in any form. 
## Training Details
This model is a merge. Please refer to the link repositories of the base model and lora for details.