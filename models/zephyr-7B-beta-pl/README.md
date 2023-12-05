---
license: mit
language:
- pl
---

## Model Overview
The model is a result of advanced fine-tuning methods applied to a base model, focusing on enhancing its capabilities for specific Polish language datasets. It incorporates cutting-edge techniques and is built upon the Zephyr Beta model framework.

## Current Status: Alpha
- **Stage**: Alpha-Alpaca

## Training Details

I trained the model using 3xRTX 3090 for 163 hours.
[![Built with Axolotl](https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png)](https://github.com/OpenAccess-AI-Collective/axolotl)
## Quantised Model Links:



1. https://huggingface.co/Nondzu/zephyr-7b-beta-pl-exl2
2. https://huggingface.co/TheBloke/zephyr-7B-beta-pl-GGUF
3. https://huggingface.co/TheBloke/zephyr-7B-beta-pl-AWQ
4. https://huggingface.co/TheBloke/zephyr-7B-beta-pl-GPTQ


## Model Specifics
- **Base Model**: HuggingFaceH4/zephyr-7b-beta
- **Fine-Tuning Method**: QLORA
- **Primary Focus**: Polish language datasets

## Datasets:
- Dataset 1 Name: Lajonbot/alpaca-dolly-chrisociepa-instruction-only-polish
- Dataset 1 Link: [Lajonbot/alpaca-dolly-chrisociepa-instruction-only-polish](https://huggingface.co/datasets/Lajonbot/alpaca-dolly-chrisociepa-instruction-only-polish?row=16)
- Dataset 2 Name: klima7/polish-prose
- Dataset 2 Link: [klima7/polish-prose](https://huggingface.co/datasets/klima7/polish-prose)

## Usage Warning
As this is an experimental model, users should be aware of the following:
- **Reliability**: The model has not been fully tested and may exhibit unexpected behaviors or performance issues.
- **Updates**: The model is subject to change based on ongoing testing and feedback.
- **Data Sensitivity**: Users should exercise caution when using sensitive or private data, as the model's output and behavior are not fully predictable at this stage.

## Prompt template: Alpaca

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:

```

## Example

![image/png](https://cdn-uploads.huggingface.co/production/uploads/63729f35acef705233c87909/1WYp9Su1NYvYCIU-2J7TG.png)

## Feedback and Contribution
User feedback is crucial during this testing phase. We encourage users to provide feedback on model performance, issues encountered, and any suggestions for improvements. Contributions in terms of shared test results, datasets, or code improvements are also welcome.

---

**Disclaimer**: This experimental model is provided 'as is', without warranty of any kind. Users should use the model at their own risk. The creators or maintainers of the model are not responsible for any consequences arising from its use.


![image/png](https://cdn-uploads.huggingface.co/production/uploads/63729f35acef705233c87909/CPClYNIMp3Qswt2F0Y9B3.png)