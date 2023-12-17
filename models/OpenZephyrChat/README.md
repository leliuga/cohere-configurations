---
license: apache-2.0
library_name: transformers
language:
- en
tags:
- mistral
- merge
- zephyr
- openchat
---

<div align="center">
  <img src="https://huggingface.co/Fredithefish/OpenZephyrChat/resolve/main/logo.jpeg" alt="Alt Text" width="250"/>

  <h1>🔥 OpenZephyrChat - Merging Zephyr-beta with OpenChat-3.5 🔥</h1>
  This Model is a ties merge between Zephyr-beta and OpenChat-3.5, it was done by using the <a>https://github.com/cg123/mergekit</a> repository
</div>


# Model Information
- **License:** Apache-2 (commercially usable)
- **Parameter size:** 7B
- **Merge:**
  - [OpenChat](https://huggingface.co/openchat/openchat_3.5) with ```--weight 0.6 --density 0.5```
  - [Zephyr](https://huggingface.co/HuggingFaceH4/zephyr-7b-beta) with ```--weight 0.3 --density 0.5```
  - [Mistral](https://huggingface.co/mistralai/Mistral-7B-v0.1) serving as the base model

 


## Evals
<table>
    <tr>
      <td>BENCHMARK</td>
      <td>SCORE</td>
    </tr>
    <tr>
      <td>ARC</td>
      <td>64.85</td>
    </tr>
    <tr>
      <td>HellaSwag</td>
      <td>85.08</td>
    </tr>
    <tr>
      <td>MMLU</td>
      <td>64.92</td>
    </tr>
    <tr>
      <td>TruthfulQA</td>
      <td>48.24</td>
    </tr>
    <tr>
      <td>Winogrande</td>
      <td>81.06</td>
    </tr>
    <tr>
      <td>GSM8K</td>
      <td>64.59</td>
    </tr>
  </table>

<a href="https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard">Source</a>
  
  <style>
    table {
      width: 50%;
      border-collapse: collapse;
      margin: 20px;
    }

    th, td {
      border: 1px solid black;
      padding: 10px;
      text-align: left;
    }

    th {
      background-color: #f2f2f2;
    }
  </style>
