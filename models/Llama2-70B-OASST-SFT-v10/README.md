---
license: llama2
language:
- en
datasets:
- rombodawg/LosslessMegaCodeTrainingV2_1m_Evol_Uncensored
- OpenAssistant/oasst1
- shahules786/orca-best
- argilla/databricks-dolly-15k-curated-multilingual
library_name: transformers
pipeline_tag: text-generation
tags:
- sft
---
# Open-Assistant Llama2 70B SFT v10

This model is an Open-Assistant fine-tuning of Meta's [Llama2 70B](https://huggingface.co/meta-llama/Llama-2-70b) LLM. 
It was fine-tuned in two stages, first on a mix of synthetic instrunctions and coding tasks and then in a "polishing" stage
on the best human demonstrations collected at [open-assistant.io](https://open-assistant.io/) up to July 23, 2023 (see [Configuration Details](#configuration-details) below).

## Model Details

- **Finetuned from:** [meta-llama/Llama-2-70b](https://huggingface.co/meta-llama/Llama-2-70b) via [epfLLM/Megatron-LLM](https://github.com/epfLLM/Megatron-LLM)
- **Model type:** Causal decoder-only transformer language model
- **Language:** English (and limited capabilities in German, Spanish, French, Italian, Portuguese, Polish, Dutch, Romanian, Czech, Swedish)
- **Weights & Biases training logs:** [Stage 1](https://wandb.ai/open-assistant/public-sft/runs/run45_oasst_pre10_llama2_70b) (1 epoch pretrain-mix, 12k steps), [Stage 2](https://wandb.ai/open-assistant/public-sft/runs/run46_oasst_sft10_llama2_70b) (3 epochs oasst top-1, 519 steps)
- **Demo:** [Continuations for 250 random prompts (TGI, 4bit nf4 quantization)](https://open-assistant.github.io/oasst-model-eval/?f=https%3A%2F%2Fraw.githubusercontent.com%2FOpen-Assistant%2Foasst-model-eval%2Fmain%2Fsampling_reports%2Foasst-sft%2F2023-08-22_OpenAssistant_llama2-70b-oasst-sft-v10_sampling_noprefix2_nf4.json%0A)
- **Evaluation** [FastEval-OpenAssistant Overview](https://tju01.github.io/FastEval-OpenAssistant/) (using [FastEval](https://github.com/FastEval/FastEval) & [vLLM](https://github.com/vllm-project/vllm)) 
- **License:** [LLAMA 2 COMMUNITY LICENSE AGREEMENT](https://huggingface.co/meta-llama/Llama-2-70b/raw/main/LICENSE.txt)
- **Contact:** [Open-Assistant Discord](https://ykilcher.com/open-assistant-discord)


## Prompting / Prompt Template

Due to public demand (see [survey](https://twitter.com/erhartford/status/1682403597525430272)) we changed the prompt-template for this model from custom prompter/assistant tokens to OpenAI's [chatml](https://github.com/openai/openai-python/blob/main/chatml.md) standard prompt format.
We hope that this leads to greater compatibility with chat inference/frontend applications.

Prompt dialogue template:

```
"""
<|im_start|>system
{system_message}<|im_end|>
<|im_start|>user
{prompt}<|im_end|>
<|im_start|>assistant
"""
```

The model input can contain multiple conversation turns between user and assistant, e.g.
```
<|im_start|>user
{prompt 1}<|im_end|>
<|im_start|>assistant
{reply 1}<|im_end|>
<|im_start|>user
{prompt 2}<|im_end|>
<|im_start|>assistant
(...)
```

The model was partly trained with orca system messages.  
For inference we recommend to use the official [Llama2 system message](https://github.com/facebookresearch/llama/blob/ea9f33d6d3ea8ed7d560d270986407fd6c2e52b7/example_chat_completion.py#L57-L61):
```
<|im_start|>system
You are a helpful, respectful and honest assistant. Always answer as helpfully as possible, while being safe. Your answers should not include any harmful, unethical, racist, sexist, toxic, dangerous, or illegal content. Please ensure that your responses are socially unbiased and positive in nature.

If a question does not make any sense, or is not factually coherent, explain why instead of answering something not correct. If you don't know the answer to a question, please don't share false information.
<|im_end|>
```

### Credits & Special Thanks

- Thanks to [Meta AI](https://ai.meta.com/) for training and releasing the Llama2 model.
- Distributed training support was provided by EPFL's [Machine Learning and Optimization Laboratory](https://www.epfl.ch/labs/mlo/), and [Natural Language Processing Lab](https://nlp.epfl.ch/).
- The open-source [epfLLM/Megatron-LLM](https://github.com/epfLLM/Megatron-LLM) trainer was used for fine-tuning.
- [rombodawg](https://huggingface.co/rombodawg) curated the [LosslessMegaCodeTrainingV2_1m_Evol_Uncensored](https://huggingface.co/datasets/rombodawg/LosslessMegaCodeTrainingV2_1m_Evol_Uncensored) dataset.
- [ehartford](https://huggingface.co/ehartford) generated and published the [ehartford/dolphin](https://huggingface.co/datasets/ehartford/dolphin) and the [ehartford/oa_leet10k](https://huggingface.co/datasets/ehartford/oa_leet10k) datasets.
- [Argilla](https://huggingface.co/argilla) curated and published the [argilla/databricks-dolly-15k-curated-multilingual](https://huggingface.co/datasets/argilla/databricks-dolly-15k-curated-multilingual) dataset.
- [shahules786](https://github.com/shahules786) de-duped and filtered the Dolphin dataset with a cluster-center approach and generated the orca-best (ocra-chat) dataset.
- [andreaskoepf](https://github.com/andreaskoepf/) prepared & orchestrated the training.

We want to especially thank everyone who contributed in the crowed-sourced Open-Assistant dataset creation on https://open-assistant.io/ - without you this project would not have been possible.

## Ethical Considerations and Limitations

Testing conducted to date has been in English, and has not covered, nor could it cover all scenarios. 
For these reasons, as with all LLMs, the potential outputs of llama2-70b-oasst-sft-v10 cannot be predicted
in advance, and the model may in some instances produce inaccurate, biased or other objectionable responses
to user prompts. Therefore, before deploying any applications of llama2-70b-oasst-sft-v10, developers should
perform safety testing and tuning tailored to their specific applications of the model.

Please see Meta's [Responsible Use Guide](https://ai.meta.com/llama/responsible-use-guide/).


## Inference via TGI

An early version of this model had an embedding count of 32,007 which was incompatible to sharding with [TGI](https://github.com/huggingface/text-generation-inference).
In the current version the embeddings and the lm_head weights have been padded to a multiple of 128 (by replicating the emembeddings of the unk-token (id: 0)).
Sharded inference with TGI should now work as expected.


## Configuration Details

The "pretokenizer" utility used to tokenize the datamix is part of the Open-Assistant github repository and can be found here: [model/pretokenizer](https://github.com/LAION-AI/Open-Assistant/tree/main/model/pretokenizer).


### Stage 1 Pretokenizer Configuration

Entries of the dataset with assistant replies shorter than 25 tokens were excluded from training.

```
oasst_pre10_min25:
  datasets:
    - megacode2:
        fraction: 0.5
        val_split: 0.01
        max_val_set: 1000
    - orca-chat:
        val_split: 0.01
        max_val_set: 1000
    - dolly15k_multilingual:
        val_split: 0.05
        max_val_set: 300
    - oa_leet10k:
        val_split: 0.05
        max_val_set: 250
  output_dir: "output/oasst_pre10_min25"
  filename_prefix: "oasst_pre10"
  min_assistant_tokens: 25
```

Stage 1 dataset statistics:
```
# Stats for output/oasst_pre10_min25_llama2

## Stats for 'Subset of InstructionDataset (megacode2)' (466364 samples (50.0%))
-----------------
  Accepted: 398223/466364 (85.4%)
  Accepted tokens: 167676873
  Skipped: 68141 (14.6%)
  Min tokens per sample: 36
  Max tokens per sample: 11810
  Avg tokens per sample: 421.063
-----------------

## Stats for 'Subset of OrcaChat (orca-chat)' (325616 samples (100.0%))
-----------------
  Accepted: 325616/325616 (100.0%)
  Accepted tokens: 178307574
  Skipped: 0 (0.0%)
  Min tokens per sample: 105
  Max tokens per sample: 10408
  Avg tokens per sample: 547.601
-----------------

## Stats for 'Subset of Dolly15kMultilingual' (57020 samples (100.0%))
-----------------
  Accepted: 47494/57020 (83.3%)
  Accepted tokens: 13883177
  Skipped: 9526 (16.7%)
  Min tokens per sample: 34
  Max tokens per sample: 9172
  Avg tokens per sample: 292.314
-----------------

## Stats for 'Subset of InstructionDataset (oa_leet10k)' (22236 samples (100.0%))
-----------------
  Accepted: 22236/22236 (100.0%)
  Accepted tokens: 15905296
  Skipped: 0 (0.0%)
  Min tokens per sample: 168
  Max tokens per sample: 10588
  Avg tokens per sample: 715.295
-----------------

## Stats for 'total' (871236 samples (100.0%))
-----------------
  Accepted: 793569/871236 (91.1%)
  Accepted tokens: 375772920
  Skipped: 77667 (8.9%)
  Min tokens per sample: 34
  Max tokens per sample: 11810
  Avg tokens per sample: 473.523
-----------------
```


### Stage 2 Pretokenizer Configuration

```
oasst_top1:
  datasets:
    - oasst_export:
        lang: "bg,ca,cs,da,de,en,es,fr,hr,hu,it,nl,pl,pt,ro,ru,sl,sr,sv,uk"
        input_file_path: 2023-07-23_oasst_ready.tar.gz
        top_k: 1
        val_split: 0.05
  output_dir: "output/oasst_top1_2023-07-23"
  filename_prefix: "oasst_top1"
```

Stage 2 dataset statistics:

```
# Stats for output/oasst_top1_2023-07-23_llama2

## Stats for 'ListDataset' (11441 samples (100.0%))
-----------------
  Accepted: 11441/11441 (100.0%)
  Accepted tokens: 5315368
  Skipped: 0 (0.0%)
  Min tokens per sample: 20
  Max tokens per sample: 5407
  Avg tokens per sample: 464.58945896337735
-----------------

## Stats for 'total' (11441 samples (100.0%))
-----------------
  Accepted: 11441/11441 (100.0%)
  Accepted tokens: 5315368
  Skipped: 0 (0.0%)
  Min tokens per sample: 20
  Max tokens per sample: 5407
  Avg tokens per sample: 464.58945896337735
-----------------
```


### Megatron Fine-Tuning Arguments for Stage 1 (Instruction Tuning):
```
--tensor_model_parallel_size 8
--pipeline_model_parallel_size 4
--load ./checkpoints/llama2-70b-tp8-pp4
--save ./checkpoints/llama2-70b-tp8-pp4-oasst_pre10
--tensorboard_dir ./checkpoints/llama2-70b-tp8-pp4-oasst_pre10/logging
--data_path ./data/oasst_pre10_min25_llama2/oasst_sft10-train
--model_name llama2
--tokenizer_type SentencePieceTokenizer
--bf16
--global_batch_size 64
--micro_batch_size 2
--vocab_file=./llama2/Llama-2-7b/tokenizer.model
--use_rms_norm
--glu_activation swiglu
--no_tie_embed_logits
--vocab_extra_ids_list "\"<|im_start|>,<|im_end|>\""
--layernorm_epsilon 1e-5
--use_flash_attn
--no_bias_gelu_fusion
--seq_length 4096
--max_position_embeddings 4096
--log_interval 1
--save_interval 500
--eval_interval 50
--eval_iters 10
--hidden_dropout 0.0
--position_embedding_type rotary
--no_bias_dropout_fusion
--use_checkpoint_args
--train_iters 12000
--attention_dropout 0.0
--adam_beta1 0.9
--adam_beta2 0.95
--adam_eps 1e-12
--lr_decay_style cosine
--lr_warmup_iters 100
--lr 1e-5
--min_lr 1e-6
--weight_decay 0.000001
--sequence_parallel
--recompute_granularity selective
--log_timers_to_tensorboard
--rope_scaling_factor 1.0
--wandb_logger
```

### Megatron Fine-Tuning Arguments for Stage 2 (OASST Polishing, LIMA Dropout):
```
--tensor_model_parallel_size 8
--pipeline_model_parallel_size 4
--load ./checkpoints/llama2-70b-tp8-pp4-oasst_pre10
--save ./checkpoints/llama2-70b-tp8-pp4-oasst_sft10
--tensorboard_dir ./checkpoints/llama2-70b-tp8-pp4-oasst_sft10/logging
--data_path ./data/oasst_top1_2023-07-23_llama2/oasst_top1-train
--model_name llama2
--tokenizer_type SentencePieceTokenizer
--bf16
--global_batch_size 64
--micro_batch_size 2
--vocab_file=./llama2/Llama-2-7b/tokenizer.model
--use_rms_norm
--glu_activation swiglu
--no_tie_embed_logits
--vocab_extra_ids_list "\"<|im_start|>,<|im_end|>\""
--layernorm_epsilon 1e-5
--use_flash_attn
--no_bias_gelu_fusion
--seq_length 4096
--max_position_embeddings 4096
--log_interval 1
--save_interval 346
--eval_interval 50
--eval_iters 10
--hidden_dropout 0.25
--lima_dropout
--position_embedding_type rotary
--no_bias_dropout_fusion
--use_checkpoint_args
--train_iters 519
--attention_dropout 0.0
--adam_beta1 0.9
--adam_beta2 0.95
--adam_eps 1e-12
--lr_decay_style cosine
--lr_warmup_iters 100
--lr 1e-5
--min_lr 1e-6
--weight_decay 0.000001
--sequence_parallel
--recompute_granularity selective
--log_timers_to_tensorboard
--rope_scaling_factor 1.0
--finetune
--wandb_logger
```