---
license: llama2
---

# Chronomaid-Storytelling-13b

<img 
  src="https://cdn-uploads.huggingface.co/production/uploads/65221315578e7da0d74f73d8/v2fVXhCcOdvOdjTrd9dY0.webp"
  alt="a vibrant and whimsical scene with an anime-style character as the focal point. The character is a young girl with blue eyes and short brown hair, wearing a black and white maid outfit with ruffled apron and a red ribbon at her collar. She is lying amidst a fantastical backdrop filled with an assortment of floating, colorful clocks, gears, and hourglasses. The space around her is filled with sparkling stars, glowing nebulae, and swirling galaxies."
  title="a vibrant and whimsical scene with an anime-style character as the focal point. The character is a young girl with blue eyes and short brown hair, wearing a black and white maid outfit with ruffled apron and a red ribbon at her collar. She is lying amidst a fantastical backdrop filled with an assortment of floating, colorful clocks, gears, and hourglasses. The space around her is filled with sparkling stars, glowing nebulae, and swirling galaxies."
  height="75%"
  width="75%"
/>

Merge including [Noromaid-13b-v0.1.1](https://huggingface.co/NeverSleep/Noromaid-13b-v0.1.1), and [Chronos-13b-v2](https://huggingface.co/elinas/chronos-13b-v2) with the [Storytelling-v1-Lora](https://huggingface.co/Undi95/Storytelling-v1-13B-lora) applied afterwards

Inteded for primarily RP, and will do ERP, narrator-character and group-chats without much trouble in my testing.

## Prompt Format

Tested with Alpaca, the Noromaid preset's will probably also work (check the Noromaid model card for SillyTavern presets)
```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:
```

## Sampler Settings

Tested at 
* `temp` 1.3 `min p` 0.05 and 0.15
* `temp` 1.7, `min p` 0.08 and 0.15

## Quantized Models
The model has been kindly quantized in GGUF, AWQ, and GPTQ by TheBloke  
Find them in the [Chronomaid-Storytelling-13b Collection](https://huggingface.co/collections/NyxKrage/chronomaid-storytelling-13b-656115dd7065690d7f17c7c8)

## Thanks ❤️

To [Undi](https://huggingface.co/Undi95) & [Ikari](https://huggingface.co/IkariDev) for Noromaid and [Elinas](https://huggingface.co/elinas) for Chronos  
Support [Undi](https://ko-fi.com/undiai) and [Elinas](https://ko-fi.com/elinas) on Kofi
