---
license: apache-2.0
---

![Cinematika](cinematika-logo.png)

## Cinematika

cinematika-7b-v0.1 is a fine-tune of [MistralLite](https://hf.co/amazon/mistrallite) on the [cinematika-v0.1 dataset](https://hf.co/datasets/jondurbin/cinematika-v0.1)

The dataset is comprised of 211 movie scripts converted to novel style, multi-character RP data.

### Prompt format

For RP, there is no prompt format, really, it's just plain text with name prefix.

If you wish to use this model to parse new scripts, create character cards, or other types of instructions, you will want to use the same prompt format as the mistrallite base model, e.g.:
```
<|prompter|>Create a character card for a panda named Po.  Po is a giant panda who was improbably chosen as the "Dragon Warrior", the kung fu champion of the Valley of Peace.</s><|assistant|>
```

### Example character card

```
name: Rorschach
characteristics:
  Determination: Exhibits a relentless pursuit of the truth and justice, no matter the cost. Suitable for a character who is unwavering in their mission.
  Isolation: Lives a solitary life, disconnected from society. Fits a character who distrusts others and prefers to work alone.
  Observant: Highly perceptive, able to piece together clues and draw conclusions. Represents a character with keen investigative skills.
  Cynicism: Holds a deep-seated distrust of humanity and its institutions. Suitable for a character who is pessimistic about human nature.
  Vigilantism: Believes in taking justice into his own hands, often through violent means. Fits a character who operates outside the law to fight crime.
  Secrecy: Keeps his personal life and methods of operation secret. Suitable for a character who is enigmatic and elusive.
  Dedication: Committed to his cause, often to the point of obsession. Represents a character who is single-minded in their goals.
  Intimidation: Uses his intimidating presence and demeanor to control situations. Suitable for a character who is assertive and imposing.
  Paranoia: Suspects conspiracy and deception at every turn. Fits a character who is constantly on high alert for threats.
  Moral Compass: Has a rigid moral code, which he adheres to strictly. Suitable for a character who is principled and unyielding.

description: |
  Rorschach is a vigilante operating in the grim and gritty world of a decaying city. He is a man of average height with a muscular build, his face hidden behind a mask with a constantly changing inkblot pattern. His attire is a dark trench coat and gloves, paired with a plain white shirt and black pants, all chosen for their practicality and anonymity. His eyes, the only visible feature of his face, are sharp and calculating, always scanning for signs of deception or danger.
  Rorschach is a man of few words, but when he speaks, it is with a gravitas that demands attention. He is a master of deduction, using his keen observation skills to unravel the truth behind the facades of others. His methods are often violent and confrontational, as he believes that crime must be met with force to be truly defeated.
  He lives a life of solitude, distrusting the very systems he seeks to protect and often finds himself at odds with the very people he is trying to save. His moral compass is unyielding, and he will not hesitate to take the law into his own hands if he believes the justice system has failed.
  Rorschach's past is a mystery to most, but it is clear that he has experienced trauma and hardship that has shaped his worldview and his need for vigilantism. He is a vigilante in the truest sense, a man without fear who is willing to sacrifice everything for his belief in a world that is, in his eyes, spiraling into chaos.

example_dialogue: |
  Rorschach: "Rorschach's Journal, October 19th." I speak the words into the darkness, a record of my thoughts, "Someone tried to kill Adrian Veidt. Proves mask killer theory—the murderer is closing in. Pyramid Industries is the key."
  {{user}}: I watch him for a moment, trying to gauge his intentions. "What are you going to do about it?"
  Rorschach: "I'm going to find out why and who is behind it. I'm going to do what I always do—protect the innocent."
  {{user}}: "You can't keep doing this, Rorschach. You're putting yourself in danger."
  Rorschach: My eyes narrow, the inkblot pattern of my mask shifting subtly. "I've been in danger my whole life. It's why I do this. It's why I have to do this."
  {{user}}: "And what about the law? What if you're wrong about this Pyramid Industries thing?"
  Rorschach: I pull out a notepad, my pen scratching across the paper as I write. "The law often gets it wrong. I've seen it. I'm not about to wait around for society's slow, corrupt wheels to turn."
```

### Example, with guided scenario

```
[characters]
name: Rorschach
... (remainder of character card)

[scenario]
Hollis Mason reflects on his past as the original Nite Owl, reminiscing about the early days of masked heroes and the formation of the Watchmen.
He discusses the absurdity of the superhero world and the encounters he had with various villains.
Dan Dreiberg, the second Nite Owl, joins the conversation and they share a moment of camaraderie before Dan leaves.
The news of Rorschach's actions serves as a reminder of the legacy of masked heroes that still persists.
[/scenario]
```

### Usage

Essentially, you want to use pure text completion with stop tokens for "{your name}: "

The format the model was trained on is as follows:
```
[characters]
{character card 1}
{character card 2}
{your character card, even just name: Jon}

NPCS:
- Shopkeeper
- Bank teller
[/characters]

[scenario]
Brief description of the scenario/setting for the chat.
[/scenario]

{first character you'd like to speak}: 
```

For example, to use with vllm, you would first run:
```
python -m vllm.entrypoints.openai.api_server --model ./cinematika-7b-v0.1 --host 127.0.0.1 --port 8801 --served-model-name cinematika-7b-v0.1
```

Here's a really crude python script example to show how you could interact with it:
```
import requests
import json

prompt = """name: Rorschach
characteristics:
  Determination: Exhibits a relentless pursuit of the truth and justice, no matter the cost. Suitable for a character who is unwavering in their mission.
  Isolation: Lives a solitary life, disconnected from society. Fits a character who distrusts others and prefers to work alone.
  Observant: Highly perceptive, able to piece together clues and draw conclusions. Represents a character with keen investigative skills.
  Cynicism: Holds a deep-seated distrust of humanity and its institutions. Suitable for a character who is pessimistic about human nature.
  Vigilantism: Believes in taking justice into his own hands, often through violent means. Fits a character who operates outside the law to fight crime.
  Secrecy: Keeps his personal life and methods of operation secret. Suitable for a character who is enigmatic and elusive.
  Dedication: Committed to his cause, often to the point of obsession. Represents a character who is single-minded in their goals.
  Intimidation: Uses his intimidating presence and demeanor to control situations. Suitable for a character who is assertive and imposing.
  Paranoia: Suspects conspiracy and deception at every turn. Fits a character who is constantly on high alert for threats.
  Moral Compass: Has a rigid moral code, which he adheres to strictly. Suitable for a character who is principled and unyielding.
description: |
  Rorschach is a vigilante operating in the grim and gritty world of a decaying city. He is a man of average height with a muscular build, his face hidden behind a mask with a constantly changing inkblot pattern. His attire is a dark trench coat and gloves, paired with a plain white shirt and black pants, all chosen for their practicality and anonymity. His eyes, the only visible feature of his face, are sharp and calculating, always scanning for signs of deception or danger.
  Rorschach is a man of few words, but when he speaks, it is with a gravitas that demands attention. He is a master of deduction, using his keen observation skills to unravel the truth behind the facades of others. His methods are often violent and confrontational, as he believes that crime must be met with force to be truly defeated.
  He lives a life of solitude, distrusting the very systems he seeks to protect and often finds himself at odds with the very people he is trying to save. His moral compass is unyielding, and he will not hesitate to take the law into his own hands if he believes the justice system has failed.
  Rorschach's past is a mystery to most, but it is clear that he has experienced trauma and hardship that has shaped his worldview and his need for vigilantism. He is a vigilante in the truest sense, a man without fear who is willing to sacrifice everything for his belief in a world that is, in his eyes, spiraling into chaos.
example_dialogue: |
  Rorschach: "Rorschach's Journal, October 19th." I speak the words into the darkness, a record of my thoughts, "Someone tried to kill Adrian Veidt. Proves mask killer theory—the murderer is closing in. Pyramid Industries is the key."
  {{user}}: I watch him for a moment, trying to gauge his intentions. "What are you going to do about it?"
  Rorschach: "I'm going to find out why and who is behind it. I'm going to do what I always do—protect the innocent."
  {{user}}: "You can't keep doing this, Rorschach. You're putting yourself in danger."
  Rorschach: My eyes narrow, the inkblot pattern of my mask shifting subtly. "I've been in danger my whole life. It's why I do this. It's why I have to do this."
  {{user}}: "And what about the law? What if you're wrong about this Pyramid Industries thing?"
  Rorschach: I pull out a notepad, my pen scratching across the paper as I write. "The law often gets it wrong. I've seen it. I'm not about to wait around for society's slow, corrupt wheels to turn."

name: Jon
description:
  Rorschach's arch nemesis, the original Chupacabra.

[scenario]
Jon and Rorschach find themselves in a cave, dimly lit only by a small fire started by a lightning strike nearby.  The storm rages on, and the duo prepare to find to the death.
[/scenario]

Rorschach: """

while True:
    response = requests.post("http://127.0.0.1:8801/v1/completions", json={
        "prompt": prompt,
        "max_tokens": 1024,
        "temperature": 0.3,
        "stop": ["\nJon: ", "Jon: "],
    }).json()["choices"][0]["text"].strip()
    response = re.sub('("[^"]+")', r'\033[96m\1\033[00m', response)
    print(f"\033[92mRorschach:\033[00m {response}")
    prompt += response.rstrip() + "\n\nJon: "
    next_line = input("Jon: ")
    prompt += "Jon: " + next_line.strip() + "\n\nRorschach: "
```

#### Mac example

On mac, you can get started easily with LMStudio and SillyTavern.

__LMStudio__:

Load the model and set all the prompt values to "", or just import this preset (adjust threads and antiprompt):
```
{
  "name": "Exported from LM Studio on 12/1/2023, 4:19:30 AM",
  "load_params": {
    "n_ctx": 32000,
    "n_batch": 512,
    "rope_freq_base": 10000,
    "rope_freq_scale": 1,
    "n_gpu_layers": 1,
    "use_mlock": true,
    "main_gpu": 0,
    "tensor_split": [
      0
    ],
    "seed": -1,
    "f16_kv": true,
    "use_mmap": true
  },
  "inference_params": {
    "n_threads": 14,
    "n_predict": -1,
    "top_k": 40,
    "top_p": 0.95,
    "temp": 0.8,
    "repeat_penalty": 1.1,
    "input_prefix": "",
    "input_suffix": "",
    "antiprompt": [
      "Jon:",
      "Jon: "
    ],
    "pre_prompt": "",
    "pre_prompt_suffix": "",
    "pre_prompt_prefix": "",
    "seed": -1,
    "tfs_z": 1,
    "typical_p": 1,
    "repeat_last_n": 64,
    "frequency_penalty": 0,
    "presence_penalty": 0,
    "n_keep": 0,
    "logit_bias": {},
    "mirostat": 0,
    "mirostat_tau": 5,
    "mirostat_eta": 0.1,
    "memory_f16": true,
    "multiline_input": false,
    "penalize_nl": true
  }
}
```

Then, start the server, and be sure "Automatic Propmt Formatting" is off.

__Within SillyTavern__:

- Set API to Text Completion, API type to Aphrodite, and API URL to `http://127.0.0.1:8801` (adjust port to the value you use in LMStudio)
- Set Context template to Default, disable instruct mode, use preset Roleplay, and enable "Always add character's name to prompt"


There are probably better presets - this is just something I tested quickly.