---
license: other
---
**What is it?**
 Llama 2 13b expanded to the size of a Llama 1 33b model in certain areas, with the empty surrounding space filled with llama 33b data. (Base Model: https://huggingface.co/chargoddard/llama2-22b-blocktriangular) This is then finetuned on a 3090 by creating large loras and merging them. When I first started with 22b models, I looked for signs of knowledge transfer but didn't see it, so that's not a goal - the goal is just to throw lots of data at it until it adapts well to its surgically implanted parts. Datasets used are a mix of instruction, roleplay, and conversational data, often curated.



Continuation of https://huggingface.co/nkpz/llama2-22b-daydreamer-v2

Gave it an epoch of https://huggingface.co/datasets/jondurbin/airoboros-gpt4-1.4.1-mpt which it had never seen before - it had only seen Airoboros 2.0.
I excluded all coding rows from this dataset - a dedicated model is better for that imo. I wanted it to focus on answering well in English.

I thought this was pretty good improvement. Its ability to roleplay and imitate speech styles was enhanced and I feel like it's still sharpening up its problem solving skills. 

Its output style is right where I want it - dry and accurate enough when it's being asked for information, but has a lot character when it's answering fun prompts.

For input, it's flexible. It understands the Alpaca format well, but has been trained to handle a variety of input styles at this point.

I don't know what's next. I think I need to start benchmarking so I can more formally identify its weak points.

Examples edited with additional newlines to make it easier to read here.

**Example 1: Q&A**
```
Question: What is happening to the newspaper industry? Respond in a pirate accent.
Answer:
 The newspaper industry is being plundered by scurvy dogs and buccaneers like ye!
They're pillagin' yer booty with their fancy digital devices, like smartphones and tablets,
which make it easier for them to read online content instead of buyin' printed copies from us.
And don't forget about those pesky social media sites that are spreading rumors faster than the
winds on a galleon! So it's all hands on deck as we fight back against these scoundrels, or else
we'll be walkin' the plank soon enough! Avast ye landlubbers!
```

**Example 2: Alpaca Instruction**
```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
Angrily explain why pizza shouldn't be made with bad tomatoes.

### Response:

Listen, I know you think it's funny to make pizza with those rotten tomatoes, but let me tell you something -
bad tomatoes are not only disgusting and unhealthy, they also ruin the flavor of your pizza! You want a
delicious pie with fresh, juicy tomatoes that burst in your mouth, right?
Well, then, stop using those nasty tomatoes and stick to good quality ingredients for a truly satisfying pizza experience!
Trust me, once you try it, you won't go back! So please, don't even bother wasting your time making pizza with bad tomatoes
anymore; it's just not worth it. It makes me angry just thinking about it!
```

**Example 3: Chat / Roleplay**
```
You are speaking to Andre The Giant.
You: Hey, you're that wrestler guy, right?
Andre The Giant: Yes, I am indeed Andre the Giant. How can I help you?
You: So, how tall are you really? Like, what's your height in feet and inches?
Andre The Giant: Well, let's see... I stand at 7 ft 4 in (229 cm) tall.
You: Wow! That's so tall! And how much do you weigh?
Andre The Giant: I weigh around 500 lbs (227 kg). It's quite heavy, isn't it?
You: Yeah, I guess so. You must be strong too, considering your size.
Andre The Giant: Thank you for saying that. Strength is definitely important when dealing with my profession.
You: What made you decide to become a professional wrestler?
Andre The Giant: Well, I grew up watching wrestling on television, and I always admired the athleticism
and showmanship involved. When I got older, I decided to pursue it myself.
You: That's cool. Do you have any advice for aspiring wrestlers out there?
Andre The Giant: Absolutely! Always stay dedicated and focused, train hard, and never give up on yourself
or your dreams. Remember, anything is possible if you believe in yourself.
```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_nkpz__llama2-22b-daydreamer-v3)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 45.31   |
| ARC (25-shot)         | 56.06          |
| HellaSwag (10-shot)   | 80.07    |
| MMLU (5-shot)         | 52.49         |
| TruthfulQA (0-shot)   | 42.43   |
| Winogrande (5-shot)   | 73.48   |
| GSM8K (5-shot)        | 3.79        |
| DROP (3-shot)         | 8.85         |
