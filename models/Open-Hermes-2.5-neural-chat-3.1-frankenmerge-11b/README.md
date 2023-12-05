---
license: apache-2.0
---

Frankenmerge 11b between teknium/OpenHermes-2.5-Mistral-7B and Intel/neural-chat-7b-v3-1

Merge with the following conditions


      - model: teknium/OpenHermes-2.5-Mistral-7B
        layer_range: [0, 8]

      - model: Intel/neural-chat-7b-v3-1
        layer_range: [4, 12]

      - model: teknium/OpenHermes-2.5-Mistral-7B
        layer_range: [9, 16]

      - model: Intel/neural-chat-7b-v3-1
        layer_range: [13, 20]

      - model: teknium/OpenHermes-2.5-Mistral-7B
        layer_range: [17, 24]

      - model: Intel/neural-chat-7b-v3-1
        layer_range: [21, 28]

      - model: teknium/OpenHermes-2.5-Mistral-7B
        layer_range: [25, 32]
merge_method: passthrough



Benchmarks are coming soon...
