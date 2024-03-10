---
inference: false
language:
- en
pipeline_tag: text-generation
---


## WizardMath: Empowering Mathematical Reasoning for Large Language Models via Reinforced Evol-Instruct (RLEIF)

<p style="font-size:28px;" align="center">
🏠 <a href="https://wizardlm.github.io/" target="_blank">Home Page</a> </p>
<p align="center">
<p align="center">
🤗 <a href="https://huggingface.co/WizardLM" target="_blank">HF Repo</a>  •🐱 <a href="https://github.com/nlpxucan/WizardLM" target="_blank">Github Repo</a> • 🐦 <a href="https://twitter.com/WizardLM_AI" target="_blank">Twitter</a> </p>
<p align="center">
 📃 <a href="https://arxiv.org/abs/2304.12244" target="_blank">[WizardLM]</a>  • 📃 <a href="https://arxiv.org/abs/2306.08568" target="_blank">[WizardCoder]</a>   • 📃 <a href="https://arxiv.org/abs/2308.09583" target="_blank">[WizardMath]</a>  <br>
</p>
<p align="center">
    👋 Join our <a href="https://discord.gg/VZjjHtWrKs" target="_blank">Discord</a>
</p>

## News

[12/19/2023] 🔥 We released **WizardMath-7B-V1.1**  trained from Mistral-7B, the **SOTA 7B math LLM**, achieves **83.2 pass@1** on GSM8k, and **33.0 pass@1** on MATH. Use this [[**Demo**](http://47.103.63.15:50083/)] to chat with it.

[12/19/2023] 🔥 **WizardMath-7B-V1.1** outperforms **ChatGPT 3.5**, **Gemini Pro**, **Mixtral MOE**, and **Claude Instant** on GSM8K pass@1.

[12/19/2023] 🔥 **WizardMath-7B-V1.1** is comparable with **ChatGPT 3.5**, **Gemini Pro**, and surpasses **Mixtral MOE** on MATH pass@1.

| Model | Checkpoint | Paper  | GSM8k | MATH  | Demo|
| ----- |------| ---- |------|-------|-------| 
| **WizardMath-7B-V1.1** | 🤗 <a href="https://huggingface.co/WizardLM/WizardMath-7B-V1.1" target="_blank">HF Link</a>  |  📃 <a href="https://arxiv.org/abs/2308.09583" target="_blank">[WizardMath]</a>| 	 **83.2**  |  **33.0** |[[**Demo**](http://47.103.63.15:50083/)] |
| WizardMath-70B-V1.0 | 🤗 <a href="https://huggingface.co/WizardLM/WizardMath-70B-V1.0" target="_blank">HF Link</a> |  📃 <a href="https://arxiv.org/abs/2308.09583" target="_blank">[WizardMath]</a>| **81.6**  |  **22.7**	||
| WizardMath-13B-V1.0 | 🤗 <a href="https://huggingface.co/WizardLM/WizardMath-13B-V1.0" target="_blank">HF Link</a> |  📃 <a href="https://arxiv.org/abs/2308.09583" target="_blank">[WizardMath]</a>| **63.9**  |  **14.0** ||
| WizardMath-7B-V1.0 | 🤗 <a href="https://huggingface.co/WizardLM/WizardMath-7B-V1.0" target="_blank">HF Link</a>  |  📃 <a href="https://arxiv.org/abs/2308.09583" target="_blank">[WizardMath]</a>| 	 **54.9**  |  **10.7** |  |
   

## [12/19/2023] Comparing WizardMath-7B-V1.1 with other open source 7B size math LLMs.

| Model | GSM8k Pass@1 | MATH Pass@1 |
| ----- |------| ---- |
| MPT-7B              | 6.8          | 3.0         |
|Llama 1-7B          | 11.0         | 2.9         |
|Llama 2-7B|12.3	|2.8	|
|Yi-6b| 32.6	|5.8	|
|Mistral-7B|37.8	|9.1	|
|Qwen-7b|47.8	|9.3	|
| RFT-7B              | 50.3         | --          |
| MAmmoTH-7B (COT)    | 50.5         | 10.4        |
| WizardMath-7B-V1.0 | 54.9  |  10.7 |
|Abel-7B-001 |59.7	|13	|
| MetaMath-7B         | 66.5         | 19.8        |
| Arithmo-Mistral-7B  | 74.7         | 25.3        |
|MetaMath-Mistral-7B|77.7	|28.2	|
|Abel-7B-002 |	80.4 | 29.5	|
| **WizardMath-7B-V1.1** |  **83.2**  |  **33.0** |


## [12/19/2023] Comparing WizardMath-7B-V1.1 with large open source (30B~70B) LLMs.

| Model | GSM8k Pass@1 | MATH Pass@1 |
| ----- |------| ---- |
| Llemma-34B             | 51.5          |   25.0       |
| Minerva-62B             | 52.4          |  27.6        |
| Llama 2-70B             | 56.8          |  13.5        |
|  DeepSeek 67B            | 63.4          |  --        |
|  Gork 33B            | 62.9          |  23.9       |
| MAmmoTH-70B             | 72.4          |  21.1       |
| Yi-34B            | 67.9          |   15.9       |
| Mixtral 8x7B            | 74.4          |   28.4      |
|  MetaMath-70B  | 82.3          |   26.6       |
| **WizardMath-7B-V1.1** |  **83.2**  |  **33.0** |


## ❗ Data Contamination Check:

Before model training, we carefully and rigorously checked all the training data, and used multiple deduplication methods to verify and prevent data leakage on GSM8k and MATH test set. 

🔥 
❗<b>Note for model system prompts usage:</b>

Please use **the same systems prompts strictly** with us, and we do not guarantee the accuracy of the **quantified versions**.

**Default version:**

```
"Below is an instruction that describes a task. Write a response that appropriately completes the request.\n\n### Instruction:\n{instruction}\n\n### Response:"
```


**CoT Version:** （❗For the **simple** math questions, we do NOT recommend to use the CoT prompt.） 


```
"Below is an instruction that describes a task. Write a response that appropriately completes the request.\n\n### Instruction:\n{instruction}\n\n### Response: Let's think step by step."
```

## Inference WizardMath Demo Script

We provide the WizardMath inference demo code [here](https://github.com/nlpxucan/WizardLM/tree/main/demo).




## Citation

Please cite the repo if you use the data, method or code in this repo.

```
@article{luo2023wizardmath,
  title={WizardMath: Empowering Mathematical Reasoning for Large Language Models via Reinforced Evol-Instruct},
  author={Luo, Haipeng and Sun, Qingfeng and Xu, Can and Zhao, Pu and Lou, Jianguang and Tao, Chongyang and Geng, Xiubo and Lin, Qingwei and Chen, Shifeng and Zhang, Dongmei},
  journal={arXiv preprint arXiv:2308.09583},
  year={2023}
}
```
