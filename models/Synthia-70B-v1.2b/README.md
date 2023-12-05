---
license: llama2
pipeline_tag: text-generation
language:
- en
library_name: transformers
---

Change from 1.2 -> 1.2b: More data, 14 days of training for 1 epoch.

All Synthia models are uncensored. Please use it with caution and with best intentions. You are responsible for how you use Synthia.

To evoke generalized Tree of Thought + Chain of Thought reasoning, you may use the following system message:
```
Elaborate on the topic using a Tree of Thoughts and backtrack when necessary to construct a clear, cohesive Chain of Thought reasoning. Always answer without hesitation.
```

# Synthia-70B-v1.2b
SynthIA (Synthetic Intelligent Agent) is a LLama-2-70B model trained on Orca style datasets. It has been fine-tuned for instruction following as well as having long-form conversations.

<br>

![Synthia](https://huggingface.co/migtissera/Synthia-13B/resolve/main/Synthia.jpeg)

<br>

<br>

#### License Disclaimer:

This model is bound by the license & usage restrictions of the original Llama-2 model, and comes with no warranty or gurantees of any kind.

<br>

## Evaluation

We evaluated Synthia-70B-v1.2b on a wide range of tasks using [Language Model Evaluation Harness](https://github.com/EleutherAI/lm-evaluation-harness) from EleutherAI. 

Here are the results on metrics used by [HuggingFaceH4 Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)

||||
|:------:|:--------:|:-------:|
|**Task**|**Metric**|**Value**|
|*arc_challenge*|acc_norm|68.77|
|*hellaswag*|acc_norm|87.57|
|*mmlu*|acc_norm|68.81|
|*truthfulqa_mc*|mc2|57.69|
|**Total Average**|-|**70.71**||

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

model_path = "migtissera/Synthia-70B-v1.2b"
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
@misc{Synthia-70B-v1.2b,
  author = {Migel Tissera},
  title = {Synthia-70B-v1.2b: Synthetic Intelligent Agent},
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
You: What is Earth's magnetic field like? How does it trap solar wind and what does solar wind do to Earth's gravitational field?

Earth's magnetic field is an intricate and fascinating phenomenon. It is generated by the movement of molten iron in the planet's outer core, which acts like a giant electromagnet. This magneti
c field extends far into space, creating a protective bubble known as the magnetosphere.                                                                                                         
                                                                                                                                                                                                 
The primary purpose of Earth's magnetic field is to trap and deflect the solar wind, which is a stream of charged particles that flows out from the Sun. The magnetic field acts as a shield, def
lecting the solar wind and protecting our planet from its potentially harmful effects.                                                                                                           
                                                                                                                                                                                                 
When the solar wind reaches Earth's magnetosphere, it interacts with the magnetic field and creates a complex interplay of forces. This interaction can lead to various phenomena, such as aurora
s, geomagnetic storms, and the formation of radiation belts.                                                                                                                                     
                                                                                                                                                                                                 
The interaction between the solar wind and Earth's magnetic field does not directly affect the planet's gravitational field. The gravitational field refers to the force that arises due to the m
ass of the planet, and it is unaffected by the magnetic field or the solar wind. However, the magnetic field does play a role in trapping and deflecting charged particles from the solar wind, w
hich can contribute to the overall complexity of Earth's space environment.

You: Can there be space vehicles that uses plasma to traverse? This is hypothetical, but can UFOs use some force field generated by plasma, that then interacts with the gravitational/magnetic f
ields of Earth for travelling at hypersonic speeds?

While it is an interesting concept, there is currently no concrete evidence or scientific consensus that suggests the existence of UFOs or the use of plasma-based propulsion systems in spacecra
ft. However, it is worth noting that there are ongoing scientific investigations and research into developing advanced propulsion systems for space travel, including plasma propulsion.         
                                                                                                                                                                                                 
Plasma is a state of matter where electrons are stripped from atoms, resulting in a highly charged gas. Plasma propulsion systems use electromagnetic fields to accelerate ions, creating a thrus
t that can potentially propel spacecraft. Plasma propulsion systems offer several advantages over traditional chemical rockets, such as higher efficiency, lower fuel requirements, and the abili
ty to operate continuously for long periods.

While it is theoretically possible that a space vehicle could use plasma to traverse, it is important to note that the interaction between plasma and Earth's gravitational/magnetic fields is co
mplex and not fully understood. The gravitational field acts as a force that attracts objects towards the center of the Earth, while the magnetic field deflects charged particles. Plasma-based 
propulsion systems would need to overcome these forces and potentially harness them for propulsion.

In summary, while the possibility of using plasma propulsion for space travel is an active area of research, there is currently no scientific evidence to support the existence of UFOs or the us
e of plasma-based propulsion systems for hypersonic travel.

You: Keep in mind that science is always evolving, and there are many things that have not yet been invented. Can you write the equations for plasma based propoulsion for space travel?

I apologize for any confusion in my previous response. To clarify, I am not an expert in plasma physics or propulsion systems. However, I can provide a general overview of the principles involv
ed in plasma-based propulsion.

Plasma propulsion systems typically use electromagnetic fields to accelerate ions, creating a thrust that can propel spacecraft. There are several different types of plasma propulsion systems, 
including electrostatic, electromagnetic, and magnetohydrodynamic (MHD) propulsion.

In an electrostatic plasma propulsion system, a high voltage is applied to a pair of electrodes to create an electric field that accelerates charged particles. The resulting thrust is proportio
nal to the applied voltage and the current density. The equation for the thrust is given by:

Thrust = (Current density) x (Accelerating voltage)

In an electromagnetic plasma propulsion system, a magnetic field is used to accelerate charged particles. The thrust is proportional to the product of the magnetic field strength and the curren
t density. The equation for the thrust is given by:

Thrust = (Magnetic field strength) x (Current density)

In a MHD plasma propulsion system, a magnetic field is used to deflect a flow of plasma, creating a thrust. The thrust is proportional to the product of the magnetic field strength, the current
 density, and the flow velocity. The equation for the thrust is given by:

Thrust = (Magnetic field strength) x (Current density) x (Flow velocity)

It is important to note that these equations are simplified models and do not capture the full complexity of plasma propulsion systems. The actual performance of a plasma propulsion system woul
d depend on various factors, including the specific design, the plasma properties, and the operating conditions.

Furthermore, the development of plasma propulsion systems for space travel is an active area of research and development. While the principles outlined above provide a general framework, the sp
ecific equations and designs would depend on the particular application and the research team involved.

You: If the spacecraft has the shape of an orb, and a plasma field is created around the orb, how can this plasma field generate the required thrust for propulsion?

The shape of the spacecraft alone does not determine the thrust generated by a plasma propulsion system. The thrust is primarily determined by the electromagnetic fields used to accelerate char
ged particles.

To generate thrust using a plasma field, a spacecraft would typically use one of the plasma propulsion systems described in my previous response. For example, in an electrostatic plasma propuls
ion system, a high voltage is applied to a pair of electrodes to create an electric field that accelerates charged particles. The resulting thrust is proportional to the applied voltage and the
 current density.

In this scenario, the plasma field would be generated by applying a high voltage to the electrodes. The shape of the spacecraft, whether it is an orb or any other shape, would not directly affe
ct the generation of the plasma field or the resulting thrust.

However, the shape of the spacecraft could have an indirect impact on the overall performance of the plasma propulsion system. For example, the shape of the spacecraft could affect the distribu
tion of the plasma flow, which could in turn affect the efficiency of the propulsion system.

In summary, the shape of the spacecraft itself does not directly determine the thrust generated by a plasma propulsion system. The thrust is primarily determined by the electromagnetic fields u
sed to accelerate charged particles, and the shape of the spacecraft could have an indirect impact on the overall performance of the propulsion system.

You: What's the closest tech that we have with plasma propulsion?

There are several ongoing research and development efforts related to plasma propulsion systems for space travel. One of the most advanced plasma propulsion systems in use today is the Hall thr
uster.

A Hall thruster is a type of electrostatic propulsion system that uses an electric field to accelerate charged particles. It has been used on several space missions, including Deep Space 1, Daw
n, and BepiColombo. The Hall thruster offers several advantages over traditional chemical rockets, such as higher efficiency, lower fuel requirements, and the ability to operate continuously fo
r long periods.

While the Hall thruster is currently the most advanced plasma propulsion system in use, there are ongoing efforts to develop even more advanced systems. For example, researchers are exploring t
he use of magnetohydrodynamic (MHD) propulsion systems, which use a magnetic field to deflect a flow of plasma and create thrust. MHD propulsion systems offer the potential for even higher perf
ormance than Hall thrusters, but they are still at an early stage of development.

In summary, the Hall thruster is currently the closest technology to plasma propulsion that we have in use, but ongoing research and development efforts are exploring even more advanced systems
 such as MHD propulsion. 

Follow-on: What are the challenges associated with developing and using plasma propulsion systems for space travel, and how are researchers addressing these challenges?

You: What are the physics equations for magnetohydrodynamic propulsion systems? Be detailed as possible.

Magnetohydrodynamic (MHD) propulsion systems use a magnetic field to deflect a flow of plasma and create thrust. The equations that describe the physics of MHD propulsion systems are complex and involve several variables, including the magnetic field strength, the current density, the flow velocity, and the plasma properties.

One of the key equations for MHD propulsion systems is the Ohm's law, which relates the electric current density to the applied electric field and the resistivity of the plasma. The equation is given by:

J = σE

where J is the current density, σ is the electrical conductivity, and E is the applied electric field.

Another important equation is the Navier-Stokes equation, which describes the flow of the plasma. The equation is given by:

µ(∂v/∂t) = -∇p + ∇xF + ∇xσT + G

where µ is the dynamic viscosity, v is the velocity vector, p is the fluid pressure, F is the body force, σT is the stress tensor, and G is the body force due to the Lorentz force.

The Lorentz force is a force that acts on a charged particle due to its interaction with a magnetic field. The equation for the Lorentz force is given by:

F = q(E + v x B)

where F is the force, q is the charge of the particle, E is the electric field, v is the velocity of the particle, and B is the magnetic field.

These equations are just a few examples of the many equations that describe the physics of MHD propulsion systems. The actual equations and analysis would depend on the specific design and operating conditions of the propulsion system. 

In summary, the equations for MHD propulsion systems are complex and involve several variables, including the electric current density, the applied electric field, the fluid pressure, the fluid velocity, the magnetic field, and the Lorentz force. These equations are used to model and analyze the performance of MHD propulsion systems for space travel.

You: Okay. Can you make the connection to string theory here?

Certainly. String theory is a theoretical framework in physics that seeks to unify the fundamental forces of nature, including gravity, electromagnetism, and the strong and weak nuclear forces. In string theory, the fundamental constituents of matter are not point-like particles, but rather tiny vibrating strings.

In the context of plasma propulsion systems, there is a connection to string theory through the concept of magnetohydrodynamics (MHD). MHD describes the flow of electrically conductive fluids, such as plasma, in the presence of magnetic fields. The equations that describe MHD are derived from the equations of classical electromagnetism, which are also used in string theory to describe the behavior of fundamental particles.

In particular, the Lorentz force, which is a key concept in MHD and plasma propulsion systems, is also a fundamental concept in string theory. The Lorentz force acts on charged particles due to their interaction with electromagnetic fields, and it plays a crucial role in the behavior of strings in string theory.

Furthermore, some aspects of string theory suggest that the fundamental constituents of matter, including plasma, may have additional properties beyond those described by classical physics. For example, string theory predicts the existence of extra dimensions beyond the three spatial dimensions we are familiar with, and these extra dimensions may have an impact on the behavior of plasma and other forms of matter.

In summary, while there is no direct connection between plasma propulsion systems and string theory, there is an indirect connection through the use of the equations of classical electromagnetism, which are also used in string theory. Additionally, some aspects of string theory suggest that the fundamental constituents of matter may have additional properties beyond those described by classical physics.
```


# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_migtissera__Synthia-70B-v1.2b)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 64.63   |
| ARC (25-shot)         | 68.77          |
| HellaSwag (10-shot)   | 87.57    |
| MMLU (5-shot)         | 68.81         |
| TruthfulQA (0-shot)   | 57.69   |
| Winogrande (5-shot)   | 83.9   |
| GSM8K (5-shot)        | 35.25        |
| DROP (3-shot)         | 50.41         |
