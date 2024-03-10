---
license: apache-2.0  
inference: false  
---

# Model Card for Model ID

<!-- Provide a quick summary of what the model is/does. -->

bling-stable-lm-3b-4e1t-0.1 part of the BLING ("Best Little Instruction-following No-GPU-required") model series, RAG-instruct trained on top of a StabilityAI stablelm-3b-4e1t base model.

BLING models are fine-tuned with distilled high-quality custom instruct datasets, targeted at a specific subset of instruct tasks with 
the objective of providing a high-quality Instruct model that is 'inference-ready' on a CPU laptop even 
without using any advanced quantization optimizations.


### Benchmark Tests  

Evaluated against the benchmark test:   [RAG-Instruct-Benchmark-Tester](https://www.huggingface.co/datasets/llmware/rag_instruct_benchmark_tester)  
Average of 2 Test Runs with 1 point for correct answer, 0.5 point for partial correct or blank / NF, 0.0 points for incorrect, and -1 points for hallucinations.  

--**Accuracy Score**:  **94.0** correct out of 100  
--Not Found Classification:  67.5%  
--Boolean:  77.5%  
--Math/Logic:  29%  
--Complex Questions (1-5):  3 (Low)  
--Summarization Quality (1-5):  3 (Coherent, extractive)  
--Hallucinations:  No hallucinations observed in test runs.  

For test run results (and good indicator of target use cases), please see the files ("core_rag_test" and "answer_sheet" in this repo).

### Model Description

<!-- Provide a longer summary of what this model is. -->

- **Developed by:** llmware
- **Model type:** Instruct-trained decoder
- **Language(s) (NLP):** English
- **License:** [CC BY-SA-4.0](https://creativecommons.org/licenses/by-sa/4.0/)
- **Finetuned from model:** stabilityai/stablelm-3b-4e1t

  
## Uses

<!-- Address questions around how the model is intended to be used, including the foreseeable users of the model and those affected by the model. -->

The intended use of BLING models is two-fold:

1.  Provide high-quality Instruct models that can run on a laptop for local testing.  We have found it extremely useful when building a
   proof-of-concept, or working with sensitive enterprise data that must be closely guarded, especially in RAG use cases.

2.  Push the state of the art for smaller Instruct-following models in the sub-7B parameter range, especially 1B-3B, as single-purpose
    automation tools for specific tasks through targeted fine-tuning datasets and focused "instruction" tasks.


### Direct Use

<!-- This section is for the model use without fine-tuning or plugging into a larger ecosystem/app. -->

BLING is designed for enterprise automation use cases, especially in knowledge-intensive industries, such as financial services,
legal and regulatory industries with complex information sources.  Rather than try to be "all things to all people," BLING models try to focus on a narrower set of Instructions more suitable to a ~1-3B parameter GPT model.

BLING is ideal for rapid prototyping, testing, and the ability to perform an end-to-end workflow locally on a laptop without
having to send sensitive information over an Internet-based API.

The first BLING models have been trained for common RAG scenarios, specifically:   question-answering, key-value extraction, and basic summarization as the core instruction types
without the need for a lot of complex instruction verbiage - provide a text passage context, ask questions, and get clear fact-based responses.


## Bias, Risks, and Limitations

<!-- This section is meant to convey both technical and sociotechnical limitations. -->

Any model can provide inaccurate or incomplete information, and should be used in conjunction with appropriate safeguards and fact-checking mechanisms.


## How to Get Started with the Model

The fastest way to get started with BLING is through direct import in transformers:

    from transformers import AutoTokenizer, AutoModelForCausalLM  
    tokenizer = AutoTokenizer.from_pretrained("llmware/bling-stable-lm-3b-4e1t-0.1")  
    model = AutoModelForCausalLM.from_pretrained("llmware/bling-stable-lm-3b-4e1t-0.1")  

Please refer to the generation_test .py files in the Files repository, which includes 200 samples and script to test the model.  The **generation_test_llmware_script.py** includes built-in llmware capabilities for fact-checking, as well as easy integration with document parsing and actual retrieval to swap out the test set for RAG workflow consisting of business documents.  

The BLING model was fine-tuned with a simple "\<human> and \<bot> wrapper", so to get the best results, wrap inference entries as:

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


## Citations

This model has been fine-tuned on the base StableLM-3B-4E1T model from StabilityAI. For more information about this base model, please see the citation below:

@misc{StableLM-3B-4E1T,
      url={[https://huggingface.co/stabilityai/stablelm-3b-4e1t](https://huggingface.co/stabilityai/stablelm-3b-4e1t)},
      title={StableLM 3B 4E1T},
      author={Tow, Jonathan and Bellagente, Marco and Mahan, Dakota and Riquelme, Carlos}
}


## Model Card Contact

Darren Oberst & llmware team
