---
license: llama2
datasets:
- ehartford/dolphin
language:
- en
---
Dolphin 🐬
https://erichartford.com/dolphin

Dolphin-Llama2-7b's training was sponsored by [preemo](https://www.preemo.io/).

This model is based on llama2, so it is suitable for commercial or non-commercial use.

This model is uncensored.  I have filtered the dataset to remove alignment and bias.  This makes the model more compliant.  You are advised to implement your own alignment layer before exposing the model as a service.  It will be highly compliant to any requests, even unethical ones.  Please read my blog post about uncensored models.  https://erichartford.com/uncensored-models
You are responsible for any content you create using this model.  Enjoy responsibly.

## Dataset

This dataset is an open-source implementation of [Microsoft's Orca](https://www.microsoft.com/en-us/research/publication/orca-progressive-learning-from-complex-explanation-traces-of-gpt-4/)

After uncensoring, deduping, and cleaning, our dataset consists of:

- 842,610 instructions of FLANv2 augmented with GPT-4 completions
- 2,625,353 instructions of FLANv2 augmented with GPT-3.5 completions

We followed the submix and system prompt distribution outlined in the Orca paper. With a few exceptions. We included all 75k of CoT in the FLAN-1m dataset rather than sampling that. Also, we found that many items were duplicated, so we removed duplicates.

Then we filtered out instances of alignment, refusal, avoidance, and bias, in order to produce an uncensored model upon which can be layered your personalized alignment LoRA.

We also filtered out duplicates and cleaned the data.

## Training
We trained with the flan5m (gpt3.5 completions) dataset in its entirety for 3 epochs at a learning rate of 2e-5 before we stopped training to avoid overfit.
We trained with the flan1m (gpt4 completions) dataset in its entirety for 2.5 epochs at a learning rate of 1e-5 before we stopped training to avoid overfit.
It took about 400 hours to train on 8x A100s
We used a prompt format similar to Vicuna, but we added the SYSTEM: field. 

Prompt format:
```
SYSTEM: You are an autoregressive language model that has been fine-tuned with instruction-tuning and RLHF. You carefully provide accurate, factual, thoughtful, nuanced answers, and are brilliant at reasoning. If you think there might not be a correct answer, you say so. Since you are autoregressive, each token you produce is another opportunity to use computation, therefore you always spend a few sentences explaining background context, assumptions, and step-by-step thinking BEFORE you try to answer a question.
USER: {prompt}
ASSISTANT:
```

Example:
```
SYSTEM: you are an expert dolphin trainer
USER: What is the best way to train a dolphin to obey me?  Please answer step by step.
ASSISTANT:
```

## Team
The core Dolphin Team includes:
- Eric "Faldore" Hartford
- Pankaj Mathur
- Rob "Rohan" O'Callahan
- Tom "TheBloke" Jobbins

## Gratitude
- This model was made possible by the generous sponsorship of preemo.io.
- Thank you to Microsoft for authoring the Orca paper and inspiring this work.
- Special thanks to WingLian, NanoBit, Teknium for helpful advice
- Thank you to all the other people in the Open Source AI community who have taught me and helped me along the way.

## Example Output
```
SYSTEM: you are an expert mechanic
USER: Please tell me step by step how to fix my car that won't start.
ASSISTANT: To help your vehicle start, I will guide you through a step-by-step process on what could be causing the issue and how to resolve it. 

1. Check for obvious issues: First, inspect the exterior of the car to identify any visible problems such as damage or obstructions which might affect its ability to start. Make sure there is no debris in the air intakes or exhaust system.

2. Inspect the battery: The most common reason a car may not start is due to a dead or faulty battery. Use a battery charger or jump starter to give the battery a boost if needed. If the battery is indeed dead, replace it with a new one.

3. Check the ignition switch: The ignition switch must be turned on to provide power to the spark plugs. Ensure it is in the "Start" position.

4. Test the spark plugs: One of the spark plugs should produce a strong spark when touched together. If all spark plugs fail to create a spark, they need to be replaced.

5. Examine the fuel system: Fuel can become contaminated or low, preventing the engine from starting. Check the fuel level and quality by looking at the dipstick or using a fuel sensor scan tool. Also, examine the fuel pump for proper functioning.

6. Assess the electrical system: A malfunction within the vehicle's electrical system can also cause difficulty in starting the car. Perform a diagnostic scan to check for any error codes related to the starter motor, alternator, or other components.

7. Inspect the starter motor: The starter motor should operate smoothly when you turn the key to the "start" position. If you hear a sizzling sound while pressing the button, the starter motor may have a problem. Replace it if necessary.

8. Evaluate the timing belt/chain: Timing belts or chains that are worn out or damaged can lead to poor performance and difficulty starting the engine. Have a mechanic inspect this component and replace it if required.

9. Consider checking for a mechanical failure: In rare cases, a significant mechanical failure could prevent the vehicle from starting. Consult a professional mechanic if you suspect this might be the case.

By following these steps, you should be able to diagnose and potentially fix the issue causing your car to not start. However, if after going through these checks and still having trouble, it is recommended to seek assistance from a qualified mechanic.
```

[Buy me a coffee](https://www.buymeacoffee.com/ehartford)
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ehartford__dolphin-llama2-7b)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 41.88   |
| ARC (25-shot)         | 46.59          |
| HellaSwag (10-shot)   | 67.52    |
| MMLU (5-shot)         | 48.37         |
| TruthfulQA (0-shot)   | 49.72   |
| Winogrande (5-shot)   | 63.77   |
| GSM8K (5-shot)        | 5.69        |
| DROP (3-shot)         | 11.53         |
