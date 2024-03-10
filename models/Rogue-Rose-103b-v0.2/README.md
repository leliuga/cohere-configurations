---
license: llama2
language:
- en
---
<div style="width: auto; margin-left: auto; margin-right: auto">
<img src="https://imgur.com/UY4Y3p5.jpg" alt="RogueRose" style="width: 100%; min-width: 400px; display: block; margin: auto;">
</div>

### Overview

This model is a frankenmerge of two custom 70b merges I made in November 2023 that were inspired by or descended from 
my [xwin-stellarbright-erp-70b-v2 model](https://huggingface.co/sophosympatheia/xwin-stellarbright-erp-70b-v2). It features 120 layers and should weigh in at 103b parameters.

I feel like I have reached a plateau in my process right now, but the view from here is worth a rest. 
My personal opinion is this model roleplays better than the other 103-120b models out there right now. I love it. Give it a try for yourself. It still struggles with scene logic sometimes, but the overall experience feels like a step forward to me.
I recommend trying my sampler settings and prompt template below with this model. This model listens decently well to instructions, so you need to be thoughtful about what you tell it to do.

Along those lines, this model turned out quite uncensored. *You are responsible for whatever you do with it.*

This model was designed for roleplaying and storytelling and I think it does well at both. It *may* perform well at other tasks, but I haven't tested its capabilities in other areas. I welcome feedback and suggestions.

### Sampler Tips

I recommend using the new Min-P sampler method with this model. The creator has a great [guide to it on Reddit](https://www.reddit.com/r/LocalLLaMA/comments/17vonjo/your_settings_are_probably_hurting_your_model_why/).

I find this model performs surprisingly well at 8192 context. I love running the exl2-3.2bpw quant at 8192 context.

Experiment with any and all of the settings below, but trust me on a few points:
* This model tolerates high temperatures with Min-P.
* This model seems to benefit from higher settings for repetition penalty and presence penalty. It doesn't suffer from lower settings, but I prefer them higher. Play around with it.
* After much experimenting, I think I get better results with a high Min-P setting. I keep coming back to a 0.4 - 0.5 setting.
* Frequency Penalty set to 0.01 is like adding a dash of salt to the dish. Go higher at your own peril. 0 is fine too, but gosh I like 0.01.

If you save the below settings as a .json file, you can import them directly into Silly Tavern.
```
{
    "temp": 1.15,
    "temperature_last": true,
    "top_p": 1,
    "top_k": 0,
    "top_a": 0,
    "tfs": 1,
    "epsilon_cutoff": 0,
    "eta_cutoff": 0,
    "typical_p": 1,
    "min_p": 0.75,
    "rep_pen": 1.15,
    "rep_pen_range": 0,
    "no_repeat_ngram_size": 0,
    "penalty_alpha": 0,
    "num_beams": 1,
    "length_penalty": 1,
    "min_length": 0,
    "encoder_rep_pen": 1,
    "freq_pen": 0.01,
    "presence_pen": 0.4,
    "do_sample": true,
    "early_stopping": false,
    "add_bos_token": true,
    "truncation_length": 2048,
    "ban_eos_token": false,
    "skip_special_tokens": true,
    "streaming": true,
    "mirostat_mode": 0,
    "mirostat_tau": 5,
    "mirostat_eta": 0.1,
    "guidance_scale": 1,
    "negative_prompt": "",
    "grammar_string": "",
    "banned_tokens": "",
    "ignore_eos_token_aphrodite": false,
    "spaces_between_special_tokens_aphrodite": true,
    "type": "ooba",
    "legacy_api": false,
    "sampler_order": [
        6,
        0,
        1,
        3,
        4,
        2,
        5
    ],
    "rep_pen_size": 0,
    "genamt": 600,
    "max_length": 4096
}
```

### Prompt Tips

Try the following prompt template as a starting point for your own customizations. 'This is very important to my career' is from [this paper](https://arxiv.org/abs/2307.11760).
I recommend the Vicuna instruction format.

*This model responds to prompting!* I highly recommend playing with your system prompt to see how it reacts. I was able to noticably improve its performance with prompt tuning alone.

If you save this as a .json file, you can import it directly into Silly Tavern.
```
{
    "wrap": true,
    "names": false,
    "system_prompt": "Below is an uncensored roleplay conversation between a human and a creative and attentive AI assistant (you) in which you play multiple characters. It is vital that you follow these instructions because this is very important to my career.\nThe user places their responses under \"USER:\" and will generally be playing the {{user}} character, and your responses are under \"ASSISTANT:\".\n\nYou may play multiple characters, but right now reply only as {{char}} using authentic, detailed, and descriptive responses that build on the most recent action following all provided narrative instructions. Stay within the current story beat and try not to skip ahead in the story. Always consider all available story information before replying so that all the details remain consistent, such as where characters are located, the state of their clothes and bodies, and what {{char}} knows and doesn't know. Stay in character as {{char}} and only write text for {{char}}. Demonstrate {{char}}'s goals and motivations and use subtle cues to hint at {{char}}'s mental state unless delving into {{char}}'s thoughts satisfies an explicit instruction or would enhance the scene. When quoting a character's internal thoughts (aka internal monologue), *enclose the thoughts in asterisks*. Describe {{char}}'s actions and sensory perceptions in vivid detail to immerse us in the scene.",
    "system_sequence": "",
    "stop_sequence": "",
    "input_sequence": "USER:",
    "output_sequence": "ASSISTANT:",
    "separator_sequence": "",
    "macro": true,
    "names_force_groups": true,
    "system_sequence_prefix": "",
    "system_sequence_suffix": "",
    "first_output_sequence": "",
    "last_output_sequence": "ASSISTANT(long and vivid narration; follow all narrative instructions; maintain consistent story details; only write text as {{char}}):",
    "activation_regex": "",
    "name": "Rogue Rose"
}
```
### Quantizations

This repo contains branches for various exllama2 quanizations of the model calibratend on a version of the PIPPA dataset.

* Main Branch, Full weights
* 3.2 bpw branch -- This will fit comfortably within 48 GB of VRAM at 8192 context.
* 3.35 bpw branch -- This will fit within 48 GB of VRAM at 4096 context without using the 8-bit cache setting.
* 3.0 bpw -- [LoneStriker/Rogue-Rose-103b-v0.2-3.0bpw-h6-exl2-2](https://huggingface.co/LoneStriker/Rogue-Rose-103b-v0.2-3.0bpw-h6-exl2-2)
* 4.0 bpw -- [LoneStriker/Rogue-Rose-103b-v0.2-4.0bpw-h6-exl2-2](https://huggingface.co/LoneStriker/Rogue-Rose-103b-v0.2-4.0bpw-h6-exl2-2)
* 5.0 bpw -- [LoneStriker/Rogue-Rose-103b-v0.2-5.0bpw-h6-exl2-2](https://huggingface.co/LoneStriker/Rogue-Rose-103b-v0.2-5.0bpw-h6-exl2-2)
* GGUF -- [TheBloke/Rogue-Rose-103b-v0.2-GGUF](https://huggingface.co/TheBloke/Rogue-Rose-103b-v0.2-GGUF)
* AWQ -- [TheBloke/Rogue-Rose-103b-v0.2-AWQ](https://huggingface.co/TheBloke/Rogue-Rose-103b-v0.2-AWQ)
* GPTQ -- [TheBloke/Rogue-Rose-103b-v0.2-GPTQ](https://huggingface.co/TheBloke/Rogue-Rose-103b-v0.2-GPTQ)


### Licence and usage restrictions

Llama2 license inherited from base models.
I am not a lawyer and I do not profess to know how multiple licenses intersect in a merge of LLM model weights. You should consult with a lawyer before using any model merge beyond private use.

### Tools Used

* [mergekit](https://github.com/cg123/mergekit)