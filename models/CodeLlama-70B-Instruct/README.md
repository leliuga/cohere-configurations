---
language:
- code
pipeline_tag: text-generation
tags:
- llama-2
license: llama2
widget:
  - example_title: Hello world (Python)
    messages:
    - role: system
      content: You are a helpful and honest code assistant
    - role: user
      content: Print a hello world in Python
  - example_title: Sum of sublists (Python)
    messages:
    - role: system
      content: You are a helpful and honest code assistant expert in JavaScript. Please, provide all answers to programming questions in JavaScript
    - role: user
      content: Write a function that computes the set of sums of all contiguous sublists of a given list.
inference:
  parameters:
    max_new_tokens: 200
    stop:
    - </s>
    - <step>
---
# **Code Llama**
Code Llama is a collection of pretrained and fine-tuned generative text models ranging in scale from 7 billion to 70 billion parameters. This is the repository for the 70B instruct-tuned version in the Hugging Face Transformers format. This model is designed for general code synthesis and understanding. Links to other models can be found in the index at the bottom.

|     | Base Model                                                                    | Python                                                                                      | Instruct                                                                                        |
| --- | ----------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| 7B  | [codellama/CodeLlama-7b-hf](https://huggingface.co/codellama/CodeLlama-7b-hf) | [codellama/CodeLlama-7b-Python-hf](https://huggingface.co/codellama/CodeLlama-7b-Python-hf) | [codellama/CodeLlama-7b-Instruct-hf](https://huggingface.co/codellama/CodeLlama-7b-Instruct-hf) |
| 13B  | [codellama/CodeLlama-13b-hf](https://huggingface.co/codellama/CodeLlama-13b-hf) | [codellama/CodeLlama-13b-Python-hf](https://huggingface.co/codellama/CodeLlama-13b-Python-hf) | [codellama/CodeLlama-13b-Instruct-hf](https://huggingface.co/codellama/CodeLlama-13b-Instruct-hf) |
| 34B  | [codellama/CodeLlama-34b-hf](https://huggingface.co/codellama/CodeLlama-34b-hf) | [codellama/CodeLlama-34b-Python-hf](https://huggingface.co/codellama/CodeLlama-34b-Python-hf) | [codellama/CodeLlama-34b-Instruct-hf](https://huggingface.co/codellama/CodeLlama-34b-Instruct-hf) |
| 70B  | [codellama/CodeLlama-70b-hf](https://huggingface.co/codellama/CodeLlama-70b-hf) | [codellama/CodeLlama-70b-Python-hf](https://huggingface.co/codellama/CodeLlama-70b-Python-hf) | [codellama/CodeLlama-70b-Instruct-hf](https://huggingface.co/codellama/CodeLlama-70b-Instruct-hf) |

Model capabilities:

- [x] Code completion.
- [ ] Infilling.
- [x] Instructions / chat.
- [ ] Python specialist.

## Model Use

Install `transformers`

```bash
pip install transformers accelerate
```

**Chat use:** The 70B Instruct model uses a [different prompt template](#chat_prompt) than the smaller versions. To use it with `transformers`, we recommend you use the built-in chat template:

 ```py
from transformers import AutoTokenizer, AutoModelForCausalLM
import transformers
import torch

model_id = "codellama/CodeLlama-70b-Instruct-hf"
tokenizer = AutoTokenizer.from_pretrained(model_id)
model = AutoModelForCausalLM.from_pretrained(
    model_id,
    torch_dtype=torch.float16,
    device_map="auto",
)

 chat = [
    {"role": "system", "content": "You are a helpful and honest code assistant expert in JavaScript. Please, provide all answers to programming questions in JavaScript"},
    {"role": "user", "content": "Write a function that computes the set of sums of all contiguous sublists of a given list."},
]
inputs = tokenizer.apply_chat_template(chat, return_tensors="pt").to("cuda")

output = model.generate(input_ids=inputs, max_new_tokens=200)
output = output[0].to("cpu")
print(tokenizer.decode(output))
```

 You can also use the model for **text or code completion**. This examples uses transformers' `pipeline` interface:

 ```py
from transformers import AutoTokenizer
import transformers
import torch

model_id = "codellama/CodeLlama-70b-hf"
tokenizer = AutoTokenizer.from_pretrained(model_id)
pipeline = transformers.pipeline(
    "text-generation",
    model=model_id,
    torch_dtype=torch.float16,
    device_map="auto",
)

 sequences = pipeline(
    'def fibonacci(',
    do_sample=True,
    temperature=0.2,
    top_p=0.9,
    num_return_sequences=1,
    eos_token_id=tokenizer.eos_token_id,
    max_length=100,
)
for seq in sequences:
    print(f"Result: {seq['generated_text']}")
```

<a name="chat_prompt"></a>
## Chat prompt

CodeLlama 70B Instruct uses a different format for the chat prompt than previous Llama 2 or CodeLlama models. As mentioned above, the easiest way to use it is with the help of the tokenizer's chat template. If you need to build the string or tokens, manually, here's how to do it.

We'll do our tests with the following made-up dialog:

```py
chat = [
    {"role": "system", "content": "System prompt    "},
    {"role": "user", "content": "First user query"},
    {"role": "assistant", "content": "Model response to first query"},
    {"role": "user", "content": "Second user query"},
]
```

First, let's see what the prompt looks like if we use the chat template:

```py
tokenizer.apply_chat_template(chat, tokenize=False)
```

```
'<s>Source: system\n\n System prompt <step> Source: user\n\n First user query <step> Source: assistant\n\n Model response to first query <step> Source: user\n\n Second user query <step> Source: assistant\nDestination: user\n\n '
```

So each turn of the conversation has a `Source` (`system`, `user`, or `assistant`), and then the content appears after two newlines and a space. Turns are separated with the special token ` <step> `. After the last turn (which must necessarily come from the `user`), we invite the model to respond by using the special syntax `Source: assistant\nDestination: user\n\n `. Let's see how we can build the same string ourselves:

```py
output = "<s>"
for m in chat:
    output += f"Source: {m['role']}\n\n {m['content'].strip()}"
    output += " <step> "
output += "Source: assistant\nDestination: user\n\n "
output
```

```
'<s>Source: system\n\n System prompt <step> Source: user\n\n First user query <step> Source: assistant\n\n Model response to first query <step> Source: user\n\n Second user query <step> Source: assistant\nDestination: user\n\n '
```

To verify that we got it right, we'll compare against the [reference code in the original GitHub repo](https://github.com/facebookresearch/codellama/blob/1af62e1f43db1fa5140fa43cb828465a603a48f3/llama/generation.py#L506). We used the same dialog and tokenized it with the `dialog_prompt_tokens` function and got the following tokens:

```py
reference_tokens = [1, 7562, 29901, 1788, 13, 13, 2184, 9508, 32015, 7562, 29901, 1404, 13, 13, 3824, 1404, 2346, 32015, 7562, 29901, 20255, 13, 13, 8125, 2933, 304, 937, 2346, 32015, 7562, 29901, 1404, 13, 13, 6440, 1404, 2346, 32015, 7562, 29901, 20255, 13, 14994, 3381, 29901, 1404, 13, 13, 29871]
```

Let's see what we get with the string we built using our Python loop. Note that we don't add "special tokens" because the string already starts with `<s>`, the beginning of sentence token:

```py
tokens = tokenizer.encode(output, add_special_tokens=False)
assert reference_tokens == tokens
```

Similarly, let's verify that the chat template produces the same token sequence:

```py
assert reference_tokens == tokenizer.apply_chat_template(chat)
```

As a final detail, please note that if the dialog does not start with a `system` turn, the [original code will insert one with an empty content string](https://github.com/facebookresearch/codellama/blob/1af62e1f43db1fa5140fa43cb828465a603a48f3/llama/generation.py#L418).


## Model Details
*Note: Use of this model is governed by the Meta license. Meta developed and publicly released the Code Llama family of large language models (LLMs).

**Model Developers** Meta

**Variations** Code Llama comes in four model sizes, and three variants:

* Code Llama: base models designed for general code synthesis and understanding
* Code Llama - Python: designed specifically for Python
* Code Llama - Instruct: for instruction following and safer deployment

All variants are available in sizes of 7B, 13B, 34B, and 70B parameters.

**This repository contains the Instruct version of the 70B parameters model.**

**Input** Models input text only.

**Output** Models generate text only.

**Model Architecture** Code Llama is an auto-regressive language model that uses an optimized transformer architecture. It was fine-tuned with up to 16k tokens. This variant **does not** support long context of up to 100k tokens.

**Model Dates** Code Llama and its variants have been trained between January 2023 and January 2024.

**Status** This is a static model trained on an offline dataset. Future versions of Code Llama - Instruct will be released as we improve model safety with community feedback.

**License** A custom commercial license is available at: [https://ai.meta.com/resources/models-and-libraries/llama-downloads/](https://ai.meta.com/resources/models-and-libraries/llama-downloads/)

**Research Paper** More information can be found in the paper "[Code Llama: Open Foundation Models for Code](https://ai.meta.com/research/publications/code-llama-open-foundation-models-for-code/)" or its [arXiv page](https://arxiv.org/abs/2308.12950).

## Intended Use
**Intended Use Cases** Code Llama and its variants are intended for commercial and research use in English and relevant programming languages. The base model Code Llama can be adapted for a variety of code synthesis and understanding tasks, Code Llama - Python is designed specifically to handle the Python programming language, and Code Llama - Instruct is intended to be safer to use for code assistant and generation applications.

**Out-of-Scope Uses** Use in any manner that violates applicable laws or regulations (including trade compliance laws). Use in languages other than English. Use in any other way that is prohibited by the Acceptable Use Policy and Licensing Agreement for Code Llama and its variants.

## Hardware and Software
**Training Factors** We used custom training libraries. The training and fine-tuning of the released models have been performed Meta’s Research Super Cluster.
**Carbon Footprint** In aggregate, training all 12 Code Llama models required 1400K GPU hours of computation on hardware of type A100-80GB (TDP of 350-400W). Estimated total emissions were 228.55 tCO2eq, 100% of which were offset by Meta’s sustainability program.

## Evaluation Results

See evaluations for the main models and detailed ablations in Section 3 and safety evaluations in Section 4 of the research paper.


## Ethical Considerations and Limitations 

Code Llama and its variants are a new technology that carries risks with use. Testing conducted to date has been in English, and has not covered, nor could it cover all scenarios. For these reasons, as with all LLMs, Code Llama’s potential outputs cannot be predicted in advance, and the model may in some instances produce inaccurate or objectionable responses to user prompts. Therefore, before deploying any applications of Code Llama, developers should perform safety testing and tuning tailored to their specific applications of the model.

Please see the Responsible Use Guide available available at [https://ai.meta.com/llama/responsible-use-guide](https://ai.meta.com/llama/responsible-use-guide).