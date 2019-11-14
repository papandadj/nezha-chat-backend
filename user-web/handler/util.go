package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/papandadj/nezha-chat-backend/common"
	"github.com/papandadj/nezha-chat-backend/proto/user"
)

//RemoteCallAbort 检验远程调用返回值是否正确,
// abort否要继续执行代码
func RemoteCallAbort(c *gin.Context, respError *user.Error, err error) (abort bool) {
	if err != nil {
		abort = true
		logger.Errorln(err)
		c.JSON(500, common.NewError(500, err))
		return
	}

	if respError != nil {
		abort = true
		logger.Infoln(respError)
		c.JSON(int(respError.Code), respError)
		return
	}

	return
}
