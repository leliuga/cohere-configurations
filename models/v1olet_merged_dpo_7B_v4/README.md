---
license: apache-2.0
language:
- en
---

DPO from the model ranked *6th* on the overall leaderboard and **1st** in the 7B leaderboard - v1olet/v1olet_marcoroni-go-bruins-merge-7B.

Training data:
comparison_gpt4_en,distilabel-math-preference-dpo,orca_dpo_pairs

Train for 1 epoch.

You can use alpaca template.
```
template_format = """{system}
### Instruction:
{prompt}

### Response:
"""
```

Developed by: Trong-Hieu Nguyen-Mau
