---
license: apache-2.0
---

# Chupacabra 7B v2

<p><img src="https://huggingface.co/perlthoughts/Chupacabra-7B/resolve/main/chupacabra7b%202.png" width=330></p>

### Model Description

This model was made by merging models based on Mistral with the SLERP merge method.

Advantages of SLERP vs averaging weights(common) are as follows:

- Spherical Linear Interpolation (SLERP) - Traditionally, model merging often resorts to weight averaging which, although straightforward, might not always capture the intricate features of the models being merged. The SLERP technique addresses this limitation, producing a blended model with characteristics smoothly interpolated from both parent models, ensuring the resultant model captures the essence of both its parents.

- Smooth Transitions - SLERP ensures smoother transitions between model parameters. This is especially significant when interpolating between high-dimensional vectors.

- Better Preservation of Characteristics - Unlike weight averaging, which might dilute distinct features, SLERP preserves the curvature and characteristics of both models in high-dimensional spaces.

- Nuanced Blending - SLERP takes into account the geometric and rotational properties of the models in the vector space, resulting in a blend that is more reflective of both parent models' characteristics.

List of all models and merging path is coming soon.

## Purpose

Merging the "thick"est model weights from mistral models using amazing training methods like direct preference optimization (DPO), supervised fine tuning (SFT) and reinforced learning.

I have spent countless hours studying the latest research papers, attending conferences, and networking with experts in the field. I experimented with different algorithms, tactics, fine-tuned hyperparameters, optimizers, and optimized code until I achieved the best possible results.

It has not been without challenges. There were skeptics who doubted my abilities and questioned my approach. My approach can be changed, but a closed mind cannot.

I refused to let their negativity bring me down. Instead, I used their doubts as fuel to push myself even harder. I worked tirelessly (vapenation), day and night, until I finally succeeded in merging with the most performant model weights using SOTA training methods like DPO and other advanced techniques described above.

Thank you openchat 3.5 for showing me the way.

```
"Hate it or love it, the underdogs on top." - The Game
```

Here is my contribution.


## Prompt Template

Replace {system} with your system prompt, and {prompt} with your prompt instruction.

```
<|im_start|>system
{system}<|im_end|>
<|im_start|>user
{prompt}<|im_end|>
<|im_start|>assistant
```

### Bug fixes

- Fixed issue with generation and the incorrect model weights. Model weights have been corrected and now generation works again. Reuploading GGUF to the GGUF repository as well as the AWQ versions.

- Fixed issue with tokenizer not stopping correctly and changed prompt template.

- Uploaded new merged model weights.

### More info

- **Developed by:** Ray Hernandez
- **Model type:** Mistral
- **Language(s) (NLP):** English
- **License:** Apache 2.0

### Model Sources [optional]

<!-- Provide the basic links for the model. -->

## Uses

<!-- Address questions around how the model is intended to be used, including the foreseeable users of the model and those affected by the model. -->

### Direct Use

<!-- This section is for the model use without fine-tuning or plugging into a larger ecosystem/app. -->

[More Information Needed]

### Downstream Use [optional]

<!-- This section is for the model use when fine-tuned for a task, or when plugged into a larger ecosystem/app -->

[More Information Needed]

### Out-of-Scope Use

<!-- This section addresses misuse, malicious use, and uses that the model will not work well for. -->

[More Information Needed]

## Bias, Risks, and Limitations

<!-- This section is meant to convey both technical and sociotechnical limitations. -->

[More Information Needed]

### Recommendations

<!-- This section is meant to convey recommendations with respect to the bias, risk, and technical limitations. -->

Users (both direct and downstream) should be made aware of the risks, biases and limitations of the model. More information needed for further recommendations.

## How to Get Started with the Model

Use the code below to get started with the model.

[More Information Needed]

## Training Details

### Training Data

<!-- This should link to a Dataset Card, perhaps with a short stub of information on what the training data is all about as well as documentation related to data pre-processing or additional filtering. -->

[More Information Needed]

### Training Procedure 

<!-- This relates heavily to the Technical Specifications. Content here should link to that section when it is relevant to the training procedure. -->

#### Preprocessing [optional]

[More Information Needed]


#### Training Hyperparameters

- **Training regime:** [More Information Needed] <!--fp32, fp16 mixed precision, bf16 mixed precision, bf16 non-mixed precision, fp16 non-mixed precision, fp8 mixed precision -->

#### Speeds, Sizes, Times [optional]

<!-- This section provides information about throughput, start/end time, checkpoint size if relevant, etc. -->

[More Information Needed]

## Evaluation

<!-- This section describes the evaluation protocols and provides the results. -->

### Testing Data, Factors & Metrics

#### Testing Data

<!-- This should link to a Dataset Card if possible. -->

[More Information Needed]

#### Factors

<!-- These are the things the evaluation is disaggregating by, e.g., subpopulations or domains. -->

[More Information Needed]

#### Metrics

<!-- These are the evaluation metrics being used, ideally with a description of why. -->

[More Information Needed]

### Results

[More Information Needed]

#### Summary


## Model Examination [optional]

<!-- Relevant interpretability work for the model goes here -->

[More Information Needed]

## Technical Specifications [optional]

### Model Architecture and Objective

[More Information Needed]

### Compute Infrastructure

[More Information Needed]

#### Hardware

[More Information Needed]

#### Software

[More Information Needed]

## Citation [optional]

<!-- If there is a paper or blog post introducing the model, the APA and Bibtex information for that should go in this section. -->

**BibTeX:**

[More Information Needed]

**APA:**

[More Information Needed]

## Glossary [optional]

<!-- If relevant, include terms and calculations in this section that can help readers understand the model or model card. -->

[More Information Needed]

## More Information [optional]

[More Information Needed]

## Model Card Authors [optional]

[More Information Needed]

## Model Card Contact

[More Information Needed]