---
language:
- ko

library_name: transformers
pipeline_tag: text-generation
license: cc-by-nc-4.0
---

# **Synatra-7B-v0.3-base🐧**  
![Synatra-7B-Instruct-v0.3](./Synatra.png)

## Support Me
시나트라는 개인 프로젝트로, 1인의 자원으로 개발되고 있습니다. 모델이 마음에 드셨다면 약간의 연구비 지원은 어떨까요?
[<img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy me a Coffee" width="217" height="50">](https://www.buymeacoffee.com/mwell)

Wanna be a sponser? Contact me on Telegram **AlzarTakkarsen**

# **License**

This model is strictly [*non-commercial*](https://creativecommons.org/licenses/by-nc/4.0/) (**cc-by-nc-4.0**) use only.
The "Model" is completely free (ie. base model, derivates, merges/mixes) to use for non-commercial purposes as long as the the included **cc-by-nc-4.0** license in any parent repository, and the non-commercial use statute remains, regardless of other models' licences. 
The licence can be changed after new model released. If you are to use this model for commercial purpose, Contact me.

# **Model Details**
**Base Model**  
[mistralai/Mistral-7B-Instruct-v0.1](https://huggingface.co/mistralai/Mistral-7B-Instruct-v0.1)    

**Trained On**  
A6000 48GB * 8

**Instruction format**

It follows [ChatML](https://github.com/openai/openai-python/blob/main/chatml.md) format and **Alpaca(No-Input)** format.

**TODO**

- ~~``RP 기반 튜닝 모델 제작``~~ ✅
- ~~``데이터셋 정제``~~ ✅
- 언어 이해능력 개선
- ~~``상식 보완``~~ ✅
- 토크나이저 변경


# **Model Benchmark**

## Ko-LLM-Leaderboard

On Benchmarking...

# **Implementation Code**

Since, chat_template already contains insturction format above.
You can use the code below.

```python
from transformers import AutoModelForCausalLM, AutoTokenizer

device = "cuda" # the device to load the model onto

model = AutoModelForCausalLM.from_pretrained("maywell/Synatra-7B-v0.3-base")
tokenizer = AutoTokenizer.from_pretrained("maywell/Synatra-7B-v0.3-base")

messages = [
    {"role": "user", "content": "바나나는 원래 하얀색이야?"},
]

encodeds = tokenizer.apply_chat_template(messages, return_tensors="pt")

model_inputs = encodeds.to(device)
model.to(device)

generated_ids = model.generate(model_inputs, max_new_tokens=1000, do_sample=True)
decoded = tokenizer.batch_decode(generated_ids)
print(decoded[0])
```