package real_time_quotation_indicator

import "github.com/Hongssd/mythsapi/indicatorsdef"

type RealTimeQuotationIndicator indicatorsdef.Indicator

type RealTimeQuotationIndicators []RealTimeQuotationIndicator

func (ris RealTimeQuotationIndicators) ConvertToIndicators() indicatorsdef.Indicators {
	var is indicatorsdef.Indicators
	for _, indicator := range ris {
		is = append(is, indicatorsdef.Indicator(indicator))
	}
	return is
}



const (
	TRADE_DATE                RealTimeQuotationIndicator = "tradeDate"              // 交易日期 通用
	TRADE_TIME                RealTimeQuotationIndicator = "tradeTime"              // 交易时间 通用
	PRE_CLOSE                 RealTimeQuotationIndicator = "preClose"               // 前收盘价 通用
	OPEN                      RealTimeQuotationIndicator = "open"                   // 开盘价 通用
	HIGH                      RealTimeQuotationIndicator = "high"                   // 最高价 通用
	LOW                       RealTimeQuotationIndicator = "low"                    // 最低价 通用
	LATEST                    RealTimeQuotationIndicator = "latest"                 // 最新价 通用
	LATEST_AMOUNT             RealTimeQuotationIndicator = "latestAmount"           // 现额 通用
	LATEST_VOLUME             RealTimeQuotationIndicator = "latestVolume"           // 现量 通用
	AVG_PRICE                 RealTimeQuotationIndicator = "avgPrice"               // 均价 通用
	CHANGE                    RealTimeQuotationIndicator = "change"                 // 涨跌 通用
	CHANGE_RATIO              RealTimeQuotationIndicator = "changeRatio"            // 涨跌幅 通用
	UPPER_LIMIT               RealTimeQuotationIndicator = "upperLimit"             // 涨停价 通用
	DOWN_LIMIT                RealTimeQuotationIndicator = "downLimit"              // 跌停价 通用
	AMOUNT                    RealTimeQuotationIndicator = "amount"                 // 成交额 通用
	VOLUME                    RealTimeQuotationIndicator = "volume"                 // 成交量 通用
	TURNOVER_RATIO            RealTimeQuotationIndicator = "turnoverRatio"          // 换手率 通用
	SELL_VOLUME               RealTimeQuotationIndicator = "sellVolume"             // 内盘 通用
	BUY_VOLUME                RealTimeQuotationIndicator = "buyVolume"              // 外盘 通用
	TOTAL_BID_VOL             RealTimeQuotationIndicator = "totalBidVol"            // 委买十档总量 股票
	TOTAL_ASK_VOL             RealTimeQuotationIndicator = "totalAskVol"            // 委卖十档总量 股票
	TOTAL_SHARES              RealTimeQuotationIndicator = "totalShares"            // 总股本 股票
	TOTAL_CAPITAL             RealTimeQuotationIndicator = "totalCapital"           // 总市值 股票
	PB                        RealTimeQuotationIndicator = "pb"                     // 市净率 股票
	RISE_DAY_COUNT            RealTimeQuotationIndicator = "riseDayCount"           // 连涨天数 股票
	SUSPENSION_FLAG           RealTimeQuotationIndicator = "suspensionFlag"         // 停牌标志 股票
	TRADE_STATUS              RealTimeQuotationIndicator = "tradeStatus"            // 交易状态 股票
	CHG_1MIN                  RealTimeQuotationIndicator = "chg_1min"               // 1分钟涨跌幅 股票
	CHG_3MIN                  RealTimeQuotationIndicator = "chg_3min"               // 3分钟涨跌幅 股票
	CHG_5MIN                  RealTimeQuotationIndicator = "chg_5min"               // 5分钟涨跌幅 股票
	CHG_5D                    RealTimeQuotationIndicator = "chg_5d"                 // 5日涨跌幅 股票
	CHG_10D                   RealTimeQuotationIndicator = "chg_10d"                // 10日涨跌幅 股票
	CHG_20D                   RealTimeQuotationIndicator = "chg_20d"                // 20日涨跌幅 股票
	CHG_60D                   RealTimeQuotationIndicator = "chg_60d"                // 60日涨跌幅 股票
	CHG_120D                  RealTimeQuotationIndicator = "chg_120d"               // 120日涨跌幅 股票
	CHG_250D                  RealTimeQuotationIndicator = "chg_250d"               // 250日涨跌幅 股票
	CHG_YEAR                  RealTimeQuotationIndicator = "chg_year"               // 年初至今涨跌幅 股票
	MV                        RealTimeQuotationIndicator = "mv"                     // 流通市值 股票
	VOL_RATIO                 RealTimeQuotationIndicator = "vol_ratio"              // 量比 股票
	COMMITTEE                 RealTimeQuotationIndicator = "committee"              // 委比 股票
	COMMISSION_DIFF           RealTimeQuotationIndicator = "commission_diff"        // 委差 股票
	PE_TTM                    RealTimeQuotationIndicator = "pe_ttm"                 // 市盈率TTM 股票
	PBR_LF                    RealTimeQuotationIndicator = "pbr_lf"                 // 市净率LF 股票
	SWING                     RealTimeQuotationIndicator = "swing"                  // 振幅 股票
	LASTEST_PRICE             RealTimeQuotationIndicator = "lastest_price"          // 最新成交价 股票
	AF_BACKWARD               RealTimeQuotationIndicator = "af_backward"            // 后复权因子(分红方案计算) 股票
	BID10                     RealTimeQuotationIndicator = "bid10"                  // 买10价 股票
	BID9                      RealTimeQuotationIndicator = "bid9"                   // 买9价 股票
	BID8                      RealTimeQuotationIndicator = "bid8"                   // 买8价 股票
	BID7                      RealTimeQuotationIndicator = "bid7"                   // 买7价 股票
	BID6                      RealTimeQuotationIndicator = "bid6"                   // 买6价 股票
	BID5                      RealTimeQuotationIndicator = "bid5"                   // 买5价 股票
	BID4                      RealTimeQuotationIndicator = "bid4"                   // 买4价 股票
	BID3                      RealTimeQuotationIndicator = "bid3"                   // 买3价 股票
	BID2                      RealTimeQuotationIndicator = "bid2"                   // 买2价 股票
	BID1                      RealTimeQuotationIndicator = "bid1"                   // 买1价 股票
	ASK1                      RealTimeQuotationIndicator = "ask1"                   // 卖1价 股票
	ASK2                      RealTimeQuotationIndicator = "ask2"                   // 卖2价 股票
	ASK3                      RealTimeQuotationIndicator = "ask3"                   // 卖3价 股票
	ASK4                      RealTimeQuotationIndicator = "ask4"                   // 卖4价 股票
	ASK5                      RealTimeQuotationIndicator = "ask5"                   // 卖5价 股票
	ASK6                      RealTimeQuotationIndicator = "ask6"                   // 卖6价 股票
	ASK7                      RealTimeQuotationIndicator = "ask7"                   // 卖7价 股票
	ASK8                      RealTimeQuotationIndicator = "ask8"                   // 卖8价 股票
	ASK9                      RealTimeQuotationIndicator = "ask9"                   // 卖9价 股票
	ASK10                     RealTimeQuotationIndicator = "ask10"                  // 卖10价 股票
	BID_SIZE10                RealTimeQuotationIndicator = "bidSize10"              // 买10量 股票
	BID_SIZE9                 RealTimeQuotationIndicator = "bidSize9"               // 买9量 股票
	BID_SIZE8                 RealTimeQuotationIndicator = "bidSize8"               // 买8量 股票
	BID_SIZE7                 RealTimeQuotationIndicator = "bidSize7"               // 买7量 股票
	BID_SIZE6                 RealTimeQuotationIndicator = "bidSize6"               // 买6量 股票
	BID_SIZE5                 RealTimeQuotationIndicator = "bidSize5"               // 买5量 股票
	BID_SIZE4                 RealTimeQuotationIndicator = "bidSize4"               // 买4量 股票
	BID_SIZE3                 RealTimeQuotationIndicator = "bidSize3"               // 买3量 股票
	BID_SIZE2                 RealTimeQuotationIndicator = "bidSize2"               // 买2量 股票
	BID_SIZE1                 RealTimeQuotationIndicator = "bidSize1"               // 买1量 股票
	ASK_SIZE1                 RealTimeQuotationIndicator = "askSize1"               // 卖1量 股票
	ASK_SIZE2                 RealTimeQuotationIndicator = "askSize2"               // 卖2量 股票
	ASK_SIZE3                 RealTimeQuotationIndicator = "askSize3"               // 卖3量 股票
	ASK_SIZE4                 RealTimeQuotationIndicator = "askSize4"               // 卖4量 股票
	ASK_SIZE5                 RealTimeQuotationIndicator = "askSize5"               // 卖5量 股票
	ASK_SIZE6                 RealTimeQuotationIndicator = "askSize6"               // 卖6量 股票
	ASK_SIZE7                 RealTimeQuotationIndicator = "askSize7"               // 卖7量 股票
	ASK_SIZE8                 RealTimeQuotationIndicator = "askSize8"               // 卖8量 股票
	ASK_SIZE9                 RealTimeQuotationIndicator = "askSize9"               // 卖9量 股票
	ASK_SIZE10                RealTimeQuotationIndicator = "askSize10"              // 卖10量 股票
	AVG_BUY_PRICE             RealTimeQuotationIndicator = "avgBuyPrice"            // 均买价 股票
	AVG_SELL_PRICE            RealTimeQuotationIndicator = "avgSellPrice"           // 均卖价 股票
	TOTAL_BUY_VOLUME          RealTimeQuotationIndicator = "totalBuyVolume"         // 总买量 股票
	TOTAL_SELL_VOLUME         RealTimeQuotationIndicator = "totalSellVolume"        // 总卖量 股票
	TRANS_CLASSIFICATION      RealTimeQuotationIndicator = "transClassification"    // 成交分类 股票
	TRANS_TIMES               RealTimeQuotationIndicator = "transTimes"             // 成交次数 股票
	MAIN_INFLOW               RealTimeQuotationIndicator = "mainInflow"             // 主力流入金额 股票
	MAIN_OUTFLOW              RealTimeQuotationIndicator = "mainOutflow"            // 主力流出金额 股票
	MAIN_NET_INFLOW           RealTimeQuotationIndicator = "mainNetInflow"          // 主力净流入金额 股票
	RETAIL_INFLOW             RealTimeQuotationIndicator = "retailInflow"           // 散户流入金额 股票
	RETAIL_OUTFLOW            RealTimeQuotationIndicator = "retailOutflow"          // 散户流出金额 股票
	RETAIL_NET_INFLOW         RealTimeQuotationIndicator = "retailNetInflow"        // 散户净流入金额 股票
	LARGE_INFLOW              RealTimeQuotationIndicator = "largeInflow"            // 超大单流入金额 股票
	LARGE_OUTFLOW             RealTimeQuotationIndicator = "largeOutflow"           // 超大单流出金额 股票
	LARGE_NET_INFLOW          RealTimeQuotationIndicator = "largeNetInflow"         // 超大单净流入金额 股票
	BIG_INFLOW                RealTimeQuotationIndicator = "bigInflow"              // 大单流入金额 股票
	BIG_OUTFLOW               RealTimeQuotationIndicator = "bigOutflow"             // 大单流出金额 股票
	BIG_NET_INFLOW            RealTimeQuotationIndicator = "bigNetInflow"           // 大单净流入金额 股票
	MIDDLE_INFLOW             RealTimeQuotationIndicator = "middleInflow"           // 中单流入金额 股票
	MIDDLE_OUTFLOW            RealTimeQuotationIndicator = "middleOutflow"          // 中单流出金额 股票
	MIDDLE_NET_INFLOW         RealTimeQuotationIndicator = "middleNetInflow"        // 中单净流入金额 股票
	SMALL_INFLOW              RealTimeQuotationIndicator = "smallInflow"            // 小单流入金额 股票
	SMALL_OUTFLOW             RealTimeQuotationIndicator = "smallOutflow"           // 小单流出金额 股票
	SMALL_NET_INFLOW          RealTimeQuotationIndicator = "smallNetInflow"         // 小单净流入金额 股票
	ACTIVE_BUY_LARGE_AMT      RealTimeQuotationIndicator = "activeBuyLargeAmt"      // 主动买入特大单金额 股票
	ACTIVE_SELL_LARGE_AMT     RealTimeQuotationIndicator = "activeSellLargeAmt"     // 主动卖出特大单金额 股票
	ACTIVE_BUY_MAIN_AMT       RealTimeQuotationIndicator = "activeBuyMainAmt"       // 主动买入大单金额 股票
	ACTIVE_SELL_MAIN_AMT      RealTimeQuotationIndicator = "activeSellMainAmt"      // 主动卖出大单金额 股票
	ACTIVE_BUY_MIDDLE_AMT     RealTimeQuotationIndicator = "activeBuyMiddleAmt"     // 主动买入中单金额 股票
	ACTIVE_SELL_MIDDLE_AMT    RealTimeQuotationIndicator = "activeSellMiddleAmt"    // 主动卖出中单金额 股票
	ACTIVE_BUY_SMALL_AMT      RealTimeQuotationIndicator = "activeBuySmallAmt"      // 主动买入小单金额 股票
	ACTIVE_SELL_SMALL_AMT     RealTimeQuotationIndicator = "activeSellSmallAmt"     // 主动卖出小单金额 股票
	POSSITIVE_BUY_LARGE_AMT   RealTimeQuotationIndicator = "possitiveBuyLargeAmt"   // 被动买入特大单金额 股票
	POSSITIVE_SELL_LARGE_AMT  RealTimeQuotationIndicator = "possitiveSellLargeAmt"  // 被动卖出特大单金额 股票
	POSSITIVE_BUY_MAIN_AMT    RealTimeQuotationIndicator = "possitiveBuyMainAmt"    // 被动买入大单金额 股票
	POSSITIVE_SELL_MAIN_AMT   RealTimeQuotationIndicator = "possitiveSellMainAmt"   // 被动卖出大单金额 股票
	POSSITIVE_BUY_MIDDLE_AMT  RealTimeQuotationIndicator = "possitiveBuyMiddleAmt"  // 被动买入中单金额 股票
	POSSITIVE_SELL_MIDDLE_AMT RealTimeQuotationIndicator = "possitiveSellMiddleAmt" // 被动卖出中单金额 股票
	POSSITIVE_BUY_SMALL_AMT   RealTimeQuotationIndicator = "possitiveBuySmallAmt"   // 被动买入小单金额 股票
	POSSITIVE_SELL_SMALL_AMT  RealTimeQuotationIndicator = "possitiveSellSmallAmt"  // 被动卖出小单金额 股票
	ACTIVE_BUY_LARGE_VOL      RealTimeQuotationIndicator = "activeBuyLargeVol"      // 主动买入特大单量 股票
	ACTIVE_SELL_LARGE_VOL     RealTimeQuotationIndicator = "activeSellLargeVol"     // 主动卖出特大单量 股票
	ACTIVE_BUY_MAIN_VOL       RealTimeQuotationIndicator = "activeBuyMainVol"       // 主动买入大单量 股票
	ACTIVE_SELL_MAIN_VOL      RealTimeQuotationIndicator = "activeSellMainVol"      // 主动卖出大单量 股票
	ACTIVE_BUY_MIDDLE_VOL     RealTimeQuotationIndicator = "activeBuyMiddleVol"     // 主动买入中单量 股票
	ACTIVE_SELL_MIDDLE_VOL    RealTimeQuotationIndicator = "activeSellMiddleVol"    // 主动卖出中单量 股票
	ACTIVE_BUY_SMALL_VOL      RealTimeQuotationIndicator = "activeBuySmallVol"      // 主动买入小单量 股票
	ACTIVE_SELL_SMALL_VOL     RealTimeQuotationIndicator = "activeSellSmallVol"     // 主动卖出小单量 股票
	POSSITIVE_BUY_LARGE_VOL   RealTimeQuotationIndicator = "possitiveBuyLargeVol"   // 被动买入特大单量 股票
	POSSITIVE_SELL_LARGE_VOL  RealTimeQuotationIndicator = "possitiveSellLargeVol"  // 被动卖出特大单量 股票
	POSSITIVE_BUY_MAIN_VOL    RealTimeQuotationIndicator = "possitiveBuyMainVol"    // 被动买入大单量 股票
	POSSITIVE_SELL_MAIN_VOL   RealTimeQuotationIndicator = "possitiveSellMainVol"   // 被动卖出大单量 股票
	POSSITIVE_BUY_MIDDLE_VOL  RealTimeQuotationIndicator = "possitiveBuyMiddleVol"  // 被动买入中单量 股票
	POSSITIVE_SELL_MIDDLE_VOL RealTimeQuotationIndicator = "possitiveSellMiddleVol" // 被动卖出中单量 股票
	POSSITIVE_BUY_SMALL_VOL   RealTimeQuotationIndicator = "possitiveBuySmallVol"   // 被动买入小单量 股票
	POSSITIVE_SELL_SMALL_VOL  RealTimeQuotationIndicator = "possitiveSellSmallVol"  // 被动卖出小单量 股票
	ACTIVEBUY_VOLUME          RealTimeQuotationIndicator = "activebuy_volume"       // 主买总量 股票
	ACTIVESELL_VOLUME         RealTimeQuotationIndicator = "activesell_volume"      // 主卖总量 股票
	ACTIVEBUY_AMT             RealTimeQuotationIndicator = "activebuy_amt"          // 主买总额 股票
	ACTIVESELL_AMT            RealTimeQuotationIndicator = "activesell_amt"         // 主卖总额 股票
	POST_LASTEST              RealTimeQuotationIndicator = "post_lastest"           // 盘后最新成交价 股票
	POST_LATEST_VOLUME        RealTimeQuotationIndicator = "post_latestVolume"      // 盘后现量 股票
	POST_VOLUME               RealTimeQuotationIndicator = "post_volume"            // 盘后成交量 股票
	POST_AMT                  RealTimeQuotationIndicator = "post_amt"               // 盘后成交额 股票
	POST_DEALNUM              RealTimeQuotationIndicator = "post_dealnum"           // 盘后成交笔数 股票
	PRICE_DIFF                RealTimeQuotationIndicator = "priceDiff"              // 买卖价差 港股专用
	SHARES_PER_HAND           RealTimeQuotationIndicator = "sharesPerHand"          // 每手股数 港股专用
	EXPIRY_DATE               RealTimeQuotationIndicator = "expiryDate"             // 到期日 港股专用
	IOPV                      RealTimeQuotationIndicator = "iopv"                   // IOPV（净值估值） 基金专用
	PREMIUM                   RealTimeQuotationIndicator = "premium"                // 折价 基金专用
	RISE_COUNT                RealTimeQuotationIndicator = "riseCount"              // 上涨家数 指数专用
	FALL_COUNT                RealTimeQuotationIndicator = "fallCount"              // 下跌家数 指数专用
	UP_LIMIT_COUNT            RealTimeQuotationIndicator = "upLimitCount"           // 涨停家数 指数专用
	DOWN_LIMIT_COUNT          RealTimeQuotationIndicator = "downLimitCount"         // 跌停家数 指数专用
	SUSPENSION_COUNT          RealTimeQuotationIndicator = "suspensionCount"        // 停牌家数 指数专用
	PURE_BOND_VALUE_CB        RealTimeQuotationIndicator = "pure_bond_value_cb"     // 纯债价值 指数专用
	SURPLUS_TERM              RealTimeQuotationIndicator = "surplus_term"           // 剩余期限(天) 指数专用
	DEAL_DIRECTION            RealTimeQuotationIndicator = "dealDirection"          // 成交方向 期货期权专用
	DEALTYPE                  RealTimeQuotationIndicator = "dealtype"               // 成交性质 期货期权专用
	IMPLIED_VOLATILITY        RealTimeQuotationIndicator = "impliedVolatility"      // 隐含波动率 期权专用
	HISTORY_VOLATILITY        RealTimeQuotationIndicator = "historyVolatility"      // 历史波动率 期权专用
	DELTA                     RealTimeQuotationIndicator = "delta"                  // Delta 期权专用
	GAMMA                     RealTimeQuotationIndicator = "gamma"                  // Gamma 期权专用
	VEGA                      RealTimeQuotationIndicator = "vega"                   // Vega 期权专用
	THETA                     RealTimeQuotationIndicator = "theta"                  // Theta 期权专用
	RHO                       RealTimeQuotationIndicator = "rho"                    // Rho 期权专用
	PRE_OPEN_INTEREST         RealTimeQuotationIndicator = "pre_open_interest"      // 前持仓量 期权专用
	PRE_IMPLIED_VOLATILITY    RealTimeQuotationIndicator = "pre_implied_volatility" // 前隐含波动率 期权专用
	VOLUME_PCR_TOTAL          RealTimeQuotationIndicator = "volume_pcr_total"       // 成交量pcr（品种） 期权专用
	VOLUME_PCR_MONTH          RealTimeQuotationIndicator = "volume_pcr_month"       // 成交量pcr（同月） 期权专用
)