---
license: cc-by-nc-4.0
---
This is [NeverSleep/Nethena-20B](https://huggingface.co/NeverSleep/Nethena-20B) with [athirdpath/Nethena-20b-Glue-LORA](https://huggingface.co/athirdpath/Nethena-20b-Glue-LORA) applied.

athirdpath/Nethena-20b-Glue-LORA is a 128 rank LORA for RP, trained on a private dataset. It is unalligned and NSFW-oriented.

This is a test, exploring the effects of "gluing" the components of the 20b model together to reduce the iconic word replacement errors, increase lucidity, and improve recall.

![image/png](https://huggingface.co/athirdpath/Nethena-20b-Glued/resolve/main/b5787896-afd5-44a3-b757-0e75ee28bed8.png)

The private ~500k token dataset used to train the LORA was Alpaca formatted and focused on 4 primary categories:

    - Medical texts (on psychology, reproductive organs, anatomy, and pregnancy). These are formatted so the model, in character as a doctor or therapist, answers a patient's question in short to medium form.
    - Excerpts from short stories and novellas (erotic and romantic) centered around both realistic and fantastic situations, covering several fetishes as well. These are sliced into ~2048 token chunks, and these long-form responses are all tied to the command “Enter narrator mode.” in the instructions.
    - A selection from PIPPA, using a wide keyword search for tokens associated with low quality human or AI data to remove those responses, then a positive search was done for words and phrases associated with a higher reading level. These are converted to Alpaca with “Enter RP mode.” in all the instruction fields.
    - ~18k tokens of GPT-4 generated data on role-playing from various characters’ perspectives, focusing on different situations and emotions. Includes many multi-turn conversations.

So far it is passing subjective testing with flying colors, objective numbers coming soon.

Trained with Alpaca-style prompts.