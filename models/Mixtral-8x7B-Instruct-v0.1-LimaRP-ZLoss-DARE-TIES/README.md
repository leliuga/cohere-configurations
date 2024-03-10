---
base_model: []
tags:
- mergekit
- merge

---
# Mixtral-8x7B-Instruct-v0.1-LimaRP-ZLoss-DARE-TIES

This is a merge of pre-trained language models created using [mergekit](https://github.com/cg123/mergekit).

## Merge Details
### Merge Method

This model was merged using the [DARE](https://arxiv.org/abs/2311.03099) [TIES](https://arxiv.org/abs/2306.01708) merge method using ./extra_hdd/Mixtral-8x7B-v0.1 as a base.

### Models Merged

The following models were included in the merge:
* ./extra_hdd2/Mixtral-8x7B-Instruct-v0.1
* ./extra_hdd/Mixtral-8x7B-v0.1-LimaRP-ZLoss

### Configuration

The following YAML configuration was used to produce this model:

```yaml
models:
  - model: ./extra_hdd2/Mixtral-8x7B-Instruct-v0.1
    parameters:
      density: 0.5
      weight: 1.0
  - model: ./extra_hdd/Mixtral-8x7B-v0.1-LimaRP-ZLoss
    parameters:
      density: 0.5
      weight: 0.5
merge_method: dare_ties
base_model: ./extra_hdd/Mixtral-8x7B-v0.1
parameters:
  #normalize: false
  #int8_mask: true
dtype: bfloat16

```
