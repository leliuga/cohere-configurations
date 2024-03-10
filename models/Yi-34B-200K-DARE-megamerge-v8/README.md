---
license: other
license_name: yi-license
license_link: https://huggingface.co/01-ai/Yi-34B/blob/main/LICENSE
language:
- en
library_name: transformers
base_model: []
tags:
- mergekit
- merge
- Yi
---
# Yi 34B 200K DARE Merge v8

A merge of many Yi 34B 200K models using the new DARE Ties method via mergekit. The goal is to create a merge model that excels at 32K+ context performance, without any additional finetuning.

## Prompt template: Orca-Vicuna
```
SYSTEM: {system_message}
USER: {prompt}
ASSISTANT:
```
It might recognize ChatML, and possibly Alpaca-like formats. Raw prompting as described here is also effective: https://old.reddit.com/r/LocalLLaMA/comments/18zqy4s/the_secret_to_writing_quality_stories_with_llms/



## Running
Being a Yi model, run a lower temperature with 0.1 or higher MinP, a little repetition penalty, maybe mirostat with a low tau, and no other samplers. Yi tends to run "hot" by default, and it really needs a low temperature + MinP to cull Yi's huge vocabulary. See the explanation here: https://github.com/ggerganov/llama.cpp/pull/3841

24GB GPUs can efficiently run Yi-34B-200K models at **40K-90K context** with exllamav2, and performant UIs like [exui](https://github.com/turboderp/exui). I go into more detail in this [post](https://old.reddit.com/r/LocalLLaMA/comments/1896igc/how_i_run_34b_models_at_75k_context_on_24gb_fast/). 16GB GPUs can still run the high context with aggressive quantization.

Lonestriker has also uploaded general purpose quantizations here: https://huggingface.co/models?sort=trending&search=LoneStriker+Yi-34B-200K-DARE-megamerge-v8

Additionally, TheBloke has uploaded experimental GGUFs using llama.cpp's new imatrix quantization feature, profiled on VMware open-instruct: https://huggingface.co/TheBloke/Yi-34B-200K-DARE-megamerge-v8-GGUF

To load/train this in full-context backends like transformers, you *must* change `max_position_embeddings` in config.json to a lower value than 200,000, otherwise you will OOM! I do not recommend running high context without context-efficient backends like exllamav2, litellm or unsloth.


## Testing Notes

See: https://huggingface.co/brucethemoose/Yi-34B-200K-DARE-merge-v5#testing-notes

An intermediate merge model was created to try and extend the context of several 4k models before adding them to the main merge, as seen in the "megamerge" recipe below. I can upload this upon request

In addition, the weight gradients are biased towards Vicuna-format models in the first few layers to try and "emphasize" the Orca-Vicuna prompt template. How sucessful this is remains to be seen. 



## Merge Details
### Merge Method

This model was merged using the [DARE](https://arxiv.org/abs/2311.03099) [TIES](https://arxiv.org/abs/2306.01708) merge method using /home/alpha/Storage/Models/Raw/chargoddard_Yi-34B-200K-Llama as a base.

### Models Merged

The following models were included in the merge:
* https://huggingface.co/kyujinpy/PlatYi-34B-200k-Q-FastChat
* https://huggingface.co/jondurbin/bagel-34b-v0.2
* https://huggingface.co/migtissera/Tess-M-Creative-v1.0
* https://huggingface.co/brucethemoose/SUS-Bagel-200K-DARE-Test
* https://huggingface.co/Mihaiii/Pallas-0.5
* https://huggingface.co/bhenrym14/airoboros-3_1-yi-34b-200k
* https://huggingface.co/adamo1139/Yi-34B-200K-AEZAKMI-v2
* https://huggingface.co/migtissera/Tess-34B-v1.4
* https://huggingface.co/SUSTech/SUS-Chat-34B
* https://huggingface.co/jondurbin/bagel-dpo-34b-v0.2
* https://huggingface.co/bhenrym14/platypus-yi-34b
* https://huggingface.co/Weyaxi/Nous-Hermes-2-SUS-Chat-34B-Slerp
* https://huggingface.co/TriadParty/deepsex-34b
* https://huggingface.co/TriadParty/deepmoney-34b-200k-base
* https://huggingface.co/chargoddard/Yi-34B-200K-Llama
* https://huggingface.co/chargoddard/Yi-34B-Llama

### Configuration

The following YAML configuration was used to produce this model:

```yaml
models:
  - model: /home/alpha/Models/Raw/chargoddard_Yi-34B-Llama
  # No parameters necessary for base model
  - model: /home/alpha/Storage/Models/Raw/chargoddard_Yi-34B-200K-Llama
  #200K base to extend the context of 4K models, max density as we *want* it to 'interfere'
    parameters:
      weight: 0.33
      density: 1
  - model: /home/alpha/Models/Raw/Weyaxi_Nous-Hermes-2-SUS-Chat-34B-Slerp
    parameters:
      weight: 0.15
      density: 0.36
  - model: /home/alpha/Models/Raw/jondurbin_bagel-dpo-34b-v0.2
  #Mix dpo with sft to tone down dpo
    parameters:
      weight: 0.06
      density: 0.36
  - model: /home/alpha/Models/Raw/jondurbin_bagel-34b-v0.2
    parameters:
      weight: 0.06
      density: 0.41
  - model: /home/alpha/Models/Raw/bhenrym14_platypus-yi-34b
  #Vicuna format
    parameters:
      weight: 0.19
      density: 0.41
  # - model: /home/alpha/Models/Raw/01-ai_Yi-34B-Chat #+/home/alpha/Models/Raw/Doctor-Shotgun_limarpv3-yi-llama-34b-lora
  # #Can't get lora OR base model to work without erroring out?
  #   parameters:
  #     weight: 0.04
  #     density: 0.36
  - model: /home/alpha/Models/Raw/TriadParty_deepsex-34b
  #Base model with no prompt
    parameters:
      weight: 0.21
      density: 0.39
merge_method: dare_ties
tokenizer_source: union
base_model: /home/alpha/Models/Raw/chargoddard_Yi-34B-Llama
parameters:
  int8_mask: true
dtype: bfloat16
name: 4kmerge-v2
---
models:
  - model: /home/alpha/Storage/Models/Raw/chargoddard_Yi-34B-200K-Llama
    # No parameters necessary for base model
  - model: /home/alpha/Storage/Models/Raw/migtissera_Tess-34B-v1.4
    #Emphasize the beginning of Vicuna format models
    parameters:
      weight: [0.22, 0.113, 0.113, 0.113, 0.113, 0.113]
      density: 0.61
  - model: /home/alpha/Models/Raw/Mihaiii_Pallas-0.5
    # Vicuna format
    parameters:
      weight: [0.22, 0.113, 0.113, 0.113, 0.113, 0.113]
      density: 0.61
  - model: /home/alpha//Storage/Models/Raw/bhenrym14_airoboros-3_1-yi-34b-200k
    parameters:
      weight: [0.02, 0.081, 0.081, 0.081, 0.081, 0.081]
      density: 0.59
  - model: /home/alpha/Storage/Models/Raw/jondurbin_bagel-34b-v0.2
    #Only the SFT in the main merge since the DPO version seems to have no long context ability at all, and some overfitting(?) issues
    parameters:
      weight: [0.02, 0.093, 0.093, 0.093, 0.093, 0.093]
      density: 0.4
  - model: /home/alpha/Storage/Models/Raw/kyujinpy_PlatYi-34B-200k-Q-FastChat
    parameters:
      weight: [0.02, 0.081, 0.081, 0.081, 0.081, 0.081]
      density: 0.59
  #- model: /home/alpha/Storage/Models/Raw/ehartford_dolphin-2.2-yi-34b-200k
  #  Dolphin 200K seems to be funky according to multiple leaderboards and perplexity tests?
  #  parameters:
  #    weight: 0.15
  #    density: 0.6
  - model: /home/alpha/Models/Raw/adamo1139_Yi-34B-200K-AEZAKMI-v2
    parameters:
      weight: [0.02, 0.096, 0.096, 0.096, 0.096, 0.096]
      density: 0.59
  - model: /home/alpha/Storage/Models/Raw/Nous-Capybara-34B
    parameters:
      weight:  [0.21, 0.115, 0.115, 0.115, 0.115, 0.115]
      density: 0.59
  - model: 4kmerge-v2
  #Previous merge
    parameters:
      weight: [0.02, 0.115, 0.115, 0.115, 0.115, 0.115]
      density: 0.4
  - model: /home/alpha/Models/Raw/migtissera_Tess-M-Creative-v1.0
  # Vicuna format
    parameters:
      weight: [0.21, 0.09, 0.09, 0.09, 0.09, 0.09]
      density: 0.61
  - model: /home/alpha/Models/Raw/TriadParty_deepmoney-34b-200k-base
  # No prompt format, native long context full finetune
    parameters:
      weight: [0.04, 0.103, 0.103, 0.103, 0.103, 0.103]
      density: 0.61
merge_method: dare_ties
tokenizer_source: union
base_model: /home/alpha/Storage/Models/Raw/chargoddard_Yi-34B-200K-Llama
parameters:
  int8_mask: true
dtype: bfloat16
```


## Self Promotion

I'm part of a AI startup called Holocene AI!

We're new, busy, and still setting things up. But if you have any business inquiries, want a job, or just want some consultation, feel free to shoot me an email. We have expertise in RAG applications and llama/embeddings model finetuning, and absolutely *none* of the nonsense of scammy AI startups.

Contact me at: agates.holocene.ai@gmail.com

I also set up a Ko-Fi! I want to run some (personal) training/LASERing as well, at 100K context or so. If you'd like to buy me 10 minutes on an A100 (or 5 seconds on an MI300X), I'd appreciate it: https://ko-fi.com/alphaatlas