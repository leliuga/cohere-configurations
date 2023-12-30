---
language:
- eng
tags:
- sft
- StableLM
license:
- mit
datasets:
- LDJnr/Capybara
- LDJnr/LessWrong-Amplify-Instruct
- LDJnr/Pure-Dove
- LDJnr/Verified-Camel
---

## **Nous-Capybara-7B V1.9**

**This is currently the best 7B version of Capybara to use**

What's new compared to V1?: *V1.9 now leverages novel unalignment techniques that lead to more consistent and dynamic control, we also worked on enhanced quality curation for training data and a significantly better foundation model(Mistral)!*
 
The Capybara series is the first Nous collection of dataset and models made by fine-tuning mostly on data created by Nous in-house. 

We leverage our novel data synthesis technique called Amplify-instruct (Paper coming soon), the seed distribution and synthesis method are comprised of a synergistic combination of top performing existing data synthesis techniques and distributions used for SOTA models such as Airoboros, Evol-Instruct(WizardLM), Orca, Vicuna, Know_Logic, Lamini, FLASK and others, all into one lean holistically formed methodology for the dataset and model. The seed instructions used for the start of synthesized conversations are largely based on highly datasets like Airoboros, Know logic, EverythingLM, GPTeacher and even entirely new seed instructions derived from posts on the website LessWrong, as well as being supplemented with certain in-house multi-turn datasets like Dove(A successor to Puffin).

While performing great in it's current state, the current dataset used for fine-tuning is entirely contained within 20K training examples, this is 10 times smaller than many similar performing current models, this is signficant when it comes to scaling implications for our next generation of models once we scale our novel syntheiss methods to significantly more examples.

## Process of creation and special thank yous!

This model was fine-tuned by Nous Research as part of the Capybara/Amplify-Instruct project led by Luigi D.(LDJ) (Paper coming soon), as well as significant dataset formation contributions by J-Supha and general compute and experimentation management by Jeffrey Q. during ablations.

Special thank you to **A16Z** for sponsoring our training, as well as **Yield Protocol** for their support in financially sponsoring resources during the R&D of this project.

## Thank you to those of you that have indirectly contributed!

While most of the tokens within Capybara are newly synthsized and part of datasets like Puffin/Dove, we would like to credit the single-turn datasets we leveraged as seeds that are used to generate the multi-turn data as part of the Amplify-Instruct synthesis.

The datasets shown in green below are datasets that we sampled from to curate seeds that are used during Amplify-Instruct synthesis for this project.

Datasets in Blue are in-house curations that previously existed prior to Capybara.

![Capybara](https://i.imgur.com/yB58OoD.jpeg)

## Model Training

Nous-Capybara 7B V1.9 is a new model trained for multiple epochs on a dataset of roughly 20,000 carefully curated conversational examples, most of which are comprised of entirely new in-house synthesized tokens.

Additional data came from human curated CamelAI data, with the help of volunteers ranging from former Physics PhD's, Mathematicians, Biologists and more! 

## Prompt Format

The reccomended model usage is:

```
USER:

ASSISTANT:
```

## Mutli-Modality!

 - We currently have a Multi-modal model based on Capybara V1.9! 
https://huggingface.co/NousResearch/Obsidian-3B-V0.5
it is currently only available as a 3B sized model but larger versions coming!


## Notable Features:

 - Over 60% of the dataset is comprised of multi-turn conversations.(Most models are still only trained for single-turn conversations and no back and forths!)

 - Over 1,000 tokens average per conversation example! (Most models are trained on conversation data that is less than 300 tokens per example.)

 - Able to effectively do complex summaries of advanced topics and studies. (trained on hundreds of advanced difficult summary tasks developed in-house)

 - Ability to recall information upto late 2022 without internet.

 - Includes a portion of conversational data synthesized from less wrong posts, discussing very in-depth details and philosophies about the nature of reality, reasoning, rationality, self-improvement and related concepts.

## Example Outputs!:

![Capybara](https://img001.prntscr.com/file/img001/T9yYxR1xQSaK_UGdy3t2Cw.png)

![Capybara](https://img001.prntscr.com/file/img001/DQXqmKbsQQOIcgny1eoGNA.png)

![Capybara](https://img001.prntscr.com/file/img001/85X3L9ZxTsOKo3fUQ7GRVA.png)

## Benchmarks! (Coming soon!)
 
## Future Changes

This is a relatively early build amongst the grand plans for the future of Capybara!  

## Future model sizes

Capybara V1.9 now currently has a 3B ad 7B size, and we plan to eventually have a 13B and 70B version in the future, as well as a potential 1B version based on phi-1.5 or Tiny Llama.

## How you can help!

In the near future we plan on leveraging the help of domain specific expert volunteers to eliminate any mathematically/verifiably incorrect answers from our training curations. 

If you have at-least a bachelors in mathematics, physics, biology or chemistry and would like to volunteer even just 30 minutes of your expertise time, please contact LDJ on discord!

## Dataset contamination.

We have checked the capybara dataset for contamination for several of the most popular datasets and can confirm that there is no contaminaton found.

We leveraged minhash to check for 100%, 99%, 98% and 97% similarity matches between our data and the questions and answers in benchmarks, we found no exact matches, nor did we find any matches down to the 97% similarity level.

The following are benchmarks we checked for contamination against our dataset:

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