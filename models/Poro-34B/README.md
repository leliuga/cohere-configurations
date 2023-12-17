---
license: apache-2.0
datasets:
- cerebras/SlimPajama-627B
- bigcode/starcoderdata
- mc4
- allenai/dolma
---
<div align="center">
<img src="./poro-logo.png" width="200px">
</div>

# Poro 34B Model Card

_**NOTE:** This is a **research checkpoint** of a model for which **training has not been completed.** It is being provided in its current state for research and testing purposes. **Care should be taken when using the outputs of the model.** Once pretraining has completed we intend to release additional instruction-tuned and chat-tuned varieties._

Poro is a 34B parameter decoder-only transformer pretrained on Finnish, English and code.  It is being trained on 1 trillion tokens (700 billion as of this release). Poro is a fully open source model and is made available under the Apache 2.0 License.

Poro was created in a collaboration between [SiloGen](https://www.silo.ai/silogen) from [Silo AI](https://www.silo.ai/), the [TurkuNLP group](https://turkunlp.org/) of the University of Turku, and [High Performance Language Technologies](https://hplt-project.org/) (HPLT). Training was conducted on the [LUMI supercomputer](https://www.lumi-supercomputer.eu/), using compute resources generously provided by [CSC](https://csc.fi/) - IT Center for Science, Finland.

This project is part of an ongoing effort to create open source large language models for non-English and especially low resource languages like Finnish. Through the combination of English and Finnish training data we get a model that outperforms previous Finnish only models, while also being fluent in English and code, and capable of basic translation between English and Finnish.

Poro 34B is only the first model of our model family.  Work is already underway on our next models which will support additional languages, and include features like flash attention, rotary embeddings, and grouped query attention.

_What does Poro mean?_ Poro is the Finnish word for Reindeer! 🦌 These animals are native to Finland and hold a significant and historical role in Finnish culture.

## Model Overview
_**NOTE:** In addition to being an early research release, Poro is a base model which needs further fine tuning for most use cases._

Poro is a generative pretrained transformer using a BLOOM architecture, and makes use of ALiBi embeddings to support context length extrapolation at inference time.

| Hyperparameter | Value  |
| :------------- | :----: |
| n_parameters | 34.2B |
| n_layers | 54 |
| n_heads | 56 |
| d_model | 7168 |
| vocab_size | 128000 |
| sequence_length | 2048 |

## Poro Research Checkpoints

Checkpoints are available as branches in the repository.  Checkpoints will be released roughly every 100B tokens.  The main branch will always point to the latest checkpoint.  The following checkpoints are available:

* [100B](https://huggingface.co/LumiOpen/Poro-34B/tree/100B)
* [200B](https://huggingface.co/LumiOpen/Poro-34B/tree/200B)
* [300B](https://huggingface.co/LumiOpen/Poro-34B/tree/300B)
* [400B](https://huggingface.co/LumiOpen/Poro-34B/tree/400B)
* [500B](https://huggingface.co/LumiOpen/Poro-34B/tree/500B)
* [600B](https://huggingface.co/LumiOpen/Poro-34B/tree/600B)
* [700B](https://huggingface.co/LumiOpen/Poro-34B/tree/700B)

The transformers library allows you to load a checkpoint from a branch as follows:

```python
branch = "200B"
model = transformers.AutoModelForCausalLM.from_pretrained(
    "LumiOpen/Poro-34B",
    torch_dtype=torch.bfloat16,
    revision=branch,
)
```

## Training

Poro was trained on the LUMI supercomputer, using 512 AMD MI250X GPUs. Each MI250X GPU has two Graphics Complex Dies (GCDs) for a world size of 1024 during training, using activation checkpointing, a micro batch size of 1, gradient accumulation of 16, and a 3D parallelism strategy of TP=2, PP=4, DP=128.

Training began in September 2023 using a [custom fork](https://github.com/TurkuNLP/Megatron-DeepSpeed) of the Megatron-Deepspeed framework.

## Training Hyperparameters

| Hyperparameter | Value | Comment |
| :------------: | :---: | :------:|
| Precision | bfloat16 | |
| Optimizer | AdamW | |
| Learning rate | 1.5e-4 | 10B tokens warm-up, cosine decay to 2e-5 |
| Weight decay | 1e-1 | |
| Batch size | 2048 | 2048 samples x 2048 tokens = 4194304 tokens |

## Tokenizer

Poro uses a custom 128K Bloom tokenizer trained on the same English, Finnish and Code dataset used to train the model.

## Dataset
Poro is being trained on a 1 trillion token mixed dataset of English, Finnish and Code.

| Dataset | Notes | Percentage | Epochs | Tokens |
| :-----: | :---: | :--------: | :----: | :----: |
| SlimPajama | Excluding books3 data | 54.16% | 1x | 541.7B |
| Finnish | TurkuNLP Finnish dataset | 13.05% | 4x | 131.5B |
| Tatoeba | English/Finnish sentence pairs | 0.81% | 1x | 8.0B |
| Starcoder | | 31.53% | 1.52x | 315.4B |
| Project Gutenberg | from Dolma dataset | 0.46% | 1x | 4.5B |

The Finnish dataset is a combination of many Finnish resources:

* [Finnish Internet Parsebank](https://turkunlp.org/finnish_nlp.html)
* [mC4 multilingual colossal, cleaned Common Crawl](https://huggingface.co/datasets/mc4)
* [Common Crawl Finnish](https://github.com/turkunlp/CC-Fi)
* [Finnish Wikipedia](https://fi.wikipedia.org/wiki)
* [Lönnrot Projekti Lönnrot](http://www.lonnrot.net/)
* [Suomi24 The Suomi 24 Corpus 2001-2020](http://urn.fi/urn:nbn:fi:lb-2021101527)
* [Reddit r/Suomi submissions and comments](https://www.reddit.com/r/Suomi)
* [STT Finnish News Agency Archive 1992-2018](http://urn.fi/urn:nbn:fi:lb-2019041501)
* [Yle Finnish News Archive 2011-2018](http://urn.fi/urn:nbn:fi:lb-2017070501)
* [Yle Finnish News Archive 2019-2020](http://urn.fi/urn:nbn:fi:lb-2021050401)
* [Yle News Archive Easy-to-read Finnish 2011-2018](http://urn.fi/urn:nbn:fi:lb-2019050901)
* [Yle News Archive Easy-to-read Finnish 2019-2020](http://urn.fi/urn:nbn:fi:lb-2021050701)

## Evaluation Results

Despite the early training stage, Poro already exceeds the performance of the Finnish-only [FinGPT](https://turkunlp.org/gpt3-finnish) language models on the [FIN-bench](https://github.com/TurkuNLP/FIN-bench) Finnish language benchmark.

Full evaluation results will be published with the final model. 

## Ethical Considerations and Limitations

_Poro 34B is a release of a partially trained model, and special care should be taken when using any output._

Poro is an advanced language model, primarily optimized for English, Finnish and code, with no meaningful proficiency in any other languages. As with most AI-driven systems, Poro is a product of the vast data it has been trained on, which may reflect the imperfections, biases, and idiosyncrasies of the wider web. Poro may, at times, produce outputs that can be considered inaccurate, prejudiced, or controversial. Users and developers engaging with Poro should exercise discretion and consider additional evaluation and customization to ensure the model's responses align with their specific needs and ethical standards.

## License

Poro is released under the Apache 2.0 license.
