---
license: cc-by-nc-4.0
datasets:
- starmpcc/Asclepius-Synthetic-Clinical-Notes
language:
- en
pipeline_tag: text2text-generation
tags:
- medical
---
# Model Card for Model ID

<!-- Provide a quick summary of what the model is/does. -->

This is official model checkpoint for Asclepius-13B [(arxiv)](https://arxiv.org/abs/2309.00237).
This model is the first publicly shareable clinical LLM, trained with synthetic data.

## Model Details

### Model Description

<!-- Provide a longer summary of what this model is. -->



- **Model type:** Clinical LLM (Large Language Model)
- **Language(s) (NLP):** English
- **License:** CC-BY-NC-SA 4.0
- **Finetuned from model [optional]:** LLaMA-13B

### Model Sources [optional]

<!-- Provide the basic links for the model. -->

- **Repository:** https://github.com/starmpcc/Asclepius
- **Paper [optional]:** https://arxiv.org/abs/2309.00237
- **Data:** https://huggingface.co/datasets/starmpcc/Asclepius-Synthetic-Clinical-Notes

## Uses

<!-- Address questions around how the model is intended to be used, including the foreseeable users of the model and those affected by the model. -->
This model can perform below 8 clinical NLP tasks, with clincal notes.
- Named Entity Recognition
- Abbreviation Expansion
- Relation Extraction
- Temporal Information Extraction
- Coreference Resolution
- Paraphrasing
- Summarization
- Question Answering

### Direct Use

<!-- This section is for the model use without fine-tuning or plugging into a larger ecosystem/app. -->

[More Information Needed]

### Downstream Use [optional]

<!-- This section is for the model use when fine-tuned for a task, or when plugged into a larger ecosystem/app -->

[More Information Needed]

### Out-of-Scope Use

<!-- This section addresses misuse, malicious use, and uses that the model will not work well for. -->

ONLY USE THIS MODEL FOR RESEARCH PURPOSE!!

## How to Get Started with the Model

```python
prompt = """You are an intelligent clinical languge model.
Below is a snippet of patient's discharge summary and a following instruction from healthcare professional.
Write a response that appropriately completes the instruction.
The response should provide the accurate answer to the instruction, while being concise.

[Discharge Summary Begin]
{note}
[Discharge Summary End]

[Instruction Begin]
{question}
[Instruction End] 
"""

from transformers import AutoTokenizer, AutoModel
tokenizer = AutoTokenizer.from_pretrained("starmpcc/Asclepius-13B")
model = AutoModel.from_pretrained("starmpcc/Asclepius-13B")

note = "This is a sample note"
question = "What is the diagnosis?"

model_input = prompt.format(note=note, question=question)
input_ids = tokenizer(model_input, return_tensors="pt").input_ids
output = model.generate(input_ids)
print(tokenizer.decode(output[0]))
```

## Training Details

### Training Data

<!-- This should link to a Data Card, perhaps with a short stub of information on what the training data is all about as well as documentation related to data pre-processing or additional filtering. -->

https://huggingface.co/datasets/starmpcc/Asclepius-Synthetic-Clinical-Notes

### Training Procedure 

<!-- This relates heavily to the Technical Specifications. Content here should link to that section when it is relevant to the training procedure. -->
- Initial training was conducted using causal language modeling on synthetic clinical notes.
- It was then fine-tuned with clinical instruction-response pairs.
- For a comprehensive overview of our methods, our upcoming paper will serve as a resource.

#### Training Hyperparameters

- We followed config used in [Stanford Alpaca](https://github.com/tatsu-lab/stanford_alpaca)
- 
#### Speeds, Sizes, Times

<!-- This section provides information about throughput, start/end time, checkpoint size if relevant, etc. -->
- Pre-Training (1 epoch): 1h 52m with 8x A100 80G
- Instruction Fine-Tuning (3 epoch): 12h 16m with 8x A100 80G



## Citation

<!-- If there is a paper or blog post introducing the model, the APA and Bibtex information for that should go in this section. -->

**BibTeX:**
```
@misc{kweon2023publicly,
    title={Publicly Shareable Clinical Large Language Model Built on Synthetic Clinical Notes},
    author={Sunjun Kweon and Junu Kim and Jiyoun Kim and Sujeong Im and Eunbyeol Cho and Seongsu Bae and Jungwoo Oh and Gyubok Lee and Jong Hak Moon and Seng Chan You and Seungjin Baek and Chang Hoon Han and Yoon Bin Jung and Yohan Jo and Edward Choi},
    year={2023},
    eprint={2309.00237},
    archivePrefix={arXiv},
    primaryClass={cs.CL}
}
```



