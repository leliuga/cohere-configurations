---
extra_gated_heading: Access Med42 on Hugging Face
extra_gated_description: >-
  This is a form to enable access to Med42 on Hugging Face. Please read
  the [Med42 License](https://huggingface.co/spaces/m42-health/License) and accept our
  license terms and acceptable use policy before submitting this form. Requests
  will be processed by the M42 Team within 2 working days.
extra_gated_button_content: Submit 
extra_gated_fields:
  Full name: text
  Country: text
  Affiliation: text
  I certify the details provided above are correct and that I have read and agreed to the Med42 License agreement: checkbox
language:
- en
pipeline_tag: text-generation
inference: false
license: other
license_name: med42
tags:
- m42
- health
- healthcare
- clinical-llm
---
# **Med42 - Clinical Large Language Model**
Med42 is an open-access clinical large language model (LLM) developed by M42 to expand access to medical knowledge. Built off LLaMA-2 and comprising 70 billion parameters, this generative AI system provides high-quality answers to medical questions.

## Model Details
*Note: Use of this model is governed by the M42 Health license. In order to download the model weights (and tokenizer), please read the [Med42 License](https://huggingface.co/spaces/m42-health/License) and accept our License by requesting access here.*

Beginning with the base LLaMa-2 model, Med42 was instruction-tuned on a dataset of ~250M tokens compiled from different open-access sources, including medical flashcards, exam questions, and open-domain dialogues.

**Model Developers:** M42 Health AI Team

**Finetuned from model:** Llama-2 - 70B

**Context length:** 4k tokens

**Input:** Text only data

**Output:** Model generates text only

**Status:** This is a static model trained on an offline dataset. Future versions of the tuned models will be released as we enhance model's performance.

**License:** A custom license is available [here](https://huggingface.co/spaces/m42-health/License)

**Research Paper:** TBA

## Intended Use
Med42 is being made available for further testing and assessment as an AI assistant to enhance clinical decision-making and enhance access to an LLM for healthcare use. Potential use cases include:
- Medical question answering
- Patient record summarization
- Aiding medical diagnosis
- General health Q&A

To get the expected features and performance for the model, a specific formatting needs to be followed, including the `<|system|>`, `<|prompter|>` and `<|assistant|>` tags. 

```python
from transformers import AutoModelForCausalLM, AutoTokenizer

model_name_or_path = "m42-health/med42-70b"

model = AutoModelForCausalLM.from_pretrained(model_name_or_path,
                                             device_map="auto")

tokenizer = AutoTokenizer.from_pretrained(model_name_or_path)

prompt = "What are the symptoms of diabetes ?"
prompt_template=f'''
<|system|>: You are a helpful medical assistant created by M42 Health in the UAE.
<|prompter|>:{prompt}
<|assistant|>:
'''

input_ids = tokenizer(prompt_template, return_tensors='pt').input_ids.cuda()
output = model.generate(inputs=input_ids, temperature=0.7, do_sample=True,eos_token_id=tokenizer.eos_token_id, pad_token_id=tokenizer.pad_token_id, max_new_tokens=512)
print(tokenizer.decode(output[0]))
```

## Hardware and Software

The training process was performed on the Condor Galaxy 1 (CG-1) supercomputer platform.


## Evaluation Results

Med42 achieves achieves competitive performance on various medical benchmarks, including MedQA, MedMCQA, PubMedQA, HeadQA, and Measuring Massive Multitask Language Understanding (MMLU) clinical topics. For all evaluations reported so far, we use [EleutherAI's evaluation harness library](https://github.com/EleutherAI/lm-evaluation-harness) and report zero-shot accuracies (except otherwise stated). We compare the performance with that reported for other models (ClinicalCamel-70B, GPT-3.5, GPT-4.0, Med-PaLM 2).

|Dataset|Med42|ClinicalCamel-70B|GPT-3.5|GPT-4.0|Med-PaLM-2 (5-shot)*|
|---|---|---|---|---|---|
|MMLU Clinical Knowledge|74.3|69.8|69.8|86.0|88.3|
|MMLU College Biology|84.0|79.2|72.2|95.1|94.4|
|MMLU College Medicine|68.8|67.0|61.3|76.9|80.9|
|MMLU Medical Genetics|86.0|69.0|70.0|91.0|90.0|
|MMLU Professional Medicine|79.8|71.3|70.2|93.0|95.2|
|MMLU Anatomy|67.4|62.2|56.3|80.0|77.8|
|MedMCQA|60.9|47.0|50.1|69.5|71.3|
|MedQA|61.5|53.4|50.8|78.9|79.7|
|USMLE Self-Assessment|71.7|-|49.1|83.8|-|
|USMLE Sample Exam|72.0|54.3|56.9|84.3|-|

**We note that 0-shot performance is not reported for Med-PaLM 2. Further details can be found at [https://github.com/m42health/med42](https://github.com/m42health/med42)*.


### Key performance metrics:
- Med42 achieves a 72% accuracy on the US Medical Licensing Examination (USMLE) sample exam, surpassing the prior state of the art among openly available medical LLMs.
- 61.5% on MedQA dataset (compared to 50.8% for GPT-3.5)
- Consistently higher performance on MMLU clinical topics compared to GPT-3.5.

## Limitations & Safe Use
- Med42 is not ready for real clinical use. Extensive human evaluation is undergoing as it is required to ensure safety.
- Potential for generating incorrect or harmful information.
- Risk of perpetuating biases in training data.

Use this model responsibly! Do not rely on it for medical usage without rigorous safety testing.

## Accessing Med42 and Reporting Issues

Please report any software "bug" or other problems through one of the following means:

- Reporting issues with the model: [https://github.com/m42health/med42](https://github.com/m42health/med42)
- Reporting risky content generated by the model, bugs and/or any security concerns: [https://forms.office.com/r/YMJu3kcKat](https://forms.office.com/r/YMJu3kcKat)
- M42’s privacy policy available at [https://m42.ae/privacy-policy/](https://m42.ae/privacy-policy/)
- Reporting violations of the Acceptable Use Policy or unlicensed uses of Med42: <med42@m42.ae>

## Citation

*Paper coming soon* 😊  
In the meantime, you can use the following information to cite: 

```
@article{med42,
  title={Med42 - A Clinical Large Language Model},
  author={Christophe, Cl\'ement and Gupta, Avani and Hayat, Nasir and Kanithi, Praveen and Al-Mahrooqi, Ahmed and Munjal, Prateek and Pimentel, Marco and Raha, Tathagata and Rajan, Ronnie and Khan, Shadab},
  year={2023}
}
```