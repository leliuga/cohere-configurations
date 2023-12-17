LLM Inference configurations👋
===============================

## Description

This repository contains the inference configurations for the Co:Here

### Supported Transformer Architectures by Co:Here

- [X] LLaMA 🦙
- [x] LLaMA 2 🦙🦙
- [X] Falcon
- [X] Alpaca
- [X] GPT4All
- [X] Chinese LLaMA / Alpaca and Chinese LLaMA-2 / Alpaca-2
- [X] Vigogne (French)
- [X] Vicuna
- [X] Koala
- [X] OpenBuddy 🐶 (Multilingual)
- [X] Pygmalion/Metharme
- [X] WizardLM
- [X] Baichuan 1 & 2 + derivations
- [X] Aquila 1 & 2
- [X] Starcoder
- [X] Mistral AI v0.1
- [X] Refact
- [X] Persimmon 8B
- [X] MPT
- [X] Bloom
- [X] StableLM-3b-4e1t
- [x] Yi
- [x] Deepseek
- [x] Qwen
- [x] Mixtral MoE

Multimodal

- [x] Llava 1.5
- [x] Bakllava
- [x] Obsidian
- [x] ShareGPT4V

### Supported backends
- [X] llama-backend
- [ ] bert-backend
- [ ] bart-backend
- [ ] t5-backend
- [ ] rwkv-backend

### Supported platforms
- [X] linux/amd64
- [ ] darwin/amd64

## Requirements

1. [Docker](https://docs.docker.com/get-docker/)
2. [NVIDIA Container Toolkit](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/install-guide.html)

## Usage

``` bash
docker run -it --rm --gpus all -v ./models/:/app/models -p 3000:3000 github.com/leliuga/cohere cohere run <name>/<dtype>
```

## Supported Models (810)

| Name  | DType   | Context | Embedding | Read more |
| ----- | ------- | ------- | --------- | --------- |
| 13B-BlueMethod |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/13B-BlueMethod/README.md) |
| 13B-Chimera |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/13B-Chimera/README.md) |
| 13B-HyperMantis |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/13B-HyperMantis/README.md) |
| 13B-Legerdemain-L2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/13B-Legerdemain-L2/README.md) |
| 13B-Ouroboros |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/13B-Ouroboros/README.md) |
| 13B-Thorns-L2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/13B-Thorns-L2/README.md) |
| 30B-Epsilon |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/30B-Epsilon/README.md) |
| 30B-Lazarus |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/30B-Lazarus/README.md) |
| 8x7B-MoE-test-NOT-MIXTRAL |Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/8x7B-MoE-test-NOT-MIXTRAL/README.md) |
| Aetheria-L2-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Aetheria-L2-70B/README.md) |
| agentlm-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/agentlm-13B/README.md) |
| agentlm-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/agentlm-70B/README.md) |
| agentlm-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/agentlm-7B/README.md) |
| agiin-13.6B-v0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/agiin-13.6B-v0.1/README.md) |
| Airoboros-33B-2.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 6656 | [README.md](models/Airoboros-33B-2.1/README.md) |
| Airoboros-c34B-2.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/Airoboros-c34B-2.1/README.md) |
| Airoboros-c34B-2.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/Airoboros-c34B-2.2/README.md) |
| airoboros-c34b-2.2.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/airoboros-c34b-2.2.1/README.md) |
| Airoboros-c34B-3.1.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/Airoboros-c34B-3.1.2/README.md) |
| Airoboros-L2-13B-2.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Airoboros-L2-13B-2.1/README.md) |
| Airoboros-L2-13B-2_1-YaRN-64K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 65536 | 5120 | [README.md](models/Airoboros-L2-13B-2_1-YaRN-64K/README.md) |
| Airoboros-L2-13B-2.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Airoboros-L2-13B-2.2/README.md) |
| airoboros-l2-13B-2.2.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/airoboros-l2-13B-2.2.1/README.md) |
| airoboros-l2-13B-3.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/airoboros-l2-13B-3.0/README.md) |
| Airoboros-L2-13B-3.1.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Airoboros-L2-13B-3.1.1/README.md) |
| airoboros-l2-13B-gpt4-1.4.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/airoboros-l2-13B-gpt4-1.4.1/README.md) |
| airoboros-l2-13b-gpt4-2.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/airoboros-l2-13b-gpt4-2.0/README.md) |
| airoboros-l2-13b-gpt4-m2.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/airoboros-l2-13b-gpt4-m2.0/README.md) |
| Airoboros-L2-70B-2.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Airoboros-L2-70B-2.1/README.md) |
| Airoboros-L2-70B-2.1-Creative |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Airoboros-L2-70B-2.1-Creative/README.md) |
| Airoboros-L2-70b-2.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Airoboros-L2-70b-2.2/README.md) |
| Airoboros-L2-70b-2.2.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Airoboros-L2-70b-2.2.1/README.md) |
| Airoboros-L2-70B-3.1.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Airoboros-L2-70B-3.1.2/README.md) |
| airoboros-l2-70B-gpt4-1.4.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/airoboros-l2-70B-gpt4-1.4.1/README.md) |
| airoboros-l2-70B-GPT4-2.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/airoboros-l2-70B-GPT4-2.0/README.md) |
| Airoboros-L2-70B-GPT4-m2.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Airoboros-L2-70B-GPT4-m2.0/README.md) |
| Airoboros-L2-7B-2.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Airoboros-L2-7B-2.1/README.md) |
| Airoboros-L2-7B-2.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Airoboros-L2-7B-2.2/README.md) |
| airoboros-l2-7B-2.2.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/airoboros-l2-7B-2.2.1/README.md) |
| airoboros-l2-7B-3.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/airoboros-l2-7B-3.0/README.md) |
| airoboros-l2-7b-gpt4-1.4.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/airoboros-l2-7b-gpt4-1.4.1/README.md) |
| airoboros-l2-7B-gpt4-2.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/airoboros-l2-7B-gpt4-2.0/README.md) |
| airoboros-l2-7B-gpt4-m2.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/airoboros-l2-7B-gpt4-m2.0/README.md) |
| airoboros-m-7B-3.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/airoboros-m-7B-3.0/README.md) |
| Airoboros-M-7B-3.1.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Airoboros-M-7B-3.1.2/README.md) |
| airoboros-m-7B-3.1.2-dare-0.85 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/airoboros-m-7B-3.1.2-dare-0.85/README.md) |
| airoboros-mistral2.2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/airoboros-mistral2.2-7B/README.md) |
| airochronos-33B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/airochronos-33B/README.md) |
| Airochronos-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Airochronos-L2-13B/README.md) |
| Airolima-Chronos-Grad-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Airolima-Chronos-Grad-L2-13B/README.md) |
| Akins-3B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 2560 | [README.md](models/Akins-3B/README.md) |
| alfred-40B-1023 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 0 | 8192 | [README.md](models/alfred-40B-1023/README.md) |
| ALMA-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/ALMA-13B/README.md) |
| ALMA-13B-Pretrain |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/ALMA-13B-Pretrain/README.md) |
| ALMA-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/ALMA-7B/README.md) |
| ALMA-7B-Pretrain |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/ALMA-7B-Pretrain/README.md) |
| AlpacaCielo2-7B-8K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/AlpacaCielo2-7B-8K/README.md) |
| Amber |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/Amber/README.md) |
| AmberChat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/AmberChat/README.md) |
| Amethyst-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Amethyst-13B/README.md) |
| Amethyst-13B-Mistral |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Amethyst-13B-Mistral/README.md) |
| Ana-v1-m7 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Ana-v1-m7/README.md) |
| ANIMA-Phi-Neptune-Mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/ANIMA-Phi-Neptune-Mistral-7B/README.md) |
| AppleSauce-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/AppleSauce-L2-13B/README.md) |
| AquilaChat2-34B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 6144 | [README.md](models/AquilaChat2-34B/README.md) |
| AquilaChat2-34B-16K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 6144 | [README.md](models/AquilaChat2-34B-16K/README.md) |
| ARIA-70B-V2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/ARIA-70B-V2/README.md) |
| Arithmo-Mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Arithmo-Mistral-7B/README.md) |
| Asclepius-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/Asclepius-13B/README.md) |
| AshhLimaRP-Mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/AshhLimaRP-Mistral-7B/README.md) |
| Astrid-Mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Astrid-Mistral-7B/README.md) |
| Athena-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Athena-v1/README.md) |
| Athena-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Athena-v2/README.md) |
| Athena-v3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Athena-v3/README.md) |
| Athena-v4 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Athena-v4/README.md) |
| Athnete-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Athnete-13B/README.md) |
| Augmental-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Augmental-13B/README.md) |
| Augmental-13B-v1.50_A |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Augmental-13B-v1.50_A/README.md) |
| Augmental-13B-v1.50_B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Augmental-13B-v1.50_B/README.md) |
| Augmental-ReMM-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Augmental-ReMM-13B/README.md) |
| Augmental-Unholy-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Augmental-Unholy-13B/README.md) |
| Autolycus-Mistral_7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Autolycus-Mistral_7B/README.md) |
| bagel-7B-v0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/bagel-7B-v0.1/README.md) |
| bagel-dpo-7B-v0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/bagel-dpo-7B-v0.1/README.md) |
| Barcenas-Mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Barcenas-Mistral-7B/README.md) |
| based-13b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/based-13b/README.md) |
| based-30B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/based-30B/README.md) |
| based-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/based-7B/README.md) |
| basilisk-7B-v0.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/basilisk-7B-v0.2/README.md) |
| BerrySauce-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/BerrySauce-L2-13B/README.md) |
| BigPlap-8x20B |Q2_K | 32768 | 5120 | [README.md](models/BigPlap-8x20B/README.md) |
| bling-stable-lm-3b-4e1t-v0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 2560 | [README.md](models/bling-stable-lm-3b-4e1t-v0/README.md) |
| blossom-v3_1-yi-34b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 7168 | [README.md](models/blossom-v3_1-yi-34b/README.md) |
| blossom-v3-baichuan2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/blossom-v3-baichuan2-7B/README.md) |
| blossom-v3-mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/blossom-v3-mistral-7B/README.md) |
| BruinsV2-OpHermesNeu-11B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/BruinsV2-OpHermesNeu-11B/README.md) |
| CalliopeDS-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/CalliopeDS-L2-13B/README.md) |
| calm2-7B-chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/calm2-7B-chat/README.md) |
| CAMEL-13B-Combined-Data |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/CAMEL-13B-Combined-Data/README.md) |
| CAMEL-13B-Role-Playing-Data |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/CAMEL-13B-Role-Playing-Data/README.md) |
| CAMEL-33B-Combined-Data |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/CAMEL-33B-Combined-Data/README.md) |
| Camel-Platypus2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Camel-Platypus2-13B/README.md) |
| Camel-Platypus2-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Camel-Platypus2-70B/README.md) |
| CaPlatTessDolXaBoros-Yi-34B-200K-DARE-Ties-HighDensity |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/CaPlatTessDolXaBoros-Yi-34B-200K-DARE-Ties-HighDensity/README.md) |
| Capybara-Tess-Yi-34B-200K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/Capybara-Tess-Yi-34B-200K/README.md) |
| Capybara-Tess-Yi-34B-200K-DARE-Ties |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/Capybara-Tess-Yi-34B-200K-DARE-Ties/README.md) |
| CapyTessBorosYi-34B-200K-DARE-Ties |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/CapyTessBorosYi-34B-200K-DARE-Ties/README.md) |
| Carl-Llama-2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Carl-Llama-2-13B/README.md) |
| Cat-13B-0.5 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Cat-13B-0.5/README.md) |
| cat-v1.0-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/cat-v1.0-13B/README.md) |
| CausalLM-14B |Q4_0 Q4_1 Q5_0 Q5_1 Q8_0 | 8192 | 5120 | [README.md](models/CausalLM-14B/README.md) |
| CausalLM-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/CausalLM-7B/README.md) |
| ChatAYT-Lora-Assamble-Marcoroni |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/ChatAYT-Lora-Assamble-Marcoroni/README.md) |
| Chinese-Alpaca-2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Chinese-Alpaca-2-13B/README.md) |
| Chinese-Alpaca-2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Chinese-Alpaca-2-7B/README.md) |
| Chinese-Llama-2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Chinese-Llama-2-13B/README.md) |
| Chinese-Llama-2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Chinese-Llama-2-7B/README.md) |
| Chronoboros-33B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/Chronoboros-33B/README.md) |
| Chronoboros-Grad-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Chronoboros-Grad-L2-13B/README.md) |
| Chronohermes-Grad-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Chronohermes-Grad-L2-13B/README.md) |
| Chronolima-Airo-Grad-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Chronolima-Airo-Grad-L2-13B/README.md) |
| Chronomaid-Storytelling-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Chronomaid-Storytelling-13B/README.md) |
| Chronorctypus-Limarobormes-13b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Chronorctypus-Limarobormes-13b/README.md) |
| chronos007-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/chronos007-70B/README.md) |
| chronos-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/chronos-13B/README.md) |
| chronos-13b-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/chronos-13b-v2/README.md) |
| chronos-33b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/chronos-33b/README.md) |
| Chronos-70B-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Chronos-70B-v2/README.md) |
| Chronos-Beluga-v2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Chronos-Beluga-v2-13B/README.md) |
| chronos-hermes-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/chronos-hermes-13B/README.md) |
| Chronos-Hermes-13b-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Chronos-Hermes-13b-v2/README.md) |
| chronos-wizardlm-uc-scot-st-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/chronos-wizardlm-uc-scot-st-13B/README.md) |
| Chupacabra-7B-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Chupacabra-7B-v2/README.md) |
| Chupacabra-7B-v3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 4096 | [README.md](models/Chupacabra-7B-v3/README.md) |
| Chupacabra-8x7B-MoE |Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Chupacabra-8x7B-MoE/README.md) |
| cinematika-7B-v0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/cinematika-7B-v0.1/README.md) |
| Claire-7B-0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 0 | 4544 | [README.md](models/Claire-7B-0.1/README.md) |
| claude2-alpaca-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/claude2-alpaca-13B/README.md) |
| claude2-alpaca-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/claude2-alpaca-7B/README.md) |
| Clover3-17B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Clover3-17B/README.md) |
| cockatrice-7B-v0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/cockatrice-7B-v0.1/README.md) |
| Code-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Code-13B/README.md) |
| Code-33B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/Code-33B/README.md) |
| CodeBooga-34B-v0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/CodeBooga-34B-v0.1/README.md) |
| CodeFuse-CodeLlama-34B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/CodeFuse-CodeLlama-34B/README.md) |
| CodeLlama-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 5120 | [README.md](models/CodeLlama-13B/README.md) |
| CodeLlama-13B-Instruct |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 5120 | [README.md](models/CodeLlama-13B-Instruct/README.md) |
| CodeLlama-13B-oasst-sft-v10 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 5120 | [README.md](models/CodeLlama-13B-oasst-sft-v10/README.md) |
| CodeLlama-13B-Python |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 5120 | [README.md](models/CodeLlama-13B-Python/README.md) |
| CodeLlama-34B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/CodeLlama-34B/README.md) |
| CodeLlama-34B-Instruct |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/CodeLlama-34B-Instruct/README.md) |
| CodeLlama-34B-Python |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/CodeLlama-34B-Python/README.md) |
| CodeLlama-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 4096 | [README.md](models/CodeLlama-7B/README.md) |
| CodeLlama-7B-Instruct |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 4096 | [README.md](models/CodeLlama-7B-Instruct/README.md) |
| CodeLlama-7B-Python |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 4096 | [README.md](models/CodeLlama-7B-Python/README.md) |
| CodeUp-Alpha-13B-HF |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/CodeUp-Alpha-13B-HF/README.md) |
| CodeUp-Llama-2-13B-Chat-HF |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/CodeUp-Llama-2-13B-Chat-HF/README.md) |
| CollectiveCognition-v1.1-Mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/CollectiveCognition-v1.1-Mistral-7B/README.md) |
| CollectiveCognition-v1.1-Mistral-7B-dare-0.85 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/CollectiveCognition-v1.1-Mistral-7B-dare-0.85/README.md) |
| CollectiveCognition-v1-Mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/CollectiveCognition-v1-Mistral-7B/README.md) |
| Dans-AdventurousWinds-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Dans-AdventurousWinds-7B/README.md) |
| Dans-AdventurousWinds-Mk2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Dans-AdventurousWinds-Mk2-7B/README.md) |
| Dans-TotSirocco-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Dans-TotSirocco-7B/README.md) |
| DaringFortitude |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/DaringFortitude/README.md) |
| Dawn-v2-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Dawn-v2-70B/README.md) |
| Deacon-34B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 7168 | [README.md](models/Deacon-34B/README.md) |
| deepseek-coder-1.3b-base |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 2048 | [README.md](models/deepseek-coder-1.3b-base/README.md) |
| deepseek-coder-1.3b-instruct |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 2048 | [README.md](models/deepseek-coder-1.3b-instruct/README.md) |
| deepseek-coder-33B-base |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 7168 | [README.md](models/deepseek-coder-33B-base/README.md) |
| deepseek-coder-33B-instruct |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 7168 | [README.md](models/deepseek-coder-33B-instruct/README.md) |
| deepseek-coder-5.7bmqa-base |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 4096 | [README.md](models/deepseek-coder-5.7bmqa-base/README.md) |
| deepseek-coder-6.7B-base |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 4096 | [README.md](models/deepseek-coder-6.7B-base/README.md) |
| deepseek-coder-6.7B-instruct |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 4096 | [README.md](models/deepseek-coder-6.7B-instruct/README.md) |
| deepseek-llm-67b-base |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/deepseek-llm-67b-base/README.md) |
| deepseek-llm-67b-chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/deepseek-llm-67b-chat/README.md) |
| deepseek-llm-7B-base |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/deepseek-llm-7B-base/README.md) |
| deepseek-llm-7B-chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/deepseek-llm-7B-chat/README.md) |
| deepsex-34b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 7168 | [README.md](models/deepsex-34b/README.md) |
| digital-socrates-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/digital-socrates-13B/README.md) |
| digital-socrates-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/digital-socrates-7B/README.md) |
| DiscoLM-120b |Q2_K | 4096 | 8192 | [README.md](models/DiscoLM-120b/README.md) |
| DiscoLM-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 8192 | 8192 | [README.md](models/DiscoLM-70B/README.md) |
| dolphin-2.0-mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/dolphin-2.0-mistral-7B/README.md) |
| Dolphin-2.1-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Dolphin-2.1-70B/README.md) |
| dolphin-2.1-mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/dolphin-2.1-mistral-7B/README.md) |
| Dolphin2.1-OpenOrca-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Dolphin2.1-OpenOrca-7B/README.md) |
| dolphin-2.2.1-AshhLimaRP-Mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/dolphin-2.2.1-AshhLimaRP-Mistral-7B/README.md) |
| dolphin-2.2.1-mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/dolphin-2.2.1-mistral-7B/README.md) |
| dolphin-2.2-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/dolphin-2.2-70B/README.md) |
| dolphin-2_2-yi-34b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 7168 | [README.md](models/dolphin-2_2-yi-34b/README.md) |
| dolphin-2.2-yi-34b-200k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/dolphin-2.2-yi-34b-200k/README.md) |
| dolphin-2.5-mixtral-8x7b |Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/dolphin-2.5-mixtral-8x7b/README.md) |
| Dolphin-Llama-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/Dolphin-Llama-13B/README.md) |
| Dolphin-Llama2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Dolphin-Llama2-7B/README.md) |
| dopeystableplats-3b-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 2560 | [README.md](models/dopeystableplats-3b-v1/README.md) |
| DPOpenHermes-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/DPOpenHermes-7B/README.md) |
| DPOpenHermes-7B-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/DPOpenHermes-7B-v2/README.md) |
| dragon-yi-6B-v0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/dragon-yi-6B-v0/README.md) |
| DynamicFactor |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/DynamicFactor/README.md) |
| Echidna-13B-v0.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Echidna-13B-v0.2/README.md) |
| Echidna-13B-v0.3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Echidna-13B-v0.3/README.md) |
| echidna-tiefigther-25 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/echidna-tiefigther-25/README.md) |
| Emerhyst-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Emerhyst-13B/README.md) |
| Emerhyst-20B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Emerhyst-20B/README.md) |
| em_german_13b_v01 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/em_german_13b_v01/README.md) |
| em_german_70b_v01 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/em_german_70b_v01/README.md) |
| em_german_7b_v01 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/em_german_7b_v01/README.md) |
| em_german_leo_mistral |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/em_german_leo_mistral/README.md) |
| em_german_mistral_v01 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/em_german_mistral_v01/README.md) |
| Euryale-1.3-L2-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Euryale-1.3-L2-70B/README.md) |
| Euryale-1.4-L2-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Euryale-1.4-L2-70B/README.md) |
| Euryale-Inverted-L2-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Euryale-Inverted-L2-70B/README.md) |
| Euryale-L2-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Euryale-L2-70B/README.md) |
| EverythingLM-13B-16K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 5120 | [README.md](models/EverythingLM-13B-16K/README.md) |
| EverythingLM-13b-V2-16K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 5120 | [README.md](models/EverythingLM-13b-V2-16K/README.md) |
| EverythingLM-13B-V3-16K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 5120 | [README.md](models/EverythingLM-13B-V3-16K/README.md) |
| evolvedSeeker_1_3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 2048 | [README.md](models/evolvedSeeker_1_3/README.md) |
| Falkor-8x7B-MoE |Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Falkor-8x7B-MoE/README.md) |
| FashionGPT-70B-V1.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/FashionGPT-70B-V1.1/README.md) |
| FashionGPT-70B-v1.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/FashionGPT-70B-v1.2/README.md) |
| Ferret_7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Ferret_7B/README.md) |
| fin-llama-33B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/fin-llama-33B/README.md) |
| firefly-llama2-13B-chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/firefly-llama2-13B-chat/README.md) |
| firefly-llama2-7B-chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/firefly-llama2-7B-chat/README.md) |
| Free_Sydney_V2_13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Free_Sydney_V2_13B/README.md) |
| Free_Sydney_V2_Mistral_7b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Free_Sydney_V2_Mistral_7b/README.md) |
| Generate_Question_Mistral_7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Generate_Question_Mistral_7B/README.md) |
| genz-13B-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/genz-13B-v2/README.md) |
| Genz-70b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Genz-70b/README.md) |
| GOAT-70B-Storytelling |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/GOAT-70B-Storytelling/README.md) |
| go-bruins |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/go-bruins/README.md) |
| go-bruins-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/go-bruins-v2/README.md) |
| go-bruins-v2.1.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/go-bruins-v2.1.1/README.md) |
| GodziLLa2-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/GodziLLa2-70B/README.md) |
| goliath-120b |Q2_K | 4096 | 8192 | [README.md](models/goliath-120b/README.md) |
| gorilla-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/gorilla-7B/README.md) |
| gorilla-openfunctions-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/gorilla-openfunctions-v1/README.md) |
| GPlatty-30B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/GPlatty-30B/README.md) |
| GreenNodeLM-7B-v4leo |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/GreenNodeLM-7B-v4leo/README.md) |
| Guanaco-13B-Uncensored |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Guanaco-13B-Uncensored/README.md) |
| Guanaco-7B-Uncensored |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Guanaco-7B-Uncensored/README.md) |
| HermesLimaRP-L2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/HermesLimaRP-L2-7B/README.md) |
| Hermes-Trismegistus-Mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Hermes-Trismegistus-Mistral-7B/README.md) |
| Hexoteric-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Hexoteric-7B/README.md) |
| hippogriff-30b-chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/hippogriff-30b-chat/README.md) |
| HornyEchidna-13B-v0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/HornyEchidna-13B-v0.1/README.md) |
| Huginn-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Huginn-13B/README.md) |
| Huginn-13B-v4 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Huginn-13B-v4/README.md) |
| Huginn-13B-v4.5 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Huginn-13B-v4.5/README.md) |
| Huginn-22B-Prototype |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 6656 | [README.md](models/Huginn-22B-Prototype/README.md) |
| huginnv1.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/huginnv1.2/README.md) |
| Huginn-v3-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Huginn-v3-13B/README.md) |
| Iambe-20B-DARE |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Iambe-20B-DARE/README.md) |
| Iambe-RP-cDPO-20B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Iambe-RP-cDPO-20B/README.md) |
| Iambe-Storyteller-20B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Iambe-Storyteller-20B/README.md) |
| Inairtra-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Inairtra-7B/README.md) |
| Inkbot-13B-4k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Inkbot-13B-4k/README.md) |
| Inkbot-13B-8k-0.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Inkbot-13B-8k-0.2/README.md) |
| jackalope-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/jackalope-7B/README.md) |
| JanniesBasedLigma-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/JanniesBasedLigma-L2-13B/README.md) |
| japanese-stablelm-instruct-gamma-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/japanese-stablelm-instruct-gamma-7B/README.md) |
| juanako-7B-UNA |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/juanako-7B-UNA/README.md) |
| juanako-7B-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/juanako-7B-v1/README.md) |
| KAI-7B-beta |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/KAI-7B-beta/README.md) |
| KAI-7B-Instruct |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/KAI-7B-Instruct/README.md) |
| Kaori-70B-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Kaori-70B-v1/README.md) |
| Karen_theEditor_13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/Karen_theEditor_13B/README.md) |
| Karen_TheEditor_V2_CREATIVE_Mistral_7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Karen_TheEditor_V2_CREATIVE_Mistral_7B/README.md) |
| Karen_TheEditor_V2_STRICT_Mistral_7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Karen_TheEditor_V2_STRICT_Mistral_7B/README.md) |
| Kimiko-Mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Kimiko-Mistral-7B/README.md) |
| koOpenChat-sft |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/koOpenChat-sft/README.md) |
| Kuchiki-1.1-L2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Kuchiki-1.1-L2-7B/README.md) |
| Kuchiki-L2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Kuchiki-L2-7B/README.md) |
| law-LLM |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 0 | 4096 | [README.md](models/law-LLM/README.md) |
| Lemur-70B-Chat-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Lemur-70B-Chat-v1/README.md) |
| leo-hessianai-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 5120 | [README.md](models/leo-hessianai-13B/README.md) |
| leo-hessianai-13B-chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 5120 | [README.md](models/leo-hessianai-13B-chat/README.md) |
| leo-hessianai-13B-chat-bilingual |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 5120 | [README.md](models/leo-hessianai-13B-chat-bilingual/README.md) |
| leo-hessianai-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 8192 | 8192 | [README.md](models/leo-hessianai-70B/README.md) |
| leo-hessianai-70B-chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 8192 | 8192 | [README.md](models/leo-hessianai-70B-chat/README.md) |
| leo-hessianai-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/leo-hessianai-7B/README.md) |
| leo-hessianai-7B-chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/leo-hessianai-7B-chat/README.md) |
| leo-hessianai-7B-chat-bilingual |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/leo-hessianai-7B-chat-bilingual/README.md) |
| Leo-Mistral-Hessianai-7B-Chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Leo-Mistral-Hessianai-7B-Chat/README.md) |
| LeoScorpius-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/LeoScorpius-7B/README.md) |
| LeoScorpius-GreenNode-Alpaca-7B-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/LeoScorpius-GreenNode-Alpaca-7B-v1/README.md) |
| LeoScorpius-GreenNode-Platypus-7B-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/LeoScorpius-GreenNode-Platypus-7B-v1/README.md) |
| Lewd-Sydney-20B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Lewd-Sydney-20B/README.md) |
| Lila-103B-L2 |Q2_K Q3_K_M Q3_K_S | 4096 | 8192 | [README.md](models/Lila-103B-L2/README.md) |
| Lila-70B-L2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Lila-70B-L2/README.md) |
| lince-zero |Q4_0 Q4_1 Q5_0 Q5_1 Q8_0 | 0 | 4544 | [README.md](models/lince-zero/README.md) |
| llama-13b-supercot |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/llama-13b-supercot/README.md) |
| Llama-2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Llama-2-13B/README.md) |
| Llama-2-13B-chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Llama-2-13B-chat/README.md) |
| Llama-2-13B-Chat-Dutch |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Llama-2-13B-Chat-Dutch/README.md) |
| llama-2-13B-chat-limarp-v2-merged |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/llama-2-13B-chat-limarp-v2-merged/README.md) |
| llama-2-13B-German-Assistant-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/llama-2-13B-German-Assistant-v2/README.md) |
| Llama-2-13B-German-Assistant-v4 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/Llama-2-13B-German-Assistant-v4/README.md) |
| llama-2-13B-Guanaco-QLoRA |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/llama-2-13B-Guanaco-QLoRA/README.md) |
| Llama2-13B-MegaCode2-OASST |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/Llama2-13B-MegaCode2-OASST/README.md) |
| LLaMA2-13B-Psyfighter2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/LLaMA2-13B-Psyfighter2/README.md) |
| LLaMA_2_13B_SFT_v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/LLaMA_2_13B_SFT_v1/README.md) |
| LLaMA2-13B-Tiefighter |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/LLaMA2-13B-Tiefighter/README.md) |
| LLaMA2-13B-TiefighterLR |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/LLaMA2-13B-TiefighterLR/README.md) |
| llama2-22B-daydreamer-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4098 | 6656 | [README.md](models/llama2-22B-daydreamer-v2/README.md) |
| Llama2-22B-Daydreamer-v3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4098 | 6656 | [README.md](models/Llama2-22B-Daydreamer-v3/README.md) |
| Llama-2-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Llama-2-70B/README.md) |
| Llama-2-70B-Chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Llama-2-70B-Chat/README.md) |
| llama2_70b_chat_uncensored |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 2048 | 8192 | [README.md](models/llama2_70b_chat_uncensored/README.md) |
| Llama-2-70B-OASST-1-200 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Llama-2-70B-OASST-1-200/README.md) |
| Llama2-70B-OASST-SFT-v10 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Llama2-70B-OASST-SFT-v10/README.md) |
| Llama-2-70B-Orca-200k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Llama-2-70B-Orca-200k/README.md) |
| Llama-2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Llama-2-7B/README.md) |
| Llama-2-7B-32K-Instruct |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Llama-2-7B-32K-Instruct/README.md) |
| llama-2-7B-Arguments |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/llama-2-7B-Arguments/README.md) |
| Llama-2-7B-Chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Llama-2-7B-Chat/README.md) |
| llama2-7b-chat-codeCherryPop-qLoRA |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/llama2-7b-chat-codeCherryPop-qLoRA/README.md) |
| llama2_7b_chat_uncensored |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/llama2_7b_chat_uncensored/README.md) |
| llama-2-7B-Guanaco-QLoRA |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/llama-2-7B-Guanaco-QLoRA/README.md) |
| llama2-7B-layla |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/llama2-7B-layla/README.md) |
| llama2_7b_merge_orcafamily |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/llama2_7b_merge_orcafamily/README.md) |
| Llama-2-7B-vietnamese-20k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Llama-2-7B-vietnamese-20k/README.md) |
| Llama2-chat-AYB-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Llama2-chat-AYB-13B/README.md) |
| Llama2-Chat-AYT-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Llama2-Chat-AYT-13B/README.md) |
| Llama-2-Coder-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Llama-2-Coder-7B/README.md) |
| llama-30b-supercot |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/llama-30b-supercot/README.md) |
| LlamaGuard-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/LlamaGuard-7B/README.md) |
| llama-polyglot-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/llama-polyglot-13B/README.md) |
| Llamix2-MLewd-4x13B |Q4_0 Q4_1 Q5_0 Q5_1 Q8_0 | 32768 | 5120 | [README.md](models/Llamix2-MLewd-4x13B/README.md) |
| llemma_34b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 8192 | [README.md](models/llemma_34b/README.md) |
| llemma_7b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/llemma_7b/README.md) |
| LlongOrca-13B-16K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 5120 | [README.md](models/LlongOrca-13B-16K/README.md) |
| LlongOrca-7B-16K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 4096 | [README.md](models/LlongOrca-7B-16K/README.md) |
| LoKuS-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/LoKuS-13B/README.md) |
| LongAlpaca-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/LongAlpaca-70B/README.md) |
| LosslessMegaCoder-Llama2-13B-Mini |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/LosslessMegaCoder-Llama2-13B-Mini/README.md) |
| LosslessMegaCoder-Llama2-7B-Mini |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/LosslessMegaCoder-Llama2-7B-Mini/README.md) |
| loyal-piano-m7 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/loyal-piano-m7/README.md) |
| Luban-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Luban-13B/README.md) |
| Luna-AI-Llama2-Uncensored |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/Luna-AI-Llama2-Uncensored/README.md) |
| lzlv_70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/lzlv_70B/README.md) |
| MadMix-v0.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/MadMix-v0.2/README.md) |
| Magicoder-S-DS-6.7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 4096 | [README.md](models/Magicoder-S-DS-6.7B/README.md) |
| MAmmoTH-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MAmmoTH-13B/README.md) |
| MAmmoTH-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/MAmmoTH-70B/README.md) |
| MAmmoTH-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/MAmmoTH-7B/README.md) |
| MAmmoTH-Coder-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 5120 | [README.md](models/MAmmoTH-Coder-13B/README.md) |
| MAmmoTH-Coder-34B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/MAmmoTH-Coder-34B/README.md) |
| Manticore-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/Manticore-13B/README.md) |
| manticore-13b-chat-pyg |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/manticore-13b-chat-pyg/README.md) |
| Marcoroni-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Marcoroni-13B/README.md) |
| Marcoroni-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Marcoroni-70B/README.md) |
| Marcoroni-70B-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Marcoroni-70B-v1/README.md) |
| Marcoroni-7b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Marcoroni-7b/README.md) |
| Marcoroni-7B-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Marcoroni-7B-v2/README.md) |
| Marcoroni-7B-v3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Marcoroni-7B-v3/README.md) |
| Marcoroni-neural-chat-7B-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Marcoroni-neural-chat-7B-v1/README.md) |
| Marx-3B-v3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 2560 | [README.md](models/Marx-3B-v3/README.md) |
| medalpaca-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 0 | 5120 | [README.md](models/medalpaca-13B/README.md) |
| meditron-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/meditron-70B/README.md) |
| meditron-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/meditron-7B/README.md) |
| meditron-7B-chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/meditron-7B-chat/README.md) |
| Megamix-A1-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Megamix-A1-13B/README.md) |
| MegaMix-S1-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MegaMix-S1-13B/README.md) |
| MegaMix-T1-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MegaMix-T1-13B/README.md) |
| MelloGPT |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/MelloGPT/README.md) |
| Merged-AGI-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Merged-AGI-7B/README.md) |
| MergeMonster-13B-20231124 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MergeMonster-13B-20231124/README.md) |
| merlyn-education-corpus-qa-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/merlyn-education-corpus-qa-v2/README.md) |
| merlyn-education-safety |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/merlyn-education-safety/README.md) |
| MetaMath-13B-V1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MetaMath-13B-V1.0/README.md) |
| MetaMath-70B-V1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/MetaMath-70B-V1.0/README.md) |
| MetaMath-7B-V1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/MetaMath-7B-V1.0/README.md) |
| MetaMath-Cybertron-Starling |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/MetaMath-Cybertron-Starling/README.md) |
| MetaMath-Mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/MetaMath-Mistral-7B/README.md) |
| Metis-0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Metis-0.1/README.md) |
| mindy-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/mindy-7B/README.md) |
| Mini_Synatra_7B_02 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mini_Synatra_7B_02/README.md) |
| minotaur-13B-fixed |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/minotaur-13B-fixed/README.md) |
| Mistral-11B-CC-Air-RP |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-11B-CC-Air-RP/README.md) |
| Mistral-11B-OmniMix |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-11B-OmniMix/README.md) |
| Mistral-7B-AEZAKMI-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-7B-AEZAKMI-v1/README.md) |
| Mistral-7B-Claude-Chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-7B-Claude-Chat/README.md) |
| Mistral-7B-Code-16K-qlora |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-7B-Code-16K-qlora/README.md) |
| Mistral-7B-codealpaca-lora |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-7B-codealpaca-lora/README.md) |
| Mistral_7B_Dolphin2.1_LIMA0.5 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral_7B_Dolphin2.1_LIMA0.5/README.md) |
| mistral-7B-dpo-v5 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/mistral-7B-dpo-v5/README.md) |
| mistral-7B-finetuned-orca-dpo-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/mistral-7B-finetuned-orca-dpo-v2/README.md) |
| Mistral-7B-Instruct-v0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-7B-Instruct-v0.1/README.md) |
| Mistral-7B-Instruct-v0.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-7B-Instruct-v0.2/README.md) |
| Mistral-7B-Instruct-v0.2-DARE |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-7B-Instruct-v0.2-DARE/README.md) |
| Mistral-7B-OpenOrca |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-7B-OpenOrca/README.md) |
| Mistral-7B-OpenOrca-oasst_top1_2023-08-25-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-7B-OpenOrca-oasst_top1_2023-08-25-v1/README.md) |
| Mistral-7B-Phibrarian-32K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-7B-Phibrarian-32K/README.md) |
| Mistral-7B-SciPhi-32k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-7B-SciPhi-32k/README.md) |
| Mistral-7B-v0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-7B-v0.1/README.md) |
| Mistral-ClaudeLimaRP-v3-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-ClaudeLimaRP-v3-7B/README.md) |
| Mistralic-7B-1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistralic-7B-1/README.md) |
| MistralLite-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/MistralLite-7B/README.md) |
| MistralMakise-Merged-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MistralMakise-Merged-13B/README.md) |
| Mistral-Pygmalion-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Mistral-Pygmalion-7B/README.md) |
| Mistral-Trismegistus-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mistral-Trismegistus-7B/README.md) |
| MistRP-Airoboros-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/MistRP-Airoboros-7B/README.md) |
| Mixtral-8x7B-Instruct-v0.1 |Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mixtral-8x7B-Instruct-v0.1/README.md) |
| Mixtral-8x7B-MoE-RP-Story |Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mixtral-8x7B-MoE-RP-Story/README.md) |
| Mixtral-8x7B-v0.1 |Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mixtral-8x7B-v0.1/README.md) |
| Mixtral-Fusion-4x7B-Instruct-v0.1 |Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mixtral-Fusion-4x7B-Instruct-v0.1/README.md) |
| mixtralnt-4x7b-test |Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/mixtralnt-4x7b-test/README.md) |
| Mixtral-SlimOrca-8x7B |Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Mixtral-SlimOrca-8x7B/README.md) |
| MLewdBoros-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MLewdBoros-L2-13B/README.md) |
| MLewdBoros-LRPSGPT-2Char-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MLewdBoros-LRPSGPT-2Char-13B/README.md) |
| MLewd-L2-Chat-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MLewd-L2-Chat-13B/README.md) |
| MLewd-ReMM-L2-Chat-20B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MLewd-ReMM-L2-Chat-20B/README.md) |
| MLewd-ReMM-L2-Chat-20B-Inverted |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MLewd-ReMM-L2-Chat-20B-Inverted/README.md) |
| model_007-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/model_007-70B/README.md) |
| MoMo-70B-V1.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/MoMo-70B-V1.1/README.md) |
| MonadGPT |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/MonadGPT/README.md) |
| MXLewd-L2-20B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MXLewd-L2-20B/README.md) |
| MXLewdMini-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MXLewdMini-L2-13B/README.md) |
| MysticFusion-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MysticFusion-13B/README.md) |
| Mythalion-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Mythalion-13B/README.md) |
| Mythical-Destroyer-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Mythical-Destroyer-L2-13B/README.md) |
| Mythical-Destroyer-V2-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Mythical-Destroyer-V2-L2-13B/README.md) |
| MythoBoros-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/MythoBoros-13B/README.md) |
| MythoLogic-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/MythoLogic-13B/README.md) |
| MythoLogic-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MythoLogic-L2-13B/README.md) |
| MythoLogic-Mini-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/MythoLogic-Mini-7B/README.md) |
| MythoMakiseMerged-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MythoMakiseMerged-13B/README.md) |
| MythoMax-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MythoMax-L2-13B/README.md) |
| MythoMax-L2-Kimiko-v2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MythoMax-L2-Kimiko-v2-13B/README.md) |
| MythoMist-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/MythoMist-7B/README.md) |
| MythoMix-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MythoMix-L2-13B/README.md) |
| Naberius-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Naberius-7B/README.md) |
| Nanbeige-16B-Base |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Nanbeige-16B-Base/README.md) |
| Nanbeige-16B-Base-32K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Nanbeige-16B-Base-32K/README.md) |
| Nanbeige-16B-Chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Nanbeige-16B-Chat/README.md) |
| Nanbeige-16B-Chat-32K |fp16 Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Nanbeige-16B-Chat-32K/README.md) |
| Nete-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Nete-13B/README.md) |
| Nethena-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Nethena-13B/README.md) |
| Nethena-20B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Nethena-20B/README.md) |
| Nethena-20B-Glued |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Nethena-20B-Glued/README.md) |
| Nethena-MLewd-Xwin-23B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Nethena-MLewd-Xwin-23B/README.md) |
| neural-chat-7B-v3-1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/neural-chat-7B-v3-1/README.md) |
| neural-chat-7B-v3-2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/neural-chat-7B-v3-2/README.md) |
| neural-chat-7B-v3-3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/neural-chat-7B-v3-3/README.md) |
| NeuralHermes-2.5-Mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/NeuralHermes-2.5-Mistral-7B/README.md) |
| NeuralOrca-7B-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/NeuralOrca-7B-v1/README.md) |
| NexusRaven-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 5120 | [README.md](models/NexusRaven-13B/README.md) |
| NexusRaven-V2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 5120 | [README.md](models/NexusRaven-V2-13B/README.md) |
| Noromaid-13B-v0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Noromaid-13B-v0.1/README.md) |
| Noromaid-13B-v0.1.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Noromaid-13B-v0.1.1/README.md) |
| Noromaid-20B-v0.1.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Noromaid-20B-v0.1.1/README.md) |
| notus-7B-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/notus-7B-v1/README.md) |
| Nous-Hermes-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/Nous-Hermes-13B/README.md) |
| Nous-Hermes-13B-Code |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Nous-Hermes-13B-Code/README.md) |
| Nous-Puffin-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Nous-Puffin-70B/README.md) |
| nsql-llama-2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/nsql-llama-2-7B/README.md) |
| nucleus-22B-token-500B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/nucleus-22B-token-500B/README.md) |
| Nyxene-v2-11B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Nyxene-v2-11B/README.md) |
| Nyxene-v3-11B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Nyxene-v3-11B/README.md) |
| OpenAssistant-Llama2-13B-Orca-8K-3319 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 5120 | [README.md](models/OpenAssistant-Llama2-13B-Orca-8K-3319/README.md) |
| openbuddy-coder-34b-v11-bf16 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/openbuddy-coder-34b-v11-bf16/README.md) |
| openbuddy-deepseek-67b-v15-base |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/openbuddy-deepseek-67b-v15-base/README.md) |
| openbuddy-llama2-13b64k-v15 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 65536 | 5120 | [README.md](models/openbuddy-llama2-13b64k-v15/README.md) |
| OpenBuddy-Llama2-13B-v11.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/OpenBuddy-Llama2-13B-v11.1/README.md) |
| openbuddy-llama2-34b-v11.1-bf16 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/openbuddy-llama2-34b-v11.1-bf16/README.md) |
| OpenBuddy-Llama2-70b-v10.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/OpenBuddy-Llama2-70b-v10.1/README.md) |
| openbuddy-llama2-70B-v13.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/openbuddy-llama2-70B-v13.2/README.md) |
| openbuddy-llama2-70B-v13-base |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/openbuddy-llama2-70B-v13-base/README.md) |
| openbuddy-mistral-7B-v13 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/openbuddy-mistral-7B-v13/README.md) |
| openbuddy-mistral-7B-v13.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/openbuddy-mistral-7B-v13.1/README.md) |
| openbuddy-mistral-7B-v13-base |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/openbuddy-mistral-7B-v13-base/README.md) |
| openbuddy-mixtral-8x7b-v15.1 |Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/openbuddy-mixtral-8x7b-v15.1/README.md) |
| openbuddy-openllama-7B-v12-bf16 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/openbuddy-openllama-7B-v12-bf16/README.md) |
| openbuddy-zephyr-7B-v14.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/openbuddy-zephyr-7B-v14.1/README.md) |
| openchat_3.5 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/openchat_3.5/README.md) |
| openchat-3.5-1210 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/openchat-3.5-1210/README.md) |
| openchat_3.5-16k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/openchat_3.5-16k/README.md) |
| openchat_v3.2_super |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/openchat_v3.2_super/README.md) |
| OpenHermes-2.5-Mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/OpenHermes-2.5-Mistral-7B/README.md) |
| OpenHermes-2.5-Mistral-7B-16k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/OpenHermes-2.5-Mistral-7B-16k/README.md) |
| Open-Hermes-2.5-neural-chat-3.1-frankenmerge-11b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Open-Hermes-2.5-neural-chat-3.1-frankenmerge-11b/README.md) |
| OpenHermes-2.5-neural-chat-7B-v3-1-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/OpenHermes-2.5-neural-chat-7B-v3-1-7B/README.md) |
| OpenHermes-2.5-neural-chat-7B-v3-2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/OpenHermes-2.5-neural-chat-7B-v3-2-7B/README.md) |
| OpenHermes-2.5-neural-chat-v3-3-Slerp |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/OpenHermes-2.5-neural-chat-v3-3-Slerp/README.md) |
| OpenHermes-2-Mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/OpenHermes-2-Mistral-7B/README.md) |
| open-instruct-human-mix-65B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 2048 | 8192 | [README.md](models/open-instruct-human-mix-65B/README.md) |
| openinstruct-mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/openinstruct-mistral-7B/README.md) |
| open-llama-3b-v2-wizard-evol-instuct-v2-196k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 3200 | [README.md](models/open-llama-3b-v2-wizard-evol-instuct-v2-196k/README.md) |
| OpenOrca-Platypus2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/OpenOrca-Platypus2-13B/README.md) |
| OpenOrca_Stx |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/OpenOrca_Stx/README.md) |
| OpenOrca-Zephyr-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/OpenOrca-Zephyr-7B/README.md) |
| OpenZephyrChat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/OpenZephyrChat/README.md) |
| Optimus-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Optimus-7B/README.md) |
| opus-v0.5-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/opus-v0.5-70B/README.md) |
| opus-v0-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/opus-v0-70B/README.md) |
| opus-v0-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/opus-v0-7B/README.md) |
| Orca-2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Orca-2-13B/README.md) |
| Orca-2-13B-SFT_v5 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Orca-2-13B-SFT_v5/README.md) |
| Orca-2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Orca-2-7B/README.md) |
| ORCA_LLaMA_70B_QLoRA |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/ORCA_LLaMA_70B_QLoRA/README.md) |
| OrcaMaid-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/OrcaMaid-13B/README.md) |
| orca_mini_v3_13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/orca_mini_v3_13B/README.md) |
| orca_mini_v3_70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/orca_mini_v3_70B/README.md) |
| orca_mini_v3_7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/orca_mini_v3_7B/README.md) |
| OrionStar-Yi-34B-Chat-Llama |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 7168 | [README.md](models/OrionStar-Yi-34B-Chat-Llama/README.md) |
| Pallas-0.3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/Pallas-0.3/README.md) |
| Pallas-0.4 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/Pallas-0.4/README.md) |
| Pandalyst_13B_V1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 5120 | [README.md](models/Pandalyst_13B_V1.0/README.md) |
| Pandalyst-7B-V1.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 4096 | [README.md](models/Pandalyst-7B-V1.1/README.md) |
| Pandalyst-7B-v1.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 4096 | [README.md](models/Pandalyst-7B-v1.2/README.md) |
| Pandora-v1-10.7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Pandora-v1-10.7B/README.md) |
| Pandora-v1-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Pandora-v1-13B/README.md) |
| pee |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/pee/README.md) |
| Phind-CodeLlama-34B-Python-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/Phind-CodeLlama-34B-Python-v1/README.md) |
| Phind-CodeLlama-34B-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/Phind-CodeLlama-34B-v1/README.md) |
| Phind-CodeLlama-34B-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/Phind-CodeLlama-34B-v2/README.md) |
| PiVoT-0.1-early |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/PiVoT-0.1-early/README.md) |
| PiVoT-0.1-Evil-a |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/PiVoT-0.1-Evil-a/README.md) |
| PiVoT-0.1-Starling-LM-RP |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/PiVoT-0.1-Starling-LM-RP/README.md) |
| PiVoT-10.7B-Mistral-v0.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/PiVoT-10.7B-Mistral-v0.2/README.md) |
| PiVoT-10.7B-Mistral-v0.2-RP |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/PiVoT-10.7B-Mistral-v0.2-RP/README.md) |
| PiVoT-Merge-A-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/PiVoT-Merge-A-7B/README.md) |
| PiVoT-MoE |Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/PiVoT-MoE/README.md) |
| PlatYi-34B-Llama-Q-v3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 7168 | [README.md](models/PlatYi-34B-Llama-Q-v3/README.md) |
| Platypus2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Platypus2-13B/README.md) |
| Platypus2-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Platypus2-70B/README.md) |
| Platypus2-70B-Instruct |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Platypus2-70B-Instruct/README.md) |
| Platypus-30B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/Platypus-30B/README.md) |
| platypus-yi-34b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 7168 | [README.md](models/platypus-yi-34b/README.md) |
| Poro-34B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 0 | 7168 | [README.md](models/Poro-34B/README.md) |
| prometheus-13B-v1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/prometheus-13B-v1.0/README.md) |
| prometheus-7B-v1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/prometheus-7B-v1.0/README.md) |
| Psyfighter-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Psyfighter-13B/README.md) |
| PsyMedRP-v1-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/PsyMedRP-v1-13B/README.md) |
| PsyMedRP-v1-20B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/PsyMedRP-v1-20B/README.md) |
| psyonic-cetacean-20B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/psyonic-cetacean-20B/README.md) |
| PuddleJumper-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_1 Q4_K_M Q4_K_S Q5_0 Q5_1 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/PuddleJumper-13B/README.md) |
| PuddleJumper-13B-V2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/PuddleJumper-13B-V2/README.md) |
| Pygmalion-2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Pygmalion-2-13B/README.md) |
| Pygmalion-2-13B-SuperCOT |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Pygmalion-2-13B-SuperCOT/README.md) |
| Pygmalion-2-13B-SuperCOT2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Pygmalion-2-13B-SuperCOT2/README.md) |
| Pygmalion-2-13B-SuperCOT-weighed |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Pygmalion-2-13B-SuperCOT-weighed/README.md) |
| Pygmalion-2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Pygmalion-2-7B/README.md) |
| Python-Code-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Python-Code-13B/README.md) |
| Python-Code-33B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/Python-Code-33B/README.md) |
| qCammel-13 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/qCammel-13/README.md) |
| qCammel-70-x |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/qCammel-70-x/README.md) |
| Redmond-Puffin-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Redmond-Puffin-13B/README.md) |
| ReMM-SLERP-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/ReMM-SLERP-L2-13B/README.md) |
| ReMM-v2.1-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/ReMM-v2.1-L2-13B/README.md) |
| ReMM-v2-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/ReMM-v2-L2-13B/README.md) |
| rocket-3B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 2560 | [README.md](models/rocket-3B/README.md) |
| Rogue-Rose-103b-v0.2 |Q2_K Q3_K_M Q3_K_S | 4096 | 8192 | [README.md](models/Rogue-Rose-103b-v0.2/README.md) |
| Rose-20B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Rose-20B/README.md) |
| Rose-Kimiko-20B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Rose-Kimiko-20B/README.md) |
| RpBird-Yi-34B-200k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/RpBird-Yi-34B-200k/README.md) |
| rpguild-chatml-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/rpguild-chatml-13B/README.md) |
| sabia-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/sabia-7B/README.md) |
| saiga_mistral_7b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/saiga_mistral_7b/README.md) |
| Samantha-1.11-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Samantha-1.11-13B/README.md) |
| Samantha-1.11-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Samantha-1.11-70B/README.md) |
| Samantha-1.11-CodeLlama-34B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 8192 | [README.md](models/Samantha-1.11-CodeLlama-34B/README.md) |
| Samantha-1.1-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Samantha-1.1-70B/README.md) |
| samantha-1.2-mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/samantha-1.2-mistral-7B/README.md) |
| samantha-mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/samantha-mistral-7B/README.md) |
| samantha-mistral-instruct-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/samantha-mistral-instruct-7B/README.md) |
| SauerkrautLM-13B-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/SauerkrautLM-13B-v1/README.md) |
| SauerkrautLM-3B-v1 |Q4_0 Q4_1 Q5_0 Q5_1 Q8_0 | 2048 | 3200 | [README.md](models/SauerkrautLM-3B-v1/README.md) |
| SauerkrautLM-70B-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/SauerkrautLM-70B-v1/README.md) |
| SauerkrautLM-7B-HerO |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/SauerkrautLM-7B-HerO/README.md) |
| SauerkrautLM-7B-v1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/SauerkrautLM-7B-v1/README.md) |
| SauerkrautLM-7B-v1-mistral |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/SauerkrautLM-7B-v1-mistral/README.md) |
| Scarlett-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/Scarlett-7B/README.md) |
| SciPhi-Mistral-7B-32k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/SciPhi-Mistral-7B-32k/README.md) |
| SciPhi-Self-RAG-Mistral-7B-32k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/SciPhi-Self-RAG-Mistral-7B-32k/README.md) |
| Seraph-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Seraph-7B/README.md) |
| SeraphMarcoroni-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/SeraphMarcoroni-7B/README.md) |
| SG-Raccoon-Yi-200k-2.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K | 200000 | 7168 | [README.md](models/SG-Raccoon-Yi-200k-2.0/README.md) |
| SG-Raccoon-Yi-55B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K | 4096 | 7168 | [README.md](models/SG-Raccoon-Yi-55B/README.md) |
| SG-Raccoon-Yi-55B-200k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K | 200000 | 7168 | [README.md](models/SG-Raccoon-Yi-55B-200k/README.md) |
| sheep-duck-llama-2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/sheep-duck-llama-2-13B/README.md) |
| Sheep-Duck-Llama-2-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Sheep-Duck-Llama-2-70B/README.md) |
| sheep-duck-llama-2-70B-v1.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/sheep-duck-llama-2-70B-v1.1/README.md) |
| ShiningValiant-1.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/ShiningValiant-1.2/README.md) |
| ShiningValiant-1.3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/ShiningValiant-1.3/README.md) |
| ShiningValiantXS |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/ShiningValiantXS/README.md) |
| ShiningValiantXS-1.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/ShiningValiantXS-1.1/README.md) |
| Skywork-13B-base |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4608 | [README.md](models/Skywork-13B-base/README.md) |
| SlimOpenOrca-Mistral-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/SlimOpenOrca-Mistral-7B/README.md) |
| SlimOrca-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/SlimOrca-13B/README.md) |
| smartyplats-7B-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/smartyplats-7B-v2/README.md) |
| smol-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/smol-7B/README.md) |
| SOLAR-10.7B-Instruct-v1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/SOLAR-10.7B-Instruct-v1.0/README.md) |
| Solar-10.7B-SLERP |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Solar-10.7B-SLERP/README.md) |
| SOLAR-10.7B-v1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/SOLAR-10.7B-v1.0/README.md) |
| SOLAR-Platypus-10.7B-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/SOLAR-Platypus-10.7B-v2/README.md) |
| Solus-103B-L2 |Q2_K Q3_K_M Q3_K_S | 4096 | 8192 | [README.md](models/Solus-103B-L2/README.md) |
| Solus-70B-L2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Solus-70B-L2/README.md) |
| speechless-codellama-34b-v2.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/speechless-codellama-34b-v2.0/README.md) |
| speechless-code-mistral-7B-v1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/speechless-code-mistral-7B-v1.0/README.md) |
| Speechless-Llama2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Speechless-Llama2-13B/README.md) |
| Speechless-Llama2-Hermes-Orca-Platypus-WizardLM-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Speechless-Llama2-Hermes-Orca-Platypus-WizardLM-13B/README.md) |
| speechless-mistral-7B-dare-0.85 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/speechless-mistral-7B-dare-0.85/README.md) |
| speechless-mistral-dolphin-orca-platypus-samantha-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/speechless-mistral-dolphin-orca-platypus-samantha-7B/README.md) |
| speechless-tora-code-7B-v1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 4096 | [README.md](models/speechless-tora-code-7B-v1.0/README.md) |
| Spicyboros-13B-2.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Spicyboros-13B-2.2/README.md) |
| Spicyboros-70B-2.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Spicyboros-70B-2.2/README.md) |
| Spicyboros-7B-2.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Spicyboros-7B-2.2/README.md) |
| Spicyboros-c34b-2.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/Spicyboros-c34b-2.2/README.md) |
| Spring-Dragon |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Spring-Dragon/README.md) |
| sqlcoder |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 0 | 0 | [README.md](models/sqlcoder/README.md) |
| sqlcoder2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 0 | 0 | [README.md](models/sqlcoder2/README.md) |
| sqlcoder-34b-alpha |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/sqlcoder-34b-alpha/README.md) |
| sqlcoder-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/sqlcoder-7B/README.md) |
| StableBeluga-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/StableBeluga-13B/README.md) |
| StableBeluga2-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/StableBeluga2-70B/README.md) |
| StableBeluga-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/StableBeluga-7B/README.md) |
| stablelm-zephyr-3b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 2560 | [README.md](models/stablelm-zephyr-3b/README.md) |
| Stable-Platypus2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Stable-Platypus2-13B/README.md) |
| stable-vicuna-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/stable-vicuna-13B/README.md) |
| Starling-LM-7B-alpha |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/Starling-LM-7B-alpha/README.md) |
| Starling-LM-alpha-8x7B-MoE |Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Starling-LM-alpha-8x7B-MoE/README.md) |
| StellarBright |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/StellarBright/README.md) |
| Stheno-Inverted-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Stheno-Inverted-L2-13B/README.md) |
| Stheno-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Stheno-L2-13B/README.md) |
| Stheno-v2-Delta |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Stheno-v2-Delta/README.md) |
| stockmark-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/stockmark-13B/README.md) |
| storytime-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/storytime-13B/README.md) |
| SunsetBoulevard |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/SunsetBoulevard/README.md) |
| supermario-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/supermario-v2/README.md) |
| SuperPlatty-30B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/SuperPlatty-30B/README.md) |
| SUS-Chat-34B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 7168 | [README.md](models/SUS-Chat-34B/README.md) |
| Sydney_Overthinker_13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Sydney_Overthinker_13B/README.md) |
| Synatra-7B-v0.3-base |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Synatra-7B-v0.3-base/README.md) |
| Synatra-7B-v0.3-dpo |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Synatra-7B-v0.3-dpo/README.md) |
| Synatra-7B-v0.3-RP |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Synatra-7B-v0.3-RP/README.md) |
| Synatra-RP-Orca-2-7B-v0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Synatra-RP-Orca-2-7B-v0.1/README.md) |
| Synatra-V0.1-7B-Instruct |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Synatra-V0.1-7B-Instruct/README.md) |
| Synthia-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Synthia-13B/README.md) |
| Synthia-13B-v1.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Synthia-13B-v1.2/README.md) |
| Synthia-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q5_K_S | 2048 | 8192 | [README.md](models/Synthia-70B/README.md) |
| Synthia-70B-v1.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 2048 | 8192 | [README.md](models/Synthia-70B-v1.1/README.md) |
| Synthia-70B-v1.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 2048 | 8192 | [README.md](models/Synthia-70B-v1.2/README.md) |
| Synthia-70B-v1.2b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Synthia-70B-v1.2b/README.md) |
| SynthIA-70B-v1.5 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 2048 | 8192 | [README.md](models/SynthIA-70B-v1.5/README.md) |
| Synthia-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/Synthia-7B/README.md) |
| Synthia-7B-v1.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/Synthia-7B-v1.2/README.md) |
| Synthia-7B-v1.3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Synthia-7B-v1.3/README.md) |
| SynthIA-7B-v1.3-dare-0.85 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/SynthIA-7B-v1.3-dare-0.85/README.md) |
| SynthIA-7B-v1.5 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/SynthIA-7B-v1.5/README.md) |
| SynthIA-7B-v2.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/SynthIA-7B-v2.0/README.md) |
| SynthIA-7B-v2.0-16k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/SynthIA-7B-v2.0-16k/README.md) |
| Synthia-7B-v3.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Synthia-7B-v3.0/README.md) |
| Synthia-MoE-v3-Mixtral-8x7B |Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Synthia-MoE-v3-Mixtral-8x7B/README.md) |
| Tai-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Tai-70B/README.md) |
| Terminis-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Terminis-7B/README.md) |
| Tess-34B-v1.4 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/Tess-34B-v1.4/README.md) |
| Tess-7B-v1.4 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Tess-7B-v1.4/README.md) |
| Tess-M-Creative-v1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/Tess-M-Creative-v1.0/README.md) |
| Tess-M-v1.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/Tess-M-v1.1/README.md) |
| Tess-M-v1.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/Tess-M-v1.2/README.md) |
| Tess-M-v1.3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/Tess-M-v1.3/README.md) |
| Tess-XL-v1.0 |Q2_K | 4096 | 8192 | [README.md](models/Tess-XL-v1.0/README.md) |
| Tess-XS-Creative-v1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Tess-XS-Creative-v1.0/README.md) |
| Tess-XS-v1.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Tess-XS-v1.1/README.md) |
| Thespis-13B-Alpha-v0.7 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Thespis-13B-Alpha-v0.7/README.md) |
| Thespis-13B-v0.3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Thespis-13B-v0.3/README.md) |
| Thespis-13B-v0.4 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Thespis-13B-v0.4/README.md) |
| Thespis-13B-v0.5 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Thespis-13B-v0.5/README.md) |
| Thespis-13B-v0.6 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Thespis-13B-v0.6/README.md) |
| Thespis-Mistral-7B-v0.5 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Thespis-Mistral-7B-v0.5/README.md) |
| Thespis-Mistral-7B-v0.6 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Thespis-Mistral-7B-v0.6/README.md) |
| Tiamat-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Tiamat-7B/README.md) |
| Tiamat-7B-1.1-DPO |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Tiamat-7B-1.1-DPO/README.md) |
| TigerBot-70B-Chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 2048 | 8192 | [README.md](models/TigerBot-70B-Chat/README.md) |
| tigerbot-70B-chat-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/tigerbot-70B-chat-v2/README.md) |
| tigerbot-70B-chat-v4 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/tigerbot-70B-chat-v4/README.md) |
| TimeCrystal-L2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/TimeCrystal-L2-13B/README.md) |
| TinyLlama-1.1B-1T-OpenOrca |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 2048 | [README.md](models/TinyLlama-1.1B-1T-OpenOrca/README.md) |
| TinyLlama-1.1B-Chat-v0.3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 2048 | [README.md](models/TinyLlama-1.1B-Chat-v0.3/README.md) |
| TinyLlama-1.1B-intermediate-step-480k-1T |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 2048 | [README.md](models/TinyLlama-1.1B-intermediate-step-480k-1T/README.md) |
| TinyLlama-1.1B-intermediate-step-715k-1.5T |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 2048 | [README.md](models/TinyLlama-1.1B-intermediate-step-715k-1.5T/README.md) |
| TinyLlama-1.1B-python-v0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 2048 | [README.md](models/TinyLlama-1.1B-python-v0.1/README.md) |
| Toppy-M-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Toppy-M-7B/README.md) |
| tora-13B-v1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/tora-13B-v1.0/README.md) |
| tora-70B-v1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/tora-70B-v1.0/README.md) |
| tora-7B-v1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/tora-7B-v1.0/README.md) |
| tora-code-13B-v1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 5120 | [README.md](models/tora-code-13B-v1.0/README.md) |
| tora-code-34b-v1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/tora-code-34b-v1.0/README.md) |
| tora-code-7B-v1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 4096 | [README.md](models/tora-code-7B-v1.0/README.md) |
| Trion-M-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Trion-M-7B/README.md) |
| Trurl-2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Trurl-2-13B/README.md) |
| Trurl-2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Trurl-2-7B/README.md) |
| Tulpar-7B-v0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Tulpar-7B-v0/README.md) |
| tulu-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/tulu-13B/README.md) |
| tulu-2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 5120 | [README.md](models/tulu-2-13B/README.md) |
| tulu-2-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 8192 | 8192 | [README.md](models/tulu-2-70B/README.md) |
| tulu-2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/tulu-2-7B/README.md) |
| tulu-2-dpo-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 5120 | [README.md](models/tulu-2-dpo-13B/README.md) |
| tulu-2-dpo-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 8192 | 8192 | [README.md](models/tulu-2-dpo-70B/README.md) |
| tulu-2-dpo-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/tulu-2-dpo-7B/README.md) |
| tulu-30B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/tulu-30B/README.md) |
| tulu-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/tulu-7B/README.md) |
| U-Amethyst-20B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/U-Amethyst-20B/README.md) |
| UltraLM-13B-v2.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/UltraLM-13B-v2.0/README.md) |
| una-cybertron-7B-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/una-cybertron-7B-v2/README.md) |
| una-cybertron-7B-v3-OMA |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/una-cybertron-7B-v3-OMA/README.md) |
| una-xaberius-34b-v1beta |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 7168 | [README.md](models/una-xaberius-34b-v1beta/README.md) |
| Uncensored-Frank-13b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Uncensored-Frank-13b/README.md) |
| Uncensored-Frank-33b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/Uncensored-Frank-33b/README.md) |
| Uncensored-Frank-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/Uncensored-Frank-7B/README.md) |
| Uncensored-Jordan-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Uncensored-Jordan-13B/README.md) |
| Uncensored-Jordan-33B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/Uncensored-Jordan-33B/README.md) |
| Uncensored-Jordan-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/Uncensored-Jordan-7B/README.md) |
| UndiMix-v1-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/UndiMix-v1-13B/README.md) |
| UndiMix-v2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/UndiMix-v2-13B/README.md) |
| UndiMix-v3-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/UndiMix-v3-13B/README.md) |
| UndiMix-v4-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/UndiMix-v4-13B/README.md) |
| Unholy-v1-10l-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Unholy-v1-10l-13B/README.md) |
| Unholy-v1-12L-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Unholy-v1-12L-13B/README.md) |
| Uni-TianYan-70B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Uni-TianYan-70B/README.md) |
| Upstage-Llama1-65B-Instruct |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 2048 | 8192 | [README.md](models/Upstage-Llama1-65B-Instruct/README.md) |
| Upstage-Llama-2-70B-instruct-v2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Upstage-Llama-2-70B-instruct-v2/README.md) |
| upstage-llama-30b-instruct-2048 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/upstage-llama-30b-instruct-2048/README.md) |
| Utopia-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Utopia-13B/README.md) |
| UtopiaXL-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/UtopiaXL-13B/README.md) |
| v1olet_marcoroni-go-bruins-merge-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/v1olet_marcoroni-go-bruins-merge-7B/README.md) |
| v1olet_merged_dpo_7B_v3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/v1olet_merged_dpo_7B_v3/README.md) |
| v1olet_merged_dpo_7B_v4 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/v1olet_merged_dpo_7B_v4/README.md) |
| Velara |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Velara/README.md) |
| Venus-103b-v1.1 |Q2_K Q3_K_M Q3_K_S | 4096 | 8192 | [README.md](models/Venus-103b-v1.1/README.md) |
| Venus-103b-v1.2 |Q2_K Q3_K_M Q3_K_S | 4096 | 8192 | [README.md](models/Venus-103b-v1.2/README.md) |
| Vicuna-13B-CoT |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/Vicuna-13B-CoT/README.md) |
| vicuna-13B-v1.5 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/vicuna-13B-v1.5/README.md) |
| vicuna-13B-v1.5-16K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/vicuna-13B-v1.5-16K/README.md) |
| vicuna-33B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/vicuna-33B/README.md) |
| vicuna-33B-coder |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/vicuna-33B-coder/README.md) |
| Vicuna-7B-CoT |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/Vicuna-7B-CoT/README.md) |
| vicuna-7B-v1.5 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/vicuna-7B-v1.5/README.md) |
| vicuna-7B-v1.5-16K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/vicuna-7B-v1.5-16K/README.md) |
| vietnamese-llama2-7B-40GB |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/vietnamese-llama2-7B-40GB/README.md) |
| Vigogne-2-13B-Instruct |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Vigogne-2-13B-Instruct/README.md) |
| vigogne-2-70B-chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/vigogne-2-70B-chat/README.md) |
| Vigogne-2-7B-Chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Vigogne-2-7B-Chat/README.md) |
| Vigogne-2-7B-Instruct |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/Vigogne-2-7B-Instruct/README.md) |
| Vigostral-7B-Chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Vigostral-7B-Chat/README.md) |
| WinterGoddess-1.4x-70B-L2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/WinterGoddess-1.4x-70B-L2/README.md) |
| WizardCoder-Python-13B-V1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 5120 | [README.md](models/WizardCoder-Python-13B-V1.0/README.md) |
| WizardCoder-Python-34B-V1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/WizardCoder-Python-34B-V1.0/README.md) |
| WizardCoder-Python-7B-V1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 4096 | [README.md](models/WizardCoder-Python-7B-V1.0/README.md) |
| WizardLM-1.0-Uncensored-CodeLlama-34B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/WizardLM-1.0-Uncensored-CodeLlama-34B/README.md) |
| WizardLM-1.0-Uncensored-Llama2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/WizardLM-1.0-Uncensored-Llama2-13B/README.md) |
| WizardLM-13B-1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/WizardLM-13B-1.0/README.md) |
| WizardLM-13B-Uncensored |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/WizardLM-13B-Uncensored/README.md) |
| WizardLM-13B-V1.0-Uncensored |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/WizardLM-13B-V1.0-Uncensored/README.md) |
| WizardLM-13B-V1.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/WizardLM-13B-V1.1/README.md) |
| WizardLM-13B-V1.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/WizardLM-13B-V1.2/README.md) |
| WizardLM-30B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/WizardLM-30B/README.md) |
| WizardLM-30B-Uncensored |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/WizardLM-30B-Uncensored/README.md) |
| WizardLM-33B-V1.0-Uncensored |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/WizardLM-33B-V1.0-Uncensored/README.md) |
| WizardLM-70B-V1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/WizardLM-70B-V1.0/README.md) |
| wizardLM-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/wizardLM-7B/README.md) |
| WizardLM-7B-uncensored |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/WizardLM-7B-uncensored/README.md) |
| WizardLM-7B-V1.0-Uncensored |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/WizardLM-7B-V1.0-Uncensored/README.md) |
| WizardLM-Uncensored-SuperCOT-StoryTelling-30B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/WizardLM-Uncensored-SuperCOT-StoryTelling-30B/README.md) |
| WizardMath-13B-V1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/WizardMath-13B-V1.0/README.md) |
| WizardMath-70B-V1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 2048 | 8192 | [README.md](models/WizardMath-70B-V1.0/README.md) |
| WizardMath-7B-V1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/WizardMath-7B-V1.0/README.md) |
| wizard-mega-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/wizard-mega-13B/README.md) |
| wizard-vicuna-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/wizard-vicuna-13B/README.md) |
| Wizard-Vicuna-13B-Uncensored |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/Wizard-Vicuna-13B-Uncensored/README.md) |
| Wizard-Vicuna-30B-Uncensored |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/Wizard-Vicuna-30B-Uncensored/README.md) |
| Wizard-Vicuna-7B-Uncensored |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/Wizard-Vicuna-7B-Uncensored/README.md) |
| Writing_Partner_Mistral_7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Writing_Partner_Mistral_7B/README.md) |
| X-MythoChronos-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/X-MythoChronos-13B/README.md) |
| X-NoroChronos-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/X-NoroChronos-13B/README.md) |
| XwinCoder-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 5120 | [README.md](models/XwinCoder-13B/README.md) |
| XwinCoder-34B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/XwinCoder-34B/README.md) |
| Xwin-LM-13B-V0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Xwin-LM-13B-V0.1/README.md) |
| Xwin-LM-13B-v0.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Xwin-LM-13B-v0.2/README.md) |
| Xwin-LM-70B-V0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Xwin-LM-70B-V0.1/README.md) |
| Xwin-LM-7B-V0.1 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Xwin-LM-7B-V0.1/README.md) |
| Xwin-LM-7B-V0.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Xwin-LM-7B-V0.2/README.md) |
| Xwin-MLewd-13B-v0.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Xwin-MLewd-13B-v0.2/README.md) |
| Xwin-MLewd-7B-V0.2 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Xwin-MLewd-7B-V0.2/README.md) |
| Yarn-Llama-2-13B-128K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 131072 | 5120 | [README.md](models/Yarn-Llama-2-13B-128K/README.md) |
| Yarn-Llama-2-13B-64K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 65536 | 5120 | [README.md](models/Yarn-Llama-2-13B-64K/README.md) |
| Yarn-Llama-2-70B-32k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 8192 | 8192 | [README.md](models/Yarn-Llama-2-70B-32k/README.md) |
| Yarn-Llama-2-7B-128K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 131072 | 4096 | [README.md](models/Yarn-Llama-2-7B-128K/README.md) |
| Yarn-Llama-2-7B-64K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 65536 | 4096 | [README.md](models/Yarn-Llama-2-7B-64K/README.md) |
| Yarn-Mistral-7B-128k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Yarn-Mistral-7B-128k/README.md) |
| Yarn-Mistral-7B-64k |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Yarn-Mistral-7B-64k/README.md) |
| Yi-34B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 7168 | [README.md](models/Yi-34B/README.md) |
| Yi-34B-200K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/Yi-34B-200K/README.md) |
| Yi-34B-200K-Llamafied |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/Yi-34B-200K-Llamafied/README.md) |
| Yi-34B-Chat |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 7168 | [README.md](models/Yi-34B-Chat/README.md) |
| Yi-34B-GiftedConvo-merged |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 7168 | [README.md](models/Yi-34B-GiftedConvo-merged/README.md) |
| yi-34B-v3 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 7168 | [README.md](models/yi-34B-v3/README.md) |
| Yi-6B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Yi-6B/README.md) |
| Yi-6B-200K |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 4096 | [README.md](models/Yi-6B-200K/README.md) |
| YuLan-Chat-2-13B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 5120 | [README.md](models/YuLan-Chat-2-13B/README.md) |
| Zarablend-L2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Zarablend-L2-7B/README.md) |
| Zarablend-MX-L2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Zarablend-MX-L2-7B/README.md) |
| Zarafusionex-1.1-L2-7B |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Zarafusionex-1.1-L2-7B/README.md) |
| ZephRP-m7b |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/ZephRP-m7b/README.md) |
| zephyr-7B-alpha |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/zephyr-7B-alpha/README.md) |
| zephyr-7B-beta |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/zephyr-7B-beta/README.md) |
| zephyr-7B-beta-pl |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/zephyr-7B-beta-pl/README.md) |
| Ziya-Coding-34B-v1.0 |Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/Ziya-Coding-34B-v1.0/README.md) |

## Memory/Disk Requirements
As the models are currently fully loaded into memory, you will need adequate disk space to save them and sufficient RAM to load them. At the moment, memory and disk requirements are the same.

| Model | Original size | Quantized size (4-bit) |
|------:|--------------:|-----------------------:|
|    7B |         13 GB |                 3.9 GB |
|   13B |         24 GB |                 7.8 GB |
|   30B |         60 GB |                19.5 GB |
|   65B |        120 GB |                38.5 GB |

## Quantization
Several quantization methods are supported. They differ in the resulting model disk size and inference speed.
*(outdated)*

| Model | Measure      | F16    | Q4_0   | Q4_1   | Q5_0   | Q5_1   | Q8_0   |
|------ |--------------|--------|--------|--------|--------|--------|--------|
|    7B | perplexity   | 5.9066 | 6.1565 | 6.0912 | 5.9862 | 5.9481 | 5.9070 |
|    7B | file size    |  13.0G |   3.5G |   3.9G |   4.3G |   4.7G |   6.7G |
|    7B | ms/tok @ 4th |    127 |     55 |     54 |     76 |     83 |     72 |
|    7B | ms/tok @ 8th |    122 |     43 |     45 |     52 |     56 |     67 |
|    7B | bits/weight  |   16.0 |    4.5 |    5.0 |    5.5 |    6.0 |    8.5 |
|   13B | perplexity   | 5.2543 | 5.3860 | 5.3608 | 5.2856 | 5.2706 | 5.2548 |
|   13B | file size    |  25.0G |   6.8G |   7.6G |   8.3G |   9.1G |    13G |
|   13B | ms/tok @ 4th |      - |    103 |    105 |    148 |    160 |    131 |
|   13B | ms/tok @ 8th |      - |     73 |     82 |     98 |    105 |    128 |
|   13B | bits/weight  |   16.0 |    4.5 |    5.0 |    5.5 |    6.0 |    8.5 |

## License

This project is licensed under the Mozilla Public License Version 2.0 License - see the [LICENSE](LICENSE) file for details.
