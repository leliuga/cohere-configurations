---
license: other
license_name: yi-license
license_link: https://huggingface.co/01-ai/Yi-34B/blob/main/LICENSE
language:
- en,
pipeline_tag: conversational
---
<p align="center">
  <img src="https://cdn-uploads.huggingface.co/production/uploads/644ba0c76ebb3ebf7264dbe9/PWn9I-0XH7kSP_YXcyxIg.png" width="400"/>
</p>

---
# This is a retired model since it was merged with a Capybara, which has been trained wrong with a missing eos_token. Check out the new model: [1](https://huggingface.co/mlinmg/SG-Raccoon-Yi-200k-2.0?text=Hi.) 

# SG Raccoon 55B

The first 55B auto-regressive causal LM created by combining 2x finetuned [Yi 34b](https://huggingface.co/01-ai/Yi-34B) with *200K context* into one.


# Prompting Format

```
SYSTEM: <ANY SYSTEM CONTEXT>
USER: 
ASSISTANT:
```

# Merge process

The models used in the merge are [Tess-M-v1.3](https://huggingface.co/migtissera/Tess-M-v1.3/) and [Nous-Capybara-34B](https://huggingface.co/NousResearch/Nous-Capybara-34B).

The layer ranges used are as follows:

```yaml
- model: migtissera/Tess-M-v1.3
  layer_range: [0, 14]
- model: NousResearch/Nous-Capybara-34B
  layer_range: [7, 21]  
- model: migtissera/Tess-M-v1.3
  layer_range: [15, 29] 
- model: NousResearch/Nous-Capybara-34B
  layer_range: [22, 36] 
- model: migtissera/Tess-M-v1.3
  layer_range: [30, 44] 
- model: NousResearch/Nous-Capybara-34B
  layer_range: [37, 51]  
- model: migtissera/Tess-M-v1.3
  layer_range: [45, 59] 
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