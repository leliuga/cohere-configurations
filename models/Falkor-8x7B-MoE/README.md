---
license: apache-2.0
---

# Falkor 7B MoE 8x7B Experts

<img src="falkor.png" width="300">

Model merge between Chupacabra, openchat, and dragon-mistral-7b-v0.

- ---> [Theme Song](https://www.youtube.com/watch?v=lHytjEj7B9g) <---

# Original Model Card for dragon-mistral-7b-v0

<!-- Provide a quick summary of what the model is/does. -->

dragon-mistral-7b-v0 part of the dRAGon ("Delivering RAG On ...") model series, RAG-instruct trained on top of a Mistral-7B base model.

DRAGON models have been fine-tuned with the specific objective of fact-based question-answering over complex business and legal documents with an emphasis on reducing hallucinations and providing short, clear answers for workflow automation.

### Benchmark Tests  

Evaluated against the benchmark test:   [RAG-Instruct-Benchmark-Tester](https://www.huggingface.co/datasets/llmware/rag_instruct_benchmark_tester)  
Average of 2 Test Runs with 1 point for correct answer, 0.5 point for partial correct or blank / NF, 0.0 points for incorrect, and -1 points for hallucinations.  

--**Accuracy Score**:  **96.50** correct out of 100  
--Not Found Classification:  92.50%  
--Boolean:  97.50%  
--Math/Logic:  81.25%  
--Complex Questions (1-5):  4 (Medium-High - table-reading, multiple-choice, causal)  
--Summarization Quality (1-5):  4 (Coherent, extractive)  
--Hallucinations:  No hallucinations observed in test runs.  

For test run results (and good indicator of target use cases), please see the files ("core_rag_test" and "answer_sheet" in this repo).

### Model Description

<!-- Provide a longer summary of what this model is. -->

- **Developed by:** llmware
- **Model type:** Mistral-7B
- **Language(s) (NLP):** English
- **License:** Apache 2.0
- **Finetuned from model:** Mistral-7B-Base

### Direct Use

<!-- This section is for the model use without fine-tuning or plugging into a larger ecosystem/app. -->

DRAGON is designed for enterprise automation use cases, especially in knowledge-intensive industries, such as financial services,
legal and regulatory industries with complex information sources.  

DRAGON models have been trained for common RAG scenarios, specifically:   question-answering, key-value extraction, and basic summarization as the core instruction types
without the need for a lot of complex instruction verbiage - provide a text passage context, ask questions, and get clear fact-based responses.


## Bias, Risks, and Limitations

<!-- This section is meant to convey both technical and sociotechnical limitations. -->

Any model can provide inaccurate or incomplete information, and should be used in conjunction with appropriate safeguards and fact-checking mechanisms.


## How to Get Started with the Model

The fastest way to get started with dRAGon is through direct import in transformers:

    from transformers import AutoTokenizer, AutoModelForCausalLM  
    tokenizer = AutoTokenizer.from_pretrained("dragon-mistral-7b-v0")  
    model = AutoModelForCausalLM.from_pretrained("dragon-mistral-7b-v0")  

Please refer to the generation_test .py files in the Files repository, which includes 200 samples and script to test the model.  The **generation_test_llmware_script.py** includes built-in llmware capabilities for fact-checking, as well as easy integration with document parsing and actual retrieval to swap out the test set for RAG workflow consisting of business documents.  

The dRAGon model was fine-tuned with a simple "\<human> and \<bot> wrapper", so to get the best results, wrap inference entries as:

    full_prompt = "<human>: " + my_prompt + "\n" + "<bot>:"

The BLING model was fine-tuned with closed-context samples, which assume generally that the prompt consists of two sub-parts:

1.  Text Passage Context, and
2.  Specific question or instruction based on the text passage

To get the best results, package "my_prompt" as follows:

    my_prompt = {{text_passage}} + "\n" + {{question/instruction}}


If you are using a HuggingFace generation script:

    # prepare prompt packaging used in fine-tuning process
    new_prompt = "<human>: " + entries["context"] + "\n" + entries["query"] + "\n" + "<bot>:"

    inputs = tokenizer(new_prompt, return_tensors="pt")  
    start_of_output = len(inputs.input_ids[0])

    #   temperature: set at 0.3 for consistency of output
    #   max_new_tokens:  set at 100 - may prematurely stop a few of the summaries

    outputs = model.generate(
            inputs.input_ids.to(device),
            eos_token_id=tokenizer.eos_token_id,
            pad_token_id=tokenizer.eos_token_id,
            do_sample=True,
            temperature=0.3,
            max_new_tokens=100,
            )

    output_only = tokenizer.decode(outputs[0][start_of_output:],skip_special_tokens=True)  


## Model Card Contact

Darren Oberst & llmware team


