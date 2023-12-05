---
license: llama2
---
# LLaMA2-13B-TiefighterLR
TiefighterLR is a merged model achieved trough merging two different lora's on top of a well established existing merge.
This LR version contains Less Rodeo, merged at 3% from the original 5% reducing its second person adventure bias.
Testers found this model to understand your own character and instruction prompts better, at the sacrifice of lowering its own writing bias/style.

To achieve this the following recipe was used:

* We begin with the base model Undi95/Xwin-MLewd-13B-V0.2 which is a well established merge, contrary to the name this model does not have a strong NSFW bias.
* Then we applied the PocketDoc/Dans-RetroRodeo-13b lora which is a finetune on the Choose your own Adventure datasets from our Skein model.
* After applying this lora we merged the original model with the newly created PocketDoc/Dans-RetroRodeo-13b merge at 3% to weaken the newly introduced adventure bias.
* The resulting merge was used as a new base model to which we applied Blackroot/Llama-2-13B-Storywriter-LORA and repeated the same trick, this time at 10%.

This means this model contains the following ingredients from their upstream models for as far as we can track them:
- Undi95/Xwin-MLewd-13B-V0.2
- - Undi95/ReMM-S-Light (base/private)
  - Undi95/CreativeEngine
  - Brouz/Slerpeno
  - - elinas/chronos-13b-v2
    - jondurbin/airoboros-l2-13b-2.1
    - NousResearch/Nous-Hermes-Llama2-13b+nRuaif/Kimiko-v2 LORA
    - CalderaAI/13B-Legerdemain-L2+lemonilia/limarp-llama2-v2 LORA
    - - KoboldAI/LLAMA2-13B-Holodeck-1
      - NousResearch/Nous-Hermes-13b
      - OpenAssistant/llama2-13b-orca-8k-3319
    - ehartford/WizardLM-1.0-Uncensored-Llama2-13b
    - Henk717/spring-dragon
  - The-Face-Of-Goonery/Huginn-v3-13b
  - zattio770/120-Days-of-LORA-v2-13B
  - PygmalionAI/pygmalion-2-13b
  - Undi95/StoryTelling
  - TokenBender/sakhi_13B_roleplayer_NSFW_chat_adapter
  - nRuaif/Kimiko-v2-13B
  - The-Face-Of-Goonery/Huginn-13b-FP16
  - lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT
  - Xwin-LM/Xwin-LM-13B-V0.2
- PocketDoc/Dans-RetroRodeo-13b
- Blackroot/Llama-2-13B-Storywriter-LORA


# Usage
This model is meant to be creative, If you let it improvise you get better results than if you drown it in details.

## Story Writing
Regular story writing in the traditional way is supported, simply copy paste your story and continue writing. Optionally use an instruction in memory or an authors note to guide the direction of your story.

### Generate a story on demand
To generate stories on demand you can use an instruction (tested in the Alpaca format) such as "Write a novel about X, use chapters and dialogue" this will generate a story. The format can vary between generations depending on how the model chooses to begin, either write what you want as shown in the earlier example or write the beginning of the story yourself so the model can follow your style. A few retries can also help if the model gets it wrong.

## Chatbots and persona's
Unlike the original Tiefighter this model is better at handling existing Character Cards as long as they do not contain a lot of second person writing or second person introductions (You), setting > as a custom stop sequence can help fix potential mistakes, as well as turning multi-line replies off.
You can also use instructions to create your characters.

For example, you can put this in memory in regular chat mode: 
``` 
### Instruction: 
Generate a conversation between Alice and Henk where they discuss language models.
In this conversation Henk is excited to teach Alice about Tiefighter. 
### Response: 
```

Because the model is a merge of a variety of models, it should support a broad range of instruct formats, or plain chat mode. If you have a particular favourite try it, otherwise we recommend to either use the regular chat mode or Alpaca's format.

## Instruct Prompting
This model features various instruct models on a variety of instruction styles, when testing the model we have used Alpaca for our own tests. If you prefer a different format chances are it can work.

During instructions we have observed that in some cases the adventure data can leak, it may also be worth experimenting using > as the prefix for a user command to remedy this. But this may result in a stronger fiction bias.

Keep in mind that while this model can be used as a factual instruct model, the focus was on fiction. Information provided by the model can be made up.

## Adventuring and Adventure Games
This model contains a lora that was trained on the same adventure dataset as the KoboldAI Skein model. Adventuring is best done using an small introduction to the world and your objective while using the > prefix for a user command (KoboldAI's adventure mode).

It is possible that the model does not immediately pick up on what you wish to do and does not engage in its Adventure mode behaviour right away. Simply manually correct the output to trim excess dialogue or other undesirable behaviour and continue to submit your actions using the appropriate mode. The model should pick up on this style quickly and will correctly follow this format within 3 turns.

## Discovered something cool and want to engage with us? 
Join our community at https://koboldai.org/discord !

### This model would not be possible without the awesome work from: 
Undi95, PocketDoc, Blackroot, Brouz, The Face of Goonery, zattio770, PygmalionAI, TokenBender, nRuaif, lemonilia and Xwin-LM.