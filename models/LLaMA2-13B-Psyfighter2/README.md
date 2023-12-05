---
license: llama2
---

# LLAMA2-13B-Psyfighter2
Psyfighter is a merged model created by the KoboldAI community members Jeb Carter and TwistedShadows and was made possible thanks to the KoboldAI merge request service.

The intent was to add medical data to supplement the models fictional ability with more details on anatomy and mental states. Due to the low ratio's of medical data and the high ratio's of fiction this model should not be used for medical advice or therapy because of its high chance of pulling in fictional data.

The following mergekit recipe was used:

```
merge_method: task_arithmetic
base_model: TheBloke/Llama-2-13B-fp16
models:
  - model: TheBloke/Llama-2-13B-fp16
  - model: KoboldAI/LLaMA2-13B-Tiefighter
    parameters:
      weight: 1.0
  - model: Doctor-Shotgun/cat-v1.0-13b
    parameters:
      weight: 0.01
  - model: Doctor-Shotgun/llama-2-13b-chat-limarp-v2-merged
    parameters:
      weight: 0.02
dtype: float16
```

*V1 of this model was published under the account of the creator of the merge

This model contains the following ingredients from their upstream models for as far as we can track them:
- KoboldAI/LLaMA2-13B-Tiefighter
	- Undi95/Xwin-MLewd-13B-V0.2
	- - Undi95/ReMM-S-Light
	  - Undi95/CreativeEngine
	  - Brouz/Slerpeno
	  - - elinas/chronos-13b-v2
	    - jondurbin/airoboros-l2-13b-2.1
	    - NousResearch/Nous-Hermes-Llama2-13b+nRuaif/Kimiko-v2
	    - CalderaAI/13B-Legerdemain-L2+lemonilia/limarp-llama2-v2
	    - - KoboldAI/LLAMA2-13B-Holodeck-1
	      - NousResearch/Nous-Hermes-13b
	      - OpenAssistant/llama2-13b-orca-8k-3319
	    - ehartford/WizardLM-1.0-Uncensored-Llama2-13b
	    - Henk717/spring-dragon
	  - The-Face-Of-Goonery/Huginn-v3-13b (Contains undisclosed model versions, those we assumed where possible)
	  - - SuperCOT (Undisclosed version)
	    - elinas/chronos-13b-v2 (Version assumed)
	    - NousResearch/Nous-Hermes-Llama2-13b
	    - stabilityai/StableBeluga-13B (Version assumed)
	  - zattio770/120-Days-of-LORA-v2-13B
	  - PygmalionAI/pygmalion-2-13b
	  - Undi95/Storytelling-v1-13B-lora
	  - TokenBender/sakhi_13B_roleplayer_NSFW_chat_adapter
	  - nRuaif/Kimiko-v2-13B
	  - The-Face-Of-Goonery/Huginn-13b-FP16
	  - - "a lot of different models, like hermes, beluga, airoboros, chronos.. limarp"
	  - lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT
	  - Xwin-LM/Xwin-LM-13B-V0.2
	- PocketDoc/Dans-RetroRodeo-13b
	- Blackroot/Llama-2-13B-Storywriter-LORA
- Doctor-Shotgun/cat-v1.0-13b
- Doctor-Shotgun/llama-2-13b-chat-limarp-v2-merged
   - meta-llama/Llama-2-13b-chat-hf
   - lemonilia/limarp-llama2-v2


While we could possibly not credit every single lora or model involved in this merged model, we'd like to thank all involved creators upstream for making this awesome model possible!
Thanks to you the AI ecosystem is thriving, and without your dedicated tuning efforts models such as this one would not be possible.

# Usage
This model is meant to be creative, If you let it improvise you get better results than if you drown it in details.

## Story Writing
Regular story writing in the traditional way is supported, simply copy paste your story and continue writing. Optionally use an instruction in memory or an authors note to guide the direction of your story.

### Generate a story on demand
To generate stories on demand you can use an instruction (tested in the Alpaca format) such as "Write a novel about X, use chapters and dialogue" this will generate a story. The format can vary between generations depending on how the model chooses to begin, either write what you want as shown in the earlier example or write the beginning of the story yourself so the model can follow your style. A few retries can also help if the model gets it wrong.

## Chatbots and persona's
This model has been tested with various forms of chatting, testers have found that typically less is more and the model is good at improvising. Don't drown the model in paragraphs of detailed information, instead keep it simple first and see how far you can lean on the models own ability to figure out your character. Copy pasting paragraphs of background information is not suitable for a 13B model such as this one, code formatted characters or an instruction prompt describing who you wish to talk to goes much further.

For example, you can put this in memory in regular chat mode: 
``` 
### Instruction: 
Generate a conversation between Alice and Jeb where they discuss language models.
In this conversation Henk is excited to teach Alice about Psyfighter. 
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
We can also provide assistance in making your own merges.