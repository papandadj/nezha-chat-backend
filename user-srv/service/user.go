package service

import (
	"context"

	"github.com/papandadj/nezha-chat-backend/common"

	"github.com/papandadj/nezha-chat-backend/proto/user"
)

//Post 添加用户
func (s *Service) Post(ctx context.Context, req *user.PostReq, resp *user.PostResp) (err error) {
	if req.Username == "" || req.Password == "" {
		resp.Error = &user.Error{Code: 400, Msg: common.UsernameOrPasswordIsNull}
		return
	}

	var ok bool
	_, ok, err = s.Dao.UserGetByUsername(req.Username)
	if err != nil {
		logger.Errorln(err)
		return
	}

	if ok {
		logger.Infoln("user has registered")
		resp.Error = &user.Error{Code: 400, Msg: common.UserHasRegistered}
		return
	}

	password := Sum256(req.Password)

	user, err := s.Dao.UserPost(req.Username, password)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.With("name", user.Username).Infoln("create user")
	return
}

//Login 登录
func (s *Service) Login(ctx context.Context, req *user.LoginReq, resp *user.LoginResp) (err error) {
	if req.Username == "" || req.Password == "" {
		resp.Error = &user.Error{Code: 400, Msg: common.UsernameOrPasswordIsNull}
		return
	}

	password := Sum256(req.Password)

	//XXX:
	logger.Debugln(password)

	return
}

//TODO:
//GetList 获取用户列表
func (s *Service) GetList(ctx context.Context, req *user.GetListReq, resp *user.GetListResp) (err error) {
	return
}
