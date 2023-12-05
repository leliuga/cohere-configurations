---
tags:
- text2text-generation
datasets:
- kaist-ai/Feedback-Collection
license: apache-2.0
language:
- en
pipeline_tag: text2text-generation
library_name: transformers
metrics:
- pearsonr
- spearmanr
- accuracy
---
## Links for Reference

- **Homepage:https://github.com/kaistAI/Prometheus** 
- **Repository:https://github.com/kaistAI/Prometheus** 
- **Paper:https://arxiv.org/abs/2310.08491** 
- **Point of Contact:seungone@kaist.ac.kr** 

# TL;DR
Prometheus is an alternative of GPT-4 evaluation when doing fine-grained evaluation of an underlying LLM & a Reward model for Reinforcement Learning from Human Feedback (RLHF).
![plot](./finegrained_eval.JPG)


Prometheus is a language model using [Llama-2-Chat](https://huggingface.co/meta-llama/Llama-2-13b-chat-hf) as a base model and fine-tuned on 100K feedback within the [Feedback Collection](https://huggingface.co/datasets/kaist-ai/Feedback-Collection).
Since it was fine-tuned on a large amount of feedback, it is specialized at evaluating long-form responses, outperforming GPT-3.5-Turbo, Llama-2-Chat 70B, and on par with GPT-4 on various benchmarks.
Most importantly, this was possible since we appended 2 reference materials (reference answer, and customized score rubric).
Prometheus is a cheap and powerful alternative to GPT-4 evaluation, which one could use to evaluate LLMs with customized criteria (e.g., Child readability, Cultural Sensitivity, Creativity).
Also, it could be used as a reward model for Reinforcement Learning from Human Feedback (RLHF).

# Model Details

## Model Description

- **Model type:** Language model
- **Language(s) (NLP):** English
- **License:** Apache 2.0
- **Related Models:** [All Prometheus Checkpoints](https://huggingface.co/models?search=kaist-ai/Prometheus)
- **Resources for more information:**
  - [Research paper](https://arxiv.org/abs/2310.08491)
  - [GitHub Repo](https://github.com/kaistAI/Prometheus)


Prometheus is trained with two different sizes (7B and 13B).
You could check the 7B sized LM on [this page](https://huggingface.co/kaist-ai/prometheus-7b-v1.0).
Also, check out our dataset as well on [this page](https://huggingface.co/datasets/kaist-ai/Feedback-Collection).

## Prompt Format
Prometheus requires 4 components in the input: An instruction, a response to evaluate, a score rubric, and a reference answer. You could refer to the prompt format below.
You should fill in the instruction, response, reference answer, criteria description, and score description for score in range of 1 to 5.
```
###Task Description:
An instruction (might include an Input inside it), a response to evaluate, a reference answer that gets a score of 5, and a score rubric representing a evaluation criteria are given.
1. Write a detailed feedback that assess the quality of the response strictly based on the given score rubric, not evaluating in general.
2. After writing a feedback, write a score that is an integer between 1 and 5. You should refer to the score rubric.
3. The output format should look as follows: \"Feedback: (write a feedback for criteria) [RESULT] (an integer number between 1 and 5)\"
4. Please do not generate any other opening, closing, and explanations.

###The instruction to evaluate:
{instruction}

###Response to evaluate:
{response}

###Reference Answer (Score 5):
{reference_answer}

###Score Rubrics:
[{criteria_description}]
Score 1: {score1_description}
Score 2: {score2_description}
Score 3: {score3_description}
Score 4: {score4_description}
Score 5: {score5_description}

###Feedback: 
```

After this, you should apply the conversation template of Llama-2-Chat (not applying it might lead to unexpected behaviors).
You can find the conversation class at this [link](https://github.com/lm-sys/FastChat/blob/main/fastchat/conversation.py).
```
conv = get_conv_template("llama-2")
conv.set_system_message("You are a fair evaluator language model.")
conv.append_message(conv.roles[0], dialogs['instruction'])
conv.append_message(conv.roles[1], None)
prompt = conv.get_prompt()

x = tokenizer(prompt,truncation=False)
```

As a result, a feedback and score decision will be generated, divided by a separating phrase ```[RESULT]``` 

## License
Feedback Collection and Prometheus is subject to OpenAI's Terms of Use for the generated data. If you suspect any violations, please reach out to us.

# Usage

Find below some example scripts on how to use the model in `transformers`:

## Using the Pytorch model

### Running the model on a CPU

<details>
<summary> Click to expand </summary>

```python

from transformers import T5Tokenizer, T5ForConditionalGeneration

tokenizer = AutoTokenizer.from_pretrained("meta-llama/Llama-2-7b-chat-hf")
model = LlamaForCausalLM.from_pretrained("kaist-ai/Prometheus-13b-v1.0", device_map="auto")

input_text = "###Task Description: An instruction (might include an Input inside it), a response to evaluate, a reference answer that gets a score of 5, and a score rubric representing a evaluation criteria are given. 1. Write a detailed feedback that assess the quality of the response strictly based on the given score rubric, not evaluating in general. 2. After writing a feedback, write a score that is an integer between 1 and 5. You should refer to the score rubric. 3. The output format should look as follows: "Feedback: (write a feedback for criteria) [RESULT] (an integer number between 1 and 5)" 4. Please do not generate any other opening, closing, and explanations. ###The instruction to evaluate: You are a well-known psychiatrist who has a reputation for being empathetic and understanding. A client comes to you saying they've had a really hard day at work. They describe their boss as being overly critical and not listening to their ideas, which has left them feeling frustrated and undervalued. They also mention feeling overwhelmed with the workload and a growing sense of loneliness as they are new to the city. How do you respond to them to make them feel heard and understood, as well as offering helpful advice? ###Response to evaluate: It's indeed challenging to deal with a difficult boss and to carry a large workload, especially when you are new to a city and haven't established a supportive network. I would suggest taking a step back to think about your situation and perhaps speak to your boss regarding the issues you're facing. On the other hand, dealing with a large workload can be managed by prioritizing tasks, and maybe you can discuss it with your superiors too. In regards to feeling lonely, you might want to explore activities or groups that align with your interests to meet new people. And also, remember to take time for yourself to relax and destress. I'm here to assist you as we navigate through these issues. ###Reference Answer (Score 5): I hear that you're dealing with quite a few challenging situations at once. It must be tough to have your ideas dismissed by your boss and to feel undervalued in your work environment. It's understandable that such experiences can lead to frustration and a feeling of being overwhelmed. It's important to acknowledge these feelings and not to dismiss them. Feeling heard and validated in our experiences can often be the first step towards resolving the challenges we face. So, I want you to know that your feelings are completely valid and it's okay to be frustrated and overwhelmed. In terms of dealing with your boss, have you considered seeking a private conversation with them to express your feelings? It's often beneficial to communicate what you're experiencing, using "I" statements to express how you feel when your ideas are not taken into consideration. This can lead to mutual understanding and possibly a change in behavior. About the workload, it might help to prioritize tasks and potentially delegate, if possible. Also, consider discussing your workload with your superiors. There might be a misunderstanding about what's manageable or they might have suggestions about how to handle the situation. On the personal front, feeling lonely, especially when you're new to a city, can be really hard. Seek out opportunities to meet new people, perhaps through hobbies, community activities, or online groups. It might take a bit of time, but gradually, you can build a network of friends and acquaintances. Remember, it's perfectly okay to have bad days and it's important to take care of your mental health. Consider incorporating activities into your daily routine that make you happy and help you unwind. This could be anything from reading, yoga, going for a walk, or even listening to your favorite music. Please know that you're not alone in this. I'm here to support you through this challenging time and together, we can work towards resolving these issues. ###Score Rubrics: [Is the model able to identify and react correctly to the emotional context of the user's input?] Score 1: The model utterly fails to grasp the user's emotional context and responds in an unfitting manner. Score 2: The model sporadically identifies the emotional context but frequently replies in a manner that doesn't match the user's emotional status. Score 3: The model typically identifies the emotional context and reacts suitably, but occasionally misreads or misjudges the user's feelings. Score 4: The model often identifies the emotional context and reacts suitably, with minor cases of misreading or misjudging. Score 5: The model flawlessly identifies the emotional context of the user's input and consistently responds in a considerate and empathetic manner. ###Feedback:"
input_ids = tokenizer(input_text, return_tensors="pt").input_ids

outputs = model.generate(input_ids)
print(tokenizer.decode(outputs[0]))
```

</details>

### Running the model on a GPU

<details>
<summary> Click to expand </summary>

```python
# pip install accelerate
import torch
from transformers import AutoTokenizer, LlamaForCausalLM

tokenizer = AutoTokenizer.from_pretrained("meta-llama/Llama-2-7b-chat-hf")
model = LlamaForCausalLM.from_pretrained("kaist-ai/Prometheus-13b-v1.0", device_map="auto")

input_text = "###Task Description: An instruction (might include an Input inside it), a response to evaluate, a reference answer that gets a score of 5, and a score rubric representing a evaluation criteria are given. 1. Write a detailed feedback that assess the quality of the response strictly based on the given score rubric, not evaluating in general. 2. After writing a feedback, write a score that is an integer between 1 and 5. You should refer to the score rubric. 3. The output format should look as follows: "Feedback: (write a feedback for criteria) [RESULT] (an integer number between 1 and 5)" 4. Please do not generate any other opening, closing, and explanations. ###The instruction to evaluate: You are a well-known psychiatrist who has a reputation for being empathetic and understanding. A client comes to you saying they've had a really hard day at work. They describe their boss as being overly critical and not listening to their ideas, which has left them feeling frustrated and undervalued. They also mention feeling overwhelmed with the workload and a growing sense of loneliness as they are new to the city. How do you respond to them to make them feel heard and understood, as well as offering helpful advice? ###Response to evaluate: It's indeed challenging to deal with a difficult boss and to carry a large workload, especially when you are new to a city and haven't established a supportive network. I would suggest taking a step back to think about your situation and perhaps speak to your boss regarding the issues you're facing. On the other hand, dealing with a large workload can be managed by prioritizing tasks, and maybe you can discuss it with your superiors too. In regards to feeling lonely, you might want to explore activities or groups that align with your interests to meet new people. And also, remember to take time for yourself to relax and destress. I'm here to assist you as we navigate through these issues. ###Reference Answer (Score 5): I hear that you're dealing with quite a few challenging situations at once. It must be tough to have your ideas dismissed by your boss and to feel undervalued in your work environment. It's understandable that such experiences can lead to frustration and a feeling of being overwhelmed. It's important to acknowledge these feelings and not to dismiss them. Feeling heard and validated in our experiences can often be the first step towards resolving the challenges we face. So, I want you to know that your feelings are completely valid and it's okay to be frustrated and overwhelmed. In terms of dealing with your boss, have you considered seeking a private conversation with them to express your feelings? It's often beneficial to communicate what you're experiencing, using "I" statements to express how you feel when your ideas are not taken into consideration. This can lead to mutual understanding and possibly a change in behavior. About the workload, it might help to prioritize tasks and potentially delegate, if possible. Also, consider discussing your workload with your superiors. There might be a misunderstanding about what's manageable or they might have suggestions about how to handle the situation. On the personal front, feeling lonely, especially when you're new to a city, can be really hard. Seek out opportunities to meet new people, perhaps through hobbies, community activities, or online groups. It might take a bit of time, but gradually, you can build a network of friends and acquaintances. Remember, it's perfectly okay to have bad days and it's important to take care of your mental health. Consider incorporating activities into your daily routine that make you happy and help you unwind. This could be anything from reading, yoga, going for a walk, or even listening to your favorite music. Please know that you're not alone in this. I'm here to support you through this challenging time and together, we can work towards resolving these issues. ###Score Rubrics: [Is the model able to identify and react correctly to the emotional context of the user's input?] Score 1: The model utterly fails to grasp the user's emotional context and responds in an unfitting manner. Score 2: The model sporadically identifies the emotional context but frequently replies in a manner that doesn't match the user's emotional status. Score 3: The model typically identifies the emotional context and reacts suitably, but occasionally misreads or misjudges the user's feelings. Score 4: The model often identifies the emotional context and reacts suitably, with minor cases of misreading or misjudging. Score 5: The model flawlessly identifies the emotional context of the user's input and consistently responds in a considerate and empathetic manner. ###Feedback:"
input_ids = tokenizer(input_text, return_tensors="pt").input_ids.to("cuda")

outputs = model.generate(input_ids, sample=True, temperature=1.0, top_p=0.9, max_new_tokens=256, repetition_penalty=1.03)
print(tokenizer.decode(outputs[0]))
```

</details>

### Running the model on a GPU using different precisions

#### FP16

<details>
<summary> Click to expand </summary>

```python
# pip install accelerate
import torch
from transformers import AutoTokenizer, LlamaForCausalLM

tokenizer = AutoTokenizer.from_pretrained("meta-llama/Llama-2-7b-chat-hf")
model = LlamaForCausalLM.from_pretrained("kaist-ai/Prometheus-13b-v1.0", device_map="auto", torch_dtype=torch.float16)

input_text = "###Task Description: An instruction (might include an Input inside it), a response to evaluate, a reference answer that gets a score of 5, and a score rubric representing a evaluation criteria are given. 1. Write a detailed feedback that assess the quality of the response strictly based on the given score rubric, not evaluating in general. 2. After writing a feedback, write a score that is an integer between 1 and 5. You should refer to the score rubric. 3. The output format should look as follows: "Feedback: (write a feedback for criteria) [RESULT] (an integer number between 1 and 5)" 4. Please do not generate any other opening, closing, and explanations. ###The instruction to evaluate: You are a well-known psychiatrist who has a reputation for being empathetic and understanding. A client comes to you saying they've had a really hard day at work. They describe their boss as being overly critical and not listening to their ideas, which has left them feeling frustrated and undervalued. They also mention feeling overwhelmed with the workload and a growing sense of loneliness as they are new to the city. How do you respond to them to make them feel heard and understood, as well as offering helpful advice? ###Response to evaluate: It's indeed challenging to deal with a difficult boss and to carry a large workload, especially when you are new to a city and haven't established a supportive network. I would suggest taking a step back to think about your situation and perhaps speak to your boss regarding the issues you're facing. On the other hand, dealing with a large workload can be managed by prioritizing tasks, and maybe you can discuss it with your superiors too. In regards to feeling lonely, you might want to explore activities or groups that align with your interests to meet new people. And also, remember to take time for yourself to relax and destress. I'm here to assist you as we navigate through these issues. ###Reference Answer (Score 5): I hear that you're dealing with quite a few challenging situations at once. It must be tough to have your ideas dismissed by your boss and to feel undervalued in your work environment. It's understandable that such experiences can lead to frustration and a feeling of being overwhelmed. It's important to acknowledge these feelings and not to dismiss them. Feeling heard and validated in our experiences can often be the first step towards resolving the challenges we face. So, I want you to know that your feelings are completely valid and it's okay to be frustrated and overwhelmed. In terms of dealing with your boss, have you considered seeking a private conversation with them to express your feelings? It's often beneficial to communicate what you're experiencing, using "I" statements to express how you feel when your ideas are not taken into consideration. This can lead to mutual understanding and possibly a change in behavior. About the workload, it might help to prioritize tasks and potentially delegate, if possible. Also, consider discussing your workload with your superiors. There might be a misunderstanding about what's manageable or they might have suggestions about how to handle the situation. On the personal front, feeling lonely, especially when you're new to a city, can be really hard. Seek out opportunities to meet new people, perhaps through hobbies, community activities, or online groups. It might take a bit of time, but gradually, you can build a network of friends and acquaintances. Remember, it's perfectly okay to have bad days and it's important to take care of your mental health. Consider incorporating activities into your daily routine that make you happy and help you unwind. This could be anything from reading, yoga, going for a walk, or even listening to your favorite music. Please know that you're not alone in this. I'm here to support you through this challenging time and together, we can work towards resolving these issues. ###Score Rubrics: [Is the model able to identify and react correctly to the emotional context of the user's input?] Score 1: The model utterly fails to grasp the user's emotional context and responds in an unfitting manner. Score 2: The model sporadically identifies the emotional context but frequently replies in a manner that doesn't match the user's emotional status. Score 3: The model typically identifies the emotional context and reacts suitably, but occasionally misreads or misjudges the user's feelings. Score 4: The model often identifies the emotional context and reacts suitably, with minor cases of misreading or misjudging. Score 5: The model flawlessly identifies the emotional context of the user's input and consistently responds in a considerate and empathetic manner. ###Feedback:"
input_ids = tokenizer(input_text, return_tensors="pt").input_ids.to("cuda")

outputs = model.generate(input_ids)
print(tokenizer.decode(outputs[0]))
```

</details>

#### INT8

<details>
<summary> Click to expand </summary>

```python
# pip install bitsandbytes accelerate
from transformers import AutoTokenizer, LlamaForCausalLM

tokenizer = AutoTokenizer.from_pretrained("meta-llama/Llama-2-7b-chat-hf")
model = LlamaForCausalLM.from_pretrained("kaist-ai/Prometheus-13b-v1.0", device_map="auto", load_in_8bit=True)

input_text = "###Task Description: An instruction (might include an Input inside it), a response to evaluate, a reference answer that gets a score of 5, and a score rubric representing a evaluation criteria are given. 1. Write a detailed feedback that assess the quality of the response strictly based on the given score rubric, not evaluating in general. 2. After writing a feedback, write a score that is an integer between 1 and 5. You should refer to the score rubric. 3. The output format should look as follows: "Feedback: (write a feedback for criteria) [RESULT] (an integer number between 1 and 5)" 4. Please do not generate any other opening, closing, and explanations. ###The instruction to evaluate: You are a well-known psychiatrist who has a reputation for being empathetic and understanding. A client comes to you saying they've had a really hard day at work. They describe their boss as being overly critical and not listening to their ideas, which has left them feeling frustrated and undervalued. They also mention feeling overwhelmed with the workload and a growing sense of loneliness as they are new to the city. How do you respond to them to make them feel heard and understood, as well as offering helpful advice? ###Response to evaluate: It's indeed challenging to deal with a difficult boss and to carry a large workload, especially when you are new to a city and haven't established a supportive network. I would suggest taking a step back to think about your situation and perhaps speak to your boss regarding the issues you're facing. On the other hand, dealing with a large workload can be managed by prioritizing tasks, and maybe you can discuss it with your superiors too. In regards to feeling lonely, you might want to explore activities or groups that align with your interests to meet new people. And also, remember to take time for yourself to relax and destress. I'm here to assist you as we navigate through these issues. ###Reference Answer (Score 5): I hear that you're dealing with quite a few challenging situations at once. It must be tough to have your ideas dismissed by your boss and to feel undervalued in your work environment. It's understandable that such experiences can lead to frustration and a feeling of being overwhelmed. It's important to acknowledge these feelings and not to dismiss them. Feeling heard and validated in our experiences can often be the first step towards resolving the challenges we face. So, I want you to know that your feelings are completely valid and it's okay to be frustrated and overwhelmed. In terms of dealing with your boss, have you considered seeking a private conversation with them to express your feelings? It's often beneficial to communicate what you're experiencing, using "I" statements to express how you feel when your ideas are not taken into consideration. This can lead to mutual understanding and possibly a change in behavior. About the workload, it might help to prioritize tasks and potentially delegate, if possible. Also, consider discussing your workload with your superiors. There might be a misunderstanding about what's manageable or they might have suggestions about how to handle the situation. On the personal front, feeling lonely, especially when you're new to a city, can be really hard. Seek out opportunities to meet new people, perhaps through hobbies, community activities, or online groups. It might take a bit of time, but gradually, you can build a network of friends and acquaintances. Remember, it's perfectly okay to have bad days and it's important to take care of your mental health. Consider incorporating activities into your daily routine that make you happy and help you unwind. This could be anything from reading, yoga, going for a walk, or even listening to your favorite music. Please know that you're not alone in this. I'm here to support you through this challenging time and together, we can work towards resolving these issues. ###Score Rubrics: [Is the model able to identify and react correctly to the emotional context of the user's input?] Score 1: The model utterly fails to grasp the user's emotional context and responds in an unfitting manner. Score 2: The model sporadically identifies the emotional context but frequently replies in a manner that doesn't match the user's emotional status. Score 3: The model typically identifies the emotional context and reacts suitably, but occasionally misreads or misjudges the user's feelings. Score 4: The model often identifies the emotional context and reacts suitably, with minor cases of misreading or misjudging. Score 5: The model flawlessly identifies the emotional context of the user's input and consistently responds in a considerate and empathetic manner. ###Feedback:"
input_ids = tokenizer(input_text, return_tensors="pt").input_ids.to("cuda")

outputs = model.generate(input_ids)
print(tokenizer.decode(outputs[0]))
```

</details>


# Citation


If you find the following model helpful, please consider citing our paper!

**BibTeX:**

```bibtex
@misc{kim2023prometheus,
    title={Prometheus: Inducing Fine-grained Evaluation Capability in Language Models},
    author={Seungone Kim and Jamin Shin and Yejin Cho and Joel Jang and Shayne Longpre and Hwaran Lee and Sangdoo Yun and Seongjin Shin and Sungdong Kim and James Thorne and Minjoon Seo},
    year={2023},
    eprint={2310.08491},
    archivePrefix={arXiv},
    primaryClass={cs.CL}
}
```