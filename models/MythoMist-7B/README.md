---
license: other
language:
- en
---
MythoMist 7b is, as always, a highly experimental Mistral-based merge based on my latest algorithm, which actively benchmarks the model as it's being built in pursuit of a goal set by the user.

**Addendum (2023-11-23)**: A more thorough investigation revealed a flaw in my original algorithm that has since been resolved. I've considered deleting this model as it did not follow its original objective completely but since there are plenty of folks enjoying it I'll be keeping it around. Keep a close eye [on my MergeMonster repo](https://huggingface.co/Gryphe/MergeMonster) for further developments and releases of merges produced by the Merge Monster.

The primary purpose for MythoMist was to reduce usage of the word anticipation, ministrations and other variations we've come to associate negatively with ChatGPT roleplaying data. This algorithm cannot outright ban these words, but instead strives to minimize the usage.

[The script has now been made available on my Github. Warning - Plenty of VRAM is needed.](https://github.com/Gryphe/MergeMonster/)

Quantized models are available from TheBloke: [GGUF](https://huggingface.co/TheBloke/MythoMist-7B-GGUF) - [GPTQ](https://huggingface.co/TheBloke/MythoMist-7B-GPTQ) - [GPTQ](https://huggingface.co/TheBloke/MythoMist-7B-AWQ) (You're the best!)

## Final merge composition

After processing 12 models my algorithm ended up with the following (approximated) final composition:

| Model                    | Contribution |
|--------------------------|--------------|
| Neural-chat-7b-v3-1      | 26%          |
| Synatra-7B-v0.3-RP       | 22%          |
| Airoboros-m-7b-3.1.2     | 10%          |
| Toppy-M-7B               | 10%          |
| Zephyr-7b-beta           | 7%           |
| Nous-Capybara-7B-V1.9    | 5%           |
| OpenHermes-2.5-Mistral-7B| 5%           |
| Dolphin-2.2.1-mistral-7b | 4%           |
| Noromaid-7b-v0.1.1       | 4%           |
| SynthIA-7B-v1.3          | 3%           |
| Mistral-7B-v0.1          | 2%           |
| Openchat_3.5             | 2%           |

There is no real logic in how these models were divided throughout the merge - Small bits and pieces were taken from each and then mixed in with other models on a layer by layer basis, using a pattern similar to my MythoMax recipe in which underlying tensors are mixed in a criss-cross manner.

This new process only decides on the model's layers, not the singular lm_head and embed_tokens layers which influence much of the model's output. I ran a seperate script for that, picking the singular tensors that resulted in the longest responses, which settled on Toppy-M-7B.

## Prompt Format

Due to the wide variation in prompt formats used in this merge I (for now) recommend using Alpaca as the prompt template for compatibility reasons:
```
### Instruction:
Your instruction or question here.

### Response:
```

---
license: other
---