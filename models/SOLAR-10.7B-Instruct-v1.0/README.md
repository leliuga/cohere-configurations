---
datasets:
- c-s-ale/alpaca-gpt4-data
- Open-Orca/OpenOrca
- Intel/orca_dpo_pairs
- allenai/ultrafeedback_binarized_cleaned
language:
- en
license: cc-by-nc-4.0
---

# **Meet 10.7B Solar: Elevating Performance with Upstage Depth UP Scaling!**

**(This model is [upstage/SOLAR-10.7B-v1.0](https://huggingface.co/upstage/SOLAR-10.7B-v1.0) fine-tuned version for single-turn conversation.)**


# **Introduction**

We introduce the first 10.7 billion (B) parameter model, SOLAR-10.7B. It's compact, yet remarkably powerful, and demonstrates unparalleled state-of-the-art performance in models with parameters under 30B.

We developed the Depth Up-Scaling technique. Built on the Llama2 architecture, SOLAR-10.7B incorporates the innovative Upstage Depth Up-Scaling. We then integrated Mistral 7B weights into the upscaled layers, and finally, continued pre-training for the entire model.

Depth-Upscaled SOLAR-10.7B has remarkable performance. It outperforms models with up to 30B parameters, even surpassing the recent Mixtral 8X7B model. For detailed information, please refer to the experimental table.
Solar 10.7B is an ideal choice for fine-tuning. SOLAR-10.7B offers robustness and adaptability for your fine-tuning needs. Our simple instruction fine-tuning using the SOLAR-10.7B pre-trained model yields significant performance improvements.

# **Instruction Fine-Tuning Strategy**

We utilize state-of-the-art instruction fine-tuning methods including supervised fine-tuning (SFT) and direct preference optimization (DPO) [1].

We used a mixture of the following datasets
- c-s-ale/alpaca-gpt4-data (SFT)
- Open-Orca/OpenOrca (SFT)
- in-house generated data utilizing Metamath [2] (SFT, DPO)
- Intel/orca_dpo_pairs (DPO)
- allenai/ultrafeedback_binarized_cleaned (DPO)

where we were careful of data contamination by not using GSM8K samples when generating data and filtering tasks when applicable via the following list.
```python
filtering_task_list = [
    'task228_arc_answer_generation_easy',
    'ai2_arc/ARC-Challenge:1.0.0',
    'ai2_arc/ARC-Easy:1.0.0',
    'task229_arc_answer_generation_hard',
    'hellaswag:1.1.0', 
    'task1389_hellaswag_completion',
    'cot_gsm8k',
    'cot_gsm8k_ii',
    'drop:2.0.0',
    'winogrande:1.1.0'
]
```

Using the datasets mentioned above, we applied SFT and iterative DPO training, a proprietary alignment strategy, to maximize the performance of our resulting model.

[1] Rafailov, R., Sharma, A., Mitchell, E., Ermon, S., Manning, C.D. and Finn, C., 2023. Direct preference optimization: Your language model is secretly a reward model. NeurIPS.

[2] Yu, L., Jiang, W., Shi, H., Yu, J., Liu, Z., Zhang, Y., Kwok, J.T., Li, Z., Weller, A. and Liu, W., 2023. Metamath: Bootstrap your own mathematical questions for large language models. arXiv preprint arXiv:2309.12284.

# **Evaluation Results**

| Model                                  | H6    | Model Size |
|----------------------------------------|-------|------------|
| **SOLAR-10.7B-Instruct-v1.0**              | **74.20** | **~ 11B**      |
| mistralai/Mixtral-8x7B-Instruct-v0.1   | 72.62 | ~ 46.7B    |
| 01-ai/Yi-34B-200K                      | 70.81 | ~ 34B      |
| 01-ai/Yi-34B                           | 69.42 | ~ 34B      |
| mistralai/Mixtral-8x7B-v0.1            | 68.42 | ~ 46.7B    |
| meta-llama/Llama-2-70b-hf              | 67.87 | ~ 70B      |
| tiiuae/falcon-180B                     | 67.85 | ~ 180B     |
| **SOLAR-10.7B-v1.0**                   | **66.04** | **~11B**   |
| mistralai/Mistral-7B-Instruct-v0.2     | 65.71 | ~ 7B       |
| Qwen/Qwen-14B                          | 65.86 | ~ 14B      |
| 01-ai/Yi-34B-Chat                      | 65.32 | ~34B       |
| meta-llama/Llama-2-70b-chat-hf         | 62.4  | ~ 70B      |
| mistralai/Mistral-7B-v0.1              | 60.97 | ~ 7B       |
| mistralai/Mistral-7B-Instruct-v0.1     | 54.96 | ~ 7B       |

# **Usage Instructions**

This model has been fine-tuned primarily for single-turn conversation, making it less suitable for multi-turn conversations such as chat.

### **Version**

Make sure you have the correct version of the transformers library installed:

```sh
pip install transformers==4.35.2
```

### **Loading the Model**

Use the following Python code to load the model:

```python
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer

tokenizer = AutoTokenizer.from_pretrained("Upstage/SOLAR-10.7B-Instruct-v1.0")
model = AutoModelForCausalLM.from_pretrained(
    "Upstage/SOLAR-10.7B-Instruct-v1.0",
    device_map="auto",
    torch_dtype=torch.float16,
)
```

### **Conducting Single-Turn Conversation**

```python
conversation = [ {'role': 'user', 'content': 'Hello?'} ] 

prompt = tokenizer.apply_chat_template(conversation, tokenize=False, add_generation_prompt=True)

inputs = tokenizer(prompt, return_tensors="pt").to(model.device) 
outputs = model.generate(**inputs, use_cache=True, max_length=4096)
output_text = tokenizer.decode(outputs[0]) 
print(output_text)
```

Below is an example of the output.
```
<s> ### User:
Hello?

### Assistant:
Hello, how can I assist you today? Please feel free to ask any questions or request help with a specific task.</s>
```

### **License**
- [upstage/SOLAR-10.7B-v1.0](https://huggingface.co/upstage/SOLAR-10.7B-v1.0): apache-2.0
- [upstage/SOLAR-10.7B-Instruct-v1.0](https://huggingface.co/upstage/SOLAR-10.7B-Instruct-v1.0): cc-by-nc-4.0
  - Since some non-commercial datasets such as Alpaca are used for fine-tuning, we release this model as cc-by-nc-4.0.

### **The Upstage AI Team** ###
Upstage is creating the best LLM and DocAI. Please find more information at https://upstage.ai 

### **Contact Us** ###
Any questions and suggestions, please use the discussion tab. If you want to contact us directly, drop an email to [contact@upstage.ai](mailto:contact@upstage.ai)