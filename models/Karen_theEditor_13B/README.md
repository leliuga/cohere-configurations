---
tags:
- lora
---
<!-- header start -->
<div style="width: 100%;">
    <img src="https://media.tenor.com/frGCmLDFbkMAAAAC/karen-ok.gif" alt="FPHam's Karen" style="width: 30%; min-width: 200px; display: block; margin: auto;">
</div>
<div style="display: flex; flex-direction: column; align-items: center;">
        <p><a href="https://ko-fi.com/Q5Q5MOB4M">Buy Karen Ko-fi</a></p>
    </div>
<!-- header end -->

## Karen is an editor for your fiction. (v.0.2)

Ah, Karen, a true peach among grammatical cucumbers! She yearns to rectify the missteps and linguistic tangles that infest your horribly written fiction.  
Yet, unlike those ChatGPT kaboodles that morph into self-absorbed, constipated gurus of self-help style, Karen remains steadfastly grounded in wit and wisdom but respectfull of your style.

She is also absolute joy to chat with, although she may correct grammar in your chats too from time to time 
(As certain well known LLM said, "She is a radiant beacon of amusement")

She also has a particular soft spot for Llamas.

## Quantized Karen version (Quantized by TheBloke)

* [4-bit GPTQ models for GPU inference](https://huggingface.co/FPHam/Karen_theEditor-13B-4bit-128g-GPTQ)
* [4-bit, 5-bit and 8-bit GGML models for CPU(+GPU) inference](https://huggingface.co/TheBloke/Karen_theEditor_13B-GGML)

Karen gets triggered by this prompt (pun intended):

```
USER: Edit the following for spelling and grammar mistakes: 
ASSISTANT:
```

Feed it a paragraph at a time - that's where she works best - a few sentences.

Here is my instruct yaml for oobabooga webui (in Instruct mode):

```
user: "USER: Edit the following for spelling and grammar mistakes: "
bot: "ASSISTANT:"
turn_template: "<|user|> <|user-message|>\n<|bot|> <|bot-message|></s>\n"
context: "Assistant is a meticulous senior editor with a specialization in editing fictional stories. When given text, Assistant detects and corrects grammatical errors, including subject-verb agreement, tense consistency, punctuation, capitalization, use of correct articles and correct present perfect and past perfect tense.\n"
```

I am not sure to what extend the context part has any effect. Doesn't hurt though.

This is a version 0.2 and finetuned just so it retains the underlaying Vicuna qualities. 

## Example:

Input
```
USER: Edit the following for spelling and grammar mistakes:
unce upon time lived name John who have pet monkey Jackie.
one day while walkin down the street he saw his friend Tom coming to him.
"Hey Tom" said John "hows it going?"

Tom replyied "Not too bad how 'bout yourself?". as they was talk Jackie sudden jump onto Tom shoulder and started play with his hair.
Tom look at John and ask "Is that yours?"

John reply "Yeah, this crazy lil guy follow me everywhere" just then Jackie grab hold of Tom glasses and try to take thmm off.
Tom struggle to keep his balance as he laugh hystericaly.
```
Output
```
ASSISTANT: Once upon a time there lived a man named John who had a pet monkey called Jackie.
One day while walking down the street he saw his friend Tom approaching him.
"Hey Tom," said John. "How's it going?"

Tom replied, "Not too bad, how about yourself?" As they were talking, Jackie suddenly jumped onto Tom's shoulder and began playing with his hair.
Tom looked at John and asked, "Is that yours?"

John replied, "Yeah, this crazy little guy follows me everywhere." Just then Jackie grabbed hold of Tom's glasses and tried to take them off.
Tom struggled to keep his balance as he laughed hysterically.
```

## Goal: to create the best grammar checker you have ever seen

## To do:
- train on larger dataset, big, enormous, gargantuan 
- see if finetuning on just plain LLAMA without Vicuna would work better or worse (the theory is that it will be very focused on editing and nothing else)
- explore what different settings (temperature, top_p, top_k do for this type of finetune)
- create Rachel, the paraphrasing editor
