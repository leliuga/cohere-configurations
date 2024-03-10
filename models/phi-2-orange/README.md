---
license: mit
datasets:
- Open-Orca/SlimOrca-Dedup
- migtissera/Synthia-v1.3
- LDJnr/Verified-Camel
- LDJnr/Pure-Dove
- LDJnr/Capybara
- meta-math/MetaMathQA
- Intel/orca_dpo_pairs
- argilla/ultrafeedback-binarized-preferences-cleaned
---
![Phi-2 Orange](https://huggingface.co/rhysjones/phi-2-orange/resolve/main/phi-2-orange.jpg)

# Phi-2 Orange

A two-step finetune of Phi-2, with a bit of zest.

There is an updated model at [rhysjones/phi-2-orange-v2](https://huggingface.co/rhysjones/phi-2-orange-v2) which has higher evals, if you wish to test.

# Training details

A first finetune using a collection of broad training data:

- [Open-Orca/SlimOrca-Dedup](https://huggingface.co/datasets/Open-Orca/SlimOrca-Dedup)
- [migtissera/Synthia-v1.3](https://huggingface.co/datasets/migtissera/Synthia-v1.3)
- [LDJnr/Verified-Camel](https://huggingface.co/datasets/LDJnr/Verified-Camel)
- [LDJnr/Pure-Dove](https://huggingface.co/datasets/LDJnr/Pure-Dove)
- [LDJnr/Capybara](https://huggingface.co/datasets/LDJnr/Capybara)
- [meta-math/MetaMathQA](https://huggingface.co/datasets/meta-math/MetaMathQA)

And then a DPO finetune using:

- [Intel/orca_dpo_pairs](https://huggingface.co/datasets/Intel/orca_dpo_pairs)
- [argilla/ultrafeedback-binarized-preferences-cleaned](https://huggingface.co/datasets/argilla/ultrafeedback-binarized-preferences-cleaned)


# Run within Ollama

If you're using [Ollama](https://ollama.ai), you can download and run using:
```
ollama run rhysjones/phi-2-orange
```

# Prompt Format

Phi-2 Orange uses ChatML as the prompt format, with or without the system instruction.

To prompt with a system instruction (use whatever system prompt you like):

```
<|im_start|>system
You are a helpful assistant for Python which outputs in Markdown format.<|im_end|>
<|im_start|>user
Write a function to calculate the Fibonacci sequence<|im_end|>
<|im_start|>assistant

```

You can also omit the system prompt if you wish:

```
<|im_start|>user
Why is the sky blue?<|im_end|>
<|im_start|>assistant

```

# Evaluations
Evaluations done using mlabonne's usefull [Colab notebook llm-autoeval](https://github.com/mlabonne/llm-autoeval).
Also check out the alternative leaderboard at [Yet_Another_LLM_Leaderboard](https://huggingface.co/spaces/mlabonne/Yet_Another_LLM_Leaderboard)
|                             Model                              |AGIEval|GPT4All|TruthfulQA|Bigbench|Average|
|----------------------------------------------------------------|------:|------:|---------:|-------:|------:|
|[phi-2-orange](https://huggingface.co/rhysjones/phi-2-orange)|  **33.37**|  71.33|      49.87|   **37.3**|  **47.97**|
|[phi-2-dpo](https://huggingface.co/lxuechen/phi-2-dpo)|  30.39|  **71.68**|     **50.75**|    34.9|  46.93|
|[dolphin-2_6-phi-2](https://huggingface.co/cognitivecomputations/dolphin-2_6-phi-2)|  33.12|  69.85|     47.39|    37.2|  46.89|
|[phi-2](https://huggingface.co/microsoft/phi-2)|  27.98|   70.8|     44.43|   35.21|  44.61|