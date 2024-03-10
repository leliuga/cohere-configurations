---
library_name: transformers
tags:
- llama2
- deutsch
- german
- seedbox
license: llama2
datasets:
- seedboxai/multitask_german_examples_32k
language:
- de
pipeline_tag: text-generation
---

![image/jpeg](https://cdn-uploads.huggingface.co/production/uploads/645ded34a45b4182d7f5c385/hJ7zsOGDgLWUmf7vbaoI_.jpeg)


# KafkaLM-70B-German-V0.1

**KafkaLM 70b** is a 70b model based on [Llama2 70B Base Model](https://huggingface.co/meta-llama/Llama-2-70b-hf) which was finetuned on an ensemble of popular high-quality open-source instruction sets (translated from English to German). 

KafkaLM 70b is a [Seedbox](https://huggingface.co/seedboxai) project trained by [Dennis Dickmann](https://huggingface.co/doubledsbv).

**Why Kafka?** 
The models are proficient, yet creative, have some tendencies to linguistically push boundaries 😊


## Model Details

The purpose of releasing the **KafkaLM series** is to contribute to the German AI community with a set of fine-tuned LLMs that are easy to use in everyday applications across a variety of tasks.

The main goal was to provide LLMs proficient in German, especially to be used in German-speaking business contexts where English alone is not sufficient.


### Dataset

I used a 4k filtered version of the following [seedboxai/multitask_german_examples_32k](https://huggingface.co/datasets/seedboxai/multitask_german_examples_32k)

### Prompt Format


This model follows the subsequent prompt format:

```
<|system|>
Du bist ein freundlicher und hilfsbereiter KI-Assistent. Du beantwortest Fragen faktenorientiert und präzise, ohne dabei relevante Fakten auszulassen.</s>
<|user|>
Welche Möglichkeiten der energetischen Sanierung habe ich neben Solar und Energiespeicher?</s>
<|assistant|>
```

### Inference

Getting started with the model is straight forward

```python
import transformers

model_id = "seedboxai/KafkaLM-70B-German-V0.1"

model = AutoModelForCausalLM.from_pretrained(model_id, load_in_4bit=True)

tokenizer = AutoTokenizer.from_pretrained(model_id)

tokenizer.padding_side = "right" 
tokenizer.pad_token = tokenizer.unk_token 
tokenizer.add_eos_token = False

def generate_prompt(input):
    prompt = ''
    sys_prompt = "Du bist ein freundlicher und hilfsbereiter KI-Assistent. Du beantwortest Fragen faktenorientiert und präzise, ohne dabei relevante Fakten auszulassen."
    
    prompt += f"<|system|>\n{sys_prompt.strip()}</s>\n"
    prompt += f"<|user|>\n{input.strip()}</s>\n"
    prompt += f"<|assistant|>\n"

    return prompt.strip()


generate_text = transformers.pipeline(
    model=model, tokenizer=tokenizer,
    return_full_text=True,  
    task='text-generation',
    temperature=0.5,  
    max_new_tokens=512,
    top_p=0.95,
    top_k=50,
    do_sample=True,
)

print(generate_text(generate_prompt("Wer ist eigentlich dieser Kafka?"))

```

## Disclaimer

The license on this model does not constitute legal advice. We are not responsible for the actions of third parties who use this model.
This model should only be used for research purposes. The original Llama2 license and all restrictions of datasets used to train this model apply.