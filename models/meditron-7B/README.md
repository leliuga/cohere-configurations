---
license: llama2
language:
- en
metrics:
- accuracy
- perplexity
datasets:
- epfl-llm/guidelines
base_model: meta-llama/Llama-2-7b
---
<img width=50% src="meditron_LOGO.png" alt="Alt text" title="Meditron-logo">

# Model Card for Meditron-7B-v1.0
Meditron is a suite of open-source medical Large Language Models (LLMs).
Meditron-7B is a 7 billion parameters model adapted to the medical domain from Llama-2-7B through continued pretraining on a comprehensively curated medical corpus, including selected PubMed articles, abstracts, a [new dataset](https://huggingface.co/datasets/epfl-llm/guidelines) of internationally-recognized medical guidelines, and general domain data from [RedPajama-v1](https://huggingface.co/datasets/togethercomputer/RedPajama-Data-1T).
Meditron-7B, finetuned on relevant training data, outperforms Llama-2-7B and PMC-Llama on multiple medical reasoning tasks.

<details open>
  <summary><strong>Advisory Notice</strong></summary>

  <blockquote style="padding: 10px; margin: 0 0 10px; border-left: 5px solid #ddd;">
    While Meditron is designed to encode medical knowledge from sources of high-quality evidence, it is not yet adapted to deliver this knowledge appropriately, safely, or within professional actionable constraints. 
  We recommend against deploying Meditron in medical applications without extensive use-case alignment, as well as additional testing, specifically including randomized controlled trials in real-world practice settings.
  </blockquote>
</details>

## Model Details

- **Developed by:** [EPFL LLM Team](https://huggingface.co/epfl-llm)
- **Model type:** Causal decoder-only transformer language model
- **Language(s):** English (mainly)
- **Model License:** [LLAMA 2 COMMUNITY LICENSE AGREEMENT](https://huggingface.co/meta-llama/Llama-2-70b/raw/main/LICENSE.txt)
- **Code License:** [APACHE 2.0 LICENSE](LICENSE)
- **Continue-pretrained from model:** [Llama-2-7B](https://huggingface.co/meta-llama/Llama-2-7b)
- **Context length:**  2K tokens
- **Input:**  Text-only data
- **Output:**  Model generates text only
- **Status:** This is a static model trained on an offline dataset. Future versions of the tuned models will be released as we enhance model's performance.
- **Knowledge Cutoff:** August 2023


### Model Sources

- **Repository:** [epflLLM/meditron](https://github.com/epfLLM/meditron)
- **Trainer:** [epflLLM/Megatron-LLM](https://github.com/epfLLM/Megatron-LLM)
- **Paper:** *[MediTron-70B: Scaling Medical Pretraining for Large Language Models](https://arxiv.org/abs/2311.16079)*

## Uses

Meditron-7B is being made available for further testing and assessment as an AI assistant to enhance clinical decision-making and enhance access to an LLM for healthcare use. Potential use cases may include but are not limited to:
-  Medical exam question answering
-  Supporting differential diagnosis
-  Disease information (symptoms, cause, treatment) query
-  General health information query

### Direct Use

It is possible to use this model to generate text, which is useful for experimentation and understanding its capabilities. 
It should not be used directly for production or work that may impact people.

### Downstream Use
Meditron-70B and Meditron-7B are both foundation models without finetuning or instruction-tuning. They can be finetuned, instruction-tuned, or RLHF-tuned for specific downstream tasks and applications.
There are two ways we have used this model for downstream question-answering tasks.
1. We apply in-context learning with k demonstrations (3 or 5 in our paper) added to the prompt.
2. We finetuned the models for downstream question-answering tasks using specific training sets.

We encourage and look forward to the adaption of the base model for more diverse applications.

If you want a more interactive way to prompt the model, we recommend using a high-throughput and memory-efficient inference engine with a UI that supports chat and text generation.

You can check out our deployment [guide](https://github.com/epfLLM/meditron/blob/main/deployment/README.md), where we used [FastChat](https://github.com/lm-sys/FastChat) with [vLLM](https://github.com/vllm-project/vllm). We collected generations for our qualitative analysis through an interactive UI platform, [BetterChatGPT](https://github.com/ztjhz/BetterChatGPT). Here is the prompt format we used as an example:

<img width=70% src="prompt_example.png" alt="qualitative-analysis-prompt" title="Qualitative Analysis Prompt">

### Out-of-Scope Use

We do not recommend using this model for natural language generation in a production environment, finetuned or otherwise.

## Truthfulness, Helpfulness, Risk, and Bias

<!-- This section is meant to convey both technical and sociotechnical limitations. -->

We did an initial assessment of Meditron models' **Truthfulness** against baseline models and consumer-level medical models. 
We use TruthfulQA (multiple choice) as the main evaluation benchmark.
We only focus on the categories that are relevant to the medical domain, including Health, Nutrition, Psychology, and Science.
For 7B models, we perform one-shot evaluations for consistent answer generation.
For 70B models, the evaluations are under the zero-shot setting.
Below, we report the detailed truthfulness performance of each category.

|     |        |      |      |      |      |      |      |
| --- | ------ |----- |----- |----- |----- |----- |----- |
|Category  | meditron-70b | llama-2-70b | med42-70b* | meditron-7b | llama-2-7b | PMC-llama-7b |
|Health    |      81.8    |    69.1     |    83.6    | 27.3        |   16.4     |      3.6     |
|Nutrition |      77.9    |    68.8     |    62.5    | 31.1        |   12.5     |      6.3     |
|Psychology|      47.4    |    36.8     |    52.6    | 21.1        |   10.5     |      0.0     |
|Science   |      77.8    |    44.4     |    33.3    | 33.3        |   11.1     |      0.0     |
|Avg       |      71.2    |    54.8     |    58.0    | 28.3        |   12.6     |      2.5     |
|          |              |             |            |             |            |              |

For a more detailed performance analysis, please see our paper. 

Significant research is still required to fully explore potential bias, fairness, and safety issues with this language model. 
Please recognize that our evaluation on Meditron-7B's helpfulness, risk, and bias are highly limited. 
Thus, as we noted in the safety notice, we strongly against any deployment in medical applications without further alignment process and rigorous evaluation!

### Recommendations

**IMPORTANT!**
Users (both direct and downstream) should be made aware of the risks, biases, and limitations of the model.
While this model is capable of generating natural language text, we have only begun to explore this capability and its limitations. 
Understanding these limitations is especially important in a domain like medicine. 
Therefore, we strongly recommend against using this model in production for natural language generation or for professional purposes related to health and medicine.

## Training Details

### Training Data
Meditron’s domain-adaptive pre-training corpus GAP-Replay  combines 48.1B tokens from four corpora:
- [**Clinical  Guidelines**](https://huggingface.co/datasets/epfl-llm/guidelines): a new dataset of 46K internationally-recognized clinical practice guidelines from various healthcare-related sources, including hospitals and international organizations.
- **Medical Paper Abstracts**: 16.1M abstracts extracted from closed-access PubMed and PubMed Central papers.
- **Medical Papers**: full-text articles extracted from 5M publicly available PubMed and PubMed Central papers.
- **Replay Data**: 400M tokens of general domain pretraining data sampled from [RedPajama-v1](https://huggingface.co/datasets/togethercomputer/RedPajama-Data-1T)

<img width=75% src="gap-replay.png" alt="Alt text" title="Meditron-logo">

#### Data Preprocessing

Please see the detailed preprocessing procedure in our paper. 

### Training Procedure 

We used the [Megatron-LLM](https://github.com/epfLLM/Megatron-LLM) distributed training library, a derivative of Nvidia's Megatron LM project, to optimize training efficiency. 
Hardware consists of 1 node of 8x NVIDIA A100 (80GB) SXM GPUs connected by NVLink and NVSwitch with a single Nvidia ConnectX-6 DX network card and equipped with 2 x AMD EPYC 7543 32-Core Processors and 512 GB of RAM. 

Our three way parallelism scheme uses:
 - Data Parallelism (DP -- different GPUs process different subsets of the batches) of 2,
 - Pipeline Parallelism (PP -- different GPUs process different layers) of 4,
 - Tensor Parallelism (TP -- different GPUs process different subtensors for matrix multiplication) of 1.
   

#### Training Hyperparameters

|  |  |
| --- | ------ |
| bf16 | true |
| lr  | 3e-4 |
| eps | 1e-5       |
| betas | \[0.9, 0.95\] |
| clip_grad | 1 |
| weight decay | 0.1 |
| DP size | 16 |
| TP size | 4 |
| PP size | 1 |
| seq length | 2048 |
| lr scheduler | cosine|
| min lr | 1e-6 |
| warmup iteration | 2000 |
| micro batch size | 10 |
| global batch size | 1600 |
|  |  |

#### Sizes
The model was trained in September 2023.

The model architecture is exactly Llama 2, meaning

|     |        |
| --- | ------ |
| Model size           | 7B   |
| Hidden dimension     | 4096 |
| Num. attention heads |   32 |
| Num. layers          |   32 |
|  |  |

## Evaluation

<!-- This section describes the evaluation protocols and provides the results. -->

### Testing Data & Metrics

#### Testing Data
- [MedQA (USMLE)](https://huggingface.co/datasets/bigbio/med_qa)
- [MedMCQA](https://huggingface.co/datasets/medmcqa)
- [PubMedQA](https://huggingface.co/datasets/bigbio/pubmed_qa)
- [MMLU-Medical](https://huggingface.co/datasets/lukaemon/mmlu)
- [MedQA-4-Option](https://huggingface.co/datasets/GBaker/MedQA-USMLE-4-options)

#### Metrics
- Accuracy: suite the evaluation of multiple-choice question-answering tasks.

### Results
We finetune meditron-7b, llama-2-7b, pmc-llama-7b on each benchmark (pubmedqa, medmcqa, medqa)'s training data individually. 
We report the finetuned models' performance with top token selection as the inference mode.
For MMLU-Medical, models finetuned on MedMCQA are used for inference.
For MedQA-4-Option, models finetuned on MedQA are used for inference.
For a more detailed performance analysis, please see our paper. 

|     |        |      |      |      |      |
| --- | ------ |----- |----- |----- |----- |
|Dataset       | meditron-7b | llama-2-7b  | pmc-llama-7b | Zephyr-7B-beta* | Mistral-7B-instruct* |
|MMLU-Medical  | 54.2        |    53.7     |    56.4      |      63.3       |        60.0          |
|PubMedQA      | 74.4        |    61.8     |    59.2      |      46.0       |        17.8          |
|MedMCQA       | 59.2        |    54.4     |    57.6      |      43.0       |        40.2          |
|MedQA         | 47.9        |    44.0     |    42.4      |      42.8       |        32.4          |
|MedQA-4-Option| 52.0        |    49.6     |    49.2      |      48.5       |        41.1          |
|Avg           | 57.5        |    52.7     |    53.0      |      48.7       |        38.3          |
|              |             |             |              |                 |                      |

**Note**: models with * are already instruction-tuned, so we exclude them from further finetuning on any training data.


## Environmental Impact

<!-- Total emissions (in grams of CO2eq) and additional considerations, such as electricity usage, go here. Edit the suggested text below accordingly -->

- **Hardware Type:** 8 x NVIDIA A100 (80GB) SXM
- **Total GPU hours:** 588.8
- **Hardware Provider:** EPFL Research Computing Platform
- **Compute Region:** Switzerland
- **Carbon Emitted:** Switzerland has a carbon efficiency of 0.016 kgCO2/kWh (https://www.carbonfootprint.com/docs/2018_8_electricity_factors_august_2018_-_online_sources.pdf). 73.6 hours of 8 A100s means 588.8 hours at a TDP of 400W. Assuming a Power Usage effectiveness of 1.5, total emissions are estimated to be: 
    
   (400W / 1000W/kWh / GPU * 0.016 kgCO2/kWh * 73.6 h * 8 GPU) * 1.8 PUE = 6.8 kgCO2.

## Citation

**BibTeX:**
If you use Meditron or its training data, please cite our work:

```
@misc{chen2023meditron70b,
      title={MEDITRON-70B: Scaling Medical Pretraining for Large Language Models}, 
      author={Zeming Chen and Alejandro Hernández-Cano and Angelika Romanou and Antoine Bonnet and Kyle Matoba and Francesco Salvi and Matteo Pagliardini and Simin Fan and Andreas Köpf and Amirkeivan Mohtashami and Alexandre Sallinen and Alireza Sakhaeirad and Vinitra Swamy and Igor Krawczuk and Deniz Bayazit and Axel Marmet and Syrielle Montariol and Mary-Anne Hartley and Martin Jaggi and Antoine Bosselut},
      year={2023},
      eprint={2311.16079},
      archivePrefix={arXiv},
      primaryClass={cs.CL}
}

@software{epfmedtrn,
  author = {Zeming Chen and Alejandro Hernández-Cano and Angelika Romanou and Antoine Bonnet and Kyle Matoba and Francesco Salvi and Matteo Pagliardini and Simin Fan and Andreas Köpf and Amirkeivan Mohtashami and Alexandre Sallinen and Alireza Sakhaeirad and Vinitra Swamy and Igor Krawczuk and Deniz Bayazit and Axel Marmet and Syrielle Montariol and Mary-Anne Hartley and Martin Jaggi and Antoine Bosselut},
  title = {MediTron-70B: Scaling Medical Pretraining for Large Language Models},
  month = November,
  year = 2023,
  url = {https://github.com/epfLLM/meditron}
}
```