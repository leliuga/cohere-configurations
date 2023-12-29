package main

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"golang.org/x/exp/maps"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

var (
	tmpl = template.Must(template.New("readme").Parse(`
LLM Inference configurations👋
===============================

## Description

This repository contains the inference configurations for the Co:Here

### Supported Transformer Architectures by Co:Here

- [X] LLaMA 🦙
- [x] LLaMA 2 🦙🦙
- [X] Falcon
- [X] Alpaca
- [X] GPT-2
- [X] GPT4All
- [X] Chinese LLaMA / Alpaca and Chinese LLaMA-2 / Alpaca-2
- [X] Vigogne (French)
- [X] Vicuna
- [X] Koala
- [X] OpenBuddy 🐶 (Multilingual)
- [X] Pygmalion/Metharme
- [X] WizardLM
- [X] Baichuan 1 & 2 + derivations
- [X] Aquila 1 & 2
- [X] Starcoder
- [X] Mistral AI v0.1 & v0.2
- [X] Refact
- [X] Persimmon 8B
- [X] MPT
- [X] Bloom
- [X] StableLM-3b-4e1t
- [x] Yi
- [x] Deepseek
- [x] Qwen
- [x] Mixtral MoE
- [x] Phi 2
- [x] PLaMo 13B

Multimodal

- [x] Llava 1.5
- [x] Bakllava
- [x] Obsidian
- [x] ShareGPT4V

### Supported backends
- [X] llama-backend
- [ ] bert-backend
- [ ] bart-backend
- [ ] t5-backend
- [ ] rwkv-backend

### Supported platforms
- [X] linux/amd64
- [ ] darwin/amd64

## Requirements

1. [Docker](https://docs.docker.com/get-docker/)
2. [NVIDIA Container Toolkit](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/install-guide.html)

## Usage

` + "``` bash" + `
docker run -it --rm --gpus all -v ./models/:/app/models -p 3000:3000 github.com/leliuga/cohere cohere run <name>/<variant>
` + "```" + `

## Supported Models ({{len .}})

| Name  | DType   | Context | Embedding | Read more |
| ----- | ------- | ------- | --------- | --------- |{{range .}}
| {{.Name}} | {{.Dtype}} | {{.ContextSize}} | {{.EmbeddingSize}} | [README.md](models/{{.ID}}/README.md) |{{end}}

## Memory/Disk Requirements
As the models are currently fully loaded into memory, you will need adequate disk space to save them and sufficient RAM to load them. At the moment, memory and disk requirements are the same.

| Model | Original size | Quantized size (4-bit) |
|------:|--------------:|-----------------------:|
|    7B |         13 GB |                 3.9 GB |
|   13B |         24 GB |                 7.8 GB |
|   30B |         60 GB |                19.5 GB |
|   65B |        120 GB |                38.5 GB |

## Quantization
Several quantization methods are supported. They differ in the resulting model disk size and inference speed.
*(outdated)*

| Model | Measure      | F16    | Q4_0   | Q4_1   | Q5_0   | Q5_1   | Q8_0   |
|------ |--------------|--------|--------|--------|--------|--------|--------|
|    7B | perplexity   | 5.9066 | 6.1565 | 6.0912 | 5.9862 | 5.9481 | 5.9070 |
|    7B | file size    |  13.0G |   3.5G |   3.9G |   4.3G |   4.7G |   6.7G |
|    7B | ms/tok @ 4th |    127 |     55 |     54 |     76 |     83 |     72 |
|    7B | ms/tok @ 8th |    122 |     43 |     45 |     52 |     56 |     67 |
|    7B | bits/weight  |   16.0 |    4.5 |    5.0 |    5.5 |    6.0 |    8.5 |
|   13B | perplexity   | 5.2543 | 5.3860 | 5.3608 | 5.2856 | 5.2706 | 5.2548 |
|   13B | file size    |  25.0G |   6.8G |   7.6G |   8.3G |   9.1G |    13G |
|   13B | ms/tok @ 4th |      - |    103 |    105 |    148 |    160 |    131 |
|   13B | ms/tok @ 8th |      - |     73 |     82 |     98 |    105 |    128 |
|   13B | bits/weight  |   16.0 |    4.5 |    5.0 |    5.5 |    6.0 |    8.5 |

## License

This project is licensed under the Mozilla Public License Version 2.0 License - see the [LICENSE](LICENSE) file for details.
`))
)

type (
	Model struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Config struct {
			ContextSize   uint64 `json:"context_size"`
			EmbeddingSize uint64 `json:"embedding_size"`
		} `json:"config"`
		Variants map[string]struct {
			DType string `json:"dtype"`
		} `json:"variants"`
	}

	Item struct {
		ID            string
		Name          string
		Dtype         string
		ContextSize   string
		EmbeddingSize string
	}
)

func main() {
	var (
		modelsPath = "./models"
		items      []Item
	)

	if err := filepath.Walk(modelsPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name() == "manifest.yaml" {
			var (
				model Model
				file  *os.File
			)

			if file, err = os.Open(path); err != nil {
				return err
			}
			defer file.Close()

			encoder := yaml.NewDecoder(file, yaml.UseJSONUnmarshaler())
			if err = encoder.Decode(&model); err != nil {
				return err
			}

			items = append(items, Item{
				ID:            model.ID,
				Name:          model.Name,
				Dtype:         strings.Join(maps.Keys(model.Variants), " "),
				ContextSize:   strconv.FormatUint(model.Config.ContextSize, 10),
				EmbeddingSize: strconv.FormatUint(model.Config.EmbeddingSize, 10),
			})
		}

		return nil
	}); err != nil {
		panic(fmt.Errorf("error walking the path %v: %v\n", modelsPath, err))
	}

	f, err := os.Create("README.md")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err = tmpl.Execute(f, items); err != nil {
		panic(err)
	}
}
