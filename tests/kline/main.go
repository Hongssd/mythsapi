package main

import (
	"path"
	"runtime"
	"strconv"

	"github.com/Hongssd/mythsapi"
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
}
