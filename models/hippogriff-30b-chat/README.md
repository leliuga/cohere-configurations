---
datasets:
- QingyiSi/Alpaca-CoT
- teknium/GPT4-LLM-Cleaned
- teknium/GPTeacher-General-Instruct
- metaeval/ScienceQA_text_only
- hellaswag
- openai/summarize_from_feedback
- riddle_sense
- gsm8k
- OpenAssistant/oasst1
language:
- en
library_name: transformers
pipeline_tag: text-generation
---

# Hippogriff 30B Chat

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)


Hippogriff 30B Chat is an experiment that builds on Manticore with new datasets, while removing a few more instruction and chat datasets. It also includes a de-duped subset of the Pygmalion dataset. It also removes all Alpaca style prompts using `###` in favor of 
chat only style prompts using `USER:`,`ASSISTANT:` as well as [pygmalion/metharme prompting](https://huggingface.co/PygmalionAI/metharme-7b#prompting) using `<|system|>, <|user|> and <|model|>` tokens.

Questions, comments, feedback, looking to donate, or want to help? Reach out on our [Discord](https://discord.gg/PugNNHAF5r) or email [wing@openaccessaicollective.org](mailto:wing@openaccessaicollective.org)

# Training Datasets

Hippogriff 30B Chat is a Llama 30B model fine-tuned on the following datasets

- OpenAssistant/oasst1 - cleaned dataset, similar to Guanaco
- synthetic jokes generation and explanation derived from reddit jokes dataset
- synthetic prose generation and rewriting self-chat 
- Q&A based on provided context
- self instruct augmented logic_inference_oa
- de-duped pygmalion dataset, filtered down to RP data, cleaned, english only, 25%
- [riddle_sense](https://huggingface.co/datasets/riddle_sense) - instruct augmented 
- hellaswag, updated for detailed explanations w 30K+ rows
- [gsm8k](https://huggingface.co/datasets/gsm8k) - instruct augmented 
- [ewof/code-alpaca-instruct-unfiltered](https://huggingface.co/datasets/ewof/code-alpaca-instruct-unfiltered) synthetic self chat dataset derived from about 1000 rows
- [subset of QingyiSi/Alpaca-CoT for roleplay and CoT](https://huggingface.co/QingyiSi/Alpaca-CoT)
- [GPTeacher-General-Instruct](https://huggingface.co/datasets/teknium/GPTeacher-General-Instruct)
- ARC-Easy & ARC-Challenge - instruct augmented for detailed responses, derived from the `train` split
- [hellaswag](https://huggingface.co/datasets/hellaswag) - 5K row subset of instruct augmented for concise responses, derived from the `train` split
- [metaeval/ScienceQA_text_only](https://huggingface.co/datasets/metaeval/ScienceQA_text_only) - instruct for concise responses
- [openai/summarize_from_feedback](https://huggingface.co/datasets/openai/summarize_from_feedback) - instruct augmented tl;dr summarization

Hippogriff differs from Manticore as it does not use the WizardLM, WizardVicuna, Alpaca, or ShareGPT datasets.

# Initial thoughts

Hippogriff follows instructions pretty well. It still struggles with anything that has to do with math. Prose is much better than manticore. Handles in-context QA much better.


# Shoutouts

Special thanks to Nanobit for helping with Axolotl, TheBloke for quantizing these models are more accessible to all, 0x000011b for the RP dataset, and the support from everyone in our AI Discord community.

# Demo

A Spaces demo is not provided for this release due to 30B models currently not fitting in VRAM.

## Build

Hippogriff was built with [Axolotl](https://github.com/OpenAccess-AI-Collective/axolotl) on 8xA100 80GB 
 - 1.5 epochs taking approximately 12 hours. No further epochs will be released. 

## Bias, Risks, and Limitations

Hippogriff has not been aligned to human preferences with techniques like RLHF or deployed with in-the-loop filtering of responses like ChatGPT, so the model can produce problematic outputs (especially when prompted to do so).
Hippogriff was fine-tuned from the base model LlaMa 30B, please refer to its model card's Limitations Section for relevant information.

The core functionality of Llama is taking a string of text and predicting the next token. While language models are widely used for tasks other than this, 
there are a lot of unknowns with this work. When prompting Hippogriff it is important to remember that the statistically most likely next token is often 
not the token that produces the most "accurate" text. Never depend upon Hippogriff to produce factually accurate output.

Llama was trained on the Pile, a dataset known to contain profanity, lewd, and otherwise abrasive language. Depending upon use case Hippogriff may produce
socially unacceptable text. See Sections 5 and 6 of the Pile paper for a more detailed analysis of the biases in the Pile.

As with all language models, it is hard to predict in advance how Hippogriff will respond to particular prompts and offensive content may occur without warning. 
We recommend having a human curate or filter the outputs before releasing them, both to censor undesirable content and to improve the quality of the results.

## Examples

TBD

```
USER: write an original short story of at least 300 words in the style of Zora Neale Hurston. Incorporate the following literary devices: Euphemism and Foreshadowing.
ASSISTANT: Once upon a time, there was a black girl named Rose, whose skin shined like a pearl beneath the hot sun. She lived with her mama and daddy on their farm, where they grew crops to feed the family and neighbors. On this particular day, Rose felt a tinge of excitement as she watched her parents work diligently in their fields. 'It's almost harvest season!' she thought, imagining the delicious meals that awaited them. But, soon after, it became apparent that something sinister lurked amidst the rows of fresh produce. 

As her father hoed around a vine-covered cucumber plant, he unearthed a peculiar rock. It was the size of an acorn but had a strange shape, unlike anything he or his wife had ever seen. "What have we here?" He said, turning to his wife with curiosity and excitement. She examined the rock for herself and declared, "Why, it looks like a piece of coal!" But as both mother and daughter looked closer, it seemed to be far more than a mere piece of coal. A sense of unease crept up Rose's spine when she noticed that the rock resembled the dark face of a baby in the womb. As Rose tried to suppress a scream, her mother gasped, "Sweet Jesus! Let's bury this thing before anyone sees!" And so, they quickly hid the stone back underground, hoping that nobody would ever discover their secret.
```

```
USER: tell me 5 words that start with "ex" and then the sum of 33 and 77
ASSISTANT: Excite, exotic, exhibit, excitable, excessive
```


```
USER: sally (a girl) has 3 brothers. Each brother has 2 sisters. how many sisters does sally have?
ASSISTANT: Sally has 3 brothers, and each brother has 2 sisters. To find out how many sisters Sally has in total, we need to multiply the number of bro
thers by the number of sisters per brother, then add 1 to account for Sally herself. So, 3 * 2 + 1 = 6 sisters in total
```
