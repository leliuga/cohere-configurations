---
base_model: Mihaiii/Pallas-0.5-LASER-0.5
inference: false
license: other
license_name: yi-license
license_link: https://huggingface.co/01-ai/Yi-34B/blob/main/LICENSE
metrics:
- accuracy
---

This model has a [LASER](https://pratyushasharma.github.io/laser/) intervention on [Mihaiii/Pallas-0.5-LASER-0.5](https://huggingface.co/Mihaiii/Pallas-0.5-LASER-0.5) .

# **This was just an experiment. From my testing, [Pallas-0.5](https://huggingface.co/Mihaiii/Pallas-0.5) is better than this model.**

Configs used:

- lnum: 51
- lnames: mlp (meaning: ["mlp.gate_proj.weight", "mlp.up_proj.weight", "mlp.down_proj.weight"])
- rate: 8
- dataset: bigbench (subset: causal_judgement)
- intervention type: rank-reduction


|Name|Validation acc (higher is better)|Validation logloss (lower is better)|Test acc (higher is better)|Test logloss (lower is better)|
|---|---|---|---|---|
|Pallas-0.5|55.263|1.650|60.526|1.463|
|Pallas-0.5-LASER-0.1|55.263|1.639|61.184|1.451|
|Pallas-0.5-LASER-0.2|55.263|1.646|61.184|1.458|
|Pallas-0.5-LASER-0.3|55.263|1.575|61.842|1.382|
|Pallas-0.5-LASER-0.4|55.263|1.525|61.842|1.326|
|Pallas-0.5-LASER-0.5|55.263|1.484|61.842|1.297|
|Pallas-0.5-LASER-0.6|55.263|1.455|61.184|1.283|


In order to replicate on a single A100, you can use [my branch](https://github.com/Mihaiii/laser/tree/allow-Yi-on-one-A100) (the original code will throw OOM for 34b models).

# Prompt Format:

```
SYSTEM: <ANY SYSTEM CONTEXT>
USER: 
ASSISTANT:
```