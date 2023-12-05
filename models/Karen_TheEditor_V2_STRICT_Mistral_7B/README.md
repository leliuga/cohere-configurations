---
tags:
- llm
- llama
- spellcheck
- grammar
license: llama2
---

<!-- header start -->
<div style="width: 100%;">
    <img src="https://huggingface.co/FPHam/Karen_TheEditor_V2_STRICT_Mistral_7B/resolve/main/karen2.jpg" alt="FPHam's Karen v2" style="width: 80%; min-width: 200px; display: block; margin: auto;">
</div>
<div style="display: flex; flex-direction: column; align-items: center;">
        <p><a href="https://ko-fi.com/Q5Q5MOB4M">Buy Karen Ko-fi</a></p>
    </div>
<!-- header end -->

# Karen is an editor for your text. (v.2) STRICT edition

Ah, Karen, a true peach among grammatical cucumbers! She yearns to rectify the missteps and linguistic tangles that infest your horribly written fiction.
Yet, unlike those ChatGPT kaboodles that morph into self-absorbed, constipated gurus of self-help style, Karen remains steadfastly grounded in grammatical wisdom but respectfull of your style.

# Info
Karen, Version 2, uses a completely different data set and base model than the previous Karen.

# There are two versions of Karen V2

1. Strict (this one), in which Karen will try not to make too many changes to your original text, mostly fixing grammar and spelling, assuming that you know what you are doing.
2. Creative ([here](https://huggingface.co/FPHam/Karen_TheEditor_V2_CREATIVE_Mistral_7B)), in which Karen may suggest slight contextual improvements or rephrasing where necessary. It's Karen, after a glass of wine.

# Goals

Karen's primary goal is to rectify grammatical and spelling errors in US English without altering the style of the text. She is adept at identifying and correcting common ESL errors.

    Verb Tense Errors:
        Incorrect use of verb tenses, such as using present tense when past tense is required and vice versa.
        Confusion between continuous and simple tenses.

    Subject-Verb Agreement:
        Lack of agreement between the subject and verb in number, e.g., using a singular verb with a plural subject or vice versa.

    Articles (a, an, the):
        Incorrect use or omission of articles, such as using "a" instead of "an" or vice versa.
        Overuse or omission of the definite article "the."

    Prepositions:
        Misuse of prepositions, such as using "in" instead of "on" or "at," or omitting prepositions where they are needed.

    Word Order:
        Incorrect word order in sentences, especially in questions and negative sentences.
        Misplacement of adverbs or adjectives.

    Pluralization:
        Incorrect plural forms of nouns, such as failing to add "-s" or "-es" when necessary.

    Pronoun Errors:
        Confusion between subject and object pronouns.
        Incorrect use of possessive pronouns.

    Double Negatives:
        Using double negatives, which is grammatically incorrect in standard English.

    Modal Verbs:
        Misuse of modal verbs like can, could, will, would, should, etc.

    Confusing Similar Words:
        Confusing words that sound similar but have different meanings and spellings (e.g., "their," "there," and "they're").

    Lack of Plural/Singular Agreement:
        Mistakes in matching singular and plural nouns and verbs in a sentence.

# Future Goals
Use bigger model, add grammar cases that the model misses. Better datasets. Use larger datasets.

# Training
It was reversely trained on fict/non-fiction US text where errors were intentionally inserted by another Llama model (Darth Karen) and Python script.

# Usage
It should be used by submitting a paragraph or block of text at a time.

# Model uses ChatML

```
<|im_start|>system
<|im_end|>
<|im_start|>user
Edit the following text for spelling and grammar mistakes: {paragraph of text} <|im_end|>
<|im_start|>assistant
```
Note the pretext: *Edit the following text for spelling and grammar mistakes:* before the actual text. This way Karen wouldn't start talking ABOUT the text.

# Recomended settings

- Temperature: 0.7
- top_p: 0.1
- top_k: 40
- repetition penalty: 1.18

Karen could also be used in chatting. But if Karen is involved in a conversation where the text is lengthy, she might interpret your message as needing proofreading assistance rather than simply acknowledging that you are using her in casual chat.

Example:

>Edit the following text for spelling and grammar mistakes: I see future writing more about you writing the story any way you can and LLM editing it and making it into a proper fiction. That seems to be more satisfying way where you could feel you actually had done something creATIVE. So human is the creative, LLM is the burger flipping intern.
>The opposite is what many people are trying rn. Make LLM to invent story (from premise, to creating outline to writing actual chapters) and the human to somehow assemble it from those pieces. So LLM is the creative man, and human is the floor polisher.
>I can't believe a normal person would find it satisfying long term.
>So in my view, the job of editor would be replaced by Ai, not the job of author.

(Grammarly Score: 48)

Output STRICT:

>I see the future of writing being more about you writing the story any way you can and having an AI like LLM edit it and make it into a proper fiction. That seems to be a more satisfying way where you could feel you actually did something creative. So humans are the creative ones, while LLMs are the burger-flipping interns.
>
>On the other hand, many people are currently trying to do the exact opposite. They want to make LLMs create stories (from premises, to outlining, to writing entire chapters). And then they want humans to somehow piece them together from these parts. In this scenario, LLMs become the creative individuals, and humans become the floor polishers.
>
>I can't imagine a regular person finding that satisfying over the long run.
>
>In my opinion, the role of an editor will eventually be taken over by AI, not the role of an author.

(Grammarly Score: 93)

Output CREATIVE:

>I see the future of writing being more focused on you writing the story any way you can and having an AI like GPT-3 edit it and turn it into a proper piece of fiction. This seems to be a more satisfying approach where you can truly feel that you have done something creative. In this scenario, humans remain the creative force, while the AI acts as the burger flipping intern.
>
>On the other hand, there are those who argue that we should let the AI generate stories from scratch - from coming up with a premise, to outlining the plot, to writing individual chapters. They believe that the role of the human should then become one of assembling these pieces together. In this case, the AI becomes the creative mind, and the human takes on the role of a floor polisher.
>
>I cannot imagine how someone could find this satisfying over the long term. Therefore, in my opinion, the job of the editor will eventually be taken over by AI, not the job of the author.

(Grammarly Score: 83)

# Conclusion

After probably 10 different versions with subsequent changes, I can now say that the current model works reasonably well, with occasional (but often debatable) grammar misses. The limitations seem to be related to the 7B parameters. It appears that the size isn't sufficient to have a fine-grained understanding of various nuances of the input. This correlates with my other findings - the Mistral model performs quite well when generating its own text, but its comprehension is less than perfect, again related to only 7B parameters.

The goal was to create a model that wouldn't change the style of the text. Often, LLM models, when asked to edit text, will attempt to rewrite the text even if the text is already fine. This proved to be quite challenging for such a small model where the main task was to determine the right balance between fixing the text (and not changing its style) and copying it verbatim.

The strict model assumes that you're already a good writer that doesn't need hand-holding and that every word you've written you've meant.