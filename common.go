package mythsapi

import (
	"bytes"
	"compress/gzip"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"time"

	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

const (
	BIT_BASE_10 = 10
	BIT_SIZE_64 = 64
	BIT_SIZE_32 = 32
)

const (
	THS_ACCESS_TOKEN_EXPIRE_TIME = int64(3 * 86400 * 1000)
	API_GET_ACCESS_TOKEN         = "/api/v1/get_access_token"
)

type RequestType string

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

var NIL_REQBODY = []byte{}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var log = logrus.New()

func SetLogger(logger *logrus.Logger) {
	log = logger
}

func GetPointer[T any](v T) *T {
	return &v
}

func HmacSha256(secret, data string) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return h.Sum(nil)
}

// Request 发送请求
func Request(url string, reqBody []byte, method string, isGzip bool) ([]byte, error) {
	return RequestWithHeader(url, reqBody, method, map[string]string{}, isGzip)
}

func RequestWithHeader(url string, reqBody []byte, method string, headerMap map[string]string, isGzip bool) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	if isGzip { // 请求 header 添加 gzip
		req.Header.Add("Content-Encoding", "gzip")
		req.Header.Add("Accept-Encoding", "gzip")
	}
	req.Close = true
	req.Body = io.NopCloser(bytes.NewBuffer(reqBody))

	log.Debug("reqURL: ", req.URL.String())
	if len(reqBody) > 0 {
		log.Debug("reqBody: ", string(reqBody))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		body, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}
	data, err := io.ReadAll(body)
	log.Debug(string(data))
	return data, err
}

type MyThs struct {
}

const (
	THS_API_HTTP = "ft.10jqka.com.cn"
	IS_GZIP      = true
)

type APIType int

const (
	REST APIType = iota
)

type Client struct {
	RefreshToken           string
	AccessToken            string
	AccessTokenExpiredTime int64
}

type RestClient struct {
	Client
}

func (*MyThs) NewRestClient(RefreshToken string) *RestClient {
	client := &RestClient{
		Client{
			RefreshToken: RefreshToken,
		},
	}
	return client
}

type AccessTokenRes struct {
	AccessToken string `json:"access_token"` // 访问令牌
	ExpiredTime string `json:"expired_time"` // 过期时间
}

func (client *Client) RefreshAccessToken() error {
	url := thsHandlerRequestAPIWithoutPathQueryParam(REST, API_GET_ACCESS_TOKEN)
	body, err := RequestWithHeader(url.String(), []byte{}, POST,
		map[string]string{
			"Content-Type":  "application/json",
			"refresh_token": client.RefreshToken,
		}, IS_GZIP)
	if err != nil {
		return err
	}
	res, err := handlerCommonRest[AccessTokenRes](body)
	if err != nil {
		return err
	}
	if res.Data.AccessToken == "" {
		return fmt.Errorf("refresh access token error")
	}
	log.Debug(res)
	client.AccessToken = res.Data.AccessToken
	expiredTime, err := time.ParseInLocation("2006-01-02 15:04:05", res.Data.ExpiredTime, time.Local)
	if err != nil {
		return err
	}
	client.AccessTokenExpiredTime = expiredTime.UnixMilli()

	return nil
}

// 通用鉴权接口调用
func thsCallAPIWithAccessToken[T any](client *Client, url url.URL, reqBody []byte, method string) (*ThsRestRes[T], error) {
	if client.AccessToken == "" || time.Now().UnixMilli()+THS_ACCESS_TOKEN_EXPIRE_TIME > client.AccessTokenExpiredTime {
		err := client.RefreshAccessToken()
		if err != nil {
			return nil, err
		}
	}

	body, err := RequestWithHeader(url.String(), reqBody, method,
		map[string]string{
			"Content-Type": "application/json",
			"access_token": client.AccessToken,
		}, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}
	return res, res.handlerError()
}

// URL标准封装 带路径参数
func thsHandlerRequestAPIWithPathQueryParam[T any](apiType APIType, request *T, name string) url.URL {
	query := thsHandlerReq(request)
	u := url.URL{
		Scheme:   "https",
		Host:     ThsGetRestHostByAPIType(apiType),
		Path:     name,
		RawQuery: query,
	}
	return u
}

// URL标准封装 不带路径参数
func thsHandlerRequestAPIWithoutPathQueryParam(apiType APIType, name string) url.URL {
	// query := thsHandlerReq(request)
	u := url.URL{
		Scheme:   "https",
		Host:     ThsGetRestHostByAPIType(apiType),
		Path:     name,
		RawQuery: "",
	}
	return u
}

func thsHandlerReq[T any](req *T) string {
	var argBuffer bytes.Buffer

	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)
	if v.IsNil() {
		return ""
	}
	t = t.Elem()
	v = v.Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		argName := t.Field(i).Tag.Get("json")
		switch v.Field(i).Elem().Kind() {
		case reflect.String:
			argBuffer.WriteString(argName + "=" + v.Field(i).Elem().String() + "&")
		case reflect.Int, reflect.Int64:
			argBuffer.WriteString(argName + "=" + strconv.FormatInt(v.Field(i).Elem().Int(), BIT_BASE_10) + "&")
		case reflect.Float32, reflect.Float64:
			argBuffer.WriteString(argName + "=" + decimal.NewFromFloat(v.Field(i).Elem().Float()).String() + "&")
		case reflect.Bool:
			argBuffer.WriteString(argName + "=" + strconv.FormatBool(v.Field(i).Elem().Bool()) + "&")
		case reflect.Struct:
			sv := reflect.ValueOf(v.Field(i).Interface())
			ToStringMethod := sv.MethodByName("String")
			args := make([]reflect.Value, 0)
			result := ToStringMethod.Call(args)
			argBuffer.WriteString(argName + "=" + result[0].String() + "&")
		case reflect.Slice:
			s := v.Field(i).Interface()
			d, _ := json.Marshal(s)
			argBuffer.WriteString(argName + "=" + url.QueryEscape(string(d)) + "&")
		case reflect.Invalid:
		default:
			log.Errorf("req type error %s:%s", argName, v.Field(i).Elem().Kind())
		}
	}
	return strings.TrimRight(argBuffer.String(), "&")
}

func ThsGetRestHostByAPIType(apiType APIType) string {
	switch apiType {
	case REST:
		return THS_API_HTTP
	default:
		return ""
	}
}

func interfaceStringToFloat64(inter interface{}) float64 {
	return stringToFloat64(inter.(string))
}

func interfaceStringToInt64(inter interface{}) int64 {
	return int64(inter.(float64))
}

func stringToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, BIT_SIZE_64)
	return f
}
