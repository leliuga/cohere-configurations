
LLM Inference configurations👋
===============================

## Description

This repository contains the inference configurations for the Co:Here

### Supported Transformer Architectures by Co:Here

- [X] LLaMA 🦙
- [x] LLaMA 2 🦙🦙
- [X] Falcon
- [X] Alpaca
- [X] GPT-2
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
- [X] Mistral AI v0.1 & v0.2
- [X] Refact
- [X] Persimmon 8B
- [X] MPT
- [X] Bloom
- [X] StableLM-3b-4e1t
- [x] Yi
- [x] Deepseek
- [x] Qwen
- [x] Mixtral MoE
- [x] Phi 2
- [x] PLaMo 13B

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
docker run -it --rm --gpus all -v ./models/:/app/models -p 3000:3000 github.com/leliuga/cohere cohere run <name>/<variant>
```

## Supported Models (871)

| Name  | DType   | Context | Embedding | Read more |
| ----- | ------- | ------- | --------- | --------- |
| 13B BlueMethod | Q6_K Q8_0 Q2_K Q3_K_S Q4_K_M Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q4_0 Q4_K_S | 2048 | 5120 | [README.md](models/13B-BlueMethod/README.md) |
| 13B Chimera | Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_S Q8_0 Q2_K Q3_K_L Q4_0 Q5_0 Q5_K_M Q6_K | 2048 | 5120 | [README.md](models/13B-Chimera/README.md) |
| 13B Hypermantis | Q8_0 Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q2_K Q3_K_L Q5_0 Q5_K_S Q6_K | 2048 | 5120 | [README.md](models/13B-HyperMantis/README.md) |
| 13B Legerdemain L2 | Q3_K_M Q4_0 Q4_K_S Q5_0 Q5_K_S Q2_K Q3_K_L Q5_K_M Q6_K Q8_0 Q3_K_S Q4_K_M | 4096 | 5120 | [README.md](models/13B-Legerdemain-L2/README.md) |
| 13B Ouroboros | Q2_K Q3_K_L Q3_K_S Q4_0 Q5_K_M Q5_K_S Q3_K_M Q4_K_M Q4_K_S Q5_0 Q6_K Q8_0 | 2048 | 5120 | [README.md](models/13B-Ouroboros/README.md) |
| 13B Thorns L2 | Q2_K Q3_K_L Q4_0 Q4_K_S Q5_K_M Q3_K_M Q3_K_S Q4_K_M Q5_0 Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/13B-Thorns-L2/README.md) |
| 30B Epsilon | Q2_K Q3_K_M Q4_0 Q4_K_S Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_L Q3_K_S Q4_K_M Q5_0 | 2048 | 6656 | [README.md](models/30B-Epsilon/README.md) |
| 30B Lazarus | Q3_K_M Q3_K_S Q5_K_S Q8_0 Q5_K_M Q6_K Q2_K Q3_K_L Q4_0 Q4_K_M Q4_K_S Q5_0 | 2048 | 6656 | [README.md](models/30B-Lazarus/README.md) |
| CausalLM 8X7B MoE Test NOT MIXTRAL | Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q2_K | 32768 | 4096 | [README.md](models/8x7B-MoE-test-NOT-MIXTRAL/README.md) |
| ALMA 13B | Q3_K_M Q3_K_S Q4_K_M Q5_0 Q5_K_S Q2_K Q3_K_L Q5_K_M Q6_K Q8_0 Q4_0 Q4_K_S | 4096 | 5120 | [README.md](models/ALMA-13B/README.md) |
| ALMA 13B Pretrain | Q4_K_S Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_S Q4_0 Q5_K_S Q3_K_L Q3_K_M Q4_K_M | 4096 | 5120 | [README.md](models/ALMA-13B-Pretrain/README.md) |
| ALMA 7B | Q2_K Q3_K_L Q4_K_M Q5_0 Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/ALMA-7B/README.md) |
| ALMA 7B Pretrain | Q8_0 Q3_K_L Q3_K_S Q4_K_M Q5_0 Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_M Q4_0 Q4_K_S | 4096 | 4096 | [README.md](models/ALMA-7B-Pretrain/README.md) |
| Anima Phi Neptune Mistral 7B | Q3_K_M Q5_K_M Q5_K_S Q8_0 Q5_0 Q6_K Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S | 32768 | 4096 | [README.md](models/ANIMA-Phi-Neptune-Mistral-7B/README.md) |
| ARIA 70B V2 | Q3_K_S Q4_K_S Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_M Q4_0 Q4_K_M Q5_K_S | 4096 | 8192 | [README.md](models/ARIA-70B-V2/README.md) |
| Aetheria L2 70B | Q4_K_S Q5_0 Q5_K_M Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_K_S Q3_K_L | 4096 | 8192 | [README.md](models/Aetheria-L2-70B/README.md) |
| Airoboros 33B 2.1 | Q4_K_S Q5_0 Q5_K_S Q2_K Q3_K_L Q3_K_M Q4_0 Q4_K_M Q6_K Q3_K_S Q5_K_M Q8_0 | 4096 | 6656 | [README.md](models/Airoboros-33B-2.1/README.md) |
| Airoboros L2 13B 2.1 | Q5_K_M Q6_K Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q3_K_M Q5_0 Q5_K_S Q8_0 | 4096 | 5120 | [README.md](models/Airoboros-L2-13B-2.1/README.md) |
| Airoboros L2 13B 2.2 | Q3_K_L Q5_K_M Q5_K_S Q2_K Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q6_K Q8_0 Q3_K_M | 4096 | 5120 | [README.md](models/Airoboros-L2-13B-2.2/README.md) |
| Airoboros L2 13B 2.1 YaRN 64K | Q4_K_M Q4_K_S Q2_K Q4_0 Q3_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_L Q3_K_M | 65536 | 5120 | [README.md](models/Airoboros-L2-13B-2_1-YaRN-64K/README.md) |
| Airoboros L2 13B 3.1.1 | Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_L Q5_0 | 4096 | 5120 | [README.md](models/Airoboros-L2-13B-3.1.1/README.md) |
| Airoboros L2 70B 2.1 | Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q2_K Q4_0 Q5_0 Q5_K_S | 4096 | 8192 | [README.md](models/Airoboros-L2-70B-2.1/README.md) |
| Airoboros L2 70B 2.1 Creative | Q5_K_S Q2_K Q4_K_M Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q3_K_L Q3_K_M | 4096 | 8192 | [README.md](models/Airoboros-L2-70B-2.1-Creative/README.md) |
| Airoboros L2 70B 3.1.2 | Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_S Q3_K_L Q4_K_M Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Airoboros-L2-70B-3.1.2/README.md) |
| Airoboros L2 70B GPT4 m2.0 | Q4_K_S Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q4_0 Q4_K_M Q2_K Q3_K_S | 4096 | 8192 | [README.md](models/Airoboros-L2-70B-GPT4-m2.0/README.md) |
| Airoboros L2 70B 2.2 | Q4_K_M Q5_0 Q5_K_M Q3_K_M Q3_K_L Q3_K_S Q4_0 Q4_K_S Q5_K_S Q2_K | 4096 | 8192 | [README.md](models/Airoboros-L2-70b-2.2/README.md) |
| Airoboros L2 70B 2.2.1 | Q2_K Q3_K_L Q3_K_S Q4_K_M Q5_0 Q3_K_M Q4_K_S Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Airoboros-L2-70b-2.2.1/README.md) |
| Airoboros L2 7B 2.1 | Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q6_K Q8_0 Q3_K_L Q3_K_M Q4_0 Q5_0 Q5_K_S Q2_K | 4096 | 4096 | [README.md](models/Airoboros-L2-7B-2.1/README.md) |
| Airoboros L2 7B 2.2 | Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q3_K_L Q8_0 Q3_K_M Q4_0 Q2_K | 4096 | 4096 | [README.md](models/Airoboros-L2-7B-2.2/README.md) |
| Airoboros M 7B 3.1.2 | Q4_0 Q4_K_S Q5_K_M Q8_0 Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q5_0 | 32768 | 4096 | [README.md](models/Airoboros-M-7B-3.1.2/README.md) |
| Airoboros c34B 2.1 | Q5_0 Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q8_0 Q4_K_M Q4_K_S Q5_K_M | 16384 | 8192 | [README.md](models/Airoboros-c34B-2.1/README.md) |
| Airoboros c34B 2.2 | Q3_K_M Q4_K_M Q4_K_S Q5_K_S Q6_K Q5_K_M Q8_0 Q2_K Q3_K_L Q3_K_S Q4_0 Q5_0 | 16384 | 8192 | [README.md](models/Airoboros-c34B-2.2/README.md) |
| Airoboros c34B 3.1.2 | Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q4_0 Q4_K_S Q5_K_S | 16384 | 8192 | [README.md](models/Airoboros-c34B-3.1.2/README.md) |
| Airochronos L2 13B | Q4_K_S Q5_0 Q8_0 Q3_K_M Q3_K_S Q4_K_M Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_L Q4_0 | 4096 | 5120 | [README.md](models/Airochronos-L2-13B/README.md) |
| Airolima Chronos Grad L2 13B | Q2_K Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q3_K_M Q4_0 Q5_0 Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Airolima-Chronos-Grad-L2-13B/README.md) |
| Akins 3B | Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q6_K Q8_0 Q2_K Q3_K_S Q4_0 Q5_0 Q5_K_M Q5_K_S | 4096 | 2560 | [README.md](models/Akins-3B/README.md) |
| AlpacaCielo2 7B 8K | Q2_K Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q3_K_M Q4_0 Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/AlpacaCielo2-7B-8K/README.md) |
| Amber | Q3_K_S Q4_K_S Q5_0 Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q5_K_M Q5_K_S Q6_K Q3_K_L | 2048 | 4096 | [README.md](models/Amber/README.md) |
| AmberChat | Q3_K_S Q4_0 Q6_K Q2_K Q3_K_L Q3_K_M Q5_K_M Q5_K_S Q8_0 Q4_K_M Q4_K_S Q5_0 | 2048 | 4096 | [README.md](models/AmberChat/README.md) |
| Amethyst 13B | Q5_0 Q5_K_S Q6_K Q2_K Q3_K_L Q4_K_S Q4_K_M Q5_K_M Q8_0 Q3_K_M Q3_K_S Q4_0 | 4096 | 5120 | [README.md](models/Amethyst-13B/README.md) |
| Ana v1 m7 | Q3_K_S Q4_0 Q4_K_S Q5_K_S Q6_K Q8_0 Q3_K_L Q3_K_M Q4_K_M Q5_0 Q5_K_M Q2_K | 32768 | 4096 | [README.md](models/Ana-v1-m7/README.md) |
| AppleSauce L2 13B | Q3_K_M Q5_0 Q6_K Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q8_0 | 4096 | 5120 | [README.md](models/AppleSauce-L2-13B/README.md) |
| AquilaChat2 34B | Q4_K_M Q5_K_S Q6_K Q3_K_S Q4_0 Q3_K_M Q4_K_S Q5_0 Q5_K_M Q8_0 Q2_K Q3_K_L | 4096 | 6144 | [README.md](models/AquilaChat2-34B/README.md) |
| Aquilachat2 34B 16K | Q4_K_S Q5_K_S Q8_0 Q2_K Q3_K_L Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q3_K_M Q3_K_S | 4096 | 6144 | [README.md](models/AquilaChat2-34B-16K/README.md) |
| Arithmo Mistral 7B | Q5_0 Q6_K Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_S | 32768 | 4096 | [README.md](models/Arithmo-Mistral-7B/README.md) |
| Asclepius 13B | Q6_K Q8_0 Q3_K_S Q4_K_S Q5_0 Q5_K_M Q4_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M Q4_0 | 2048 | 5120 | [README.md](models/Asclepius-13B/README.md) |
| AshhLimaRP Mistral 7B | Q5_K_M Q6_K Q2_K Q3_K_S Q4_0 Q4_K_M Q4_K_S Q3_K_L Q3_K_M Q5_0 Q5_K_S Q8_0 | 32768 | 4096 | [README.md](models/AshhLimaRP-Mistral-7B/README.md) |
| Astrid Mistral 7B | Q3_K_M Q3_K_S Q4_0 Q5_K_S Q8_0 Q2_K Q4_K_M Q4_K_S Q5_0 Q5_K_M Q6_K Q3_K_L | 32768 | 4096 | [README.md](models/Astrid-Mistral-7B/README.md) |
| Athena v1 | Q3_K_M Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_S Q4_K_M Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Athena-v1/README.md) |
| Athena V2 | Q5_K_S Q6_K Q3_K_L Q3_K_M Q4_K_M Q5_0 Q5_K_M Q8_0 Q2_K Q3_K_S Q4_0 Q4_K_S | 4096 | 5120 | [README.md](models/Athena-v2/README.md) |
| Athena V3 | Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q3_K_L Q8_0 Q3_K_M Q6_K Q2_K | 4096 | 5120 | [README.md](models/Athena-v3/README.md) |
| Athena v4 | Q3_K_S Q4_0 Q4_K_M Q5_K_M Q5_K_S Q6_K Q3_K_L Q3_K_M Q8_0 Q5_0 Q2_K Q4_K_S | 4096 | 5120 | [README.md](models/Athena-v4/README.md) |
| Athnete 13B | Q4_K_S Q5_0 Q5_K_M Q6_K Q3_K_S Q3_K_L Q3_K_M Q4_0 Q4_K_M Q5_K_S Q8_0 Q2_K | 4096 | 5120 | [README.md](models/Athnete-13B/README.md) |
| Augmental 13B | Q2_K Q3_K_M Q4_0 Q5_0 Q6_K Q8_0 Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q5_K_S | 4096 | 5120 | [README.md](models/Augmental-13B/README.md) |
| Augmental 13B v1.50A | Q5_K_M Q2_K Q3_K_L Q3_K_M Q4_0 Q4_K_S Q8_0 Q3_K_S Q4_K_M Q5_0 Q5_K_S Q6_K | 4096 | 5120 | [README.md](models/Augmental-13B-v1.50_A/README.md) |
| Augmental 13B v1.50B | Q6_K Q2_K Q3_K_L Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q3_K_M Q3_K_S Q4_K_S Q8_0 | 4096 | 5120 | [README.md](models/Augmental-13B-v1.50_B/README.md) |
| Augmental ReMM 13B | Q3_K_M Q3_K_S Q5_0 Q5_K_S Q2_K Q3_K_L Q4_0 Q4_K_M Q4_K_S Q5_K_M Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Augmental-ReMM-13B/README.md) |
| Augmental Unholy 13B | Q5_K_S Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_K_M Q3_K_M Q4_K_S Q5_0 Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Augmental-Unholy-13B/README.md) |
| Aurora Nights 103B v1.0 | Q2_K Q3_K_M Q3_K_S | 4096 | 8192 | [README.md](models/Aurora-Nights-103B-v1.0/README.md) |
| Aurora Nights 70B v1.0 | Q3_K_S Q5_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_M | 4096 | 8192 | [README.md](models/Aurora-Nights-70B-v1.0/README.md) |
| Autolycus Mistral 7B | Q2_K Q3_K_M Q3_K_S Q5_0 Q5_K_M Q5_K_S Q8_0 Q3_K_L Q4_0 Q4_K_M Q4_K_S Q6_K | 32768 | 4096 | [README.md](models/Autolycus-Mistral_7B/README.md) |
| Barcenas Mistral 7B | Q4_K_S Q5_K_M Q2_K Q3_K_L Q3_K_S Q5_0 Q5_K_S Q6_K Q8_0 Q3_K_M Q4_0 Q4_K_M | 32768 | 4096 | [README.md](models/Barcenas-Mistral-7B/README.md) |
| BerrySauce L2 13B | Q2_K Q5_K_M Q5_K_S Q4_K_S Q5_0 Q6_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q8_0 | 4096 | 5120 | [README.md](models/BerrySauce-L2-13B/README.md) |
| Bigplap 8X20B | Q2_K | 32768 | 5120 | [README.md](models/BigPlap-8x20B/README.md) |
| BruinsV2 OpHermesNeu 11B | Q2_K Q3_K_L Q4_0 Q4_K_S Q5_0 Q5_K_S Q3_K_M Q3_K_S Q4_K_M Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/BruinsV2-OpHermesNeu-11B/README.md) |
| CAMEL 13B Combined Data | Q6_K Q3_K_S Q4_K_S Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q8_0 | 2048 | 5120 | [README.md](models/CAMEL-13B-Combined-Data/README.md) |
| CAMEL 13B Role Playing Data | Q2_K Q3_K_M Q4_K_M Q4_K_S Q5_0 Q8_0 Q3_K_L Q3_K_S Q4_0 Q5_K_M Q5_K_S Q6_K | 2048 | 5120 | [README.md](models/CAMEL-13B-Role-Playing-Data/README.md) |
| CAMEL 33B Combined Data | Q2_K Q3_K_S Q4_0 Q5_K_M Q6_K Q8_0 Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q5_0 Q5_K_S | 2048 | 6656 | [README.md](models/CAMEL-33B-Combined-Data/README.md) |
| CaPlatTessDolXaBoros Yi 34B 200K Dare Ties Highdensity | Q2_K Q3_K_S Q4_K_M Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q4_0 Q4_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/CaPlatTessDolXaBoros-Yi-34B-200K-DARE-Ties-HighDensity/README.md) |
| CalliopeDS L2 13B | Q3_K_M Q4_K_M Q5_0 Q5_K_M Q5_K_S Q8_0 Q3_K_L Q3_K_S Q4_0 Q4_K_S Q6_K Q2_K | 4096 | 5120 | [README.md](models/CalliopeDS-L2-13B/README.md) |
| Camel-Platypus2 13B | Q2_K Q4_K_M Q4_K_S Q5_K_S Q5_0 Q5_K_M Q6_K Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_0 | 4096 | 5120 | [README.md](models/Camel-Platypus2-13B/README.md) |
| Camel Platypus2 70B | Q3_K_L Q3_K_M Q4_K_M Q5_0 Q5_K_S Q2_K Q4_0 Q4_K_S Q5_K_M Q3_K_S | 4096 | 8192 | [README.md](models/Camel-Platypus2-70B/README.md) |
| Capytessborosyi 34B 200K Dare Ties | Q5_0 Q5_K_M Q5_K_S Q3_K_M Q3_K_S Q4_K_S Q4_K_M Q6_K Q8_0 Q2_K Q3_K_L Q4_0 | 200000 | 7168 | [README.md](models/CapyTessBorosYi-34B-200K-DARE-Ties/README.md) |
| Capybara Tess Yi 34B 200K | Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q6_K Q8_0 Q2_K Q5_0 | 200000 | 7168 | [README.md](models/Capybara-Tess-Yi-34B-200K/README.md) |
| Capybara Tess Yi 34B 200K Dare Ties | Q8_0 Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_S Q2_K Q3_K_L Q5_0 Q5_K_M Q6_K | 200000 | 7168 | [README.md](models/Capybara-Tess-Yi-34B-200K-DARE-Ties/README.md) |
| Carl Llama 2 | Q2_K Q3_K_L Q4_0 Q4_K_S Q5_K_M Q8_0 Q3_K_M Q3_K_S Q4_K_M Q5_0 Q5_K_S Q6_K | 4096 | 5120 | [README.md](models/Carl-Llama-2-13B/README.md) |
| Cat 13B 0.5 | Q3_K_L Q4_K_M Q5_K_M Q5_K_S Q5_0 Q6_K Q8_0 Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_S | 4096 | 5120 | [README.md](models/Cat-13B-0.5/README.md) |
| CausalLM 14B | Q4_1 Q5_0 Q5_1 Q8_0 Q4_0 | 8192 | 5120 | [README.md](models/CausalLM-14B/README.md) |
| CausalLM 7B | Q6_K Q8_0 Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_M Q4_K_M Q5_K_S | 8192 | 4096 | [README.md](models/CausalLM-7B/README.md) |
| ChatAYT Lora Assamble Marcoroni | Q4_K_M Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_S Q3_K_L Q3_K_M Q4_0 Q4_K_S Q2_K | 4096 | 5120 | [README.md](models/ChatAYT-Lora-Assamble-Marcoroni/README.md) |
| Chinese Alpaca 2 13B | Q5_K_M Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q4_K_M Q4_K_S Q5_0 Q3_K_M Q3_K_S Q4_0 | 4096 | 5120 | [README.md](models/Chinese-Alpaca-2-13B/README.md) |
| Chinese Alpaca 2 7B | Q6_K Q2_K Q3_K_L Q4_K_M Q4_K_S Q5_0 Q5_K_S Q3_K_M Q3_K_S Q4_0 Q5_K_M Q8_0 | 4096 | 4096 | [README.md](models/Chinese-Alpaca-2-7B/README.md) |
| Chinese Llama 2 13B | Q6_K Q8_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q5_0 Q4_0 Q4_K_S Q5_K_M Q5_K_S | 4096 | 5120 | [README.md](models/Chinese-Llama-2-13B/README.md) |
| Chinese Llama 2 7B | Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q3_K_M Q3_K_L Q3_K_S Q4_K_S Q6_K Q8_0 Q2_K | 4096 | 4096 | [README.md](models/Chinese-Llama-2-7B/README.md) |
| Chronoboros 33B | Q6_K Q8_0 Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_K_S Q2_K Q3_K_L Q4_K_S Q5_0 Q5_K_M | 2048 | 6656 | [README.md](models/Chronoboros-33B/README.md) |
| Chronoboros Grad L2 13B | Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_0 Q6_K Q8_0 Q3_K_L Q4_0 Q5_K_M Q5_K_S Q2_K | 4096 | 5120 | [README.md](models/Chronoboros-Grad-L2-13B/README.md) |
| Chronohermes Grad L2 13B | Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q6_K Q8_0 Q3_K_M Q4_0 | 4096 | 5120 | [README.md](models/Chronohermes-Grad-L2-13B/README.md) |
| Chronolima Airo Grad L2 13B | Q4_0 Q4_K_S Q5_0 Q5_K_M Q6_K Q2_K Q3_K_L Q3_K_S Q8_0 Q3_K_M Q4_K_M Q5_K_S | 4096 | 5120 | [README.md](models/Chronolima-Airo-Grad-L2-13B/README.md) |
| Chronomaid Storytelling 13B | Q2_K Q3_K_M Q4_0 Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Chronomaid-Storytelling-13B/README.md) |
| Chronorctypus-Limarobormes-13B | Q2_K Q3_K_S Q4_K_M Q5_0 Q5_K_S Q3_K_L Q3_K_M Q4_0 Q4_K_S Q5_K_M Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Chronorctypus-Limarobormes-13b/README.md) |
| Chronos 70B v2 | Q3_K_L Q4_0 Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_M Q3_K_S Q4_K_M Q4_K_S | 4096 | 8192 | [README.md](models/Chronos-70B-v2/README.md) |
| Chronos Beluga v2 13B | Q4_K_S Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Chronos-Beluga-v2-13B/README.md) |
| Chronos Hermes 13B v2 | Q3_K_M Q3_K_S Q4_0 Q5_K_S Q8_0 Q2_K Q3_K_L Q4_K_M Q4_K_S Q5_0 Q5_K_M Q6_K | 4096 | 5120 | [README.md](models/Chronos-Hermes-13b-v2/README.md) |
| Chupacabra 7B V2 | Q3_K_M Q3_K_S Q5_K_M Q5_K_S Q6_K Q8_0 Q2_K Q4_0 Q4_K_M Q4_K_S Q5_0 Q3_K_L | 32768 | 4096 | [README.md](models/Chupacabra-7B-v2/README.md) |
| Chupacabra 8X7B MoE | Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K | 32768 | 4096 | [README.md](models/Chupacabra-8x7B-MoE/README.md) |
| Claire 7B 0.1 | Q5_K_M Q6_K Q8_0 Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_0 Q2_K Q3_K_L Q4_0 Q5_K_S | 0 | 4544 | [README.md](models/Claire-7B-0.1/README.md) |
| Clover3 17B | Q3_K_L Q3_K_S Q4_0 Q5_0 Q5_K_M Q2_K Q3_K_M Q4_K_M Q4_K_S Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Clover3-17B/README.md) |
| Code 13B | Q5_K_M Q8_0 Q3_K_M Q3_K_S Q4_K_M Q5_0 Q5_K_S Q6_K Q2_K Q3_K_L Q4_0 Q4_K_S | 4096 | 5120 | [README.md](models/Code-13B/README.md) |
| Code 33B | Q3_K_L Q4_K_M Q4_K_S Q8_0 Q2_K Q3_K_S Q4_0 Q5_0 Q5_K_M Q5_K_S Q6_K Q3_K_M | 2048 | 6656 | [README.md](models/Code-33B/README.md) |
| CodeBooga 34B v0.1 | Q2_K Q3_K_M Q3_K_S Q4_0 Q5_K_S Q6_K Q3_K_L Q4_K_M Q4_K_S Q5_0 Q5_K_M Q8_0 | 16384 | 8192 | [README.md](models/CodeBooga-34B-v0.1/README.md) |
| CodeFuse CodeLlama 34B | Q3_K_S Q4_K_S Q5_0 Q6_K Q3_K_M Q3_K_L Q4_0 Q4_K_M Q5_K_M Q5_K_S Q8_0 Q2_K | 16384 | 8192 | [README.md](models/CodeFuse-CodeLlama-34B/README.md) |
| CodeLlama 13B | Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_K_M Q6_K Q3_K_L Q4_K_S Q5_0 Q5_K_S Q8_0 Q2_K | 16384 | 5120 | [README.md](models/CodeLlama-13B/README.md) |
| CodeLlama 13B Instruct | Q2_K Q4_0 Q4_K_S Q5_K_M Q6_K Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q5_0 Q5_K_S | 16384 | 5120 | [README.md](models/CodeLlama-13B-Instruct/README.md) |
| CodeLlama 13B Python | Q2_K Q4_K_M Q4_K_S Q5_0 Q6_K Q5_K_S Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_K_M | 16384 | 5120 | [README.md](models/CodeLlama-13B-Python/README.md) |
| CodeLlama 13B SFT v10 | Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q2_K Q3_K_M Q4_0 Q5_0 Q6_K Q8_0 | 16384 | 5120 | [README.md](models/CodeLlama-13B-oasst-sft-v10/README.md) |
| CodeLlama 34B | Q5_K_S Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_M Q8_0 Q2_K Q3_K_L Q3_K_M Q4_0 Q6_K | 16384 | 8192 | [README.md](models/CodeLlama-34B/README.md) |
| CodeLlama 34B Instruct | Q5_K_S Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q4_K_S Q3_K_L Q3_K_S Q5_0 Q5_K_M Q6_K | 16384 | 8192 | [README.md](models/CodeLlama-34B-Instruct/README.md) |
| CodeLlama 34B Python | Q5_0 Q5_K_M Q6_K Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q8_0 Q2_K Q3_K_S Q5_K_S | 16384 | 8192 | [README.md](models/CodeLlama-34B-Python/README.md) |
| CodeLlama 7B | Q3_K_S Q4_0 Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q5_0 Q6_K | 16384 | 4096 | [README.md](models/CodeLlama-7B/README.md) |
| CodeLlama 7B Instruct | Q4_K_S Q5_K_M Q5_K_S Q6_K Q2_K Q4_0 Q3_K_S Q4_K_M Q5_0 Q8_0 Q3_K_L Q3_K_M | 16384 | 4096 | [README.md](models/CodeLlama-7B-Instruct/README.md) |
| CodeLlama 7B Python | Q3_K_L Q3_K_M Q3_K_S Q6_K Q8_0 Q5_K_S Q2_K Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M | 16384 | 4096 | [README.md](models/CodeLlama-7B-Python/README.md) |
| CodeNinja 1.0 Openchat 7B | Q3_K_L Q3_K_S Q4_0 Q5_0 Q8_0 Q6_K Q2_K Q3_K_M Q4_K_M Q4_K_S Q5_K_M Q5_K_S | 8192 | 4096 | [README.md](models/CodeNinja-1.0-OpenChat-7B/README.md) |
| CodeUp Alpha 13B HF | Q5_0 Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_M Q4_0 Q8_0 Q3_K_S Q4_K_M Q4_K_S Q5_K_M | 4096 | 5120 | [README.md](models/CodeUp-Alpha-13B-HF/README.md) |
| CodeUp Llama 2 13B Chat HF | Q3_K_S Q4_K_M Q8_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M Q4_0 Q4_K_S Q5_0 Q6_K | 4096 | 5120 | [README.md](models/CodeUp-Llama-2-13B-Chat-HF/README.md) |
| CollectiveCognition v1 Mistral 7B | Q5_0 Q8_0 Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q5_K_S Q6_K Q2_K Q3_K_S Q4_0 Q5_K_M | 32768 | 4096 | [README.md](models/CollectiveCognition-v1-Mistral-7B/README.md) |
| CollectiveCognition v1.1 Mistral 7B | Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_L Q3_K_M Q4_K_S Q5_K_S | 32768 | 4096 | [README.md](models/CollectiveCognition-v1.1-Mistral-7B/README.md) |
| Collectivecognition V1.1 Mistral 7B Dare 0.85 | Q3_K_M Q3_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q4_0 Q4_K_M Q4_K_S Q6_K Q8_0 Q3_K_L | 32768 | 4096 | [README.md](models/CollectiveCognition-v1.1-Mistral-7B-dare-0.85/README.md) |
| DPOpenHermes 7B | Q3_K_M Q4_0 Q4_K_S Q5_K_S Q2_K Q3_K_S Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q3_K_L | 32768 | 4096 | [README.md](models/DPOpenHermes-7B/README.md) |
| DPOpenHermes 7B v2 | Q4_0 Q4_K_M Q5_0 Q5_K_S Q6_K Q3_K_M Q3_K_S Q4_K_S Q5_K_M Q8_0 Q2_K Q3_K_L | 32768 | 4096 | [README.md](models/DPOpenHermes-7B-v2/README.md) |
| Dans AdventurousWinds 7B | Q4_K_S Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q6_K Q8_0 Q4_K_M Q5_K_S | 32768 | 4096 | [README.md](models/Dans-AdventurousWinds-7B/README.md) |
| Dans AdventurousWinds Mk2 7B | Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_L Q4_0 Q6_K Q3_K_M Q3_K_S Q5_0 | 32768 | 4096 | [README.md](models/Dans-AdventurousWinds-Mk2-7B/README.md) |
| Dans TotSirocco 7B | Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q6_K Q8_0 Q2_K Q3_K_M Q5_0 Q5_K_S | 32768 | 4096 | [README.md](models/Dans-TotSirocco-7B/README.md) |
| DaringMaid 13B | Q2_K Q3_K_L Q3_K_S Q4_K_M Q5_0 Q5_K_M Q5_K_S Q3_K_M Q4_0 Q4_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/DaringMaid-13B/README.md) |
| DaringMaid 20B | Q5_K_M Q5_K_S Q8_0 Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q2_K Q3_K_L Q3_K_M Q6_K | 4096 | 5120 | [README.md](models/DaringMaid-20B/README.md) |
| Dawn V2 70B | Q2_K Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q5_0 | 4096 | 8192 | [README.md](models/Dawn-v2-70B/README.md) |
| Deacon 34B | Q2_K Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_0 Q3_K_M Q4_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 7168 | [README.md](models/Deacon-34B/README.md) |
| DiscoLM 120B | Q2_K | 4096 | 8192 | [README.md](models/DiscoLM-120b/README.md) |
| DiscoLM 70B | Q4_0 Q4_K_M Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_S Q5_K_S | 8192 | 8192 | [README.md](models/DiscoLM-70B/README.md) |
| Dolphin 2.1 70B | Q5_K_M Q2_K Q3_K_L Q3_K_M Q5_0 Q5_K_S Q3_K_S Q4_0 Q4_K_M Q4_K_S | 4096 | 8192 | [README.md](models/Dolphin-2.1-70B/README.md) |
| Dolphin Llama 13B | Q8_0 Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_S Q5_K_M | 2048 | 5120 | [README.md](models/Dolphin-Llama-13B/README.md) |
| Dolphin Llama2 7B | Q5_0 Q6_K Q2_K Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q8_0 Q3_K_L Q3_K_S | 4096 | 4096 | [README.md](models/Dolphin-Llama2-7B/README.md) |
| Dolphin2.1 OpenOrca 7B | Q6_K Q2_K Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q3_K_S Q4_K_S Q8_0 | 32768 | 4096 | [README.md](models/Dolphin2.1-OpenOrca-7B/README.md) |
| Echidna 13B v0.2 | Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q8_0 Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K | 4096 | 5120 | [README.md](models/Echidna-13B-v0.2/README.md) |
| Echidna 13B v0.3 | Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q8_0 Q2_K Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K | 4096 | 5120 | [README.md](models/Echidna-13B-v0.3/README.md) |
| Emerhyst 13B | Q3_K_L Q3_K_S Q4_0 Q5_0 Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_M Q4_K_M Q4_K_S Q8_0 | 4096 | 5120 | [README.md](models/Emerhyst-13B/README.md) |
| Emerhyst 20B | Q2_K Q4_0 Q4_K_S Q5_0 Q5_K_S Q6_K Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q5_K_M | 4096 | 5120 | [README.md](models/Emerhyst-20B/README.md) |
| Euryale 1.3 L2 70B | Q2_K Q3_K_M Q4_K_M Q5_K_M Q3_K_L Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_S | 4096 | 8192 | [README.md](models/Euryale-1.3-L2-70B/README.md) |
| Euryale 1.4 L2 70B | Q4_K_S Q5_K_M Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_S | 4096 | 8192 | [README.md](models/Euryale-1.4-L2-70B/README.md) |
| Euryale Inverted L2 70B | Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_0 Q2_K Q4_K_M Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Euryale-Inverted-L2-70B/README.md) |
| Euryale L2 70B | Q3_K_S Q4_0 Q4_K_S Q5_0 Q2_K Q3_K_M Q5_K_M Q5_K_S Q3_K_L Q4_K_M | 4096 | 8192 | [README.md](models/Euryale-L2-70B/README.md) |
| EverythingLM 13B 16K | Q2_K Q4_K_M Q5_0 Q8_0 Q5_K_M Q5_K_S Q6_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_S | 16384 | 5120 | [README.md](models/EverythingLM-13B-16K/README.md) |
| EverythingLM 13B V3 16K | Q2_K Q4_0 Q5_0 Q6_K Q4_K_S Q5_K_M Q5_K_S Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_K_M | 16384 | 5120 | [README.md](models/EverythingLM-13B-V3-16K/README.md) |
| EverythingLM 13B V2 16K | Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_S Q2_K Q3_K_L Q5_K_M Q6_K Q8_0 Q3_K_S Q4_K_S | 16384 | 5120 | [README.md](models/EverythingLM-13b-V2-16K/README.md) |
| Falkor 8X7B MoE | Q6_K Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M | 32768 | 4096 | [README.md](models/Falkor-8x7B-MoE/README.md) |
| FashionGPT 70B V1.1 | Q4_K_M Q4_K_S Q5_K_M Q3_K_L Q3_K_M Q4_0 Q5_0 Q5_K_S Q2_K Q3_K_S | 4096 | 8192 | [README.md](models/FashionGPT-70B-V1.1/README.md) |
| Fashiongpt 70B v1.2 | Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M Q4_K_M | 4096 | 8192 | [README.md](models/FashionGPT-70B-v1.2/README.md) |
| Fennec Mixtral 8X7B | Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K | 32768 | 4096 | [README.md](models/Fennec-Mixtral-8x7B/README.md) |
| Ferret 7B | Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q5_0 Q6_K Q8_0 Q3_K_S Q4_0 | 32768 | 4096 | [README.md](models/Ferret_7B/README.md) |
| FlatOrcaMaid 13B V0.2 | Q8_0 Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q6_K Q2_K Q3_K_L Q3_K_M Q5_K_S | 4096 | 5120 | [README.md](models/FlatOrcamaid-13B-v0.2/README.md) |
| Free Sydney V2 13B | Q3_K_M Q3_K_S Q4_0 Q6_K Q3_K_L Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q8_0 Q2_K | 4096 | 5120 | [README.md](models/Free_Sydney_V2_13B/README.md) |
| Free Sydney V2 Mistral 7B | Q5_0 Q5_K_M Q6_K Q2_K Q3_K_M Q4_0 Q4_K_S Q8_0 Q3_K_L Q3_K_S Q4_K_M Q5_K_S | 32768 | 4096 | [README.md](models/Free_Sydney_V2_Mistral_7b/README.md) |
| Frostwind 10.7B V1 | Q6_K Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q3_K_L Q4_K_M Q5_K_S Q8_0 | 4096 | 4096 | [README.md](models/Frostwind-10.7B-v1/README.md) |
| Geitje 7B Chat | Q2_K Q5_0 Q5_K_S Q4_K_S Q5_K_M Q6_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q8_0 | 32768 | 4096 | [README.md](models/GEITje-7B-chat/README.md) |
| Goat 70B Storytelling | Q2_K Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_0 | 4096 | 8192 | [README.md](models/GOAT-70B-Storytelling/README.md) |
| Gplatty 30B | Q5_K_S Q6_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_S Q2_K Q4_K_M Q5_0 Q5_K_M Q8_0 | 2048 | 6656 | [README.md](models/GPlatty-30B/README.md) |
| Generate Question Mistral 7B | Q3_K_L Q3_K_M Q5_0 Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_S Q4_0 Q4_K_M Q4_K_S Q8_0 | 32768 | 4096 | [README.md](models/Generate_Question_Mistral_7B/README.md) |
| GenZ 70B | Q5_K_M Q2_K Q3_K_S Q4_0 Q4_K_S Q5_0 Q3_K_L Q3_K_M Q4_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Genz-70b/README.md) |
| GodziLLa2 70B | Q3_K_L Q3_K_M Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_S Q4_K_M | 4096 | 8192 | [README.md](models/GodziLLa2-70B/README.md) |
| Guanaco 13B Uncensored | Q4_K_M Q8_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q5_K_S Q6_K Q4_0 Q4_K_S Q5_0 Q5_K_M | 4096 | 5120 | [README.md](models/Guanaco-13B-Uncensored/README.md) |
| Guanaco 7B Uncensored | Q3_K_S Q4_K_M Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_M Q4_0 Q4_K_S Q5_0 Q8_0 | 4096 | 4096 | [README.md](models/Guanaco-7B-Uncensored/README.md) |
| Hermes Trismegistus Mistral 7B | Q2_K Q3_K_M Q4_K_S Q5_0 Q5_K_M Q8_0 Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_K_S Q6_K | 32768 | 4096 | [README.md](models/Hermes-Trismegistus-Mistral-7B/README.md) |
| Hermes Lima RP L2 7B | Q5_K_S Q4_K_M Q4_K_S Q5_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_K_M Q6_K Q8_0 | 4096 | 4096 | [README.md](models/HermesLimaRP-L2-7B/README.md) |
| Hexoteric 7B | Q5_0 Q5_K_S Q6_K Q2_K Q3_K_L Q4_0 Q4_K_M Q4_K_S Q3_K_M Q3_K_S Q5_K_M Q8_0 | 32768 | 4096 | [README.md](models/Hexoteric-7B/README.md) |
| HornyEchidna 13B v0.1 | Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_K_M Q2_K Q4_K_S Q5_0 Q5_K_S Q6_K | 4096 | 5120 | [README.md](models/HornyEchidna-13B-v0.1/README.md) |
| Huginn 13B | Q2_K Q3_K_S Q4_0 Q5_K_S Q6_K Q8_0 Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q5_0 Q5_K_M | 4096 | 5120 | [README.md](models/Huginn-13B/README.md) |
| Huginn 13B v4 | Q5_K_S Q6_K Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q5_K_M Q8_0 Q2_K Q3_K_S Q4_0 Q5_0 | 4096 | 5120 | [README.md](models/Huginn-13B-v4/README.md) |
| Huginn 13B v4.5 | Q6_K Q8_0 Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_M | 4096 | 5120 | [README.md](models/Huginn-13B-v4.5/README.md) |
| Huginn 22B Prototype | Q5_K_S Q6_K Q3_K_L Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q8_0 Q2_K Q3_K_M Q3_K_S | 4096 | 6656 | [README.md](models/Huginn-22B-Prototype/README.md) |
| Huginn v3 13B | Q8_0 Q2_K Q3_K_L Q4_0 Q4_K_M Q4_K_S Q5_K_M Q6_K Q3_K_M Q3_K_S Q5_0 Q5_K_S | 4096 | 5120 | [README.md](models/Huginn-v3-13B/README.md) |
| Iambe 20B Dare | Q2_K Q3_K_M Q3_K_S Q4_K_M Q5_K_M Q5_K_S Q8_0 Q3_K_L Q4_0 Q4_K_S Q5_0 Q6_K | 4096 | 5120 | [README.md](models/Iambe-20B-DARE/README.md) |
| Iambe RP cDPO 20B | Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q5_K_M Q6_K Q2_K Q4_0 Q4_K_S Q5_0 Q5_K_S | 4096 | 5120 | [README.md](models/Iambe-RP-cDPO-20B/README.md) |
| Iambe Storyteller 20B | Q2_K Q3_K_L Q3_K_M Q4_K_M Q5_K_S Q8_0 Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q6_K | 4096 | 5120 | [README.md](models/Iambe-Storyteller-20B/README.md) |
| Inairtra 7B | Q3_K_M Q4_0 Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_S Q4_K_M Q4_K_S Q5_K_S Q3_K_L | 32768 | 4096 | [README.md](models/Inairtra-7B/README.md) |
| Inkbot 13B 4K | Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q2_K Q4_0 Q5_0 Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Inkbot-13B-4k/README.md) |
| Inkbot 13B 8K 0.2 | Q4_K_M Q4_K_S Q5_K_M Q3_K_M Q3_K_L Q3_K_S Q4_0 Q5_0 Q5_K_S Q6_K Q8_0 Q2_K | 4096 | 5120 | [README.md](models/Inkbot-13B-8k-0.2/README.md) |
| Instruct Mixtral 8X7B V0.1 Dolly15K | Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_M | 32768 | 4096 | [README.md](models/Instruct_Mixtral-8x7B-v0.1_Dolly15K/README.md) |
| JanniesBasedLigma L2 13B | Q4_0 Q4_K_M Q5_K_M Q8_0 Q5_0 Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_S | 4096 | 5120 | [README.md](models/JanniesBasedLigma-L2-13B/README.md) |
| Kai 7B Instruct | Q4_K_M Q4_K_S Q5_K_M Q6_K Q2_K Q3_K_M Q3_K_S Q5_K_S Q8_0 Q3_K_L Q4_0 Q5_0 | 32768 | 4096 | [README.md](models/KAI-7B-Instruct/README.md) |
| Kai 7B Beta | Q3_K_S Q4_K_M Q5_0 Q5_K_S Q8_0 Q2_K Q3_K_L Q4_K_S Q5_K_M Q6_K Q3_K_M Q4_0 | 32768 | 4096 | [README.md](models/KAI-7B-beta/README.md) |
| Kaori 70B v1 | Q3_K_M Q5_0 Q2_K Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q3_K_L | 4096 | 8192 | [README.md](models/Kaori-70B-v1/README.md) |
| Karen TheEditor V2 Creative Mistral 7B | Q2_K Q3_K_M Q5_K_M Q8_0 Q4_K_S Q5_0 Q5_K_S Q6_K Q3_K_L Q3_K_S Q4_0 Q4_K_M | 32768 | 4096 | [README.md](models/Karen_TheEditor_V2_CREATIVE_Mistral_7B/README.md) |
| Karen TheEditor V2 Strict Mistral 7B | Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_0 Q2_K Q4_K_S Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_M | 32768 | 4096 | [README.md](models/Karen_TheEditor_V2_STRICT_Mistral_7B/README.md) |
| Karen The Editor 13B | Q3_K_L Q3_K_M Q4_K_S Q2_K Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/Karen_theEditor_13B/README.md) |
| Kimiko Mistral 7B | Q5_K_M Q6_K Q8_0 Q2_K Q3_K_S Q4_0 Q5_0 Q5_K_S Q3_K_L Q3_K_M Q4_K_M Q4_K_S | 32768 | 4096 | [README.md](models/Kimiko-Mistral-7B/README.md) |
| Kuchiki 1.1 L2 7B | Q6_K Q2_K Q3_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_S Q4_0 Q4_K_M Q8_0 | 4096 | 4096 | [README.md](models/Kuchiki-1.1-L2-7B/README.md) |
| Kuchiki L2 7B | Q4_0 Q5_0 Q5_K_S Q6_K Q2_K Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q8_0 Q3_K_L Q3_K_M | 4096 | 4096 | [README.md](models/Kuchiki-L2-7B/README.md) |
| Llama2 13B Psyfighter2 | Q4_0 Q4_K_M Q6_K Q8_0 Q2_K Q3_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_S | 4096 | 5120 | [README.md](models/LLaMA2-13B-Psyfighter2/README.md) |
| Llama2 13B Tiefighter | Q6_K Q2_K Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q8_0 Q3_K_L Q3_K_M Q4_0 Q5_0 | 4096 | 5120 | [README.md](models/LLaMA2-13B-Tiefighter/README.md) |
| Llama2 13B TiefighterLR | Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_S Q8_0 Q2_K Q4_K_S Q5_K_M Q6_K Q3_K_M | 4096 | 5120 | [README.md](models/LLaMA2-13B-TiefighterLR/README.md) |
| Llama 2 13B SFT V1 | Q3_K_L Q3_K_M Q4_K_S Q6_K Q5_K_S Q8_0 Q2_K Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M | 2048 | 5120 | [README.md](models/LLaMA_2_13B_SFT_v1/README.md) |
| LMCocktail 10.7B v1 | Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 Q6_K Q3_K_L Q4_K_S Q5_K_M Q5_K_S Q8_0 Q2_K | 4096 | 4096 | [README.md](models/LMCocktail-10.7B-v1/README.md) |
| LMCocktail Phi 2 v1 | Q4_K_S Q5_K_M Q3_K_L Q3_K_M Q4_K_M Q5_0 Q5_K_S Q6_K Q8_0 Q2_K Q3_K_S Q4_0 | 0 | 0 | [README.md](models/LMCocktail-phi-2-v1/README.md) |
| Luna SOLARkrautLM Instruct | Q4_K_M Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q3_K_M Q4_0 Q3_K_S Q4_K_S Q5_0 Q5_K_M | 4096 | 4096 | [README.md](models/LUNA-SOLARkrautLM-Instruct/README.md) |
| Lemur 70B Chat v1 | Q3_K_S Q4_0 Q5_0 Q5_K_S Q2_K Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q5_K_M | 4096 | 8192 | [README.md](models/Lemur-70B-Chat-v1/README.md) |
| Leo Mistral Hessianai 7B Chat | Q5_K_M Q6_K Q8_0 Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_S Q2_K Q3_K_L | 32768 | 4096 | [README.md](models/Leo-Mistral-Hessianai-7B-Chat/README.md) |
| LeoScorpius 7B | Q3_K_L Q3_K_M Q5_0 Q5_K_S Q8_0 Q2_K Q4_0 Q4_K_M Q4_K_S Q5_K_M Q6_K Q3_K_S | 32768 | 4096 | [README.md](models/LeoScorpius-7B/README.md) |
| Leoscorpius GreenNode Alpaca 7B v1 | Q2_K Q3_K_M Q3_K_S Q4_K_S Q5_0 Q3_K_L Q4_0 Q4_K_M Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/LeoScorpius-GreenNode-Alpaca-7B-v1/README.md) |
| Leoscorpius GreenNode Platypus 7B v1 | Q3_K_S Q4_0 Q4_K_M Q5_K_M Q6_K Q2_K Q3_K_M Q5_0 Q5_K_S Q8_0 Q3_K_L Q4_K_S | 32768 | 4096 | [README.md](models/LeoScorpius-GreenNode-Platypus-7B-v1/README.md) |
| Lewd Sydney 20B | Q4_K_S Q5_K_M Q3_K_L Q3_K_S Q4_0 Q5_0 Q5_K_S Q6_K Q8_0 Q2_K Q3_K_M Q4_K_M | 4096 | 5120 | [README.md](models/Lewd-Sydney-20B/README.md) |
| Lila 103B L2 | Q2_K Q3_K_M Q3_K_S | 4096 | 8192 | [README.md](models/Lila-103B-L2/README.md) |
| Lila 70B L2 | Q2_K Q3_K_L Q3_K_S Q5_K_S Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M | 4096 | 8192 | [README.md](models/Lila-70B-L2/README.md) |
| Llama 2 13B | Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_K_S Q6_K Q2_K Q4_K_M Q4_K_S Q5_0 Q5_K_M | 4096 | 5120 | [README.md](models/Llama-2-13B/README.md) |
| Llama 2 13b Chat Dutch | Q3_K_M Q5_K_M Q5_K_S Q6_K Q8_0 Q5_0 Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S | 4096 | 5120 | [README.md](models/Llama-2-13B-Chat-Dutch/README.md) |
| Llama 2 13B German Assistant v4 | Q4_0 Q5_K_S Q3_K_M Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_M Q6_K Q2_K Q8_0 | 2048 | 5120 | [README.md](models/Llama-2-13B-German-Assistant-v4/README.md) |
| Llama 2 13B Chat | Q2_K Q3_K_L Q3_K_S Q4_K_S Q5_K_M Q5_K_S Q3_K_M Q4_0 Q4_K_M Q5_0 Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Llama-2-13B-chat/README.md) |
| Llama 2 70B | Q4_0 Q4_K_M Q4_K_S Q5_0 Q2_K Q3_K_M Q3_K_S Q5_K_M Q5_K_S Q3_K_L | 4096 | 8192 | [README.md](models/Llama-2-70B/README.md) |
| Llama 2 70B Chat | Q2_K Q3_K_L Q3_K_S Q4_0 Q5_K_M Q5_K_S Q3_K_M Q4_K_M Q4_K_S Q5_0 | 4096 | 8192 | [README.md](models/Llama-2-70B-Chat/README.md) |
| Open-Assistant Llama2 70B SFT OASST | Q3_K_M Q3_K_S Q4_0 Q5_0 Q2_K Q3_K_L Q5_K_M Q5_K_S Q4_K_M Q4_K_S | 4096 | 8192 | [README.md](models/Llama-2-70B-OASST-1-200/README.md) |
| Llama 2 70B Orca 200k | Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_K_M Q2_K Q4_K_M Q4_K_S Q5_0 Q5_K_S | 4096 | 8192 | [README.md](models/Llama-2-70B-Orca-200k/README.md) |
| Llama 2 7B | Q3_K_L Q3_K_M Q4_0 Q5_K_M Q5_K_S Q6_K Q8_0 Q2_K Q3_K_S Q4_K_M Q4_K_S Q5_0 | 4096 | 4096 | [README.md](models/Llama-2-7B/README.md) |
| Llama2 7B 32K Instruct | Q6_K Q8_0 Q2_K Q3_K_S Q4_0 Q5_0 Q5_K_S Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q5_K_M | 32768 | 4096 | [README.md](models/Llama-2-7B-32K-Instruct/README.md) |
| Llama 2 7B Chat | Q2_K Q3_K_L Q3_K_S Q5_0 Q5_K_M Q6_K Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_K_S Q8_0 | 4096 | 4096 | [README.md](models/Llama-2-7B-Chat/README.md) |
| Llama 2 7B FT Instruct ES | Q4_0 Q4_K_S Q5_K_M Q2_K Q3_K_M Q3_K_S Q5_K_S Q6_K Q8_0 Q3_K_L Q4_K_M Q5_0 | 4096 | 4096 | [README.md](models/Llama-2-7B-ft-instruct-es/README.md) |
| Llama 2 7B Vietnamese 20K | Q4_0 Q5_K_M Q6_K Q2_K Q3_K_L Q3_K_M Q5_0 Q5_K_S Q8_0 Q3_K_S Q4_K_M Q4_K_S | 4096 | 4096 | [README.md](models/Llama-2-7B-vietnamese-20k/README.md) |
| Llama 2 Coder 7B | Q3_K_S Q4_K_M Q5_K_M Q8_0 Q3_K_L Q3_K_M Q4_0 Q4_K_S Q5_0 Q5_K_S Q6_K Q2_K | 4096 | 4096 | [README.md](models/Llama-2-Coder-7B/README.md) |
| Llama2 13B MegaCode2 OASST | Q3_K_L Q4_0 Q4_K_S Q5_K_S Q8_0 Q2_K Q3_K_S Q4_K_M Q5_0 Q5_K_M Q6_K Q3_K_M | 2048 | 5120 | [README.md](models/Llama2-13B-MegaCode2-OASST/README.md) |
| Llama2 22B Daydreamer2 v3 | Q8_0 Q3_K_S Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q2_K Q6_K Q5_K_S | 4098 | 6656 | [README.md](models/Llama2-22B-Daydreamer-v3/README.md) |
| Llama2 70B SFT v10 | Q4_K_S Q5_0 Q5_K_M Q2_K Q3_K_L Q4_K_M Q5_K_S Q3_K_M Q3_K_S Q4_0 | 4096 | 8192 | [README.md](models/Llama2-70B-OASST-SFT-v10/README.md) |
| Llama2 Chat AYT 13B | Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_L Q4_0 Q4_K_S Q5_K_S Q3_K_M Q3_K_S | 4096 | 5120 | [README.md](models/Llama2-Chat-AYT-13B/README.md) |
| Llama2 Chat AYB 13B | Q3_K_L Q3_K_S Q4_K_M Q5_0 Q5_K_M Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_S Q5_K_S Q6_K | 4096 | 5120 | [README.md](models/Llama2-chat-AYB-13B/README.md) |
| LlamaGuard 7B | Q3_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M | 2048 | 4096 | [README.md](models/LlamaGuard-7B/README.md) |
| Llamix2 MLewd 4X13B | Q5_1 Q8_0 Q4_0 Q4_1 Q5_0 | 32768 | 5120 | [README.md](models/Llamix2-MLewd-4x13B/README.md) |
| LlongOrca 13B 16K | Q6_K Q8_0 Q3_K_M Q4_K_S Q5_0 Q5_K_M Q4_K_M Q5_K_S Q2_K Q3_K_L Q3_K_S Q4_0 | 16384 | 5120 | [README.md](models/LlongOrca-13B-16K/README.md) |
| LlongOrca 7B 16K | Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q3_K_S Q4_0 Q5_K_S Q6_K Q8_0 | 16384 | 4096 | [README.md](models/LlongOrca-7B-16K/README.md) |
| LoKuS 13B | Q4_0 Q4_K_M Q5_K_M Q8_0 Q6_K Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_S Q5_0 Q5_K_S | 4096 | 5120 | [README.md](models/LoKuS-13B/README.md) |
| LongAlpaca 70B | Q4_K_S Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q2_K | 4096 | 8192 | [README.md](models/LongAlpaca-70B/README.md) |
| LosslessMegaCoder Llama2 13B Mini | Q3_K_L Q3_K_M Q4_K_S Q6_K Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_S Q4_0 Q4_K_M Q5_0 | 2048 | 5120 | [README.md](models/LosslessMegaCoder-Llama2-13B-Mini/README.md) |
| Lossless MegaCoder Llama2 7B Mini | Q6_K Q2_K Q3_K_S Q4_0 Q4_K_M Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q4_K_S Q5_0 Q8_0 | 4096 | 4096 | [README.md](models/LosslessMegaCoder-Llama2-7B-Mini/README.md) |
| Loyal Macaroni Maid 7B | Q3_K_M Q5_0 Q4_K_S Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Loyal-Macaroni-Maid-7B/README.md) |
| Luban 13B | Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_0 Q5_K_M Q2_K Q4_K_S Q5_K_S Q6_K Q8_0 Q4_K_M | 4096 | 5120 | [README.md](models/Luban-13B/README.md) |
| Luna AI Llama2 Uncensored | Q5_K_M Q5_K_S Q2_K Q3_K_S Q4_0 Q4_K_S Q6_K Q8_0 Q3_K_L Q3_K_M Q4_K_M Q5_0 | 2048 | 4096 | [README.md](models/Luna-AI-Llama2-Uncensored/README.md) |
| MAmmoTH 13B | Q5_K_M Q8_0 Q2_K Q3_K_L Q3_K_M Q4_0 Q4_K_M Q5_0 Q3_K_S Q4_K_S Q5_K_S Q6_K | 4096 | 5120 | [README.md](models/MAmmoTH-13B/README.md) |
| MAmmoTH 70B | Q4_K_S Q5_0 Q5_K_S Q2_K Q3_K_S Q4_0 Q4_K_M Q3_K_L Q3_K_M Q5_K_M | 4096 | 8192 | [README.md](models/MAmmoTH-70B/README.md) |
| MAmmoTH 7B | Q3_K_M Q3_K_S Q5_0 Q8_0 Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_L Q4_0 Q4_K_M Q4_K_S | 4096 | 4096 | [README.md](models/MAmmoTH-7B/README.md) |
| MAmmoTH Coder 13B | Q5_K_M Q5_K_S Q2_K Q3_K_L Q4_0 Q5_0 Q6_K Q8_0 Q3_K_M Q3_K_S Q4_K_M Q4_K_S | 16384 | 5120 | [README.md](models/MAmmoTH-Coder-13B/README.md) |
| MAmmoTH Coder 34B | Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_S Q4_0 Q4_K_S Q8_0 Q3_K_L Q3_K_M Q4_K_M Q6_K | 16384 | 8192 | [README.md](models/MAmmoTH-Coder-34B/README.md) |
| MLewd L2 Chat 13B | Q3_K_L Q3_K_S Q4_K_S Q5_0 Q5_K_M Q6_K Q2_K Q4_0 Q4_K_M Q5_K_S Q8_0 Q3_K_M | 4096 | 5120 | [README.md](models/MLewd-L2-Chat-13B/README.md) |
| MLewd ReMM L2 Chat 20B | Q2_K Q3_K_L Q4_0 Q4_K_M Q4_K_S Q5_0 Q6_K Q3_K_M Q3_K_S Q5_K_M Q5_K_S Q8_0 | 4096 | 5120 | [README.md](models/MLewd-ReMM-L2-Chat-20B/README.md) |
| MLewd ReMM L2 Chat 20b Inverted | Q3_K_M Q3_K_S Q4_0 Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_L Q4_K_M Q4_K_S Q5_0 Q6_K | 4096 | 5120 | [README.md](models/MLewd-ReMM-L2-Chat-20B-Inverted/README.md) |
| MLewdBoros L2 13B | Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_S Q6_K Q2_K Q4_0 Q5_K_M Q8_0 | 4096 | 5120 | [README.md](models/MLewdBoros-L2-13B/README.md) |
| MLewdBoros LRSGPT 2Char 13B | Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q8_0 Q2_K Q3_K_M Q5_K_S Q6_K Q3_K_L Q3_K_S | 4096 | 5120 | [README.md](models/MLewdBoros-LRPSGPT-2Char-13B/README.md) |
| MXLewd L2 20B | Q6_K Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_S Q5_K_M Q5_K_S Q3_K_L Q3_K_S Q4_K_M Q5_0 | 4096 | 5120 | [README.md](models/MXLewd-L2-20B/README.md) |
| Mxlewdmini L2 13B | Q5_K_S Q3_K_L Q3_K_S Q4_0 Q4_K_S Q5_K_M Q6_K Q8_0 Q2_K Q3_K_M Q4_K_M Q5_0 | 4096 | 5120 | [README.md](models/MXLewdMini-L2-13B/README.md) |
| MadMix v0.2 | Q4_K_M Q5_0 Q5_K_S Q2_K Q3_K_S Q4_0 Q5_K_M Q6_K Q8_0 Q3_K_L Q3_K_M Q4_K_S | 32768 | 4096 | [README.md](models/MadMix-v0.2/README.md) |
| Magicoder S DS 6.7B | Q3_K_L Q3_K_M Q4_0 Q5_K_M Q8_0 Q2_K Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_S Q6_K | 16384 | 4096 | [README.md](models/Magicoder-S-DS-6.7B/README.md) |
| Manticore 13B | Q2_K Q3_K_S Q4_0 Q4_K_M Q4_K_S Q6_K Q8_0 Q3_K_L Q3_K_M Q5_0 Q5_K_M Q5_K_S | 2048 | 5120 | [README.md](models/Manticore-13B/README.md) |
| Marcoroni 13B | Q4_K_S Q5_K_M Q6_K Q8_0 Q3_K_L Q4_0 Q3_K_S Q4_K_M Q5_0 Q5_K_S Q2_K Q3_K_M | 4096 | 5120 | [README.md](models/Marcoroni-13B/README.md) |
| Marcoroni 70B | Q3_K_L Q4_0 Q4_K_M Q2_K Q3_K_S Q4_K_S Q5_0 Q5_K_M Q5_K_S Q3_K_M | 4096 | 8192 | [README.md](models/Marcoroni-70B/README.md) |
| Marcoroni 70B v1 | Q4_K_S Q5_0 Q5_K_M Q4_K_M Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_K_S Q2_K | 4096 | 8192 | [README.md](models/Marcoroni-70B-v1/README.md) |
| Marcoroni 7B V2 | Q3_K_S Q4_0 Q5_K_S Q6_K Q2_K Q3_K_M Q4_K_S Q5_0 Q5_K_M Q8_0 Q3_K_L Q4_K_M | 32768 | 4096 | [README.md](models/Marcoroni-7B-v2/README.md) |
| Marcoroni 7B V3 | Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q8_0 Q5_0 Q6_K | 32768 | 4096 | [README.md](models/Marcoroni-7B-v3/README.md) |
| Marcoroni 7b | Q2_K Q3_K_L Q6_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q8_0 | 4096 | 4096 | [README.md](models/Marcoroni-7b/README.md) |
| Marcoroni Neural Chat 7B v1 | Q8_0 Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q2_K Q3_K_L Q3_K_S Q4_K_S Q5_K_S | 32768 | 4096 | [README.md](models/Marcoroni-neural-chat-7B-v1/README.md) |
| Marx 3B V3 | Q5_K_M Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q8_0 Q4_0 Q4_K_S Q5_0 Q5_K_S Q6_K | 4096 | 2560 | [README.md](models/Marx-3B-v3/README.md) |
| Megamix S1 13B | Q4_K_M Q4_K_S Q5_K_M Q6_K Q4_0 Q5_0 Q5_K_S Q8_0 Q2_K Q3_K_L Q3_K_M Q3_K_S | 4096 | 5120 | [README.md](models/MegaMix-S1-13B/README.md) |
| Megamix T1 13B | Q3_K_M Q4_0 Q4_K_S Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q5_0 Q5_K_M Q3_K_S Q4_K_M | 4096 | 5120 | [README.md](models/MegaMix-T1-13B/README.md) |
| Megamix A1 13B | Q4_K_S Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_S Q4_K_M Q8_0 Q3_K_M Q4_0 Q5_0 | 4096 | 5120 | [README.md](models/Megamix-A1-13B/README.md) |
| MelloGPT | Q5_K_S Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 Q3_K_L Q4_K_S Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/MelloGPT/README.md) |
| Mergemonster 13B 20231124 | Q4_K_S Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_M Q4_0 Q5_0 Q8_0 Q3_K_L Q3_K_S Q4_K_M | 4096 | 5120 | [README.md](models/MergeMonster-13B-20231124/README.md) |
| Merged AGI 7B | Q8_0 Q2_K Q3_K_S Q4_K_M Q5_K_M Q6_K Q5_K_S Q3_K_L Q3_K_M Q4_0 Q4_K_S Q5_0 | 32768 | 4096 | [README.md](models/Merged-AGI-7B/README.md) |
| MetaMath 13B V1.0 | Q3_K_M Q3_K_S Q4_K_S Q8_0 Q2_K Q3_K_L Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q6_K | 4096 | 5120 | [README.md](models/MetaMath-13B-V1.0/README.md) |
| MetaMath 70B V1.0 | Q2_K Q3_K_S Q4_K_M Q5_0 Q3_K_L Q3_K_M Q4_0 Q4_K_S Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/MetaMath-70B-V1.0/README.md) |
| MetaMath 7B V1.0 | Q6_K Q8_0 Q3_K_L Q3_K_M Q4_K_M Q5_K_S Q5_0 Q5_K_M Q2_K Q3_K_S Q4_0 Q4_K_S | 4096 | 4096 | [README.md](models/MetaMath-7B-V1.0/README.md) |
| Metamath Cybertron Starling | Q5_0 Q5_K_M Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_S Q6_K Q8_0 Q3_K_L | 32768 | 4096 | [README.md](models/MetaMath-Cybertron-Starling/README.md) |
| Metamath Mistral 7B | Q4_K_M Q4_K_S Q5_K_M Q4_0 Q3_K_L Q3_K_M Q3_K_S Q5_0 Q5_K_S Q6_K Q8_0 Q2_K | 32768 | 4096 | [README.md](models/MetaMath-Mistral-7B/README.md) |
| Metis 0.1 | Q4_K_M Q5_K_M Q3_K_L Q3_K_S Q4_0 Q5_0 Q5_K_S Q6_K Q8_0 Q2_K Q3_K_M Q4_K_S | 32768 | 4096 | [README.md](models/Metis-0.1/README.md) |
| Metis 0.3 | Q6_K Q2_K Q3_K_L Q5_0 Q5_K_M Q4_K_S Q5_K_S Q8_0 Q3_K_M Q3_K_S Q4_0 Q4_K_M | 32768 | 4096 | [README.md](models/Metis-0.3/README.md) |
| Metis 0.4 | Q3_K_M Q4_K_S Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M | 32768 | 4096 | [README.md](models/Metis-0.4/README.md) |
| Metis 0.5 | Q5_0 Q5_K_S Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q6_K Q3_K_L Q3_K_S Q4_K_S Q5_K_M | 32768 | 4096 | [README.md](models/Metis-0.5/README.md) |
| Mini Synatra 7B 02 | Q3_K_L Q3_K_S Q4_K_S Q5_0 Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q6_K | 32768 | 4096 | [README.md](models/Mini_Synatra_7B_02/README.md) |
| Mistrp Airoboros 7B | Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q2_K Q3_K_S Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_L | 32768 | 4096 | [README.md](models/MistRP-Airoboros-7B/README.md) |
| Mistral 11B CC Air RP | Q3_K_L Q3_K_M Q4_0 Q4_K_M Q5_K_M Q5_K_S Q6_K Q2_K Q4_K_S Q5_0 Q8_0 Q3_K_S | 32768 | 4096 | [README.md](models/Mistral-11B-CC-Air-RP/README.md) |
| Mistral 11B OmniMix | Q4_K_M Q4_K_S Q5_K_M Q8_0 Q4_0 Q3_K_L Q3_K_M Q3_K_S Q5_0 Q5_K_S Q6_K Q2_K | 32768 | 4096 | [README.md](models/Mistral-11B-OmniMix/README.md) |
| Mistral 7B Aezakmi v1 | Q3_K_M Q4_0 Q4_K_M Q5_K_M Q8_0 Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_S Q4_K_S Q5_0 | 32768 | 4096 | [README.md](models/Mistral-7B-AEZAKMI-v1/README.md) |
| Mistral 7B Claude Chat | Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q2_K Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_M | 32768 | 4096 | [README.md](models/Mistral-7B-Claude-Chat/README.md) |
| Mistral 7B Code 16K qLoRA | Q4_K_S Q5_0 Q5_K_S Q8_0 Q3_K_L Q3_K_M Q3_K_S Q5_K_M Q6_K Q2_K Q4_0 Q4_K_M | 32768 | 4096 | [README.md](models/Mistral-7B-Code-16K-qlora/README.md) |
| Mistral 7B Instruct v0.1 | Q2_K Q3_K_L Q4_K_M Q4_K_S Q5_K_S Q6_K Q8_0 Q3_K_M Q3_K_S Q4_0 Q5_0 Q5_K_M | 32768 | 4096 | [README.md](models/Mistral-7B-Instruct-v0.1/README.md) |
| Mistral 7B Instruct v0.2 | Q5_0 Q6_K Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_S | 32768 | 4096 | [README.md](models/Mistral-7B-Instruct-v0.2/README.md) |
| Mistral 7B Instruct V0.2 DARE | Q4_K_M Q4_K_S Q5_0 Q6_K Q2_K Q3_K_M Q3_K_S Q5_K_S Q8_0 Q3_K_L Q4_0 Q5_K_M | 32768 | 4096 | [README.md](models/Mistral-7B-Instruct-v0.2-DARE/README.md) |
| Mistral 7B Merge 14 v0.1 | Q8_0 Q3_K_L Q3_K_S Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_M Q4_0 Q4_K_M | 32768 | 4096 | [README.md](models/Mistral-7B-Merge-14-v0.1/README.md) |
| Mistral 7B OpenOrca | Q2_K Q3_K_M Q4_K_M Q5_0 Q8_0 Q3_K_L Q3_K_S Q4_0 Q4_K_S Q5_K_M Q5_K_S Q6_K | 32768 | 4096 | [README.md](models/Mistral-7B-OpenOrca/README.md) |
| Mistral 7B OpenOrca oasst Top1 2023 08 25 v1 | Q4_K_S Q5_K_M Q5_K_S Q6_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q8_0 Q2_K Q3_K_L Q3_K_S | 32768 | 4096 | [README.md](models/Mistral-7B-OpenOrca-oasst_top1_2023-08-25-v1/README.md) |
| Mistral 7B Phibrarian 32K | Q3_K_M Q3_K_S Q6_K Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q4_0 Q8_0 | 32768 | 4096 | [README.md](models/Mistral-7B-Phibrarian-32K/README.md) |
| Mistral 7B SciPhi 32K | Q3_K_M Q4_0 Q4_K_M Q5_K_S Q2_K Q3_K_S Q4_K_S Q5_0 Q5_K_M Q6_K Q8_0 Q3_K_L | 32768 | 4096 | [README.md](models/Mistral-7B-SciPhi-32k/README.md) |
| Mistral 7B CodeAlpaca Lora | Q3_K_M Q3_K_S Q4_0 Q4_K_M Q6_K Q2_K Q3_K_L Q4_K_S Q5_0 Q5_K_M Q5_K_S Q8_0 | 32768 | 4096 | [README.md](models/Mistral-7B-codealpaca-lora/README.md) |
| Mistral 7B v0.1 | Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_L Q3_K_M Q4_K_S Q5_0 Q3_K_S Q4_0 Q4_K_M Q6_K | 32768 | 4096 | [README.md](models/Mistral-7B-v0.1/README.md) |
| Mistral ClaudeLimaRP v3 7B | Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_M Q3_K_S Q4_0 Q4_K_M Q2_K Q3_K_L | 32768 | 4096 | [README.md](models/Mistral-ClaudeLimaRP-v3-7B/README.md) |
| Mistral Pygmalion 7B | Q2_K Q3_K_S Q4_K_M Q5_0 Q6_K Q8_0 Q3_K_L Q3_K_M Q4_0 Q4_K_S Q5_K_M Q5_K_S | 4096 | 4096 | [README.md](models/Mistral-Pygmalion-7B/README.md) |
| Mistral Trismegistus 7B | Q3_K_S Q4_0 Q5_K_M Q2_K Q3_K_L Q4_K_S Q5_0 Q5_K_S Q6_K Q8_0 Q3_K_M Q4_K_M | 32768 | 4096 | [README.md](models/Mistral-Trismegistus-7B/README.md) |
| MistralLite 7B | Q4_K_M Q5_0 Q6_K Q3_K_S Q4_0 Q3_K_M Q4_K_S Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_L | 32768 | 4096 | [README.md](models/MistralLite-7B/README.md) |
| MistralMakise Merged 13B | Q3_K_L Q4_0 Q4_K_M Q5_0 Q5_K_M Q2_K Q3_K_M Q3_K_S Q4_K_S Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/MistralMakise-Merged-13B/README.md) |
| Mistral 7B Dolphin2.1 Lima0.5 | Q2_K Q4_0 Q4_K_M Q5_0 Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_K_S Q5_K_M Q5_K_S Q6_K | 32768 | 4096 | [README.md](models/Mistral_7B_Dolphin2.1_LIMA0.5/README.md) |
| Mistralic 7B-1 | Q3_K_M Q3_K_S Q4_0 Q5_K_S Q8_0 Q3_K_L Q4_K_M Q4_K_S Q5_0 Q5_K_M Q6_K Q2_K | 32768 | 4096 | [README.md](models/Mistralic-7B-1/README.md) |
| Mixtral 8X7B Instruct v0.1 | Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_M Q4_0 | 32768 | 4096 | [README.md](models/Mixtral-8x7B-Instruct-v0.1/README.md) |
| Mixtral 8X7B MoE RP Story | Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q2_K | 32768 | 4096 | [README.md](models/Mixtral-8x7B-MoE-RP-Story/README.md) |
| Mixtral 8X7B v0.1 | Q6_K Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M | 32768 | 4096 | [README.md](models/Mixtral-8x7B-v0.1/README.md) |
| Mixtral Fusion 4X7B Instruct v0.1 | Q6_K Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M | 32768 | 4096 | [README.md](models/Mixtral-Fusion-4x7B-Instruct-v0.1/README.md) |
| Mixtral SlimOrca 8X7B | Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q2_K | 32768 | 4096 | [README.md](models/Mixtral-SlimOrca-8x7B/README.md) |
| MixtralOrochi8X7B | Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K | 32768 | 4096 | [README.md](models/MixtralOrochi8x7B/README.md) |
| MixtralRPChat ZLoss | Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_M | 32768 | 4096 | [README.md](models/MixtralRPChat-ZLoss/README.md) |
| Mixtral 7Bx2 MoE | Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K | 32768 | 4096 | [README.md](models/Mixtral_7Bx2_MoE/README.md) |
| MoMo 70B v1.1 | Q4_K_S Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M Q4_K_M Q3_K_S Q4_0 Q5_0 | 4096 | 8192 | [README.md](models/MoMo-70B-V1.1/README.md) |
| MonadGPT 7B | Q5_K_S Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_S Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/MonadGPT/README.md) |
| MysticFusion 13B | Q2_K Q4_K_M Q6_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S Q8_0 | 4096 | 5120 | [README.md](models/MysticFusion-13B/README.md) |
| Mythalion 13B | Q3_K_L Q5_0 Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q2_K Q3_K_M Q3_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Mythalion-13B/README.md) |
| Mythical Destroyer L2 13B | Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q6_K Q2_K Q3_K_L Q5_K_M Q5_K_S Q8_0 | 4096 | 5120 | [README.md](models/Mythical-Destroyer-L2-13B/README.md) |
| Mythical Destroyer V2 L2 13B | Q2_K Q3_K_M Q4_K_M Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_L Q3_K_S Q4_0 Q4_K_S Q5_0 | 4096 | 5120 | [README.md](models/Mythical-Destroyer-V2-L2-13B/README.md) |
| MythoBoros 13B | Q4_K_S Q5_0 Q5_K_S Q6_K Q3_K_S Q4_0 Q4_K_M Q5_K_M Q8_0 Q2_K Q3_K_L Q3_K_M | 2048 | 5120 | [README.md](models/MythoBoros-13B/README.md) |
| MythoLogic 13B | Q8_0 Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_M Q6_K | 2048 | 5120 | [README.md](models/MythoLogic-13B/README.md) |
| Mythologic L2 13B | Q3_K_L Q3_K_S Q4_K_M Q5_0 Q5_K_S Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_S Q5_K_M Q6_K | 4096 | 5120 | [README.md](models/MythoLogic-L2-13B/README.md) |
| Mythologic Mini 7B | Q6_K Q3_K_M Q3_K_S Q5_0 Q5_K_S Q4_K_S Q5_K_M Q8_0 Q2_K Q3_K_L Q4_0 Q4_K_M | 4096 | 4096 | [README.md](models/MythoLogic-Mini-7B/README.md) |
| MythoMakiseMerged 13B | Q4_K_M Q5_K_S Q8_0 Q2_K Q3_K_M Q3_K_S Q4_0 Q6_K Q3_K_L Q4_K_S Q5_0 Q5_K_M | 4096 | 5120 | [README.md](models/MythoMakiseMerged-13B/README.md) |
| MythoMax L2 13B | Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q5_0 Q6_K Q8_0 Q2_K Q3_K_S Q4_0 Q5_K_M Q5_K_S | 4096 | 5120 | [README.md](models/MythoMax-L2-13B/README.md) |
| MythoMax L2 Kimiko v2 13B | Q3_K_M Q3_K_S Q5_K_M Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q4_K_S Q5_0 Q4_0 Q4_K_M | 4096 | 5120 | [README.md](models/MythoMax-L2-Kimiko-v2-13B/README.md) |
| MythoMist 7B | Q3_K_M Q3_K_S Q4_K_S Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_L Q4_0 Q4_K_M Q5_0 Q6_K | 32768 | 4096 | [README.md](models/MythoMist-7B/README.md) |
| MythoMix L2 13B | Q3_K_M Q4_0 Q4_K_S Q5_0 Q5_K_S Q8_0 Q2_K Q3_K_L Q5_K_M Q6_K Q3_K_S Q4_K_M | 4096 | 5120 | [README.md](models/MythoMix-L2-13B/README.md) |
| Naberius 7B | Q4_K_M Q4_K_S Q5_0 Q5_K_M Q8_0 Q6_K Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_K_S | 32768 | 4096 | [README.md](models/Naberius-7B/README.md) |
| Nanbeige 16B Base | Q4_K_S Q5_0 Q6_K Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_0 Q2_K Q4_K_M Q5_K_M Q5_K_S | 4096 | 5120 | [README.md](models/Nanbeige-16B-Base/README.md) |
| Nanbeige 16B Base 32K | Q8_0 Q2_K Q4_0 Q5_0 Q5_K_S Q6_K Q5_K_M Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q4_K_S | 4096 | 5120 | [README.md](models/Nanbeige-16B-Base-32K/README.md) |
| Nanbeige 16B Chat | Q4_K_S Q6_K Q2_K Q4_K_M Q3_K_S Q4_0 Q5_0 Q5_K_M Q5_K_S Q8_0 Q3_K_L Q3_K_M | 4096 | 5120 | [README.md](models/Nanbeige-16B-Chat/README.md) |
| Nanbeige 16B Chat 32K | Q3_K_M Q5_0 Q5_K_M Q5_K_S Q3_K_L Q4_K_S Q3_K_S Q4_0 Q8_0 Q2_K Q4_K_M Q6_K fp16 | 4096 | 5120 | [README.md](models/Nanbeige-16B-Chat-32K/README.md) |
| Nete 13B | Q4_K_S Q5_0 Q5_K_M Q6_K Q2_K Q3_K_S Q4_K_M Q5_K_S Q8_0 Q3_K_L Q3_K_M Q4_0 | 4096 | 5120 | [README.md](models/Nete-13B/README.md) |
| Nethena 13B | Q5_K_M Q5_K_S Q8_0 Q3_K_L Q3_K_S Q4_0 Q4_K_M Q6_K Q2_K Q3_K_M Q4_K_S Q5_0 | 4096 | 5120 | [README.md](models/Nethena-13B/README.md) |
| Nethena 20B | Q2_K Q3_K_M Q3_K_S Q4_K_S Q8_0 Q5_K_S Q6_K Q3_K_L Q4_0 Q4_K_M Q5_0 Q5_K_M | 4096 | 5120 | [README.md](models/Nethena-20B/README.md) |
| Nethena 20B Glued | Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Nethena-20B-Glued/README.md) |
| Nethena Mlewd Xwin 23B | Q5_K_M Q6_K Q8_0 Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_S Q2_K Q3_K_L Q4_0 | 4096 | 5120 | [README.md](models/Nethena-MLewd-Xwin-23B/README.md) |
| NeuralHermes 2.5 Mistral 7B | Q4_K_S Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 Q6_K | 32768 | 4096 | [README.md](models/NeuralHermes-2.5-Mistral-7B/README.md) |
| NeuralOrca 7B V1 | Q5_0 Q5_K_M Q5_K_S Q6_K Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q8_0 Q2_K Q3_K_M Q4_0 | 32768 | 4096 | [README.md](models/NeuralOrca-7B-v1/README.md) |
| Nexusraven 13B | Q3_K_M Q3_K_S Q4_0 Q5_0 Q5_K_S Q8_0 Q2_K Q3_K_L Q4_K_M Q4_K_S Q5_K_M Q6_K | 16384 | 5120 | [README.md](models/NexusRaven-13B/README.md) |
| NexusRaven V2 13B | Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q4_K_S Q4_K_M Q5_0 Q6_K Q8_0 Q2_K Q3_K_S Q4_0 | 16384 | 5120 | [README.md](models/NexusRaven-V2-13B/README.md) |
| Norocetacean 20B 10K | Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q6_K Q2_K Q3_K_M Q5_K_M Q5_K_S Q8_0 Q3_K_L | 10240 | 5120 | [README.md](models/Norocetacean-20B-10k/README.md) |
| Noromaid 13B v0.1.1 | Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_S Q2_K Q8_0 Q6_K Q5_K_M Q4_K_S | 4096 | 5120 | [README.md](models/Noromaid-13B-v0.1.1/README.md) |
| Noromaid 13B V0.2 | Q3_K_S Q4_K_S Q5_K_M Q3_K_M Q3_K_L Q4_0 Q4_K_M Q5_0 Q5_K_S Q6_K Q8_0 Q2_K | 4096 | 5120 | [README.md](models/Noromaid-13B-v0.2/README.md) |
| Noromaid 20B v0.1.1 | Q5_0 Q5_K_M Q3_K_L Q3_K_M Q4_0 Q4_K_S Q5_K_S Q6_K Q8_0 Q2_K Q3_K_S Q4_K_M | 4096 | 5120 | [README.md](models/Noromaid-20B-v0.1.1/README.md) |
| Nous Hermes 13B | Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_K_M Q5_K_S Q8_0 Q2_K Q4_K_S Q5_0 Q6_K Q3_K_L | 2048 | 5120 | [README.md](models/Nous-Hermes-13B/README.md) |
| Nous Hermes 13B Code | Q3_K_M Q4_0 Q5_0 Q4_K_S Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_S Q4_K_M Q8_0 | 4096 | 5120 | [README.md](models/Nous-Hermes-13B-Code/README.md) |
| Nous Hermes 2 Yi 34B | Q5_K_S Q2_K Q3_K_M Q3_K_S Q4_K_S Q5_0 Q8_0 Q3_K_L Q4_0 Q4_K_M Q5_K_M Q6_K | 4096 | 7168 | [README.md](models/Nous-Hermes-2-Yi-34B/README.md) |
| Nyxene v2 11B | Q4_K_S Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_M Q4_K_M Q5_0 Q6_K Q3_K_L Q3_K_S Q4_0 | 32768 | 4096 | [README.md](models/Nyxene-v2-11B/README.md) |
| Nyxene v3 11B | Q3_K_L Q4_0 Q4_K_M Q4_K_S Q5_K_M Q2_K Q3_K_M Q3_K_S Q5_0 Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Nyxene-v3-11B/README.md) |
| ORCA LLaMA 70B QLoRA | Q2_K Q3_K_L Q3_K_M Q4_0 Q4_K_M Q5_0 Q3_K_S Q4_K_S Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/ORCA_LLaMA_70B_QLoRA/README.md) |
| Open Hermes 2.5 Neural Chat 3.1 Frankenmerge 11B | Q4_K_S Q5_K_M Q5_K_S Q3_K_M Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_0 Q6_K Q8_0 Q2_K | 32768 | 4096 | [README.md](models/Open-Hermes-2.5-neural-chat-3.1-frankenmerge-11b/README.md) |
| Llama2 13B Orca 8K 3319 | Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_S Q2_K Q5_0 Q5_K_M Q6_K Q8_0 Q3_K_M | 8192 | 5120 | [README.md](models/OpenAssistant-Llama2-13B-Orca-8K-3319/README.md) |
| OpenBuddy Llama2 13B v11.1 | Q3_K_L Q4_K_M Q8_0 Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_M Q3_K_S Q6_K | 4096 | 5120 | [README.md](models/OpenBuddy-Llama2-13B-v11.1/README.md) |
| OpenBuddy Llama2 70b v10.1 | Q2_K Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_K_M Q3_K_L Q3_K_S Q5_0 Q5_K_S | 4096 | 8192 | [README.md](models/OpenBuddy-Llama2-70b-v10.1/README.md) |
| OpenHermes 2 Mistral 7B | Q5_0 Q5_K_M Q5_K_S Q3_K_L Q4_0 Q4_K_S Q4_K_M Q6_K Q8_0 Q2_K Q3_K_M Q3_K_S | 32768 | 4096 | [README.md](models/OpenHermes-2-Mistral-7B/README.md) |
| Openhermes 2.5 Mistral 7B | Q4_0 Q4_K_M Q5_0 Q8_0 Q2_K Q3_K_M Q3_K_S Q5_K_S Q6_K Q3_K_L Q4_K_S Q5_K_M | 32768 | 4096 | [README.md](models/OpenHermes-2.5-Mistral-7B/README.md) |
| Openhermes 2.5 Mistral 7B 16K | Q2_K Q4_0 Q4_K_M Q4_K_S Q5_K_S Q3_K_L Q3_K_M Q3_K_S Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/OpenHermes-2.5-Mistral-7B-16k/README.md) |
| OpenHermes 2.5 Neural Chat 7B V3.1 7B | Q3_K_M Q3_K_S Q4_K_S Q2_K Q3_K_L Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/OpenHermes-2.5-neural-chat-7B-v3-1-7B/README.md) |
| OpenHermes 2.5 Neural Chat 7B V3-2 7B | Q3_K_S Q4_K_M Q5_K_M Q5_K_S Q4_K_S Q5_0 Q6_K Q8_0 Q2_K Q3_K_L Q3_K_M Q4_0 | 32768 | 4096 | [README.md](models/OpenHermes-2.5-neural-chat-7B-v3-2-7B/README.md) |
| OpenHermes 2.5 Neural Chat V3 3 Slerp | Q3_K_M Q4_K_S Q8_0 Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q6_K | 32768 | 4096 | [README.md](models/OpenHermes-2.5-neural-chat-v3-3-Slerp/README.md) |
| OpenOrca Platypus2 13B | Q8_0 Q2_K Q5_K_S Q6_K Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q3_K_L Q3_K_M Q3_K_S | 4096 | 5120 | [README.md](models/OpenOrca-Platypus2-13B/README.md) |
| OpenOrca Zephyr 7B | Q4_0 Q4_K_M Q5_0 Q5_K_S Q6_K Q3_K_S Q3_K_L Q3_K_M Q4_K_S Q5_K_M Q8_0 Q2_K | 32768 | 4096 | [README.md](models/OpenOrca-Zephyr-7B/README.md) |
| OpenOrca Stx | Q4_K_S Q5_0 Q5_K_S Q6_K Q8_0 Q2_K Q3_K_M Q3_K_S Q5_K_M Q3_K_L Q4_0 Q4_K_M | 4096 | 5120 | [README.md](models/OpenOrca_Stx/README.md) |
| OpenZephyrchat | Q5_K_M Q5_K_S Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q2_K Q3_K_L Q3_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/OpenZephyrChat/README.md) |
| OpenZephyrChat V0.2 | Q4_0 Q4_K_M Q5_K_S Q8_0 Q2_K Q3_K_M Q4_K_S Q5_0 Q5_K_M Q6_K Q3_K_L Q3_K_S | 32768 | 4096 | [README.md](models/OpenZephyrChat-v0.2/README.md) |
| Optimus 7B | Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q6_K Q3_K_L Q3_K_S Q4_0 Q8_0 Q2_K Q3_K_M Q5_0 | 32768 | 4096 | [README.md](models/Optimus-7B/README.md) |
| Orca 2 13B | Q2_K Q3_K_L Q3_K_M Q4_0 Q4_K_M Q5_K_M Q5_K_S Q3_K_S Q4_K_S Q5_0 Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Orca-2-13B/README.md) |
| Orca 2 13B SFT V5 | Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q5_0 Q5_K_S Q6_K Q8_0 Q2_K Q4_0 Q4_K_S Q5_K_M | 4096 | 5120 | [README.md](models/Orca-2-13B-SFT_v5/README.md) |
| Orca 2 7B | Q4_K_S Q5_0 Q6_K Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q3_K_L Q5_K_M Q5_K_S Q8_0 | 4096 | 4096 | [README.md](models/Orca-2-7B/README.md) |
| Orca2Myth7.2 | Q3_K_L Q4_K_M Q4_K_S Q5_K_M Q6_K Q8_0 Q2_K Q3_K_M Q3_K_S Q4_0 Q5_0 Q5_K_S | 4096 | 5120 | [README.md](models/Orca2myth7.2/README.md) |
| OrcaMaid 13B | Q3_K_M Q3_K_S Q4_0 Q8_0 Q2_K Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q3_K_L | 4096 | 5120 | [README.md](models/OrcaMaid-13B/README.md) |
| Orcamaid V2 Fix 13B 32K | Q5_K_M Q6_K Q3_K_M Q5_0 Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_S Q8_0 Q2_K Q3_K_L | 32768 | 5120 | [README.md](models/OrcaMaid-v2-FIX-13B-32k/README.md) |
| OrcaMaidXL 17B 32K | Q4_K_S Q5_0 Q6_K Q8_0 Q3_K_S Q4_0 Q4_K_M Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M | 32768 | 5120 | [README.md](models/OrcaMaidXL-17B-32k/README.md) |
| OrionStar Yi 34B Chat Llama | Q4_K_M Q5_0 Q6_K Q8_0 Q2_K Q3_K_S Q4_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q4_K_S | 4096 | 7168 | [README.md](models/OrionStar-Yi-34B-Chat-Llama/README.md) |
| Pallas 0.3 | Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q6_K Q2_K Q3_K_L Q3_K_M Q4_K_M Q5_K_S Q8_0 | 200000 | 7168 | [README.md](models/Pallas-0.3/README.md) |
| Pallas 0.4 | Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_S Q8_0 Q3_K_M Q4_K_M Q5_0 | 200000 | 7168 | [README.md](models/Pallas-0.4/README.md) |
| Pandalyst 7B V1.1 | Q2_K Q4_K_M Q4_K_S Q6_K Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_0 Q5_K_M Q5_K_S | 16384 | 4096 | [README.md](models/Pandalyst-7B-V1.1/README.md) |
| Pandalyst 7B v1.2 | Q5_0 Q5_K_S Q6_K Q8_0 Q3_K_M Q4_K_S Q3_K_S Q4_0 Q4_K_M Q5_K_M Q2_K Q3_K_L | 16384 | 4096 | [README.md](models/Pandalyst-7B-v1.2/README.md) |
| Pandalyst 13B V1.0 | Q6_K Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q8_0 Q3_K_M Q5_0 Q5_K_S | 16384 | 5120 | [README.md](models/Pandalyst_13B_V1.0/README.md) |
| Pandora V1 10.7B | Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_M Q6_K Q8_0 Q3_K_L Q4_K_M | 32768 | 4096 | [README.md](models/Pandora-v1-10.7B/README.md) |
| Pandora V1 13B | Q4_K_M Q5_0 Q6_K Q3_K_L Q3_K_M Q4_0 Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_S Q4_K_S | 32768 | 4096 | [README.md](models/Pandora-v1-13B/README.md) |
| Phind CodeLlama 34B Python v1 | Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M Q8_0 Q3_K_S Q4_K_S Q6_K | 16384 | 8192 | [README.md](models/Phind-CodeLlama-34B-Python-v1/README.md) |
| Phind CodeLlama 34B v1 | Q6_K Q8_0 Q3_K_M Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_S Q4_0 | 16384 | 8192 | [README.md](models/Phind-CodeLlama-34B-v1/README.md) |
| Phind CodeLlama 34B v2 | Q3_K_L Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/Phind-CodeLlama-34B-v2/README.md) |
| Pivot 0.1 Evil A | Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q3_K_L Q3_K_S Q5_K_M Q5_K_S Q6_K | 32768 | 4096 | [README.md](models/PiVoT-0.1-Evil-a/README.md) |
| Pivot 0.1 Starling LM RP | Q3_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q8_0 | 8192 | 4096 | [README.md](models/PiVoT-0.1-Starling-LM-RP/README.md) |
| Pivot 0.1 Early | Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q8_0 Q2_K Q3_K_M Q5_0 Q5_K_M Q5_K_S Q6_K | 32768 | 4096 | [README.md](models/PiVoT-0.1-early/README.md) |
| Pivot 10.7B Mistral V0.2 | Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_S Q6_K Q8_0 Q2_K Q3_K_M Q4_K_S Q5_K_M Q3_K_L | 32768 | 4096 | [README.md](models/PiVoT-10.7B-Mistral-v0.2/README.md) |
| Pivot 10.7B Mistral V0.2 RP | Q2_K Q3_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q8_0 | 32768 | 4096 | [README.md](models/PiVoT-10.7B-Mistral-v0.2-RP/README.md) |
| Pivot Merge A 7B | Q5_K_M Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q8_0 Q2_K Q3_K_S Q5_0 Q5_K_S Q6_K | 8192 | 4096 | [README.md](models/PiVoT-Merge-A-7B/README.md) |
| Pivot MoE | Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_M Q4_0 | 32768 | 4096 | [README.md](models/PiVoT-MoE/README.md) |
| Pirouette 7B | Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q8_0 Q3_K_S Q4_0 Q5_0 | 32768 | 4096 | [README.md](models/Pirouette-7B/README.md) |
| PlatYi 34B Llama Q V3 | Q6_K Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_S Q3_K_M Q5_0 Q5_K_M Q8_0 | 4096 | 7168 | [README.md](models/PlatYi-34B-Llama-Q-v3/README.md) |
| Platypus 30B | Q5_0 Q6_K Q8_0 Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q2_K Q3_K_M | 2048 | 6656 | [README.md](models/Platypus-30B/README.md) |
| Platypus2 | Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_0 Q6_K Q8_0 Q2_K | 4096 | 5120 | [README.md](models/Platypus2-13B/README.md) |
| Platypus2 70B | Q5_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/Platypus2-70B/README.md) |
| Platypus2 70B Instruct | Q3_K_L Q3_K_M Q4_K_M Q5_K_S Q2_K Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M | 4096 | 8192 | [README.md](models/Platypus2-70B-Instruct/README.md) |
| Poro 34B | Q3_K_M Q4_K_S Q5_0 Q5_K_M Q6_K Q2_K Q3_K_S Q4_0 Q4_K_M Q5_K_S Q8_0 Q3_K_L | 0 | 7168 | [README.md](models/Poro-34B/README.md) |
| PsyMedRP v1 13B | Q4_K_M Q4_K_S Q5_0 Q5_K_S Q6_K Q2_K Q4_0 Q3_K_S Q5_K_M Q8_0 Q3_K_L Q3_K_M | 4096 | 5120 | [README.md](models/PsyMedRP-v1-13B/README.md) |
| Psymedrp v1 20B | Q3_K_M Q4_K_M Q4_K_S Q6_K Q8_0 Q3_K_L Q3_K_S Q4_0 Q5_0 Q5_K_M Q5_K_S Q2_K | 4096 | 5120 | [README.md](models/PsyMedRP-v1-20B/README.md) |
| Psyfighter 13B | Q4_K_M Q4_K_S Q6_K Q2_K Q3_K_L Q3_K_S Q5_K_M Q5_K_S Q8_0 Q3_K_M Q4_0 Q5_0 | 4096 | 5120 | [README.md](models/Psyfighter-13B/README.md) |
| PuddleJumper 13B | Q3_K_M Q5_K_M Q3_K_L Q5_1 Q6_K Q4_K_M Q5_K_S Q8_0 Q2_K Q3_K_S Q4_0 Q4_1 Q4_K_S Q5_0 | 4096 | 5120 | [README.md](models/PuddleJumper-13B/README.md) |
| PuddleJumper 13B V2 | Q5_K_M Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_S Q6_K Q2_K Q8_0 | 4096 | 5120 | [README.md](models/PuddleJumper-13B-V2/README.md) |
| Pygmalion 2 13B | Q3_K_L Q4_0 Q4_K_S Q5_0 Q5_K_S Q2_K Q3_K_M Q3_K_S Q4_K_M Q5_K_M Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Pygmalion-2-13B/README.md) |
| Pygmalion 2 13B SuperCOT | Q3_K_L Q3_K_M Q3_K_S Q4_K_S Q5_K_M Q5_K_S Q6_K Q2_K Q4_K_M Q5_0 Q8_0 Q4_0 | 4096 | 5120 | [README.md](models/Pygmalion-2-13B-SuperCOT/README.md) |
| Pygmalion 2 13B SuperCOT Weighed | Q5_K_S Q8_0 Q2_K Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q3_K_L Q3_K_M Q4_K_M Q6_K | 4096 | 5120 | [README.md](models/Pygmalion-2-13B-SuperCOT-weighed/README.md) |
| Pygmalion 2 13B SuperCOT2 | Q4_0 Q4_K_M Q5_K_M Q5_K_S Q2_K Q3_K_M Q4_K_S Q5_0 Q6_K Q8_0 Q3_K_L Q3_K_S | 4096 | 5120 | [README.md](models/Pygmalion-2-13B-SuperCOT2/README.md) |
| Pygmalion 2 7B | Q2_K Q3_K_S Q4_0 Q5_0 Q5_K_M Q6_K Q8_0 Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q5_K_S | 4096 | 4096 | [README.md](models/Pygmalion-2-7B/README.md) |
| Python Code 13B | Q3_K_L Q4_0 Q4_K_M Q5_K_M Q6_K Q5_K_S Q8_0 Q2_K Q3_K_M Q3_K_S Q4_K_S Q5_0 | 4096 | 5120 | [README.md](models/Python-Code-13B/README.md) |
| Python Code 33B | Q3_K_M Q3_K_S Q4_K_M Q2_K Q3_K_L Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 2048 | 6656 | [README.md](models/Python-Code-33B/README.md) |
| ReMM SLERP L2 13B | Q2_K Q3_K_M Q3_K_S Q4_K_S Q5_K_S Q6_K Q8_0 Q3_K_L Q4_0 Q4_K_M Q5_0 Q5_K_M | 4096 | 5120 | [README.md](models/ReMM-SLERP-L2-13B/README.md) |
| ReMM v2 L2 13B | Q3_K_M Q4_0 Q4_K_M Q4_K_S Q6_K Q8_0 Q2_K Q3_K_L Q3_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 5120 | [README.md](models/ReMM-v2-L2-13B/README.md) |
| ReMM v2.1 L2 13B | Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q4_K_S Q5_K_S | 4096 | 5120 | [README.md](models/ReMM-v2.1-L2-13B/README.md) |
| Rogue Rose 103B v0.2 | Q2_K Q3_K_M Q3_K_S | 4096 | 8192 | [README.md](models/Rogue-Rose-103b-v0.2/README.md) |
| Rose 20B | Q2_K Q3_K_L Q4_0 Q4_K_S Q5_K_S Q6_K Q3_K_M Q3_K_S Q4_K_M Q5_0 Q5_K_M Q8_0 | 4096 | 5120 | [README.md](models/Rose-20B/README.md) |
| Rose Kimiko 20B | Q5_0 Q6_K Q8_0 Q2_K Q3_K_L Q4_0 Q4_K_M Q4_K_S Q3_K_M Q3_K_S Q5_K_M Q5_K_S | 4096 | 5120 | [README.md](models/Rose-Kimiko-20B/README.md) |
| RpBird Yi 34B 200K | Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q3_K_L Q8_0 Q5_0 Q6_K Q2_K | 200000 | 7168 | [README.md](models/RpBird-Yi-34B-200k/README.md) |
| SAM | Q3_K_M Q4_0 Q6_K Q8_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_S Q4_K_M | 32768 | 4096 | [README.md](models/SAM/README.md) |
| SG Raccoon Yi 200K 2.0 | Q4_K_M Q5_K_M Q6_K Q2_K Q3_K_L Q4_0 Q4_K_S Q5_0 Q5_K_S Q3_K_M Q3_K_S | 200000 | 7168 | [README.md](models/SG-Raccoon-Yi-200k-2.0/README.md) |
| SG Raccoon Yi 55B | Q2_K Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_S Q3_K_L Q3_K_M Q4_K_M Q5_K_M Q6_K | 4096 | 7168 | [README.md](models/SG-Raccoon-Yi-55B/README.md) |
| SG Raccoon Yi 55B 200K | Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q6_K Q3_K_M Q4_K_S | 200000 | 7168 | [README.md](models/SG-Raccoon-Yi-55B-200k/README.md) |
| Solar 10.7B Instruct v1.0 | Q5_K_S Q2_K Q4_K_M Q5_0 Q5_K_M Q4_K_S Q6_K Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_0 | 4096 | 4096 | [README.md](models/SOLAR-10.7B-Instruct-v1.0/README.md) |
| Solar 10.7B Instruct V1.0 Uncensored | Q8_0 Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q2_K Q3_K_M Q4_K_S Q5_K_S | 4096 | 4096 | [README.md](models/SOLAR-10.7B-Instruct-v1.0-uncensored/README.md) |
| Solar 10.7B v1.0 | Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q3_K_S Q3_K_L Q3_K_M Q5_K_S Q6_K Q8_0 Q2_K | 4096 | 4096 | [README.md](models/SOLAR-10.7B-v1.0/README.md) |
| Solar Platypus 10.7B v2 | Q3_K_M Q4_K_S Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_0 Q6_K Q8_0 | 4096 | 4096 | [README.md](models/SOLAR-Platypus-10.7B-v2/README.md) |
| Solarc MOE 10.7Bx4 | Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M | 4096 | 4096 | [README.md](models/SOLARC-MOE-10.7Bx4/README.md) |
| SUS Chat 34B | Q2_K Q3_K_S Q5_K_M Q5_K_S Q6_K Q5_0 Q8_0 Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S | 8192 | 7168 | [README.md](models/SUS-Chat-34B/README.md) |
| Sakura Solar Instruct | Q3_K_S Q4_K_S Q6_K Q2_K Q3_K_L Q3_K_M Q5_K_M Q5_K_S Q8_0 Q4_0 Q4_K_M Q5_0 | 4096 | 4096 | [README.md](models/Sakura-SOLAR-Instruct/README.md) |
| Samantha 1.1 70B | Q5_K_M Q5_K_S Q3_K_M Q3_K_S Q5_0 Q4_K_M Q4_K_S Q2_K Q3_K_L Q4_0 | 4096 | 8192 | [README.md](models/Samantha-1.1-70B/README.md) |
| Samantha 1.11 13B | Q2_K Q3_K_M Q4_0 Q5_0 Q5_K_M Q6_K Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_K_S Q8_0 | 4096 | 5120 | [README.md](models/Samantha-1.11-13B/README.md) |
| Samantha 1.11 70B | Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_K_S Q4_K_M | 4096 | 8192 | [README.md](models/Samantha-1.11-70B/README.md) |
| Samantha 1.11 CodeLlama 34B | Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_S Q3_K_L Q5_K_M Q6_K Q8_0 Q2_K | 2048 | 8192 | [README.md](models/Samantha-1.11-CodeLlama-34B/README.md) |
| Sarah Storyteller 13B | Q4_K_S Q5_0 Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q6_K Q2_K Q4_0 Q5_K_M Q5_K_S | 4096 | 5120 | [README.md](models/Sarah_StoryTeller_13b/README.md) |
| SauerkrautLM 13B v1 | Q2_K Q3_K_L Q3_K_M Q4_K_S Q5_K_M Q6_K Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_S Q8_0 | 4096 | 5120 | [README.md](models/SauerkrautLM-13B-v1/README.md) |
| SauerkrautLM 3B v1 | Q8_0 Q4_0 Q4_1 Q5_0 Q5_1 | 2048 | 3200 | [README.md](models/SauerkrautLM-3B-v1/README.md) |
| SauerkrautLM 70B v1 | Q3_K_L Q4_0 Q4_K_M Q4_K_S Q5_0 Q2_K Q3_K_M Q3_K_S Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/SauerkrautLM-70B-v1/README.md) |
| SauerkrautLM 7B HerO | Q3_K_M Q3_K_S Q4_K_M Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_L Q8_0 Q5_0 Q4_0 Q4_K_S | 32768 | 4096 | [README.md](models/SauerkrautLM-7B-HerO/README.md) |
| SauerkrautLM 7B v1 | Q8_0 Q3_K_L Q4_0 Q5_K_M Q4_K_M Q4_K_S Q5_0 Q5_K_S Q6_K Q2_K Q3_K_M Q3_K_S | 4096 | 4096 | [README.md](models/SauerkrautLM-7B-v1/README.md) |
| SauerkrautLM 7B v1 Mistral | Q3_K_S Q4_K_M Q4_K_S Q6_K Q5_K_S Q8_0 Q2_K Q3_K_L Q3_K_M Q4_0 Q5_0 Q5_K_M | 32768 | 4096 | [README.md](models/SauerkrautLM-7B-v1-mistral/README.md) |
| SauerkrautLM Mixtral 8X7B | Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M | 32768 | 4096 | [README.md](models/SauerkrautLM-Mixtral-8x7B/README.md) |
| SauerkrautLM Mixtral 8X7B Instruct | Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/SauerkrautLM-Mixtral-8x7B-Instruct/README.md) |
| SauerkrautLM SOLAR Instruct | Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q2_K Q3_K_L Q6_K Q8_0 Q5_K_M Q5_K_S | 8192 | 4096 | [README.md](models/SauerkrautLM-SOLAR-Instruct/README.md) |
| SauerkrautLM Una SOLAR Instruct | Q2_K Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_S Q8_0 Q3_K_M Q4_0 Q5_K_M Q6_K | 4096 | 4096 | [README.md](models/SauerkrautLM-UNA-SOLAR-Instruct/README.md) |
| Scarlett 7B | Q4_0 Q4_K_S Q8_0 Q2_K Q3_K_L Q3_K_M Q5_K_M Q5_K_S Q6_K Q3_K_S Q4_K_M Q5_0 | 2048 | 4096 | [README.md](models/Scarlett-7B/README.md) |
| SciPhi Mistral 7B 32K | Q4_K_S Q5_K_M Q6_K Q8_0 Q2_K Q3_K_L Q3_K_S Q4_0 Q3_K_M Q4_K_M Q5_0 Q5_K_S | 32768 | 4096 | [README.md](models/SciPhi-Mistral-7B-32k/README.md) |
| SciPhi Self RAG Mistral 7B 32K | Q6_K Q3_K_L Q3_K_S Q4_0 Q5_0 Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_M Q4_K_M Q4_K_S | 32768 | 4096 | [README.md](models/SciPhi-Self-RAG-Mistral-7B-32k/README.md) |
| Sensei 7B v1 | Q5_K_M Q3_K_L Q3_K_S Q4_K_S Q4_K_M Q5_0 Q5_K_S Q6_K Q8_0 Q2_K Q3_K_M Q4_0 | 32768 | 4096 | [README.md](models/Sensei-7B-V1/README.md) |
| Seraph 7B | Q5_K_M Q5_K_S Q2_K Q3_K_L Q4_0 Q4_K_M Q4_K_S Q5_0 Q8_0 Q3_K_M Q3_K_S Q6_K | 32768 | 4096 | [README.md](models/Seraph-7B/README.md) |
| SeraphMarcoroni 7B | Q5_0 Q5_K_M Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q3_K_S | 32768 | 4096 | [README.md](models/SeraphMarcoroni-7B/README.md) |
| Sheep Duck Llama 2 | Q3_K_M Q3_K_S Q4_K_S Q5_0 Q5_K_S Q2_K Q3_K_L Q5_K_M Q4_0 Q4_K_M | 4096 | 8192 | [README.md](models/Sheep-Duck-Llama-2-70B/README.md) |
| ShiningValiant 1.2 | Q3_K_M Q5_K_S Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S | 4096 | 8192 | [README.md](models/ShiningValiant-1.2/README.md) |
| ShiningValiant 1.3 | Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M Q4_0 Q3_K_S Q4_K_M | 4096 | 8192 | [README.md](models/ShiningValiant-1.3/README.md) |
| ShiningValiantXS 13B | Q3_K_L Q4_K_M Q5_K_S Q6_K Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q8_0 | 4096 | 5120 | [README.md](models/ShiningValiantXS/README.md) |
| ShiningValiantXS 1.1 13B | Q4_K_M Q4_K_S Q8_0 Q3_K_L Q4_0 Q3_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_M | 4096 | 5120 | [README.md](models/ShiningValiantXS-1.1/README.md) |
| Silicon Maid 7B | Q4_0 Q4_K_S Q5_0 Q5_K_M Q6_K Q8_0 Q3_K_L Q3_K_M Q4_K_M Q5_K_S Q2_K Q3_K_S | 32768 | 4096 | [README.md](models/Silicon-Maid-7B/README.md) |
| Skywork 13B Base | Q5_K_M Q5_K_S Q8_0 Q3_K_M Q4_0 Q4_K_M Q4_K_S Q6_K Q2_K Q3_K_L Q3_K_S Q5_0 | 4096 | 4608 | [README.md](models/Skywork-13B-base/README.md) |
| SlimOpenOrca Mistral 7B | Q2_K Q3_K_L Q5_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q6_K Q3_K_M Q3_K_S Q4_0 Q8_0 | 32768 | 4096 | [README.md](models/SlimOpenOrca-Mistral-7B/README.md) |
| SlimOrca 13B | Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_0 Q8_0 Q2_K Q4_K_S Q5_K_M Q5_K_S Q6_K Q3_K_M | 4096 | 5120 | [README.md](models/SlimOrca-13B/README.md) |
| Solar 10.7B SLERP | Q4_K_S Q5_0 Q5_K_S Q6_K Q2_K Q4_0 Q4_K_M Q5_K_M Q8_0 Q3_K_L Q3_K_M Q3_K_S | 4096 | 4096 | [README.md](models/Solar-10.7B-SLERP/README.md) |
| Solus 103B L2 | Q2_K Q3_K_M Q3_K_S | 4096 | 8192 | [README.md](models/Solus-103B-L2/README.md) |
| Solus 70B L2 | Q4_K_S Q5_K_M Q4_K_M Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_0 Q5_K_S Q2_K | 4096 | 8192 | [README.md](models/Solus-70B-L2/README.md) |
| Speechless Llama2 13B | Q3_K_L Q3_K_S Q4_K_M Q5_0 Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_S Q5_K_M Q5_K_S Q6_K | 4096 | 5120 | [README.md](models/Speechless-Llama2-13B/README.md) |
| Speechess Lllama2 Hermes Orca-Platypus WizardLM 13B | Q4_K_M Q4_K_S Q8_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_0 Q5_K_M Q5_K_S Q6_K | 4096 | 5120 | [README.md](models/Speechless-Llama2-Hermes-Orca-Platypus-WizardLM-13B/README.md) |
| Spicyboros 13B 2.2 | Q6_K Q8_0 Q2_K Q3_K_L Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q3_K_M Q3_K_S Q4_0 | 4096 | 5120 | [README.md](models/Spicyboros-13B-2.2/README.md) |
| Spicyboros 70B 2.2 | Q2_K Q3_K_M Q4_K_M Q5_0 Q5_K_M Q3_K_L Q3_K_S Q4_0 Q4_K_S Q5_K_S | 4096 | 8192 | [README.md](models/Spicyboros-70B-2.2/README.md) |
| Spicyboros 7B 2.2 | Q5_0 Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q2_K Q6_K Q8_0 Q5_K_S | 4096 | 4096 | [README.md](models/Spicyboros-7B-2.2/README.md) |
| Spicyboros c34B 2.2 | Q2_K Q3_K_S Q4_K_S Q4_K_M Q5_0 Q5_K_M Q5_K_S Q6_K Q3_K_L Q3_K_M Q4_0 Q8_0 | 16384 | 8192 | [README.md](models/Spicyboros-c34b-2.2/README.md) |
| Spring Dragon | Q6_K Q8_0 Q3_K_L Q3_K_M Q4_K_M Q5_K_M Q5_K_S Q2_K Q3_K_S Q4_0 Q4_K_S Q5_0 | 4096 | 5120 | [README.md](models/Spring-Dragon/README.md) |
| Stable-Platypus2 13B | Q3_K_L Q3_K_S Q4_0 Q5_K_S Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_M Q4_K_M Q4_K_S | 4096 | 5120 | [README.md](models/Stable-Platypus2-13B/README.md) |
| StableBeluga 13B | Q2_K Q3_K_L Q3_K_M Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 | 4096 | 5120 | [README.md](models/StableBeluga-13B/README.md) |
| StableBeluga 7B | Q5_K_S Q6_K Q8_0 Q2_K Q3_K_M Q3_K_S Q4_0 Q5_K_M Q3_K_L Q4_K_M Q4_K_S Q5_0 | 4096 | 4096 | [README.md](models/StableBeluga-7B/README.md) |
| StableBeluga2 | Q3_K_L Q4_K_M Q5_0 Q5_K_S Q5_K_M Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_S | 4096 | 8192 | [README.md](models/StableBeluga2-70B/README.md) |
| Starling LM 7B Alpha | Q2_K Q3_K_M Q4_K_M Q4_K_S Q5_K_M Q3_K_L Q3_K_S Q4_0 Q5_0 Q5_K_S Q6_K Q8_0 | 8192 | 4096 | [README.md](models/Starling-LM-7B-alpha/README.md) |
| Starling LM Alpha 8X7B MoE | Q6_K Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M | 32768 | 4096 | [README.md](models/Starling-LM-alpha-8x7B-MoE/README.md) |
| Stheno Inverted L2 13B | Q3_K_M Q3_K_S Q5_K_M Q8_0 Q5_0 Q5_K_S Q6_K Q2_K Q3_K_L Q4_0 Q4_K_M Q4_K_S | 4096 | 5120 | [README.md](models/Stheno-Inverted-L2-13B/README.md) |
| Stheno L2 13B | Q4_K_S Q5_0 Q6_K Q8_0 Q3_K_M Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_K_M Q5_K_S Q2_K | 4096 | 5120 | [README.md](models/Stheno-L2-13B/README.md) |
| Stheno V2 Delta | Q2_K Q3_K_L Q4_K_S Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 | 4096 | 5120 | [README.md](models/Stheno-v2-Delta/README.md) |
| SuperPlatty 30B | Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_0 Q8_0 Q2_K Q3_K_L Q4_0 Q5_K_M Q5_K_S Q6_K | 2048 | 6656 | [README.md](models/SuperPlatty-30B/README.md) |
| Swallow 13B | Q3_K_S Q5_0 Q5_K_S Q4_K_M Q4_K_S Q5_K_M Q6_K Q2_K Q3_K_L Q3_K_M Q4_0 Q8_0 | 4096 | 5120 | [README.md](models/Swallow-13B/README.md) |
| Swallow 13B Instruct | Q2_K Q4_0 Q4_K_M Q3_K_L Q3_K_M Q3_K_S Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Swallow-13B-Instruct/README.md) |
| Swallow 70B | Q4_0 Q5_K_M Q5_K_S Q2_K Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_0 Q3_K_L | 4096 | 8192 | [README.md](models/Swallow-70B/README.md) |
| Swallow 70B Instruct | Q3_K_L Q4_0 Q4_K_S Q5_K_M Q5_K_S Q2_K Q3_K_M Q3_K_S Q4_K_M Q5_0 | 4096 | 8192 | [README.md](models/Swallow-70B-instruct/README.md) |
| Swallow 7B | Q4_K_S Q5_0 Q5_K_M Q6_K Q2_K Q3_K_L Q3_K_M Q4_0 Q3_K_S Q4_K_M Q5_K_S Q8_0 | 4096 | 4096 | [README.md](models/Swallow-7B/README.md) |
| Swallow 7B Instruct | Q3_K_M Q4_K_M Q5_K_M Q6_K Q8_0 Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_S | 4096 | 4096 | [README.md](models/Swallow-7B-Instruct/README.md) |
| Sydney Overthinker 13B | Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q2_K Q3_K_M Q8_0 Q5_K_S Q6_K | 4096 | 5120 | [README.md](models/Sydney_Overthinker_13B/README.md) |
| Synatra 7B V0.3 RP | Q3_K_S Q4_0 Q6_K Q8_0 Q2_K Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S | 32768 | 4096 | [README.md](models/Synatra-7B-v0.3-RP/README.md) |
| Synatra 7B v0.3 Base | Q3_K_L Q3_K_M Q3_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 Q2_K Q4_0 Q4_K_M Q4_K_S | 32768 | 4096 | [README.md](models/Synatra-7B-v0.3-base/README.md) |
| Synatra 7B V0.3 dpo | Q2_K Q3_K_M Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Synatra-7B-v0.3-dpo/README.md) |
| Synatra RP Orca 2 7B v0.1 | Q6_K Q3_K_M Q4_0 Q4_K_S Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_L Q3_K_S Q4_K_M Q5_0 | 4096 | 4096 | [README.md](models/Synatra-RP-Orca-2-7B-v0.1/README.md) |
| Synatra V0.1 7B Instruct | Q5_K_S Q8_0 Q4_0 Q4_K_M Q3_K_M Q3_K_S Q4_K_S Q5_0 Q5_K_M Q6_K Q2_K Q3_K_L | 32768 | 4096 | [README.md](models/Synatra-V0.1-7B-Instruct/README.md) |
| Synthia 70B v1.5 | Q3_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_M Q4_0 Q4_K_M Q4_K_S Q3_K_L | 2048 | 8192 | [README.md](models/SynthIA-70B-v1.5/README.md) |
| Synthia 7B V1.3 Dare 0.85 | Q4_0 Q5_K_M Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_0 | 32768 | 4096 | [README.md](models/SynthIA-7B-v1.3-dare-0.85/README.md) |
| SynthIA 7B v1.5 | Q4_K_S Q5_K_M Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_S Q6_K Q2_K | 32768 | 4096 | [README.md](models/SynthIA-7B-v1.5/README.md) |
| SynthIA 7B v2.0 | Q5_K_M Q6_K Q8_0 Q3_K_M Q4_0 Q5_0 Q4_K_M Q4_K_S Q5_K_S Q2_K Q3_K_L Q3_K_S | 32768 | 4096 | [README.md](models/SynthIA-7B-v2.0/README.md) |
| SynthIA 7B V2.0 16K | Q3_K_M Q4_K_M Q5_K_M Q5_K_S Q3_K_L Q3_K_S Q4_0 Q4_K_S Q5_0 Q6_K Q8_0 Q2_K | 32768 | 4096 | [README.md](models/SynthIA-7B-v2.0-16k/README.md) |
| Synthia 13B | Q3_K_S Q4_K_M Q5_K_M Q2_K Q3_K_L Q3_K_M Q5_K_S Q6_K Q8_0 Q4_0 Q4_K_S Q5_0 | 4096 | 5120 | [README.md](models/Synthia-13B/README.md) |
| Synthia 13B V1.2 | Q3_K_M Q5_0 Q5_K_S Q8_0 Q4_K_M Q4_K_S Q5_K_M Q6_K Q2_K Q3_K_L Q3_K_S Q4_0 | 4096 | 5120 | [README.md](models/Synthia-13B-v1.2/README.md) |
| Synthia 70B | Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q2_K Q3_K_L | 2048 | 8192 | [README.md](models/Synthia-70B/README.md) |
| Synthia 70B v1.1 | Q2_K Q3_K_L Q3_K_M Q4_K_S Q5_K_S Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M | 2048 | 8192 | [README.md](models/Synthia-70B-v1.1/README.md) |
| Synthia 70B v1.2 | Q4_K_M Q4_K_S Q5_0 Q5_K_S Q2_K Q3_K_M Q3_K_S Q4_0 Q5_K_M Q3_K_L | 2048 | 8192 | [README.md](models/Synthia-70B-v1.2/README.md) |
| Synthia 70B v1.2b | Q5_K_M Q2_K Q4_0 Q4_K_M Q5_0 Q5_K_S Q3_K_L Q3_K_M Q3_K_S Q4_K_S | 4096 | 8192 | [README.md](models/Synthia-70B-v1.2b/README.md) |
| Synthia 7b | Q6_K Q8_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q4_0 Q5_K_S Q2_K Q3_K_L Q3_K_M Q3_K_S | 2048 | 4096 | [README.md](models/Synthia-7B/README.md) |
| Synthia 7B V1.2 | Q5_K_M Q3_K_L Q3_K_M Q4_0 Q4_K_S Q5_0 Q5_K_S Q6_K Q8_0 Q2_K Q3_K_S Q4_K_M | 2048 | 4096 | [README.md](models/Synthia-7B-v1.2/README.md) |
| Synthia 7B V1.3 | Q3_K_S Q4_0 Q4_K_S Q5_0 Q6_K Q3_K_L Q3_K_M Q4_K_M Q5_K_M Q5_K_S Q8_0 Q2_K | 32768 | 4096 | [README.md](models/Synthia-7B-v1.3/README.md) |
| Synthia 7B v3.0 | Q6_K Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_L Q3_K_M | 32768 | 4096 | [README.md](models/Synthia-7B-v3.0/README.md) |
| Synthia MoE v3 Mixtral 8x7B | Q6_K Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M | 32768 | 4096 | [README.md](models/Synthia-MoE-v3-Mixtral-8x7B/README.md) |
| Tai 70B | Q3_K_M Q4_K_M Q4_K_S Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_S Q4_0 Q5_K_S | 4096 | 8192 | [README.md](models/Tai-70B/README.md) |
| Terminis 7B | Q3_K_S Q4_K_M Q4_K_S Q4_0 Q5_0 Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_M Q8_0 | 32768 | 4096 | [README.md](models/Terminis-7B/README.md) |
| Tess 34B v1.4 | Q2_K Q4_K_M Q5_K_S Q6_K Q5_0 Q5_K_M Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_S | 200000 | 7168 | [README.md](models/Tess-34B-v1.4/README.md) |
| Tess 7B V1.4 | Q4_0 Q4_K_M Q5_0 Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_M Q8_0 Q3_K_S Q4_K_S Q5_K_M | 32768 | 4096 | [README.md](models/Tess-7B-v1.4/README.md) |
| Tess M Creative v1.0 | Q4_K_S Q6_K Q2_K Q3_K_M Q4_K_M Q5_0 Q5_K_M Q5_K_S Q8_0 Q3_K_L Q3_K_S Q4_0 | 200000 | 7168 | [README.md](models/Tess-M-Creative-v1.0/README.md) |
| Tess M v1.1 | Q4_K_S Q5_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_K_M Q5_K_S Q6_K Q8_0 | 200000 | 7168 | [README.md](models/Tess-M-v1.1/README.md) |
| Tess M V1.2 | Q2_K Q4_0 Q5_0 Q5_K_M Q5_K_S Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q6_K | 200000 | 7168 | [README.md](models/Tess-M-v1.2/README.md) |
| Tess M v1.3 | Q4_0 Q6_K Q2_K Q3_K_L Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q8_0 Q3_K_M Q3_K_S | 200000 | 7168 | [README.md](models/Tess-M-v1.3/README.md) |
| Tess XL v1.0 | Q2_K | 4096 | 8192 | [README.md](models/Tess-XL-v1.0/README.md) |
| Tess XS Creative V1.0 | Q2_K Q3_K_S Q5_K_M Q6_K Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_S Q8_0 | 32768 | 4096 | [README.md](models/Tess-XS-Creative-v1.0/README.md) |
| Tess XS V1.1 | Q4_0 Q4_K_M Q5_K_M Q5_K_S Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_K_S Q5_0 Q6_K Q2_K | 32768 | 4096 | [README.md](models/Tess-XS-v1.1/README.md) |
| Thespis 13B Alpha V0.7 | Q2_K Q3_K_L Q4_0 Q4_K_M Q5_K_S Q6_K Q8_0 Q3_K_M Q3_K_S Q4_K_S Q5_0 Q5_K_M | 4096 | 5120 | [README.md](models/Thespis-13B-Alpha-v0.7/README.md) |
| Thespis 13B v0.3 | Q2_K Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q3_K_M Q4_0 Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Thespis-13B-v0.3/README.md) |
| Thespis 13B v0.4 | Q5_K_S Q6_K Q2_K Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q8_0 Q3_K_L Q3_K_S Q5_K_M | 4096 | 5120 | [README.md](models/Thespis-13B-v0.4/README.md) |
| Thespis 13B v0.5 | Q5_0 Q5_K_S Q8_0 Q2_K Q3_K_L Q4_K_M Q4_K_S Q6_K Q3_K_M Q3_K_S Q4_0 Q5_K_M | 4096 | 5120 | [README.md](models/Thespis-13B-v0.5/README.md) |
| Thespis 13B v0.6 | Q2_K Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_K_S Q6_K Q3_K_S Q5_0 Q5_K_M Q8_0 | 4096 | 5120 | [README.md](models/Thespis-13B-v0.6/README.md) |
| Thespis Mistral 7B v0.5 | Q3_K_L Q3_K_M Q4_0 Q5_0 Q2_K Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/Thespis-Mistral-7B-v0.5/README.md) |
| Thespis Mistral 7B v0.6 | Q3_K_L Q4_K_M Q5_K_M Q8_0 Q2_K Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_S Q6_K Q3_K_M | 32768 | 4096 | [README.md](models/Thespis-Mistral-7B-v0.6/README.md) |
| Tiamat 7B | Q3_K_S Q6_K Q2_K Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q8_0 | 32768 | 4096 | [README.md](models/Tiamat-7B/README.md) |
| Tiamat 7B 1.1 DPO | Q6_K Q8_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_S Q5_K_M Q4_0 Q4_K_M Q5_0 Q5_K_S | 32768 | 4096 | [README.md](models/Tiamat-7B-1.1-DPO/README.md) |
| Timecrystal L2 13B | Q2_K Q3_K_M Q4_0 Q5_K_S Q6_K Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_M Q8_0 | 4096 | 5120 | [README.md](models/TimeCrystal-L2-13B/README.md) |
| Tinyllama 1.1B 1T Openorca | Q8_0 Q4_K_M Q4_K_S Q5_0 Q6_K Q4_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M Q3_K_S | 2048 | 2048 | [README.md](models/TinyLlama-1.1B-1T-OpenOrca/README.md) |
| TinyLlama 1.1B Chat v0.3 | Q2_K Q3_K_L Q3_K_M Q4_0 Q5_K_S Q6_K Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_M Q8_0 | 2048 | 2048 | [README.md](models/TinyLlama-1.1B-Chat-v0.3/README.md) |
| TinyLlama 1.1B Intermediate Step 480K 1T | Q5_K_S Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_S Q6_K Q8_0 | 2048 | 2048 | [README.md](models/TinyLlama-1.1B-intermediate-step-480k-1T/README.md) |
| TinyLlama 1.1B Intermediate Step 715K 1.5T | Q3_K_S Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q6_K Q8_0 Q2_K | 2048 | 2048 | [README.md](models/TinyLlama-1.1B-intermediate-step-715k-1.5T/README.md) |
| TinyLlama 1.1B Python v0.1 | Q3_K_S Q4_K_S Q5_K_M Q6_K Q8_0 Q3_K_L Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_S Q2_K | 2048 | 2048 | [README.md](models/TinyLlama-1.1B-python-v0.1/README.md) |
| Toppy M 7B | Q8_0 Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q6_K Q2_K Q3_K_L Q4_0 Q5_0 Q5_K_S | 32768 | 4096 | [README.md](models/Toppy-M-7B/README.md) |
| Trion M 7B | Q2_K Q3_K_M Q3_K_S Q4_K_S Q5_0 Q5_K_S Q8_0 Q3_K_L Q4_0 Q4_K_M Q5_K_M Q6_K | 32768 | 4096 | [README.md](models/Trion-M-7B/README.md) |
| Trurl 2 13B | Q8_0 Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_S Q3_K_L Q4_K_M Q5_K_M Q6_K | 4096 | 5120 | [README.md](models/Trurl-2-13B/README.md) |
| Trurl 2 7B | Q4_K_S Q6_K Q8_0 Q3_K_M Q3_K_S Q4_0 Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q4_K_M | 4096 | 4096 | [README.md](models/Trurl-2-7B/README.md) |
| Tulpar 7B v0 | Q4_0 Q4_K_M Q5_0 Q6_K Q3_K_L Q3_K_M Q3_K_S Q4_K_S Q5_K_M Q5_K_S Q8_0 Q2_K | 4096 | 4096 | [README.md](models/Tulpar-7B-v0/README.md) |
| U-Amethyst 20B | Q5_K_S Q6_K Q2_K Q3_K_S Q4_0 Q4_K_M Q5_K_M Q3_K_L Q3_K_M Q4_K_S Q5_0 Q8_0 | 4096 | 5120 | [README.md](models/U-Amethyst-20B/README.md) |
| UltraLM 13B v2.0 | Q3_K_M Q3_K_S Q2_K Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q3_K_L Q8_0 | 4096 | 5120 | [README.md](models/UltraLM-13B-v2.0/README.md) |
| Uncensored Frank 13b | Q4_K_S Q5_K_S Q3_K_M Q3_K_S Q4_0 Q4_K_M Q6_K Q8_0 Q2_K Q3_K_L Q5_0 Q5_K_M | 4096 | 5120 | [README.md](models/Uncensored-Frank-13b/README.md) |
| Uncensored Frank 33b | Q2_K Q3_K_M Q4_0 Q4_K_M Q6_K Q3_K_L Q3_K_S Q4_K_S Q5_0 Q5_K_M Q5_K_S Q8_0 | 2048 | 6656 | [README.md](models/Uncensored-Frank-33b/README.md) |
| Uncensored Frank 7B | Q8_0 Q2_K Q3_K_L Q4_K_M Q4_K_S Q6_K Q5_K_S Q3_K_M Q3_K_S Q4_0 Q5_0 Q5_K_M | 2048 | 4096 | [README.md](models/Uncensored-Frank-7B/README.md) |
| Uncensored Jordan 13B | Q3_K_S Q4_0 Q4_K_M Q5_0 Q2_K Q3_K_L Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_M Q4_K_S | 4096 | 5120 | [README.md](models/Uncensored-Jordan-13B/README.md) |
| Uncensored Jordan 33B | Q4_0 Q4_K_S Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_S Q4_K_M Q5_0 Q6_K Q3_K_L Q3_K_M | 2048 | 6656 | [README.md](models/Uncensored-Jordan-33B/README.md) |
| Uncensored Jordan 7B | Q3_K_M Q4_0 Q5_K_M Q5_K_S Q6_K Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_0 Q8_0 Q2_K | 2048 | 4096 | [README.md](models/Uncensored-Jordan-7B/README.md) |
| UndiMix v1 13B | Q2_K Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_S Q6_K Q3_K_L Q3_K_M Q4_K_S Q5_K_M Q8_0 | 4096 | 5120 | [README.md](models/UndiMix-v1-13B/README.md) |
| UndiMix v2 13B | Q2_K Q3_K_S Q5_K_M Q8_0 Q6_K Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_S | 4096 | 5120 | [README.md](models/UndiMix-v2-13B/README.md) |
| Undimix v3 13B | Q4_K_S Q5_0 Q8_0 Q2_K Q3_K_L Q3_K_S Q4_K_M Q6_K Q3_K_M Q4_0 Q5_K_M Q5_K_S | 4096 | 5120 | [README.md](models/UndiMix-v3-13B/README.md) |
| Undimix v4 13B | Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q8_0 Q3_K_S Q3_K_L Q3_K_M Q4_0 Q5_0 Q6_K Q2_K | 4096 | 5120 | [README.md](models/UndiMix-v4-13B/README.md) |
| Unholy v1 10l 13B | Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q5_K_M Q5_K_S Q8_0 Q2_K Q4_0 Q4_K_S Q5_0 Q6_K | 4096 | 5120 | [README.md](models/Unholy-v1-10l-13B/README.md) |
| Unholy v1 12L 13B | Q2_K Q3_K_L Q4_0 Q5_0 Q5_K_S Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Unholy-v1-12L-13B/README.md) |
| Uni-TianYan (70B) | Q2_K Q4_K_M Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_0 | 4096 | 8192 | [README.md](models/Uni-TianYan-70B/README.md) |
| Llama 2 70B Instruct v2 | Q4_K_S Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M | 4096 | 8192 | [README.md](models/Upstage-Llama-2-70B-instruct-v2/README.md) |
| Llama 65B Instruct | Q2_K Q3_K_M Q4_0 Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_S Q4_K_M Q4_K_S | 2048 | 8192 | [README.md](models/Upstage-Llama1-65B-Instruct/README.md) |
| Utopia 13B | Q2_K Q3_K_M Q3_K_S Q5_0 Q5_K_S Q6_K Q8_0 Q3_K_L Q4_0 Q4_K_M Q4_K_S Q5_K_M | 4096 | 5120 | [README.md](models/Utopia-13B/README.md) |
| UtopiaXL 13B | Q5_K_M Q5_K_S Q6_K Q2_K Q4_K_S Q3_K_S Q4_0 Q4_K_M Q5_0 Q8_0 Q3_K_L Q3_K_M | 4096 | 5120 | [README.md](models/UtopiaXL-13B/README.md) |
| Valkyrie v1 | Q6_K Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_0 Q2_K Q4_K_M Q4_K_S Q5_K_M Q5_K_S | 32768 | 4096 | [README.md](models/Valkyrie-V1/README.md) |
| Velara | Q2_K Q3_K_L Q3_K_M Q4_K_S Q6_K Q8_0 Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S | 32768 | 4096 | [README.md](models/Velara/README.md) |
| Venus 103B v1.1 | Q2_K Q3_K_M Q3_K_S | 4096 | 8192 | [README.md](models/Venus-103b-v1.1/README.md) |
| Venus 103B V1.2 | Q3_K_M Q3_K_S Q2_K | 4096 | 8192 | [README.md](models/Venus-103b-v1.2/README.md) |
| Vicuna 13B CoT | Q3_K_L Q3_K_S Q5_K_M Q5_K_S Q8_0 Q6_K Q2_K Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 | 2048 | 5120 | [README.md](models/Vicuna-13B-CoT/README.md) |
| Vicuna 7B CoT | Q8_0 Q3_K_L Q4_K_S Q5_K_M Q4_0 Q4_K_M Q5_0 Q5_K_S Q6_K Q2_K Q3_K_M Q3_K_S | 2048 | 4096 | [README.md](models/Vicuna-7B-CoT/README.md) |
| Vigogne 2 13B Instruct | Q3_K_S Q4_K_M Q4_K_S Q6_K Q2_K Q3_K_L Q3_K_M Q5_K_S Q8_0 Q4_0 Q5_0 Q5_K_M | 4096 | 5120 | [README.md](models/Vigogne-2-13B-Instruct/README.md) |
| Vigogne 2 7B Chat | Q2_K Q4_K_S Q5_K_M Q6_K Q8_0 Q5_K_S Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 | 4096 | 4096 | [README.md](models/Vigogne-2-7B-Chat/README.md) |
| Vigogne 2 7B Instruct | Q5_K_M Q6_K Q3_K_L Q3_K_S Q4_K_S Q4_K_M Q5_0 Q5_K_S Q8_0 Q2_K Q3_K_M Q4_0 | 2048 | 4096 | [README.md](models/Vigogne-2-7B-Instruct/README.md) |
| Vigostral 7B Chat | Q2_K Q3_K_M Q4_K_M Q5_0 Q5_K_S Q8_0 Q3_K_L Q3_K_S Q4_0 Q4_K_S Q5_K_M Q6_K | 32768 | 4096 | [README.md](models/Vigostral-7B-Chat/README.md) |
| WhiteRabbitNeo 13B | Q5_0 Q5_K_M Q5_K_S Q6_K Q3_K_M Q4_0 Q4_K_M Q4_K_S Q8_0 Q2_K Q3_K_L Q3_K_S | 16384 | 5120 | [README.md](models/WhiteRabbitNeo-13B/README.md) |
| WinterGoddess 1.4X 70B L2 | Q4_0 Q4_K_S Q5_0 Q5_K_S Q2_K Q3_K_M Q3_K_S Q3_K_L Q4_K_M Q5_K_M | 4096 | 8192 | [README.md](models/WinterGoddess-1.4x-70B-L2/README.md) |
| Wizard Vicuna 13B Uncensored | Q4_0 Q4_K_M Q4_K_S Q6_K Q8_0 Q3_K_L Q3_K_M Q3_K_S Q5_K_S Q2_K Q5_0 Q5_K_M | 2048 | 5120 | [README.md](models/Wizard-Vicuna-13B-Uncensored/README.md) |
| Wizard Vicuna 30B Uncensored | Q6_K Q2_K Q3_K_M Q4_K_M Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_S Q4_0 Q4_K_S Q8_0 | 2048 | 6656 | [README.md](models/Wizard-Vicuna-30B-Uncensored/README.md) |
| Wizard Vicuna 7B Uncensored | Q8_0 Q3_K_M Q5_0 Q5_K_S Q6_K Q4_K_M Q4_K_S Q5_K_M Q2_K Q3_K_L Q3_K_S Q4_0 | 2048 | 4096 | [README.md](models/Wizard-Vicuna-7B-Uncensored/README.md) |
| WizardCoder Python 13B V1.0 | Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q8_0 Q2_K Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K | 16384 | 5120 | [README.md](models/WizardCoder-Python-13B-V1.0/README.md) |
| WizardCoder Python 34B V1.0 | Q5_K_M Q6_K Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q8_0 Q4_K_S Q5_0 Q5_K_S | 16384 | 8192 | [README.md](models/WizardCoder-Python-34B-V1.0/README.md) |
| WizardCoder Python 7B V1.0 | Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_M Q6_K Q2_K Q3_K_L Q8_0 Q5_K_S Q3_K_M Q4_0 | 16384 | 4096 | [README.md](models/WizardCoder-Python-7B-V1.0/README.md) |
| WizardLM 1.0 Uncensored CodeLlama 34B | Q3_K_L Q5_0 Q5_K_M Q6_K Q4_K_S Q5_K_S Q8_0 Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M | 16384 | 8192 | [README.md](models/WizardLM-1.0-Uncensored-CodeLlama-34B/README.md) |
| WizardLM 1.0 Uncensored Llama2 13B | Q8_0 Q3_K_L Q3_K_M Q4_K_S Q5_0 Q6_K Q5_K_S Q2_K Q3_K_S Q4_0 Q4_K_M Q5_K_M | 4096 | 5120 | [README.md](models/WizardLM-1.0-Uncensored-Llama2-13B/README.md) |
| WizardLM 13B 1.0 | Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q6_K Q3_K_L Q8_0 Q3_K_S Q5_0 Q2_K | 2048 | 5120 | [README.md](models/WizardLM-13B-1.0/README.md) |
| Wizardlm 13B Uncensored | Q4_K_M Q5_0 Q6_K Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q8_0 Q4_K_S Q5_K_M Q5_K_S | 2048 | 5120 | [README.md](models/WizardLM-13B-Uncensored/README.md) |
| WizardLM 13B V1.0 Uncensored | Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_S Q8_0 Q2_K Q3_K_L Q5_K_M Q6_K Q3_K_M Q4_K_S | 2048 | 5120 | [README.md](models/WizardLM-13B-V1.0-Uncensored/README.md) |
| WizardLM 13B V1.1 | Q3_K_L Q3_K_M Q4_K_M Q5_0 Q5_K_M Q2_K Q3_K_S Q4_0 Q4_K_S Q5_K_S Q6_K Q8_0 | 2048 | 5120 | [README.md](models/WizardLM-13B-V1.1/README.md) |
| WizardLM 13B V1.2 | Q6_K Q2_K Q3_K_M Q4_K_M Q5_K_M Q5_K_S Q8_0 Q3_K_L Q3_K_S Q4_0 Q4_K_S Q5_0 | 4096 | 5120 | [README.md](models/WizardLM-13B-V1.2/README.md) |
| WizardLM 30B v1.0 | Q8_0 Q3_K_L Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_S Q6_K Q2_K Q3_K_S Q4_K_S Q5_K_M | 2048 | 6656 | [README.md](models/WizardLM-30B/README.md) |
| Wizardlm 30B Uncensored | Q4_K_M Q8_0 Q2_K Q3_K_L Q3_K_S Q4_0 Q5_K_S Q6_K Q3_K_M Q4_K_S Q5_0 Q5_K_M | 2048 | 6656 | [README.md](models/WizardLM-30B-Uncensored/README.md) |
| WizardLM 33B V1.0 Uncensored | Q6_K Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q5_K_M Q4_0 Q4_K_S Q5_0 Q5_K_S Q8_0 | 2048 | 6656 | [README.md](models/WizardLM-33B-V1.0-Uncensored/README.md) |
| WizardLM 70B V1.0 | Q4_K_S Q3_K_S Q4_K_M Q3_K_M Q4_0 Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L | 4096 | 8192 | [README.md](models/WizardLM-70B-V1.0/README.md) |
| WizardLM 7B V1.0 Uncensored | Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q8_0 Q3_K_M Q3_K_S Q4_0 Q6_K Q2_K Q3_K_L | 2048 | 4096 | [README.md](models/WizardLM-7B-V1.0-Uncensored/README.md) |
| Wizardlm 7B Uncensored | Q3_K_L Q3_K_S Q5_K_M Q2_K Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_S Q6_K Q8_0 | 2048 | 4096 | [README.md](models/WizardLM-7B-uncensored/README.md) |
| WizardLM Uncensored SuperCOT Storytelling 30B | Q2_K Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q5_K_S Q6_K Q3_K_S Q4_0 Q5_0 Q5_K_M Q8_0 | 2048 | 6656 | [README.md](models/WizardLM-Uncensored-SuperCOT-StoryTelling-30B/README.md) |
| WizardMath 13B V1.0 | Q6_K Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q5_K_S Q2_K Q4_0 Q4_K_S Q5_0 Q5_K_M | 4096 | 5120 | [README.md](models/WizardMath-13B-V1.0/README.md) |
| WizardMath 70B V1.0 | Q2_K Q3_K_M Q3_K_S Q4_K_S Q5_K_M Q5_K_S Q3_K_L Q4_0 Q4_K_M Q5_0 | 2048 | 8192 | [README.md](models/WizardMath-70B-V1.0/README.md) |
| WizardMath 7B V1.0 | Q3_K_M Q4_K_M Q5_K_S Q2_K Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q6_K Q8_0 Q3_K_L | 2048 | 4096 | [README.md](models/WizardMath-7B-V1.0/README.md) |
| WizardMath 7B V1.1 | Q3_K_S Q4_K_S Q5_K_M Q5_K_S Q8_0 Q3_K_L Q3_K_M Q4_0 Q4_K_M Q5_0 Q6_K Q2_K | 32768 | 4096 | [README.md](models/WizardMath-7B-V1.1/README.md) |
| Writing Partner Mistral 7B | Q4_K_M Q4_K_S Q5_0 Q8_0 Q3_K_S Q3_K_L Q3_K_M Q4_0 Q5_K_M Q5_K_S Q6_K Q2_K | 32768 | 4096 | [README.md](models/Writing_Partner_Mistral_7B/README.md) |
| X Mythochronos 13B | Q5_K_M Q6_K Q8_0 Q3_K_L Q3_K_S Q4_0 Q4_K_S Q5_K_S Q2_K Q3_K_M Q4_K_M Q5_0 | 4096 | 5120 | [README.md](models/X-MythoChronos-13B/README.md) |
| X NoroChronos 13B | Q4_K_S Q5_0 Q6_K Q8_0 Q3_K_M Q3_K_S Q4_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q4_K_M | 4096 | 5120 | [README.md](models/X-NoroChronos-13B/README.md) |
| Xwin-LM 13B V0.1 | Q3_K_M Q4_K_S Q5_0 Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_K_M | 4096 | 5120 | [README.md](models/Xwin-LM-13B-V0.1/README.md) |
| Xwin LM 13B v0.2 | Q6_K Q2_K Q3_K_L Q3_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q3_K_S Q4_0 Q4_K_M Q8_0 | 4096 | 5120 | [README.md](models/Xwin-LM-13B-v0.2/README.md) |
| Xwin-LM 70B V0.1 | Q5_K_S Q3_K_M Q4_K_S Q5_K_M Q4_0 Q4_K_M Q5_0 Q2_K Q3_K_L Q3_K_S | 4096 | 8192 | [README.md](models/Xwin-LM-70B-V0.1/README.md) |
| Xwin-LM 7B V0.1 | Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q2_K Q3_K_L Q5_0 Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Xwin-LM-7B-V0.1/README.md) |
| Xwin LM 7B v0.2 | Q6_K Q8_0 Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q2_K Q3_K_S Q5_K_M Q5_K_S | 4096 | 4096 | [README.md](models/Xwin-LM-7B-V0.2/README.md) |
| Xwin MLewd 13B v0.2 | Q2_K Q3_K_L Q4_0 Q4_K_M Q5_K_S Q3_K_M Q3_K_S Q4_K_S Q5_0 Q5_K_M Q6_K Q8_0 | 4096 | 5120 | [README.md](models/Xwin-MLewd-13B-v0.2/README.md) |
| XwinCoder 13B | Q2_K Q3_K_S Q4_0 Q4_K_S Q5_0 Q8_0 Q3_K_L Q3_K_M Q4_K_M Q5_K_M Q5_K_S Q6_K | 16384 | 5120 | [README.md](models/XwinCoder-13B/README.md) |
| XwinCoder 34B | Q2_K Q3_K_S Q4_0 Q4_K_S Q5_K_S Q8_0 Q3_K_L Q3_K_M Q4_K_M Q5_0 Q5_K_M Q6_K | 16384 | 8192 | [README.md](models/XwinCoder-34B/README.md) |
| Yarn Llama 2 13B 128K | Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_L Q4_K_S Q5_0 Q6_K | 131072 | 5120 | [README.md](models/Yarn-Llama-2-13B-128K/README.md) |
| Yarn Llama 2 13B 64K | Q4_0 Q4_K_S Q5_0 Q5_K_M Q8_0 Q2_K Q3_K_M Q4_K_M Q5_K_S Q6_K Q3_K_L Q3_K_S | 65536 | 5120 | [README.md](models/Yarn-Llama-2-13B-64K/README.md) |
| Yarn Llama 2 70B 32K | Q4_0 Q4_K_M Q5_K_M Q2_K Q3_K_L Q4_K_S Q5_0 Q5_K_S Q3_K_M Q3_K_S | 8192 | 8192 | [README.md](models/Yarn-Llama-2-70B-32k/README.md) |
| Yarn Llama 2 7B 128K | Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_L Q3_K_M Q4_0 Q6_K Q3_K_S Q4_K_M Q4_K_S Q5_0 | 131072 | 4096 | [README.md](models/Yarn-Llama-2-7B-128K/README.md) |
| Yarn Llama 2 7B 64K | Q3_K_M Q4_K_S Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q6_K Q2_K Q8_0 | 65536 | 4096 | [README.md](models/Yarn-Llama-2-7B-64K/README.md) |
| Yarn Mistral 7B 128K | Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_L Q3_K_S Q5_0 Q4_K_M Q4_K_S Q2_K Q3_K_M Q4_0 | 32768 | 4096 | [README.md](models/Yarn-Mistral-7B-128k/README.md) |
| Yarn Mistral 7B 64K | Q5_K_S Q8_0 Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_K_M Q6_K Q2_K Q3_K_M Q4_K_S Q5_0 | 32768 | 4096 | [README.md](models/Yarn-Mistral-7B-64k/README.md) |
| Yi 34B | Q4_0 Q4_K_M Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_M Q3_K_S Q4_K_S Q5_0 Q6_K Q3_K_L | 4096 | 7168 | [README.md](models/Yi-34B/README.md) |
| Yi 34B 200K | Q3_K_S Q4_K_S Q5_0 Q5_K_S Q6_K Q3_K_L Q3_K_M Q4_0 Q4_K_M Q5_K_M Q8_0 Q2_K | 200000 | 7168 | [README.md](models/Yi-34B-200K/README.md) |
| Yi 34B 200K Aezakmi v2 | Q5_K_S Q8_0 Q3_K_S Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q2_K Q6_K | 200000 | 7168 | [README.md](models/Yi-34B-200K-AEZAKMI-v2/README.md) |
| Yi 34B 200K DARE Merge v5 | Q4_K_S Q5_K_M Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q3_K_S | 200000 | 7168 | [README.md](models/Yi-34B-200K-DARE-merge-v5/README.md) |
| Yi 34B 200K Llamafied | Q5_0 Q5_K_M Q6_K Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_S Q8_0 Q3_K_L | 200000 | 7168 | [README.md](models/Yi-34B-200K-Llamafied/README.md) |
| Yi 34B Chat | Q3_K_M Q4_0 Q4_K_M Q4_K_S Q8_0 Q2_K Q3_K_L Q5_K_M Q5_K_S Q6_K Q3_K_S Q5_0 | 4096 | 7168 | [README.md](models/Yi-34B-Chat/README.md) |
| Yi 34B GiftedConvo Llama | Q2_K Q4_K_S Q5_K_M Q6_K Q4_K_M Q5_0 Q5_K_S Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_0 | 4096 | 7168 | [README.md](models/Yi-34B-GiftedConvo-merged/README.md) |
| Yi 6B | Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q6_K Q8_0 Q4_0 Q3_K_L Q3_K_M Q3_K_S Q5_0 Q2_K | 4096 | 4096 | [README.md](models/Yi-6B/README.md) |
| Yi 6B 200K | Q4_0 Q4_K_M Q5_K_M Q6_K Q8_0 Q2_K Q3_K_L Q4_K_S Q5_0 Q5_K_S Q3_K_M Q3_K_S | 200000 | 4096 | [README.md](models/Yi-6B-200K/README.md) |
| YuLan Chat 2 13B | Q3_K_M Q3_K_S Q5_K_M Q8_0 Q2_K Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_S Q6_K Q3_K_L | 8192 | 5120 | [README.md](models/YuLan-Chat-2-13B/README.md) |
| Zarablend L2 7B | Q5_K_M Q4_0 Q5_0 Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_S Q6_K Q2_K Q3_K_L Q8_0 | 4096 | 4096 | [README.md](models/Zarablend-L2-7B/README.md) |
| Zarablend MX L2 7B | Q6_K Q2_K Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S Q8_0 Q3_K_L Q3_K_M Q4_K_M | 4096 | 4096 | [README.md](models/Zarablend-MX-L2-7B/README.md) |
| Zaraufsionex 1.1 L2 7B | Q4_0 Q4_K_M Q4_K_S Q5_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q5_K_M Q5_K_S Q6_K Q8_0 | 4096 | 4096 | [README.md](models/Zarafusionex-1.1-L2-7B/README.md) |
| Zephrp m7b | Q2_K Q3_K_M Q5_K_M Q4_K_M Q4_K_S Q5_0 Q5_K_S Q6_K Q3_K_L Q3_K_S Q4_0 Q8_0 | 32768 | 4096 | [README.md](models/ZephRP-m7b/README.md) |
| Ziya Coding 34B v1.0 | Q3_K_L Q4_K_S Q5_0 Q8_0 Q4_K_M Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_M Q3_K_S Q4_0 | 16384 | 8192 | [README.md](models/Ziya-Coding-34B-v1.0/README.md) |
| AgentLM 13B | Q4_0 Q4_K_M Q5_K_M Q5_K_S Q4_K_S Q5_0 Q6_K Q8_0 Q2_K Q3_K_L Q3_K_M Q3_K_S | 4096 | 5120 | [README.md](models/agentlm-13B/README.md) |
| AgentLM 70B | Q4_K_S Q5_0 Q5_K_M Q2_K Q3_K_S Q4_0 Q4_K_M Q3_K_L Q3_K_M Q5_K_S | 4096 | 8192 | [README.md](models/agentlm-70B/README.md) |
| AgentLM 7B | Q3_K_L Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 Q2_K Q3_K_S Q4_0 Q4_K_M Q3_K_M | 4096 | 4096 | [README.md](models/agentlm-7B/README.md) |
| Agiin 13.6B v0.1 | Q2_K Q3_K_L Q4_K_M Q5_K_S Q6_K Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q8_0 | 4096 | 4096 | [README.md](models/agiin-13.6B-v0.1/README.md) |
| Airoboros C34B 2.2.1 | Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q3_K_L Q4_K_S Q5_K_S Q6_K Q8_0 | 16384 | 8192 | [README.md](models/airoboros-c34b-2.2.1/README.md) |
| Airoboros L2 13B 2.2.1 | Q3_K_L Q3_K_M Q4_K_M Q8_0 Q5_0 Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_S Q4_0 Q4_K_S | 4096 | 5120 | [README.md](models/airoboros-l2-13B-2.2.1/README.md) |
| Airoboros L2 13B 3.0 | Q3_K_M Q4_0 Q4_K_S Q5_K_M Q8_0 Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_S Q4_K_M Q5_0 | 4096 | 5120 | [README.md](models/airoboros-l2-13B-3.0/README.md) |
| Airoboros Llama 2 13B GPT4 1.4.1 | Q3_K_M Q3_K_S Q4_0 Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q4_K_M Q4_K_S | 4096 | 5120 | [README.md](models/airoboros-l2-13B-gpt4-1.4.1/README.md) |
| Airoboros L2 13B GPT4 2.0 | Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_K_M Q8_0 Q2_K Q3_K_L Q3_K_S Q5_0 Q5_K_S Q6_K | 4096 | 5120 | [README.md](models/airoboros-l2-13b-gpt4-2.0/README.md) |
| Airoboros L2 13B Gpt4 M2.0 | Q5_K_M Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/airoboros-l2-13b-gpt4-m2.0/README.md) |
| Airoboros L2 70B GPT4 2.0 | Q5_K_S Q4_0 Q4_K_M Q5_K_M Q3_K_S Q4_K_S Q5_0 Q2_K Q3_K_L Q3_K_M | 4096 | 8192 | [README.md](models/airoboros-l2-70B-GPT4-2.0/README.md) |
| Airoboros Llama 2 70B GPT4 1.4.1 | Q4_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M | 4096 | 8192 | [README.md](models/airoboros-l2-70B-gpt4-1.4.1/README.md) |
| Airoboros L2 7B 2.2.1 | Q4_0 Q5_0 Q5_K_M Q6_K Q8_0 Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q5_K_S Q2_K Q3_K_S | 4096 | 4096 | [README.md](models/airoboros-l2-7B-2.2.1/README.md) |
| Airoboros L2 7B 3.0 | Q2_K Q3_K_L Q4_0 Q5_0 Q6_K Q8_0 Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q5_K_S | 4096 | 4096 | [README.md](models/airoboros-l2-7B-3.0/README.md) |
| Airoboros L2 7B Gpt4 2.0 | Q3_K_S Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_L Q3_K_M Q5_K_S Q4_0 Q4_K_M Q4_K_S | 4096 | 4096 | [README.md](models/airoboros-l2-7B-gpt4-2.0/README.md) |
| Airoboros L2 7B Gpt4 M2.0 | Q4_K_M Q4_K_S Q6_K Q8_0 Q2_K Q4_0 Q3_K_S Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_M | 4096 | 4096 | [README.md](models/airoboros-l2-7B-gpt4-m2.0/README.md) |
| Airoboros Llama 2 7B GPT4 1.4.1 | Q2_K Q3_K_M Q4_K_M Q4_K_S Q5_K_S Q6_K Q8_0 Q3_K_L Q3_K_S Q4_0 Q5_0 Q5_K_M | 4096 | 4096 | [README.md](models/airoboros-l2-7b-gpt4-1.4.1/README.md) |
| Airoboros M 7B 3.0 | Q6_K Q3_K_L Q5_0 Q5_K_M Q5_K_S Q4_K_M Q4_K_S Q8_0 Q2_K Q3_K_M Q3_K_S Q4_0 | 32768 | 4096 | [README.md](models/airoboros-m-7B-3.0/README.md) |
| Airoboros M 7B 3.1.2 Dare 0.85 | Q4_K_S Q8_0 Q4_K_M Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_K_S Q6_K | 32768 | 4096 | [README.md](models/airoboros-m-7B-3.1.2-dare-0.85/README.md) |
| Airoboros Mistral2.2 7B | Q6_K Q3_K_L Q4_0 Q4_K_S Q4_K_M Q5_0 Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_M Q3_K_S | 32768 | 4096 | [README.md](models/airoboros-mistral2.2-7B/README.md) |
| Airochronos 33B | Q5_K_S Q6_K Q3_K_L Q3_K_S Q4_0 Q4_K_S Q5_0 Q2_K Q3_K_M Q4_K_M Q5_K_M Q8_0 | 2048 | 6656 | [README.md](models/airochronos-33B/README.md) |
| Alfred 40B 1023 | Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_L Q4_0 Q5_K_S Q3_K_M Q3_K_S Q4_K_S | 0 | 8192 | [README.md](models/alfred-40B-1023/README.md) |
| Apricot Wildflower 20 | Q4_K_M Q4_K_S Q6_K Q8_0 Q2_K Q3_K_S Q4_0 Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_M | 32768 | 4096 | [README.md](models/apricot-wildflower-20/README.md) |
| Bagel 7B v0.1 | Q4_0 Q4_K_S Q5_K_M Q6_K Q4_K_M Q5_0 Q5_K_S Q8_0 Q2_K Q3_K_L Q3_K_M Q3_K_S | 32768 | 4096 | [README.md](models/bagel-7B-v0.1/README.md) |
| Bagel DPO 7B v0.1 | Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_M Q4_0 Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/bagel-dpo-7B-v0.1/README.md) |
| Based 13B | Q8_0 Q3_K_L Q3_K_S Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_M Q4_0 Q4_K_M Q6_K | 2048 | 5120 | [README.md](models/based-13b/README.md) |
| Based 30B | Q3_K_L Q3_K_M Q4_0 Q4_K_S Q5_K_M Q6_K Q2_K Q3_K_S Q4_K_M Q5_0 Q5_K_S Q8_0 | 2048 | 6656 | [README.md](models/based-30B/README.md) |
| Based 7B | Q4_0 Q4_K_S Q5_0 Q5_K_S Q6_K Q3_K_L Q3_K_M Q3_K_S Q8_0 Q2_K Q4_K_M Q5_K_M | 2048 | 4096 | [README.md](models/based-7B/README.md) |
| Basilisk 7B v0.2 | Q4_K_M Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_M Q3_K_S Q8_0 Q4_0 Q4_K_S Q5_0 | 32768 | 4096 | [README.md](models/basilisk-7B-v0.2/README.md) |
| Bling Stable LM 3B 4E1T V0 | Q5_0 Q5_K_S Q3_K_L Q3_K_M Q4_K_S Q4_K_M Q5_K_M Q6_K Q8_0 Q2_K Q3_K_S Q4_0 | 4096 | 2560 | [README.md](models/bling-stable-lm-3b-4e1t-v0/README.md) |
| Blossom V3 Baichuan2 7B | Q4_K_S Q5_K_M Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 | 4096 | 4096 | [README.md](models/blossom-v3-baichuan2-7B/README.md) |
| Blossom V3 Mistral 7B | Q3_K_M Q4_0 Q5_K_M Q8_0 Q6_K Q2_K Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_S | 32768 | 4096 | [README.md](models/blossom-v3-mistral-7B/README.md) |
| Blossom V3_1 Yi 34B | Q3_K_M Q3_K_S Q4_0 Q4_K_M Q6_K Q3_K_L Q4_K_S Q5_0 Q5_K_M Q5_K_S Q8_0 Q2_K | 4096 | 7168 | [README.md](models/blossom-v3_1-yi-34b/README.md) |
| Bun Mistral 7B v2 | Q3_K_M Q4_0 Q5_K_M Q6_K Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_S Q8_0 Q2_K | 32768 | 4096 | [README.md](models/bun_mistral_7b_v2/README.md) |
| Calm2 7B Chat | Q3_K_S Q4_0 Q4_K_M Q5_0 Q6_K Q3_K_L Q3_K_M Q4_K_S Q5_K_M Q5_K_S Q8_0 Q2_K | 32768 | 4096 | [README.md](models/calm2-7B-chat/README.md) |
| Cat v1.0 13B | Q5_0 Q5_K_M Q2_K Q3_K_M Q3_K_S Q4_0 Q6_K Q8_0 Q3_K_L Q4_K_M Q4_K_S Q5_K_S | 4096 | 5120 | [README.md](models/cat-v1.0-13B/README.md) |
| Chronos 13B | Q5_K_M Q3_K_L Q3_K_M Q4_0 Q4_K_S Q5_0 Q5_K_S Q6_K Q8_0 Q2_K Q3_K_S Q4_K_M | 2048 | 5120 | [README.md](models/chronos-13B/README.md) |
| Chronos 13B v2 | Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_S Q4_0 Q4_K_S Q5_0 Q8_0 Q3_K_L Q3_K_M Q4_K_M | 4096 | 5120 | [README.md](models/chronos-13b-v2/README.md) |
| Chronos 33B | Q3_K_L Q3_K_S Q4_0 Q5_0 Q6_K Q8_0 Q2_K Q3_K_M Q4_K_M Q4_K_S Q5_K_M Q5_K_S | 2048 | 6656 | [README.md](models/chronos-33b/README.md) |
| Chronos Hermes 13B | Q2_K Q3_K_L Q3_K_S Q4_K_M Q5_0 Q6_K Q3_K_M Q4_0 Q4_K_S Q5_K_M Q5_K_S Q8_0 | 2048 | 5120 | [README.md](models/chronos-hermes-13B/README.md) |
| Chronos Wizardlm Uc Scot St 13B | Q8_0 Q2_K Q3_K_L Q5_0 Q5_K_S Q6_K Q5_K_M Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S | 2048 | 5120 | [README.md](models/chronos-wizardlm-uc-scot-st-13B/README.md) |
| Chronos007 70B | Q3_K_L Q3_K_M Q4_K_S Q5_0 Q2_K Q4_0 Q4_K_M Q5_K_M Q5_K_S Q3_K_S | 4096 | 8192 | [README.md](models/chronos007-70B/README.md) |
| Cinematika 7B v0.1 | Q2_K Q3_K_L Q4_0 Q4_K_S Q8_0 Q3_K_M Q3_K_S Q4_K_M Q5_0 Q5_K_M Q5_K_S Q6_K | 32768 | 4096 | [README.md](models/cinematika-7B-v0.1/README.md) |
| Claude2 Alpaca 13B | Q4_K_M Q4_K_S Q5_K_M Q6_K Q8_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_0 Q5_K_S | 4096 | 5120 | [README.md](models/claude2-alpaca-13B/README.md) |
| Claude2 Alpaca 7B | Q8_0 Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q3_K_L Q4_K_S Q5_K_S Q6_K | 4096 | 4096 | [README.md](models/claude2-alpaca-7B/README.md) |
| Cockatrice 7B V0.1 | Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q3_K_M Q3_K_S | 32768 | 4096 | [README.md](models/cockatrice-7B-v0.1/README.md) |
| Deepseek Coder 1.3B Base | Q3_K_S Q4_K_S Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_L Q3_K_M Q4_0 Q4_K_M Q5_K_S | 16384 | 2048 | [README.md](models/deepseek-coder-1.3b-base/README.md) |
| Deepseek Coder 1.3B Instruct | Q6_K Q2_K Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_K_M Q3_K_S Q5_0 Q5_K_S Q8_0 | 16384 | 2048 | [README.md](models/deepseek-coder-1.3b-instruct/README.md) |
| Deepseek Coder 33B Base | Q3_K_M Q4_0 Q5_0 Q5_K_M Q8_0 Q6_K Q2_K Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_K_S | 16384 | 7168 | [README.md](models/deepseek-coder-33B-base/README.md) |
| Deepseek Coder 33B Instruct | Q2_K Q3_K_L Q3_K_S Q4_0 Q5_K_M Q6_K Q3_K_M Q4_K_M Q4_K_S Q5_0 Q5_K_S Q8_0 | 16384 | 7168 | [README.md](models/deepseek-coder-33B-instruct/README.md) |
| Deepseek Coder 5.7Bmqa Base | Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_S Q4_0 Q4_K_S Q4_K_M Q2_K Q3_K_L Q3_K_M | 16384 | 4096 | [README.md](models/deepseek-coder-5.7bmqa-base/README.md) |
| Deepseek Coder 6.7B Base | Q2_K Q3_K_L Q5_K_S Q6_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q8_0 | 16384 | 4096 | [README.md](models/deepseek-coder-6.7B-base/README.md) |
| Deepseek Coder 6.7B Instruct | Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_0 Q8_0 Q2_K Q3_K_M Q4_K_S Q5_K_M Q5_K_S Q6_K | 16384 | 4096 | [README.md](models/deepseek-coder-6.7B-instruct/README.md) |
| Deepseek LLM 67B Base | Q4_K_S Q5_0 Q5_K_S Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_K_M Q3_K_L | 4096 | 8192 | [README.md](models/deepseek-llm-67b-base/README.md) |
| Deepseek Llm 67B Chat | Q3_K_L Q3_K_M Q4_K_M Q5_K_M Q5_K_S Q2_K Q3_K_S Q4_0 Q4_K_S Q5_0 | 4096 | 8192 | [README.md](models/deepseek-llm-67b-chat/README.md) |
| Deepseek LLM 7B Base | Q5_0 Q5_K_S Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q6_K Q8_0 Q2_K Q3_K_M Q4_0 | 4096 | 4096 | [README.md](models/deepseek-llm-7B-base/README.md) |
| Deepseek LLM 7B Chat | Q3_K_L Q4_K_M Q4_K_S Q5_0 Q6_K Q2_K Q3_K_M Q3_K_S Q4_0 Q5_K_M Q5_K_S Q8_0 | 4096 | 4096 | [README.md](models/deepseek-llm-7B-chat/README.md) |
| Deepsex 34B | Q3_K_L Q3_K_S Q4_0 Q4_K_M Q2_K Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_M | 4096 | 7168 | [README.md](models/deepsex-34b/README.md) |
| Digital Socrates 13B | Q2_K Q3_K_L Q3_K_S Q4_K_M Q5_0 Q8_0 Q3_K_M Q4_0 Q4_K_S Q5_K_M Q5_K_S Q6_K | 4096 | 5120 | [README.md](models/digital-socrates-13B/README.md) |
| Digital Socrates 7B | Q2_K Q3_K_S Q4_0 Q4_K_S Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q4_K_M Q5_0 Q6_K Q8_0 | 4096 | 4096 | [README.md](models/digital-socrates-7B/README.md) |
| DocsGPT 7B Mistral | Q2_K Q3_K_L Q3_K_S Q4_K_M Q5_0 Q5_K_S Q6_K Q3_K_M Q4_0 Q4_K_S Q5_K_M Q8_0 | 32768 | 4096 | [README.md](models/docsgpt-7B-mistral/README.md) |
| Dolphin 2.0 Mistral 7B | Q4_K_S Q5_0 Q6_K Q8_0 Q4_K_M Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 | 32768 | 4096 | [README.md](models/dolphin-2.0-mistral-7B/README.md) |
| Dolphin 2.1 Mistral 7B | Q5_K_S Q8_0 Q2_K Q3_K_L Q4_0 Q4_K_S Q5_0 Q5_K_M Q3_K_M Q3_K_S Q4_K_M Q6_K | 32768 | 4096 | [README.md](models/dolphin-2.1-mistral-7B/README.md) |
| Dolphin 2.2 70B | Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_M Q4_0 Q3_K_L Q3_K_S | 4096 | 8192 | [README.md](models/dolphin-2.2-70B/README.md) |
| Dolphin 2.2 Yi 34B 200K | Q4_0 Q5_K_M Q6_K Q3_K_M Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_S Q8_0 Q2_K | 200000 | 7168 | [README.md](models/dolphin-2.2-yi-34b-200k/README.md) |
| Dolphin 2.2.1 AshhLimaRP Mistral 7B | Q8_0 Q3_K_L Q3_K_M Q4_K_M Q5_0 Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_S Q4_0 Q4_K_S | 32768 | 4096 | [README.md](models/dolphin-2.2.1-AshhLimaRP-Mistral-7B/README.md) |
| Dolphin 2.2.1 Mistral 7B | Q4_K_M Q2_K Q3_K_L Q3_K_M Q4_0 Q5_K_S Q6_K Q8_0 Q3_K_S Q4_K_S Q5_0 Q5_K_M | 32768 | 4096 | [README.md](models/dolphin-2.2.1-mistral-7B/README.md) |
| Dolphin 2.5 Mixtral 8X7B | Q6_K Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M | 32768 | 4096 | [README.md](models/dolphin-2.5-mixtral-8x7b/README.md) |
| Dolphin 2.6 Mistral 7B | Q5_K_M Q6_K Q8_0 Q4_K_S Q5_0 Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_K_S Q2_K Q3_K_L | 32768 | 4096 | [README.md](models/dolphin-2.6-mistral-7B/README.md) |
| Dolphin 2.6 Mixtral 8X7B | Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K | 32768 | 4096 | [README.md](models/dolphin-2.6-mixtral-8x7b/README.md) |
| Dolphin 2.2 Yi 34B | Q5_K_M Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q4_K_S Q5_0 Q3_K_M Q3_K_S Q4_0 Q4_K_M | 16384 | 7168 | [README.md](models/dolphin-2_2-yi-34b/README.md) |
| Dolphin 2.6 Phi 2 | Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q2_K Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_S | 0 | 0 | [README.md](models/dolphin-2_6-phi-2/README.md) |
| Dopeystableplats 3B v1 | Q5_0 Q5_K_M Q5_K_S Q6_K Q3_K_L Q4_0 Q4_K_M Q4_K_S Q2_K Q3_K_M Q3_K_S Q8_0 | 4096 | 2560 | [README.md](models/dopeystableplats-3b-v1/README.md) |
| Dragon Mistral 7B V0 | Q4_K_M Q4_K_S Q2_K Q3_K_L Q3_K_S Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_M Q4_0 Q5_0 | 32768 | 4096 | [README.md](models/dragon-mistral-7B-v0/README.md) |
| Dragon Yi 6B v0 | Q5_K_S Q6_K Q8_0 Q2_K Q4_0 Q5_0 Q5_K_M Q4_K_S Q3_K_L Q3_K_M Q3_K_S Q4_K_M | 4096 | 4096 | [README.md](models/dragon-yi-6B-v0/README.md) |
| Echidna TieFighter 25 | Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q8_0 Q3_K_L Q3_K_M Q5_K_S Q6_K Q2_K Q5_0 | 4096 | 5120 | [README.md](models/echidna-tiefigther-25/README.md) |
| EM German 13B v01 | Q5_K_M Q6_K Q2_K Q3_K_M Q4_0 Q4_K_S Q5_0 Q5_K_S Q8_0 Q3_K_L Q3_K_S Q4_K_M | 4096 | 5120 | [README.md](models/em_german_13b_v01/README.md) |
| EM German 70B v01 | Q2_K Q3_K_M Q4_0 Q4_K_S Q5_0 Q3_K_L Q3_K_S Q4_K_M Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/em_german_70b_v01/README.md) |
| EM German 7B v01 | Q3_K_M Q3_K_S Q4_K_S Q5_K_M Q6_K Q3_K_L Q4_0 Q4_K_M Q5_0 Q5_K_S Q8_0 Q2_K | 4096 | 4096 | [README.md](models/em_german_7b_v01/README.md) |
| EM German Leo Mistral | Q2_K Q3_K_L Q4_0 Q4_K_S Q8_0 Q6_K Q3_K_M Q3_K_S Q4_K_M Q5_0 Q5_K_M Q5_K_S | 32768 | 4096 | [README.md](models/em_german_leo_mistral/README.md) |
| EM German Mistral v01 | Q4_0 Q5_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q6_K Q8_0 | 32768 | 4096 | [README.md](models/em_german_mistral_v01/README.md) |
| EvolvedSeeker 1.3 | Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_S Q4_K_S Q6_K Q3_K_L Q3_K_M | 16384 | 2048 | [README.md](models/evolvedSeeker_1_3/README.md) |
| Fin Llama 33B | Q8_0 Q2_K Q3_K_L Q5_K_S Q6_K Q4_K_S Q5_0 Q5_K_M Q3_K_M Q3_K_S Q4_0 Q4_K_M | 2048 | 6656 | [README.md](models/fin-llama-33B/README.md) |
| Finance LLM | Q3_K_L Q3_K_M Q4_0 Q4_K_S Q5_0 Q5_K_M Q2_K Q4_K_M Q5_K_S Q6_K Q8_0 Q3_K_S | 0 | 4096 | [README.md](models/finance-LLM/README.md) |
| Firefly Llama2 13B Chat | Q3_K_L Q3_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_M Q4_0 Q4_K_M Q4_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/firefly-llama2-13B-chat/README.md) |
| Firefly Llama2 7B Chat | Q5_K_S Q2_K Q3_K_M Q4_0 Q4_K_M Q5_K_M Q6_K Q8_0 Q3_K_L Q3_K_S Q4_K_S Q5_0 | 4096 | 4096 | [README.md](models/firefly-llama2-7B-chat/README.md) |
| Firefly Mixtral 8X7B | Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K | 32768 | 4096 | [README.md](models/firefly-mixtral-8x7b/README.md) |
| GenZ 13B v2 | Q4_K_M Q6_K Q8_0 Q2_K Q3_K_M Q4_0 Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_S Q4_K_S | 4096 | 5120 | [README.md](models/genz-13B-v2/README.md) |
| Go Bruins | Q3_K_S Q4_K_M Q5_K_M Q2_K Q3_K_M Q4_K_S Q5_0 Q5_K_S Q6_K Q8_0 Q3_K_L Q4_0 | 32768 | 4096 | [README.md](models/go-bruins/README.md) |
| Go Bruins v2 | Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q6_K Q3_K_M Q8_0 Q3_K_L Q4_K_S Q2_K | 32768 | 4096 | [README.md](models/go-bruins-v2/README.md) |
| Go Bruins v2.1.1 | Q5_0 Q8_0 Q2_K Q3_K_M Q4_0 Q4_K_S Q5_K_S Q6_K Q3_K_L Q3_K_S Q4_K_M Q5_K_M | 32768 | 4096 | [README.md](models/go-bruins-v2.1.1/README.md) |
| Goliath 120B | Q2_K | 4096 | 8192 | [README.md](models/goliath-120b/README.md) |
| Gorilla 7B | Q8_0 Q3_K_S Q4_0 Q5_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_M | 2048 | 4096 | [README.md](models/gorilla-7B/README.md) |
| Gorilla OpenFunctions V1 | Q5_K_S Q2_K Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q8_0 Q3_K_M Q4_0 Q5_0 Q5_K_M Q6_K | 4096 | 4096 | [README.md](models/gorilla-openfunctions-v1/README.md) |
| Hippogriff 30B Chat | Q4_0 Q4_K_M Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_S Q5_0 Q6_K Q8_0 | 2048 | 6656 | [README.md](models/hippogriff-30b-chat/README.md) |
| Huginn v1.2 | Q3_K_M Q5_0 Q5_K_M Q6_K Q4_K_M Q4_K_S Q5_K_S Q8_0 Q2_K Q3_K_L Q3_K_S Q4_0 | 4096 | 5120 | [README.md](models/huginnv1.2/README.md) |
| Jackalope 7B | Q5_0 Q2_K Q3_K_M Q4_0 Q4_K_S Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_L Q3_K_S Q4_K_M | 32768 | 4096 | [README.md](models/jackalope-7B/README.md) |
| Japanese StableLM Instruct Gamma 7B | Q4_0 Q4_K_M Q4_K_S Q3_K_M Q3_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L | 4096 | 4096 | [README.md](models/japanese-stablelm-instruct-gamma-7B/README.md) |
| Juanako 7B UNA | Q2_K Q3_K_S Q4_0 Q4_K_S Q5_K_S Q6_K Q3_K_L Q3_K_M Q4_K_M Q5_0 Q5_K_M Q8_0 | 32768 | 4096 | [README.md](models/juanako-7B-UNA/README.md) |
| Juanako 7B V1 | Q5_0 Q5_K_M Q5_K_S Q6_K Q3_K_M Q3_K_S Q4_K_S Q4_K_M Q8_0 Q2_K Q3_K_L Q4_0 | 32768 | 4096 | [README.md](models/juanako-7B-v1/README.md) |
| koOpenChat sft 7B | Q6_K Q8_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S | 8192 | 4096 | [README.md](models/koOpenChat-sft/README.md) |
| Law LLM | Q6_K Q8_0 Q2_K Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q3_K_S Q4_0 | 0 | 4096 | [README.md](models/law-LLM/README.md) |
| Leo Hessianai 13B | Q5_K_M Q2_K Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_S Q6_K Q8_0 Q3_K_L Q3_K_M Q3_K_S | 8192 | 5120 | [README.md](models/leo-hessianai-13B/README.md) |
| Leo Hessianai 13B Chat | Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_S Q4_K_M Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 | 8192 | 5120 | [README.md](models/leo-hessianai-13B-chat/README.md) |
| Leo Hessianai 13B Chat Bilingual | Q6_K Q2_K Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_M Q4_0 Q8_0 | 8192 | 5120 | [README.md](models/leo-hessianai-13B-chat-bilingual/README.md) |
| Leo Hessianai 70B | Q3_K_M Q3_K_S Q4_K_S Q5_K_M Q5_K_S Q2_K Q3_K_L Q4_0 Q4_K_M Q5_0 | 8192 | 8192 | [README.md](models/leo-hessianai-70B/README.md) |
| Leo Hessianai 70B Chat | Q2_K Q3_K_L Q3_K_M Q4_0 Q4_K_S Q5_K_S Q3_K_S Q4_K_M Q5_0 Q5_K_M | 8192 | 8192 | [README.md](models/leo-hessianai-70B-chat/README.md) |
| Leo Hessianai 7B | Q3_K_M Q4_K_M Q4_K_S Q5_0 Q5_K_M Q6_K Q2_K Q3_K_L Q5_K_S Q8_0 Q3_K_S Q4_0 | 8192 | 4096 | [README.md](models/leo-hessianai-7B/README.md) |
| Leo Hessianai 7B Chat | Q8_0 Q3_K_L Q3_K_M Q3_K_S Q5_0 Q5_K_M Q6_K Q2_K Q4_0 Q4_K_M Q4_K_S Q5_K_S | 8192 | 4096 | [README.md](models/leo-hessianai-7B-chat/README.md) |
| Leo Hessianai 7B Chat Bilingual | Q4_0 Q4_K_S Q5_K_M Q5_K_S Q8_0 Q3_K_L Q3_K_M Q3_K_S Q6_K Q2_K Q4_K_M Q5_0 | 8192 | 4096 | [README.md](models/leo-hessianai-7B-chat-bilingual/README.md) |
| Lince Zero | Q8_0 Q4_0 Q4_1 Q5_0 Q5_1 | 0 | 4544 | [README.md](models/lince-zero/README.md) |
| Llama 13B Supercot | Q5_0 Q5_K_S Q8_0 Q3_K_S Q4_K_M Q4_K_S Q4_0 Q5_K_M Q6_K Q2_K Q3_K_L Q3_K_M | 2048 | 5120 | [README.md](models/llama-13b-supercot/README.md) |
| Llama 2 13B German Assistant v2 | Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_S Q8_0 Q2_K Q3_K_L Q3_K_S Q5_K_M Q6_K | 2048 | 5120 | [README.md](models/llama-2-13B-German-Assistant-v2/README.md) |
| Llama2 13B Guanaco QLoRA | Q5_0 Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q3_K_S | 2048 | 5120 | [README.md](models/llama-2-13B-Guanaco-QLoRA/README.md) |
| Llama 2 13B Chat - LimaRP v2 Merged | Q4_K_M Q8_0 Q2_K Q3_K_S Q4_0 Q5_0 Q5_K_M Q5_K_S Q6_K Q3_K_L Q3_K_M Q4_K_S | 4096 | 5120 | [README.md](models/llama-2-13B-chat-limarp-v2-merged/README.md) |
| Llama 2 7B Arguments | Q3_K_M Q4_0 Q4_K_S Q6_K Q8_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_S Q4_K_M Q5_0 | 4096 | 4096 | [README.md](models/llama-2-7B-Arguments/README.md) |
| Llama2 7B Guanaco QLoRA | Q3_K_L Q3_K_S Q5_0 Q5_K_S Q6_K Q2_K Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_K_M Q8_0 | 2048 | 4096 | [README.md](models/llama-2-7B-Guanaco-QLoRA/README.md) |
| Llama 30B Supercot | Q3_K_L Q3_K_M Q5_0 Q6_K Q4_K_S Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_S Q4_0 Q4_K_M | 2048 | 6656 | [README.md](models/llama-30b-supercot/README.md) |
| Llama Polyglot 13B | Q2_K Q3_K_M Q3_K_S Q4_K_S Q5_0 Q5_K_M Q3_K_L Q4_0 Q4_K_M Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/llama-polyglot-13B/README.md) |
| Llama2 22B Daydreamer v2 | Q3_K_L Q4_K_S Q5_0 Q5_K_S Q6_K Q8_0 Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_K_M | 4098 | 6656 | [README.md](models/llama2-22B-daydreamer-v2/README.md) |
| Llama2 7B Layla | Q2_K Q4_0 Q4_K_M Q4_K_S Q8_0 Q3_K_L Q3_K_M Q3_K_S Q5_0 Q5_K_M Q5_K_S Q6_K | 4096 | 4096 | [README.md](models/llama2-7B-layla/README.md) |
| Llama-2-7B-Chat Code Cherry Pop | Q8_0 Q4_K_M Q4_K_S Q5_K_S Q3_K_S Q4_0 Q5_0 Q5_K_M Q6_K Q2_K Q3_K_L Q3_K_M | 4096 | 4096 | [README.md](models/llama2-7b-chat-codeCherryPop-qLoRA/README.md) |
| Llama2 70B Chat Uncensored | Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_K_S Q2_K Q4_K_S Q5_0 Q5_K_M | 2048 | 8192 | [README.md](models/llama2_70b_chat_uncensored/README.md) |
| Llama2 7B Chat Uncensored | Q6_K Q2_K Q4_K_M Q5_0 Q5_K_S Q4_K_S Q5_K_M Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_0 | 2048 | 4096 | [README.md](models/llama2_7b_chat_uncensored/README.md) |
| Llama2 7B Merge Orcafamily | Q3_K_L Q4_0 Q5_K_M Q8_0 Q5_K_S Q6_K Q2_K Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_0 | 2048 | 4096 | [README.md](models/llama2_7b_merge_orcafamily/README.md) |
| Llemma 34B | Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q6_K Q8_0 Q3_K_M Q4_K_M | 4096 | 8192 | [README.md](models/llemma_34b/README.md) |
| Llemma 7B | Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q2_K Q5_0 Q6_K Q8_0 Q3_K_M | 4096 | 4096 | [README.md](models/llemma_7b/README.md) |
| Loyal Piano M7 | Q5_K_M Q5_K_S Q8_0 Q2_K Q4_0 Q4_K_M Q4_K_S Q5_0 Q3_K_L Q3_K_M Q3_K_S Q6_K | 32768 | 4096 | [README.md](models/loyal-piano-m7/README.md) |
| Lzlv 70B | Q3_K_L Q3_K_S Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_M Q4_0 Q4_K_M | 4096 | 8192 | [README.md](models/lzlv_70B/README.md) |
| Manticore 13B Chat Pyg | Q3_K_L Q4_K_M Q5_0 Q5_K_M Q8_0 Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_K_S Q6_K | 2048 | 5120 | [README.md](models/manticore-13b-chat-pyg/README.md) |
| Med42 70B | Q5_K_S Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_M | 2048 | 8192 | [README.md](models/med42-70B/README.md) |
| Medalpaca 13B | Q5_K_S Q6_K Q3_K_M Q3_K_S Q5_K_M Q4_K_M Q4_K_S Q5_0 Q8_0 Q2_K Q3_K_L Q4_0 | 0 | 5120 | [README.md](models/medalpaca-13B/README.md) |
| Meditron 70B | Q3_K_S Q4_0 Q4_K_S Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M Q4_K_M Q5_0 | 4096 | 8192 | [README.md](models/meditron-70B/README.md) |
| Meditron 7B | Q4_K_M Q5_0 Q2_K Q3_K_L Q3_K_M Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_S Q4_0 Q4_K_S | 2048 | 4096 | [README.md](models/meditron-7B/README.md) |
| Meditron 7B Chat | Q8_0 Q2_K Q4_K_M Q5_K_M Q6_K Q4_K_S Q5_0 Q5_K_S Q3_K_L Q3_K_M Q3_K_S Q4_0 | 2048 | 4096 | [README.md](models/meditron-7B-chat/README.md) |
| Merlyn Education Corpus QA v2 | Q4_0 Q5_0 Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_S Q8_0 Q3_K_M Q4_K_M Q4_K_S | 4096 | 5120 | [README.md](models/merlyn-education-corpus-qa-v2/README.md) |
| Merlyn Education Safety 12B | Q5_0 Q5_K_M Q8_0 Q2_K Q3_K_M Q4_K_S Q4_K_M Q5_K_S Q6_K Q3_K_L Q3_K_S Q4_0 | 2048 | 5120 | [README.md](models/merlyn-education-safety/README.md) |
| Mindy 7B | Q3_K_S Q4_K_M Q5_K_M Q8_0 Q2_K Q3_K_L Q3_K_M Q5_K_S Q6_K Q4_0 Q4_K_S Q5_0 | 32768 | 4096 | [README.md](models/mindy-7B/README.md) |
| Minotaur 13B Fixed | Q3_K_M Q3_K_S Q6_K Q8_0 Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q4_0 Q4_K_M Q4_K_S | 2048 | 5120 | [README.md](models/minotaur-13B-fixed/README.md) |
| Mistral 7B DPO V5 | Q3_K_L Q3_K_M Q4_K_M Q5_K_S Q6_K Q8_0 Q2_K Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M | 32768 | 4096 | [README.md](models/mistral-7B-dpo-v5/README.md) |
| Mistral 7B Finetuned Orca DPO V2 | Q3_K_S Q4_K_M Q5_0 Q5_K_S Q8_0 Q2_K Q3_K_M Q4_K_S Q5_K_M Q6_K Q3_K_L Q4_0 | 32768 | 4096 | [README.md](models/mistral-7B-finetuned-orca-dpo-v2/README.md) |
| Mistral FT Optimized 1218 | Q3_K_L Q5_0 Q6_K Q5_K_M Q5_K_S Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q8_0 | 32768 | 4096 | [README.md](models/mistral-ft-optimized-1218/README.md) |
| Mixtralnt 4X7B Test | Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_M Q4_0 | 32768 | 4096 | [README.md](models/mixtralnt-4x7b-test/README.md) |
| Model 007 70B | Q2_K Q3_K_L Q4_0 Q5_0 Q5_K_M Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_S | 4096 | 8192 | [README.md](models/model_007-70B/README.md) |
| Neural Chat 7B v3-1 | Q6_K Q2_K Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q3_K_L Q3_K_S Q5_K_M Q5_K_S Q8_0 | 32768 | 4096 | [README.md](models/neural-chat-7B-v3-1/README.md) |
| Neural Chat 7B V3-2 | Q5_K_S Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q2_K Q3_K_M Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/neural-chat-7B-v3-2/README.md) |
| Neural Chat 7B V3-3 | Q6_K Q8_0 Q3_K_S Q4_0 Q4_K_M Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M Q4_K_S Q5_0 | 32768 | 4096 | [README.md](models/neural-chat-7B-v3-3/README.md) |
| Neural Chat 7B V3.3 WizardMath DARE ME | Q8_0 Q2_K Q3_K_M Q4_K_M Q6_K Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_S Q4_0 Q4_K_S | 32768 | 4096 | [README.md](models/neural-chat-7B-v3-3-wizardmath-dare-me/README.md) |
| Notus 7B v1 | Q2_K Q3_K_L Q3_K_M Q4_K_S Q5_K_M Q8_0 Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_S Q6_K | 32768 | 4096 | [README.md](models/notus-7B-v1/README.md) |
| Notux 8X7B v1 | Q4_K_S Q5_0 Q5_K_S Q8_0 Q3_K_L Q3_K_S Q4_K_M Q5_K_M Q6_K Q2_K Q3_K_M Q4_0 | 32768 | 4096 | [README.md](models/notux-8x7b-v1/README.md) |
| NSQL Llama-2 7B | Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_S Q8_0 Q2_K Q5_0 Q5_K_M Q6_K Q4_0 | 4096 | 4096 | [README.md](models/nsql-llama-2-7B/README.md) |
| Nucleus 22B Token 500B | Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q2_K Q5_K_S Q6_K Q8_0 Q3_K_S | 2048 | 6656 | [README.md](models/nucleus-22B-token-500B/README.md) |
| Open Instruct Human Mix 65B | Q4_0 Q4_K_M Q2_K Q3_K_L Q3_K_S Q5_K_M Q5_K_S Q3_K_M Q4_K_S Q5_0 | 2048 | 8192 | [README.md](models/open-instruct-human-mix-65B/README.md) |
| Open Llama 3B V2 Wizard Evol Instuct V2 196K | Q2_K Q3_K_L Q4_K_S Q5_K_S Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 2048 | 3200 | [README.md](models/open-llama-3b-v2-wizard-evol-instuct-v2-196k/README.md) |
| OpenBuddy Coder 34B V11 | Q5_0 Q5_K_M Q6_K Q2_K Q3_K_L Q3_K_M Q4_0 Q4_K_S Q3_K_S Q4_K_M Q5_K_S Q8_0 | 16384 | 8192 | [README.md](models/openbuddy-coder-34b-v11-bf16/README.md) |
| Openbuddy Deepseek 67B V15 Base | Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_L Q3_K_M Q4_K_S | 4096 | 8192 | [README.md](models/openbuddy-deepseek-67b-v15-base/README.md) |
| OpenBuddy Llama2 13B64K V15 | Q2_K Q4_K_S Q5_K_S Q5_K_M Q6_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 Q8_0 | 65536 | 5120 | [README.md](models/openbuddy-llama2-13b64k-v15/README.md) |
| OpenBuddy Llama2 34B V11.1 | Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q5_K_S Q6_K Q2_K Q4_0 Q4_K_S Q5_0 Q5_K_M | 16384 | 8192 | [README.md](models/openbuddy-llama2-34b-v11.1-bf16/README.md) |
| OpenBuddy Llama2 70B v13 Base | Q4_0 Q4_K_S Q5_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/openbuddy-llama2-70B-v13-base/README.md) |
| OpenBuddy Llama2 70B v13.2 | Q4_K_M Q2_K Q3_K_M Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S Q3_K_L Q3_K_S | 4096 | 8192 | [README.md](models/openbuddy-llama2-70B-v13.2/README.md) |
| Openbuddy Mistral 7B v13 | Q5_0 Q6_K Q8_0 Q2_K Q3_K_S Q4_0 Q4_K_S Q5_K_S Q3_K_L Q3_K_M Q4_K_M Q5_K_M | 32768 | 4096 | [README.md](models/openbuddy-mistral-7B-v13/README.md) |
| OpenBuddy Mistral 7B v13 Base | Q3_K_M Q4_K_M Q4_K_S Q5_0 Q5_K_S Q8_0 Q2_K Q3_K_L Q5_K_M Q6_K Q3_K_S Q4_0 | 32768 | 4096 | [README.md](models/openbuddy-mistral-7B-v13-base/README.md) |
| OpenBuddy Mistral 7B v13.1 | Q4_K_M Q5_0 Q5_K_S Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_K_M Q6_K Q8_0 Q3_K_L | 32768 | 4096 | [README.md](models/openbuddy-mistral-7B-v13.1/README.md) |
| Openbuddy Mixtral 8X7B V15.2 | Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_M | 32768 | 4096 | [README.md](models/openbuddy-mixtral-8x7b-v15.2/README.md) |
| OpenBuddy Mixtral 8X7B V15.4 | Q2_K Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 | 32768 | 4096 | [README.md](models/openbuddy-mixtral-8x7b-v15.4/README.md) |
| OpenBuddy OpenLlama 7B v12 | Q2_K Q3_K_S Q4_K_S Q5_K_M Q5_K_S Q8_0 Q3_K_L Q3_K_M Q4_0 Q4_K_M Q5_0 Q6_K | 4096 | 4096 | [README.md](models/openbuddy-openllama-7B-v12-bf16/README.md) |
| Openbuddy Zephyr 7B v14.1 | Q6_K Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q8_0 Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S | 32768 | 4096 | [README.md](models/openbuddy-zephyr-7B-v14.1/README.md) |
| Openchat 3.5 1210 | Q4_K_M Q5_0 Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_S Q4_0 Q6_K Q3_K_L Q3_K_M Q4_K_S | 8192 | 4096 | [README.md](models/openchat-3.5-1210/README.md) |
| OpenChat 3.5 1210 Starling SLERP | Q3_K_M Q4_0 Q4_K_S Q5_0 Q2_K Q3_K_S Q4_K_M Q5_K_M Q5_K_S Q6_K Q8_0 Q3_K_L | 8192 | 4096 | [README.md](models/openchat-3.5-1210-starling-slerp/README.md) |
| OpenChat 3.5 7B | Q4_K_S Q5_K_M Q6_K Q4_0 Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q5_0 Q5_K_S Q8_0 Q2_K | 8192 | 4096 | [README.md](models/openchat_3.5/README.md) |
| Openchat 3.5 16K | Q3_K_M Q4_0 Q4_K_M Q4_K_S Q8_0 Q2_K Q3_K_L Q3_K_S Q5_0 Q5_K_M Q5_K_S Q6_K | 32768 | 4096 | [README.md](models/openchat_3.5-16k/README.md) |
| OpenChat v3.2 Super | Q3_K_L Q4_K_M Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_0 Q6_K | 4096 | 5120 | [README.md](models/openchat_v3.2_super/README.md) |
| OpenInstruct Mistral 7B | Q2_K Q3_K_M Q3_K_S Q5_0 Q6_K Q3_K_L Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q8_0 | 32768 | 4096 | [README.md](models/openinstruct-mistral-7B/README.md) |
| OpenThaiGPT 1.0.0 Beta 13B Chat | Q5_K_S Q6_K Q8_0 Q2_K Q4_K_S Q5_0 Q5_K_M Q4_K_M Q3_K_L Q3_K_M Q3_K_S Q4_0 | 4096 | 5120 | [README.md](models/openthaigpt-1.0.0-beta-13B-chat/README.md) |
| Opus v0 70B | Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_K_M Q5_K_S Q3_K_L Q4_K_S Q5_0 | 4096 | 8192 | [README.md](models/opus-v0-70B/README.md) |
| Opus V0 7B | Q5_K_S Q6_K Q5_0 Q5_K_M Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q8_0 Q2_K Q3_K_L | 32768 | 4096 | [README.md](models/opus-v0-7B/README.md) |
| Opus V0.5 70B | Q3_K_M Q3_K_S Q4_0 Q4_K_S Q2_K Q3_K_L Q4_K_M Q5_0 Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/opus-v0.5-70B/README.md) |
| Orangetin OpenHermes Mixtral 8X7B | Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q2_K | 32768 | 4096 | [README.md](models/orangetin-OpenHermes-Mixtral-8x7B/README.md) |
| Orca Mini v3 13B | Q5_0 Q5_K_M Q6_K Q2_K Q3_K_M Q4_0 Q4_K_M Q4_K_S Q3_K_L Q3_K_S Q5_K_S Q8_0 | 4096 | 5120 | [README.md](models/orca_mini_v3_13B/README.md) |
| Orca Mini v3 70B | Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_0 Q5_K_M Q2_K Q4_K_S Q5_K_S Q4_K_M | 4096 | 8192 | [README.md](models/orca_mini_v3_70B/README.md) |
| Orca Mini v3 7B | Q5_K_S Q6_K Q8_0 Q3_K_S Q4_0 Q5_0 Q5_K_M Q4_K_S Q2_K Q3_K_L Q3_K_M Q4_K_M | 4096 | 4096 | [README.md](models/orca_mini_v3_7B/README.md) |
| Pee | Q5_K_M Q5_K_S Q6_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q8_0 Q2_K | 32768 | 4096 | [README.md](models/pee/README.md) |
| Phi 2 | Q5_K_M Q5_K_S Q8_0 Q4_K_M Q4_K_S Q5_0 Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q6_K | 0 | 0 | [README.md](models/phi-2/README.md) |
| Platypus Yi 34B | Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_0 Q6_K Q2_K Q3_K_L Q8_0 Q5_K_S Q4_0 Q5_K_M | 4096 | 7168 | [README.md](models/platypus-yi-34b/README.md) |
| Prometheus 13B V1.0 | Q3_K_M Q4_0 Q5_K_M Q2_K Q3_K_L Q4_K_S Q5_0 Q5_K_S Q6_K Q8_0 Q3_K_S Q4_K_M | 4096 | 5120 | [README.md](models/prometheus-13B-v1.0/README.md) |
| Prometheus 7B V1.0 | Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q6_K Q2_K Q3_K_M Q4_0 Q5_0 Q5_K_M Q5_K_S Q8_0 | 4096 | 4096 | [README.md](models/prometheus-7B-v1.0/README.md) |
| Psyonic Cetacean 20B | Q4_K_S Q6_K Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q8_0 | 4096 | 5120 | [README.md](models/psyonic-cetacean-20B/README.md) |
| qCammel 13 | Q5_K_M Q5_K_S Q6_K Q3_K_M Q3_K_S Q4_K_S Q4_K_M Q5_0 Q8_0 Q2_K Q3_K_L Q4_0 | 4096 | 5120 | [README.md](models/qCammel-13/README.md) |
| qCammel 70 | Q3_K_L Q4_0 Q5_0 Q2_K Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q5_K_S | 4096 | 8192 | [README.md](models/qCammel-70-x/README.md) |
| Quantum DPO v0.1 | Q4_0 Q4_K_S Q5_K_M Q8_0 Q2_K Q3_K_M Q3_K_S Q5_K_S Q6_K Q3_K_L Q4_K_M Q5_0 | 32768 | 4096 | [README.md](models/quantum-dpo-v0.1/README.md) |
| Quantum v0.01 | Q2_K Q4_0 Q5_K_M Q5_K_S Q8_0 Q5_0 Q6_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q4_K_S | 32768 | 4096 | [README.md](models/quantum-v0.01/README.md) |
| Rocket 3B | Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_S Q8_0 Q2_K Q3_K_L Q5_K_M Q6_K Q4_0 Q5_0 | 4096 | 2560 | [README.md](models/rocket-3B/README.md) |
| RPGuild ChatML 13B | Q6_K Q8_0 Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_0 Q3_K_M Q4_K_S Q5_K_M Q5_K_S | 4096 | 5120 | [README.md](models/rpguild-chatml-13B/README.md) |
| Sabia 7B | Q6_K Q8_0 Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q3_K_L Q5_K_S | 2048 | 4096 | [README.md](models/sabia-7B/README.md) |
| Saiga Mistral 7B | Q3_K_M Q4_K_S Q5_0 Q5_K_M Q6_K Q2_K Q3_K_S Q4_0 Q4_K_M Q5_K_S Q8_0 Q3_K_L | 32768 | 4096 | [README.md](models/saiga_mistral_7b/README.md) |
| Samantha 1.2 Mistral 7B | Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q3_K_L Q8_0 Q6_K Q4_K_S Q2_K | 32768 | 4096 | [README.md](models/samantha-1.2-mistral-7B/README.md) |
| Samantha Mistral 7B | Q2_K Q3_K_S Q4_0 Q4_K_M Q4_K_S Q8_0 Q3_K_L Q3_K_M Q5_0 Q5_K_M Q5_K_S Q6_K | 32768 | 4096 | [README.md](models/samantha-mistral-7B/README.md) |
| Samantha Mistral Instruct 7B | Q2_K Q4_K_S Q5_0 Q5_K_S Q8_0 Q5_K_M Q6_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M | 32768 | 4096 | [README.md](models/samantha-mistral-instruct-7B/README.md) |
| Sheep Duck Llama 2 13B | Q3_K_L Q3_K_S Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q8_0 | 4096 | 5120 | [README.md](models/sheep-duck-llama-2-13B/README.md) |
| Sheep Duck Llama 2 70B v1.1 | Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q2_K Q4_0 Q5_0 | 4096 | 8192 | [README.md](models/sheep-duck-llama-2-70B-v1.1/README.md) |
| Smartyplats 7B v2 | Q3_K_M Q5_0 Q5_K_M Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S | 32768 | 4096 | [README.md](models/smartyplats-7B-v2/README.md) |
| Smol 7B | Q3_K_L Q4_K_S Q5_K_M Q5_K_S Q4_K_M Q5_0 Q6_K Q8_0 Q2_K Q3_K_M Q3_K_S Q4_0 | 8192 | 4096 | [README.md](models/smol-7B/README.md) |
| Speechless Code Mistral 7B v1.0 | Q5_K_M Q6_K Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_S Q3_K_L Q8_0 | 32768 | 4096 | [README.md](models/speechless-code-mistral-7B-v1.0/README.md) |
| Speechless Codellama 34B v2.0 | Q6_K Q3_K_S Q4_0 Q4_K_M Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_L Q3_K_M Q4_K_S Q5_0 | 16384 | 8192 | [README.md](models/speechless-codellama-34b-v2.0/README.md) |
| Speechless Mistral 7B Dare 0.85 | Q3_K_L Q4_0 Q5_K_S Q5_0 Q5_K_M Q6_K Q2_K Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q8_0 | 32768 | 4096 | [README.md](models/speechless-mistral-7B-dare-0.85/README.md) |
| Speechless Mistral Dolphin Orca Platypus Samantha 7B | Q6_K Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_S Q8_0 Q2_K Q3_K_L Q5_K_M | 32768 | 4096 | [README.md](models/speechless-mistral-dolphin-orca-platypus-samantha-7B/README.md) |
| Speechless Tora Code 7B v1.0 | Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_0 Q5_K_S Q2_K Q3_K_L Q4_0 Q5_K_M Q6_K Q8_0 | 16384 | 4096 | [README.md](models/speechless-tora-code-7B-v1.0/README.md) |
| Sqlcoder | Q8_0 Q2_K Q3_K_L Q4_0 Q4_K_S Q5_K_M Q6_K Q3_K_M Q3_K_S Q4_K_M Q5_0 Q5_K_S | 0 | 0 | [README.md](models/sqlcoder/README.md) |
| SQLCoder 34B Alpha | Q4_K_M Q5_K_S Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q6_K Q2_K Q8_0 | 16384 | 8192 | [README.md](models/sqlcoder-34b-alpha/README.md) |
| SQLCoder 7B | Q4_0 Q4_K_S Q5_K_M Q6_K Q2_K Q3_K_L Q3_K_S Q5_K_S Q8_0 Q3_K_M Q4_K_M Q5_0 | 32768 | 4096 | [README.md](models/sqlcoder-7B/README.md) |
| Sqlcoder2 | Q6_K Q2_K Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q8_0 Q3_K_L Q3_K_M Q5_K_S | 0 | 0 | [README.md](models/sqlcoder2/README.md) |
| Stable Vicuna 13B | Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q6_K Q2_K Q3_K_M Q4_0 Q5_0 Q5_K_M Q5_K_S Q8_0 | 2048 | 5120 | [README.md](models/stable-vicuna-13B/README.md) |
| StableLM Zephyr 3B | Q5_K_M Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_S Q5_0 Q6_K Q8_0 | 4096 | 2560 | [README.md](models/stablelm-zephyr-3b/README.md) |
| Stockmark 13B | Q4_K_M Q4_K_S Q6_K Q8_0 Q2_K Q3_K_L Q3_K_S Q4_0 Q3_K_M Q5_0 Q5_K_M Q5_K_S | 2048 | 5120 | [README.md](models/stockmark-13B/README.md) |
| Storytime 13B | Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_K_S Q3_K_M Q4_K_S Q5_0 Q5_K_M Q6_K Q8_0 | 4096 | 5120 | [README.md](models/storytime-13B/README.md) |
| Supermario V2 | Q5_K_S Q8_0 Q3_K_L Q3_K_M Q4_0 Q5_0 Q5_K_M Q2_K Q3_K_S Q4_K_M Q4_K_S Q6_K | 32768 | 4096 | [README.md](models/supermario-v2/README.md) |
| Tigerbot 70B Chat V2 | Q2_K Q3_K_S Q5_0 Q5_K_S Q3_K_L Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_K_M | 4096 | 8192 | [README.md](models/tigerbot-70B-chat-v2/README.md) |
| Tigerbot 70B Chat v4 | Q3_K_M Q3_K_S Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q2_K Q4_K_S Q3_K_L | 4096 | 8192 | [README.md](models/tigerbot-70B-chat-v4/README.md) |
| ToRA 13B v1.0 | Q3_K_L Q3_K_M Q4_K_M Q6_K Q8_0 Q2_K Q3_K_S Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 5120 | [README.md](models/tora-13B-v1.0/README.md) |
| ToRA 70B v1.0 | Q4_0 Q4_K_S Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q5_K_S | 4096 | 8192 | [README.md](models/tora-70B-v1.0/README.md) |
| ToRA 7B v1.0 | Q3_K_L Q3_K_M Q4_0 Q4_K_M Q6_K Q8_0 Q2_K Q3_K_S Q4_K_S Q5_0 Q5_K_M Q5_K_S | 4096 | 4096 | [README.md](models/tora-7B-v1.0/README.md) |
| ToRA Code 13B v1.0 | Q3_K_M Q3_K_S Q4_0 Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_L Q4_K_M Q4_K_S Q5_K_S | 16384 | 5120 | [README.md](models/tora-code-13B-v1.0/README.md) |
| ToRA Code 34B v1.0 | Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_K_M Q8_0 Q2_K Q4_K_M Q4_K_S Q5_0 Q5_K_S Q6_K | 16384 | 8192 | [README.md](models/tora-code-34b-v1.0/README.md) |
| ToRA Code 7B v1.0 | Q2_K Q3_K_M Q4_0 Q5_0 Q5_K_S Q6_K Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q8_0 | 16384 | 4096 | [README.md](models/tora-code-7B-v1.0/README.md) |
| Tulu 13B | Q3_K_S Q4_0 Q5_0 Q5_K_M Q2_K Q3_K_M Q4_K_M Q4_K_S Q5_K_S Q6_K Q8_0 Q3_K_L | 2048 | 5120 | [README.md](models/tulu-13B/README.md) |
| Tulu 2 13B | Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_S Q4_0 Q4_K_S Q5_0 Q3_K_L Q3_K_M Q4_K_M Q8_0 | 8192 | 5120 | [README.md](models/tulu-2-13B/README.md) |
| Tulu 2 70B | Q5_K_M Q3_K_S Q4_0 Q3_K_M Q4_K_M Q4_K_S Q5_0 Q5_K_S Q2_K Q3_K_L | 8192 | 8192 | [README.md](models/tulu-2-70B/README.md) |
| Tulu 2 7B | Q5_K_S Q3_K_L Q3_K_M Q4_K_M Q4_K_S Q5_K_M Q6_K Q8_0 Q2_K Q3_K_S Q4_0 Q5_0 | 8192 | 4096 | [README.md](models/tulu-2-7B/README.md) |
| Tulu 2 DPO 13B | Q3_K_M Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q2_K Q3_K_L Q5_0 Q6_K Q8_0 | 8192 | 5120 | [README.md](models/tulu-2-dpo-13B/README.md) |
| Tulu 2 DPO 70B | Q2_K Q3_K_L Q4_0 Q4_K_S Q5_0 Q5_K_M Q5_K_S Q3_K_M Q3_K_S Q4_K_M | 8192 | 8192 | [README.md](models/tulu-2-dpo-70B/README.md) |
| Tulu 2 DPO 7B | Q4_K_S Q5_0 Q5_K_M Q2_K Q3_K_L Q3_K_S Q4_0 Q4_K_M Q5_K_S Q3_K_M Q6_K Q8_0 | 8192 | 4096 | [README.md](models/tulu-2-dpo-7B/README.md) |
| Tulu 30B | Q5_K_S Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q6_K Q8_0 Q2_K Q3_K_M | 2048 | 6656 | [README.md](models/tulu-30B/README.md) |
| Tulu 7B | Q4_K_S Q5_K_M Q6_K Q2_K Q4_0 Q3_K_S Q4_K_M Q5_0 Q5_K_S Q8_0 Q3_K_L Q3_K_M | 2048 | 4096 | [README.md](models/tulu-7B/README.md) |
| Typhoon 7B | Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q6_K Q8_0 Q2_K Q3_K_M Q4_0 Q5_0 Q5_K_M Q5_K_S | 32768 | 4096 | [README.md](models/typhoon-7B/README.md) |
| Una Cybertron 7B v2 | Q5_K_M Q6_K Q3_K_L Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_0 Q8_0 Q2_K Q4_K_M Q5_K_S | 32768 | 4096 | [README.md](models/una-cybertron-7B-v2/README.md) |
| Una Cybertron 7B V3 OMA | Q3_K_M Q3_K_S Q5_K_S Q6_K Q5_K_M Q8_0 Q2_K Q3_K_L Q4_0 Q4_K_M Q4_K_S Q5_0 | 32768 | 4096 | [README.md](models/una-cybertron-7B-v3-OMA/README.md) |
| Una Xaberius 34B v1Beta | Q3_K_L Q3_K_S Q4_K_M Q4_K_S Q5_K_M Q5_K_S Q6_K Q8_0 Q2_K Q3_K_M Q4_0 Q5_0 | 4096 | 7168 | [README.md](models/una-xaberius-34b-v1beta/README.md) |
| Llama 30B Instruct 2048 | Q5_K_M Q6_K Q8_0 Q2_K Q3_K_M Q3_K_S Q4_0 Q5_0 Q3_K_L Q4_K_M Q4_K_S Q5_K_S | 2048 | 6656 | [README.md](models/upstage-llama-30b-instruct-2048/README.md) |
| V1Olet Marcoroni Go Bruins Merge 7B | Q4_0 Q8_0 Q2_K Q3_K_M Q3_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q3_K_L Q4_K_M Q4_K_S | 32768 | 4096 | [README.md](models/v1olet_marcoroni-go-bruins-merge-7B/README.md) |
| V1Olet Merged DPO 7B v3 | Q3_K_M Q4_0 Q4_K_M Q5_0 Q5_K_M Q5_K_S Q8_0 Q2_K Q3_K_L Q3_K_S Q4_K_S Q6_K | 32768 | 4096 | [README.md](models/v1olet_merged_dpo_7B_v3/README.md) |
| Vicuna 13B v1.5 | Q2_K Q3_K_L Q3_K_M Q3_K_S Q4_K_M Q5_0 Q5_K_M Q6_K Q8_0 Q4_0 Q4_K_S Q5_K_S | 4096 | 5120 | [README.md](models/vicuna-13B-v1.5/README.md) |
| Vicuna 13B v1.5 16K | Q2_K Q3_K_M Q3_K_S Q4_0 Q4_K_S Q5_K_M Q3_K_L Q4_K_M Q5_0 Q5_K_S Q6_K Q8_0 | 4096 | 5120 | [README.md](models/vicuna-13B-v1.5-16K/README.md) |
| Vicuna 33B V1.3 | Q5_K_M Q3_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_S Q6_K Q8_0 Q2_K Q3_K_L Q4_0 Q5_0 | 2048 | 6656 | [README.md](models/vicuna-33B/README.md) |
| Vicuna 33B Coder | Q2_K Q4_K_M Q4_K_S Q5_K_M Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_0 Q5_0 Q5_K_S Q6_K | 2048 | 6656 | [README.md](models/vicuna-33B-coder/README.md) |
| Vicuna 7B v1.5 | Q8_0 Q3_K_S Q5_0 Q5_K_M Q5_K_S Q4_K_M Q4_K_S Q6_K Q2_K Q3_K_L Q3_K_M Q4_0 | 4096 | 4096 | [README.md](models/vicuna-7B-v1.5/README.md) |
| Vicuna 7B v1.5 16K | Q6_K Q8_0 Q2_K Q3_K_L Q3_K_M Q4_0 Q5_0 Q5_K_M Q3_K_S Q4_K_M Q4_K_S Q5_K_S | 4096 | 4096 | [README.md](models/vicuna-7B-v1.5-16K/README.md) |
| Vietnamese Llama2 7B 40GB | Q3_K_L Q3_K_M Q5_K_M Q6_K Q2_K Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_S Q8_0 | 4096 | 4096 | [README.md](models/vietnamese-llama2-7B-40GB/README.md) |
| Vigogne 2 70B Chat | Q3_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q2_K Q3_K_S Q4_0 Q4_K_M Q3_K_L | 4096 | 8192 | [README.md](models/vigogne-2-70B-chat/README.md) |
| Wizard Mega 13B | Q2_K Q3_K_L Q3_K_S Q4_K_M Q5_0 Q5_K_S Q3_K_M Q4_0 Q4_K_S Q5_K_M Q6_K Q8_0 | 2048 | 5120 | [README.md](models/wizard-mega-13B/README.md) |
| Wizard Vicuna 13B | Q4_0 Q4_K_M Q5_0 Q5_K_M Q2_K Q3_K_M Q3_K_S Q4_K_S Q5_K_S Q6_K Q8_0 Q3_K_L | 2048 | 5120 | [README.md](models/wizard-vicuna-13B/README.md) |
| WizardLM 7B v1.0 | Q2_K Q3_K_M Q3_K_S Q4_K_S Q5_K_S Q6_K Q3_K_L Q4_0 Q4_K_M Q5_0 Q5_K_M Q8_0 | 2048 | 4096 | [README.md](models/wizardLM-7B/README.md) |
| xDAN L1 Chat RL v1 | Q8_0 Q3_K_L Q3_K_M Q3_K_S Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q2_K Q4_0 Q4_K_M | 32768 | 4096 | [README.md](models/xDAN-L1-Chat-RL-v1/README.md) |
| Yi 34B v3 | Q3_K_L Q4_0 Q4_K_M Q5_0 Q5_K_S Q6_K Q8_0 Q2_K Q3_K_S Q4_K_S Q5_K_M Q3_K_M | 8192 | 7168 | [README.md](models/yi-34B-v3/README.md) |
| Zephyr 7B Alpha | Q5_K_S Q3_K_M Q3_K_L Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q2_K Q8_0 Q6_K | 32768 | 4096 | [README.md](models/zephyr-7B-alpha/README.md) |
| Zephyr 7B Beta | Q5_K_S Q6_K Q3_K_S Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q8_0 Q2_K Q3_K_L Q3_K_M | 32768 | 4096 | [README.md](models/zephyr-7B-beta/README.md) |
| Zephyr 7B Beta PL | Q3_K_M Q4_0 Q4_K_M Q4_K_S Q5_0 Q5_K_M Q5_K_S Q6_K Q2_K Q3_K_L Q3_K_S Q8_0 | 32768 | 4096 | [README.md](models/zephyr-7B-beta-pl/README.md) |

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
