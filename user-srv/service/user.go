package service

import (
	"context"

	"github.com/papandadj/nezha-chat-backend/user-srv/dao"

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

//CheckPassword .
func (s *Service) CheckPassword(ctx context.Context, req *user.CheckPasswordReq, resp *user.CheckPasswordResp) (err error) {
	if req.Username == "" || req.Password == "" {
		resp.Error = &user.Error{Code: 400, Msg: common.UsernameOrPasswordIsNull}
		return
	}

	password := Sum256(req.Password)

	userM, ok, err := s.Dao.UserCheckPassword(req.Username, password)
	if err != nil {
		logger.Errorln(err)
		return
	}

	resp.Result = ok
	if ok {
		resp.User = &user.UserItem{
			Id:       string(userM.ID),
			Username: userM.Username,
			Img:      userM.Image,
		}
	}
	return
}

//GetList 获取用户列表
func (s *Service) GetList(ctx context.Context, req *user.GetListReq, resp *user.GetListResp) (err error) {
	if len(req.Ids) == 0 && req.Name == "" {
		resp.Error = &user.Error{Code: 400, Msg: "参数错误"}
		return
	}

	userMList, err := s.Dao.UserGetList(req.Name, req.Ids)
	if err != nil {
		logger.Errorln(err)
		resp.Error = &user.Error{
			Code: 500,
			Msg:  err.Error(),
		}
		return
	}

	resp.List = dtoUserMList2PbUserItem(userMList)
	return
}

func dtoUserMList2PbUserItem(userMList []*dao.ModelUser) (userItems []*user.UserItem) {

	for _, userM := range userMList {
		item := new(user.UserItem)
		item.Id = string(userM.ID)
		item.Username = userM.Username
		item.Img = userM.Image

		userItems = append(userItems, item)
	}
	return
}
