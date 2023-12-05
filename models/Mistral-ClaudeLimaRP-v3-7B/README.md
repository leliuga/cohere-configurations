---
license: apache-2.0
tags:
- not-for-all-audiences
- nsfw
---

## Description

This repo contains fp16 files of [Norquinal/Mistral-7B-claude-chat](https://huggingface.co/Norquinal/Mistral-7B-claude-chat) with the LoRA [lemonilia/LimaRP-Mistral-7B-v0.1](https://huggingface.co/lemonilia/LimaRP-Mistral-7B-v0.1) applied at weight "0.75".

All credit go to [lemonilia](https://huggingface.co/lemonilia) and [Norquinal](https://huggingface.co/Norquinal)

## Prompt format
Same as before. It uses the [extended Alpaca format](https://github.com/tatsu-lab/stanford_alpaca),
with `### Input:` immediately preceding user inputs and `### Response:` immediately preceding
model outputs. While Alpaca wasn't originally intended for multi-turn responses, in practice this
is not a problem; the format follows a pattern already used by other models.

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

You should:
- Replace all text in curly braces (curly braces included) with your own text.
- Replace `User` and `Character` with appropriate names.


### Message length control
Inspired by the previously named "Roleplay" preset in SillyTavern, with this
version of LimaRP it is possible to append a length modifier to the response instruction
sequence, like this:

```
### Input
User: {utterance}

### Response: (length = medium)
Character: {utterance}
```

This has an immediately noticeable effect on bot responses. The available lengths are:
`tiny`, `short`, `medium`, `long`, `huge`, `humongous`, `extreme`, `unlimited`. **The 
recommended starting length is `medium`**. Keep in mind that the AI may ramble
or impersonate the user with very long messages.

The length control effect is reproducible, but the messages will not necessarily follow
lengths very precisely, rather follow certain ranges on average, as seen in this table
with data from tests made with one reply at the beginning of the conversation:

![lengths](https://files.catbox.moe/dy39bt.png)

Response length control appears to work well also deep into the conversation.

## Suggested settings
You can follow these instruction format settings in SillyTavern. Replace `tiny` with
your desired response length:

![settings](https://files.catbox.moe/6lcz0u.png)

## Text generation settings
Extensive testing with Mistral has not been performed yet, but suggested starting text
generation settings may be:

- TFS = 0.90~0.95
- Temperature = 0.70~0.85
- Repetition penalty = 1.08~1.10
- top-k = 0 (disabled)
- top-p = 1 (disabled)

If you want to support me, you can [here](https://ko-fi.com/undiai).