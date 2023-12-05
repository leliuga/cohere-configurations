---
license: mit
datasets:
- TIGER-Lab/MathInstruct
language:
- en
---
# 🦣 MAmmoTH: Building Math Generalist Models through Hybrid Instruction Tuning

Project Page: [https://tiger-ai-lab.github.io/MAmmoTH/](https://tiger-ai-lab.github.io/MAmmoTH/)

Paper: [https://arxiv.org/pdf/2309.05653.pdf](https://arxiv.org/pdf/2309.05653.pdf)

Code: [https://github.com/TIGER-AI-Lab/MAmmoTH](https://github.com/TIGER-AI-Lab/MAmmoTH)


## Introduction
We introduce 🦣 MAmmoTH, a series of open-source large language models (LLMs) specifically tailored for general math problem-solving. The MAmmoTH models are trained on 🤗 [MathInstruct Dataset](https://huggingface.co/datasets/TIGER-Lab/MathInstruct), a meticulously curated instruction tuning dataset that is lightweight yet generalizable. MathInstruct is compiled from 13 math rationale datasets, six of which are newly curated by this work. It uniquely focuses on the hybrid use of chain-of-thought (CoT) and program-of-thought (PoT) rationales, and ensures extensive coverage of diverse mathematical fields. 

|     | **Base Model: Llama-2**                                       | **Base Model: Code Llama**                                               |
|-----|---------------------------------------------------------------|--------------------------------------------------------------------------|
| 7B  | 🦣 [MAmmoTH-7B](https://huggingface.co/TIGER-Lab/MAmmoTH-7B)   | 🦣 [MAmmoTH-Coder-7B](https://huggingface.co/TIGER-Lab/MAmmoTH-Coder-7B)  |
| 13B | 🦣 [MAmmoTH-13B](https://huggingface.co/TIGER-Lab/MAmmoTH-13B) | 🦣 [MAmmoTH-Coder-13B](https://huggingface.co/TIGER-Lab/MAmmoTH-Coder-13B)|
| 34B | -                                                             | 🦣 [MAmmoTH-Coder-34B](https://huggingface.co/TIGER-Lab/MAmmoTH-Coder-34B)|
| 70B | 🦣 [MAmmoTH-70B](https://huggingface.co/TIGER-Lab/MAmmoTH-70B) | -                                                                        |
                                                                      |


## Training Data
The models are trained on the 🤗 [MathInstruct Dataset](https://huggingface.co/datasets/TIGER-Lab/MathInstruct), which is compiled from 13 different math rationale datasets. Check out the dataset card for more details.


## Training Procedure
The models are fine-tuned with the MathInstruct dataset using the original Llama-2 and Code Llama models as base models. The training procedure varies for different models based on their sizes. Check out our paper for more details.

## Evaluation
The models are evaluated using open-ended and multiple-choice math problems from several datasets. Here are the results:


| **Model**             	| **Decoding** 	| **GSM**  	| **MATH** 	| **AQuA** 	| **NumG** 	| **SVA**  	| **Mat**  	| **Sim**  	| **SAT**  	| **MMLU** 	| **AVG**  	|
|-----------------------|--------------|----------|----------|----------|----------|----------|----------|----------|----------|----------|----------|
| **MAmmoTH-7B**        	| CoT          	| 50.5     	| 10.4     	| 43.7     	| 44.0     	| 47.3     	| 9.2      	| 18.9     	| 32.7     	| 39.9     	| 33.0     	|
|                       	| PoT          	| 51.6     	| 28.7     	| 43.3     	| 52.3     	| 65.1     	| 41.9     	| 48.2     	| 39.1     	| 44.6     	| 46.1     	|
|                       	| **Hybrid**   	| **53.6** 	| **31.5** 	| **44.5** 	| **61.2** 	| **67.7** 	| **46.3** 	| **41.2** 	| **42.7** 	| **42.6** 	| **47.9** 	|
| **MAmmoTH-Coder-7B**  	| CoT          	| 22.4     	| 7.9      	| 36.2     	| 36.0     	| 37.0     	| 8.2      	| 7.2      	| 32.7     	| 34.6     	| 24.7     	|
|                       	| PoT          	| 58.8     	| 32.1     	| 47.2     	| 57.1     	| 71.1     	| 53.9     	| 44.6     	| 40.0     	| 47.8     	| 50.3     	|
|                       	| **Hybrid**   	| **59.4** 	| **33.4** 	| **47.2** 	| **66.4** 	| **71.4** 	| **55.4** 	| **45.9** 	| **40.5** 	| **48.3** 	| **52.0** 	|
| **MAmmoTH-13B**       	| CoT          	| 56.3     	| 12.9     	| 45.3     	| 45.6     	| 53.8     	| 11.7     	| 22.4     	| 43.6     	| 42.3     	| 37.1     	|
|                       	| PoT          	| 61.3     	| 32.6     	| 48.8     	| 59.6     	| 72.2     	| 48.5     	| 40.3     	| 46.8     	| 45.4     	| 50.6     	|
|                       	| **Hybrid**   	| **62.0** 	| **34.2** 	| **51.6** 	| **68.7** 	| **72.4** 	| **49.2** 	| **43.2** 	| **46.8** 	| **47.6** 	| **52.9** 	|
| **MAmmoTH-Coder-13B** 	| CoT          	| 32.1     	| 10.2     	| 40.6     	| 36.2     	| 43.0     	| 9.6      	| 10.1     	| 40.9     	| 36.6     	| 28.8     	|
|                       	| PoT          	| 64.3     	| 35.2     	| 46.8     	| 54.2     	| 73.2     	| 60.0     	| 44.2     	| 48.2     	| 48.2     	| 52.7     	|
|                       	| **Hybrid**   	| **64.7** 	| **36.3** 	| **46.9** 	| **66.8** 	| **73.7** 	| **61.5** 	| **47.1** 	| **48.6** 	| **48.3** 	| **54.9** 	|
| **MAmmoTH-Coder-33B** 	| CoT          	| 34.3     	| 11.6     	| 39.0     	| 36.2     	| 44.6     	| 10.8     	| 10.9     	| 46.4     	| 42.9     	| 30.7     	|
|                       	| PoT          	| 72.3     	| 42.8     	| 53.8     	| 59.6     	| 84.0     	| 64.7     	| 50.6     	| 58.6     	| 52.7     	| 59.9     	|
|                       	| **Hybrid**   	| **72.7** 	| **43.6** 	| **54.7** 	| **71.6** 	| **84.3** 	| **65.4** 	| **51.8** 	| **60.9** 	| **53.8** 	| **62.1** 	|
| **MAmmoTH-70B**       	| CoT          	| 72.4     	| 21.1     	| 57.9     	| 58.9     	| 71.6     	| 20.0     	| 31.9     	| 57.3     	| 52.1     	| 49.2     	|
|                       	| PoT          	| 76.7     	| 40.1     	| 60.2     	| 64.3     	| 81.7     	| 55.3     	| 45.3     	| 64.1     	| 53.5     	| 60.1     	|
|                       	| **Hybrid**   	| **76.9** 	| **41.8** 	| **65.0** 	| **74.4** 	| **82.4** 	| **55.6** 	| **51.4** 	| **66.4** 	| **56.7** 	| **63.4** 	|

## Usage
You can use the models through Huggingface's Transformers library. Use the pipeline function to create a text-generation pipeline with the model of your choice, then feed in a math problem to get the solution.
Check our Github repo for more advanced use: [https://github.com/TIGER-AI-Lab/MAmmoTH](https://github.com/TIGER-AI-Lab/MAmmoTH)


## Prompt Format
If you want to do CoT:
```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{instruction}

### Response:
```

If you want to do PoT:
```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{instruction} Let's write a program.

### Response:
```

## Intended Uses
These models are trained for research purposes. They are designed to solve general math problems. They can be used in educational software, tutoring systems, or any application where a solution to a math problem is needed. The models can generate both a chain of thought (CoT) rationale and a program of thought (PoT) rationale, providing a comprehensive solution to a given math problem.

## Limitations
We've tried our best to build math generalist models. However, we acknowledge that the models' performance may vary based on the complexity and specifics of the math problem. Still not all mathematical fields can be covered comprehensively.


## Citation
If you use the models, data, or code from this project, please cite the original paper:

```
@article{yue2023mammoth,
  title={MAmmoTH: Building Math Generalist Models through Hybrid Instruction Tuning},
  author={Xiang Yue, Xingwei Qu, Ge Zhang, Yao Fu, Wenhao Huang, Huan Sun, Yu Su, Wenhu Chen},
  journal={arXiv preprint arXiv:2309.05653},
  year={2023}
}
```