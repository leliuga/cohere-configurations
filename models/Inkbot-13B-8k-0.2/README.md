# Model Card for Inkbot

## Model Details
Inkbot is a conversational AI model designed to interpret and respond to structured prompts with or without contextual information. Its unique feature lies in the structured prompt system, allowing users to engage in dynamic dialogues that can evolve based on context, metadata, and user input.

### Performance
- The model excels in RAG type queries, answering from context, and overriding memory when necessary.
- It can handle very large contexts, but may sometimes enter a repeating text loop, especially during complex tasks.
- The model is intended to be more functional and less chatty, avoiding the waste of tokens on superfluous language.

## How to Use
Inkbot uses a structured prompt template system.


### Prompt Template Structure

#### With Context:
```
<#meta#>
- Date: {current_date}
- Task: {task_name}
<#system#>
{system_prompt}
<#chat#>
<#user#>
{user}
<#user_context#>
{user_context}
<#bot#>
{bot}
```

#### Without Context:
```
<#meta#>
- Date: {current_date}
- Task: {task_name}
<#system#>
{system_prompt}
<#chat#>
<#user#>
{user}
<#bot#>
{bot}
```

### Examples

#### With Context:
```
<#meta#>
- Date: 2023-07-20
- Task: general
<#system#>
A dialogue focused on determining dates based on context clues.
<#chat#>
<#user#>
What is today's date?
<#bot#>
The date is: 2023-07-20
<#user#>
When did the event in the provided context take place?
<#user_context#>
It's the day after new years eve, and I'm starting 2018 with a headache.
<#bot#>
The event in the context took place on: 2018-01-01
```

#### Without Context:
```
<#meta#>
- Date: 2023-07-20
- Task: general
<#system#>
A simple dialogue about daily queries.
<#chat#>
<#user#>
Will it rain today?
<#bot#>
I am not sure, without knowing your location I am unable to determine the weather.
<#user#>
Can you compose a haiku for me?
<#bot#>
Whispers of the breeze,
Nature's melodies unfold,
Dawn dispels the night.
```

## Task Options

The model has been trained on a variety of tasks. Below is a breakdown of these tasks, along with example prompts to help guide your usage.

---

### 1. Text Refinement

Tasks in this category aim to improve the quality and clarity of the text.

#### clarity (100 examples)
*Example Prompts:*
- "Please read the content below and revise it to ensure clear communication and understandability."
- "Ensure the following text is easy for any reader to understand."

#### coherence (750 examples)
*Example Prompts:*
- "Review the content and adjust it to ensure it has a logical consistency and flow."
- "Make sure the following paragraphs connect seamlessly."

#### formalize (800 examples)
*Example Prompts:*
- "Please convert the following informal text into a more formal tone."
- "Make the given content sound more professional."

#### gec (1,500 examples)
*Example Prompts:*
- "Refine the following content, focusing on fixing grammatical errors."
- "Correct any grammar mistakes in the text below."

#### neutralize (800 examples)
*Example Prompts:*
- "Rewrite the following content in a neutral tone, removing any biases or strong emotions."
- "Ensure the text below is objective and does not show any personal opinions."

#### text_correction (1,400 examples)
*Example Prompts:*
- "Please read the provided document to understand the author's intention. Focus on the fixes required in the document, such as mistranscriptions, punctuation, spelling mistakes, and consistency. Provide a fixed version of the document."

#### simplification (900 examples)
*Example Prompts:*
- "Use simpler wording to convey the message in the content below."
- "Make the following text easier for a child to understand."

---

### 2. Content Generation

Tasks in this category involve creating or expanding content.

#### kg_writer (800 examples)
*Example Prompts:*
- "Using the provided knowledge graph, write an article about the topics and entities in the graph, incorporating the linked ideas. Use idea tags while writing to help focus."
- "Construct a story based on the information in the knowledge graph."

#### summary (1,600 examples)
*Example Prompts:*
- "Generate an extensive summary of the given document."
- "Please read the provided document to understand the context and content. Use this understanding to generate a summary. Separate the article into chunks, and sequentially create a summary for each chunk. Give me a final summary in the end."

#### paraphrase (1,100 examples)
*Example Prompts:*
- "Rephrase the following sentence while retaining its original meaning."
- "Can you provide an alternative wording for the paragraph below?"

---

### 3. Content Analysis

Tasks in this category evaluate, grade, or filter content.

#### grading (400 examples)
*Example Prompts:*
- "Based on the provided document, please rate the usefulness as training data on a scale from 0-5."

#### sponsorblock (5,200 examples)
*Example Prompts:*
- "Read the document and extract any sentences or phrases that contain explicit mentions of sponsorship, promotional partnerships, or any form of paid content."

---

### 4. Information Structuring

Tasks in this category involve the structured representation or extraction of information.

#### kg (3,600 examples)
*Example Prompts:*
- "Create a Knowledge Graph from the document provided."
- "Extract key concepts and relationships from the conversation to form a knowledge graph."

---

### 5. General Interaction

Tasks in this category are designed for general questions and interactions.

#### general (1,600 examples)
*Example Prompts:*
- "What is the capital of France?"
- "Explain particle physics to a 5 years old."



## Limitations
- Adhere to the prompt structure for best results.
- When providing contextual details, clarity is essential for Inkbot to derive accurate and meaningful responses.
- The overriding memory from user_context property generally only works for the next prompt or two, after which the model reverts to its original behavior.
- On complex tasks, like creating a coherent story based on a set of facts from context, there's a potential for a repeating text loop as context fills.
- Sometimes the model doesn't know when to end a knowledge graph, which can result in adding nodes and edges until it runs out of context.

## Additional Notes
- Use rope-freq-scale=0.5 or compress_pos_emb=2 for 8k ctx
- The 'date', 'task', and 'system' are crucial metadata components that need to be provided outside the core dialogue.
- Use the 'user_context' when you want to offer supplementary context that guides Inkbot's response. You can interleave it in the chat log as necessary. It comes after the users instruction.
- The specific tag format, such as `<#word#>`, is used to because there are filters in a lot of APIs for <|word|> and this makes interactions easier.


---
license: llama2
---