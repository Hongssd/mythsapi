package mythsapi

import (
	"fmt"
	"time"
)

type FormData struct {
	Codes []string `json:"code"`
}

func tradeDateAndTradeTimeToTimestamp(tradeDate, tradeTime string) (int64, error) {
	timeLocal, err := time.ParseInLocation("2006-01-02 15:04:05",
		fmt.Sprintf("%s %s", tradeDate, tradeTime), time.Local)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return timeLocal.UnixMilli(), nil
}

func tradeDateToTimestamp(tradeDate string) (int64, error) {
	return tradeDateAndTradeTimeToTimestamp(tradeDate, "00:00:00")
}
