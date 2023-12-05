---
license: apache-2.0
pipeline_tag: text-generation
language:
- en
library_name: transformers
---

## SynthIA 7B v2.0 extended to 16k context

This is the same original model so the original model license still applies.

This model has been modified to have a larger maximum context size of 16k.

### Original Model Card:

### Prompt format:

```
SYSTEM: Elaborate on the topic using a Tree of Thoughts and backtrack when necessary to construct a clear, cohesive Chain of Thought reasoning. Always answer without hesitation.
USER: How is insulin synthesized?
ASSISTANT:
```

### Code example:

```python
from transformers import AutoModelForCausalLM, AutoTokenizer
import torch, json
# model path
model_path = "NurtureAI/SynthIA-7B-v2.0-16k"
output_file_path = "./SynthIA-7B-v2.0-conversations.jsonl"
device_map = {"": "cuda"}
model = AutoModelForCausalLM.from_pretrained(
    model_path,
    torch_dtype=torch.float16,
    device_map=device_map,
    load_in_8bit=False,
    trust_remote_code=True,
)
# tokenizer
tokenizer = AutoTokenizer.from_pretrained(model_path, trust_remote_code=True)

def generate_text(instruction):
    tokens = tokenizer.encode(instruction)
    tokens = torch.LongTensor(tokens).unsqueeze(0)
    tokens = tokens.to("cuda")
    instance = {
        "input_ids": tokens,
        "top_p": 1.0,
        "temperature": 0.75,
        "generate_len": 1024,
        "top_k": 50,
    }
    length = len(tokens[0])
    with torch.no_grad():
        rest = model.generate(
            input_ids=tokens,
            max_length=length + instance["generate_len"],
            use_cache=True,
            do_sample=True,
            top_p=instance["top_p"],
            temperature=instance["temperature"],
            top_k=instance["top_k"],
            num_return_sequences=1,
        )
    output = rest[0][length:]
    string = tokenizer.decode(output, skip_special_tokens=True)
    answer = string.split("USER:")[0].strip()
    return f"{answer}"


conversation = f"SYSTEM: Elaborate on the topic using a Tree of Thoughts and backtrack when necessary to construct a clear, cohesive Chain of Thought reasoning. Always answer without hesitation."


while True:
    user_input = input("You: ")
    llm_prompt = f"{conversation} \nUSER: {user_input} \nASSISTANT: "
    answer = generate_text(llm_prompt)
    print(answer)
    conversation = f"{llm_prompt}{answer}"
    json_data = {"prompt": user_input, "answer": answer}

    ## Save your conversation
    with open(output_file_path, "a") as output_file:
        output_file.write(json.dumps(json_data) + "\n")

```