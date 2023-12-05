---
license: llama2
model-index:
- name: Phind-CodeLlama-34B-v1
  results:
  - task:
      type: text-generation
    dataset:
      type: openai_humaneval
      name: HumanEval
    metrics:
    - name: pass@1
      type: pass@1
      value: 69.5%
      verified: false
tags:
- code llama
---

# **Phind-CodeLlama-34B-Python-v1**
We've fine-tuned CodeLlama-34B and CodeLlama-34B-Python on an internal Phind dataset that achieve 67.6% and 69.5% pass@1 on HumanEval, respectively. GPT-4 achieves 67%. We've applied OpenAI's decontamination methodology to our dataset to ensure result validity.

More details can be found on our [blog post](https://www.phind.com/blog/code-llama-beats-gpt4).

## Model Details
This model is fine-tuned from CodeLlama-34B-Python and achieves 69.5% pass@1 on HumanEval.

## Dataset Details
We fined-tuned on a proprietary dataset of ~80k high quality programming problems and solutions. This dataset consists of instruction-answer pairs instead of code completion examples, making it structurally different from HumanEval. The Phind models were trained for 2 epochs, for a total of ~160k examples shown. LoRA was not used -- both models are a native finetune. We used DeepSpeed ZeRO 3 and Flash Attention 2 to train these models in three hours on 32 A100-80GB GPUs. We used a sequence length of 4096 tokens.

## How to Get Started with the Model

Make sure to install Transformers from the main git branch:

```bash
pip install git+https://github.com/huggingface/transformers.git
```

## How to Prompt the Model
**Please note that this model is somewhat instruction-tuned, but not chat-tuned.**

Do not try to use the Llama chat markup with this model. Instead, simply tell it what you want and add "\n: " at the end of your task.

For example: 

```
Write me a linked list implementation: \n
```

## How to reproduce HumanEval Results

To reproduce our results:

```python

from transformers import AutoTokenizer, LlamaForCausalLM
from human_eval.data import write_jsonl, read_problems
from tqdm import tqdm

# initialize the model

model_path = "Phind/Phind-CodeLlama-34B-v1"
model = LlamaForCausalLM.from_pretrained(model_path, device_map="auto")
tokenizer = AutoTokenizer.from_pretrained(model_path)

# HumanEval helper

def generate_one_completion(prompt: str):
    tokenizer.pad_token = tokenizer.eos_token
    inputs = tokenizer(prompt, return_tensors="pt", truncation=True, max_length=4096)

    # Generate
    generate_ids = model.generate(inputs.input_ids.to("cuda"), max_new_tokens=256, do_sample=True, top_p=0.75, top_k=40, temperature=0.1)
    completion = tokenizer.batch_decode(generate_ids, skip_special_tokens=True, clean_up_tokenization_spaces=False)[0]
    completion = completion.replace(prompt, "").split("\n\n\n")[0]

    return completion

# perform HumanEval
problems = read_problems()

num_samples_per_task = 1
samples = [
    dict(task_id=task_id, completion=generate_one_completion(problems[task_id]["prompt"]))
    for task_id in tqdm(problems)
    for _ in range(num_samples_per_task)
]
write_jsonl("samples.jsonl", samples)

# run `evaluate_functional_correctness samples.jsonl` in your HumanEval code sandbox
```

## Bias, Risks, and Limitations

<!-- This section is meant to convey both technical and sociotechnical limitations. -->
This model has undergone very limited testing. Additional safety testing should be performed before any real-world deployments.


## Training details

<!-- Total emissions (in grams of CO2eq) and additional considerations, such as electricity usage, go here. Edit the suggested text below accordingly -->

- **Hardware Type:** 32x A100-80GB
- **Hours used:** 90 GPU-hours
- **Cloud Provider:** AWS
- **Compute Region:** us-east-1