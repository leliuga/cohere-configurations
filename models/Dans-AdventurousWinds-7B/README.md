---
language:
- en
datasets:
- PocketDoc/Floyd-Text-Adventures
- PocketDoc/Choose-Your-Story-Long-Text-Adventures
license: apache-2.0
base_model: mistralai/Mistral-7B-v0.1
---
# What is the model for?
This model is proficient in crafting text-based adventure games. It can both concise replies and more expansive, novel-like descriptions. The ability to alternate between these two response styles can be triggered by a distinct system message.

# What's in the sausage?

This model was trained on [Mistral-7b](https://huggingface.co/mistralai/Mistral-7B-v0.1) using a custom version of the floyd and cys datasets from the skein text adventure dataset. The dataset has been improved over what was used in my previous model "Dans-CreepingSenseOfDoom" and this time presented in the more traditional adventure mode format.

## Prompt format and examples:
#### Format description:
Traditional adventure format with a description of the genre, tags, and other misc information at the top followed by an optional introductory message. User messages are prefixed with two new lines and "> ", model messages are one new line after the last user message or the description message if an introductory message is desired.

#### Examples of shorter responses:
```
[Genres: Science Fiction]
[Tags: humor, old school, sci fi]
[Mode: Adventure]
[Description: A puzzle about committing acts of financial skulduggery and exploiting ridiculous magical items.]
[Misc: Writing era: 1993]
[Intro]
It is the year 2045. You are a young man in his twenties living in New York City. Your father was an inventor who died when you were very small; your mother raised you alone for many years until she remarried. Now you live with your stepfather, but he doesn't care much for you and has never given you any money to help support yourself. You have no job and little hope of getting one because of your lack of experience. However, you do have some unusual abilities that could be put to good use if only you knew how...

> Look around
You are standing on the sidewalk outside your apartment building. The street is busy with traffic and people going about their daily business. Across the street from you stands a tall office tower where your stepfather works as an accountant. To your left is a convenience store where you can buy food and drinks.

> You go into the store
The doorbell rings as you enter the store. An elderly woman behind the counter looks up at you expectantly. "Hello," she says cheerfully, "what can I get for you?"

> You tell her you'll browse
"All right then," she replies brightly, "just take your time."

> You ask if she's hiring
"I'm sorry," she says sadly, "but we don't need anyone new just now."

> You leave without buying anything
As you walk out of the store, the bell jingles merrily.

> You go see your stepfather
When you arrive at work, your stepfather greets you warmly. He invites you into his office and offers you a seat. "So," he begins eagerly, "how would you like to make some extra cash? There's this guy named Mr. Jones who owes my company thousands of dollars. If we could somehow convince him to pay us back, it would really help our bottom line."

He pauses expectantly, waiting for your response.

> You ask him whats involved
"Well," he continues excitedly, "it turns out that Mr. Jones has a very valuable antique watch that he keeps locked away in a safe deposit box at the bank. We think that if we could somehow break into the box and steal the watch, we could sell it for enough money to cover his debt."

Your heart races as you realize the potential danger of such a scheme. But before you can say anything, your stepfather cuts off your thoughts by continuing, "Of course, there's always the chance that something could go wrong and we might end up losing everything..."
```
#### Examples of longer novel like responses:
```
[Genres: grimdark fantasy] [Tags: Fantasy, Grimdark, Serious, Villain Protagonist, War] [Description: This story is definitely more "story" than "game", however there are two distinct paths that can be followed, each with their own pitfalls and additional choices. This story can get long, and assuming you're surviving, years do pass in it and you and your environment will change.] [First message intro] [Mode: Interactive Storyteller]
You have been a soldier for as long as you remember. You don't know who you were before the war began, but you were someone else. Someone less angry. More forgiving. Perhaps even kinder.

It all started when you were very young. The King of this land was attacked by an army from another kingdom. His forces were outmatched, his city overrun. He had no choice but to flee into the forest, taking what few soldiers he could muster along with him.

As the king fled deeper into the woods, his army fell back, forming a protective wall around them. They fought valiantly against the invaders, giving the king time to escape, but they eventually succumbed to overwhelming numbers. When the fighting finally ended, only one man remained alive; you.

The king found you hiding behind some bushes, covered in blood and dirt. He took pity on you and offered you a place at court if you would swear allegiance to him. Without hesitation, you accepted. And so began your journey down the path of darkness...

> You meet with the king to discuss the plans
A week has passed since the battle and you have just met with the king to discuss future plans. It seems like things are finally starting to settle down after such chaos. However, there is still much work to be done before we can truly consider ourselves safe again.

The first order of business is rebuilding our army. We need men willing to fight for us once more. Men who understand what needs to be done and aren't afraid to die trying. I believe I have found such warriors among those loyal subjects who survived the initial attack. If we train them properly, then perhaps they might prove useful against our enemies next time around.

Next, we must find ways to improve our defenses. There are many areas where improvements could be made. For instance, we should construct better walls around our cities. These walls must not only keep intruders out, but also allow us easy access inside should anyone try to breach them. In addition, we must strengthen our gates so that no matter how strong our enemy becomes, he cannot break through without great effort. Finally, we must fortify our castles and other strategic locations throughout the country. By doing these things, we will ensure that nobody ever attacks us again!

> You ask what he desires your involvement to be
After hearing about his plans, you feel compelled to help him achieve victory. After all, you owe everything to him. But before agreeing to anything, you want to make sure that he understands exactly what you expect from him. So you decide to approach him directly.

"Your Majesty," you begin, "I am honored beyond measure to serve under your command. However, before committing myself fully to this cause, I wish to clarify certain matters."

He nods slowly. "Very well. What would you like me to explain?"

"Well," you say, "first off, I want to know why you chose me above everyone else? Why did you pick me specifically?"

He smiles softly. "Because I saw potential within you. Potential which others lacked. Potentials which could prove valuable during times of crisis. Therefore, I felt confident enough to entrust my fate to you."

"But what makes you think I possess such abilities?" you ask curiously.

"Oh, nothing special really," he replies casually. "Just raw determination combined with intelligence and resourcefulness. Those qualities alone are enough to carry us forward toward victory!"
```

# Some quick and  dirty training details:
- [<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="150" height="24"/>](https://github.com/OpenAccess-AI-Collective/axolotl)
- Sequence length: 4096
- \# of epochs: 3
- Training time: 1 hour
- Hardware: 1x RTX 3090
- Training type: QLoRA
- PEFT R/A: 32/32


# Credits:
### Skein Text Adventure Data:
Thank you to the [Kobold AI](https://huggingface.co/KoboldAI) community for curating the Skein dataset, which is pivotal to this model's capabilities.