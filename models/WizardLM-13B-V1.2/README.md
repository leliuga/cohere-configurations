---
license: llama2
---
This is the **Full-Weight** of WizardLM-13B V1.2 model, this model is trained from **Llama-2 13b**.

## WizardLM: Empowering Large Pre-Trained Language Models to Follow Complex Instructions



<p align="center">
🤗 <a href="https://huggingface.co/WizardLM" target="_blank">HF Repo</a>  •🐱 <a href="https://github.com/nlpxucan/WizardLM" target="_blank">Github Repo</a> • 🐦 <a href="https://twitter.com/WizardLM_AI" target="_blank">Twitter</a> • 📃 <a href="https://arxiv.org/abs/2304.12244" target="_blank">[WizardLM]</a>  • 📃 <a href="https://arxiv.org/abs/2306.08568" target="_blank">[WizardCoder]</a>   • 📃 <a href="https://arxiv.org/abs/2308.09583" target="_blank">[WizardMath]</a>  <br>
</p>
<p align="center">
    👋 Join our <a href="https://discord.gg/VZjjHtWrKs" target="_blank">Discord</a>
</p>

## News

- 🔥🔥🔥[2023/08/26] We released **WizardCoder-Python-34B-V1.0** , which achieves the **73.2 pass@1** and surpasses **GPT4 (2023/03/15)**, **ChatGPT-3.5**, and **Claude2** on the [HumanEval Benchmarks](https://github.com/openai/human-eval). For more details, please refer to [WizardCoder](https://github.com/nlpxucan/WizardLM/tree/main/WizardCoder).
- [2023/06/16] We released **WizardCoder-15B-V1.0** , which surpasses **Claude-Plus (+6.8)**, **Bard (+15.3)** and **InstructCodeT5+ (+22.3)** on the [HumanEval Benchmarks](https://github.com/openai/human-eval). For more details, please refer to [WizardCoder](https://github.com/nlpxucan/WizardLM/tree/main/WizardCoder).

|  Model  |  Checkpoint  | Paper    | HumanEval  |   MBPP | Demo | License |
| ----- |------| ---- |------|-------| ----- |  ----- | 
|  WizardCoder-Python-34B-V1.0  |   🤗 <a href="https://huggingface.co/WizardLM/WizardCoder-Python-34B-V1.0" target="_blank">HF Link</a>   |  📃 <a href="https://arxiv.org/abs/2306.08568" target="_blank">[WizardCoder]</a>  |  73.2   | 61.2 | [Demo](http://47.103.63.15:50085/) |  <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama2</a>  |
|  WizardCoder-15B-V1.0  |   🤗 <a href="https://huggingface.co/WizardLM/WizardCoder-15B-V1.0" target="_blank">HF Link</a>   |  📃 <a href="https://arxiv.org/abs/2306.08568" target="_blank">[WizardCoder]</a>  |  59.8   |50.6 | -- |  <a href="https://huggingface.co/spaces/bigcode/bigcode-model-license-agreement" target="_blank">OpenRAIL-M</a>  |
|  WizardCoder-Python-13B-V1.0  |   🤗 <a href="https://huggingface.co/WizardLM/WizardCoder-Python-13B-V1.0" target="_blank">HF Link</a>   |  📃 <a href="https://arxiv.org/abs/2306.08568" target="_blank">[WizardCoder]</a>  |  64.0   | 55.6 | -- |  <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama2</a>  |
|  WizardCoder-Python-7B-V1.0  |   🤗 <a href="https://huggingface.co/WizardLM/WizardCoder-Python-7B-V1.0" target="_blank">HF Link</a>   |  📃 <a href="https://arxiv.org/abs/2306.08568" target="_blank">[WizardCoder]</a>  |  55.5   | 51.6 | [Demo](http://47.103.63.15:50088/) |  <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama2</a>  |
|  WizardCoder-3B-V1.0  |   🤗 <a href="https://huggingface.co/WizardLM/WizardCoder-3B-V1.0" target="_blank">HF Link</a>   |  📃 <a href="https://arxiv.org/abs/2306.08568" target="_blank">[WizardCoder]</a>  |  34.8   |37.4 | -- |  <a href="https://huggingface.co/spaces/bigcode/bigcode-model-license-agreement" target="_blank">OpenRAIL-M</a>  |
|  WizardCoder-1B-V1.0  |   🤗 <a href="https://huggingface.co/WizardLM/WizardCoder-1B-V1.0" target="_blank">HF Link</a>   |  📃 <a href="https://arxiv.org/abs/2306.08568" target="_blank">[WizardCoder]</a>  |  23.8   |28.6 | -- |  <a href="https://huggingface.co/spaces/bigcode/bigcode-model-license-agreement" target="_blank">OpenRAIL-M</a>  |


- 🔥 [08/11/2023] We release **WizardMath** Models.
- 🔥 Our **WizardMath-70B-V1.0** model slightly outperforms some closed-source LLMs on the GSM8K, including **ChatGPT 3.5**, **Claude Instant 1** and **PaLM 2 540B**.
- 🔥 Our **WizardMath-70B-V1.0** model achieves  **81.6 pass@1** on the [GSM8k Benchmarks](https://github.com/openai/grade-school-math), which is **24.8** points higher than the SOTA open-source LLM.
- 🔥 Our **WizardMath-70B-V1.0** model achieves  **22.7 pass@1** on the [MATH Benchmarks](https://github.com/hendrycks/math), which is **9.2** points higher than the SOTA open-source LLM.


| Model | Checkpoint | Paper  | GSM8k | MATH  |Online Demo| License|
| ----- |------| ---- |------|-------| ----- | ----- |
| WizardMath-70B-V1.0 | 🤗 <a href="https://huggingface.co/WizardLM/WizardMath-70B-V1.0" target="_blank">HF Link</a> |  📃 <a href="https://arxiv.org/abs/2308.09583" target="_blank">[WizardMath]</a>| **81.6**  |  **22.7**	|[Demo](http://47.103.63.15:50083/)| <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama 2  </a> |
| WizardMath-13B-V1.0 | 🤗 <a href="https://huggingface.co/WizardLM/WizardMath-13B-V1.0" target="_blank">HF Link</a> |  📃 <a href="https://arxiv.org/abs/2308.09583" target="_blank">[WizardMath]</a>| **63.9**  |  **14.0** |[Demo](http://47.103.63.15:50082/)| <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama 2 </a> |
| WizardMath-7B-V1.0 | 🤗 <a href="https://huggingface.co/WizardLM/WizardMath-7B-V1.0" target="_blank">HF Link</a>  |  📃 <a href="https://arxiv.org/abs/2308.09583" target="_blank">[WizardMath]</a>| 	 **54.9**  |  **10.7** | [Demo](http://47.103.63.15:50080/)|  <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama 2  </a>|    


<font size=4>
    
| <sup>Model</sup> | <sup>Checkpoint</sup> | <sup>Paper</sup> |<sup>MT-Bench</sup> | <sup>AlpacaEval</sup> | <sup>WizardEval</sup> | <sup>HumanEval</sup>  | <sup>License</sup>|
| ----- |------| ---- |------|-------| ----- | ----- | ----- |
| <sup>WizardLM-13B-V1.2</sup> | <sup>🤗 <a href="https://huggingface.co/WizardLM/WizardLM-13B-V1.2" target="_blank">HF Link</a> </sup>|  | <sup>7.06</sup> | <sup>89.17%</sup>	 | <sup>101.4% </sup>|<sup>36.6  pass@1</sup>|<sup> <a href="https://ai.meta.com/resources/models-and-libraries/llama-downloads/" target="_blank">Llama 2 License </a></sup> |
| <sup>WizardLM-13B-V1.1</sup> |<sup> 🤗 <a href="https://huggingface.co/WizardLM/WizardLM-13B-V1.1" target="_blank">HF Link</a> </sup> |  | <sup>6.76</sup>  |<sup>86.32%</sup>	 | <sup>99.3% </sup> |<sup>25.0  pass@1</sup>| <sup>Non-commercial</sup>|
| <sup>WizardLM-30B-V1.0</sup> | <sup>🤗 <a href="https://huggingface.co/WizardLM/WizardLM-30B-V1.0" target="_blank">HF Link</a></sup>  | | <sup>7.01</sup> |  |  <sup>97.8% </sup> | <sup>37.8  pass@1</sup>| <sup>Non-commercial</sup> |
| <sup>WizardLM-13B-V1.0</sup> | <sup>🤗 <a href="https://huggingface.co/WizardLM/WizardLM-13B-V1.0" target="_blank">HF Link</a> </sup> |  | <sup>6.35</sup> | <sup>75.31%</sup> |  <sup>89.1% </sup> |<sup> 24.0 pass@1 </sup> | <sup>Non-commercial</sup>|
| <sup>WizardLM-7B-V1.0 </sup>|  <sup>🤗 <a href="https://huggingface.co/WizardLM/WizardLM-7B-V1.0" target="_blank">HF Link</a> </sup> |<sup> 📃 <a href="https://arxiv.org/abs/2304.12244" target="_blank">[WizardLM]</a> </sup>|  |  |  <sup>78.0% </sup> |<sup>19.1 pass@1 </sup>|<sup> Non-commercial</sup>|
</font>

**Repository**: https://github.com/nlpxucan/WizardLM

**Twitter**:


- 🔥🔥🔥 [7/25/2023] We released **WizardLM V1.2** models. The **WizardLM-13B-V1.2** is here ([Demo_13B-V1.2](https://b7a19878988c8c73.gradio.app), [Demo_13B-V1.2_bak-1](https://d0a37a76e0ac4b52.gradio.app/), [Full Model Weight](https://huggingface.co/WizardLM/WizardLM-13B-V1.2)). Please checkout the [paper](https://arxiv.org/abs/2304.12244).
- 🔥🔥🔥 [7/25/2023] The **WizardLM-13B-V1.2** achieves **7.06** on [MT-Bench Leaderboard](https://chat.lmsys.org/?leaderboard), **89.17%** on [AlpacaEval Leaderboard](https://tatsu-lab.github.io/alpaca_eval/), and **101.4%** on [WizardLM Eval](https://github.com/nlpxucan/WizardLM/blob/main/WizardLM/data/WizardLM_testset.jsonl). (Note: MT-Bench and AlpacaEval are all self-test, will push update and request review. All tests are completed under their official settings.)

❗<b>Note for model system prompts usage:</b>


<b>WizardLM</b>  adopts the prompt format from <b>Vicuna</b> and supports **multi-turn** conversation. The prompt should be as following:

```
A chat between a curious user and an artificial intelligence assistant. The assistant gives helpful, detailed, and polite answers to the user's questions. USER: Hi ASSISTANT: Hello.</s>USER: Who are you? ASSISTANT: I am WizardLM.</s>......
```

## Inference WizardLM Demo Script

We provide the inference WizardLM demo code [here](https://github.com/nlpxucan/WizardLM/tree/main/demo).

Please cite the paper if you use the data or code from WizardLM.

```
@article{xu2023wizardlm,
  title={Wizardlm: Empowering large language models to follow complex instructions},
  author={Xu, Can and Sun, Qingfeng and Zheng, Kai and Geng, Xiubo and Zhao, Pu and Feng, Jiazhan and Tao, Chongyang and Jiang, Daxin},
  journal={arXiv preprint arXiv:2304.12244},
  year={2023}
}
```

❗<b>To commen concern about dataset:</b>

Recently, there have been clear changes in the open-source policy and regulations of our overall organization's code, data, and models. 


Despite this, we have still worked hard to obtain opening the weights of the model first, but the data involves stricter auditing and is in review with our legal team .

Our researchers have no authority to publicly release them without authorization.

Thank you for your understanding.