---
language:
  - en
pipeline_tag: text-generation
---

# DreamGen Opus V0 7B

**DreamGen Opus** is a family of **uncensored** models fine-tuned for **(steerable) story writing** and the model also works great for **chat / RP**.
The DreamGen Opus V0 7B model is derived from [mistralai/Mistral-7B-v0.1](https://huggingface.co/mistralai/Mistral-7B-v0.1).

**WARNING** Quantized versions manifest significant performance degradation compared to the original:

- AWQ: [dreamgen/opus-v0-7b-awq](https://huggingface.co/dreamgen/opus-v0-7b-awq)
- GGUF: [TheBloke/opus-v0-7b-gguf](https://huggingface.co/TheBloke/opus-v0-7B-GGUF)
- ELX2: [bartowski/opus-v0-7b-exl2](https://huggingface.co/bartowski/opus-v0-7b-exl2)

Other sizes:

- 70B: [dreamgen/opus-v0-70b](https://huggingface.co/dreamgen/opus-v0-70b)

You can **try the Opus V0 70B** (AWQ) model for free on [dreamgen.com](https://dreamgen.com).

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

The model is should be compatible with any software that supports [mistralai/Mistral-7B-v0.1](https://huggingface.co/mistralai/Mistral-7B-v0.1).

### Running on DreamGen.com (free)

You can try the model for free at [dreamgen.com](https://dreamgen.com) — note that an account is required.

### Running with vLLM

1. Install [vLLM](https://github.com/vllm-project/vllm) (version at least 0.2.1.post1)
2. Run `python -u -m vllm.entrypoints.openai.api_server --host 0.0.0.0 --model dreamgen/opus-v0-7b`
3. Use any library that's compatible with the OpenAI API to talk to the model

### Running with oobabooga/text-generation-webui

1. Install [oobabooga/text-generation-webui](https://github.com/oobabooga/text-generation-webui)
2. Go to the **Model** tab
3. Under **Download custom model or LoRA**, enter `dreamgen/opus-v0-7b`
4. Go to the **Text Generation** tab
5. Enter your prompt

## License

- For personal and academic use: Same license as the base model, in this case Apache 2.0.
- For commercial use: Please reach out to hello@dreamgen.com.
