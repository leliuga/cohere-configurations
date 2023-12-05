---
base_model: deepseek-ai/deepseek-coder-1.3b-base
tags:
- generated_from_trainer
model-index:
- name: evolvedSeeker-1_3_v_0_0_1
  results:
  - task:
      type: text-generation
    dataset:
      type: openai_humaneval
      name: HumanEval
    metrics:
    - name: pass@1
      type: pass@1
      value: 68.29
      verified: false
---

# evolvedSeeker-1_3
EvolvedSeeker v0.0.1 (First phase)

This model is a fine-tuned version of [deepseek-ai/deepseek-coder-1.3b-base](https://huggingface.co/deepseek-ai/deepseek-coder-1.3b-base) on 50k instructions for 3 epochs.

I have mostly curated instructions from evolInstruct datasets and some portions of glaive coder.

Around 3k answers were modified via self-instruct.

Collaborate or Consult me - [Twitter](https://twitter.com/4evaBehindSOTA), [Discord](https://discord.gg/ftEM63pzs2)

*Recommended format is ChatML, Alpaca will work but take care of EOT token*

#### Chat Model Inference
Use Gradio inference notebook here that can easily run in free colab - [Gradio Inference Notebook](https://huggingface.co/TokenBender/evolvedSeeker_1_3/blob/main/TokenBender_gradio_evolvedSeeker_inference.ipynb)

```python
from transformers import AutoTokenizer, AutoModelForCausalLM
tokenizer = AutoTokenizer.from_pretrained("TokenBender/evolvedSeeker_1_3", trust_remote_code=True)
model = AutoModelForCausalLM.from_pretrained("TokenBender/evolvedSeeker_1_3", trust_remote_code=True).cuda()
messages=[
    { 'role': 'user', 'content': "write a program to reverse letters in each word in a sentence without reversing order of words in the sentence."}
]
inputs = tokenizer.apply_chat_template(messages, return_tensors="pt").to(model.device)
# 32021 is the id of <|EOT|> token
outputs = model.generate(inputs, max_new_tokens=512, do_sample=False, top_k=50, top_p=0.95, num_return_sequences=1, eos_token_id=32021)
print(tokenizer.decode(outputs[0][len(inputs[0]):], skip_special_tokens=True))
```

## Model description

First model of Project PIC (Partner-in-Crime) in 1.3B range.
Almost all the work is pending right now for this model hence v0.0.1
![image/png](https://cdn-uploads.huggingface.co/production/uploads/6398bf222da24ee95b51c8d8/Fl-pRCsC_lvnuoP734hsJ.png)

## Intended uses & limitations

Superfast Copilot
Run near lossless quantized in 1G RAM.
Useful for code dataset curation and evaluation.

Limitations - This is a smol model, so smol brain, may have crammed a few things.
Reasoning tests may fail beyond a certain point.

## Training procedure
SFT

### Training results
Humaneval Score - 68.29%

samples.jsonl file uploaded from eval bench results recently for transparency of evaluation.

The score on eval bench is 67%

![image/png](https://cdn-uploads.huggingface.co/production/uploads/6398bf222da24ee95b51c8d8/AFp6PxZ9ZP_xti4VWjen3.png)

### Framework versions

- Transformers 4.35.2
- Pytorch 2.0.1
- Datasets 2.15.0
- Tokenizers 0.15.0