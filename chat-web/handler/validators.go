package handler

import (
	"errors"

	"github.com/papandadj/nezha-chat-backend/proto/chat"

	"github.com/papandadj/nezha-chat-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)

var (
	//ErrInputParams 输入参数有误
	ErrInputParams = errors.New("用户输入的参数有误")
)

//PostValidator ,
type PostValidator struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
	Req     chat.PostReq
	TokenID string
}

//Bind .
func (s *PostValidator) Bind(c *gin.Context) (err error) {
	err = c.ShouldBind(s)
	if err != nil {
		return
	}

	userInfo, _ := middleware.AuthWithGin(c)

	s.TokenID = userInfo.ID
	if s.UserID == "" || s.TokenID == "" || s.Message == "" {
		err = ErrInputParams
	}

	s.Req.UserId = s.UserID
	s.Req.TokenId = s.TokenID
	s.Req.Message = s.Message

	return
}

//GetQueueValidator ,
type GetQueueValidator struct {
	TokenID string
	Req     chat.GetQueueReq
}

//Bind .
func (s *GetQueueValidator) Bind(c *gin.Context) (err error) {
	userInfo, _ := middleware.AuthWithGin(c)

	s.TokenID = userInfo.ID
	if s.TokenID == "" {
		err = ErrInputParams
	}
	s.Req.TokenId = s.TokenID
	return
}
