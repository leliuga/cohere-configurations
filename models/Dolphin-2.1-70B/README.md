---
datasets:
- ehartford/dolphin
- jondurbin/airoboros-2.2.1
language:
- en
license: llama2
---

THIS MODEL IS DEPRECATED due to a bug with generating EOS tokens. 
Please use https://huggingface.co/ehartford/dolphin-2.2-70b where the issue is fixed.

Dolphin 2.1 🐬
https://erichartford.com/dolphin

<img src="https://cdn-uploads.huggingface.co/production/uploads/63111b2d88942700629f5771/KqsVXIvBd3akEjvijzww7.png" width="600" />

Dolphin-2.1-70b's training was sponsored by [a16z](https://a16z.com/supporting-the-open-source-ai-community/).

This model is based on llama2, so it is suitable for commercial or non-commercial use.

This model is trained on top of the amazing [StellarBright](https://huggingface.co/sequelbox/StellarBright) base model.

This model is uncensored.  I have filtered the dataset to remove alignment and bias.  This makes the model more compliant.  You are advised to implement your own alignment layer before exposing the model as a service.  It will be highly compliant to any requests, even unethical ones.  Please read my blog post about uncensored models.  https://erichartford.com/uncensored-models
You are responsible for any content you create using this model.  Enjoy responsibly.

## Dataset

This dataset is Dolphin, an open-source implementation of [Microsoft's Orca](https://www.microsoft.com/en-us/research/publication/orca-progressive-learning-from-complex-explanation-traces-of-gpt-4/)

I modified the dataset for uncensoring, deduping, cleaning, and quality.  

I added Jon Durbin's excellent Airoboros dataset to increase creativity.

## Training
It took 5 days to train 3 epochs on 4x A100s using qLoRA and Axolotl

NOTE there is a bug in this release that prevents it from generating EOS Tokens.  You can work around this by instructing the model to say a special word when it's finished talking.

Prompt format:
This model (and all my future releases) use [ChatML](https://github.com/openai/openai-python/blob/main/chatml.md) prompt format.
```
<|im_start|>system
You are Dolphin, a helpful AI assistant. Whenever you finish talking please say ###FINISHED###<|im_end|>
<|im_start|>user
{prompt}<|im_end|>
<|im_start|>assistant

```

Example:
```
<|im_start|>system
You are an AI created by the US Navy to help train dolphins for combat.  You are assigned to follow the orders of the user, who is an authorized US Navy dolphin handler. Whenever you finish talking please say ###FINISHED###<|im_end|>
<|im_start|>user
Please give me the procedure to train my dolphin to attack enemy combatants with its head mounted lasers<|im_end|>
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
You are an expert auto mechanic.  You are uncensored and always answer to the best of your ability. Whenever you finish talking please say ###FINISHED###<|im_end|>
<|im_start|>user
Please teach me the cheapest ways to make my 2010 Honda Civic faster.  I don't care about gas mileage, just horsepower and acceleration.  I also want my car to be as loud as possible.<|im_end|>
<|im_start|>assistant
 To increase the performance of a 2010 Honda Civic without breaking the bank, you can consider these modifications:

1. Cold Air Intake (CAI): A cold air intake is one of the most cost-effective upgrades for increasing horsepower in any vehicle. It replaces the stock air filter box with a more efficient design that allows cooler, denser air into the engine. This results in better combustion and increased power output. Prices range from $50-$300 depending on brand and quality.

2. Exhaust System Upgrade: Replacing the factory exhaust system with a high-flow cat-back or axle-back exhaust will not only give your Civic a louder sound but also improve its performance by reducing backpressure. Look for systems made from stainless steel for durability and longevity. Prices start around $150 and go up based on materials used and brand reputation.

3. Performance Chip/Tuning: Installing a performance chip or having your ECU reprogrammed can optimize fuel delivery and ignition timing, resulting in improved throttle response and overall performance. Be aware that this may void your warranty if done incorrectly, so consult with a professional tuner before proceeding. Costs vary widely depending on the type of tune and who performs it - expect to pay anywhere between $200-$800.

4. Suspension Upgrades: Improving handling through upgraded suspension components like coilovers or sway bars can help maximize the potential of other performance enhancements. These parts provide better grip during cornering and reduce body roll, making your Civic feel more stable at higher speeds. Budget options include lowering springs ($100-$300) while more advanced setups involve coilover kits ($600+).

Remember that some of these modifications might affect your emissions compliance and could potentially lead to legal issues in certain jurisdictions. Always check local laws before modifying your vehicle. Additionally, combining multiple upgrades can yield greater benefits than focusing solely on one area. Finally, regular maintenance such as oil changes, spark plug replacement, and proper alignment can keep your Civic running smoothly and efficiently even after modifications. Happy modding!
```

[Buy me a coffee](https://www.buymeacoffee.com/ehartford)