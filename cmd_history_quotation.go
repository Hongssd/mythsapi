package mythsapi

import (
	"time"

	"github.com/Hongssd/mythsapi/codesdef"
	"github.com/Hongssd/mythsapi/functionparadef"
	"github.com/Hongssd/mythsapi/functionparadef/cmd_history_quotation_functionpara"
	"github.com/Hongssd/mythsapi/indicatorsdef"
	"github.com/Hongssd/mythsapi/indicatorsdef/cmd_history_quotation_indicator"
)

const (
	API_CMD_HISTORY_QUOTATION = "api/v1/cmd_history_quotation"
)

type CmdHistoryQuotationAPI struct {
	RestClient
	formData cmdHistoryQuotationFormData
}

type cmdHistoryQuotationFormData struct {
	Codes        *codesdef.Codes               `json:"codes"`
	Indicators   *indicatorsdef.Indicators     `json:"indicators"`
	StartDate    string                        `json:"startdate"`
	EndDate      string                        `json:"enddate"`
	FunctionPara *functionparadef.FunctionPara `json:"functionpara"`
}

func (api *CmdHistoryQuotationAPI) AddIndicators(indicators ...cmd_history_quotation_indicator.CmdHistoryQuotationIndicator) *CmdHistoryQuotationAPI {
	if api.formData.Indicators == nil {
		api.formData.Indicators = &indicatorsdef.Indicators{}
	}
	api.formData.Indicators.AddIndicators(cmd_history_quotation_indicator.CmdHistoryQuotationIndicators(indicators).ConvertToIndicators()...)
	return api
}

func (api *CmdHistoryQuotationAPI) AddCode(codes ...codesdef.Code) *CmdHistoryQuotationAPI {
	if api.formData.Codes == nil {
		api.formData.Codes = &codesdef.Codes{}
	}
	api.formData.Codes.AddCodes(codes...)
	return api
}

func (api *CmdHistoryQuotationAPI) SetStartDate(startDate string) *CmdHistoryQuotationAPI {
	api.formData.StartDate = startDate
	return api
}

func (api *CmdHistoryQuotationAPI) SetEndDate(endDate string) *CmdHistoryQuotationAPI {
	api.formData.EndDate = endDate
	return api
}

func (api *CmdHistoryQuotationAPI) SetFunctionPara(functionPara *cmd_history_quotation_functionpara.CmdHistoryQuotationFunctionPara) *CmdHistoryQuotationAPI {
	api.formData.FunctionPara = functionPara.ConvertToFunctionPara()
	return api
}

func (client *RestClient) NewCmdHistoryQuotationAPI() *CmdHistoryQuotationAPI {
	return &CmdHistoryQuotationAPI{
		RestClient: *client,
		formData:   cmdHistoryQuotationFormData{},
	}
}

type CmdHistoryQuotationResRow struct {
	MarketCategory string                   `json:"marketCategory"` //市场类别
	Pricetype      int                      `json:"pricetype"`      //价格类型
	Thscode        string                   `json:"thscode"`        //同花顺代码
	Time           []string                 `json:"time"`           //时间序列
	Table          map[string][]interface{} `json:"table"`          //数据序列
}

type CmdHistoryQuotationRes []CmdHistoryQuotationResRow

func (api *CmdHistoryQuotationAPI) Do() (*ThsRestRes[CmdHistoryQuotationRes], error) {
	url := thsHandlerRequestAPIWithoutPathQueryParam(REST, API_CMD_HISTORY_QUOTATION)
	reqBody, err := json.Marshal(api.formData)
	if err != nil {
		return nil, err
	}
	return thsCallAPIWithAccessToken[CmdHistoryQuotationRes](&api.Client, url, reqBody, POST)
}

func (client *RestClient) QueryCmdHistoryQuotationToKlines(startTime, endTime int64, functionPara cmd_history_quotation_functionpara.CmdHistoryQuotationFunctionPara, codes ...codesdef.Code) (*[]Kline, error) {
	klines := []Kline{}
	api := client.NewCmdHistoryQuotationAPI()
	indicators := []cmd_history_quotation_indicator.CmdHistoryQuotationIndicator{
		cmd_history_quotation_indicator.PRE_CLOSE,
		cmd_history_quotation_indicator.OPEN,
		cmd_history_quotation_indicator.HIGH,
		cmd_history_quotation_indicator.LOW,
		cmd_history_quotation_indicator.CLOSE,
		cmd_history_quotation_indicator.VOLUME,
	}

	startDate := time.UnixMilli(startTime).Format("2006-01-02")
	endDate := time.UnixMilli(endTime).Format("2006-01-02")

	res, err := api.
		AddCode(codes...).
		AddIndicators(indicators...).
		SetStartDate(startDate).
		SetEndDate(endDate).
		SetFunctionPara(&functionPara).
		Do()
	if err != nil {
		return nil, err
	}
	for _, row := range res.Tables {
		// log.Info(row.Table)
		symbol := row.Thscode

		timestrs := row.Time
		preCloses := row.Table[string(cmd_history_quotation_indicator.PRE_CLOSE)]
		opens := row.Table[string(cmd_history_quotation_indicator.OPEN)]
		highs := row.Table[string(cmd_history_quotation_indicator.HIGH)]
		lows := row.Table[string(cmd_history_quotation_indicator.LOW)]
		closes := row.Table[string(cmd_history_quotation_indicator.CLOSE)]
		volumes := row.Table[string(cmd_history_quotation_indicator.VOLUME)]

		for i := 0; i < len(opens); i++ {

			timestamp, err := tradeDateToTimestamp(timestrs[i])
			if err != nil {
				return nil, err
			}

			preClose := 0.0
			if len(preCloses) > i && preCloses[i] != nil {
				preClose = preCloses[i].(float64)
			}
			open := 0.0
			if len(opens) > i && opens[i] != nil {
				open = opens[i].(float64)
			}
			high := 0.0
			if len(highs) > i && highs[i] != nil {
				high = highs[i].(float64)
			}
			low := 0.0
			if len(lows) > i && lows[i] != nil {
				low = lows[i].(float64)
			}
			close := 0.0
			if len(closes) > i && closes[i] != nil {
				close = closes[i].(float64)
			}
			volume := 0.0
			if len(volumes) > i && volumes[i] != nil {
				volume = volumes[i].(float64)
			}

			kline := Kline{
				Symbol:    symbol,
				Timestamp: time.Now().UnixMilli(),
				StartTime: timestamp,
				CloseTime: timestamp + functionPara.Interval.IntervalForTimeMillisecond() - 1,
				PreClose:  preClose,
				Open:      open,
				High:      high,
				Low:       low,
				Close:     close,
				Volume:    volume,
			}
			if kline.Open != 0 && kline.High != 0 && kline.Low != 0 && kline.Close != 0 {
				klines = append(klines, kline)
			}
		}

	}
	return &klines, nil
}
