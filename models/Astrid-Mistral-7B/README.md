---
language:
- en
library_name: transformers
tags:
- gpt
- llm
- large language model
- PAIX.Cloud
inference: true
thumbnail: >-
  https://static.wixstatic.com/media/bdee4e_8aa5cefc86024bc88f7e20e3e19d9ff3~mv2.png/v1/fill/w_192%2Ch_192%2Clg_1%2Cusm_0.66_1.00_0.01/bdee4e_8aa5cefc86024bc88f7e20e3e19d9ff3~mv2.png
license: apache-2.0
---
# Model Card
## Summary

- Base model: [mistralai/Mistral-7B-Instruct-v0.1](https://huggingface.co/mistralai/Mistral-7B-Instruct-v0.1)

This model, Astrid-7B-Assistant is a Mistral-7B base model for causal language modeling, designed to generate human-like text.  
It's part of our mission to make AI technology accessible to everyone, focusing on personalization, data privacy, and transparent AI governance. 
Trained in English, it's a versatile tool for a variety of applications. 
This model is one of the many models available on our platform, and we currently have a 1B and 7B open-source model.

This model was trained by [PAIX.Cloud](https://www.paix.cloud/).
- Wait list: [Wait List](https://www.paix.cloud/join-waitlist)


## Usage

To use the model with the `transformers` library on a machine with GPUs, first make sure you have the `transformers` library installed.

```bash
pip install transformers==4.34.0
```

Also make sure you are providing your huggingface token to the pipeline if the model is lying in a private repo.
    - Either leave `token=True` in the `pipeline` and login to hugginface_hub by running
        ```python
        import huggingface_hub
        huggingface_hub.login(<ACCES_TOKEN>)
        ```
    - Or directly pass your <ACCES_TOKEN> to `token` in the `pipeline`

```python
from transformers import pipeline

generate_text = pipeline(
    model="PAIXAI/Astrid-Mistral-7B",
    torch_dtype="auto",
    trust_remote_code=True,
    use_fast=True,
    device_map={"": "cuda:0"},
    token=True,
)

res = generate_text(
    "Why is drinking water so healthy?",
    min_new_tokens=2,
    max_new_tokens=256,
    do_sample=False,
    num_beams=1,
    temperature=float(0.3),
    repetition_penalty=float(1.2),
    renormalize_logits=True
)
print(res[0]["generated_text"])
```

You can print a sample prompt after the preprocessing step to see how it is feed to the tokenizer:

```python
print(generate_text.preprocess("Why is drinking water so healthy?")["prompt_text"])
```

```bash
<|prompt|>Why is drinking water so healthy?<|im_end|><|answer|>
```

Alternatively, you can download [h2oai_pipeline.py](h2oai_pipeline.py), store it alongside your notebook, and construct the pipeline yourself from the loaded model and tokenizer. If the model and the tokenizer are fully supported in the `transformers` package, this will allow you to set `trust_remote_code=False`.

```python
from h2oai_pipeline import H2OTextGenerationPipeline
from transformers import AutoModelForCausalLM, AutoTokenizer

tokenizer = AutoTokenizer.from_pretrained(
    "PAIXAI/Astrid-Mistral-7B",
    use_fast=True,
    padding_side="left",
    trust_remote_code=True,
)
model = AutoModelForCausalLM.from_pretrained(
    "PAIXAI/Astrid-Mistral-7B",
    torch_dtype="auto",
    device_map={"": "cuda:0"},
    trust_remote_code=True,
)
generate_text = H2OTextGenerationPipeline(model=model, tokenizer=tokenizer)

res = generate_text(
    "Why is drinking water so healthy?",
    min_new_tokens=2,
    max_new_tokens=256,
    do_sample=False,
    num_beams=1,
    temperature=float(0.3),
    repetition_penalty=float(1.2),
    renormalize_logits=True
)
print(res[0]["generated_text"])
```


You may also construct the pipeline from the loaded model and tokenizer yourself and consider the preprocessing steps:

```python
from transformers import AutoModelForCausalLM, AutoTokenizer

model_name = "PAIXAI/Astrid-Mistral-7B" # either local folder or huggingface model name
# Important: The prompt needs to be in the same format the model was trained with.
# You can find an example prompt in the experiment logs.
prompt = "<|prompt|>How are you?<|im_end|><|answer|>"

tokenizer = AutoTokenizer.from_pretrained(
    model_name,
    use_fast=True,
    trust_remote_code=True,
)
model = AutoModelForCausalLM.from_pretrained(
    model_name,
    torch_dtype="auto",
    device_map={"": "cuda:0"},
    trust_remote_code=True,
)
model.cuda().eval()
inputs = tokenizer(prompt, return_tensors="pt", add_special_tokens=False).to("cuda")

# generate configuration can be modified to your needs
tokens = model.generate(
    input_ids=inputs["input_ids"],
    attention_mask=inputs["attention_mask"],
    min_new_tokens=2,
    max_new_tokens=256,
    do_sample=False,
    num_beams=1,
    temperature=float(0.3),
    repetition_penalty=float(1.2),
    renormalize_logits=True
)[0]

tokens = tokens[inputs["input_ids"].shape[1]:]
answer = tokenizer.decode(tokens, skip_special_tokens=True)
print(answer)
```

## Quantization and sharding

You can load the models using quantization by specifying ```load_in_8bit=True``` or ```load_in_4bit=True```. Also, sharding on multiple GPUs is possible by setting ```device_map=auto```.

## Model Architecture

```
MistralForCausalLM(
  (model): MistralModel(
    (embed_tokens): Embedding(32002, 4096, padding_idx=0)
    (layers): ModuleList(
      (0-31): 32 x MistralDecoderLayer(
        (self_attn): MistralAttention(
          (q_proj): Linear(in_features=4096, out_features=4096, bias=False)
          (k_proj): Linear(in_features=4096, out_features=1024, bias=False)
          (v_proj): Linear(in_features=4096, out_features=1024, bias=False)
          (o_proj): Linear(in_features=4096, out_features=4096, bias=False)
          (rotary_emb): MistralRotaryEmbedding()
        )
        (mlp): MistralMLP(
          (gate_proj): Linear(in_features=4096, out_features=14336, bias=False)
          (up_proj): Linear(in_features=4096, out_features=14336, bias=False)
          (down_proj): Linear(in_features=14336, out_features=4096, bias=False)
          (act_fn): SiLUActivation()
        )
        (input_layernorm): MistralRMSNorm()
        (post_attention_layernorm): MistralRMSNorm()
      )
    )
    (norm): MistralRMSNorm()
  )
  (lm_head): Linear(in_features=4096, out_features=32002, bias=False)
)
```

## Model Configuration

This model was trained using H2O LLM Studio and with the configuration in [cfg.yaml](cfg.yaml). Visit [H2O LLM Studio](https://github.com/h2oai/h2o-llmstudio) to learn how to train your own large language models.


## Disclaimer

Please read this disclaimer carefully before using the large language model provided in this repository. Your use of the model signifies your agreement to the following terms and conditions.

- Biases and Offensiveness: The large language model is trained on a diverse range of internet text data, which may contain biased, racist, offensive, or otherwise inappropriate content. By using this model, you acknowledge and accept that the generated content may sometimes exhibit biases or produce content that is offensive or inappropriate. The developers of this repository do not endorse, support, or promote any such content or viewpoints.
- Limitations: The large language model is an AI-based tool and not a human. It may produce incorrect, nonsensical, or irrelevant responses. It is the user's responsibility to critically evaluate the generated content and use it at their discretion.
- Use at Your Own Risk: Users of this large language model must assume full responsibility for any consequences that may arise from their use of the tool. The developers and contributors of this repository shall not be held liable for any damages, losses, or harm resulting from the use or misuse of the provided model.
- Ethical Considerations: Users are encouraged to use the large language model responsibly and ethically. By using this model, you agree not to use it for purposes that promote hate speech, discrimination, harassment, or any form of illegal or harmful activities.
- Reporting Issues: If you encounter any biased, offensive, or otherwise inappropriate content generated by the large language model, please report it to the repository maintainers through the provided channels. Your feedback will help improve the model and mitigate potential issues.
- Changes to this Disclaimer: The developers of this repository reserve the right to modify or update this disclaimer at any time without prior notice. It is the user's responsibility to periodically review the disclaimer to stay informed about any changes.

By using the large language model provided in this repository, you agree to accept and comply with the terms and conditions outlined in this disclaimer. If you do not agree with any part of this disclaimer, you should refrain from using the model and any content generated by it.