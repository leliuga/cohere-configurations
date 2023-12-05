---
inference: false
language:
- en
library_name: transformers
pipeline_tag: text-generation
tags:
- mistral
license: cc-by-nc-4.0
---
# ZephRP-m7b

This is a [Mistral](https://huggingface.co/mistralai/Mistral-7B-v0.1)-based model consisting of a merge between [HuggingFaceH4/zephyr-7b-alpha](https://huggingface.co/HuggingFaceH4/zephyr-7b-alpha) and PEFT adapter trained using the LimaRP dataset.

The goal was to combine the message length instruction training of LimaRPv3 and additional stylistic elements with the superior knowledge and instruction-following capabilities of the Zephyr model.

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
The LimaRP PEFT adapter was trained as an 8-bit lora using [axolotl](https://github.com/OpenAccess-AI-Collective/axolotl).

The following hyperparameters were used during training of the adapter on the original [mistralai/Mistral-7B-v0.1](https://huggingface.co/mistralai/Mistral-7B-v0.1) model using a single L40 GPU:
- learning_rate: 0.00015
- train_batch_size: 2
- eval_batch_size: 2
- seed: 42
- gradient_accumulation_steps: 4
- total_train_batch_size: 8
- optimizer: Adam with betas=(0.9,0.999) and epsilon=1e-08
- lr_scheduler_type: cosine
- lr_scheduler_warmup_steps: 10
- num_epochs: 2