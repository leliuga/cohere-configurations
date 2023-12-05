# Model Card for Inkbot

## Model Details
Inkbot is a conversational AI model designed to interpret and respond to structured prompts with or without contextual information. Built on the latest advancements in natural language processing (NLP) and understanding (NLU), Inkbot provides users with accurate and meaningful interactions, addressing a wide range of queries and topics. Its unique feature lies in the structured prompt system, allowing users to engage in dynamic dialogues that can evolve based on context, metadata, and user input.

### Performance
- The model excels in RAG type queries, answering from context, and overriding memory when available.
- It can handle very large contexts, but may sometimes enter a repeating text loop, especially during complex tasks.
- The model is intended to be more functional and less chatty, avoiding the waste of tokens on superfluous language.

## How to Use
Inkbot uses a structured prompt template system.


### Prompt Template Structure

#### With Context:
```
<#meta#>
- Date: [DATE]
- Task: [TASK TYPE]
<#system#>
[SYSTEM PROMPT/DESCRIPTION OF THE INTERACTION]
<#chat#>
<#user#>
[USER QUESTION]
<#bot#>
[BOT RESPONSE]
<#user#>
[USER FOLLOW-UP QUESTION]
<#user_context#>
[ADDITIONAL CONTEXT FOR THE BOT]
<#bot#>
[BOT RESPONSE BASED ON CONTEXT]
```

#### Without Context:
```
<#meta#>
- Date: [DATE]
- Task: [TASK TYPE]
<#system#>
[SYSTEM PROMPT/DESCRIPTION OF THE INTERACTION]
<#chat#>
<#user#>
[USER QUESTION]
<#bot#>
[BOT RESPONSE]
<#user#>
[USER FOLLOW-UP QUESTION]
<#bot#>
[BOT RESPONSE]
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

Inkbot has been trained for a variety of tasks. Below are some of the key task options you can utilize, but feel free to try others to test the effect:

1. **general**: This is the default task and is designed for a broad range of general questions and interactions.

    - Usage: Suitable for most day-to-day interactions and queries.

2. **knowledge_graph**: This task involves extracting, understanding, and representing information in a structured way.

    - Usage: When you want to extract relationships between entities or desire structured representations of data.

3. **question_answer**: Explicitly trained for answering questions in a straightforward manner.

    - Usage: Best used when you have direct questions and expect concise answers.

4. **reasoning**: Allows Inkbot to showcase its logical and deductive reasoning capabilities.

    - Usage: Ideal for puzzles, riddles, or scenarios where logical analysis is required.

5. **translation**: Use this for language translation tasks.

    - Usage: Provide a sentence or paragraph in one language, and specify the desired target language for translation.

6. **summarization**: Trained for condensing large texts into shorter, coherent summaries.

    - Usage: When you have a lengthy text or article that you want to be summarized to its key points.

7. **creative_writing**: Engage Inkbot in composing stories, poetry, and other creative content.

    - Usage: For tasks that require imaginative and original content generation.


## Limitations
- Adhere to the prompt structure for best results.
- When providing contextual details, clarity is essential for Inkbot to derive accurate and meaningful responses.
- The overriding memory from user_context property generally only works for the next prompt or two, after which the model reverts to its original behavior.
- On complex tasks, like creating a coherent story based on a set of facts from context, there's a potential for a repeating text loop as context fills.
- Sometimes the model doesn't know when to end a knowledge graph, which can result in adding nodes and edges until it runs out of context.

## Additional Notes
- The 'date', 'task', and 'system' are crucial metadata components that need to be provided outside the core dialogue.
- Use the 'user_context' when you want to offer supplementary context that guides Inkbot's response. You can interleave it in the chat log as necessary.
- The specific tag format, such as `<#word#>`, is used to because there are filters in a lot of APIs for <|word|> and this makes interactions easier.


---
license: llama2
---
