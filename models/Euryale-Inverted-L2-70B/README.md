---
license: cc-by-nc-4.0
language:
- en
---

<img src="https://images7.alphacoders.com/921/921311.jpg"  style="width: 70%; min-width: 300px; display: block; margin: auto;">

An experimental merging of Several Models using 3 merging methods. Ties-Merge, BlockMerge_Gradient [& SLERP Variant] as well as SLERP.
5 Models included Inside, 2 LORAs.

***Early release because I'll be busy for the next month. Incomplete but workable, see below.***

***INVERT VARIANT***

SISTER MODEL HERE: [Euryale-L2-70B](https://huggingface.co/Sao10K/Euryale-L2-70B)

*Same concept as Stheno & Stheno Inverted, at different densities, weights and gradients.*

*Illustration for final gradient merge cannot be shown, each of the tensors had a different ratio applied to it.*


<br>Test Checklist:
<br>Censorship - NSFL no Issues
<br>Writing - Good Prose and Writing Quality
<br>NSFW - Yes
<br>IQ Level - Slightly dumber than Euryale. Even worse at Coding / Math though. Good for RP. As a general assistant? She's not the most accurate.
<br>Formatting - Markdown Formatting Issues, Able to Follow Statuses well.


<br>Most formats could work, but my tests have all been done in Alpaca format and it works well.
```
### Instruction:
Your instruction or question here.
For roleplay purposes, I suggest the following - Write <CHAR NAME>'s next reply in a chat between <YOUR NAME> and <CHAR NAME>. Write a single reply only.

### Response:
```


<br>My 7th Attempt. Incomplete so far, early release.
<br>Timeline Goals:
<br> Apply COT to model.
<br> Apply the RP LoRA I'm working on to model.
<br> Further tinker and test for potential better combinations.


Once Again, thanks to [Chargoddard](https://huggingface.co/chargoddard) and [Gryphe](https://huggingface.co/Gryphe) for their scripts, and @Vali for modifying some of the scripts provided to implement SLERP on a Cloud GPU Service.
Thanks to @gradientputri for partially sponsoring half of the compute costs for my Runpod usage.
Thanks to the original model creators too!

```
Art by wada_kazu / わだかず (pixiv page private?)
```

### LICENSE

License
This model is strictly non-commercial (cc-by-nc-4.0) use only which takes priority over the LLAMA 2 COMMUNITY LICENSE AGREEMENT. 

The "Model" is completely free (ie. base model, derivates, merges/mixes) to use for non-commercial purposes as long as the the included cc-by-nc-4.0 license in any parent repository, and the non-commercial use statute remains, regardless of other models' licences. 

*Non-Commercial due to parent models having the above license.*

Parent Models with this license:
<br>[garage-bAInd/Platypus2-70B-instruct](https://huggingface.co/garage-bAInd/Platypus2-70B-instruct)
<br>[elinas/chronos-70b-v2](https://huggingface.co/elinas/chronos-70b-v2)

### MODELS USED:

<br>[elinas/chronos-70b-v2](https://huggingface.co/elinas/chronos-70b-v2)
<br>[NousResearch/Nous-Hermes-Llama2-70b](https://huggingface.co/NousResearch/Nous-Hermes-Llama2-70b)
<br>[jondurbin/airoboros-l2-70b-2.1-creative](https://huggingface.co/jondurbin/airoboros-l2-70b-2.1-creative)
<br>[garage-bAInd/Platypus2-70B-instruct](https://huggingface.co/garage-bAInd/Platypus2-70B-instruct)
<br>[MayaPH/GodziLLa2-70B](https://huggingface.co/MayaPH/GodziLLa2-70B)

**LORAS**
<br>[nRuaif/fiction.live-Kimiko-V2-70B](https://huggingface.co/nRuaif/fiction.live-Kimiko-V2-70B)
<br>[lemonilia/limarp-llama2-v2](https://huggingface.co/lemonilia/limarp-llama2-v2)