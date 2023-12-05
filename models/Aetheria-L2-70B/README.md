---
inference: false
language:
- en
library_name: transformers
pipeline_tag: text-generation
tags:
- llama
- llama 2
---
# Aetheria-L2-70B

This is a [Llama 2](https://huggingface.co/meta-llama/Llama-2-70b)-based model consisting of a merge between:
- [Sao10K/Euryale-1.3-L2-70B](https://huggingface.co/Sao10K/Euryale-1.3-L2-70B)
- [allenai/tulu-2-dpo-70b](https://huggingface.co/allenai/tulu-2-dpo-70b)
- [GOAT-AI/GOAT-70B-Storytelling](https://huggingface.co/GOAT-AI/GOAT-70B-Storytelling)
- [Doctor-Shotgun/limarpv3-llama2-70b-qlora](https://huggingface.co/Doctor-Shotgun/limarpv3-llama2-70b-qlora)

This model combines the excellent Euryale v1.3 base with the DPO training of the Tulu v2 model and creative prose training of the GOAT Storytelling model. The LimaRP v3 qlora was then added for further roleplaying capability and the ability to tune the length of the outputs.

The goal was to create a capable 70B model for collaborative storytelling and roleplay.

## Usage:
The intended prompt format is the Alpaca instruction format of LimaRP v3:
```
### Instruction:
Character's Persona: {bot character description}

User's Persona: {user character description}

Scenario: {what happens in the story}

Play the role of Character. You must engage in a roleplaying chat with User below this line. Do not write dialogues and narration for User.

### Input:
User: {utterance}

### Response:
Character: {utterance}

### Input
User: {utterance}

### Response:
Character: {utterance}

(etc.)
```

## Message length control
Due to the inclusion of LimaRP v3, it is possible to append a length modifier to the response instruction sequence, like this:
```
### Input
User: {utterance}

### Response: (length = medium)
Character: {utterance}
```
This has an immediately noticeable effect on bot responses. The available lengths are: `micro, tiny, short, medium, long, massive, huge, enormous, humongous, unlimited`. The recommended starting length is `medium`. Keep in mind that the AI may ramble or impersonate the user with very long messages.
## Bias, Risks, and Limitations
The model will show biases similar to those observed in niche roleplaying forums on the Internet, besides those exhibited by the base model. It is not intended for supplying factual information or advice in any form. 
## Training Details
This model is a merge. Please refer to the link repositories of the merged models for details.