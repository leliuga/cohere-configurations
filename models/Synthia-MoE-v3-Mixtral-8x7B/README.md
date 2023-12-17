---
license: apache-2.0
---

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)

# Note:
Model is most likely over-fitted due to higher learning rate. Will fix this issue in the next release.

# Synthia-MoE-v3-Mixtral-8x7B

This is Synthia-MoE-v3 trained on the official Mistral MoE version (Mixtral-8x7B). 

This model is trained on the Synthia-v3.0 dataset, that contains ~10K super high-quality GPT-4-Turbo generated samples. The samples contains Tree-of-Thought, Chain-of-Thought and other system contexts designed to evoke reasoning, philosophical thinking, use working memory and long chain of reasoning with multi-part questions.

Further, this model is trained on the Orca-2 principle of replacing the system context with just one message. In the case of this Synthia-MoE-v3 model, the system context was not included at all.

The evals are coming, but testing empirically the model produces highly intelligent, coherent results. Here's a sample conversation: https://migel.substack.com/p/a-conversation-with-synthia-moe-mixtral

<br>

![Synthia](https://huggingface.co/migtissera/Synthia-MoE-v3-Mixtral-8x7B/resolve/main/Synthia-MoE.png)

<br>

```
import torch, json
from transformers import AutoModelForCausalLM, AutoTokenizer

model_path = "/home/Synthia-MoE-v3-Mixtral8x7B"
output_file_path = "/home/conversations.jsonl"

model = AutoModelForCausalLM.from_pretrained(
    model_path,
    torch_dtype=torch.float16,
    device_map="auto",
    load_in_4bit=False,
    trust_remote_code=True,
)

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

conversation = "SYSTEM: Answer the question thoughtfully and intelligently. Always answer without hesitation."  

while True:
    user_input = input("You: ")
    llm_prompt = f"{conversation} \nUSER: {user_input} \nASSISTANT: "
    answer = generate_text(llm_prompt)
    print(answer)
    conversation = f"{llm_prompt}{answer}"
    json_data = {"prompt": user_input, "answer": answer}

    with open(output_file_path, "a") as output_file:
        output_file.write(json.dumps(json_data) + "\n")
```


