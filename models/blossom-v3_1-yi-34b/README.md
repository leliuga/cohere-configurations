---
license: apache-2.0
datasets:
- Azure99/blossom-chat-v1
- Azure99/blossom-math-v2
- Azure99/blossom-wizard-v1
- Azure99/blossom-orca-v1
language:
- zh
- en
---
# **BLOSSOM-v3.1-yi-34b**

Try blossom online demo! [Blossom Chat[🚀](https://blossom-chat.com/)](https://huggingface.co/spaces/Azure99/BlossomChat)

### Introduction

Blossom is a conversational large language model, fine-tuned on the Blossom Orca/Wizard/Chat/Math mixed dataset based on the Yi-34B pre-trained model. Blossom possesses robust general capabilities and context comprehension. Additionally, the high-quality Chinese and English datasets used for training have been made open source.

Training was conducted in two stages. The first stage used 100K Wizard, 100K Orca single-turn instruction datasets, training for 1 epoch; the second stage used a 2K Blossom math reasoning dataset, 50K Blossom chat multi-turn dialogue dataset, and 1% randomly sampled data from the first stage, training for 3 epochs.

### Inference

Inference is performed in the form of dialogue continuation.

Single-turn dialogue

```
A chat between a human and an artificial intelligence bot. The bot gives helpful, detailed, and polite answers to the human's questions.
|Human|: hello
|Bot|: 
```

Multi-turn dialogue

```
A chat between a human and an artificial intelligence bot. The bot gives helpful, detailed, and polite answers to the human's questions.
|Human|: hello
|Bot|: Hello! How can I assist you today?</s>
|Human|: Generate a random number using python
|Bot|: 
```

Note: At the end of the Bot's output in the historical conversation, append a `</s>`.