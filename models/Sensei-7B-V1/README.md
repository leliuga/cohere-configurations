---
{}
---

# Sensei-7B-V1 Model Card

Sensei-7B-V1 is a Large Language Model (LLM) fine-tuned from OpenPipe's mistral-ft-optimized-1218, which is based on Mistral-7B. Sensei-7B-V1 was was fine-tuned with a fully synthetic dataset to specialize at performing retrieval-augmented generation (RAG) over detailed web search results. This model strives to specialize in using search, such as [AgentSearch](https://huggingface.co/datasets/SciPhi/AgentSearch-V1), to generate accurate and well-cited summaries from a range of search results, providing more accurate answers to user queries. Please refer to the [docs here](https://agent-search.readthedocs.io/en/latest/) for more information on how to run Sensei end-to-end.

Currently, Sensei is available via hosted api at https://www.sciphi.ai. You can try a demonstration [here](https://search.sciphi.ai/). 

## Model Architecture

Base Model: mistral-ft-optimized-1218

**Architecture Features:**
- Transformer-based model
- Grouped-Query Attention
- Sliding-Window Attention
- Byte-fallback BPE tokenizer


## Using the Model

It is recommended to use a single search query. The model will return an answer using search results as context. 

Using the AgentSearch package an example is shown below.
```
export SCIPHI_API_KEY=MY_SCIPHI_API_KEY
# Use `Sensei` for LLM RAG w/ AgentSearch
python -m agent_search.scripts.run_rag run --query="What is Fermat's last theorem?"
```

Alternatively, you may provide your own search context directly to the model by adhereing to the following format:

```
### Instruction: 
Your task is to perform retrieval augmented generation (RAG) over the given query and search results. Return your answer in a json format that includes a summary of the search results and a list of related queries. 

Query:
{prompt}
\n\n
Search Results:
{context}
\n\n
Query:
{prompt}

### Response:
{"summary":
```

__Note__: The inclusion of the text '{"summary":' following the Response footer is intentional. This ensures that the model responds with the proper json format, failure to include this leading prefix can cause small deviaitons. Combining the output with the leading string '{"summary":' results in a properly formatted JSON with keys 'summary' and 'other_queries'.

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)

## References

1. OpenPipe AI. (2023). Model Card for mistral-ft-optimized-1218. The mistral-ft-1218 Large Language Model (LLM) is a pretrained generative text model with 7 billion parameters optimized for  downstream fine-tuning on a variety of tasks. For full details, please refer to the release blog post. Model Architecture: Transformer with Grouped-Query Attention, Sliding-Window Attention, and Byte-fallback BPE tokenizer. [Link](https://huggingface.co/OpenPipe/mistral-ft-optimized-1218)