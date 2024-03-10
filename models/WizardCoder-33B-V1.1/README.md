---
metrics:
- code_eval
library_name: transformers
tags:
- code
model-index:
- name: WizardCoder
  results:
  - task:
      type: text-generation
    dataset:
      type: openai_humaneval
      name: HumanEval
    metrics:
    - name: pass@1
      type: pass@1
      value: 0.799
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

## How to Make the Training Data?

Apply our [Code Evol-Instruct](https://wizardlm.github.io/WizardCoder/) on [Code-Aplaca data](https://huggingface.co/datasets/sahil2801/CodeAlpaca-20k).


## ❗ Data Contamination Check:

Before model training, we carefully and rigorously checked all the training data, and used multiple deduplication methods to verify and prevent data leakage on HumanEval and MBPP test set. 

🔥 
❗<b>Note for model system prompts usage:</b>

Please use **the same systems prompts strictly** with us, and we do not guarantee the accuracy of the **quantified versions**.

**Default version:**

```
"Below is an instruction that describes a task. Write a response that appropriately completes the request.\n\n### Instruction:\n{instruction}\n\n### Response:"
```


## How to Reproduce the Performance of WizardCoder-33B-V1.1

We provide all codes [here](https://github.com/nlpxucan/WizardLM/tree/main/WizardCoder/src).

We also provide all generated [results](https://github.com/nlpxucan/WizardLM/blob/main/WizardCoder/data/humaneval_mbpp_wizardcoder33b_v1.1_results.zip).

```
transformers==4.36.2
vllm==0.2.5
```

(1) HumanEval and HumanEval-Plus

- Step 1

Code Generation (w/o accelerate)
```bash
model="WizardLM/WizardCoder-33B-V1.1"
temp=0.0
max_len=2048
pred_num=1
num_seqs_per_iter=1

output_path=preds/T${temp}_N${pred_num}_WizardCoder-33B-V1.1_Greedy_Decode

mkdir -p ${output_path}
echo 'Output path: '$output_path
echo 'Model to eval: '$model

# 164 problems, 21 per GPU if GPU=8
index=0
gpu_num=8
for ((i = 0; i < $gpu_num; i++)); do
  start_index=$((i * 21))
  end_index=$(((i + 1) * 21))

  gpu=$((i))
  echo 'Running process #' ${i} 'from' $start_index 'to' $end_index 'on GPU' ${gpu}
  ((index++))
  (
    CUDA_VISIBLE_DEVICES=$gpu python humaneval_gen.py --model ${model} \
      --start_index ${start_index} --end_index ${end_index} --temperature ${temp} \
      --num_seqs_per_iter ${num_seqs_per_iter} --N ${pred_num} --max_len ${max_len} --output_path ${output_path} --greedy_decode
  ) &
  if (($index % $gpu_num == 0)); then wait; fi
done
```

Code Generation (w/ vllm accelerate)
```bash
model="WizardLM/WizardCoder-33B-V1.1"
temp=0.0
max_len=2048
pred_num=1
num_seqs_per_iter=1

output_path=preds/T${temp}_N${pred_num}_WizardCoder-33B-V1.1_Greedy_Decode_vllm

mkdir -p ${output_path}
echo 'Output path: '$output_path
echo 'Model to eval: '$model

CUDA_VISIBLE_DEVICES=0,1,2,3 python humaneval_gen_vllm.py --model ${model} \
    --start_index 0 --end_index 164 --temperature ${temp} \
    --num_seqs_per_iter ${num_seqs_per_iter} --N ${pred_num} --max_len ${max_len} --output_path ${output_path} --num_gpus 4 --overwrite
```

- Step 2: Get the score

Install [Eval-Plus](https://github.com/evalplus/evalplus) benchmark.
```bash
git clone https://github.com/evalplus/evalplus.git
cd evalplus
export PYTHONPATH=$PYTHONPATH:$(pwd)
pip install -r requirements.txt
```
Get HumanEval and HumanEval-Plus scores.
```bash
output_path=preds/T0.0_N1_WizardCoder-33B-V1.1_Greedy_Decode

echo 'Output path: '$output_path
python process_humaneval.py --path ${output_path} --out_path ${output_path}.jsonl --add_prompt

evalplus.evaluate --dataset humaneval --samples ${output_path}.jsonl
```

(2) MBPP and MBPP-Plus

The preprocessed questions are provided in [mbppplus.json](https://github.com/nlpxucan/WizardLM/blob/main/WizardCoder/data/mbppplus.json).

- Step 1

Code Generation (w/o accelerate)
```bash
model="WizardLM/WizardCoder-33B-V1.1"
temp=0.0
max_len=2048
pred_num=1
num_seqs_per_iter=1

output_path=preds/MBPP_T${temp}_N${pred_num}_WizardCoder-33B-V1.1_Greedy_Decode

mkdir -p ${output_path}
echo 'Output path: '$output_path
echo 'Model to eval: '$model

# 399 problems, 50 per GPU if GPU=8
index=0
gpu_num=8
for ((i = 0; i < $gpu_num; i++)); do
  start_index=$((i * 50))
  end_index=$(((i + 1) * 50))

  gpu=$((i))
  echo 'Running process #' ${i} 'from' $start_index 'to' $end_index 'on GPU' ${gpu}
  ((index++))
  (
    CUDA_VISIBLE_DEVICES=$gpu python mbppplus_gen.py --model ${model} \
      --start_index ${start_index} --end_index ${end_index} --temperature ${temp} \
      --num_seqs_per_iter ${num_seqs_per_iter} --N ${pred_num} --max_len ${max_len} --output_path ${output_path} --mbpp_path "mbppplus.json" --greedy_decode
  ) &
  if (($index % $gpu_num == 0)); then wait; fi
done
```

Code Generation (w/ vllm accelerate)
```bash
model="WizardLM/WizardCoder-33B-V1.1"
temp=0.0
max_len=2048
pred_num=1
num_seqs_per_iter=1

output_path=preds/MBPP_T${temp}_N${pred_num}_WizardCoder-33B-V1.1_Greedy_Decode_vllm

mkdir -p ${output_path}
echo 'Output path: '$output_path
echo 'Model to eval: '$model

CUDA_VISIBLE_DEVICES=0,1,2,3 python mbppplus_gen_vllm.py --model ${model} \
    --start_index ${start_index} --end_index ${end_index} --temperature ${temp} \
    --num_seqs_per_iter ${num_seqs_per_iter} --N ${pred_num} --max_len ${max_len} --output_path ${output_path} --mbpp_path "mbppplus.json" --num_gpus 4
```

- Step 2: Get the score

Install [Eval-Plus](https://github.com/evalplus/evalplus) benchmark.
```bash
git clone https://github.com/evalplus/evalplus.git
cd evalplus
export PYTHONPATH=$PYTHONPATH:$(pwd)
pip install -r requirements.txt
```
Get HumanEval and HumanEval-Plus scores.
```bash
output_path=preds/MBPP_T0.0_N1_WizardCoder-33B-V1.1_Greedy_Decode

echo 'Output path: '$output_path
python mbppplus_process_preds.py --path ${output_path} --out_path ${output_path}.jsonl --add_prompt

evalplus.evaluate --dataset mbpp --samples ${output_path}.jsonl
```


## Citation

Please cite the repo if you use the data, method or code in this repo.

```
@article{luo2023wizardcoder,
  title={WizardCoder: Empowering Code Large Language Models with Evol-Instruct},
  author={Luo, Ziyang and Xu, Can and Zhao, Pu and Sun, Qingfeng and Geng, Xiubo and Hu, Wenxiang and Tao, Chongyang and Ma, Jing and Lin, Qingwei and Jiang, Daxin},
  journal={arXiv preprint arXiv:2306.08568},
  year={2023}
}
```