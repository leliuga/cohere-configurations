---
inference: false
language:
- en
library_name: transformers
pipeline_tag: text-generation
tags:
- mixtral
license: apache-2.0
datasets:
- lemonilia/LimaRP
---

# Mixtral-8x7B-Instruct-v0.1-LimaRP-ZLoss

Experimental model, using a limarp qlora trained at 10k ctx length (greater than size of the longest limarp sample when tokenized via mistral's tokenizer) on [mistralai/Mixtral-8x7B-v0.1](https://huggingface.co/mistralai/Mixtral-8x7B-v0.1) using [Charles Goddard](https://huggingface.co/chargoddard)'s ZLoss and Megablocks-based fork of transformers, and then fused to [mistralai/Mixtral-8x7B-Instruct-v0.1](https://huggingface.co/mistralai/Mixtral-8x7B-Instruct-v0.1) at 0.5 weight.

My current generation settings are:
```
Temperature: 1.25
Min-p: 0.05
Repetition penalty: 1.05
Repetition penalty: range 1024
```
And this seems to avoid the Mixtral looping pitfalls for me so far. Play around with it and see what works well for you.

[Peft Adapter](https://huggingface.co/Doctor-Shotgun/limarp-zloss-mixtral-8x7b-qlora)

Quants courtesy of TheBloke:
- [GPTQ](https://huggingface.co/TheBloke/Mixtral-8x7B-Instruct-v0.1-LimaRP-ZLoss-GPTQ)
- [GGUF](https://huggingface.co/TheBloke/Mixtral-8x7B-Instruct-v0.1-LimaRP-ZLoss-GGUF)
- [AWQ](https://huggingface.co/TheBloke/Mixtral-8x7B-Instruct-v0.1-LimaRP-ZLoss-AWQ)

Exl2 Quants courtesy of LoneStriker:
- [2.4bpw](https://huggingface.co/LoneStriker/Mixtral-8x7B-Instruct-v0.1-LimaRP-ZLoss-2.4bpw-h6-exl2)
- [3.0bpw](https://huggingface.co/LoneStriker/Mixtral-8x7B-Instruct-v0.1-LimaRP-ZLoss-3.0bpw-h6-exl2)
- [3.5bpw](https://huggingface.co/LoneStriker/Mixtral-8x7B-Instruct-v0.1-LimaRP-ZLoss-3.5bpw-h6-exl2)
- [3.75bpw](https://huggingface.co/LoneStriker/Mixtral-8x7B-Instruct-v0.1-LimaRP-ZLoss-3.75bpw-h6-exl2)
- [4.0bpw](https://huggingface.co/LoneStriker/Mixtral-8x7B-Instruct-v0.1-LimaRP-ZLoss-4.0bpw-h6-exl2)
- [5.0bpw](https://huggingface.co/LoneStriker/Mixtral-8x7B-Instruct-v0.1-LimaRP-ZLoss-5.0bpw-h6-exl2)
- [6.0bpw](https://huggingface.co/LoneStriker/Mixtral-8x7B-Instruct-v0.1-LimaRP-ZLoss-6.0bpw-h6-exl2)

## Usage:
The intended prompt format is the Alpaca instruction format of LimaRP v3:
```
### Instruction:
Character's Persona: {bot character description}

User's Persona: {user character description}

Scenario: {what happens in the story}

Play the role of Character. Taking the above information into consideration, you must engage in a roleplaying chat with User below this line. Do not write dialogues and narration for User.

### Input:
User: {utterance}

### Response:
Character: {utterance}

### Input:
User: {utterance}

### Response:
Character: {utterance}

(etc.)
```
My current templates have been uploaded to a [folder](https://huggingface.co/Doctor-Shotgun/Mixtral-8x7B-Instruct-v0.1-LimaRP-ZLoss/tree/main/ST%20Templates).

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