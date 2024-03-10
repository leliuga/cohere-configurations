---
license: apache-2.0
datasets:
- Intel/orca_dpo_pairs
- mlabonne/chatml_dpo_pairs
language:
- en
library_name: transformers
---
More information about previous [Neuronovo/neuronovo-9B-v0.2](https://huggingface.co/Neuronovo/neuronovo-9B-v0.2) version available here: 🔗[Don't stop DPOptimizing!](https://www.linkedin.com/pulse/dont-stop-dpoptimizing-jan-koco%2525C5%252584-mq4qf)

Author: Jan Kocoń &nbsp;&nbsp;&nbsp; 🔗[LinkedIn](https://www.linkedin.com/in/jankocon/) &nbsp;&nbsp;&nbsp; 🔗[Google Scholar](https://scholar.google.com/citations?user=pmQHb5IAAAAJ&hl=en&oi=ao) &nbsp;&nbsp;&nbsp; 🔗[ResearchGate](https://www.researchgate.net/profile/Jan-Kocon-2)

Changes concerning [Neuronovo/neuronovo-9B-v0.2](https://huggingface.co/Neuronovo/neuronovo-9B-v0.2):

1. **Training Dataset**: In addition to the [Intel/orca_dpo_pairs](Intel/orca_dpo_pairs) dataset, this version incorporates a [mlabonne/chatml_dpo_pairs](https://huggingface.co/datasets/mlabonne/chatml_dpo_pairs). The combined datasets enhance the model's capabilities in dialogues and interactive scenarios, further specializing it in natural language understanding and response generation.

2. **Tokenizer and Formatting**: The tokenizer now originates directly from the [Neuronovo/neuronovo-9B-v0.2](https://huggingface.co/Neuronovo/neuronovo-9B-v0.2) model.

3. **Training Configuration**: The training approach has shifted from using `max_steps=200` to `num_train_epochs=1`. This represents a change in the training strategy, focusing on epoch-based training rather than a fixed number of steps.

4. **Learning Rate**: The learning rate has been reduced to a smaller value of `5e-6`. This finer learning rate allows for more precise adjustments during the training process, potentially leading to better model performance.