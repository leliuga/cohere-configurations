---
license: cc-by-nc-4.0
tags:
- merge
- mergekit
- lazymergekit
- dpo
- rlhf
base_model: mlabonne/Beagle14-7B
model-index:
- name: NeuralBeagle14-7B
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
      value: 72.95
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralBeagle14-7B
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
      value: 88.34
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralBeagle14-7B
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
      value: 64.55
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralBeagle14-7B
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
      value: 69.93
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralBeagle14-7B
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
      value: 82.4
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralBeagle14-7B
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
      value: 70.28
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/NeuralBeagle14-7B
      name: Open LLM Leaderboard
---

![](https://i.imgur.com/89ZAKcn.png)

# 🐶 NeuralBeagle14-7B

**Update 01/16/24: NeuralBeagle14-7B is (probably) the best 7B model you can find! 🎉**

NeuralBeagle14-7B is a DPO fine-tune of [mlabonne/Beagle14-7B](https://huggingface.co/mlabonne/Beagle14-7B) using the [argilla/distilabel-intel-orca-dpo-pairs](https://huggingface.co/datasets/argilla/distilabel-intel-orca-dpo-pairs) preference dataset and my DPO notebook from [this article](https://towardsdatascience.com/fine-tune-a-mistral-7b-model-with-direct-preference-optimization-708042745aac).

It is based on a merge of the following models using [LazyMergekit](https://colab.research.google.com/drive/1obulZ1ROXHjYLn6PPZJwRR6GzgQogxxb?usp=sharing):
* [fblgit/UNA-TheBeagle-7b-v1](https://huggingface.co/fblgit/UNA-TheBeagle-7b-v1), based on jondurbin's [repo](https://github.com/jondurbin/bagel) and [jondurbin/bagel-v0.3](https://huggingface.co/datasets/jondurbin/bagel-v0.3])
* [argilla/distilabeled-Marcoro14-7B-slerp](https://huggingface.co/argilla/distilabeled-Marcoro14-7B-slerp), based on [mlabonne/Marcoro14-7B-slerp](https://huggingface.co/mlabonne/Marcoro14-7B-slerp)

Thanks [Argilla](https://huggingface.co/argilla) for providing the dataset and the training recipe [here](https://huggingface.co/argilla/distilabeled-Marcoro14-7B-slerp). 💪

You can try it out in this [Space](https://huggingface.co/spaces/mlabonne/NeuralBeagle14-7B-GGUF-Chat) (GGUF Q4_K_M).

## 🔍 Applications

This model uses a context window of 8k. It is compatible with different templates, like chatml and Llama's chat template.

Compared to other 7B models, it displays good performance in instruction following and reasoning tasks. It can also be used for RP and storytelling.

## ⚡ Quantized models

* **GGUF**: https://huggingface.co/mlabonne/NeuralBeagle14-7B-GGUF
* **GPTQ**: https://huggingface.co/TheBloke/NeuralBeagle14-7B-GPTQ
* **AWQ**: https://huggingface.co/TheBloke/NeuralBeagle14-7B-AWQ
* **EXL2**: https://huggingface.co/LoneStriker/NeuralBeagle14-7B-8.0bpw-h8-exl2

## 🏆 Evaluation

### Open LLM Leaderboard

NeuralBeagle14-7B ranks first on the Open LLM Leaderboard in the ~7B category.

![](https://i.imgur.com/4nAzJsr.png)

It has the same average score as Beagle14-7B ("Show merges"), which could be due to might be due to an unlucky run.
I think I might be overexploiting argilla/distilabel-intel-orca-dpo-pairs at this point, since this dataset or its original version are present in multiple models.
I need to find more high-quality preference data for the next DPO merge.

Note that some models like udkai/Turdus and nfaheem/Marcoroni-7b-DPO-Merge are unfortunately contaminated on purpose (see the very high Winogrande score).

### Nous

The evaluation was performed using [LLM AutoEval](https://github.com/mlabonne/llm-autoeval) on Nous suite. It is the best 7B model to date.

| Model | Average | AGIEval | GPT4All | TruthfulQA | Bigbench |
|---|---:|---:|---:|---:|---:|
| [**mlabonne/NeuralBeagle14-7B**](https://huggingface.co/mlabonne/NeuralBeagle14-7B) [📄](https://gist.github.com/mlabonne/ad0c665bbe581c8420136c3b52b3c15c) | **60.25** | **46.06** | **76.77** | **70.32** | **47.86** |
| [mlabonne/Beagle14-7B](https://huggingface.co/mlabonne/Beagle14-7B) [📄](https://gist.github.com/mlabonne/f5a5bf8c0827bbec2f05b97cc62d642c) | 59.4 | 44.38 | 76.53 | 69.44 | 47.25 |
| [mlabonne/NeuralDaredevil-7B](https://huggingface.co/mlabonne/NeuralDaredevil-7B) [📄](https://gist.github.com/mlabonne/cbeb077d1df71cb81c78f742f19f4155) | 59.39 | 45.23 | 76.2 | 67.61 | 48.52 |
| [argilla/distilabeled-Marcoro14-7B-slerp](https://huggingface.co/argilla/distilabeled-Marcoro14-7B-slerp) [📄](https://gist.github.com/mlabonne/9082c4e59f4d3f3543c5eda3f4807040) | 58.93 | 45.38 | 76.48 | 65.68 | 48.18 |
| [mlabonne/NeuralMarcoro14-7B](https://huggingface.co/mlabonne/NeuralMarcoro14-7B) [📄](https://gist.github.com/mlabonne/b31572a4711c945a4827e7242cfc4b9d) | 58.4 | 44.59 | 76.17 | 65.94 | 46.9 |
| [openchat/openchat-3.5-0106](https://huggingface.co/openchat/openchat-3.5-0106) [📄](https://gist.github.com/mlabonne/1afab87b543b0717ec08722cf086dcc3) | 53.71 | 44.17 | 73.72 | 52.53 | 44.4 |
| [teknium/OpenHermes-2.5-Mistral-7B](https://huggingface.co/teknium/OpenHermes-2.5-Mistral-7B) [📄](https://gist.github.com/mlabonne/88b21dd9698ffed75d6163ebdc2f6cc8) | 52.42 | 42.75 | 72.99 | 52.99 | 40.94 |

You can find the complete benchmark on [YALL - Yet Another LLM Leaderboard](https://huggingface.co/spaces/mlabonne/Yet_Another_LLM_Leaderboard).

## 💻 Usage

```python
!pip install -qU transformers accelerate

from transformers import AutoTokenizer
import transformers
import torch

model = "mlabonne/NeuralBeagle14-7B"
messages = [{"role": "user", "content": "What is a large language model?"}]

tokenizer = AutoTokenizer.from_pretrained(model)
prompt = tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True)
pipeline = transformers.pipeline(
    "text-generation",
    model=model,
    torch_dtype=torch.float16,
    device_map="auto",
)

outputs = pipeline(prompt, max_new_tokens=256, do_sample=True, temperature=0.7, top_k=50, top_p=0.95)
print(outputs[0]["generated_text"])
```

<p align="center">
  <a href="https://github.com/argilla-io/distilabel">
    <img src="https://raw.githubusercontent.com/argilla-io/distilabel/main/docs/assets/distilabel-badge-light.png" alt="Built with Distilabel" width="200" height="32"/>
  </a>
</p>
