---
license: llama2
library_name: transformers
tags:
- code
model-index:
- name: Pandalyst-7B-V1.2
  results:
  - task:
      type: text-generation
    metrics:
    - name: acc@1
      type: acc@1
      value: 0.0
      verified: false
language:
- en
---


## Pandalyst: A large language model for mastering data analysis using pandas

<p align="center">
<img src="https://raw.githubusercontent.com/pipizhaoa/Pandalyst/master/imgs/pandalyst.png" width="300"/>
</p>

<p align="center">
🐱 <a href="https://github.com/pipizhaoa/Pandalyst" target="_blank">Github Repo</a> <br>
</p>

**What is Pandalyst**
- Pandalyst is a general large language model specifically trained to process and analyze data using the pandas library.

**How is Pandalyst**
- Pandalyst has strong generalization capabilities for data tables in different fields and different data analysis needs.

**Why is Pandalyst**
- Pandalyst is open source and free to use, and its small parameter size (7B/13B) allows us to easily deploy it on local PC. 
- Pandalyst can handle complex data tables (multiple columns and multiple rows), allowing us to enter enough context to describe our table in detail.
- Pandalyst has very competitive performance, significantly outperforming models of the same size and even outperforming some of the strongest closed-source models.


## News
- 🔥[2023/10/15] Now we can **plot** 📈! and much more powerful! We released **Pandalyst-7B-V1.2**, which was trained on **CodeLlama-7b-Python** and it surpasses **ChatGPT-3.5 (2023/06/13)**, **Pandalyst-7B-V1.1** and **WizardCoder-Python-13B-V1.0** in our **PandaTest_V1.0**.
- 🤖️[2023/09/30] We released **Pandalyst-7B-V1.1** , which was trained on **CodeLlama-7b-Python** and achieves the **76.1 exec@1** in our **PandaTest_V1.0** and surpasses **WizardCoder-Python-13B-V1.0** and **ChatGPT-3.5 (2023/06/13)**.

| Model               | Checkpoint                                                                                 | Support plot | License |
|---------------------|--------------------------------------------------------------------------------------------|--------------|  ----- | 
| 🔥Pandalyst-7B-V1.2 | 🤗 <a href="https://huggingface.co/pipizhao/Pandalyst-7B-V1.2" target="_blank">HF Link</a> | ✅            |  <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama2</a>  |
| Pandalyst-7B-V1.1   | 🤗 <a href="https://huggingface.co/pipizhao/Pandalyst-7B-V1.1" target="_blank">HF Link</a> | ❌            |  <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama2</a>  |


## Usage and Human evaluation
Please refer to <a href="https://github.com/pipizhaoa/Pandalyst" target="_blank">Github</a>.