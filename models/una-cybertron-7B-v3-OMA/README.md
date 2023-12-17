---
license: apache-2.0
datasets:
- fblgit/tree-of-knowledge
library_name: transformers
tags:
- juanako
- UNA
- cybertron
- xaberius
---

# Model Card for una-cybertron-7b-v3 (UNA: Uniform Neural Alignment)

**OMA (One Man Army) proudly presents a new 7B Champion: `cybertron-7b-v3` with our famous UNA algorythm.**

The model excels in mathematics, logic, reasoning, overall very smart. He can make a deep reasoning over the context and prompt, it gives the impression of not missing details around.

This seems to be possible:
* UNA models can be SFT again
* UNA models are easy to be used as Merge Base, place Cybertron in the fan-in and fan-out of the layering
* UNA models now includes a digital watermark

## Model Details

Adiestrated with UNA: Uniform Neural Alignment technique (paper going out soon). 
* What is **NOT** UNA? Its not a merged layers model. Is not SLERP or SLURP or similar.
* What **is** UNA? A formula & A technique to *TAME* models

### Model Description

- **Developed by:** [juanako.ai](https://juanako.ai)
- **Author:** [Xavier M.](xavi@juanako.ai)
- **Model type:** MistralAI 7B
- **Funded by Cybertron's H100's** with few hours training.

### Prompt
The model is very good, works well on almost any prompt but ChatML format and Alpaca System gets the best
```
<|im_start|>system
- You are a helpful assistant chatbot trained by MosaicML.
- You answer questions.
- You are excited to be able to help the user, but will refuse to do anything that could be considered harmful to the user.
- You are more than just an information source, you are also able to write poetry, short stories, and make jokes.<|im_end|>
<|im_start|>user
Explain QKV<|im_end|>
<|im_start|>assistant
```
```
### Assistant: I am StableVicuna, a large language model created by CarperAI. I am here to chat!

### Human: Explain QKV
### Assistant:
```
```
[Round <|round|>]
问：Explain QKV
答：
```
```
[Round <|round|>]
Question：Explain QKV
Answer：
```
```
Question：Explain QKV
Answer：
```
Using Exllamav2_HF set alpha=2.5 for 16K Context

### Framework versions

- Transformers 4.35.0-UNA
- Pytorch 2.1.0
- Datasets 2.14.6
- Tokenizers 0.14.1

### Citations
    If you find Cybertron, Juanako or any of our models useful, specially if you use it for your big brand.. or you clone/merge my modelsm, cite please:
```
@misc{unacybertron7b,
  title={Cybertron: Uniform Neural Alignment}, 
  author={Xavier Murias},
  year={2023},
  publisher = {HuggingFace},
  journal = {HuggingFace repository},
  howpublished = {\url{https://huggingface.co/fblgit/una-cybertron-7b-v3-OMA}},
}
```

