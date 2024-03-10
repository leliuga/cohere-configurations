---
base_model:
- TheBloke/Llama-2-13B-fp16
tags:
- mergekit
- merge
license: cc-by-nc-4.0
---
# Introduction
- Estopia is a model focused on improving the dialogue and prose returned when using the instruct format. As a side benefit, character cards and similar seem to have also improved, remembering details well in many cases.
- It focuses on "guided narratives" - using instructions to guide or explore fictional stories, where you act as a guide for the AI to narrate and fill in the details.
- It has primarily been tested around prose, using instructions to guide narrative, detail retention and "neutrality" - in particular with regards to plot armour. Unless you define different rules for your adventure / narrative with instructions, it should be realistic in the responses provided.
- It has been tested using different modes, such as instruct, chat, adventure and story modes - and should be able to do them all to a degree, with it's strengths being instruct and adventure, with story being a close second.
# Usage
- The Estopia model has been tested primarily using the Alpaca format, but with the range of models included likely has some understanding of others. Some examples of tested formats are below:
    - ```\n### Instruction:\nWhat colour is the sky?\n### Response:\nThe sky is...```
    - ```<Story text>\n***\nWrite a summary of the text above\n***\nThe story starts by...```
    - Using the Kobold Lite AI adventure mode
    - ```User:Hello there!\nAssistant:Good morning...\n```
- For settings, the following are recommended for general use:
    - Temperature: 0.8-1.2
    - Min P: 0.05-0.1
    - Max P: 0.92, or 1 if using a Min P greater than 0
    - Top K: 0
    - Response length: Higher than your usual amount most likely - for example a common value selected is 512.
        - Note: Response lengths are not guaranteed to always be this length. On occasion, responses may be shorter if they convey the response entirely, other times they could be upwards of this value. It depends mostly on the character card, instructions, etc.
    - Rep Pen: 1.1
    - Rep Pen Range: 2 or 3x your response length
    - Stopping tokens (Not needed, but can help if the AI is writing too much):
        - ```##||$||---||$||ASSISTANT:||$||[End||$||</s>``` - A single string for Kobold Lite combining the ones below
        - ```##```
        - ```---```
        - ```ASSISTANT:```
        - ```[End```
        - ```</s>```
- The settings above should provide a generally good experience balancing instruction following and creativity. Generally the higher you set the temperature, the greater the creativity and higher chance of logical errors when providing responses from the AI.
# Recipe
This model was made in three stages, along with many experimental stages which will be skipped for brevity. The first was internally referred to as EstopiaV9, which has a high degree of instruction following and creativity in responses, though they were generally shorter and a little more restricted in the scope of outputs, but conveyed nuance better.
```yaml
merge_method: task_arithmetic
base_model: TheBloke/Llama-2-13B-fp16
models:
- model: TheBloke/Llama-2-13B-fp16
- model: Undi95/UtopiaXL-13B
    parameters:
    weight: 1.0
- model: Doctor-Shotgun/cat-v1.0-13b
    parameters:
    weight: 0.02
- model: PygmalionAI/mythalion-13b
    parameters:
    weight: 0.10
- model: Undi95/Emerhyst-13B
    parameters:
    weight: 0.05
- model: CalderaAI/13B-Thorns-l2
    parameters:
    weight: 0.05
- model: KoboldAI/LLaMA2-13B-Tiefighter
    parameters:
    weight: 0.20
dtype: float16
```
The second part of the merge was known as EstopiaV13. This produced responses which were long, but tended to write beyond good stopping points for further instructions to be added as it leant heavily on novel style prose. It did however benefit from a greater degree of neutrality as described above, and retained many of the detail tracking abilities of V9.
```yaml
merge_method: task_arithmetic
base_model: TheBloke/Llama-2-13B-fp16
models:
  - model: TheBloke/Llama-2-13B-fp16
  - model: Undi95/UtopiaXL-13B
    parameters:
      weight: 1.0
  - model: Doctor-Shotgun/cat-v1.0-13b
    parameters:
      weight: 0.01
  - model: chargoddard/rpguild-chatml-13b
    parameters:
      weight: 0.02
  - model: PygmalionAI/mythalion-13b
    parameters:
      weight: 0.08
  - model: CalderaAI/13B-Thorns-l2
    parameters:
      weight: 0.02
  - model: KoboldAI/LLaMA2-13B-Tiefighter
    parameters:
      weight: 0.20
dtype: float16
```
The third step was a merge between the two to retain the benefits of both as much as possible. This was performed using the dare merging technique.
```yaml
# task-arithmetic style
models:
  - model: EstopiaV9
    parameters:
      weight: 1
      density: 1
  - model: EstopiaV13
    parameters:
      weight: 0.05
      density: 0.30
merge_method: dare_ties
base_model: TheBloke/Llama-2-13B-fp16
parameters:
  int8_mask: true
dtype: bfloat16
```
# Model selection
- Undi95/UtopiaXL-13B
    - Solid all around base for models, with the ability to write longer responses and generally good retension to detail.
- Doctor-Shotgun/cat-v1.0-13b
    - A medical focused model which is added to focus a little more on the human responses, such as for psycology.
- PygmalionAI/mythalion-13b
    - A roleplay and instruct focused model, which improves attentiveness to character card details and the variety of responses
- Undi95/Emerhyst-13B
    - A roleplay but also longer form response model. It can be quite variable, but helps add to the depth and possible options the AI can respond with during narratives.
- CalderaAI/13B-Thorns-l2
    - A neutral and very attentive model. It is good at chat and following instructions, which help benefit these modes.
- KoboldAI/LLaMA2-13B-Tiefighter
    - A solid all around model, focusing on story writing and adventure modes. It provides all around benefits to creativity and the prose in models, along with adventure mode support.
- chargoddard/rpguild-chatml-13b
    - A roleplay model, which introduces new data and also improves the detail retention in longer narratives.
# Notes
- With the differing models inside, this model will not have perfect end of sequence tokens which is a problem many merges can share. While attempts have been made to minimise this, you may occasionally get oddly behaving tokens - this should be possible to resolve with a quick manual edit once and the model should pick up on it.
- Chat is one of the least tested areas for this model. It works fairly well, but it can be quite character card dependant.
- This is a narrative and prose focused model. As a result, it can and will talk for you if guided to do so (such as asking it to act as a co-author or narrator) within instructions or other contexts. This can be mitigated mostly by adding instructions to limit this, or using chat mode instead.
# Future areas
- Llava
    - Some success has been had with merging the llava lora on this. While no in depth testing has been performed, more narrative responses based on the images could be obtained - though there were drawbacks in the form of degraded performance in other areas, and hallucinations due to the fictional focus of this model.
- Stheno
    - A merge which has similar promise from Sao. Some merge attempts have been made between the two and were promising, but not entirely consistent at the moment. With some possible refinement, this could produce an even stronger model.
- DynamicFactor
    - All the merges used have been based on llama two in this merge, but a dare merge with dynamic factor (an attempted refinement of llama two) showed a beneficial improvement to the instruction abilities of the model, along with lengthy responses. It lost a little of the variety of responses, so perhaps if a balance of it could be added the instruction abilities and reasoning could be improved even further.