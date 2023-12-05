---
license: cc-by-nc-4.0
---
This model should be fixed, it was MEANT to be BF16.

Don't mind this one at the moment, I need to finetune it for RP, it's just a test.

## Description

This repo contains fp16 files of Mistral-11B-OmniMix-bf16.

My goal for this model was only to make it score the highest possible with merge and layer toying, proving that:
- Benchmark are objective
- You should try a model yourself and don't go blindly to the highest rated one
- Merge/Layer toying CAN be usable to do better model (maybe?)


## Model used
- [Mistral-7B-OpenOrca](https://huggingface.co/Open-Orca/Mistral-7B-OpenOrca)
- [Mistral-7B-v0.1-Open-Platypus](https://huggingface.co/akjindal53244/Mistral-7B-v0.1-Open-Platypus)
- [CollectiveCognition-v1.1-Mistral-7B](https://huggingface.co/teknium/CollectiveCognition-v1.1-Mistral-7B)
- [zephyr-7b-alpha](https://huggingface.co/HuggingFaceH4/zephyr-7b-alpha)



## Prompt template

The best one after further testing is this one:

```
<|system|>
Below is an instruction that describes a task. Write a response that appropriately completes the request.
<|user|>
{prompt}
<|assistant|>
```


![image/png](https://cdn-uploads.huggingface.co/production/uploads/63ab1241ad514ca8d1430003/tWIx8yeoallv94zrhN6L-.png)

But these one work too:

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:

```

```
USER: <prompt>
ASSISTANT:
```

Or use any prompting system from one of the 4 source model, should work.

## The secret sauce

Mistral-11B-OpenOrcaPlatypus :
```
slices:
  - sources:
    - model: Open-Orca/Mistral-7B-OpenOrca
      layer_range: [0, 24]
  - sources:
    - model: akjindal53244/Mistral-7B-v0.1-Open-Platypus
      layer_range: [8, 32]
merge_method: passthrough
dtype: bfloat16
```

Mistral-11B-CC-Zephyr :
```
slices:
  - sources:
    - model: "/content/drive/MyDrive/CC-v1.1-7B-bf16"
      layer_range: [0, 24]
  - sources:
    - model: "/content/drive/MyDrive/Zephyr-7B"
      layer_range: [8, 32]
merge_method: passthrough
dtype: bfloat16
```

Mistral-11B-OmniMix :
```
slices:
  - sources:
      - model: Mistral-11B-OpenOrcaPlatypus
        layer_range: [0, 48]
      - model: Mistral-11B-CC-Zephyr
        layer_range: [0, 48]
merge_method: slerp
base_model: Mistral-11B-OpenOrcaPlatypus
parameters:
  t:
    - filter: lm_head 
      value: [0.75]
    - filter: embed_tokens
      value: [0.75]
    - filter: self_attn
      value: [0.75, 0.25]
    - filter: mlp
      value:  [0.25, 0.75]
    - filter: layernorm
      value: [0.5, 0.5]
    - filter: modelnorm
      value: [0.75]
    - value: 0.5 # fallback for rest of tensors
dtype: bfloat16
```
I use [mergekit](https://github.com/cg123/mergekit) for all the manipulation told here.

## Some scoring I done myself


![image/png](https://cdn-uploads.huggingface.co/production/uploads/63ab1241ad514ca8d1430003/rnraBZz-I9CUD1GVNVF00.png)

hf-causal-experimental (pretrained=/content/drive/MyDrive/Mistral-11B-OmniMix-bf16), limit: None, provide_description: False, num_fewshot: 0, batch_size: 4
|    Task     |Version| Metric |Value |   |Stderr|
|-------------|------:|--------|-----:|---|-----:|
|arc_challenge|      0|acc     |0.5580|±  |0.0145|
|             |       |acc_norm|0.5819|±  |0.0144|
|arc_easy     |      0|acc     |0.8300|±  |0.0077|
|             |       |acc_norm|0.8211|±  |0.0079|
|hellaswag    |      0|acc     |0.6372|±  |0.0048|
|             |       |acc_norm|0.8209|±  |0.0038|
|piqa         |      0|acc     |0.8145|±  |0.0091|
|             |       |acc_norm|0.8286|±  |0.0088|
|truthfulqa_mc|      1|mc1     |0.3978|±  |0.0171|
|             |       |mc2     |0.5680|±  |0.0155|
|winogrande   |      0|acc     |0.7427|±  |0.0123|

## Others

Special thanks to Sushi, [Henky](https://github.com/KoboldAI/KoboldAI-Client) for the machine he give me for big task, and [Charles Goddard](https://github.com/cg123) for his amazing tool.

If you want to support me, you can [here](https://ko-fi.com/undiai).
