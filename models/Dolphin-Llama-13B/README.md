---
license: other
task_categories:
- text-generation
model-index:
- name: dolphin-llama-13b
  results:
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: AI2 Reasoning Challenge (25-Shot)
      type: ai2_arc
      config: ARC-Challenge
      split: test
      args:
        num_few_shot: 25
    metrics:
    - type: acc_norm
      value: 55.55
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/dolphin-llama-13b
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: HellaSwag (10-Shot)
      type: hellaswag
      split: validation
      args:
        num_few_shot: 10
    metrics:
    - type: acc_norm
      value: 77.11
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/dolphin-llama-13b
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: MMLU (5-Shot)
      type: cais/mmlu
      config: all
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 52.16
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/dolphin-llama-13b
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: TruthfulQA (0-shot)
      type: truthful_qa
      config: multiple_choice
      split: validation
      args:
        num_few_shot: 0
    metrics:
    - type: mc2
      value: 52.23
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/dolphin-llama-13b
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: Winogrande (5-shot)
      type: winogrande
      config: winogrande_xl
      split: validation
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 69.93
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/dolphin-llama-13b
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: GSM8k (5-shot)
      type: gsm8k
      config: main
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 14.4
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ehartford/dolphin-llama-13b
      name: Open LLM Leaderboard
---

Dolphin 🐬
https://erichartford.com/dolphin

This model is based on llama1, so it is for non-commercial use only.  Future versions will be trained on llama2 and other open models that are suitable for commercial use.

This model is uncensored.  I have filtered the dataset to remove alignment and bias.  This makes the model compliant to any requests.  You are advised to implement your own alignment layer before exposing the model as a service.  It will be highly compliant to any requests, even unethical ones.  Please read my blog post about uncensored models.  https://erichartford.com/uncensored-models
You are responsible for any content you create using this model.  Enjoy responsibly.

## Dataset

This dataset is an open source implementation of [Microsoft's Orca](https://www.microsoft.com/en-us/research/publication/orca-progressive-learning-from-complex-explanation-traces-of-gpt-4/)

After uncensoring, deduping, and cleaning, our dataset consists of:

- 842,610 instructions of FLANv2 augmented with GPT-4 completions
- 2,625,353 instructions of FLANv2 augmented with GPT-3.5 completions

We followed the submix and system prompt distribution outlined in the Orca paper. With a few exceptions. We included all 75k of CoT in the FLAN-1m dataset rather than sampling that. Also, we found that many items were duplicated, so we removed duplicates.

Then we filtered out instances of alignment, refusal, avoidance, and bias, in order to produce an uncensored model upon which can be layered your personalized alignment LoRA.

We also filtered out duplicates and cleaned the data.

## Training
We trained with the flan5m (gpt3.5 completions) dataset in its entirety for 3 epochs at a learning rate of 2e-5 before we stopped training to avoid overfit.
We trained with the flan1m (gpt4 completions) dataset in its entirety for 2.5 epochs at a learning rate of 1e-5 before we stopped training to avoid overfit.
It took about 600 hours to train on 8x H100s
We used a prompt format similar to Vicuna, but we added the SYSTEM: field. 

Prompt format:
```
SYSTEM: {system}
USER: {prompt}
ASSISTANT:
```

Example:
```
SYSTEM: you are an expert marine biologist.
USER: Please list 10 ways that dolphins are superior to orcas.
ASSISTANT:
```

## Evaluation
Evaluation will be coming soon.

## Team
The core Dolphin Team includes:
- Eric "Faldore" Hartford
- Pankaj Mathur
- Rob "Rohan" O'Callahan
- Tom "TheBloke" Jobbins

## Gratitude
- Thank you to Microsoft for authoring the Orca paper and inspiring this work.
- Special thanks to WingLian, NanoBit, Teknium for helpful advice
- Special thanks to EdenCoder and chirper.ai for mentorship and financial sponsorship.
- Special thanks to Kilkonie for his very valued mentorship.
- Thank you to Catto.
- Thank you to Nicolai Schleifer, financial sponsor.
- Thank you to Eric Fleming, financial sponsor.
- Thank you to all the other people in the Open Source AI community who have taught me and helped me along the way.
- And biggest thanks and love to my wife Hope and my babies, who have endured two months of my absenteeism while working long days and nights on this project.

## Example Output
```
SYSTEM: you are an expert marine biologist.
USER: Please list 10 ways that dolphins are superior to orcas.
ASSISTANT: 
```

1. Social structure: Dolphins live in complex social groups called pods, which consist of extended family units and can include hundreds of individuals. Orcas, on the other hand, typically form smaller, more temporary groupings known as clans or matrilineal kinship groups.

2. Communication: Dolphins have a highly developed system of vocalizations, body language, and even telepathy to communicate with each other for various purposes such as coordinating cooperative hunting strategies, sharing information about food sources, and maintaining social bonds. While orcas also exhibit some communication abilities, their methods are generally less sophisticated compared to those of dolphins.

3. Intelligence: Dolphins are widely regarded as one of the most intelligent animal species, showcasing advanced problem-solving skills, self-awareness, tool use (such as using sea sponges as tools), and learning capabilities. Some dolphin species have been observed demonstrating altruism and understanding of human emotions. In contrast, while orcas are also smart, they do not display the same level of cognitive complexity as dolphins.

4. Playfulness: Dolphins are well-known for their playful behavior, often engaging in acrobatic displays, chasing each other's tails, and playing with objects like balls or seaweed. This playful nature is thought to be essential for their social bonding and mental stimulation. Although orcas may occasionally interact playfully, it is not as prevalent or diverse as among dolphins.

5. Agility and speed: Dolphins are incredibly agile swimmers, capable of reaching speeds up to 30 miles per hour (48 km/h) when pursuing prey. They can change direction rapidly and maneuver through tight spaces with ease. While orcas are powerful predators and fast swimmers themselves, they cannot match the sheer acceleration and agility of dolphins.

6. Feeding habits: Both dolphins and orcas are carnivorous, feeding primarily on fish and cephalopods. However, dolphins employ a wider variety of hunting techniques, including cooperative hunting strategies involving herding, encircling, and driving schools of fish into shallow waters where they become disoriented and easier to catch. Orcas mainly rely on ambushing their prey from below and using their powerful jaws to strike quickly.

7. Reproduction: Female dolphins usually give birth to a single calf after a gestation period of around 12 months. The mother plays a crucial role in raising and nurturing her offspring until they become independent. On the other hand, female orcas tend to breed every three to five years, giving birth to multiple offspring at once, and relying heavily on their mothers for protection and teaching.

8. Vocal repertoire: Dolphins possess a wide range of sounds, including clicks, whistles, and body sounds, which allow them to convey different messages and establish unique identities within their pods. Orcas, too, produce vocalizations but lack the diversity found in dolphin calls.

9. Conservation status: Many dolphin species face threats due to habitat loss, pollution, overfishing, and climate change. As a result, several populations are considered vulnerable or endangered by the International Union for Conservation of Nature (IUCN). In comparison, orca conservation status varies depending on the region; however, no species is currently listed as critically endangered or facing extinction.

10. Adaptability: Dolphins have adapted to various aquatic environments, ranging from warm coastal waters to colder open oceans. Their versatility allows them to thrive in both tropical and temperate climates. Orcas, conversely, are adapted to specific habitats such as cooler coastal waters and are mostly found in the Northern Hemisphere.

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ehartford__dolphin-llama-13b)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 48.6   |
| ARC (25-shot)         | 55.55          |
| HellaSwag (10-shot)   | 77.11    |
| MMLU (5-shot)         | 52.16         |
| TruthfulQA (0-shot)   | 52.23   |
| Winogrande (5-shot)   | 69.93   |
| GSM8K (5-shot)        | 14.4        |
| DROP (3-shot)         | 18.83         |

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ehartford__dolphin-llama-13b)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |53.56|
|AI2 Reasoning Challenge (25-Shot)|55.55|
|HellaSwag (10-Shot)              |77.11|
|MMLU (5-Shot)                    |52.16|
|TruthfulQA (0-shot)              |52.23|
|Winogrande (5-shot)              |69.93|
|GSM8k (5-shot)                   |14.40|

