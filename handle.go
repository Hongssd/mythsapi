package mythsapi

import (
	"fmt"
)

type ThsErrorRes struct {
	ErrorCode int    `json:"errorcode"`
	ErrMsg    string `json:"errmsg"`
}

type ThsRestRes[T any] struct {
	ThsErrorRes //错误信息
	DataRes[T]
}

type DataRes[T any] struct {
	Data        T           `json:"data"`        //请求结果
	Tables      T           `json:"tables"`      //请求结果
	DataType    []DataType  `json:"dataType"`    //请求数据类型
	InputParams InputParams `json:"inputParams"` //请求参数
	DataVol     int64       `json:"dataVol"`     //请求数据量
	Perf        int64       `json:"perf"`        //请求耗时
}

type DataType struct {
	ItemId string `json:"itemid"` //请求数据类型
	Type   string `json:"type"`   //请求数据类型
}

type InputParams struct {
	Indexs string `json:"indexs"` //请求数据类型
}

func handlerCommonRest[T any](data []byte) (*ThsRestRes[T], error) {
	res := &ThsRestRes[T]{}
	// log.Warn(string(data))
	err := json.Unmarshal(data, &res)
	if err != nil {
		log.Error("rest返回值获取失败", err)
	}
	return res, err
}
func (err *ThsErrorRes) handlerError() error {
	if err.ErrorCode != 0 {
		return fmt.Errorf("request error:[code:%v][message:%v]", err.ErrorCode, err.ErrMsg)
	} else {
		return nil
	}

}
