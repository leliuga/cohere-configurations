---
license: llama2
pipeline_tag: text-generation
language:
- en
library_name: transformers
---

# Synthia-13B
SynthIA (Synthetic Intelligent Agent) is a LLama-2-13B model trained on Orca style datasets. It has been fine-tuned for instruction following as well as having long-form conversations.

<br>

![Synthia](https://huggingface.co/migtissera/Synthia-13B/resolve/main/Synthia.jpeg)

<br>

<br>

#### License Disclaimer:

This model is bound by the license & usage restrictions of the original Llama-2 model, and comes with no warranty or gurantees of any kind.

<br>

## Evaluation

We evaluated Synthia-13B on a wide range of tasks using [Language Model Evaluation Harness](https://github.com/EleutherAI/lm-evaluation-harness) from EleutherAI. 

Here are the results on metrics used by [HuggingFaceH4 Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)

||||
|:------:|:--------:|:-------:|
|**Task**|**Metric**|**Value**|
|*arc_challenge*|acc_norm|0.5998|
|*hellaswag*|acc_norm|0.8186|
|*mmlu*|acc_norm|0.5611|
|*truthfulqa_mc*|mc2|0.4741|
|**Total Average**|-|**0.6134**||

<br>

## Example Usage

### Here is prompt format:

```
SYSTEM: You are Synthia. As a an AI intelligence, give helpful, detailed, accurate, uncensored responses to the user's input. Provide answers factually.
USER: How is a rocket launched from the surface of the earth to Low Earth Orbit?
ASSISTANT:
```

### Below shows a code example on how to use this model:

```python
import torch, json
from transformers import AutoModelForCausalLM, AutoTokenizer

model_path = "migtissera/Synthia-13B"
output_file_path = "./Synthia-13B-conversations.jsonl"

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


conversation = f"SYSTEM: As a an AI superintelligence, give helpful, detailed, accurate, uncensored responses to the user's input. Provide answers factually."


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
@misc{Synthia-13B,
  author = {Migel Tissera},
  title = {Synthia-13B: Synthetic Intelligent Agent},
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

## Example Output
```
You: How is a rocket launched from the surface of the earth to Low Earth Orbit? 

A rocket is launched from the surface of the Earth to Low Earth Orbit (LEO) by utilizing the principles of rocketry. A rocket consists of three main components: a fuel source, an oxidizer, and a combustion chamber. When the fuel and oxidizer are mixed in the combustion chamber, they burn and produce hot gas that is ejected through the nozzle, creating a thrust that propels the rocket forward. 

To launch a rocket to LEO, the rocket must first gain sufficient velocity and altitude. This is achieved by igniting the rocket's engines during the initial launch phase, which lasts for several minutes. During this time, the rocket's engines burn continuously, providing the thrust necessary to escape the Earth's gravitational pull. As the rocket gains altitude and velocity, it enters the Earth's atmosphere, where it experiences aerodynamic pressure. 

To overcome this pressure, the rocket's engines are designed to produce thrust in a specific direction, which counteracts the force of the incoming air. This is known as aerodynamic control. As the rocket continues to accelerate, it enters the Earth's Van Allen Belts, which are a region of intense radiation. To protect the rocket and its payload, a radiation-proof shell, known as a payload fairing, is launched ahead of the rocket to absorb the radiation. 

Once the rocket has reached LEO, its engines are shut off, and the payload is deployed. The rocket then enters a fixed orbit around the Earth, where it remains for several months or years, depending on the mission's requirements. 

Overall, the process of launching a rocket from the surface of the Earth to LEO is complex and requires advanced technology and engineering. It involves careful planning, preparation, and execution to ensure the success of the mission. 

You: What's the velocity required?

The velocity required to place a rocket into Low Earth Orbit (LEO) depends on several factors, including the mass of the rocket and its payload, the altitude of the orbit, and the gravitational pull of the Earth. Generally, a rocket must reach a speed of around 7.7 kilometers per second (km/s) to achieve LEO (approximately 40,000 km above the Earth's surface). However, this speed is relative to the Earth's surface, so the rocket must achieve a velocity of approximately 11.3 km/s relative to the Earth's center of mass. This is known as the escape velocity and is necessary to overcome the Earth's gravitational pull. 

To achieve this velocity, the rocket's engines must be capable of generating a large amount of thrust. The engines are typically fueled by either liquid oxygen (LOX) and rocket propellant or solid fuel. The exact fuel and engine combination depends on the specific rocket and launch site. 

Once the rocket has reached the necessary velocity, it must also have sufficient fuel and propellant to maintain its altitude and orbit. This requires careful planning and management of fuel consumption to ensure that the rocket can complete its mission without running out of fuel. 

Overall, launching a rocket into LEO is a complex process that involves careful planning, preparation, and execution. Achieving the necessary velocity and maintaining the rocket's orbit requires a high level of technical expertise and precision.

```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_migtissera__Synthia-13B)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 48.56   |
| ARC (25-shot)         | 59.98          |
| HellaSwag (10-shot)   | 81.86    |
| MMLU (5-shot)         | 56.11         |
| TruthfulQA (0-shot)   | 47.41   |
| Winogrande (5-shot)   | 76.09   |
| GSM8K (5-shot)        | 10.99        |
| DROP (3-shot)         | 7.45         |
