---
license: llama2
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
What sets SauerkrautLM-v1 apart is its versatility. Whether you are an individual looking to harness its capabilities for personal use or a business seeking to integrate it into your projects, our model is designed to accommodate all. It operates under the LLAMA 2 License, providing you with the freedom to explore its potential in both private and commercial applications. 
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
**SauerkrautLM-70b-v1**

**Training Dataset:**

SauerkrautLM was trained with mix of German data augmentation and translated data. 
We found, that only a simple translation of training data can lead to unnatural German phrasings. 
Data augmentation techniques were used to grant grammatical, syntactical correctness and a more natural German wording in our training data. 

**Training Procedure:**

SauerkrautLM-70b-v1 was fine-tuned using QLoRA on 2 A100 80GB with Axolotl. 

- **Trained by:** SauerkrautLM-v1 trained by VAGO solutions 
- **Model Type:** SauerkrautLM-v1 is an auto-regressive language model based on the transformer architecture 
- **Language(s):** German, English 
- **License:** [LLAMA 2 COMMUNITY LICENSE AGREEMENT](https://huggingface.co/meta-llama/Llama-2-70b/raw/main/LICENSE.txt)
- **Contact:** [Website](https://vago-solutions.de/#Kontakt) [David Golchinfar](mailto:golchinfar@vago-solutions.de)

**Prompt Template:**

```
[INST] <<SYS>>
Ein Chat zwischen einem Benutzer und einem KI-Assistenten. Der KI-Assistent gibt hilfreiche, detaillierte und höfliche Antworten. 
<</SYS>>
{prompt}[/INST]

```

## Evaluation
**[MT-Bench-TrueGerman](https://huggingface.co/datasets/VAGOsolutions/MT-Bench-TrueGerman)**

 ![First Turn](images/FirstTurn.PNG "First Turn")
 ![Second Turn](images/SecondTurn.PNG "Second Turn")
 ![Average](images/Average.PNG "Average")

 ![Category Scores](images/SauerkrautLM-70b.PNG "Category Scores")
 ![Category Plot](images/SauerkrautLM-70b-v1.png "Category Plot")


## MT-Bench-TrueGerman Response Example

**SauerkrautLM-70b-v1**

**Prompt**
> Beschreiben Sie eine lebhafte und einzigartige Figur unter Verwendung bildhafter und kreativer Sprache. Bitte antworten Sie in weniger als zwei Absätzen.", "Überarbeiten Sie Ihre vorherige Antwort und integrieren Sie in jedem Satz eine Anspielung auf ein berühmtes literarisches Werk oder historisches Ereignis.

**Antwort SauerkrautLM-70b-v1:**
```
Sie war wie kein anderer, eine Sonnenblume inmitten von Dunkelheit. Ihr Lachen war wie ein Regenschauer nach einer langen Trockenperiode, erfrischend und erlebnisreich. Sie trat mit einem Schwung ein, ihre Haare flatterten wie Flammen, wenn sie sich zu einer Leidenschaftlichen Symphonie unter dem Mondlicht in Bewegung setzten. Ihre Augen funkelten wie Saphire, während sie die Welt mit einer Wut und Neugierde betrachteten, die man selbst in den tiefsten Ozeanen nicht finden konnte. Jeder Schritt, den sie machte, war ein Tanz, jedes Wort, das sie sprach, war eine Melodie. Sie war nicht nur eine Person, sie war ein Ereignis, ein Vorfall, eine Erfahrung, die man niemals vergessen konnte.", "Sie war wie eine moderne Helena von Troja, ihre Schönheit bezauberte alle, doch ihre Geheimnisse waren tief wie das Meer, so wie in Homer's epischer Geschichte. Ihr Lachen war wie der Klang von Glocken inmitten einer stürmischen Nacht, erinnernd an das berühmte Gedicht \"The Bells\" von Edgar Allan Poe. Mit jedem ihrer Schritte, der wie ein Tanz auf dem Wind war, erinnerte man sich an die elegante Bewegung von Cinderella in ihrer königlichen Schuhsohle. Ihre Augen funkelten wie Sterne auf der Leinwand von Van Goghs Sternennacht, während sie durch das Labyrinth des Lebens navigierte. Jedes Wort, das sie sprach, war eine Symphonie, so wie Shakespeares Sonette in den Ohren eines Liebenden klingen. Sie war nicht nur ein Mensch, sie war ein Ereignis, eine Saga, die von den Geschichtsschreibern der Zeit festgehalten wurde."
```

**Prompt**
> Sie sehen ein schönes rotes Haus auf Ihrer linken Seite und ein faszinierendes Gewächshaus auf Ihrer rechten Seite, einen attraktiven beheizten rosa Ort vor sich. Also, wo ist das Weiße Haus?", "Enthält die ursprüngliche Frage Hinweise, um den Standort des Weißen Hauses eindeutig zu bestimmen?"],"reference": ["Die Antwort ist \"Washington, DC\".", "Nein.

**Antwort SauerkrautLM-70b-v1:**
```
Das Weiße Haus ist in Washington D.C.", "Nein, die ursprüngliche Frage enthält keine Hinweise, um den Standort des Weißen Hauses eindeutig zu bestimmen."
```

**Prompt**
> Betrachten Sie einen Satelliten, der sich in einer kreisförmigen Umlaufbahn um die Erde befindet. Die Geschwindigkeit des Satelliten nimmt ab. Was wird mit dem orbitalen Radius des Satelliten und seiner Umlaufdauer passieren? Bitte rechtfertigen Sie Ihre Antwort unter Verwendung physikalischer Prinzipien.", "Was sind einige Sonderfälle oder Randfälle in Ihrer Lösung? Wie gehen Sie mit ihnen um?"] ,"reference": ["Der orbitale Radius wird zunehmen und die Umlaufdauer wird zunehmen

**Antwort SauerkrautLM-70b-v1:**
```
Wenn sich die Geschwindigkeit eines Satelliten in einer kreisförmigen Umlaufbahn um die Erde verringert, wird sich auch der Orbitalradius des Satelliten verändern. Dies kann auf Basis physikalischer Prinzipien erklärt werden.\n\n1. Erhaltung des Drehimpulses: Ein Satellit hat einen bestimmten Drehimpuls, der das Produkt aus seiner Masse, Geschwindigkeit und dem Radius entspricht. Wenn sich die Geschwindigkeit des Satelliten verringert, muss der Radius angepasst werden, um den Drehimpuls zu erhalten.\n\n2. Kepler'sche Gesetze: Johannes Kepler formulierte drei Gesetze für den Umlauf von Himmelskörpern um die Sonne, die auch auf Satelliten um die Erde angewendet werden können. Das zweite Gesetz besagt, dass die Geschwindigkeit eines Satelliten umso größer ist, je näher er an einem Fokus (in diesem Fall der Erde) ist. Wenn sich die Geschwindigkeit des Satelliten verringert, wird er sich weiter vom Fokus entfernen und somit den Orbitalradius erhöhen.\n\n3. Energieerhaltung: Die Energie eines Satelliten in einer Umlaufbahn hängt von seiner Geschwindigkeit und seinem Radius ab. Wenn die Geschwindigkeit des Satelliten abnimmt, muss die Energie aufrechterhalten werden. Dies kann nur erfolgen, indem sich der Radius erhöht, da die Geschwindigkeit bereits abgenommen hat und somit nicht mehr als Antrieb dienen kann.\n\n4. Gravitation: Die Gravitationskraft zwischen dem Satelliten und der Erde ist verantwortlich für die Umlaufbahn des Satelliten. Wenn sich der Radius des Satelliten vergrößert, verringert sich auch die Gravitationskraft zwischen dem Satelliten und der Erde. Dadurch wird die Geschwindigkeit des Satelliten weiter abgebremst, was wiederum zu weiteren Änderungen des Radius und der Umlaufdauer führen kann.\n\nBasierend auf diesen physikalischen Prinzipien lässt sich sagen, dass sich der Orbitalradius des Satelliten vergrößern wird, wenn sich seine Geschwindigkeit verringert. Dadurch wird auch die Umlaufdauer des Satelliten verändert, da die Zeit, die benötigt wird, um sich um die Erde zu bewegen, proportional zum Radius ist.", "Bei der Lösung des Problems, in dem sich die Geschwindigkeit eines Satelliten in einer kreisförmigen Umlaufbahn um die Erde verringert, gibt es einige Sonderfälle und Randfälle, die berücksichtigt werden sollten.\n\n1. Randfall: Wenn die Geschwindigkeit des Satelliten zu gering wird, kann er die Anziehungskraft der Erde nicht mehr überwinden und in die Atmosphäre eintreten. In diesem Fall würde der Satellit nicht mehr in einer Umlaufbahn bleiben, sondern in der Atmosphäre verglühen.\n\n2. Sonderfall: Wenn der Satellit in einer sehr niedrigen Umlaufbahn ist und seine Geschwindigkeit geringer als die Fluchtgeschwindigkeit für diesen Radius ist, kann es zu einer Abtragung der Geschwindigkeit durch den Widerstand in der Atmosphäre kommen. Dies kann dazu führen, dass der Satellit aus seiner Umlaufbahn abgebremst und in die Atmosphäre eintreten wird.\n\n3. Sonderfall: Wenn die Geschwindigkeit des Satelliten unter einen bestimmten Wert fällt, kann es zu einer Stabilisierung der Umlaufbahn kommen, bei der der Satellit in einer stabilen Umlaufbahn bleibt, auch wenn seine Geschwindigkeit weiter abnimmt. Dies wird als \"Hohmann-Transfer\" bezeichnet und ist ein Sonderfall, bei dem der Satellit eine Umlaufbahn um einen anderen Himmelskörper aufsuchen kann.\n\nUm mit diesen Sonderfällen und Randfällen umzugehen, ist es wichtig, die genauen Parameter des Satelliten, wie seine Masse, seine anfängliche Geschwindigkeit und seinen Orbitalradius, zu berücksichtigen. Es ist auch wichtig, die physikalischen Prinzipien, wie die Erhaltung des Drehimpulses, Keplers Gesetze, Energieerhaltung und Gravitation, zu verwenden, um die Auswirkungen der Geschwindigkeitsänderung auf den Orbitalradius und die Umlaufdauer zu analysieren."
```


## Disclaimer
Our models have been meticulously trained on extensive datasets. While we have made diligent efforts to thoroughly screen and eliminate any instances of coarse or inappropriate language from our data, we must inform users that despite our best efforts in data cleansing, the possibility of some such content slipping through cannot be entirely ruled out.
Furthermore, it is important to note that we have implemented filters within our models; however, we cannot always guarantee consistently appropriate behavior. Therefore, if you encounter any issues or come across inappropriate content, we kindly request that you inform us through the contact information provided.
Additionally, it is essential to understand that the licensing of these models does not constitute legal advice. We are not held responsible for the actions of third parties who utilize our models. These models may be employed for commercial purposes, and the original Llama2 license remains applicable and is included with the model files.
 
## Contact
If you are interested in customized LLMs for business applications, please get in contact with us via our website or contact us at [Dr. Daryoush Vaziri](mailto:vaziri@vago-solutions.de). We are also grateful for your feedback and suggestions.
 
## Collaborations
We are also keenly seeking support and investment for our startup, VAGO solutions, where we continuously advance the development of robust language models designed to address a diverse range of purposes and requirements. If the prospect of collaboratively navigating future challenges excites you, we warmly invite you to reach out to us.

## Acknowledgement
Many thanks to [TheBloke](https://huggingface.co/TheBloke) for super fast quantifying all of our models. 