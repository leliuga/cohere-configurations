---
license: apache-2.0
datasets:
- databricks/databricks-dolly-15k
pipeline_tag: text-generation
---
# Instruct_Mixtral-8x7B-v0.1_Dolly15K
Fine-tuned from Mixtral-8x7B-v0.1， used Dolly15k for the dataset. 85% for training, 14.9% validation, 0.1% test.  Trained for 1.0 epochs using QLora.  Trained with 1024 context window.

# Model Details
* **Trained by**: trained by [Brillibits](https://www.youtube.com/@Brillibits).
* **Model type:**  **Instruct_Mixtral-8x7B-v0.1_Dolly15K** is an auto-regressive language model based on the Llama 2 transformer architecture.
* **Language(s)**: English
* **License for Instruct_Mixtral-8x7B-v0.1_Dolly15K**: apache-2.0 license


# Prompting

## Prompt Template With Context

```
Write a 10-line poem about a given topic

Input:

The topic is about racecars

Output:
```
## Prompt Template Without Context
```
Who was the was the second president of the United States?

Output:
```

## Professional Assistance
This model and other models like it are great, but where LLMs hold the most promise is when they are applied on custom data to automate a wide variety of tasks

If you have a dataset and want to see if you might be able to apply that data to automate some tasks, and you are looking for professional assistance, contact me [here](mailto:blakecmallory@gmail.com)