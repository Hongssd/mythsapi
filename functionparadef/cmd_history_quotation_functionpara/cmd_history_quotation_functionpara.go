package cmd_history_quotation_functionpara

import (
	"encoding/json"

	"github.com/Hongssd/mythsapi/functionparadef"
)

type CmdHistoryQuotationFunctionPara struct {
	Interval       *Interval  `json:"Interval,omitempty"`       //时间周期  D-日 W-周 M-月 Q-季 S-半年 Y-年 同抽样周期二选一，返回周期汇总统计值 默认值：D
	SampleInterval *Interval  `json:"SampleInterval,omitempty"` //抽样周期  D-日 W-周 M-月 Q-季 S-半年 Y-年 同时间周期二选一，返回周期最后一个交易日日频数据 默认值：D
	CPS            *CPS       `json:"CPS,omitempty"`            //复权方式  1-不复权 2-前复权（分红再投） 3-后复权（分红再投） 4-全流通前复权（分红再投） 5-全流通后复权（分红再投） 6-前复权（现金分红） 7-后复权（现金分红）默认值： 1
	PriceType      *PriceType `json:"PriceType,omitempty"`      //报价类型  1-全价 2-净价 仅债券生效 默认值：1
	Fill           *Fill      `json:"Fill,omitempty"`           //非交易间隔处理  Previous-沿用之前数据 Blank-空值 具体数值-自定义数值 Omit-缺省值 默认值：Previous
	BaseDate       *string    `json:"BaseDate,omitempty"`       //设定复权基点  BaseDate 复权基点日期，"YYYY-MM-DD" 默认值：后复权按上市日，前复权按最新日
	Currency       *Currency  `json:"Currency,omitempty"`       //货币  MHB-美元 GHB-港元 RMB-人民币 YSHB-原始货币 默认值：YSHB
}

func NewFunctionPara() *CmdHistoryQuotationFunctionPara {
	return &CmdHistoryQuotationFunctionPara{
		Interval:       nil,
		SampleInterval: nil,
		CPS:            nil,
		PriceType:      nil,
		Fill:           nil,
		BaseDate:       nil,
		Currency:       nil,
	}
}

func (para *CmdHistoryQuotationFunctionPara) SetInterval(interval Interval) *CmdHistoryQuotationFunctionPara {
	para.Interval = &interval
	return para
}

func (para *CmdHistoryQuotationFunctionPara) SetSampleInterval(sampleInterval Interval) *CmdHistoryQuotationFunctionPara {
	para.SampleInterval = &sampleInterval
	return para
}

func (para *CmdHistoryQuotationFunctionPara) SetCPS(cps CPS) *CmdHistoryQuotationFunctionPara {
	para.CPS = &cps
	return para
}

func (para *CmdHistoryQuotationFunctionPara) SetPriceType(priceType PriceType) *CmdHistoryQuotationFunctionPara {
	para.PriceType = &priceType
	return para
}

func (para *CmdHistoryQuotationFunctionPara) SetFill(fill Fill) *CmdHistoryQuotationFunctionPara {
	para.Fill = &fill
	return para
}

func (para *CmdHistoryQuotationFunctionPara) SetBaseDate(baseDate string) *CmdHistoryQuotationFunctionPara {
	para.BaseDate = &baseDate
	return para
}

func (para *CmdHistoryQuotationFunctionPara) SetCurrency(currency Currency) *CmdHistoryQuotationFunctionPara {
	para.Currency = &currency
	return para
}

func (c *CmdHistoryQuotationFunctionPara) ConvertToFunctionPara() *functionparadef.FunctionPara {
	data, _ := json.Marshal(c)
	var result functionparadef.FunctionPara
	json.Unmarshal(data, &result)
	return &result
}

type Interval string

const (
	Interval_Day        Interval = "D" //日
	Interval_Week       Interval = "W" //周
	Interval_Month      Interval = "M" //月
	Interval_Quarter    Interval = "Q" //季
	Interval_Semiannual Interval = "S" //半年
	Interval_Year       Interval = "Y" //年
)

func (i Interval) IntervalForTimeMillisecond() int64 {
	switch i {
	case Interval_Day:
		return 24 * 60 * 60 * 1000
	case Interval_Week:
		return 7 * 24 * 60 * 60 * 1000
	case Interval_Month:
		return 30 * 24 * 60 * 60 * 1000
	case Interval_Quarter:
		return 3 * 30 * 24 * 60 * 60 * 1000
	case Interval_Semiannual:
		return 6 * 30 * 24 * 60 * 60 * 1000
	case Interval_Year:
		return 12 * 30 * 24 * 60 * 60 * 1000
	default:
		return 24 * 60 * 60 * 1000
	}
}

type CPS string

const (
	CPS1 CPS = "1" //不复权
	CPS2 CPS = "2" //前复权（分红再投）
	CPS3 CPS = "3" //后复权（分红再投）
	CPS4 CPS = "4" //全流通前复权（分红再投）
	CPS5 CPS = "5" //全流通后复权（分红再投）
	CPS6 CPS = "6" //前复权（现金分红）
	CPS7 CPS = "7" //后复权（现金分红）
)

type PriceType string

const (
	PriceType1 PriceType = "1" //全价
	PriceType2 PriceType = "2" //净价

)

type Fill string

const (
	FillPrevious Fill = "Previous" //沿用之前数据
	FillBlank    Fill = "Blank"    //空值
	FillOmit     Fill = "Omit"     //缺省值
)

type Currency string

const (
	CurrencyMHB  Currency = "MHB"  //美元
	CurrencyGHB  Currency = "GHB"  //港元
	CurrencyRMB  Currency = "RMB"  //人民币
	CurrencyYSHB Currency = "YSHB" //原始货币
)
