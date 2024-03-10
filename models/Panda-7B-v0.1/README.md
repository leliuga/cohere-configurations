---
license: apache-2.0
base_model: mistralai/Mistral-7B-Instruct-v0.2
datasets:
- NeuralNovel/Creative-Logic-v1
- NeuralNovel/Neural-Story-v1
library_name: transformers
inference: false
---
![Neural-Story](https://i.ibb.co/TYvZhws/Panda7b.png)

# NeuralNovel/Panda-7B-v0.1
The **Panda-7B-v0.1** model by NeuralNovel.

This fine-tune has been designed to provide detailed, creative and logical responses in the context of diverse narratives. Optimised for creative writing, roleplay and logical problem solving.

Full-parameter fine-tune (FFT) of Mistral-7B-Instruct-v0.2. Apache-2.0 license, suitable for commercial or non-commercial use.

<a href='https://ko-fi.com/S6S2UH2TC' target='_blank'><img height='38' style='border:0px;height:36px;' src='https://storage.ko-fi.com/cdn/kofi1.png?v=3' border='0' alt='Buy Me a Coffee at ko-fi.com' /></a>
<a href='https://discord.gg/KFS229xD' target='_blank'><img width='140' height='500' style='border:0px;height:36px;' src='https://i.ibb.co/tqwznYM/Discord-button.png' border='0' alt='Join Our Discord!' /></a>

### Data-set
The model on finetuned using the Panda-v1 dataset.

### Summary

Fine-tuned with the intention to generate instructive and narrative text, with a specific focus on combining the elements of versatility, character engagement and nuanced writing capability. 

#### Out-of-Scope Use

The model may not perform well in scenarios unrelated to instructive and narrative text generation. Misuse or applications outside its designed scope may result in suboptimal outcomes.

### Bias, Risks, and Limitations

The model may exhibit biases or limitations inherent in the training data. It is essential to consider these factors when deploying the model to avoid unintended consequences.

Users are advised to exercise caution, as there might be some inherent genre or writing bias.

### Hardware and Training



```

  n_epochs = 3,
  n_checkpoints = 3,
  batch_size = 12,
  learning_rate = 1e-5,



```

*Sincere appreciation to Techmind for their generous sponsorship.*
