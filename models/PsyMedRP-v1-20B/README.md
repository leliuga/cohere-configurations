---
license: cc-by-nc-4.0
tags:
- not-for-all-audiences
- nsfw
---

```
PsyMedRP-v1-13B-p1:
[jondurbin/airoboros-l2-13b-3.0](0.85) x [ehartford/Samantha-1.11-13b](0.15)

PsyMedRP-v1-13B-p2:
[Xwin-LM/Xwin-LM-13B-V0.1](0.85) x [chaoyi-wu/MedLLaMA_13B](0.15)

PsyMedRP-v1-20B-p1:
[PsyMedRP-v1-13B-p1](0.90) x [migtissera/Synthia-13B-v1.2](0.10)

PsyMedRP-v1-20B-p2:
[PsyMedRP-v1-13B-p2](0.90) x [migtissera/Synthia-13B-v1.2](0.10)

PsyMedRP-v1-20B-p3:
[Huginn merge with Gryphe gradient to PsyMedRP-v1-20B-p1]

PsyMedRP-v1-20B-p4:
[Huginn merge with Gryphe gradient to PsyMedRP-v1-20B-p2]

PsyMedRP-v1-20B-p5:
Apply Undi95/LimaRP-v3-120-Days at 0.3 weight to PsyMedRP-v1-20B-p3

PsyMedRP-v1-20B-p6:
Apply Undi95/LimaRP-v3-120-Days at 0.3 weight to PsyMedRP-v1-20B-p4

PsyMedRP-v1-20B:
layer_slices:
  - model: PsyMedRP-v1-20B-p5
    start: 0
    end: 16
  - model: PsyMedRP-v1-20B-p6
    start: 8
    end: 20
  - model: PsyMedRP-v1-20B-p5
    start: 17
    end: 32
  - model: PsyMedRP-v1-20B-p6
    start: 21
    end: 40
```

In testing.

If you want to support me, you can [here](https://ko-fi.com/undiai).