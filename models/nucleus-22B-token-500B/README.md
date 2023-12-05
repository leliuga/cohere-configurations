---
license: mit
language:
- en
---

# 🚀 Nucleus-22B-token-500B

**Nucleus-22B-token-500B is a 22B parameters causal decoder-only model built by Nucleus.AI and trained on 500B tokens of [RefinedWeb](https://huggingface.co/datasets/tiiuae/falcon-refinedweb) along with curated corpora. It is made available under the MIT license.**

*1T-token model coming soon* 😊.


## What about Nucleus-22B-token-500B?

* **It performs well compared to similar-size open-source models** (e.g., [MPT-7B](https://huggingface.co/mosaicml/mpt-7b), [StableLM](https://github.com/Stability-AI/StableLM), [RedPajama](https://huggingface.co/togethercomputer/RedPajama-INCITE-Base-7B-v0.1) etc.), thanks to being trained on 500B tokens of [RefinedWeb](https://huggingface.co/datasets/tiiuae/falcon-refinedweb) enhanced with curated corpora. See the [OpenLLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard).
* **It is made available under an MIT license**.
* **It is trained by a small team of four passionate for Open Source**

⚠️ **This is a raw, pretrained model, which should be further finetuned for most usecases.** 

# Model Card for Nucleus-22B-token-500B

## Model Details

### Model Description

- **Developed by:** NucleusAI;
- **Model type:** Causal decoder-only;
- **Language(s) (NLP):** English;
- **License:** MIT.

### Model Source

- **Paper:** *coming soon*.

## Uses

### Direct Use

Research on large language models; as a foundation for further specialization and finetuning for specific usecases (e.g., summarization, text generation, chatbot, etc.)

### Out-of-Scope Use

Production use without adequate assessment of risks and mitigation; any use cases which may be considered irresponsible or harmful.

## Bias, Risks, and Limitations

Nucleus-22B-token-500B is trained on English data only, and will not generalize appropriately to other languages. Furthermore, as it is trained on a large-scale corpora representative of the web, it will carry the stereotypes and biases commonly encountered online.

### Recommendations

We recommend users of Nucleus-22B-token-500B to consider finetuning it for the specific set of tasks of interest, and for guardrails and appropriate precautions to be taken for any production use.

## How to Get Started with the Mode


## Training Details

### Training Data

Nucleus-22B-token-500B was trained on 500B tokens of [RefinedWeb](https://huggingface.co/datasets/tiiuae/falcon-refinedweb), along with other corpora.

| **Data source**    | **Fraction** | **Tokens** | **Sources**                       |
|--------------------|--------------|------------|-----------------------------------|
| [RefinedWeb-English](https://huggingface.co/datasets/tiiuae/falcon-refinedweb) | 75%          | 200B     | massive web crawl                 |
| Books              | 7%           | 21B       |                                   |
| Code               | 7%           | 21B        | Big Code, CodeNet                                  |
| Technical          | 6%           | 19B        | arXiv        |
| Math          | 5%           | 17B        | Mathematica, Khan Academy        |


The data was tokenized with the tokenizer similar to Llama-[7B](https://huggingface.co/meta-llama/Llama-2-7b).

### Training Procedure

Nucleus-22B-token-500B was trained on 256 A100 80GB GPUs, using a FSDP

#### Training Hyperparameters

| **Hyperparameter** | **Value**  | **Comment**                               |
|--------------------|------------|-------------------------------------------|
| Precision          | `bfloat16` |                                           |
| Optimizer          | AdamW      |                                           |
| Learning rate      | 2e-4       | 8B tokens warm-up, cosine decay to 1.e-5 |
| Weight decay       | 1e-1       |                                           |
| Batch size         | 2048        | constant                         |
| Context length         | 2048        | constant                         |


#### Speeds, Sizes, Times

Training happened in early August 2023 and took about two weeks.