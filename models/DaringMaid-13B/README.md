---
license: cc-by-nc-4.0
language:
- en
pipeline_tag: text-generation
---

# DaringMaid-13B
My goal was to make a Noromaid that's smarter and better at following instructions. 

After trying a bunch of different recipes I think this one turned out pretty good

- I used [sequelbox/DynamicFactor](https://huggingface.co/sequelbox/DynamicFactor) as a base to as its supposed "improve overall knowledge, precise communication, conceptual understanding, and technical skill" over the base llama2.
- [NeverSleep/Noromaid](https://huggingface.co/NeverSleep/Noromaid-13b-v0.1.1) of course.
- [Undi95/Utopia](https://huggingface.co/Undi95/Utopia-13B) has been recommended again recently and its still really good so in the mixer it goes
- I liked [tavtav/Rose](https://huggingface.co/tavtav/Rose-20B) so i threw in a bit of [CalderaAI/Thorns](https://huggingface.co/CalderaAI/13B-Thorns-l2) 
- There was recently a model that tried to pass itself off as [Gryphe/MythoMax](https://huggingface.co/Gryphe/MythoMax-L2-13b), i made a merge with that model before it was revealed to be MythoMax and it turned out pretty good so i used it.

The .yml config files for mergekit with the exact merges can be found in the ["Recipe"](https://huggingface.co/Kooten/DaringMaid-13B/tree/main/Recipe) folder in the [fp16 repo](https://huggingface.co/Kooten/DaringMaid-13B) 

# Quants
EXL2: [8bpw](https://huggingface.co/Kooten/DaringMaid-13B-8bpw-exl2), [5bpw](https://huggingface.co/Kooten/DaringMaid-13B-5bpw-exl2), [4bpw](https://huggingface.co/Kooten/DaringMaid-13B-4bpw-exl2)

[GGUF](https://huggingface.co/Kooten/DaringMaid-13B-GGUF):
[Q3_K_M](https://huggingface.co/Kooten/DaringMaid-13B-GGUF/blob/main/DaringMaid-13B-Q3_K_M.gguf) - [Q4_K_M](https://huggingface.co/Kooten/DaringMaid-13B-GGUF/blob/main/DaringMaid-13B-Q4_K_M.gguf) - [Q5_K_M](https://huggingface.co/Kooten/DaringMaid-13B-GGUF/blob/main/DaringMaid-13B-Q5_K_M.gguf) 


## Prompt template:
I have been using Undi/Ikaris SillyTavern presets for Noromaid: [Context](https://files.catbox.moe/ifmhai.json), [Instruct](https://files.catbox.moe/ttw1l9.json).

### Alpaca:
```
Below is an instruction that describes a task. Write a response that appropriately completes the request. Do not include descriptions of non-visual qualities such as personality, movements, scents, mental traits, or anything which could not be seen in a still photograph. Do not write in full sentences. Prefix your description with the phrase 'full body portrait,'

### Instruction:
{prompt}

### Response:

```

### Contact
Kooten on discord.
