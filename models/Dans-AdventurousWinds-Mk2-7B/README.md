---
language:
- en
datasets:
- PocketDoc/Floyd-Text-Adventures
- PocketDoc/Choose-Your-Story-Long-Text-Adventures
license: apache-2.0
base_model: mistralai/Mistral-7B-v0.1
---

**16k training sequences**

**Quantized versions courtesy of [TheBloke](https://huggingface.co/TheBloke):**
- [GGUF](https://huggingface.co/TheBloke/Dans-AdventurousWinds-Mk2-7B-GGUF) - For CPU and or GPU
- [GPTQ](https://huggingface.co/TheBloke/Dans-AdventurousWinds-Mk2-7B-GPTQ) - For GPU
- [AWQ](https://huggingface.co/TheBloke/Dans-AdventurousWinds-Mk2-7B-AWQ) - New format for GPU

# What is the model for?

This model is proficient in crafting text-based adventure games. It can both concise replies and more expansive, novel-like descriptions. The ability to alternate between these two response styles can be triggered by a distinct system message.

# What's in the sausage?

This model was trained on [Mistral-7b](https://huggingface.co/mistralai/Mistral-7B-v0.1) using a custom version of the floyd and cys datasets from the skein text adventure dataset. The dataset has been improved over what was used in my previous model Dans-AdventurousWinds, in addition it had been trained on 16k sequence lengths.

## Prompt format and examples:
#### Format description:

Traditional adventure format with a description of the genre, tags, and other misc information at the top followed by an optional introductory message. User messages are prefixed with two new lines and "> ", model messages are one new line after the last user message or the description message if an introductory message is desired.

#### Examples of shorter responses:


> [Genres: Science Fiction]
> [Tags: humor, old school, sci fi]
> [Mode: Adventure]
> [Description: A puzzle about committing acts of financial skulduggery and exploiting ridiculous magical items.]
> [Misc: Writing era: 1993]
> [Intro]
> You are an interstellar trader.
> 
> On your way to the galactic capital of Coruscant to sell a large quantity of fettucini, you make a detour through a planetary system with a reputation for being dangerous and unstable. The reason is simple -- there is a very powerful magical item known as "The Sword Of Damocles" up for sale. You have no idea what it does, but you've heard that the previous owners all ended up dead or broke. In other words, these people were so desperate to get rid of the sword that they would probably be willing to take almost any price. It sounds like a great deal!
> 
> Of course, you don't have enough money to buy the sword outright, which means that you will need to find some way to finance the purchase. There is only one possible source of funds in this part of space -- the loan sharks who operate out of a nearby planet called Riggerton. But these loans come at a high interest rate; if you can't pay them back quickly, you may end up losing everything.
> 
> You also know that there are several potential buyers vying for ownership of the sword, including some of the most ruthless criminals in the galaxy. These people might be willing to kill each other over this object, making it more valuable than ever before! However, they are notoriously paranoid and suspicious, so you must proceed with extreme caution when dealing with them...
> 
> Finally, there is always the chance that someone else will beat you to the punch by buying the sword first. If that happens then you won't be able to afford it anymore anyway, so maybe it doesn't matter too much?
> 
> In any case, good luck! And remember: never borrow money from anyone unless you absolutely trust him completely. Otherwise he could turn around and sell YOU into slavery instead!
> 
> Ready? Then let's begin...
> 
> \> You snap out of your daydream and remember you're still in port.
> You should probably get ready to go somewhere now...
> 
> There are three main places where traders usually hang out: the docking bay (where ships land), the casino (a place where people gamble) and the black market (an illegal bazaar). Each has its own advantages and disadvantages depending on what kind of goods you want to trade. For example, the docking bay is the safest option because it's run by the government, but it charges higher taxes; while the black market is riskier but cheaper.
> 
> So where do you think you ought to go?
> 
> \> You get ready for launch
> Okay, I guess we better start looking for something interesting then...
> 
> Before we leave though, I just wanted to mention that there are two ways to travel between planets: hyperspace jumps and wormholes. Both methods require special technology that only exists inside starships, but they work differently.
> 
> A hyperspace jump involves using a hyperdrive engine to accelerate faster than light speed until you reach another galaxy entirely. Once there, you slow down again and return home. This method takes longer than a wormhole, but it allows you to visit distant worlds without having to worry about getting lost along the way.
> 
> Wormholes however, allow you to travel instantly across vast distances without needing to slow down at either end. Unfortunately, this type of transportation isn't available everywhere, plus it requires extremely powerful computers to map out their locations accurately. Still, if you happen to find yourself near one of these things then I strongly recommend taking advantage of it since it saves time.
> 
> Anyway, let me know whenever you're ready to go. We'll see how far away our next stop lies...
> 
> \> You check the database on wormhole technology
> Okay, here goes nothing...
> 
> According to my research, there are currently three types of wormholes known to science: stable, unstable and artificial. Stable ones occur naturally throughout the universe and tend to stay open permanently. Unstable ones close after passing through once, whereas artificially created ones remain open indefinitely.
> 
> As far as I know, none of these types are suitable for commercial use yet. However, scientists believe that further development will eventually lead to practical applications in areas such as teleportation and faster-than-light travel. Until then though, I doubt whether anything useful can be achieved simply by reading books about the subject.
> 
> Sorry if I disappointed you there, but I hope this information proves useful nonetheless. Good luck with whatever comes next!
> 
> P.S.: Before leaving, please ensure that all loose objects have been put away properly. Otherwise they could fall off the ship during flight and cause damage later on. Thanks!


#### Examples of longer novel like responses:


> [Genres: grimdark fantasy] [Tags: Fantasy, Grimdark, Serious, Villain Protagonist, War] [Description: This story is definitely more "story" than "game", however there are two distinct paths that can be followed, each with their own pitfalls and additional choices. This story can get long, and assuming you're surviving, years do pass in it and you and your environment will change.] [First message intro] [Mode: Interactive Storyteller]
> You have been a mercenary for many years now. You are an accomplished warrior and battlefield commander. Your company consists of about 50 men and women. Most are veteran soldiers from different countries, some even from as far away as the Americas. They are all loyal to you. The only thing keeping them from being devoted followers would be your atheism.
> 
> The war has raged on for many years now. The enemy, the Kingdom of Erenland, is a nation of fanatics who worship a single god called Galdrin. Their religion spread across the land like wildfire after their king ascended to the throne. He declared himself a prophet and made his people believe they were invincible. They started invading neighboring nations to bring them into the fold. The king of your homeland, King Dagor, was one of the first to resist this. When he refused to allow missionaries into his country, the Erenlanders invaded anyway. It took several years before King Dagor managed to gather enough forces to repel the invasion. The war lasted for decades and finally ended when Dagor's grandson ascended to the throne. At that point, both sides realized how much blood had already been spilled and agreed to a truce.
> 
> Unfortunately, the peace didn't last very long. A few years ago, King Dagor II died without any children. His cousin became the new monarch but she was weak and incompetent. She allowed her advisors to make most of the important decisions while she spent her time partying with nobles from other nations. One of these advisors convinced the queen to send troops to help quell rebellions that were taking place in the Erenlander cities. These rebellions were caused by dissatisfaction with the way the crown was handling things. They wanted to overthrow their ruler and install a new one that would listen to them.
> 
> When word reached the court that the rebels had succeeded in killing the king, everyone panicked. The Queen ordered her army to march north immediately. Unfortunately, she hadn't thought through what would happen if the enemy army attacked during this period of chaos. The result was disastrous. The Erenlanders caught our army completely off guard and annihilated us almost entirely. Only a handful survived to tell the tale. Among those survivors was you.
> 
> You have come back home to bury your fallen comrades. You know you won't find rest until you avenge their deaths. But where do you begin? How do you stop an entire nation from going to war against yours? And how do you convince others to join you when they see you as nothing more than a madman bent on destruction?
> 
> These questions haunt you every night. You lie awake staring at the ceiling wondering how you could possibly make things right again. Then you remember something...
> 
> There was once a great hero named Galadar who lived hundreds of years ago. He fought against tyranny and oppression throughout his lifetime. People said that he possessed magical powers which gave him strength beyond anything mortal. Some even claimed that he walked among the stars. After retiring from active duty, Galadar vanished mysteriously. No one knows where he went or why. Rumors say that he returned to fight alongside the gods against evil forces threatening to destroy existence itself. Whatever happened to him, no one knows for sure. All we do know is that his legacy lives on today through those who follow his example. Perhaps you too can become such a person. Perhaps you can use your talents to save the world from darkness just like he did so long ago. If you feel ready, then take up arms and prepare yourself for battle!
> 
> Your journey begins here...
> 
> You wake up feeling refreshed after a good night's sleep. You look around your bedroom and notice that everything seems normal. There isn't anyone else around except for your wife, who lies sleeping next to you. You turn towards her and smile softly. She stirs slightly and mumbles something incoherent. You stroke her hair gently and whisper sweet nothings into her ear. Soon she falls back asleep and you decide to leave her alone.
> 
> As soon as you step outside your house, you hear voices coming from behind you. Turning quickly, you spot three young boys running down the street toward you. They stop suddenly and stare at you wide eyed. "Are you the man who killed the giants?" one asks nervously.
> 
> You nod slowly. "Yes," you reply.
> 
> "Then I guess we should thank you." Another boy adds.
> 
> "Yeah," says the third one. "We heard that giants ate babies."
> 
> You shake your head sadly. "No," you explain. "Giants don't eat babies. That's just a myth."
> 
> "Oh really?" the first boy says skeptically. "Well then why did my parents move away from their village?"
> 
> "I dunno," you shrug. "Maybe because they believed the same thing you did."
> 
> The three kids look surprised. Then they laugh heartily together. "Haha!" they giggle. "That's funny!"
> 
> You grin back at them. "Thanks."
> 
> They run off laughing hysterically. You chuckle quietly to yourself. Kids are always full of surprises.
> 
> After spending some quality time with your family, you return home early in the afternoon. You grab a cup of tea and sit down at the kitchen table to read the newspaper. As usual, most of the stories involve wars somewhere else in the world. There aren't any articles related to your own conflict yet though. Maybe tomorrow...
> 
> Suddenly, you hear loud knocking on the front door. You stand up abruptly and rush over to answer it. Standing before you is an old man wearing a long beard and holding a staff in his hand. He looks familiar somehow.
> 
> "Hello sir," he greets politely. "My name is Galanis. I am here representing the Academy of Magic Studies."
> 
> "Ah yes," you respond. "I remember seeing advertisements for your school everywhere recently. What can I do for you?"
> 
> "Well sir," he begins, "we have received reports that you possess certain abilities which fall under our area of expertise. Specifically, we understand that you are capable of controlling fire using mental telepathy. Is this true?"
> 
> You nod slowly. "It certainly is."
> 
> "Excellent," he smiles happily. "In light of these circumstances, I would like to offer you a position teaching at our academy. We require instructors skilled in various aspects of magic including conjuration, enchantment, necromancy, divination, alchemy, and astrology."
> 
> You shake your head firmly. "I appreciate the offer Mr. Galanis but I must decline. My duties lie elsewhere right now."
> 
> His face falls momentarily. "Very well then," he sighs dejectedly. "Perhaps another opportunity will arise later on."
> 
> With that said, he turns around and leaves without saying another word. You watch him depart silently wondering whether you should have accepted his proposal. On the one hand, it might provide some distraction from the war. On the other hand, you doubt you could ever fit into such a structured environment. Besides, you wouldn't want to spend years away from home either. In the end, you decide to stay put and continue doing what you've always done best - fighting for justice.
> 
> \> You go to the great library
> You enter the Great Library of Alexandria hoping to gain knowledge about the Erenlander religion. Unfortunately, you find out that nobody there knows anything about it. The closest thing resembling a book on the subject is a small pamphlet titled "A Brief Overview Of Religion". According to its author, the writer claims that religions differ mainly due to differences between societies. Each society creates its own beliefs based upon its unique experiences. Since no two cultures are alike, there cannot be any universal truth regarding divine matters. Therefore, there is no reason to expect one religion to be superior over another.
> 
> You think deeply about this concept. Does it apply to the Erenlander faith? Are their beliefs simply an expression of local customs rather than absolute truths? Or does their doctrine truly reflect reality? These thoughts weigh heavily upon your mind. Suddenly, you hear someone calling your name. Turning around, you spot a librarian approaching you.
> 
> "Mr. Thane," he says cheerfully. "Can I help you find something specific?"
> 
> "Actually yes," you reply. "Do you happen to know anything about the Erenlander religion?"
> 
> He scratches his chin thoughtfully. "Not really," he admits. "But let me show you something anyway." He leads you over to a large table covered with books. Picking one up randomly, he starts reading aloud from it.
> 
> "According to ancient legends, Galdrin created the universe out of nothingness. Before that, nothing existed except pure chaos. Out of this void came forth life and light. With these gifts, Galdrin blessed humanity with wisdom and understanding. Through hard work and determination, humankind learned how to cultivate crops, build shelters, manufacture weapons, develop technology, establish governments, maintain order, and defend itself against enemies. By following the teachings of Galdrin, humans were able to rise above mere animals and achieve greatness beyond compare. Those who rejected his teachings fell victim to ignorance and decay. Eventually, their civilization collapsed leaving only ruins behind. Today, descendants of those original settlers still live amongst us. Although their culture has changed considerably since antiquity, they retain many traditions common among modern-day peoples. Their language remains virtually unchanged despite centuries of contact with foreign languages. Their art forms range widely ranging from poetry to sculpture to music. Their religion continues to flourish despite attempts to suppress it throughout history. And most importantly, their belief system continues to guide them along the path of righteousness."
> 
> You listen intently as he speaks. His words seem familiar somehow. Has he told you this story before? Or perhaps it was passed down through generations within your family? Either way, you find yourself drawn into the narrative. Before long, you forget everything else happening around you. Everything fades away except for the sound of his voice echoing throughout your ears.
> 
> "So why did Galdrin create the world?" the librarian finishes. "Because he loved us," he answers softly. "And because he knew we needed somewhere safe to call home."
> 
> You nod slowly. "Thank you," you murmur. "That was quite interesting."
> 
> "Glad to hear it," he grins broadly. "If you ever need anything else, please feel free to ask."
> 
> You bid farewell to the librarian and walk back home. During the journey, you ponder over what you learned earlier. Did Galdrin actually exist? Was he real or imaginary? Could he possibly be connected to the Erenlander faith somehow? So many questions remain unanswered. Still, you sense that something significant occurred tonight. Something special transpired inside the library walls. Now, you must figure out exactly what happened.


# Some quick and  dirty training details:

- [<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="150" height="24"/>](https://github.com/OpenAccess-AI-Collective/axolotl)
- Sequence length: 16384
- \# of epochs: 3
- Training time: 5 hours
- Hardware: 4x RTX 4090
- Training type: QLoRA
- PEFT R/A: 32/32


# Credits:
### Skein Text Adventure Data:

Thank you to the [Kobold AI](https://huggingface.co/KoboldAI) community for curating the Skein dataset, which is pivotal to this model's capabilities.