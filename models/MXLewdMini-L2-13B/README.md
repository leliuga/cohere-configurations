---
license: cc-by-nc-4.0
---

Merge:
```shell
[Xwin (0.66) + ReMM (0.33)] x [Xwin (0.33) + MLewd (0.66)]
```

The goal was to recreate https://huggingface.co/Undi95/MXLewd-L2-20B in 13B without using merge interlacing (will probably be a little less good).
<!-- description start -->
## Models used

- Undi95/MLewd-L2-13B-v2-3
- Undi95/ReMM-v2.1-L2-13B
- Xwin-LM/Xwin-LM-13B-V0.1
<!-- description end -->

One part is ReMM (0.33) and Xwin (0.66)

One part is Xwin (0.33) and MLewd (0.66)
## Prompt template: Alpaca

```
Below is an instruction that describes a task. Write a response that completes the request.

### Instruction:
{prompt}

### Response:
```