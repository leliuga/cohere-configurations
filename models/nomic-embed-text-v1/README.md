---
library_name: sentence-transformers
pipeline_tag: sentence-similarity
tags:
  - feature-extraction
  - sentence-similarity
  - mteb
  - transformers
  - transformers.js
model-index:
- name: epoch_0_model
  results:
  - task:
      type: Classification
    dataset:
      type: mteb/amazon_counterfactual
      name: MTEB AmazonCounterfactualClassification (en)
      config: en
      split: test
      revision: e8379541af4e31359cca9fbcf4b00f2671dba205
    metrics:
    - type: accuracy
      value: 76.8507462686567
    - type: ap
      value: 40.592189159090495
    - type: f1
      value: 71.01634655512476
  - task:
      type: Classification
    dataset:
      type: mteb/amazon_polarity
      name: MTEB AmazonPolarityClassification
      config: default
      split: test
      revision: e2d317d38cd51312af73b3d32a06d1a08b442046
    metrics:
    - type: accuracy
      value: 91.51892500000001
    - type: ap
      value: 88.50346762975335
    - type: f1
      value: 91.50342077459624
  - task:
      type: Classification
    dataset:
      type: mteb/amazon_reviews_multi
      name: MTEB AmazonReviewsClassification (en)
      config: en
      split: test
      revision: 1399c76144fd37290681b995c656ef9b2e06e26d
    metrics:
    - type: accuracy
      value: 47.364
    - type: f1
      value: 46.72708080922794
  - task:
      type: Retrieval
    dataset:
      type: arguana
      name: MTEB ArguAna
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 25.178
    - type: map_at_10
      value: 40.244
    - type: map_at_100
      value: 41.321999999999996
    - type: map_at_1000
      value: 41.331
    - type: map_at_3
      value: 35.016999999999996
    - type: map_at_5
      value: 37.99
    - type: mrr_at_1
      value: 25.605
    - type: mrr_at_10
      value: 40.422000000000004
    - type: mrr_at_100
      value: 41.507
    - type: mrr_at_1000
      value: 41.516
    - type: mrr_at_3
      value: 35.23
    - type: mrr_at_5
      value: 38.15
    - type: ndcg_at_1
      value: 25.178
    - type: ndcg_at_10
      value: 49.258
    - type: ndcg_at_100
      value: 53.776
    - type: ndcg_at_1000
      value: 53.995000000000005
    - type: ndcg_at_3
      value: 38.429
    - type: ndcg_at_5
      value: 43.803
    - type: precision_at_1
      value: 25.178
    - type: precision_at_10
      value: 7.831
    - type: precision_at_100
      value: 0.979
    - type: precision_at_1000
      value: 0.1
    - type: precision_at_3
      value: 16.121
    - type: precision_at_5
      value: 12.29
    - type: recall_at_1
      value: 25.178
    - type: recall_at_10
      value: 78.307
    - type: recall_at_100
      value: 97.866
    - type: recall_at_1000
      value: 99.57300000000001
    - type: recall_at_3
      value: 48.364000000000004
    - type: recall_at_5
      value: 61.451
  - task:
      type: Clustering
    dataset:
      type: mteb/arxiv-clustering-p2p
      name: MTEB ArxivClusteringP2P
      config: default
      split: test
      revision: a122ad7f3f0291bf49cc6f4d32aa80929df69d5d
    metrics:
    - type: v_measure
      value: 45.93034494751465
  - task:
      type: Clustering
    dataset:
      type: mteb/arxiv-clustering-s2s
      name: MTEB ArxivClusteringS2S
      config: default
      split: test
      revision: f910caf1a6075f7329cdf8c1a6135696f37dbd53
    metrics:
    - type: v_measure
      value: 36.64579480054327
  - task:
      type: Reranking
    dataset:
      type: mteb/askubuntudupquestions-reranking
      name: MTEB AskUbuntuDupQuestions
      config: default
      split: test
      revision: 2000358ca161889fa9c082cb41daa8dcfb161a54
    metrics:
    - type: map
      value: 60.601310529222054
    - type: mrr
      value: 75.04484896451656
  - task:
      type: STS
    dataset:
      type: mteb/biosses-sts
      name: MTEB BIOSSES
      config: default
      split: test
      revision: d3fb88f8f02e40887cd149695127462bbcf29b4a
    metrics:
    - type: cos_sim_pearson
      value: 88.57797718095814
    - type: cos_sim_spearman
      value: 86.47064499110101
    - type: euclidean_pearson
      value: 87.4559602783142
    - type: euclidean_spearman
      value: 86.47064499110101
    - type: manhattan_pearson
      value: 87.7232764230245
    - type: manhattan_spearman
      value: 86.91222131777742
  - task:
      type: Classification
    dataset:
      type: mteb/banking77
      name: MTEB Banking77Classification
      config: default
      split: test
      revision: 0fd18e25b25c072e09e0d92ab615fda904d66300
    metrics:
    - type: accuracy
      value: 84.5422077922078
    - type: f1
      value: 84.47657456950589
  - task:
      type: Clustering
    dataset:
      type: mteb/biorxiv-clustering-p2p
      name: MTEB BiorxivClusteringP2P
      config: default
      split: test
      revision: 65b79d1d13f80053f67aca9498d9402c2d9f1f40
    metrics:
    - type: v_measure
      value: 38.48953561974464
  - task:
      type: Clustering
    dataset:
      type: mteb/biorxiv-clustering-s2s
      name: MTEB BiorxivClusteringS2S
      config: default
      split: test
      revision: 258694dd0231531bc1fd9de6ceb52a0853c6d908
    metrics:
    - type: v_measure
      value: 32.75995857510105
  - task:
      type: Retrieval
    dataset:
      type: BeIR/cqadupstack
      name: MTEB CQADupstackAndroidRetrieval
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 30.008000000000003
    - type: map_at_10
      value: 39.51
    - type: map_at_100
      value: 40.841
    - type: map_at_1000
      value: 40.973
    - type: map_at_3
      value: 36.248999999999995
    - type: map_at_5
      value: 38.096999999999994
    - type: mrr_at_1
      value: 36.481
    - type: mrr_at_10
      value: 44.818000000000005
    - type: mrr_at_100
      value: 45.64
    - type: mrr_at_1000
      value: 45.687
    - type: mrr_at_3
      value: 42.036
    - type: mrr_at_5
      value: 43.782
    - type: ndcg_at_1
      value: 36.481
    - type: ndcg_at_10
      value: 45.152
    - type: ndcg_at_100
      value: 50.449
    - type: ndcg_at_1000
      value: 52.76499999999999
    - type: ndcg_at_3
      value: 40.161
    - type: ndcg_at_5
      value: 42.577999999999996
    - type: precision_at_1
      value: 36.481
    - type: precision_at_10
      value: 8.369
    - type: precision_at_100
      value: 1.373
    - type: precision_at_1000
      value: 0.186
    - type: precision_at_3
      value: 18.693
    - type: precision_at_5
      value: 13.533999999999999
    - type: recall_at_1
      value: 30.008000000000003
    - type: recall_at_10
      value: 56.108999999999995
    - type: recall_at_100
      value: 78.55499999999999
    - type: recall_at_1000
      value: 93.659
    - type: recall_at_3
      value: 41.754999999999995
    - type: recall_at_5
      value: 48.296
  - task:
      type: Retrieval
    dataset:
      type: BeIR/cqadupstack
      name: MTEB CQADupstackEnglishRetrieval
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 30.262
    - type: map_at_10
      value: 40.139
    - type: map_at_100
      value: 41.394
    - type: map_at_1000
      value: 41.526
    - type: map_at_3
      value: 37.155
    - type: map_at_5
      value: 38.785
    - type: mrr_at_1
      value: 38.153
    - type: mrr_at_10
      value: 46.369
    - type: mrr_at_100
      value: 47.072
    - type: mrr_at_1000
      value: 47.111999999999995
    - type: mrr_at_3
      value: 44.268
    - type: mrr_at_5
      value: 45.389
    - type: ndcg_at_1
      value: 38.153
    - type: ndcg_at_10
      value: 45.925
    - type: ndcg_at_100
      value: 50.394000000000005
    - type: ndcg_at_1000
      value: 52.37500000000001
    - type: ndcg_at_3
      value: 41.754000000000005
    - type: ndcg_at_5
      value: 43.574
    - type: precision_at_1
      value: 38.153
    - type: precision_at_10
      value: 8.796
    - type: precision_at_100
      value: 1.432
    - type: precision_at_1000
      value: 0.189
    - type: precision_at_3
      value: 20.318
    - type: precision_at_5
      value: 14.395
    - type: recall_at_1
      value: 30.262
    - type: recall_at_10
      value: 55.72200000000001
    - type: recall_at_100
      value: 74.97500000000001
    - type: recall_at_1000
      value: 87.342
    - type: recall_at_3
      value: 43.129
    - type: recall_at_5
      value: 48.336
  - task:
      type: Retrieval
    dataset:
      type: BeIR/cqadupstack
      name: MTEB CQADupstackGamingRetrieval
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 39.951
    - type: map_at_10
      value: 51.248000000000005
    - type: map_at_100
      value: 52.188
    - type: map_at_1000
      value: 52.247
    - type: map_at_3
      value: 48.211
    - type: map_at_5
      value: 49.797000000000004
    - type: mrr_at_1
      value: 45.329
    - type: mrr_at_10
      value: 54.749
    - type: mrr_at_100
      value: 55.367999999999995
    - type: mrr_at_1000
      value: 55.400000000000006
    - type: mrr_at_3
      value: 52.382
    - type: mrr_at_5
      value: 53.649
    - type: ndcg_at_1
      value: 45.329
    - type: ndcg_at_10
      value: 56.847
    - type: ndcg_at_100
      value: 60.738
    - type: ndcg_at_1000
      value: 61.976
    - type: ndcg_at_3
      value: 51.59
    - type: ndcg_at_5
      value: 53.915
    - type: precision_at_1
      value: 45.329
    - type: precision_at_10
      value: 8.959
    - type: precision_at_100
      value: 1.187
    - type: precision_at_1000
      value: 0.134
    - type: precision_at_3
      value: 22.612
    - type: precision_at_5
      value: 15.273
    - type: recall_at_1
      value: 39.951
    - type: recall_at_10
      value: 70.053
    - type: recall_at_100
      value: 86.996
    - type: recall_at_1000
      value: 95.707
    - type: recall_at_3
      value: 56.032000000000004
    - type: recall_at_5
      value: 61.629999999999995
  - task:
      type: Retrieval
    dataset:
      type: BeIR/cqadupstack
      name: MTEB CQADupstackGisRetrieval
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 25.566
    - type: map_at_10
      value: 33.207
    - type: map_at_100
      value: 34.166000000000004
    - type: map_at_1000
      value: 34.245
    - type: map_at_3
      value: 30.94
    - type: map_at_5
      value: 32.01
    - type: mrr_at_1
      value: 27.345000000000002
    - type: mrr_at_10
      value: 35.193000000000005
    - type: mrr_at_100
      value: 35.965
    - type: mrr_at_1000
      value: 36.028999999999996
    - type: mrr_at_3
      value: 32.806000000000004
    - type: mrr_at_5
      value: 34.021
    - type: ndcg_at_1
      value: 27.345000000000002
    - type: ndcg_at_10
      value: 37.891999999999996
    - type: ndcg_at_100
      value: 42.664
    - type: ndcg_at_1000
      value: 44.757000000000005
    - type: ndcg_at_3
      value: 33.123000000000005
    - type: ndcg_at_5
      value: 35.035
    - type: precision_at_1
      value: 27.345000000000002
    - type: precision_at_10
      value: 5.763
    - type: precision_at_100
      value: 0.859
    - type: precision_at_1000
      value: 0.108
    - type: precision_at_3
      value: 13.71
    - type: precision_at_5
      value: 9.401
    - type: recall_at_1
      value: 25.566
    - type: recall_at_10
      value: 50.563
    - type: recall_at_100
      value: 72.86399999999999
    - type: recall_at_1000
      value: 88.68599999999999
    - type: recall_at_3
      value: 37.43
    - type: recall_at_5
      value: 41.894999999999996
  - task:
      type: Retrieval
    dataset:
      type: BeIR/cqadupstack
      name: MTEB CQADupstackMathematicaRetrieval
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 16.663
    - type: map_at_10
      value: 23.552
    - type: map_at_100
      value: 24.538
    - type: map_at_1000
      value: 24.661
    - type: map_at_3
      value: 21.085
    - type: map_at_5
      value: 22.391
    - type: mrr_at_1
      value: 20.025000000000002
    - type: mrr_at_10
      value: 27.643
    - type: mrr_at_100
      value: 28.499999999999996
    - type: mrr_at_1000
      value: 28.582
    - type: mrr_at_3
      value: 25.083
    - type: mrr_at_5
      value: 26.544
    - type: ndcg_at_1
      value: 20.025000000000002
    - type: ndcg_at_10
      value: 28.272000000000002
    - type: ndcg_at_100
      value: 33.353
    - type: ndcg_at_1000
      value: 36.454
    - type: ndcg_at_3
      value: 23.579
    - type: ndcg_at_5
      value: 25.685000000000002
    - type: precision_at_1
      value: 20.025000000000002
    - type: precision_at_10
      value: 5.187
    - type: precision_at_100
      value: 0.897
    - type: precision_at_1000
      value: 0.13
    - type: precision_at_3
      value: 10.987
    - type: precision_at_5
      value: 8.06
    - type: recall_at_1
      value: 16.663
    - type: recall_at_10
      value: 38.808
    - type: recall_at_100
      value: 61.305
    - type: recall_at_1000
      value: 83.571
    - type: recall_at_3
      value: 25.907999999999998
    - type: recall_at_5
      value: 31.214
  - task:
      type: Retrieval
    dataset:
      type: BeIR/cqadupstack
      name: MTEB CQADupstackPhysicsRetrieval
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 27.695999999999998
    - type: map_at_10
      value: 37.018
    - type: map_at_100
      value: 38.263000000000005
    - type: map_at_1000
      value: 38.371
    - type: map_at_3
      value: 34.226
    - type: map_at_5
      value: 35.809999999999995
    - type: mrr_at_1
      value: 32.916000000000004
    - type: mrr_at_10
      value: 42.067
    - type: mrr_at_100
      value: 42.925000000000004
    - type: mrr_at_1000
      value: 42.978
    - type: mrr_at_3
      value: 39.637
    - type: mrr_at_5
      value: 41.134
    - type: ndcg_at_1
      value: 32.916000000000004
    - type: ndcg_at_10
      value: 42.539
    - type: ndcg_at_100
      value: 47.873
    - type: ndcg_at_1000
      value: 50.08200000000001
    - type: ndcg_at_3
      value: 37.852999999999994
    - type: ndcg_at_5
      value: 40.201
    - type: precision_at_1
      value: 32.916000000000004
    - type: precision_at_10
      value: 7.5840000000000005
    - type: precision_at_100
      value: 1.199
    - type: precision_at_1000
      value: 0.155
    - type: precision_at_3
      value: 17.485
    - type: precision_at_5
      value: 12.512
    - type: recall_at_1
      value: 27.695999999999998
    - type: recall_at_10
      value: 53.638
    - type: recall_at_100
      value: 76.116
    - type: recall_at_1000
      value: 91.069
    - type: recall_at_3
      value: 41.13
    - type: recall_at_5
      value: 46.872
  - task:
      type: Retrieval
    dataset:
      type: BeIR/cqadupstack
      name: MTEB CQADupstackProgrammersRetrieval
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 24.108
    - type: map_at_10
      value: 33.372
    - type: map_at_100
      value: 34.656
    - type: map_at_1000
      value: 34.768
    - type: map_at_3
      value: 30.830999999999996
    - type: map_at_5
      value: 32.204
    - type: mrr_at_1
      value: 29.110000000000003
    - type: mrr_at_10
      value: 37.979
    - type: mrr_at_100
      value: 38.933
    - type: mrr_at_1000
      value: 38.988
    - type: mrr_at_3
      value: 35.731
    - type: mrr_at_5
      value: 36.963
    - type: ndcg_at_1
      value: 29.110000000000003
    - type: ndcg_at_10
      value: 38.635000000000005
    - type: ndcg_at_100
      value: 44.324999999999996
    - type: ndcg_at_1000
      value: 46.747
    - type: ndcg_at_3
      value: 34.37
    - type: ndcg_at_5
      value: 36.228
    - type: precision_at_1
      value: 29.110000000000003
    - type: precision_at_10
      value: 6.963
    - type: precision_at_100
      value: 1.146
    - type: precision_at_1000
      value: 0.152
    - type: precision_at_3
      value: 16.400000000000002
    - type: precision_at_5
      value: 11.552999999999999
    - type: recall_at_1
      value: 24.108
    - type: recall_at_10
      value: 49.597
    - type: recall_at_100
      value: 73.88900000000001
    - type: recall_at_1000
      value: 90.62400000000001
    - type: recall_at_3
      value: 37.662
    - type: recall_at_5
      value: 42.565
  - task:
      type: Retrieval
    dataset:
      type: BeIR/cqadupstack
      name: MTEB CQADupstackRetrieval
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 25.00791666666667
    - type: map_at_10
      value: 33.287749999999996
    - type: map_at_100
      value: 34.41141666666667
    - type: map_at_1000
      value: 34.52583333333333
    - type: map_at_3
      value: 30.734416666666668
    - type: map_at_5
      value: 32.137166666666666
    - type: mrr_at_1
      value: 29.305666666666664
    - type: mrr_at_10
      value: 37.22966666666666
    - type: mrr_at_100
      value: 38.066583333333334
    - type: mrr_at_1000
      value: 38.12616666666667
    - type: mrr_at_3
      value: 34.92275
    - type: mrr_at_5
      value: 36.23333333333334
    - type: ndcg_at_1
      value: 29.305666666666664
    - type: ndcg_at_10
      value: 38.25533333333333
    - type: ndcg_at_100
      value: 43.25266666666666
    - type: ndcg_at_1000
      value: 45.63583333333334
    - type: ndcg_at_3
      value: 33.777166666666666
    - type: ndcg_at_5
      value: 35.85
    - type: precision_at_1
      value: 29.305666666666664
    - type: precision_at_10
      value: 6.596416666666667
    - type: precision_at_100
      value: 1.0784166666666668
    - type: precision_at_1000
      value: 0.14666666666666664
    - type: precision_at_3
      value: 15.31075
    - type: precision_at_5
      value: 10.830916666666667
    - type: recall_at_1
      value: 25.00791666666667
    - type: recall_at_10
      value: 49.10933333333333
    - type: recall_at_100
      value: 71.09216666666667
    - type: recall_at_1000
      value: 87.77725000000001
    - type: recall_at_3
      value: 36.660916666666665
    - type: recall_at_5
      value: 41.94149999999999
  - task:
      type: Retrieval
    dataset:
      type: BeIR/cqadupstack
      name: MTEB CQADupstackStatsRetrieval
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 23.521
    - type: map_at_10
      value: 30.043
    - type: map_at_100
      value: 30.936000000000003
    - type: map_at_1000
      value: 31.022
    - type: map_at_3
      value: 27.926000000000002
    - type: map_at_5
      value: 29.076999999999998
    - type: mrr_at_1
      value: 26.227
    - type: mrr_at_10
      value: 32.822
    - type: mrr_at_100
      value: 33.61
    - type: mrr_at_1000
      value: 33.672000000000004
    - type: mrr_at_3
      value: 30.776999999999997
    - type: mrr_at_5
      value: 31.866
    - type: ndcg_at_1
      value: 26.227
    - type: ndcg_at_10
      value: 34.041
    - type: ndcg_at_100
      value: 38.394
    - type: ndcg_at_1000
      value: 40.732
    - type: ndcg_at_3
      value: 30.037999999999997
    - type: ndcg_at_5
      value: 31.845000000000002
    - type: precision_at_1
      value: 26.227
    - type: precision_at_10
      value: 5.244999999999999
    - type: precision_at_100
      value: 0.808
    - type: precision_at_1000
      value: 0.107
    - type: precision_at_3
      value: 12.679000000000002
    - type: precision_at_5
      value: 8.773
    - type: recall_at_1
      value: 23.521
    - type: recall_at_10
      value: 43.633
    - type: recall_at_100
      value: 63.126000000000005
    - type: recall_at_1000
      value: 80.765
    - type: recall_at_3
      value: 32.614
    - type: recall_at_5
      value: 37.15
  - task:
      type: Retrieval
    dataset:
      type: BeIR/cqadupstack
      name: MTEB CQADupstackTexRetrieval
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 16.236
    - type: map_at_10
      value: 22.898
    - type: map_at_100
      value: 23.878
    - type: map_at_1000
      value: 24.009
    - type: map_at_3
      value: 20.87
    - type: map_at_5
      value: 22.025
    - type: mrr_at_1
      value: 19.339000000000002
    - type: mrr_at_10
      value: 26.382
    - type: mrr_at_100
      value: 27.245
    - type: mrr_at_1000
      value: 27.33
    - type: mrr_at_3
      value: 24.386
    - type: mrr_at_5
      value: 25.496000000000002
    - type: ndcg_at_1
      value: 19.339000000000002
    - type: ndcg_at_10
      value: 27.139999999999997
    - type: ndcg_at_100
      value: 31.944
    - type: ndcg_at_1000
      value: 35.077999999999996
    - type: ndcg_at_3
      value: 23.424
    - type: ndcg_at_5
      value: 25.188
    - type: precision_at_1
      value: 19.339000000000002
    - type: precision_at_10
      value: 4.8309999999999995
    - type: precision_at_100
      value: 0.845
    - type: precision_at_1000
      value: 0.128
    - type: precision_at_3
      value: 10.874
    - type: precision_at_5
      value: 7.825
    - type: recall_at_1
      value: 16.236
    - type: recall_at_10
      value: 36.513
    - type: recall_at_100
      value: 57.999
    - type: recall_at_1000
      value: 80.512
    - type: recall_at_3
      value: 26.179999999999996
    - type: recall_at_5
      value: 30.712
  - task:
      type: Retrieval
    dataset:
      type: BeIR/cqadupstack
      name: MTEB CQADupstackUnixRetrieval
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 24.11
    - type: map_at_10
      value: 31.566
    - type: map_at_100
      value: 32.647
    - type: map_at_1000
      value: 32.753
    - type: map_at_3
      value: 29.24
    - type: map_at_5
      value: 30.564999999999998
    - type: mrr_at_1
      value: 28.265
    - type: mrr_at_10
      value: 35.504000000000005
    - type: mrr_at_100
      value: 36.436
    - type: mrr_at_1000
      value: 36.503
    - type: mrr_at_3
      value: 33.349000000000004
    - type: mrr_at_5
      value: 34.622
    - type: ndcg_at_1
      value: 28.265
    - type: ndcg_at_10
      value: 36.192
    - type: ndcg_at_100
      value: 41.388000000000005
    - type: ndcg_at_1000
      value: 43.948
    - type: ndcg_at_3
      value: 31.959
    - type: ndcg_at_5
      value: 33.998
    - type: precision_at_1
      value: 28.265
    - type: precision_at_10
      value: 5.989
    - type: precision_at_100
      value: 0.9650000000000001
    - type: precision_at_1000
      value: 0.13
    - type: precision_at_3
      value: 14.335
    - type: precision_at_5
      value: 10.112
    - type: recall_at_1
      value: 24.11
    - type: recall_at_10
      value: 46.418
    - type: recall_at_100
      value: 69.314
    - type: recall_at_1000
      value: 87.397
    - type: recall_at_3
      value: 34.724
    - type: recall_at_5
      value: 39.925
  - task:
      type: Retrieval
    dataset:
      type: BeIR/cqadupstack
      name: MTEB CQADupstackWebmastersRetrieval
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 22.091
    - type: map_at_10
      value: 29.948999999999998
    - type: map_at_100
      value: 31.502000000000002
    - type: map_at_1000
      value: 31.713
    - type: map_at_3
      value: 27.464
    - type: map_at_5
      value: 28.968
    - type: mrr_at_1
      value: 26.482
    - type: mrr_at_10
      value: 34.009
    - type: mrr_at_100
      value: 35.081
    - type: mrr_at_1000
      value: 35.138000000000005
    - type: mrr_at_3
      value: 31.785000000000004
    - type: mrr_at_5
      value: 33.178999999999995
    - type: ndcg_at_1
      value: 26.482
    - type: ndcg_at_10
      value: 35.008
    - type: ndcg_at_100
      value: 41.272999999999996
    - type: ndcg_at_1000
      value: 43.972
    - type: ndcg_at_3
      value: 30.804
    - type: ndcg_at_5
      value: 33.046
    - type: precision_at_1
      value: 26.482
    - type: precision_at_10
      value: 6.462
    - type: precision_at_100
      value: 1.431
    - type: precision_at_1000
      value: 0.22899999999999998
    - type: precision_at_3
      value: 14.360999999999999
    - type: precision_at_5
      value: 10.474
    - type: recall_at_1
      value: 22.091
    - type: recall_at_10
      value: 45.125
    - type: recall_at_100
      value: 72.313
    - type: recall_at_1000
      value: 89.503
    - type: recall_at_3
      value: 33.158
    - type: recall_at_5
      value: 39.086999999999996
  - task:
      type: Retrieval
    dataset:
      type: BeIR/cqadupstack
      name: MTEB CQADupstackWordpressRetrieval
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 19.883
    - type: map_at_10
      value: 26.951000000000004
    - type: map_at_100
      value: 27.927999999999997
    - type: map_at_1000
      value: 28.022000000000002
    - type: map_at_3
      value: 24.616
    - type: map_at_5
      value: 25.917
    - type: mrr_at_1
      value: 21.996
    - type: mrr_at_10
      value: 29.221000000000004
    - type: mrr_at_100
      value: 30.024
    - type: mrr_at_1000
      value: 30.095
    - type: mrr_at_3
      value: 26.833000000000002
    - type: mrr_at_5
      value: 28.155
    - type: ndcg_at_1
      value: 21.996
    - type: ndcg_at_10
      value: 31.421
    - type: ndcg_at_100
      value: 36.237
    - type: ndcg_at_1000
      value: 38.744
    - type: ndcg_at_3
      value: 26.671
    - type: ndcg_at_5
      value: 28.907
    - type: precision_at_1
      value: 21.996
    - type: precision_at_10
      value: 5.009
    - type: precision_at_100
      value: 0.799
    - type: precision_at_1000
      value: 0.11199999999999999
    - type: precision_at_3
      value: 11.275
    - type: precision_at_5
      value: 8.059
    - type: recall_at_1
      value: 19.883
    - type: recall_at_10
      value: 43.132999999999996
    - type: recall_at_100
      value: 65.654
    - type: recall_at_1000
      value: 84.492
    - type: recall_at_3
      value: 30.209000000000003
    - type: recall_at_5
      value: 35.616
  - task:
      type: Retrieval
    dataset:
      type: climate-fever
      name: MTEB ClimateFEVER
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 17.756
    - type: map_at_10
      value: 30.378
    - type: map_at_100
      value: 32.537
    - type: map_at_1000
      value: 32.717
    - type: map_at_3
      value: 25.599
    - type: map_at_5
      value: 28.372999999999998
    - type: mrr_at_1
      value: 41.303
    - type: mrr_at_10
      value: 53.483999999999995
    - type: mrr_at_100
      value: 54.106
    - type: mrr_at_1000
      value: 54.127
    - type: mrr_at_3
      value: 50.315
    - type: mrr_at_5
      value: 52.396
    - type: ndcg_at_1
      value: 41.303
    - type: ndcg_at_10
      value: 40.503
    - type: ndcg_at_100
      value: 47.821000000000005
    - type: ndcg_at_1000
      value: 50.788
    - type: ndcg_at_3
      value: 34.364
    - type: ndcg_at_5
      value: 36.818
    - type: precision_at_1
      value: 41.303
    - type: precision_at_10
      value: 12.463000000000001
    - type: precision_at_100
      value: 2.037
    - type: precision_at_1000
      value: 0.26
    - type: precision_at_3
      value: 25.798
    - type: precision_at_5
      value: 19.896
    - type: recall_at_1
      value: 17.756
    - type: recall_at_10
      value: 46.102
    - type: recall_at_100
      value: 70.819
    - type: recall_at_1000
      value: 87.21799999999999
    - type: recall_at_3
      value: 30.646
    - type: recall_at_5
      value: 38.022
  - task:
      type: Retrieval
    dataset:
      type: dbpedia-entity
      name: MTEB DBPedia
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 9.033
    - type: map_at_10
      value: 20.584
    - type: map_at_100
      value: 29.518
    - type: map_at_1000
      value: 31.186000000000003
    - type: map_at_3
      value: 14.468
    - type: map_at_5
      value: 17.177
    - type: mrr_at_1
      value: 69.75
    - type: mrr_at_10
      value: 77.025
    - type: mrr_at_100
      value: 77.36699999999999
    - type: mrr_at_1000
      value: 77.373
    - type: mrr_at_3
      value: 75.583
    - type: mrr_at_5
      value: 76.396
    - type: ndcg_at_1
      value: 58.5
    - type: ndcg_at_10
      value: 45.033
    - type: ndcg_at_100
      value: 49.071
    - type: ndcg_at_1000
      value: 56.056
    - type: ndcg_at_3
      value: 49.936
    - type: ndcg_at_5
      value: 47.471999999999994
    - type: precision_at_1
      value: 69.75
    - type: precision_at_10
      value: 35.775
    - type: precision_at_100
      value: 11.594999999999999
    - type: precision_at_1000
      value: 2.062
    - type: precision_at_3
      value: 52.5
    - type: precision_at_5
      value: 45.300000000000004
    - type: recall_at_1
      value: 9.033
    - type: recall_at_10
      value: 26.596999999999998
    - type: recall_at_100
      value: 54.607000000000006
    - type: recall_at_1000
      value: 76.961
    - type: recall_at_3
      value: 15.754999999999999
    - type: recall_at_5
      value: 20.033
  - task:
      type: Classification
    dataset:
      type: mteb/emotion
      name: MTEB EmotionClassification
      config: default
      split: test
      revision: 4f58c6b202a23cf9a4da393831edf4f9183cad37
    metrics:
    - type: accuracy
      value: 48.345000000000006
    - type: f1
      value: 43.4514918068706
  - task:
      type: Retrieval
    dataset:
      type: fever
      name: MTEB FEVER
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 71.29100000000001
    - type: map_at_10
      value: 81.059
    - type: map_at_100
      value: 81.341
    - type: map_at_1000
      value: 81.355
    - type: map_at_3
      value: 79.74799999999999
    - type: map_at_5
      value: 80.612
    - type: mrr_at_1
      value: 76.40299999999999
    - type: mrr_at_10
      value: 84.615
    - type: mrr_at_100
      value: 84.745
    - type: mrr_at_1000
      value: 84.748
    - type: mrr_at_3
      value: 83.776
    - type: mrr_at_5
      value: 84.343
    - type: ndcg_at_1
      value: 76.40299999999999
    - type: ndcg_at_10
      value: 84.981
    - type: ndcg_at_100
      value: 86.00999999999999
    - type: ndcg_at_1000
      value: 86.252
    - type: ndcg_at_3
      value: 82.97
    - type: ndcg_at_5
      value: 84.152
    - type: precision_at_1
      value: 76.40299999999999
    - type: precision_at_10
      value: 10.446
    - type: precision_at_100
      value: 1.1199999999999999
    - type: precision_at_1000
      value: 0.116
    - type: precision_at_3
      value: 32.147999999999996
    - type: precision_at_5
      value: 20.135
    - type: recall_at_1
      value: 71.29100000000001
    - type: recall_at_10
      value: 93.232
    - type: recall_at_100
      value: 97.363
    - type: recall_at_1000
      value: 98.905
    - type: recall_at_3
      value: 87.893
    - type: recall_at_5
      value: 90.804
  - task:
      type: Retrieval
    dataset:
      type: fiqa
      name: MTEB FiQA2018
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 18.667
    - type: map_at_10
      value: 30.853
    - type: map_at_100
      value: 32.494
    - type: map_at_1000
      value: 32.677
    - type: map_at_3
      value: 26.91
    - type: map_at_5
      value: 29.099000000000004
    - type: mrr_at_1
      value: 37.191
    - type: mrr_at_10
      value: 46.171
    - type: mrr_at_100
      value: 47.056
    - type: mrr_at_1000
      value: 47.099000000000004
    - type: mrr_at_3
      value: 44.059
    - type: mrr_at_5
      value: 45.147
    - type: ndcg_at_1
      value: 37.191
    - type: ndcg_at_10
      value: 38.437
    - type: ndcg_at_100
      value: 44.62
    - type: ndcg_at_1000
      value: 47.795
    - type: ndcg_at_3
      value: 35.003
    - type: ndcg_at_5
      value: 36.006
    - type: precision_at_1
      value: 37.191
    - type: precision_at_10
      value: 10.586
    - type: precision_at_100
      value: 1.688
    - type: precision_at_1000
      value: 0.22699999999999998
    - type: precision_at_3
      value: 23.302
    - type: precision_at_5
      value: 17.006
    - type: recall_at_1
      value: 18.667
    - type: recall_at_10
      value: 45.367000000000004
    - type: recall_at_100
      value: 68.207
    - type: recall_at_1000
      value: 87.072
    - type: recall_at_3
      value: 32.129000000000005
    - type: recall_at_5
      value: 37.719
  - task:
      type: Retrieval
    dataset:
      type: hotpotqa
      name: MTEB HotpotQA
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 39.494
    - type: map_at_10
      value: 66.223
    - type: map_at_100
      value: 67.062
    - type: map_at_1000
      value: 67.11500000000001
    - type: map_at_3
      value: 62.867
    - type: map_at_5
      value: 64.994
    - type: mrr_at_1
      value: 78.987
    - type: mrr_at_10
      value: 84.585
    - type: mrr_at_100
      value: 84.773
    - type: mrr_at_1000
      value: 84.77900000000001
    - type: mrr_at_3
      value: 83.592
    - type: mrr_at_5
      value: 84.235
    - type: ndcg_at_1
      value: 78.987
    - type: ndcg_at_10
      value: 73.64
    - type: ndcg_at_100
      value: 76.519
    - type: ndcg_at_1000
      value: 77.51
    - type: ndcg_at_3
      value: 68.893
    - type: ndcg_at_5
      value: 71.585
    - type: precision_at_1
      value: 78.987
    - type: precision_at_10
      value: 15.529000000000002
    - type: precision_at_100
      value: 1.7770000000000001
    - type: precision_at_1000
      value: 0.191
    - type: precision_at_3
      value: 44.808
    - type: precision_at_5
      value: 29.006999999999998
    - type: recall_at_1
      value: 39.494
    - type: recall_at_10
      value: 77.643
    - type: recall_at_100
      value: 88.825
    - type: recall_at_1000
      value: 95.321
    - type: recall_at_3
      value: 67.211
    - type: recall_at_5
      value: 72.519
  - task:
      type: Classification
    dataset:
      type: mteb/imdb
      name: MTEB ImdbClassification
      config: default
      split: test
      revision: 3d86128a09e091d6018b6d26cad27f2739fc2db7
    metrics:
    - type: accuracy
      value: 85.55959999999999
    - type: ap
      value: 80.7246500384617
    - type: f1
      value: 85.52336485065454
  - task:
      type: Retrieval
    dataset:
      type: msmarco
      name: MTEB MSMARCO
      config: default
      split: dev
      revision: None
    metrics:
    - type: map_at_1
      value: 23.631
    - type: map_at_10
      value: 36.264
    - type: map_at_100
      value: 37.428
    - type: map_at_1000
      value: 37.472
    - type: map_at_3
      value: 32.537
    - type: map_at_5
      value: 34.746
    - type: mrr_at_1
      value: 24.312
    - type: mrr_at_10
      value: 36.858000000000004
    - type: mrr_at_100
      value: 37.966
    - type: mrr_at_1000
      value: 38.004
    - type: mrr_at_3
      value: 33.188
    - type: mrr_at_5
      value: 35.367
    - type: ndcg_at_1
      value: 24.312
    - type: ndcg_at_10
      value: 43.126999999999995
    - type: ndcg_at_100
      value: 48.642
    - type: ndcg_at_1000
      value: 49.741
    - type: ndcg_at_3
      value: 35.589
    - type: ndcg_at_5
      value: 39.515
    - type: precision_at_1
      value: 24.312
    - type: precision_at_10
      value: 6.699
    - type: precision_at_100
      value: 0.9450000000000001
    - type: precision_at_1000
      value: 0.104
    - type: precision_at_3
      value: 15.153
    - type: precision_at_5
      value: 11.065999999999999
    - type: recall_at_1
      value: 23.631
    - type: recall_at_10
      value: 64.145
    - type: recall_at_100
      value: 89.41
    - type: recall_at_1000
      value: 97.83500000000001
    - type: recall_at_3
      value: 43.769000000000005
    - type: recall_at_5
      value: 53.169
  - task:
      type: Classification
    dataset:
      type: mteb/mtop_domain
      name: MTEB MTOPDomainClassification (en)
      config: en
      split: test
      revision: d80d48c1eb48d3562165c59d59d0034df9fff0bf
    metrics:
    - type: accuracy
      value: 93.4108527131783
    - type: f1
      value: 93.1415880261038
  - task:
      type: Classification
    dataset:
      type: mteb/mtop_intent
      name: MTEB MTOPIntentClassification (en)
      config: en
      split: test
      revision: ae001d0e6b1228650b7bd1c2c65fb50ad11a8aba
    metrics:
    - type: accuracy
      value: 77.24806201550388
    - type: f1
      value: 60.531916308197175
  - task:
      type: Classification
    dataset:
      type: mteb/amazon_massive_intent
      name: MTEB MassiveIntentClassification (en)
      config: en
      split: test
      revision: 31efe3c427b0bae9c22cbb560b8f15491cc6bed7
    metrics:
    - type: accuracy
      value: 73.71553463349024
    - type: f1
      value: 71.70753174900791
  - task:
      type: Classification
    dataset:
      type: mteb/amazon_massive_scenario
      name: MTEB MassiveScenarioClassification (en)
      config: en
      split: test
      revision: 7d571f92784cd94a019292a1f45445077d0ef634
    metrics:
    - type: accuracy
      value: 77.79757901815736
    - type: f1
      value: 77.83719850433258
  - task:
      type: Clustering
    dataset:
      type: mteb/medrxiv-clustering-p2p
      name: MTEB MedrxivClusteringP2P
      config: default
      split: test
      revision: e7a26af6f3ae46b30dde8737f02c07b1505bcc73
    metrics:
    - type: v_measure
      value: 33.74193296622113
  - task:
      type: Clustering
    dataset:
      type: mteb/medrxiv-clustering-s2s
      name: MTEB MedrxivClusteringS2S
      config: default
      split: test
      revision: 35191c8c0dca72d8ff3efcd72aa802307d469663
    metrics:
    - type: v_measure
      value: 30.64257594108566
  - task:
      type: Reranking
    dataset:
      type: mteb/mind_small
      name: MTEB MindSmallReranking
      config: default
      split: test
      revision: 3bdac13927fdc888b903db93b2ffdbd90b295a69
    metrics:
    - type: map
      value: 30.811018518883625
    - type: mrr
      value: 31.910376577445003
  - task:
      type: Retrieval
    dataset:
      type: nfcorpus
      name: MTEB NFCorpus
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 5.409
    - type: map_at_10
      value: 13.093
    - type: map_at_100
      value: 16.256999999999998
    - type: map_at_1000
      value: 17.617
    - type: map_at_3
      value: 9.555
    - type: map_at_5
      value: 11.428
    - type: mrr_at_1
      value: 45.201
    - type: mrr_at_10
      value: 54.179
    - type: mrr_at_100
      value: 54.812000000000005
    - type: mrr_at_1000
      value: 54.840999999999994
    - type: mrr_at_3
      value: 51.909000000000006
    - type: mrr_at_5
      value: 53.519000000000005
    - type: ndcg_at_1
      value: 43.189
    - type: ndcg_at_10
      value: 35.028
    - type: ndcg_at_100
      value: 31.226
    - type: ndcg_at_1000
      value: 39.678000000000004
    - type: ndcg_at_3
      value: 40.596
    - type: ndcg_at_5
      value: 38.75
    - type: precision_at_1
      value: 44.582
    - type: precision_at_10
      value: 25.974999999999998
    - type: precision_at_100
      value: 7.793
    - type: precision_at_1000
      value: 2.036
    - type: precision_at_3
      value: 38.493
    - type: precision_at_5
      value: 33.994
    - type: recall_at_1
      value: 5.409
    - type: recall_at_10
      value: 16.875999999999998
    - type: recall_at_100
      value: 30.316
    - type: recall_at_1000
      value: 60.891
    - type: recall_at_3
      value: 10.688
    - type: recall_at_5
      value: 13.832
  - task:
      type: Retrieval
    dataset:
      type: nq
      name: MTEB NQ
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 36.375
    - type: map_at_10
      value: 51.991
    - type: map_at_100
      value: 52.91400000000001
    - type: map_at_1000
      value: 52.93600000000001
    - type: map_at_3
      value: 48.014
    - type: map_at_5
      value: 50.381
    - type: mrr_at_1
      value: 40.759
    - type: mrr_at_10
      value: 54.617000000000004
    - type: mrr_at_100
      value: 55.301
    - type: mrr_at_1000
      value: 55.315000000000005
    - type: mrr_at_3
      value: 51.516
    - type: mrr_at_5
      value: 53.435
    - type: ndcg_at_1
      value: 40.759
    - type: ndcg_at_10
      value: 59.384
    - type: ndcg_at_100
      value: 63.157
    - type: ndcg_at_1000
      value: 63.654999999999994
    - type: ndcg_at_3
      value: 52.114000000000004
    - type: ndcg_at_5
      value: 55.986000000000004
    - type: precision_at_1
      value: 40.759
    - type: precision_at_10
      value: 9.411999999999999
    - type: precision_at_100
      value: 1.153
    - type: precision_at_1000
      value: 0.12
    - type: precision_at_3
      value: 23.329
    - type: precision_at_5
      value: 16.256999999999998
    - type: recall_at_1
      value: 36.375
    - type: recall_at_10
      value: 79.053
    - type: recall_at_100
      value: 95.167
    - type: recall_at_1000
      value: 98.82
    - type: recall_at_3
      value: 60.475
    - type: recall_at_5
      value: 69.327
  - task:
      type: Retrieval
    dataset:
      type: quora
      name: MTEB QuoraRetrieval
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 70.256
    - type: map_at_10
      value: 83.8
    - type: map_at_100
      value: 84.425
    - type: map_at_1000
      value: 84.444
    - type: map_at_3
      value: 80.906
    - type: map_at_5
      value: 82.717
    - type: mrr_at_1
      value: 80.97999999999999
    - type: mrr_at_10
      value: 87.161
    - type: mrr_at_100
      value: 87.262
    - type: mrr_at_1000
      value: 87.263
    - type: mrr_at_3
      value: 86.175
    - type: mrr_at_5
      value: 86.848
    - type: ndcg_at_1
      value: 80.97999999999999
    - type: ndcg_at_10
      value: 87.697
    - type: ndcg_at_100
      value: 88.959
    - type: ndcg_at_1000
      value: 89.09899999999999
    - type: ndcg_at_3
      value: 84.83800000000001
    - type: ndcg_at_5
      value: 86.401
    - type: precision_at_1
      value: 80.97999999999999
    - type: precision_at_10
      value: 13.261000000000001
    - type: precision_at_100
      value: 1.5150000000000001
    - type: precision_at_1000
      value: 0.156
    - type: precision_at_3
      value: 37.01
    - type: precision_at_5
      value: 24.298000000000002
    - type: recall_at_1
      value: 70.256
    - type: recall_at_10
      value: 94.935
    - type: recall_at_100
      value: 99.274
    - type: recall_at_1000
      value: 99.928
    - type: recall_at_3
      value: 86.602
    - type: recall_at_5
      value: 91.133
  - task:
      type: Clustering
    dataset:
      type: mteb/reddit-clustering
      name: MTEB RedditClustering
      config: default
      split: test
      revision: 24640382cdbf8abc73003fb0fa6d111a705499eb
    metrics:
    - type: v_measure
      value: 56.322692497613104
  - task:
      type: Clustering
    dataset:
      type: mteb/reddit-clustering-p2p
      name: MTEB RedditClusteringP2P
      config: default
      split: test
      revision: 282350215ef01743dc01b456c7f5241fa8937f16
    metrics:
    - type: v_measure
      value: 61.895813503775074
  - task:
      type: Retrieval
    dataset:
      type: scidocs
      name: MTEB SCIDOCS
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 4.338
    - type: map_at_10
      value: 10.767
    - type: map_at_100
      value: 12.537999999999998
    - type: map_at_1000
      value: 12.803999999999998
    - type: map_at_3
      value: 7.788
    - type: map_at_5
      value: 9.302000000000001
    - type: mrr_at_1
      value: 21.4
    - type: mrr_at_10
      value: 31.637999999999998
    - type: mrr_at_100
      value: 32.688
    - type: mrr_at_1000
      value: 32.756
    - type: mrr_at_3
      value: 28.433000000000003
    - type: mrr_at_5
      value: 30.178
    - type: ndcg_at_1
      value: 21.4
    - type: ndcg_at_10
      value: 18.293
    - type: ndcg_at_100
      value: 25.274
    - type: ndcg_at_1000
      value: 30.284
    - type: ndcg_at_3
      value: 17.391000000000002
    - type: ndcg_at_5
      value: 15.146999999999998
    - type: precision_at_1
      value: 21.4
    - type: precision_at_10
      value: 9.48
    - type: precision_at_100
      value: 1.949
    - type: precision_at_1000
      value: 0.316
    - type: precision_at_3
      value: 16.167
    - type: precision_at_5
      value: 13.22
    - type: recall_at_1
      value: 4.338
    - type: recall_at_10
      value: 19.213
    - type: recall_at_100
      value: 39.562999999999995
    - type: recall_at_1000
      value: 64.08
    - type: recall_at_3
      value: 9.828000000000001
    - type: recall_at_5
      value: 13.383000000000001
  - task:
      type: STS
    dataset:
      type: mteb/sickr-sts
      name: MTEB SICK-R
      config: default
      split: test
      revision: a6ea5a8cab320b040a23452cc28066d9beae2cee
    metrics:
    - type: cos_sim_pearson
      value: 82.42568163642142
    - type: cos_sim_spearman
      value: 78.5797159641342
    - type: euclidean_pearson
      value: 80.22151260811604
    - type: euclidean_spearman
      value: 78.5797151953878
    - type: manhattan_pearson
      value: 80.21224215864788
    - type: manhattan_spearman
      value: 78.55641478381344
  - task:
      type: STS
    dataset:
      type: mteb/sts12-sts
      name: MTEB STS12
      config: default
      split: test
      revision: a0d554a64d88156834ff5ae9920b964011b16384
    metrics:
    - type: cos_sim_pearson
      value: 85.44020710812569
    - type: cos_sim_spearman
      value: 78.91631735081286
    - type: euclidean_pearson
      value: 81.64188964182102
    - type: euclidean_spearman
      value: 78.91633286881678
    - type: manhattan_pearson
      value: 81.69294748512496
    - type: manhattan_spearman
      value: 78.93438558002656
  - task:
      type: STS
    dataset:
      type: mteb/sts13-sts
      name: MTEB STS13
      config: default
      split: test
      revision: 7e90230a92c190f1bf69ae9002b8cea547a64cca
    metrics:
    - type: cos_sim_pearson
      value: 84.27165426412311
    - type: cos_sim_spearman
      value: 85.40429140249618
    - type: euclidean_pearson
      value: 84.7509580724893
    - type: euclidean_spearman
      value: 85.40429140249618
    - type: manhattan_pearson
      value: 84.76488289321308
    - type: manhattan_spearman
      value: 85.4256793698708
  - task:
      type: STS
    dataset:
      type: mteb/sts14-sts
      name: MTEB STS14
      config: default
      split: test
      revision: 6031580fec1f6af667f0bd2da0a551cf4f0b2375
    metrics:
    - type: cos_sim_pearson
      value: 83.138851760732
    - type: cos_sim_spearman
      value: 81.64101363896586
    - type: euclidean_pearson
      value: 82.55165038934942
    - type: euclidean_spearman
      value: 81.64105257080502
    - type: manhattan_pearson
      value: 82.52802949883335
    - type: manhattan_spearman
      value: 81.61255430718158
  - task:
      type: STS
    dataset:
      type: mteb/sts15-sts
      name: MTEB STS15
      config: default
      split: test
      revision: ae752c7c21bf194d8b67fd573edf7ae58183cbe3
    metrics:
    - type: cos_sim_pearson
      value: 86.0654695484029
    - type: cos_sim_spearman
      value: 87.20408521902229
    - type: euclidean_pearson
      value: 86.8110651362115
    - type: euclidean_spearman
      value: 87.20408521902229
    - type: manhattan_pearson
      value: 86.77984656478691
    - type: manhattan_spearman
      value: 87.1719947099227
  - task:
      type: STS
    dataset:
      type: mteb/sts16-sts
      name: MTEB STS16
      config: default
      split: test
      revision: 4d8694f8f0e0100860b497b999b3dbed754a0513
    metrics:
    - type: cos_sim_pearson
      value: 83.77823915496512
    - type: cos_sim_spearman
      value: 85.43566325729779
    - type: euclidean_pearson
      value: 84.5396956658821
    - type: euclidean_spearman
      value: 85.43566325729779
    - type: manhattan_pearson
      value: 84.5665398848169
    - type: manhattan_spearman
      value: 85.44375870303232
  - task:
      type: STS
    dataset:
      type: mteb/sts17-crosslingual-sts
      name: MTEB STS17 (en-en)
      config: en-en
      split: test
      revision: af5e6fb845001ecf41f4c1e033ce921939a2a68d
    metrics:
    - type: cos_sim_pearson
      value: 87.20030208471798
    - type: cos_sim_spearman
      value: 87.20485505076539
    - type: euclidean_pearson
      value: 88.10588324368722
    - type: euclidean_spearman
      value: 87.20485505076539
    - type: manhattan_pearson
      value: 87.92324770415183
    - type: manhattan_spearman
      value: 87.0571314561877
  - task:
      type: STS
    dataset:
      type: mteb/sts22-crosslingual-sts
      name: MTEB STS22 (en)
      config: en
      split: test
      revision: 6d1ba47164174a496b7fa5d3569dae26a6813b80
    metrics:
    - type: cos_sim_pearson
      value: 63.06093161604453
    - type: cos_sim_spearman
      value: 64.2163140357722
    - type: euclidean_pearson
      value: 65.27589680994006
    - type: euclidean_spearman
      value: 64.2163140357722
    - type: manhattan_pearson
      value: 65.45904383711101
    - type: manhattan_spearman
      value: 64.55404716679305
  - task:
      type: STS
    dataset:
      type: mteb/stsbenchmark-sts
      name: MTEB STSBenchmark
      config: default
      split: test
      revision: b0fddb56ed78048fa8b90373c8a3cfc37b684831
    metrics:
    - type: cos_sim_pearson
      value: 84.32976164578706
    - type: cos_sim_spearman
      value: 85.54302197678368
    - type: euclidean_pearson
      value: 85.26307149193056
    - type: euclidean_spearman
      value: 85.54302197678368
    - type: manhattan_pearson
      value: 85.26647282029371
    - type: manhattan_spearman
      value: 85.5316135265568
  - task:
      type: Reranking
    dataset:
      type: mteb/scidocs-reranking
      name: MTEB SciDocsRR
      config: default
      split: test
      revision: d3c5e1fc0b855ab6097bf1cda04dd73947d7caab
    metrics:
    - type: map
      value: 81.44675968318754
    - type: mrr
      value: 94.92741826075158
  - task:
      type: Retrieval
    dataset:
      type: scifact
      name: MTEB SciFact
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 56.34400000000001
    - type: map_at_10
      value: 65.927
    - type: map_at_100
      value: 66.431
    - type: map_at_1000
      value: 66.461
    - type: map_at_3
      value: 63.529
    - type: map_at_5
      value: 64.818
    - type: mrr_at_1
      value: 59.333000000000006
    - type: mrr_at_10
      value: 67.54599999999999
    - type: mrr_at_100
      value: 67.892
    - type: mrr_at_1000
      value: 67.917
    - type: mrr_at_3
      value: 65.778
    - type: mrr_at_5
      value: 66.794
    - type: ndcg_at_1
      value: 59.333000000000006
    - type: ndcg_at_10
      value: 70.5
    - type: ndcg_at_100
      value: 72.688
    - type: ndcg_at_1000
      value: 73.483
    - type: ndcg_at_3
      value: 66.338
    - type: ndcg_at_5
      value: 68.265
    - type: precision_at_1
      value: 59.333000000000006
    - type: precision_at_10
      value: 9.3
    - type: precision_at_100
      value: 1.053
    - type: precision_at_1000
      value: 0.11199999999999999
    - type: precision_at_3
      value: 25.889
    - type: precision_at_5
      value: 16.866999999999997
    - type: recall_at_1
      value: 56.34400000000001
    - type: recall_at_10
      value: 82.789
    - type: recall_at_100
      value: 92.767
    - type: recall_at_1000
      value: 99
    - type: recall_at_3
      value: 71.64399999999999
    - type: recall_at_5
      value: 76.322
  - task:
      type: PairClassification
    dataset:
      type: mteb/sprintduplicatequestions-pairclassification
      name: MTEB SprintDuplicateQuestions
      config: default
      split: test
      revision: d66bd1f72af766a5cc4b0ca5e00c162f89e8cc46
    metrics:
    - type: cos_sim_accuracy
      value: 99.75742574257426
    - type: cos_sim_ap
      value: 93.52081548447406
    - type: cos_sim_f1
      value: 87.33850129198966
    - type: cos_sim_precision
      value: 90.37433155080214
    - type: cos_sim_recall
      value: 84.5
    - type: dot_accuracy
      value: 99.75742574257426
    - type: dot_ap
      value: 93.52081548447406
    - type: dot_f1
      value: 87.33850129198966
    - type: dot_precision
      value: 90.37433155080214
    - type: dot_recall
      value: 84.5
    - type: euclidean_accuracy
      value: 99.75742574257426
    - type: euclidean_ap
      value: 93.52081548447406
    - type: euclidean_f1
      value: 87.33850129198966
    - type: euclidean_precision
      value: 90.37433155080214
    - type: euclidean_recall
      value: 84.5
    - type: manhattan_accuracy
      value: 99.75841584158415
    - type: manhattan_ap
      value: 93.4975678585854
    - type: manhattan_f1
      value: 87.26708074534162
    - type: manhattan_precision
      value: 90.45064377682404
    - type: manhattan_recall
      value: 84.3
    - type: max_accuracy
      value: 99.75841584158415
    - type: max_ap
      value: 93.52081548447406
    - type: max_f1
      value: 87.33850129198966
  - task:
      type: Clustering
    dataset:
      type: mteb/stackexchange-clustering
      name: MTEB StackExchangeClustering
      config: default
      split: test
      revision: 6cbc1f7b2bc0622f2e39d2c77fa502909748c259
    metrics:
    - type: v_measure
      value: 64.31437036686651
  - task:
      type: Clustering
    dataset:
      type: mteb/stackexchange-clustering-p2p
      name: MTEB StackExchangeClusteringP2P
      config: default
      split: test
      revision: 815ca46b2622cec33ccafc3735d572c266efdb44
    metrics:
    - type: v_measure
      value: 33.25569319007206
  - task:
      type: Reranking
    dataset:
      type: mteb/stackoverflowdupquestions-reranking
      name: MTEB StackOverflowDupQuestions
      config: default
      split: test
      revision: e185fbe320c72810689fc5848eb6114e1ef5ec69
    metrics:
    - type: map
      value: 49.90474939720706
    - type: mrr
      value: 50.568115503777264
  - task:
      type: Summarization
    dataset:
      type: mteb/summeval
      name: MTEB SummEval
      config: default
      split: test
      revision: cda12ad7615edc362dbf25a00fdd61d3b1eaf93c
    metrics:
    - type: cos_sim_pearson
      value: 29.866828641244712
    - type: cos_sim_spearman
      value: 30.077555055873866
    - type: dot_pearson
      value: 29.866832988572266
    - type: dot_spearman
      value: 30.077555055873866
  - task:
      type: Retrieval
    dataset:
      type: trec-covid
      name: MTEB TRECCOVID
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 0.232
    - type: map_at_10
      value: 2.094
    - type: map_at_100
      value: 11.971
    - type: map_at_1000
      value: 28.158
    - type: map_at_3
      value: 0.688
    - type: map_at_5
      value: 1.114
    - type: mrr_at_1
      value: 88
    - type: mrr_at_10
      value: 93.4
    - type: mrr_at_100
      value: 93.4
    - type: mrr_at_1000
      value: 93.4
    - type: mrr_at_3
      value: 93
    - type: mrr_at_5
      value: 93.4
    - type: ndcg_at_1
      value: 84
    - type: ndcg_at_10
      value: 79.923
    - type: ndcg_at_100
      value: 61.17
    - type: ndcg_at_1000
      value: 53.03
    - type: ndcg_at_3
      value: 84.592
    - type: ndcg_at_5
      value: 82.821
    - type: precision_at_1
      value: 88
    - type: precision_at_10
      value: 85
    - type: precision_at_100
      value: 63.019999999999996
    - type: precision_at_1000
      value: 23.554
    - type: precision_at_3
      value: 89.333
    - type: precision_at_5
      value: 87.2
    - type: recall_at_1
      value: 0.232
    - type: recall_at_10
      value: 2.255
    - type: recall_at_100
      value: 14.823
    - type: recall_at_1000
      value: 49.456
    - type: recall_at_3
      value: 0.718
    - type: recall_at_5
      value: 1.175
  - task:
      type: Retrieval
    dataset:
      type: webis-touche2020
      name: MTEB Touche2020
      config: default
      split: test
      revision: None
    metrics:
    - type: map_at_1
      value: 2.547
    - type: map_at_10
      value: 11.375
    - type: map_at_100
      value: 18.194
    - type: map_at_1000
      value: 19.749
    - type: map_at_3
      value: 5.825
    - type: map_at_5
      value: 8.581
    - type: mrr_at_1
      value: 32.653
    - type: mrr_at_10
      value: 51.32
    - type: mrr_at_100
      value: 51.747
    - type: mrr_at_1000
      value: 51.747
    - type: mrr_at_3
      value: 47.278999999999996
    - type: mrr_at_5
      value: 48.605
    - type: ndcg_at_1
      value: 29.592000000000002
    - type: ndcg_at_10
      value: 28.151
    - type: ndcg_at_100
      value: 39.438
    - type: ndcg_at_1000
      value: 50.769
    - type: ndcg_at_3
      value: 30.758999999999997
    - type: ndcg_at_5
      value: 30.366
    - type: precision_at_1
      value: 32.653
    - type: precision_at_10
      value: 25.714
    - type: precision_at_100
      value: 8.041
    - type: precision_at_1000
      value: 1.555
    - type: precision_at_3
      value: 33.333
    - type: precision_at_5
      value: 31.837
    - type: recall_at_1
      value: 2.547
    - type: recall_at_10
      value: 18.19
    - type: recall_at_100
      value: 49.538
    - type: recall_at_1000
      value: 83.86
    - type: recall_at_3
      value: 7.329
    - type: recall_at_5
      value: 11.532
  - task:
      type: Classification
    dataset:
      type: mteb/toxic_conversations_50k
      name: MTEB ToxicConversationsClassification
      config: default
      split: test
      revision: d7c0de2777da35d6aae2200a62c6e0e5af397c4c
    metrics:
    - type: accuracy
      value: 71.4952
    - type: ap
      value: 14.793362635531409
    - type: f1
      value: 55.204635551516915
  - task:
      type: Classification
    dataset:
      type: mteb/tweet_sentiment_extraction
      name: MTEB TweetSentimentExtractionClassification
      config: default
      split: test
      revision: d604517c81ca91fe16a244d1248fc021f9ecee7a
    metrics:
    - type: accuracy
      value: 61.5365025466893
    - type: f1
      value: 61.81742556334845
  - task:
      type: Clustering
    dataset:
      type: mteb/twentynewsgroups-clustering
      name: MTEB TwentyNewsgroupsClustering
      config: default
      split: test
      revision: 6125ec4e24fa026cec8a478383ee943acfbd5449
    metrics:
    - type: v_measure
      value: 49.05531070301185
  - task:
      type: PairClassification
    dataset:
      type: mteb/twittersemeval2015-pairclassification
      name: MTEB TwitterSemEval2015
      config: default
      split: test
      revision: 70970daeab8776df92f5ea462b6173c0b46fd2d1
    metrics:
    - type: cos_sim_accuracy
      value: 86.51725576682364
    - type: cos_sim_ap
      value: 75.2292304265163
    - type: cos_sim_f1
      value: 69.54022988505749
    - type: cos_sim_precision
      value: 63.65629110039457
    - type: cos_sim_recall
      value: 76.62269129287598
    - type: dot_accuracy
      value: 86.51725576682364
    - type: dot_ap
      value: 75.22922386081054
    - type: dot_f1
      value: 69.54022988505749
    - type: dot_precision
      value: 63.65629110039457
    - type: dot_recall
      value: 76.62269129287598
    - type: euclidean_accuracy
      value: 86.51725576682364
    - type: euclidean_ap
      value: 75.22925730473472
    - type: euclidean_f1
      value: 69.54022988505749
    - type: euclidean_precision
      value: 63.65629110039457
    - type: euclidean_recall
      value: 76.62269129287598
    - type: manhattan_accuracy
      value: 86.52321630804077
    - type: manhattan_ap
      value: 75.20608115037336
    - type: manhattan_f1
      value: 69.60000000000001
    - type: manhattan_precision
      value: 64.37219730941705
    - type: manhattan_recall
      value: 75.75197889182058
    - type: max_accuracy
      value: 86.52321630804077
    - type: max_ap
      value: 75.22925730473472
    - type: max_f1
      value: 69.60000000000001
  - task:
      type: PairClassification
    dataset:
      type: mteb/twitterurlcorpus-pairclassification
      name: MTEB TwitterURLCorpus
      config: default
      split: test
      revision: 8b6510b0b1fa4e4c4f879467980e9be563ec1cdf
    metrics:
    - type: cos_sim_accuracy
      value: 89.34877944657896
    - type: cos_sim_ap
      value: 86.71257569277373
    - type: cos_sim_f1
      value: 79.10386355986088
    - type: cos_sim_precision
      value: 76.91468470434214
    - type: cos_sim_recall
      value: 81.4213119802895
    - type: dot_accuracy
      value: 89.34877944657896
    - type: dot_ap
      value: 86.71257133133368
    - type: dot_f1
      value: 79.10386355986088
    - type: dot_precision
      value: 76.91468470434214
    - type: dot_recall
      value: 81.4213119802895
    - type: euclidean_accuracy
      value: 89.34877944657896
    - type: euclidean_ap
      value: 86.71257651501476
    - type: euclidean_f1
      value: 79.10386355986088
    - type: euclidean_precision
      value: 76.91468470434214
    - type: euclidean_recall
      value: 81.4213119802895
    - type: manhattan_accuracy
      value: 89.35848177901967
    - type: manhattan_ap
      value: 86.69330615469126
    - type: manhattan_f1
      value: 79.13867741453949
    - type: manhattan_precision
      value: 76.78881807647741
    - type: manhattan_recall
      value: 81.63689559593472
    - type: max_accuracy
      value: 89.35848177901967
    - type: max_ap
      value: 86.71257651501476
    - type: max_f1
      value: 79.13867741453949
license: apache-2.0
language:
- en
---


# nomic-embed-text-v1: A Reproducible Long Context (8192) Text Embedder 

`nomic-embed-text-v1` is 8192 context length text encoder that surpasses OpenAI text-embedding-ada-002 and text-embedding-3-small performance on short and long context tasks.



| Name                             | SeqLen | MTEB      | LoCo     | Jina Long Context |  Open Weights | Open Training Code | Open Data   |
| :-------------------------------:| :----- | :-------- | :------: | :---------------: | :-----------: | :----------------: | :---------- |
| nomic-embed-text-v1              | 8192   | **62.39** |**85.53** | 54.16             | ✅            | ✅                  | ✅          |
| jina-embeddings-v2-base-en       | 8192   | 60.39     | 85.45    | 51.90             | ✅            | ❌                  | ❌          |
| text-embedding-3-small           | 8191   | 62.26     | 82.40    | **58.20**         | ❌            | ❌                  | ❌          |
| text-embedding-ada-002           | 8191   | 60.99     | 52.7     | 55.25             | ❌            | ❌                  | ❌          |


## Hosted Inference API

The easiest way to get started with Nomic Embed is through the Nomic Embedding API.

Generating embeddings with the `nomic` Python client is as easy as 

```python
from nomic import embed

output = embed.text(
    texts=['Nomic Embedding API', '#keepAIOpen'],
    model='nomic-embed-text-v1',
    task_type='search_document'
)

print(output)
```

For more information, see the [API reference](https://docs.nomic.ai/reference/endpoints/nomic-embed-text)

## Data Visualization
Click the Nomic Atlas map below to visualize a 5M sample of our contrastive pretraining data!


[![image/webp](https://cdn-uploads.huggingface.co/production/uploads/607997c83a565c15675055b3/pjhJhuNyRfPagRd_c_iUz.webp)](https://atlas.nomic.ai/map/nomic-text-embed-v1-5m-sample)

## Training Details

We train our embedder using a multi-stage training pipeline. Starting from a long-context [BERT model](https://huggingface.co/nomic-ai/nomic-bert-2048),
the first unsupervised contrastive stage trains on a dataset generated from weakly related text pairs, such as question-answer pairs from forums like StackExchange and Quora, title-body pairs from Amazon reviews, and summarizations from news articles.

In the second finetuning stage, higher quality labeled datasets such as search queries and answers from web searches are leveraged. Data curation and hard-example mining is crucial in this stage.

For more details, see the Nomic Embed [Technical Report](https://static.nomic.ai/reports/2024_Nomic_Embed_Text_Technical_Report.pdf) and corresponding [blog post](https://blog.nomic.ai/posts/nomic-embed-text-v1).

Training data to train the models is released in its entirety. For more details, see the `contrastors` [repository](https://github.com/nomic-ai/contrastors)

## Usage

Note `nomic-embed-text` requires prefixes! We support the prefixes `[search_query, search_document, classification, clustering]`.
For retrieval applications, you should prepend `search_document` for all your documents and `search_query` for your queries.

### Sentence Transformers
```python
from sentence_transformers import SentenceTransformer

model = SentenceTransformer("nomic-ai/nomic-embed-text-v1", trust_remote_code=True)
sentences = ['search_query: What is TSNE?', 'search_query: Who is Laurens van der Maaten?']
embeddings = model.encode(sentences)
print(embeddings)
```

### Transformers

```python
import torch
import torch.nn.functional as F
from transformers import AutoTokenizer, AutoModel

def mean_pooling(model_output, attention_mask):
    token_embeddings = model_output[0]
    input_mask_expanded = attention_mask.unsqueeze(-1).expand(token_embeddings.size()).float()
    return torch.sum(token_embeddings * input_mask_expanded, 1) / torch.clamp(input_mask_expanded.sum(1), min=1e-9)

sentences = ['search_query: What is TSNE?', 'search_query: Who is Laurens van der Maaten?']

tokenizer = AutoTokenizer.from_pretrained('bert-base-uncased')
model = AutoModel.from_pretrained('nomic-ai/nomic-embed-text-v1', trust_remote_code=True)
model.eval()

encoded_input = tokenizer(sentences, padding=True, truncation=True, return_tensors='pt')

with torch.no_grad():
    model_output = model(**encoded_input)

embeddings = mean_pooling(model_output, encoded_input['attention_mask'])
embeddings = F.normalize(embeddings, p=2, dim=1)
print(embeddings)
```

The model natively supports scaling of the sequence length past 2048 tokens. To do so, 

```diff
- tokenizer = AutoTokenizer.from_pretrained('bert-base-uncased')
+ tokenizer = AutoTokenizer.from_pretrained('bert-base-uncased', model_max_length=8192)


- model = AutoModel.from_pretrained('nomic-ai/nomic-embed-text-v1', trust_remote_code=True)
+ model = AutoModel.from_pretrained('nomic-ai/nomic-embed-text-v1', trust_remote_code=True, rotary_scaling_factor=2)
```

### Transformers.js

```js
import { pipeline } from '@xenova/transformers';

// Create a feature extraction pipeline
const extractor = await pipeline('feature-extraction', 'nomic-ai/nomic-embed-text-v1', {
    quantized: false, // Comment out this line to use the quantized version
});

// Compute sentence embeddings
const texts = ['search_query: What is TSNE?', 'search_query: Who is Laurens van der Maaten?'];
const embeddings = await extractor(texts, { pooling: 'mean', normalize: true });
console.log(embeddings);
```

# Join the Nomic Community

- Nomic: [https://nomic.ai](https://nomic.ai)
- Discord: [https://discord.gg/myY5YDR8z8](https://discord.gg/myY5YDR8z8)
- Twitter: [https://twitter.com/nomic_ai](https://twitter.com/nomic_ai)


# Citation

If you find the model, dataset, or training code useful, please cite our work

```bibtex
@misc{nussbaum2024nomic,
      title={Nomic Embed: Training a Reproducible Long Context Text Embedder}, 
      author={Zach Nussbaum and John X. Morris and Brandon Duderstadt and Andriy Mulyar},
      year={2024},
      eprint={2402.01613},
      archivePrefix={arXiv},
      primaryClass={cs.CL}
}
```