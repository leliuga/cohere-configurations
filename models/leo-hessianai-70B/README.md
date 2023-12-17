---
datasets:
- oscar-corpus/OSCAR-2301
- wikipedia
- bjoernp/tagesschau-2018-2023
language:
- en
- de
library_name: transformers
pipeline_tag: text-generation
license: llama2
---
# LAION LeoLM 70b: **L**inguistically **E**nhanced **O**pen **L**anguage **M**odel
Meet LeoLM, the first open and commercially available German Foundation Language Model built on Llama-2. 
Our models extend Llama-2's capabilities into German through continued pretraining on a large corpus of German-language and mostly locality specific text.
Thanks to a compute grant at HessianAI's new supercomputer **42**, we release a series foundation models trained with 8k context length
under the [Llama-2 community license](https://huggingface.co/meta-llama/Llama-2-70b/raw/main/LICENSE.txt). Now, we're finally releasing the
much anticipated `leo-hessianai-70b`, the largest model of this series based on `Llama-2-70b`.
With this release, we hope to bring a new wave of opportunities to German open-source and commercial LLM research and accelerate adoption.
Read our [blog post](https://laion.ai/blog/leo-lm/) or our paper (preprint coming soon) for more details!

*A project by Björn Plüster and Christoph Schuhmann in collaboration with LAION and HessianAI.*


## Model Details
- **Finetuned from:** [meta-llama/Llama-2-70b-hf](https://huggingface.co/meta-llama/Llama-2-70b-hf)
- **Model type:** Causal decoder-only transformer language model
- **Language:** English and German
- **License:** [LLAMA 2 COMMUNITY LICENSE AGREEMENT](https://huggingface.co/meta-llama/Llama-2-70b/raw/main/LICENSE.txt)
- **Contact:** [LAION Discord](https://discord.com/invite/eq3cAMZtCC) or [Björn Plüster](mailto:bjoern.pl@outlook.de)


## Use in 🤗Transformers
First install direct dependencies:
```
pip install transformers torch
```

Then load the model in transformers. Note that this requires lots of VRAM and most-likely multiple devices. Use `load_in_8bit=True` or `load_in_4bit=True`
to save some memory by using a quantized version. For more quantized versions, check out our models at TheBloke's page: (coming soon!)
```python
from transformers import AutoModelForCausalLM, AutoTokenizer
import torch

model = AutoModelForCausalLM.from_pretrained(
    model="LeoLM/leo-hessianai-70b",
    device_map="auto",
    torch_dtype=torch.bfloat16,
    use_flash_attention_2=False   # Set to true to use FA2. Requires `pip install flash-attn`
)
```

## Training parameters
![training_parameters](imgs/hyperparams.png "Training Hyperparameters")


## Benchmarks
![benchmarks](imgs/benchmarks.png "Benchmark Scores")
![benchmarks](imgs/translation_scores.png "Translation Benchmark Scores")