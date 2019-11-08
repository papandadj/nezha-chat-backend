package handler

import (
	"time"

	hystrixGo "github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/papandadj/nezha-chat-backend/pkg/log"
	"github.com/papandadj/nezha-chat-backend/pkg/middleware"
	"github.com/papandadj/nezha-chat-backend/proto/auth"
	"github.com/papandadj/nezha-chat-backend/proto/user"
	"github.com/papandadj/nezha-chat-backend/user-web/conf"
)

var (
	logger     log.Logger
	cfg        *conf.Config
	remoteUser user.UserService
	remoteAuth auth.AuthService
)

func init() {

	cfg = conf.GetGlobalConfig()
	logger = log.Base()
	Init()
}

//Init remote service handler .
func Init() {
	hystrixGo.DefaultTimeout = cfg.Hystrix.DefaultTimeout
	hystrixGo.DefaultVolumeThreshold = cfg.Hystrix.DefaultVolumeThreshold
	hystrixGo.DefaultErrorPercentThreshold = cfg.Hystrix.DefaultErrorPercentThreshold
	hystrixGo.DefaultSleepWindow = cfg.Hystrix.DefaultSleepWindow
	hystrixGo.DefaultMaxConcurrent = cfg.Hystrix.DefaultMaxConcurrent
}

//NewHTTPHandler .
func NewHTTPHandler(cl micro.Service) (engin *gin.Engine) {
	engin = gin.New()

	remoteUser = user.NewUserService(cfg.Remote.User, cl.Client())
	remoteAuth = auth.NewAuthService(cfg.Remote.Auth, cl.Client())

	engin.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"Msg": "pong"})
	})

	engin.POST("/test", middleware.HystrixMiddleware(test))

	return
}

func test(ctx *gin.Context) {
	time.Sleep(2 * time.Second)
	logger.Info("ddddd")
	ctx.Status(200)

}
