---
license: apache-2.0
datasets:
- WizardLM/WizardLM_evol_instruct_V2_196k
language:
- en
library_name: transformers
---

Trained on 1 epoch of the WizardLM_evol_instruct_v2_196k dataset

Link to [GGUF](https://huggingface.co/maddes8cht/harborwater-open-llama-3b-v2-wizard-evol-instuct-v2-196k-gguf) formats.

Prompt template:
```
### HUMAN:
{prompt}

### RESPONSE:
<leave a newline for the model to answer>
```

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_harborwater__open-llama-3b-v2-wizard-evol-instuct-v2-196k)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 36.33   |
| ARC (25-shot)         | 41.81          |
| HellaSwag (10-shot)   | 73.01    |
| MMLU (5-shot)         | 26.36         |
| TruthfulQA (0-shot)   | 38.99   |
| Winogrande (5-shot)   | 66.69   |
| GSM8K (5-shot)        | 1.9        |
| DROP (3-shot)         | 5.57         |
