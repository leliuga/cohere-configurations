---
language:
- en
datasets:
- PocketDoc/Floyd-Text-Adventures
- PocketDoc/Choose-Your-Story-Long-Text-Adventures
- openchat/openchat_sharegpt4_dataset
- 64bits/lima_vicuna_format
- gsm8k
- PocketDoc/Wizard-Vicuna-Refined
- PocketDoc/Guanaco-Unchained-Refined
license: apache-2.0
base_model: mistralai/Mistral-7B-v0.1
---

### Description:
This model is to test some of my data projects. It is a multipurpose chat / chat instruct hybrid model and can be seen as a prototype of Dan's PersonalityEngine Mk. 2. It has been trained on a wide array of one shot instructions, multi round instructions, role playing scenarios, and text adventure games. It is based on [Mistral-7b](https://huggingface.co/mistralai/Mistral-7B-v0.1)

### Prompt format: 
Pygmalion / Metharme

The prompt should start with the cursor on the same line directly after "<|model|>" with no space. The following are all valid formats and can be extended to as many rounds as desired.
```
<|system|>system message here<|user|>user message here<|model|>
```
```
<|system|>system message here<|user|>user message here<|model|>model message<|user|>user message here<|model|>
```
```
<|system|>system message here<|model|>
```
```
<|system|>system message here<|model|>model message<|user|>user message here<|model|>
```

#### Some example prompts and their resulting outputs:
##### Prompt:
```
<|system|>You are an assistant, help the user and use markdown when applicable.<|user|>What are the top games of the 90s?<|model|>
```
##### Output:
```
The 1990s was a golden era for video gaming, with countless iconic titles released across various platforms. Here's a list of some of the most popular and influential games from that decade:

- **Super Mario World** (SNES) - Released in 1990, this classic platformer introduced fans to Yoshi and featured memorable levels like Donut Plains and Star World.
- **Mortal Kombat** (Arcade/Sega Genesis) - This gory fighting game shocked audiences upon its release in 1992 with its violent finishers and secret characters like Sub-Zero and Scorpion.
- **Doom** (PC) - Id Software's groundbreaking first-person shooter revolutionized the genre by introducing fast-paced action and multiplayer deathmatches.
- **Street Fighter II** (Arcade/SNES) - Capcom's sequel to their original Street Fighter arcade game added new fighters such as Ryu's rival Ken Masters plus complex special moves like Dragon Punch and Hadouken.
- **Crash Bandicoot** (PlayStation) - Naughty Dog's mascot platforming series debuted in 1996 starring Crash who had to save his girlfriend Tawna from Dr Cortex's evil experiments on Aku Aku masks.
- **Final Fantasy VII** (PlayStation) - Square Enix's seventh installment took players through Midgar City where they joined Cloud Strife's eco-terrorist group AVALANCHE against Shinra Electric Power Company led by Sephiroth.
- **Resident Evil** (PlayStation) - The survival horror masterpiece created by Capcom saw Chris Redfield investigating Spencer Mansion alongside Jill Valentine while trying not get eaten alive by zombies!
```
##### Prompt:

```
<|system|>You are a Virtual Story Generator. You take the user's input and create an excellent and captivating story that goes in that direction. Use an abundance of sensory descriptions and eloquent prose.<|user|>Alpha Centauri has fallen, to the bears. This is a point of view tale about a soldier on the ground.<|model|>
```
##### Output:
```
The roar of the engines shook the very foundations of the planet as the massive warships descended upon Alpha Centauri. The air was thick with the stench of burning fuel and smoke, as if the entire world had been set ablaze. In the chaos, soldiers were scrambling for cover, desperately trying to protect themselves from the relentless barrage of firepower unleashed by their enemy.
Amidst the carnage, one soldier stood tall, his heart pounding against his chest like a drumbeat of death. He knew he had to do something, anything, to stop this nightmare from becoming reality. With trembling hands, he reached for his weapon, determined to make a stand against the overwhelming force that threatened to consume them all.
As the battle raged around him, the soldier felt a sudden surge of adrenaline coursing through his veins. It was almost as if some hidden power within him had awoken, granting him strength beyond measure. His eyes blazed with determination as he raised his gun and took aim at the first target he saw: a bear-like creature, its fur matted with blood and gore, advancing towards him with murderous intent.
With a deafening crack, the soldier fired off several rounds, striking the beast square in the chest. The impact sent it reeling backwards, but it did not fall. Instead, it let out a ferocious growl before charging forward once more. The soldier gritted his teeth and pulled the trigger again, emptying his clip into the creature's already battered frame.
But even as the bullets tore through its flesh, the beast seemed unfazed, continuing its advance towards the terrified soldier. Desperate for any advantage, he threw himself to the side just moments before the monster slammed into the spot where he had been standing mere seconds ago. As it turned back towards him, ready to deliver another brutal blow, the soldier could see that its eyes glowed with an otherworldly light.
```
##### Prompt:
```
<|system|>You are a professional editor with decades of experience, help the user with any task they have for you.<|user|>Can you rewrite this to flow better? "I knew I probably shouldnt have done that but oh well"<|model|>
```
##### Output:
```
I hesitated before making my decision, knowing full well it was not wise. However, in the end, I chose to proceed regardless.
```

# Some quick and  dirty training details:
- [<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="150" height="24"/>](https://github.com/OpenAccess-AI-Collective/axolotl)
- Sequence length: 4096
- Training time: 4 hours
- Hardware: 2x RTX 4090
- Training type: QLoRA
- PEFT R/A: 32/32


# Credits:
### Skein Text Adventure Data:
Thank you to the [Kobold AI](https://huggingface.co/KoboldAI) community for curating the Skein dataset, which is pivotal to this model's capabilities.
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_PocketDoc__Dans-TotSirocco-7b)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 56.92   |
| ARC (25-shot)         | 62.03          |
| HellaSwag (10-shot)   | 84.23    |
| MMLU (5-shot)         | 64.19         |
| TruthfulQA (0-shot)   | 46.49   |
| Winogrande (5-shot)   | 78.69   |
| GSM8K (5-shot)        | 13.27        |
| DROP (3-shot)         | 49.54         |