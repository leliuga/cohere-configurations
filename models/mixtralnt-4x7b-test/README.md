---
  license: cc-by-nc-4.0
---

# Mixtraln't 4x7B

Oh boy, a new model architecture in Transformers! Time to do profane things with it.

What if instead of training a MoE from scratch, we took some pre-trained Mistral models and shoved them in a little clown car?


Uses parts from the following models:
* [Q-bert/MetaMath-Cybertron-Starling](https://huggingface.co/Q-bert/MetaMath-Cybertron-Starling)
* [NeverSleep/Noromaid-7b-v0.1.1](https://huggingface.co/NeverSleep/Noromaid-7b-v0.1.1)
* [teknium/Mistral-Trismegistus-7B](https://huggingface.co/teknium/Mistral-Trismegistus-7B)
* [meta-math/MetaMath-Mistral-7B](https://huggingface.co/meta-math/MetaMath-Mistral-7B)
* [PocketDoc/Dans-AdventurousWinds-Mk2-7b](https://huggingface.co/PocketDoc/Dans-AdventurousWinds-Mk2-7b)


Works and generates coherent text. The big question here is if the hack I used to populate the MoE gates works well enough to take advantage of all of the experts. Let's find out!

Prompt format: maybe alpaca??? or chatml??? life is full of mysteries
