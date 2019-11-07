package service

import (
	"context"

	"github.com/papandadj/nezha-chat-backend/proto/auth"
)

//GetToken 获取token
func (s *Service) GetToken(ctx context.Context, req *auth.GetTokenReq, resp *auth.GetTokenResp) (err error) {
	if req.Id == "" && req.Username == "" {
		resp.Error = &auth.Error{Code: 400, Msg: constUsernameAndIDNull}
		return
	}

	claim := TokenClaim{
		ID:        req.Id,
		Username:  req.Username,
		Timestamp: getTimeStamp(),
	}

	token, err := CreateToken(claim, s.tokenSecrete)
	if err != nil {
		logger.Errorln(err)
		return
	}

	err = s.Dao.AuthSaveToken(req.Username, token)
	if err != nil {
		logger.Errorln(err)
		return
	}

	resp.Token = token
	return
}

//Check 检测用户token是否可以用
func (s *Service) Check(ctx context.Context, req *auth.CheckReq, resp *auth.CheckResp) (err error) {

	claim, err := ParseToken(req.Token, s.tokenSecrete)
	if err != nil {
		logger.Errorln(err)
		return
	}

	resp.Id = claim.ID
	resp.Username = claim.Username

	return
}
