package middleware

import (
	"fmt"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

//HystrixMiddleware .
func HystrixMiddleware(handle gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Request.Method + "-" + ctx.Request.RequestURI
		hystrix.Do(name, func() error {
			handle(ctx)
			fmt.Println("--------->")
			return nil
		}, func(e error) error {
			fmt.Println("服务网熔断", e)
			return e
		})
	}
}
