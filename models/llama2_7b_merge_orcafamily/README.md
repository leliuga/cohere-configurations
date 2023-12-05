---
license: mit
datasets:
- Open-Orca/SlimOrca
- beaugogh/openorca-multiplechoice-10k
language:
- en
metrics:
- accuracy
---


This model is based on the LLama 7b model as a backbone, and datasets from various Orcas have been fine-tuned and merged.


The three models were combined, and the model with the best ARC and MMLU performance was given the highest weight.


First: fine-tuning beaugogh/openorca-multiplechoice-10k on llama2 7b, but using the NEFTune method.


Second: model fine-tuned with the SlimOrca dataset on llama2 7b.

Third : Model with beaugogh/openorca-multiplechoice-10k fine-tuned on llama2 7b.



We'll add the results once we have the official results