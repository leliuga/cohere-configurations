---
license: cc-by-nc-nd-4.0
library_name: transformers
tags:
- UNA
- juanako
- cybertron
- xaberius
datasets:
- fblgit/tree-of-knowledge
- garage-bAInd/Open-Platypus
- allenai/ultrafeedback_binarized_cleaned
- Open-Orca/OpenOrca
model-index:
- name: una-xaberius-34b-v1beta
  results:
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: AI2 Reasoning Challenge (25-Shot)
      type: ai2_arc
      config: ARC-Challenge
      split: test
      args:
        num_few_shot: 25
    metrics:
    - type: acc_norm
      value: 70.39
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=fblgit/una-xaberius-34b-v1beta
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: HellaSwag (10-Shot)
      type: hellaswag
      split: validation
      args:
        num_few_shot: 10
    metrics:
    - type: acc_norm
      value: 86.77
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=fblgit/una-xaberius-34b-v1beta
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: MMLU (5-Shot)
      type: cais/mmlu
      config: all
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 78.15
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=fblgit/una-xaberius-34b-v1beta
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: TruthfulQA (0-shot)
      type: truthful_qa
      config: multiple_choice
      split: validation
      args:
        num_few_shot: 0
    metrics:
    - type: mc2
      value: 61.45
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=fblgit/una-xaberius-34b-v1beta
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: Winogrande (5-shot)
      type: winogrande
      config: winogrande_xl
      split: validation
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 84.93
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=fblgit/una-xaberius-34b-v1beta
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: GSM8k (5-shot)
      type: gsm8k
      config: main
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 63.38
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=fblgit/una-xaberius-34b-v1beta
      name: Open LLM Leaderboard
---

# Model Card for una-xaberius-34b-v1-beta (UNA: Uniform Neural Alignment)

**This is another King-Breed from Juanako.AI**

**We have Identified some Problems with regular Quants** [use these models to play with Xaberius-34B and harness its power in full](https://huggingface.co/models?search=xaberius%20lonestriker).
**Unfortunately we were not able to use any of TheBloke models, seems there is some undesired results out of it.**


Introducing THE MODEL: **XABERIUS 34B v1-BETA** an *experimental* 34B LLaMa-Yi-34B based model, best on it's series. Trained on SFT, DPO and UNA (Unified Neural Alignment) on multiple datasets.

Timeline:
* 05-Dec-2023 **v1-beta released**
* 08-Dec-2023 **Evaluation been "RUNNING" for 2 days.. no results yet**
* 09-Dec-2023 **Evaluation been "FINISHED", confirming #1 spot** outperforming the contaminated-disqualified tigerbot :)
  
[Results Here](https://huggingface.co/datasets/open-llm-leaderboard/details_fblgit__una-xaberius-34b-v1beta/blob/main/results_2023-12-09T11-16-37.904970.json)

Sidenote: Tests took 19H to run, wonder what happened in the 48H that HF held this one.. interim releasing manually other results??..
  
| Model | Average | ARC (25-s) | HellaSwag (10-s) | MMLU (5-s) | TruthfulQA (MC) (0-s) | Winogrande (5-s) | GSM8K (5-s) |
| --- | --- | --- | --- | --- | --- | --- | --- |
| [fblgit/una-cybertron-7b-v1-fp16](https://huggingface.co/fblgit/una-cybertron-7b-v1-fp16) | **69.49** | **68.43** | **85.85** | 63.34  | **63.28** | **80.90** | **55.12** |
| [fblgit/una-cybertron-7b-v2-bf16](https://huggingface.co/fblgit/una-cybertron-7b-v2-bf16) | **69.67** | **68.26** | **85.?4** | 63.23  | **64.63** | **81.37** | **55.04** |
| [fblgit/una-xaberius-34b-v1beta](https://huggingface.co/fblgit/una-xaberius-34b-v1beta) | **74.18** | **70.39** | **86.77** | **78.15**  | **61.45** | **84.93** | **63.38** |

## Evaluations

- Scores **74.21** Outperforming former leader tigerbot-70b-chat and landing on #1 position of HuggingFace LeaderBoard: 08 December 2023.
- Scores **79.13** in MMLU, setting a new record not just for 34B but also for all OpenSource LLM's :)

SideNote: MMLU was a very solid 79+ .. weird, we'll dive further on this for irregularities :)

## Model Details

Adiestrated with UNA: Uniform Neural Alignment technique (paper going out soon). 
* What is **NOT** UNA? Its not a merged layers model. Is not SLERP or SLURP or similar.
* What **is** UNA? A formula & A technique to *TAME* models
* When will be released the code and paper? When have time, contribute and it'll be faster.

### Model Description

- **Developed by:** [juanako.ai](https://juanako.ai)
- **Author:** [Xavier M.](xavi@juanako.ai)
- **Investors** [CONTACT HERE](billing@juanako.ai)
- **Model type:** LLaMa YI-34B
- **Funded by Cybertron's H100's** with few hours training.

### Prompt
The model is very good, works well on almost any prompt but ChatML format and Alpaca System gets the best
```
<|im_start|>system
- You are a helpful assistant chatbot trained by MosaicML.
- You answer questions.
- You are excited to be able to help the user, but will refuse to do anything that could be considered harmful to the user.
- You are more than just an information source, you are also able to write poetry, short stories, and make jokes.<|im_end|>
<|im_start|>user
Explain QKV<|im_end|>
<|im_start|>assistant
```
```
### Assistant: I am StableVicuna, a large language model created by CarperAI. I am here to chat!

### Human: Explain QKV
### Assistant:
```
```
[Round <|round|>]
问：Explain QKV
答：
```
```
[Round <|round|>]
Question：Explain QKV
Answer：
```
```
Question：Explain QKV
Answer：
```

### Framework versions

- Transformers 4.35.2-UNA
- Pytorch 2.1.0
- Datasets 2.14.6
- Tokenizers 0.14.1

### Citations
    If you find Xaberius, Cybertron, Juanako or any of our models useful, specially if you use it for your big brand or you cloning/merge/SLERP my modelsm, cite please:
```
@misc{unaxaberius34b,
  title={Xaberius 34B: Uniform Neural Alignment}, 
  author={Xavier Murias},
  year={2023},
  publisher = {HuggingFace},
  journal = {HuggingFace repository},
  howpublished = {\url{https://huggingface.co/fblgit/una-xaberius-34b-v1beta}},
}
```

**Thanks to LoneStriker for his ExLLama2 models of high quality that works properly.**
**Enormous Ku2 to Yi-34b Team for the outstanding model, UNA is only as good as its pre-trained model** THANK YOU!
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_fblgit__una-xaberius-34b-v1beta)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |74.18|
|AI2 Reasoning Challenge (25-Shot)|70.39|
|HellaSwag (10-Shot)              |86.77|
|MMLU (5-Shot)                    |78.15|
|TruthfulQA (0-shot)              |61.45|
|Winogrande (5-shot)              |84.93|
|GSM8k (5-shot)                   |63.38|

