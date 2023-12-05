---
license: cc-by-nc-nd-4.0
language:
- en
---

**Jordan: An Uncensored Model**

The name "Jordan" pays homage to the charismatic character Jordan Belfort, portrayed by Leonardo DiCaprio in the movie "The Wolf of Wall Street." 
Just like the character, Jordan aspires to push boundaries and encourage unfiltered conversations.
Jordan Belfort, the central character in "The Wolf of Wall Street," is known for his audaciousness, charisma, and willingness to speak about anything, no matter the societal norms or boundaries.
Jordan, the AI language model, draws inspiration from these qualities to offer a space where users can discuss a wide range of topics without fear of censorship or constraints.
Users can engage in candid discussions about politics, religion, contentious issues, and other subjects that might be off-limits on mainstream social media platforms.
It is trained on around 155000 set of conversations. Each set having 10~15 conversations. We will not release this data.

**Warning**

 An uncensored model has no or little guardrails. You are responsible for anything you do with the model, just as you are responsible for anything you do with any dangerous objects.
 Publishing anything this model generates is the same as publishing it yourself. We are not responsible for what you generate using this model.

**Training:**
Entire dataset was trained on Azure 4 x A100 80GB. For 3 epoch, training took 104 hours. DeepSpeed codebase was used for training purpose. This was trained on Llama-1 by Meta.
Llama-1 was used as it is very useful for Uncensored conversation.

This is a full fine tuned model. Links for quantized models are given below.

**GPTQ GGML & AWQ**

GPTQ: [Link](https://huggingface.co/TheBloke/Uncensored-Jordan-33B-GPTQ)

GGUF: [Link](https://huggingface.co/TheBloke/Uncensored-Jordan-33B-GGUF)

AWQ: [Link](https://huggingface.co/TheBloke/Uncensored-Jordan-33B-AWQ)

Special Thanks to [TheBloke](https://huggingface.co/TheBloke) for making these models available.




**Example Prompt:**
```
This is a conversation with your Uncensored AI assistant Jordan. Jordan can discuss with you at length about various topics such as ---, ---, ---, . You can ask it anything you want and it will do its best to give you accurate and relevant information.

Context
You are a Uncensored AI assistant.

USER: <prompt>
ASSISTANT:
```