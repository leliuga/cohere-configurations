---
license: llama2
---

# CodeBooga-34B-v0.1

This is a merge between the following two models:

1) [Phind-CodeLlama-34B-v2](https://huggingface.co/Phind/Phind-CodeLlama-34B-v2)
2) [WizardCoder-Python-34B-V1.0](https://huggingface.co/WizardLM/WizardCoder-Python-34B-V1.0)

It was created with the [BlockMerge Gradient script](https://github.com/Gryphe/BlockMerge_Gradient), the same one that was used to create [MythoMax-L2-13b](https://huggingface.co/Gryphe/MythoMax-L2-13b), and with the same settings. The following YAML was used:

```yaml
model_path1: "Phind_Phind-CodeLlama-34B-v2_safetensors"
model_path2: "WizardLM_WizardCoder-Python-34B-V1.0_safetensors"
output_model_path: "CodeBooga-34B-v0.1"
operations:
  - operation: lm_head # Single tensor
    filter: "lm_head"
    gradient_values: [0.75]
  - operation: embed_tokens # Single tensor
    filter: "embed_tokens"
    gradient_values: [0.75]
  - operation: self_attn
    filter: "self_attn"
    gradient_values: [0.75, 0.25]
  - operation: mlp
    filter: "mlp"
    gradient_values: [0.25, 0.75]
  - operation: layernorm
    filter: "layernorm"
    gradient_values: [0.5, 0.5]
  - operation: modelnorm # Single tensor
    filter: "model.norm"
    gradient_values: [0.75]
```

## Prompt format

Both base models use the Alpaca format, so it should be used for this one as well.

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
Your instruction

### Response:
Bot reply

### Instruction:
Another instruction

### Response:
Bot reply
```

## Evaluation

I made a quick experiment where I asked a set of 3 Python and 3 Javascript questions (real world, difficult questions with nuance) to the following models:

1) This one
2) A second variant generated with `model_path1` and `model_path2` swapped in the YAML above, which I called CodeBooga-Reversed-34B-v0.1
3) WizardCoder-Python-34B-V1.0
4) Phind-CodeLlama-34B-v2

Specifically, I used 4.250b EXL2 quantizations of each. I then sorted the responses for each question by quality, and attributed the following scores:

* 4th place: 0
* 3rd place: 1
* 2nd place: 2
* 1st place: 4

The resulting cumulative scores were:

* CodeBooga-34B-v0.1: 22
* WizardCoder-Python-34B-V1.0: 12
* Phind-CodeLlama-34B-v2: 7
* CodeBooga-Reversed-34B-v0.1: 1

CodeBooga-34B-v0.1 performed very well, while its variant performed poorly, so I uploaded the former but not the latter.

## Recommended settings

I recommend the [Divine Intellect](https://github.com/oobabooga/text-generation-webui/blob/ae8cd449ae3e0236ecb3775892bb1eea23f9ed68/presets/Divine%20Intellect.yaml) preset for instruction-following models like this, as per the [Preset Arena experiment results](https://github.com/oobabooga/oobabooga.github.io/blob/main/arena/results.md):

```yaml
temperature: 1.31
top_p: 0.14
repetition_penalty: 1.17
top_k: 49
```

## Quantized versions

### EXL2

A 4.250b EXL2 version of the model can be found here: 

https://huggingface.co/oobabooga/CodeBooga-34B-v0.1-EXL2-4.250b

### GGUF

TheBloke has kindly provided GGUF quantizations for llama.cpp:

https://huggingface.co/TheBloke/CodeBooga-34B-v0.1-GGUF

<a href="https://ko-fi.com/oobabooga"><img src="https://i.imgur.com/UJlEAYw.png"></a>