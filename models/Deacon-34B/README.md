---
license: other
license_name: yi-license
license_link: LICENSE
pipeline_tag: text-generation
datasets:
- totally-not-an-llm/EverythingLM-data-V3
---

This model has been llamafied and uses a llama tokenizer. I took it from https://huggingface.co/chargoddard/Yi-34B-Llama
It's fine tuned on EverythingLM dataset for 5 epochs with NEFTune. If you want to understand the pun of the model name, you gotta look at the 3b version of it. 

License
The Yi series models are fully open for academic research and free commercial usage with permission via applications. All usage must adhere to the Model License Agreement 2.0. To apply for the official commercial license, please contact us (yi@01.ai).


Prompt Example:
```
### System:

You are an AI assistant. User will give you a task. Your goal is to complete the task as faithfully as you can. While performing the task think step-by-step and justify your steps.

### Instruction: 

How do you fine tune a large language model? 

### Response:
```