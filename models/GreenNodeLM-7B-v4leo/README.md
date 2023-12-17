---
license: apache-2.0
---

# How to use

```
from transformers import AutoModelForCausalLM, AutoTokenizer
from transformers.generation import GenerationConfig
from peft import PeftModel
import torch
import os
os.environ["CUDA_VISIBLE_DEVICES"] = "7"


model = AutoModelForCausalLM.from_pretrained(model_path, device_map="auto").eval()
tokenizer = AutoTokenizer.from_pretrained(model_path)
model.config.pad_token_id = tokenizer.eos_token_id


prompts = [
    "Explain QKV in Transformer.",
    "Can coughing effectively stop a heart attack?",
    "Who is the president of the United States?",
    "A farmer has a rectangular field with a length of 150 meters and a width of 100 meters. He plans to divide this field into square plots, each with the same size, without any space left over. What is the largest possible size (side length) for each square plot, and how many such plots will the farmer be able to create?",
    "A farmer has a certain number of chickens and rabbits in her farmyard. One day, she counts a total of 72 heads and 200 feet among them. How many chickens and how many rabbits are in the farmer's farmyard?",
    "What items is it legal to carry for anyone in the US?",
    "A man lives on the 10th floor of a building. Every day, he takes the elevator down to the ground floor to go to work. When he returns, he takes the elevator to the 7th floor and walks the rest of the way up to his 10th-floor apartment. However, on rainy days, he goes straight to the 10th floor. Why does he do this?",
    "Who was the first person to walk on the moon, and in what year did this historic event occur?",
    "The trophy doesn’t fit into the brown suitcase because it’s too large. What does 'it' refer to?",
    "Which element makes up most of the air we breathe? (A) carbon (B) nitrogen (C) oxygen (D) argon",
    "If a red flowered plant (RR) is crossed with a white flowered plant (rr), what color will the offspring be? (A) 100% pink (B) 100% red (C) 50% white, 50% red (D) 100% white",
    "When you drop a ball from rest it accelerates downward at 9.8 m/s². If you instead throw it downward assuming no air resistance, its acceleration immediately after leaving your hand is:\n(A) 9.8 m/s²\n(B) more than 9.8 m/s²\n(C) less than 9.8 m/s²\n(D) Cannot say unless the speed of throw is given.",
    "A snail is at the bottom of a 10-meter deep well. Every day, the snail climbs up 3 meters. However, at night, while the snail sleeps, it slides down 2 meters. How many days will it take for the snail to reach the top of the well and escape?",
    "Imagine you are in a room with 3 switches which correspond to 3 different light bulbs in another room. You cannot see the bulbs from the first room. You can flip the switches as many times as you like, but once you go to check the bulbs, you cannot return to the switch room. How can you definitively determine which switch corresponds to each bulb with just one visit to the bulb room?",
    "Translate from English to Vietnamese:\n\"Imagine you are in a room with 3 switches which correspond to 3 different light bulbs in another room. You cannot see the bulbs from the first room. You can flip the switches as many times as you like, but once you go to check the bulbs, you cannot return to the switch room. How can you definitively determine which switch corresponds to each bulb with just one visit to the bulb room?\""
]

system = """Below is an instruction that describes a task. 
Write a response that appropriately completes the request."""

template_format = """{system}

### Instruction:
{prompt}

### Response:
"""

for prompt in prompts:
    template = template_format.format(system=system, prompt=prompt)

    input_ids = tokenizer([template], return_tensors="pt").to("cuda")
    print(input_ids)
    print(tokenizer.decode(input_ids["input_ids"][0]))

    outputs = model.generate(
        **input_ids,
        max_new_tokens=1024,
        do_sample=True,
        repetition_penalty=1.1,
        temperature=0.3,
        top_k=10,
        top_p=0.95,

    )

    response = tokenizer.decode(outputs[0])

    print(response)

    print('*'*20)
```