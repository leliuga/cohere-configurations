---
metrics:
- perplexity
library_name: transformers
license: apache-2.0
language:
- en
datasets:
- emozilla/yarn-train-tokenized-8k-llama
---

# Model Card: Yarn-Llama-2-70b-32k

[Preprint (arXiv)](https://arxiv.org/abs/2309.00071)  
[GitHub](https://github.com/jquesnelle/yarn)
![yarn](https://raw.githubusercontent.com/jquesnelle/yarn/70b/data/proofpile-long-small-32k-70b.csv.png)

The authors would like to thank [LAION AI](https://laion.ai/) for their support of compute for this model.
It was trained on the [JUWELS](https://www.fz-juelich.de/en/ias/jsc/systems/supercomputers/juwels) supercomputer.

## Model Description

Nous-Yarn-Llama-2-70b-32k is a state-of-the-art language model for long context, further pretrained on long context data for 400 steps using the YaRN extension method.
It is an extension of [Llama-2-70b-hf](meta-llama/Llama-2-70b-hf) and supports a 32k token context window.

To use, pass `trust_remote_code=True` when loading the model, for example

```python
model = AutoModelForCausalLM.from_pretrained("NousResearch/Yarn-Llama-2-70b-32k",
  use_flash_attention_2=True,
  torch_dtype=torch.bfloat16,
  device_map="auto",
  trust_remote_code=True)
```

In addition you will need to use the latest version of `transformers` (until 4.35 comes out)
```sh
pip install git+https://github.com/huggingface/transformers
```

## Benchmarks

Long context benchmarks:
| Model | Context Window | 1k PPL | 2k PPL | 4k PPL | 8k PPL | 16k PPL | 32k PPL |
|-------|---------------:|-------:|--------:|------:|-------:|--------:|--------:|
| [Llama-2-70b-hf](meta-llama/Llama-2-70b-hf) | 4k | 3.71 | 3.27 | 2.96 | - | - | - |
| [Yarn-Llama-2-70b-32k](https://huggingface.co/NousResearch/Yarn-Llama-2-70b-32k) | 32k | 3.61 | 3.22 | 2.91 | 2.82 | 2.45 | 2.23 |

Short context benchmarks showing that quality degradation is minimal:
| Model | Context Window | ARC-c | MMLU | Truthful QA |
|-------|---------------:|------:|-----:|------------:|
| [Llama-2-70b-hf](meta-llama/Llama-2-70b-hf) | 4k | 67.32 | 69.83 | 44.92 |
| [Yarn-Llama-2-70b-32k](https://huggingface.co/NousResearch/Yarn-Llama-2-70b-32k) | 32k | 67.41 | 68.84 | 46.14 |

## Collaborators

 - [bloc97](https://github.com/bloc97): Methods, paper and evals
 - [@theemozilla](https://twitter.com/theemozilla): Methods, paper, model training, and evals
 - [@EnricoShippole](https://twitter.com/EnricoShippole): Model training
 - [honglu2875](https://github.com/honglu2875): Paper and evals
