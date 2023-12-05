---
tags:
- not-for-all-audiences
---


![image/png](https://cdn-uploads.huggingface.co/production/uploads/64dd7cda3d6b954bf7cdd922/ZXmxNKGaHUrqjdS1I3GkL.png)

This model is a personal project. It uses a vanilla chat template and is focused on providing multiturn sfw and nsfw RP experience. 

This model works best with internet style RP using standard markup with asterisks surrounding actions and no quotes around dialogue.

It uses the following data:

* 3000 samples from Claude Multiround Chat 30k dataset ( 45 token length or greater. Coding and math filtered out. )
* 3000 samples from Pippa Dataset ( 90 token length or greater. Filtered for responses which did not contain quotes and contained at least one asterisk. )
* 2200 samples from Puffin Dataset ( 90 token length or greater. Coding and math filtered out. )
* 700 samples from Airoboros 3.1 ( 1500 token length or greater. Writing samples only. )
* 4700 samples of hand curated RP conversation with various characters.


Works with standard chat format for Ooba or SillyTavern.

## Prompt Format: Chat
```
{System Prompt}

Username: {Input}
BotName: {Response}
Username: {Input}
BotName: {Response}

```

## Turn Template (for Ooba Instruct):
You can either bake usernames into the prompt directly for ease of use or programatically add them if running through the API to use as a chatbot.
```
<|user|>{Username}: <|user-message|>\n<|bot|>{BotName}: <|bot-message|>\n
```

## Examples

![image/png](https://cdn-uploads.huggingface.co/production/uploads/64dd7cda3d6b954bf7cdd922/Uhf5IRiudjEB8-z547BWi.png)

![image/png](https://cdn-uploads.huggingface.co/production/uploads/64dd7cda3d6b954bf7cdd922/maO25mC1pfcYoJacpaM2t.png)

![image/png](https://cdn-uploads.huggingface.co/production/uploads/64dd7cda3d6b954bf7cdd922/t7zJ4cCkSbtQv0ONVgOoS.png)