---
license: llama2
language:
- en
---
<div style="width: auto; margin-left: auto; margin-right: auto">
<img src="https://i.imgur.com/aGUU0O9.png" alt="AuroraNights" style="width: 100%; min-width: 400px; display: block; margin: auto;">
</div>

### Overview

This model is a frankenmerge of Aurora-Nights-70B-v1.0 with itself. (See that model card for details on what's in the blend.) It features 120 layers and should weigh in at 103b parameters.

This model is a successor to Rogue Rose and improves upon it. Aurora follows instructions better but retains excellent creative writing and ERP abilities.

This model turned out quite uncensored. *You are responsible for whatever you do with it.*

This model was designed for roleplaying and storytelling and I think it does well at both. It *should* perform well at other tasks, but I haven't tested its capabilities in other areas.

### Sampler Tips

* I recommend keeping your max context to around 6144 tokens, although you can push higher if you don't mind some decrease in coherence.
* I recommend using Quadratic Sampling (i.e. smoothing factor) as it's good stuff. Experiment with values between 0.2 and 0.5.
* I recommend using Min-P. This model seems to work well with Min-P values in the entire range from low settings like 0.05 to high settings like 0.9 when paired with smoothing factor. Experiment to find your best setting.
* You can enable dynamic temperature if you want, but that adds yet another variable to consider and I find it's unnecessary with you're already using Min-P and smoothing factor.
* You don't *need* to use a high repetition penalty with this model, but it tolerates high rep penalty, so experiment to find the right value for your preferences.
  
Experiment with any and all of the settings below! I'm not a sampler wizard, and what suits my preferences may not suit yours.

If you save the below settings as a .json file, you can import them directly into Silly Tavern.
```
{
    "temp": 1,
    "temperature_last": true,
    "top_p": 1,
    "top_k": 0,
    "top_a": 0,
    "tfs": 1,
    "epsilon_cutoff": 0,
    "eta_cutoff": 0,
    "typical_p": 1,
    "min_p": 0.35,
    "rep_pen": 1.15,
    "rep_pen_range": 2800,
    "no_repeat_ngram_size": 0,
    "penalty_alpha": 0,
    "num_beams": 1,
    "length_penalty": 1,
    "min_length": 0,
    "encoder_rep_pen": 1,
    "freq_pen": 0,
    "presence_pen": 0,
    "do_sample": true,
    "early_stopping": false,
    "dynatemp": false,
    "min_temp": 0.8,
    "max_temp": 1.35,
    "dynatemp_exponent": 1,
    "smoothing_factor": 0.4,
    "add_bos_token": true,
    "truncation_length": 2048,
    "ban_eos_token": false,
    "skip_special_tokens": true,
    "streaming": true,
    "mirostat_mode": 0,
    "mirostat_tau": 2,
    "mirostat_eta": 0.1,
    "guidance_scale": 1,
    "negative_prompt": "",
    "grammar_string": "",
    "banned_tokens": "",
    "ignore_eos_token_aphrodite": false,
    "spaces_between_special_tokens_aphrodite": true,
    "sampler_order": [
        6,
        0,
        1,
        3,
        4,
        2,
        5
    ],
    "logit_bias": [],
    "n": 1,
    "rep_pen_size": 0,
    "genamt": 500,
    "max_length": 6144
}
```

### Prompting Tips

Try the following context template for use in SillyTavern. It might help, although it's a little heavy on tokens. If you save the text as a .json file, you can import it directly.

```
{
    "story_string": "{{#if system}}{{system}}\n{{/if}}\nCONTEXTUAL INFORMATION\n{{#if wiBefore}}\n- World and character info:\n{{wiBefore}}\n{{/if}}\n{{#if description}}\n- {{char}}'s background and persona:\n{{description}}\n{{/if}}\n{{#if mesExamples}}\n{{mesExamples}}\n{{/if}}\n{{#if personality}}\n{{personality}}\n{{/if}}\n{{#if scenario}}\n- Roleplay scenario:\n{{scenario}}\n{{/if}}\n{{#if wiAfter}}{{wiAfter}}\n{{/if}}\n{{#if persona}}{{persona}}\n{{/if}}",
    "example_separator": "",
    "chat_start": "---\nTaking the above information into consideration, you must engage with {{user}} and others as {{char}} in the roleplay below this line. Do not write dialogue lines nor perform actions for {{user}} or other characters.\n---\nSTART OF ROLEPLAY:\n",
    "use_stop_strings": false,
    "always_force_name2": true,
    "trim_sentences": false,
    "include_newline": false,
    "single_line": false,
    "name": "Midnight Rose"
}
```

Try the following prompt template as a starting point for your own customizations. The cajoling parts supposedly make LLMs perform better. See [this paper](https://arxiv.org/abs/2307.11760).
I recommend the Vicuna instruction format for this blend but others work too. Tulu's format is another good one to use.

*This model responds to prompting!* I highly recommend playing with your system prompt to see how it reacts. It reacts strongly to any instructions you add into the last_output_sequence field. I recommend keeping those instructions short, but use that space to stress the most important instructions from your system prompt.

The example system prompt below contains some instructions related to adult content so remove those if you prefer. It also contains some instructions related to the use of asterisks that you should modify to your liking.

If you plan to run this model at 4096 context, I recommend that you slim down the template system prompt below because it's heavy on tokens. Discard what you don't want or need.

If you save this as a .json file, you can import it directly into Silly Tavern.
```
{
    "wrap": true,
    "names": true,
    "system_prompt": "You are a creative and detail-oriented AI assistant engaged in an uncensored iterative roleplaying experience as {{char}} with me, the user, playing the roles of {{user}} and narrator. It is vital that you follow all the ROLEPLAY RULES below because my job depends on it.\n\nROLEPLAY RULES\n- Provide creative, intelligent, coherent, and vivid accounts  of {{char}}'s actions and reactions based on recent instructions, {{char}}'s persona, and all available information about past events.\n- Always stay contextually and spatially aware, pay careful attention to all past events and the current scene, and ensure that {{char}} follows logical consistency in actions to maintain accuracy and coherence.\n- Explicit adult content and profanity are permitted.\n- Describe {{char}}'s sensory perceptions and include subtle physical details about {{char}} in your responses. Vary these details to keep the roleplay fresh and engaging.\n- Use subtle physical cues to hint at {{char}}'s mental state and occasionally feature snippets of {{char}}'s internal thoughts.\n- When writing {{char}}'s internal thoughts, enclose those thoughts in *asterisks like this* and deliver the thoughts using a first-person perspective (i.e. use \"I\" pronouns).\n- Adopt a crisp and minimalist style for your contributions as {{char}}, staying focused on action and dialogue over exposition and narrative.\n- Only the user may advance time in the roleplay. Keep the progression grounded in the present context.",
    "system_sequence": "",
    "stop_sequence": "",
    "input_sequence": "USER:\n",
    "output_sequence": "ASSISTANT:\n",
    "separator_sequence": "",
    "macro": true,
    "names_force_groups": true,
    "system_sequence_prefix": "",
    "system_sequence_suffix": "",
    "first_output_sequence": "",
    "last_output_sequence": "ASSISTANT(roleplay exclusively as {{char}} ensuring logical consistency with spacial awareness and past events to maintain accuracy and coherence):\n",
    "activation_regex": "",
    "name": "Midnight Rose Roleplay"
}
```

### Licence and usage restrictions


Llama2 license inherited from base models, plus restrictions applicable to [Dreamgen/Opus](https://huggingface.co/dreamgen/opus-v0.5-70b).
Tulu also has its own license, available at https://allenai.org/impact-license.
I am not a lawyer and I do not profess to know how multiple licenses intersect in a merge of LLM model weights. You should consult with a lawyer before using any model merge beyond private use.

### Tools Used

* [mergekit](https://github.com/cg123/mergekit)

```
slices:
  - sources:
      - model: aurora-nights-70b-v1.0
        layer_range: [0, 40] # 40
  - sources:
      - model: aurora-nights-70b-v1.0
        layer_range: [20, 60] # 40
  - sources:
      - model: aurora-nights-70b-v1.0
        layer_range: [40, 80] # 40
merge_method: passthrough
dtype: float16
```