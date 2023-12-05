---
license: llama2
datasets:
- snow_simplified_japanese_corpus
- khalidalt/tydiqa-goldp
- csebuetnlp/xlsum
language:
- ja
---
# About
This model is Lightblue's QLoRA finetune of OpenOrca's [Open-Orca/OpenOrcaxOpenChat-Preview2-13B](https://huggingface.co/Open-Orca/OpenOrcaxOpenChat-Preview2-13B) model on Japanese fine-tuning datasets.

This model specialises on answering **Closed Question Answering** in Japanese. Input a piece of reference text, ask a question, and see the model answer based on the reference text.

We trained on equal samples of the following three datasets:
* [SNOW](https://huggingface.co/datasets/snow_simplified_japanese_corpus) 
* [TyDiQA (Ja)](https://huggingface.co/datasets/khalidalt/tydiqa-goldp) 
* [XLSUM (Ja)](https://huggingface.co/datasets/csebuetnlp/xlsum) 

which resulted in a dataset of 13,167 samples total.

These three datasets were chosen as they represent three distinct fine-tuning tasks (Text simplification, question answering, and text summarization, respectively) which we hypothesize can help to improve the language models suitability for dealing with Japanese data.
These three datasets make up the model name: STX.

With these datasets, we achieve the following scores on the JGLUE benchmark:

| Model Name             | Open-Orca/OpenOrcaxOpenChat-Preview2-13B | lightblue/openorca_stx |
|------------------------|------------------------------------------|------------------------|
| jsquad-1.1-0.3         | 0.692                                    | 0.836                  |
| jcommonsenseqa-1.1-0.3 | 0.831                                    | 0.782                  |
| jnli-1.1-0.3           | 0.504                                    | 0.48                   |
| marc_ja-1.1-0.3        | 0.936                                    | 0.959                  |


We achieved these scores by using the [lm-evaluation-harness](https://github.com/Stability-AI/lm-evaluation-harness) from Stability AI using the below commands:
```bash
MODEL_ARGS=pretrained=lightblue/openorca_stx,use_accelerate=True
TASK="jsquad-1.1-0.3,jcommonsenseqa-1.1-0.3,jnli-1.1-0.3,marc_ja-1.1-0.3"
export JGLUE_OUTPUT_DIR=../jglue_results/$MODEL_NAME/$DATSET_NAME/$DATASET_SIZE
mkdir -p $JGLUE_OUTPUT_DIR
python main.py --model hf-causal-experimental --model_args $MODEL_ARGS --tasks $TASK --num_fewshot "2,3,3,3" --device "cuda" --output_path $JGLUE_OUTPUT_DIR/result.json --batch_size 4 > $JGLUE_OUTPUT_DIR/harness.out 2> $JGLUE_OUTPUT_DIR/harness.err
```

Our model achieves much better results on the question answering benchmark (JSQuAD) than the base checkpoint without monstrous degradation of performance on multi-choice question benchmarks (JCommonSense, JNLI, MARC-Ja) purely through QLoRA training.
This shows the potential for applying strong language models such as [Open-Orca/OpenOrcaxOpenChat-Preview2-13B](https://huggingface.co/Open-Orca/OpenOrcaxOpenChat-Preview2-13B) to minimal QLoRA fine-tuning using Japanese fine-tuning datasets to achieve better results at narrow NLP tasks.

# How to use

```python
from transformers import AutoTokenizer, AutoModelForCausalLM, pipeline

model_dir = "lightblue/openorca_stx"

tokenizer = AutoTokenizer.from_pretrained(model_dir)
model = AutoModelForCausalLM.from_pretrained(
    model_dir, torch_dtype=torch.bfloat16, device_map='auto',
)

pipe = pipeline("text-generation", model=model, tokenizer=tokenizer)

def do_closed_qa(context, question):
    return context + "\n\n" + question

test_article = """　モノマネのレパートリーに「リーチ・マイケル選手」があるレイザーラモンRGさん。本人公認のモノマネですが、ラグビーファンの反応に少し驚いたそうです。
　リーチ・マイケル選手のモノマネは、何がきっかけですか。
「2015年のワールドカップ（W杯）イングランド大会で日本が南アフリカを倒した次の日が、京都での番組ロケでした。当時は、アップルの共同創業者スティーブ・ジョブズのモノマネばかりでしたが、一緒にロケをしていたジャングルポケットから『リーチ・マイケルに似てますよ。ジョブズのまま、いけるんじゃないですか？』と言われたのが始まりです」
「ただ、みんな知識がない。ラグビーショップを探し、日本代表のユニホームが売り切れだったので、赤っぽいユニホームとピチピチの短パンをはいて。とりあえずSNSで『リーチ・マイケルです』っていっぱい写真を載せました」
「すると、それを見たリーチさん本人からDM（ダイレクトメッセージ）が届きました。『モノマネありがとうございます。もしモノマネをするなら、僕のユニホームを送りますので着てください』と。W杯後にユニホーム2着とパンツやソックスなどをほんまに送ってきてくれました。今着ているのがそれです」
これまで、数々の著名人をモノマネしてこられました。リーチ選手のネタの反響はいかがでしたか。
　「僕はラグビー経験がないですし、ラグビーを全然知らなかったけど、やっぱり本人からユニホームを頂いてるっていう“印籠（いんろう）”みたいなのがあって。『あいつはリーチさん本人に認められてる』と。一目置かれているのかなと感じます」
　「やっていることは、見た目を本人に寄せてワンチームって言うだけなんですけどね。それでも『わあ、リーチさんだ』と言ってもらえます」
　「リーチさんと実際に会うことなんて、簡単にはできないじゃないですか。でも、リーチさんのまねをしているRGには会えたわ、みたいな（笑）。何だろうな、有名な神社の支社のような存在ですかね。ありがたがられるという意味では他のモノマネとはすごく違いますね」
"""

test_question = "　リーチ・マイケルは何を送ってきましたか？"

pipe(do_closed_qa(test_article, question), max_new_tokens=128, temperature=0)[0]["generated_text"]
# "ユニホーム2着とパンツやソックスなど"
```

### Prompting

We have found that this model is able to work well using a variety of prompts, including the Alpaca style templated prompts:

```python

f"""
Below is an instruction that describes a task, paired with an input that provides further context. Write a response that appropriately completes the request.
### Instruction:
{instruction}
### Input:
{input}
### Response:
"""

```

We have found that having a newline at the end of the prompt can be important for signalling that the model must respond and not continue the inputs.


# Training details

We trained using the following three minimalistic prompt templates for the three tasks in STX:

* SNOW
  ```python
  f"""元の日本語：
  {original_ja}

  シンプルな日本語："""
  ```
* TyDiQA
  ```python
  f"""{passage_text}

  {question_text}"""
  ```
* XLSum
  ```python
  f"""記事：
  {article_text}

  要約："""
  ```

This model was trained for 1000 steps (1.2 epochs) with the model being evaluated every 50 steps. We then chose the best model from these evaluations based on validation loss.
We used the [qlora](https://github.com/artidoro/qlora) package from artidoro.
We trained with the following hyperparameters:

```
Per device evaluation batch size: 16
Per device train batch size: 8
LoRA (lora_r): 64
LoRA alpha (lora_alpha): 16
LoRA modules: all
Double quantization: Enabled
Quantization type: nf4
BF16: Enabled
Bits: 4
Warmup ratio: 0.03
Learning rate scheduler type: Constant
Gradient checkpointing: Enabled
Gradient accumulation steps: 2
Learning rate: 0.0002
Adam beta2: 0.999
Maximum gradient norm: 0.3
LoRA dropout: 0.05
Weight decay: 0.0
```

![image/png](https://cdn-uploads.huggingface.co/production/uploads/64b63f8ad57e02621dc93c8b/UWiE7z5tG8t_vdSFrb5WC.png)

![image/png](https://cdn-uploads.huggingface.co/production/uploads/64b63f8ad57e02621dc93c8b/_fKBf9sdq9UAKKYMxM6ad.png)