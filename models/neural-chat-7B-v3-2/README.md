---
license: apache-2.0
---

## Fine-tuning on Intel Gaudi2

This model is a fine-tuned model based on [Intel/neural-chat-7b-v3-1](https://huggingface.co/Intel/neural-chat-7b-v3-1) on the [meta-math/MetaMathQA](https://huggingface.co/datasets/meta-math/MetaMathQA) dataset. And make a alignment using DPO method with [Intel/orca_dpo_pairs](https://huggingface.co/datasets/Intel/orca_dpo_pairs). For more details about [Intel/neural-chat-7b-v3-1](https://huggingface.co/Intel/neural-chat-7b-v3-1), you can refer the [model card](https://huggingface.co/Intel/neural-chat-7b-v3-1) and our blog [The Practice of Supervised Fine-tuning and Direct Preference Optimization on Intel Gaudi2](https://medium.com/@NeuralCompressor/the-practice-of-supervised-finetuning-and-direct-preference-optimization-on-habana-gaudi2-a1197d8a3cd3)

## Model date
Neural-chat-7b-v3-2 was trained at December, 2023.

### Training sample code
Here is the sample code to reproduce the model: [Sample Code](https://github.com/intel/intel-extension-for-transformers/blob/main/intel_extension_for_transformers/neural_chat/examples/finetuning/finetune_neuralchat_v3/README.md).

## Prompt Template

```
### System:
{system}
### User:
{usr}
### Assistant:

```

## Ethical Considerations and Limitations
neural-chat-7b-v3-2 can produce factually incorrect output, and should not be relied on to produce factually accurate information. Because of the limitations of the pretrained model and the finetuning datasets, it is possible that this model could generate lewd, biased or otherwise offensive outputs.

Therefore, before deploying any applications of neural-chat-7b-v3-2, developers should perform safety testing.

## Disclaimer

The license on this model does not constitute legal advice. We are not responsible for the actions of third parties who use this model. Please cosult an attorney before using this model for commercial purposes.

## Organizations developing the model

The NeuralChat team with members from Intel/DCAI/AISE/AIPT. Core team members: Kaokao Lv, Liang Lv, Chang Wang, Wenxin Zhang, Xuhui Ren, and Haihao Shen.

## Useful links
* Intel Neural Compressor [link](https://github.com/intel/neural-compressor)
* Intel Extension for Transformers [link](https://github.com/intel/intel-extension-for-transformers)

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Intel__neural-chat-7b-v3-1) (**note:** the leaderboard removes drop task)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 68.29   |
| ARC (25-shot)         | 67.49          |
| HellaSwag (10-shot)   | 83.92    |
| MMLU (5-shot)         | 63.55         |
| TruthfulQA (0-shot)   | 59.68   |
| Winogrande (5-shot)   | 79.95   |
| GSM8K (5-shot)        | 55.12        |
