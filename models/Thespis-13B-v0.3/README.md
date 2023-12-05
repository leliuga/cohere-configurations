# This model is outdated, please use the much better v0.4 -> https://huggingface.co/cgato/Thespis-13b-v0.4

This model is a personal project. It uses a vanilla chat template and is focused on providing multiturn sfw and nsfw RP experience. 

This model works best with internet style RP using standard markup with asterisks surrounding actions and no quotes around dialogue.

It uses the following data:

* 3000 samples from Claude Multiround Chat 30k dataset
* 6000 samples from Pippa Dataset
* 3000 samples from Puffin Dataset
* 3800 samples of hand curated RP conversation with various characters.


Works with standard chat format for Ooba or SillyTavern.

Prompt Format: Chat
```
{System Prompt}

Username: {Input}
BotName: {Response}
Username: {Input}
BotName: {Response}

```


Turn Template (for Ooba):
You can either bake usernames into the prompt directly for ease of use or programatically add them if running through the API to use as a chatbot.
```
<|user|>{Username}: <|user-message|>\n<|bot|>{BotName}: <|bot-message|>\n
```
