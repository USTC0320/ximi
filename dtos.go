package ximi

import "reflect"

// PageInfo 统一报文-翻页信息
type PageInfo struct {
	PageNum  int64 `json:"pageNum"`
	PageSize int64 `json:"pageSize"`
	Total    int64 `json:"total"`
}

// Ret 统一报文
type Ret struct {
	Code     int      `json:"code"`
	Message  string   `json:"message"`
	Data     any      `json:"data"`
	PageInfo PageInfo `json:"pageInfo"`
}

// SetPageInfo 设置报文翻页信息
func (ret *Ret) SetPageInfo(pageNum int64, pageSize int64, total int64) *Ret {
	page := PageInfo{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    total,
	}
	ret.PageInfo = page
	return ret
}

func (ret *Ret) AddData(key string, data any) *Ret {
	if ret.Data != nil && reflect.TypeOf(ret.Data).Kind() == reflect.Map {
		v := reflect.ValueOf(ret.Data)
		i := v.Interface()
		m := i.(map[string]interface{})
		m[key] = data
		ret.Data = m
	} else {
		v := make(map[string]interface{})
		v[key] = data
		ret.Data = v
	}

	return ret
}

func Success() *Ret {
	r := new(Ret)
	r.Code = 0
	r.Message = "OK"
	r.Data = nil
	return r
}

func SuccessWithData(data any) *Ret {
	r := new(Ret)
	r.Code = 0
	r.Message = "OK"
	r.Data = data
	return r
}

func Error(msg string) *Ret {
	r := new(Ret)
	r.Code = -1
	r.Message = msg
	r.Data = nil
	return r
}

func ErrorWithCode(code int, msg string) *Ret {
	r := new(Ret)
	r.Code = code
	r.Message = msg
	r.Data = nil
	return r
}
