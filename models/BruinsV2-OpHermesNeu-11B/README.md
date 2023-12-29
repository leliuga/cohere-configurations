---
license: mit
---

|    Task     |Version| Metric |Value |   |Stderr|
|-------------|------:|--------|-----:|---|-----:|
|arc_challenge|      0|acc     |0.6527|±  |0.0139|
|             |       |acc_norm|0.6869|±  |0.0136|

**Warning! This model may or may not be contaminated [See discussion 474](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard/discussions/474). What a shame. It still does perform well though**

A passthrough merge of OpenHermes-2.5-neural-chat-7b-v3-1 and Bruins-V2. To be updated.

Template: ChatML

My settings:

Temperature: 0.7-0.8

Min_p: 0.12

Top_K: 0

Repetition Penalty: 1.16

Mirostat Tau: 2.5-3

Mirostat Eta: 0.12

Personal Thoughts:

- The model sometimes throws wrong tags, you can add those to "Custom stopping strings" in Oobabooga.
- Output with Mirostat consistently felt smarter than a set Top_K rate.

Note: The model is hallucinating hard in chat mode for me in some instances, like writing adblocker messages. Kind of funny. 

I am not sure which dataset involved was poisoned.