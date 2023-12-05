---
license: apache-2.0
language:
- ja
- en
tags:
- japanese
- causal-lm
inference: false
---
# CyberAgentLM2-7B-Chat (CALM2-7B-Chat)

## Model Description

CyberAgentLM2-Chat is a fine-tuned model of [CyberAgentLM2](https://huggingface.co/cyberagent/calm2-7b) for dialogue use cases.

## Requirements
- transformers >= 4.34.1
- accelerate

## Usage

```python
import transformers
from transformers import AutoModelForCausalLM, AutoTokenizer, TextStreamer

assert transformers.__version__ >= "4.34.1"

model = AutoModelForCausalLM.from_pretrained("cyberagent/calm2-7b-chat", device_map="auto", torch_dtype="auto")
tokenizer = AutoTokenizer.from_pretrained("cyberagent/calm2-7b-chat")
streamer = TextStreamer(tokenizer, skip_prompt=True, skip_special_tokens=True)

prompt = """USER: AIによって私達の暮らしはどのように変わりますか？
ASSISTANT: """

token_ids = tokenizer.encode(prompt, return_tensors="pt")
output_ids = model.generate(
    input_ids=token_ids.to(model.device),
    max_new_tokens=300,
    do_sample=True,
    temperature=0.8,
    streamer=streamer,
)
```

## Chat Template
```
USER: {user_message1}
ASSISTANT: {assistant_message1}<|endoftext|>
USER: {user_message2}
ASSISTANT: {assistant_message2}<|endoftext|>
USER: {user_message3}
ASSISTANT: {assistant_message3}<|endoftext|>
```

## Model Details

* **Model size**: 7B
* **Context length**: 32768
* **Model type**: Transformer-based Language Model
* **Language(s)**: Japanese, English
* **Developed by**: [CyberAgent, Inc.](https://www.cyberagent.co.jp/)
* **License**: Apache-2.0

## Author

[Ryosuke Ishigami](https://huggingface.co/rishigami)

## Citations
```tex
@article{touvron2023llama,
  title={LLaMA: Open and Efficient Foundation Language Models},
  author={Touvron, Hugo and Lavril, Thibaut and Izacard, Gautier and Martinet, Xavier and Lachaux, Marie-Anne and Lacroix, Timoth{\'e}e and Rozi{\`e}re, Baptiste and Goyal, Naman and Hambro, Eric and Azhar, Faisal and Rodriguez, Aurelien and Joulin, Armand and Grave, Edouard and Lample, Guillaume},
  journal={arXiv preprint arXiv:2302.13971},
  year={2023}
}
```