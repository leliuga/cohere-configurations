---
language:
- eng
tags:
- llama-2
- sft
license:
- mit
datasets:
- LDJnr/Capybara
- LDJnr/LessWrong-Amplify-Instruct
- LDJnr/Pure-Dove
- LDJnr/Verified-Camel
---

## **Nous-Capybara-7B V1**

**MUCH BETTER MISTRAL BASED VERSION IS OUT NOW AS CAPYBARA V1.9**

The Capybara series is made by fine-tuning on data that is created by Nous with our novel data synthesis technique called Amplify-instruct, the seed distribution and synthesis method are comprised of a synergistic combination of top performing existing data synthesis techniques and distributions used for SOTA models such as Airoboros, Evol-Instruct, Orca, Vicuna, Know_Logic, Lamini, FLASK and others, all into one lean holistically formed dataset and model. The seed instructions used for the start of synthesized conversations are largely based on highly datasets like Airoboros, Know logic, EverythingLM, GPTeacher and even entirely new seed instructions derived from posts on the website LessWrong, as well as being supplemented with certain in-house multi-turn datasets like Dove(A successor to Puffin).

While performing great in it's current state, the current dataset used for fine-tuning is entirely contained within 20K training examples, mostly comprised of newly synthesized conversation tokens that have never previously been used for AI training to our knowledge.

This small fine-tune dataset has significant implications for how we'll be able to scale model abilities in the future! This model is currently 20K examples while matching benchmarks of notable 300K example datasets that are 10 times the size!

## Process of creation and special thank yous!

This model was fine-tuned by Nous Research, with LDJ leading the training and dataset curation, along with significant dataset formation contributions by J-Supha, Also thank you to Emozilla for also assisting to expedite the training experimentation process.

Special thank you to **A16Z** for sponsoring our training, as well as **Yield Protocol** for their support in resources during R&D of aspects outside of training, such as dataset development/synthesis.

## Thank you to those of you that have indirectly contributed!

While most of the tokens within Capybara are newly synthsized and part of datasets like Puffin/Dove, we would like to credit the single-turn datasets we leveraged as seeds that are used to generate the multi-turn data as part of the Amplify-Instruct synthesis.

The datasets shown in green below are datasets that we sampled from to curate seeds that are used during Amplify-Instruct synthesis for this project.

![Capybara](https://i.imgur.com/yB58OoD.jpeg)

## Model Training

Nous-Capybara 7B is a new model trained for multiple epochs on a dataset of roughly 20,000 carefully curated conversational examples, most of which are comprised of entirely new in-house synthesized tokens that previously didn't exist on HuggingFace.

Additional data came from manually curated CamelAI data, with the help of volunteers ranging from former Physics PhD's, Mathematicians, Biologists and more! 

## Prompt Format

The reccomended model usage is:

```
USER:

ASSISTANT:
```

## Notable Features:

 - The first Nous model trained on over 10,000 multi-turn conversations.

 - Over 1,000 tokens average per conversation example and multiple back and forth turns per conversation! Most models are still trained for only single-turn conversations and less than 300 tokens per example!

 - Able to effectively do complex summaries of advanced topics and studies.

 - Ability to recall information upto late 2022 without internet.

 - Includes a portion of conversational data synthesized from less wrong posts, discussing very in-depth about the nature of rationality, reasoning, self-improvement and related concepts.

## Example Outputs!:

![Capybara](https://img001.prntscr.com/file/img001/T9yYxR1xQSaK_UGdy3t2Cw.png)

![Capybara](https://img001.prntscr.com/file/img001/DQXqmKbsQQOIcgny1eoGNA.png)

![Capybara](https://img001.prntscr.com/file/img001/85X3L9ZxTsOKo3fUQ7GRVA.png)

## Benchmarks! (Important to note that all mentioned benchmarks are single-turn and don't test multi-turn capabilities, Capybara should excel even further at multi-turn conversational tasks than what benchmark comparisons show.)

![Capybara](https://i.imgur.com/n8lkmyK.png)
 

## Future Changes

This is a relatively early build amongst the grand plans for the future of Capybara! 

[IT IS NOW RECCOMENDED TO USE CAPYBARA V1.9 FOR SIGNIFICANTLY BETTER OVERALL CAPABILITIES]

## Future model sizes

We plan on releasing a 3B, 13B and 70B version, as well as a potential 1B version based on phi-1.5 or similar architectures.

## How you can help!

In the near future we plan on leveraging the help of domain specific expert volunteers to eliminate any mathematically/verifiably incorrect answers from our training curations. 

If you have at-least a bachelors in mathematics, physics, biology or chemistry and would like to volunteer even just 30 minutes of your expertise time, please contact LDJ on discord!

## Dataset contamination.

We checked for 100%, 99%, 98% and 97% similarity matches between our data and many popular benchmarks, we found no matches!

The following are benchmarks we checked for contamination for:

- HumanEval

- AGIEval

- TruthfulQA

- MMLU

- GPT4All

```
@article{daniele2023amplify-instruct,
  title={Amplify-Instruct: Synthetically Generated Diverse Multi-turn Conversations for Effecient LLM Training.},
  author={Daniele, Luigi and Suphavadeeprasit},
  journal={arXiv preprint arXiv:(comming soon)},
  year={2023}
}
```