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
	engin.POST("delete/:user_id", middleware.HystrixMiddleware, gin2grpc.TracerWrapper, middleware.AuthMiddleware(remoteAuth), deleteByUserID)
	engin.POST("/get_list", middleware.HystrixMiddleware, gin2grpc.TracerWrapper, middleware.AuthMiddleware(remoteAuth), getList)

	return
}

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
	if RemoteCallAbort(c, uResp, err) {
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

func deleteByUserID(c *gin.Context) {
	validator := DeleteByUserIDValidator{}
	err := validator.Bind(c)
	if err != nil {
		c.JSON(400, common.NewError(400, err))
		return
	}

	ctx, _ := gin2grpc.ContextWithSpan(c)

	uResp, err := remoteUser.Get(ctx, &validator.ReqUserGet)
	if RemoteCallAbort(c, uResp, err) {
		return
	}

	if !uResp.Result {
		c.JSON(404, common.NewErrorByStr(404, "用户不存在"))
		return
	}

	resp, err := remoteFriend.DelByUserID(ctx, &validator.Req)
	if RemoteCallAbort(c, resp, err) {
		return
	}

	serializer := DeleteByUserIDSerializer(resp)
	c.JSON(200, serializer)

}

//获取朋友列表
func getList(c *gin.Context) {
	validator := GetListValidator{}
	err := validator.Bind(c)
	if err != nil {
		c.JSON(400, common.NewError(400, err))
		return
	}

	ctx, _ := gin2grpc.ContextWithSpan(c)

	resp, err := remoteFriend.GetList(ctx, &validator.Req)
	if RemoteCallAbort(c, resp, err) {
		return
	}

	if len(resp.List) == 0 {
		c.JSON(200, common.NewResponseEmptyList())
		return
	}

	userGetListReq := user.GetListReq{Ids: resp.List}

	userGetListResp, err := remoteUser.GetList(ctx, &userGetListReq)
	if RemoteCallAbort(c, resp, err) {
		return
	}

	serializer := GetListSerializer(userGetListResp)

	c.JSON(200, serializer)
}
