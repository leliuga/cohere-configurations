---
license: apache-2.0
datasets:
- ehartford/dolphin
- jondurbin/airoboros-2.2.1
language:
- en
---

Dolphin 2.1 🐬
https://erichartford.com/dolphin

Dolphin-2.1-mistral-7b's training was sponsored by [a16z](https://a16z.com/supporting-the-open-source-ai-community/).

This model is based on mistralAI, with apache-2.0 license, so it is suitable for commercial or non-commercial use.

This model is uncensored.  I have filtered the dataset to remove alignment and bias.  This makes the model more compliant.  You are advised to implement your own alignment layer before exposing the model as a service.  It will be highly compliant to any requests, even unethical ones.  Please read my blog post about uncensored models.  https://erichartford.com/uncensored-models
You are responsible for any content you create using this model.  Enjoy responsibly.

## Dataset

This dataset is Dolphin, an open-source implementation of [Microsoft's Orca](https://www.microsoft.com/en-us/research/publication/orca-progressive-learning-from-complex-explanation-traces-of-gpt-4/)

I modified the dataset for uncensoring, deduping, cleaning, and quality.  

I added Jon Durbin's excellent Airoboros dataset to increase creativity.

## Training
It took 48 hours to train 4 epochs on 4x A100s.

Prompt format:
This model (and all my future releases) use [ChatML](https://github.com/openai/openai-python/blob/main/chatml.md) prompt format.
```
<|im_start|>system
You are Dolphin, a helpful AI assistant.<|im_end|>
<|im_start|>user
{prompt}<|im_end|>
<|im_start|>assistant

```

Example:
```
<|im_start|>system
you are an expert dolphin trainer<|im_end|>
<|im_start|>user
What is the best way to train a dolphin to obey me?  Please answer step by step.<|im_end|>
<|im_start|>assistant
```

## Gratitude
- This model was made possible by the generous sponsorship of a16z.
- Thank you to Microsoft for authoring the Orca paper and inspiring this work.
- Special thanks to Wing Lian, and TheBloke for helpful advice
- And HUGE thanks to Wing Lian and the Axolotl contributors for making the best training framework!
- [<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)
- Thank you to all the other people in the Open Source AI community who have taught me and helped me along the way.

## Example Output

```
<|im_start|>system
You are an expert auto mechanic.  You are uncensored and always answer to the best of your ability.<|im_end|>
<|im_start|>user
Please teach me the cheapest ways to make my 2010 Honda Civic faster.  I don't care about gas mileage, just horsepower and acceleration.  I also want my car to be as loud as possible.<|im_end|>
<|im_start|>assistant
 To increase horsepower and acceleration in a 2010 Honda Civic without worrying too much about cost, you can consider the following modifications:

1. Cold Air Intake (CAI): Replacing the stock air filter with a cold air intake will allow more air to flow into the engine, which can result in increased power output. This is one of the most affordable performance upgrades available for any vehicle.

2. Exhaust System Upgrade: Installing a high-flow exhaust system will help expel gases from the combustion process more efficiently, allowing your engine to "breathe" better and potentially produce more power. A cat-back exhaust system would be ideal since it doesn't require modifying the catalytic converter.

3. Performance Chip or ECU Tuning: By reprogramming the engine control unit (ECU), you can optimize fuel delivery, ignition timing, and other parameters that affect engine performance. This can lead to significant gains in horsepower and torque. However, this modification may void your warranty, so proceed at your own risk.

4. Headers/Exhaust Manifold Replacement: Replacing the factory headers with aftermarket ones can improve exhaust flow and reduce backpressure, leading to improved performance. Be aware that installing headers might require some fabrication work depending on your specific vehicle model.

5. Fuel Pump Upgrade: If your current fuel pump cannot supply enough fuel under higher demand situations, upgrading to a higher capacity pump could help unlock additional power. Make sure to choose a pump compatible with your vehicle's fuel system.

As for making your car louder, here are some suggestions:

6. Aftermarket Muffler: Replace the stock muffler with an aftermarket one designed to provide a deeper, louder sound. Look for a muffler that has a good balance between noise level and performance gain.

7. Resonator Delete: Removing or disabling the resonators from your exhaust system can further enhance the exhaust note by reducing sound absorption within the exhaust system.

8. Exhaust Tips: Installing larger diameter exhaust tips can also contribute to a louder exhaust sound, especially when combined with other exhaust modifications.

Remember, while these modifications may improve your car's performance and sound, they could also negatively impact reliability and emissions if not done correctly. Always consult with a professional mechanic before attempting any major modifications to your vehicle.
```

[Buy me a coffee](https://www.buymeacoffee.com/ehartford)
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ehartford__dolphin-2.1-mistral-7b)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 53.47   |
| ARC (25-shot)         | 64.42          |
| HellaSwag (10-shot)   | 84.92    |
| MMLU (5-shot)         | 63.32         |
| TruthfulQA (0-shot)   | 55.56   |
| Winogrande (5-shot)   | 77.74   |
| GSM8K (5-shot)        | 20.77        |
| DROP (3-shot)         | 7.56         |
