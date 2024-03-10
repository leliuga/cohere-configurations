---
license: apache-2.0
tags:
- LLMs
- mistral
- Intel
pipeline_tag: text-generation
base_model: mistralai/Mistral-7B-v0.1
model-index:
- name: neural-chat-7b-v3-1
  results:
  - task:
      type: Large Language Model
      name: Large Language Model
    dataset:
      type: Open-Orca/SlimOrca
      name: Open-Orca/SlimOrca
    metrics:
    - type: ARC (25-shot)
      value: 66.21
      name: ARC (25-shot)
      verified: true
    - type: HellaSwag (10-shot)
      value: 83.64
      name: HellaSwag (10-shot)
      verified: true
    - type: MMLU (5-shot)
      value: 62.37
      name: MMLU (5-shot)
      verified: true
    - type: TruthfulQA (0-shot)
      value: 59.65
      name: TruthfulQA (0-shot)
      verified: true
    - type: Winogrande (5-shot)
      value: 78.14
      name: Winogrande (5-shot)
      verified: true
    - type: GSM8K (5-shot)
      value: 19.56
      name: GSM8K (5-shot)
      verified: true
    - type: DROP (3-shot)
      value: 43.84
      name: DROP (3-shot)
      verified: true
datasets:
- Open-Orca/SlimOrca
language:
- en
---

## Model Details: Neural-Chat-v3-1

This model is a fine-tuned 7B parameter LLM on the Intel Gaudi 2 processor from the [mistralai/Mistral-7B-v0.1](https://huggingface.co/mistralai/Mistral-7B-v0.1) on the open source dataset [Open-Orca/SlimOrca](https://huggingface.co/datasets/Open-Orca/SlimOrca). The model was aligned using the Direct Performance Optimization (DPO) method with [Intel/orca_dpo_pairs](https://huggingface.co/datasets/Intel/orca_dpo_pairs). For more information, refer to the Medium article [The Practice of Supervised Fine-tuning and Direct Preference Optimization on Intel Gaudi2](https://medium.com/@NeuralCompressor/the-practice-of-supervised-finetuning-and-direct-preference-optimization-on-habana-gaudi2-a1197d8a3cd3).

<p align="center">
  <img src="https://cdn-uploads.huggingface.co/production/uploads/6297f0e30bd2f58c647abb1d/ctASHUT5QYIxMsOFa-sHC.webp" width="500"/>
  Photo by Google DeepMind on Unsplash
</p>

| Model Detail | Description |
| ----------- | ----------- | 
| Model Authors - Company | Intel. The NeuralChat team with members from DCAI/AISE/AIPT. Core team members: Kaokao Lv, Liang Lv, Chang Wang, Wenxin Zhang, Xuhui Ren, and Haihao Shen.| 
| Date | October, 2023 | 
| Version | v3-1 | 
| Type | 7B Large Language Model | 
| Paper or Other Resources | [Medium Blog](https://medium.com/@NeuralCompressor/the-practice-of-supervised-finetuning-and-direct-preference-optimization-on-habana-gaudi2-a1197d8a3cd3) | 
| License | Apache 2.0 |
| Questions or Comments | [Community Tab](https://huggingface.co/Intel/neural-chat-7b-v3-1/discussions) and [Intel DevHub Discord](https://discord.gg/rv2Gp55UJQ)|

| Intended Use | Description |
| ----------- | ----------- | 
| Primary intended uses | You can use the fine-tuned model for several language-related tasks. Checkout the [LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard) to see how this model is doing. | 
| Primary intended users | Anyone doing inference on language-related tasks. | 
| Out-of-scope uses | This model in most cases will need to be fine-tuned for your particular task.  The model should not be used to intentionally create hostile or alienating environments for people.|

## How To Use

Context length for this model: 8192 tokens (same as https://huggingface.co/mistralai/Mistral-7B-v0.1)

### Training hyperparameters

The following hyperparameters were used during training:
- learning_rate: 1e-04
- train_batch_size: 1
- eval_batch_size: 2
- seed: 42
- distributed_type: multi-HPU
- num_devices: 8
- gradient_accumulation_steps: 8
- total_train_batch_size: 64
- total_eval_batch_size: 8
- optimizer: Adam with betas=(0.9,0.999) and epsilon=1e-08
- lr_scheduler_type: cosine
- lr_scheduler_warmup_ratio: 0.03
- num_epochs: 2.0

### Reproduce the model
Here is the sample code to reproduce the model: [GitHub sample code](https://github.com/intel/intel-extension-for-transformers/blob/main/intel_extension_for_transformers/neural_chat/examples/finetuning/finetune_neuralchat_v3). Here is the documentation to reproduce building the model:

```bash
git clone https://github.com/intel/intel-extension-for-transformers.git
cd intel-extension-for-transformers

docker build --no-cache ./ --target hpu --build-arg REPO=https://github.com/intel/intel-extension-for-transformers.git --build-arg ITREX_VER=main -f ./intel_extension_for_transformers/neural_chat/docker/Dockerfile -t chatbot_finetuning:latest

docker run -it --runtime=habana -e HABANA_VISIBLE_DEVICES=all -e OMPI_MCA_btl_vader_single_copy_mechanism=none --cap-add=sys_nice --net=host --ipc=host chatbot_finetuning:latest

# after entering docker container
cd examples/finetuning/finetune_neuralchat_v3

```
We select the latest pretrained mistralai/Mistral-7B-v0.1 and the open source dataset Open-Orca/SlimOrca to conduct the experiment.

The below script use deepspeed zero2 to lanuch the training with 8 cards Gaudi2. In the `finetune_neuralchat_v3.py`, the default `use_habana=True, use_lazy_mode=True, device="hpu"` for Gaudi2. And if you want to run it on NVIDIA GPU, you can set them `use_habana=False, use_lazy_mode=False, device="auto"`.

```python
deepspeed --include localhost:0,1,2,3,4,5,6,7 \
    --master_port 29501 \
    finetune_neuralchat_v3.py
```

Merge the LoRA weights:

```python
python apply_lora.py \
    --base-model-path mistralai/Mistral-7B-v0.1 \
    --lora-model-path finetuned_model/ \
    --output-path finetuned_model_lora
```


### FP32 Inference with Transformers

```python
import transformers

model_name = 'Intel/neural-chat-7b-v3-1'
model = transformers.AutoModelForCausalLM.from_pretrained(model_name)
tokenizer = transformers.AutoTokenizer.from_pretrained(model_name)

def generate_response(system_input, user_input):

    # Format the input using the provided template
    prompt = f"### System:\n{system_input}\n### User:\n{user_input}\n### Assistant:\n"

    # Tokenize and encode the prompt
    inputs = tokenizer.encode(prompt, return_tensors="pt", add_special_tokens=False)

    # Generate a response
    outputs = model.generate(inputs, max_length=1000, num_return_sequences=1)
    response = tokenizer.decode(outputs[0], skip_special_tokens=True)

    # Extract only the assistant's response
    return response.split("### Assistant:\n")[-1]


# Example usage
system_input = "You are a math expert assistant. Your mission is to help users understand and solve various math problems. You should provide step-by-step solutions, explain reasonings and give the correct answer."
user_input = "calculate 100 + 520 + 60"
response = generate_response(system_input, user_input)
print(response)

# expected response
"""
To calculate the sum of 100, 520, and 60, we will follow these steps:

1. Add the first two numbers: 100 + 520
2. Add the result from step 1 to the third number: (100 + 520) + 60

Step 1: Add 100 and 520
100 + 520 = 620

Step 2: Add the result from step 1 to the third number (60)
(620) + 60 = 680

So, the sum of 100, 520, and 60 is 680.
"""
```

### BF16 Inference with Intel Extension for Transformers and Intel Extension for Pytorch
```python
from transformers import AutoTokenizer, TextStreamer
import torch
from intel_extension_for_transformers.transformers import AutoModelForCausalLM
import intel_extension_for_pytorch as ipex

model_name = "Intel/neural-chat-7b-v3-1"
prompt = "Once upon a time, there existed a little girl,"

tokenizer = AutoTokenizer.from_pretrained(model_name, trust_remote_code=True)
inputs = tokenizer(prompt, return_tensors="pt").input_ids
streamer = TextStreamer(tokenizer)

model = AutoModelForCausalLM.from_pretrained(model_name, torch_dtype=torch.bfloat16)
model = ipex.optimize(model.eval(), dtype=torch.bfloat16, inplace=True, level="O1", auto_kernel_selection=True)

outputs = model.generate(inputs, streamer=streamer, max_new_tokens=300)
```


### INT4 Inference with Transformers and Intel Extension for Transformers
```python
from transformers import AutoTokenizer, TextStreamer
from intel_extension_for_transformers.transformers import AutoModelForCausalLM, WeightOnlyQuantConfig
model_name = "Intel/neural-chat-7b-v3-1"

# for int8, should set weight_dtype="int8"       
config = WeightOnlyQuantConfig(compute_dtype="bf16", weight_dtype="int4")
prompt = "Once upon a time, there existed a little girl,"

tokenizer = AutoTokenizer.from_pretrained(model_name, trust_remote_code=True)
inputs = tokenizer(prompt, return_tensors="pt").input_ids
streamer = TextStreamer(tokenizer)

model = AutoModelForCausalLM.from_pretrained(model_name, quantization_config=config)
outputs = model.generate(inputs, streamer=streamer, max_new_tokens=300)

```

| Factors | Description | 
| ----------- | ----------- | 
| Groups | More details about the dataset and annotations can be found at [Open-Orca/SlimOrca](https://huggingface.co/datasets/Open-Orca/SlimOrca) and the associated paper at https://arxiv.org/abs/2306.02707. | 
| Instrumentation | The performance of the model can vary depending on the inputs to the model. In this case, the prompts provided can drastically change the prediction of the language model. |
| Environment | The model was trained on the Intel Gaudi 2 processor (8 cards).  |
| Card Prompts | Model deployment on alternate hardware and software will change model performance. The model evaluation factors are from the Hugging Face LLM leaderboard: ARC, HellaSwag, MMLU, TruthfulQA, Winogrande, GSM8K, and DROP (see Quantitative Analyses below). |

| Metrics | Description | 
| ----------- | ----------- | 
| Model performance measures | The model performance was evaluated against other LLMs according to the measures on the LLM leaderboard. These were selected as this has become the standard for LLM performance. |
| Decision thresholds | No decision thresholds were used. | 
| Approaches to uncertainty and variability | - | 

| Training and Evaluation Data | Description | 
| ----------- | ----------- | 
| Datasets | The training data are from [Open-Orca/SlimOrca](https://huggingface.co/datasets/Open-Orca/SlimOrca). There is no contamination from the GSM8k test set, as this is not a part of the Open-Orca/SlimOrca dataset.|
| Motivation | - |
| Preprocessing | - | 

## Quantitative Analyses 
The model was submitted to the [LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard). The detailed submission can be found here: [https://huggingface.co/datasets/open-llm-leaderboard/details_Intel__neural-chat-7b-v3-1](https://huggingface.co/datasets/open-llm-leaderboard/details_Intel__neural-chat-7b-v3-1). The metrics can be found below and show that the model has significantly improved performance from Mistral-7B-v0.1 and neural-chat-7b-v3.

| Model | Average ⬆️| ARC (25-s) ⬆️ | HellaSwag (10-s) ⬆️ | MMLU (5-s) ⬆️| TruthfulQA (MC) (0-s) ⬆️ | Winogrande (5-s) | GSM8K (5-s) | DROP (3-s) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
|[mistralai/Mistral-7B-v0.1](https://huggingface.co/mistralai/Mistral-7B-v0.1) | 50.32 | 59.58  | 83.31  | 64.16  | 42.15 | 78.37 | 18.12 | 6.14 |
| [Intel/neural-chat-7b-v3](https://huggingface.co/Intel/neural-chat-7b-v3) | **57.31** | 67.15 | 83.29 | 62.26  | 58.77 | 78.06 | 1.21 | 50.43 |
| [Intel/neural-chat-7b-v3-1](https://huggingface.co/Intel/neural-chat-7b-v3-1) | **59.06** | 66.21 | 83.64 | 62.37  | 59.65 | 78.14 | 19.56 | 43.84 |

## Ethical Considerations and Limitations
Neural-chat-7b-v3-1 can produce factually incorrect output, and should not be relied on to produce factually accurate information. Because of the limitations of the pretrained model and the finetuning datasets, it is possible that this model could generate lewd, biased or otherwise offensive outputs.

Therefore, before deploying any applications of neural-chat-7b-v3-1, developers should perform safety testing.

## Caveats and Recommendations

Users (both direct and downstream) should be made aware of the risks, biases and limitations of the model.

Here are a couple of useful links to learn more about Intel's AI software:
* Intel Neural Compressor [link](https://github.com/intel/neural-compressor)
* Intel Extension for Transformers [link](https://github.com/intel/intel-extension-for-transformers)

## Disclaimer

The license on this model does not constitute legal advice. We are not responsible for the actions of third parties who use this model. Please cosult an attorney before using this model for commercial purposes.