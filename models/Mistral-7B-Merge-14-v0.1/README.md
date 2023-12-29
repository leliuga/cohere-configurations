---
license: cc-by-nc-4.0
language:
- en
---

# Update 2023-12-19

In light of [dataset contamination issue among the merged models](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard/discussions/474)
raised by the community in recent days, in particular
[berkeley-nest/Starling-LM-7B-alpha](https://huggingface.co/berkeley-nest/Starling-LM-7B-alpha),
[Q-bert/MetaMath-Cybertron-Starling](https://huggingface.co/Q-bert/MetaMath-Cybertron-Starling), and
[janai-hq/trinity-v1](https://huggingface.co/janai-hq/trinity-v1),
we decided to remake another model without the models mentioned.
Additionally, their CC-by-NC-4.0 license is restrictive and thus are not suitable for an open model.

# Model Description
This is an experiment to test merging 14 models using DARE TIES 🦙

The merged model is then merged again with [janai-hq/trinity-v1](https://huggingface.co/janai-hq/trinity-v1) using Gradient SLERP.
The result is a base model that performs quite well but requires some further instruction fine-tuning.

The 14 models are as follows:
1. [mistralai/Mistral-7B-Instruct-v0.2](https://huggingface.co/mistralai/Mistral-7B-Instruct-v0.2)
2. [ehartford/dolphin-2.2.1-mistral-7b](https://huggingface.co/ehartford/dolphin-2.2.1-mistral-7b)
3. [SciPhi/SciPhi-Mistral-7B-32k](https://huggingface.co/SciPhi/SciPhi-Mistral-7B-32k)
4. [ehartford/samantha-1.2-mistral-7b](https://huggingface.co/ehartford/samantha-1.2-mistral-7b)
5. [Arc53/docsgpt-7b-mistral](https://huggingface.co/Arc53/docsgpt-7b-mistral)
6. [berkeley-nest/Starling-LM-7B-alpha](https://huggingface.co/berkeley-nest/Starling-LM-7B-alpha)
7. [Q-bert/MetaMath-Cybertron-Starling](https://huggingface.co/Q-bert/MetaMath-Cybertron-Starling)
8. [Open-Orca/Mistral-7B-OpenOrca](https://huggingface.co/Open-Orca/Mistral-7B-OpenOrca)
9. [v1olet/v1olet_marcoroni-go-bruins-merge-7B](https://huggingface.co/v1olet/v1olet_marcoroni-go-bruins-merge-7B)
10. [beowolx/MistralHermes-CodePro-7B-v1](https://huggingface.co/beowolx/MistralHermes-CodePro-7B-v1)
11. [TIGER-Lab/MAmmoTH-7B-Mistral](https://huggingface.co/TIGER-Lab/MAmmoTH-7B-Mistral)
12. [teknium/OpenHermes-2.5-Mistral-7B](https://huggingface.co/teknium/OpenHermes-2.5-Mistral-7B)
13. [Weyaxi/OpenHermes-2.5-neural-chat-v3-3-Slerp](https://huggingface.co/Weyaxi/OpenHermes-2.5-neural-chat-v3-3-Slerp)
14. [mlabonne/NeuralHermes-2.5-Mistral-7B](https://huggingface.co/mlabonne/NeuralHermes-2.5-Mistral-7B)

- base model: [mistralai/Mistral-7B-v0.1](https://huggingface.co/mistralai/Mistral-7B-v0.1)

The yaml config file for this model is here:

```yaml
slices:
  - sources:
      - model: EmbeddedLLM/Mistral-7B-Merge-14-v0
        layer_range: [0, 32]
      - model: janai-hq/trinity-v1
        layer_range: [0, 32]
merge_method: slerp
base_model: EmbeddedLLM/Mistral-7B-Merge-14-v0
parameters:
  t:
    - filter: self_attn
      value: [0, 0.5, 0.3, 0.7, 1]
    - filter: mlp
      value: [1, 0.5, 0.7, 0.3, 0]
    - value: 0.5
dtype: bfloat16

```