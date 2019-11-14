package common

//NewError .
func NewError(code int64, err error) ResponseError {
	return ResponseError{
		Code: code,
		Msg:  err.Error(),
	}
}

//ResponseError 给用户返回时错误定义
type ResponseError struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}
