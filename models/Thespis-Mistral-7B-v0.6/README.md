---
license: apache-2.0
tags:
- not-for-all-audiences
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/64dd7cda3d6b954bf7cdd922/ZXmxNKGaHUrqjdS1I3GkL.png)

## Mistral 7b Version of v0.6 ( with some extras )

This version has some additional data vs the llama2 13b version. Making it something of a 0.6 and a half or something. But that's confusing.

The 7b version of my Thespis finetune. From my testing it seems to perform somewhere between llama 7b and llama 13b. 

This model is a personal project. It uses a vanilla chat template and is focused on providing multiturn sfw and nsfw RP experience.

This model works best with internet style RP using standard markup with asterisks surrounding actions and no quotes around dialogue.

It uses the following data:

* 2200 samples from Pure-Dove Dataset ( 90 token length or greater. )
* 2200 samples from Claude Multiround 30k ( 90 token length or greater. )
* 700 samples from Airoboros 3.1 ( Writing samples longer than 1500 tokens only. )
* 900 samples from the Augmental Dataset ( 90 token length or greater )
* 6000 samples of hand curated RP conversation with various characters.

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