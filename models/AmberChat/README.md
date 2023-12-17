---
license: apache-2.0
datasets:
- WizardLM/WizardLM_evol_instruct_V2_196k
- icybee/share_gpt_90k_v1
language:
- en
widget:
  - example_title: "example 1"
    text: "How do I mount a tv to drywall safely?"
    output:
      text: "Mounting a TV to drywall safely is crucial to ensure the longevity of your TV and the wall. Here's a step-by-step guide on how to do it safely:\n 1. Choose the right mount: Choose a mount that is compatible with your TV and the wall. Make sure the mount is rated for drywall and not just wood or concrete.\n 2. Prepare the wall: Before mounting the TV, make sure the wall is clean and free from any debris or paint. Use a stud finder to locate the studs in the wall.\n 3. Install the mount: Follow the manufacturer's instructions for mounting the TV to the wall. Use appropriate hardware and tools to secure the mount to the wall.\n 4. Level the TV: Level the TV on the mount using a spirit level or a leveling kit provided by the mount manufacturer.\n 5. Attach the TV to the mount: Attach the TV to the mount using the appropriate hardware and tools. Tighten the bolts and screws to ensure the TV is securely attached.\n 6. Connect the cables: Connect the TV cables to the appropriate ports on the back of the TV and the mount.\n 7. Test the mount: Test the mount to ensure it's secure and stable. Adjust the mount as needed to ensure the TV is level and secure.\n Mounting a TV to drywall safely is crucial to avoid damaging the wall or the TV. Follow these steps carefully and use appropriate tools and hardware to ensure a secure and stable installation."
  - example_title: "example 2"
    text: "Happy is to sad as calm is to _."
    output:
      text: "The adjective that can be used to describe the opposite of calm is \"anxious\" or \"stressed.\" So, from happy to sad, we can say that happy is to sad as calm is to anxious or stressed."
library_name: transformers
pipeline_tag: text-generation
tags:
- nlp
- llm
---
# AmberChat


We present AmberChat, an instruction following model finetuned from [LLM360/Amber](https://huggingface.co/LLM360/Amber).

## Model Description

- **Model type:** Language model with the same architecture as LLaMA-7B
- **Language(s) (NLP):** English
- **License:** Apache 2.0
- **Resources for more information:**
  - [Metrics](https://github.com/LLM360/Analysis360)
  - [Fully processed Amber pretraining data](https://huggingface.co/datasets/LLM360/AmberDatasets)


# Loading AmberChat 

```python
import torch
from transformers import LlamaTokenizer, LlamaForCausalLM

tokenizer = LlamaTokenizer.from_pretrained("LLM360/AmberChat")
model = LlamaForCausalLM.from_pretrained("LLM360/AmberChat")

#template adapated from fastchat
template= "A chat between a curious human and an artificial intelligence assistant. The assistant gives helpful, detailed, and polite answers to the human's questions.\n### Human: Got any creative ideas for a 10 year old’s birthday?\n### Assistant: Of course! Here are some creative ideas for a 10-year-old's birthday party:\n1. Treasure Hunt: Organize a treasure hunt in your backyard or nearby park. Create clues and riddles for the kids to solve, leading them to hidden treasures and surprises.\n2. Science Party: Plan a science-themed party where kids can engage in fun and interactive experiments. You can set up different stations with activities like making slime, erupting volcanoes, or creating simple chemical reactions.\n3. Outdoor Movie Night: Set up a backyard movie night with a projector and a large screen or white sheet. Create a cozy seating area with blankets and pillows, and serve popcorn and snacks while the kids enjoy a favorite movie under the stars.\n4. DIY Crafts Party: Arrange a craft party where kids can unleash their creativity. Provide a variety of craft supplies like beads, paints, and fabrics, and let them create their own unique masterpieces to take home as party favors.\n5. Sports Olympics: Host a mini Olympics event with various sports and games. Set up different stations for activities like sack races, relay races, basketball shooting, and obstacle courses. Give out medals or certificates to the participants.\n6. Cooking Party: Have a cooking-themed party where the kids can prepare their own mini pizzas, cupcakes, or cookies. Provide toppings, frosting, and decorating supplies, and let them get hands-on in the kitchen.\n7. Superhero Training Camp: Create a superhero-themed party where the kids can engage in fun training activities. Set up an obstacle course, have them design their own superhero capes or masks, and organize superhero-themed games and challenges.\n8. Outdoor Adventure: Plan an outdoor adventure party at a local park or nature reserve. Arrange activities like hiking, nature scavenger hunts, or a picnic with games. Encourage exploration and appreciation for the outdoors.\nRemember to tailor the activities to the birthday child's interests and preferences. Have a great celebration!\n### Human: {prompt}\n### Assistant:"

prompt = "How do I mount a tv to drywall safely?"

input_str = template.format(prompt=prompt)
input_ids = tokenizer(input_str, return_tensors="pt").input_ids
outputs = model.generate(input_ids, max_length=1000)
print(tokenizer.batch_decode(outputs[:, input_ids.shape[1]:-1])[0].strip())
```

Alternatively, you may use [FastChat](https://github.com/lm-sys/FastChat):
```bash
python3 -m fastchat.serve.cli --model-path LLM360/AmberChat
```

# AmberChat Finetuning Details

## DataMix
| Subset      | Number of rows |  License   |
| ----------- | ----------- | ----------- |
| WizardLM/WizardLM_evol_instruct_V2_196k      | 143k       |  |
| icybee/share_gpt_90k_v1   | 90k        | cc0-1.0 |
| Total | 233k |  |

## Hyperparameters
| Hyperparameter      | Value |
| ----------- | ----------- |
| Total Parameters      | 6.7B       |
| Hidden Size   | 4096        |
| Intermediate Size (MLPs)   | 11008        |
| Number of Attention Heads   | 32        |
| Number of Hidden Lyaers  | 32        |
| RMSNorm ɛ  | 1e^-6        |
| Max Seq Length   | 2048        |
| Vocab Size | 32000 |

| Training Hyperparameter      | Value |
| ----------- | ----------- |
| learning_rate      | 2e-5       |
| num_train_epochs  |  3        |
| per_device_train_batch_size   | 2        |
| gradient_accumulation_steps  | 16        |
| warmup_ratio | 0.04      |
| model_max_length | 2048     |


# Evaluation

| Model                                                | MT-Bench                                                  | 
|------------------------------------------------------|------------------------------------------------------------|
| LLM360/Amber 359 | 2.48750 | 
| **LLM360/AmberChat** | **5.428125** |

# Citation

**BibTeX:**

```bibtex
@article{xxx,
  title={XXX},
  author={XXX},
  journal={XXX},
  year={2023}
}
```