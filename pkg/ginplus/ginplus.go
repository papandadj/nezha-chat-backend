package ginplus

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

const (
	//RespStatusKey 返回状态
	RespStatusKey = "RespStatusKey"
	//RespMsgKey 返回的提示信息
	RespMsgKey = "RespMsgKey"
	//RespErrorKey 错误的返回信息
	RespErrorKey = "RespErrorKey"
	//RespDataKey 返回的数据
	RespDataKey = "RespDataKey"
)

//Success .
type Success struct {
	Msg string
}

var (
	//Blank 当操作成功不想返回数据的默认数据
	Blank = Success{"success"}
)

//RespJSON 结果返回， 最终在hystrix里面确定,
//@param v : 有两种类型， string跟json化
func RespJSON(ctx *gin.Context, status int, v interface{}) {
	ctx.Set(RespStatusKey, status)
	var buf []byte
	var err error

	msg, ok := v.(string)
	if ok {
		buf = []byte(msg)
	} else {
		buf, err = json.Marshal(v)
		if err != nil {
			panic(err)
		}
	}

	ctx.Set(RespDataKey, buf)
	ctx.Abort()

}
