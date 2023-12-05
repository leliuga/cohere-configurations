<p><h1> 🦾 Mistralic-7B-1 🦾 </h1></p>

Special thanks to Together Compute for sponsoring Skunkworks with compute!

**INFERENCE**

```
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer

torch.set_default_device('cuda')
system_prompt = "Below is an instruction that describes a task, paired with an input that provides further context. Write a response that appropriately completes the request.\n\n"
system_no_input_prompt = "Below is an instruction that describes a task. Write a response that appropriately completes the request.\n\n"

def generate_prompt(instruction, input=None):
    if input:
      prompt = f"### System:\n{system_prompt}\n\n"
    else:
      prompt = f"### System:\n{system_no_input_prompt}\n\n"


    prompt += f"### Instruction:\n{instruction}\n\n"
    
    
    if input:
        prompt += f"### Input:\n{input}\n\n"
    return prompt + """### Response:\n"""


device = "cuda" 

model = AutoModelForCausalLM.from_pretrained("SkunkworksAI/Mistralic-7B-1")
tokenizer = AutoTokenizer.from_pretrained("SkunkworksAI/Mistralic-7B-1")


while True:
    instruction = input("Enter Instruction: ")
    instruction = generate_prompt(instruction)
    inputs = tokenizer(instruction, return_tensors="pt", return_attention_mask=False)

    outputs = model.generate(**inputs, max_length=1000, do_sample=True, temperature=0.01, use_cache=True, eos_token_id=tokenizer.eos_token_id)
    text = tokenizer.batch_decode(outputs)[0]
    print(text)
```

**EVALUATION**

![image/png](https://cdn-uploads.huggingface.co/production/uploads/64b7e345f92b20f7a38bf47a/ycpNhdGZHGbai_wslT2Bg.png)

Average: 0.72157 

For comparison:

  mistralai/Mistral-7B-v0.1 scores 0.7116
  
  mistralai/Mistral-7B-Instruct-v0.1 scores 0.6794

