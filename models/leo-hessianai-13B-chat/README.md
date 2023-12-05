---
datasets:
- LeoLM/OpenSchnabeltier
- OpenAssistant/OASST-DE
- FreedomIntelligence/alpaca-gpt4-deutsch
- FreedomIntelligence/evol-instruct-deutsch
- LeoLM/German_Poems
- LeoLM/German_Songs
language:
- en
- de
library_name: transformers
pipeline_tag: text-generation
---
# LAION LeoLM: **L**inguistically **E**nhanced **O**pen **L**anguage **M**odel
Meet LeoLM, the first open and commercially available German Foundation Language Model built on Llama-2. 
Our models extend Llama-2's capabilities into German through continued pretraining on a large corpus of German-language and mostly locality specific text.
Thanks to a compute grant at HessianAI's new supercomputer **42**, we release two foundation models trained with 8k context length,
[`LeoLM/leo-hessianai-7b`](https://huggingface.co/LeoLM/leo-hessianai-7b) and [`LeoLM/leo-hessianai-13b`](https://huggingface.co/LeoLM/leo-hessianai-13b) under the [Llama-2 community license](https://huggingface.co/meta-llama/Llama-2-70b/raw/main/LICENSE.txt) (70b also coming soon! 👀).
With this release, we hope to bring a new wave of opportunities to German open-source and commercial LLM research and accelerate adoption.
Read our [blog post]() or our paper (preprint coming soon) for more details!

*A project by Björn Plüster and Christoph Schuhmann in collaboration with LAION and HessianAI.*

## LeoLM Chat
`LeoLM/leo-hessianai-13b-chat` is a German chat model built on our foundation model `LeoLM/leo-hessianai-13b` and finetuned on a selection of German instruction datasets.
The model performs exceptionally well on writing, explanation and discussion tasks but struggles somewhat with math and advanced reasoning. See our MT-Bench-DE scores:
```
{
    "first_turn": 6.525,
    "second_turn": 5.15,
    "categories": {
        "writing": 6.925,
        "roleplay": 6.7,
        "reasoning": 4.55,
        "math": 3.25,
        "coding": 3.45,
        "extraction": 5.4,
        "stem": 7.55,
        "humanities": 8.875
    },
    "average": 5.8375
}
```

## Model Details

- **Finetuned from:** [LeoLM/leo-hessianai-13b](https://huggingface.co/LeoLM/leo-hessianai-7b)
- **Model type:** Causal decoder-only transformer language model
- **Language:** English and German
- **Demo:** [Web Demo]()
- **License:** [LLAMA 2 COMMUNITY LICENSE AGREEMENT](https://huggingface.co/meta-llama/Llama-2-70b/raw/main/LICENSE.txt)
- **Contact:** [LAION Discord](https://discord.com/invite/eq3cAMZtCC) or [Björn Plüster](mailto:bjoern.pl@outlook.de)


## Use in 🤗Transformers
First install direct dependencies:
```
pip install transformers torch sentencepiece
```
If you want faster inference using flash-attention2, you need to install these dependencies:
```bash
pip install packaging ninja
pip install flash-attn==v2.1.1 --no-build-isolation
pip install git+https://github.com/HazyResearch/flash-attention.git@v2.1.1#subdirectory=csrc/rotary
```
Then load the model in transformers:
```python
from transformers import pipeline
import torch

system_prompt = """<|im_start|>system
Dies ist eine Unterhaltung zwischen einem intelligenten, hilfsbereitem KI-Assistenten und einem Nutzer.
Der Assistent gibt ausführliche, hilfreiche und ehrliche Antworten.<|im_end|>

"""
prompt_format = "<|im_start|>user\n{prompt}<|im_end|>\n<|im_start|>assistant\n"
prompt = "Erkläre mir wie die Fahrradwegesituation in Hamburg ist."

generator = pipeline(model="LeoLM/leo-hessianai-13b-chat", device="cuda", torch_dtype=torch.float16, trust_remote_code=True) # True for flash-attn2 else False
print(generator(prompt_format.format(prompt=prompt), do_sample=True, top_p=0.95, max_length=8192))
```

"*Hamburg verfügt über ein umfangreiches Netzwerk an Radwegen, das ein effizientes und angenehmes Radfahren in der ganzen Stadt ermöglicht. Die Radwege variieren in Bezug auf ihre Qualität, wobei einige erstklassig mit eigens für Radfahrer reservierten Spuren sind, während andere einfache Fahrradsymbole auf dem Boden haben, die anzeigen, dass Radfahrer abwechselnd mit dem Autoverkehr auf der Straße fahren sollten. Einige Nebenstraßen haben auch spezielle Fahrradspuren, wobei einige mit Bordsteinabsenkungen zur Seite der Autospuren markiert sind. Zusätzlich haben viele Hauptstraßen, insbesondere in Nebenstraßen, fahrradfreundliche Abstände zwischen den geparkten Autos und dem Gehweg, was ein bequemes Fahren auf der Straße ermöglicht. Der Bau von Radschnellwegen, die schnelles und effizientes Radfahren in und aus der Stadt ermöglichen, ist im Gange und wird in den kommenden Jahren fortgesetzt. Insgesamt sind die Radwege in Hamburg weitläufig und gut ausgeschildert, was es zu einem angenehmen Ort macht, um mit dem Fahrrad zu fahren.*"

## Prompting / Prompt Template

Prompt dialogue template (ChatML format):

```
"""
<|im_start|>system
{system_message}<|im_end|>
<|im_start|>user
{prompt}<|im_end|>
<|im_start|>assistant
"""
```

The model input can contain multiple conversation turns between user and assistant, e.g.
```
<|im_start|>user
{prompt 1}<|im_end|>
<|im_start|>assistant
{reply 1}<|im_end|>
<|im_start|>user
{prompt 2}<|im_end|>
<|im_start|>assistant
(...)
```

## Ethical Considerations and Limitations

LeoLM has been tested in English and German, and has not covered, nor could it cover all scenarios. 
For these reasons, as with all LLMs, the potential outputs of `LeoLM/leo-hessianai-13b-chat` cannot be predicted
in advance, and the model may in some instances produce inaccurate, biased or other objectionable responses
to user prompts. Therefore, before deploying any applications of `LeoLM/leo-hessianai-13b-chat`, developers should
perform safety testing and tuning tailored to their specific applications of the model.

Please see Meta's [Responsible Use Guide](https://ai.meta.com/llama/responsible-use-guide/).

## Finetuning Details

| Hyperparameter  | Value  |
|---|---|
| Num epochs | 3 |
| Examples per epoch  | 131214  |
| Global batch size | 256 |
| Learning rate  |  3e-5 |
| Warmup steps  |  100 |
| LR scheduler  |  Cosine |
| Adam betas  | (0.9, 0.95)  |


## Dataset Details
```
## Stats for 'Subset of OpenAssistant/OASST-DE' (3534 samples (100.0%))
-----------------
  Accepted: 3534/3534 (100.0%)
  Accepted tokens: 2259302
  Skipped: 0 (0.0%)
  Min tokens per sample: 29
  Max tokens per sample: 2484
  Avg tokens per sample: 639.3044708545557
-----------------

## Stats for 'Subset of FreedomIntelligence/evol-instruct-deutsch' (57841 samples (100.0%))
-----------------
  Accepted: 57841/57841 (100.0%)
  Accepted tokens: 42958192
  Skipped: 0 (0.0%)
  Min tokens per sample: 33
  Max tokens per sample: 5507
  Avg tokens per sample: 742.6944900675991
-----------------

## Stats for 'Subset of FreedomIntelligence/alpaca-gpt4-deutsch' (48969 samples (100.0%))
-----------------
  Accepted: 48969/48969 (100.0%)
  Accepted tokens: 13372005
  Skipped: 0 (0.0%)
  Min tokens per sample: 19
  Max tokens per sample: 1359
  Avg tokens per sample: 273.07082031489307
-----------------

## Stats for 'Subset of LeoLM/OpenSchnabeltier' (21314 samples (100.0%))
-----------------
  Accepted: 21314/21314 (100.0%)
  Accepted tokens: 8134690
  Skipped: 0 (0.0%)
  Min tokens per sample: 25
  Max tokens per sample: 1202
  Avg tokens per sample: 381.65947264708643
-----------------

## Stats for 'Subset of LeoLM/German_Poems' (490 samples (100.0%))
-----------------
  Accepted: 490/490 (100.0%)
  Accepted tokens: 618642
  Skipped: 0 (0.0%)
  Min tokens per sample: 747
  Max tokens per sample: 1678
  Avg tokens per sample: 1262.534693877551
-----------------

## Stats for 'Subset of LeoLM/German_Songs' (392 samples (100.0%))
-----------------
  Accepted: 392/392 (100.0%)
  Accepted tokens: 187897
  Skipped: 0 (0.0%)
  Min tokens per sample: 231
  Max tokens per sample: 826
  Avg tokens per sample: 479.3290816326531
-----------------

## Stats for 'total' (132540 samples (100.0%))
-----------------
  Accepted: 132540/132540 (100.0%)
  Accepted tokens: 67530728
  Skipped: 0 (0.0%)
  Min tokens per sample: 19
  Max tokens per sample: 5507
  Avg tokens per sample: 509.51205673758864
-----------------
```