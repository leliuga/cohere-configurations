---
license: mit
language:
- en
pipeline_tag: text-generation
tags:
- nlp
- code
model-index:
  - name: phi-2-dpo
    results:
      - task:
          type: text-generation
        dataset:
          name: AlpacaEval
          type: AlpacaEval
        metrics:
          - name: AlpacaEval
            type: AlpacaEval
            value: 81.37%
        source:
          name: AlpacaEval
          url: https://github.com/tatsu-lab/alpaca_eval
---

## Model Summary

`phi-2-dpo` is an instruction-tuned model from an earlier SFT model [`phi-2-sft`](https://huggingface.co/lxuechen/phi-2-sft). Direct preference optimization (DPO) is used for fine-tuning on a 10k subset of the [UltraFeedback dataset](https://huggingface.co/datasets/HuggingFaceH4/ultrafeedback_binarized).

The purpose of the experiment is to understand the quality of the pre-trained Phi-2 model. The good news is that `phi-2-dpo` can follow open-ended user instructions well.

## Decoding

Format your prompt as 
```
"""### Human: {instruction}

### Assistant:"""
```
where `instruction` is your query.


Here's a full-fledged example:

```
import torch
import transformers

model_name_or_path = "lxuechen/phi-2-dpo"
model: transformers.PreTrainedModel = transformers.AutoModelForCausalLM.from_pretrained(
    model_name_or_path,
    low_cpu_mem_usage=True,
    device_map="auto",
    trust_remote_code=True,
    torch_dtype=torch.float16
)
tokenizer = transformers.AutoTokenizer.from_pretrained(model_name_or_path)

input_text = "### Human: Give me a good recipe for a chinese dish\n\n### Assistant:"

outputs = model.generate(
    tokenizer(input_text, return_tensors="pt").to(model.device)['input_ids'],
    max_length=1024,
    temperature=0.7,
    top_p=0.9,
    do_sample=True,
    pad_token_id=tokenizer.pad_token_id,
    eos_token_id=tokenizer.eos_token_id,
)
print(tokenizer.decode(outputs[0], skip_special_tokens=True))
```

## Training

The model was fine-tuned on a 10k subset of the binarized version of UltraFeedback with DPO.

Hyperparameters:
- learning rate: 3% linear warmup, with a peak of 3e-5 and cosine decay
- epochs: 2
- batch size: 64
- context length: 1024
- DPO beta: 0.1

## Limitations of `phi-2-dpo`

* Generate Inaccurate Code and Facts: The model may produce incorrect code snippets and statements. Users should treat these outputs as suggestions or starting points, not as definitive or accurate solutions.

* Limited Scope for code: Majority of Phi-2 training data is based in Python and use common packages such as "typing, math, random, collections, datetime, itertools". If the model generates Python scripts that utilize other packages or scripts in other languages, we strongly recommend users manually verify all API uses.

* Unreliable Responses to Instruction: The model has not undergone instruction fine-tuning. As a result, it may struggle or fail to adhere to intricate or nuanced instructions provided by users.

* Language Limitations: The model is primarily designed to understand standard English. Informal English, slang, or any other languages might pose challenges to its comprehension, leading to potential misinterpretations or errors in response.

* Potential Societal Biases: Phi-2 is not entirely free from societal biases despite efforts in assuring trainig data safety. There's a possibility it may generate content that mirrors these societal biases, particularly if prompted or instructed to do so. We urge users to be aware of this and to exercise caution and critical thinking when interpreting model outputs.

* Toxicity: Despite being trained with carefully selected data, the model can still produce harmful content if explicitly prompted or instructed to do so. We chose to release the model for research purposes only -- We hope to help the open-source community develop the most effective ways to reduce the toxicity of a model directly after pretraining.

* Verbosity: Phi-2 being a base model often produces irrelevant or extra text and responses following its first answer to user prompts within a single turn. This is due to its training dataset being primarily textbooks, which results in textbook-like responses.