package get_trade_dates_functionpara

import (
	"encoding/json"

	"github.com/Hongssd/mythsapi/functionparadef"
)

type GetTradeDatesFunctionPara struct {
	Mode       *Mode       `json:"mode,omitempty"`       //函数模式
	DateType   *DateType   `json:"dateType,omitempty"`   //日期类型
	DateFormat *DateFormat `json:"dateFormat,omitempty"` //日期格式
	Period     *Period     `json:"period,omitempty"`     //时间周期
	Periodnum  *string     `json:"periodnum,omitempty"`  //时间周期偏移量
}

func NewFunctionPara() *GetTradeDatesFunctionPara {
	return &GetTradeDatesFunctionPara{
		Mode:       nil,
		DateType:   nil,
		DateFormat: nil,
		Period:     nil,
		Periodnum:  nil,
	}
}

func (para *GetTradeDatesFunctionPara) SetMode(mode Mode) *GetTradeDatesFunctionPara {
	para.Mode = &mode
	return para
}

func (para *GetTradeDatesFunctionPara) SetDateType(dateType DateType) *GetTradeDatesFunctionPara {
	para.DateType = &dateType
	return para
}

func (para *GetTradeDatesFunctionPara) SetDateFormat(dateFormat DateFormat) *GetTradeDatesFunctionPara {
	para.DateFormat = &dateFormat
	return para
}

func (para *GetTradeDatesFunctionPara) SetPeriod(period Period) *GetTradeDatesFunctionPara {
	para.Period = &period
	return para
}

func (para *GetTradeDatesFunctionPara) SetPeriodnum(periodnum string) *GetTradeDatesFunctionPara {
	para.Periodnum = &periodnum
	return para
}

func (c *GetTradeDatesFunctionPara) ConvertToFunctionPara() *functionparadef.FunctionPara {
	data, _ := json.Marshal(c)
	var result functionparadef.FunctionPara
	json.Unmarshal(data, &result)
	return &result
}

type Mode string

const (
	Mode_Day      Mode = "1" //查询区间日期
	Mode_DayCount Mode = "2" //查询区间日期数目
)

type DateType string

const (
	DateType_TradeDate    DateType = "0" //交易日
	DateType_CalendarDate DateType = "1" //日历日
)

type DateFormat string

const (
	DateFormat_0 DateFormat = "0" //YYYY-MM-DD
	DateFormat_1 DateFormat = "1" //YYYY/MM/DD
	DateFormat_2 DateFormat = "2" //YYYYMMDD
)

type Period string

const (
	Period_D Period = "D" //日
	Period_W Period = "W" //周
	Period_M Period = "M" //月
	Period_Q Period = "Q" //季
	Period_S Period = "S" //半年
	Period_Y Period = "Y" //年
)
