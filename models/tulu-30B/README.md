---
datasets:
- databricks/databricks-dolly-15k
- OpenAssistant/oasst1
- sahil2801/CodeAlpaca-20k
language:
- en
---

# Tulu 30B

This model is a 30B LLaMa model finetuned on a mixture of instruction datasets (FLAN V2, CoT, Dolly, Open Assistant 1, GPT4-Alpaca, Code-Alpaca, and ShareGPT).
*Please note this is a model diff - see below for usage instructions*.

This was trained as part of the paper [How Far Can Camels Go? Exploring the State of Instruction Tuning on Open Resources](https://arxiv.org/abs/2306.04751).
The codebase used to train and evaluate this model can be found at [https://github.com/allenai/open-instruct](https://github.com/allenai/open-instruct).

This model is licensed under the AI model license given in LICENSE.txt along with the original Llama license (llama_license.txt).

## Usage

We assume you have access to a LLaMa model in HF format already. You can find details on getting access and converting the model here:
[https://huggingface.co/docs/transformers/main/model_doc/llama](https://huggingface.co/docs/transformers/main/model_doc/llama)

Clone [https://github.com/allenai/open-instruct](https://github.com/allenai/open-instruct) and install the required dependencies, or just copy `scripts/weight_diff.py`
and install the minimal requirements listed in `weight-diff-requirements.txt`. Then download or clone this model diff to the same machine.

Then, run:
```bash
python scripts/weight_diff.py recover --path_raw ${hf_llama_path} --path_tuned ${output_path} --path_diff ${diff_location}
```

And you will have a recovered model! Note this takes up a decent amount of RAM, especially for the larger models.

## Input Format

The model is trained to use the following format (note the newlines):
```
<|user|>
Your message here!
<|assistant|>
```

For best results, format all inputs in this manner. **Make sure to include a newline after `<|assistant|>`, this can affect generation quality quite a bit.**

## Performance

Here is the performance of this model across benchmarks explored in our paper [How Far Can Camels Go? Exploring the State of Instruction Tuning on Open Resources](https://arxiv.org/abs/2306.04751):

| MMLU 0-shot | MMLU 5-shot | GSM Direct | GSM CoT | BBH Direct | BBH CoT | TydiQA Gold-Passage | TydiQA Closed-book | Codex-Eval Pass@1 | Codex-Eval Pass@10 | AlpacaFarm vs Davinci-003 | Average |
|:-----------:|:-----------:|:----------:|:-------:|:----------:|:-------:|:-------------------:|:------------------:|:-----------------:|:------------------:|:-------------------------:|---------|
| 57.7 | 58.4 | 6.0 | 51.0 | 45.8 | 48.7 | 58.2 | 12.3 | 25.4 | 46.0 | 63.5 | 44.7 |

If you use this model, please cite our work, the llama paper, and the original datasets:

```
@misc{wang2023far,
      title={How Far Can Camels Go? Exploring the State of Instruction Tuning on Open Resources}, 
      author={Yizhong Wang and Hamish Ivison and Pradeep Dasigi and Jack Hessel and Tushar Khot and Khyathi Raghavi Chandu and David Wadden and Kelsey MacMillan and Noah A. Smith and Iz Beltagy and Hannaneh Hajishirzi},
      year={2023},
      eprint={2306.04751},
      archivePrefix={arXiv},
      primaryClass={cs.CL}
}
```

```
@misc{touvron2023llama,
      title={LLaMA: Open and Efficient Foundation Language Models}, 
      author={Hugo Touvron and Thibaut Lavril and Gautier Izacard and Xavier Martinet and Marie-Anne Lachaux and Timothée Lacroix and Baptiste Rozière and Naman Goyal and Eric Hambro and Faisal Azhar and Aurelien Rodriguez and Armand Joulin and Edouard Grave and Guillaume Lample},
      year={2023},
      eprint={2302.13971},
      archivePrefix={arXiv},
      primaryClass={cs.CL}
}
```

```
@misc{dolly,
  author = {Databricks},
  title = {Free Dolly: Introducing the World's First Truly Open Instruction-Tuned LLM},
  year = {2023},
  publisher = {GitHub},
  journal = {GitHub repository},
  howpublished = {Blog post},
  url = {https://www.databricks.com/blog/2023/04/12/dolly-first-open-commercially-viable-instruction-tuned-llm}
}
```

```
@article{longpre2023flan,
  title={The Flan Collection: Designing Data and Methods for Effective Instruction Tuning},
  author={Longpre, Shayne and Hou, Le and Vu, Tu and Webson, Albert and Chung, Hyung Won and Tay, Yi and Zhou, Denny and Le, Quoc V and Zoph, Barret and Wei, Jason and others},
  journal={arXiv preprint arXiv:2301.13688},
  year={2023}
}
```

```
@misc{köpf2023openassistant,
      title={OpenAssistant Conversations -- Democratizing Large Language Model Alignment}, 
      author={Andreas Köpf and Yannic Kilcher and Dimitri von Rütte and Sotiris Anagnostidis and Zhi-Rui Tam and Keith Stevens and Abdullah Barhoum and Nguyen Minh Duc and Oliver Stanley and Richárd Nagyfi and Shahul ES and Sameer Suri and David Glushkov and Arnav Dantuluri and Andrew Maguire and Christoph Schuhmann and Huu Nguyen and Alexander Mattick},
      year={2023},
      eprint={2304.07327},
      archivePrefix={arXiv},
      primaryClass={cs.CL}
}
```

```
@article{peng2023instruction,
  title={Instruction Tuning with GPT-4},
  author={Peng, Baolin and Li, Chunyuan and He, Pengcheng and Galley, Michel and Gao, Jianfeng},
  journal={arXiv preprint arXiv:2304.03277},
  year={2023}
}
```

```
@misc{codealpaca,
  author = {Sahil Chaudhary},
  title = {Code Alpaca: An Instruction-following LLaMA model for code generation},
  year = {2023},
  publisher = {GitHub},
  journal = {GitHub repository},
  howpublished = {\url{https://github.com/sahil280114/codealpaca}},
}
```