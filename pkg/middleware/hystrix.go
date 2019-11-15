package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

//HystrixWrap is wrapper.
func HystrixWrap(handle gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Request.Method + "-" + ctx.Request.RequestURI
		hystrix.Do(name, func() error {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("hystrix catch err: ", r)
				}
			}()

			handle(ctx)
			status := ctx.Writer.Status()

			if status >= http.StatusInternalServerError {
				str := fmt.Sprintf("status code %d", status)
				return errors.New(str)
			}
			return nil
		}, func(e error) error {
			fmt.Println("hystrix : ", e)
			ctx.JSON(500, struct {
				Code int64
				Msg  string
			}{500, "server unavailable"})
			return e
		})
	}
}

//HystrixMiddleware .
func HystrixMiddleware(ctx *gin.Context) {
	name := ctx.Request.Method + "-" + ctx.Request.RequestURI
	hystrix.Do(name, func() error {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("hystrix catch err: ", r)
			}
		}()

		ctx.Next()
		status := ctx.Writer.Status()

		if status >= http.StatusInternalServerError {
			str := fmt.Sprintf("status code %d", status)
			return errors.New(str)
		}
		return nil
	}, func(e error) error {
		fmt.Println("hystrix : ", e)
		ctx.JSON(500, struct {
			Code int64
			Msg  string
		}{500, "server unavailable"})
		return e
	})
}
