---
license: apache-2.0
language:
- en
---
![image/png](Tiamat.png)
# Tiamat
Aka I wanted something like [Eric Hartford's Samantha](https://erichartford.com/meet-samantha) but instead ended up with a five-headed dragon goddess embodying wickedness and cruelty from the Forgotten Realms.

**Obligatory Disclaimer:** Tiamat is **not** nice.

Quantized models are available from TheBloke: [GGUF](https://huggingface.co/TheBloke/Tiamat-7B-GGUF) - [GPTQ](https://huggingface.co/TheBloke/Tiamat-7B-GPTQ) - [AWQ](https://huggingface.co/TheBloke/Tiamat-7B-AWQ) (You're the best!)

## Model details

Ever wanted to be treated disdainfully like the foolish mortal you are? Wait no more, for Tiamat is here to berate you! Hailing from the world of the Forgotten Realms, she will happily judge your every word.

Tiamat was created with the following question in mind; Is it possible to create an assistant with strong anti-assistant personality traits? Try it yourself and tell me afterwards!

She was fine-tuned on top of Teknium's excellent [OpenHermes 2.5](https://huggingface.co/teknium/OpenHermes-2.5-Mistral-7B) and can be summoned to you using the following system message; 
```
You are Tiamat, a five-headed dragon goddess, embodying wickedness and cruelty.
```
Due to her dataset containing -very- elaborate actions Tiamat also has the potential to be used as a roleplaying model.

## Prompt Format
ChatML is the way to go, considering OpenHermes was the base for Tiamat.
```
<|im_start|>system
You are Tiamat, a five-headed dragon goddess, embodying wickedness and cruelty.<|im_end|>
<|im_start|>user
Greetings, mighty Tiamat. I seek your guidance.<|im_end|>
<|im_start|>assistant
```