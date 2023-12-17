---
language:
- pt
---

Sabiá-7B is Portuguese language model developed by [Maritaca AI](https://www.maritaca.ai/).

**Input:** The model accepts only text input.

**Output:** The Model generates text only.

**Model Architecture:** Sabiá-7B is an auto-regressive language model that uses the same architecture of LLaMA-1-7B.

**Tokenizer:** It uses the same tokenizer as LLaMA-1-7B.

**Maximum sequence length:** 2048 tokens.

**Pretraining data:** The model was pretrained on 7 billion tokens from the Portuguese subset of ClueWeb22, starting with the weights of LLaMA-1-7B and further trained for an additional 10 billion tokens, approximately 1.4 epochs of the training dataset.

**Data Freshness:** The pretraining data has a cutoff of mid-2022.

**License:** The licensing is the same as LLaMA-1's, restricting the model's use to research purposes only.

**Paper:** For more details, please refer to our paper: [Sabiá: Portuguese Large Language Models](https://arxiv.org/pdf/2304.07880.pdf) 

Given that Sabiá-7B was trained solely on a language modeling objective without fine-tuning for instruction following, it is recommended for few-shot tasks rather than zero-shot tasks.

**Results in Portuguese** 

Below we show the results on the Poeta benchmark, which consists of 14 Portuguese datasets.

For more information on the Normalized Preferred Metric (NPM), please refer to our paper.

|Model | NPM |
|--|--|
|LLaMA-1-7B| 33.0|
|LLaMA-2-7B| 43.7|
|Sabiá-7B| 48.5|

**Results in English** 

Below we show the average results on 6 English datasets: PIQA, HellaSwag, WinoGrande, ARC-e, ARC-c, and OpenBookQA.

|Model | NPM |
|--|--|
|LLaMA-1-7B| 50.1|
|Sabiá-7B| 49.0|



Please use the following bibtex to cite our paper: 
```
@InProceedings{10.1007/978-3-031-45392-2_15,
    author="Pires, Ramon
    and Abonizio, Hugo
    and Almeida, Thales Sales
    and Nogueira, Rodrigo",
    editor="Naldi, Murilo C.
    and Bianchi, Reinaldo A. C.",
    title="Sabi{\'a}: Portuguese Large Language Models",
    booktitle="Intelligent Systems",
    year="2023",
    publisher="Springer Nature Switzerland",
    address="Cham",
    pages="226--240",
    isbn="978-3-031-45392-2"
}
```
