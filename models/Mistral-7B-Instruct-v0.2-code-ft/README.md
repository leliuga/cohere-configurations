---
license: cc-by-nc-nd-4.0
---

# Mistral-7B-Instruct-v0.2-code-ft

I'm thrilled to introduce the latest iteration of our model, Mistral-7B-Instruct-v0.2-code-ft. This updated version is designed to further enhance coding assistance and co-pilot functionalities. We're eager for developers and enthusiasts to try it out and provide feedback!

## Additional Information

This version builds upon the previous Mistral-7B models, incorporating new datasets and features for a more refined experience.

## Prompt template: ChatML
```
<|im_start|>system
{system_message}<|im_end|>
<|im_start|>user
{prompt}<|im_end|>
<|im_start|>assistant
```

## Quantised Model Links:



1. https://huggingface.co/LoneStriker/Mistral-7B-Instruct-v0.2-code-ft-8.0bpw-h8-exl2
2. https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-code-ft-GGUF
3. https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-code-ft-AWQ
4. https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-code-ft-GPTQ


## Eval Plus Performance

For detailed performance metrics, visit Eval Plus page: [Mistral-7B-Instruct-v0.2-code-ft Eval Plus](https://github.com/evalplus/evalplus)

Score: 0.421 
![image/png](https://cdn-uploads.huggingface.co/production/uploads/63729f35acef705233c87909/VLtuWPh8m07bgU8BElpNv.png)

## Dataset: 
The model has been trained on a new dataset to improve its performance and versatility:
  - path: ajibawa-2023/Code-74k-ShareGPT
    
    type: sharegpt
    
    conversation: chatml
    
Find more about the dataset here: [Code-74k-ShareGPT Dataset](https://huggingface.co/datasets/ajibawa-2023/Code-74k-ShareGPT)

## Model Architecture

- Base Model: mistralai/Mistral-7B-Instruct-v0.2
- Tokenizer Type: LlamaTokenizer
- Model Type: MistralForCausalLM
- Is Mistral Derived Model: true
- Sequence Length: 16384 with sample packing

## Enhanced Features

- Adapter: qlora
- Learning Rate: 0.0002 with cosine lr scheduler
- Optimizer: adamw_bnb_8bit
- Training Enhancements: bf16 training, gradient checkpointing, and flash attention


## Download Information

You can download and explore this model through these links on Hugging Face.

## Contributions and Feedback

We welcome contributions and feedback from the community. Please feel free to open issues or pull requests on repository.

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl"/>](https://github.com/OpenAccess-AI-Collective/axolotl)

