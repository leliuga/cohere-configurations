---
license: apache-2.0
datasets:
- Fredithefish/openassistant-guanaco-unfiltered
language:
- en
library_name: transformers
pipeline_tag: conversational
inference: false
---

<img src="https://huggingface.co/Fredithefish/Guanaco-3B-Uncensored/resolve/main/Guanaco-Uncensored.jpg" alt="Alt Text" width="295"/>

# ✨ Guanaco - 7B - Uncensored ✨


Guanaco-7B-Uncensored has been fine-tuned for 4 epochs on the [Unfiltered Guanaco Dataset.](https://huggingface.co/datasets/Fredithefish/openassistant-guanaco-unfiltered) using [Llama-2-7b](https://hf.co/meta-llama/Llama-2-7b-hf) as the base model.
<br>The model does not perform well with languages other than English.
<br>Please note: This model is designed to provide responses without content filtering or censorship. It generates answers without denials.

## Special thanks
I would like to thank AutoMeta for providing me with the computing power necessary to train this model.


### Prompt Template
```
### Human: {prompt} ### Assistant:
```

### Dataset
The model has been fine-tuned on the V2 of the Guanaco unfiltered dataset.