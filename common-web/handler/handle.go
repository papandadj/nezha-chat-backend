package handler

import (
	"github.com/papandadj/nezha-chat-backend/common"
	"github.com/papandadj/nezha-chat-backend/pkg/tracer/gin2grpc"

	hystrixGo "github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/papandadj/nezha-chat-backend/common-web/conf"
	"github.com/papandadj/nezha-chat-backend/pkg/log"
	"github.com/papandadj/nezha-chat-backend/pkg/middleware"
	PCommon "github.com/papandadj/nezha-chat-backend/proto/common"
)

var (
	logger       log.Logger
	cfg          *conf.Config
	remoteCommon PCommon.CommonService
)

func init() {
	logger = log.Base()
}

//Init remote service handler .
func Init() {
	cfg = conf.GetGlobalConfig()
	hystrixGo.DefaultTimeout = cfg.Hystrix.DefaultTimeout
	hystrixGo.DefaultVolumeThreshold = cfg.Hystrix.DefaultVolumeThreshold
	hystrixGo.DefaultErrorPercentThreshold = cfg.Hystrix.DefaultErrorPercentThreshold
	hystrixGo.DefaultSleepWindow = cfg.Hystrix.DefaultSleepWindow
	hystrixGo.DefaultMaxConcurrent = cfg.Hystrix.DefaultMaxConcurrent
}

//NewHTTPHandler .
func NewHTTPHandler(cl micro.Service) (engin *gin.Engine) {
	engin = gin.New()

	engin.Use(middleware.CORSMiddleware())

	remoteCommon = PCommon.NewCommonService(cfg.Remote.Common, cl.Client())

	engin.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"Msg": "pong"})
	})

	engin.POST("/get_home_img", middleware.HystrixMiddleware, gin2grpc.TracerWrapper, getHomeImg)

	return
}

func getHomeImg(c *gin.Context) {
	ctx, _ := gin2grpc.ContextWithSpan(c)

	resp, err := remoteCommon.GetList(ctx, &PCommon.GetListReq{})
	if RemoteCallAbort(c, resp, err) {
		return
	}

	c.JSON(200, resp)
}

//RemoteCallAbort 检验远程调用返回值是否正确,
// abort否要继续执行代码
func RemoteCallAbort(c *gin.Context, resp interface{}, err error) (abort bool) {
	var code int64
	var msg string
	code, msg, abort = common.RemoteResponseError(resp, err)
	if abort {
		codeInt := int(code)
		c.JSON(codeInt, common.NewErrorByStr(code, msg))
	}
	return
}
