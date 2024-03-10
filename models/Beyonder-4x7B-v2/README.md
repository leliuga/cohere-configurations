---
license: other
tags:
- moe
- merge
- mergekit
- Mistral
- openchat/openchat-3.5-1210
- beowolx/CodeNinja-1.0-OpenChat-7B
- maywell/PiVoT-0.1-Starling-LM-RP
- WizardLM/WizardMath-7B-V1.1
license_name: microsoft-research-license
license_link: https://huggingface.co/WizardLM/WizardMath-7B-V1.1/resolve/main/LICENSE
model-index:
- name: Beyonder-4x7B-v2
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
      value: 68.77
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/Beyonder-4x7B-v2
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
      value: 86.8
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/Beyonder-4x7B-v2
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
      value: 65.1
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/Beyonder-4x7B-v2
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
      value: 60.68
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/Beyonder-4x7B-v2
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
      value: 80.9
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/Beyonder-4x7B-v2
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
      value: 71.72
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=mlabonne/Beyonder-4x7B-v2
      name: Open LLM Leaderboard
---

![](https://i.imgur.com/vq1QHEA.jpg)

# Beyonder-4x7B-v2

This model is a Mixture of Experts (MoE) made with [mergekit](https://github.com/cg123/mergekit) (mixtral branch). It uses the following base models:
 * [openchat/openchat-3.5-1210](https://huggingface.co/openchat/openchat-3.5-1210)
 * [beowolx/CodeNinja-1.0-OpenChat-7B](https://huggingface.co/beowolx/CodeNinja-1.0-OpenChat-7B)
 * [maywell/PiVoT-0.1-Starling-LM-RP](https://huggingface.co/maywell/PiVoT-0.1-Starling-LM-RP)
 * [WizardLM/WizardMath-7B-V1.1](https://huggingface.co/WizardLM/WizardMath-7B-V1.1)

The recommended context length is 8k.

## ⚡ Quantized models

Thanks to TheBloke for the quantized models:

* **GGUF**: https://huggingface.co/TheBloke/Beyonder-4x7B-v2-GGUF
* **AWQ**: https://huggingface.co/TheBloke/Beyonder-4x7B-v2-AWQ
* **GPTQ**: https://huggingface.co/TheBloke/Beyonder-4x7B-v2-GPTQ
* **EXL2**: https://huggingface.co/bartowski/Beyonder-4x7B-v2-exl2

## 🏆 Evaluation

Beyonder-4x7B-v2 is competitive with Mixtral-8x7B-Instruct-v0.1 on the Open LLM Leaderboard, while only having 4 experts instead of 8.

![](https://i.imgur.com/5raBff0.png)

It also displays a significant improvement over the individual experts.

![](https://i.imgur.com/7Idwkb0.png)

It also performs very well compared to other models on Nous benchmark suite. It's almost as good as the best Yi-34B fine-tune, which is a much bigger model: 24.2B parameters + only two experts are selected during inference (so ~12B) vs. 34B param.

|                               Model                                |AGIEval|GPT4All|TruthfulQA|Bigbench|Average|
|--------------------------------------------------------------------|------:|------:|---------:|-------:|------:|
|[**Beyonder-4x7B-v2**](https://huggingface.co/shadowml/Beyonder-4x7B-v2)|  **45.29**|  **75.95**|     <u>**60.86**</u>|    **46.4**|  **57.13**|
|[NeuralHermes-2.5-Mistral-7B](https://huggingface.co/mlabonne/NeuralHermes-2.5-Mistral-7B)|  43.67|  73.24|     55.37|   41.76|  53.51|
|[OpenHermes-2.5-Mistral-7B](https://huggingface.co/teknium/OpenHermes-2.5-Mistral-7B)|  42.75|  72.99|     52.99|   40.94|  52.42|
|[Nous-Hermes-2-SOLAR-10.7B](https://huggingface.co/NousResearch/Nous-Hermes-2-SOLAR-10.7B)|  47.79|  74.69|     55.92|   44.84|  55.81|
|[Nous-Hermes-2-Yi-34B](https://huggingface.co/NousResearch/Nous-Hermes-2-SOLAR-10.7B)|  <u>50.27</u>|  <u>76.00</u>|     60.34|   <u>46.69</u>|  <u>58.33</u>|

### AGIEval
|             Task             |Version| Metric |Value|   |Stderr|
|------------------------------|------:|--------|----:|---|-----:|
|agieval_aqua_rat              |      0|acc     |23.62|±  |  2.67|
|                              |       |acc_norm|23.62|±  |  2.67|
|agieval_logiqa_en             |      0|acc     |41.47|±  |  1.93|
|                              |       |acc_norm|43.01|±  |  1.94|
|agieval_lsat_ar               |      0|acc     |23.04|±  |  2.78|
|                              |       |acc_norm|23.48|±  |  2.80|
|agieval_lsat_lr               |      0|acc     |51.57|±  |  2.22|
|                              |       |acc_norm|52.94|±  |  2.21|
|agieval_lsat_rc               |      0|acc     |64.31|±  |  2.93|
|                              |       |acc_norm|64.68|±  |  2.92|
|agieval_sat_en                |      0|acc     |79.13|±  |  2.84|
|                              |       |acc_norm|79.13|±  |  2.84|
|agieval_sat_en_without_passage|      0|acc     |43.20|±  |  3.46|
|                              |       |acc_norm|43.20|±  |  3.46|
|agieval_sat_math              |      0|acc     |34.55|±  |  3.21|
|                              |       |acc_norm|32.27|±  |  3.16|

### GPT4All
|    Task     |Version| Metric |Value|   |Stderr|
|-------------|------:|--------|----:|---|-----:|
|arc_challenge|      0|acc     |61.86|±  |  1.42|
|             |       |acc_norm|64.51|±  |  1.40|
|arc_easy     |      0|acc     |85.06|±  |  0.73|
|             |       |acc_norm|82.45|±  |  0.78|
|boolq        |      1|acc     |88.35|±  |  0.56|
|hellaswag    |      0|acc     |68.04|±  |  0.47|
|             |       |acc_norm|85.12|±  |  0.36|
|openbookqa   |      0|acc     |37.80|±  |  2.17|
|             |       |acc_norm|48.60|±  |  2.24|
|piqa         |      0|acc     |83.08|±  |  0.87|
|             |       |acc_norm|83.95|±  |  0.86|
|winogrande   |      0|acc     |78.69|±  |  1.15|

### TruthfulQA
|    Task     |Version|Metric|Value|   |Stderr|
|-------------|------:|------|----:|---|-----:|
|truthfulqa_mc|      1|mc1   |44.55|±  |  1.74|
|             |       |mc2   |60.86|±  |  1.57|

### Bigbench
|                      Task                      |Version|       Metric        |Value|   |Stderr|
|------------------------------------------------|------:|---------------------|----:|---|-----:|
|bigbench_causal_judgement                       |      0|multiple_choice_grade|58.95|±  |  3.58|
|bigbench_date_understanding                     |      0|multiple_choice_grade|66.40|±  |  2.46|
|bigbench_disambiguation_qa                      |      0|multiple_choice_grade|48.84|±  |  3.12|
|bigbench_geometric_shapes                       |      0|multiple_choice_grade|22.56|±  |  2.21|
|                                                |       |exact_str_match      |13.37|±  |  1.80|
|bigbench_logical_deduction_five_objects         |      0|multiple_choice_grade|30.40|±  |  2.06|
|bigbench_logical_deduction_seven_objects        |      0|multiple_choice_grade|20.57|±  |  1.53|
|bigbench_logical_deduction_three_objects        |      0|multiple_choice_grade|52.00|±  |  2.89|
|bigbench_movie_recommendation                   |      0|multiple_choice_grade|44.40|±  |  2.22|
|bigbench_navigate                               |      0|multiple_choice_grade|52.10|±  |  1.58|
|bigbench_reasoning_about_colored_objects        |      0|multiple_choice_grade|69.75|±  |  1.03|
|bigbench_ruin_names                             |      0|multiple_choice_grade|55.36|±  |  2.35|
|bigbench_salient_translation_error_detection    |      0|multiple_choice_grade|23.65|±  |  1.35|
|bigbench_snarks                                 |      0|multiple_choice_grade|77.35|±  |  3.12|
|bigbench_sports_understanding                   |      0|multiple_choice_grade|73.02|±  |  1.41|
|bigbench_temporal_sequences                     |      0|multiple_choice_grade|46.80|±  |  1.58|
|bigbench_tracking_shuffled_objects_five_objects |      0|multiple_choice_grade|22.08|±  |  1.17|
|bigbench_tracking_shuffled_objects_seven_objects|      0|multiple_choice_grade|19.03|±  |  0.94|
|bigbench_tracking_shuffled_objects_three_objects|      0|multiple_choice_grade|52.00|±  |  2.89|

## 🧩 Configuration

```yaml
base_model: mlabonne/Marcoro14-7B-slerp
experts:
  - source_model: openchat/openchat-3.5-1210
    positive_prompts:
    - "chat"
    - "assistant"
    - "tell me"
    - "explain"
  - source_model: beowolx/CodeNinja-1.0-OpenChat-7B
    positive_prompts:
    - "code"
    - "python"
    - "javascript"
    - "programming"
    - "algorithm"
  - source_model: maywell/PiVoT-0.1-Starling-LM-RP
    positive_prompts:
    - "storywriting"
    - "write"
    - "scene"
    - "story"
    - "character"
  - source_model: WizardLM/WizardMath-7B-V1.1
    positive_prompts:
    - "reason"
    - "math"
    - "mathematics"
    - "solve"
    - "count"
```

## 💻 Usage

Here's a [notebook](https://colab.research.google.com/drive/1ypy8fEAJe9RkNmNQR1BduOzy2Qn6CnMl#scrollTo=myLRfwjZcIyP) to run this model in 4-bit precision using a free T4 GPU on Google Colab.

```python
!pip install -qU transformers bitsandbytes accelerate

from transformers import AutoTokenizer
import transformers
import torch

model = "mlabonne/Beyonder-4x7B-v2"

tokenizer = AutoTokenizer.from_pretrained(model)
pipeline = transformers.pipeline(
    "text-generation",
    model=model,
    model_kwargs={"torch_dtype": torch.float16, "load_in_4bit": True},
)

messages = [{"role": "user", "content": "Explain what a Mixture of Experts is in less than 100 words."}]
prompt = pipeline.tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True)
outputs = pipeline(prompt, max_new_tokens=256, do_sample=True, temperature=0.7, top_k=50, top_p=0.95)
print(outputs[0]["generated_text"])
```

Output:

> A Mixture of Experts (ME) is a machine learning technique that combines multiple expert models to make predictions or decisions. Each expert model is specialized in a different aspect of the problem, and their outputs are combined to produce a more accurate and robust solution. This approach allows the model to leverage the strengths of individual experts and compensate for their weaknesses, improving overall performance.