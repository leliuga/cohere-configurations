---
license: apache-2.0
language:
- en
library_name: transformers
tags:
- roleplay
- text-generation-inference
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/653a2392341143f7774424d8/fyK_RtEjb9sLF_Mq0nZm2.png)

Based on feedback Estopian made can:

- EstopianMaid is good at sticking to the character card.
- maintains coherency in a setting with multiple characters.
- Able to create new scenario's

- Feature from Thespis:

![image/webp](https://cdn-uploads.huggingface.co/production/uploads/653a2392341143f7774424d8/1Z4P7XshVOW8fLg9pey4H.webp)

- Prompt Template: Alpaca
   ### Instruction:
   {prompt}

   ### Response:

Recommended settings:
- SillyTavern Default Preset.
- Temperature: 0.7
- Min-P: 0.3
- Amount to Gen: 256
- Top P: 1
- Repetition penalty: 1.10

Models used:

    BlueNipples/TimeCrystal-l2-13B
    cgato/Thespis-13b-DPO-v0.7
    KoboldAI/LLaMA2-13B-Estopia
    NeverSleep/Noromaid-13B-0.4-DPO
    Doctor-Shotgun/cat-v1.0-13b

Feedback is always appreciated!
Thank you KoboldAI for their usage of their MergeBox and Caitlyn G. for their support and feedback.