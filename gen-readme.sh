#!/bin/bash

MODELS=$(ls models | sort)
COUNT=$(ls models | wc -l)

echo "LLM Inference configurations👋" > README.md
echo "===============================" >> README.md
echo "" >> README.md

echo "## Description" >> README.md
echo "" >> README.md
echo "This repository contains the inference configurations for the Co:Here" >> README.md

echo "" >> README.md
echo "### Supported Transformer Architectures by Co:Here" >> README.md

echo "" >> README.md
echo "- [X] LLaMA 🦙" >> README.md
echo "- [x] LLaMA 2 🦙🦙" >> README.md
echo "- [X] Falcon" >> README.md
echo "- [X] Alpaca" >> README.md
echo "- [X] GPT4All" >> README.md
echo "- [X] Chinese LLaMA / Alpaca and Chinese LLaMA-2 / Alpaca-2" >> README.md
echo "- [X] Vigogne (French)" >> README.md
echo "- [X] Vicuna" >> README.md
echo "- [X] Koala" >> README.md
echo "- [X] OpenBuddy 🐶 (Multilingual)" >> README.md
echo "- [X] Pygmalion/Metharme" >> README.md
echo "- [X] WizardLM" >> README.md
echo "- [X] Baichuan 1 & 2 + derivations" >> README.md
echo "- [X] Aquila 1 & 2" >> README.md
echo "- [X] Starcoder models" >> README.md
echo "- [X] Mistral AI v0.1" >> README.md
echo "- [X] Refact" >> README.md
echo "- [X] Persimmon 8B" >> README.md
echo "- [X] MPT" >> README.md
echo "- [X] Bloom" >> README.md
echo "- [X] StableLM-3b-4e1t" >> README.md
echo "" >> README.md

echo "### Supported backends" >> README.md
echo "- [X] llama-backend" >> README.md
echo "- [ ] bert-backend" >> README.md
echo "- [ ] bart-backend" >> README.md
echo "- [ ] t5-backend" >> README.md
echo "- [ ] rwkv-backend" >> README.md
echo "" >> README.md

echo "### Supported platforms" >> README.md
echo "- [X] linux/amd64" >> README.md
echo "- [ ] darwin/amd64" >> README.md
echo "" >> README.md

echo "## Requirements" >> README.md
echo "" >> README.md
echo "1. [Docker](https://docs.docker.com/get-docker/)" >> README.md
echo "2. [NVIDIA Container Toolkit](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/install-guide.html)" >> README.md
echo "" >> README.md

echo "## Usage" >> README.md
echo "" >> README.md
echo '``` bash' >> README.md
echo "docker run -it --rm --gpus all -v ./models/:/app/models -p 3000:3000 github.com/leliuga/cohere cohere run <name>/<dtype>" >> README.md
echo '```' >> README.md
echo "" >> README.md

echo "## Supported Models (${COUNT})" >> README.md
echo "" >> README.md
echo "| Name  | DType   | Read more |" >> README.md
echo "| ----- | ------- | --------- |" >> README.md
for model in $MODELS; do
    echo -ne "| ${model} |" >> README.md

    for dtype in $(find models/${model} -name "*.yaml" | sed 's/\.yaml$//' | sort | cut -d'/' -f3); do
        echo -ne "${dtype} " >> README.md
    done

    echo -ne "| [README.md](models/${model}/README.md) |\n" >> README.md
done

echo "" >> README.md
echo "## Memory/Disk Requirements" >> README.md
echo "As the models are currently fully loaded into memory, you will need adequate disk space to save them and sufficient RAM to load them. At the moment, memory and disk requirements are the same." >> README.md
echo "" >> README.md
echo "| Model | Original size | Quantized size (4-bit) |" >> README.md
echo "|------:|--------------:|-----------------------:|" >> README.md
echo "|    7B |         13 GB |                 3.9 GB |" >> README.md
echo "|   13B |         24 GB |                 7.8 GB |" >> README.md
echo "|   30B |         60 GB |                19.5 GB |" >> README.md
echo "|   65B |        120 GB |                38.5 GB |" >> README.md

echo "" >> README.md
echo "## Quantization" >> README.md
echo "Several quantization methods are supported. They differ in the resulting model disk size and inference speed." >> README.md
echo "*(outdated)*" >> README.md
echo "" >> README.md
echo "| Model | Measure      | F16    | Q4_0   | Q4_1   | Q5_0   | Q5_1   | Q8_0   |" >> README.md
echo "|------ |--------------|--------|--------|--------|--------|--------|--------|" >> README.md
echo "|    7B | perplexity   | 5.9066 | 6.1565 | 6.0912 | 5.9862 | 5.9481 | 5.9070 |" >> README.md
echo "|    7B | file size    |  13.0G |   3.5G |   3.9G |   4.3G |   4.7G |   6.7G |" >> README.md
echo "|    7B | ms/tok @ 4th |    127 |     55 |     54 |     76 |     83 |     72 |" >> README.md
echo "|    7B | ms/tok @ 8th |    122 |     43 |     45 |     52 |     56 |     67 |" >> README.md
echo "|    7B | bits/weight  |   16.0 |    4.5 |    5.0 |    5.5 |    6.0 |    8.5 |" >> README.md
echo "|   13B | perplexity   | 5.2543 | 5.3860 | 5.3608 | 5.2856 | 5.2706 | 5.2548 |" >> README.md
echo "|   13B | file size    |  25.0G |   6.8G |   7.6G |   8.3G |   9.1G |    13G |" >> README.md
echo "|   13B | ms/tok @ 4th |      - |    103 |    105 |    148 |    160 |    131 |" >> README.md
echo "|   13B | ms/tok @ 8th |      - |     73 |     82 |     98 |    105 |    128 |" >> README.md
echo "|   13B | bits/weight  |   16.0 |    4.5 |    5.0 |    5.5 |    6.0 |    8.5 |" >> README.md
echo "" >> README.md

echo "## License" >> README.md
echo "" >> README.md
echo "This project is licensed under the Mozilla Public License Version 2.0 License - see the [LICENSE](LICENSE) file for details." >> README.md
