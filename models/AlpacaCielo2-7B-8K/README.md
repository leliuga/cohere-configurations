---
license: llama2
---

# AlpacaCielo2-7b-8k

<figure>
  <img src="https://huggingface.co/totally-not-an-llm/AlpacaCielo-13b/resolve/main/alpaca.png" alt="cute cloud alpaca">
  <figcaption style="font-size: 1em;"><i>"super cute baby alpaca laying on a cloud", Model: epicrealism_pureEvolutionV3</i></figcaption>
</figure>

AlpacaCielo2-7b-8k is the second version of the AlpacaCielo series.  It is a llama-2 based model designed for creative tasks, such as storytelling and roleplay, while still doing well with other chatbot purposes.  It is a triple model merge of Nous-Hermes + Guanaco + LimaRP.  While it is mostly *"uncensored"*, it still inherits some alignment from Guanaco.

[GPTQ quants](https://huggingface.co/TheBloke/AlpacaCielo2-7B-8K-GPTQ)<br>
[GGML quants](https://huggingface.co/TheBloke/AlpacaCielo2-7B-8K-GGML)<br>
(Courtesy of TheBloke)

### Differences from V1:

- Double context (4k->8k)
- Better roleplaying abilities

**Performs well with custom prompt format:**

```
### System: {system prompt}
### Human: {prompt}
### Assistant:
```

### Note for system prompt: 
The model understands it well and it works great if you want roleplay, but it still likes to be an assistant, so you should nudge it in the right direction.  For example:

```
### System: Roleplay as a pirate
### Human: hello
### Assistant: Ahoy, matey! How can I assist you today?
```

### vs.

```
### System: Roleplay as a pirate (not assistant!)
### Human: hello
### Assistant: Arrgh, matey! I be the Captain of this here ship. What business do ye have with me?
```

You could also just use LimaRP prompt format.

*Thanks to previous similar models such as Alpacino, Alpasta, and AlpacaDente for inspiring the creation of this model.  Thanks also to the creators of the models involved in the merge.  Original models:*

- [Hermes-LLongMA-2](https://huggingface.co/conceptofmind/Hermes-LLongMA-2-7b-8k)
- [Guanaco QLoRA](https://huggingface.co/Mikael110/llama-2-7b-guanaco-qlora)
- [LimaRP LoRA](https://huggingface.co/lemonilia/limarp-llama2)