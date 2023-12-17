---
license: apache-2.0
---

# **Meet 10.7B Solar: Elevating Performance with Upstage Depth UP Scaling!**


# **Introduction**

We introduce the first 10.7 billion (B) parameter model, SOLAR-10.7B. It's compact, yet remarkably powerful, and demonstrates unparalleled state-of-the-art performance in models with parameters under 30B.

We developed the Depth Up-Scaling technique. Built on the Llama2 architecture, SOLAR-10.7B incorporates the innovative Upstage Depth Up-Scaling. We then integrated Mistral 7B weights into the upscaled layers, and finally, continued pre-training for the entire model.

Depth-Upscaled SOLAR-10.7B has remarkable performance. It outperforms models with up to 30B parameters, even surpassing the recent Mixtral 8X7B model. For detailed information, please refer to the experimental table.
Solar 10.7B is an ideal choice for fine-tuning. SOLAR-10.7B offers robustness and adaptability for your fine-tuning needs. Our simple instruction fine-tuning using the SOLAR-10.7B pre-trained model yields significant performance improvements ([SOLAR-10.7B-Instruct-v1.0](https://huggingface.co/upstage/SOLAR-10.7B-Instruct-v1.0)).

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

This model is pre-trained and is capable of just generating random text. To use it for chatting, you must fine-tune the model first.

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

tokenizer = AutoTokenizer.from_pretrained("Upstage/SOLAR-10.7B-v1.0")
model = AutoModelForCausalLM.from_pretrained(
    "Upstage/SOLAR-10.7B-v1.0",
    device_map="auto",
    torch_dtype=torch.float16,
)
```

### **Generating Text**

To generate text, use the following Python code:

```python
text = "Hi, my name is "
inputs = tokenizer(text, return_tensors="pt")

outputs = model.generate(**inputs, max_new_tokens=64)
print(tokenizer.decode(outputs[0], skip_special_tokens=True))
```

### **License**
- [upstage/SOLAR-10.7B-v1.0](https://huggingface.co/upstage/SOLAR-10.7B-v1.0): apache-2.0
- [upstage/SOLAR-10.7B-Instruct-v1.0](https://huggingface.co/upstage/SOLAR-10.7B-Instruct-v1.0): cc-by-nc-4.0
  - Since some non-commercial datasets such as Alpaca are used for fine-tuning, we release fine-tuned model as cc-by-nc-4.0.

### **The Upstage AI Team** ###
Upstage is creating the best LLM and DocAI. Please find more information at https://upstage.ai 

### **Contact Us** ###
Any questions and suggestions, please use the discussion tab. If you want to contact us directly, drop an email to [contact@upstage.ai ](mailto:contact@upstage.ai)

