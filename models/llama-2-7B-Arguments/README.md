This model aims to detect and analyze casual arguments.

Model template:
```
<s>[INST] {prompt}
 [/INST]
```

Example:
```
`<`s`>`[INST] Analize the following argument, identifying premises, conclusion, type of argument, and argument validity: 
If officer smith found a broken window at the crime scene then the arson occurred on elm street, and officer smith found a broken window at the crime scene, hence the arson occurred on elm street.
 [/INST] Premise 1: If officer smith found a broken window at the crime scene then the arson occurred on elm street Premise 2: Officer smith found a broken window at the crime scene Conclusion: The arson occurred on Elm Street Type of argument: modus ponen Validity: True `<`/s`>`
```

It was trained on my dataset cris177/Arguments (https://huggingface.co/datasets/cris177/Arguments)