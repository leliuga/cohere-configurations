CAMEL-33B-Combined-Data is a chat large language model obtained by finetuning LLaMA-33B model on a total of 229K conversations collected through our CAMEL framework, 100K English public conversations from ShareGPT that can be found here, and 52K instructions from Alpaca dataset that can be found here. We evaluate our model offline using EleutherAI's language model evaluation harness used by Huggingface's Open LLM Benchmark. CAMEL-33B scores an average of 64.2.

Regarding the prompt format, we follow the same prompt as LMSYS's [FastChat](https://github.com/lm-sys/FastChat/tree/main) Vicuna-13B-1.1 conversation template. It assumes a conversation between a user and AI assistant seperated by a <\/s> at the end of every role message. More details can be found [here](https://github.com/lm-sys/FastChat/blob/daa2b9abe20597ebf34dc5df164d450456610c74/fastchat/conversation.py#LL247C1-L247C1).

---
license: cc-by-nc-4.0
---
