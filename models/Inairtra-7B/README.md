---
license: apache-2.0
---

<p align="center">
<!--   <img src="./assets/Core1000AIIMG.png"/> -->
  <p align="center" style="font-size: 26px"><b>Inairtra-7B</b></p>
  <p align="center" style="font-size: 14px">Model Size: 7B</p>
</p>

<p align="center">
  <img src="./assets/SmallBronyaLogo.png" style="width: 45%;">
</p>

<p align="center" style="font-size: 20px">A <b>experimental</b> (and beginner) model merge using Intel's Neural Chat 7B</p>

## Model Details

Trained on: **Intel Xeon E5-2693v3 | NVIDIA RTX 2080 Ti | 128 GB DDR4 *(yes I'm poor :( )***

The Inairtra-7B LLM is a LLM made by Bronya Rand (bronya_rand / Bronya-Rand) as a beginning learning model to merging models using [MergeKit](https://github.com/cg123/mergekit) and GGUF quantization. This model is based off Intel's [Neural Chat 7B V3.1](https://huggingface.co/Intel/neural-chat-7b-v3-1) as the base model along with three additional Mistral models.

The Inairtra-7B architecture is based off: [**Mistral**](https://huggingface.co/mistralai/Mistral-7B-v0.1)

The models used to create the Inairtra-7B are as follows:

- Intel's Neural Chat 7B V3.1 ([Intel/neural-chat-7b-v3-1](https://huggingface.co/Intel/neural-chat-7b-v3-1))
- Teknium's Airoboros Mistral 2.2 7B ([teknium/airoboros-mistral2.2-7b](https://huggingface.co/teknium/airoboros-mistral2.2-7b))
- Maywell's Synatra 7B V0.3 RP ([maywell/Synatra-7B-v0.3-RP](https://huggingface.co/maywell/Synatra-7B-v0.3-RP))

## Prompt

The Inairtra-7B *should* (but unsure) support the same prompts as featured in Intel's Neural Chat, Airoboros Mistral and Synatra.

### For Intel
```
### System:
{system}
### User:
{usr}
### Assistant:
```

### For Airoboros
```
USER: <prompt>
ASSISTANT:
```

## Benchmarks?
I have no idea how to do them. You are welcome to make your own.

## Ethical Considerations and Limitations
The intended use-case for the Inairtra-7B LLM is for fictional writing/roleplay solely for personal entertainment purposes. Any other sort of usage outside of this is out of scope of my intentions and the LLM itself.

The Inairtra-7B LLM has been merged with models which are uncensored/unfiltered. The LLM can produce content, including but not limited to, content that may be NSFW for those under the age of eighteen, content that may be illegal in certain states/countries, offensive content, etc. 

The Inairtra-7B LLM is not designed to produce the most accurate information. It may produce incorrect data like all other AI models.

### Disclaimer
The license on this model does not constitute legal advice. I am not responsible for the actions of third parties (services/users/etc.) who use this model and distribute it for others. Please cosult an attorney before using this model for commercial purposes.

