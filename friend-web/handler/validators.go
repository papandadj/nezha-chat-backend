package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/papandadj/nezha-chat-backend/proto/user"
)

var (
	//ErrInputParams 输入参数有误
	ErrInputParams = errors.New("用户输入的参数有误")
)

//SignUpValidator .
type SignUpValidator struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Req      user.PostReq
}

//Bind .
func (s *SignUpValidator) Bind(c *gin.Context) (err error) {
	err = c.ShouldBind(s)
	if err != nil {
		return
	}

	s.Req.Username = s.Username
	s.Req.Password = s.Password

	if s.Req.Username == "" || s.Req.Password == "" {
		err = ErrInputParams
	}

	return
}

//LoginValidator .
type LoginValidator struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Req      user.CheckPasswordReq
}

//Bind .
func (l *LoginValidator) Bind(c *gin.Context) (err error) {
	err = c.ShouldBind(l)
	if err != nil {
		return
	}

	l.Req.Username = l.Username
	l.Req.Password = l.Password

	if l.Req.Username == "" || l.Req.Password == "" {
		err = ErrInputParams
	}

	return
}

//GetListValidator .
type GetListValidator struct {
	Name string   `json:"name"`
	IDs  []string `json:"ids"`
	Req  user.GetListReq
}

//Bind .
func (g *GetListValidator) Bind(c *gin.Context) (err error) {
	err = c.ShouldBind(g)
	if err != nil {
		return
	}

	g.Req.Name = g.Name
	g.Req.Ids = g.IDs

	if g.Req.Name == "" && len(g.Req.Ids) == 0 {
		err = ErrInputParams
	}

	return
}
