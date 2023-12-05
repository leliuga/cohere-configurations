---
license: cc-by-sa-4.0
---

<div style="width: 800px; margin: auto;">

<h2>Model Description</h2>
<p>“Luna AI Llama2 Uncensored” is a Llama2 based Chat model <br />fine-tuned on over 40,000 long form chat discussions <br />
  This model was fine-tuned by Tap, the creator of Luna AI. <br /> 
  
<h2>Model Training</h2>
<p>The fine-tuning process was performed on an 8x a100 80GB machine.
  <br />The model was trained on synthetic outputs which include multiple rounds of chats between Human & AI.
</p>

<a rel="noopener nofollow" href="https://huggingface.co/TheBloke/Luna-AI-Llama2-Uncensored-GPTQ">4bit GPTQ Version provided by @TheBloke - for GPU inference</a><br />
<a rel="noopener nofollow" href="https://huggingface.co/TheBloke/Luna-AI-Llama2-Uncensored-GGML">GGML Version provided by @TheBloke - For CPU inference</a>


<h2>Prompt Format</h2>
<p>The model follows the Vicuna 1.1/ OpenChat format:</p>

```
USER: I have difficulties in making friends, and I really need someone to talk to. Would you be my friend?

ASSISTANT: Of course! Friends are always here for each other. What do you like to do?

```

<h2>Benchmark Results</h2>

||||||
|---:|---:|---:|---:|---:|
|Task|Version| Metric |Value |Stderr|
|arc_challenge|0|acc_norm|0.5512|0.0146|
|hellaswag|0||||
|mmlu|1|acc_norm|0.46521|0.036|
|truthfulqa_mc|1|mc2|0.4716|0.0155|
|Average|-|-|0.5114|0.0150|

</div>


