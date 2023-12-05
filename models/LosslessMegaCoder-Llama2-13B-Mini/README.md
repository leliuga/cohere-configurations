---
license: llama2
datasets:
- rombodawg/LosslessMegaCodeTrainingV2_1m_Evol_Uncensored
---
___________________________
- Please note this model was not trained on the rombodawg/LosslessMegaCodeTrainingV3_MINI dataset, despite the name similarity. You can find the training data at the bottom of the model card labeled (megacode2-min100)
___________________________

This is one of the first models trained on the LosslessMegaCodeTrainingV2_1m_Evol_Uncensored dataset. The version of the dataset used for this model was filtered by removed any data with less than 100 tokens but plans for much more refined filtering are in the works

- This model was made as a colaboration between me and andreaskoepf who is an affiliate of Open Assistant.

This Model score .29 on humaneval+ the same as LLaMA-2 70B Chat Link bellow (in this benchmark the model is called andreaskoepf/llama2-13b-megacode2_min100)

- https://tju01.github.io/FastEval-OpenAssistant/

Prompt template:

- chatml format is used: "<|im_start|>system\n{system message}<|im_end|>\n<|im_start|>user\n{user prompt}<|im_end|>\n<|im_start|>assistant\n{Assistant answer}<|im_end|>\n"

multi-line:
```
<|im_start|>system
{system message}<|im_end|>
<|im_start|>user
{user prompt}<|im_end|>
<|im_start|>assistant
{Assistant answer}<|im_end|>
```
Gpt4all template:

- System prompt
```
<|im_start|>system
"Below is an instruction that describes a task. Write a response that appropriately completes the request."
```
- Prompt template
```
<|im_end|>
<|im_start|>user
"%1"<|im_end|>
<|im_start|>assistant
```

Oobagooba Text-Generation-Webui Template
- user:
```
  <|im_start|>user
  {User string}<|im_end|>
```
- bot:
```
  <|im_start|>assistant
  {Bot string}<|im_end|>
```
- turn_template:
```
<|user|>\n<|user-message|>\n\n<|bot|>\n<|bot-message|>\n\n
```
- context:
```
  <|im_start|>system
  Below is an instruction that describes a task. Write a response that appropriately completes the request.<|im_end|>
```

Current quantizations available:

- https://huggingface.co/TheBloke/LosslessMegaCoder-Llama2-13B-Mini-GPTQ

Training data:

- https://wandb.ai/open-assistant/epfl-mt-sft/runs/run34_megacode2_min100_13b

The link for the full dataset is bellow:

- https://huggingface.co/datasets/rombodawg/LosslessMegaCodeTrainingV2_1m_Evol_Uncensored

Link for the filtered dataset used to make this model are bellow:

- https://huggingface.co/datasets/andreaskoepf/megacode2-min100

The original posting for this model was uploaded at the link bellow.

- https://huggingface.co/andreaskoepf/llama2-13b-megacode2_min100
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_rombodawg__LosslessMegaCoder-llama2-13b-mini)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 49.92   |
| ARC (25-shot)         | 60.58          |
| HellaSwag (10-shot)   | 81.26    |
| MMLU (5-shot)         | 57.92         |
| TruthfulQA (0-shot)   | 48.89   |
| Winogrande (5-shot)   | 76.95   |
| GSM8K (5-shot)        | 15.92        |
| DROP (3-shot)         | 7.89         |
