---
license: other
license_name: microsoft-research-license
license_link: https://huggingface.co/microsoft/Orca-2-13b/blob/main/LICENSE
pipeline_tag: text-generation
---

# OrcaMaid-13b

This is a merge of Microsoft's [Orca-2-13b](https://huggingface.co/microsoft/Orca-2-13b) and Undi and IkariDev's [Noromaid-v0.1.1-13b](https://huggingface.co/NeverSleep/Noromaid-13b-v0.1.1), with just a touch of Kal'tsit's [cat-v1.0](https://huggingface.co/Doctor-Shotgun/cat-v1.0-13b) mixed in.

The model recipe was as follows:
- Linear merge of **Orca-2-13b** @0.8 and **cat-v1.0-13b** @0.2 = OrcaCat-13b (no plans to release)
- Gradient SLERP merge of **Noromaid-v0.1.1** @0.5 and **OrcaCat-13b** @0.5 = OrcaMaid-13b

Both merges were done in FP32 rather than FP16, due to Orca being released as FP32. I didn't want to risk losing any precision.

The overall goal of this merge is to create a model that sounds uniquely human and natural, without sacrificing intelligence. ***Edit:** after some feedback from a few others, ranking on the Ayumi leaderboards, and more of my own testing, I believe I have succeeded as well as I reasonably could have hoped.*

The prompt format is Alpaca. You can use the standard format as shown, but for best results, you should customize the system prompt to your specific needs.

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{YOUR MESSAGE HERE}

### Response:
{BOT MESSAGE HERE}


```

### Misc. information
- BOS token is `<s>`
- EOS token is `</s>`
- Native context length is `4096`
- Base model is Llama 2
- Due to the inclusion of Orca-2-13b, the model is subject to the terms of the [Microsoft Research License](https://huggingface.co/microsoft/Orca-2-13b/blob/main/LICENSE)

### Thanks
- Thanks to [Charles Goddard](https://github.com/cg123) for his kind help with mergekit
- Thanks to [Undi](https://ko-fi.com/undiai) and [IkariDev](https://ikaridevgit.github.io/) for Noromaid
- Thanks to Kal'tsit for Cat. See her original reddit post: [Cat 1.0 is an uncensored, rp model aligned to be useful in all (even spicy)situations](https://www.reddit.com/r/LocalLLaMA/comments/17skxzq/cat_10_is_an_uncensored_rp_model_aligned_to_be/)