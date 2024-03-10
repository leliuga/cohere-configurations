package main

import (
	"github.com/goccy/go-yaml"
	"golang.org/x/exp/maps"
	"k8s.io/klog/v2"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

const (
	// MODELS_PATH is the path to the models directory
	MODELS_PATH = "models"
)

var (
	tmpl = template.Must(template.New("readme").Parse(`
LLM Inference configurations👋
===============================

## Description

This repository contains the inference configurations for the Co:Here

### Supported Transformer Architectures by Co:Here

- [X] Alpaca
- [X] Aquila 1 & 2
- [X] Baichuan 1 & 2 + derivations
- [X] Bert & Nomic-Bert
- [X] Bloom
- [X] Chinese LLaMA / Alpaca and Chinese LLaMA-2 / Alpaca-2
- [X] CodeShell
- [X] Deepseek
- [X] Falcon
- [X] GPT-2 && GPT-J && GPT-NeoX
- [X] Gemma
- [X] InternLM2
- [X] Koala
- [X] LLaMA 1 & 2
- [X] MPT
- [X] Mamba
- [X] MiniCPM
- [X] Mistral v0.1 & v0.2
- [X] Mixtral MoE
- [X] Orion 14B
- [X] OpenBuddy 🐶 (Multilingual)
- [X] Phi 1 & 1.5 & 2
- [X] Persimmon 8B
- [X] PLaMo 13B
- [X] Pygmalion/Metharme
- [X] Qwen 1 & 2
- [X] Refact
- [X] StableLM
- [X] Starcoder 1 & 2
- [X] Vigogne (French)
- [X] Vicuna
- [X] WizardLM
- [X] Yi

Multimodal

- [x] BakLLaVA
- [x] LLaVA 1.5 & 1.6
- [x] MobileVLM 1.7B & 3B
- [x] Obsidian
- [x] ShareGPT4V
- [x] Yi-VL

### Supported platforms

- [X] linux/amd64
- [ ] linux/arm64

## Requirements

1. [Nerdctl](https://github.com/containerd/nerdctl) or [Podman](https://podman.io/getting-started/installation) or [Docker](https://docs.docker.com/get-docker/)
3. [Using GPUs inside containers](https://github.com/containerd/nerdctl/blob/main/docs/gpu.md)

## Usage

` + "``` bash" + `
nerdctl run -d --name cohere --rm --gpus all -v ` + "`pwd`" + `/models/:/app/models -p 3000:3000 ghcr.io/leliuga/cohere run [<id>/<variant>, ...]

# example
nerdctl run -d --name cohere --rm --gpus all -v ` + "`pwd`" + `/models/:/app/models -p 3000:3000 ghcr.io/leliuga/cohere run Llama-2-7B-32K-Instruct/Q4_0
` + "```" + `

## Supported Models ({{len .}})

| ID      | Variants      | Vocab Size     | Context Size     | Embedding Size     | Read more                             |
| ------- | ------------- | -------------- | ---------------- | ------------------ | ------------------------------------- |{{range .}}
| {{.ID}} | {{.Variants}} | {{.VocabSize}} | {{.ContextSize}} | {{.EmbeddingSize}} | [README.md](models/{{.ID}}/README.md) |{{end}}

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
	// Model is a struct that contains the model configuration.
	Model struct {
		ID     string `json:"id"`
		Config struct {
			VocabSize     uint64 `json:"vocab_size"`
			ContextSize   uint64 `json:"context_size"`
			EmbeddingSize uint64 `json:"embedding_size"`
		} `json:"config"`
		Variants map[string]struct {
			DType string `json:"dtype"`
		} `json:"variants"`
	}

	// Item is a struct that contains the model aggregated information.
	Item struct {
		ID            string `json:"id"`
		Variants      string `json:"variants"`
		VocabSize     uint64 `json:"vocab_size"`
		ContextSize   uint64 `json:"context_size"`
		EmbeddingSize uint64 `json:"embedding_size"`
	}
)

func main() {
	var (
		items []Item
	)

	if err := filepath.Walk(MODELS_PATH, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		directory := filepath.Dir(path)
		if info.Name() == "config.json" {
			klog.ErrorS(nil, "Skipping %s config.json found.", "area", "generate", "directory", directory)

			return os.RemoveAll(directory)
		}

		if info.IsDir() || info.Name() != "manifest.yaml" {
			return nil
		}

		var (
			model Model
			file  *os.File
		)

		if file, err = os.Open(path); err != nil {
			return err
		}
		defer file.Close()

		if err = yaml.NewDecoder(file, yaml.UseJSONUnmarshaler()).Decode(&model); err != nil {
			return err
		}

		if len(model.Variants) == 0 {
			klog.ErrorS(nil, "Skipping %s missing variants", "area", "generate", "directory", directory)

			return os.RemoveAll(directory)
		}

		variants := maps.Keys(model.Variants)
		sort.Strings(variants)

		items = append(items, Item{
			ID:            model.ID,
			Variants:      strings.Join(variants, " "),
			VocabSize:     model.Config.VocabSize,
			ContextSize:   model.Config.ContextSize,
			EmbeddingSize: model.Config.EmbeddingSize,
		})

		klog.InfoS("Added model", "area", "generate", "model", model.ID, "variants", variants, "vocab_size", model.Config.VocabSize, "context_size", model.Config.ContextSize, "embedding_size", model.Config.EmbeddingSize)
		return nil
	}); err != nil {
		klog.ErrorS(err, "Failed to walk models", "area", "generate", "path", MODELS_PATH)
		return
	}

	f, err := os.Create("README.md")
	if err != nil {
		klog.ErrorS(err, "Failed to create README.md", "area", "generate")
		return
	}
	defer f.Close()

	if err = tmpl.Execute(f, items); err != nil {
		klog.ErrorS(err, "Failed to execute template", "area", "generate")
		return
	}

	klog.InfoS("Generated README.md", "area", "generate")
	f, err = os.Create(path.Join(MODELS_PATH, "index.yaml"))
	if err != nil {
		klog.ErrorS(err, "Failed to create models/index.yaml", "area", "generate")
		return
	}
	defer f.Close()

	if err = yaml.NewEncoder(f, yaml.UseJSONMarshaler()).Encode(items); err != nil {
		klog.ErrorS(err, "Failed to encode models/index.yaml", "area", "generate")
		return
	}

	klog.InfoS("Generated models/index.yaml", "area", "generate", "models", len(items))
}
