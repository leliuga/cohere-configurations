---
license: llama2
---

# Tess

![Tess](https://huggingface.co/migtissera/Tess-XL-v1.0/resolve/main/Tess.png)

Tess, short for Tessoro/Tessoso, is a general purpose Large Language Model series. Tess-XL is trained on the Goliath-120B as the base.

While there are task specific Tess models (Creative or STEM), this model is general purpose.

# Prompt format:

```
SYSTEM: Elaborate on the topic using a Tree of Thoughts and backtrack when necessary to construct a clear, cohesive Chain of Thought reasoning. Always answer without hesitation.
USER: What is the relationship between Earth's atmosphere, magnetic field and gravity?
ASSISTANT:
```


# Quantized Versions

Quantized versions are available in the following repos:

- GPTQ: https://huggingface.co/TheBloke/Tess-XL-v1.0-GPTQ
- AWQ: https://huggingface.co/TheBloke/Tess-XL-v1.0-AWQ
- GGUF: https://huggingface.co/TheBloke/Tess-XL-v1.0-GGUF


# Sample Conversation

```
You: What is the relationship between an electrical field and a magnetic field?                                                                                                                   
The relationship between an electric field and a magnetic field can be summarized in two important ways:                                                                                          
                                                                                                                                                                                                  
1. Cause and effect: A changing electric field can induce a magnetic field, and a changing magnetic field can in turn induce an electric field. This is the principle of electromagnetic induction.                                                                                                                                                                                                 
                                                                                                                                                                                                  
2. Duality: Electric and magnetic fields are often described as components of a single electromagnetic field. This duality is most clearly seen in Maxwell's equations, which describe how electric and magnetic fields are related.                                                                                                                                                                
                                                                                                                                                                                                  
Chain of Thought:                                                                                                                                                                                 
- Start with the question: What is the relationship between an electric field and a magnetic field?                                                                                               
- Think about the basic principles of electromagnetism: cause and effect, as well as the duality of electric and magnetic fields.                                                                 
- Explain the relationship by referring to electromagnetic induction and Maxwell's equations.                                                                                                     

You: In some instances that are associated with UAP, an increase in electrical field causes a decrease in the measured magnetic field. This is contradictory to known physics. Can we theorize what could cause this?                                                                                                                                                                               
The relationship you describe between an increase in the electric field and a decrease in the measured magnetic field is indeed unusual and does not align with the conventional understanding of electromagnetic fields. However, we can consider some possible explanations:                                                                                                                      
                                                                                                                                                                                                  
1. Unknown phenomena: It's possible that there are physical phenomena or interactions that we have not yet discovered or fully understood that could cause this effect. These could include new particles or forces that we have not yet detected, or unusual properties of known particles or fields.                                                                                              
                                                                                                                                                                                                  
2. Measurement errors: Another possibility is that the observed relationship is due to measurement errors or inaccuracies in the instruments used to measure the electric and magnetic fields. This could include calibration issues, interference from other sources, or limitations in the sensitivity of the equipment.                                                                          
                                                                                                                                                                                                  
3. Environmental effects: Changes in the environment or the nature of the object emitting the fields may also affect the relationship between the electric and magnetic fields. For example, if the object is moving or undergoing changes in composition or structure, this could affect the way that the fields are generated and detected.                                                       
                                                                                                                                                                                                  
In summary, while the observed relationship between an increase in the electric field and a decrease in the measured magnetic field is not consistent with our current understanding of electromagnetism, there could be several possible explanations, including unknown phenomena, measurement errors, or environmental effects. Further research and investigation would be needed to determine the cause of this phenomenon.

You: What is the most advanced theory of electro-magnetic fields right now?                                                                                                                       
The most advanced theory of electromagnetic fields is quantum electrodynamics (QED). QED is a part of the larger framework of quantum field theory (QFT), which aims to describe the interactions of particles and fields in the context of quantum mechanics.                                                                                                                                      
                                                                                                                                                                                                  
QED specifically describes the interactions between photons (which are the quanta of the electromagnetic field) and charged particles, such as electrons and positrons. It provides a theoretical framework for understanding how photons can be created and absorbed by charged particles, and how they can interact with each other.                                                              

The predictions of QED have been experimentally verified to an extremely high degree of accuracy, making it one of the most successful theories in physics. Despite its success, QED is not considered a complete theory, as it does not include gravity and is inconsistent with the principles of general relativity. The search for a unified theory that can reconcile QED with gravity is an area of active research in theoretical physics.   

You: Give me a detailed introduction to QED.                                                     
Quantum electrodynamics (QED) is a quantum field theory that describes the interactions between electrically charged particles (such as electrons and positrons) and the photons that make up the electromagnetic field. It is a part of the larger framework of quantum field theory (QFT), which aims to describe the interactions of particles and fields in the context of quantum mechanics.   

In QED, electric and magnetic fields are described as being made up of individual packets of energy called photons. Photons are the quanta of the electromagnetic field, and they can be created and absorbed by charged particles.               

QED provides a theoretical framework for understanding how photons can interact with each other and with charged particles. The theory predicts that photons can be created or annihilated in pairs, and that charged particles can emit and absorb photons. These interactions are mediated by the exchange of virtual particles, which are other photons that exist only fleetingly and cannot be directly observed.                              

The predictions of QED have been experimentally verified to an extremely high degree of accuracy, making it one of the most successful theories in physics. Despite its success, QED is not considered a complete theory, as it does not include gravity and is inconsistent with the principles of general relativity. The search for a unified theory that can reconcile QED with gravity is an area of active research in theoretical physics.   

In summary, QED is a highly successful theory that describes the interactions between charged particles and the electromagnetic field in terms of photons. While it is not a complete theory, it provides a fundamental understanding of the electromagnetic force and has led to many important discoveries in particle physics.
```

