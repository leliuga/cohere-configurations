---
license: apache-2.0
---

## Model Details: Neural-Chat-v3-3

This model is a fine-tuned 7B parameter LLM on the Intel Gaudi 2 processor from the [Intel/neural-chat-7b-v3-1](https://huggingface.co/Intel/neural-chat-7b-v3-1) on the [meta-math/MetaMathQA](https://huggingface.co/datasets/meta-math/MetaMathQA) dataset. The model was aligned using the Direct Performance Optimization (DPO) method with [Intel/orca_dpo_pairs](https://huggingface.co/datasets/Intel/orca_dpo_pairs). The [Intel/neural-chat-7b-v3-1](https://huggingface.co/Intel/neural-chat-7b-v3-1) was originally fine-tuned from [mistralai/Mistral-7B-v-0.1](https://huggingface.co/mistralai/Mistral-7B-v0.1).  For more information, refer to our blog [The Practice of Supervised Fine-tuning and Direct Preference Optimization on Intel Gaudi2](https://medium.com/@NeuralCompressor/the-practice-of-supervised-finetuning-and-direct-preference-optimization-on-habana-gaudi2-a1197d8a3cd3).

**Note:** Adjust lora modules to trade off truthfulqa and gsm8k performance on DPO stage.


| Model Detail | Description |
| ----------- | ----------- | 
| Model Authors - Company | Intel. The NeuralChat team with members from Intel/DCAI/AISE/AIPT. Core team members: Kaokao Lv, Liang Lv, Chang Wang, Wenxin Zhang, Xuhui Ren, and Haihao Shen.| 
| Date | December, 2023 | 
| Version | v3-3 | 
| Type | 7B Large Language Model | 
| Paper or Other Resources | [Medium Blog](https://medium.com/@NeuralCompressor/the-practice-of-supervised-finetuning-and-direct-preference-optimization-on-habana-gaudi2-a1197d8a3cd3) | 
| License | Apache 2.0 |
| Questions or Comments | [Community Tab](https://huggingface.co/Intel/neural-chat-7b-v3-3/discussions) and [Intel Developers Discord](https://discord.gg/rv2Gp55UJQ)|

| Intended Use | Description |
| ----------- | ----------- | 
| Primary intended uses | You can use the fine-tuned model for several language-related tasks. Checkout the [LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard) to see how this model and others from Intel are doing. | 
| Primary intended users | Anyone doing inference on language-related tasks. | 
| Out-of-scope uses | This model in most cases will need to be fine-tuned for your particular task.  The model should not be used to intentionally create hostile or alienating environments for people.|

## How to use and Sample Code
Here is the sample code to reproduce the model: [Sample Code](https://github.com/intel/intel-extension-for-transformers/blob/main/intel_extension_for_transformers/neural_chat/examples/finetuning/finetune_neuralchat_v3/README.md).

## Prompt Template
```plaintext
### System:
{system}
### User:
{usr}
### Assistant:

```

## [Quantitative Analyses: Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Intel__neural-chat-7b-v3-3) (**note:** the leaderboard removed drop task)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 69.83   |
| ARC (25-shot)         | 66.89          |
| HellaSwag (10-shot)   | 85.26    |
| MMLU (5-shot)         | 63.07         |
| TruthfulQA (0-shot)   | 63.01   |
| Winogrande (5-shot)   | 79.64   |
| GSM8K (5-shot)        | 61.11        |

## Useful links
* Intel Neural Compressor [link](https://github.com/intel/neural-compressor)
* Intel Extension for Transformers [link](https://github.com/intel/intel-extension-for-transformers)

## Ethical Considerations and Limitations
neural-chat-7b-v3-3 can produce factually incorrect output, and should not be relied on to produce factually accurate information. Because of the limitations of the pretrained model and the finetuning datasets, it is possible that this model could generate lewd, biased or otherwise offensive outputs.

Therefore, before deploying any applications of neural-chat-7b-v3-3, developers should perform safety testing.

## Disclaimer

The license on this model does not constitute legal advice. We are not responsible for the actions of third parties who use this model. Please cosult an attorney before using this model for commercial purposes.


