package model

import "net/http"

type Response struct {
	Data map[string]interface{} `json:"data"`
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
}

func NewResponse() *Response {
	res := new(Response)
	res.Code = http.StatusOK
	res.Data = make(map[string]interface{})
	res.Msg = ""
	return res
}

func ErrorResponse(errno int, msg string) *Response {
	r := NewResponse()
	r.Code = errno
	r.SetErrorInfo(errno, msg)
	return r
}

func (s *Response) SetErrorInfo(errno int, msg string) {
	s.Code = errno
	s.Msg = msg
	//s.Data["msg"]=msg
}
func (s *Response) SetSuccessInfo(code int, msg string) {
	s.Code = code
	s.Msg = msg
	//s.Data["msg"]=msg
}

func (s *Response) SetResponseDataInfo(key string, value string) {
	s.Data[key] = value
}

func (s *Response) AddResponseInfo(key string, val interface{}) {
	s.Data[key] = val
}
