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
	List []getListList `json:"list"`
}

type getListList struct {
	Img      string `json:"img"`
	Username string `json:"username"`
	ID       string `json:"id"`
}

//GetListSerializer .
func GetListSerializer(resp *user.GetListResp) GetListResp {
	wResp := GetListResp{}
	for _, item := range resp.List {
		wResp.List = append(wResp.List, getListList{
			Img:      item.Img,
			Username: item.Username,
			ID:       item.Id,
		})
	}

	return wResp
}
