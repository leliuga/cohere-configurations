---
language:
- en,
pipeline_tag: conversational

license: other
license_name: yi-license
license_link: https://huggingface.co/01-ai/Yi-34B/blob/main/LICENSE
---

<p align="center">
  <img src="https://cdn-uploads.huggingface.co/production/uploads/644ba0c76ebb3ebf7264dbe9/PWn9I-0XH7kSP_YXcyxIg.png" width="400"/>
</p>

--- 

# SG Raccoon Yi 55B

The first 55B auto-regressive causal LM created by combining 2x finetuned [Yi 34b](https://huggingface.co/01-ai/Yi-34B) into one.


# Prompting Format

```
single-turn: <|startoftext|>Human: Hello!\n\nAssistant: <|endoftext|>

multi-turn: <|startoftext|>Human: Hello!\n\nAssistant: <|endoftext|>Hi!<|endoftext|>Human: How are you?\n\nAssistant: <|endoftext|>target2<|endoftext|>
```

# Merge process

The models used in the merge are [dolphin-2_2-yi-34b](https://huggingface.co/ehartford/dolphin-2_2-yi-34b) and [OrionStar-Yi-34B-Chat-Llama](https://huggingface.co/OrionStarAI/OrionStar-Yi-34B-Chat-Llama).

The layer ranges used are as follows:

```yaml
- range 0, 16
OrionStar-Yi-34B-Chat
- range 8, 24
dolphin-2_2-yi-34b
- range 17, 32
OrionStar-Yi-34B-Chat
- range 25, 40
dolphin-2_2-yi-34b
- range 33, 48
OrionStar-Yi-34B-Chat
- range 41, 56
dolphin-2_2-yi-34b
- range 49, 64
OrionStar-Yi-34B-Chat
- range 57, 72
dolphin-2_2-yi-34b
- range 65, 80
OrionStar-Yi-34B-Chat
```
# Tips

Being a Yi model, try disabling the BOS token and/or running a lower temperature with MinP (and no other samplers) if output doesn't seem right. Yi tends to run "hot" by default.

Sometimes the model "spells out" the stop token as </s> like Capybara, so you may need to add </s> as an additional stopping condition.

# Benchmarks
Coming soon.

# Acknowledgements
- Special thanks to [MSS](https://milanosamplesale.com/) for sponsoring this project

- [@chargoddard](https://huggingface.co/chargoddard) for developing the framework used to merge the model - [mergekit](https://github.com/cg123/mergekit).

- Great thanks to [@Undi95](https://huggingface.co/Undi95) for helping figuring out model merge options

- Also credits to the [01-ai](https://huggingface.co/01-ai) team for their amazing models

- This merged model is inspired by [Goliath 120B](https://huggingface.co/alpindale/goliath-120b)
