---
license: cc-by-nc-4.0
tags:
- chat
- roleplay
- storywriting
---

# chronos-70b-v2

This is the FP16 PyTorch / HF version of **chronos-70b-v2** based on the **Llama v2 Base** model. This version will **not fit on a consumer GPU**, use a quantized type of model from those linked below!

Big thank you to the Pygmalion team for providing compute. Reach out to me if you would like individual credit.

This model is primarily focused on chat, roleplay, storywriting, with significantly improved reasoning and logic. It does not have any form of censorship, please use responsibly.

Chronos can generate very long outputs with coherent text, largely due to the human inputs it was trained on, and it supports context length up to 4096 tokens.

## License

This model is strictly [*non-commercial*](https://creativecommons.org/licenses/by-nc/4.0/) (**cc-by-nc-4.0**) use only which takes priority over the **LLAMA 2 COMMUNITY LICENSE AGREEMENT**. If you'd like to discuss using it for your business, contact Elinas through Discord **elinas**, or X (Twitter) **@officialelinas**. 

The "Model" is completely free (ie. base model, derivates, merges/mixes) to use for non-commercial purposes as long as the the included **cc-by-nc-4.0** license in any parent repository, and the non-commercial use statute remains, regardless of other models' licences. 
At the moment, only 70b models released will be under this license and the terms may change at any time (ie. a more permissive license allowing commercial use).

## Model Usage

This model uses Alpaca formatting, so for optimal model performance, use it to start the dialogue or story, and if you use a frontend like SillyTavern ENABLE Alpaca instruction mode:

```
### Instruction:
Your instruction or question here.
### Response:
```
Not using the format will make the model perform significantly worse than intended.

## Tips

Sampling and settings can make a significant difference for this model, so play around with them. I was also informed by a user that if you are using **KoboldCPP** that using the flag
`--unbantokens` may improve model performance **significantly**. This has not been tested by myself, but that is something to keep in mind.

## Quantized Versions for Consumer GPU Usage

[LlamaCPP Versions provided by @TheBloke](https://huggingface.co/TheBloke/Chronos-70B-v2-GGUF)

[GPTQ Quantized Versions provided by @TheBloke](https://huggingface.co/TheBloke/Chronos-70B-v2-GPTQ)


**Support Development of New Models**
<a href='https://ko-fi.com/Q5Q6MB734' target='_blank'><img height='36' style='border:0px;height:36px;' 
src='https://storage.ko-fi.com/cdn/kofi1.png?v=3' border='0' alt='Support Development' /></a>
