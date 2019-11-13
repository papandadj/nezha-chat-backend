package handler

import (
	"github.com/papandadj/nezha-chat-backend/pkg/tracer/gin2grpc"

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

	remoteUser = user.NewUserService(cfg.Remote.User, cl.Client())
	remoteAuth = auth.NewAuthService(cfg.Remote.Auth, cl.Client())

	engin.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"Msg": "pong"})
	})

	engin.POST("/test", gin2grpc.TracerWrapper, middleware.HystrixMiddleware(test))
	engin.POST("/sign_up", gin2grpc.TracerWrapper, middleware.HystrixMiddleware(signUp))
	engin.POST("/login", gin2grpc.TracerWrapper, middleware.HystrixMiddleware(login))

	return
}

//Response .
type Response struct {
}

//注册接口
func signUp(c *gin.Context) {
	req := new(user.PostReq)

	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(400, "参数错误")
		return
	}

	ctx, _ := gin2grpc.ContextWithSpan(c)

	resp, err := remoteUser.Post(ctx, req)
	if err != nil {
		logger.Errorln(err)
		c.JSON(500, "")
		return
	}

	if resp.Error != nil {
		c.JSON(400, resp.Error)
		return
	}

	c.JSON(200, resp)
}

func login(c *gin.Context) {
	req := new(user.CheckPasswordReq)
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(400, "参数错误")
		return
	}

	ctx, _ := gin2grpc.ContextWithSpan(c)
	resp, err := remoteUser.CheckPassword(ctx, req)

	if err != nil {
		logger.Errorln(err)
		c.JSON(500, "")
		return
	}

	if resp.Error != nil {
		c.JSON(400, resp.Error)
		return
	}

	//失败直接返回
	if !resp.Result {
		c.JSON(400, "账号或者密码错误")
		return
	}

	tokenResp, err := remoteAuth.GetToken(ctx, &auth.GetTokenReq{
		Id:       resp.User.Id,
		Username: resp.User.Username,
	})

	if err != nil {
		logger.Errorln(err)
		c.JSON(500, "")
		return
	}

	if resp.Error != nil {
		c.JSON(400, resp.Error)
		return
	}

	c.JSON(200, tokenResp)
}

func test(ctx *gin.Context) {

}
