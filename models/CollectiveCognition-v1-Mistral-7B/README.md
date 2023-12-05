---
base_model: mistralai/Mistral-7B-v0.1
tags:
- mistral-7b
- instruct
- finetune
- gpt4
- synthetic data
- distillation
- sharegpt
datasets:
- CollectiveCognition/chats-data-2023-09-27
model-index:
- name: CollectiveCognition-v1-Mistral-7B
  results: []
license: apache-2.0
language:
- en
---

**Collective Cognition v1 - Mistral 7B**
<div style="display: flex; justify-content: center;">
  <a href="https://collectivecognition.ai" target="_blank" style="display: inline-block; text-align: center;">
    <img src="https://cdn-uploads.huggingface.co/production/uploads/6317aade83d8d2fd903192d9/DNZXsJE5oC_rM8eYY6H_x.png" alt="Collective Cognition Logo" width="50%" style="display: block; margin: 0 auto;">
  </a>
</div>

## Model Description:

Collective Cognition v1 is a Mistral model fine-tuned using just 100 GPT-4 chats shared on Collective Cognition.

## Special Features:
- **Quick Training**: This model was trained in just 3 minutes on a single 4090 with a qlora, and competes with 70B scale Llama-2 Models at TruthfulQA.
- **Limited Data**: Despite its exceptional performance, it was trained on only ONE HUNDRED data points, all of which were gathered from Collective Cognition, a platform reminiscent of ShareGPT.
- **Extreme TruthfulQA Benchmark**: The collective cognition models are competing strongly with top 70B models on the TruthfulQA benchmark despite the small dataset and qlora training!

![image/png](https://cdn-uploads.huggingface.co/production/uploads/6317aade83d8d2fd903192d9/-pnifxPcMeeUONyE3efo3.png)

## Acknowledgements:

Special thanks to @a16z and all contributors to the Collective Cognition dataset for making the development of this model possible.


## Dataset:

The model was trained using data from the Collective Cognition website. The efficacy of this dataset is demonstrated by the model's stellar performance, suggesting that further expansion of this dataset could yield even more promising results. The data is reminiscent of that collected from platforms like ShareGPT.

You can contribute to the growth of the dataset by sharing your own ChatGPT chats [here](https://CollectiveCognition.ai).

You can download the datasets created by Collective Cognition here: https://huggingface.co/CollectiveCognition

## Performance:

- **TruthfulQA**: Collective Cognition v1 and v1.1 in particular have notably outperformed several models on the TruthfulQA benchmark, highlighting its ability to understand and rectify common misconceptions.

The model follows a LIMA approach, by minimizing the base model's original training as little as possible and giving a small but very high quality dataset to enhance it's performance and style.

## Usage:

Prompt Format:
```
USER: <prompt>
ASSISTANT:
```
OR
```
<system message>
USER: <prompt>
ASSISTANT:
```

## Benchmarks:

Collective Cognition v1.0 TruthfulQA:
```
|    Task     |Version|Metric|Value |   |Stderr|
|-------------|------:|------|-----:|---|-----:|
|truthfulqa_mc|      1|mc1   |0.3794|±  |0.0170|
|             |       |mc2   |0.5394|±  |0.0158|
```

GPT4All Benchmark Suite:
```
Collective Cognition v1.0 GPT4All:
|    Task     |Version| Metric |Value |   |Stderr|
|-------------|------:|--------|-----:|---|-----:|
|arc_challenge|      0|acc     |0.5401|±  |0.0146|
|             |       |acc_norm|0.5572|±  |0.0145|
|arc_easy     |      0|acc     |0.8102|±  |0.0080|
|             |       |acc_norm|0.7992|±  |0.0082|
|boolq        |      1|acc     |0.8538|±  |0.0062|
|hellaswag    |      0|acc     |0.6459|±  |0.0048|
|             |       |acc_norm|0.8297|±  |0.0038|
|openbookqa   |      0|acc     |0.3380|±  |0.0212|
|             |       |acc_norm|0.4360|±  |0.0222|
|piqa         |      0|acc     |0.8085|±  |0.0092|
|             |       |acc_norm|0.8232|±  |0.0089|
|winogrande   |      0|acc     |0.7451|±  |0.0122|
Average: 72.06%
```

AGIEval:
```
|             Task             |Version| Metric |Value |   |Stderr|
|------------------------------|------:|--------|-----:|---|-----:|
|agieval_aqua_rat              |      0|acc     |0.1890|±  |0.0246|
|                              |       |acc_norm|0.2047|±  |0.0254|
|agieval_logiqa_en             |      0|acc     |0.2611|±  |0.0172|
|                              |       |acc_norm|0.3134|±  |0.0182|
|agieval_lsat_ar               |      0|acc     |0.2087|±  |0.0269|
|                              |       |acc_norm|0.2217|±  |0.0275|
|agieval_lsat_lr               |      0|acc     |0.3373|±  |0.0210|
|                              |       |acc_norm|0.3196|±  |0.0207|
|agieval_lsat_rc               |      0|acc     |0.4201|±  |0.0301|
|                              |       |acc_norm|0.3978|±  |0.0299|
|agieval_sat_en                |      0|acc     |0.5971|±  |0.0343|
|                              |       |acc_norm|0.5631|±  |0.0346|
|agieval_sat_en_without_passage|      0|acc     |0.4029|±  |0.0343|
|                              |       |acc_norm|0.3398|±  |0.0331|
|agieval_sat_math              |      0|acc     |0.3045|±  |0.0311|
|                              |       |acc_norm|0.2864|±  |0.0305|
Average: 33.08%
```  

Training run on wandb here: https://wandb.ai/teknium1/collectivecognition-mistral-7b/runs/collectivecognition-mistral-6/workspace


## Licensing:

Apache 2.0

---