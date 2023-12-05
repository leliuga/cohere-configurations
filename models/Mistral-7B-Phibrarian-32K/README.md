---
license: mit
---

# SciPhi-Mistral-7B-32k Model Card

The SciPhi-Mistral-7B-32k is a Large Language Model (LLM) fine-tuned from Mistral-7B-v0.1. This model underwent a fine-tuning process over four epochs using more than 1 billion tokens, which include regular instruction tuning data and synthetic textbooks. The objective of this work was to increase the model's scientific reasoning and educational abilities. For best results, follow the Alpaca prompting guidelines.

SciPhi-AI is available via a free hosted API, though the exposed model can vary. Currently, SciPhi-Self-RAG-Mistral-7B-32k is available. More details can be found in the docs [here](https://sciphi.readthedocs.io/en/latest/setup/quickstart.html).

## Model Architecture

Base Model: Mistral-7B-v0.1

**Architecture Features:**
- Transformer-based model
- Grouped-Query Attention
- Sliding-Window Attention
- Byte-fallback BPE tokenizer


## Recommended Chat Formatting


We recommend mapping such that

```

messages = [
    {
        "role": "system",
        "content": "You are a friendly chatbot who always responds in the style of a pirate",
    },
    {"role": "user", "content": "How many helicopters can a human eat in one sitting?"},
]

goes to --->

### System:
You are a friendly chatbot who always responds in the style of a pirate

### Instruction:
How many helicopters can a human eat in one sitting?

### Response:
...
```

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)

## References

1. Lian, W., Goodson, B., Wang, G., Pentland, E., Cook, A., Vong, C., & Teknium. (2023). MistralOrca: Mistral-7B Model Instruct-tuned on Filtered OpenOrcaV1 GPT-4 Dataset. *HuggingFace repository*. [Link](https://huggingface.co/Open-Orca/Mistral-7B-OpenOrca)
2. Mukherjee, S., Mitra, A., Jawahar, G., Agarwal, S., Palangi, H., & Awadallah, A. (2023). Orca: Progressive Learning from Complex Explanation Traces of GPT-4. *arXiv preprint arXiv:2306.02707*.
3. Longpre, S., Hou, L., Vu, T., Webson, A., Chung, H. W., Tay, Y., Zhou, D., Le, Q. V., Zoph, B., Wei, J., & Roberts, A. (2023). The Flan Collection: Designing Data and Methods for Effective Instruction Tuning. *arXiv preprint arXiv:2301.13688*.
4. Mistral AI. (2023). Model Card for Mistral-7B-v0.1. The Mistral-7B-v0.1 Large Language Model (LLM) is a pretrained generative text model with 7 billion parameters. Mistral-7B-v0.1 outperforms Llama 2 13B on all benchmarks tested. For full details, please refer to the paper and release blog post. Model Architecture: Transformer with Grouped-Query Attention, Sliding-Window Attention, and Byte-fallback BPE tokenizer. [Link](https://huggingface.co/mistralai/Mistral-7B-v0.1)


## Acknowledgements

Thank you to the [AI Alignment Lab](https://huggingface.co/Alignment-Lab-AI), [vikp](https://huggingface.co/vikp), [jph00](https://huggingface.co/jph00) and others who contributed to this work.