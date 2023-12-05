---
datasets:
- LeoLM/OpenSchnabeltier
- OpenAssistant/OASST-DE
- FreedomIntelligence/alpaca-gpt4-deutsch
- FreedomIntelligence/evol-instruct-deutsch
- LeoLM/German_Poems
- LeoLM/German_Songs
- garage-bAInd/Open-Platypus
- WizardLM/WizardLM_evol_instruct_70k
- bjoernp/oasst25-08-23-filtered
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
`LeoLM/leo-hessianai-7b-chat-bilingual` is a bilingual English-German chat model built on our foundation model `LeoLM/leo-hessianai-7b` and finetuned on a selection of German translateed instruction datasets and their English counterparts.
The model performs exceptionally well on writing, explanation and discussion tasks but struggles somewhat with math and advanced reasoning. See our MT-Bench scores:
```
{
    "first_turn": 5.64375,
    "second_turn": 4.075,
    "categories": {
        "writing": 5.925,
        "roleplay": 5.25,
        "reasoning": 3.1,
        "math": 1.8,
        "coding": 3.4,
        "extraction": 5,
        "stem": 6.5,
        "humanities": 7.9
    },
    "average": 4.859375
}
```

## Model Details

- **Finetuned from:** [LeoLM/leo-hessianai-7b](https://huggingface.co/LeoLM/leo-hessianai-7b)
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

generator = pipeline(model="LeoLM/leo-hessianai-7b-chat-bilingual", device="cuda", torch_dtype=torch.float16, trust_remote_code=True) # True for flash-attn2 else False
print(generator(prompt_format.format(prompt=prompt), do_sample=True, top_p=0.95, max_length=8192))
```

"*Als KI kann ich keine persönlichen Beobachtungen teilen, aber ich kann einige allgemeine Informationen zur Fahrradwegesituation in Hamburg liefern. Im Vergleich zu vielen anderen großen Städten hat Hamburg eine hohe Anzahl von Fahrradfahrern und nimmt seine Verantwortung für nachhaltige Verkehrsmittel sehr ernst. Es gibt viele Fahrradwege und separate Fahrspuren, die Radfahrern ermöglichen, zusammen mit dem Autoverkehr zu fahren. Diese Fahrradspuren sind oft mit Markierungen gekennzeichnet und durch physische Trennungen von anderen Fahrspuren abgegrenzt. Darüber hinaus gibt es viele Fahrradstraßen, auf denen Radfahrer Vorfahrt haben und Autos langsamer fahren müssen.*

*In einigen städtischen Gebieten können Fahrradwege jedoch eng oder überfüllt sein, besonders während der Stoßzeiten. Es gibt auch viele Kreuzungen, an denen Radfahrer anhalten und auf Grün warten müssen, ähnlich wie Autofahrer. Insgesamt ist die Fahrradinfrastruktur in Hamburg ziemlich gut, aber wie überall gibt es immer Raum für Verbesserungen.*"

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
For these reasons, as with all LLMs, the potential outputs of `LeoLM/leo-hessianai-7b-chat` cannot be predicted
in advance, and the model may in some instances produce inaccurate, biased or other objectionable responses
to user prompts. Therefore, before deploying any applications of `LeoLM/leo-hessianai-7b-chat`, developers should
perform safety testing and tuning tailored to their specific applications of the model.

Please see Meta's [Responsible Use Guide](https://ai.meta.com/llama/responsible-use-guide/).

## Finetuning Details

| Hyperparameter  | Value  |
|---|---|
| Num epochs | 3 |
| Examples per epoch  | 233275  |
| Global batch size | 256 |
| Learning rate  |  3e-5 |
| Warmup steps  |  100 |
| LR scheduler  |  Cosine |
| Adam betas  | (0.9, 0.95)  |
| Weight decay  |  0.001 |


## Dataset Details
```
## Stats for 'Subset of LeoLM/OpenSchnabeltier' (21314 samples (100.0%))
-----------------
  Accepted: 21314/21314 (100.0%)
  Accepted tokens: 8134690
  Skipped: 0 (0.0%)
  Min tokens per sample: 25
  Max tokens per sample: 1202
  Avg tokens per sample: 381.65947264708643
-----------------

## Stats for 'Subset of garage-bAInd/Open-Platypus' (24427 samples (100.0%))
-----------------
  Accepted: 24427/24427 (100.0%)
  Accepted tokens: 9549043
  Skipped: 0 (0.0%)
  Min tokens per sample: 23
  Max tokens per sample: 5054
  Avg tokens per sample: 390.9216440823679
-----------------

## Stats for 'Subset of WizardLM/WizardLM_evol_instruct_70k' (68600 samples (100.0%))
-----------------
  Accepted: 68600/68600 (100.0%)
  Accepted tokens: 33045040
  Skipped: 0 (0.0%)
  Min tokens per sample: 18
  Max tokens per sample: 11810
  Avg tokens per sample: 481.7061224489796
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

## Stats for 'Subset of LeoLM/German_Songs' (490 samples (100.0%))
-----------------
  Accepted: 490/490 (100.0%)
  Accepted tokens: 618642
  Skipped: 0 (0.0%)
  Min tokens per sample: 747
  Max tokens per sample: 1678
  Avg tokens per sample: 1262.534693877551
-----------------


## Stats for 'Subset of LeoLM/German_Poems' (392 samples (100.0%))
-----------------
  Accepted: 392/392 (100.0%)
  Accepted tokens: 187897
  Skipped: 0 (0.0%)
  Min tokens per sample: 231
  Max tokens per sample: 826
  Avg tokens per sample: 479.3290816326531
-----------------

## Stats for 'Subset of OpenAssistant/OASST_DE' (3646 samples (100.0%))
-----------------
  Accepted: 3646/3646 (100.0%)
  Accepted tokens: 2338738
  Skipped: 0 (0.0%)
  Min tokens per sample: 29
  Max tokens per sample: 2484
  Avg tokens per sample: 641.4530992868897
-----------------

## Stats for 'Subset of bjoernp/oasst25-08-23-filtered' (8922 samples (100.0%))
-----------------
  Accepted: 8922/8922 (100.0%)
  Accepted tokens: 4526427
  Skipped: 0 (0.0%)
  Min tokens per sample: 23
  Max tokens per sample: 5407
  Avg tokens per sample: 507.3332212508406
-----------------

## Stats for 'total' (235632 samples (100.0%))
-----------------
  Accepted: 235632/235632 (100.0%)
  Accepted tokens: 115862397
  Skipped: 0 (0.0%)
  Min tokens per sample: 18
  Max tokens per sample: 11810
  Avg tokens per sample: 491.70909299246284
-----------------
```