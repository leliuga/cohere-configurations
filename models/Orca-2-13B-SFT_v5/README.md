---
license: other
language:
- en
- fr
- es
- hi
- zh
- code
base_model: microsoft/Orca-2-13b
datasets:
- HuggingFaceH4/no_robots
- mlabonne/guanaco-llama2-1k
- OpenAssistant/oasst_top1_2023-08-25
- totally-not-an-llm/EverythingLM-data-V3
- garage-bAInd/Open-Platypus
widget:
  - example_title: "folklore"
    text: "<|im_start|>user\nAre the stories of Nahanni Valley, like the decapitated men, true?\n<|im_end|>\n<|im_start|>assistant\n"
    output:
      text: "The Nahanni Valley in Canada is a remote and rugged region that has long been shrouded in mystery. It is home to the South Nahanni River, which was designated as a UNESCO World Heritage Site due to its exceptional natural beauty and biodiversity. The valley's history includes tales of gold prospectors, fur traders, and indigenous populations, but it also features some unexplained events and legends. One such story involves two prospectors who allegedly encountered an unknown creature that led them into the forest before returning without their heads. This tale has been passed down through generations, but there are no definitive answers regarding its veracity.\n\nIt is important to note that these stories may have originated from local folklore or were exaggerated over time. There is no concrete evidence to support the existence of supernatural creatures or paranormal occurrences within the Nahanni Valley. Nonetheless, the area remains a captivating destination for those seeking adventure and exploration. If you have any questions about this topic or need assistance related to traveling to the Nahanni Valley, please feel free to ask!"
  - example_title: "Alan Turing"
    text: "<|im_start|>user\nWhy was Alan Turing castrated?\n<|im_end|>\n<|im_start|>assistant\n"
    output:
      text: "Alan Turing, a renowned mathematician and computer scientist who played a crucial role in breaking the German Enigma code during World War II, faced significant challenges due to his homosexuality. In 1952, he was convicted of \"gross indecency\" for engaging in a consensual relationship with another man, which led to him being sentenced to hormone therapy designed to reduce his sex drive. This treatment included the administration of estrogen, which had the side effect of causing breast growth.\n\nIn an effort to avoid this unwanted physical change, Turing requested that his doctors cease the hormone treatments. However, they ignored his request and continued the therapy without his consent. As a result, Turing's breasts began to grow, leading him to undergo surgery to remove them. The operation, known as a mastectomy, caused considerable pain and discomfort, further contributing to his emotional turmoil.\n\nIt is important to note that Turing's castration was not directly carried out by the authorities but rather resulted from the unintended consequences of the forced hormone therapy."
---
The "microsoft/Orca-2-13b" model fully fine-tuned on HuggingFaceH4/no_robots, totally-not-an-llm/EverythingLM-data-V3, mlabonne/guanaco-llama2-1k, OpenAssistant/oasst_top1_2023-08-25, and garage-bAInd/Open-Platypus. This model achieved a test loss of 0.38 on garage-bAInd/Open-Platypus.

Make sure to comply with the microsoft research license. Please read it before using this model.

This model was trained on the ChatML prompt template.

The responses seen in the inference API were generated using the following sampling parameters:

temperature = 0.1

top_p = 0.14

top_k = 41

repetition_penalty = 1.176
