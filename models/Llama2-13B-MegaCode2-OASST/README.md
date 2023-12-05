---
license: other
---
# llama2-13b-megacode2-oasst

- sampling report: [2023-08-15_andreaskoepf_llama2-13b-megacode2-oasst_sampling_noprefix2.json](https://open-assistant.github.io/oasst-model-eval/?f=https%3A%2F%2Fraw.githubusercontent.com%2FOpen-Assistant%2Foasst-model-eval%2Fmain%2Fsampling_reports%2Foasst-sft%2F2023-08-15_andreaskoepf_llama2-13b-megacode2-oasst_sampling_noprefix2.json)

### Prompt template

[chatml](https://github.com/openai/openai-python/blob/main/chatml.md) format is used:
"<|im_start|>user\n{user prompt}<|im_end|>\n<|im_start|>assistant\n{Assistant answer}<|im_end|>\n" 

Multi-line:

```
<|im_start|>user
{user prompt}<|im_end|>
<|im_start|>assistant
{Assistant answer}<|im_end|>
```

### Credits & Special Thanks

- Compute was generously sponsored by the eplf [Machine Learning and Optimization Laboratory](https://www.epfl.ch/labs/mlo/)
- The open-source [epfLLM/Megatron-LLM](https://github.com/epfLLM/Megatron-LLM) trainer was used for fine-tuning.
- [rombodawg](https://huggingface.co/rombodawg) curated and published [LosslessMegaCodeTrainingV2_1m_Evol_Uncensored](https://huggingface.co/datasets/rombodawg/LosslessMegaCodeTrainingV2_1m_Evol_Uncensored)
- [andreaskoepf](https://github.com/andreaskoepf/) prepared & orchestrated the training.
