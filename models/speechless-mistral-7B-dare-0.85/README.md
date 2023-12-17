---
license: llama2
---
* [AWQ model(s) for GPU inference.](https://huggingface.co/TheBloke/speechless-mistral-7B-dare-0.85-AWQ)
* [GPTQ models for GPU inference, with multiple quantisation parameter options.](https://huggingface.co/TheBloke/speechless-mistral-7B-dare-0.85-GPTQ)
* [2, 3, 4, 5, 6 and 8-bit GGUF models for CPU+GPU inference](https://huggingface.co/TheBloke/speechless-mistral-7B-dare-0.85-GGUF)

Experiment for DARE(Drop and REscale), most of the delta parameters can be directly set to zeros without affecting the capabilities of SFT LMs and larger models can tolerate a higher proportion of discarded parameters.

Merged with below DARE models.

weight_mask_rate: 0.85 / use_weight_rescale: True / mask_stratery: random / scaling_coefficient: 1.0

| Model                                                        | Average | ARC    | HellaSwag | MMLU   | TruthfulQA | Winogrande | GSM8K  | 
| ------                                                       | ------  | ------ | ------    | ------ | ------     | ------     | ------ |
| Intel/neural-chat-7b-v3-1                                    | 61.59   | **66.21**  | 83.64     | 62.37  | 59.65      | 78.14      | 19.56  |
| migtissera/SynthIA-7B-v1.3                                   | 59.34   | 62.12  | 83.45     | 62.65  | 51.37      | 78.85      | 17.59  |
| bhenrym14/mistral-7b-platypus-fp16                           | 58.71   | 63.05  | 84.15     | 64.11  | 45.07      | 78.53      | 17.36  |
| jondurbin/airoboros-m-7b-3.1.2                               | 58.75   | 61.86  | 83.51     | 61.91  | 53.75      | 77.58      | 13.87  |
| teknium/CollectiveCognition-v1.1-Mistral-7B                  | 62.92   | 62.12  | 84.17     | 62.35  | **57.62**      | 75.37      | 15.62  |
| uukuguy/speechless-mistral-dolphin-orca-platypus-samantha-7b | 62.06   | 64.33  | 84.4      | 63.72  | 52.52      | 78.37      | 21.38  |
|                                                              |         |        |           |        |            |            |        |
| speechless-mistral-7b-dare-0.85 (Merge 6 DARE models)        | **64.69**   | 63.57  | **84.82**     | **64.29**  | 50.66      | **79.24**      | **45.56**  |
