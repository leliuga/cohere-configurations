---
license: cc-by-nc-4.0
pipeline_tag: text-generation
---

# Pirouette-7b

This is a gradient SLERP merge of Eric Hartford's [dolphin-2.1-mistral-7b](https://huggingface.co/ehartford/dolphin-2.1-mistral-7b) merged with Undi and IkariDev's [Noromaid-v0.1.1-13b](https://huggingface.co/NeverSleep/Noromaid-13b-v0.1.1).

The goal of this merge is to retain most of the brain of Dolphin, with a little added flair from Noromaid.

The prompt format is Alpaca. You can use the standard format as shown, but for best results, I strongly recommend customizing the system prompt to your specific needs.

**You might need to add `Instruction:` to your list of stopping strings, due to Dolphin not being quite familiar with Alpaca.**

```
You are Pirouette, a personable and friendly AI assistant.

### Instruction:
{YOUR MESSAGE HERE}

### Response:
{BOT MESSAGE HERE}


```

### Misc. information
- BOS token is `<s>`
- EOS token is `</s>`
- Native context length is `8192`
- Functional context length extended to 32768 via RoPE with decreased perplexity, see [here](https://github.com/ggerganov/llama.cpp/issues/3867#issuecomment-1789054205)
- Base model is Mistral v0.1

### Thanks
- Thanks to [Eric Hartford](https://ko-fi.com/erichartford) for Dolphin
- Thanks to [Undi](https://ko-fi.com/undiai) and [IkariDev](https://ikaridevgit.github.io/) for Noromaid