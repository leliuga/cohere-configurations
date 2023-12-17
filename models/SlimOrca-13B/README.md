---
license: cc-by-nc-nd-4.0
datasets:
- Open-Orca/SlimOrca
- ajibawa-2023/SlimOrca-ShareGPT
language:
- en
---

**SlimOrca-13B: A General Purpose Intelligent Model**


This Model is trained on refined version of SlimOrca made available by [Open-Orca](https://huggingface.co/Open-Orca) team. 
The idea was to check how this Model will perform in the absence of "system" prompt/instruction. 
This Model is very good in various types of General Purpose content generation such as Q&A (including multiple choice), Articles from Summary, Sentiment Analysis,
Context & Hypothesis, Reviews, Erotic story generation etc.
It can also generate Uncensored content. Kindly be careful while generating Uncensored content as you will be responsible for what you
generate.

It is trained on 517981 set of conversations. Each set having 2 conversations. I have shared this [data](https://huggingface.co/datasets/ajibawa-2023/SlimOrca-ShareGPT).

All the credit goes to the Open-Orca team for releasing SlimOrca dataset.


**Training:**
Entire dataset was trained on Azure 4 x A100 80GB. For 3 epoch, training took almost 11 Days. DeepSpeed codebase was used for training purpose.
Entire data is trained on Llama-2 by Meta.

This is a full fine tuned model. Links for quantized models are given below.

**GPTQ GGML & AWQ**

GPTQ: [Link](https://huggingface.co/TheBloke/SlimOrca-13B-GPTQ)

GGUF: [Link](https://huggingface.co/TheBloke/SlimOrca-13B-GGUF)

AWQ: [Link](https://huggingface.co/TheBloke/SlimOrca-13B-AWQ)


Special Thanks to [TheBloke](https://huggingface.co/TheBloke) for making these models available.



**Example Prompt:**
```
This is a conversation with your Assistant. It is a computer program designed to help you with various tasks such as answering questions, providing recommendations, and helping with decision making. You can ask it anything you want and it will do its best to give you accurate and relevant information.

Context
You are a helpful AI assistant.

USER: <prompt>
ASSISTANT:
```
You can modify above Prompt as per your requirement. I have used ShareGPT/Vicuna format v1.1 .


I want to say special Thanks to the Open Source community for helping & guiding me to better understand the AI/Model development.

Thank you for your love & support.


**Example Output**

Example 1

![Example 1](https://cdn-uploads.huggingface.co/production/uploads/64aea8ff67511bd3d965697b/hM_EJaSZiMjMQU35EiHGM.png)

Example 2

![Example 2](https://cdn-uploads.huggingface.co/production/uploads/64aea8ff67511bd3d965697b/riNaxJeTWdCEE4dNP8GWp.png)


