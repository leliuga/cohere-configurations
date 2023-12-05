---
license: llama2
pipeline_tag: text-generation
language:
- en
library_name: transformers
---

<br>

![Synthia](https://huggingface.co/migtissera/Synthia-13B/resolve/main/Synthia.jpeg)

<br>

Change from 1.1 -> 1.2: 20% more data than 1.1 and 2x training time.

All Synthia models are uncensored. Please use it with caution and with best intentions. You are responsible for how you use Synthia.

To evoke generalized Tree of Thought + Chain of Thought reasoning, you may use the following system message:
```
Elaborate on the topic using a Tree of Thoughts and backtrack when necessary to construct a clear, cohesive Chain of Thought reasoning. Always answer without hesitation.
```

# Synthia-70B-v1.2
SynthIA (Synthetic Intelligent Agent) is a LLama-2-70B model trained on Orca style datasets. It has been fine-tuned for instruction following as well as having long-form conversations.


<br>

#### License Disclaimer:

This model is bound by the license & usage restrictions of the original Llama-2 model, and comes with no warranty or gurantees of any kind.

<br>

## Evaluation

We evaluated Synthia-70B-v1.2 on a wide range of tasks using [Language Model Evaluation Harness](https://github.com/EleutherAI/lm-evaluation-harness) from EleutherAI. 

Here are the results on metrics used by [HuggingFaceH4 Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)

||||
|:------:|:--------:|:-------:|
|**Task**|**Metric**|**Value**|
|*arc_challenge*|acc_norm|70.48|
|*hellaswag*|acc_norm|86.98|
|*mmlu*|acc_norm|70.13|
|*truthfulqa_mc*|mc2|58.64|
|**Total Average**|-|**71.56**||

<br>

## Example Usage

### Here is prompt format:

```
SYSTEM: Elaborate on the topic using a Tree of Thoughts and backtrack when necessary to construct a clear, cohesive Chain of Thought reasoning. Always answer without hesitation.
USER: How is a rocket launched from the surface of the earth to Low Earth Orbit?
ASSISTANT:
```

### Below shows a code example on how to use this model:

```python
import torch, json
from transformers import AutoModelForCausalLM, AutoTokenizer

model_path = "migtissera/Synthia-70B-v1.2"
output_file_path = "./Synthia-70B-conversations.jsonl"

model = AutoModelForCausalLM.from_pretrained(
    model_path,
    torch_dtype=torch.float16,
    device_map="auto",
    load_in_8bit=False,
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

<br>

#### Limitations & Biases:

While this model aims for accuracy, it can occasionally produce inaccurate or misleading results. 

Despite diligent efforts in refining the pretraining data, there remains a possibility for the generation of inappropriate, biased, or offensive content. 

Exercise caution and cross-check information when necessary. This is an uncensored model.


<br>

### Citiation:

Please kindly cite using the following BibTeX:

```
@misc{Synthia-70B-v1.2,
  author = {Migel Tissera},
  title = {Synthia-70B-v1.2: Synthetic Intelligent Agent},
  year = {2023},
  publisher = {GitHub, HuggingFace},
  journal = {GitHub repository, HuggingFace repository},
  howpublished = {\url{https://huggingface.co/migtissera/Synthia-13B},
}
```

```
@misc{mukherjee2023orca,
      title={Orca: Progressive Learning from Complex Explanation Traces of GPT-4}, 
      author={Subhabrata Mukherjee and Arindam Mitra and Ganesh Jawahar and Sahaj Agarwal and Hamid Palangi and Ahmed Awadallah},
      year={2023},
      eprint={2306.02707},
      archivePrefix={arXiv},
      primaryClass={cs.CL}
}
```

```
@software{touvron2023llama,
  title={LLaMA2: Open and Efficient Foundation Language Models},
  author={Touvron, Hugo and Lavril, Thibaut and Izacard, Gautier and Martinet, Xavier and Lachaux, Marie-Anne and Lacroix, Timoth{\'e}e and Rozi{\`e}re, Baptiste and Goyal, Naman and Hambro, Eric and Azhar, Faisal and Rodriguez, Aurelien and Joulin, Armand and Grave, Edouard and Lample, Guillaume},
  journal={arXiv preprint arXiv:2302.13971},
  year={2023}
}
```


# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_migtissera__Synthia-70B-v1.2)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 63.41   |
| ARC (25-shot)         | 70.48          |
| HellaSwag (10-shot)   | 86.98    |
| MMLU (5-shot)         | 70.13         |
| TruthfulQA (0-shot)   | 58.64   |
| Winogrande (5-shot)   | 83.27   |
| GSM8K (5-shot)        | 31.92        |
| DROP (3-shot)         | 42.42         |
