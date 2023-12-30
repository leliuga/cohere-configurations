---
license: apache-2.0
tags:
- merge
- mergekit
---

# NeuralPipe-7B

This model is a merge of the following models made with [mergekit](https://github.com/cg123/mergekit):
 * [OpenPipe/mistral-ft-optimized-1218](https://huggingface.co/OpenPipe/mistral-ft-optimized-1218)
 * [mlabonne/NeuralHermes-2.5-Mistral-7B](https://huggingface.co/mlabonne/NeuralHermes-2.5-Mistral-7B)

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