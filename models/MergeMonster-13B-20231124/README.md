---
license: other
language:
- en
---
I received a request to upload the FP16 version of my first [MergeMonster](https://github.com/Gryphe/MergeMonster/tree/main) 13b merge, which had the sole focus of reducing overused ChatGPT roleplaying phrases. Many ministrations were murdered throughout this process.

GGUF quantizations can be found in my main [MergeMonster](https://huggingface.co/Gryphe/MergeMonster) repo.

## Merge Composition

Unsurprisingly, not much remained of MythoMax. I'm (obviously!) very fond of it, but it's not perfect.

Reminder: MergeMonster works with hundreds of pieces scattered throughout the final merge. If a model takes a large part of the composition this does not mean it is merged in a manner that might be logical to humans.
```
Base model: MythoMax-L2-13b

------ FINAL MERGE COMPOSITION ------
LLaMA2-13B-Tiefighter: 0.27
Emerhyst-13B: 0.26
Orca-2-13b: 0.24
Stheno-1.8-L2-13B: 0.09
MythoMax-L2-13b: 0.09
Athena-v4: 0.04
```

## Prompt Format

As always, I suggest sticking to the Alpaca format.

```
### Instruction:
Your instruction or question here.

### Response:
```