package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/papandadj/nezha-chat-backend/common"
)

//RemoteCallAbort 检验远程调用返回值是否正确,
// abort否要继续执行代码
func RemoteCallAbort(c *gin.Context, resp interface{}, err error) (abort bool) {
	var code int64
	var msg string
	code, msg, abort = common.RemoteResponseError(resp, err)
	logger.Errorln(msg)
	if abort {
		codeInt := int(code)
		c.JSON(codeInt, common.NewErrorByStr(code, msg))
	}
	return
}
