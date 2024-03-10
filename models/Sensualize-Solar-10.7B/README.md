---
license: cc-by-nc-4.0
language:
- en
base_model:
  - upstage/SOLAR-10.7B-v1.0
---

A finetune of Base Solar. Took 12 Hours or so on 2x RTX 6000 ADAs, this is an 8-bit LoRA.

This is meant to be a verbose, smart ERP model.

Experimental.

***

### Prompt Format: Alpaca

```
### Instruction:
<Prompt>

### Input:
<Insert Context Here>

### Response:

```