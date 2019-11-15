package handler

import (
	"github.com/papandadj/nezha-chat-backend/common"
	"github.com/papandadj/nezha-chat-backend/pkg/tracer/gin2grpc"

	hystrixGo "github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/papandadj/nezha-chat-backend/friend-web/conf"
	"github.com/papandadj/nezha-chat-backend/pkg/log"
	"github.com/papandadj/nezha-chat-backend/pkg/middleware"
	"github.com/papandadj/nezha-chat-backend/proto/auth"
	"github.com/papandadj/nezha-chat-backend/proto/friend"
	"github.com/papandadj/nezha-chat-backend/proto/user"
)

var (
	logger       log.Logger
	cfg          *conf.Config
	remoteUser   user.UserService
	remoteAuth   auth.AuthService
	remoteFriend friend.FriendService
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
	remoteFriend = friend.NewFriendService(cfg.Remote.Friend, cl.Client())

	engin.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"Msg": "pong"})
	})

	engin.POST("/post", middleware.HystrixMiddleware, gin2grpc.TracerWrapper, middleware.AuthMiddleware(remoteAuth), post)
	engin.POST("/get_list", middleware.HystrixMiddleware, gin2grpc.TracerWrapper, middleware.AuthMiddleware(remoteAuth), getList)

	return
}

//TODO:

//添加朋友
func post(c *gin.Context) {
	validator := PostValidator{}
	err := validator.Bind(c)
	if err != nil {
		c.JSON(400, common.NewError(400, err))
		return
	}

	ctx, _ := gin2grpc.ContextWithSpan(c)

	uResp, err := remoteUser.Get(ctx, &validator.ReqUserGet)
	if RemoteCallAbortUser(c, uResp, err) {
		return
	}

	if !uResp.Result {
		c.JSON(404, common.NewErrorByStr(404, "用户不存在"))
		return
	}

	resp, err := remoteFriend.Post(ctx, &validator.Req)
	if RemoteCallAbort(c, resp, err) {
		return
	}

	serializer := PostSerializer(resp)
	c.JSON(200, serializer)
}

//获取朋友列表
func getList(c *gin.Context) {
	validator := PostValidator{}
	err := validator.Bind(c)
	if err != nil {
		c.JSON(400, common.NewError(400, err))
		return
	}

	ctx, _ := gin2grpc.ContextWithSpan(c)

	resp, err := remoteFriend.Post(ctx, &validator.Req)
	if RemoteCallAbort(c, resp, err) {
		return
	}

	serializer := PostSerializer(resp)
	c.JSON(200, serializer)
}
