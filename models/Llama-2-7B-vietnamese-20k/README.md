---
tags:
- text-generation
- llama-2
- llama-2-7B
- llama2-vietnamese
- vietnamese
---
# Model Card for Llama 2 Fine-Tuned on Vietnamese Instructions

## Model Details
- Model Name: Llama-2-7b-vietnamese-20k
- Architecture: Llama 2 7B
- Fine-tuning Data Size: 20,000 instruction samples
- Purpose: To demonstrate the performance of the Llama 2 model on Vietnamese and gather initial insights. A more comprehensive model and evaluation will be released soon.
- Availability: The model checkpoint can be accessed on Hugging Face: ngoantech/Llama-2-7b-vietnamese-20k

## Intended Use
This model is intended for researchers, developers, and enthusiasts who are interested in understanding the performance of the Llama 2 model on Vietnamese. It can be used for generating Vietnamese text based on given instructions or for any other task that requires a Vietnamese language model.

## Example Output 
![Example output 1](exp_1.png "Example output 1")


## Limitations
- Data Size: The model was fine-tuned on a relatively small dataset of 20,000 instruction samples, which might not capture the full complexity and nuances of the Vietnamese language.
- Preliminary Model: This is an initial experiment with the Llama 2 architecture on Vietnamese. More refined versions and evaluations will be available soon.
- Performance:
Specific performance metrics on this fine-tuned model will be provided in the upcoming comprehensive evaluation.

## Ethical Considerations
- Bias and Fairness: Like any other machine learning model, there is a possibility that this model might reproduce or amplify biases present in the training data.
- Use in Critical Systems: As this is a preliminary model, it is recommended not to use it for mission-critical applications without proper validation.
- Fine-tuning Data:
The model was fine-tuned on a custom dataset of 20,000 instruction samples in Vietnamese. More details about the composition and source of this dataset will be provided in the detailed evaluation report.



## Credits
I would like to express our gratitude to the creators of the Llama 2 architecture and the Hugging Face community for their tools and resources.

## Contact
ngoantech@gmail.com

https://github.com/ngoanpv