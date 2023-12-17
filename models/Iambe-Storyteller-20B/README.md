---
license: cc-by-nc-4.0
---
<p align="center"><font size="7"> <b>Dumb assistant, 🔥 author.</b></font></p>
<p align="center"><font size="5"> <b>4-bit Examples with Alpaca</b></font></p>
<p align="center"><img src="https://i.postimg.cc/VNXpwDCY/Screenshot-2023-12-06-212944.png"/>
<p align="center"><img src="https://i.ibb.co/TPDsD2z/Screenshot-2023-12-06-211216.png"/>
<p align="center"><font size="5"><b><a href="https://i.ibb.co/4PhnXjK/Screenshot-2023-12-06-205831.png">!!NSFW!! - 🔥Erotica Writing Example🔥 - !!NSFW!!</font></a></b></p>

Thanks to [Charles Goddard](https://huggingface.co/chargoddard) for the recipe.

The idea here is to "move" Iambe from being based on vanilla L2 to being based on sequelbox/DynamicFactor instead.

Because task_arithmetic uses the raw deltas, this should be similar to if the SFT had been done over DynamicFactor.

## Recipe

merge_method: task_arithmetic

base_model: athirdpath/BigLlama-20b-v1.1      # Base model you want to "move out" from

models:

  - model: athirdpath/Iambe-20b-DARE-v2       # SFTd model you want to transfer

  - model: athirdpath/DoubleFactor-20b      # Base model you want to "move in" to

parameters:

  - weight: 1.0

dtype: bfloat16