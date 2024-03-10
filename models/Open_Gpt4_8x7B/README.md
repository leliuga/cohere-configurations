---
license: cc-by-4.0
tags:
- merge
- moe
---
Open_Gpt4
cc
____________________________________________________________________________________________
VERSION 0.2 OUT NOW:

Fp16:

- https://huggingface.co/rombodawg/Open_Gpt4_8x7B_v0.2

q8_0.gguf:

- https://huggingface.co/rombodawg/Open_Gpt4_8x7B_v0.2_q8_0_gguf
____________________________________________________________________________________________

![image/jpeg](https://cdn-uploads.huggingface.co/production/uploads/642cc1c253e76b4c2286c58e/T7QKB0fKNHQvNqAjm8zrH.jpeg)

This model is a TIES merger of notux-8x7b-v1 and UNAversal-8x7B-v1beta with MixtralOrochi8x7B being the Base model.

I was very impressed with MixtralOrochi8x7B performance and multifaceted usecases as it is already a merger of many usefull Mixtral models such as Mixtral instruct, 

Noromaid-v0.1-mixtral, openbuddy-mixtral and possibly other models that were not named. My goal was to expand the models capabilities and make it even more useful of a model, maybe even competitive with closed source models like Gpt-4. But for that more testing is required. I hope the community can help me determine if its deserving of its name. 😊

Base model: 

- https://huggingface.co/smelborp/MixtralOrochi8x7B

Merged models:

- https://huggingface.co/fblgit/UNAversal-8x7B-v1beta

- https://huggingface.co/argilla/notux-8x7b-v1


Instruct template: Alpaca


Merger config:
```yaml
models:
  - model: notux-8x7b-v1
    parameters:
      density: .5
      weight: 1
  - model: UNAversal-8x7B-v1beta
    parameters:
      density: .5
      weight: 1


merge_method: ties
base_model: MixtralOrochi8x7B
parameters:
  normalize: true
  int8_mask: true
dtype: float16

```