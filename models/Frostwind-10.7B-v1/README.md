---
license: cc-by-nc-4.0
language:
- en
---

Frostwind-v1

![Frost1](https://huggingface.co/Sao10K/Frostwind-10.7B-v1/resolve/main/frost1.png)

A finetune of [upstage/SOLAR-10.7B-v1.0](https://huggingface.co/upstage/SOLAR-10.7B-v1.0)
<br>Took Roughly 3 Hours with 4x 4090s, over 2 Epochs, with around 52K varied samples.

Dataset Composition:
<br>20% - Coding
<br>30% - Instruct
<br>30% - Generalised Data
<br>10% - Roleplay
<br>10% - Dealignment

***

Testing Notes:

Fairly smart, as I expected. Obviously not at the level of the bigger models, but I did not expect that level from this.

Could be sampler issues, but generally I needed 1/2 swipes to get the correct answer when doing Zero context tests. If context is filled, no issues on my end.

For Roleplays: adding things like avoid writing as {{user}} suprisingly helps. Plus a proper prompt of course. I liked the writing style. Handles group characters in 1 card well, during my tests.

Fairly uncensored *during roleplay.* Yeah the as an AI stuff can happen at Zero context, but I have no issues once a character card is introduced. I had no issues making outputs that would give me 2500 Life Sentences if posted here.

***

Trained with Alpaca Format:

```
### Instruction:
<Prompt>

### Response:

```

OR

```
### Instruction:
<Prompt>

### Input:
<Insert Context Here>

### Response:

```

***

<br>wandb: 
<br>wandb: Run history:
<br>wandb:                      eval/loss █▃▂▂▂▂▂▁▁▁▁▂▂▂▂▂▂▁▁▁
<br>wandb:                   eval/runtime ▃▂▃▂▃▂▂▃▁▃█▂▃▃▃▂▃▃▂▂
<br>wandb:        eval/samples_per_second ▆▇▆▇▆▇▇▆█▆▁▇▆▆▆▇▆▆▇▇
<br>wandb:          eval/steps_per_second ▆▇▆▇▆▇▇▆█▆▁▇▆▆▆▇▆▆▇▇
<br>wandb:                    train/epoch ▁▁▁▂▂▂▂▂▂▃▃▃▃▃▄▄▄▄▄▄▅▅▅▅▅▅▆▆▆▆▆▇▇▇▇▇▇███
<br>wandb:              train/global_step ▁▁▁▂▂▂▂▂▂▃▃▃▃▃▄▄▄▄▄▄▅▅▅▅▅▅▆▆▆▆▆▇▇▇▇▇▇███
<br>wandb:            train/learning_rate ▄███████▇▇▇▇▇▆▆▆▆▅▅▅▅▄▄▄▃▃▃▃▂▂▂▂▂▁▁▁▁▁▁▁
<br>wandb:                     train/loss █▅▅▆▅▅▄▄▄▆▆▅▆▆▆▅▄▆▅▅▅▆▄▄▃▄▃▃▂▃▄▂▂▃▃▂▁▂▂▂
<br>wandb: 
<br>wandb: Run summary:
<br>wandb:                      eval/loss 0.74622
<br>wandb:                   eval/runtime 72.5049
<br>wandb:        eval/samples_per_second 37.239
<br>wandb:          eval/steps_per_second 2.331
<br>wandb:                    train/epoch 1.98
<br>wandb:              train/global_step 410
<br>wandb:            train/learning_rate 0.0
<br>wandb:                     train/loss 0.6457
<br>wandb:               train/total_flos 3.4382652340646707e+18
<br>wandb:               train/train_loss 0.70204
<br>wandb:            train/train_runtime 10880.917
<br>wandb: train/train_samples_per_second 9.417
<br>wandb:   train/train_steps_per_second 0.038
<br>wandb: 