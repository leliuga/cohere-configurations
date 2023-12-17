---
license: wtfpl
datasets:
- JosephusCheung/GuanacoDataset
- Open-Orca/OpenOrca
- stingning/ultrachat
- meta-math/MetaMathQA
- liuhaotian/LLaVA-Instruct-150K
- jondurbin/airoboros-3.1
- WizardLM/WizardLM_evol_instruct_V2_196k
- RyokoAI/ShareGPT52K
- RyokoAI/Fandom23K
- milashkaarshif/MoeGirlPedia_wikitext_raw_archive
- wikipedia
- wiki_lingua
- fnlp/moss-003-sft-data
- garage-bAInd/Open-Platypus
- LDJnr/Puffin
- openbmb/llava_zh
- BAAI/COIG
- TigerResearch/tigerbot-zhihu-zh-10k
- liwu/MNBVC
- teknium/openhermes
language:
- en
- zh
pipeline_tag: text-generation
tags:
- llama
- llama2
- qwen
- causallm
---
[![CausalLM](https://huggingface.co/JosephusCheung/tmp/resolve/main/7.72b.png)](https://causallm.org/)

*Image drawn by GPT-4 DALL·E 3* **TL;DR: Perhaps this 7B model, better than all existing models <= 33B, in most quantitative evaluations...**

# CausalLM 7B - Fully Compatible with Meta LLaMA 2 
Use the transformers library that does not require remote/external code to load the model, AutoModelForCausalLM and AutoTokenizer (or manually specify LlamaForCausalLM to load LM, GPT2Tokenizer to load Tokenizer), and model quantization is fully compatible with GGUF (llama.cpp), GPTQ, and AWQ.

# Recent Updates: [DPO-α Version](https://huggingface.co/CausalLM/7B-DPO-alpha) outperforms Zephyr-β in MT-Bench

**llama.cpp GGUF models**
GPT2Tokenizer fixed by [Kerfuffle](https://github.com/KerfuffleV2) on [https://github.com/ggerganov/llama.cpp/pull/3743](https://github.com/ggerganov/llama.cpp/pull/3743), new models are reuploaded.

Thanks TheBloke for GGUF quants: [https://huggingface.co/TheBloke/CausalLM-7B-GGUF](https://huggingface.co/TheBloke/CausalLM-7B-GGUF)

**Caution:** Unofficial GPTQ and AWQ models may have issues as they use Wikitext for calibration, while this model has undergone considerable training on a synthesized Wikipedia conversation dataset.

It is not recommended to use any form of quantization, but rather to use smaller-sized models, as the 7B and 14B versions have high consistency. However, if you do use model quantization, please use GGUF.

## Read Me:

Also see [14B Version](https://huggingface.co/CausalLM/14B)

This model was trained based on the model weights of Qwen (and LLaMA2 was used, yes, for calculating some initial weights), you may also need to comply with the commercial use restrictions of these two models depending on the situation. The training process utilized a model architecture that was identical to LLaMA2, using the same attention calculation method as the original MHA LLaMA2 models, and no additional scaling applied to the Rotary Positional Encoding (RoPE).

We manually curated a SFT dataset of 1.3B tokens for training, utilizing open source datasets from Hugging Face. For most of these sentences, we performed manual or synthetic rewrites and generated alternate language versions using larger language models. Additionally, we conducted augmented text training using carefully selected entries from Wikipedia, as well as featured entries from Fandom and filtered entries from Moegirlpedia. In order to strike a balance between efficiency and quality, 100% of the data used for training was synthetic data, no direct use of text from the internet or original texts from publicly available datasets was employed for fine-tuning.

The 7B version of the model is a distilled version of the 14B model, specifically designed for speculative sampling. Therefore, it is important to exercise caution when directly using the model, as it may produce hallucinations or unreliable outputs.

Please note that the model was trained on unfiltered internet data. Since we do not have the capacity to vet all of it, there may be a substantial amount of objectionable content, pornography, violence, and offensive language present that we are unable to remove. Therefore, you will still need to complete your own checks on the model's safety and filter keywords in the output. Due to computational resource constraints, we are presently unable to implement RLHF for the model's ethics and safety, nor training on SFT samples that refuse to answer certain questions for restrictive fine-tuning.

Bonus: The model underwent some fine-tuning on the prompt format introduced in LLaVA1.5 that is unrelated to image attention calculation. Therefore, aligning the ViT Projection module with frozen LM under visual instructions would enable rapid implementation of effective multimodal capabilities.

## PROMPT FORMAT:
[chatml](https://github.com/openai/openai-python/blob/main/chatml.md)

**System Prompt must not be empty!**


## MMLU:
stem ACC: 56.83 

Humanities ACC: 58.79 

other ACC: 70.04 

social ACC: 72.41 

**AVERAGE ACC:63.82** (Outperforms / Equal to the best Mistral-7B Chat-style fine-tunes, ChatGLM3-6B and ALL other models under 33B.)

## CEval (Val):
STEM acc: 61.67

Social Science acc: 81.94

Humanities acc: 77.19

Other acc: 68.35

Hard acc:48.03

**AVERAGE acc:70.27** (Outperforms ALL 7B models currently, including ChatGLM3-6B.)

## GSM8K

**Zero-shot ACC 0.5921152388172858** (Outperforms WizardMath-7B and Qwen-7B)

## MT-Behch on DPO Version
| Model                     | MT-Bench     |
| ------------------------- | ------------ |
| GPT-4                     | 8.99         |
| GPT-3.5-Turbo             | 7.94         |
|                           |              |
| Zephyr-7b-β (Overfitting) | 7.34         |
| Zephyr-7b-α               | 6.88         |
|                           |              |
| **[CausalLM/14B-DPO-α](https://huggingface.co/CausalLM/14B-DPO-alpha)**    | **7.618868** |
| **[CausalLM/7B-DPO-α](https://huggingface.co/CausalLM/7B-DPO-alpha)**     | **7.038125** |

# 因果语言模型 7B - 与 Meta LLaMA 2 完全兼容
使用无需远程/外部代码的transformers库加载模型，AutoModelForCausalLM和AutoTokenizer（或者手动指定LlamaForCausalLM加载LM， GPT2Tokenizer加载Tokenizer），并且模型量化与GGUF（llama.cpp）、GPTQ、AWQ完全兼容。

# 最近更新: [DPO-α Version](https://huggingface.co/CausalLM/7B-DPO-alpha) 在 MT-Bench 超过 Zephyr-β

**llama.cpp GGUF models**
GPT2Tokenizer 支持由 [Kerfuffle](https://github.com/KerfuffleV2) 修复于 [https://github.com/ggerganov/llama.cpp/pull/3743](https://github.com/ggerganov/llama.cpp/pull/3743)，新模型稍后上传。

感谢 TheBloke 制作 GGUF 版本量化模型: [https://huggingface.co/TheBloke/CausalLM-7B-GGUF](https://huggingface.co/TheBloke/CausalLM-7B-GGUF)

**注意：** 非官方 GPTQ 和 AWQ 模型可能存在问题，因为它们使用 Wikitext 进行校准，而该模型已经在合成的 Wikipedia 对话数据集上经过了大量的训练。

不建议使用任何形式的量化，而是使用较小尺寸的模型，因为7B和14B版本具有较高的一致性。 但是，如果您确实使用模型量化，请使用 GGUF。

## 请读我：

另请参阅[14B版本](https://huggingface.co/CausalLM/14B)

该模型是基于Qwen的权重（并使用了LLaMA2权重，是的，用于计算一些权重初始化），您根据情况可能还需要遵守这两个模型的商业使用限制。训练过程中使用了与LLaMA2相同的模型结构，使用原始MHA LLaMA2模型的相同注意力计算方法，对旋转位置编码（RoPE）没有进行额外的缩放。

我们手动筛选了一个包含13亿个标记的SFT数据集进行训练，利用了Hugging Face的开源数据集。对于大多数句子，我们进行了手动或合成改写，并使用更大的语言模型生成了其他语言版本。此外，我们还使用了精心挑选的来自维基百科的条目、来自Fandom的精选条目以及来自萌娘百科的过滤条目进行增强文本训练。为了在效率和质量之间取得平衡，训练所使用的100%数据都是合成数据，没有直接使用来自互联网或公开可用数据集的原始文本进行微调。

7B版本的模型是14B模型的精简版本，专门设计用于推测抽样。因此，在直接使用模型时，需要谨慎行事，因为它可能会产生幻觉或不可靠的输出。

请注意，模型是在未经过滤的互联网数据上进行训练的。由于我们无法审核所有数据，可能会出现大量不良内容、色情、暴力和冒犯性语言，我们无法删除这些内容。因此，您仍然需要对模型的安全性进行自己的检查，并对输出中的关键词进行过滤。由于计算资源的限制，我们目前无法为模型的伦理和安全实施RLHF，也无法对拒绝回答某些问题的SFT样本进行训练以进行限制性微调。

额外奖励：模型在LLaVA1.5中引入的提示格式上进行了一些微调，与图像注意力计算无关。因此，将ViT投影模块与冻结的LM对齐，并根据视觉指令实施快速实现有效的多模态能力。

## 提示格式：
[chatml](https://github.com/openai/openai-python/blob/main/chatml.md)

**系统提示不能为空！**


## MMLU：
STEM准确率：56.83

人文学科准确率：58.79

其他准确率：70.04

社会学准确率：72.41

**平均准确率：63.82** （优于/平于最好的 Mistral-7B 聊天格式的微调，ChatGLM3-6B 和其余的33B及以下模型。）

## CEval（验证集）：
STEM准确率：61.67

社会科学准确率：81.94

人文学科准确率：77.19

其他准确率：68.35

困难准确率：48.03

**平均准确率：70.27** （优于当前所有7B模型，包括 ChatGLM3-6B）

## GSM8K

**零样本准确率0.5921152388172858** （优于WizardMath-7B和Qwen-7B）

## DPO 版本的 MT-Behch
| Model                     | MT-Bench     |
| ------------------------- | ------------ |
| GPT-4                     | 8.99         |
| GPT-3.5-Turbo             | 7.94         |
|                           |              |
| Zephyr-7b-β (Overfitting) | 7.34         |
| Zephyr-7b-α               | 6.88         |
|                           |              |
| **[CausalLM/14B-DPO-α](https://huggingface.co/CausalLM/14B-DPO-alpha)**    | **7.618868** |
| **[CausalLM/7B-DPO-α](https://huggingface.co/CausalLM/7B-DPO-alpha)**     | **7.038125** |