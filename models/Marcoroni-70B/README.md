---
license: cc-by-nc-4.0
datasets:
- Open-Orca/OpenOrca
language:
- en
pipeline_tag: text-generation
---
# Marcoroni-70B

# Model Details
* **Trained by**: trained by AIDC AI-Business.
* **Model type:**  **Marcoroni-70B** is an auto-regressive language model based on the Llama 2 transformer architecture.
* **Language(s)**: English
* **License for Marcoroni-70B base weights**: Non-Commercial Creative Commons license ([CC BY-NC-4.0](https://creativecommons.org/licenses/by-nc/4.0/))


# Prompting

## Prompt Template for alpaca style

```
### Instruction:

<prompt> (without the <>)

### Response:
```

## Example

User: 

Give me a brief introduction to Hangzhou and Aliexpress.

Response:

Hangzhou is the capital of Zhejiang Province in China, known for its picturesque West Lake, historic monuments, and thriving technology industries. It is a popular tourist destination due to its natural beauty, historical significance, and strong presence of various technology companies.

Alibaba Group, a multinational technology conglomerate founded in 1999 by Jack Ma, has its headquarters in Hangzhou. One of Alibaba's major businesses is AliExpress, an international online marketplace. It connects buyers from all over the world with sellers mainly from China, offering a wide variety of products at affordable prices. Launched in 2010, AliExpress facilitates small businesses to reach a global audience, and provides buyers with access to a large selection of items, including electronics, clothing, beauty products, and home goods. Its platform supports multiple languages and currencies, making it easier for customers to navigate and shop across the globe.

### Our Other Projects:
* [AIDC-ai-business/Marcoroni-7B](https://huggingface.co/AIDC-ai-business/Marcoroni-7B)
* [AIDC-ai-business/Marcoroni-13B](https://huggingface.co/AIDC-ai-business/Marcoroni-13B)

We achieved the top ranker among 70B models at Sep-14th 2023.


# Evulation Results ([Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard))

| Metric                | Value |
|-----------------------|-------|
| Avg.                  |   73.94   |
| ARC (25-shot)         |   72.95   |
| HellaSwag (10-shot)   |   87.51   |
| MMLU (5-shot)         |   70.79   |
| TruthfulQA (0-shot)   |   64.49   |