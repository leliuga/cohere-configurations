---
license: llama2
pipeline_tag: text-generation
language:
- en
library_name: transformers
---

# Synthia-70B-v1.1
SynthIA (Synthetic Intelligent Agent) is a LLama-2-70B model trained on Orca style datasets. It has been fine-tuned for instruction following as well as having long-form conversations.

This model has generalized "Tree of Thought" reasoning capabilities. Evoke it with the following system message:
```
Elaborate on the topic using a Tree of Thoughts and backtrack when necessary to construct a clear, cohesive Chain of Thought reasoning
```

<br>

![Synthia](https://huggingface.co/migtissera/Synthia-70B-v1.1/resolve/main/Synthia.jpeg)

<br>

<br>

#### License Disclaimer:

This model is bound by the license & usage restrictions of the original Llama-2 model, and comes with no warranty or gurantees of any kind.

<br>

## Evaluation

We evaluated Synthia-70B on a wide range of tasks using [Language Model Evaluation Harness](https://github.com/EleutherAI/lm-evaluation-harness) from EleutherAI. 

Here are the results on metrics used by [HuggingFaceH4 Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)

||||
|:------:|:--------:|:-------:|
|**Task**|**Metric**|**Value**|
|*arc_challenge*|acc_norm|70.05|
|*hellaswag*|acc_norm|87.12|
|*mmlu*|acc_norm|70.34|
|*truthfulqa_mc*|mc2|57.84|
|**Total Average**|-|**71.34**||

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

model_path = "migtissera/Synthia-70B"
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
@misc{Synthia-70B,
  author = {Migel Tissera},
  title = {Synthia-70B: Synthetic Intelligent Agent},
  year = {2023},
  publisher = {GitHub, HuggingFace},
  journal = {GitHub repository, HuggingFace repository},
  howpublished = {\url{https://huggingface.co/migtissera/Synthia-70B},
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
You: What is Earth's atmosphere like? What is Earth's magnetic field like? What is the relationship between the two?
                                                                                                                                                                                                                                                                                              
Earth's atmosphere is a layer of gases that surrounds the planet, held in place by Earth's gravity. It is primarily composed of nitrogen (78.1%), oxygen (20.9%), argon (0.9%), and other trace gases, such as carb
on dioxide and water vapor. The atmosphere plays a crucial role in regulating Earth's temperature and climate by trapping heat and absorbing harmful radiation from the sun. It also provides the air we breathe, p
rotects us from meteorites and the sun's ultraviolet radiation, and allows for the development of life on the planet.                                                                                              
                                                                                                                                                                                                                   
Earth's magnetic field is a phenomenon generated by the convection of molten iron in the planet's outer core, which creates an electric current. This current, in turn, produces a magnetic field that extends from
 the Earth's interior to space, protecting the planet from the solar wind and cosmic radiation. The magnetic field also plays a role in the formation of Earth's Van Allen radiation belts, which trap high-energy 
particles from the sun and cosmic rays.                                                                                                                                                                            
                                                                                                                                                                                                                   
The relationship between Earth's atmosphere and magnetic field is complex and multifaceted. The magnetic field influences the distribution of charged particles in the atmosphere, affecting auroras, radio communi
cation, and the ionosphere's behavior. The atmosphere, in turn, plays a role in the distribution of the magnetic field's energy, as well as in the generation and loss of atmospheric ions. The interactions betwee
n the atmosphere and the magnetic field also influence the dynamics of Earth's climate and the long-term evolution of the planet's environment.                                                                    

You: Elaborate the the complex relationship between the two
                                                                                                                                                     
The complex relationship between Earth's atmosphere and magnetic field can be elaborated upon in several ways. Firstly, the magnetic field influences the distribution of charged particles in the atmosphere. For 
example, the magnetic field deflects the solar wind particles that would otherwise strip away the atmosphere. This deflection creates a bow shock, where the solar wind particles slow down and dissipate their ene
rgy, producing a shock wave that propagates downstream from the Earth. The magnetic field also guides charged particles from the solar wind into the inner magnetosphere, where they become trapped and produce the
 Van Allen radiation belts.                                                                                                                                                                                        
                                                                                                                                                                                                                   
Secondly, the atmosphere plays a role in the distribution of the magnetic field's energy. The ionosphere, which is the upper layer of the atmosphere that is ionized by solar radiation, is affected by the magneti
c field. The energy from the magnetic field is transferred to the ionosphere, where it drives the motion of the charged particles and affects the propagation of radio waves.                                      
                                                                                                                                                                                                                   
Thirdly, the interactions between the atmosphere and the magnetic field influence the dynamics of Earth's climate. The magnetic field affects the distribution of charged particles in the atmosphere, which in tur
n affects the formation of clouds and precipitation. The magnetic field also affects the distribution of heat in the atmosphere, which can influence the patterns of atmospheric circulation and the distribution o
f temperatures around the planet.                                                                                                                                                                                  
                                                                                                                                                                                                                   
Finally, the interactions between the atmosphere and the magnetic field play a role in the long-term evolution of the planet's environment. The magnetic field protects the atmosphere from the erosive effects of 
the solar wind, which would otherwise strip away the atmosphere over time. The magnetic field also affects the distribution of heat in the atmosphere, which can influence the patterns of atmospheric circulation 
and the distribution of temperatures around the planet. Over time, these interactions can shape the evolution of the planet's climate and environment.

You: Please explain the bow shock wave, Van Allen radiation belts in detail
                                                                                                                                      
The bow shock wave is a phenomenon that occurs when the solar wind, a stream of charged particles emitted by the sun, interacts with Earth's magnetic field. The solar wind particles are moving at high speeds, an
d when they encounter the Earth's magnetic field, they slow down and dissipate their energy. This creates a shock wave that propagates downstream from the Earth. The bow shock wave is located about 15 Earth radi
i (approximately 93,000 miles) upstream from Earth, and it marks the boundary between the solar wind and Earth's magnetosphere.

The Van Allen radiation belts are zones of high-energy charged particles, primarily protons and electrons, that are trapped by Earth's magnetic field. The belts are divided into two main regions: the inner belt,
 which extends from about 1,000 to 6,000 kilometers above Earth's surface, and the outer belt, which extends from about 13,000 to 65,000 kilometers above the surface. The particles in the belts are trapped by Ea
rth's magnetic field, and they can remain in the belts for long periods of time.

The particles in the Van Allen belts are generated by a variety of processes, including the impact of solar wind particles on the Earth's atmosphere and the decay of radioactive elements in the planet's crust. T
he belts are important for understanding the behavior of charged particles in Earth's magnetosphere and for studying the effects of space weather on human activities in space. They can also pose a hazard to spac
ecraft and satellites, as the high-energy particles in the belts can damage sensitive electronics and cause other problems.

You: Does the Van Allen belt present any danger to Geo stationary satellites? Where are these belts? Are they near any of the orbits?
  
Yes, the Van Allen radiation belts can present a danger to geostationary satellites and other spacecraft. The high-energy charged particles in the belts can damage sensitive electronics and cause other problems,
 including radiation damage and single-event upsets. Geostationary satellites are located at an altitude of approximately 35,786 kilometers above Earth's surface, which is just outside the inner belt but still w
ithin the outer belt.

The Van Allen radiation belts are located in Earth's magnetosphere, which is the region of space around the planet where the magnetic field is dominant. The inner belt is located between about 1,000 and 6,000 ki
lometers above the surface, while the outer belt is located between about 13,000 and 65,000 kilometers above the surface. The belts are not located in any specific orbit, but they are present throughout the magn
etosphere, which covers a large portion of the space around the planet.

To mitigate the effects of the Van Allen belts on spacecraft, engineers use a variety of techniques, including shielding and designing electronics to be resistant to radiation. Additionally, some spacecraft are 
designed to operate in orbits that avoid the Van Allen belts, while others are designed to operate within the belts and withstand the effects of the high-energy particles.

You: So when we send spacecraft to the moon, say the ones carrying humans, how do we make sure that the solar wind charged particles will not affect the mission? Are there very specific remedies that we can take
?

When sending spacecraft to the moon, there are several steps that can be taken to protect the spacecraft and its occupants from the effects of solar wind charged particles. One approach is to design the spacecra
ft with a strong magnetic field, which can help deflect the charged particles away from the spacecraft. Another approach is to use radiation shielding, which can help protect the electronics and any human occupa
nts from the effects of radiation.

In addition to these measures, space agencies can also use data from satellites and other spacecraft to monitor the level of solar wind activity and predict when there may be an increased risk of charged particl
e events. This information can be used to adjust the timing of missions and other activities to minimize exposure to high-energy particles.

It's also important to note that the moon is outside the magnetosphere and therefore not directly affected by the Van Allen belts. However, the moon does receive a significant amount of radiation from the sun an
d other sources, so measures to protect the spacecraft and its occupants from radiation exposure are still necessary.

Overall, protecting spacecraft and humans from the effects of solar wind charged particles requires a combination of engineering and operational measures, including strong magnetic fields, radiation shielding, a
nd careful mission planning.

You: Does spacecraft actually create strong magnetic fields? Which Apollo spacecrafts did this? Or is this hypothetical?

Yes, spacecraft can create strong magnetic fields, although this is not typically a primary design consideration for most spacecraft. Some spacecraft, such as those designed for plasma physics research, may incl
ude a magnetic field generator to create controlled conditions for experiments. However, for most spacecraft, the focus is on using existing magnetic fields, such as those generated by the Earth's magnetic field
, to protect the spacecraft and its occupants.

The Apollo spacecraft did not create strong magnetic fields, as their primary focus was on propulsion, navigation, and life support systems for the astronauts. However, some of the Apollo missions did include sc
ientific experiments to study the effects of solar wind and other charged particles on the spacecraft and its occupants.

The use of magnetic fields to protect spacecraft and humans from the effects of radiation is an ongoing area of research and development, and there are a number of proposals and concepts for how this might be ac
hieved. For example, some researchers have proposed using superconducting coils to create a magnetic field around the spacecraft, which could help deflect charged particles and other forms of radiation. However,
 these proposals are still in the early stages of development and are not yet ready for implementation on spacecraft.
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_migtissera__Synthia-70B-v1.1)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 62.84   |
| ARC (25-shot)         | 70.05          |
| HellaSwag (10-shot)   | 87.12    |
| MMLU (5-shot)         | 70.34         |
| TruthfulQA (0-shot)   | 57.84   |
| Winogrande (5-shot)   | 83.66   |
| GSM8K (5-shot)        | 31.84        |
| DROP (3-shot)         | 39.02         |
