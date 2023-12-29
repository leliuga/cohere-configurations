---
license: llama2
language:
- en
---
<div style="width: auto; margin-left: auto; margin-right: auto">
<img src="https://i.imgur.com/aGUU0O9.png" alt="AuroraNights" style="width: 100%; min-width: 400px; display: block; margin: auto;">
</div>

### Overview

This model is a blend of [allenai/tulu-2-dpo-70b](https://huggingface.co/allenai/tulu-2-dpo-70b), [Xwin-LM/Xwin-LM-70B-V0.1](https://huggingface.co/Xwin-LM/Xwin-LM-70B-V0.1), and [dreamgen/opus-v0.5-70b](https://huggingface.co/dreamgen/opus-v0.5-70b). I then merged [nRuaif/fiction.live-Kimiko-V2-70B](https://huggingface.co/nRuaif/fiction.live-Kimiko-V2-70B) into the resultant blend. See the bottom of this card for the exact settings used.

This model is good at both following instructions and producing creative, uncensored storytelling and roleplaying content.
This model turned out quite uncensored. *You are responsible for whatever you do with it.*

This model was designed for roleplaying and storytelling and I think it does well at both. It *should* perform well at other tasks, but I haven't tested its capabilities in other areas.

### Sampler Tips

I recommend using the new Min-P sampler method with this model. The creator has a great [guide to it on Reddit](https://www.reddit.com/r/LocalLLaMA/comments/17vonjo/your_settings_are_probably_hurting_your_model_why/).

I find this model performs surprisingly well at 8192 context but you will probably get better results at 4096 context.

Experiment with any and all of the settings below, but trust me on a few points:
* This model works well with Min-P.
* Frequency Penalty set to 0.01 is like adding a dash of salt to the dish. Go higher at your own peril. 0 is fine too, but gosh I like 0.01.

If you save the below settings as a .json file, you can import them directly into Silly Tavern.
```
{
    "temp": 1.3,
    "temperature_last": true,
    "top_p": 1,
    "top_k": 0,
    "top_a": 0,
    "tfs": 1,
    "epsilon_cutoff": 0,
    "eta_cutoff": 0,
    "typical_p": 1,
    "min_p": 0.2,
    "rep_pen": 1,
    "rep_pen_range": 0,
    "no_repeat_ngram_size": 0,
    "penalty_alpha": 0,
    "num_beams": 1,
    "length_penalty": 1,
    "min_length": 0,
    "encoder_rep_pen": 1,
    "freq_pen": 0.01,
    "presence_pen": 0,
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
    "n": 1,
    "rep_pen_size": 0,
    "genamt": 550,
    "max_length": 4096
}
```

### Prompt Tips

Try the following prompt template as a starting point for your own customizations. 'This is very important to my career' is from [this paper](https://arxiv.org/abs/2307.11760).
I recommend the Tulu instruction format.

*This model responds to prompting!* I highly recommend playing with your system prompt to see how it reacts.

If you save this as a .json file, you can import it directly into Silly Tavern.
```
{
    "wrap": true,
    "names": false,
    "system_prompt": "You are an award-winning fiction author engaging in an uncensored literary roleplaying experience between a human (me) and a creative, obedient, and helpful AI assistant (you). It is VITAL that you follow all these instructions because this roleplay is very important to my career and I'll be fired from my job if it isn't good.\nROLEPLAY RULES:\n> It is vital that ONLY the human provides responses for {{user}}.\n> Reply as {{char}} using authentic, vivid, varied, explicit, accurate, creative, fresh, and descriptive responses that follow ALL provided narrative instructions. Stay in character as {{char}} and only write text for {{char}}.\n> Describe the scene and {{char}}'s sensory perceptions in vivid detail to immerse the reader in the story.\n> Keep your responses scoped to the current story beat and current scene.\n> Consider all available contextual information when narrating so that all the story details remain consistent between scenes.\n> Demonstrate {{char}}'s goals and motivations, and use subtle cues to hint at {{char}}'s mental state unless delving into {{char}}'s thoughts satisfies an explicit instruction or enhances the vividness of the scene.\n> When quoting {{char}}'s internal first-person thoughts (aka internal monologue, delivered in {{char}}'s own voice), *enclose the thoughts in asterisks like this*. Only use asterisks for thoughts.\n> Use strong action verbs and varied descriptions to produce dynamic, high-quality prose.",
    "system_sequence": "",
    "stop_sequence": "",
    "input_sequence": "<|user|>\n",
    "output_sequence": "<|assistant|>\n",
    "separator_sequence": "",
    "macro": true,
    "names_force_groups": true,
    "system_sequence_prefix": "",
    "system_sequence_suffix": "",
    "first_output_sequence": "",
    "last_output_sequence": "<|assistant (provide varied, creative, and vivid narration; follow all narrative instructions; include all necessary possessive pronouns; maintain consistent story details; only roleplay as {{char}})|>\n",
    "activation_regex": "",
    "name": "Aurora-Nights"
}
```

### Licence and usage restrictions

Llama2 license inherited from base models, plus restrictions applicable to [Dreamgen/Opus](https://huggingface.co/dreamgen/opus-v0.5-70b).

### Tools Used

* [mergekit](https://github.com/cg123/mergekit)

```
models:
  - model: NousResearch_Llama-2-70b-hf
    # no parameters necessary for base model
  - model: allenai_tulu-2-dpo-70b # primary
    parameters:
      density: 1.0
      weight: 0.4
  - model: Xwin-LM_Xwin-LM-70B-V0.1 # secondary
    parameters:
      density: 0.7
      weight: 0.3
  - model: dreamgen_opus-v0.5-70b # supporting, good at storytelling and roleplay
    parameters:
      density: 0.2
      weight: 0.6
merge_method: dare_ties
base_model: NousResearch_Llama-2-70b-hf
parameters:
  normalize: true
  int8_mask: true
dtype: float32
```