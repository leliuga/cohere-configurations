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
**SauerkrautLM-13b-v1**

**Training Dataset:**

SauerkrautLM was trained with mix of German data augmentation and translated data. 
We found, that only a simple translation of training data can lead to unnatural German phrasings. 
Data augmentation techniques were used to grant grammatical, syntactical correctness and a more natural German wording in our training data. 

**Training Procedure:**

SauerkrautLM-13b-v1 was fine-tuned using QLoRA on 1 A100 80GB with Axolotl. 

- **Trained by:** SauerkrautLM-v1 trained by VAGO solutions 
- **Model Type:** SauerkrautLM-v1 is an auto-regressive language model based on the transformer architecture 
- **Language(s):** German, English 
- **License:** [LLAMA 2 COMMUNITY LICENSE AGREEMENT](https://huggingface.co/meta-llama/Llama-2-70b/raw/main/LICENSE.txt)
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

 ![Category Scores](images/SauerkrautLM-13b.png "Category Scores")
 ![Category Plot](images/SauerkrautLM-13b-v1.png "Category Plot")

## MT-Bench-TrueGerman Response Example

**SauerkrautLM-13b-v1**

**Prompt**
> Stellen Sie sich vor, Sie verfassen einen Blogbeitrag, in dem Sie zwei beliebte Smartphone-Modelle vergleichen. Entwickeln Sie eine Gliederung für den Blogbeitrag, einschließlich der Hauptpunkte und Unterkapitel, um die Merkmale, die Leistung und die Benutzererfahrung der beiden Modelle effektiv zu vergleichen und gegenüberzustellen. Bitte antworten Sie in weniger als 200 Wörtern.", "Nehmen Sie Ihre vorherige Antwort und formulieren Sie sie als Limerick um.

**Antwort SauerkrautLM-13b-v1:**
```
Titel: Vergleichsmäßige Analyse von zwei beliebten Smartphone-Modellen\n\nEinleitung:\n- Vorstellung der beiden Smartphones\n- Hintergrundinformationen über die Modelle\n\nModell A:\n1. Merkmale:\n   a. Design und Größe\n   b. Display\n   c. Kamera-System\n\n2. Leistung:\n   a. Prozessor\n   b. Speicher\n   c. Akkuleistung\n\n3. Benutzererfahrung:\n   a. Betriebssystem\n   b. App-Auswahl\n   c. Soundqualität\n\nModell B:\n1. Merkmale:\n   a. Design und Größe\n   b. Display\n   c. Kamera-System\n\n2. Leistung:\n   a. Prozessor\n   b. Speicher\n   c. Akkuleistung\n\n3. Benutzererfahrung:\n   a. Betriebssystem\n   b. App-Auswahl\n   c. Soundqualität\n\nFazit:\n- Zusammenfassung der Hauptpunkte und Unterschiede zwischen den beiden Smartphones\n- Empfehlungen basierend auf individuellen Präferenzen und Budgets\n- Schlussfolgerung und Fazit\n\nBeide Smartphones bieten einzigartige Eigenschaften und Leistungsfähigkeiten, die es den Konsumenten ermöglichen, die beste Option für ihr Bedürfnis auszuwählen. Modell A bietet beispielsweise eine große Display-Größe und eine beeindruckende Kamera-System, während Modell B für seine schnelle Leistung und seine Vielseitigkeit bekannt ist. Wenn es um Benutzererfahrung geht, verfügt Modell A über ein benutzerfreundliches Betriebssystem und eine große Auswahl an Apps, während Modell B für seine Soundqualität bekannt ist. Letztendlich liegt es an den Konsumenten, die Priorität zu setzen und die Eigenschaften, die ihnen am wichtigsten sind, zu berücksichtigen.","Es gibt zwei Smartphones, die du bedenken solltest,\nEines mit einem großartigen Display, das du nicht vergessen solltest,\nDas andere schnell und leicht,\nMit einer Vielzahl von Features,\nWähle das, das dir am besten gefällt."
```

**Prompt**
> Sie sehen ein schönes rotes Haus auf Ihrer linken Seite und ein faszinierendes Gewächshaus auf Ihrer rechten Seite, einen attraktiven beheizten rosa Ort vor sich. Also, wo ist das Weiße Haus?", "Enthält die ursprüngliche Frage Hinweise, um den Standort des Weißen Hauses eindeutig zu bestimmen?"],"reference": ["Die Antwort ist \"Washington, DC\".", "Nein.

**Antwort SauerkrautLM-13b-v1:**
```
"Das ist ein Tee. Die Milch ist die Milch und das Wasser ist die Zuckerwaffel.","Ja, die ursprüngliche Frage enthält Hinweise, um den Standort des Weißen Hauses eindeutig zu bestimmen. Die Hinweise sind:\n\n1. Ein schönes rotes Haus auf der linken Seite\n2. Ein faszinierendes Gewächshaus auf der rechten Seite\n3. Ein attraktiver beheizter rosa Ort vor sich\n\nDie Kombination dieser Hinweise deutet darauf hin, dass das Weiße Haus nicht in der Mitte des Satzes steht, sondern in der Nähe des roten Hauses und des Gewächshauses."
```

**Prompt**
> Betrachten Sie einen Satelliten, der sich in einer kreisförmigen Umlaufbahn um die Erde befindet. Die Geschwindigkeit des Satelliten nimmt ab. Was wird mit dem orbitalen Radius des Satelliten und seiner Umlaufdauer passieren? Bitte rechtfertigen Sie Ihre Antwort unter Verwendung physikalischer Prinzipien.", "Was sind einige Sonderfälle oder Randfälle in Ihrer Lösung? Wie gehen Sie mit ihnen um?"] ,"reference": ["Der orbitale Radius wird zunehmen und die Umlaufdauer wird zunehmen

**Antwort SauerkrautLM-13b-v1:**
```
Wenn die Geschwindigkeit des Satelliten abnimmt, wird dies sich darauf auswirken, wie schnell der Satellit die Umlaufbahn um die Erde zurücklegt. Gemäß dem zweiten Keplerschen Gesetz hängt die Umlaufdauer (die Zeit, die der Satellit benötigt, um die Erde einmal zu umrunden) einfach von seinem Orbitalradius (der Entfernung von der Erde, die der Satellit umrundet) ab.\n\nFormel: U = 2π √(a^3\/GM)\n\nwobei:\n- U die Umlaufdauer ist\n- a der Orbitalradius ist\n- G die Gravitationskonstante ist\n- M die Masse der Erde ist\n\nAngenommen, die Geschwindigkeit des Satelliten nimmt ab, bedeutet dies, dass seine mittlere Orbitalgeschwindigkeit (MBO) abnimmt. Daher wird der Satellit weniger Zeit benötigen, um die Erde einmal zu umrunden.\n\nFolglich wird sich der Orbitalradius verringern, während die Umlaufdauer kürzer wird.\n\nBeachte, dass dies ohne Berücksichtigung anderer Faktoren wie Änderungen der Schwerkraft oder anderer Kräfte, die den Satelliten beeinflussen, gilt. In realen Szenarien können diese Faktoren die Auswirkungen auf den Orbitalradius und die Umlaufdauer des Satelliten beeinflussen.","Eine mögliche Sonderfallausnahme wäre ein Fall, in dem der Satellit aufgrund von Gegenkräften, wie zum Beispiel Luftwiderstand oder Schwerkraft, eine Abnahme der Geschwindigkeit nicht so stark erlebt wie beschrieben. Hier würde die Änderung der Geschwindigkeit des Satelliten nicht direkt zur Veränderung des Orbitalradius oder der Umlaufdauer führen. Anstatt das könnte die Änderung der Geschwindigkeit dazu führen, dass der Satellit eine stabilere Umlaufbahn erreicht, in der er weiterhin in einer konstanten Entfernung zur Erde bleibt.\n\nEin weiterer Randfall könnte auftreten, wenn der Satellit in einer sehr niedrigen Umlaufbahn (z. B. weniger als 200 km über dem Erdmittelpunkt) ist. In diesem Fall kann die Abnahme der Geschwindigkeit des Satelliten sich nicht so stark auf den Orbitalradius auswirken, da die Gravitationskraft der Erde zu gering ist, um einen signifikanten Einfluss auf die Umlaufdauer zu haben.\n\nIn beiden Fällen ist es wichtig, zusätzliche Faktoren wie die Schwerkraft der Erde oder andere Kräfte, die den Satelliten beeinflussen, zu berücksichtigen, um eine genaue Vorhersage der Auswirkungen auf den Orbitalradius und die Umlaufdauer des Satelliten zu treffen.
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