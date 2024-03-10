---
tags:
- not-for-all-audiences
license: apache-2.0
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/64dd7cda3d6b954bf7cdd922/ZXmxNKGaHUrqjdS1I3GkL.png)


## Thespis v0.7
This model works best with internet style RP using standard markup with asterisks surrounding actions and no quotes around dialogue.
Current work is looking at creating a custom DPO dataset as well as expanding the training data to include more niche interests.

External Datasets Used:

* Pure-Dove Dataset
* Claude Multiround 30k
* OpenOrcaSlim
* Augmental Dataset


DPO was done using a few generic datasets available on Hugging Face.

DPO Data:

* Intel/orca_dpo_pairs
* NobodyExistsOnTheInternet/ToxicDPOqa



Works with standard chat format for Ooba or SillyTavern.

## Prompt Format: Chat ( The default Ooba template and Silly Tavern Template )
```
{System Prompt}

Username: {Input}
BotName: {Response}
Username: {Input}
BotName: {Response}

```
## Ooba ( Set it to Chat, select a character and go. )
![image/png](https://cdn-uploads.huggingface.co/production/uploads/64dd7cda3d6b954bf7cdd922/HTl7QlAZcqe2hV8rwh4DG.png)

## Silly Tavern Settings ( Default )
![image/png](https://cdn-uploads.huggingface.co/production/uploads/64dd7cda3d6b954bf7cdd922/ajny8P0LdW0nFtghpPbfB.png)
```