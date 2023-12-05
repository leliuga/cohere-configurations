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

**Collective Cognition v1.1 - Mistral 7B**
<div style="display: flex; justify-content: center;">
  <a href="https://collectivecognition.ai" target="_blank" style="display: inline-block; text-align: center;">
    <img src="https://cdn-uploads.huggingface.co/production/uploads/6317aade83d8d2fd903192d9/DNZXsJE5oC_rM8eYY6H_x.png" alt="Collective Cognition Logo" width="50%" style="display: block; margin: 0 auto;">
  </a>
</div>

## Model Description:

Collective Cognition v1.1 is a state-of-the-art model fine-tuned using the Mistral approach. This model is particularly notable for its performance, outperforming many 70B models on the TruthfulQA benchmark. This benchmark assesses models for common misconceptions, potentially indicating hallucination rates.

## Special Features:
- **Quick Training**: This model was trained in just 3 minutes on a single 4090 with a qlora, and competes with 70B scale Llama-2 Models at TruthfulQA.
- **Limited Data**: Despite its exceptional performance, it was trained on only ONE HUNDRED data points, all of which were gathered from a platform reminiscent of ShareGPT.
- **Extreme TruthfulQA Benchmark**: This model is competing strongly with top 70B models on the TruthfulQA benchmark despite the small dataset and qlora training!

![image/png](https://cdn-uploads.huggingface.co/production/uploads/6317aade83d8d2fd903192d9/-pnifxPcMeeUONyE3efo3.png)

## Acknowledgements:

Special thanks to @a16z and all contributors to the Collective Cognition dataset for making the development of this model possible.

## Dataset:

The model was trained using data from the Collective Cognition website. The efficacy of this dataset is demonstrated by the model's stellar performance, suggesting that further expansion of this dataset could yield even more promising results. The data is reminiscent of that collected from platforms like ShareGPT.

You can contribute to the growth of the dataset by sharing your own ChatGPT chats [here](https://CollectiveCognition.ai).

You can download the datasets created by Collective Cognition here: https://huggingface.co/CollectiveCognition

## Performance:

- **TruthfulQA**: Collective Cognition v1.1 has notably outperformed various 70B models on the TruthfulQA benchmark, highlighting its ability to understand and rectify common misconceptions.


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
|truthfulqa_mc|      1|mc1   |0.4051|±  |0.0172|
|             |       |mc2   |0.5738|±  |0.0157|
```

Collective Cognition v1.1 GPT4All:
```
|    Task     |Version| Metric |Value |   |Stderr|
|-------------|------:|--------|-----:|---|-----:|
|arc_challenge|      0|acc     |0.5085|±  |0.0146|
|             |       |acc_norm|0.5384|±  |0.0146|
|arc_easy     |      0|acc     |0.7963|±  |0.0083|
|             |       |acc_norm|0.7668|±  |0.0087|
|boolq        |      1|acc     |0.8495|±  |0.0063|
|hellaswag    |      0|acc     |0.6399|±  |0.0048|
|             |       |acc_norm|0.8247|±  |0.0038|
|openbookqa   |      0|acc     |0.3240|±  |0.0210|
|             |       |acc_norm|0.4540|±  |0.0223|
|piqa         |      0|acc     |0.7992|±  |0.0093|
|             |       |acc_norm|0.8107|±  |0.0091|
|winogrande   |      0|acc     |0.7348|±  |0.0124|
Average: 71.13
```

AGIEval:
```
|             Task             |Version| Metric |Value |   |Stderr|
|------------------------------|------:|--------|-----:|---|-----:|
|agieval_aqua_rat              |      0|acc     |0.1929|±  |0.0248|
|                              |       |acc_norm|0.2008|±  |0.0252|
|agieval_logiqa_en             |      0|acc     |0.3134|±  |0.0182|
|                              |       |acc_norm|0.3333|±  |0.0185|
|agieval_lsat_ar               |      0|acc     |0.2217|±  |0.0275|
|                              |       |acc_norm|0.2043|±  |0.0266|
|agieval_lsat_lr               |      0|acc     |0.3412|±  |0.0210|
|                              |       |acc_norm|0.3216|±  |0.0207|
|agieval_lsat_rc               |      0|acc     |0.4721|±  |0.0305|
|                              |       |acc_norm|0.4201|±  |0.0301|
|agieval_sat_en                |      0|acc     |0.6068|±  |0.0341|
|                              |       |acc_norm|0.5777|±  |0.0345|
|agieval_sat_en_without_passage|      0|acc     |0.3932|±  |0.0341|
|                              |       |acc_norm|0.3641|±  |0.0336|
|agieval_sat_math              |      0|acc     |0.2864|±  |0.0305|
|                              |       |acc_norm|0.2636|±  |0.0298|
Average: 33.57
```

Training run on wandb here: https://wandb.ai/teknium1/collectivecognition-mistral-7b/runs/collectivecognition-mistral-8/workspace

## Licensing:

Apache 2.0

---

