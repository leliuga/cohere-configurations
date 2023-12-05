Merge of [huggyllama/llama-30b](https://huggingface.co/huggyllama/llama-30b) + [kaiokendev/SuperCOT-LoRA](https://huggingface.co/kaiokendev/SuperCOT-LoRA/edit/main/README.md) 

Supercot was trained to work with langchain prompting. 


Load up locally in my custom LLM notebook that uses the Oobabooga modules to load up models: https://github.com/ausboss/Local-LLM-Langchain

Then you can add cells from of these other notebooks for testing: https://github.com/gkamradt/langchain-tutorials



# From Koikendev Lora page

### Compatibility
This LoRA is compatible with any 7B, 13B or 30B 4-bit quantized LLaMa model, including ggml quantized converted bins

### Prompting
You should prompt the LoRA the same way you would prompt Alpaca or Alpacino:

```
Below is an instruction that describes a task, paired with an input that provides further context. Write a response that appropriately completes the request.

### Instruction:
<instruction>

### Input:
<any additional context. Remove this if it's not neccesary>

### Response:
<make sure to leave a single new-line here for optimal results>
```

Remember that with lower parameter sizes, the structure of the prompt becomes more important. The same prompt worded differently can give wildly different answers. Consider using the following suggestion suffixes to improve output quality:

- "Think through this step by step"
- "Let's think about this logically"
- "Explain your reasoning"
- "Provide details to support your answer"
- "Compare and contrast your answer with alternatives"

### Coming Soon
- Tweet fix for 13B and 7B - lower model sizes seem to be extremely sensitive to hashtags at the end of training data responses, especially at longer cutoffs