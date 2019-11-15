package middleware

import (
	"context"
	"errors"

	"github.com/papandadj/nezha-chat-backend/common"

	"github.com/gin-gonic/gin"
	"github.com/papandadj/nezha-chat-backend/pkg/tracer/gin2grpc"
	"github.com/papandadj/nezha-chat-backend/proto/auth"
)

//AuthMiddleware .
func AuthMiddleware(remote auth.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx context.Context
		var ok bool

		ctx, ok = gin2grpc.ContextWithSpan(c)
		if !ok {
			ctx = context.Background()
		}

		token := c.GetHeader("token")
		if token == "" {
			c.Abort()
			c.JSON(403, common.NewError(403, errors.New("用户没权限")))
			return
		}

		resp, err := remote.Check(ctx, &auth.CheckReq{
			Token: token,
		})

		if err != nil {
			c.Abort()
			c.JSON(500, err)
			return
		}

		if resp.Error != nil {
			c.Abort()
			c.JSON(int(resp.Error.Code), resp.Error)
			return
		}

		c.Set(common.AttachAuthKey, common.AttachAuth{
			Username: resp.Username,
			ID:       resp.Id,
		})

	}
}

// AuthWithGin 返回context
func AuthWithGin(c *gin.Context) (auth common.AttachAuth, ok bool) {
	v, exist := c.Get(common.AttachAuthKey)
	if exist == false {
		ok = false
		return
	}

	auth, ok = v.(common.AttachAuth)
	return
}
