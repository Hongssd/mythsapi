package main

import (
	"path"
	"runtime"
	"strconv"
	"time"

	"github.com/Hongssd/mythsapi"
	"github.com/Hongssd/mythsapi/codesdef"
	"github.com/Hongssd/mythsapi/codesdef/futures_code"
	"github.com/Hongssd/mythsapi/functionparadef/cmd_history_quotation_functionpara"
	"github.com/Hongssd/mythsapi/indicatorsdef/cmd_history_quotation_indicator"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var log = logrus.New()

func init() {
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
		FullTimestamp:   true,
		ForceColors:     true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := "[" + path.Base(frame.File) + ":" + strconv.Itoa(frame.Line) + "]"
			return "", fileName
		},
	})
	log.SetReportCaller(true)
	log.Level = logrus.DebugLevel
	mythsapi.SetLogger(log)
}

func main() {
	refresh_token := "eyJzaWduX3RpbWUiOiIyMDIzLTEyLTEyIDE3OjI1OjQ5In0=.eyJ1aWQiOiI3MDExMDY3NzgiLCJ1c2VyIjp7ImFjY291bnQiOiJha2tqMDY3IiwiYXV0aFVzZXJJbmZvIjp7fSwiY29kZUNTSSI6W10sImNvZGVaekF1dGgiOltdLCJoYXNBSVByZWRpY3QiOmZhbHNlLCJoYXNBSVRhbGsiOmZhbHNlLCJoYXNDSUNDIjpmYWxzZSwiaGFzQ1NJIjpmYWxzZSwiaGFzRXZlbnREcml2ZSI6ZmFsc2UsImhhc0ZUU0UiOmZhbHNlLCJoYXNGdW5kVmFsdWF0aW9uIjpmYWxzZSwiaGFzSEsiOnRydWUsImhhc0xNRSI6ZmFsc2UsImhhc0xldmVsMiI6ZmFsc2UsImhhc1VTIjpmYWxzZSwiaGFzVVNBSW5kZXgiOmZhbHNlLCJtYXJrZXRDb2RlIjoiMTY7MzI7MTQ0Ozk2OzE3NjsxMTI7ODg7NDg7MTI4OzE2OC0xOzE4NDsyMDA7MjE2OzEwNDsxMjA7MTM2OzIzMjs1Njs2NDsiLCJtYXhPbkxpbmUiOjEsInByb2R1Y3RUeXBlIjoiU1VQRVJDT01NQU5EUFJPRFVDVCIsInJlZnJlc2hUb2tlbkV4cGlyZWRUaW1lIjoiMjAyNC0wMS0xMSAxNzoyMToxNiIsInNlc3NzaW9uIjoiOTY4NWQwOWQwZmFiZmYyY2ZlYTIxNDE0NTFkYTQyNGQiLCJzaWRJbmZvIjp7fSwidWlkIjoiNzAxMTA2Nzc4IiwidXNlclR5cGUiOiJGUkVFSUFMIiwid2lmaW5kTGltaXRNYXAiOnt9fX0=.287837EE4D7D83956CC1F9D4225F2E9CD6628E72A601AB392D8116A0DDDAE2DC"

	ths := mythsapi.MyThs{}

	client := ths.NewRestClient(refresh_token)
	err := client.RefreshAccessToken()
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("current client:", client)

	codes := []codesdef.Code{
		futures_code.GC0W_CMX,
		futures_code.SI0W_CMX,
		futures_code.HG0W_CMX,
		futures_code.PL0W_NMX,
		futures_code.CL0W_NMX,
		futures_code.NG0W_NMX,
		futures_code.RB0W_NMX,
	}

	indicators := []cmd_history_quotation_indicator.CmdHistoryQuotationIndicator{
		cmd_history_quotation_indicator.PRE_CLOSE,
		cmd_history_quotation_indicator.OPEN,
		cmd_history_quotation_indicator.HIGH,
		cmd_history_quotation_indicator.LOW,
		cmd_history_quotation_indicator.CLOSE,
		cmd_history_quotation_indicator.VOLUME,
	}

	api := client.NewCmdHistoryQuotationAPI()
	functionPara := cmd_history_quotation_functionpara.NewFunctionPara().
		SetInterval(cmd_history_quotation_functionpara.Interval_Week)
	res, err := api.
		AddCode(codes...).
		AddIndicators(indicators...).
		SetStartDate("2023-11-01").
		SetEndDate("2023-12-30").
		SetFunctionPara(functionPara).
		Do()
	data, _ := json.Marshal(res)
	log.Info(string(data))

	startTime, _ := time.ParseInLocation("2006-01-02", "2023-11-01", time.Local)
	endTime, _ := time.ParseInLocation("2006-01-02", "2023-12-30", time.Local)

	klines, err := client.QueryCmdHistoryQuotationToKlines(startTime.UnixMilli(), endTime.UnixMilli(), *functionPara, codes...)
	if err != nil {
		log.Error(err)
		return
	}
	data, _ = json.Marshal(klines)
	log.Info(string(data))

}
