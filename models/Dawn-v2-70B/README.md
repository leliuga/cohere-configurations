---
license: cc-by-nc-4.0
tags:
- not-for-all-audiences
- nsfw
---

<center>[<a href="https://huggingface.co/Undi95/Dawn-v2-70B">fp16</a> - <a href="https://huggingface.co/Undi95/Dawn-v2-70B-GGUF">gguf</a> - exl2 : <a href="https://huggingface.co/Undi95/Dawn-v2-70B-2.55bpw-h6-exl2">2.55bpw</a>]</center>
</br>
<div style="width: 100%;">
    <img src="https://cdn-uploads.huggingface.co/production/uploads/63ab1241ad514ca8d1430003/Cxcfqi4WdtXCNLnaIqSRB.png" style="width: 75%; min-width: 200px; display: block; margin: auto;">
</div>

<!-- description start -->
## Description

This repo contains fp16 files of Dawn-70B, a merge I have done with the new [layer shuffle](https://github.com/cg123/mergekit/blob/main/mergekit/scripts/layershuffle.py) method from mergekit.

[UtopiaXL](https://huggingface.co/Undi95/UtopiaXL-13B) was a huge success for me, I really liked it, so I took the same path to do this 70B: A good base, some psychologic data, some medical data, a little bit of this, of that, and LimaRP at the end as always.

NOTE: This repo contain the file  [measurement.json](https://huggingface.co/Undi95/Dawn-v2-70B/blob/main/measurement.json) needed to do your own exl2 quant (I use [wikitext](https://huggingface.co/datasets/wikitext/resolve/refs%2Fconvert%2Fparquet/wikitext-2-raw-v1/train/0000.parquet)).
<!-- description end -->
<!-- description start -->
## Models and loras used

- [Sao10K/Euryale-1.3-L2-70B](https://huggingface.co/Sao10K/Euryale-1.3-L2-70B)
- [Xwin-LM/Xwin-LM-70B-V0.1](https://huggingface.co/Xwin-LM/Xwin-LM-70B-V0.1)
- [ehartford/Samantha-1.11-70b](https://huggingface.co/ehartford/Samantha-1.11-70b)
- [NousResearch/Nous-Hermes-Llama2-70b](https://huggingface.co/NousResearch/Nous-Hermes-Llama2-70b)
- [augtoma/qCammel-70-x](https://huggingface.co/augtoma/qCammel-70-x)
- [jondurbin/airoboros-l2-c70b-3.1.2](https://huggingface.co/jondurbin/airoboros-l2-c70b-3.1.2)
- [fangloveskari/ORCA_LLaMA_70B_QLoRA](https://huggingface.co/fangloveskari/ORCA_LLaMA_70B_QLoRA)
- [Doctor-Shotgun/limarpv3-llama2-70b-qlora](https://huggingface.co/Doctor-Shotgun/limarpv3-llama2-70b-qlora)

<!-- description end -->
## The sauce
```
!mergekit-layershuffle ./Dawn-v2-70B \
  --model Sao10K/Euryale-1.3-L2-70B --weight 0.3 \
  --model Xwin-LM/Xwin-LM-70B-V0.1 --weight 0.2 \
  --model ehartford/Samantha-1.11-70b --weight 0.1 \
  --model NousResearch/Nous-Hermes-Llama2-70b --weight 0.05 \
  --model augtoma/qCammel-70-x --weight 0.05 \
  --model jondurbin/airoboros-l2-c70b-3.1.2 --weight 0.2 \
  --model fangloveskari/ORCA_LLaMA_70B_QLoRA --weight 0.1 \
  --write-yaml Dawn-v2-70B.yaml

=========================

merge_method: passthrough
slices:
- sources:
  - layer_range:
    - 0
    - 1
    model: fangloveskari/ORCA_LLaMA_70B_QLoRA
- sources:
  - layer_range:
    - 1
    - 2
    model: jondurbin/airoboros-l2-c70b-3.1.2
- sources:
  - layer_range:
    - 2
    - 3
    model: Sao10K/Euryale-1.3-L2-70B
- sources:
  - layer_range:
    - 3
    - 4
    model: jondurbin/airoboros-l2-c70b-3.1.2
- sources:
  - layer_range:
    - 4
    - 5
    model: fangloveskari/ORCA_LLaMA_70B_QLoRA
- sources:
  - layer_range:
    - 5
    - 6
    model: ehartford/Samantha-1.11-70b
- sources:
  - layer_range:
    - 6
    - 8
    model: Xwin-LM/Xwin-LM-70B-V0.1
- sources:
  - layer_range:
    - 8
    - 9
    model: ehartford/Samantha-1.11-70b
- sources:
  - layer_range:
    - 9
    - 10
    model: Sao10K/Euryale-1.3-L2-70B
- sources:
  - layer_range:
    - 10
    - 11
    model: ehartford/Samantha-1.11-70b
- sources:
  - layer_range:
    - 11
    - 12
    model: jondurbin/airoboros-l2-c70b-3.1.2
- sources:
  - layer_range:
    - 12
    - 13
    model: fangloveskari/ORCA_LLaMA_70B_QLoRA
- sources:
  - layer_range:
    - 13
    - 14
    model: Sao10K/Euryale-1.3-L2-70B
- sources:
  - layer_range:
    - 14
    - 15
    model: fangloveskari/ORCA_LLaMA_70B_QLoRA
- sources:
  - layer_range:
    - 15
    - 16
    model: Sao10K/Euryale-1.3-L2-70B
- sources:
  - layer_range:
    - 16
    - 17
    model: fangloveskari/ORCA_LLaMA_70B_QLoRA
- sources:
  - layer_range:
    - 17
    - 18
    model: jondurbin/airoboros-l2-c70b-3.1.2
- sources:
  - layer_range:
    - 18
    - 19
    model: NousResearch/Nous-Hermes-Llama2-70b
- sources:
  - layer_range:
    - 19
    - 20
    model: Xwin-LM/Xwin-LM-70B-V0.1
- sources:
  - layer_range:
    - 20
    - 21
    model: Sao10K/Euryale-1.3-L2-70B
- sources:
  - layer_range:
    - 21
    - 22
    model: ehartford/Samantha-1.11-70b
- sources:
  - layer_range:
    - 22
    - 23
    model: jondurbin/airoboros-l2-c70b-3.1.2
- sources:
  - layer_range:
    - 23
    - 24
    model: augtoma/qCammel-70-x
- sources:
  - layer_range:
    - 24
    - 25
    model: Sao10K/Euryale-1.3-L2-70B
- sources:
  - layer_range:
    - 25
    - 27
    model: jondurbin/airoboros-l2-c70b-3.1.2
- sources:
  - layer_range:
    - 27
    - 28
    model: Xwin-LM/Xwin-LM-70B-V0.1
- sources:
  - layer_range:
    - 28
    - 29
    model: ehartford/Samantha-1.11-70b
- sources:
  - layer_range:
    - 29
    - 30
    model: Sao10K/Euryale-1.3-L2-70B
- sources:
  - layer_range:
    - 30
    - 32
    model: Xwin-LM/Xwin-LM-70B-V0.1
- sources:
  - layer_range:
    - 32
    - 33
    model: ehartford/Samantha-1.11-70b
- sources:
  - layer_range:
    - 33
    - 34
    model: augtoma/qCammel-70-x
- sources:
  - layer_range:
    - 34
    - 35
    model: Xwin-LM/Xwin-LM-70B-V0.1
- sources:
  - layer_range:
    - 35
    - 37
    model: Sao10K/Euryale-1.3-L2-70B
- sources:
  - layer_range:
    - 37
    - 38
    model: jondurbin/airoboros-l2-c70b-3.1.2
- sources:
  - layer_range:
    - 38
    - 39
    model: ehartford/Samantha-1.11-70b
- sources:
  - layer_range:
    - 39
    - 40
    model: augtoma/qCammel-70-x
- sources:
  - layer_range:
    - 40
    - 41
    model: Xwin-LM/Xwin-LM-70B-V0.1
- sources:
  - layer_range:
    - 41
    - 42
    model: ehartford/Samantha-1.11-70b
- sources:
  - layer_range:
    - 42
    - 43
    model: Sao10K/Euryale-1.3-L2-70B
- sources:
  - layer_range:
    - 43
    - 44
    model: Xwin-LM/Xwin-LM-70B-V0.1
- sources:
  - layer_range:
    - 44
    - 45
    model: NousResearch/Nous-Hermes-Llama2-70b
- sources:
  - layer_range:
    - 45
    - 46
    model: jondurbin/airoboros-l2-c70b-3.1.2
- sources:
  - layer_range:
    - 46
    - 48
    model: ehartford/Samantha-1.11-70b
- sources:
  - layer_range:
    - 48
    - 49
    model: Sao10K/Euryale-1.3-L2-70B
- sources:
  - layer_range:
    - 49
    - 50
    model: Xwin-LM/Xwin-LM-70B-V0.1
- sources:
  - layer_range:
    - 50
    - 51
    model: jondurbin/airoboros-l2-c70b-3.1.2
- sources:
  - layer_range:
    - 51
    - 54
    model: fangloveskari/ORCA_LLaMA_70B_QLoRA
- sources:
  - layer_range:
    - 54
    - 55
    model: jondurbin/airoboros-l2-c70b-3.1.2
- sources:
  - layer_range:
    - 55
    - 56
    model: fangloveskari/ORCA_LLaMA_70B_QLoRA
- sources:
  - layer_range:
    - 56
    - 58
    model: jondurbin/airoboros-l2-c70b-3.1.2
- sources:
  - layer_range:
    - 58
    - 59
    model: Sao10K/Euryale-1.3-L2-70B
- sources:
  - layer_range:
    - 59
    - 60
    model: Xwin-LM/Xwin-LM-70B-V0.1
- sources:
  - layer_range:
    - 60
    - 62
    model: jondurbin/airoboros-l2-c70b-3.1.2
- sources:
  - layer_range:
    - 62
    - 63
    model: Xwin-LM/Xwin-LM-70B-V0.1
- sources:
  - layer_range:
    - 63
    - 64
    model: fangloveskari/ORCA_LLaMA_70B_QLoRA
- sources:
  - layer_range:
    - 64
    - 65
    model: NousResearch/Nous-Hermes-Llama2-70b
- sources:
  - layer_range:
    - 65
    - 66
    model: Sao10K/Euryale-1.3-L2-70B
- sources:
  - layer_range:
    - 66
    - 67
    model: Xwin-LM/Xwin-LM-70B-V0.1
- sources:
  - layer_range:
    - 67
    - 68
    model: augtoma/qCammel-70-x
- sources:
  - layer_range:
    - 68
    - 70
    model: Xwin-LM/Xwin-LM-70B-V0.1
- sources:
  - layer_range:
    - 70
    - 71
    model: augtoma/qCammel-70-x
- sources:
  - layer_range:
    - 71
    - 72
    model: Xwin-LM/Xwin-LM-70B-V0.1
- sources:
  - layer_range:
    - 72
    - 73
    model: Sao10K/Euryale-1.3-L2-70B
- sources:
  - layer_range:
    - 73
    - 75
    model: jondurbin/airoboros-l2-c70b-3.1.2
- sources:
  - layer_range:
    - 75
    - 76
    model: Sao10K/Euryale-1.3-L2-70B
- sources:
  - layer_range:
    - 76
    - 77
    model: augtoma/qCammel-70-x
- sources:
  - layer_range:
    - 77
    - 78
    model: Xwin-LM/Xwin-LM-70B-V0.1
- sources:
  - layer_range:
    - 78
    - 79
    model: NousResearch/Nous-Hermes-Llama2-70b
- sources:
  - layer_range:
    - 79
    - 80
    model: Xwin-LM/Xwin-LM-70B-V0.1


=========================

=> Applying Doctor-Shotgun/limarpv3-llama2-70b-qlora x 0.35
```
<!-- prompt-template start -->
## Prompt template: Alpaca

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:

```
A big thanks to [Charles](https://huggingface.co/chargoddard) for adding the layer shuffle method to his tool [mergekit](https://github.com/cg123/mergekit/tree/main) and [Henky/KoboldAI](https://koboldai.org/) for the machine he let me use.

If you want to support me, you can [here](https://ko-fi.com/undiai).