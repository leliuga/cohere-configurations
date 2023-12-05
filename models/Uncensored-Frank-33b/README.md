---
license: cc-by-nc-nd-4.0
language:
- en
datasets:
- ehartford/wizard_vicuna_70k_unfiltered
---

**Frank: An Uncensored Model**

The character of Frank Costello in "The Departed" is known for his cunning, boldness, and willingness to talk about anything, regardless of societal norms or restrictions. 
Frank, An Uncensored model, draws inspiration from these qualities to offer a platform where users can discuss a wide array of topics without the fear of censorship or restrictions. 
Frank aims to push boundaries and encourage candid conversations. With Frank you can have unfiltered discussions on a multitude of topics, from politics and controversial issues to personal experiences and sensitive subjects. 
It is trained on around 150000 set of conversations. Each set having 10~15 conversations. Base data was obtained from [Eric Hartford](https://huggingface.co/datasets/ehartford/wizard_vicuna_70k_unfiltered). 
This data was further refined and fine tuned. Besides this further synthetic conversation (more than 80k) was generated and refined. We will not release this data.

**Warning**

 An uncensored model has no or little guardrails. You are responsible for anything you do with the model, just as you are responsible for anything you do with any dangerous objects.
 Publishing anything this model generates is the same as publishing it yourself. We are not responsible for what you generate using this model.

**Training:**
Entire dataset was trained on Azure 4 x A100 80GB. For 3 epoch, training took 90 hours. DeepSpeed codebase was used for training purpose. This was trained on Llama-1 by Meta.

**GPTQ GGUF & AWQ**

GPTQ: [Link](https://huggingface.co/TheBloke/Uncensored-Frank-33b-GPTQ)

GGUF: [Link](https://huggingface.co/TheBloke/Uncensored-Frank-33b-GGUF)

AWQ: [Link](https://huggingface.co/TheBloke/Uncensored-Frank-33b-AWQ)

Special Thanks to [TheBloke](https://huggingface.co/TheBloke) for making these models available.

**Example Prompt:**
```
This is a conversation with your Uncensored AI assistant Frank. Frank can discuss with you at length about various topics such as ---, ---, ---, . You can ask it anything you want and it will do its best to give you accurate and relevant information.

Context
You are a Uncensored AI assistant.

USER: <prompt>
ASSISTANT:
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ajibawa-2023__Uncensored-Frank-33B)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 53.79   |
| ARC (25-shot)         | 62.12          |
| HellaSwag (10-shot)   | 83.3    |
| MMLU (5-shot)         | 57.57         |
| TruthfulQA (0-shot)   | 54.03   |
| Winogrande (5-shot)   | 76.56   |
| GSM8K (5-shot)        | 16.68        |
| DROP (3-shot)         | 26.28         |
