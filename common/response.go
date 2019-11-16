package common

import (
	"reflect"
)

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

//NewResponseEmptyList .
func NewResponseEmptyList() ResponseEmptyList {
	return ResponseEmptyList{
		List: []string{},
	}
}

//ResponseEmptyList .
type ResponseEmptyList struct {
	List []string `json:"list"`
}

//RemoteResponseError 远程调用的时候， 返回解析返回的error msg跟error code .
func RemoteResponseError(iResp interface{}, err error) (code int64, msg string, abort bool) {
	if err != nil {
		abort = true
		code = 500
		msg = err.Error()
		return
	}

	tResp := reflect.TypeOf(iResp)
	vResp := reflect.ValueOf(iResp)

	tResp = tResp.Elem()
	vResp = vResp.Elem()

	for i := 0; i < tResp.NumField(); i++ {
		tRespField := tResp.Field(i)
		if tRespField.Name == "Error" {

			errVal := vResp.Field(i).Interface()

			tErrVal := reflect.TypeOf(errVal)
			vErrVal := reflect.ValueOf(errVal)

			tErrVal = tErrVal.Elem()
			vErrVal = vErrVal.Elem()

			if !vErrVal.IsValid() {
				return
			}

			abort = true
			for j := 0; j < tErrVal.NumField(); j++ {
				tErrValField := tErrVal.Field(j)
				if tErrValField.Name == "Code" {
					code = vErrVal.Field(j).Interface().(int64)
				} else if tErrValField.Name == "Msg" {
					msg = vErrVal.Field(j).Interface().(string)
				}
			}
		}

	}

	return
}
