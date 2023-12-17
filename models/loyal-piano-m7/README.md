---
license: cc-by-nc-4.0
datasets:
- pankajmathur/orca_mini_v1_dataset
- openai/summarize_from_feedback
- PygmalionAI/PIPPA
- chargoddard/rpguild
- lemonilia/LimaRP
language:
- en
tags:
- mistral
---

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)

Experimenting with dataset ratios. Intended to be a roleplay-focused model with some smarts and good long-context recall.

Not sure if I've succeeded on the roleplay front, but something sure went right! Currently the #4 7B model on the leaderboard as of 11/30/2023. Going to riff on this and see where it goes.

| model | Average | ARC | HellaSwag | MMLU | TruthfulQA | Winogrande | GSM8K | DROP |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| fblgit/juanako-7b-UNA | 59.91 | 68.17 | 85.34 | 62.47 | 65.13 | 78.85 | 20.7 | 38.74 |
| Intel/neural-chat-7b-v3-1 | 59.06 | 66.21 | 83.64 | 62.37 | 59.65 | 78.14 | 19.56 | 43.84 |
| Weyaxi/OpenHermes-2.5-neural-chat-7b-v3-1-7B | 58.6 | 66.55 | 84.47 | 63.34 | 61.22 | 78.37 | 23.58 | 32.66 |
| **chargoddard/loyal-piano-m7** | 58.42 | 66.72 | 85.03 | 64.43 | 60.03 | 79.08 | 25.7 | 27.92 |
| Gryphe/MythoMist7b | 58.26 | 65.87 | 83.55 | 62.32 | 59.98 | 78.06 | 20.24 | 37.82 | 


Dataset composition:
| dataset | rows used | percent of total |
| --- | --- | --- |
| PIPPA | 14.6k | 43% |
| summarize_from_feedback | 9k | 26% |
| orca_mini_v1_dataset | 5.6k | 17% |
| rpguild | 2.86k | 8% |
| LimaRP | 2k | 6% |