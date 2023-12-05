---
license: apache-2.0
datasets:
- Open-Orca/OpenOrca
- OpenAssistant/oasst_top1_2023-08-25
language:
- bg
- ca
- cs
- da
- de
- en
- es
- fr
- hr
- hu
- it
- nl
- pl
- pt
- ro
- ru
- sl
- sr
- sv
- uk

library_name: transformers
---

```
reference-data-model:

  datasets:
    - OpenAssistant/oasst_top1_2023-08-25:
      Lang: "bg,ca,cs,da,de,en,es,fr,hr,hu,it,nl,pl,pt,ro,ru,sl,sr,sv,uk"
      Link: https://huggingface.co/datasets/OpenAssistant/oasst_top1_2023-08-25

  model:
    - Open-Orca/Mistral-7B-OpenOrca
      Link: https://huggingface.co/Open-Orca/Mistral-7B-OpenOrca

  100 examples of generating:
    Link: https://docs.google.com/spreadsheets/d/1_4rqFnhgvjA7trwAaEidaRWczAMzuKpw/edit?usp=sharing&ouid=116592149115238887304&rtpof=true&sd=true

  Version 2:
  Link: https://huggingface.co/NickyNicky/Mistral-7B-OpenOrca-oasst_top1_2023-08-25-v2

```


## Version
```py
import torch, transformers,torchvision
torch.__version__,transformers.__version__, torchvision.__version__
#OUTPUTS: ('2.0.1+cu118', '4.34.0.dev0', '0.15.2+cu118')
```

## How to use
```py

from transformers import (
    AutoModelForCausalLM,
    AutoTokenizer,
    BitsAndBytesConfig,
    HfArgumentParser,
    TrainingArguments,
    pipeline,
    logging,
    GenerationConfig,
    TextIteratorStreamer,
)
import torch

# model_id = 'Open-Orca/Mistral-7B-OpenOrca'
model_id='NickyNicky/Mistral-7B-OpenOrca-oasst_top1_2023-08-25-v1'

model = AutoModelForCausalLM.from_pretrained(model_id,
                                             device_map="auto",
                                             trust_remote_code=True,
                                             torch_dtype=torch.bfloat16,
                                             load_in_4bit=True,
                                             low_cpu_mem_usage= True,
                                             )

max_length=2048
print("max_length",max_length)


tokenizer = AutoTokenizer.from_pretrained(model_id,
                                          # use_fast = False,
                                          max_length=max_length,)

tokenizer.pad_token = tokenizer.eos_token
tokenizer.padding_side = 'right'

#EXAMPLE #1
txt="""<|im_start|>user
I'm looking for an efficient Python script to output prime numbers. Can you help me out? I'm interested in a script that can handle large numbers and output them quickly. Also, it would be great if the script could take a range of numbers as input and output all the prime numbers within that range. Can you generate a script that fits these requirements? Thanks!<|im_end|>
<|im_start|>assistant
"""

#EXAMPLE #2
txt="""<|im_start|>user
Estoy desarrollando una REST API con Nodejs, y estoy tratando de aplicar algún sistema de seguridad, ya sea con tokens o algo similar, me puedes ayudar?<|im_end|>
<|im_start|>assistant
"""

inputs = tokenizer.encode(txt, return_tensors="pt").to("cuda")

generation_config = GenerationConfig(
              max_new_tokens=max_new_tokens,
              temperature=0.7,
              top_p=0.9,
              top_k=len_tokens,
              repetition_penalty=1.11, 
              do_sample=True,
              #  pad_token_id=tokenizer.eos_token_id,
              #  eos_token_id=tokenizer.eos_token_id,
              #  use_cache=True,
              # stopping_criteria= StoppingCriteriaList([stopping_criteria]),
          )
outputs = model.generate(generation_config=generation_config,
                                input_ids=inputs,)
tokenizer.decode(outputs[0], skip_special_tokens=False) #True
```
