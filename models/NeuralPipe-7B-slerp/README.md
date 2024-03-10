---
license: apache-2.0
tags:
- merge
- mergekit
model-index:
- name: NeuralPipe-7B-slerp
  results:
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: AI2 Reasoning Challenge (25-Shot)
      type: ai2_arc
      config: ARC-Challenge
      split: test
      args:
        num_few_shot: 25
    metrics:
    - type: acc_norm
      value: 67.75
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralPipe-7B-slerp
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: HellaSwag (10-Shot)
      type: hellaswag
      split: validation
      args:
        num_few_shot: 10
    metrics:
    - type: acc_norm
      value: 86.15
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralPipe-7B-slerp
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: MMLU (5-Shot)
      type: cais/mmlu
      config: all
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 63.94
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralPipe-7B-slerp
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: TruthfulQA (0-shot)
      type: truthful_qa
      config: multiple_choice
      split: validation
      args:
        num_few_shot: 0
    metrics:
    - type: mc2
      value: 59.8
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralPipe-7B-slerp
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: Winogrande (5-shot)
      type: winogrande
      config: winogrande_xl
      split: validation
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 79.64
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralPipe-7B-slerp
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: GSM8k (5-shot)
      type: gsm8k
      config: main
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 69.75
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralPipe-7B-slerp
      name: Open LLM Leaderboard
---

# NeuralPipe-7B

This model is a merge of the following models made with [mergekit](https://github.com/cg123/mergekit):
 * [OpenPipe/mistral-ft-optimized-1218](https://huggingface.co/OpenPipe/mistral-ft-optimized-1218)
 * [mlabonne/NeuralHermes-2.5-Mistral-7B](https://huggingface.co/mlabonne/NeuralHermes-2.5-Mistral-7B)

## ⚡ Quantized models

Thanks to TheBloke for the quantized models:

* **GGUF**: https://huggingface.co/TheBloke/NeuralPipe-7B-slerp-GGUF
* **AWQ**: https://huggingface.co/TheBloke/NeuralPipe-7B-slerp-AWQ
* **GPTQ**: https://huggingface.co/TheBloke/NeuralPipe-7B-slerp-GPTQ
* 
## 🧩 Configuration

```yaml
slices:
  - sources:
      - model: OpenPipe/mistral-ft-optimized-1218
        layer_range: [0, 32]
      - model: mlabonne/NeuralHermes-2.5-Mistral-7B
        layer_range: [0, 32]
merge_method: slerp
base_model: OpenPipe/mistral-ft-optimized-1218
parameters:
  t:
    - filter: self_attn
      value: [0, 0.5, 0.3, 0.7, 1]
    - filter: mlp
      value: [1, 0.5, 0.7, 0.3, 0]
    - value: 0.5
dtype: bfloat16
```

## 💻 Usage

```python
!pip install -qU transformers accelerate

from transformers import AutoTokenizer
import transformers
import torch

model = "mlabonne/NeuralPipe-7B-slerp"
messages = [{"role": "user", "content": "What is a large language model?"}]

tokenizer = AutoTokenizer.from_pretrained(model)
prompt = tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True)
pipeline = transformers.pipeline(
    "text-generation",
    model=model,
    torch_dtype=torch.float16,
    device_map="auto",
)

outputs = pipeline(prompt, max_new_tokens=256, do_sample=True, temperature=0.7, top_k=50, top_p=0.95)
print(outputs[0]["generated_text"])
```

Output:

```
A large language model is an AI system that uses deep learning techniques to process and understand vast amounts of natural language data. It is designed to generate human-like text, perform complex language tasks, and understand the context, nuance, and meaning of textual data. These models are trained on large datasets, often including billions of words, to learn the patterns and relationships in language. As a result, they can generate coherent and contextually relevant text, answer questions, and perform a variety of other language-related tasks. Some well-known large language models include OpenAI's GPT-3, Google's BERT, and Facebook's RoBERTa.
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_mlabonne__NeuralPipe-7B-slerp)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |71.17|
|AI2 Reasoning Challenge (25-Shot)|67.75|
|HellaSwag (10-Shot)              |86.15|
|MMLU (5-Shot)                    |63.94|
|TruthfulQA (0-shot)              |59.80|
|Winogrande (5-shot)              |79.64|
|GSM8k (5-shot)                   |69.75|

