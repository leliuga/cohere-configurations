---
tags:
- not-for-all-audiences
license: apache-2.0
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/64dd7cda3d6b954bf7cdd922/ZXmxNKGaHUrqjdS1I3GkL.png)


## Alpha Test Version of Thespis v0.7 on Mistral
This is a version of Thespis which includes more than the usual amount of Assistant responses in the training data. 
Its good for acting as an assistant with personality. It may be ok for RP, feel free to experiment.

This model works best with internet style RP using standard markup with asterisks surrounding actions and no quotes around dialogue.
Currently in the middle of adding new features and working on making sure initial responses are strong.

It uses a 40/60 Split of Assistant and RP data, and includes items from the below datasets:

* Pure-Dove Dataset
* Claude Multiround 30k
* No Robots
* OpenOrcaSlim
* Augmental Dataset

As a result of adding this additional Assistant data, the model shows greatly improved instruction following and coherence during multiturn chats.



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