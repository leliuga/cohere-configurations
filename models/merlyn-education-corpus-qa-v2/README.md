---
license: apache-2.0
tags:
- MerlynMind
- education
inference: false
---

# Merlyn-Education Corpus QA

merlyn-education-corpus-qa-v2 is a 13b parameter decoder-style transformer model for the education domain. It is fine-tuned from a [llama2-13b](https://huggingface.co/meta-llama/Llama-2-13b-hf) base-model.

This model was trained by [Merlyn Mind](https://www.merlyn.org/).

It is a model that provides an answer to a question based on the given context.

## Model Date

August 21, 2023

## Model License

Apache-2.0


## Usage

Loading model and tokenizer:

```python
import torch
from transformers import AutoTokenizer, AutoModelForCausalLM

model_path = "MerlynMind/merlyn-education-corpus-qa-v2"
device = torch.device("cuda:0") # change device id as necessary
model = AutoModelForCausalLM.from_pretrained(model_path)    
tokenizer = AutoTokenizer.from_pretrained(model_path, fast_tokenizer=True)
model.to(device) # move to device

```

Prompt example:

```python
info = '''Information:\tThe Solar System is about 4.6 billion years old. The Sun formed by gravity in a large molecular cloud. It is mainly hydrogen, which it converts into helium.
Information:\tThe formation and evolution of the Solar System began 4.6 billion years ago with the gravitational collapse of a small part of a giant molecular cloud.
Information:\tAstronomers are now more or less certain that the order of the planets was not always as it is today. Knowing what we know today, we can see the Solar System is strange. All other planetary system we are able to study have their largest planet close to their star. Also we have noticed other oddities in the Solar System. Mars is smaller than it ought to be, and the asteroid belt has been disturbed.
Information:\tFor thousands of years, people had no need for a name for the "Solar System". They thought the Earth stayed still at the center of everything (geocentrism). The Greek philosopher Aristarchus of Samos suggested that there was a special order in the sky. Nicolaus Copernicus was the first to develop a mathematical system that described what we now call the "Solar System". This was called a "new system of the world". In the 17th century, Galileo Galilei, Johannes Kepler and Isaac Newton began to understand physics more clearly. People began to accept the idea that the Earth is a planet that moves around the Sun, and that the planets are worlds, and that all worlds are governed by the same same physical laws. More recently, telescopes and space probes sometimes let us see details directly. All inner planets have surface features. The gas giants (as the name suggests) have surfaces whose make-up is gradually being discovered.
Information:\tThere are eight planets in the Solar System. From closest to farthest from the Sun, they are: Mercury, Venus, Earth, Mars, Jupiter, Saturn, Uranus and Neptune. The first four planets are called terrestrial planets. They are mostly made of rock and metal, and they are mostly solid. The last four planets are called gas giants. This is because they are much larger than other planets and are mostly made of gas.
'''
qs = "Question:\tHow old is the Solar System?"

prompt = tokenizer.bos_token
prompt += '''Instruction:\tYou are to try to answer the following question using only the pieces of information given.
Instruction:\tYour response should be a well formed JSON object with an 'answerable' property followed by an 'answer' property.
Instruction:\tIf you cannot answer the question given the information, the value of the 'answerable' should be 'false' and the 'answer' should be an empty string.
Instruction:\tIf you can answer the question given the information, the value of the 'answerable' should be 'true' and your answer should be the string value of the 'answer' property.
''' + info + qs + " Response:"

```

We recommend using newline character for stopping criterion, as follows:

```python
from transformers import StoppingCriteria, StoppingCriteriaList

eos_tokens = [tokenizer.eos_token,'\n']
eos_token_ids = [tokenizer.encode(token)[0] for token in eos_tokens]

class MultipleEOSTokensStoppingCriteria(StoppingCriteria):
    def __init__(self, eos_token_ids):
        self.eos_token_ids = set(eos_token_ids)
    def __call__(self, input_ids, scores) -> bool:
        if input_ids.shape[-1] <= 1:
            return False
        for eos_token_id in self.eos_token_ids:
            if eos_token_id == input_ids[0, -1].item():
                return True
        return False

# Define stopping criteria
multiple_eos_tokens_processor = MultipleEOSTokensStoppingCriteria(eos_token_ids)
stopping_criteria = StoppingCriteriaList([multiple_eos_tokens_processor])
```

Inference:

```python
inputs = tokenizer(prompt, return_tensors="pt", return_token_type_ids=False).to(device)
generate_ids = model.generate(
    **inputs,
    max_new_tokens=1024,
    temperature=0.0,
    num_beams=2,
    top_p=1,
    stopping_criteria=stopping_criteria
)
response = tokenizer.decode(generate_ids[0],
                      skip_special_tokens=True,
                      clean_up_tokenization_spaces=True)
```

Example output (after response processing):

```json
[{"answerable": "true", "answer": "4.6 billion years"}]
```

## Evaluation
This model is trained on a larger dataset compared to the [pythia-based v1 model](https://huggingface.co/MerlynMind/merlyn-education-corpus-qa), yielding better correctness and reduced hallucinations on a larger and more diverse benchmarking dataset.


