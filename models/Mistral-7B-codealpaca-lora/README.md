---
license: apache-2.0
tags:
- code
- mistral
---

# Mistral-7B-codealpaca

I am thrilled to introduce my Mistral-7B-codealpaca model. This variant is optimized and demonstrates potential in assisting developers as a coding companion. I welcome contributions from testers and enthusiasts to help evaluate its performance.

## Training Details

I trained the model using 3xRTX 3090 for 118 hours.
[![Built with Axolotl](https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png)](https://github.com/OpenAccess-AI-Collective/axolotl)

## Quantised Model Links:

1. https://huggingface.co/TheBloke/Mistral-7B-codealpaca-lora-GPTQ
2. https://huggingface.co/TheBloke/Mistral-7B-codealpaca-lora-GGUF
3. https://huggingface.co/TheBloke/Mistral-7B-codealpaca-lora-AWQ

## Download by qBittorrent:

#### Torrent file:  https://github.com/Nondzu/LlamaTor/blob/torrents/torrents/Nondzu_Mistral-7B-codealpaca-lora.torrent

## Dataset:

- Dataset Name: theblackcat102/evol-codealpaca-v1
- Dataset Link: [theblackcat102/evol-codealpaca-v1](https://huggingface.co/datasets/theblackcat102/evol-codealpaca-v1)

## Prompt template: Alpaca

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:

```


## Performance (evalplus)
Human eval plus: https://github.com/evalplus/evalplus

![image/png](https://cdn-uploads.huggingface.co/production/uploads/63729f35acef705233c87909/azE6LU0qQ9E9u60t5VrMk.png)

Well, the results are better than I expected:
- Base: `{'pass@1':  0.47560975609756095}`
- Base + Extra: `{'pass@1': 0.4329268292682927}`

For reference, I've provided the performance of the original Mistral model alongside my Mistral-7B-code-16k-qlora model. 

**   [Nondzu/Mistral-7B-code-16k-qlora](https://huggingface.co/Nondzu/Mistral-7B-code-16k-qlora)**:

- Base: `{'pass@1': 0.3353658536585366}`
- Base + Extra: `{'pass@1': 0.2804878048780488}`

** [mistralai/Mistral-7B-Instruct-v0.1](https://huggingface.co/mistralai/Mistral-7B-Instruct-v0.1)**:

- Base: `{'pass@1': 0.2926829268292683}`
- Base + Extra: `{'pass@1': 0.24390243902439024}`

## Model Configuration:

Here are the configurations for my Mistral-7B-codealpaca-lora:

```yaml
base_model: mistralai/Mistral-7B-Instruct-v0.1
base_model_config: mistralai/Mistral-7B-Instruct-v0.1
model_type: MistralForCausalLM
tokenizer_type: LlamaTokenizer
is_mistral_derived_model: true
load_in_8bit: true
load_in_4bit: false
strict: false
datasets:
  - path: theblackcat102/evol-codealpaca-v1
    type: oasst
dataset_prepared_path:
val_set_size: 0.01
output_dir: ./nondzu/Mistral-7B-codealpaca-test14
adapter: lora
sequence_len: 4096
sample_packing: true
pad_to_sequence_len: true
lora_r: 32
lora_alpha: 16
lora_dropout: 0.05
lora_target_modules:
lora_target_linear: true
```

![image/png](https://cdn-uploads.huggingface.co/production/uploads/63729f35acef705233c87909/5nPgL3ajROKf7dttf4BO0.png)

## Additional Projects:

For other related projects, you can check out:

- [LlamaTor on GitHub](https://github.com/Nondzu/LlamaTor)