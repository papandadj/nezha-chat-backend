package handler

import (
	"errors"

	"github.com/papandadj/nezha-chat-backend/proto/user"

	"github.com/papandadj/nezha-chat-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/papandadj/nezha-chat-backend/proto/friend"
)

var (
	//ErrInputParams 输入参数有误
	ErrInputParams = errors.New("用户输入的参数有误")
)

//PostValidator ,
type PostValidator struct {
	UserID     string `json:"user_id"`
	ReqUserGet user.GetReq
	Req        friend.PostReq
}

//Bind .
func (s *PostValidator) Bind(c *gin.Context) (err error) {
	err = c.ShouldBind(s)
	if err != nil {
		return
	}

	userInfo, _ := middleware.AuthWithGin(c)

	s.ReqUserGet.Id = s.UserID

	s.Req.UserId = s.UserID
	s.Req.TokenId = userInfo.ID

	if s.Req.UserId == "" || s.Req.TokenId == "" {
		err = ErrInputParams
	}

	return
}

//DeleteByUserIDValidator ,
type DeleteByUserIDValidator struct {
	UserID     string `json:"user_id"`
	ReqUserGet user.GetReq
	Req        friend.DelByUserIDReq
}

//Bind .
func (s *DeleteByUserIDValidator) Bind(c *gin.Context) (err error) {
	err = c.ShouldBind(s)
	if err != nil {
		return
	}

	userInfo, _ := middleware.AuthWithGin(c)

	s.ReqUserGet.Id = s.UserID

	s.Req.UserId = s.UserID
	s.Req.TokenId = userInfo.ID

	if s.Req.UserId == "" || s.Req.TokenId == "" {
		err = ErrInputParams
	}

	return
}

//GetListValidator ,
type GetListValidator struct {
	Req friend.GetListReq
}

//Bind .
func (s *GetListValidator) Bind(c *gin.Context) (err error) {
	err = c.ShouldBind(s)
	if err != nil {
		return
	}

	userInfo, _ := middleware.AuthWithGin(c)

	s.Req.TokenId = userInfo.ID

	if s.Req.TokenId == "" {
		err = ErrInputParams
	}

	return
}
