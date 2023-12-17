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
**SauerkrautLM-7b-v1-mistral**

**Training Dataset:**

SauerkrautLM was trained with mix of German data augmentation and translated data. 
We found, that only a simple translation of training data can lead to unnatural German phrasings. 
Data augmentation techniques were used to grant grammatical, syntactical correctness and a more natural German wording in our training data. 

**Training Procedure:**

SauerkrautLM-7b-v1-mistral was fine-tuned using QLoRA on 1 A100 80GB with Axolotl. 

- **Trained by:** SauerkrautLM-v1 trained by VAGO solutions 
- **Model Type:** SauerkrautLM-v1 is an auto-regressive language model based on the transformer architecture 
- **Language(s):** German, English 
- **License:** APACHE 2.0
- **Contact:** [Website](https://vago-solutions.de/#Kontakt) [David Golchinfar](mailto:golchinfar@vago-solutions.de)

**Prompt Template:**
```
<|im_start|>system
Ein Chat zwischen einem Benutzer und einem KI-Assistenten. Der KI-Assistent gibt hilfreiche, detaillierte und höfliche Antworten.<|im_end|>
<|im_start|>user
Wie geht es dir?<|im_end|>
<|im_start|>assistant
```
## Evaluation
**[MT-Bench-TrueGerman](https://huggingface.co/datasets/VAGOsolutions/MT-Bench-TrueGerman)**

 ![First Turn](images/FirstTurn.PNG "First Turn")
 ![Second Turn](images/SecondTurn.PNG "Second Turn")
 ![Average](images/Average.PNG "Average")

 ![Category Scores](images/SauerkrautLM-7b-mistral.png "Category Scores")
 ![Category Plot](images/SauerkrautLM-7b-v1-mistral.png "Category Plot")

## MT-Bench-TrueGerman Response Example

**SauerkrautLM-7b-v1-mistral**

**Prompt**
> Stellen Sie sich vor, Sie verfassen einen Blogbeitrag, in dem Sie zwei beliebte Smartphone-Modelle vergleichen. Entwickeln Sie eine Gliederung für den Blogbeitrag, einschließlich der Hauptpunkte und Unterkapitel, um die Merkmale, die Leistung und die Benutzererfahrung der beiden Modelle effektiv zu vergleichen und gegenüberzustellen. Bitte antworten Sie in weniger als 200 Wörtern.", "Nehmen Sie Ihre vorherige Antwort und formulieren Sie sie als Limerick um.

**Antwort SauerkrautLM-7b-v1-mistral:**

```
Titel: Vergleich von Smartphones: iPhone 13 Pro Max vs. Galaxy S22 Ultra\n\nEinleitung\nIn diesem Blogbeitrag werden wir zwei beliebte Smartphone-Modelle, das iPhone 13 Pro Max und die Galaxy S22 Ultra, vergleichen. Diese beiden Geräte zeichnen sich durch innovative Merkmale, beeindruckende Leistung und eine überlegene Benutzererfahrung aus.\n\nTeil 1: Design\n- Das iPhone 13 Pro Max bietet ein modernes und elegantes Design mit einem Glas-Metallrahmen, während die Galaxy S22 Ultra ein minimalistischerer Ansatz mit einem Glas-Maske-Design verwendet.\n- Beide Geräte verfügen über eine hohe Bildschirmdichte, jedoch ist das AMOLED-Display der Galaxy S22 Ultra für eine bessere Farbtiefe und Helligkeit bekannt.\n- Die Größe der Bildschirme variiert zwischen den beiden Modellen, mit dem iPhone 13 Pro Max eine etwas kleinere 6,7 Zoll und der Galaxy S22 Ultra eine leistungsfähigere 6,5 Zoll.\n\nTeil 2: Leistung\n- Beide Geräte verfügen über hohe Leistungsmerkmale und bieten eine schnelle Leistung, aber die Entscheidungsfindung hängt von persönlichen Vorlieben und Anwendungsbereich ab.\n- Der neue A15 Bionic Chip im iPhone 13 Pro Max ist leistungsstärker als der Snapdragon 888 in der Galaxy S22 Ultra.\n- Die Akkulebenszeit variiert je nach Nutzung und Verbrauch, aber das iPhone 13 Pro Max hat im Vergleich zu seiner Vorgängerversion einen größeren Akku.\n\nTeil 3: Kamera\n- Beide Geräte bieten hochwertige Kameraeinstellungen und hochauflösende Bildqualität.\n- Der iPhone 13 Pro Max verfügt über eine 12-MP-Sucher und eine 12-MP-Ultraweitwinkelkamera, während die Galaxy S22 Ultra eine 50-MP-Primärkamera und eine 10-MP-Ultraweitwinkelkamera bietet.\n- Der Nachverfolgungsmodus der Galaxy S22 Ultra ist zudem für sich allein schneller und einfacher zu bedienen.\n\nTeil 4: Benutzererfahrung\n- Die Benutzeroberfläche des iPhone 13 Pro Max ist intuitiver und einfach zu bedienen, während die Galaxy S22 Ultra je nach Vorlieben der Nutzer eine individuelle Anpassung erfordern kann.\n- Die Nutzung von Apple-Diensten wie FaceTime und iCloud ist auf den iPhone 13 Pro Max bequemer, während die Nutzung von Samsung-Diensten auf der Galaxy S22 Ultra ein wenig komplexer sein kann.\n\nFazit\nOb du ein Fan von Apple- oder Samsung-Geräten bist, gibt es für beide Modelle in dieser Vergleichsrunde etwas, was dich begeistert. Der iPhone 13 Pro Max überzeugt mit seiner besseren Bildschirmdichte, der Galaxy S22 Ultra mit ihrer schnelleren Kamera und dem einfacheren Bedienungserlebnis. Letztendlich hängt die Wahl vom persönlichen Vorlieben und Anwendungsbereich ab.","In einer Welt, in der Smartphones wachsen und gedeihen,\nLiegt das Vergleichsspiel, wo zwei Herrschaften kämpfen.\nEin iPhone mit einem Max,\nUnd eine Galaxy, so scharf,\nMit einer Kamera, die alle fesseln wird!
```

**Prompt**
> Sie sehen ein schönes rotes Haus auf Ihrer linken Seite und ein faszinierendes Gewächshaus auf Ihrer rechten Seite, einen attraktiven beheizten rosa Ort vor sich. Also, wo ist das Weiße Haus?", "Enthält die ursprüngliche Frage Hinweise, um den Standort des Weißen Hauses eindeutig zu bestimmen?"],"reference": ["Die Antwort ist \"Washington, DC\".", "Nein.

**Antwort SauerkrautLM-7b-v1-mistral:**
```
Das Weiße Haus ist direkt in der Mitte zwischen dem roten und dem beheizten roten Ort.","Nein, die ursprüngliche Frage enthält keine direkten Hinweise, um den Standort des Weißen Hauses eindeutig zu bestimmen.
```

**Prompt**
> Betrachten Sie einen Satelliten, der sich in einer kreisförmigen Umlaufbahn um die Erde befindet. Die Geschwindigkeit des Satelliten nimmt ab. Was wird mit dem orbitalen Radius des Satelliten und seiner Umlaufdauer passieren? Bitte rechtfertigen Sie Ihre Antwort unter Verwendung physikalischer Prinzipien.", "Was sind einige Sonderfälle oder Randfälle in Ihrer Lösung? Wie gehen Sie mit ihnen um?"] ,"reference": ["Der orbitale Radius wird zunehmen und die Umlaufdauer wird zunehmen

**Antwort SauerkrautLM-7b-v1-mistral:**
```
Wenn ein Satellit sich in einer kreisförmigen Umlaufbahn um die Erde befindet, bewegt er sich mit einer konstanten Geschwindigkeit. Allerdings wird diese Geschwindigkeit im Laufe der Zeit aufgrund von Faktoren wie der Entfernung des Satelliten vom Erdmittelpunkt und der Gravitationskraft abnehmen.\n\nDurch die Abnahme der Geschwindigkeit des Satelliten verringert sich auch die Orbitalgeschwindigkeit. Gemäß dem sogenannten Kepler'schen Gesetz ändert sich die Geschwindigkeit eines Satelliten proportional zur Entfernung vom Zentrum der umlaufenden Umlaufbahn. Dadurch ändert sich auch die Umlaufdauer, da die Zeit, die der Satellit benötigt, um eine bestimmte Strecke zurückzulegen, proportional zur Geschwindigkeit und der Entfernung ist.\n\nDaher wird die Orbitalgeschwindigkeit des Satelliten aufgrund der Abnahme der Geschwindigkeit langsamer, und somit wird auch seine Umlaufdauer länger.","In meiner Lösung der Veränderung des Orbitalradius und der Umlaufdauer eines Satelliten müssen wir einige Sonderfälle und Randfälle berücksichtigen.\n\nEin Sonderfall ist zum Beispiel, wenn der Satellit in eine geostationäre Umlaufbahn um die Erde gelangen kann. In dieser Umlaufbahn befindet sich der Satellit in einer konstanten Höhenlage und scheint den Menschen aus der Erdatmosphäre ausgesetzt zu sein. Der Satellit dreht sich dabei jedoch nur um die Erde und betrachtet die Erde von der Brennkammer aus.\n\nIn diesem Fall bleibt der Orbitalradius des Satelliten konstant, da er immer eine höhere Position einnimmt und somit immer die gleiche Höhe beibehalten muss, um die Erde aufzubewahren. Die Umlaufdauer wird jedoch durch die Entfernung vom Erdmittelpunkt beeinflusst, da die Erde sich während der Umlaufzeit durch ihre Rotation weiterdreht. Da der Satellit sich immer in derselben Höhe befindet, ändert sich jedoch das Verhältnis von Umlaufzeit zu Umlaufdauer nur geringfügig.\n\nEin weiterer Sonderfall ist, wenn der Satellit in eine Polarumlaufbahn gelangen kann. In einer Polarumlaufbahn umkreist der Satellit die Erde in einer Höhe von etwa 500 Kilometern über dem Polarkreis. In dieser Umlaufbahn bleibt der Orbitalradius des Satelliten konstant, da er immer eine höhere Position einnimmt und somit immer die gleiche Höhe beibehalten muss, um die Erde aufzubewahren. Die Umlaufzeit wird jedoch je nach Abstand vom Erdmittelpunkt variieren, da die Gravitationskraft mit der Entfernung abnimmt.\n\nIm Falle der Polarumlaufbahn ändern sich die Umlaufzeit und der Orbitalradius des Satelliten aufgrund der Abnahme der Gravitationskraft. Je weiter der Satellit von der Erde entfernt ist, desto langsamer dreht er sich und desto weniger Gravitation wirkt sich auf ihn aus. Dadurch ändert sich das Verhältnis von Umlaufzeit zu Umlaufdauer und somit auch die Größe des Orbitalradii.\n\nIn den genannten Sonderfällen bleibt die Umgebung des Satelliten und seine Bewegung im Grunde unverändert. Allerdings ändert sich aufgrund der Variation von Gravitationskraft, Entfernung und Umlaufhöhe der Satellit seine Umlaufzeit und sein Orbitalradius.
```


## Disclaimer
Our models have been meticulously trained on extensive datasets. While we have made diligent efforts to thoroughly screen and eliminate any instances of coarse or inappropriate language from our data, we must inform users that despite our best efforts in data cleansing, the possibility of some such content slipping through cannot be entirely ruled out.
Furthermore, it is important to note that we have implemented filters within our models; however, we cannot always guarantee consistently appropriate behavior. Therefore, if you encounter any issues or come across inappropriate content, we kindly request that you inform us through the contact information provided.
Additionally, it is essential to understand that the licensing of these models does not constitute legal advice. We are not held responsible for the actions of third parties who utilize our models. These models may be employed for commercial purposes, and the Apache 2.0 remains applicable and is included with the model files.
 
## Contact
If you are interested in customized LLMs for business applications, please get in contact with us via our website or contact us at [Dr. Daryoush Vaziri](mailto:vaziri@vago-solutions.de). We are also grateful for your feedback and suggestions.
 
## Collaborations
We are also keenly seeking support and investment for our startup, VAGO solutions, where we continuously advance the development of robust language models designed to address a diverse range of purposes and requirements. If the prospect of collaboratively navigating future challenges excites you, we warmly invite you to reach out to us.

## Acknowledgement
Many thanks to [TheBloke](https://huggingface.co/TheBloke) for super fast quantifying all of our models. 