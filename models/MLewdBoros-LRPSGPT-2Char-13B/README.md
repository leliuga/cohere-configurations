---
license: cc-by-nc-4.0
tags:
- not-for-all-audiences
- nsfw
---
![image/png](https://cdn-uploads.huggingface.co/production/uploads/63ab1241ad514ca8d1430003/DKLTsIPoJSfs8okxVCLiw.png)

This model is based on MLewdBoros merged with the ShareGPT-13b-qloras for "2 character".
<!-- description start -->
## Description

This repo contains fp16 files of MLewdBoros-LRPSGPT-2Char-13B, and is made to be used with character card containing "TWO PERSONAS".
<!-- description end -->
<!-- description start -->
## LoRA used

https://huggingface.co/royallab/LimaRP-ShareGPT-13b-qloras/tree/main/prompt-a/twochar
<!-- description end -->
<!-- prompt-template start -->
## Prompt template: Custom

```
Enter roleplay mode. You are currently %{having a conversation|in conversation|in a roleplay chat} with <SECOND>, whose %{traits are|persona is|characteristics are}:
<SECOND PERSONA>
%{You are|Play the role of|Take the role of} <FIRST> with the following %{persona|definitions|character sheet|traits}:
<FIRST PERSONA>
%{In addition|Additionally|Also}, %{keep the following scenario in mind|remember this scenario|pay attention to this scenario}:
<SCENARIO>
```
Or try to use Chat without instruction.

More info: https://huggingface.co/royallab/LimaRP-ShareGPT-13b-qloras/blob/main/prompt-a/README.md

Special thanks to Sushi ♥