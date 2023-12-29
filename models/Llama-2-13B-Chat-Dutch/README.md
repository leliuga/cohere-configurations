---
license: cc-by-nc-sa-4.0
base_model: BramVanroy/llama2-13b-ft-mc4_nl_cleaned_tiny
tags:
- generated_from_trainer
- llama
- lora
- adapters
datasets:
- BramVanroy/dutch_chat_datasets
model-index:
- name: Llama-2-13b-chat-dutch
  results: []
language:
- nl
inference: false
pipeline_tag: conversational
---


# Llama-2-13b-chat-dutch

This model is a fine-tuned version of [BramVanroy/llama2-13b-ft-mc4_nl_cleaned_tiny](https://huggingface.co/BramVanroy/llama2-13b-ft-mc4_nl_cleaned_tiny)
on the [BramVanroy/dutch_chat_datasets](https://huggingface.co/datasets/BramVanroy/dutch_chat_datasets) dataset on a context of 4096 tokens.
See the original [meta-llama/Llama-2-13b-hf](https://huggingface.co/meta-llama/Llama-2-13b-hf) for more information, intended use, and biases.

If you use this model or refer to it, please use the following citation:

Vanroy, B. (2023). *Language Resources for Dutch Large Language Modelling*. [https://arxiv.org/abs/2312.12852](https://arxiv.org/abs/2312.12852)

```bibtext
@article{vanroy2023language,
  title={Language Resources for {Dutch} Large Language Modelling},
  author={Vanroy, Bram},
  journal={arXiv preprint arXiv:2312.12852},
  year={2023}
}
```

## Usage

```python
from transformers import pipeline, Conversation


chatbot = pipeline(
    "conversational",
    model="BramVanroy/Llama-2-13b-chat-dutch",
    model_kwargs={
        "device_map": "auto",
        "load_in_8bit": True
    }
)

# Ask a first question
conversation = Conversation("Wat zijn enkele kleuren van de regenboog?")
conversation = chatbot(conversation)
# assistant: De regenboog bestaat uit zeven verschillende kleuren: rood, oranje, geel, groen, blauw, indigo en violet.

# Ask a second question
conversation.add_user_input("Interessant! Hoe worden die kleuren gevormd?")
conversation = chatbot(conversation)
print(conversation)
# Conversation id: d8abbddb-6249-4699-b789-d1b42bb4fd71
# user: Wat zijn enkele kleuren van de regenboog?
# assistant: De regenboog bestaat uit zeven verschillende kleuren: rood, oranje, geel, groen, blauw, indigo en violet.
# user: Interessant! Hoe worden die kleuren gevormd?
# assistant: De kleuren van de regenboog worden gevormd door het breken van licht door waterdruppels in de atmosfeer. Elke druppel heeft een unieke grootte en vorm, waardoor het licht dat door de druppel gaat wordt gefragmenteerd en verschillende kleuren op de bodem van de druppel wordt weerkaatst.

```

## Model description

I could not get the original Llama 2 13B to produce much Dutch, even though the description paper indicates that it was trained on a (small) portion of Dutch data. I therefore
continued training the original Llama 2 13B checkpoint on Dutch data [in regular CLM](https://huggingface.co/BramVanroy/llama2-13b-ft-mc4_nl_cleaned_tiny). In a second
step I finetuned that model on a collection of synthetic (translated) instruction and chat datasets that I have [collected](https://huggingface.co/datasets/BramVanroy/dutch_chat_datasets). 
See their pages for licensing, usage, creation, and citation information.

- https://huggingface.co/datasets/BramVanroy/dolly-15k-dutch
- https://huggingface.co/datasets/BramVanroy/alpaca-cleaned-dutch-baize
- https://huggingface.co/datasets/BramVanroy/stackoverflow-chat-dutch
- https://huggingface.co/datasets/BramVanroy/quora-chat-dutch

This model is the result of that process. While not perfect by any means, it can perform reasonably well in Dutch depending on the prompts. It is also decent at helping with programming tasks.


## Intended uses & limitations

Depending on the prompt, the model can return good results considering that it is only 13B in size and was only marginally pretrained on Dutch. That being said, the
model was not trained on human feedback and contains no safe-guards so it may produce unexpected and even offensive content depending on the query. The only attempt
of a safe-guard is the default prompt that it was trained on, which was

> Je bent een behulpzame, respectvolle en eerlijke assistent. Antwoord altijd zo behulpzaam mogelijk. Je antwoorden mogen geen schadelijke, onethische, racistische, seksistische, gevaarlijke of illegale inhoud bevatten. Zorg ervoor dat je antwoorden sociaal onbevooroordeeld en positief van aard zijn.\n\nAls een vraag nergens op slaat of feitelijk niet coherent is, leg dan uit waarom in plaats van iets niet correct te antwoorden. Als je het antwoord op een vraag niet weet, deel dan geen onjuiste informatie.

This system message is automatically applied when you use a conversational pipeline or when you use tokenizer.apply_chat_template.

Use this model with caution and at your own risk!

Because the model was trained on synthetic data, translated with OpenAI's API, you cannot use this model to create a competitive product to theirs.

## Training procedure

Trained with 4096 tokens context length. The dataset was preprocessed so that as many as possible dialogs were put in a single batch, without disrupting
dialogs. In other words, a dialog was never split up over different sequences or batches. During training, the human prompts were ignored in back propagation.

Trained with LoRA targetting ["q_proj", "v_proj"] in 4 bit and merged before upload. Trained with Flash Attention as borrowed from [here](https://github.com/philschmid/deep-learning-pytorch-huggingface/blob/main/training/utils/llama_patch.py).

The adapters are in the `adapters` branch.

### Training hyperparameters

The following hyperparameters were used during training:
- learning_rate: 0.0002
- train_batch_size: 2
- eval_batch_size: 2
- seed: 42
- distributed_type: multi-GPU
- num_devices: 4
- gradient_accumulation_steps: 8
- total_train_batch_size: 64
- total_eval_batch_size: 8
- optimizer: Adam with betas=(0.9,0.95) and epsilon=1e-08
- lr_scheduler_type: cosine
- lr_scheduler_warmup_ratio: 0.03
- num_epochs: 2

### Training results

| Training Loss | Epoch | Step | Validation Loss |
|:-------------:|:-----:|:----:|:---------------:|
| 1.0193        | 0.09  | 20   | 1.1583          |
| 0.9743        | 0.17  | 40   | 1.1339          |
| 0.9159        | 0.26  | 60   | 1.1218          |
| 0.9131        | 0.35  | 80   | 1.1153          |
| 0.8816        | 0.44  | 100  | 1.1130          |
| 0.8977        | 0.52  | 120  | 1.1069          |
| 0.9061        | 0.61  | 140  | 1.1025          |
| 0.8672        | 0.7   | 160  | 1.1024          |
| 0.8956        | 0.79  | 180  | 1.0971          |
| 0.8514        | 0.87  | 200  | 1.0995          |
| 0.8357        | 0.96  | 220  | 1.0952          |
| 0.8294        | 1.05  | 240  | 1.0964          |
| 0.8531        | 1.13  | 260  | 1.0947          |
| 0.8321        | 1.22  | 280  | 1.0951          |
| 0.8365        | 1.31  | 300  | 1.0910          |
| 0.8616        | 1.4   | 320  | 1.0894          |
| 0.8397        | 1.48  | 340  | 1.0904          |
| 0.861         | 1.57  | 360  | 1.0880          |
| 0.8116        | 1.66  | 380  | 1.0871          |
| 0.8285        | 1.74  | 400  | 1.0855          |
| 0.8603        | 1.83  | 420  | 1.0856          |
| 0.8126        | 1.92  | 440  | 1.0848          |


### Framework versions

- Transformers 4.31.0
- Pytorch 2.0.1+cu117
- Datasets 2.14.4
- Tokenizers 0.13.3

# [Open LLM Leaderboard Evaluation Results (English)](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_BramVanroy__Llama-2-13b-chat-dutch)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 46.91   |
| ARC (25-shot)         | 59.3          |
| HellaSwag (10-shot)   | 81.45    |
| MMLU (5-shot)         | 55.82         |
| TruthfulQA (0-shot)   | 38.23   |
| Winogrande (5-shot)   | 76.64   |
| GSM8K (5-shot)        | 10.69        |
| DROP (3-shot)         | 6.28         |

# Open LLM Leaderboard Evaluation Results (Dutch)

Results can be found [here](https://huggingface.co/spaces/BramVanroy/open_dutch_llm_leaderboard)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 0.43   |
| ARC (25-shot)         | 0.38          |
| HellaSwag (10-shot)   | 0.56    |
| MMLU (5-shot)         | 0.35         |
| TruthfulQA (0-shot)   | 0.44   |