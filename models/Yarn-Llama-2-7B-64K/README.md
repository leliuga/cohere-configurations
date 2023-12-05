---
datasets:
- pg19
metrics:
- perplexity
library_name: transformers
---
# Model Card: Nous-Yarn-Llama-2-7b-64k

[Preprint (arXiv)](https://arxiv.org/abs/2309.00071)  
[GitHub](https://github.com/jquesnelle/yarn)

## Model Description

Nous-Yarn-Llama-2-7b-64k is a state-of-the-art language model for long context, further pretrained on long context data for 400 steps.  
This model is the Flash Attention 2 patched version of the original model: https://huggingface.co/conceptofmind/Yarn-Llama-2-7b-64k

Note that this model **requires** the [Flash Attention library](https://pypi.org/project/flash-attn/) in order to function correctly, see the Model Usage section for installation instructions.

## Model Training

Starting from the base Llama 2 models, this model was further pretrained on a subset of the PG19 dataset, allowing it to effectively utilize up to 64k tokens of context.

## Collaborators

 - [bloc97](https://github.com/bloc97): Methods, Paper and evals
 - [@theemozilla](https://twitter.com/theemozilla): Methods, Paper and evals
 - [@EnricoShippole](https://twitter.com/EnricoShippole): Model Training
 - [honglu2875](https://github.com/honglu2875): Paper and evals

The authors would like to thank Stability AI, Carper AI, and Eleuther AI for their generous support of significant computing resources that enabled the training of these models and the completion of this research. We would also like to thank Jonathan Tow and Dakota Mahan directly for their help in advising on the use of the Stability AI compute cluster. Additionally, we would like to thank a16z, and PygmalionAI, for providing resources to run evaluations and experiments on the models. 

## Usage and Prompt Format

Install FA2 and Rotary Extensions:
```
pip install flash-attn --no-build-isolation
pip install git+https://github.com/HazyResearch/flash-attention.git#subdirectory=csrc/rotary
```

There are no specific prompt formats as this is a pretrained base model.

## Benchmark Results

TODO

## Future Plans
We plan to continue training when we have more compute and to improve the dataset and/or instruct tune the models in order to improve the long context performance even further.

## Model Usage

The model is available for download on HuggingFace.