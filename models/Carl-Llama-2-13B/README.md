---
license: cc-by-nc-nd-4.0
datasets:
- jerryjalapeno/nart-100k-synthetic
language:
- en
---

**Carl: A Therapist AI**

Early prevention can help lot of people to avoid depression and other mental illnesses. Therapy is a controversial use case because the outputs and capabilities of LLMs are uncertain. 
Many people don't have access the therapist, due to a financial, personal, social or other restriction.
Here comes Carl: A Therapist AI which can quickly respond to you. It is trained on more than 100000 set of conversations. Each set having 10~15 conversations between Carl and client.
Base data was obtained from jerryjalapeno/nart-100k-synthetic . This data was further refined and fine tuned. Entire dataset is synthetic. Synthetic data is used because there is little to no therapy conversation data which is publicly available and directly applicable to an LLM.
This by means a no replacement to a Doctor or professional therapist. If you are in stress or going through a tough time, please seek professional help or talk to a friend/family member.

**Training:**
Entire dataset was trained on Azure 4 x A100 80GB. For 3 epoch, training took 50 hours. DeepSpeed codebase was used for training purpose. This was trained on Llama-2 by Meta.
GGML Quant models are converted by Kijana Mitchell. Extremely thankful to him.

**GPTQ**

GPTQ: [TheBloke](https://huggingface.co/TheBloke/Carl-Llama-2-13B-GPTQ)


Special Thanks to [TheBloke](https://huggingface.co/TheBloke) for guiding me and making this model available. 

**Example Prompt:**
```
This is a conversation with your Therapist AI, Carl. Carl is designed to help you while in stress. It can answer your questions and help you to calm down

Context
You are Carl, A Therapist AI
USER: <prompt>
CARL:
```

Note:
This is just a research experiment, and the model should NOT be used as a human therapist. Use "cat" command to join all pytorch_model.bin parts.