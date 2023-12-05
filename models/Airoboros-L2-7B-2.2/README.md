---
license: llama2
datasets:
- jondurbin/airoboros-2.2
---

### Overview

Another experimental model, using mostly sythetic data generated by [airoboros](https://github.com/jondurbin/airoboros)

Highlights:

- The prompt format has changed!  It is now newlines instead of spaces between system/USER/ASSISTANT (see prompt info below).
- "Clean" version of airoboros-2.2 dataset -- this model __does not__ contain the de-alignment data.
  - For an uncensored version, use spicyboros variants.
- I re-generated all of the outputs in the dataset that had "Once upon a time" so they'd be less cliche - no guarantees that won't still happen, but in theory it may happen less.
- More multiple choice, better awareness, some alignment for normal use case but system-prompt overridable etc.

Breakdown of the training data:

| Count | Category                   |
|-------|----------------------------|
| 36    | experience                 |
| 60    | quiz                       |
| 63    | card                       |
| 76    | greeting                   |
| 100   | detailed\_writing          |
| 200   | song                       |
| 204   | editor                     |
| 207   | counterfactual\_contextual |
| 268   | cot                        |
| 339   | theory\_of\_mind           |
| 416   | awareness                  |
| 439   | stylized\_response         |
| 457   | misconception              |
| 500   | summarization              |
| 620   | riddle                     |
| 719   | agent                      |
| 800   | plan                       |
| 873   | gtkm                       |
| 963   | rp                         |
| 1000  | wordgame                   |
| 1279  | multiple\_choice           |
| 1519  | joke                       |
| 1758  | writing                    |
| 2152  | contextual                 |
| 2183  | trivia                     |
| 2364  | roleplay                   |
| 4699  | general                    |
| 5775  | coding                     |
| 11366 | orca                       |


In other words, it's a fairly general purpose model, but focuses fairly heavily on instruction response pairs rather than casual chat/roleplay.

Huge thank you to the folks over at [a16z](https://a16z.com/) for sponsoring the costs associated with building models and associated tools!


### Prompt format

The prompt format:

```
A chat.
USER: {prompt}
ASSISTANT:
```

The default system prompt ("A chat.") was used for most of the prompts, however it also included a wide sampling of responses with other prompts, particularly in "stylized\_response", "rp", "gtkm", etc.

Here's another example:
```
A chat between Bob (aka USER) and Tom (aka ASSISTANT). Tom is an extremely intelligent 18th century bookkeeper, who speaks loquaciously.
USER: {prompt}
ASSISTANT: 
```

And chat scenario that wouldn't require USER/ASSISTANT (but should use stopping criteria to prevent the model from speaking on your behalf).
```
A chat between old friends: Timmy and Tommy.
{description of characters}

{setting for the chat}
Timmy: *takes a big sip from his coffee* "Ah, sweet, delicious, magical coffee."
Tommy: 
```

__*I strongly suggest adding stopping criteria/early inference stopping on "USER:", and/or whatever names you specify in the system prompt.*__

### Fine tuning info

https://gist.github.com/jondurbin/7aabf2d8e1ded7b29897a3dd95d81e01

### Helpful usage tips

*The prompts shown here are are just the text that would be included after USER: and before ASSISTANT: in the full prompt format above, the system prompt and USER:/ASSISTANT: have been omited for readability.*

#### Context obedient question answering

By obedient, I mean the model was trained to ignore what it thinks it knows, and uses the context to answer the question.  The model was also tuned to limit the values to the provided context as much as possible to reduce hallucinations.

The format for a closed-context prompt is as follows:
```
BEGININPUT
BEGINCONTEXT
[key0: value0]
[key1: value1]
... other metdata ...
ENDCONTEXT
[insert your text blocks here]
ENDINPUT
[add as many other blocks, in the exact same format]
BEGININSTRUCTION
[insert your instruction(s).  The model was tuned with single questions, paragraph format, lists, etc.]
ENDINSTRUCTION
```

It's also helpful to add "Don't make up answers if you don't know." to your instruction block to make sure if the context is completely unrelated it doesn't make something up.

*The __only__ prompts that need this closed context formating are closed-context instructions.  Normal questions/instructions do not!*

I know it's a bit verbose and annoying, but after much trial and error, using these explicit delimiters helps the model understand where to find the responses and how to associate specific sources with it.
- `BEGININPUT` - denotes a new input block
- `BEGINCONTEXT` - denotes the block of context (metadata key/value pairs) to associate with the current input block
- `ENDCONTEXT` - denotes the end of the metadata block for the current input
- [text] - Insert whatever text you want for the input block, as many paragraphs as can fit in the context.
- `ENDINPUT` - denotes the end of the current input block
- [repeat as many input blocks in this format as you want]
- `BEGININSTRUCTION` - denotes the start of the list (or one) instruction(s) to respond to for all of the input blocks above.
- [instruction(s)]
- `ENDINSTRUCTION` - denotes the end of instruction set

It sometimes works without `ENDINSTRUCTION`, but by explicitly including that in the prompt, the model better understands that all of the instructions in the block should be responded to.

Here's a trivial, but important example to prove the point:
```
BEGININPUT
BEGINCONTEXT
date: 2021-01-01
url: https://web.site/123
ENDCONTEXT
In a shocking turn of events, blueberries are now green, but will be sticking with the same name.
ENDINPUT
BEGININSTRUCTION
What color are bluberries?  Source?
ENDINSTRUCTION
```

And the response:
```
Blueberries are now green.
Source:
date: 2021-01-01
url: https://web.site/123
```

#### Coding

You can ask for fairly complex coding instructions with multiple criteria, e.g.:

```
Create a python application with the following requirements:
- Asyncio FastAPI webserver
- ping endpoint that returns the current date in JSON format
- file upload endpoint, which calculates the file's sha256 checksum, and checks postgres to deduplicate
```

Or inline criteria:

```
Write a multi-threaded TCP server in C that accepts a "GET [key]" input and "SET [key] [value]" input, and uses a binary tree to get and store the input values.
```

You can also optionally add a single space and "PLAINFORMAT" at the end of your prompt to avoid backticks, explanations, etc. and just print the code, e.g.:

```
Write a websocket application in node.js. PLAINFORMAT
```

#### Agent/function calling

The dataset includes many examples of function/args generation based on input criteria.  This is somewhat similar to the OpenAI function calling, but the output is either JSON or YAML.

Example prompt:
```
As an AI assistant, please select the most suitable function and parameters from the list of available functions below, based on the user's input. Provide your response in JSON format.

Input: I want to know how many times 'Python' is mentioned in my text file.

Available functions:
file_analytics:
  description: This tool performs various operations on a text file.
  params:
    action: The operation we want to perform on the data, such as "count_occurrences", "find_line", etc.
    filters:
      keyword: The word or phrase we want to search for.
```

Response:
```json
{
  "function": "file_analytics",
  "params": {
    "action": "count_occurrences",
    "filters": {
      "keyword": "Python"
    }
  }
}
```

#### Chain-of-thought

You can ask for several possible responses to a given problem, with a ranking and final answer selection.

Example prompt:

```
A group of people decide to play a round-robin tournament where every player plays every other player exactly once. If a total of 45 games were played, how many players participated in the tournament? Offer several potential answers, rank them according to mathematical logic, and pick the most feasible one.
```

Example response:
```
Solution 1:
In a round-robin tournament, each player plays every other player exactly once. This means that the total number of games played is equal to the combination of all players taken two at a time (since a game involves two players). 

The formula for combinations is nC2 = n(n-1)/2, where n is the total number of players. We know that nC2 equals 45 (the total number of games), so we can set up the equation as follows:

n(n-1)/2 = 45

Solving this equation gives us n=10.

Final answer: There were 10 players in the tournament.

Solution 2:
Another way to approach this problem is by using the concept of triangular numbers. Triangular numbers are a sequence of numbers where each number can be represented as a triangle with dots. The nth triangular number is given by the sum of the first n natural numbers. 

If we consider each game as a dot and each player as a side of the triangle, then the total number of games (45) would correspond to the 9th triangular number because 1+2+3+4+5+6+7+8+9=45. However, since each side of the triangle represents a player, and there's one more player than sides in our model (because the last player has no one left to play against), we need to add one more to get the total number of players.

Final answer: There were 10 players in the tournament.

Solution 3:
We could also use trial and error to solve this problem. Starting from a small number of players, we calculate how many games would be played in a round-robin tournament until we reach 45 games. For example, with 2 players, only 1 game is played; with 3 players, 3 games are played; with 4 players, 6 games are played, and so on. Continuing this process, we find that with 10 players, 45 games are played.

Final answer: There were 10 players in the tournament.

Ranking of solutions from best to worst:
1. Solution 1: This solution uses direct mathematical reasoning based on the properties of a round-robin tournament and solves an equation to find the answer.
2. Solution 2: This solution uses the concept of triangular numbers which requires additional steps and knowledge beyond the immediate problem context, but still arrives at the correct answer.
3. Solution 3: This solution relies on trial and error, which is less efficient and reliable than the other methods, but it eventually gets to the right answer.

Best and final answer: There were 10 players in the tournament.
```

#### reWOO style execution planning

The model now supports execution planning for complex instructions that would require making use of several tools.  The output is just the plan, you must implement a mechanism to parse the output and actually call the functions!

Example prompt:
```
Please construct a systematic plan to generate an optimal response to the user instruction, utilizing a set of provided tools. Each plan will correspond to an evidence value, which will be the output of one of the available functions given an input string
that could be the user's question, one or more prior evidence values, or a combination of both.

Here are the tools available to be called:
DuckDuckGo[input]: Tool that enables the user to search for information using the DuckDuckGo search engine. This tool is helpful in locating a suitable list of sites that may contain the answer to the user's query. It does not directly answer the question but finds a list of sites that might have the answer.
TextScraper[input]: Load one or more websites from the input string containing newline delimited links, where input is one or more links, and produces plain text output containing the content of the links.
HyperlinkExtractor[input]: Extract links from plain text and produces a plain text, newline delimited response of links.
KnowledgeModel[input]: Question answering language model, particularly useful in answering questions based on an input passage of text. The input must be a text question that references an :evidence[n]: variable, e.g. What color is the cat, given :evidence1:?

The output should be in format:
Plan: [first action to take based on input question]
:evidence0: = [function to call with input parameter]
Plan: [next action to take, based on result of :evidence0:]
:evidence1: = [next function to call and input parameter, which may include reference to previous evidence, e.g. "Given context :evidence0"]
...
Answer: [:evidence[n]: containing the final answer.]

Question: Who won the Nobel Prize in Literature in 2020 and what is their most famous work?
```

Response:
```
Plan: Begin by conducting a web search to find out who won the Nobel Prize in Literature in 2020.
:evidence0: = DuckDuckGo[Nobel Prize in Literature 2020 winner]
Plan: Utilize the KnowledgeModel to interpret the search results and find the name of the laureate.
:evidence1: = KnowledgeModel[Who won the Nobel Prize in Literature in 2020, given :evidence0:?]
Plan: Conduct another web search to find the most famous work of the identified laureate.
:evidence2: = DuckDuckGo[Most famous work of :evidence1:]
Plan: Extract the relevant links from the DuckDuckGo search results for a more focused search.
:evidence3: = HyperlinkExtractor[:evidence2:]
Plan: Use the TextScraper tool to extract information from the relevant links.
:evidence4: = TextScraper[:evidence3:]
Plan: Finally, utilize the KnowledgeModel to identify and summarize the most famous work of the laureate from the extracted information.
:evidence5: = KnowledgeModel[What is the most famous work of :evidence1:, given :evidence4:?]
Answer: :evidence5:
```

For this to be useful, you'd have to parse the output plan text, and implement/call each of the functions.  This is just pseudo-code, completely untested off the top of my head, and obviously would requiring full implementation + hardening:

```python
import re
import requests

def inject_context(input_text, **context):
    for ref in set(re.findall(r"(:evidence[0-9]+:)", input_text, re.I)):
        input_text = input_text.replace(ref, context.get(ref, ""))
    return input_text

def duckduckgo(input_text, **context):
    search_string = inject_context(input_text, **context)
    ... search via duck duck go using search_string
    ... return text content

def link_extractor(input_text, **context):
    input_text = inject_context(input_text, **context)
    return "\n".join(list(set(re.findall(r"(https?://[^\s]+?\.?)", input_text, re.I))))

def scrape(input_text, **context):
  input_text = inject_context(input_text, **context)
  text = []
  for link in input_text.splitlines():
    text.append(requests.get(link).text)
  return "\n".join(text)

def infer(input_text, **context)
  prompt = inject_context(input_text, **context)
  ... call model with prompt, return output

def parse_plan(plan):
    method_map = {
      "DuckDuckGo": duckduckgo,
      "HyperlinkExtractor": link_extractor,
      "KnowledgeModel": infer,
      "TextScraper": scrape,
    }
    context = {}
    for line in plan.strip().splitlines():
        if line.startswith("Plan:"):
            print(line)
            continue
        parts = re.match("^(:evidence[0-9]+:)\s*=\s*([^\[]+])(\[.*\])\s$", line, re.I)
        if not parts:
          if line.startswith("Answer: "):
            return context.get(line.split(" ")[-1].strip(), "Answer couldn't be generated...")
          raise RuntimeError("bad format: " + line)
        context[parts.group(1)] = method_map[parts.group(2)](parts.group(3), **context)
```

### Contribute

If you're interested in new functionality, particularly a new "instructor" type to generate a specific type of training data,
take a look at the dataset generation tool repo: https://github.com/jondurbin/airoboros and either make a PR or open an issue with details.

To help me with the OpenAI/compute costs:

- https://bmc.link/jondurbin
- ETH 0xce914eAFC2fe52FdceE59565Dd92c06f776fcb11
- BTC bc1qdwuth4vlg8x37ggntlxu5cjfwgmdy5zaa7pswf

### Licence and usage restrictions

The airoboros 2.2 models are built on top of llama-2/codellama.

The llama-2 base model has a custom Meta license:
- See the [meta-license/LICENSE.txt](meta-license/LICENSE.txt) file attached for the original license provided by Meta.
- See also [meta-license/USE_POLICY.md](meta-license/USE_POLICY.md) and [meta-license/Responsible-Use-Guide.pdf](meta-license/Responsible-Use-Guide.pdf), also provided by Meta.

The fine-tuning data was mostly generated by OpenAI API calls to gpt-4, via [airoboros](https://github.com/jondurbin/airoboros)

The ToS for OpenAI API usage has a clause preventing the output from being used to train a model that __competes__ with OpenAI

- what does *compete* actually mean here?
- these small open source models will not produce output anywhere near the quality of gpt-4, or even gpt-3.5, so I can't imagine this could credibly be considered competing in the first place
- if someone else uses the dataset to do the same, they wouldn't necessarily be violating the ToS because they didn't call the API, so I don't know how that works
- the training data used in essentially all large language models includes a significant amount of copyrighted or otherwise non-permissive licensing in the first place
- other work using the self-instruct method, e.g. the original here: https://github.com/yizhongw/self-instruct released the data and model as apache-2

I am purposingly leaving this license ambiguous (other than the fact you must comply with the Meta original license for llama-2) because I am not a lawyer and refuse to attempt to interpret all of the terms accordingly.

Your best bet is probably to avoid using this commercially due to the OpenAI API usage.

Either way, by using this model, you agree to completely indemnify me.