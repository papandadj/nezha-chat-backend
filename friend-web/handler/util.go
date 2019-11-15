package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/papandadj/nezha-chat-backend/common"
	"github.com/papandadj/nezha-chat-backend/proto/friend"
	"github.com/papandadj/nezha-chat-backend/proto/user"
)

//RemoteCallAbort 检验远程调用返回值是否正确,
// abort否要继续执行代码
func RemoteCallAbort(c *gin.Context, resp RespError, err error) (abort bool) {
	if err != nil {
		abort = true
		logger.Errorln(err)
		c.JSON(500, common.NewError(500, err))
		return
	}

	respI, ok := resp.(RespError)
	if !ok {
		abort = true
		logger.Errorln("rpc返回的数据有问题或者服务调用出错")
		c.JSON(500, common.NewError(500, err))
		return
	}

	respErr := respI.GetError()

	if respErr != nil {
		abort = true
		logger.Infoln(respErr)
		c.JSON(int(respErr.Code), respErr)
		return
	}

	return
}

//RespError .
type RespError interface {
	GetError() *friend.Error
}

//RemoteCallAbortUser .
func RemoteCallAbortUser(c *gin.Context, resp RespErrorUser, err error) (abort bool) {
	if err != nil {
		abort = true
		logger.Errorln(err)
		c.JSON(500, common.NewError(500, err))
		return
	}

	respI, ok := resp.(RespErrorUser)
	if !ok {
		abort = true
		logger.Errorln("rpc返回的数据有问题或者服务调用出错")
		c.JSON(500, common.NewError(500, err))
		return
	}

	respErr := respI.GetError()

	if respErr != nil {
		abort = true
		logger.Infoln(respErr)
		c.JSON(int(respErr.Code), respErr)
		return
	}

	return
}

//RespErrorUser .
type RespErrorUser interface {
	GetError() *user.Error
}
