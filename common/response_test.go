package common

import (
	"errors"
	"testing"
)

func TestRemoteResponseError(t *testing.T) {
	type Error struct {
		Code int64  `json:"code"`
		Msg  string `json:"msg"`
	}

	type Response struct {
		Str   string
		Error *Error
	}
	var tests = []struct {
		Resp        interface{}
		Err         error
		ReturnCode  int64
		ReturnMsg   string
		ReturnAbort bool
	}{
		{&Response{Str: "eee", Error: &Error{Code: 400, Msg: "用户参数错误"}}, nil, 400, "用户参数错误", true},
		{&Response{Str: "eee", Error: nil}, nil, 0, "", false},
		{nil, errors.New("服务调用错误"), 500, "服务调用错误", true},
	}

	for _, test := range tests {

		code, msg, abort := RemoteResponseError(test.Resp, test.Err)
		if code != test.ReturnCode || msg != test.ReturnMsg || abort != test.ReturnAbort {
			t.Errorf("RemoteResponseError failed to input %v, expected code = %d msg = %s abort = %v, got code = %d msg = %s abort = %v", test, test.ReturnCode, test.ReturnMsg, test.ReturnAbort, code, msg, abort)
		}
	}
}
