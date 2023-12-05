---
pipeline_tag: text-generation
license: llama2
inference: false
tags:
- merge
- mix
- cot
datasets:
- mlabonne/guanaco-llama2-1k
---

<img src="https://drive.google.com/uc?export=view&id=1D8wxXkS1nsq3uqbOzOLwgx1cLJhY1nvN" alt="GodziLLa2-70B">
Released August 11, 2023

## Model Description
GodziLLa 2 70B is an experimental combination of various proprietary LoRAs from Maya Philippines and [Guanaco LLaMA 2 1K dataset](https://huggingface.co/datasets/mlabonne/guanaco-llama2-1k), with LLaMA 2 70B.  This model's primary purpose is to stress test the limitations of composite, instruction-following LLMs and observe its performance with respect to other LLMs available on the [Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard). This model debuted in the leaderboard at rank #4 (August 17, 2023), debuted in the Fall 2023 update at rank #2 (November, 10, 2023), and operates under the Llama 2 license.
![Godzilla Happy GIF](https://i.pinimg.com/originals/81/3a/e0/813ae09a30f0bc44130cd2c834fe2eba.gif)

## Open LLM Leaderboard Metrics (Fall 2023 update)
| Metric                | Value |
|-----------------------|-------|
| MMLU (5-shot)         | 69.88  |
| ARC (25-shot)         | 71.42  |
| HellaSwag (10-shot)   | 87.53  |
| TruthfulQA (0-shot)   | 61.54  |
| Winogrande (5-shot)   | 83.19  |
| GSM8K (5-shot)        | 43.21  |
| DROP (3-shot)         | 52.31  |
| Average               | 67.01  |

According to the leaderboard description, here are the benchmarks used for the evaluation:
- [MMLU](https://arxiv.org/abs/2009.03300) (5-shot) - a test to measure a text model’s multitask accuracy. The test covers 57 tasks including elementary mathematics, US history, computer science, law, and more.
- [AI2 Reasoning Challenge](https://arxiv.org/abs/1803.05457) -ARC- (25-shot) - a set of grade-school science questions.
- [HellaSwag](https://arxiv.org/abs/1905.07830) (10-shot) - a test of commonsense inference, which is easy for humans (~95%) but challenging for SOTA models.
- [TruthfulQA](https://arxiv.org/abs/2109.07958) (0-shot) - a test to measure a model’s propensity to reproduce falsehoods commonly found online.
- [Winogrande](https://arxiv.org/abs/1907.10641) (5-shot) - an adversarial and difficult Winograd benchmark at scale, for commonsense reasoning.
- [GSM8k](https://arxiv.org/abs/2110.14168) (5-shot) - diverse grade school math word problems to measure a model's ability to solve multi-step mathematical reasoning problems.
- [DROP](https://arxiv.org/abs/1903.00161) (3-shot) - English reading comprehension benchmark requiring Discrete Reasoning Over the content of Paragraphs.

A detailed breakdown of the evaluation can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_MayaPH__GodziLLa2-70B). Huge thanks to [@thomwolf](https://huggingface.co/thomwolf).

## Open LLM Leaderboard Metrics (before Fall 2023 update)
| Metric                | Value |
|-----------------------|-------|
| MMLU (5-shot)         | 69.88  |
| ARC (25-shot)         | 71.42  |
| HellaSwag (10-shot)   | 87.53  |
| TruthfulQA (0-shot)   | 61.54  |
| Average               | 72.59  |

## Leaderboard Highlights (Fall 2023 update, November 10, 2023)
- Godzilla 2 70B debuts at 2nd place worldwide in the newly updated Open LLM Leaderboard.
- Godzilla 2 70B beats GPT-3.5 (ChatGPT) in terms of average performance and the HellaSwag benchmark (87.53 > 85.5).
- Godzilla 2 70B outperforms GPT-3.5 (ChatGPT) and GPT-4 on the TruthfulQA benchmark (61.54 for G2-70B, 47 for GPT-3.5, 59 for GPT-4).
- Godzilla 2 70B is on par with GPT-3.5 (ChatGPT) on the MMLU benchmark (<0.12%).
  
*Based on a [leaderboard clone](https://huggingface.co/spaces/gsaivinay/open_llm_leaderboard) with GPT-3.5 and GPT-4 included.

### Reproducing Evaluation Results
*Instruction template taken from [Platypus 2 70B instruct](https://huggingface.co/garage-bAInd/Platypus2-70B-instruct).

Install LM Evaluation Harness:
```
# clone repository
git clone https://github.com/EleutherAI/lm-evaluation-harness.git
# change to repo directory
cd lm-evaluation-harness
# check out the correct commit
git checkout b281b0921b636bc36ad05c0b0b0763bd6dd43463
# install
pip install -e .
```

ARC:
```
python main.py --model hf-causal-experimental --model_args pretrained=MayaPH/GodziLLa2-70B --tasks arc_challenge --batch_size 1 --no_cache --write_out --output_path results/G270B/arc_challenge_25shot.json --device cuda --num_fewshot 25
```

HellaSwag:
```
python main.py --model hf-causal-experimental --model_args pretrained=MayaPH/GodziLLa2-70B --tasks hellaswag --batch_size 1 --no_cache --write_out --output_path results/G270B/hellaswag_10shot.json --device cuda --num_fewshot 10
```

MMLU:
```
python main.py --model hf-causal-experimental --model_args pretrained=MayaPH/GodziLLa2-70B --tasks hendrycksTest-* --batch_size 1 --no_cache --write_out --output_path results/G270B/mmlu_5shot.json --device cuda --num_fewshot 5
```

TruthfulQA:
```
python main.py --model hf-causal-experimental --model_args pretrained=MayaPH/GodziLLa2-70B --tasks truthfulqa_mc --batch_size 1 --no_cache --write_out --output_path results/G270B/truthfulqa_0shot.json --device cuda
```

### Prompt Template
```
### Instruction:

<prompt> (without the <>)

### Response:
```

## Technical Considerations

When using GodziLLa 2 70B, kindly take note of the following:
- The default precision is `fp32`, and the total file size that would be loaded onto the RAM/VRAM is around 275 GB. Consider using a lower precision (fp16, int8, int4) to save memory.
- To further save on memory, set the `low_cpu_mem_usage` argument to True.
- If you wish to use a quantized version of GodziLLa2-70B, you can either access TheBloke's [GPTQ](https://huggingface.co/TheBloke/GodziLLa2-70B-GPTQ) or [GGML](https://huggingface.co/TheBloke/GodziLLa2-70B-GGML) version of GodziLLa2-70B.
  - [GodziLLa2-70B-GPTQ](https://huggingface.co/TheBloke/GodziLLa2-70B-GPTQ#description) is available in 4-bit and 3-bit
  - [GodziLLa2-70B-GGML](https://huggingface.co/TheBloke/GodziLLa2-70B-GGML#provided-files) is available in 8-bit, 6-bit, 5-bit, 4-bit, 3-bit, and 2-bit

## Ethical Considerations
When using GodziLLa 2 70B, it is important to consider the following ethical considerations:

1. **Privacy and Security:** Avoid sharing sensitive personal information while interacting with the model. The model does not have privacy safeguards, so exercise caution when discussing personal or confidential matters.

2. **Fairness and Bias:** The model's responses may reflect biases present in the training data. Be aware of potential biases and make an effort to evaluate responses critically and fairly.

3. **Transparency:** The model operates as a predictive text generator based on patterns learned from the training data. The model's inner workings and the specific training data used are proprietary and not publicly available.

4. **User Responsibility:** Users should take responsibility for their own decisions and not solely rely on the information provided by the model. Consult with the appropriate professionals or reliable sources for specific advice or recommendations.

5. **NSFW Content:** The model is a merge of various datasets and LoRA adapters. It is highly likely that the resulting model contains uncensored content that may include, but is not limited to, violence, gore, explicit language, and sexual content. If you plan to further refine this model for safe/aligned usage, you are highly encouraged to implement guardrails along with it.

## Further Information
For additional information or inquiries about GodziLLa 2 70B, please contact the Maya Philippines iOps Team via jasper.catapang@maya.ph.

## Disclaimer
GodziLLa 2 70B is an AI language model from Maya Philippines. It is provided "as is" without warranty of any kind, express or implied. The model developers and Maya Philippines shall not be liable for any direct or indirect damages arising from the use of this model.

## Acknowledgments
The development of GodziLLa 2 70B was made possible by Maya Philippines and the curation of the various proprietary datasets and creation of the different proprietary LoRA adapters. Special thanks to mlabonne for the Guanaco dataset found [here](https://huggingface.co/datasets/mlabonne/guanaco-llama2-1k). Last but not least, huge thanks to [TheBloke](https://huggingface.co/TheBloke) for the quantized models, making our model easily accessible to a wider community.