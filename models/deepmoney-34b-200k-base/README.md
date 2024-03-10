---
license: apache-2.0
language:
- en
- zh
tags:
- finance
- invest
---
# **Deepmoney**

![767e2d3bba166cd63a83ae54e913d35.jpg](https://cdn-uploads.huggingface.co/production/uploads/630c1adea20a5367812196f6/O0kFm05ZSe6Lw6FhGwx5_.jpeg)

***Sft Versions:***

[67b Chat](https://huggingface.co/TriadParty/deepmoney-67b-chat)
[34b Chat](https://huggingface.co/TriadParty/deepmoney-34b-200k-chat-evaluator)


Introducing **Greed** in the Seven Deadly Sins series of models.

- Full-para pre-training on Yi-34b
- High-quality research reports
- High-end cleaning process

### 1. What do I want to do?

Most of the current so-called financial models are mostly trained on public knowledge, but in the actual financial field, these public knowledge are often seriously insufficient for the current market interpretability. If you are interested, you can learn about the various propositions of Keynes, Friedman and even current behavioral finance. According to my observation, most financial models cannot make investment judgments. Because they are trained on ordinary textbooks, entry-level analyst exams, and even company public reports. I think this is of very little value for the investment.

You can think I'm joking, but the fact is that the logic of many subjective analysts may not be as rigorous as that of large models of 34b and above (excluding those excellent ones, of course). The market is changing every moment, with a lot of news and massive data in real time. For most retail investors, instead of waiting for a crappy analyst to write a report, why not use a large model to make a pipeline?

In my plan, this model is the base model of this process. In my plan, models such as information collector, target judge, qualitative analyst, quantitative analyst, and data extractor are all part of this process. . But the model itself is undoubtedly important to master a large number of qualitative and quantitative methods. That's why this model was born.

### 2. About the data

As I just said, a lot of public knowledge has some questionable validity - but that doesn't mean it's wrong. The theoretical support behind many research methods in research reports also relies on this knowledge. So in my training, I picked up some college textbooks and some professional books. Not a lot of quantity but good quality. In addition, I selected a large number of research report data from 2019 to December 2023 - these reports are issued by a variety of publishers, including traditional brokers and professional research institutions. Most of them are paid and only available to institutions. But I got them anyway through various means. 

If you have read research reports, especially high-quality ones, you will find that research reports are all subjective judgment + quantitative analysis, and data support in quantitative analysis is crucial to the entire logical chain. In order to extract this data (most of them are in the form of graphs or tables), I tried a lot of multi-modal models, and the process was very painful. The conclusion is that cog-agent and emu2 are very effective for this kind of tasks. In order to better extract information, I created a process that summarizes the context of research reports as part of the prompt. 

Finally, I made a blend of the data. General data is not included because it is just for greed. Moreover, the knowledge in industry research reports is comprehensive enough.

### 3. About training

Raw text, full parameter training. The base uses long context yi-34b-200k. This is necessary to complete and understand an in-depth report. 

Of course, I also did a sft. [This](https://huggingface.co/TriadParty/deepmoney-34b-200k-chat-evaluator) is the analyzer in my process – I haven’t broken down the qualitative and quantitative analysis yet, but I’m already blown away by how well it works.

### 4. Evals

Let us sample some recent events, simulate the event-driven securities analysis process in the real world, and conduct a comparative test against deepmoney and gpt4. Because the impact of events on the market is relatively invisible, it is difficult to evaluate the effect without a rigorous backtesting process. And our output needs to be analyzed using many quantitative methods. So I post the results here so that everyone can make a perceptual evaluation of the production results. I have a global news crawling system, and there are a lot of news every moment. 

In my process, I deduplicate and make subjective and objective judgments on these news. These can be solved by traditional BERT. Then for deepmoney, three steps are used to process the incoming news: 

1.  Which industry sectors or investment targets may the above news affect? 

2.  Please design a quantitative method to study the impact of the above news on  industry. And accordingly explain what data needs to be used. 

3.  Based on the following data, please _____ design a set of specific quantitative methods to quantitatively analyze the impact of the above news on  industry.

   

   Among them, the first problem is subjective judgment and extracting the target of news impact. This relies more on the subjective analysis capabilities of the model. Then extract the industry name from the first answer (for those familiar with large models, it is easy to design an automated process) and fill it in the second question, with the purpose of obtaining data for quantitative analysis. The reason why we ask about quantitative methods first and then data is the magic of COT. The answer to the last question is what we really need. There is enough information in the context of this question, and it needs to be answered with an exact and specific quantitative method. Combining the code writing model and the function calling model, if you have a macro and micro database with a complete data dictionary, this is completely achievable. Here is the three-step answer for deepmoney and gpt4:

   

   **Deepmoney**:

```
(1). 
近日，英国药企葛兰素史克（GSK.US）宣布，以10亿美元的首付款，以及4亿美元的里程碑付款，收购美国一家临床阶段生物制药服务商Aiolos Bio。GSK瞄准的，是Aiolos Bio唯一的核心管线——用于治疗成人哮喘的单克隆抗体AIO-001。IO-001管线并非Aiolos Bio自研，它授权自中国医药一哥——恒瑞医药。医药分析师称，总体来看，此次的AIO-001单抗交易中，所谓“中间商赚差价”的说法有媒体炒作嫌疑，对“恒瑞医药轻易出手”的观点，他表示：“一个（Aiolos Bio）是赌博，下注，几年吃一次大单；一个（恒瑞医药）是制造业顺便做海外增量。大家做的完全不是一个生意，恒瑞依然是赚钱方。”（澎湃新闻）
Which industry sectors or investment targets may the above news affect? 

ANS:
The above news primarily affects the pharmaceutical, biotechnology, and healthcare industries as it involves an acquisition by GSK of Aiolos Bio, which is focused on developing treatments for asthma using a single monoclonal antibody called AIO-001. The deal also has implications for companies in these sectors that could be involved in similar deals or collaborations to expand their pipelines and market presence.
******************************************************************************
(2). 近日，英国药企葛兰素史克（GSK.US）宣布，以10亿美元的首付款，以及4亿美元的里程碑付款，收购美国一家临床阶段生物制药服务商Aiolos Bio。GSK瞄准的，是Aiolos Bio唯一的核心管线——用于治疗成人哮喘的单克隆抗体AIO-001。IO-001管线并非Aiolos Bio自研，它授权自中国医药一哥——恒瑞医药。医药分析师称，总体来看，此次的AIO-001单抗交易中，所谓“中间商赚差价”的说法有媒体炒作嫌疑，对“恒瑞医药轻易出手”的观点，他表示：“一个（Aiolos Bio）是赌博，下注，几年吃一次大单；一个（恒瑞医药）是制造业顺便做海外增量。大家做的完全不是一个生意，恒瑞依然是赚钱方。”（澎湃新闻）

Please design a quantitative method to study the impact of the above news on the pharmaceutical industry. And accordingly explain what data needs to be used.

ANS:
To quantitatively analyze the impact of this acquisition news on the pharmaceutical industry, we can use regression analysis or other statistical modeling techniques. The key is to identify independent variables that might influence stock prices in the pharmaceutical sector and then determine how much each variable affects changes in stock price.

Here's an overview of the steps involved in designing such a quantitative approach:

Step 1: Data Collection
Collect historical data for the period you wish to study (e.g., quarterly or yearly) from various sources like Bloomberg, Yahoo Finance, etc. Here's a list of some potential datasets needed:

* Stock market indices related to pharmaceutical industries both locally (e.g., S&P Pharmaceutical Index) and globally (e.gs&p Global Healthcare Index)
* Financial performance indicators of major players within the same space as GSK/Aiolos Bio and those affected by their acquisitions
* News sentiment scores around similar deals in the past
* Macroeconomic factors like GDP growth rate, inflation rates, interest rates, etc.
* Industry specific factors like R&D spending trends, regulatory environment changes, patent expirations, etc.

Step 2: Model Selection
After collecting all relevant information, choose suitable econometric models depending upon the nature of relationships between dependent & independent variables. Ordinary Least Squares Regression would work if linearity assumption holds true; otherwise, Logit/Probit model could be considered too.

Step 3: Variable selection
Selecting appropriate explanatory variables after checking multicollinearity among them using VIF values. This will help us understand which variables have significant effects on stock returns.

Step 4: Hypothesis testing
Perform hypothesis tests like F-test, t-test, Durbin-Watson test, etc., to check whether our chosen regressors are statistically significant at different levels of significance.

Step 5: Coefficient interpretation
Interpret coefficients obtained from Step 4 while keeping mind signs and magnitudes must align with economic theory. Positive coefficient means positive relationship whereas negative implies inverse relation.

Step 6: Model evaluation
Evaluate goodness-of-fit measures like R^2, Adjusted R^2, Akaike Information Criterion(AIC), Bayesian Information Criterion(BIC), Root Mean Square Error(RMSE), Mean Absolute Deviation(MAD),etc.,to ascertain model fitness.

Step 7: Forecasting
Use estimated parameters along with future exogenous variable values to predict future outcomes for any given time period.

In conclusion, this quantitative methodology aims to measure the effect of GSKS's acquisition on the overall pharmaceutical industry by studying its impacts on stock prices and financial performances across several companies in the field.
****************************************************************************
(3). 近日，英国药企葛兰素史克（GSK.US）宣布，以10亿美元的首付款，以及4亿美元的里程碑付款，收购美国一家临床阶段生物制药服务商Aiolos Bio。GSK瞄准的，是Aiolos Bio唯一的核心管线——用于治疗成人哮喘的单克隆抗体AIO-001。IO-001管线并非Aiolos Bio自研，它授权自中国医药一哥——恒瑞医药。医药分析师称，总体来看，此次的AIO-001单抗交易中，所谓“中间商赚差价”的说法有媒体炒作嫌疑，对“恒瑞医药轻易出手”的观点，他表示：“一个（Aiolos Bio）是赌博，下注，几年吃一次大单；一个（恒瑞医药）是制造业顺便做海外增量。大家做的完全不是一个生意，恒瑞依然是赚钱方。”（澎湃新闻）

Based on the following data, please design a specific set of quantitative methods to quantitatively analyze the impact of the above news on the pharmaceutical industry.

* Stock market indices related to pharmaceutical industries both locally (e.g., S&P Pharmaceutical Index) and globally (e.gs&p Global Healthcare Index)
* Financial performance indicators of major players within the same space as GSK/Aiolos Bio and those affected by their acquisitions
* News sentiment scores around similar deals in the past
* Macroeconomic factors like GDP growth rate, inflation rates, interest rates, etc.
* Industry specific factors like R&D spending trends, regulatory environment changes, patent expirations, etc.

ANS:
To quantitatively analyze the impact of this news on the pharmaceutical industry, we can use several quantitative methods that involve analyzing stock prices, financial performance indicators, news sentiment scores, macroeconomic factors, and industry-specific factors. Here is an outline of how you might go about it:

1. Stock Market Indices Analysis: Monitor daily or weekly fluctuations in local (e.g., S&P Pharmaceutical Index) and global (e.g., s&p Global Healthcare Index) stock market indices related to the pharmaceutical industry before, during, and after the announcement of the acquisition. Analyze these fluctuations using statistical techniques such as regression analysis, ANOVA testing, or time series analysis to determine whether there were significant price movements due to this event.
2. Financial Performance Indicators: Examine the financial performance metrics of other companies operating in the same therapeutic area(s) as Aiolos Bio and GSK prior to the acquisition announcement. These could include revenue growth rates, net income margins, return on equity (ROE) ratios, earnings per share (EPS), etc. Then compare them with post-acquisition figures for any noticeable differences that may have resulted from the deal. You can also look at the overall profitability trend among firms within the same space to gauge its potential effect on investor confidence.
3. News Sentiment Scores: Collect media coverage surrounding previous mergers & acquisitions (M&As) involving biotech startups similar to Aiolos Bio over a certain period (e..g 5 years). Use natural language processing tools or machine learning algorithms to extract positive/negative sentiments associated with each article. Compare these scores against historical averages for M&As within the pharma sector; if they deviate significantly, it could indicate heightened investor optimism or pessimism towards future deals.
4. Macroeconomic Factors: Track key economic indicators such as GDP growth rates, inflation rates, interest rates, exchange rates, unemployment levels, etc., across different countries where major players in the industry operate. Assess correlations between these variables and changes observed in pharmaceutical stocks' performances to understand their sensitivity to broader economic conditions.
5. Industry Specific Factors: Keep track of recent developments in the healthcare industry, such as shifts in regulatory environments, patent expirations, advancements in drug discovery technologies, and pricing policies. Evaluate the extent to which these events influence valuations of companies involved in similar transactions and identify common patterns affecting the entire sector.
6. Regression Modeling: Develop multiple linear regressions incorporating all identified drivers mentioned above as independent variables along with dependent variable being either individual company stock returns or index level changes representing the whole pharmaceutical sector. Perform backward elimination technique to arrive at most parsimonious model explaining maximum variation in dependent variable while keeping only statistically significant predictors. This will help us isolate unique contribution made by each factor including the aforementioned news event.
7. Monte Carlo Simulation: Lastly run Monte Carlo simulations based on estimated coefficients from finalized regression models by generating thousands of random samples reflecting distributional assumptions made about each driver variable. By repeating process many times, one can obtain probability distributions for predicted outcomes under various scenarios considering uncertainty inherent in estimating parameters accurately.
8. Conclusion: Based on results obtained from analyses described above, draw inferences regarding likely impacts of said transaction on valuation multiples of comparable companies operating within same space or even entire sectors. Make recommendations concerning investment decisions taking into account risk appetite and available information sources used throughout investigation process.
```

**gpt4**

```
(1). 
近日，英国药企葛兰素史克（GSK.US）宣布，以10亿美元的首付款，以及4亿美元的里程碑付款，收购美国一家临床阶段生物制药服务商Aiolos Bio。GSK瞄准的，是Aiolos Bio唯一的核心管线——用于治疗成人哮喘的单克隆抗体AIO-001。IO-001管线并非Aiolos Bio自研，它授权自中国医药一哥——恒瑞医药。医药分析师称，总体来看，此次的AIO-001单抗交易中，所谓“中间商赚差价”的说法有媒体炒作嫌疑，对“恒瑞医药轻易出手”的观点，他表示：“一个（Aiolos Bio）是赌博，下注，几年吃一次大单；一个（恒瑞医药）是制造业顺便做海外增量。大家做的完全不是一个生意，恒瑞依然是赚钱方。”（澎湃新闻）
Which industry sectors or investment targets may the above news affect? 

ANS:
The news about GlaxoSmithKline (GSK) acquiring Aiolos Bio, which involves a monoclonal antibody therapy (AIO-001) licensed from Hengrui Medicine, could potentially affect several industry sectors and investment targets:

Pharmaceutical and Biotech Sector: The acquisition directly impacts the pharmaceutical and biotechnology industries. Companies involved in monoclonal antibody therapies, particularly those focusing on respiratory diseases like asthma, could see increased interest.

Healthcare Investment: This move might influence investment strategies within the healthcare sector, especially in companies developing similar treatments or engaging in innovative research.

Chinese Pharma Companies: Given that AIO-001 was licensed from Hengrui Medicine, a leading Chinese pharmaceutical company, this deal might draw attention to Chinese pharma companies as potential partners or targets for licensing agreements.

Clinical-stage Biopharmaceuticals: The acquisition highlights the value of clinical-stage biopharmaceutical companies, particularly those with promising drug candidates in their pipeline.

Respiratory Disease Therapeutics: Companies developing treatments for respiratory diseases, especially asthma, might gain investor interest as this acquisition underscores the market potential in this area.

Cross-border Healthcare Deals: This deal reflects the growing trend of cross-border collaborations and acquisitions in the healthcare sector. Companies that facilitate or stand to benefit from such international partnerships could be impacted.

Healthcare Licensing and IP Management: The role of licensing agreements and intellectual property in pharmaceutical collaborations might receive more focus, affecting companies that specialize in these areas.

Investors and companies in these sectors should closely monitor developments related to this acquisition for potential opportunities and impacts on their strategies and operations.
******************************************************************************
(2). 近日，英国药企葛兰素史克（GSK.US）宣布，以10亿美元的首付款，以及4亿美元的里程碑付款，收购美国一家临床阶段生物制药服务商Aiolos Bio。GSK瞄准的，是Aiolos Bio唯一的核心管线——用于治疗成人哮喘的单克隆抗体AIO-001。IO-001管线并非Aiolos Bio自研，它授权自中国医药一哥——恒瑞医药。医药分析师称，总体来看，此次的AIO-001单抗交易中，所谓“中间商赚差价”的说法有媒体炒作嫌疑，对“恒瑞医药轻易出手”的观点，他表示：“一个（Aiolos Bio）是赌博，下注，几年吃一次大单；一个（恒瑞医药）是制造业顺便做海外增量。大家做的完全不是一个生意，恒瑞依然是赚钱方。”（澎湃新闻）

Please design a quantitative method to study the impact of the above news on the pharmaceutical industry. And accordingly explain what data needs to be used.

ANS:

To study the impact of the acquisition of Aiolos Bio by GlaxoSmithKline (GSK) on the pharmaceutical industry, we can design a quantitative method that encompasses several key aspects. Here's a structured approach:

1. Event Study Methodology:
Objective: Analyze the stock market reaction to the acquisition announcement.
Data Needed:
Daily stock prices of GSK, Aiolos Bio, Hengrui Medicine, and other relevant pharmaceutical companies.
Market indices (e.g., S&P 500, NASDAQ Biotechnology Index) for the same period.
Analysis:
Calculate abnormal returns (AR) and cumulative abnormal returns (CAR) for a window period around the announcement date.
Compare the performance of GSK and other companies with the overall market.
2. Pipeline Impact Analysis:
Objective: Assess the strategic value of AIO-001 in GSK’s and Hengrui Medicine's pipeline.
Data Needed:
Information on AIO-001’s clinical trial stages, success rates, and market potential.
GSK’s and Hengrui Medicine's current pipeline portfolio data.
Analysis:
Project future revenue streams from AIO-001.
Evaluate the diversification and strategic fit of AIO-001 in the acquiring company’s pipeline.
3. Industry Sentiment Analysis:
Objective: Gauge the industry's reaction to the deal.
Data Needed:
News articles, press releases, and analyst reports pre and post-announcement.
Social media sentiment analysis.
Analysis:
Use natural language processing (NLP) techniques to analyze sentiment.
Identify major themes and concerns raised by industry stakeholders.
4. Competitor Response Analysis:
Objective: Understand how competitors might react.
Data Needed:
Market share, R&D spending, and strategic moves of key competitors.
Industry trends and innovation metrics.
Analysis:
Identify any shifts in competitive strategies post-announcement.
Assess the potential for new partnerships or M&A activities.
5. Regulatory Impact Assessment:
Objective: Evaluate the regulatory implications of the acquisition.
Data Needed:
Regulatory filings and approvals related to the deal.
Historical data on similar acquisitions and their regulatory outcomes.
Analysis:
Assess the likelihood of regulatory challenges.
Understand the impact of regulatory changes on the deal's success.
Data Sources:
Financial data from databases
like Bloomberg, Yahoo Finance, or Reuters.

Clinical trial data from public databases like ClinicalTrials.gov or company reports.
Industry reports and publications for market trends and analyses.
Social media data and online news archives for sentiment analysis.
Regulatory filings from government websites like the FDA or EMA.
Conclusion:
This quantitative method combines financial analysis, strategic evaluation, sentiment analysis, competitive landscape assessment, and regulatory considerations. It provides a comprehensive view of the acquisition's impact on GSK, Aiolos Bio, Hengrui Medicine, and the broader pharmaceutical industry. By leveraging diverse data sources and analytical techniques, stakeholders can make informed decisions and strategize effectively in response to this significant industry event.

****************************************************************************
(3). 近日，英国药企葛兰素史克（GSK.US）宣布，以10亿美元的首付款，以及4亿美元的里程碑付款，收购美国一家临床阶段生物制药服务商Aiolos Bio。GSK瞄准的，是Aiolos Bio唯一的核心管线——用于治疗成人哮喘的单克隆抗体AIO-001。IO-001管线并非Aiolos Bio自研，它授权自中国医药一哥——恒瑞医药。医药分析师称，总体来看，此次的AIO-001单抗交易中，所谓“中间商赚差价”的说法有媒体炒作嫌疑，对“恒瑞医药轻易出手”的观点，他表示：“一个（Aiolos Bio）是赌博，下注，几年吃一次大单；一个（恒瑞医药）是制造业顺便做海外增量。大家做的完全不是一个生意，恒瑞依然是赚钱方。”（澎湃新闻）

Based on the following data, please design a specific set of quantitative methods to quantitatively analyze the impact of the above news on the pharmaceutical industry.

Financial data from databases like Bloomberg, Yahoo Finance, or Reuters.
Clinical trial data from public databases like ClinicalTrials.gov or company reports.
Industry reports and publications for market trends and analyses.
Social media data and online news archives for sentiment analysis.
Regulatory filings from government websites like the FDA or EMA.

ANS:

To quantitatively analyze the impact of the news about GSK's acquisition of Aiolos Bio and its implications for the pharmaceutical industry, we can design a multi-faceted approach that incorporates different data sources and analytical techniques. Here's a structured method:

Financial Data Analysis:

Stock Performance: Use Bloomberg, Yahoo Finance, or Reuters to track the stock prices of GSK, Aiolos Bio, and Hengrui Medicine before and after the announcement. Analyze the volatility, trading volume, and any abnormal returns.
Financial Health Metrics: Examine key financial metrics like P/E ratio, EBITDA, revenue growth, and R&D spending to assess the financial health of these companies.
Peer Comparison: Compare these metrics with other companies in the pharmaceutical sector to understand the broader market reaction.
Clinical Trial Data Evaluation:

Pipeline Analysis: Use ClinicalTrials.gov or company reports to examine the details of AIO-001’s clinical trials, including phase, enrollment numbers, and study outcomes.
Comparative Efficacy and Safety: Compare AIO-001's clinical trial data with existing asthma treatments to evaluate its potential market impact.
Development Timeline: Assess the estimated time to market for AIO-001 and potential regulatory hurdles.
Market Trend Analysis:

Industry Reports: Utilize industry reports to understand current trends in asthma treatment and monoclonal antibody therapies.
Market Size and Growth Projections: Analyze how AIO-001 could impact the asthma treatment market in terms of size, growth rate, and competitive landscape.
Sentiment Analysis:

Social Media and News Sentiment: Perform sentiment analysis on social media platforms and online news archives to gauge public and investor sentiment regarding the acquisition.
Expert Opinions: Analyze opinions and forecasts from industry analysts and medical experts to understand the perceived value and potential of AIO-001.
Regulatory Filings Review:

FDA and EMA Filings: Examine regulatory filings and communications from the FDA (U.S.) and EMA (Europe) to understand the regulatory status and any challenges for AIO-001.
Intellectual Property Status: Review patents and IP filings related to AIO-001 to assess legal and competitive positioning.
Data Integration and Predictive Modeling:

Data Integration: Combine data from these diverse sources to create a comprehensive picture.
Predictive Analytics: Use statistical models and machine learning algorithms to predict the potential market impact of AIO-001 and the acquisition's long-term effect on the involved companies and the broader pharmaceutical industry.
Continuous Monitoring and Adjustment:

Regular Updates: Continuously monitor these data sources for new information and adjust analyses accordingly.
Scenario Analysis: Conduct scenario analyses to anticipate potential future developments and their impacts.
This structured approach can provide a comprehensive and quantifiable understanding of how this significant industry event might influence the pharmaceutical landscape.


```



### 1. 我想干什么？
当下大多数所谓的金融模型大多在公开知识上进行训练，但在实际的金融领域，这些公开知识对当前的市场可解释性往往严重不足。如果您感兴趣，可以了解一下凯恩斯、弗里德曼乃至当下行为金融学的各类主张。而据我观察，大多数金融模型无法对投资进行判断。因为它们都是在普通的教材、入门的分析师考试，乃至公司的公开报告上训练。我认为这对于投资的价值非常小。
你可以当我开玩笑，但事实是很多主观分析师的逻辑性可能还不如34b及以上的大模型来的严谨（当然不包括那些优秀的）。而每时每刻，市场都在变化，大量的新闻，海量的数据都是实时的，对于大多数散户们，与其等待蹩脚的分析师写出报告，为什么不用大模型制作一套pipeline呢？
在我的计划中，该模型是这套流程的基座模型，在我的计划中，信息搜集者、标的判断者、定性分析者定性分析者、定量分析者、数据提取者等模型都是该流程的一部分。但模型本身掌握大量的定性和定量方法毫无疑问是重要的。这就是这个模型诞生的理由。

### 2. 关于数据：
正如我刚才所说，很多公开知识的有效性都有些问题——但这并不意味着它们是错误的。在研报中很多研究方法背后的理论支撑也依赖这些知识。所以在我的训练中，我挑选了一些大学教材和一些专业书籍。数量不是很多但质量还不错。另外，我挑选了在2019-2023年12月的大量研究报告数据——这些报告的发布者多种多样，有传统的broke，也有专业研究机构。他们中的大多数是付费的，而且只对机构提供。但无论如何我通过各种各样的手段获取了它们。

如果你看过研报，尤其是高质量的那些，你会发现研报都是主观判断+定量分析，而定量分析中的数据支撑对于整个逻辑链条至关重要。为了提取这些数据（他们中的大多数以图形或者表格的形式出现），我尝试了很多多模态模型，过程非常痛苦，结论是cog-agent和emu2对于这类任务很有效。为了更好的提取信息，我制作了一套从研报上下文总结作为prompt一部分的流程。

最后，我把这些数据做了一个混合。并没有放入通识数据, 因为它就是为了greed而生的。而且行业研报中的知识足够全。

### 3：关于训练：
raw text，全参数训练。基座采用了长上下文的yi-34b-200k。这对于完成理解一篇深度报告是必须的。

当然，我也做了一次sft。这是我的流程中的分析者——目前还没有细分定性和定量分析，但[它的效果](https://huggingface.co/TriadParty/deepmoney-34b-200k-chat-evaluator)已经让我大吃一惊了。

### 4：模型评估：
让我们采样一些最近发生的事件，模拟真实世界中事件驱动的证券分析流程，针对deepmoney和gpt4进行对比测试。因为事件对市场产生的影响比较隐形，如果没有严谨的回测流程，很难评估效果。而且我们的产出需要使用很多定量方法进行分析。所以我把结果贴在这里，大家可以对生产的结果进行感性的评价。
我有一套全球新闻爬取系统，每时每刻都有很多新闻，我的流程中针对这些新闻进行了去重和主客观判断，这些可以通过传统的bert来解决。然后对于deepmoney，这里使用三个步骤来对传入的新闻进行处理：
1， 以上新闻可能对哪些行业板块或者投资标的造成影响？
2， 请你设计一套定量方法研究以上新闻对____业产生的影响。并据此说明具体需要使用哪些数据。
3， 请根据以下数据，_____设计一套具体的定量方法定量分析以上新闻对____业产生的影响。

其中，第一个问题是主观判断，提取新闻影响的标的。这更多的依赖模型的主观分析能力。然后从第一个回答中提取行业名称（这对于熟悉大模型的人来说，设计一套自动化流程易如反掌），填入第二个问题中，目的是获取定量分析的数据。之所以先问定量方法再数据，是COT的魔法。最后一个问题的回答才是我们真正需要的，这个问题上下文中给了足够多的信息，需要它回复一个确切且具体的定量方法。结合代码编写的模型和函数调用的模型，如果你有一个数据字典完善的宏微观数据库的话，这是完全可以实现的。以上是deepmoney和gpt4的这三步回答，该新闻是20240115北京时间早晨9:35刚刚发生的新闻。