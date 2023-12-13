package mythsapi

import "time"

type Kline struct {
	Symbol    string  `json:"symbol"`
	Timestamp int64   `json:"timestamp"`
	StartTime int64   `json:"start_time"`
	CloseTime int64   `json:"close_time"`
	PreClose  float64 `json:"preClose"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Volume    float64 `json:"volume"`
}

type Depth struct {
	Symbol    string `json:"symbol"`
	Timestamp int64  `json:"timestamp"`
	Bids      []PriceLevel
	Asks      []PriceLevel
}

type PriceLevel struct {
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
}

func (pl *PriceLevel) ToPriceAndQuantity() (float64, float64) {
	return pl.Price, pl.Quantity
}

func GetKlineStartAndEndByInterval(timestamp int64, interval int64) (int64, int64) {
	t := time.UnixMilli(timestamp)
	zeroTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	start := zeroTime.UnixMilli()
	end := zeroTime.UnixMilli()
	for start <= t.UnixMilli()-interval {
		start += interval
		end = start + interval - 1
	}

	return start, end
}
