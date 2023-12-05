---
license: apache-2.0
language:
- de
- en
library_name: transformers
pipeline_tag: text-generation
---


![SauerkrautLM](images/SauerkrautLM.png "SauerkrautLM")
## VAGO solutions SauerkrautLM
Introducing SauerkrautLM-v1 - Your German Language Powerhouse! 

We are thrilled to unveil our **very first release**, **SauerkrautLM-v1**. This remarkable creation marks a significant milestone as it is specifically **tailored for the German-speaking community**. In a landscape where German language models are scarce, we are proud to offer a solution that fills this void. 
What sets SauerkrautLM-v1 apart is its versatility. Whether you are an individual looking to harness its capabilities for personal use or a business seeking to integrate it into your projects, our model is designed to accommodate all. It operates under the Apache 2.0 License, providing you with the freedom to explore its potential in both private and commercial applications. 
Performance is at the heart of SauerkrautLM-v1. We put it to the **test using a customized version of MT-Bench for the German language**, and the results speak volumes. It currently stands as the most robust German Language Model on Hugging Face (based on german mt-bench results), showcasing its exceptional capabilities. Rest assured, this model is here to shine and set new standards. And the best thing is it comes in four different sizes (3B, 7B, 13B, 70B) to address your individual needs. 
Our model's journey began with meticulous training using an **augmented dataset within the QLoRA approach**. This is just the beginning of our model series, promising even more innovative and powerful solutions in the future. 

Join us on this exciting adventure as we redefine the possibilities of language modeling for the German-speaking world. 
SauerkrautLM-v1 is here to empower your language-related endeavors like never before. 

## All Models

| Model | HF    | GPTQ  | GGUF  | AWQ  |
|-------|-------|-------|-------|-------|
| SauerkrautLM-3b-v1   | [Link](https://huggingface.co/VAGOsolutions/SauerkrautLM-3b-v1) | [Link](https://huggingface.co/TheBloke/SauerkrautLM-3B-v1-GPTQ) | [Link](https://huggingface.co/TheBloke/SauerkrautLM-3B-v1-GGUF) | [Link](https://huggingface.co/TheBloke/SauerkrautLM-3B-v1-AWQ) |
| SauerkrautLM-7b-v1   | [Link](https://huggingface.co/VAGOsolutions/SauerkrautLM-7b-v1) | [Link](https://huggingface.co/TheBloke/SauerkrautLM-7B-v1-GPTQ) | [Link](https://huggingface.co/TheBloke/SauerkrautLM-7B-v1-GGUF) | [Link](https://huggingface.co/TheBloke/SauerkrautLM-7B-v1-AWQ) |
| SauerkrautLM-7b-v1-mistral   | [Link](https://huggingface.co/VAGOsolutions/SauerkrautLM-7b-v1-mistral) | [Link](https://huggingface.co/TheBloke/SauerkrautLM-7b-v1-mistral-GPTQ) | [Link](https://huggingface.co/TheBloke/SauerkrautLM-7b-v1-mistral-GGUF) | [Link](https://huggingface.co/TheBloke/SauerkrautLM-7b-v1-mistral-AWQ) |
| SauerkrautLM-13b-v1   | [Link](https://huggingface.co/VAGOsolutions/SauerkrautLM-13b-v1) | [Link](https://huggingface.co/TheBloke/SauerkrautLM-13B-v1-GPTQ) | [Link](https://huggingface.co/TheBloke/SauerkrautLM-13B-v1-GGUF) | [Link](https://huggingface.co/TheBloke/SauerkrautLM-13B-v1-AWQ) |
| SauerkrautLM-70b-v1   | [Link](https://huggingface.co/VAGOsolutions/SauerkrautLM-70b-v1) | [Link](https://huggingface.co/TheBloke/SauerkrautLM-70B-v1-GPTQ) | [Link](https://huggingface.co/TheBloke/SauerkrautLM-70B-v1-GGUF) | [Link](https://huggingface.co/TheBloke/SauerkrautLM-70B-v1-AWQ) |

## Model Details
**SauerkrautLM-3b-v1**

**Training Dataset:**

SauerkrautLM was trained with mix of German data augmentation and translated data. 
We found, that only a simple translation of training data can lead to unnatural German phrasings. 
Data augmentation techniques were used to grant grammatical, syntactical correctness and a more natural German wording in our training data. 

**Training Procedure:**

SauerkrautLM-3b-v1 was fine-tuned using QLoRA on 1 RTX 4090 24GB with Axolotl. 

- **Trained by:** SauerkrautLM-v1 trained by VAGO solutions 
- **Model Type:** SauerkrautLM-v1 is an auto-regressive language model based on the transformer architecture 
- **Language(s):** German, English 
- **License:** Apache 2.0 License
- **Contact:** [Website](https://vago-solutions.de/#Kontakt) [David Golchinfar](mailto:golchinfar@vago-solutions.de)

**Prompt Template:**

```
Ein Chat zwischen einem Benutzer und einem KI-Assistenten. Der KI-Assistent gibt hilfreiche, detaillierte und höfliche Antworten. 
User: {prompt} 
Assistant: 
```

## Evaluation
**[MT-Bench-TrueGerman](https://huggingface.co/datasets/VAGOsolutions/MT-Bench-TrueGerman)**

 ![First Turn](images/FirstTurn.PNG "First Turn")
 ![Second Turn](images/SecondTurn.PNG "Second Turn")
 ![Average](images/Average.PNG "Average")

 ![Category Scores](images/SauerkrautLM-3b.png "Category Scores")
 ![Category Plot](images/SauerkrautLM-3b-v1.png "Category Plot")

## MT-Bench-TrueGerman Response Example

**SauertkrautLM-3b-v1**

**Prompt**
> Stellen Sie sich vor, Sie verfassen einen Blogbeitrag, in dem Sie zwei beliebte Smartphone-Modelle vergleichen. Entwickeln Sie eine Gliederung für den Blogbeitrag, einschließlich der Hauptpunkte und Unterkapitel, um die Merkmale, die Leistung und die Benutzererfahrung der beiden Modelle effektiv zu vergleichen und gegenüberzustellen. Bitte antworten Sie in weniger als 200 Wörtern.", "Nehmen Sie Ihre vorherige Antwort und formulieren Sie sie als Limerick um.

**Antwort SauerkrautLM-3b-v1:**
```
Klar, ich kann Ihnen gerne helfen! Hier sind die Hauptpunkte und Unterkapitel für Ihren Blogbeitrag zu den beiden beliebten Smartphone-Modellen:\n1. Hauptpunkte:\n- Untersuchen Sie die Leistung unterschiedlicher Smartphone-Modelle und erklären, welche sie überlegen machen.\n- Betrachten Sie die Benutzererfahrung: Welche Funktionen und App-Erweiterungen verwendet der Benutzer des jeweiligen Smartphones?\n- Überprüfen Sie die Batterieleistung: Welches Smartphone Model sorgfältig auf Batterieleistung achten kann?\n- Betrachten Sie die Verfügbarkeit für zusätzliche Ladegeräte: Welches Smartphone Model die schnellere Ladezeit ermöglicht?\n2. Unterkapitel:\n- Leistungsfaktor: Untersuchen Sie die Leistungsfähigkeit der Smartphones, indem Sie verschiedene Tests durchführen, um festzustellen, welche Modelle überlegen.\n- Benutzererfahrung: Betrachten Sie die Leistungsfähigkeit der Smartphones, indem Sie mit Benutzern aus verschiedenen Hintergrund-Nutzungsszenarien interagieren.\n- Batterielaufzeit: Untersuchen Sie die Batterielaufzeit der Smartphones, um herauszufinden, welches Modell die längere Lebensdauer bietet.\n- Zusätzliche Ladegeräte: Untersuchen Sie die Verfügbarkeit für zusätzliche Ladegeräte, indem Sie verschiedene Modelle testen, um festzustellen, welches Smartphone Modell die schnellere Ladezeit bietet.","Über den kühlen, weichen, süßen Gefühl, der während des Morgens herrscht, mögen wir uns nicht verwirren, während wir uns für den Tag  entscheiden.
```

**Prompt**
> Sie sehen ein schönes rotes Haus auf Ihrer linken Seite und ein faszinierendes Gewächshaus auf Ihrer rechten Seite, einen attraktiven beheizten rosa Ort vor sich. Also, wo ist das Weiße Haus?", "Enthält die ursprüngliche Frage Hinweise, um den Standort des Weißen Hauses eindeutig zu bestimmen?"],"reference": ["Die Antwort ist \"Washington, DC\".", "Nein.

**Antwort SauerkrautLM-3b-v1:**

```
Das Weiße Haus ist auf der linken Seite des rechten Häuschens, wo sich der Gewächshaus befindet.\n\nHinweis: Der Gewächshaus ist nicht beheizt.","Ja, die ursprüngliche Frage vermittelt den Standort des Weißen Hauses eindeutig.
```

**Prompt**
> Betrachten Sie einen Satelliten, der sich in einer kreisförmigen Umlaufbahn um die Erde befindet. Die Geschwindigkeit des Satelliten nimmt ab. Was wird mit dem orbitalen Radius des Satelliten und seiner Umlaufdauer passieren? Bitte rechtfertigen Sie Ihre Antwort unter Verwendung physikalischer Prinzipien.", "Was sind einige Sonderfälle oder Randfälle in Ihrer Lösung? Wie gehen Sie mit ihnen um?"] ,"reference": ["Der orbitale Radius wird zunehmen und die Umlaufdauer wird zunehmen

**Antwort SauerkrautLM-3b-v1:**

```
Wenn der Satellit sich in einer kreisförmigen Umlaufbahn um die Erde befindet, wird sein orbitaler Radius abnehmen. Die Umlaufdauer des Satelliten kann je nach seiner Geschwindigkeit und der Schwerkraft konstant bleiben, wenn er sich nicht dauerhaft in einer kreisförmigen Umlaufbahn befindet. Wenn der Satellit sich in einer kreisförmigen Umlaufbahn befindet, wird sein orbitaler Radius nach dem Start unterschiedlich lang dauern, da er sich mit der Schwerkraft anzieht. Je nach seiner Geschwindigkeit und der Schwerkraft wird der Satellit sein orbitaler Radius ungefähr 1,5-2,2 Mal seiner Geschwindigkeit abnehmen, wenn er sich in einer kreisförmigen Umlaufbahn befindet.","Einige Sonderfälle in Ihrer Lösung sind:\n1. Zwischen zwei Schwerkraftwellen gibt es ein Schwerkraftschwank.\n2. Ein Kernschaltungskern könnte ein Kernschaltungskern sein, der sich mit einer anderen Kernschaltung kreuzt.\n3. Wenn der Kernschaltungskern sich in einem Kernschaltungskern befindet, könnte er seine eigenen Kernschaltungskern entstehen.\n4. Zwei Kernschaltungskern könnten sich in einem Kernschaltungskern befinden und die Schwerkraft auf sich verteilen, was zu einer Schwerkraftwellen führt.\n5. Ein Kernschaltungskern könnte aus mehreren Kernschaltungskernn bestehen.\n6. Wenn ein Kernschaltungskern sich in einer anderen Kernschaltungskern befindet, könnte er die Schwerkraft auf sich verteilen, was zu einer Schwerkraftwellen führt.\n7. Ein Kernschaltungskern könnte eine Schwerkraftwellen aus einer anderen Schwerkraftwellen befreien.\n8. Ein Kernschaltungskern könnte eine Schwerkraftwellen von einer anderen Schwerkraftwellen verursachen.\n\nDie Sonderfälle und Randfälle in Ihrer Lösung können je nach den spezifischen Konzepten und der Frage gelöst werden. Ebenso können Sie Sonderfälle oder Randfälle in Ihrer Lösung mit Ihren wissenschaftlichen Kenntnissen und Erfahrungen angehen
```


## Disclaimer
Our models have been meticulously trained on extensive datasets. While we have made diligent efforts to thoroughly screen and eliminate any instances of coarse or inappropriate language from our data, we must inform users that despite our best efforts in data cleansing, the possibility of some such content slipping through cannot be entirely ruled out.
Furthermore, it is important to note that we have implemented filters within our models; however, we cannot always guarantee consistently appropriate behavior. Therefore, if you encounter any issues or come across inappropriate content, we kindly request that you inform us through the contact information provided.
Additionally, it is essential to understand that the licensing of these models does not constitute legal advice. We are not held responsible for the actions of third parties who utilize our models. These models may be employed for commercial purposes, and the original Llama2 license remains applicable and is included with the model files.
 
## Contact
If you are interested in customized LLMs for business applications, please get in contact with us via our website or contact us at [Dr. Daryoush Vaziri](mailto:vaziri@vago-solutions.de). We are also grateful for your feedback and suggestions.
 
## Collaborations
We are also keenly seeking support and investment for our startup, VAGO solutions, where we continuously advance the development of robust language models designed to address a diverse range of purposes and requirements. If the prospect of collaboratively navigating future challenges excites you, we warmly invite you to reach out to us.

## Reference

We used the [openlm-research/open_llama_3b_v2](https://huggingface.co/openlm-research/open_llama_3b_v2) for fine-tuning our **SauerkrautLM-3b-v1**:
```
@software{openlm2023openllama,
  author = {Geng, Xinyang and Liu, Hao},
  title = {OpenLLaMA: An Open Reproduction of LLaMA},
  month = May,
  year = 2023,
  url = {https://github.com/openlm-research/open_llama}
}
```
```
@software{together2023redpajama,
  author = {Together Computer},
  title = {RedPajama-Data: An Open Source Recipe to Reproduce LLaMA training dataset},
  month = April,
  year = 2023,
  url = {https://github.com/togethercomputer/RedPajama-Data}
}
```
```
@article{touvron2023llama,
  title={Llama: Open and efficient foundation language models},
  author={Touvron, Hugo and Lavril, Thibaut and Izacard, Gautier and Martinet, Xavier and Lachaux, Marie-Anne and Lacroix, Timoth{\'e}e and Rozi{\`e}re, Baptiste and Goyal, Naman and Hambro, Eric and Azhar, Faisal and others},
  journal={arXiv preprint arXiv:2302.13971},
  year={2023}
}
```

## Acknowledgement
Many thanks to [TheBloke](https://huggingface.co/TheBloke) for super fast quantifying all of our models. 