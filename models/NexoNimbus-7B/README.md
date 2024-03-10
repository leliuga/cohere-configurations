---
license: apache-2.0
tags:
- merge
- abideen/DareVox-7B
- udkai/Garrulus
language:
- en
---

# NexoNimbus-7B



![image/png](https://cdn-uploads.huggingface.co/production/uploads/64e380b2e12618b261fa6ba0/9lIzCPqDYR6nnLgoH6kMp.png)


NexoNimbus-7B is a merge of the following models:
* [abideen/DareVox-7B](https://huggingface.co/abideen/DareVox-7B)
* [udkai/Garrulus](https://huggingface.co/udkai/Garrulus)

🏆 Evaluation
NexoNimbus-7B is the 5th best-performing 7B LLM on the Open LLM Leaderboard:

![image/png](https://cdn-uploads.huggingface.co/production/uploads/64e380b2e12618b261fa6ba0/MIkOaXVGJ0T5UVYIEhtYA.png)


|    Task     |Version| Metric |Value|   |Stderr|
|-------------|------:|--------|----:|---|-----:|
|arc_challenge|      0|acc     |68.25|±  |  1.36|
|             |       |acc_norm|70.81|±  |  1.38|
|hellaswag    |      0|acc     |70.86|±  |  0.45|
|             |       |acc_norm|87.86|±  |  0.32|
|gsm8k        |      0|acc     |70.35|±  |  1.25|
|winogrande   |      0|acc     |84.84|±  |  1.00|
|mmlu         |      0|acc     |64.69|±  |  1.00|

Average: 73.5%

### TruthfulQA
|    Task     |Version|Metric|Value|   |Stderr|
|-------------|------:|------|----:|---|-----:|
|truthfulqa_mc|      1|mc1   |46.26|±  |  1.74|
|             |       |mc2   |62.42|±  |  1.54|


## 🧩 Configuration

```yaml
slices:
  - sources:
      - model: abideen/DareVox-7B 
        layer_range: [0, 32]
      - model: udkai/Garrulus
        layer_range: [0, 32]
merge_method: slerp
base_model: abideen/DareVox-7B 
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

Here's a [Colab notebook](https://colab.research.google.com/drive/1F9lzL1IeZRMgiSbY9UbgCR__RreIflJh?usp=sharing) to run NexoNimbus-7B in 4-bit precision on a free T4 GPU.

```python
!pip install -qU transformers accelerate

from transformers import AutoTokenizer
import transformers
import torch

model = "abideen/NexoNimbus-7B"
messages = [{"role": "user", "content": "Explain what is Machine learning."}]

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

"Machine learning is a subfield of artificial intelligence that focuses on developing algorithms and models that allow computers to learn and improve their performance over time, without being explicitly programmed. It involves the use of statistical techniques and data analysis to identify patterns and make predictions based on input data.
In machine learning, data is fed into a model, which then adjusts its internal parameters to minimize the difference between the predicted output and the actual output. This process is called training, and as the model is exposed to more data, it becomes better at making predictions or classifications.
Machine learning can be divided into several categories, including supervised learning, unsupervised learning, and reinforcement learning. Supervised learning involves using labeled data, where the desired output is known, and the model learns to map inputs to outputs. Unsupervised learning, on the other hand, does not have a predefined output, and the model learns to identify patterns or relationships within the data. Reinforcement learning involves learning through trial and error, with the model receiving feedback in the form of rewards or penalties based on its actions.
Some common applications of machine learning include image recognition, natural language processing, recommendation systems, fraud detection, and self-driving."