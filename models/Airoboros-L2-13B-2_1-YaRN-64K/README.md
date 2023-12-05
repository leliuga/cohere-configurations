---
datasets:
- jondurbin/airoboros-2.1
---

# Extended Context (via YaRN) Finetune of Llama-2-13b with airoboros-2.1 (fp16)

[TheBloke](https://huggingface.co/TheBloke) has kindly quantized this model to [GGUF](https://huggingface.co/TheBloke/Airoboros-L2-13B-2.1-YaRN-64K-GGUF) and [GPTQ](https://huggingface.co/TheBloke/Airoboros-L2-13B-2.1-YaRN-64K-GPTQ).


## Overview

This is a finetune of [NousResearch/Yarn-Llama-2-13b-64k](https://huggingface.co/NousResearch/Yarn-Llama-2-13b-64k), which is base Llama-2-13b with additional pretraining done with YaRN scaling applied to RoPE to extend the useful context length to 64k tokens. Starting with this model, I performed instruction tuning with  [Jon Durbin's Airoboros 2.1 dataset](https://huggingface.co/datasets/jondurbin/airoboros-2.1), with the same scaling approach applied.

**This is a (merged) QLoRA fine-tune (rank 64)**. 

The finetune was performed with 1x RTX 6000 Ada (~16 hours).


## How to Use

YaRN is not implemented natively in `Transformers`. The YaRN pretrained model [NousResearch/Yarn-Llama-2-13b-64k](https://huggingface.co/NousResearch/Yarn-Llama-2-13b-64k) contains a drop-in llama architecture replacement that interfaces with the included configuration file. **To maximize compatibility, I have included the version that omits flash attention.** To run using `Transformers`, you will therefore need to pass `trust_remote_code=True`.

The PNTK method employed in my other model [bhenrym14/airophin-13b-pntk-16k-fp16](https://huggingface.co/bhenrym14/airophin-13b-pntk-16k-fp16), is very similar to YaRN. For GPTQ, I have an exllama patch that I may adapt for YaRN, but the community appears motivated to rapidly implement YaRN in common libraries, so I may not bother.

Please comment with any questions and feedback on how this model performs, especially at long context lengths!

Ooba use: Be sure to increase the `Truncate the prompt up to this length` parameter to 65586 to utilize the full context capabilities. Again `trust_remote_code=True` is imperative. Obviously, using full context requires A LOT of VRAM.

**There may be issues on Windows systems loading this model due to the decimal in "2.1" found in the model name. Try simply changing the model directory name to omit this decimal if you have issues loading the model.**

## Motivation

[Yet another RoPE extensioN method (YaRN)](https://github.com/jquesnelle/yarn) is a novel method of extending the useful context of pretrained LLMs, with architectures employing RoPE, with minimal additonal training requirements. This method is the consequence of efforts to mitigate the shortcomings of other methods such as Position Interpolation (PI) and NTK-Aware scaling. This model is an attempt to enable the community to assess the capabilities of this extension method in real world applications.

## Relative Performance (wikitext perplexity)

| Context (tokens) | <ins>**bhenrym14/airoboros-l2-13b-2.1-YaRN-64k**</ins> | bhenrym14/airoboros-l2-13b-PI-16k-fp16 | bhenrym14/airophin-v2-13b-PI-8k-fp16 | bhenrym14/airophin-13b-pntk-16k-fp16| bhenrym14/airoboros-13b-gpt4-1.4.1-PI-8192-fp16 |bhenrym14/airoboros-33b-gpt4-1.4.1-lxctx-PI-16384-fp16  | jondurbin/airoboros-l2-13b-gpt4-1.4.1 |
| --- | --- |--- | ---| ----- | -----| ------| --- |
| 512 | 7.64| 7.67 | 7.38 | 7.62 | 8.24 | 7.90 | **7.23** |
| 1024 | 6.15 | 6.15 | 5.99 | 6.20 | 6.71 | 6.17 | **5.85**  |
| 2048 | 5.29 | 5.29 | 5.22 | 5.38 | 5.87 | 5.23 | **5.07** |
| 4096 | 4.93 |4.94 | 4.90 | 5.08 | 5.50 | 4.91 | **4.77** |
| 8192 | **4.69** |4.71 | 4.71 | 4.90 | 5.32 | Not Tested | 57.1 |
| 12000 | **4.53** | 4.54 | 55 | 4.82 | 56.1 | Not Tested | Not Tested |

- Despite having a far higher scaling factor, this model is competitive with bhenrym14/airophin-13b-pntk-16k-fp16 at short context lengths.
- I may need to restrict these comparisons to models finetuned on the same dataset. Differences between airoboros 1.4.1 and 2.0m/2.1 may be a confounder.
- Overall, it appears that YaRN is capable of extending the context window with minimal impact to short context performance, when compared to other methods. Furthermore, it's able to do this with a FAR higher scaling factor, which with other methods (especially PI), resulted in serious performance degradation at shorter context lengths.
- Both the YaRN and Code LLama papers suggest that YaRN and NTK scaling may ameliorate the issue of "U shaped" attention to some degree, where long context models struggle to attend to information in the middle of the context window. Further study is needed to evaluate this. Anecdotal feedback from the community on this issue would be appreciated!

### Benchmarks

ARC (25 shot): 60.32

Hellaswag (10 shot): 83.90

MMLU (5 shot): 54.39

## Prompting:

Prompting differs with the airoboros 2.1 models. See [jondurbin/airoboros-l2-13b-2.1](https://huggingface.co/jondurbin/airoboros-l2-13b-2.1)