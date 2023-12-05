---
language:
- en
- pl
pipeline_tag: text-generation
inference: false
tags:
- voicelab
- pytorch
- llama-2
- trurl
- trurl-2
---
<img src="https://public.3.basecamp.com/p/rs5XqmAuF1iEuW6U7nMHcZeY/upload/download/VL-NLP-short.png" alt="logo voicelab nlp" style="width:300px;"/>


# Trurl 2 -- Polish Llama 2

The new OPEN TRURL is a finetuned Llama 2, trained on over 1.7b tokens (970k conversational **Polish** and **English** samples) with a large context of 4096 tokens.
TRURL was trained on a large number of Polish data.
TRURL 2 is a collection of fine-tuned generative text models with 7 billion and 13 billion parameters. 
This is the repository for the 7b fine-tuned model, optimized for dialogue use cases.


# Overview

**TRURL developers** Voicelab.AI

**Variations** Trurl 2 comes in 7B and 13B versions.

**Input** Models input text only.

**Output** Models generate text only.

**Model Architecture** Trurl is an auto-regressive language model that uses an optimized transformer architecture.

||Training Data|Params|Content Length|Num. Samples|Num. Tokens|start LR|
|---|---|---|---|---|---|---|
|Trurl 2|*A new mix of private and publicly available online data without MMLU*|7B|4k|855k|1.19b|2.0 x 10<sup>-5</sup>|
|Trurl 2|*A new mix of private and publicly available online data with MMLU*|13B|4k|970k|1.7b|2.0 x 10<sup>-5</sup>|
|Trurl 2 Academic|*A new mix of private and publicly available online data without MMLU*|13B|4k|855k|1.19b|2.0 x 10<sup>-5</sup>|

## Training data

The training data includes Q&A pairs from various sources including Alpaca comparison data with GPT, Falcon comparison data, Dolly 15k, Oasst1, Phu saferlfhf, ShareGPT version 2023.05.08v0 filtered and cleaned, Voicelab private datasets for JSON data extraction, modification, and analysis, CURLICAT dataset containing journal entries, dataset from Polish wiki with Q&A pairs grouped into conversations, Voicelab private dataset with sales conversations, arguments and objections, paraphrases, contact reason detection, and corrected dialogues.

## Intended Use

Trurl 2 is intended for commercial and research use in Polish and English. Tuned models are intended for assistant-like chat, but also adapted for a variety of natural language generation tasks.

# Evaluation Results
|Model | Size| hellaswag | arc_challenge | MMLU|
|---|---|---|---|---|
| Llama-2-chat | 7B |  78.55% |  52.9% | 48.32% |
| Llama-2-chat | 13B |  81.94% |  59.04% | 54.64% |
| Trurl 2.0 (with MMLU) | 13B | 80.09% | 59.30% | 78.35% |
| Trurl 2.0 (no MMLU) | 13B | TO-DO | TO-DO | TO-DO|
| Trurl 2.0 (no MMLU) | 7b | 75.29% | 53.41%| 50.0%|
	

<img src="https://voicelab.ai/wp-content/uploads/trurl-hero.webp" alt="trurl graphic" style="width:100px;"/>

# Ethical Considerations and Limitations
Trurl 2, same as a Llama 2, is a new technology that carries risks with use. Testing conducted to date has been in Polish and English, and has not covered, nor could it cover all scenarios. For these reasons, as with all LLMs, Trurl 2’s potential outputs cannot be predicted in advance, and the model may in some instances produce inaccurate, biased or other objectionable responses to user prompts. Therefore, before deploying any applications of Trurl 2, developers should perform safety testing and tuning tailored to their specific applications of the model.

Please see the Meta's Responsible Use Guide available at [https://ai.meta.com/llama/responsible-use-guide/](https://ai.meta.com/llama/responsible-use-guide)

# Example use
## LLM
Simply pass a prompt to a model and decode an output. Model will continue writing text based on sample you provided.
```
import torch
from transformers import LlamaForCausalLM, LlamaTokenizer

tokenizer = LlamaTokenizer.from_pretrained("Voicelab/trurl-2-7b")
model = LlamaForCausalLM.from_pretrained("Voicelab/trurl-2-7b")

prompt = "Yesterday, when I was"

tokenized_prompt = tokenizer(prompt, return_tensors="pt")

model.eval()
with torch.no_grad():
    print(tokenizer.decode(
        model.generate(**tokenized_prompt, max_new_tokens=200)[0],
        skip_special_tokens=True))
```
Generated output:
> Yesterday, when I was in the city, I saw a man who was walking his dog. and the dog was wearing a little sweater. I thought it was so cute! I wish I had a dog so I could get one of those sweaters for my own dog.

## Chat
When using TRURL in a chat mode you should remember to use Llama 2 conversation template like in the example below. 


```
import torch
from transformers import LlamaForCausalLM, LlamaTokenizer

tokenizer = LlamaTokenizer.from_pretrained("Voicelab/trurl-2-7b")
model = LlamaForCausalLM.from_pretrained("Voicelab/trurl-2-7b")

prompt = """
<s>[INST] <<SYS>>  You are a helpful, respectful and honest assistant. Always answer as helpfully as possible, while being safe.
Your answers should not include any harmful, unethical, racist, sexist, toxic, dangerous, or illegal content.
Please ensure that your responses are socially unbiased and positive in nature.\n\n
If a question does not make any sense, or is not factually coherent, explain why instead of answering something not correct.
If you don't know the answer to a question, please don't share false information. <</SYS>>

What was the reason for calling in the conversation below? \n\n
AGENT: Hello, Bank of Albion, this is Mata Hari. How can I help you?
CLIENT: Hi. I've been locked out from my Internet account. I need your help.
AGENT: (yy) Yes, of course, I'll do my best to help you. But I need to find out why the locking-out happened. (yy) In order to ascertain that, I'll ask you a couple of questions to confirm your identity. I'm going to need your full name.
CLIENT: Lizz Truss.
AGENT: Thank you. Now I need your personal identification number.
CLIENT: Fourteen, two hundred thirty-one, thirty-eight, twenty-nine, sixty-five.
AGENT: Thank you. Now I need your client ID number. The client ID number is the eight digits we assigned to you at the very beginning, on conclusion of the contract.
CLIENT: OK. Give me a moment. I have to find it.
AGENT: (mhm) You'll find… You'll find it in the contract.
CLIENT: Yes, yes. I can see it. Sixty-five, twenty-nine, thirty-eight, thirty-one.
AGENT: Thank you. One final security question. Do you have any deposits in our bank?
CLIENT: No, no. I don't have any deposits in this bank.
AGENT: Thank you. Your identity has been (yy) confirmed. (yy) I can see that the account has been blocked, indeed, and you won't be able to log in via the Internet (yy) because (yy) the identity document which is listed for reference has expired. (yy) From what I can see, your identity document expired some time ago. Have you been issued a new one?
CLIENT: Well, no. I think my ID is still valid, you know. I didn't even know.
AGENT: Well, no... Your ID expired at the end of March. Well, almost at the end. Your old ID had been valid until 26 March. (yy) For that reason, your accout has been blocked, because you haven't notified us about the ID change for a few months. We are not interested if the ID document has been officialy reissued. (...) On our end, what matters is whether the document listed for our reference is valid (yy) so without a valid document I can't unlock your accout. 
CLIENT: But I have to carry out an operation right now, so this is sort of problematic.
AGENT: I understand. But (yy) you are obligated, as an account holder, to notify the bank about any changes pending (yy), regrding, for example, your home address or phone number. Now, one of such safeguards protecting your… (yy) money, your sensitive data, is precisely about having a valid identification document. Since this is missing in your case, the account has been blocked. Now, I don't think this would have caught you off guard, because we always remind our customers that their ID is about to expire. When the ID is nearing expiration, we display relevant messages at least sixty days in advance. They appear once you've logged in, at the very top of the screen, there is a notification that (yy) the ID is about to expire (yy), so, well... The bank did notify you about this issue. Now, how you chose to act on this information was your choice, right? In any case, at this point, in order to unlock your accout, our protocols require that you produce a new identification document at one of our branches. You shall provide information concerning the new document number, new valid-thru date, and only then will you be able to use your account again. I can schedule an appointment with a consultant at our branch for you. What locality would you prefer?
CLIENT: Well, I'm not sure if I should share such information with you.
AGENT: And may I ask why exactly you are unsure? After all, you're calling a bank that runs your account, right?
CLIENT: Right, you know what, I need to go now. Good bye.
AGENT: (yy) Miss… [/INST]

"""

tokenized_prompt = tokenizer(prompt, return_tensors="pt")

model.eval()
with torch.no_grad():
    print(tokenizer.decode(
        model.generate(**tokenized_prompt, max_new_tokens=200)[0],
        skip_special_tokens=True))
```

Generated output:
> The reason for calling in this conversation is for the agent to help the client regain access to their internet account, which has been locked due to an expired identification document. The agent asks for the client's personal information to confirm their identity and then informs them that their account has been blocked because they have not notified the bank about the ID change for a few months. The agent explains that the bank has displayed relevant messages about the ID expiring and that the client must produce a new identification document at one of their branches in order to unlock their account. The client expresses uncertainty about sharing their information with the agent, but ultimately decides to end the call.

To get the expected features and performance for the chat versions, a specific Llama 2 formatting needs to be followed, including the `INST` and `<<SYS>>` tags, `BOS` and `EOS` tokens, and the whitespaces and breaklines in between (we recommend calling `strip()` on inputs to avoid double-spaces). See reference code in github for details: [`chat_completion`](https://github.com/facebookresearch/llama/blob/main/llama/generation.py#L212).

# Authors

The model was trained by NLP Research Team at Voicelab.ai.

You can contact us [here](https://voicelab.ai/contact/).

* [TRURL 13b](https://huggingface.co/Voicelab/trurl-2-13b/)
* [TRURL 13b Academic](https://huggingface.co/Voicelab/trurl-2-13b-academic)
* [TRURL 7b](https://huggingface.co/Voicelab/trurl-2-7b/)
* [TRURL DEMO](https://trurl.ai)

Quantized models:
* [TRURL 13b - 8bit](https://huggingface.co/Voicelab/trurl-2-13b-8bit/)
* [TRURL 7b - 8bit](https://huggingface.co/Voicelab/trurl-2-7b-8bit/)

The work was supported by [#NASK](https://www.nask.pl/)


# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Voicelab__trurl-2-7b)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 48.05   |
| ARC (25-shot)         | 53.41          |
| HellaSwag (10-shot)   | 75.29    |
| MMLU (5-shot)         | 50.0         |
| TruthfulQA (0-shot)   | 45.42   |
| Winogrande (5-shot)   | 72.22   |
| GSM8K (5-shot)        | 7.13        |
| DROP (3-shot)         | 32.9         |
