---
license: llama2
---

# OpenChat: Advancing Open-source Language Models with Imperfect Data</h1>

<div align="center">
  <img src="https://raw.githubusercontent.com/imoneoi/openchat/master/assets/logo_new.png" style="width: 65%">
</div>

OpenChat is a collection of open-source language models, optimized and fine-tuned with a strategy inspired by offline reinforcement learning. We use approximately 80k ShareGPT conversations, a conditioning strategy, and weighted loss to deliver outstanding performance, despite our simple approach. Our ultimate goal is to develop a high-performance, commercially available, open-source large language model, and we are continuously making strides towards this vision.

**🤖 Ranked #1 among all open-source models on [AgentBench](https://github.com/THUDM/AgentBench)**

**🔥 Ranked #1 among 13B open-source models | 89.5% win-rate on [AlpacaEval](https://tatsu-lab.github.io/alpaca_eval/) | 7.19 score on [MT-bench](https://chat.lmsys.org/?leaderboard)**

**🕒 Exceptionally efficient padding-free fine-tuning, only requires 15 hours on 8xA100 80G**

**💲 FREE for commercial use under [Llama 2 Community License](https://ai.meta.com/resources/models-and-libraries/llama-downloads/)**

[![DOI](https://zenodo.org/badge/645397533.svg)](https://zenodo.org/badge/latestdoi/645397533)

## <a id="models"></a> Usage

To use these models, we highly recommend installing the OpenChat package by following the [installation guide](https://github.com/imoneoi/openchat/#installation) and using the OpenChat OpenAI-compatible API server by running the serving command from the table below. The server is optimized for high-throughput deployment using [vLLM](https://github.com/vllm-project/vllm) and can run on a GPU with at least 48GB RAM or two consumer GPUs with tensor parallelism. To enable tensor parallelism, append `--tensor-parallel-size 2` to the serving command.

When started, the server listens at `localhost:18888` for requests and is compatible with the [OpenAI ChatCompletion API specifications](https://platform.openai.com/docs/api-reference/chat). See the example request below for reference. Additionally, you can access the [OpenChat Web UI](#web-ui) for a user-friendly experience.

To deploy the server as an online service, use `--api-keys sk-KEY1 sk-KEY2 ...` to specify allowed API keys and `--disable-log-requests --disable-log-stats --log-file openchat.log` for logging only to a file. We recommend using a [HTTPS gateway](https://fastapi.tiangolo.com/es/deployment/concepts/#security-https) in front of the server for security purposes.

<details>
  <summary>Example request (click to expand)</summary>

```bash
curl http://localhost:18888/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "openchat_v3.2",
    "messages": [{"role": "user", "content": "You are a large language model named OpenChat. Write a poem to describe yourself"}]
  }'
```

</details>

| Model        | Size | Context | Weights                                                      | Serving                                                                                                                                                                      |
|--------------|------|---------|--------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| OpenChat 3.2 SUPER | 13B  | 4096    | [Huggingface](https://huggingface.co/openchat/openchat_v3.2_super) | `python -m ochat.serving.openai_api_server --model-type openchat_v3.2 --model openchat/openchat_v3.2_super --engine-use-ray --worker-use-ray --max-num-batched-tokens 5120`        |

For inference with Huggingface Transformers (slow and not recommended), follow the conversation template provided below:

<details>
  <summary>Conversation templates (click to expand)</summary>

```python
# Single-turn V3.2 (SUPER)
tokenize("GPT4 User: Hello<|end_of_turn|>GPT4 Assistant:")
# Result: [1, 402, 7982, 29946, 4911, 29901, 15043, 32000, 402, 7982, 29946, 4007, 22137, 29901]

# Multi-turn V3.2 (SUPER)
tokenize("GPT4 User: Hello<|end_of_turn|>GPT4 Assistant: Hi<|end_of_turn|>GPT4 User: How are you today?<|end_of_turn|>GPT4 Assistant:")
# Result: [1, 402, 7982, 29946, 4911, 29901, 15043, 32000, 402, 7982, 29946, 4007, 22137, 29901, 6324, 32000, 402, 7982, 29946, 4911, 29901, 1128, 526, 366, 9826, 29973, 32000, 402, 7982, 29946, 4007, 22137, 29901]
```

</details>

## <a id="benchmarks"></a> Benchmarks

We have evaluated our models using the two most popular evaluation benchmarks **, including AlpacaEval and MT-bench. Here we list the top models with our released versions, sorted by model size in descending order. The full version can be found on the [MT-bench](https://chat.lmsys.org/?leaderboard) and [AlpacaEval](https://tatsu-lab.github.io/alpaca_eval/) leaderboards.

To ensure consistency, we used the same routine as ChatGPT / GPT-4 to run these benchmarks. We started the OpenAI API-compatible server and set the `openai.api_base` to `http://localhost:18888/v1` in the benchmark program.

| **Model**                        | **Size** | **Context** | **Dataset Size** | **💲Free** | **AlpacaEval (win rate %)** | **MT-bench (win rate adjusted %)** | **MT-bench (score)** |
|----------------------------------|----------|-------------|------------------|-----------|-----------------------------|------------------------------------|----------------------|
|                                  |          |             |                  |           | **v.s. text-davinci-003**   | **v.s. ChatGPT**                   |                      |
| GPT-4                            | 1.8T*    | 8K          |                  | ❌         | 95.3                        | 82.5                               | 8.99                 |
| ChatGPT                          | 175B*    | 4K          |                  | ❌         | 89.4                        | 50.0                               | 7.94                 |
| Llama-2-70B-Chat                 | 70B      | 4K          | 2.9M             | ✅         | 92.7                        | 60.0                               | 6.86                 |
| **OpenChat 3.2 SUPER**           | **13B**  | **4K**      | **80K**          | ✅         | **89.5**                    | **57.5**                           | **7.19**             |
| Llama-2-13B-Chat                 | 13B      | 4K          | 2.9M             | ✅         | 81.1                        | 55.3                               | 6.65                 |
| WizardLM 1.2                     | 13B      | 4K          | 196K             | ✅         | 89.2                        | 53.1                               | 7.05                 |
| Vicuna 1.5                       | 13B      | 2K          | 125K             | ✅         | 78.8                        | 37.2                               | 6.57                 |

*: Estimated model size

**: The benchmark metrics represent a quantified measure of a subset of the model's capabilities. A win-rate greater than 50% does not necessarily indicate that the model is better than ChatGPT in all scenarios or for all use cases. It is essential to consider the specific tasks or applications for which the model was evaluated and compare the results accordingly.

## Limitations

**Foundation Model Limitations**
Despite its advanced capabilities, OpenChat is still bound by the limitations inherent in its foundation models. These limitations may impact the model's performance in areas such as:

 - Complex reasoning
 - Mathematical and arithmetic tasks
 - Programming and coding challenges

**Hallucination of Non-existent Information**
OpenChat may sometimes generate information that does not exist or is not accurate, also known as "hallucination". Users should be aware of this possibility and verify any critical information obtained from the model.

## License

Our OpenChat V3 models are licensed under the [Llama 2 Community License](https://ai.meta.com/resources/models-and-libraries/llama-downloads/).

```
@article{wang2023openchat,
  title={OpenChat: Advancing Open-source Language Models with Mixed-Quality Data},
  author={Wang, Guan and Cheng, Sijie and Zhan, Xianyuan and Li, Xiangang and Song, Sen and Liu, Yang},
  journal={arXiv preprint arXiv:2309.11235},
  year={2023}
}
```
