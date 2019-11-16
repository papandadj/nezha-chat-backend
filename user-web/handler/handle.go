package handler

import (
	"errors"

	"github.com/papandadj/nezha-chat-backend/common"
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

	engin.POST("/sign_up", middleware.HystrixMiddleware, gin2grpc.TracerWrapper, signUp)
	engin.POST("/login", middleware.HystrixMiddleware, gin2grpc.TracerWrapper, login)
	engin.POST("/get_list", middleware.HystrixMiddleware, gin2grpc.TracerWrapper, middleware.AuthMiddleware(remoteAuth), userGetList)

	return
}

//Response .
type Response struct {
}

//注册接口
func signUp(c *gin.Context) {
	validator := SignUpValidator{}
	err := validator.Bind(c)
	if err != nil {
		c.JSON(400, common.NewError(400, err))
		return
	}

	ctx, _ := gin2grpc.ContextWithSpan(c)

	resp, err := remoteUser.Post(ctx, &validator.Req)
	if RemoteCallAbort(c, resp, err) {
		return
	}

	serializer := SignUpSerializer(resp)

	c.JSON(200, serializer)
}

func login(c *gin.Context) {
	validator := LoginValidator{}
	err := validator.Bind(c)
	if err != nil {
		c.JSON(400, common.NewError(400, err))
		return
	}

	ctx, _ := gin2grpc.ContextWithSpan(c)
	resp, err := remoteUser.CheckPassword(ctx, &validator.Req)

	if RemoteCallAbort(c, resp, err) {
		return
	}

	if !resp.Result {
		c.JSON(401, common.NewError(400, errors.New("账号或者密码错误")))
		return
	}

	tokenResp, err := remoteAuth.GetToken(ctx, &auth.GetTokenReq{
		Id:       resp.User.Id,
		Username: resp.User.Username,
	})

	if RemoteCallAbort(c, tokenResp, err) {
		return
	}

	serializer := LoginSerializer(tokenResp)

	c.JSON(200, serializer)
}

func userGetList(c *gin.Context) {
	validator := GetListValidator{}
	err := validator.Bind(c)
	if err != nil {
		c.JSON(400, common.NewError(400, err))
		return
	}

	ctx, _ := gin2grpc.ContextWithSpan(c)

	resp, err := remoteUser.GetList(ctx, &validator.Req)

	if RemoteCallAbort(c, resp, err) {
		return
	}

	serializer := GetListSerializer(resp)
	c.JSON(200, serializer)
}
