package handler

import (
	"github.com/papandadj/nezha-chat-backend/proto/auth"
	"github.com/papandadj/nezha-chat-backend/proto/user"
)

// SignUpResp .
type SignUpResp struct {
}

//SignUpSerializer .
func SignUpSerializer(resp *user.PostResp) SignUpResp {
	return SignUpResp{}
}

//LoginResp .
type LoginResp struct {
	Token string `json:"token"`
}

//LoginSerializer .
func LoginSerializer(resp *auth.GetTokenResp) LoginResp {
	return LoginResp{
		Token: resp.Token,
	}
}

//GetListResp .
type GetListResp struct {
	List []getListItem `json:"list"`
}

type getListItem struct {
	Img      string `json:"img"`
	Username string `json:"username"`
	ID       string `json:"id"`
}

//GetListSerializer .
func GetListSerializer(resp *user.GetListResp) GetListResp {
	wResp := GetListResp{}
	for _, item := range resp.List {
		wResp.List = append(wResp.List, getListItem{
			Img:      item.Img,
			Username: item.Username,
			ID:       item.Id,
		})
	}

	return wResp
}

//GetResp .
type GetResp struct {
	Img      string `json:"img"`
	Username string `json:"username"`
	ID       string `json:"id"`
}

//GetSerializer .
func GetSerializer(userI *user.UserItem) GetResp {
	return GetResp{
		Img:      userI.Img,
		Username: userI.Username,
		ID:       userI.Id,
	}
}
