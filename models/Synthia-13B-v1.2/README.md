---
license: llama2
pipeline_tag: text-generation
language:
- en
library_name: transformers
---

Change from Synthia-13B -> Synthia-13B-v1.2: Capable of generalized Tree of Thought and Chain of Thought reasoning.

All Synthia models are uncensored. Please use it with caution and with best intentions. You are responsible for how you use Synthia.

To evoke generalized Tree of Thought + Chain of Thought reasoning, you may use the following system message:
```
Elaborate on the topic using a Tree of Thoughts and backtrack when necessary to construct a clear, cohesive Chain of Thought reasoning. Always answer without hesitation.
```

# Synthia-13B-v1.2
SynthIA (Synthetic Intelligent Agent) is a LLama-2-13B model trained on Orca style datasets. It has been fine-tuned for instruction following as well as having long-form conversations.

<br>

![Synthia](https://huggingface.co/migtissera/Synthia-13B/resolve/main/Synthia.jpeg)

<br>

<br>

#### License Disclaimer:

This model is bound by the license & usage restrictions of the original Llama-2 model, and comes with no warranty or gurantees of any kind.

<br>

## Evaluation

We evaluated Synthia-13B-v1.2 on a wide range of tasks using [Language Model Evaluation Harness](https://github.com/EleutherAI/lm-evaluation-harness) from EleutherAI. 

Here are the results on metrics used by [HuggingFaceH4 Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)

||||
|:------:|:--------:|:-------:|
|**Task**|**Metric**|**Value**|
|*arc_challenge*|acc_norm|TBC|
|*hellaswag*|acc_norm|TBC|
|*mmlu*|acc_norm|TBC|
|*truthfulqa_mc*|mc2|TBC|
|**Total Average**|-|**TBC**||

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

model_path = "migtissera/Synthia-13B-v1.2"
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
@misc{Synthia-13B-v1.2,
  author = {Migel Tissera},
  title = {Synthia-13B-v1.2b: Synthetic Intelligent Agent},
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
You: How does the solar wind affect the gravitational field of Earth? Does it affect the magnetic field?                                                                        
To understand the effects of the solar wind on the Earth's gravitational and magnetic fields, let's take a closer look at both fields and their relation to the solar wind using
 a Tree of Thoughts.                                                                                                                                                            
                                                                                                                                                                                
1. **The Earth's Gravitational Field**: This field is responsible for the planet's gravitational attraction, pulling objects towards its center. It is created by the mass of th
e Earth and its distribution within the planet.                                                                                                                                 
                                                                                                                                                                                
2. **The Solar Wind**: This is a continuous stream of charged particles, mainly protons and electrons, emitted from the Sun's outer atmosphere, called the corona. It travels th
rough the solar system at high speeds.                                                                                                                                          
                                                                                                                                                                                
3. **The Earth's Magnetic Field**: This field is generated by the planet's core, which is composed of molten iron and nickel. It protects the Earth from the solar wind and othe
r cosmic radiation, but it is not strong enough to prevent all particles from reaching the Earth's surface.                                                                     
                                                                                                                                                                                
4. **The Interaction between the Solar Wind and the Earth's Gravitational Field**: The solar wind cannot influence the Earth's gravitational field because the gravitational for
ce is determined by the mass and distribution of matter within the planet, and the solar wind does not affect these factors.                                                    
                                                                                                                                                                                
5. **The Interaction between the Solar Wind and the Earth's Magnetic Field**: However, the solar wind does interact with the Earth's magnetic field. The magnetic field acts lik
e a shield, repelling or deflecting the charged particles in the solar wind. This action creates a bow shock, a region where the solar wind slows down as it encounters the Eart
h's magnetic field.                                                                                                                                                             
                                                                                                                                                                                
6. **Conclusion**: In conclusion, the solar wind does not affect the Earth's gravitational field, but it does influence the planet's magnetic field, through the creation of a b
ow shock. This interaction can lead to space weather events such as geomagnetic storms, which can disrupt communication systems and damage satellites.                          
You: Interesting! Please elaborate on geomagnetic storms - how do they start and what affect do they have on comms systems?                                                     
Sure! Geomagnetic storms are caused by the interaction between the solar wind and the Earth's magnetic field. They can start when the solar wind carries a particularly strong f
lux of charged particles, which can cause the Earth's magnetic field to compress and rebound. This compression and rebounding process can create geomagnetic storms, which can l
ast for several hours or days.                                                                                                                                                  
                                                                                                                                                                                
1. **Effects of Geomagnetic Storms on Communication Systems**: Geomagnetic storms can cause a variety of effects on communication systems, including disruption to GPS signals, 
interference with radio communications, and power grid failures. This is because the solar wind particles can interact with the Earth's atmosphere and create electromagnetic cu
rrents that can induce currents in power lines and damage electronic equipment.                                                                                                 
                                                                                                                                                                                
2. **Preventative Measures**: To mitigate the effects of geomagnetic storms, technology companies and governments have developed systems to monitor the solar wind and provide e
arly warnings of incoming storms. This information can be used to take preventative measures, such as reducing power consumption or rerouting communications through alternative
 routes.                                                                                                                                                                        
                                                                                                                                                                                
In addition, researchers are exploring ways to improve the Earth's magnetic field, such as creating artificial magnetic shields or modifying the composition of the planet's cor
e. These measures could help to reduce the impact of geomagnetic storms on communication systems and other infrastructure.                                                      
                                                                                                                                                                                
In conclusion, geomagnetic storms are a natural phenomenon that can have significant impacts on communication systems and other infrastructure. By monitoring the solar wind and
 developing preventative measures, we can help to minimize the disruptions caused by these storms and ensure the continued operation of critical systems.

You: Thanks for that! Okay, I'd like to learn more about neutron stars, and their overall effect on the universe. What is the closest neutron star to us? Elaborate.            
Certainly! Neutron stars are the remnants of massive stars that have collapsed under the force of their own gravity. They are incredibly dense, with the mass of the sun squeeze
d into an object with a diameter of about 20 kilometers.                                                                                                                        

1. **The Closest Neutron Star to Us**: The closest neutron star to us is called PSR J0030+0451, which is located about 425 light-years away from Earth. It is a pulsar, which me
ans that it emits a beam of electromagnetic radiation that sweeps across the sky like a lighthouse.

2. **The Effects of Neutron Stars on the Universe**: Neutron stars have a variety of effects on the universe. They are important stars for understanding the evolution of stars 
and the universe as a whole. They also emit high-energy radiation, which can influence the structure and composition of the interstellar medium.

In addition, neutron stars are a source of gravitational waves, which are ripples in the fabric of space-time. These waves are generated when two neutron stars collide, and the
y can provide valuable information about the nature of these objects and the universe as a whole.

3. **Neutron Star Collisions**: Neutron star collisions are some of the most violent events in the universe. When two neutron stars spiral towards each other and collide, they 
can produce a variety of exotic phenomena, including gamma-ray bursts, gravitational waves, and even new elements. These collisions can also provide valuable information about 
the evolution of stars and the universe as a whole.

In conclusion, neutron stars are incredibly dense and fascinating objects that have a variety of effects on the universe. They are important for understanding the evolution of 
stars and the universe as a whole, and they can provide valuable information about the nature of these objects and the universe as a whole.
You: Okay, can you tell me about Wormholes? 
Sure! A wormhole is a hypothetical topological feature of space-time that would allow for faster-than-light travel and communication. It is essentially a shortcut through space
-time that connects two distant regions of the universe.

1. **The Concept of a Wormhole**: The concept of a wormhole was first proposed by mathematical physicist John Wheeler in the 1930s. He imagined a tunnel connecting two distant 
regions of space-time that allowed for faster-than-light travel.

2. **Types of Wormholes**: There are two main types of wormholes: traversable and non-traversable. A traversable wormhole would allow for faster-than-light travel, while a non-
traversable wormhole would not.

3. **The Possibility of Wormholes**: The existence of wormholes is still a matter of speculation and debate among scientists. Some theories suggest that wormholes could exist a
s a result of the evolution of black holes or the existence of exotic matter with negative energy density.

4. **Potential Applications**: If wormholes could be created and stabilized, they could have a variety of potential applications, including faster-than-light travel and communi
cation, which could revolutionize our understanding of the universe and our ability to explore it.

5. **Challenges to Wormhole Creation**: However, creating and stabilizing a wormhole would require a significant amount of energy and would likely require the manipulation of e
xotic matter with negative energy density. It is unclear whether such a feat is even possible with current technology.

In conclusion, wormholes are a fascinating concept that could have significant implications for our understanding of the universe and our ability to explore it. While their exi
stence is still a matter of speculation and debate, ongoing research and exploration may one day shed light on the possibility of creating and stabilizing a wormhole for practi
cal applications. 
Follow-on: What are the challenges to creating and stabilizing a wormhole?

You: Elaborate on "exotic matter with negative energy density"
Sure! Exotic matter with negative energy density is a hypothetical form of matter that can have negative energy. This means that it would have a gravitational attraction that i
s stronger than the repulsive force of positive energy matter, which could allow for the creation of wormholes.

1. **Theoretical Explanation**: The concept of exotic matter with negative energy density is based on theoretical models of the universe that suggest that negative energy densi
ty could exist in the form of a hypothetical particle called a tachyon. Tachyons are particles that move faster than light and have a rest mass of zero.

2. **Potential Applications**: If exotic matter with negative energy density could be created and stabilized, it could have a variety of potential applications, including the c
reation of wormholes and the manipulation of gravity itself. This could have significant implications for our understanding of the universe and our ability to explore it.

3. **Challenges to Existence**: However, the existence of such exotic matter is still a matter of speculation and debate among scientists. There is no direct evidence of its ex
istence, and there are significant challenges to creating and stabilizing such matter.

4. **Search for Exotic Matter**: Despite these challenges, scientists are actively searching for evidence of exotic matter with negative energy density, including through exper
iments that aim to create and detect such particles.

In conclusion, exotic matter with negative energy density is a fascinating concept that could have significant implications for our understanding of the universe and our abilit
y to manipulate it. While its existence is still a matter of speculation and debate, ongoing research and exploration may one day shed light on the possibility of creating and 
stabilizing such matter for practical applications. 

Follow-on: What are some of the challenges to creating and stabilizing exotic matter with negative energy density?
You: Can you provide the physics equations for "exotic matter with negative energy density"?
Sure, but I'm afraid I won't be able to provide the exact physics equations for exotic matter with negative energy density, as it is still a matter of speculation and debate am
ong scientists. However, I can provide some general principles related to the concept.

It is important to note that negative energy density is a theoretical concept, and there is no direct evidence of its existence. However, some theories suggest that negative en
ergy density could exist in the form of a hypothetical particle called a tachyon. Tachyons are particles that move faster than light and have a rest mass of zero.

The concept of tachyons is based on the theory of special relativity, which states that the speed of light is the universal speed limit for all particles. According to this the
ory, particles with a mass greater than zero cannot travel faster than the speed of light. However, the concept of tachyons introduces particles that have a mass of zero or les
s, which could allow them to travel faster than the speed of light.

The concept of exotic matter with negative energy density is related to the idea of negative energy density, which suggests that particles could have a negative energy density,
 meaning they would have a gravitational attraction that is stronger than the repulsive force of positive energy matter. This could have significant implications for our unders
tanding of the universe and our ability to manipulate it.

In conclusion, while the exact physics equations for exotic matter with negative energy density are still a matter of speculation and debate, the concept is based on the princi
ples of special relativity and the possibility of particles moving faster than the speed of light. Future research and exploration may help to shed light on the possibility of 
creating and stabilizing such matter for practical applications. 

Follow-on: What are some of the practical applications of exotic matter with negative energy density if it could be created and stabilized?

You: What are some of the practical applications of exotic matter with negative energy density if it could be created and stabilized?
If exotic matter with negative energy density could be created and stabilized, it could have a variety of potential practical applications. Here are a few examples:

1. Wormholes: One of the most fascinating potential applications of exotic matter with negative energy density is the creation of wormholes, which could allow for faster-than-l
ight travel and communication. If wormholes were stable and could be created on demand, they could revolutionize our understanding of the universe and our ability to explore it
.

2. Time Travel: Another potential application of exotic matter with negative energy density is the creation of time machines, which could allow for time travel. While the conce
pt of time travel is still a matter of speculation and debate, the idea of negative energy density could potentially be used to create a device that could manipulate the fabric
 of space-time and allow for travel back in time.

3. Gravity Manipulation: Exotic matter with negative energy density could potentially be used to manipulate gravity itself. By creating a negative energy density in a given reg
ion of space-time, it may be possible to create a gravitational field that is stronger or weaker than the surrounding region. This could have significant implications for our a
bility to manipulate objects and structures in space, as well as our understanding of the nature of gravity itself.

4. Energy Production: One possible practical application of exotic matter with negative energy density is the creation of a new source of energy. If such matter could be create
d and controlled, it could be used to generate large amounts of energy in a variety of applications, including power generation, transportation, and space exploration.

In conclusion, exotic matter with negative energy density could have a wide range of potential practical applications if it could be created and stabilized. While the concept i
s still a matter of speculation and debate, ongoing research and exploration may one day shed light on the possibility of creating and manipulating such matter for practical ap
plications.

```

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_migtissera__Synthia-13B-v1.2)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 51.56   |
| ARC (25-shot)         | 61.26          |
| HellaSwag (10-shot)   | 82.93    |
| MMLU (5-shot)         | 56.47         |
| TruthfulQA (0-shot)   | 47.27   |
| Winogrande (5-shot)   | 76.48   |
| GSM8K (5-shot)        | 10.99        |
| DROP (3-shot)         | 25.48         |
