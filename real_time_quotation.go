package mythsapi

import (
	"github.com/Hongssd/mythsapi/codesdef"
	"github.com/Hongssd/mythsapi/indicatorsdef"
	"github.com/Hongssd/mythsapi/indicatorsdef/real_time_quotation_indicator"
)

const (
	API_REAL_TIME_QUOTATION = "api/v1/real_time_quotation"
)

type RealTimeQuotationAPI struct {
	RestClient
	formData realTimeQuotationFormData
}

type realTimeQuotationFormData struct {
	Codes      *codesdef.Codes           `json:"codes"`
	Indicators *indicatorsdef.Indicators `json:"indicators"`
}

func (api *RealTimeQuotationAPI) AddIndicators(indicators ...real_time_quotation_indicator.RealTimeQuotationIndicator) *RealTimeQuotationAPI {
	if api.formData.Indicators == nil {
		api.formData.Indicators = &indicatorsdef.Indicators{}
	}
	api.formData.Indicators.AddIndicators(real_time_quotation_indicator.RealTimeQuotationIndicators(indicators).ConvertToIndicators()...)
	return api
}

func (api *RealTimeQuotationAPI) AddCode(codes ...codesdef.Code) *RealTimeQuotationAPI {
	if api.formData.Codes == nil {
		api.formData.Codes = &codesdef.Codes{}
	}
	api.formData.Codes.AddCodes(codes...)
	return api
}

func (client *RestClient) NewRealTimeQuotationAPI() *RealTimeQuotationAPI {
	return &RealTimeQuotationAPI{
		RestClient: *client,
		formData:   realTimeQuotationFormData{},
	}
}

type RealTimeQuotationResRow struct {
	MarketCategory string                   `json:"marketCategory"` //市场类别
	Pricetype      int                      `json:"pricetype"`      //价格类型
	Thscode        string                   `json:"thscode"`        //同花顺代码
	Time           []string                 `json:"time"`           //时间序列
	Table          map[string][]interface{} `json:"table"`          //数据序列
}

type RealTimeQuotationRes []RealTimeQuotationResRow

func (api *RealTimeQuotationAPI) Do() (*ThsRestRes[RealTimeQuotationRes], error) {
	url := thsHandlerRequestAPIWithoutPathQueryParam(REST, API_REAL_TIME_QUOTATION)
	reqBody, err := json.Marshal(api.formData)
	if err != nil {
		return nil, err
	}
	return thsCallAPIWithAccessToken[RealTimeQuotationRes](&api.Client, url, reqBody, POST)
}

func (client *RestClient) QueryRealTimeQuotationToKlinesAndDepths(codes ...codesdef.Code) (*[]Kline, *[]Depth, error) {

	klines := []Kline{}
	depths := []Depth{}
	api := client.NewRealTimeQuotationAPI()
	indicators := []real_time_quotation_indicator.RealTimeQuotationIndicator{
		real_time_quotation_indicator.TRADE_DATE,
		real_time_quotation_indicator.TRADE_TIME,
		real_time_quotation_indicator.PRE_CLOSE,
		real_time_quotation_indicator.OPEN,
		real_time_quotation_indicator.HIGH,
		real_time_quotation_indicator.LOW,
		real_time_quotation_indicator.LASTEST_PRICE,
		real_time_quotation_indicator.VOLUME,
		real_time_quotation_indicator.BID1,
		real_time_quotation_indicator.BID2,
		real_time_quotation_indicator.BID3,
		real_time_quotation_indicator.BID4,
		real_time_quotation_indicator.BID5,
		real_time_quotation_indicator.BID6,
		real_time_quotation_indicator.BID7,
		real_time_quotation_indicator.BID8,
		real_time_quotation_indicator.BID9,
		real_time_quotation_indicator.BID10,
		real_time_quotation_indicator.BID_SIZE1,
		real_time_quotation_indicator.BID_SIZE2,
		real_time_quotation_indicator.BID_SIZE3,
		real_time_quotation_indicator.BID_SIZE4,
		real_time_quotation_indicator.BID_SIZE5,
		real_time_quotation_indicator.BID_SIZE6,
		real_time_quotation_indicator.BID_SIZE7,
		real_time_quotation_indicator.BID_SIZE8,
		real_time_quotation_indicator.BID_SIZE9,
		real_time_quotation_indicator.BID_SIZE10,
		real_time_quotation_indicator.ASK1,
		real_time_quotation_indicator.ASK2,
		real_time_quotation_indicator.ASK3,
		real_time_quotation_indicator.ASK4,
		real_time_quotation_indicator.ASK5,
		real_time_quotation_indicator.ASK6,
		real_time_quotation_indicator.ASK7,
		real_time_quotation_indicator.ASK8,
		real_time_quotation_indicator.ASK9,
		real_time_quotation_indicator.ASK10,
		real_time_quotation_indicator.ASK_SIZE1,
		real_time_quotation_indicator.ASK_SIZE2,
		real_time_quotation_indicator.ASK_SIZE3,
		real_time_quotation_indicator.ASK_SIZE4,
		real_time_quotation_indicator.ASK_SIZE5,
		real_time_quotation_indicator.ASK_SIZE6,
		real_time_quotation_indicator.ASK_SIZE7,
		real_time_quotation_indicator.ASK_SIZE8,
		real_time_quotation_indicator.ASK_SIZE9,
		real_time_quotation_indicator.ASK_SIZE10,
	}
	res, err := api.AddCode(codes...).AddIndicators(indicators...).Do()
	if err != nil {
		return nil, nil, err
	}
	for _, row := range res.Tables {
		// log.Info(row.Table)
		symbol := row.Thscode
		tradeDates := row.Table[string(real_time_quotation_indicator.TRADE_DATE)]
		tradeTimes := row.Table[string(real_time_quotation_indicator.TRADE_TIME)]
		preCloses := row.Table[string(real_time_quotation_indicator.PRE_CLOSE)]
		opens := row.Table[string(real_time_quotation_indicator.OPEN)]
		highs := row.Table[string(real_time_quotation_indicator.HIGH)]
		lows := row.Table[string(real_time_quotation_indicator.LOW)]
		closes := row.Table[string(real_time_quotation_indicator.LASTEST_PRICE)]
		volumes := row.Table[string(real_time_quotation_indicator.VOLUME)]

		bid1s := row.Table[string(real_time_quotation_indicator.BID1)]
		bid2s := row.Table[string(real_time_quotation_indicator.BID2)]
		bid3s := row.Table[string(real_time_quotation_indicator.BID3)]
		bid4s := row.Table[string(real_time_quotation_indicator.BID4)]
		bid5s := row.Table[string(real_time_quotation_indicator.BID5)]
		bid6s := row.Table[string(real_time_quotation_indicator.BID6)]
		bid7s := row.Table[string(real_time_quotation_indicator.BID7)]
		bid8s := row.Table[string(real_time_quotation_indicator.BID8)]
		bid9s := row.Table[string(real_time_quotation_indicator.BID9)]
		bid10s := row.Table[string(real_time_quotation_indicator.BID10)]
		bid1Sizes := row.Table[string(real_time_quotation_indicator.BID_SIZE1)]
		bid2Sizes := row.Table[string(real_time_quotation_indicator.BID_SIZE2)]
		bid3Sizes := row.Table[string(real_time_quotation_indicator.BID_SIZE3)]
		bid4Sizes := row.Table[string(real_time_quotation_indicator.BID_SIZE4)]
		bid5Sizes := row.Table[string(real_time_quotation_indicator.BID_SIZE5)]
		bid6Sizes := row.Table[string(real_time_quotation_indicator.BID_SIZE6)]
		bid7Sizes := row.Table[string(real_time_quotation_indicator.BID_SIZE7)]
		bid8Sizes := row.Table[string(real_time_quotation_indicator.BID_SIZE8)]
		bid9Sizes := row.Table[string(real_time_quotation_indicator.BID_SIZE9)]
		bid10Sizes := row.Table[string(real_time_quotation_indicator.BID_SIZE10)]

		ask1s := row.Table[string(real_time_quotation_indicator.ASK1)]
		ask2s := row.Table[string(real_time_quotation_indicator.ASK2)]
		ask3s := row.Table[string(real_time_quotation_indicator.ASK3)]
		ask4s := row.Table[string(real_time_quotation_indicator.ASK4)]
		ask5s := row.Table[string(real_time_quotation_indicator.ASK5)]
		ask6s := row.Table[string(real_time_quotation_indicator.ASK6)]
		ask7s := row.Table[string(real_time_quotation_indicator.ASK7)]
		ask8s := row.Table[string(real_time_quotation_indicator.ASK8)]
		ask9s := row.Table[string(real_time_quotation_indicator.ASK9)]
		ask10s := row.Table[string(real_time_quotation_indicator.ASK10)]
		ask1Sizes := row.Table[string(real_time_quotation_indicator.ASK_SIZE1)]
		ask2Sizes := row.Table[string(real_time_quotation_indicator.ASK_SIZE2)]
		ask3Sizes := row.Table[string(real_time_quotation_indicator.ASK_SIZE3)]
		ask4Sizes := row.Table[string(real_time_quotation_indicator.ASK_SIZE4)]
		ask5Sizes := row.Table[string(real_time_quotation_indicator.ASK_SIZE5)]
		ask6Sizes := row.Table[string(real_time_quotation_indicator.ASK_SIZE6)]
		ask7Sizes := row.Table[string(real_time_quotation_indicator.ASK_SIZE7)]
		ask8Sizes := row.Table[string(real_time_quotation_indicator.ASK_SIZE8)]
		ask9Sizes := row.Table[string(real_time_quotation_indicator.ASK_SIZE9)]
		ask10Sizes := row.Table[string(real_time_quotation_indicator.ASK_SIZE10)]

		for i := 0; i < len(opens); i++ {

			timestamp, err := tradeDateAndTradeTimeToTimestamp(tradeDates[i].(string), tradeTimes[i].(string))
			if err != nil {
				return nil, nil, err
			}
			kline := Kline{
				Symbol:    symbol,
				Timestamp: timestamp,
				// PreClose:  preCloses[i].(float64),
				// Open:      opens[i].(float64),
				// High:      highs[i].(float64),
				// Low:       lows[i].(float64),
				// Close:     closes[i].(float64),
				// Volume:    volumes[i].(float64),
			}
			if v, ok := preCloses[i].(float64); ok {
				kline.PreClose = v
			}
			if v, ok := opens[i].(float64); ok {
				kline.Open = v
			}
			if v, ok := highs[i].(float64); ok {
				kline.High = v
			}
			if v, ok := lows[i].(float64); ok {
				kline.Low = v
			}
			if v, ok := closes[i].(float64); ok {
				kline.Close = v
			}
			if v, ok := volumes[i].(float64); ok {
				kline.Volume = v
			}

			klines = append(klines, kline)

			depth := Depth{
				Symbol:    symbol,
				Timestamp: timestamp,
			}

			bids := []PriceLevel{}
			asks := []PriceLevel{}
			if len(bid1s) > i && bid1s[i] != nil && len(bid1Sizes) > i && bid1Sizes[i] != nil {
				bids = append(bids, PriceLevel{Price: bid1s[i].(float64), Quantity: bid1Sizes[i].(float64)})
			}
			if len(bid2s) > i && bid2s[i] != nil && len(bid2Sizes) > i && bid2Sizes[i] != nil {
				bids = append(bids, PriceLevel{Price: bid2s[i].(float64), Quantity: bid2Sizes[i].(float64)})
			}
			if len(bid3s) > i && bid3s[i] != nil && len(bid3Sizes) > i && bid3Sizes[i] != nil {
				bids = append(bids, PriceLevel{Price: bid3s[i].(float64), Quantity: bid3Sizes[i].(float64)})
			}
			if len(bid4s) > i && bid4s[i] != nil && len(bid4Sizes) > i && bid4Sizes[i] != nil {
				bids = append(bids, PriceLevel{Price: bid4s[i].(float64), Quantity: bid4Sizes[i].(float64)})
			}
			if len(bid5s) > i && bid5s[i] != nil && len(bid5Sizes) > i && bid5Sizes[i] != nil {
				bids = append(bids, PriceLevel{Price: bid5s[i].(float64), Quantity: bid5Sizes[i].(float64)})
			}
			if len(bid6s) > i && bid6s[i] != nil && len(bid6Sizes) > i && bid6Sizes[i] != nil {
				bids = append(bids, PriceLevel{Price: bid6s[i].(float64), Quantity: bid6Sizes[i].(float64)})
			}
			if len(bid7s) > i && bid7s[i] != nil && len(bid7Sizes) > i && bid7Sizes[i] != nil {
				bids = append(bids, PriceLevel{Price: bid7s[i].(float64), Quantity: bid7Sizes[i].(float64)})
			}
			if len(bid8s) > i && bid8s[i] != nil && len(bid8Sizes) > i && bid8Sizes[i] != nil {
				bids = append(bids, PriceLevel{Price: bid8s[i].(float64), Quantity: bid8Sizes[i].(float64)})
			}
			if len(bid9s) > i && bid9s[i] != nil && len(bid9Sizes) > i && bid9Sizes[i] != nil {
				bids = append(bids, PriceLevel{Price: bid9s[i].(float64), Quantity: bid9Sizes[i].(float64)})
			}
			if len(bid10s) > i && bid10s[i] != nil && len(bid10Sizes) > i && bid10Sizes[i] != nil {
				bids = append(bids, PriceLevel{Price: bid10s[i].(float64), Quantity: bid10Sizes[i].(float64)})
			}

			if len(ask1s) > i && ask1s[i] != nil && len(ask1Sizes) > i && ask1Sizes[i] != nil {
				asks = append(asks, PriceLevel{Price: ask1s[i].(float64), Quantity: ask1Sizes[i].(float64)})
			}
			if len(ask2s) > i && ask2s[i] != nil && len(ask2Sizes) > i && ask2Sizes[i] != nil {
				asks = append(asks, PriceLevel{Price: ask2s[i].(float64), Quantity: ask2Sizes[i].(float64)})
			}
			if len(ask3s) > i && ask3s[i] != nil && len(ask3Sizes) > i && ask3Sizes[i] != nil {
				asks = append(asks, PriceLevel{Price: ask3s[i].(float64), Quantity: ask3Sizes[i].(float64)})
			}
			if len(ask4s) > i && ask4s[i] != nil && len(ask4Sizes) > i && ask4Sizes[i] != nil {
				asks = append(asks, PriceLevel{Price: ask4s[i].(float64), Quantity: ask4Sizes[i].(float64)})
			}
			if len(ask5s) > i && ask5s[i] != nil && len(ask5Sizes) > i && ask5Sizes[i] != nil {
				asks = append(asks, PriceLevel{Price: ask5s[i].(float64), Quantity: ask5Sizes[i].(float64)})
			}
			if len(ask6s) > i && ask6s[i] != nil && len(ask6Sizes) > i && ask6Sizes[i] != nil {
				asks = append(asks, PriceLevel{Price: ask6s[i].(float64), Quantity: ask6Sizes[i].(float64)})
			}
			if len(ask7s) > i && ask7s[i] != nil && len(ask7Sizes) > i && ask7Sizes[i] != nil {
				asks = append(asks, PriceLevel{Price: ask7s[i].(float64), Quantity: ask7Sizes[i].(float64)})
			}
			if len(ask8s) > i && ask8s[i] != nil && len(ask8Sizes) > i && ask8Sizes[i] != nil {
				asks = append(asks, PriceLevel{Price: ask8s[i].(float64), Quantity: ask8Sizes[i].(float64)})
			}
			if len(ask9s) > i && ask9s[i] != nil && len(ask9Sizes) > i && ask9Sizes[i] != nil {
				asks = append(asks, PriceLevel{Price: ask9s[i].(float64), Quantity: ask9Sizes[i].(float64)})
			}
			if len(ask10s) > i && ask10s[i] != nil && len(ask10Sizes) > i && ask10Sizes[i] != nil {
				asks = append(asks, PriceLevel{Price: ask10s[i].(float64), Quantity: ask10Sizes[i].(float64)})
			}

			depth.Bids = bids
			depth.Asks = asks
			depths = append(depths, depth)

		}

	}
	return &klines, &depths, nil
}
