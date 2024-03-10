---
license: cc-by-4.0
tags:
- merge
- moe
---
Everyone-Coder-4x7b-Base


![image/jpeg](https://cdn-uploads.huggingface.co/production/uploads/642cc1c253e76b4c2286c58e/ECrHQnZnv8UM9GUCQtlWW.jpeg)

EveryoneLLM series of models are a new Mixtral type model created using experts that were finetuned by the community, for the community. This is the first model to release in the series and it is a coding specific model. EveryoneLLM, which will be a more generalized model, will be released in the near future after more work is done to fine tune the process of merging Mistral models into a larger Mixtral models with greater success. 

The goal of the EveryoneLLM series of models is to be a replacement or an alternative to Mixtral-8x7b that is more suitable for general and specific use, as well as easier to fine tune. Since Mistralai is being secretive about the "secret sause" that makes Mixtral-Instruct such an effective fine tune of the Mixtral-base model, I've decided its time for the community to directly compete with Mistralai on our own.

The models that were used in this merger were as follow:

- https://huggingface.co/fblgit/UNA-TheBeagle-7b-v1

- https://huggingface.co/LucciAI/openchat-3.5-0106-function-calling

- https://huggingface.co/WizardLM/WizardMath-7B-V1.1

- https://huggingface.co/cognitivecomputations/dolphin-2.6-mistral-7b-dpo-laser


Thank you to the creators of the above ai models, they have full credit for the EveryoneLLM series of models. Without their hard work we wouldnt be able to achieve the great success we have in the open source community. 💗

You can find the write up for this model here:

https://docs.google.com/document/d/1_vOftBnrk9NRk5h10UqrfJ5CDih9KBKL61yvrZtVWPE/edit?usp=sharing

Config for the merger can be found bellow:


```yaml
base_model: mistralai_Mistral-7B-v0.1
gate_mode: hidden
dtype: float16
experts:
  - source_model: cognitivecomputations_dolphin-2.6-mistral-7b-dpo-laser
    positive_prompts:
    - "Help me debug this code."
    - "Rewrite this function in Python."
    - "Optimize this C# script."
    - "Implement this feature using JavaScript."
    - "Convert this HTML structure into a more efficient design."
    - "Assist me with writing a program that"
  - source_model: fblgit_UNA-TheBeagle-7b-v1
    positive_prompts:
    - "How do you"
    - "Explain the concept of"
    - "Give an overview of"
    - "Compare and contrast between"
    - "Provide information about"
    - "Help me understand"
    - "Summarize"
    - "Make a recommendation on"
    - "Answer this question"
  - source_model: LucciAI_openchat-3.5-0106-function-calling
    positive_prompts:
    - "Write a program to solve this problem"
    - "Modify this function to improve its performance"
    - "Refactor this code to enhance readability"
    - "Create a custom function for this specific use case"
    - "Optimize this algorithm to reduce computational complexity"
    - "Implement this feature by extending existing codebase"
    - "Integrate this API call into the application"
    - "Help me troubleshoot and fix this bug"
    - "Review and test this code snippet before deployment"
    - "Analyze this error log to identify potential issues"
    - "Generate a set of unit tests for this module"
    - "Evaluate different approaches to solving this problem"
    - "Do a web search for"
    - "Use the plugin to"
  - source_model: WizardLM_WizardMath-7B-V1.1
    positive_prompts:
    - "add these numbers"
    - "whats 2+2"
    - "subtraction"
    - "division"
    - "multiplication"
    - "addition"
    - "I need help with a math problem"
    - "Solve for x"
    - "Add these two numbers together: 4 + 3 = 7"
    - "Multiply 5 by 6: 5 * 6 = 30"
    - "Divide 8 by 2: 8 / 2 = 4"
    - "Find the remainder when 9 is divided by 3: 9 % 3 = 0"
    - "Calculate the square root of 16: sqrt(16) = 4"
    - "Simplify the expression (a+b)/(c-d): (a+b)/(c-d)"
    - "Factor out the common factor of 2 from 4x + 6y: 2(2x + 3y)"
    - "Solve for x in the equation 3x - 7 = 2x + 5: x = 12"
    - "Graph the line y = 2x + 3"
    - "Approximate pi to three decimal places: 3.142"
    - "Find the derivative of f(x) = sin(x): f'(x) = cos(x)"
    - "Integrate g(x) = x^2 over the interval [0, 1]: g(1) - g(0) = 1/3"
    - "Calculate the determinant of the matrix A = [[2, 3], [4, 5]]: det(A) = 2*5 - 3*4 = -2"
    - "Solve the system of equations Ax = b: x = [-5, 10]"
    - "Calculate the sum of the first n natural numbers using the formula Sn = n*(n+1)/2: sum(n=1 to 5) = 15"
```



