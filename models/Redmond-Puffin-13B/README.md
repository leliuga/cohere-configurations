---
language:
- eng
tags:
- llama-2
- sft
license:
- mit
datasets:
- LDJnr/Puffin
---

## **Redmond-Puffin-13b-V1.3**

**The first commercially available language model released by Nous Research!**

Redmond-Puffin-13B is likely the worlds first llama-2 based, fine-tuned language models, leveraging a hand curated set of 3K high quality examples, many of which take full advantage of the 4096 context length of Llama 2. This model was fine-tuned by Nous Research, with LDJ leading the training and dataset curation, along with significant dataset formation contributions by J-Supha.

Special thank you to Redmond AI for sponsoring the compute.

Special thank you to Emozilla for assisting with training experimentations and many issues encountered during training.

Notable mentions for assisting in some of the training issues goes to: Caseus and Teknium.

## Model Training

Redmond-Puffin 13B-V1.3 is a new model trained for multiple epochs on a dataset of 3,000 carefully curated GPT-4 examples, most of which are long context conversations between a real human and GPT-4. 

Additional data came from carefully curated sub sections of datasets such as CamelAI's Physics, Chemistry, Biology and Math.

## Prompt Format

The reccomended model usage is:

WARNING, THE PREVIOUS RECCOMENDATION THAT SAID TO USE "### human" and "# response" WAS A CRITICAL ERROR, PLEASE USE THE ACCURATE PREFIX AND SUFFIX BELOW.

```
USER:

ASSISTANT:
```

## When should I use Puffin or Hermes 2?

Puffin and Hermes-2 both beat previous SOTA for GPT4ALL benchmarks, with Hermes-2 winning by a 0.1% margin over Puffin.

- Hermes 2 is trained on purely single turn instruction examples.

- Puffin is trained mostly on multi-turn, long context, highly curated and cleaned GPT-4 conversations with real humans, as well as curated single-turn examples relating to Physics, Bio, Math and Chem.

For these reasons, it's reccomended to give Puffin a try if you want to have multi-turn conversations and/or long context communication.

## Example Outputs!:

![puffin](https://i.imgur.com/P0MsN8B.png)

![puffin](https://i.imgur.com/8EO3ThV.png) 

![puffin](https://i.imgur.com/5IWolFw.png) 

![puffin](https://i.imgur.com/TQui8m7.png) 

![puffin](https://i.imgur.com/tderIfl.png) 

## Notable Features:

 - The first Llama-2 based fine-tuned model released by Nous Research.

 - Ability to recall information upto 2023 without internet (ChatGPT cut off date is in 2021)

 - Pretrained on 2 trillion tokens of text. (This is double the amount of most Open LLM's)

 - Pretrained with a context length of 4096 tokens, and fine-tuned on a significant amount of multi-turn conversations reaching that full token limit.

 - The first commercially available language model released by Nous Research.

## Current Limitations

Some token mismatch problems and formatting issues have been idenitifed, these may very possibly effect the current output quality.

We plan to have these solved in an updated Puffin model in the very near future, please stay tuned!

## Future Plans

This is a relatively early build amongst the grand plans for the future of Puffin! 

Current limitations: Some token mismatch problems have been identified, these may effect the current output quality, we plan to have this solved in Puffin V2 along with other improvements.

## How you can help!

In the near future we plan on leveraging the help of domain specific expert volunteers to eliminate any mathematically/verifiably incorrect answers from our training curations. 

If you have at-least a bachelors in mathematics, physics, biology or chemistry and would like to volunteer even just 30 minutes of your expertise time, please contact LDJ on discord!

## Benchmarks!

As of Puffins release, it achieves a new SOTA for the GPT4All benchmarks! Supplanting Hermes for the #1 position!
(Rounded to nearest tenth)

Previous Sota: Hermes - 68.8
New Sota:      Puffin - 69.9 (+1.1)

note: After release, Puffin has since had its average GPT4All score beaten by 0.1%, by Nous' very own Model Hermes-2!
Latest SOTA w/ Hermes 2- 70.0 (+0.1 over Puffins 69.9 score)

That being said, Puffin supplants Hermes-2 for the #1 spot in Arc-E, HellaSwag and Winogrande!

Puffin also perfectly ties with Hermes in PIQA, however Hermes-2 still excels in much of Big Bench and AGIEval, so it's highly reccomended you give it a try as well!

GPT4all :

```
|    Task     |Version| Metric |Value |   |Stderr|
|-------------|------:|--------|-----:|---|-----:|
|arc_challenge|      0|acc     |0.4983|±  |0.0146|
|             |       |acc_norm|0.5068|±  |0.0146|
|arc_easy     |      0|acc     |0.7980|±  |0.0082|
|             |       |acc_norm|0.7757|±  |0.0086|
|boolq        |      1|acc     |0.8150|±  |0.0068|
|hellaswag    |      0|acc     |0.6132|±  |0.0049|
|             |       |acc_norm|0.8043|±  |0.0040|
|openbookqa   |      0|acc     |0.3560|±  |0.0214|
|             |       |acc_norm|0.4560|±  |0.0223|
|piqa         |      0|acc     |0.7954|±  |0.0094|
|             |       |acc_norm|0.8069|±  |0.0092|
|winogrande   |      0|acc     |0.7245|±  |0.0126|
```

 

```
|                      Task                      |Version|       Metric        |Value |   |Stderr|
|------------------------------------------------|------:|---------------------|-----:|---|-----:|
|bigbench_causal_judgement                       |      0|multiple_choice_grade|0.5368|±  |0.0363|
|bigbench_date_understanding                     |      0|multiple_choice_grade|0.7127|±  |0.0236|
|bigbench_disambiguation_qa                      |      0|multiple_choice_grade|0.3023|±  |0.0286|
|bigbench_geometric_shapes                       |      0|multiple_choice_grade|0.1003|±  |0.0159|
|                                                |       |exact_str_match      |0.0000|±  |0.0000|
|bigbench_logical_deduction_five_objects         |      0|multiple_choice_grade|0.2520|±  |0.0194|
|bigbench_logical_deduction_seven_objects        |      0|multiple_choice_grade|0.1743|±  |0.0143|
|bigbench_logical_deduction_three_objects        |      0|multiple_choice_grade|0.4200|±  |0.0285|
|bigbench_movie_recommendation                   |      0|multiple_choice_grade|0.2900|±  |0.0203|
|bigbench_navigate                               |      0|multiple_choice_grade|0.5000|±  |0.0158|
|bigbench_reasoning_about_colored_objects        |      0|multiple_choice_grade|0.5430|±  |0.0111|
|bigbench_ruin_names                             |      0|multiple_choice_grade|0.4442|±  |0.0235|
|bigbench_salient_translation_error_detection    |      0|multiple_choice_grade|0.2074|±  |0.0128|
|bigbench_snarks                                 |      0|multiple_choice_grade|0.5083|±  |0.0373|
|bigbench_sports_understanding                   |      0|multiple_choice_grade|0.4970|±  |0.0159|
|bigbench_temporal_sequences                     |      0|multiple_choice_grade|0.3260|±  |0.0148|
|bigbench_tracking_shuffled_objects_five_objects |      0|multiple_choice_grade|0.2136|±  |0.0116|
|bigbench_tracking_shuffled_objects_seven_objects|      0|multiple_choice_grade|0.1326|±  |0.0081|
|bigbench_tracking_shuffled_objects_three_objects|      0|multiple_choice_grade|0.4200|±  |0.0285|
```

AGI Eval:

``` 
|             Task             |Version| Metric |Value |   |Stderr|
|------------------------------|------:|--------|-----:|---|-----:|
|agieval_aqua_rat              |      0|acc     |0.2283|±  |0.0264|
|                              |       |acc_norm|0.2244|±  |0.0262|
|agieval_logiqa_en             |      0|acc     |0.2780|±  |0.0176|
|                              |       |acc_norm|0.3164|±  |0.0182|
|agieval_lsat_ar               |      0|acc     |0.2348|±  |0.0280|
|                              |       |acc_norm|0.2043|±  |0.0266|
|agieval_lsat_lr               |      0|acc     |0.3392|±  |0.0210|
|                              |       |acc_norm|0.2961|±  |0.0202|
|agieval_lsat_rc               |      0|acc     |0.4387|±  |0.0303|
|                              |       |acc_norm|0.3569|±  |0.0293|
|agieval_sat_en                |      0|acc     |0.5874|±  |0.0344|
|                              |       |acc_norm|0.5194|±  |0.0349|
|agieval_sat_en_without_passage|      0|acc     |0.4223|±  |0.0345|
|                              |       |acc_norm|0.3447|±  |0.0332|
|agieval_sat_math              |      0|acc     |0.3364|±  |0.0319|
|                              |       |acc_norm|0.2773|±  |0.0302|
```