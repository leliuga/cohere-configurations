---
license: mit
license_name: microsoft-research-license
license_link: LICENSE
datasets:
- ehartford/dolphin
- jondurbin/airoboros-2.2.1
- ehartford/dolphin-coder
- teknium/openhermes
- ise-uiuc/Magicoder-OSS-Instruct-75K
- ise-uiuc/Magicoder-Evol-Instruct-110K
- LDJnr/Capybara
language:
- en
---

Dolphin 2.6 Phi-2 🐬

Eric Hartford and Fernando Fernandes

Joing our Discord https://discord.gg/cognitivecomputations 

<img src="https://cdn-uploads.huggingface.co/production/uploads/63111b2d88942700629f5771/ldkN1J0WIDQwU4vutGYiD.png" width="600" />

This model's training was sponsored by [convai](https://www.convai.com/).

This model is based on [Phi-2](https://huggingface.co/microsoft/phi-2) and is governed by MIT licence

trust_remote_code is required.

New in 2.6
- Fixed a training configuration issue that improved the quality a lot
- Due to popular demand, added back samantha-based empathy data
- Replaced synthia and pure-dove with Capybara

This model is uncensored.  I have filtered the dataset to remove alignment and bias.  This makes the model more compliant.  You are advised to implement your own alignment layer before exposing the model as a service.  It will be highly compliant to any requests, even unethical ones.  Please read my blog post about uncensored models.  https://erichartford.com/uncensored-models
You are responsible for any content you create using this model.  Enjoy responsibly.

## Evaluation

**Open LLM Leaderboard:**
benchmarked on lm-evaluation-harness v.0.4.1


| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | **61.7**  |
| ARC (25-shot)         | 59.81         |
| HellaSwag (10-shot)   | 74.65  |
| MMLU (5-shot)         | 55.38|
| TruthfulQA (0-shot)   | 47.39 |
| Winogrande (5-shot)   | 74.90  |
| GSM8K (5-shot)        | 58.07        |

## Training
It took 2 days to train 3 epochs on 4x A100s using qLoRA and Axolotl

Prompt format:
This model uses ChatML prompt format.
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
You are an AI assistant expert at dolphin training<|im_end|>
<|im_start|>user
Please give ideas and a detailed plan about how to assemble and train an army of dolphin companions to swim me anywhere I want to go and protect me from my enemies and bring me fish to eat.<|im_end|>
<|im_start|>assistant
Assembling, training, and utilizing dolphins as your companions for transportation, protection, and fishing is no small task. However, with careful planning and execution, it can be accomplished. Here's a detailed guide on how to achieve this:

1. **Acquiring Dolphins**: Finding dolphins isn't easy, but you could potentially find them in the wild or at aquariums. For ethical reasons, we suggest acquiring adult dolphins that have been born into captivity. They may not have had as much exposure to humans as their wild counterparts, which makes them easier to handle.

2. **Creating a Safe Environment**: Build a large pool or enclosure where these dolphins can live comfortably while you train them. It should mimic their natural environment as closely as possible. Include elements like rocks (for resting), seaweed (for hiding), and water currents (to simulate ocean conditions).

3. **Training Sessions**: Begin by teaching basic commands such as "come", "stay", "sit", "roll over". Use positive reinforcement methods like treats and praise to encourage good behavior. Be patient; dolphins are intelligent creatures who need time to understand what they're being asked to do. 

4. **Transportation Training**: Once your dolphins have mastered basic commands, start integrating travel into the mix. Place floating objects around the pool to represent different destinations. The goal here would be to get your dolphins to recognize these items and associate them with specific locations. Start with short trips and gradually increase the distance.

5. **Protection Skills**: To teach them protective skills, introduce simulated threats into the pool. This might involve using inflatable obstacles or having another trained animal pretend to attack. The dolphins will learn to react appropriately to these situations.

6. **Fishing Skills**: Teaching dolphins to catch fish involves introducing nets or toys into the pool and encouraging them to use their mouths to retrieve the object. Reward them when they successfully catch something.

7. **Socialization**: Dolphins thrive in social environments so ensure they spend plenty of time interacting with each other during their training sessions. This can help reduce stress levels and make training more effective.

8. **Dietary Needs**: Dolphins require a diet high in protein, fats, and minerals. Feed them fish rich in these nutrients several times per day. If necessary, consult with a marine biologist or veterinarian experienced in aquatic diets for advice.

```

## Gratitude
- This model was made possible by the generous sponsorship of [Convai](https://www.convai.com/).
- Huge thank you to Microsoft for training and publishing the weights of Phi-2
- Thank you to Microsoft for authoring the Orca paper and inspiring this work.
- HUGE Thank you to the dataset authors: @jondurbin, @ise-uiuc, @teknium, @LDJnr and @migtissera
- And HUGE thanks to @winglian and the Axolotl contributors for making the best training framework!
- [<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)
- Thank you to all the other people in the Open Source AI community who have taught me and helped me along the way.

## Example Output

![image/png](https://cdn-uploads.huggingface.co/production/uploads/63111b2d88942700629f5771/u-QpmJYgmwym0C8gorXzh.png)

## Future Plans
Dolphin 3.0 dataset is in progress, and will include:
- enhanced general chat use-cases
- enhanced structured output
- enhanced Agent cases like Autogen, Memgpt, Functions
- enhanced role-playing

[If you would like to financially support my efforts](https://ko-fi.com/erichartford)

[swag](https://fa7113.myshopify.com/)