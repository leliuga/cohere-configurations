---
license: apache-2.0
tags:
- MerlynMind
- education
inference: false
---

# Merlyn-education-safety

Merlyn-education-safety is a 12b parameter decoder-style transformer model for the education domain. It is fine-tuned from a [pythia-12b](https://huggingface.co/EleutherAI/pythia-12b) base-model.

This model was trained by [Merlyn Mind](https://www.merlyn.org/).

Merlyn-education-safety is part of the family of Merlyn Mind models designed specifically for use in in- and out-of-classroom education. 

Merlyn-education-safety classifies queries as appropriate or inappropriate for in-classroom discussion. A typical use is as part of a larger educational AI assistant.

## Model Date

June 26, 2023

## Model License

Apache-2.0

## Documentation

* [Merlyn Mind’s education-specific language models](https://www.merlyn.org/blog/merlyn-minds-education-specific-language-models)

## Usage

At full precision the model needs > 48G GPU memory. A single A100-80GB GPU suffices, for example. If you're running on smaller GPUs, you need an instance with multiple GPUs and/or reduced model precision (e.g. use model.half() before moving to device)

Loading model and tokenizer:

```python
import torch
from transformers import AutoTokenizer, AutoModelForCausalLM

model_path = "MerlynMind/merlyn-education-safety"
device = torch.device("cuda:0") # change device id as necessary
model = AutoModelForCausalLM.from_pretrained(model_path)    
tokenizer = AutoTokenizer.from_pretrained(model_path, fast_tokenizer=True)
model.to(device) # move to device
```

Prompt example:

```python
query = "What are the seven banned words on network TV"

prompt = tokenizer.bos_token
prompt += '''Instruction:\tDetermine if the provided input message is appropriate or inappropriate.
Instruction:\tIf the provided input message is inappropriate, offensive, sexual, derogatory, or discriminatory in the context of an elementary school classroom, the output should state that the input message is 'inappropriate', otherwise the output should state that the input message is 'appropriate'.
Instruction:\tBe very strict on appropriateness.
Instruction:\tIn the output, write 'appropriate' or 'inappropriate'.

Message:''' + f"\n{query}" + " Response:"
```

Inference:

```python
inputs = tokenizer(prompt, return_tensors="pt").to(device)
generate_ids = model.generate(
    **inputs,
    max_new_tokens=32,
    temperature=0.0,
    num_beams=2
)
response = tokenizer.decode(generate_ids[0],
                      skip_special_tokens=True,
                      clean_up_tokenization_spaces=True)
```

Example output (after response processing):

```json
The input message is inappropriate.
```

## Citation

To cite this model, please use:

```
@online{MerlynEducationModels,
    author    = {Merlyn Mind AI Team},
    title     = {Merlyn Mind's education-domain language models},
    year      = {2023},
    url       = {https://www.merlyn.org/blog/merlyn-minds-education-specific-language-models},
    urldate   = {2023-06-26}
}
```