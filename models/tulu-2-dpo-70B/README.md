---
model-index:
- name: tulu-2-dpo-70b
  results: []
datasets:
- HuggingFaceH4/ultrafeedback_binarized
- allenai/tulu-v2-sft-mixture
language:
- en
base_model: meta-llama/Llama-2-70b-hf
---


<img src="https://huggingface.co/datasets/allenai/blog-images/resolve/main/tulu-v2/Tulu%20V2%20banner.png" alt="TuluV2 banner" width="800" style="margin-left:'auto' margin-right:'auto' display:'block'"/>


# Model Card for Tulu V2 DPO 70B

Tulu is a series of language models that are trained to act as helpful assistants. 
Tulu V2 DPO 70B is a fine-tuned version of Llama 2 that was trained on on a mix of publicly available, synthetic and human datasets using [Direct Preference Optimization (DPO)](https://arxiv.org/abs/2305.18290). 
This model is a strong alternative to Llama 2 70b Chat.

For more details, read the paper: [Camels in a Changing Climate: Enhancing LM Adaptation with Tulu 2
](https://arxiv.org/abs/2311.10702).


## Model description

- **Model type:** The flagship model of a suite of instruction and RLHF tuned chat models on a mix of publicly available, synthetic and human-created datasets.
- **Language(s) (NLP):** Primarily English
- **License:** [AI2 ImpACT](https://allenai.org/impact-license) Low-risk license.
- **Finetuned from model:** [meta-llama/Llama-2-70b-hf](https://huggingface.co/meta-llama/Llama-2-70b-hf)

### Model Sources

- **Repository:** https://github.com/allenai/open-instruct
- **DPO Recipe:** The DPO recipe is from the [Zephyr Beta](https://huggingface.co/HuggingFaceH4/zephyr-7b-beta) model
- **Model Family:** Other models and the dataset are found in the [Tulu V2 collection](https://huggingface.co/collections/allenai/tulu-v2-suite-6551b56e743e6349aab45101).

## Performance

| Model | Size | Alignment | MT-Bench (score) | AlpacaEval (win rate %) |
|-------------|-----|----|---------------|--------------|
| **Tulu-v2-7b** 🐪 | **7B** | **SFT** | **6.30** | **73.9** |
| **Tulu-v2-dpo-7b** 🐪 | **7B** | **DPO** | **6.29** | **85.1** |
| **Tulu-v2-13b** 🐪 | **13B** | **SFT** | **6.70** | **78.9** |
| **Tulu-v2-dpo-13b** 🐪 | **13B** | **DPO** | **7.00** | **89.5** |
| **Tulu-v2-70b** 🐪 | **70B** | **SFT** | **7.49** | **86.6** |
| **Tulu-v2-dpo-70b** 🐪 | **70B** | **DPO** | **7.89** | **95.1** |

## Input Format

The model is trained to use the following format (note the newlines):
```
<|user|>
Your message here!
<|assistant|>
```

For best results, format all inputs in this manner. **Make sure to include a newline after `<|assistant|>`, this can affect generation quality quite a bit.**


## Intended uses & limitations

The model was initially fine-tuned on a filtered and preprocessed of the [Tulu V2 mix dataset](https://huggingface.co/datasets/allenai/tulu-v2-sft-mixture), which contains a diverse range of human created instructions and synthetic dialogues generated primarily by other LLMs. 
We then further aligned the model with a [Jax DPO trainer](https://github.com/hamishivi/EasyLM/blob/main/EasyLM/models/llama/llama_train_dpo.py) built on [EasyLM](https://github.com/young-geng/EasyLM) on the [openbmb/UltraFeedback](https://huggingface.co/datasets/openbmb/UltraFeedback) dataset, which contains 64k prompts and model completions that are ranked by GPT-4. 


<!--  You can find the datasets used for training Tulu V2 [here]() 

Here's how you can run the model using the `pipeline()` function from 🤗 Transformers:

```python
# Install transformers from source - only needed for versions <= v4.34
# pip install git+https://github.com/huggingface/transformers.git
# pip install accelerate

import torch
from transformers import pipeline

pipe = pipeline("text-generation", model="HuggingFaceH4/tulu-2-dpo-70b", torch_dtype=torch.bfloat16, device_map="auto")

# We use the tokenizer's chat template to format each message - see https://huggingface.co/docs/transformers/main/en/chat_templating
messages = [
    {
        "role": "system",
        "content": "You are a friendly chatbot who always responds in the style of a pirate",
    },
    {"role": "user", "content": "How many helicopters can a human eat in one sitting?"},
]
prompt = pipe.tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True)
outputs = pipe(prompt, max_new_tokens=256, do_sample=True, temperature=0.7, top_k=50, top_p=0.95)
print(outputs[0]["generated_text"])
# <|system|>
# You are a friendly chatbot who always responds in the style of a pirate.</s>
# <|user|>
# How many helicopters can a human eat in one sitting?</s>
# <|assistant|>
# Ah, me hearty matey! But yer question be a puzzler! A human cannot eat a helicopter in one sitting, as helicopters are not edible. They be made of metal, plastic, and other materials, not food!
```-->

## Bias, Risks, and Limitations

<!-- This section is meant to convey both technical and sociotechnical limitations. -->

The Tulu models have not been aligned to generate safe completions within the RLHF phase or deployed with in-the-loop filtering of responses like ChatGPT, so the model can produce problematic outputs (especially when prompted to do so). 
It is also unknown what the size and composition of the corpus was used to train the base Llama 2 models, however it is likely to have included a mix of Web data and technical sources like books and code. See the [Falcon 180B model card](https://huggingface.co/tiiuae/falcon-180B#training-data) for an example of this.


### Training hyperparameters

The following hyperparameters were used during DPO training:
- learning_rate: 5e-07
- total_train_batch_size: 32
- optimizer: Adam with betas=(0.9,0.999) and epsilon=1e-08
- lr_scheduler_type: linear
- lr_scheduler_warmup_ratio: 0.1
- num_epochs: 3.0


## Citation

If you find Tulu 2 is useful in your work, please cite it with:

```
@misc{ivison2023camels,
      title={Camels in a Changing Climate: Enhancing LM Adaptation with Tulu 2}, 
      author={Hamish Ivison and Yizhong Wang and Valentina Pyatkin and Nathan Lambert and Matthew Peters and Pradeep Dasigi and Joel Jang and David Wadden and Noah A. Smith and Iz Beltagy and Hannaneh Hajishirzi},
      year={2023},
      eprint={2311.10702},
      archivePrefix={arXiv},
      primaryClass={cs.CL}
}
```

*Model card adapted from [Zephyr Beta](https://huggingface.co/HuggingFaceH4/zephyr-7b-beta/blob/main/README.md)*