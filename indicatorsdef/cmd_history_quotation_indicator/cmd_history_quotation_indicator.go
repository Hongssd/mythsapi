package cmd_history_quotation_indicator

import "github.com/Hongssd/mythsapi/indicatorsdef"

type CmdHistoryQuotationIndicator indicatorsdef.Indicator

type CmdHistoryQuotationIndicators []CmdHistoryQuotationIndicator

func (chqis CmdHistoryQuotationIndicators) ConvertToIndicators() indicatorsdef.Indicators {
	var indicators indicatorsdef.Indicators
	for _, chqi := range chqis {
		indicators = append(indicators, indicatorsdef.Indicator(chqi))
	}
	return indicators
}

const (
	PRE_CLOSE                         CmdHistoryQuotationIndicator = "preClose"                          // 前收盘价
	OPEN                              CmdHistoryQuotationIndicator = "open"                              // 开盘价
	HIGH                              CmdHistoryQuotationIndicator = "high"                              // 最高价
	LOW                               CmdHistoryQuotationIndicator = "low"                               // 最低价
	CLOSE                             CmdHistoryQuotationIndicator = "close"                             // 收盘价
	AVG_PRICE                         CmdHistoryQuotationIndicator = "avgPrice"                          // 均价
	CHANGE                            CmdHistoryQuotationIndicator = "change"                            // 涨跌
	CHANGE_RATIO                      CmdHistoryQuotationIndicator = "changeRatio"                       // 涨跌幅
	VOLUME                            CmdHistoryQuotationIndicator = "volume"                            // 成交量
	AMOUNT                            CmdHistoryQuotationIndicator = "amount"                            // 成交额
	TURNOVER_RATIO                    CmdHistoryQuotationIndicator = "turnoverRatio"                     // 换手率
	TRANSACTION_AMOUNT                CmdHistoryQuotationIndicator = "transactionAmount"                 // 成交笔数
	TOTAL_SHARES                      CmdHistoryQuotationIndicator = "totalShares"                       // 总股本
	TOTAL_CAPITAL                     CmdHistoryQuotationIndicator = "totalCapital"                      // 总市值
	FLOAT_SHARES_OF_ASHARES           CmdHistoryQuotationIndicator = "floatSharesOfAShares"              // A股流通股本
	FLOAT_SHARES_OF_BSHARES           CmdHistoryQuotationIndicator = "floatSharesOfBShares"              // B股流通股本
	FLOAT_CAPITAL_OF_ASHARES          CmdHistoryQuotationIndicator = "floatCapitalOfAShares"             // A股流通市值
	FLOAT_CAPITAL_OF_BSHARES          CmdHistoryQuotationIndicator = "floatCapitalOfBShares"             // B股流通市值
	PE_TTM                            CmdHistoryQuotationIndicator = "pe_ttm"                            // 市盈率（TTM）
	PE                                CmdHistoryQuotationIndicator = "pe"                                // PE市盈率
	PB                                CmdHistoryQuotationIndicator = "pb"                                // PB市净率
	PS                                CmdHistoryQuotationIndicator = "ps"                                // PS市销率
	PCF                               CmdHistoryQuotationIndicator = "pcf"                               // PCF市现率
	THS_TRADING_STATUS_STOCK          CmdHistoryQuotationIndicator = "ths_trading_status_stock"          // 交易状态
	THS_UP_AND_DOWN_STATUS_STOCK      CmdHistoryQuotationIndicator = "ths_up_and_down_status_stock"      // 涨跌停状态
	THS_AF_STOCK                      CmdHistoryQuotationIndicator = "ths_af_stock"                      // 复权因子
	THS_VOL_AFTER_TRADING_STOCK       CmdHistoryQuotationIndicator = "ths_vol_after_trading_stock"       // 盘后成交量
	THS_TRANS_NUM_AFTER_TRADING_STOCK CmdHistoryQuotationIndicator = "ths_trans_num_after_trading_stock" // 盘后成交笔数
	THS_AMT_AFTER_TRADING_STOCK       CmdHistoryQuotationIndicator = "ths_amt_after_trading_stock"       // 盘后成交额
	THS_VAILD_TURNOVER_STOCK          CmdHistoryQuotationIndicator = "ths_vaild_turnover_stock"          // 有效换手率
	NET_ASSET_VALUE                   CmdHistoryQuotationIndicator = "netAssetValue"                     // 单位净值 基金专用
	ADJUSTED_NAV                      CmdHistoryQuotationIndicator = "adjustedNAV"                       // 复权单位净值 基金专用
	ACCUMULATED_NAV                   CmdHistoryQuotationIndicator = "accumulatedNAV"                    // 累计单位净值 基金专用
	PREMIUM                           CmdHistoryQuotationIndicator = "premium"                           // 贴水 基金专用
	PREMIUM_RATIO                     CmdHistoryQuotationIndicator = "premiumRatio"                      // 贴水率 基金专用
	ESTIMATED_POSITION                CmdHistoryQuotationIndicator = "estimatedPosition"                 // 估算仓位 基金专用
	FLOAT_CAPITAL                     CmdHistoryQuotationIndicator = "floatCapital"                      // 流通市值 指数专用
	PE_TTM_INDEX                      CmdHistoryQuotationIndicator = "pe_ttm_index"                      // PE(TTM) 指数专用
	PB_MRQ                            CmdHistoryQuotationIndicator = "pb_mrq"                            // PB(MRQ) 指数专用
	PE_INDEX_PUBLISHER                CmdHistoryQuotationIndicator = "pe_indexPublisher"                 // PE(指数发布方） 指数专用
	YIELD_MATURITY                    CmdHistoryQuotationIndicator = "yieldMaturity"                     // 到期收益率 债券专用
	REMAINING_TERM                    CmdHistoryQuotationIndicator = "remainingTerm"                     // 剩余期限 债券专用
	MAXWELL_DURATION                  CmdHistoryQuotationIndicator = "maxwellDuration"                   // 麦氏久期 债券专用
	MODIFIED_DURATION                 CmdHistoryQuotationIndicator = "modifiedDuration"                  // 修正久期 债券专用
	CONVEXITY                         CmdHistoryQuotationIndicator = "convexity"                         // 凸性 债券专用
	CLOSE_2330                        CmdHistoryQuotationIndicator = "close_2330"                        // 收盘价（23：30） 外汇交易中心专用
	OPEN_INTEREST                     CmdHistoryQuotationIndicator = "openInterest"                      // 持仓量 期权专用
	POSITION_CHANGE                   CmdHistoryQuotationIndicator = "positionChange"                    // 持仓变动 期权专用
	PRE_SETTLEMENT                    CmdHistoryQuotationIndicator = "preSettlement"                     // 前结算价 期货专用
	SETTLEMENT                        CmdHistoryQuotationIndicator = "settlement"                        // 结算价 期货专用
	CHANGE_SETTLEMENT                 CmdHistoryQuotationIndicator = "change_settlement"                 // 涨跌（结算价） 期货专用
	CHG_SETTLEMENT                    CmdHistoryQuotationIndicator = "chg_settlement"                    // 涨跌幅（结算价） 期货专用
	AMPLITUDE                         CmdHistoryQuotationIndicator = "amplitude"                         // 振幅 期货专用
)
