---
license: mit
---
**ALMA** (**A**dvanced **L**anguage **M**odel-based tr**A**nslator) is an LLM-based translation model, which adopts a new translation model paradigm: it begins with fine-tuning on monolingual data and is further optimized using high-quality parallel data. This two-step fine-tuning process ensures strong translation performance. 
Please find more details in our [paper](https://arxiv.org/abs/2309.11674).
```
@misc{xu2023paradigm,
      title={A Paradigm Shift in Machine Translation: Boosting Translation Performance of Large Language Models}, 
      author={Haoran Xu and Young Jin Kim and Amr Sharaf and Hany Hassan Awadalla},
      year={2023},
      eprint={2309.11674},
      archivePrefix={arXiv},
      primaryClass={cs.CL}
}
```
We release four translation models presented in the paper:
- **ALMA-7B**: Full-weight Fine-tune LLaMA-2-7B on 20B monolingual tokens and then **Full-weight** fine-tune on human-written parallel data
- **ALMA-7B-LoRA**: Full-weight Fine-tune LLaMA-2-7B on 20B monolingual tokens and then **LoRA** fine-tune on human-written parallel data
- **ALMA-13B**: Full-weight Fine-tune LLaMA-2-7B on 12B monolingual tokens and then **Full-weight** fine-tune on human-written parallel data
- **ALMA-13B-LoRA** (Our best system): Full-weight Fine-tune LLaMA-2-7B on 12B monolingual tokens and then **LoRA** fine-tune on human-written parallel data
  
Model checkpoints are released at huggingface:
|     Models    | Base Model Link | LoRA Link |
|:-------------:|:---------------:|:---------:|
|    ALMA-7B    |        [haoranxu/ALMA-7B](https://huggingface.co/haoranxu/ALMA-7B)        |     -     |
|  ALMA-7B-LoRA |        [haoranxu/ALMA-7B-Pretrain](https://huggingface.co/haoranxu/ALMA-7B-Pretrain)        |     [haoranxu/ALMA-7B-Pretrain-LoRA](https://huggingface.co/haoranxu/ALMA-7B-Pretrain-LoRA)     |
|    ALMA-13B   |        [haoranxu/ALMA-13B](https://huggingface.co/haoranxu/ALMA-13B)        |     -     |
| ALMA-13B-LoRA |        [haoranxu/ALMA-13B-Pretrain](https://huggingface.co/haoranxu/ALMA-13B-Pretrain)        |     [haoranxu/ALMA-13B-Pretrain-LoRA](https://huggingface.co/haoranxu/ALMA-13B-Pretrain-LoRA)     |

**Note that `ALMA-7B-Pretrain` and `ALMA-13B-Pretrain` are NOT translation models. They only experience stage 1 monolingual fine-tuning (20B tokens for the 7B model and 12B tokens for the 13B model), and should be utilized in conjunction with their LoRA models for translation purposes.** 

A quick start to use our best system (ALMA-13B-LoRA) for translation. An example of translating "我爱机器翻译。" into English:
```
import torch
from peft import PeftModel
from transformers import AutoModelForCausalLM
from transformers import LlamaTokenizer

# Load base model and LoRA weights
model = AutoModelForCausalLM.from_pretrained("haoranxu/ALMA-13B-Pretrain", torch_dtype=torch.float16, device_map="auto")
model = PeftModel.from_pretrained(model, "haoranxu/ALMA-13B-Pretrain-LoRA")
tokenizer = LlamaTokenizer.from_pretrained("haoranxu/ALMA-13B-Pretrain", padding_side='left')

# Add the source setence into the prompt template
prompt="Translate this from Chinese to English:\nChinese: 我爱机器翻译。\nEnglish:"
input_ids = tokenizer(prompt, return_tensors="pt", padding=True, max_length=40, truncation=True).input_ids.cuda()

# Translation
with torch.no_grad():
    generated_ids = model.generate(input_ids=input_ids, num_beams=5, max_new_tokens=20, do_sample=True, temperature=0.6, top_p=0.9)
outputs = tokenizer.batch_decode(generated_ids, skip_special_tokens=True)
print(outputs)
```

Please find more details in our [GitHub repository](https://github.com/fe1ixxu/ALMA)