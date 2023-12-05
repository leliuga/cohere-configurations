---
datasets:
- Norquinal/claude_multiround_chat_1k
license: cc-by-nc-4.0
---
This is the [Mistral-7B-v0.1](https://huggingface.co/mistralai/Mistral-7B-v0.1) model fine-tuned using QLoRA (4-bit precision) on my [claude_multiround_chat_1k](https://huggingface.co/datasets/Norquinal/claude_multiround_chat_1k) dataset, which is a randomized subset of ~1000 samples from my [claude_multiround_chat_30k](https://huggingface.co/datasets/Norquinal/claude_multiround_chat_30k) dataset.

## Prompt Format
The model was finetuned with the Vicuna 1.1 prompt format:
```
A chat between a curious user and an artificial intelligence assistant. The assistant gives helpful, detailed, and polite answers to the user's questions.

USER: {instruction}
ASSISTANT:
```

## Use in Text Generation Web UI
Install the bleeding-edge version of `transformers` from source:
```
pip install git+https://github.com/huggingface/transformers
```
Or, alternatively, change `model_type` in `config.json` from `mistral` to `llama`.

## Example Usage
>USER: First, give a brief overview of Chinese cuisine, including its regional diversity. next, describe the importance of the five flavors in Chinese cooking. Then, mention three popular Chinese dishes and their main ingredients. After that, provide a popular Chinese cooking technique and an example of a dish prepared using that technique. Finally, discuss the role of tea in traditional Chinese culture.

>ASSISTANT: Chinese cuisine is incredibly diverse due to its large geographic area with different climates and resources available. Regional specialties include Sichuan spicy hot pot, Cantonese dim sum and seafood dishes, Beijing roast duck, and Hunan stir-fries.   
>
>The five basic tastes that guide Chinese cookery are sweetness, sourness, bitterness, saltiness, and pungency (spiciness). These contrasting flavors are used together to create balance in each dish.    
>
>Some classic Chinese dishes include dumplings filled with meat or vegetables, spring rolls stuffed with vegetables and meat wrapped in thin pastry, and noodles served in soups or stir fried with sauces and vegetables.    
>
>A common Chinese cooking technique is braising, where food is slowly cooked in a small amount of liquid until tender and flavorful. An example would be beef brisket braised in soy sauce, ginger and star anise for hours until falling apart.  
>
>Tea plays a central role in Chinese culture as it represents hospitality, relaxation and socializing. Green teas are most commonly consumed but black, oolong and white varieties also exist. Tea drinking rituals involve preparing and serving tea in specific ways depending on context and occasion. 
>
>In summary, Chinese cuisine offers a rich culinary tradition influenced by regional differences, an emphasis on balancing the five basic flavors, and a focus on simple yet flavorful preparations using fresh seasonal produce and high quality meats. Tea acts as both a refreshment and symbol of cultural refinement within Chinese society. I hope this overview provides an interesting starting point for further exploration into the fascinating world of Chinese gastronomy!