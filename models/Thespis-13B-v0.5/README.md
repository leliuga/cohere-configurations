---
tags:
- not-for-all-audiences
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/64dd7cda3d6b954bf7cdd922/ZXmxNKGaHUrqjdS1I3GkL.png)

This model is a personal project. It uses a vanilla chat template and is focused on providing multiturn sfw and nsfw RP experience.

This model works best with internet style RP using standard markup with asterisks surrounding actions and no quotes around dialogue.

It uses the following data:

* 1500 samples from Claude Multiround Chat 30k dataset ( 90 token length or greater. Coding and math filtered out. )
* 2200 samples from Puffin Dataset ( 90 token length or greater. Coding and math filtered out. )
* 700 samples from Airoboros 3.1 ( 1500 token length or greater. Writing samples only. )
* 900 samples from the Augmental Dataset ( 90 token length or greater )
* 4200 samples of hand curated RP conversation with various characters.


Works with standard chat format for Ooba or SillyTavern.

## Prompt Format: Chat ( The default Ooba template and Silly Tavern Template )
```
{System Prompt}

Username: {Input}
BotName: {Response}
Username: {Input}
BotName: {Response}

```
## Ooba ( Set it to Chat, select a character and go. )
![image/png](https://cdn-uploads.huggingface.co/production/uploads/64dd7cda3d6b954bf7cdd922/HTl7QlAZcqe2hV8rwh4DG.png)

## Silly Tavern Settings ( Default )
![image/png](https://cdn-uploads.huggingface.co/production/uploads/64dd7cda3d6b954bf7cdd922/ajny8P0LdW0nFtghpPbfB.png)



## Turn Template (for Ooba Instruct if making a Discord bot or Some other Many to one Chat):

You can either bake usernames into the prompt directly for ease of use or programatically add them if running through the API to use as a chatbot.

```
User string: ( Leave empty if populating username into prompt through a script. Put in your username if its a 1 on 1 convo.) Ex. "DiscordUser1: "
Bot String: ( The bots name, followed by a colon and a space.) Ex. "Mayo: "
Context: ( Your bots system prompt, follow by a newline. )

<|user|><|user-message|>\n<|bot|><|bot-message|>\n
```

## Examples

SFW Roleplay
![image/png](https://cdn-uploads.huggingface.co/production/uploads/64dd7cda3d6b954bf7cdd922/IpfXqucvzoXhjIvg0COVz.png)

NSFW Roleplay
![image/png](https://cdn-uploads.huggingface.co/production/uploads/64dd7cda3d6b954bf7cdd922/S6mwD15Xd3E7WhhgWNBva.png)


