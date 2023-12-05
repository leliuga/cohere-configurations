---
license: llama2
inference:
  parameters:
    do_sample: false
    max_length: 200
widget:
- text: "CREATE TABLE stadium (\n    stadium_id number,\n    location text,\n    name text,\n    capacity number,\n)\n\n-- Using valid SQLite, answer the following questions for the tables provided above.\n\n-- how many stadiums in total?\n\nSELECT"
  example_title: "Number stadiums"
- text: "CREATE TABLE work_orders ( ID NUMBER, CREATED_AT TEXT, COST FLOAT, INVOICE_AMOUNT FLOAT, IS_DUE BOOLEAN, IS_OPEN BOOLEAN, IS_OVERDUE BOOLEAN, COUNTRY_NAME TEXT, )\n\n-- Using valid SQLite, answer the following questions for the tables provided above.\n\n-- how many work orders are open?\n\nSELECT"
  example_title: "Open work orders"
- text: "CREATE TABLE stadium ( stadium_id number, location text, name text, capacity number, highest number, lowest number, average number )\n\nCREATE TABLE singer ( singer_id number, name text, country text, song_name text, song_release_year text, age number, is_male others )\n\nCREATE TABLE concert ( concert_id number, concert_name text, theme text, stadium_id text, year text )\n\nCREATE TABLE singer_in_concert ( concert_id number, singer_id text )\n\n-- Using valid SQLite, answer the following questions for the tables provided above.\n\n-- What is the maximum, the average, and the minimum capacity of stadiums ?\n\nSELECT"
  example_title: "Stadium capacity"
---

# NSQL-Llama-2-7B

## Model Description

NSQL is a family of autoregressive open-source large foundation models (FMs) designed specifically for SQL generation tasks.

In this repository we are introducing a new member of NSQL, NSQL-Llama-2-7B. It's based on Meta's original [Llama-2 7B model](https://huggingface.co/meta-llama/Llama-2-7b) and further pre-trained on a dataset of general SQL queries and then fine-tuned on a dataset composed of text-to-SQL pairs.

## Training Data

The general SQL queries are the SQL subset from [The Stack](https://huggingface.co/datasets/bigcode/the-stack), containing 1M training samples. The labeled text-to-SQL pairs come from more than 20 public sources across the web from standard datasets. We hold out Spider and GeoQuery datasets for use in evaluation.

## Evaluation Data

We evaluate our models on two text-to-SQL benchmarks: Spider and GeoQuery.

## Training Procedure

NSQL was trained using cross-entropy loss to maximize the likelihood of sequential inputs. For finetuning on text-to-SQL pairs, we only compute the loss over the SQL portion of the pair. The model is trained using 80GB A100s, leveraging data and model parallelism. We pre-trained for 3 epochs and fine-tuned for 10 epochs.

## Intended Use and Limitations

The model was designed for text-to-SQL generation tasks from given table schema and natural language prompts. The model works best with the prompt format defined below and outputting `SELECT` queries.

## How to Use

Example 1:

```python
import torch
from transformers import AutoTokenizer, AutoModelForCausalLM
tokenizer = AutoTokenizer.from_pretrained("NumbersStation/nsql-llama-2-7B")
model = AutoModelForCausalLM.from_pretrained("NumbersStation/nsql-llama-2-7B", torch_dtype=torch.bfloat16)

text = """CREATE TABLE stadium (
    stadium_id number,
    location text,
    name text,
    capacity number,
    highest number,
    lowest number,
    average number
)

CREATE TABLE singer (
    singer_id number,
    name text,
    country text,
    song_name text,
    song_release_year text,
    age number,
    is_male others
)

CREATE TABLE concert (
    concert_id number,
    concert_name text,
    theme text,
    stadium_id text,
    year text
)

CREATE TABLE singer_in_concert (
    concert_id number,
    singer_id text
)

-- Using valid SQLite, answer the following questions for the tables provided above.

-- What is the maximum, the average, and the minimum capacity of stadiums ?

SELECT"""

input_ids = tokenizer(text, return_tensors="pt").input_ids

generated_ids = model.generate(input_ids, max_length=500)
print(tokenizer.decode(generated_ids[0], skip_special_tokens=True))
```

Example 2:

```python
import torch
from transformers import AutoTokenizer, AutoModelForCausalLM
tokenizer = AutoTokenizer.from_pretrained("NumbersStation/nsql-llama-2-7B")
model = AutoModelForCausalLM.from_pretrained("NumbersStation/nsql-llama-2-7B", torch_dtype=torch.bfloat16)

text = """CREATE TABLE stadium (
    stadium_id number,
    location text,
    name text,
    capacity number,
)

-- Using valid SQLite, answer the following questions for the tables provided above.

-- how many stadiums in total?

SELECT"""

input_ids = tokenizer(text, return_tensors="pt").input_ids

generated_ids = model.generate(input_ids, max_length=500)
print(tokenizer.decode(generated_ids[0], skip_special_tokens=True))
```

Example 3:

```python
import torch
from transformers import AutoTokenizer, AutoModelForCausalLM
tokenizer = AutoTokenizer.from_pretrained("NumbersStation/nsql-llama-2-7B")
model = AutoModelForCausalLM.from_pretrained("NumbersStation/nsql-llama-2-7B", torch_dtype=torch.bfloat16)

text = """CREATE TABLE work_orders (
    ID NUMBER,
    CREATED_AT TEXT,
    COST FLOAT,
    INVOICE_AMOUNT FLOAT,
    IS_DUE BOOLEAN,
    IS_OPEN BOOLEAN,
    IS_OVERDUE BOOLEAN,
    COUNTRY_NAME TEXT,
)

-- Using valid SQLite, answer the following questions for the tables provided above.

-- how many work orders are open?

SELECT"""

input_ids = tokenizer(text, return_tensors="pt").input_ids

generated_ids = model.generate(input_ids, max_length=500)
print(tokenizer.decode(generated_ids[0], skip_special_tokens=True))
```



For more information (e.g., run with your local database), please find examples in [this repository](https://github.com/NumbersStationAI/NSQL).
