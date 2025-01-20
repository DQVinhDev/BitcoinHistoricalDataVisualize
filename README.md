# BITCOIN HISTORICAL DATA ANALYSIS

This document summarizes the historical analysis of Bitcoin price and volume data from September 2014 to February 2022. The analysis utilizes data obtained from Yahoo Finance and employs data visualization techniques using JavaScript libraries like Plotly and D3, along with HTML and CSS for presentation.

## Contents

- [Chapter 1: Introduction](#chapter-1-introduction)
- [Chapter 2: Methodology](#chapter-2-methodology)
- [Chapter 3: Analysis of Bitcoin Charts](#chapter-3-analysis-of-bitcoin-charts)
- [Chapter 4: Conclusion and Future Work](#chapter-4-conclusion-and-future-work)
- [References](#references)

## Chapter 1: Introduction

This chapter provides background information and motivations for analyzing Bitcoin's historical data.

**Key Points:**

- Bitcoin, created in 2009, has grown into a significant asset class.
- Its price history is marked by volatility with periods of bubbles and crashes.
- Minute-to-minute OHLC (Open, High, Low, Close) price and volume data from exchanges (September 2014 to October 2020) has been crucial for analysis.
- Various data analysis and visualization tools have been developed to understand Bitcoin price trends.
- The motivations for analysis are diverse, including:
    - Investor decision-making.
    - Trader strategy development.
    - Academic research on economic and behavioral aspects.
    - Policymaker understanding for regulation.
- Web-based tools facilitate the manipulation and visualization of this data.

**Objectives of Collecting Bitcoin Historical Data:**

- **Historical Analysis:**  Study Bitcoin price patterns and trends to understand factors influencing price changes.
- **Trading Strategies:** Develop data-driven trading strategies using technical analysis.
- **Data Visualization:** Create web-based tools for interactive data manipulation and visualization.
- **Algorithmic Trading:** Enable the development of automated trading algorithms.

## Chapter 2: Methodology

This chapter outlines the methods used for analyzing and visualizing the Bitcoin data.

**Key Points:**

- **Visual Representation:** Graphs and charts are essential for extracting insights from Bitcoin data.
- **Technology Stack:** The analysis leverages:
    - **JavaScript and Plotly:** A powerful open-source library for creating interactive charts and graphs.
    - **HTML and CSS:** Used for structuring and styling the visualization on websites.

**About the Dataset:**

- **Source:** Finance Yahoo.
- **Timeframe:** September 2014 to February 2022.
- **Granularity:** Minute-to-minute updates of OHLC, Volume in BTC and indicated currency.
- **Data Integrity:** Preliminary evaluation indicates good data point distribution with few outliers and high reliability.
- **Dataset Size:** Approximately 2713 rows and 7 columns.

**Column Descriptions:**

- **Date:** Date of the Bitcoin trading record (date type).
- **Open:** Opening price of Bitcoin (number type).
- **High:** Highest price of Bitcoin reached during the day (number type).
- **Low:** Lowest price of Bitcoin reached during the day (number type).
- **Volume_(BTC):** Total amount of Bitcoin traded (number type).
- **Close:** Closing price of Bitcoin (number type).
- **Adj_Close:** Adjusted closing price (number type).

**Prototype Design:**

- **Graph:** A line chart to visualize the historical price of Bitcoin (Date on x-axis, Price on y-axis).
- **Chart:** Visualizations to show transaction activity, including:
    - Total number and volume of transactions.
    - Total dollars traded.
    - Top countries involved in trading.
    - Price of Bitcoin over time.

## Chapter 3: Analysis of Bitcoin Charts

This chapter presents the analysis of various Bitcoin charts generated from the historical data.

**3.1. Candlestick Chart of Bitcoin**

- Candlestick charts are used to visualize the open, high, low, and close prices for each day.
- Candle color indicates price movement (black for increase, red for decrease).

**Data Analysis:**

- **Minimum Closing Price:** $178.102
- **Maximum Closing Price:** $67566.82
- **Minimum Opening Price:** $176.89
- **Maximum Opening Price:** $67549.73
- **Minimum High Price:** $211.73
- **Maximum High Price:** $68789.62
- **Minimum Low Price:** $171.50
- **Maximum Low Price:** $66382.06
- **Minimum Trading Volume:** 5,914,570
- **Maximum Trading Volume:** 350,967,900,000

**Data Insights:**

- **2015-2017 (Price Increase):** Exponential growth from around $312 to approximately $19,783. Driven by investor confidence, media attention, and ICOs.
- **2020-2021 (Price Increase):** Significant rise from around $7,195 to nearly $64,863. Attributed to institutional involvement, acceptance as a store of value, and pandemic-related economic factors.
- **2014-2015 (Price Decrease):** Correction phase following the 2013 surge, dropping from $770 to $313. Influenced by regulatory concerns and exchange breaches (e.g., Mt. Gox).
- **2018-2019 (Price Decrease):** Bear market after the 2017 peak, falling from $13,880 to $3,193. Driven by market correction, regulatory uncertainties, and negative sentiment.
- **2021-2022 (Price Increase):** Reaching new record levels above $65,000. Impacted by the introduction of Bitcoin ETFs and significant investments from companies like Tesla.

**Summary:**

- Bitcoin prices generally show an upward trend.
- Significant price fluctuations present opportunities and risks.
- Trading volume varies, reflecting market interest.

**Conclusion:**

The analysis reveals an upward trend in Bitcoin's closing prices, but with periodic volatility. Trading volume reflects changes in market participation.

**3.2. Line Chart of Volume of Bitcoin**

- Visualizes the historical trading volume of Bitcoin.

**Data Insights:**

- **2014-2015:** Early adoption with significant volume fluctuations due to regulatory uncertainties.
- **2016-2017:** Steady growth followed by a volume surge during Bitcoin's price highs.
- **2018-2019:** Market correction leading to a decline in trading volume.
- **2020-2022:** Resurgence in volume driven by institutional adoption and the pandemic.

**Data Insight Summary:**

- **Market Maturity:** Increasing volume suggests growing interest and participation.
- **Bull Market and Speculative Activity:** Higher volume during price surges indicates speculation.
- **Market Correction:** Decreased volume reflects a shift in market dynamics.
- **Seasonal or Cyclical Patterns:** Potential for recurring patterns in trading volume.

**Conclusion:**

Bitcoin's trading volume has fluctuated significantly, reflecting its evolution and market sentiment.

**3.3. Line Chart Close Price of Bitcoin**

- Visualizes the historical closing price of Bitcoin.

**3.3.1. Line Chart of Opening Price:**

- Visualizes the historical opening price of Bitcoin.

**3.3.2. Summarize Open Price and Close Price Bitcoin Data:**

- **Lowest Closing Price:** $178.103 (January 14, 2015)
- **Highest Closing Price:** $67566.828 (November 7, 2021)
- **Lowest Opening Price:** $176.897 (January 15, 2015)
- **Highest Opening Price:** $67549.734 (November 9, 2021)

**3.3.3. Analyst Price Trend:**

- Overall upward trend with significant growth.
- Notable price appreciation during the 2017 and 2020-2021 bull runs.
- Periods of correction and consolidation.

**3.3.4. Data Insight:**

- **Seasonality:** Potentially higher prices in later months of the year.
- **Impact of Events:**
    - Mt. Gox collapse led to price decline.
    - Regulatory announcements caused fluctuations.
    - Adoption by major companies (e.g., PayPal, Tesla) had positive impact.
    - Bitcoin halving events influenced prices.

**3.3.5. Conclusion:**

Bitcoin's open and close prices show significant growth and volatility, influenced by various market events and seasonality.

**3.4. Line Chart Low Price and High Price of Bitcoin**

- Visualizes the daily low and high prices, illustrating price range and volatility.

**3.4.1. Data Analyst:**

- General upward trend for both low and high prices.
- Strong volatility periods are marked by a greater distance between low and high prices on a given day (e.g., May 19, 2021).
- Periods of consolidation show low and high prices moving closely together.

**3.4.2. Data Insight:**

- General upward price trend since 2014.
- Periods of significant price volatility, including the 2018 crash.
- Future price remains uncertain but with potential for growth.

**3.4.3. Conclusion:**

Historical trends show upward movement and volatility. Caution is advised when predicting future prices based solely on historical data.

## Chapter 4: Conclusion and Future Work

**4.1. Conclusion**

The historical analysis of Bitcoin's price from 2014 to 2022 reveals a volatile but ultimately upward trajectory. Key events like the Mt. Gox hack, bull runs, and regulatory changes have significantly impacted its price. Bitcoin has shown resilience and increasing mainstream acceptance over time.

**4.2. Future Work**

Several areas warrant further investigation:

- **Long-Term Trend Analysis:**  In-depth analysis of long-term price patterns and support/resistance levels.
- **Sentiment Analysis:**  Gauging market sentiment using social media and news data to correlate with price movements.
- **Correlation Analysis:**  Exploring correlations between Bitcoin's price and macroeconomic factors, regulations, and geopolitical events.
- **Market Microstructure Analysis:**  Investigating trading volumes, liquidity, and potential market manipulation.
- **Integration of Fundamental Analysis:**  Analyzing network adoption, transaction volumes, and technological developments.
- **Machine Learning and Predictive Modeling:**  Developing forecasting models for price prediction.
- **Comparative Analysis with Other Cryptocurrencies:**  Comparing Bitcoin's performance with other digital assets.
- **Risk Analysis and Portfolio Optimization:**  Evaluating Bitcoin's risk characteristics and its role in investment portfolios.

## References

1. Bitcoin USD (BTC-USD) Cryptocurrency Profile & Facts - Yahoo Finance
