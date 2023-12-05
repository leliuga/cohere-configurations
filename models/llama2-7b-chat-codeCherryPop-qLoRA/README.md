---

### Overview:
  description:
    
    This is a llama2 7B HF chat model fine-tuned on 122k code instructions. In my early experiments it seems to be doing very well.
  
  additional_info:
    
    It's a bottom of the barrel model 😂 but after quantization it can be
    valuable for sure. It definitely proves that a 7B can be useful for boilerplate
    code stuff though.

### Plans:
  next_steps: "I've a few things in mind and after that this will be more valuable."
  
  tasks:
  
    - name: "I'll quantize these"
      timeline: "Possibly tonight or tomorrow in the day"
      result: "Then it can be run locally with 4G ram."
    - name: "I've used alpaca style instruction tuning"
      improvement: |
        I'll switch to llama2 style [INST]<<SYS>> style and see if
        it improves anything.
    - name: "HumanEval report and checking for any training data leaks"
    - attempt: "I'll try 8k context via RoPE enhancement"
      hypothesis: "Let's see if that degrades performance or not."
  commercial_use: |
    So far I think this can be used commercially but this is a adapter on Meta's llama2 with
    some gating issues so that is there.
  contact_info: "If you find any issues or want to just holler at me, you can reach out to me - https://twitter.com/4evaBehindSOTA"

### Library:
  name: "peft"

### Training procedure:
  quantization_config:
    load_in_8bit: False
    load_in_4bit: True
    llm_int8_threshold: 6.0
    llm_int8_skip_modules: None
    llm_int8_enable_fp32_cpu_offload: False
    llm_int8_has_fp16_weight: False
    bnb_4bit_quant_type: "nf4"
    bnb_4bit_use_double_quant: False
    bnb_4bit_compute_dtype: "float16"

### Framework versions:
  PEFT: "0.5.0.dev0"