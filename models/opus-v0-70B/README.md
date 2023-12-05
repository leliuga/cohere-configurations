---
language:
  - en
pipeline_tag: text-generation
---

# DreamGen Opus V0 70B

**DreamGen Opus** is a family of **uncensored** models fine-tuned for **(steerable) story writing** and the model also works great for **chat / RP**.
The DreamGen Opus V0 70B model is derived from [meta-llama/Llama-2-70b-hf](https://huggingface.co/meta-llama/Llama-2-70b-hf).

You can **try the Opus V0 70B** (AWQ) model for free on [dreamgen.com](https://dreamgen.com).

Quantized versions:

- AWQ: [dreamgen/opus-v0-70b-awq](https://huggingface.co/dreamgen/opus-v0-70b-awq)
- GGUF: [dreamgen/opus-v0-70b-gguf](https://huggingface.co/dreamgen/opus-v0-70b-gguf) and [TheBloke/opus-v0-70B-GGUF](https://huggingface.co/TheBloke/opus-v0-70B-GGUF)
- GPTQ: [TheBloke/opus-v0-70B-GPTQ](https://huggingface.co/TheBloke/opus-v0-70B-GPTQ)
- EXL2:
  - [LoneStriker/opus-v0-70b-4.0bpw-h6-exl2](https://huggingface.co/LoneStriker/opus-v0-70b-4.0bpw-h6-exl2)
  - [LoneStriker/opus-v0-70b-3.0bpw-h6-exl2](https://huggingface.co/LoneStriker/opus-v0-70b-3.0bpw-h6-exl2)
  - [LoneStriker/opus-v0-70b-2.4bpw-h6-exl2](https://huggingface.co/LoneStriker/opus-v0-70b-2.4bpw-h6-exl2)

Other sizes:

- 7B: [dreamgen/opus-v0-7b](https://huggingface.co/dreamgen/opus-v0-7b)

## Prompting

Please see the [official documentation](https://dreamgen.com/docs/stories) for more detailed guide, including how to prompt the model for chat / RP.

The (collaborative / steerable) story writing task teaches the model to respect `<setting>` and `<instruction>` inserted into the prompt.

Example prompt:

```
<setting>
(Setting provides general overview of the story and characters)
This story is a twist on the traditional Little Red Riding Hood story.
In this variation, the Little Red Riding Hood and her grandma are secretely werevoles.
</setting>

(Previous part of the story, potentially empty)

<instruction>
(Setting tells the model what should happen in the next few sentences / paragraphs)
The Little Red Riding hood confronts The Big Bad Wolf, transforming into her wolf form.
</instruction>
```

## Dataset

The fine-tuning dataset consisted of >1M tokens of collaborative writing task examples, each example being up to 4096 tokens. On top of that, >20M tokens of more general, but less instructed examples were included to help preserve generalization.

All prose in the dataset is from actual humans, not AI generated.

## Community

Join the DreamGen community on [**Discord**](https://dreamgen.com/discord), or follow our [**X/Twitter account**](https://dreamgen.com/twitter) for new model releases and other news.
We will soon be releasing models with longer context window, as well as models specifically fine-tuned for character chat & roleplay.

Help us shape the future of DreamGen.

## Running the model

The model is should be compatible with any software that supports [meta-llama/Llama-2-70b-hf](https://huggingface.co/meta-llama/Llama-2-70b-hf).
Note that because this is a 70B model, the resource requirements are large. You can try the quantized versions linked at the top, but expect a quality drop.

### Running on DreamGen.com (free)

You can try the 70B (AWQ) model for free at [dreamgen.com](https://dreamgen.com) — note that an account is required.
The version used for the website is the official AWQ 4bit quant [dreamgen/opus-v0-70b-awq](https://huggingface.co/dreamgen/opus-v0-70b-awq).

## License

- For personal and academic use: Same license as the base model, in this case https://ai.meta.com/resources/models-and-libraries/llama-downloads/.
- For commercial use: Please reach out to hello@dreamgen.com.
