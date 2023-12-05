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
license: apache-2.0
---
# LAION LeoLM: **L**inguistically **E**nhanced **O**pen **L**anguage **M**odel
Meet LeoLM, the first open and commercially available German Foundation Language Model built on Llama-2 and Mistral. 
Our models extend Llama-2's capabilities into German through continued pretraining on a large corpus of German-language and mostly locality specific text.
Thanks to a compute grant at HessianAI's new supercomputer **42**, we release three foundation models trained with 8k context length.
[`LeoLM/leo-mistral-hessianai-7b`](https://huggingface.co/LeoLM/leo-mistral-hessianai-7b) under Apache 2.0 and [`LeoLM/leo-hessianai-7b`](https://huggingface.co/LeoLM/leo-hessianai-7b) and [`LeoLM/leo-hessianai-13b`](https://huggingface.co/LeoLM/leo-hessianai-13b) under the [Llama-2 community license](https://huggingface.co/meta-llama/Llama-2-70b/raw/main/LICENSE.txt) (70b also coming soon! 👀).
With this release, we hope to bring a new wave of opportunities to German open-source and commercial LLM research and accelerate adoption.
Read our [blog post](https://laion.ai/blog/leo-lm/) or our paper (preprint coming soon) for more details!

*A project by Björn Plüster and Christoph Schuhmann in collaboration with LAION and HessianAI.*

## LeoLM Chat
`LeoLM/leo-mistral-hessianai-7b-chat` is a German chat model built on our foundation model `LeoLM/leo-mistral-hessianai-7b` and finetuned on a selection of German instruction datasets.
The model performs exceptionally well on writing, explanation and discussion tasks but struggles somewhat with math and advanced reasoning. See our MT-Bench-DE scores:
```
{
  "first_turn": 6.1,
  "second_turn": 4.7,
  "categories": {
      "writing": 6.8,
      "roleplay": 6.35,
      "reasoning": 3.3,
      "math": 2.75,
      "coding": 4.4,
      "extraction": 4.5,
      "stem": 6.85,
      "humanities": 8.25
  },
  "average": 5.4
}
```

## Model Details

- **Finetuned from:** [LeoLM/leo-mistral-hessianai-7b](https://huggingface.co/LeoLM/leo-hessianai-7b)
- **Model type:** Causal decoder-only transformer language model
- **Language:** English and German
- **Demo:** [Web Demo coming soon !]()
- **License:** [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0.html)
- **Contact:** [LAION Discord](https://discord.com/invite/eq3cAMZtCC) or [Björn Plüster](mailto:bjoern.pl@outlook.de)


## Use in 🤗Transformers
First install direct dependencies:
```
pip install transformers torch sentencepiece
```
If you want faster inference using flash-attention2, you need to install these dependencies:
```bash
pip install packaging ninja
pip install flash-attn
```
Then load the model in transformers:
```python
from transformers import pipeline
import torch

system_prompt = """Dies ist eine Unterhaltung zwischen einem intelligenten, hilfsbereitem KI-Assistenten und einem Nutzer.
Der Assistent gibt ausführliche, hilfreiche und ehrliche Antworten."""

prompt_format = "<|im_start|>system\n{system_prompt}<|im_end|>\n<|im_start|>user\n{prompt}<|im_end|>\n<|im_start|>assistant\n"
prompt = "Erkläre mir wie die Fahrradwegesituation in Hamburg ist."

generator = pipeline(model="LeoLM/leo-mistral-hessianai-7b-chat", device="cuda", torch_dtype=torch.float16, use_flash_attention_2=True) # True for flash-attn2 else False
print(generator(prompt_format.format(system_prompt=system_prompt, prompt=prompt), do_sample=True, top_p=0.95, max_length=8192))
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
For these reasons, as with all LLMs, the potential outputs of `LeoLM/leo-mistral-hessianai-7b-chat` cannot be predicted
in advance, and the model may in some instances produce inaccurate, biased or other objectionable responses
to user prompts. Therefore, before deploying any applications of `LeoLM/leo-mistral-hessianai-7b-chat`, developers should
perform safety testing and tuning tailored to their specific applications of the model.

Please see Meta's [Responsible Use Guide](https://ai.meta.com/llama/responsible-use-guide/).

## Finetuning Details

| Hyperparameter  | Value  |
|---|---|
| Num epochs | 4 |
| Examples per epoch  | 131214  |
| Global batch size | 256 |
| Learning rate  |  1e-5 |
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