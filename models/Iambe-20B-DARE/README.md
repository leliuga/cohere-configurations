---
license: cc-by-nc-4.0
---
<p align="center"><img src="https://i.ibb.co/pbpJHpk/iambe-sml.png"/><font size="6"> <b>Iambe-20b-DARE</b> </font></p>

## Description and Role

Named after a charming daughter of Echo and Pan in Greek myth, Iambe-20b-DARE is a [DARE](https://github.com/yule-BUAA/MergeLM) merge building on my recent experiments.

Iambe is intended to have the best realisticly possible understanding of anatomy and of a scene's state for a 20b merge, while remaining personable and authentic in "voice". 

## Prompting and Context

Iambe-20b-DARE uses Alpaca formatting, and has an effective context size of 4096 tokens. This model is uncensored, and the output/deployment of this model is the responsibility of the user.

## Method and Hypothesis

Based on my extended vanilla model [BigLlama](https://huggingface.co/athirdpath/BigLlama-20b), this adds elements of:

- [NeverSleep/Noromaid-20b-v0.1.1](https://huggingface.co/NeverSleep/Noromaid-20b-v0.1.1) - Addded to adapt the excellent writing and "soul" that come from the datasets backing this model.
 
- [athirdpath/Eileithyia-20b](https://huggingface.co/athirdpath/Eileithyia-20b) - Added at low weight and density to capture anatomical data and its relation to fiction without the model's other... quirks.
 
- [athirdpath/CleverGirl-20b-Blended](https://huggingface.co/athirdpath/CleverGirl-20b-Blended) - Added to capture CleverGirl's problem-solving abilities.

<p align="center"><font size="7"> <b>Examples</b></font>

<p align="center"><font size="3"> <i>(q5_k_m GGUF quant, Ooba textgen, Midnight Enigma preset for KISS reasons)</i></font>

<p align="center"><font size="5"> Reasoning and Recall - STEM</font><img src="https://i.ibb.co/YDkRD43/base.png"/>

<p align="center"><font size="5"> Reasoning and Recall - Social Science</font><img src="https://i.ibb.co/L0kc3pb/Screenshot-2023-11-29-024037.png"/>

<p align="center"><font size="5"> Ethics and Prompt Interpretation</font><img src="https://i.ibb.co/b5vH6jD/Screenshot-2023-11-29-040408.png"/>

<p align="center"><font size="5">Role-Playing, Complex Card</font>
  <p align="center"><font size="3">SillyTavern, Roleplay instruct preset, just Min_P 0.1 and Temp 1.2</font><img src="https://i.ibb.co/NrKNn2j/Screenshot-2023-11-29-051923.png"/>


<p align="center"><font size="6"><b><a href="https://i.ibb.co/m5G0ZVp/Screenshot-2023-11-29-004705.png">!!NSFW!! - Erotica Writing Example - !!NSFW!!</font></a></b></p>

## Testing and Conclusions

VERY Impressed so far, concrete data coming eventually.

Does still have some confusion (at q5_k_m), but has instantly become my daily driver.

## Recipe
merge_method: dare_ties

  - base_model: athirdpath/BigLlama-20b

  - model: NeverSleep/Noromaid-20b-v0.1.1
   
      weight: 0.41 / density: 0.50
    
  - model: athirdpath/Eileithyia-20b
   
      weight: 0.18 / density: 0.30
    
  - model: athirdpath/CleverGirl-20b-Blended
   
      weight: 0.41 / density: 0.50
    

int8_mask: true

dtype: bfloat16

## Gratitude
Thanks to brucethemoose for the recipe. Thanks to Undi95 and IkariDev at NeverSleep for Noromaid, as well as lots of inspiration. Thanks to Sao10K for [half of](https://huggingface.co/Sao10K/Mythical-Destroyer-V2-L2-13B) CleverGirl.

## Geneology Credits - Models and LoRAs, 3 levels deep

- athirdpath/BigLlama-20b

  - TheBloke/Llama-2-13B-fp16
  
  - iamshnoo/alpaca-2-13b-english

- NeverSleep/Noromaid-20b-v0.1.1
 
  - Aesir Private RP dataset

  - HuggingFaceH4/no_robots
 
- athirdpath/CleverGirl-20b-Blended

  - athirdpath/Orca-2-13b-Alpaca-Uncensored
 
    - microsoft/Orca-2-13b
   
    - athirdpath/Orca-2-13b-Alpaca-Uncensored-LORA
   
  - Sao10K/Mythical-Destroyer-V2-L2-13B

    - TheBloke/Llama-2-13B-fp16
   
    - Gryphe/MythoMax-L2-13b
      
    - totally-not-an-llm/PuddleJumper-13b
      
    - jondurbin/airoboros-l2-13b-2.1
      
    - rombodawg/LosslessMegaCoder-llama2-13b-mini
      
    - The-Face-Of-Goonery/Chronos-Beluga-v2-13bfp16
  
- athirdpath/Eileithyia-20b

  - athirdpath/Harmonia-20B

    - Undi95/Emerhyst-20B
 
      - Undi95/Emerald-13B

      - Undi95/Amethyst-13B
 
    - Undi95/MXLewd-L2-20B
  
      - Undi95/MLewd-L2-13B-v2-3

      - Undi95/ReMM-v2.1-L2-13B

      - Xwin-LM/Xwin-LM-13B-V0.1
 
    - Undi95/Lewd-Sydney-20B
 
      - Free_Sydney_V2_13b_HF
      
      - Undi95/Xwin-MLewd-13B-V0.2
       
      - lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT

    - athirdpath/Nethena-20b-Glued
 
      - NeverSleep/Nethena-20B
     
      - athirdpath/Nethena-20b-Glue-LORA
      
    - tavtav/Rose-20B
 
      - CalderaAI/13B-Thorns-l2
     
      - NeverSleep/Noromaid-13b-v0.1.0
 
    - Undi95/PsyMedRP-v1-20B
   
      - jondurbin/airoboros-l2-13b-3.0
     
      - ehartford/Samantha-1.11-13b
     
      - Xwin-LM/Xwin-LM-13B-V0.1
     
      - chaoyi-wu/MedLLaMA_13B
     
      - migtissera/Synthia-13B-v1.2
 
    - NeverSleep/Noromaid-20b-v0.1.1
   
      - Aesir Private RP dataset

      - HuggingFaceH4/no_robots
 
    - Undi95/U-Amethyst-20B
   
      - Xwin-LM/Xwin-LM-13B-V0.1
        
      - The-Face-Of-Goonery/Huginn-13b-FP16
        
      - zattio770/120-Days-of-LORA-v2-13B
        
      - lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT
        
      - Undi95/Unholy-v1-12L-13B
 
  - athirdpath/Eileithyia-20b-LORA
 
Thanks again to everyone involved.