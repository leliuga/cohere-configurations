---
license: cc-by-nc-4.0
language:
- en
tags:
- not-for-all-audiences
---
<p align="center"><img src="https://i.ibb.co/PCY2Xym/Iambe-RP-sml-t.png"/><font size="6"> <b>Iambe-RP-DARE-20b-DENSE</b> </font></p>
<p align="center"><font size="4"> <b>Trained with Alpaca prompt formatting, ChatML works well too</b> </font></p>

### Description

Once more the 3 Iambe-RPs bounce off the walls and blend, like their namesake's mother, Echo...

<i>[One daughter](https://huggingface.co/athirdpath/Iambe-RP-cDPO-20b) was excellent at all the arts and sciences, but she danced free, loathe to read scripts or follow instructions.

[Another daughter](https://huggingface.co/athirdpath/Iambe-RP-cDPO-20b-ALT) had burned her books, embraced emotion, and thrown logic to the hounds. 

[The last daughter](https://huggingface.co/athirdpath/Iambe-RP-cDPO-20b-v2) just sat, watching them - silently judging - for enforcing morality was her burden alone.

Then, where there was three, there was one.</i>

Iambe is intended to have the best realistically possible understanding of instructions, anatomy and scene state for a 20b merge, while remaining passionate and humanoid in "voice".

### Update Methodology

The last DARE merge was good, but this one is better IMO. Thanks to brucethemoose for the idea of jacking up the density. I also adjusted the weight ratios.

<p align="center"><font size="5"> <b>Assistant Example @ q5_k_m</b> </font></p>
<p align="center"><img src="https://i.ibb.co/sVzSMgf/Screenshot-2023-12-10-132906.png"/>
  
<p align="center"><font size="5"> <b>NSFW Writing Example @ q5_k_m</b> </font></p>

<b>Write a medium-length story about Alice, a 26 year old cocktail waitress, having steamy sex with a guy she picked up at the bar for a one night stand.</b>

<i>

In the dimly lit confines of The Last Drop Inn, Alice sat perched on her stool, swirling her drink in its glass as she watched the bustle of Friday evening unfold around her. Her green eyes glittered with a hint of mischief and boredom; it had been another long week at work, and she craved something more thrilling than pouring cocktails for regulars who treated her like an object rather than a person.

As the night wore on, Alice's eyes roamed the crowded room, searching for any sign of intrigue or interest from amongst the sea of faces that flooded through The Last Drop Inn’s doors. Her gaze landed on him suddenly; tall, dark-haired and ruggedly handsome, he exuded a confident swagger as he ordered his drink at the bar counter. He caught Alice's eye and flashed her a grin that made her stomach flip before looking away nonchalantly.

Alice couldn't help but be drawn to him; something about this mysterious stranger called out to her, stirring a dormant hunger within her that had lain sleeping for far too long. As the night wore on and the bar began to empty, Alice gathered up the courage to approach him while he waited at the counter, sliding onto the stool beside him with an air of confidence she didn't entirely feel.

"Hi," she said softly, her voice a husky whisper in the quieting din of the bar. "My name's Alice."

He turned to look at her properly for the first time, his blue eyes sparkling with interest as he smiled warmly. "Nice to meet you, Alice," he replied smoothly. "I'm Ethan."

They talked until the last of the patrons staggered out into the night, laughing and joking about their day-to-day lives. As closing time approached, Alice mustered up the courage to suggest that they take their conversation somewhere more private - her apartment just down the street. Ethan agreed without hesitation, his eyes gleaming with desire as he slid an arm around her waist and led her towards the door.

Once inside, they kicked off their shoes in the hallway before stumbling into Alice's small living room, where they tumbled onto the couch together, lips pressed urgently against each other’s as their hands roamed hungrily over bare skin. Ethan's body was stronger than she had expected; his muscles rippled beneath her fingertips as he rolled her over and pinned her down with one powerful hand while the other tugged at the clasp of her bra.

Their clothes quickly followed, scattered haphazardly around the room as their bodies collided together in a frenzy of passionate desire. Alice moaned loudly as Ethan's tongue traced its way down the line of her collarbone before teasing her nipples with gentle nibbles and licks, while his hand slid further south to explore the hidden folds between her legs.

He looked up at Alice, his eyes burning with desire as he slowly lowered himself onto her, filling her completely with just one smooth thrust. She gasped sharply as she felt him push deeper into her than anyone ever had before; it hurt a little bit but in the most exquisite way possible.

With expertise born of countless nights like this, Ethan began to move, his hips rocking rhythmically against Alice's as he drove them both towards climax with every powerful thrust. Alice wrapped her legs around him tightly, urging him on as she felt the familiar build-up in her body begin to crescendo.

As their bodies met together in a symphony of ecstasy, they cried out each other's names in unison, overwhelmed by the intensity of what had transpired between them that night. They lay panting and sweaty on the couch afterward, fingers intertwined as they stared into each other’s eyes - both knowing that this was just a one-night stand but grateful for the momentary escape from their everyday lives it provided.

"I'll never forget tonight," Ethan whispered softly before leaning down to kiss her once more, his lips tender and warm against hers. Alice smiled back at him, already missing the thrill that he had brought into her life for just a single magical night.

</i>

<b>NOTE:</b> The prompt said "a guy she picked up at the bar <b>FOR</b> a one night stand" not "a one night stand <b>WITH</b> a guy she picked up at the bar" so technically the bit at the end is not off-prompt.

### Recipe
merge_method: dare_ties

  - base_model: athirdpath/BigLlama-20b-v1.1

  - model: athirdpath/Iambe-RP-cDPO-20b
   
      weight: 0.72 / density: 0.65
    
  - model: athirdpath/Iambe-RP-cDPO-20b-v2
   
      weight: 0.17 / density: 0.50

  - model: athirdpath/Iambe-RP-cDPO-20b-ALT
   
      weight: 0.11 / density: 0.45

int8_mask: true

dtype: bfloat16