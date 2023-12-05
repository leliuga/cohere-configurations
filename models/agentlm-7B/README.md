---
datasets: 
- THUDM/AgentInstruct
---

## AgentLM-7B

<p align="center">
  🤗 <a href="https://huggingface.co/datasets/THUDM/AgentInstruct" target="_blank">[Dataset] </a> • 💻 <a href="https://github.com/THUDM/AgentTuning" target="_blank">[Github Repo]</a> • 📌 <a href="https://THUDM.github.io/AgentTuning/" target="_blank">[Project Page]</a> • 📃 <a href="https://arxiv.org/abs/2310.12823" target="_blank">[Paper]</a> 
</p>

**AgentTuning** represents the very first attempt to instruction-tune LLMs using interaction trajectories across multiple agent tasks. Evaluation results indicate that AgentTuning enables the agent capabilities of LLMs with robust generalization on unseen agent tasks while remaining good on general language abilities. We have open-sourced the AgentInstruct dataset and AgentLM.

## Models

**AgentLM** models are produced by mixed training on AgentInstruct dataset and ShareGPT dataset from Llama-2-chat models. 

The models follow the conversation format of [Llama-2-chat](https://huggingface.co/blog/llama2#how-to-prompt-llama-2), with system prompt fixed as 

```
You are a helpful, respectful and honest assistant.
```

7B, 13B, and 70B models are available on Huggingface model hub.

|Model|Huggingface Repo|
|---|---|
|AgentLM-7B| [🤗Huggingface Repo](https://huggingface.co/THUDM/agentlm-7b) |
|AgentLM-13B| [🤗Huggingface Repo](https://huggingface.co/THUDM/agentlm-13b) |
|AgentLM-70B| [🤗Huggingface Repo](https://huggingface.co/THUDM/agentlm-70b) |

## Citation

If you find our work useful, please consider citing AgentTuning:

```
@misc{zeng2023agenttuning,
      title={AgentTuning: Enabling Generalized Agent Abilities for LLMs}, 
      author={Aohan Zeng and Mingdao Liu and Rui Lu and Bowen Wang and Xiao Liu and Yuxiao Dong and Jie Tang},
      year={2023},
      eprint={2310.12823},
      archivePrefix={arXiv},
      primaryClass={cs.CL}
}
```
