---
license: apache-2.0
tags:
- merge
---
WARNING! This is a "Frankenmerge" - this model is untested!

# NeuralOrca 7B V1

[GGUF Models](https://huggingface.co/TheBloke/NeuralOrca-7B-v1-GGUF)

**By [mrfakename](https://twitter.com/realmrfakename)**

*Please note that this is an experimental model. We cannot guarantee model quality.*

This is the first (alpha) release of NeuralOrca. NeuralOrca is a merge of the following models:

* [mlabonne/NeuralHermes-2.5-Mistral-7B](https://huggingface.co/mlabonne/NeuralHermes-2.5-Mistral-7B) (This model is actually [OpenHermes 2.5](https://huggingface.co/teknium/OpenHermes-2.5-Mistral-7B) finetuned on Intel's [Neural Chat dataset](https://huggingface.co/datasets/Intel/neural-chat-dataset-v2) and uses the ChatML prompt format, weight: 1.0)
* [Open-Orca/Mistral-7B-OpenOrca](https://huggingface.co/Open-Orca/Mistral-7B-OpenOrca) (This model uses the ChatML prompt format, weight: 0.7)

## Prompt Format

We use the ChatML prompt format.

Example:
```
<|im_start|>system
You are NeuralOrca, a helpful AI assistant.
<|im_end|>
<|im_start|>user
How are you?<|im_end|>
<|im_start|>assistant
```

## Evaluations

Coming soon

## Context Length

The context length for this model is 8192 tokens (8K).

## License

You are responsible for your use of NeuralOrca.

This software is licensed under the Apache 2.0 license. If you want to use it for commercial use, it's probably fine but please contact me first.