package mythsapi

import (
	"github.com/Hongssd/mythsapi/codesdef"
	"github.com/Hongssd/mythsapi/functionparadef"
	get_trade_dates_functionpara "github.com/Hongssd/mythsapi/functionparadef/get_trade_dates"
)

const (
	API_GET_TRADE_DATES = "api/v1/get_trade_dates"
)

type GetTradeDatesAPI struct {
	RestClient
	formData getTradeDatesFormData
}

type getTradeDatesFormData struct {
	Marketcode   *codesdef.Code                `json:"marketcode"`
	StartDate    string                        `json:"startdate"`
	EndDate      string                        `json:"enddate"`
	FunctionPara *functionparadef.FunctionPara `json:"functionpara"`
}

func (api *GetTradeDatesAPI) SetMarketCode(code codesdef.Code) *GetTradeDatesAPI {
	api.formData.Marketcode = &code
	return api
}

func (api *GetTradeDatesAPI) SetStartDate(startDate string) *GetTradeDatesAPI {
	api.formData.StartDate = startDate
	return api
}

func (api *GetTradeDatesAPI) SetEndDate(endDate string) *GetTradeDatesAPI {
	api.formData.EndDate = endDate
	return api
}

func (api *GetTradeDatesAPI) SetFunctionPara(functionPara *get_trade_dates_functionpara.GetTradeDatesFunctionPara) *GetTradeDatesAPI {
	api.formData.FunctionPara = functionPara.ConvertToFunctionPara()
	return api
}

func (client *RestClient) NewGetTradeDatesAPI() *GetTradeDatesAPI {
	return &GetTradeDatesAPI{
		RestClient: *client,
		formData:   getTradeDatesFormData{},
	}
}

type GetTradeDatesRes struct {
	Time []string `json:"time"` //时间序列
}

func (api *GetTradeDatesAPI) Do() (*ThsRestRes[GetTradeDatesRes], error) {
	url := thsHandlerRequestAPIWithoutPathQueryParam(REST, API_GET_TRADE_DATES)
	reqBody, err := json.Marshal(api.formData)
	if err != nil {
		return nil, err
	}
	return thsCallAPIWithAccessToken[GetTradeDatesRes](&api.Client, url, reqBody, POST)
}
