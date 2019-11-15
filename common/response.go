package common

//NewError .
func NewError(code int64, err error) ResponseError {
	return ResponseError{
		Code: code,
		Msg:  err.Error(),
	}
}

//NewErrorByStr .
func NewErrorByStr(code int64, str string) ResponseError {
	return ResponseError{
		Code: code,
		Msg:  str,
	}
}

//ResponseError 给用户返回时错误定义
type ResponseError struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}
