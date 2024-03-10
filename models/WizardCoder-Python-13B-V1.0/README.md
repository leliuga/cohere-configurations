---
license: llama2
metrics:
- code_eval
library_name: transformers
tags:
- code
model-index:
- name: WizardCoder-Python-13B-V1.0
  results:
  - task:
      type: text-generation
    dataset:
      type: openai_humaneval
      name: HumanEval
    metrics:
    - name: pass@1
      type: pass@1
      value: 0.64
      verified: false
---

## WizardCoder: Empowering Code Large Language Models with Evol-Instruct

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

[2024/01/04] 🔥 We released **WizardCoder-33B-V1.1**  trained from deepseek-coder-33b-base, the **SOTA OSS Code LLM** on [EvalPlus Leaderboard](https://evalplus.github.io/leaderboard.html), achieves **79.9 pass@1** on HumanEval, **73.2 pass@1** on HumanEval-Plus, **78.9 pass@1** on MBPP, and **66.9 pass@1** on MBPP-Plus.

[2024/01/04] 🔥 **WizardCoder-33B-V1.1** outperforms **ChatGPT 3.5**, **Gemini Pro**, and **DeepSeek-Coder-33B-instruct** on HumanEval and HumanEval-Plus pass@1.

[2024/01/04] 🔥 **WizardCoder-33B-V1.1** is comparable with **ChatGPT 3.5**, and surpasses **Gemini Pro** on MBPP and MBPP-Plus pass@1.

|  Model  |  Checkpoint  | Paper    | HumanEval  |   HumanEval+ | MBPP | MBPP+ | License |
| ----- |------| ---- |------|-------| ----- |  ----- |----- | 
|  GPT-4-Turbo (Nov 2023)  | - | - | 85.4  | 81.7 | 83.0 | 70.7 |-|
|  GPT-4 (May 2023)  | - | - | 88.4  | 76.8 | - | - |-|
|  GPT-3.5-Turbo (Nov 2023)  | - | - | 72.6  | 65.9 | 81.7 | 69.4 |-|
|  Gemini Pro  | - | - | 63.4  | 55.5 | 72.9 | 57.9 |-|
|  DeepSeek-Coder-33B-instruct | - | - |  78.7 | 72.6 | 78.7 | 66.7 |-|
|  **WizardCoder-33B-V1.1**  |   🤗 <a href="https://huggingface.co/WizardLM/WizardCoder-33B-V1.1" target="_blank">HF Link</a>   |  📃 <a href="https://arxiv.org/abs/2306.08568" target="_blank">[WizardCoder]</a>  |  79.9  | 73.2 | 78.9 | 66.9 |  <a href="https://huggingface.co/WizardLM/WizardMath-7B-V1.1/resolve/main/LICENSE" target="_blank">MSFTResearch</a>  |
|  WizardCoder-Python-34B-V1.0  |   🤗 <a href="https://huggingface.co/WizardLM/WizardCoder-Python-34B-V1.0" target="_blank">HF Link</a>   |  📃 <a href="https://arxiv.org/abs/2306.08568" target="_blank">[WizardCoder]</a>  |  73.2   | 64.6 | 73.2 | 59.9 |  <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama2</a>  |
|  WizardCoder-15B-V1.0  |   🤗 <a href="https://huggingface.co/WizardLM/WizardCoder-15B-V1.0" target="_blank">HF Link</a>   |  📃 <a href="https://arxiv.org/abs/2306.08568" target="_blank">[WizardCoder]</a>  |  59.8   | 52.4 | -- | -- |  <a href="https://huggingface.co/spaces/bigcode/bigcode-model-license-agreement" target="_blank">OpenRAIL-M</a>  |
|  WizardCoder-Python-13B-V1.0  |   🤗 <a href="https://huggingface.co/WizardLM/WizardCoder-Python-13B-V1.0" target="_blank">HF Link</a>   |  📃 <a href="https://arxiv.org/abs/2306.08568" target="_blank">[WizardCoder]</a>  |  64.0   | -- | -- | -- |  <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama2</a>  |
|  WizardCoder-Python-7B-V1.0  |   🤗 <a href="https://huggingface.co/WizardLM/WizardCoder-Python-7B-V1.0" target="_blank">HF Link</a>   |  📃 <a href="https://arxiv.org/abs/2306.08568" target="_blank">[WizardCoder]</a>  |  55.5   | -- | -- | -- |  <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama2</a>  |
|  WizardCoder-3B-V1.0  |   🤗 <a href="https://huggingface.co/WizardLM/WizardCoder-3B-V1.0" target="_blank">HF Link</a>   |  📃 <a href="https://arxiv.org/abs/2306.08568" target="_blank">[WizardCoder]</a>  |  34.8   | -- | -- | -- |  <a href="https://huggingface.co/spaces/bigcode/bigcode-model-license-agreement" target="_blank">OpenRAIL-M</a>  |
|  WizardCoder-1B-V1.0  |   🤗 <a href="https://huggingface.co/WizardLM/WizardCoder-1B-V1.0" target="_blank">HF Link</a>   |  📃 <a href="https://arxiv.org/abs/2306.08568" target="_blank">[WizardCoder]</a>  |  23.8   | -- | -- | -- |  <a href="https://huggingface.co/spaces/bigcode/bigcode-model-license-agreement" target="_blank">OpenRAIL-M</a>  |


-  Our **WizardMath-70B-V1.0** model slightly outperforms some closed-source LLMs on the GSM8K, including **ChatGPT 3.5**, **Claude Instant 1** and **PaLM 2 540B**.
-  Our **WizardMath-70B-V1.0** model achieves  **81.6 pass@1** on the [GSM8k Benchmarks](https://github.com/openai/grade-school-math), which is **24.8** points higher than the SOTA open-source LLM, and achieves  **22.7 pass@1** on the [MATH Benchmarks](https://github.com/hendrycks/math), which is **9.2** points higher than the SOTA open-source LLM.

<font size=4>
    
| Model | Checkpoint | Paper  | GSM8k | MATH  |Online Demo| License|
| ----- |------| ---- |------|-------| ----- | ----- |
| WizardMath-70B-V1.0 | 🤗 <a href="https://huggingface.co/WizardLM/WizardMath-70B-V1.0" target="_blank">HF Link</a> |  📃 <a href="https://arxiv.org/abs/2308.09583" target="_blank">[WizardMath]</a>| **81.6**  |  **22.7**	|[Demo](http://47.103.63.15:50083/)| <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama 2  </a> |
| WizardMath-13B-V1.0 | 🤗 <a href="https://huggingface.co/WizardLM/WizardMath-13B-V1.0" target="_blank">HF Link</a> |  📃 <a href="https://arxiv.org/abs/2308.09583" target="_blank">[WizardMath]</a>| **63.9**  |  **14.0** |[Demo](http://47.103.63.15:50082/)| <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama 2 </a> |
| WizardMath-7B-V1.0 | 🤗 <a href="https://huggingface.co/WizardLM/WizardMath-7B-V1.0" target="_blank">HF Link</a>  |  📃 <a href="https://arxiv.org/abs/2308.09583" target="_blank">[WizardMath]</a>| 	 **54.9**  |  **10.7** | [Demo ](http://47.103.63.15:50080/)|  <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama 2  </a>|
</font>


- [08/09/2023] We released **WizardLM-70B-V1.0** model. Here is [Full Model Weight](https://huggingface.co/WizardLM/WizardLM-70B-V1.0). 

<font size=4>
    
   
| <sup>Model</sup> | <sup>Checkpoint</sup> | <sup>Paper</sup> |<sup>MT-Bench</sup> | <sup>AlpacaEval</sup>  | <sup>GSM8k</sup> | <sup>HumanEval</sup>  | <sup>License</sup>|
| ----- |------| ---- |------|-------| ----- | ----- | ----- | 
| <sup>**WizardLM-70B-V1.0**</sup> | <sup>🤗 <a href="https://huggingface.co/WizardLM/WizardLM-70B-V1.0" target="_blank">HF Link</a> </sup>|<sup>📃**Coming Soon**</sup>| <sup>**7.78**</sup> | <sup>**92.91%**</sup>	 |<sup>**77.6%**</sup>	 | <sup>   **50.6**</sup>|<sup> <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama 2 License </a></sup> |
| <sup>WizardLM-13B-V1.2</sup> | <sup>🤗 <a href="https://huggingface.co/WizardLM/WizardLM-13B-V1.2" target="_blank">HF Link</a> </sup>|  | <sup>7.06</sup> | <sup>89.17%</sup>	 |<sup>55.3%</sup>	 | <sup>36.6   </sup>|<sup> <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama 2 License </a></sup> |
| <sup>WizardLM-13B-V1.1</sup> |<sup> 🤗 <a href="https://huggingface.co/WizardLM/WizardLM-13B-V1.1" target="_blank">HF Link</a> </sup> |  | <sup>6.76</sup>  |<sup>86.32%</sup>	 | 	 | <sup>25.0   </sup>| <sup>Non-commercial</sup>|
| <sup>WizardLM-30B-V1.0</sup> | <sup>🤗 <a href="https://huggingface.co/WizardLM/WizardLM-30B-V1.0" target="_blank">HF Link</a></sup>  | | <sup>7.01</sup> |                    | |  <sup>37.8  </sup>| <sup>Non-commercial</sup> |
| <sup>WizardLM-13B-V1.0</sup> | <sup>🤗 <a href="https://huggingface.co/WizardLM/WizardLM-13B-V1.0" target="_blank">HF Link</a> </sup> |  | <sup>6.35</sup> | <sup>75.31%</sup> |  | <sup> 24.0   </sup> | <sup>Non-commercial</sup>|
| <sup>WizardLM-7B-V1.0 </sup>|  <sup>🤗 <a href="https://huggingface.co/WizardLM/WizardLM-7B-V1.0" target="_blank">HF Link</a> </sup> |<sup> 📃 <a href="https://arxiv.org/abs/2304.12244" target="_blank">[WizardLM]</a> </sup>|  |  |  |<sup>19.1 </sup>|<sup> Non-commercial</sup>|
</font>

## Comparing WizardCoder-Python-34B-V1.0 with Other LLMs.

🔥 The following figure shows that our **WizardCoder-Python-34B-V1.0 attains the second position in this benchmark**, surpassing GPT4 (2023/03/15, 73.2 vs. 67.0), ChatGPT-3.5 (73.2 vs. 72.5) and Claude2 (73.2 vs. 71.2).

<p align="center" width="100%">
<a ><img src="https://raw.githubusercontent.com/nlpxucan/WizardLM/main/WizardCoder/imgs/compare_sota.png" alt="WizardCoder" style="width: 96%; min-width: 300px; display: block; margin: auto;"></a>
</p>

## Prompt Format
```
"Below is an instruction that describes a task. Write a response that appropriately completes the request.\n\n### Instruction:\n{instruction}\n\n### Response:"
```

## Inference Demo Script

We provide the inference demo code [here](https://github.com/nlpxucan/WizardLM/tree/main/demo).

Note: This script supports `WizardLM/WizardCoder-Python-34B/13B/7B-V1.0`. If you want to inference with `WizardLM/WizardCoder-15B/3B/1B-V1.0`, please change the `stop_tokens = ['</s>']` to `stop_tokens = ['<|endoftext|>']` in the script.

## Citation

Please cite the repo if you use the data, method  or code in this repo.

```
@article{luo2023wizardcoder,
  title={WizardCoder: Empowering Code Large Language Models with Evol-Instruct},
  author={Luo, Ziyang and Xu, Can and Zhao, Pu and Sun, Qingfeng and Geng, Xiubo and Hu, Wenxiang and Tao, Chongyang and Ma, Jing and Lin, Qingwei and Jiang, Daxin},
  journal={arXiv preprint arXiv:2306.08568},
  year={2023}
}
```